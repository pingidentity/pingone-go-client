// Copyright © 2025 Ping Identity Corporation

package oauth2

type GrantType string

const (
	// GrantTypeAuthorizationCode GrantType = "authorization_code"
	GrantTypeClientCredentials GrantType = "client_credentials"
	// GrantTypeRefreshToken      GrantType = "refresh_token"
	// GrantTypeDeviceCode        GrantType = "urn:ietf:params:oauth:grant-type:device_code"
)

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
