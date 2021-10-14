package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"simple-ecommerce-rest-api/app/setup"
	"testing"
)

func TestGeneralCustomerAllProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		generalCustomer := setup.CustomerProduct()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products", nil)
		recorder := httptest.NewRecorder()
		generalCustomer.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}

func TestGeneralCustomerDetailProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		generalCustomer := setup.CustomerProduct()
		request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/44", nil)
		recorder := httptest.NewRecorder()
		generalCustomer.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
}
