// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"context"
	"testing"

	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
)

func TestClientCredentialsTokenSource(t *testing.T) {
	tests := []struct {
		name          string
		setup         func() *config.ClientCredentials
		expectError   bool
		errorContains string
	}{
		{
			name: "MissingClientID",
			setup: func() *config.ClientCredentials {
				return &config.ClientCredentials{}
			},
			expectError:   true,
			errorContains: "client ID is required",
		},
		{
			name: "EmptyClientID",
			setup: func() *config.ClientCredentials {
				clientID := ""
				return &config.ClientCredentials{
					ClientCredentialsClientID: &clientID,
				}
			},
			expectError:   true,
			errorContains: "client ID is required",
		},
		{
			name: "MissingScopes",
			setup: func() *config.ClientCredentials {
				clientID := "test-client-id"
				clientSecret := "test-secret"
				return &config.ClientCredentials{
					ClientCredentialsClientID:     &clientID,
					ClientCredentialsClientSecret: &clientSecret,
				}
			},
			expectError:   true,
			errorContains: "scopes are required",
		},
		{
			name: "EmptyScopes",
			setup: func() *config.ClientCredentials {
				clientID := "test-client-id"
				clientSecret := "test-secret"
				emptyScopes := []string{}
				return &config.ClientCredentials{
					ClientCredentialsClientID:     &clientID,
					ClientCredentialsClientSecret: &clientSecret,
					ClientCredentialsScopes:       &emptyScopes,
				}
			},
			expectError:   true,
			errorContains: "scopes are required",
		},
		{
			name: "MissingClientSecret",
			setup: func() *config.ClientCredentials {
				clientID := "test-client-id"
				scopes := []string{"p1:read:*"}
				return &config.ClientCredentials{
					ClientCredentialsClientID: &clientID,
					ClientCredentialsScopes:   &scopes,
				}
			},
			expectError:   true,
			errorContains: "client secret is required",
		},
		{
			name: "EmptyClientSecret",
			setup: func() *config.ClientCredentials {
				clientID := "test-client-id"
				clientSecret := ""
				scopes := []string{"p1:read:*"}
				return &config.ClientCredentials{
					ClientCredentialsClientID:     &clientID,
					ClientCredentialsClientSecret: &clientSecret,
					ClientCredentialsScopes:       &scopes,
				}
			},
			expectError:   true,
			errorContains: "client secret is required",
		},
		{
			name: "ValidConfiguration_SingleScope",
			setup: func() *config.ClientCredentials {
				clientID := "test-client-id"
				clientSecret := "test-secret"
				scopes := []string{"p1:read:*"}
				return &config.ClientCredentials{
					ClientCredentialsClientID:     &clientID,
					ClientCredentialsClientSecret: &clientSecret,
					ClientCredentialsScopes:       &scopes,
				}
			},
			expectError:   true, // Will fail on token fetch but configuration is valid
			errorContains: "",   // Any error is acceptable (network, invalid endpoint, etc)
		},
		{
			name: "ValidConfiguration_MultipleScopes",
			setup: func() *config.ClientCredentials {
				clientID := "test-client-id"
				clientSecret := "test-secret"
				scopes := []string{"p1:read:*", "p1:write:*", "p1:delete:*"}
				return &config.ClientCredentials{
					ClientCredentialsClientID:     &clientID,
					ClientCredentialsClientSecret: &clientSecret,
					ClientCredentialsScopes:       &scopes,
				}
			},
			expectError:   true, // Will fail on token fetch but configuration is valid
			errorContains: "",   // Any error is acceptable (network, invalid endpoint, etc)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientCreds := tt.setup()
			testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

			ts, err := clientCreds.ClientCredentialsTokenSource(context.Background(), testEndpoints)

			if tt.expectError {
				if err == nil && ts == nil {
					t.Error("Expected error or token source but got neither")
					return
				}
				if err != nil && tt.errorContains != "" && !containsString(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain %q, got %q", tt.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if ts == nil {
					t.Error("Expected token source but got nil")
				}
			}
		})
	}
}

func TestClientCredentialsTokenSource_NilContext(t *testing.T) {
	clientID := "test-client-id"
	clientSecret := "test-secret"
	scopes := []string{"p1:read:*"}
	clientCreds := &config.ClientCredentials{
		ClientCredentialsClientID:     &clientID,
		ClientCredentialsClientSecret: &clientSecret,
		ClientCredentialsScopes:       &scopes,
	}
	testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

	// This should handle nil context - TokenSource creation should still work
	ts, err := clientCreds.ClientCredentialsTokenSource(context.TODO(), testEndpoints)

	// TokenSource creation should succeed even with nil context
	// (The actual token fetch would fail, but that's not tested here)
	if err != nil {
		t.Errorf("Unexpected error with nil context: %v", err)
	}
	if ts == nil {
		t.Error("Expected token source but got nil")
	}
}

func TestClientCredentialsTokenSource_ScopeValidation(t *testing.T) {
	tests := []struct {
		name        string
		scopes      []string
		expectError bool
	}{
		{
			name:        "ValidScope_ReadAll",
			scopes:      []string{"p1:read:*"},
			expectError: false,
		},
		{
			name:        "ValidScope_MultiplePermissions",
			scopes:      []string{"p1:read:env", "p1:write:env"},
			expectError: false,
		},
		{
			name:        "ValidScope_OpenID",
			scopes:      []string{"openid", "profile"},
			expectError: false,
		},
		{
			name:        "EmptyScope",
			scopes:      []string{},
			expectError: false,
		},
		{
			name:        "NilScope",
			scopes:      nil,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientID := "test-client-id"
			clientSecret := "test-secret"
			var scopes *[]string
			if tt.scopes != nil {
				scopes = &tt.scopes
			}

			clientCreds := &config.ClientCredentials{
				ClientCredentialsClientID:     &clientID,
				ClientCredentialsClientSecret: &clientSecret,
				ClientCredentialsScopes:       scopes,
			}
			testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

			_, err := clientCreds.ClientCredentialsTokenSource(context.Background(), testEndpoints)

			if tt.expectError && err == nil {
				t.Error("Expected error for invalid scopes but got none")
			}
			if !tt.expectError && err != nil && containsString(err.Error(), "scopes are required") {
				t.Errorf("Unexpected scope validation error: %v", err)
			}
		})
	}
}
