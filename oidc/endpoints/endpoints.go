// Copyright © 2025 Ping Identity Corporation

// Package endpoints provides OpenID Connect (OIDC) endpoint construction utilities for PingOne services.
// It handles the creation of OIDC endpoint configurations including discovery, userinfo, and signoff
// endpoints for different PingOne deployment scenarios and regions.
package endpoints

import (
	"net/url"
	"strings"

	"github.com/pingidentity/pingone-go-client/oauth2/endpoints"
)

const (
	// OIDCDiscoveryURLPath is the URL path for OpenID Connect discovery endpoints.
	// This endpoint provides metadata about the OpenID Connect provider configuration.
	OIDCDiscoveryURLPath = "/as/.well-known/openid-configuration"
	// SignoffURLPath is the URL path for OpenID Connect signoff endpoints.
	// This endpoint handles user logout and session termination.
	SignoffURLPath = "/as/signoff"
	// UserInfoURLPath is the URL path for OpenID Connect userinfo endpoints.
	// This endpoint provides user profile information for authenticated users.
	UserInfoURLPath = "/as/userinfo"
)

// OIDCEndpoint represents a complete set of OpenID Connect endpoints for PingOne services.
// It extends the OAuth2 ExtendedEndpoint with OIDC-specific endpoints including discovery,
// userinfo, and signoff functionality.
type OIDCEndpoint struct {
	endpoints.ExtendedEndpoint
	// OIDCDiscoveryURLPath is the URL for OpenID Connect provider discovery.
	OIDCDiscoveryURLPath string
	// SignoffURLPath is the URL for OpenID Connect logout/signoff.
	SignoffURLPath string
	// UserInfoURLPath is the URL for OpenID Connect userinfo endpoint.
	UserInfoURLPath string
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
		ExtendedEndpoint:     endpoints.PingOneExtendedEndpoint(host),
		OIDCDiscoveryURLPath: u.JoinPath(OIDCDiscoveryURLPath).String(),
		SignoffURLPath:       u.JoinPath(SignoffURLPath).String(),
		UserInfoURLPath:      u.JoinPath(UserInfoURLPath).String(),
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
		ExtendedEndpoint:     endpoints.PingOneEnvironmentExtendedEndpoint(rootDomain, environmentID),
		OIDCDiscoveryURLPath: u.JoinPath(OIDCDiscoveryURLPath).String(),
		SignoffURLPath:       u.JoinPath(SignoffURLPath).String(),
		UserInfoURLPath:      u.JoinPath(UserInfoURLPath).String(),
	}
}
