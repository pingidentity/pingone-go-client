package config_test

import (
	"context"
	"testing"

	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
)

func TestDeviceAuthTokenSource(t *testing.T) {
	tests := []struct {
		name          string
		setup         func() *config.DeviceCode
		expectError   bool
		errorContains string
	}{
		{
			name: "MissingClientID",
			setup: func() *config.DeviceCode {
				return &config.DeviceCode{}
			},
			expectError:   true,
			errorContains: "client ID is required",
		},
		{
			name: "EmptyClientID",
			setup: func() *config.DeviceCode {
				clientID := ""
				return &config.DeviceCode{
					DeviceCodeClientID: &clientID,
				}
			},
			expectError:   true,
			errorContains: "client ID is required",
		},
		{
			name: "ValidClientID_WithScopes",
			setup: func() *config.DeviceCode {
				clientID := "test-client-id"
				scopes := []string{"openid", "profile"}
				return &config.DeviceCode{
					DeviceCodeClientID: &clientID,
					DeviceCodeScopes:   &scopes,
				}
			},
			expectError:   true, // Will fail due to invalid endpoint but validates inputs
			errorContains: "device auth request failed",
		},
		{
			name: "ValidClientID_WithoutScopes",
			setup: func() *config.DeviceCode {
				clientID := "test-client-id"
				emptyScopes := []string{}
				return &config.DeviceCode{
					DeviceCodeClientID: &clientID,
					DeviceCodeScopes:   &emptyScopes,
				}
			},
			expectError:   true, // Will fail due to invalid endpoint but validates inputs
			errorContains: "device auth request failed",
		},
		{
			name: "ValidClientID_WithMultipleScopes",
			setup: func() *config.DeviceCode {
				clientID := "test-client-id"
				scopes := []string{"openid", "profile", "email"}
				return &config.DeviceCode{
					DeviceCodeClientID: &clientID,
					DeviceCodeScopes:   &scopes,
				}
			},
			expectError:   true, // Valid config but will fail auth with fake credentials
			errorContains: "device auth request failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deviceCode := tt.setup()
			testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

			_, err := deviceCode.DeviceAuthTokenSource(context.Background(), testEndpoints)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
					return
				}
				if tt.errorContains != "" && !containsString(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain %q, got %q", tt.errorContains, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}

func TestDeviceAuthTokenSource_NilContext(t *testing.T) {
	clientID := "test-client-id"
	scopes := []string{"openid"}
	deviceCode := &config.DeviceCode{
		DeviceCodeClientID: &clientID,
		DeviceCodeScopes:   &scopes,
	}
	testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

	// This should handle nil context gracefully
	_, err := deviceCode.DeviceAuthTokenSource(nil, testEndpoints)
	if err == nil {
		t.Error("Expected error with nil context")
	}
}

func TestDeviceAuthTokenSource_CanceledContext(t *testing.T) {
	clientID := "test-client-id"
	scopes := []string{"openid"}
	deviceCode := &config.DeviceCode{
		DeviceCodeClientID: &clientID,
		DeviceCodeScopes:   &scopes,
	}
	testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

	// Create a context that's already canceled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := deviceCode.DeviceAuthTokenSource(ctx, testEndpoints)
	if err == nil {
		t.Error("Expected error with canceled context")
	}
}
