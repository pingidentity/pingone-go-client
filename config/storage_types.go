// Copyright © 2025 Ping Identity Corporation

package config

// StorageType represents the type of storage backend for tokens
type StorageType string

const (
	// StorageTypeKeychain uses the system keychain for token storage
	StorageTypeKeychain StorageType = "keychain"

	// StorageTypeNone defines no token storage
	StorageTypeNone StorageType = "none"
)

// String returns the string representation of the StorageType
func (s StorageType) String() string {
	return string(s)
}

// IsValid checks if the StorageType is valid
func (s StorageType) IsValid() bool {
	switch s {
	case StorageTypeKeychain:
		return true
	default:
		return false
	}
}
