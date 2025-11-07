// Copyright © 2025 Ping Identity Corporation
package config_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
)

func TestAuthCodeTokenSource(t *testing.T) {
	tests := []struct {
		name          string
		setup         func() *config.AuthCode
		expectError   bool
		errorContains string
	}{
		{
			name: "MissingClientID",
			setup: func() *config.AuthCode {
				return &config.AuthCode{}
			},
			expectError:   true,
			errorContains: "client ID is required",
		},
		{
			name: "EmptyClientID",
			setup: func() *config.AuthCode {
				clientID := ""
				return &config.AuthCode{
					AuthCodeClientID: &clientID,
				}
			},
			expectError:   true,
			errorContains: "client ID is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authCode := tt.setup()
			testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

			// Use a short timeout to prevent tests from hanging
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			_, err := authCode.AuthCodeTokenSource(ctx, testEndpoints)

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

func TestCallbackHandling(t *testing.T) {
	tests := []struct {
		name        string
		queryParams string
		expectCode  string
		expectError bool
	}{
		{
			name:        "SuccessfulCallbackWithCode",
			queryParams: "?code=test-auth-code&state=test-state",
			expectCode:  "test-auth-code",
			expectError: false,
		},
		{
			name:        "CallbackWithError",
			queryParams: "?error=access_denied&error_description=User%20denied%20access",
			expectError: true,
		},
		{
			name:        "CallbackWithErrorNoDescription",
			queryParams: "?error=access_denied",
			expectError: true,
		},
		{
			name:        "CallbackWithoutCode",
			queryParams: "?state=test-state",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expectError {
				errorReceived := make(chan bool, 1)

				// Use a test server to simulate the callback
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// Simulate error callback
					query := r.URL.Query()
					if errCode := query.Get("error"); errCode != "" {
						errorReceived <- true
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte("error"))
					} else if query.Get("code") == "" {
						// Missing code is also an error
						errorReceived <- true
						w.WriteHeader(http.StatusBadRequest)
						_, _ = w.Write([]byte("missing code"))
					}
				}))
				defer server.Close() // Simulate error callback request
				resp, err := http.Get(server.URL + tt.queryParams)
				if err != nil {
					t.Fatalf("Failed to make request: %v", err)
				}
				defer func() {
					if closeErr := resp.Body.Close(); closeErr != nil {
						t.Logf("Warning: failed to close response body: %v", closeErr)
					}
				}() // Check that error was received
				select {
				case <-errorReceived:
					// Expected error received
				case <-time.After(100 * time.Millisecond):
					t.Fatal("timeout waiting for error")
				}
			} else {
				codeChan := make(chan string, 1)

				// Use a test server to simulate the callback
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					// Simulate authorization callback
					query := r.URL.Query()
					code := query.Get("code")
					if code != "" {
						codeChan <- code
						w.WriteHeader(http.StatusOK)
						_, _ = w.Write([]byte("success"))
					}
				}))
				defer server.Close() // Simulate callback request
				resp, err := http.Get(server.URL + tt.queryParams)
				if err != nil {
					t.Fatalf("Failed to make request: %v", err)
				}
				defer func() {
					if closeErr := resp.Body.Close(); closeErr != nil {
						t.Logf("Warning: failed to close response body: %v", closeErr)
					}
				}() // Check that code was received
				select {
				case code := <-codeChan:
					if code != tt.expectCode {
						t.Errorf("Expected code %q, got %q", tt.expectCode, code)
					}
				case <-time.After(1 * time.Second):
					t.Fatal("timeout waiting for authorization code")
				}
			}
		})
	}
}
