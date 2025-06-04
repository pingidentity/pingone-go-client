// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"math/rand/v2"
)

// randomString generates a random string of a given length.
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]rune, length)
	for i := range b {
		b[i] = rune(charset[rand.IntN(len(charset))])
	}
	return string(b)
}

func RandomResourceName() string {
	// Generate a random string of length 10
	length := 10
	return randomString(length)
}
