package config

import "testing"

func TestOIDCScopeString(t *testing.T) {
	if got := OIDCScopeOpenID.String(); got != "openid" {
		t.Fatalf("OIDCScopeOpenID.String() = %q, want %q", got, "openid")
	}
}

func TestOIDCScopeIsValid_OpenID(t *testing.T) {
	if !OIDCScopeOpenID.IsValid() {
		t.Fatalf("OIDCScopeOpenID.IsValid() = false, want true")
	}
}

func TestOIDCScopeIsValid_Invalid(t *testing.T) {
	invalid := OIDCScope("invalid")
	if invalid.IsValid() {
		t.Fatalf("invalid.IsValid() = true, want false")
	}
}
