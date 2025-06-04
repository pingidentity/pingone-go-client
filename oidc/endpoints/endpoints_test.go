// Copyright © 2025 Ping Identity Corporation

package endpoints_test

import (
	"testing"

	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"github.com/stretchr/testify/assert"
)

func TestPingOneOIDCEndpoint(t *testing.T) {
	host := "auth.bxretail.org"
	endpoint := endpoints.PingOneOIDCEndpoint(host)

	// Base OAuth2 endpoints
	assert.Equal(t, "https://auth.bxretail.org/as/authorize", endpoint.AuthURL)
	assert.Equal(t, "https://auth.bxretail.org/as/token", endpoint.TokenURL)
	assert.Equal(t, "https://auth.bxretail.org/as/device_authorization", endpoint.DeviceAuthURL)

	// Extended OAuth2 endpoints
	assert.Equal(t, "https://auth.bxretail.org/as/introspect", endpoint.IntrospectionURL)
	assert.Equal(t, "https://auth.bxretail.org/as", endpoint.IssuerURLPath)
	assert.Equal(t, "https://auth.bxretail.org/as/jwks", endpoint.JWKSURL)
	assert.Equal(t, "https://auth.bxretail.org/as/par", endpoint.PARURL)
	assert.Equal(t, "https://auth.bxretail.org/as/revoke", endpoint.TokenRevocationURL)

	// OIDC specific endpoints
	assert.Equal(t, "https://auth.bxretail.org/as/.well-known/openid-configuration", endpoint.OIDCDiscoveryURLPath)
	assert.Equal(t, "https://auth.bxretail.org/as/signoff", endpoint.SignoffURLPath)
	assert.Equal(t, "https://auth.bxretail.org/as/userinfo", endpoint.UserInfoURLPath)
}

func TestPingOneEnvironmentOIDCEndpoint(t *testing.T) {
	tests := []struct {
		name          string
		rootDomain    string
		environmentID string
		expectPanic   bool
	}{
		{
			name:          "Valid root domain and environment ID",
			rootDomain:    "pingone.com",
			environmentID: "env-id",
			expectPanic:   false,
		},
		{
			name:          "Root domain with leading dot",
			rootDomain:    ".pingone.com",
			environmentID: "env-id",
			expectPanic:   false,
		},
		{
			name:          "Empty environment ID",
			rootDomain:    "pingone.com",
			environmentID: "",
			expectPanic:   true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectPanic {
				assert.Panics(t, func() {
					endpoints.PingOneEnvironmentOIDCEndpoint(tc.rootDomain, tc.environmentID)
				})
			} else {
				endpoint := endpoints.PingOneEnvironmentOIDCEndpoint(tc.rootDomain, tc.environmentID)

				expectedHost := "auth." + trimPrefix(tc.rootDomain, ".")
				basePath := "https://" + expectedHost + "/" + tc.environmentID

				// Base OAuth2 endpoints
				assert.Equal(t, basePath+"/as/authorize", endpoint.AuthURL)
				assert.Equal(t, basePath+"/as/token", endpoint.TokenURL)
				assert.Equal(t, basePath+"/as/device_authorization", endpoint.DeviceAuthURL)

				// Extended OAuth2 endpoints
				assert.Equal(t, basePath+"/as/introspect", endpoint.IntrospectionURL)
				assert.Equal(t, basePath+"/as", endpoint.IssuerURLPath)
				assert.Equal(t, basePath+"/as/jwks", endpoint.JWKSURL)
				assert.Equal(t, basePath+"/as/par", endpoint.PARURL)
				assert.Equal(t, basePath+"/as/revoke", endpoint.TokenRevocationURL)

				// OIDC specific endpoints
				assert.Equal(t, basePath+"/as/.well-known/openid-configuration", endpoint.OIDCDiscoveryURLPath)
				assert.Equal(t, basePath+"/as/signoff", endpoint.SignoffURLPath)
				assert.Equal(t, basePath+"/as/userinfo", endpoint.UserInfoURLPath)
			}
		})
	}
}

func TestConstantPaths(t *testing.T) {
	assert.Equal(t, "/as/.well-known/openid-configuration", endpoints.OIDCDiscoveryURLPath)
	assert.Equal(t, "/as/signoff", endpoints.SignoffURLPath)
	assert.Equal(t, "/as/userinfo", endpoints.UserInfoURLPath)
}

// Helper function that replicates strings.TrimPrefix since we don't want to import strings directly
func trimPrefix(s, prefix string) string {
	if len(s) >= len(prefix) && s[0:len(prefix)] == prefix {
		return s[len(prefix):]
	}
	return s
}
