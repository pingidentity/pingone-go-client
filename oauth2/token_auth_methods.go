package oauth2

type TokenAuthType string

const (
	TokenAuthTypeAuthCode          TokenAuthType = "AUTH_CODE"
	TokenAuthTypeClientSecretBasic TokenAuthType = "CLIENT_SECRET_BASIC"
	TokenAuthTypeClientSecretPost  TokenAuthType = "CLIENT_SECRET_POST"
	TokenAuthTypeDeviceCode        TokenAuthType = "DEVICE_CODE"
	TokenAuthTypeNone              TokenAuthType = "NONE"
)
