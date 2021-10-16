package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"simple-ecommerce-rest-api/app"
	"simple-ecommerce-rest-api/app/setup"
	"simple-ecommerce-rest-api/model/web"
	"strings"
	"testing"
)

func TestCustomerCartPost(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		requestCart := web.CartRequest{
			IdProduct: 55,
			Jumlah:    9,
		}
		requestCartJson, _ := json.Marshal(requestCart)
		reader := strings.NewReader(string(requestCartJson))
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodPost, app.CART, reader)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()

		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("greater than product stok", func(t *testing.T) {
		requestCart := web.CartRequest{
			IdProduct: 54,
			Jumlah:    100,
		}
		requestCartJson, _ := json.Marshal(requestCart)
		reader := strings.NewReader(string(requestCartJson))
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodPost, app.CART, reader)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()

		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("not found", func(t *testing.T) {
		requestCart := web.CartRequest{
			IdProduct: 1001,
			Jumlah:    2,
		}
		requestCartJson, _ := json.Marshal(requestCart)
		reader := strings.NewReader(string(requestCartJson))
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodPost, app.CART, reader)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()

		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func TestCustomerCartUpdate(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		requestCart := web.CartRequest{
			Jumlah: 0,
		}
		requestCartJson, _ := json.Marshal(requestCart)
		reader := strings.NewReader(string(requestCartJson))
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/cart/10", reader)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()

		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("greater than product stok", func(t *testing.T) {
		requestCart := web.CartRequest{
			Jumlah: 100,
		}
		requestCartJson, _ := json.Marshal(requestCart)
		reader := strings.NewReader(string(requestCartJson))
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/cart/10", reader)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()

		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("not found", func(t *testing.T) {
		requestCart := web.CartRequest{
			Jumlah: 100,
		}
		requestCartJson, _ := json.Marshal(requestCart)
		reader := strings.NewReader(string(requestCartJson))
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/cart/91", reader)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()

		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func TestCustomerCartGet(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodGet, app.CART, nil)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()
		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("not confirmed customer account", func(t *testing.T) {
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodGet, app.CART, nil)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()
		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})

}

func TestCustomerCartDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/cart/9", nil)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()

		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("not found", func(t *testing.T) {
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/cart/1001", nil)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		router := setup.CartRouter()

		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}
