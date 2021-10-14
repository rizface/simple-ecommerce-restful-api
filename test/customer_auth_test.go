package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"simple-ecommerce-rest-api/app/setup"
	"simple-ecommerce-rest-api/model/web"
	"strings"
	"testing"
)

func TestRegister(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		requestData := web.RequestCustomer{
			NamaCustomer: "handoko",
			Email:        "handoko@gmail.com",
			NoHp:         "081212342134",
			Password:     "rahasia",
		}
		requestDataJson, _ := json.Marshal(requestData)
		reader := strings.NewReader(string(requestDataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/customer/register", reader)
		recorder := httptest.NewRecorder()
		router := setup.CustomerAuthRouter()
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusOK, recorder.Code)
	})
	t.Run("bad request", func(t *testing.T) {
		requestData := web.RequestCustomer{
			NamaCustomer: "",
			Email:        "malfarizzi13@gmail.com",
			NoHp:         "081212342134",
			Password:     "rahasia",
		}
		requestDataJson, _ := json.Marshal(requestData)
		reader := strings.NewReader(string(requestDataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/customer/register", reader)
		recorder := httptest.NewRecorder()
		router := setup.CustomerAuthRouter()
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusBadRequest, recorder.Code)
	})
	t.Run("duplicate", func(t *testing.T) {
		requestData := web.RequestCustomer{
			NamaCustomer: "muhammad al farizzi",
			Email:        "malfarizzi13@gmail.com",
			NoHp:         "081212342134",
			Password:     "rahasia",
		}
		requestDataJson, _ := json.Marshal(requestData)
		reader := strings.NewReader(string(requestDataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/customer/register", reader)
		recorder := httptest.NewRecorder()
		router := setup.CustomerAuthRouter()
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusUnprocessableEntity, recorder.Code)
	})
}

func TestLogin(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		requestData := web.RequestCustomer{
			Email:    "handoko@gmail.com",
			Password: "rahasia",
		}
		requestDataJson, _ := json.Marshal(requestData)
		reader := strings.NewReader(string(requestDataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/customer/login", reader)
		recorder := httptest.NewRecorder()
		router := setup.CustomerAuthRouter()
		router.ServeHTTP(recorder, request)
		resBody, _ := io.ReadAll(recorder.Body)
		fmt.Println(string(resBody))
		assert.Equal(t, http.StatusOK, recorder.Code)
	})

	t.Run("failed", func(t *testing.T) {
		requestData := web.RequestCustomer{
			Email:    "handoko@gmail.com",
			Password: "rahasia1",
		}
		requestDataJson, _ := json.Marshal(requestData)
		reader := strings.NewReader(string(requestDataJson))
		request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/customer/login", reader)
		recorder := httptest.NewRecorder()
		router := setup.CustomerAuthRouter()
		router.ServeHTTP(recorder, request)
		assert.Equal(t, http.StatusUnauthorized, recorder.Code)
	})
}
