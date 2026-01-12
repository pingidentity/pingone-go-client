// Copyright © 2025 Ping Identity Corporation

package oauth2

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
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
	// Test if keychain is available by attempting to create and use storage
	testStorage, err := NewKeychainStorage("pingcli-test-availability", "test-check")
	if err != nil {
		t.Skipf("Skipping keychain tests: keychain storage not available (%v)", err)
	}

	// Try a simple operation to verify keychain is actually working
	testToken := &oauth2.Token{AccessToken: "test"}
	err = testStorage.SaveToken(testToken)
	_ = testStorage.ClearToken() // Clean up
	if err != nil {
		t.Skipf("Skipping keychain tests: keychain service not available (%v)", err)
	}

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
				require.NotNil(t, loadedToken, "Loaded token should not be nil")

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
			defer func() { _ = storage.ClearToken() }() // Clean up after test

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

func TestGenerateKeychainAccountNameWithSuffix(t *testing.T) {
	tests := []struct {
		name          string
		environmentID string
		clientID      string
		grantType     string
		suffix        string
		expected      string // We can't predict hash, but we can check format
}{
{
name:          "Basic usage",
environmentID: "env1",
clientID:      "client1",
grantType:     "client_credentials",
suffix:        "suffix1",
},
{
name:          "Empty suffix",
environmentID: "env1",
clientID:      "client1",
grantType:     "client_credentials",
suffix:        "",
},
{
name:          "All empty",
environmentID: "",
clientID:      "",
grantType:     "",
suffix:        "",
},
{
name:          "Empty base, with suffix",
environmentID: "",
clientID:      "",
grantType:     "",
suffix:        "suffix2",
},
}

for _, tt := range tests {
t.Run(tt.name, func(t *testing.T) {
result := GenerateKeychainAccountNameWithSuffix(tt.environmentID, tt.clientID, tt.grantType, tt.suffix)

// Basic format checks
if tt.environmentID == "" && tt.clientID == "" && tt.grantType == "" {
// Should use default base
if tt.suffix == "" {
if result != "default-token" {
t.Errorf("Expected 'default-token', got %q", result)
}
} else {
expected := "default-token_" + tt.suffix
if result != expected {
t.Errorf("Expected %q, got %q", expected, result)
}
}
} else {
// Should be hashed
if tt.suffix == "" {
// Format: token-<hash>
if len(result) < 14 { // token- + 8 chars
t.Errorf("Result too short: %q", result)
}
} else {
// Format: token-<hash>_<suffix>
if len(result) < 14+1+len(tt.suffix) {
t.Errorf("Result too short: %q", result)
}
}
}

if tt.suffix != "" {
// Verify suffix is present at end
expectedSuffix := "_" + tt.suffix
if len(result) < len(expectedSuffix) || result[len(result)-len(expectedSuffix):] != expectedSuffix {
t.Errorf("Result %q does not end with suffix %q", result, expectedSuffix)
}
}
})
}

// Consistency check
t.Run("Consistency", func(t *testing.T) {
r1 := GenerateKeychainAccountNameWithSuffix("env", "client", "grant", "suf")
r2 := GenerateKeychainAccountNameWithSuffix("env", "client", "grant", "suf")
if r1 != r2 {
t.Error("Function should be deterministic")
}
})
}
