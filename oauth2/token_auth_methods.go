// Copyright © 2025 Ping Identity Corporation

package oauth2

type TokenAuthType string

const (
	TokenAuthTypeNone              TokenAuthType = "NONE"
	TokenAuthTypeClientSecretBasic TokenAuthType = "CLIENT_SECRET_BASIC"
	TokenAuthTypeClientSecretPost  TokenAuthType = "CLIENT_SECRET_POST"
	// TokenAuthTypeClientSecretJWT  TokenAuthType = "CLIENT_SECRET_JWT"
	// TokenAuthTypePrivateKeyJWT  TokenAuthType = "PRIVATE_KEY_JWT"
)
