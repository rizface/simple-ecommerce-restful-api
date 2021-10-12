package setup

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple-ecommerce-rest-api/app"
	"simple-ecommerce-rest-api/app/middleware"
	"simple-ecommerce-rest-api/controller"
	"simple-ecommerce-rest-api/repository"
	"simple-ecommerce-rest-api/service"
)

func Seller() controller.SellerController {
	db := app.Connection()
	sellerRepo := repository.NewSellerRepoImpl()
	sellerProductRepo := repository.NewSellerProductRepositoryImpl()
	sellerService := service.NewSellerServiceImpl(app.Validator, db, sellerRepo)
	sellerProductService := service.NewSellerProductServiceImpl(db,app.Validator,sellerProductRepo)
	sellerController := controller.NewSellerControllerImpl(sellerService,sellerProductService)
	return sellerController
}

func GeneralCustomer() controller.CustomerProductController{
	db := app.Connection()
	customerProductRepo := repository.NewCustomerProductRepoImpl()
	productImages := repository.NewProductImagesRepoImpl()
	customerProductService := service.NewCustomerProductServiceImpl(db,productImages,customerProductRepo)
	customerProductController := controller.NewCustomerProductControllerImpl(customerProductService)
	return customerProductController
}

func SellerAuth() *mux.Router {
	SellerAuth := app.Mux.NewRoute().Subrouter()
	SellerAuth.Use(middleware.PanicHandler)

	sellerController := Seller()

	SellerAuth.HandleFunc(app.SELLER_REGITER,sellerController.Register).Methods(http.MethodPost)
	SellerAuth.HandleFunc(app.SELLER_LOGIN, sellerController.Login).Methods(http.MethodPost)
	return SellerAuth
}

func AuthenticatedSeller() *mux.Router {
	AuthenticatedSeller := app.Mux.NewRoute().Subrouter()
	AuthenticatedSeller.Use(middleware.PanicHandler)
	AuthenticatedSeller.Use(middleware.AuthenticatedSeller)

	sellerController := Seller()
	AuthenticatedSeller.HandleFunc(app.SELLER_PRODUCT, sellerController.GetProducts).Methods(http.MethodGet)
	AuthenticatedSeller.HandleFunc(app.SELLER_PRODUCT,sellerController.PostProduct).Methods(http.MethodPost)
	AuthenticatedSeller.HandleFunc(app.SELLER_PROUDUCT_MANIPULATION,sellerController.DeleteProduct).Methods(http.MethodDelete)
	AuthenticatedSeller.HandleFunc(app.SELLER_PROUDUCT_MANIPULATION,sellerController.UpdateProduct).Methods(http.MethodPut)
	AuthenticatedSeller.HandleFunc(app.SELLER_PROUDUCT_MANIPULATION,sellerController.GetDetailProduct).Methods(http.MethodGet)
	return AuthenticatedSeller
}

func CustomerProduct() *mux.Router {
	generalCustomerRouter := app.Mux.NewRoute().Subrouter()
	generalCustomerRouter.Use(middleware.PanicHandler)

	generalCustomerController := GeneralCustomer()
	generalCustomerRouter.HandleFunc(app.PRODUCTS, generalCustomerController.Get).Methods(http.MethodGet)
	generalCustomerRouter.HandleFunc(app.PRODUCT_DETAIL, generalCustomerController.GetDetail).Methods(http.MethodGet)
	return generalCustomerRouter
}