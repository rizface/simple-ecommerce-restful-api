package controller

import (
	"encoding/json"
	"net/http"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/service"
)

type sellerControllerImpl struct {
	service service.SellerService
}

func NewSellerControllerImpl(service service.SellerService) SellerController {
	return sellerControllerImpl{service: service}
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
