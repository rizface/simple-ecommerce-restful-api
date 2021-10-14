package helper

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func Hash(password string) string {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIfError(err)
	return string(result)
}

func Compare(plain string, hashed string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	if err != nil {
		return errors.New("username / password salah")
	}
	return nil
}
