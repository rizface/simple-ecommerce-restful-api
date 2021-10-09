package helper

import (
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	result,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIfError(err)
	return string(result)
}

func Compare(plain string, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed),[]byte(plain))
	return err
}