package main

import (
	"context"
	"net/http"
	"simple-ecommerce-rest-api/app"
	"simple-ecommerce-rest-api/app/setup"
	"simple-ecommerce-rest-api/helper"
)

func main() {
	setup.SellerAuth()
	setup.AuthenticatedSeller()
	setup.CustomerProduct()
	setup.CustomerAuthRouter()
	setup.CartRouter()

	_, err := helper.Rdb.Ping(context.Background()).Result()

	server := http.Server{
		Addr:    ":8080",
		Handler: app.Mux,
	}
	err = server.ListenAndServe()
	helper.PanicIfError(err)
}
