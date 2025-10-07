package oauth2

type TokenAuthType string

const (
	TokenAuthTypeClientSecretBasic TokenAuthType = "CLIENT_SECRET_BASIC"
	TokenAuthTypeClientSecretPost  TokenAuthType = "CLIENT_SECRET_POST"
	TokenAuthTypeNone              TokenAuthType = "NONE"
)
