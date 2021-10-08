package middleware

import (
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				if badRequest,badRequestOK := err.(exception.BadRequest); badRequestOK {
					helper.JsonWriter(writer,http.StatusBadRequest,badRequest.Error.(string),nil)
				}
				if duplicate,duplicateOK := err.(exception.Duplicate); duplicateOK {
					helper.JsonWriter(writer,http.StatusUnprocessableEntity,duplicate.Error.(string),nil)
				}
			}
		}()
		next.ServeHTTP(writer,request)
	})
}