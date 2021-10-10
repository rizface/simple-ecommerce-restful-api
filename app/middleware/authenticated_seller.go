package middleware

import (
	"context"
	"errors"
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"strings"
)

func AuthenticatedSeller(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		token := strings.Split(request.Header.Get("Authorization"), " ")
		if token[0] != "Bearer" || len(token) != 2{
			exception.PanicBadRequest(errors.New("Token Tidak Valid"))
		} else {
			claims,err := helper.VerifyToken(token[1])
			exception.PanicBadRequest(err)
			sellerData := context.WithValue(request.Context(),"seller-data",claims)
			request = request.WithContext(sellerData)
			next.ServeHTTP(writer,request)
		}
	})
}
