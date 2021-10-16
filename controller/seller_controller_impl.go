package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/service"
	"strconv"
)

type sellerControllerImpl struct {
	service       service.SellerService
	sellerProduct service.SellerProductService
}

func NewSellerControllerImpl(service service.SellerService, sellerProduct service.SellerProductService) SellerController {
	return sellerControllerImpl{service: service, sellerProduct: sellerProduct}
}

func (s sellerControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	request := web.RequestSeller{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	seller := s.service.Register(r.Context(), request)
	json.NewEncoder(w).Encode(model.StandardResponse{
		Code:   http.StatusOK,
		Status: "seller registration success, open your email to verification email",
		Data:   seller,
	})
}

func (s sellerControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	request := web.RequestSeller{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	token := s.service.Login(r.Context(), request)
	helper.JsonWriter(w, http.StatusOK, "Login Success", map[string]string{
		"token": token,
	})
}

func (s sellerControllerImpl) Confirm(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	token := params["token"]
	result := s.service.Confirm(r.Context(),token)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (s sellerControllerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("seller-data").(*helper.SellerCustom)
	products := s.sellerProduct.GetProducts(r.Context(), data.Id)
	helper.JsonWriter(w, http.StatusOK, "Success", map[string]interface{}{
		"products": products,
	})
}

func (s sellerControllerImpl) GetDetailProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idProduct, _ := strconv.Atoi(params["idProduct"])
	product := s.sellerProduct.GetDetailProduct(r.Context(), idProduct)
	helper.JsonWriter(w, http.StatusOK, "success", product)
}

func (s sellerControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request) {
	request := web.ProductRequest{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	exception.PanicIfInternalServerError(err)
	idSeller := r.Context().Value("seller-data").(*helper.SellerCustom).Id
	s.sellerProduct.PostProduct(r.Context(), idSeller, request)
}

func (s sellerControllerImpl) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idProduct, err := strconv.Atoi(params["idProduct"])
	exception.PanicIfInternalServerError(err)
	s.sellerProduct.DeleteProduct(r.Context(), idProduct)
	helper.JsonWriter(w, http.StatusOK, "product delete success", nil)
}

func (s sellerControllerImpl) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	request := web.ProductRequest{}
	params := mux.Vars(r)
	idProduct, err := strconv.Atoi(params["idProduct"])
	exception.PanicBadRequest(err)
	seller := r.Context().Value("seller-data").(*helper.SellerCustom)
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&request)
	exception.PanicBadRequest(err)
	result := s.sellerProduct.UpdateProduct(r.Context(), idProduct, seller.Id, request)
	helper.JsonWriter(w, http.StatusOK, result, nil)
}
