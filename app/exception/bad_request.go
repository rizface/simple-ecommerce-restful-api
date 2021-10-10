package exception

import (
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

type BadRequest struct {
	Error interface{}
}

func PanicBadRequest(err interface{}) {
	if err != nil {
		if validatorErr,validatorErrOK := err.(validator.ValidationErrors); validatorErrOK {
			panic(BadRequest{
				Error: validatorErr[0].Error(),
			})
		} else if jwtErr,jwtErrOK := err.(*jwt.ValidationError); jwtErrOK{
			panic(BadRequest{
				Error: jwtErr.Error(),
			})
		} else {
			panic(BadRequest{
				Error: err.(error).Error(),
			})
		}
	}
}