package oidc

import (
	"net/url"
	"strings"

	"github.com/pingidentity/pingone-go-client/oauth2"
)

const (
	OIDCDiscoveryURLPath = "/as/.well-known/openid-configuration"
	SignoffURLPath       = "/as/signoff"
	UserInfoURLPath      = "/as/userinfo"
)

type OIDCEndpoint struct {
	oauth2.ExtendedEndpoint
	OIDCDiscoveryURLPath string
	SignoffURLPath       string
	UserInfoURLPath      string
}

// PingOneOIDCEndpoint returns a new OIDCEndpoint object for the given custom domain configured on the PingOne environment.
//
// The host parameter should be the custom domain, with the subdomain (if required), but without the protocol (e.g., "bxretail.org" or "auth.bxretail.org").
func PingOneOIDCEndpoint(host, environmentID string) OIDCEndpoint {
	u := url.URL{
		Scheme: "https",
		Host:   host,
	}
	return OIDCEndpoint{
		ExtendedEndpoint:     oauth2.PingOneExtendedEndpoint(host, environmentID),
		OIDCDiscoveryURLPath: u.JoinPath(OIDCDiscoveryURLPath).String(),
		SignoffURLPath:       u.JoinPath(SignoffURLPath).String(),
		UserInfoURLPath:      u.JoinPath(UserInfoURLPath).String(),
	}
}

// PingOneEnvironmentOIDCEndpoint returns a new OIDCEndpoint object when not using a custom domain in a PingOne environment.
// The endpoints are constructed using the top level domain and environment ID, which result in the URL format:
// https://auth.pingone.<topLevelDomain>/<environmentID>/<endpoint path>.
//
// The topLevelDomain parameter should be a valid PingOne top level domain that applies to the region of the PingOne tenant (e.g., "com" or "eu").
// Top level domains with leading `.` are also supported (e.g., ".com" or ".eu").
//
// The environmentID parameter should be a valid PingOne environment ID.
func PingOneEnvironmentOIDCEndpoint(topLevelDomain, environmentID string) OIDCEndpoint {
	u := url.URL{
		Scheme: "https",
		Host:   "auth.pingone." + strings.TrimPrefix(topLevelDomain, "."),
	}
	if environmentID == "" {
		panic("endpoints: invalid environment ID")
	}
	return OIDCEndpoint{
		ExtendedEndpoint:     oauth2.PingOneExtendedEndpoint(topLevelDomain, environmentID),
		OIDCDiscoveryURLPath: u.JoinPath(OIDCDiscoveryURLPath).String(),
		SignoffURLPath:       u.JoinPath(SignoffURLPath).String(),
		UserInfoURLPath:      u.JoinPath(UserInfoURLPath).String(),
	}
}
