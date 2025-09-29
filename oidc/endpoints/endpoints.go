// Copyright © 2025 Ping Identity Corporation

package endpoints

import (
	"net/url"
	"strings"

	"github.com/pingidentity/pingone-go-client/oauth2/endpoints"
)

const (
	AuthorizationURLPath       = "/as/authorize"
	DeviceAuthorizationURLPath = "/as/device_authorization"
	OIDCDiscoveryURLPath       = "/as/.well-known/openid-configuration"
	SignoffURLPath             = "/as/signoff"
	TokenURLPath               = "/as/token"
	UserInfoURLPath            = "/as/userinfo"
)

type OIDCEndpoint struct {
	endpoints.ExtendedEndpoint
	AuthorizationURLPath       string
	DeviceAuthorizationURLPath string
	OIDCDiscoveryURLPath       string
	SignoffURLPath             string
	TokenURLPath               string
	UserInfoURLPath            string
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
		AuthorizationURLPath:       u.JoinPath(AuthorizationURLPath).String(),
		DeviceAuthorizationURLPath: u.JoinPath(DeviceAuthorizationURLPath).String(),
		ExtendedEndpoint:           endpoints.PingOneExtendedEndpoint(host),
		OIDCDiscoveryURLPath:       u.JoinPath(OIDCDiscoveryURLPath).String(),
		SignoffURLPath:             u.JoinPath(SignoffURLPath).String(),
		TokenURLPath:               u.JoinPath(TokenURLPath).String(),
		UserInfoURLPath:            u.JoinPath(UserInfoURLPath).String(),
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
		AuthorizationURLPath:       u.JoinPath(AuthorizationURLPath).String(),
		DeviceAuthorizationURLPath: u.JoinPath(DeviceAuthorizationURLPath).String(),
		ExtendedEndpoint:           endpoints.PingOneEnvironmentExtendedEndpoint(rootDomain, environmentID),
		OIDCDiscoveryURLPath:       u.JoinPath(OIDCDiscoveryURLPath).String(),
		SignoffURLPath:             u.JoinPath(SignoffURLPath).String(),
		TokenURLPath:               u.JoinPath(TokenURLPath).String(),
		UserInfoURLPath:            u.JoinPath(UserInfoURLPath).String(),
	}
}
