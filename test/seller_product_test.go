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

func TestSellerProductGet(t *testing.T) {
	t.Run("valid token", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodGet,app.SELLER_PRODUCT,nil)
		request.Header.Add("Authorization","Bearer " + string(token))
		recorder := httptest.NewRecorder()
		sellerAuth := setup.AuthenticatedSeller()
		sellerAuth.ServeHTTP(recorder,request)
		response := recorder
		assert.Equal(t, http.StatusOK,response.Code)
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

func TestSellerProductGetDetailProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/seller/products/24",nil)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()
		authenticatedSeller := setup.AuthenticatedSeller()
		authenticatedSeller.ServeHTTP(recorder,request)
		resBody,_ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
	t.Run("not found", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodGet,"http://localhost:8080/seller/products/1",nil)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()
		authenticatedSeller := setup.AuthenticatedSeller()
		authenticatedSeller.ServeHTTP(recorder,request)
		resBody,_ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusNotFound, recorder.Code)
	})
}

func TestSellerProductPost(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		img,_ := ioutil.ReadFile("img.txt")
		product := web.ProductRequest{
			NamaBarang:  "baju tidur",
			HargaBarang: 10000,
			Stokbarang:  10,
			Deskripsi:   "baju ini bagus",
			Gambar:      []string{
				string(img),
			},
		}
		productJson,_ := json.Marshal(product)
		reader := strings.NewReader(string(productJson))
		request := httptest.NewRequest(http.MethodPost,app.SELLER_PRODUCT, reader)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()

		authenticatedSeller := setup.AuthenticatedSeller()
		authenticatedSeller.ServeHTTP(recorder,request)

		response := recorder

		assert.Equal(t, http.StatusOK,response.Code)
	})

	t.Run("invalid empty nama_barang", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		img,_ := ioutil.ReadFile("img.txt")
		product := web.ProductRequest{
			NamaBarang:  "",
			HargaBarang: 10000,
			Stokbarang:  10,
			Deskripsi:   "baju ini bagus",
			Gambar:      []string{
				string(img),
			},
		}
		productJson,_ := json.Marshal(product)
		reader := strings.NewReader(string(productJson))
		request := httptest.NewRequest(http.MethodPost,app.SELLER_PRODUCT, reader)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()

		authenticatedSeller := setup.AuthenticatedSeller()
		authenticatedSeller.ServeHTTP(recorder,request)

		response := recorder

		assert.Equal(t, http.StatusBadRequest,response.Code)
	})

	t.Run("invalid empty gambar", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		//img,_ := ioutil.ReadFile("img.txt")
		product := web.ProductRequest{
			NamaBarang:  "sempak",
			HargaBarang: 10000,
			Stokbarang:  10,
			Deskripsi:   "baju ini bagus",
			Gambar:      []string{},
		}
		productJson,_ := json.Marshal(product)
		reader := strings.NewReader(string(productJson))
		request := httptest.NewRequest(http.MethodPost,app.SELLER_PRODUCT, reader)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()
		authenticatedSeller := setup.AuthenticatedSeller()
		authenticatedSeller.ServeHTTP(recorder,request)

		response := recorder

		assert.Equal(t, http.StatusBadRequest,response.Code)
	})
}

func TestSellerProductUpdate(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		updatedData := web.ProductRequest{
			NamaBarang:  "baju tidur lagi lagi ",
			HargaBarang: 10000,
			Stokbarang:  10,
			Deskripsi:   "baju ini bagus",
		}
		token,_ := ioutil.ReadFile("token.txt")
		jsonUpdate,_ := json.Marshal(updatedData)
		reader := strings.NewReader(string(jsonUpdate))
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/seller/products/23",reader)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()
		authSeller := setup.AuthenticatedSeller()
		authSeller.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusOK,recorder.Code)
	})
	t.Run("invalid", func(t *testing.T) {
		updatedData := web.ProductRequest{
			NamaBarang:  "",
			HargaBarang: 10000,
			Stokbarang:  10,
			Deskripsi:   "baju ini bagus",
		}
		token,_ := ioutil.ReadFile("token.txt")
		jsonUpdate,_ := json.Marshal(updatedData)
		reader := strings.NewReader(string(jsonUpdate))
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/seller/products/23",reader)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()
		authSeller := setup.AuthenticatedSeller()
		authSeller.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusBadRequest,recorder.Code)
	})
	t.Run("not found", func(t *testing.T) {
		updatedData := web.ProductRequest{
			NamaBarang:  "baju tidur lagi lagi ",
			HargaBarang: 10000,
			Stokbarang:  10,
			Deskripsi:   "baju ini bagus",
		}
		token,_ := ioutil.ReadFile("token.txt")
		jsonUpdate,_ := json.Marshal(updatedData)
		reader := strings.NewReader(string(jsonUpdate))
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/seller/products/212",reader)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()
		authSeller := setup.AuthenticatedSeller()
		authSeller.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusNotFound,recorder.Code)
	})
	t.Run("invalid token", func(t *testing.T) {
		updatedData := web.ProductRequest{
			NamaBarang:  "baju tidur lagi lagi ",
			HargaBarang: 10000,
			Stokbarang:  10,
			Deskripsi:   "baju ini bagus",
		}
		jsonUpdate,_ := json.Marshal(updatedData)
		reader := strings.NewReader(string(jsonUpdate))
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/seller/products/212",reader)
		request.Header.Add("Authorization", "Bearer " + " token")
		recorder := httptest.NewRecorder()
		authSeller := setup.AuthenticatedSeller()
		authSeller.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusBadRequest,recorder.Code)
	})
	t.Run("empty token", func(t *testing.T) {
		updatedData := web.ProductRequest{
			NamaBarang:  "baju tidur lagi lagi ",
			HargaBarang: 10000,
			Stokbarang:  10,
			Deskripsi:   "baju ini bagus",
		}
		jsonUpdate,_ := json.Marshal(updatedData)
		reader := strings.NewReader(string(jsonUpdate))
		request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/seller/products/212",reader)
		request.Header.Add("Authorization", "")
		recorder := httptest.NewRecorder()
		authSeller := setup.AuthenticatedSeller()
		authSeller.ServeHTTP(recorder,request)
		assert.Equal(t, http.StatusBadRequest,recorder.Code)
	})
}

func TestSellerProductDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodDelete,"http://localhost:8080/seller/products/23",nil)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()
		authSeller := setup.AuthenticatedSeller()
		authSeller.ServeHTTP(recorder,request)
		assert.Equal(t, 200,recorder.Code)
	})
	t.Run("not found", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodDelete,"http://localhost:8080/seller/products/100",nil)
		request.Header.Add("Authorization", "Bearer " + string(token))
		recorder := httptest.NewRecorder()
		authSeller := setup.AuthenticatedSeller()
		authSeller.ServeHTTP(recorder,request)
		assert.Equal(t, 404,recorder.Code)
	})
}

