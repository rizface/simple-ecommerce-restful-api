package exception

import (
	"github.com/go-playground/validator/v10"
)

type BadRequest struct {
	Error interface{}
}

func PanicBadRequest(err interface{}) {
	if err != nil {
		panic(BadRequest{
			Error: err.(validator.ValidationErrors)[0].Error(),
		})
	}
}