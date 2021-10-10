package controller

import (
	"encoding/json"
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/service"
)

type sellerControllerImpl struct {
	service service.SellerService
	sellerProduct service.SellerProductService
}

func NewSellerControllerImpl(service service.SellerService, sellerProduct service.SellerProductService) SellerController {
	return sellerControllerImpl{service: service, sellerProduct:sellerProduct}
}

func (s sellerControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	request := web.RequestSeller{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	seller := s.service.Register(r.Context(),request)
	json.NewEncoder(w).Encode(model.StandardResponse{
		Code:   http.StatusOK,
		Status: "Registrasi Seller Berhasil",
		Data:   seller,
	})
}

func (s sellerControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	request := web.RequestSeller{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&request)
	token := s.service.Login(r.Context(),request)
	helper.JsonWriter(w,http.StatusOK,"Login Success", map[string]string{
		"token": token,
	})
}

func (s sellerControllerImpl) GetProducts(w http.ResponseWriter, r *http.Request) {
	data := r.Context().Value("seller-data").(*helper.SellerCustom)
	products := s.sellerProduct.GetProducts(r.Context(),data.Id)
	helper.JsonWriter(w,http.StatusOK,"Success", map[string]interface{} {
		"products": products,
	})
}

func (s sellerControllerImpl) PostProduct(w http.ResponseWriter, r *http.Request) {
	request := web.NewProduct{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	exception.PanicIfInternalServerError(err)
	idSeller := r.Context().Value("seller-data").(*helper.SellerCustom).Id
	s.sellerProduct.PostProduct(r.Context(),idSeller,request)
}
