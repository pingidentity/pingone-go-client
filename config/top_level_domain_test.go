// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingone-go-client/config"
)

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
			result := config.IsValidTopLevelDomain(tt.domain)
			if result != tt.expected {
				t.Errorf("IsValidTopLevelDomain(%q) = %v, expected %v", tt.domain, result, tt.expected)
			}
		})
	}
}
