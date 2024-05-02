package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var rnd *rand.Rand

func init() {
	rsource := rand.NewSource(time.Now().UnixNano())
	rnd = rand.New(rsource)
}

// RandInt generate rand int with max and min
func RandInt(min int64, max int64) int64 {
	return min + rnd.Int63n(max-min+1)
}

// RandStr generate randon string by length
func RandString(num int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < num; i++ {
		sb.WriteByte(alphabet[rnd.Intn(k)])
	}
	return sb.String()
}

func RandUsername() string {
	return RandString(6)
}

func RandAmount() int64 {
	return RandInt(0, 1000)
}

func RandCurrency() string {
	allCurrencies := []string{"USD", "EUR"}
	return allCurrencies[rnd.Intn(len(allCurrencies))]
}
