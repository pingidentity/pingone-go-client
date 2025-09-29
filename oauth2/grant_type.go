// Copyright © 2025 Ping Identity Corporation

package oauth2

type GrantType string

const (
	GrantTypeAuthCode          GrantType = "auth_code"
	GrantTypeClientCredentials GrantType = "client_credentials"
	GrantTypeDeviceCode        GrantType = "device_code"
)

var AllowedTokenAuthMethods = map[GrantType][]TokenAuthType{
	GrantTypeAuthCode: {
		TokenAuthTypeAuthCode,
	},
	GrantTypeClientCredentials: {
		TokenAuthTypeClientSecretBasic,
		TokenAuthTypeClientSecretPost,
	},
	GrantTypeDeviceCode: {
		TokenAuthTypeDeviceCode,
	},
}

func IsValidGrantType(gt string) bool {
	switch GrantType(gt) {
	case GrantTypeAuthCode, GrantTypeClientCredentials, GrantTypeDeviceCode:
		return true
	}
	return false
}
