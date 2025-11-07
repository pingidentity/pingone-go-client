// Copyright © 2025 Ping Identity Corporation

package endpoints_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pingidentity/pingone-go-client/oauth2/endpoints"
)

func TestPingOneEndpoint(t *testing.T) {
	host := "auth.bxretail.org"
	endpoint := endpoints.PingOneEndpoint(host)

	assert.Equal(t, "https://auth.bxretail.org/as/authorize", endpoint.AuthURL)
	assert.Equal(t, "https://auth.bxretail.org/as/token", endpoint.TokenURL)
	assert.Equal(t, "https://auth.bxretail.org/as/device_authorization", endpoint.DeviceAuthURL)
}

func TestPingOneEnvironmentEndpoint(t *testing.T) {
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
					endpoints.PingOneEnvironmentEndpoint(tc.rootDomain, tc.environmentID)
				})
			} else {
				endpoint := endpoints.PingOneEnvironmentEndpoint(tc.rootDomain, tc.environmentID)

				expectedHost := "auth." + trimPrefix(tc.rootDomain, ".")
				assert.Equal(t, "https://"+expectedHost+"/"+tc.environmentID+"/as/authorize", endpoint.AuthURL)
				assert.Equal(t, "https://"+expectedHost+"/"+tc.environmentID+"/as/token", endpoint.TokenURL)
				assert.Equal(t, "https://"+expectedHost+"/"+tc.environmentID+"/as/device_authorization", endpoint.DeviceAuthURL)
			}
		})
	}
}

func TestPingOneExtendedEndpoint(t *testing.T) {
	host := "auth.bxretail.org"
	endpoint := endpoints.PingOneExtendedEndpoint(host)

	// Base OAuth2 endpoints
	assert.Equal(t, "https://auth.bxretail.org/as/authorize", endpoint.AuthURL)
	assert.Equal(t, "https://auth.bxretail.org/as/token", endpoint.TokenURL)
	assert.Equal(t, "https://auth.bxretail.org/as/device_authorization", endpoint.DeviceAuthURL)

	// Extended endpoints
	assert.Equal(t, "https://auth.bxretail.org/as/introspect", endpoint.IntrospectionURL)
	assert.Equal(t, "https://auth.bxretail.org/as", endpoint.IssuerURLPath)
	assert.Equal(t, "https://auth.bxretail.org/as/jwks", endpoint.JWKSURL)
	assert.Equal(t, "https://auth.bxretail.org/as/par", endpoint.PARURL)
	assert.Equal(t, "https://auth.bxretail.org/as/revoke", endpoint.TokenRevocationURL)
}

func TestPingOneEnvironmentExtendedEndpoint(t *testing.T) {
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
					endpoints.PingOneEnvironmentExtendedEndpoint(tc.rootDomain, tc.environmentID)
				})
			} else {
				endpoint := endpoints.PingOneEnvironmentExtendedEndpoint(tc.rootDomain, tc.environmentID)

				expectedHost := "auth." + trimPrefix(tc.rootDomain, ".")
				basePath := "https://" + expectedHost + "/" + tc.environmentID

				// Base OAuth2 endpoints
				assert.Equal(t, basePath+"/as/authorize", endpoint.AuthURL)
				assert.Equal(t, basePath+"/as/token", endpoint.TokenURL)
				assert.Equal(t, basePath+"/as/device_authorization", endpoint.DeviceAuthURL)

				// Extended endpoints
				assert.Equal(t, basePath+"/as/introspect", endpoint.IntrospectionURL)
				assert.Equal(t, basePath+"/as", endpoint.IssuerURLPath)
				assert.Equal(t, basePath+"/as/jwks", endpoint.JWKSURL)
				assert.Equal(t, basePath+"/as/par", endpoint.PARURL)
				assert.Equal(t, basePath+"/as/revoke", endpoint.TokenRevocationURL)
			}
		})
	}
}

func TestConstantPaths(t *testing.T) {
	assert.Equal(t, "/as/authorize", endpoints.AuthURLPath)
	assert.Equal(t, "/as/device_authorization", endpoints.DeviceAuthURLPath)
	assert.Equal(t, "/as/introspect", endpoints.IntrospectionURLPath)
	assert.Equal(t, "/as", endpoints.IssuerURLPath)
	assert.Equal(t, "/as/jwks", endpoints.JWKSURLPath)
	assert.Equal(t, "/as/par", endpoints.PARURLPath)
	assert.Equal(t, "/as/revoke", endpoints.TokenRevocationURLPath)
	assert.Equal(t, "/as/token", endpoints.TokenURLPath)
}

// Helper function that replicates strings.TrimPrefix since we don't want to import strings directly
func trimPrefix(s, prefix string) string {
	if len(s) >= len(prefix) && s[0:len(prefix)] == prefix {
		return s[len(prefix):]
	}
	return s
}
