package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var r = rand.New(rand.NewSource(time.Now().UnixNano()))


func RandomInt (min, max int64) int64 {
	return min + r.Int63n(max - min + 1)
}

func RandomString (n int) string {
	var result strings.Builder

	for i := n; i>0; i-- {
		result.WriteByte(alphabet[r.Intn(26)])
	}

	return result.String()
}

func RandomOwner(n int) string {
	return RandomString(n)
}

func RandomAmount() int64 {
	return RandomInt(0, 10000)
}

func RandomCurrency() string {
	currencies := []string{"INR", "USD", "CAD", "YEN", "THAI"}
	n := len(currencies)

	return currencies[r.Intn(n)]
}