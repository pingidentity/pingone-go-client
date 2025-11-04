// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"testing"

	"github.com/pingidentity/pingone-go-client/config"
)

func TestStorageType_String(t *testing.T) {
	tests := []struct {
		name        string
		storageType config.StorageType
		expected    string
	}{
		{
			name:        "StorageTypeKeychain",
			storageType: config.StorageTypeKeychain,
			expected:    "keychain",
		},
		{
			name:        "StorageTypeFile",
			storageType: config.StorageTypeFile,
			expected:    "file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.storageType.String()
			if result != tt.expected {
				t.Errorf("Expected String() to return %q, got %q", tt.expected, result)
			}
		})
	}
}

func TestStorageType_IsValid(t *testing.T) {
	tests := []struct {
		name        string
		storageType config.StorageType
		expected    bool
	}{
		{
			name:        "StorageTypeKeychain is valid",
			storageType: config.StorageTypeKeychain,
			expected:    true,
		},
		{
			name:        "StorageTypeFile is valid",
			storageType: config.StorageTypeFile,
			expected:    true,
		},
		{
			name:        "Empty string is invalid",
			storageType: config.StorageType(""),
			expected:    false,
		},
		{
			name:        "Random string is invalid",
			storageType: config.StorageType("random"),
			expected:    false,
		},
		{
			name:        "Invalid storage type",
			storageType: config.StorageType("invalid"),
			expected:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.storageType.IsValid()
			if result != tt.expected {
				t.Errorf("Expected IsValid() to return %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestStorageTypeConstants(t *testing.T) {
	// Verify the constant values are as expected
	if config.StorageTypeKeychain != "keychain" {
		t.Errorf("Expected StorageTypeKeychain to be 'keychain', got %q", config.StorageTypeKeychain)
	}

	if config.StorageTypeFile != "file" {
		t.Errorf("Expected StorageTypeFile to be 'file', got %q", config.StorageTypeFile)
	}
}
