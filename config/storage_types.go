// Copyright © 2025 Ping Identity Corporation

package config

// StorageType represents the type of storage backend for tokens
type StorageType string

const (
	// StorageTypeKeychain uses the system keychain for token storage
	StorageTypeKeychain StorageType = "keychain"

	// StorageTypeFile uses file system storage for tokens
	StorageTypeFile StorageType = "file"
)

// String returns the string representation of the StorageType
func (s StorageType) String() string {
	return string(s)
}

// IsValid checks if the StorageType is valid
func (s StorageType) IsValid() bool {
	switch s {
	case StorageTypeKeychain, StorageTypeFile:
		return true
	default:
		return false
	}
}
