// Copyright © 2025 Ping Identity Corporation

package oauth2

import (
	"testing"
	"time"

	"golang.org/x/oauth2"
)

func TestKeychainStorageDefaults(t *testing.T) {
	_, err := NewKeychainStorage("", "")
	if err == nil {
		t.Error("Expected error when creating keychain storage with empty strings, got nil")
	}

	storage, err := NewKeychainStorage("pingcli", "test-user")
	if err != nil {
		t.Fatalf("Unexpected error creating keychain storage: %v", err)
	}
	if storage.serviceName != "pingcli" {
		t.Errorf("Expected serviceName to be 'pingcli', got %q", storage.serviceName)
	}
	if storage.username != "test-user" {
		t.Errorf("Expected username to be 'test-user', got %q", storage.username)
	}
}

func TestKeychainStorageOperations(t *testing.T) {
	tests := []struct {
		name      string
		operation func(t *testing.T, storage *KeychainStorage)
	}{
		{
			name: "SaveTokenAndLoadToken",
			operation: func(t *testing.T, storage *KeychainStorage) {
				// Create a test token
				token := &oauth2.Token{
					AccessToken:  "test-access-token",
					RefreshToken: "test-refresh-token",
					TokenType:    "Bearer",
					Expiry:       time.Now().Add(time.Hour),
				}

				// Save the token
				err := storage.SaveToken(token)
				if err != nil {
					t.Fatalf("Failed to save token: %v", err)
				}

				// Load the token
				loadedToken, err := storage.LoadToken()
				if err != nil {
					t.Fatalf("Failed to load token: %v", err)
				}
				if loadedToken == nil {
					t.Fatal("Loaded token should not be nil")
				}

				// Verify token contents
				if token.AccessToken != loadedToken.AccessToken {
					t.Errorf("Expected AccessToken %q, got %q", token.AccessToken, loadedToken.AccessToken)
				}
				if token.RefreshToken != loadedToken.RefreshToken {
					t.Errorf("Expected RefreshToken %q, got %q", token.RefreshToken, loadedToken.RefreshToken)
				}
				if token.TokenType != loadedToken.TokenType {
					t.Errorf("Expected TokenType %q, got %q", token.TokenType, loadedToken.TokenType)
				}
				if !token.Expiry.Equal(loadedToken.Expiry) {
					t.Errorf("Expected Expiry %v, got %v", token.Expiry, loadedToken.Expiry)
				}
			},
		},
		{
			name: "HasToken",
			operation: func(t *testing.T, storage *KeychainStorage) {
				// Create a test token
				token := &oauth2.Token{
					AccessToken:  "test-access-token",
					RefreshToken: "test-refresh-token",
					TokenType:    "Bearer",
					Expiry:       time.Now().Add(time.Hour),
				}

				// Save the token
				err := storage.SaveToken(token)
				if err != nil {
					t.Fatalf("Failed to save token: %v", err)
				}

				// Should have token
				hasToken, err := storage.HasToken()
				if err != nil {
					t.Fatalf("Failed to check token: %v", err)
				}
				if !hasToken {
					t.Error("Expected HasToken to be true")
				}
			},
		},
		{
			name: "ClearToken",
			operation: func(t *testing.T, storage *KeychainStorage) {
				// Create a test token
				token := &oauth2.Token{
					AccessToken:  "test-access-token",
					RefreshToken: "test-refresh-token",
					TokenType:    "Bearer",
					Expiry:       time.Now().Add(time.Hour),
				}

				// Save the token
				err := storage.SaveToken(token)
				if err != nil {
					t.Fatalf("Failed to save token: %v", err)
				}

				// Clear the token
				err = storage.ClearToken()
				if err != nil {
					t.Fatalf("Failed to clear token: %v", err)
				}

				// Verify token is gone
				hasToken, err := storage.HasToken()
				if err != nil {
					t.Fatalf("Failed to check token: %v", err)
				}
				if hasToken {
					t.Error("Expected HasToken to be false")
				}

				// Load should fail since token doesn't exist
				loadedToken, err := storage.LoadToken()
				if err == nil {
					t.Error("Expected error when loading non-existent token, got nil")
				}
				if loadedToken != nil {
					t.Error("Expected loaded token to be nil")
				}
			},
		},
		{
			name: "SaveTokenWithNilToken",
			operation: func(t *testing.T, storage *KeychainStorage) {
				err := storage.SaveToken(nil)
				if err == nil {
					t.Error("Expected error when saving nil token")
					return
				}
				if !containsString(err.Error(), "token cannot be nil") {
					t.Errorf("Expected error to contain 'token cannot be nil', got %q", err.Error())
				}
			},
		},
		{
			name: "ClearTokenWhenNoTokenExists",
			operation: func(t *testing.T, storage *KeychainStorage) {
				// Should not error when clearing non-existent token
				err := storage.ClearToken()
				if err != nil {
					t.Errorf("Unexpected error when clearing non-existent token: %v", err)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Use a test-specific service name to avoid conflicts
			storage, err := NewKeychainStorage("pingcli-test", "test-user-"+tt.name)
			if err != nil {
				t.Fatalf("Failed to create keychain storage: %v", err)
			}

			// Clean up any existing test data
			_ = storage.ClearToken()
			defer storage.ClearToken() // Clean up after test

			tt.operation(t, storage)
		})
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(substr) > 0 && stringContains(s, substr)))
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
