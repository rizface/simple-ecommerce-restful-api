package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/service"
	"strconv"
)

type costomerProductControllerImpl struct {
	service service.CustomerProductService
}

func NewCustomerProductControllerImpl(service service.CustomerProductService) CustomerProductController {
	return costomerProductControllerImpl{service: service}
}

func (c costomerProductControllerImpl) Get(w http.ResponseWriter, r *http.Request) {
	products := c.service.Get(r.Context())
	helper.JsonWriter(w,http.StatusOK,"success", products)
}

func (c costomerProductControllerImpl) GetDetail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idProduct,_ := strconv.Atoi(params["idProduct"])
	product := c.service.GetDetail(r.Context(),idProduct)
	helper.JsonWriter(w, http.StatusOK, "success", product)
}
