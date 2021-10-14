package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/service"
	"strconv"
)

type cartControllerImpl struct {
	service service.CartService
}

func NewCartController(service service.CartService) CartController {
	return cartControllerImpl{
		service: service,
	}
}

func (c cartControllerImpl) PostItem(w http.ResponseWriter, r *http.Request) {
	request := web.CartRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	exception.PanicBadRequest(err)
	customer := r.Context().Value("customer-data").(*helper.CustomerCustom)
	request.IdCustomer = customer.Id
	result := c.service.PostItem(r.Context(),request)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (c cartControllerImpl) GetItems(w http.ResponseWriter, r *http.Request) {
	customer := r.Context().Value("customer-data").(*helper.CustomerCustom)
	items := c.service.GetItems(r.Context(),customer.Id)
	helper.JsonWriter(w,http.StatusOK,"success",items)
}

func (c cartControllerImpl) UpdateItem(w http.ResponseWriter, r *http.Request) {
	request := web.CartRequest{}
	params := mux.Vars(r)
	err := json.NewDecoder(r.Body).Decode(&request)
	exception.PanicBadRequest(err)
	customer := r.Context().Value("customer-data").(*helper.CustomerCustom)
	idCart,_ := strconv.Atoi(params["idCart"])
	request.IdCustomer = customer.Id
	result := c.service.UpdateItem(r.Context(),request,idCart)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

func (c cartControllerImpl) DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idCart,_ := strconv.Atoi(params["idCart"])
	result := c.service.DeleteItem(r.Context(),idCart)
	helper.JsonWriter(w,http.StatusOK,result,nil)
}

