// Copyright © 2025 Ping Identity Corporation
package config_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
)

func TestAuthorizationCodeTokenSource(t *testing.T) {
	tests := []struct {
		name          string
		setup         func() *config.AuthorizationCode
		expectError   bool
		errorContains string
	}{
		{
			name: "MissingClientID",
			setup: func() *config.AuthorizationCode {
				return &config.AuthorizationCode{}
			},
			expectError:   true,
			errorContains: "client ID is required",
		},
		{
			name: "EmptyClientID",
			setup: func() *config.AuthorizationCode {
				clientID := ""
				return &config.AuthorizationCode{
					AuthorizationCodeClientID: &clientID,
				}
			},
			expectError:   true,
			errorContains: "client ID is required",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authorizationCode := tt.setup()
			testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

			// Use a short timeout to prevent tests from hanging
			ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
			defer cancel()

			_, err := authorizationCode.AuthorizationCodeTokenSource(ctx, testEndpoints)

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

func TestStateParameterValidation(t *testing.T) {
	tests := []struct {
		name          string
		expectedState string
		receivedState string
		expectError   bool
		errorContains string
	}{
		{
			name:          "ValidStateMatch",
			expectedState: "valid-state-parameter",
			receivedState: "valid-state-parameter",
			expectError:   false,
		},
		{
			name:          "StateMismatch",
			expectedState: "expected-state",
			receivedState: "different-state",
			expectError:   true,
			errorContains: "invalid state parameter",
		},
		{
			name:          "MissingStateInCallback",
			expectedState: "expected-state",
			receivedState: "",
			expectError:   true,
			errorContains: "invalid state parameter",
		},
		{
			name:          "EmptyExpectedState",
			expectedState: "",
			receivedState: "",
			expectError:   false,
		},
		{
			name:          "StateWithBase64Characters",
			expectedState: "state+base64/encoding==",
			receivedState: "state+base64/encoding==",
			expectError:   false,
		},
		{
			name:          "StateWithBase64CharactersMismatch",
			expectedState: "state+base64/encoding==",
			receivedState: "state+base64/encoding=different",
			expectError:   true,
			errorContains: "invalid state parameter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errorReceived := make(chan error, 1)
			codeReceived := make(chan string, 1)

			// Create a test server that simulates the callback handler
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				query := r.URL.Query()
				receivedState := query.Get("state")

				// Simulate state validation logic from startCallbackServer
				if receivedState != tt.expectedState {
					errorReceived <- http.ErrAbortHandler
					w.WriteHeader(http.StatusBadRequest)
					_, _ = w.Write([]byte("Invalid state parameter"))
					return
				}

				// If state is valid, return the code
				code := query.Get("code")
				if code != "" {
					codeReceived <- code
					w.WriteHeader(http.StatusOK)
					_, _ = w.Write([]byte("Success"))
				}
			}))
			defer server.Close()

			// Simulate callback with the received state (properly URL-encoded)
			queryParams := "?code=test-code"
			if tt.receivedState != "" {
				queryParams += "&state=" + url.QueryEscape(tt.receivedState)
			}

			resp, err := http.Get(server.URL + queryParams)
			if err != nil {
				t.Fatalf("Failed to make request: %v", err)
			}
			defer func() {
				if closeErr := resp.Body.Close(); closeErr != nil {
					t.Logf("Warning: failed to close response body: %v", closeErr)
				}
			}()

			if tt.expectError {
				// Should receive error
				select {
				case <-errorReceived:
					// Expected error received
					if resp.StatusCode != http.StatusBadRequest {
						t.Errorf("Expected status %d, got %d", http.StatusBadRequest, resp.StatusCode)
					}
				case <-time.After(100 * time.Millisecond):
					t.Fatal("timeout waiting for error")
				}
			} else {
				// Should receive code
				select {
				case code := <-codeReceived:
					if code != "test-code" {
						t.Errorf("Expected code 'test-code', got %q", code)
					}
					if resp.StatusCode != http.StatusOK {
						t.Errorf("Expected status %d, got %d", http.StatusOK, resp.StatusCode)
					}
				case <-time.After(100 * time.Millisecond):
					t.Fatal("timeout waiting for code")
				}
			}
		})
	}
}

func TestGenerateState(t *testing.T) {
	// Test that generateState produces unique values
	states := make(map[string]bool)
	iterations := 100

	for i := 0; i < iterations; i++ {
		// We can't directly call generateState as it's not exported,
		// but we can test the behavior through AuthorizationCodeTokenSource
		// For now, we'll verify that multiple calls should produce different URLs

		clientID := "test-client-id"
		authCode := &config.AuthorizationCode{
			AuthorizationCodeClientID: &clientID,
			OnOpenBrowser: func(authURL string) error {
				// Extract state from URL
				// This is a simplified test - in real scenario we'd parse the URL
				states[authURL] = true
				return nil
			},
		}

		testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")
		ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)

		// This will fail due to timeout, but will call OnOpenBrowser first
		_, _ = authCode.AuthorizationCodeTokenSource(ctx, testEndpoints)
		cancel()
	}

	// We should have multiple unique URLs (due to different state parameters)
	if len(states) < 2 {
		t.Errorf("Expected multiple unique authorization URLs, got %d", len(states))
	}
}

func TestStateParameterInAuthURL(t *testing.T) {
	clientID := "test-client-id"
	var capturedURL string

	authCode := &config.AuthorizationCode{
		AuthorizationCodeClientID: &clientID,
		OnOpenBrowser: func(authURL string) error {
			capturedURL = authURL
			return nil
		},
	}

	testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	// This will timeout but will capture the auth URL
	_, _ = authCode.AuthorizationCodeTokenSource(ctx, testEndpoints)

	// Verify that the URL contains a state parameter
	if capturedURL == "" {
		t.Fatal("No authorization URL was captured")
	}

	if !containsString(capturedURL, "state=") {
		t.Error("Authorization URL does not contain state parameter")
	}

	// Verify state parameter is not empty
	// Parse the URL to get the state value
	if idx := containsString(capturedURL, "state="); idx {
		// Find the state value (simplified check)
		if containsString(capturedURL, "state=&") || containsString(capturedURL, "state= ") {
			t.Error("State parameter appears to be empty")
		}
	}
}
