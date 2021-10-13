package controller

import (
	"encoding/json"
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/service"
)

type customerAuthControllerImpl struct{
	service service.CustomerAuthService
}

func NewCustomerAuthController(service service.CustomerAuthService) CustomerAuthController {
	return customerAuthControllerImpl{
		service: service,
	}
}
func (c customerAuthControllerImpl) Register(w http.ResponseWriter, r *http.Request) {
	request := web.RequestCustomer{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	exception.PanicIfInternalServerError(err)
	result := c.service.RegisterCustomer(r.Context(),request)
	if result == true {
		helper.JsonWriter(w,http.StatusOK,"registrasi customer success", nil)
	} else {
		helper.JsonWriter(w,http.StatusOK,"registrasi customer failed", nil)
	}
}

func (c customerAuthControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	request := web.RequestCustomer{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	exception.PanicIfInternalServerError(err)
	result := c.service.LoginCustomer(r.Context(),request)
	helper.JsonWriter(w,http.StatusOK,"login success", map[string]interface{} {
		"token":result,
	})
}
