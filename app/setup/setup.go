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

func SellerAuth() *mux.Router {
	db := app.Connection()
	SellerAuth := app.Mux.NewRoute().Subrouter()
	SellerAuth.Use(middleware.PanicHandler)
	sellerRepo := repository.NewSellerRepoImpl()
	sellerService := service.NewSellerServiceImpl(app.Validator, db,sellerRepo)
	sellerController := controller.NewSellerControllerImpl(sellerService)

	SellerAuth.HandleFunc(app.SELLER_REGITER,sellerController.Register).Methods(http.MethodPost)
	SellerAuth.HandleFunc(app.SELLER_LOGIN, sellerController.Login).Methods(http.MethodPost)

	return SellerAuth
}