package helper

import (
	"encoding/json"
	"net/http"
	"simple-ecommerce-rest-api/model"
)

func JsonWriter(writer http.ResponseWriter, code int, status string, data interface{}) {
	writer.WriteHeader(code)
	json.NewEncoder(writer).Encode(model.StandardResponse{
		Code:   code,
		Status: status,
		Data:   data,
	})
}
