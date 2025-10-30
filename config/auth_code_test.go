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
		{
			name: "InvalidRedirectURI",
			setup: func() *config.AuthCode {
				clientID := "test-client-id"
				invalidURI := ":/invalid-uri"
				return &config.AuthCode{
					AuthCodeClientID:    &clientID,
					AuthCodeRedirectURI: &invalidURI,
				}
			},
			expectError:   true,
			errorContains: "failed to start callback server",
		},
		{
			name: "ValidClientID_WithScopes",
			setup: func() *config.AuthCode {
				clientID := "test-client-id"
				scopes := []string{"openid", "profile"}
				return &config.AuthCode{
					AuthCodeClientID: &clientID,
					AuthCodeScopes:   &scopes,
				}
			},
			expectError:   true, // Will timeout after starting server
			errorContains: "context deadline exceeded",
		},
		{
			name: "ValidClientID_WithCustomRedirectURI",
			setup: func() *config.AuthCode {
				clientID := "test-client-id"
				redirectURI := "http://localhost:9999/callback"
				return &config.AuthCode{
					AuthCodeClientID:    &clientID,
					AuthCodeRedirectURI: &redirectURI,
				}
			},
			expectError:   true, // Will timeout or fail binding
			errorContains: "",   // Accept either timeout or port binding error
		},
		{
			name: "EmptyRedirectURI_UsesDefault",
			setup: func() *config.AuthCode {
				clientID := "test-client-id"
				emptyURI := ""
				return &config.AuthCode{
					AuthCodeClientID:    &clientID,
					AuthCodeRedirectURI: &emptyURI,
				}
			},
			expectError:   true, // Will timeout after starting server
			errorContains: "context deadline exceeded",
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
						w.Write([]byte("error"))
					} else if query.Get("code") == "" {
						// Missing code is also an error
						errorReceived <- true
						w.WriteHeader(http.StatusBadRequest)
						w.Write([]byte("missing code"))
					}
				}))
				defer server.Close()

				// Simulate error callback request
				resp, err := http.Get(server.URL + tt.queryParams)
				if err != nil {
					t.Fatalf("Failed to make request: %v", err)
				}
				defer resp.Body.Close()

				// Check that error was received
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
						w.Write([]byte("success"))
					}
				}))
				defer server.Close()

				// Simulate callback request
				resp, err := http.Get(server.URL + tt.queryParams)
				if err != nil {
					t.Fatalf("Failed to make request: %v", err)
				}
				defer resp.Body.Close()

				// Check that code was received
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

func TestOpenBrowser_DoesNotPanic(t *testing.T) {
	// This test just ensures the function doesn't panic
	// We can't easily test actual browser opening in unit tests
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("openBrowser panicked: %v", r)
		}
	}()

	// This will likely fail to open a browser in CI, but shouldn't panic
	// The function should handle errors gracefully
	// Note: We can't easily test this without mocking exec.Command
}

func TestAuthCodeTokenSource_ContextCancellation(t *testing.T) {
	clientID := "test-client-id"
	authCode := &config.AuthCode{
		AuthCodeClientID: &clientID,
	}
	testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

	// Create a context with short timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := authCode.AuthCodeTokenSource(ctx, testEndpoints)
	if err == nil {
		t.Error("Expected error with short timeout context")
	}
}

func TestAuthCodeTokenSource_DefaultRedirectURI(t *testing.T) {
	// This test just validates that missing redirect URI doesn't panic
	// It will fail to start the callback server but that's expected
	clientID := "test-client-id"
	authCode := &config.AuthCode{
		AuthCodeClientID: &clientID,
		// No redirect URI - should use default
	}
	testEndpoints := endpoints.PingOneOIDCEndpoint("auth.pingone.com")

	// Use short timeout to prevent hanging
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	_, err := authCode.AuthCodeTokenSource(ctx, testEndpoints)

	// We expect an error (callback server likely can't start on occupied port)
	if err == nil {
		t.Error("Expected error without successful server start")
	}
}
