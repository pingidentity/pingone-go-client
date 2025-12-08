// Copyright © 2025 Ping Identity Corporation

package oauth2

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zalando/go-keyring"
	"golang.org/x/oauth2"
)

// KeychainStorage implements TokenStorage using the system keychain
type KeychainStorage struct {
	serviceName string
	username    string
}

// NewKeychainStorage creates a new KeychainStorage instance
// Both serviceName and username must be non-empty strings
func NewKeychainStorage(serviceName, username string) (*KeychainStorage, error) {
	if serviceName == "" {
		return nil, fmt.Errorf("serviceName cannot be empty")
	}
	if username == "" {
		return nil, fmt.Errorf("username cannot be empty")
	}
	return &KeychainStorage{
		serviceName: serviceName,
		username:    username,
	}, nil
}

// SaveToken stores the OAuth2 token in the system keychain
func (k *KeychainStorage) SaveToken(token *oauth2.Token) error {
	if token == nil {
		return errors.New("token cannot be nil")
	}

	tokenJSON, err := json.Marshal(token)
	if err != nil {
		return fmt.Errorf("failed to marshal token: %w", err)
	}

	err = keyring.Set(k.serviceName, k.username, string(tokenJSON))
	if err != nil {
		return fmt.Errorf("failed to save token to keychain: %w", err)
	}

	return nil
}

// LoadToken retrieves the OAuth2 token from the system keychain
func (k *KeychainStorage) LoadToken() (*oauth2.Token, error) {
	tokenJSON, err := keyring.Get(k.serviceName, k.username)
	if err != nil {
		return nil, fmt.Errorf("failed to load token from keychain: %w", err)
	}

	var token oauth2.Token
	err = json.Unmarshal([]byte(tokenJSON), &token)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal token: %w", err)
	}

	return &token, nil
}

// ClearToken removes the OAuth2 token from the system keychain
func (k *KeychainStorage) ClearToken() error {
	err := keyring.Delete(k.serviceName, k.username)
	if err != nil && !errors.Is(err, keyring.ErrNotFound) {
		return fmt.Errorf("failed to clear token from keychain: %w", err)
	}
	return nil
}

// HasToken checks if a token exists in the system keychain
func (k *KeychainStorage) HasToken() (bool, error) {
	_, err := keyring.Get(k.serviceName, k.username)
	if err != nil {
		if errors.Is(err, keyring.ErrNotFound) {
			return false, nil
		}
		return false, fmt.Errorf("failed to check token in keychain: %w", err)
	}
	return true, nil
}

// GenerateKeychainAccountName creates a unique account name based on environment ID, client ID, and grant type.
// Optionally, a suffix can be provided to append to the generated token key for disambiguation across contexts
// (e.g., "_pingone_device_code_default"). Existing callers remain compatible.
func GenerateKeychainAccountName(environmentID, clientID, grantType string, optionalSuffix ...string) string {
	if environmentID == "" && clientID == "" && grantType == "" {
		// When no inputs are provided, return a stable default (with optional suffix if specified)
		base := "default-token"
		if len(optionalSuffix) > 0 && optionalSuffix[0] != "" {
			return base + optionalSuffix[0]
		}
		return base
	}

	// Create a hash of environment ID + client ID + grant type for uniqueness
	var b []byte
	b = fmt.Appendf(b, "%s:%s:%s", environmentID, clientID, grantType)
	hash := sha256.Sum256(b)
	base := fmt.Sprintf("token-%x", hash[:8]) // Use first 8 bytes of hash for shorter key
	if len(optionalSuffix) > 0 && optionalSuffix[0] != "" {
		return base + optionalSuffix[0]
	}
	return base
}
