// Copyright © 2025 Ping Identity Corporation

package oauth2_test

import (
	"testing"

	"github.com/pingidentity/pingone-go-client/oauth2"
)

func TestIsValidGrantType(t *testing.T) {
	tests := []struct {
		name      string
		grantType string
		expected  bool
	}{
		{
			name:      "ValidClientCredentials",
			grantType: string(oauth2.GrantTypeClientCredentials),
			expected:  true,
		},
		{
			name:      "ValidAuthorizationCode",
			grantType: string(oauth2.GrantTypeAuthCode),
			expected:  true,
		},
		{
			name:      "ValidDeviceCode",
			grantType: string(oauth2.GrantTypeDeviceCode),
			expected:  true,
		},
		{
			name:      "InvalidGrantType",
			grantType: "invalid_grant_type",
			expected:  false,
		},
		{
			name:      "EmptyGrantType",
			grantType: "",
			expected:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := oauth2.IsValidGrantType(tt.grantType)
			if result != tt.expected {
				t.Errorf("IsValidGrantType(%q) = %v, expected %v", tt.grantType, result, tt.expected)
			}
		})
	}
}

func TestIsValidTopLevelDomain(t *testing.T) {
	tests := []struct {
		name     string
		domain   string
		expected bool
	}{
		{
			name:     "ValidComDomain",
			domain:   "com",
			expected: true,
		},
		{
			name:     "ValidCaDomain",
			domain:   "ca",
			expected: true,
		},
		{
			name:     "ValidEuDomain",
			domain:   "eu",
			expected: true,
		},
		{
			name:     "ValidAsiaDomain",
			domain:   "asia",
			expected: true,
		},
		{
			name:     "InvalidDomain",
			domain:   "invalid",
			expected: false,
		},
		{
			name:     "EmptyDomain",
			domain:   "",
			expected: false,
		},
		{
			name:     "CaseSensitiveTest",
			domain:   "COM",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := oauth2.IsValidTopLevelDomain(tt.domain)
			if result != tt.expected {
				t.Errorf("IsValidTopLevelDomain(%q) = %v, expected %v", tt.domain, result, tt.expected)
			}
		})
	}
}

// Removed TestNewTokenStorageConfig - TokenStorageConfig was simplified away

func TestPKCECodeGeneration(t *testing.T) {
	tests := []struct {
		name      string
		operation func(t *testing.T)
	}{
		{
			name: "GenerateCodeVerifier",
			operation: func(t *testing.T) {
				verifier, err := oauth2.GenerateCodeVerifier()
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}
				if verifier == "" {
					t.Error("Expected non-empty verifier")
				}
				if len(verifier) < 40 {
					t.Errorf("Code verifier should be at least 40 characters, got %d", len(verifier))
				}

				// Verify it's URL-safe base64
				if !matchesRegexp(`^[A-Za-z0-9_\-]+$`, verifier) {
					t.Errorf("Code verifier should be URL-safe, got %q", verifier)
				}
			},
		},
		{
			name: "GenerateCodeChallenge",
			operation: func(t *testing.T) {
				verifier, err := oauth2.GenerateCodeVerifier()
				if err != nil {
					t.Fatalf("Expected no error, got %v", err)
				}

				challenge := oauth2.GenerateCodeChallenge(verifier)

				if challenge == "" {
					t.Error("Expected non-empty challenge")
				}
				if verifier == challenge {
					t.Error("Challenge should be different from verifier")
				}

				// Verify it's URL-safe base64
				if !matchesRegexp(`^[A-Za-z0-9_\-]+$`, challenge) {
					t.Errorf("Code challenge should be URL-safe, got %q", challenge)
				}
			},
		},
		{
			name: "GenerateCodeChallengeWithEmptyVerifier",
			operation: func(t *testing.T) {
				challenge := oauth2.GenerateCodeChallenge("")
				if challenge == "" {
					t.Error("Should handle empty verifier gracefully")
				}
			},
		},
		{
			name: "GenerateMultipleVerifiersAreDifferent",
			operation: func(t *testing.T) {
				verifier1, err1 := oauth2.GenerateCodeVerifier()
				verifier2, err2 := oauth2.GenerateCodeVerifier()

				if err1 != nil {
					t.Errorf("Expected no error for first verifier, got %v", err1)
				}
				if err2 != nil {
					t.Errorf("Expected no error for second verifier, got %v", err2)
				}
				if verifier1 == verifier2 {
					t.Error("Each generated verifier should be unique")
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.operation(t)
		})
	}
}

// Helper function to check if a string matches a regex pattern
func matchesRegexp(pattern, s string) bool {
	for _, char := range s {
		isValid := (char >= 'A' && char <= 'Z') ||
			(char >= 'a' && char <= 'z') ||
			(char >= '0' && char <= '9') ||
			char == '_' || char == '-'
		if !isValid {
			return false
		}
	}
	return true
}
