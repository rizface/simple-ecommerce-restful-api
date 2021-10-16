package middleware

import (
	"context"
	"errors"
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
	"strings"
)

func AuthenticatedCustomer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		items := strings.Split(request.Header.Get("Authorization"), " ")
		if len(items) != 2 || items[0] != "Bearer" {
			exception.PanicBadRequest(errors.New("token invalid"))
		} else {
			claims, err := helper.VerifyTokenCustomer(items[1])
			exception.PanicBadRequest(err)
			helper.Confirmed(claims.(*helper.CustomerCustom).Confirmed)
			request = request.WithContext(context.WithValue(request.Context(), "customer-data", claims))
			next.ServeHTTP(writer, request)
		}
	})
}
