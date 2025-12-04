package config

// OIDCScope represents an OIDC scope value.
// Extend this enum with additional scopes as needed.
type OIDCScope string

const (
	// OIDCScopeOpenID is the base OIDC scope required for OpenID Connect.
	OIDCScopeOpenID OIDCScope = "openid"
)

// String returns the string representation of the OIDCScope
func (s OIDCScope) String() string {
	return string(s)
}

// IsValid checks if the OIDCScope is valid
func (s OIDCScope) IsValid() bool {
	switch s {
	case OIDCScopeOpenID:
		return true
	default:
		return false
	}
}
