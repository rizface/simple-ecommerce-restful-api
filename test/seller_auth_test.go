package test

import (
	"context"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"simple-ecommerce-rest-api/app"
	"simple-ecommerce-rest-api/app/setup"
	"simple-ecommerce-rest-api/controller"
	"simple-ecommerce-rest-api/helper"
	"simple-ecommerce-rest-api/model/web"
	"simple-ecommerce-rest-api/repository"
	"simple-ecommerce-rest-api/service"
	"strings"
	"testing"
)

var ctx = context.Background()
var db = app.Connection()

var dataValid = web.RequestSeller{
NamaToko:   "toko_sejahtera",
Email:      "sejahtera@gmail.com",
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

// Test SellerRepository For Register
func TestSellerRepoRegister(t *testing.T) {
	t.SkipNow()
	tx,err := db.Begin()
	helper.PanicIfError(err)
	repoImpl := repository.NewSellerRepoImpl()
	id := repoImpl.Register(ctx,tx,dataValid)
	tx.Commit()
	result := id > 0
	assert.Equal(t, true,result)
}

// Test SellerService For Register
func TestSellerServiceRegister(t *testing.T) {
	t.SkipNow()
	t.Run("success", func(t *testing.T) {
		sellerRepo := repository.NewSellerRepoImpl()
		sellerProductRepo := repository.NewSellerProductRepositoryImpl()
		serviceImpl := service.NewSellerServiceImpl(app.Validator,db,sellerProductRepo,sellerRepo)
		result := serviceImpl.Register(ctx,dataValid)
		assert.Equal(t, true,result.Id > 0)
	})
}

// Test SellerController For Register
func TestSellerControllerRegister(t *testing.T)  {
	t.SkipNow()
	t.Run("success", func(t *testing.T) {
		dataJson,err := json.Marshal(dataValid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, app.SELLER_REGITER,reader)
		recorder := httptest.NewRecorder()

		sellerRepo := repository.NewSellerRepoImpl()
		sellerProductRepo := repository.NewSellerProductRepositoryImpl()
		sellerService := service.NewSellerServiceImpl(app.Validator,db,sellerProductRepo,sellerRepo)
		sellerController := controller.NewSellerControllerImpl(sellerService)
		setup.SellerAuth()
		sellerController.Register(recorder,request)
		response := recorder
 		helper.PanicIfError(err)
		assert.Equal(t, 200,response.Code)
	} )
}

// Seller HTTP Test
func TestSellerIntegrtionRegister(t *testing.T)  {
	//t.SkipNow()
	t.Run("success", func(t *testing.T) {
		dataJson,err := json.Marshal(dataValid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/seller/register",reader)
		recorder := httptest.NewRecorder()

		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder,request)

		response := recorder
		assert.Equal(t, 200,response.Code)
	} )

	t.Run("bad request", func(t *testing.T) {
		dataJson,err := json.Marshal(dataInvalid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/seller/register",reader)
		recorder := httptest.NewRecorder()

		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder,request)

		response := recorder
		assert.Equal(t, 400,response.Code)
	})

	t.Run("duplicate email" , func(t *testing.T) {
		dataJson,err := json.Marshal(dataValid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/seller/register",reader)
		recorder := httptest.NewRecorder()

		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder,request)

		response := recorder
		assert.Equal(t, 422,response.Code)
	} )

	t.Run("duplicate seller" , func(t *testing.T) {
		dataValid.Email = "otherseller@gmail.com"
		dataJson,err := json.Marshal(dataValid)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(dataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/seller/register",reader)
		recorder := httptest.NewRecorder()

		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder,request)

		response := recorder
		assert.Equal(t, 422,response.Code)
	})
}

func TestSellerIntegrationLogin(t *testing.T) {

	t.Run("login success", func(t *testing.T) {
		payload := web.RequestSeller{
			Email: "sejahtera@gmail.com",
			Password: "sejahtera123",
		}
		jsonPayload,err := json.Marshal(payload)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(jsonPayload))
		request := httptest.NewRequest(http.MethodPost, app.SELLER_LOGIN, reader)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder,request)
		response := recorder
		assert.Equal(t, 200,response.Code)
	})

	t.Run("login failed", func(t *testing.T) {
		payload := web.RequestSeller{
			Email: "sejahtera@gmail.com",
			Password: "sejahtera1235",
		}
		jsonPayload,err := json.Marshal(payload)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(jsonPayload))
		request := httptest.NewRequest(http.MethodPost, app.SELLER_LOGIN, reader)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder,request)
		response := recorder
		assert.Equal(t, 401,response.Code)
	})

	t.Run("account not exist", func(t *testing.T) {
		payload := web.RequestSeller{
			Email: "sejahasdtera2@gmail.com",
			Password: "sejahtera1235",
		}
		jsonPayload,err := json.Marshal(payload)
		helper.PanicIfError(err)
		reader := strings.NewReader(string(jsonPayload))
		request := httptest.NewRequest(http.MethodPost, app.SELLER_LOGIN, reader)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.SellerAuth()
		sellerAuth.ServeHTTP(recorder,request)
		response := recorder
		assert.Equal(t, 404,response.Code)
	})
}

func TestSellerVerifyToken(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		token := "Bearer" + " eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MjcwLCJuYW1hX3Rva28iOiJ0b2tvX3NlamFodGVyYTIiLCJlbWFpbCI6InNlamFodGVyYTJAZ21haWwuY29tIiwiZGVza3JpcHNpIjoidG9rbyBrYW1pIGtlcmVuIiwic2VsbGVyIjp0cnVlLCJjcmVhdGVkX2F0IjoiNiBPY3RvYmVyIDIwMjEiLCJpc3MiOiJNdWhhbW1hZCBBbCBGYXJpenppIiwiZXhwIjoxNjMzNzc0NjE4fQ.mt2ZAd_9WV-7t8-YSzI-r5KVdJzvB9yLs4sAaPSWtGk"
		request := httptest.NewRequest(http.MethodGet,app.SELLER_PRODUCT,nil)
		request.Header.Add("Authorization", token)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.AuthenticatedSeller()
		sellerAuth.ServeHTTP(recorder,request)
		response := recorder
		assert.Equal(t, http.StatusOK,response.Code)

		//resBody,_ := io.ReadAll(response.Body)
		//fmt.Println(string(resBody))

	})

	t.Run("invalid token", func(t *testing.T) {
		//t.SkipNow()
		token := "Bearer" + " token"
		request := httptest.NewRequest(http.MethodGet,app.SELLER_PRODUCT,nil)
		request.Header.Add("Authorization", token)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.AuthenticatedSeller()
		sellerAuth.ServeHTTP(recorder,request)
		response := recorder
		assert.Equal(t, http.StatusBadRequest,response.Code)
	})

	t.Run("empty token", func(t *testing.T) {
		//t.SkipNow()
		token := ""
		request := httptest.NewRequest(http.MethodGet,app.SELLER_PRODUCT,nil)
		request.Header.Add("Authorization", token)
		recorder := httptest.NewRecorder()
		sellerAuth := setup.AuthenticatedSeller()
		sellerAuth.ServeHTTP(recorder,request)
		response := recorder
		assert.Equal(t, http.StatusBadRequest,response.Code)
	})
}