package main

import (
	"context"
	"net/http"
	"simple-ecommerce-rest-api/app"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/app/setup"
	"simple-ecommerce-rest-api/helper"
)

func main() {
	setup.SellerAuth()
	setup.AuthenticatedSeller()
	setup.CustomerProduct()
	setup.CustomerAuthRouter()
	setup.CartRouter()

	app.Mux.MethodNotAllowedHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		helper.JsonWriter(writer,http.StatusMethodNotAllowed,"method not allowed",nil)
	})

	app.Mux.NotFoundHandler = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		helper.JsonWriter(writer,http.StatusNotFound,"endpoint / page not found",nil)
	})

	_, err := helper.Rdb.Ping(context.Background()).Result()
	exception.PanicIfInternalServerError(err)
	server := http.Server{
		Addr:    ":8080",
		Handler: app.Mux,
	}
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
