// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oauth2/endpoints"
	"golang.org/x/oauth2"
)

// newDeviceAuthServer returns an httptest server that answers both legs of the RFC 8628 flow —
// the device-authorization request and the token request — with fixed responses, so
// DeviceAuthTokenSource can run to completion without a live PingOne server.
// verificationURI deliberately pairs a non-http(s) scheme with a loopback host so that a default
// handler exercising it does not attempt to open a real browser. interval is the polling interval
// (in seconds) the token leg advertises, letting callers trade poll delay against test runtime.
func newDeviceAuthServer(t *testing.T, deviceCode, userCode, verificationURI string, interval int64) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if err := r.ParseForm(); err != nil || r.PostForm.Get("grant_type") == "" {
			// Device authorization request (no grant_type).
			_ = json.NewEncoder(w).Encode(map[string]any{
				"device_code":      deviceCode,
				"user_code":        userCode,
				"verification_uri": verificationURI,
				"expires_in":       600,
				"interval":         interval,
			})
			return
		}

		// Token request (grant_type=urn:ietf:params:oauth:grant-type:device_code).
		_ = json.NewEncoder(w).Encode(map[string]any{
			"access_token": "device-access-token",
			"token_type":   "Bearer",
			"expires_in":   3600,
		})
	}))
	t.Cleanup(srv.Close)
	return srv
}

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
			testEndpoints := endpoints.PingOneEndpoint("auth.pingone.com")

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

// TestDeviceAuthTokenSource_DefaultHandlerHonorsOutput verifies that with no custom handler,
// the default device code prompt handler is selected and writes its progress messages to Output.
// The flow runs to completion against the stub server, with the non-http(s) verification URI
// keeping browser.Open from opening a real browser.
func TestDeviceAuthTokenSource_DefaultHandlerHonorsOutput(t *testing.T) {
	srv := newDeviceAuthServer(t, "test-device-code", "test-user-code", "ftp://127.0.0.1/device", 1)

	clientID := "test-client-id"
	var buf bytes.Buffer

	deviceCode := &config.DeviceCode{
		DeviceCodeClientID: &clientID,
		Output:             &buf,
	}
	testEndpoint := oauth2.Endpoint{DeviceAuthURL: srv.URL, TokenURL: srv.URL}

	ts, err := deviceCode.DeviceAuthTokenSource(context.Background(), testEndpoint)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	token, err := ts.Token()
	if err != nil {
		t.Fatalf("unexpected error getting token: %v", err)
	}
	if token.AccessToken != "device-access-token" {
		t.Errorf("expected device-access-token, got %q", token.AccessToken)
	}

	// The default handler is invoked with the stubbed verification URI and user code.
	out := buf.String()
	if !strings.Contains(out, "Device Authorization Required") {
		t.Errorf("expected default handler progress output, got %q", out)
	}
	if !strings.Contains(out, "ftp://127.0.0.1/device") {
		t.Errorf("expected verification URI in default handler output, got %q", out)
	}
	if !strings.Contains(out, "test-user-code") {
		t.Errorf("expected user code in default handler output, got %q", out)
	}
}

func TestDeviceAuthTokenSource_NilContext(t *testing.T) {
	clientID := "test-client-id"
	scopes := []string{"openid"}
	deviceCode := &config.DeviceCode{
		DeviceCodeClientID: &clientID,
		DeviceCodeScopes:   &scopes,
	}
	testEndpoints := endpoints.PingOneEndpoint("auth.pingone.com")

	// This should handle nil context gracefully
	_, err := deviceCode.DeviceAuthTokenSource(context.TODO(), testEndpoints)
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
	testEndpoints := endpoints.PingOneEndpoint("auth.pingone.com")

	// Create a context that's already canceled
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := deviceCode.DeviceAuthTokenSource(ctx, testEndpoints)
	if err == nil {
		t.Error("Expected error with canceled context")
	}
}
