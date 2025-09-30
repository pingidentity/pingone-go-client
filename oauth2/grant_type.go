// Copyright © 2025 Ping Identity Corporation

// Package oauth2 provides OAuth2 authentication utilities for the PingOne Go Client SDK.
// It includes grant type definitions, token authentication methods, and endpoint configuration
// for OAuth2 flows supported by PingOne services.
package oauth2

// GrantType represents the OAuth2 grant type used for token acquisition.
// Grant types define the method by which applications obtain access tokens from the authorization server.
type GrantType string

const (
	// GrantTypeAuthorizationCode represents the authorization code grant type (commented out - not yet implemented)
	// GrantTypeAuthorizationCode GrantType = "authorization_code"

	// GrantTypeClientCredentials represents the client credentials grant type.
	// This grant type is used for server-to-server authentication where the client
	// authenticates directly with the authorization server using its client credentials.
	GrantTypeClientCredentials GrantType = "client_credentials"

	// GrantTypeRefreshToken represents the refresh token grant type (commented out - not yet implemented)
	// GrantTypeRefreshToken      GrantType = "refresh_token"

	// GrantTypeDeviceCode represents the device code grant type (commented out - not yet implemented)
	// GrantTypeDeviceCode        GrantType = "urn:ietf:params:oauth:grant-type:device_code"
)

// AllowedTokenAuthMethods maps each grant type to its supported token authentication methods.
// This mapping ensures that only compatible authentication methods are used with each grant type.
// The map helps validate authentication configurations and provides available options for each flow.
var AllowedTokenAuthMethods = map[GrantType][]TokenAuthType{
	// GrantTypeAuthorizationCode: []TokenAuthType{
	// 	TokenAuthTypeNone,
	// 	TokenAuthTypeClientSecretBasic,
	// 	TokenAuthTypeClientSecretPost,
	// 	// TokenAuthTypeClientSecretJWT,
	// 	// TokenAuthTypePrivateKeyJWT,
	// },
	GrantTypeClientCredentials: {
		TokenAuthTypeClientSecretBasic,
		TokenAuthTypeClientSecretPost,
		// TokenAuthTypeClientSecretJWT,
		// TokenAuthTypePrivateKeyJWT,
	},
	// GrantTypeRefreshToken: []TokenAuthType{
	// 	TokenAuthTypeNone,
	// 	TokenAuthTypeClientSecretBasic,
	// 	TokenAuthTypeClientSecretPost,
	// 	// TokenAuthTypeClientSecretJWT,
	// 	// TokenAuthTypePrivateKeyJWT,
	// },
	// GrantTypeDeviceCode:   []TokenAuthType{},
}
