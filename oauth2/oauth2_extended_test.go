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
			grantType: string(oauth2.GrantTypeAuthorizationCode),
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
