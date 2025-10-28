// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"crypto/rand"
	"math/big"
)

// randomString generates a random string of the specified length using lowercase letters.
// The length parameter determines how many characters the returned string will contain.
// This function uses cryptographically secure random number generation and panics if
// random number generation fails, which is appropriate for utility functions.
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

// RandomResourceName generates a random resource name suitable for testing purposes.
// It returns a 10-character string composed of lowercase letters. This function is
// useful for creating unique test resource names to avoid conflicts in test environments.
func RandomResourceName() string {
	// Generate a random string of length 10
	length := 10
	return randomString(length)
}
