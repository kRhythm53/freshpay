package utilities

import (
	"math/rand"
	"time"
)

func RandomString(n int,prefix string) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[seededRand.Intn(len(letters))]
	}
	return prefix+"_"+string(s)
}
