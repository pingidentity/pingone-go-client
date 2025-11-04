// Copyright © 2025 Ping Identity Corporation

package endpoints

import (
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

const (
	AuthorizationURLPath       = "/as/authorize"
	DeviceAuthorizationURLPath = "/as/device_authorization"
	IntrospectionURLPath       = "/as/introspect"
	IssuerURLPath              = "/as"
	JWKSURLPath                = "/as/jwks"
	OIDCDiscoveryURLPath       = "/as/.well-known/openid-configuration"
	PARURLPath                 = "/as/par"
	SignoffURLPath             = "/as/signoff"
	TokenRevocationURLPath     = "/as/revoke"
	TokenURLPath               = "/as/token"
	UserInfoURLPath            = "/as/userinfo"
)

type OIDCEndpoint struct {
	oauth2.Endpoint
	IntrospectionURL   string
	IssuerURLPath      string
	JWKSURL            string
	PARURL             string
	TokenRevocationURL string
	// OIDC-specific fields
	OIDCDiscoveryURL string
	SignoffURL       string
	UserInfoURL      string
}

// PingOneOIDCEndpoint returns a new OIDCEndpoint object for the given custom domain configured on the PingOne environment.
//
// The host parameter should be the custom domain, with the subdomain (if required), but without the protocol (e.g., "bxretail.org" or "auth.bxretail.org").
func PingOneOIDCEndpoint(host string) OIDCEndpoint {
	u := url.URL{
		Scheme: "https",
		Host:   host,
	}
	return OIDCEndpoint{
		Endpoint: oauth2.Endpoint{
			AuthURL:       u.JoinPath(AuthorizationURLPath).String(),
			TokenURL:      u.JoinPath(TokenURLPath).String(),
			DeviceAuthURL: u.JoinPath(DeviceAuthorizationURLPath).String(),
		},
		IntrospectionURL:   u.JoinPath(IntrospectionURLPath).String(),
		IssuerURLPath:      u.JoinPath(IssuerURLPath).String(),
		JWKSURL:            u.JoinPath(JWKSURLPath).String(),
		PARURL:             u.JoinPath(PARURLPath).String(),
		TokenRevocationURL: u.JoinPath(TokenRevocationURLPath).String(),
		OIDCDiscoveryURL:   u.JoinPath(OIDCDiscoveryURLPath).String(),
		SignoffURL:         u.JoinPath(SignoffURLPath).String(),
		UserInfoURL:        u.JoinPath(UserInfoURLPath).String(),
	}
}

// PingOneEnvironmentOIDCEndpoint returns a new OIDCEndpoint object when not using a custom domain in a PingOne environment.
// The endpoints are constructed using the root domain and environment ID, which result in the URL format:
// https://auth.<rootDomain>/<environmentID>/<endpoint path>.
//
// The rootDomain parameter should be a valid PingOne root domain that applies to the region of the PingOne tenant (e.g., "pingone.com" or "pingone.eu").
// Root domains with leading `.` are also supported (e.g., ".pingone.com" or ".pingone.eu").
//
// The environmentID parameter should be a valid PingOne environment ID.
func PingOneEnvironmentOIDCEndpoint(rootDomain, environmentID string) OIDCEndpoint {
	u := url.URL{
		Scheme: "https",
		Host:   "auth." + strings.TrimPrefix(rootDomain, "."),
		Path:   environmentID,
	}
	if environmentID == "" {
		panic("endpoints: invalid environment ID")
	}
	return OIDCEndpoint{
		Endpoint: oauth2.Endpoint{
			AuthURL:       u.JoinPath(AuthorizationURLPath).String(),
			TokenURL:      u.JoinPath(TokenURLPath).String(),
			DeviceAuthURL: u.JoinPath(DeviceAuthorizationURLPath).String(),
		},
		IntrospectionURL:   u.JoinPath(IntrospectionURLPath).String(),
		IssuerURLPath:      u.JoinPath(IssuerURLPath).String(),
		JWKSURL:            u.JoinPath(JWKSURLPath).String(),
		PARURL:             u.JoinPath(PARURLPath).String(),
		TokenRevocationURL: u.JoinPath(TokenRevocationURLPath).String(),
		OIDCDiscoveryURL:   u.JoinPath(OIDCDiscoveryURLPath).String(),
		SignoffURL:         u.JoinPath(SignoffURLPath).String(),
		UserInfoURL:        u.JoinPath(UserInfoURLPath).String(),
	}
}
