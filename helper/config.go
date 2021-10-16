package helper

import (
	"github.com/joho/godotenv"
	"simple-ecommerce-rest-api/app/exception"
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
		exception.PanicIfInternalServerError(err)
	}
}