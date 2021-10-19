package test

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"net/http/httptest"
	"simple-ecommerce-rest-api/app"
	"simple-ecommerce-rest-api/controller"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/repository"
	"simple-ecommerce-rest-api/service"
	"testing"
	"time"
)

func TestPostOrders(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		orderRequest := web.OrderRequest{
			Items:  []web.OrderProduct{
				web.OrderProduct{
					IdProduct: 59,
					Jumlah:    1,
				},
			},
			Alamat: "kecamatan lima puluh kota, sumatera barat" ,
		}
		orderJson,_ := json.Marshal(orderRequest)
		reader := bytes.NewReader(orderJson)
		claims := &helper.CustomerCustom{
			Id:               41,
			NamaCustomer:     "fariz",
			Email:            "malfarizzi13@gmail.com",
			NoHp:             "123",
			Confirmed:        1,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer: "Fariz",
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			},
		}
		request := httptest.NewRequest(http.MethodPost,"/cart",reader)
		request = request.WithContext(context.WithValue(request.Context(),"customer-data",claims))
		recorder := httptest.NewRecorder()
		c := controller.NewOrderController(service.NewOrderService(app.Connection(),app.Validator,repository.NewSellerProductRepositoryImpl(),repository.NewOrderRepository(),repository.NewInvoiceRepo()))
		c.PostOrders(recorder,request)
	})
}
