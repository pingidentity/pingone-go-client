// Copyright © 2025 Ping Identity Corporation
package config

import (
	"context"
	"testing"
	"time"

	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"golang.org/x/oauth2"
)

func TestKeychainTokenSource_SavesRefreshedToken(t *testing.T) {
	// Test if keychain is available by attempting to create and use storage
	testStorage, err := svcOAuth2.NewKeychainStorage("pingcli-test-availability", "test-check")
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

	// Create a mock token source that returns a new token
	baseTokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken:  "refreshed-access-token",
		RefreshToken: "refreshed-refresh-token",
		Expiry:       time.Now().Add(1 * time.Hour),
		TokenType:    "Bearer",
	})

	// Create keychain storage for testing
	keychainStorage, err := svcOAuth2.NewKeychainStorage("test-service", "test-token-key")
	if err != nil {
		t.Fatalf("Failed to create keychain storage: %v", err)
	}

	// Clean up before and after test
	_ = keychainStorage.ClearToken()
	defer func() {
		_ = keychainStorage.ClearToken()
	}()

	// Create the keychain token source wrapper
	ts := &keychainTokenSource{
		base:            baseTokenSource,
		keychainStorage: keychainStorage,
		tokenKey:        "test-token-key",
	}

	// Get token - this should save it to keychain
	token, err := ts.Token()
	if err != nil {
		t.Fatalf("Failed to get token: %v", err)
	}

	if token.AccessToken != "refreshed-access-token" {
		t.Errorf("Expected access token 'refreshed-access-token', got %q", token.AccessToken)
	}

	// Verify token was saved to keychain
	savedToken, err := keychainStorage.LoadToken()
	if err != nil {
		t.Fatalf("Failed to load token from keychain: %v", err)
	}

	if savedToken.AccessToken != "refreshed-access-token" {
		t.Errorf("Expected saved access token 'refreshed-access-token', got %q", savedToken.AccessToken)
	}

	if savedToken.RefreshToken != "refreshed-refresh-token" {
		t.Errorf("Expected saved refresh token 'refreshed-refresh-token', got %q", savedToken.RefreshToken)
	}
}

func TestTokenRefresh_WithExpiredAccessToken(t *testing.T) {
	// This is an integration test demonstrating the refresh flow
	t.Skip("Integration test - requires actual OAuth2 server")

	// Create a token with expired access token but valid refresh token
	expiredToken := &oauth2.Token{
		AccessToken:  "expired-access-token",
		RefreshToken: "valid-refresh-token",
		Expiry:       time.Now().Add(-1 * time.Hour), // Expired 1 hour ago
		TokenType:    "Bearer",
	}

	// In a real scenario, oauth2.Config.TokenSource would use the refresh token
	// to get a new access token when the old one is expired
	config := &oauth2.Config{
		ClientID: "test-client-id",
		Endpoint: oauth2.Endpoint{
			TokenURL: "https://auth.pingone.com/test-env/as/token",
		},
	}

	// Create token source that will refresh automatically
	baseTokenSource := config.TokenSource(context.Background(), expiredToken)

	// Wrap with keychain saving
	keychainStorage, _ := svcOAuth2.NewKeychainStorage("test-service", "test-token-key")
	defer func() {
		_ = keychainStorage.ClearToken()
	}()

	ts := &keychainTokenSource{
		base:            baseTokenSource,
		keychainStorage: keychainStorage,
		tokenKey:        "test-token-key",
	}

	// Get token - this would trigger refresh in real scenario
	token, err := ts.Token()
	if err != nil {
		t.Logf("Expected error in test mode without real server: %v", err)
		return
	}

	// In real scenario, this would be a new access token
	if token.AccessToken == "expired-access-token" {
		t.Error("Token was not refreshed")
	}
}
