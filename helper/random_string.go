package helper

import (
	"math/rand"
	"time"
)

func GenerateRandomString() string {
	rand.Seed(time.Now().Unix())
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	random := make([]byte,5)
	for i, _ := range random {
		random[i] = letters[rand.Intn(len(letters))]
	}
	return string(random)
}
