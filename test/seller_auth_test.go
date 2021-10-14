package test

import (
	"context"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"simple-ecommerce-rest-api/app"
	"simple-ecommerce-rest-api/app/setup"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/web"
	"strings"
	"testing"
)

var ctx = context.Background()
var db = app.Connection()

var dataValid = web.RequestSeller{
	NamaToko:   "toko_sejahtera20",
	Email:      "sejahtera4@gmail.com",
	Password:   helper.Hash("sejahtera123"),
	AlamatToko: "jakarta",
	Deskripsi:  "toko kami keren",
}

var dataInvalid = web.RequestSeller{
	Email:      "sejahtera@gmail.com",
	Password:   helper.Hash("sejahtera123"),
	AlamatToko: "jakarta",
	Deskripsi:  "toko kami keren",
}

// Test Seller Register
func TestSellerIntegrtionRegister(t *testing.T) {
	//t.SkipNow()
	t.Run("success", func(t *testing.T) {
		dataJson, err := json.Marshal(dataValid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/seller/register", reader)
		recorder := httptest.NewRecorder()

		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder, request)

		response := recorder
		assert.Equal(t, 200, response.Code)
	})

	t.Run("bad request", func(t *testing.T) {
		dataJson, err := json.Marshal(dataInvalid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/seller/register", reader)
		recorder := httptest.NewRecorder()

		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder, request)

		response := recorder
		assert.Equal(t, 400, response.Code)
	})

	t.Run("duplicate email", func(t *testing.T) {
		dataJson, err := json.Marshal(dataValid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/seller/register", reader)
		recorder := httptest.NewRecorder()

		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder, request)

		response := recorder
		assert.Equal(t, 422, response.Code)
	})

	t.Run("duplicate seller", func(t *testing.T) {
		dataValid.Email = "otherseller@gmail.com"
		dataJson, err := json.Marshal(dataValid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/seller/register", reader)
		recorder := httptest.NewRecorder()

		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder, request)

		response := recorder
		assert.Equal(t, 422, response.Code)
	})
}

// Test Seller Login
func TestSellerIntegrationLogin(t *testing.T) {

	t.Run("login success", func(t *testing.T) {
		payload := web.RequestSeller{
			Email:    "sejahtera4@gmail.com",
			Password: "sejahtera123",
		}
		jsonPayload, err := json.Marshal(payload)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(jsonPayload))
		request := httptest.NewRequest(http.MethodPost, app.SELLER_LOGIN, reader)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder, request)
		response := recorder
		assert.Equal(t, 200, response.Code)
	})

	t.Run("login failed", func(t *testing.T) {
		payload := web.RequestSeller{
			Email:    "sejahtera4@gmail.com",
			Password: "sejahtera1235",
		}
		jsonPayload, err := json.Marshal(payload)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(jsonPayload))
		request := httptest.NewRequest(http.MethodPost, app.SELLER_LOGIN, reader)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder, request)
		response := recorder
		assert.Equal(t, 401, response.Code)
	})

	t.Run("account not exist", func(t *testing.T) {
		payload := web.RequestSeller{
			Email:    "sejahasdtera2@gmail.com",
			Password: "sejahtera1235",
		}
		jsonPayload, err := json.Marshal(payload)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(jsonPayload))
		request := httptest.NewRequest(http.MethodPost, app.SELLER_LOGIN, reader)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder, request)
		response := recorder
		assert.Equal(t, 404, response.Code)
	})
}

var test int

// Test Verify Seller Token
func TestSellerVerifyToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		token, _ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodGet, app.SELLER_PRODUCT, nil)
		request.Header.Add("Authorization", "Bearer "+string(token))
		recorder := httptest.NewRecorder()
		sellerAuth := setup.AuthenticatedSeller()
		sellerAuth.ServeHTTP(recorder, request)
		response := recorder
		assert.Equal(t, http.StatusOK, response.Code)

		//resBody,_ := io.ReadAll(response.Body)
		//fmt.Println(string(resBody))

	})

	t.Run("invalid token", func(t *testing.T) {
		//t.SkipNow()
		token := "Bearer" + " token"
		request := httptest.NewRequest(http.MethodGet, app.SELLER_PRODUCT, nil)
		request.Header.Add("Authorization", token)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.AuthenticatedSeller()
		sellerAuth.ServeHTTP(recorder, request)
		response := recorder
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})

	t.Run("empty token", func(t *testing.T) {
		//t.SkipNow()
		token := ""
		request := httptest.NewRequest(http.MethodGet, app.SELLER_PRODUCT, nil)
		request.Header.Add("Authorization", token)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.AuthenticatedSeller()
		sellerAuth.ServeHTTP(recorder, request)
		response := recorder
		assert.Equal(t, http.StatusBadRequest, response.Code)
	})
}
