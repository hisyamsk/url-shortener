package helpers

import (
	"math/rand"
)

var strLength = 6

func GenerateRandomString() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	str := make([]rune, strLength)
	for i := range str {
		str[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(str)
}
