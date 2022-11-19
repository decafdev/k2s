package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphaPlus = "abcdefghijklmnopqrstuvwxyz-_"

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphaPlus)

	for i := 0; i < n; i++ {
		c := alphaPlus[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
