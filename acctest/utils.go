// Copyright © 2025 Ping Identity Corporation

package acctest

import (
	"math/rand"
)

// randomString generates a random string of a given length.
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]rune, length)
	for i := range b {
		b[i] = rune(charset[rand.Intn(len(charset))])
	}
	return string(b)
}

func RandomResourceName() string {
	// Generate a random string of length 10
	length := 10
	return randomString(length)
}
