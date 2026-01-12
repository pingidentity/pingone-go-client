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
			name:        "StorageTypeSecureLocal",
			storageType: config.StorageTypeSecureLocal,
			expected:    "secure_local",
		},
		{
			name:        "StorageTypeFileSystem",
			storageType: config.StorageTypeFileSystem,
			expected:    "file_system",
		},
		{
			name:        "StorageTypeSecureRemote",
			storageType: config.StorageTypeSecureRemote,
			expected:    "secure_remote",
		},
		{
			name:        "StorageTypeNone",
			storageType: config.StorageTypeNone,
			expected:    "none",
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
			name:        "StorageTypeSecureLocal is valid",
			storageType: config.StorageTypeSecureLocal,
			expected:    true,
		},
		{
			name:        "StorageTypeFileSystem is valid",
			storageType: config.StorageTypeFileSystem,
			expected:    true,
		},
		{
			name:        "StorageTypeSecureRemote is valid",
			storageType: config.StorageTypeSecureRemote,
			expected:    true,
		},
		{
			name:        "StorageTypeNone is valid",
			storageType: config.StorageTypeNone,
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
	if config.StorageTypeSecureLocal != "secure_local" {
		t.Errorf("Expected StorageTypeSecureLocal to be 'secure_local', got %q", config.StorageTypeSecureLocal)
	}
	if config.StorageTypeFileSystem != "file_system" {
		t.Errorf("Expected StorageTypeFileSystem to be 'file_system', got %q", config.StorageTypeFileSystem)
	}
	if config.StorageTypeSecureRemote != "secure_remote" {
		t.Errorf("Expected StorageTypeSecureRemote to be 'secure_remote', got %q", config.StorageTypeSecureRemote)
	}
	if config.StorageTypeNone != "none" {
		t.Errorf("Expected StorageTypeNone to be 'none', got %q", config.StorageTypeNone)
	}
}
