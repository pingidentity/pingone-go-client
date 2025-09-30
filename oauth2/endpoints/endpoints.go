// Copyright © 2025 Ping Identity Corporation

// Package endpoints provides OAuth2 endpoint construction utilities for PingOne services.
// It handles the creation of OAuth2 endpoint configurations for different PingOne deployment
// scenarios including custom domains and environment-specific endpoints across regions.
package endpoints

import (
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

const (
	// AuthURLPath is the URL path for OAuth2 authorization endpoints.
	AuthURLPath = "/as/authorize"
	// DeviceAuthURLPath is the URL path for OAuth2 device authorization endpoints.
	DeviceAuthURLPath = "/as/device_authorization"
	// IntrospectionURLPath is the URL path for OAuth2 token introspection endpoints.
	IntrospectionURLPath = "/as/introspect"
	// IssuerURLPath is the URL path for OAuth2 issuer endpoints.
	IssuerURLPath = "/as"
	// JWKSURLPath is the URL path for JSON Web Key Set endpoints.
	JWKSURLPath = "/as/jwks"
	// PARURLPath is the URL path for Pushed Authorization Request endpoints.
	PARURLPath = "/as/par"
	// TokenRevocationURLPath is the URL path for OAuth2 token revocation endpoints.
	TokenRevocationURLPath = "/as/revoke"
	// TokenURLPath is the URL path for OAuth2 token endpoints.
	TokenURLPath = "/as/token"
)

// PingOneEndpoint returns a new oauth2.Endpoint object for the given custom domain configured on the PingOne environment.
//
// The host parameter should be the custom domain, with the subdomain (if required), but without the protocol (e.g., "bxretail.org" or "auth.bxretail.org").
func PingOneEndpoint(host string) oauth2.Endpoint {
	u := url.URL{
		Scheme: "https",
		Host:   host,
	}
	return oauth2.Endpoint{
		AuthURL:       u.JoinPath(AuthURLPath).String(),
		TokenURL:      u.JoinPath(TokenURLPath).String(),
		DeviceAuthURL: u.JoinPath(DeviceAuthURLPath).String(),
	}
}

// PingOneEnvironmentEndpoint returns a new oauth2.Endpoint object when not using a custom domain in a PingOne environment.
// The endpoints are constructed using the root domain and environment ID, which result in the URL format:
// https://auth.<rootDomain>/<environmentID>/<endpoint path>.
//
// The rootDomain parameter should be a valid PingOne top level domain that applies to the region of the PingOne tenant (e.g., "pingone.com" or "pingone.eu").
// Root domains with leading `.` are also supported (e.g., ".pingone.com" or ".pingone.eu").
//
// The environmentID parameter should be a valid PingOne environment ID.
func PingOneEnvironmentEndpoint(rootDomain, environmentID string) oauth2.Endpoint {
	u := url.URL{
		Scheme: "https",
		Host:   "auth." + strings.TrimPrefix(rootDomain, "."),
		Path:   environmentID,
	}
	if environmentID == "" {
		panic("endpoints: invalid environment ID")
	}
	return oauth2.Endpoint{
		AuthURL:       u.JoinPath(AuthURLPath).String(),
		TokenURL:      u.JoinPath(TokenURLPath).String(),
		DeviceAuthURL: u.JoinPath(DeviceAuthURLPath).String(),
	}
}

// ExtendedEndpoint represents an OAuth2 endpoint with additional PingOne-specific endpoints.
// It extends the standard OAuth2 endpoint with additional URLs for token introspection,
// JWKS, PAR (Pushed Authorization Requests), and token revocation functionality.
type ExtendedEndpoint struct {
	oauth2.Endpoint
	// IntrospectionURL is the endpoint URL for OAuth2 token introspection.
	IntrospectionURL string
	// IssuerURLPath is the issuer URL path for the OAuth2 authorization server.
	IssuerURLPath string
	// JWKSURL is the endpoint URL for retrieving JSON Web Key Sets.
	JWKSURL string
	// PARURL is the endpoint URL for Pushed Authorization Requests.
	PARURL string
	// TokenRevocationURL is the endpoint URL for OAuth2 token revocation.
	TokenRevocationURL string
}

// PingOneExtendedEndpoint returns a new ExtendedEndpoint object for the given custom domain configured on the PingOne environment.
//
// The host parameter should be the custom domain, with the subdomain (if required), but without the protocol (e.g., "bxretail.org" or "auth.bxretail.org").
func PingOneExtendedEndpoint(host string) ExtendedEndpoint {
	u := url.URL{
		Scheme: "https",
		Host:   host,
	}
	return ExtendedEndpoint{
		Endpoint:           PingOneEndpoint(host),
		IntrospectionURL:   u.JoinPath(IntrospectionURLPath).String(),
		IssuerURLPath:      u.JoinPath(IssuerURLPath).String(),
		JWKSURL:            u.JoinPath(JWKSURLPath).String(),
		PARURL:             u.JoinPath(PARURLPath).String(),
		TokenRevocationURL: u.JoinPath(TokenRevocationURLPath).String(),
	}
}

// PingOneEnvironmentExtendedEndpoint returns a new ExtendedEndpoint object when not using a custom domain in a PingOne environment.
// The endpoints are constructed using the root domain and environment ID, which result in the URL format:
// https://auth.<rootDomain>/<environmentID>/<endpoint path>.
//
// The rootDomain parameter should be a valid PingOne root domain that applies to the region of the PingOne tenant (e.g., "pingone.com" or "pingone.eu").
// Top level domains with leading `.` are also supported (e.g., ".pingone.com" or ".pingone.eu").
//
// The environmentID parameter should be a valid PingOne environment ID.
func PingOneEnvironmentExtendedEndpoint(rootDomain, environmentID string) ExtendedEndpoint {
	u := url.URL{
		Scheme: "https",
		Host:   "auth." + strings.TrimPrefix(rootDomain, "."),
		Path:   environmentID,
	}
	if environmentID == "" {
		panic("endpoints: invalid environment ID")
	}
	return ExtendedEndpoint{
		Endpoint:           PingOneEnvironmentEndpoint(rootDomain, environmentID),
		IntrospectionURL:   u.JoinPath(IntrospectionURLPath).String(),
		IssuerURLPath:      u.JoinPath(IssuerURLPath).String(),
		JWKSURL:            u.JoinPath(JWKSURLPath).String(),
		PARURL:             u.JoinPath(PARURLPath).String(),
		TokenRevocationURL: u.JoinPath(TokenRevocationURLPath).String(),
	}
}
