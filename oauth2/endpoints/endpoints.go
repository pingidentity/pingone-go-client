// Copyright © 2025 Ping Identity Corporation

package endpoints

import (
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

const (
	AuthURLPath            = "/as/authorize"
	DeviceAuthURLPath      = "/as/device_authorization"
	IntrospectionURLPath   = "/as/introspect"
	IssuerURLPath          = "/as"
	JWKSURLPath            = "/as/jwks"
	PARURLPath             = "/as/par"
	TokenRevocationURLPath = "/as/revoke"
	TokenURLPath           = "/as/token"
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

type ExtendedEndpoint struct {
	oauth2.Endpoint
	IntrospectionURL   string
	IssuerURLPath      string
	JWKSURL            string
	PARURL             string
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
