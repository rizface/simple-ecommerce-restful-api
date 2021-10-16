package helper

import (
	"errors"
	"simple-ecommerce-rest-api/app/exception"
)

func Confirmed(confirmed int) {
	if confirmed < 1 {
		exception.PanicUnauthorized(errors.New("account not confirmed yet"))
	}
}