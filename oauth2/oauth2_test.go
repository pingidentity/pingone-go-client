// Copyright © 2025 Ping Identity Corporation

package oauth2_test

import (
	"testing"

	"github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/stretchr/testify/assert"
)

func TestGrantTypeConstants(t *testing.T) {
	assert.Equal(t, oauth2.GrantType("client_credentials"), oauth2.GrantTypeClientCredentials)
}

func TestTokenAuthTypeConstants(t *testing.T) {
	assert.Equal(t, oauth2.TokenAuthType("NONE"), oauth2.TokenAuthTypeNone)
	assert.Equal(t, oauth2.TokenAuthType("CLIENT_SECRET_BASIC"), oauth2.TokenAuthTypeClientSecretBasic)
	assert.Equal(t, oauth2.TokenAuthType("CLIENT_SECRET_POST"), oauth2.TokenAuthTypeClientSecretPost)
}

func TestAllowedTokenAuthMethods(t *testing.T) {
	// Test that client credentials grant type has the expected auth methods
	clientCredentialsAuthMethods, exists := oauth2.AllowedTokenAuthMethods[oauth2.GrantTypeClientCredentials]
	assert.True(t, exists, "Client credentials grant type should be in the map")

	// Verify that client credentials has the correct auth methods
	assert.Contains(t, clientCredentialsAuthMethods, oauth2.TokenAuthTypeClientSecretBasic)
	assert.Contains(t, clientCredentialsAuthMethods, oauth2.TokenAuthTypeClientSecretPost)
	assert.Len(t, clientCredentialsAuthMethods, 2, "Client credentials should have exactly 2 auth methods")
}
