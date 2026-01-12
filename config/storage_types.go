// Copyright © 2025 Ping Identity Corporation

package config

// StorageType represents the type of storage backend for tokens
type StorageType string

const (
	// StorageTypeFileSystem uses the local file system for token storage
	StorageTypeFileSystem StorageType = "file_system"

	// StorageTypeSecureLocal uses a secure local storage (e.g., OS keychain)
	StorageTypeSecureLocal StorageType = "secure_local"

	// StorageTypeSecureRemote uses a secure remote storage service
	StorageTypeSecureRemote StorageType = "secure_remote"

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
	case StorageTypeFileSystem:
		return true
	case StorageTypeSecureLocal:
		return true
	case StorageTypeSecureRemote:
		return true
	case StorageTypeNone:
		return true
	default:
		return false
	}
}
