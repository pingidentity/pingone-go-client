// Copyright © 2025 Ping Identity Corporation

package oauth2

type GrantType string

const (
	GrantTypeAuthCode          GrantType = "authorization_code"
	GrantTypeClientCredentials GrantType = "client_credentials"
	GrantTypeDeviceCode        GrantType = "device_code"

	// GrantTypeRefreshToken represents the refresh token grant type (commented out - not yet implemented)
	// GrantTypeRefreshToken      GrantType = "refresh_token"
)

var AllowedTokenAuthMethods = map[GrantType][]TokenAuthType{
	GrantTypeAuthCode: {
		TokenAuthTypeNone,
	},
	GrantTypeClientCredentials: {
		TokenAuthTypeClientSecretBasic,
		TokenAuthTypeClientSecretPost,
	},
	GrantTypeDeviceCode: {
		TokenAuthTypeNone,
	},
	// GrantTypeRefreshToken: []TokenAuthType{
	// 	TokenAuthTypeNone,
	// 	TokenAuthTypeClientSecretBasic,
	// 	TokenAuthTypeClientSecretPost,
	// 	// TokenAuthTypeClientSecretJWT,
	// 	// TokenAuthTypePrivateKeyJWT,
	// },
}

func IsValidGrantType(gt string) bool {
	switch GrantType(gt) {
	case GrantTypeAuthCode, GrantTypeClientCredentials, GrantTypeDeviceCode:
		return true
	}
	return false
}
