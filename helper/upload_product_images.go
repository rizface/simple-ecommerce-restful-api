package helper

import (
	"encoding/json"
	"errors"
	"net/http"
	"simple-ecommerce-rest-api/app"
	"simple-ecommerce-rest-api/app/exception"
	"simple-ecommerce-rest-api/model/outbond"
	"strings"
)

func UploadProductImages(images []string) []string {

	if len(images) < 1 {
		exception.PanicBadRequest(errors.New("masukan gambar product"))
	}

	client := &http.Client{}
	jsonPayload, err := json.Marshal(outbond.FileUpload{
		Files: images,
	})
	exception.PanicIfInternalServerError(err)

	reader := strings.NewReader(string(jsonPayload))
	request, err := http.NewRequest(http.MethodPost, app.UPLOAD_IMAGES, reader)
	exception.PanicIfInternalServerError(err)

	response, err := client.Do(request)
	exception.PanicIfInternalServerError(err)

	uploadResponse := outbond.FileUploadResponse{}
	err = json.NewDecoder(response.Body).Decode(&uploadResponse)
	exception.PanicIfInternalServerError(err)

	return uploadResponse.Data
}
