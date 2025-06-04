// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"crypto/rand"
	"math/big"
)

// randomString generates a random string of a given length.
func randomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	b := make([]rune, length)
	for i := range b {
		// Generate a random index within the charset
		maxBigInt := big.NewInt(int64(len(charset)))
		n, err := rand.Int(rand.Reader, maxBigInt)
		if err != nil {
			panic(err) // In a utility function, we panic on errors
		}
		b[i] = rune(charset[n.Int64()])
	}
	return string(b)
}

func RandomResourceName() string {
	// Generate a random string of length 10
	length := 10
	return randomString(length)
}
