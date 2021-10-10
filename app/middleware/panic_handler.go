package middleware

import (
	"fmt"
	"net/http"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/helper"
)

func PanicHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			err := recover()
			fmt.Println(err)
			if err != nil {
				if badRequest,badRequestOK := err.(exception.BadRequest); badRequestOK {
					helper.JsonWriter(writer,http.StatusBadRequest,badRequest.Error.(string),nil)
				}
				if duplicate,duplicateOK := err.(exception.Duplicate); duplicateOK {
					helper.JsonWriter(writer,http.StatusUnprocessableEntity,duplicate.Error.(string),nil)
				}
				if notFound,notFoundOK := err.(exception.NotFound); notFoundOK {
					helper.JsonWriter(writer,http.StatusNotFound,notFound.Error.(string),nil)
				}
				if internalerror,internalErrorOK := err.(exception.InternalServerError); internalErrorOK {
					helper.JsonWriter(writer,http.StatusInternalServerError,internalerror.Error.(string),nil)
				}
				if unauth,unauthOK := err.(exception.Unauthorized); unauthOK {
					helper.JsonWriter(writer,http.StatusUnauthorized,unauth.Error.(string),nil)
				}
			}
		}()
		next.ServeHTTP(writer,request)
	})
}