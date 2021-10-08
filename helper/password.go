package helper

import "golang.org/x/crypto/bcrypt"

func Hash(password string) string {
	result,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	PanicIfError(err)
	return string(result)
}
