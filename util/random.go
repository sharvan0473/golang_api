package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphobet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphobet)
	for i := 0; i < n; i++ {
		c := alphobet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomMoney() int64 {
	return RandomInt(0, 100)
}

func RandomCurrency() string {
	curr := []string{"EUR", "USD", "CAD"}
	n := len(curr)
	return curr[rand.Intn(n)]
}
