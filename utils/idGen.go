package utils

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

func GenerateID() string {
	t := time.Now().UnixNano()/1000
	ans := ""
	for i := 0;i<10;i++ {
		char := letters[t%int64(len(letters))]
		t /= 10
		ans = string(char) + ans
	}
	return ans
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
