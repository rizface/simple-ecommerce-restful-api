package controller

import (
	"encoding/json"
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/service"
)

type orderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return orderController{
		service: service,
	}
}

func (o orderController) GetOrders(w http.ResponseWriter, r *http.Request) {
	panic("implement me")
}

func (o orderController) PostOrders(w http.ResponseWriter, r *http.Request) {
	request := web.OrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	exception.PanicBadRequest(err)
	claims,_ := r.Context().Value("customer-data").(*helper.CustomerCustom)
	o.service.PostOrders(r.Context(),claims.Id,request)
}

