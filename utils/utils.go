package utils

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var strings = `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`

func RandString(long int) string {
	if long <= 0 {
		return ""
	}
	b := make([]byte, long)
	for i := range b {
		b[i] = strings[rand.Intn(len(strings))]
	}
	return string(b)
}
