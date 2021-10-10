package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
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

func TestSellerProductPost(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		img,_ := ioutil.ReadFile("img.txt")
		product := web.NewProduct{
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
		product := web.NewProduct{
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
		product := web.NewProduct{
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

func TestSellerProductDelete(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		token,_ := ioutil.ReadFile("token.txt")
		request := httptest.NewRequest(http.MethodDelete,"http://localhost:8080/seller/products/22",nil)
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