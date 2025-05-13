package pingone

type Configuration struct {
	Auth struct {
		ClientID     *string `envconfig:"PINGONE_CLIENT_ID" json:"clientId,omitempty"`
		ClientSecret *string `envconfig:"PINGONE_CLIENT_SECRET" json:"clientSecret,omitempty"`
		AccessToken  *string `envconfig:"PINGONE_API_ACCESS_TOKEN" json:"accessToken,omitempty"`
	} `json:"auth"`
	Endpoint struct {
		AuthEnvironmentID *string `envconfig:"PINGONE_ENVIRONMENT_ID" json:"environmentId,omitempty"`
		TopLevelDomain    *string `envconfig:"PINGONE_TOP_LEVEL_DOMAIN" json:"topLevelDomain,omitempty"`
		RootDomain        *string `envconfig:"PINGONE_ROOT_DOMAIN" json:"rootDomain,omitempty"`
		APIDomain         *string `envconfig:"PINGONE_API_DOMAIN" json:"apiDomain,omitempty"`
		CustomDomain      *string `envconfig:"PINGONE_CUSTOM_DOMAIN" json:"customDomain,omitempty"`
	} `json:"endpoint"`
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func (c *Configuration) WithAuthEnvironmentID(environmentID string) *Configuration {
	c.Endpoint.AuthEnvironmentID = &environmentID
	return c
}

func (c *Configuration) WithClientID(clientID string) *Configuration {
	c.Auth.ClientID = &clientID
	return c
}

func (c *Configuration) WithClientSecret(clientSecret string) *Configuration {
	c.Auth.ClientSecret = &clientSecret
	return c
}

func (c *Configuration) WithAccessToken(accessToken string) *Configuration {
	c.Auth.AccessToken = &accessToken
	return c
}

func (c *Configuration) WithTopLevelDomain(tld string) *Configuration {
	c.Endpoint.TopLevelDomain = &tld
	return c
}

func (c *Configuration) WithRootDomain(rootDomain string) *Configuration {
	c.Endpoint.RootDomain = &rootDomain
	return c
}

func (c *Configuration) WithAPIDomain(apiDomain string) *Configuration {
	c.Endpoint.APIDomain = &apiDomain
	return c
}

func (c *Configuration) WithCustomDomain(customDomain string) *Configuration {
	c.Endpoint.CustomDomain = &customDomain
	return c
}

func (c *Configuration) GetBearerToken() *string {
	return c.Auth.AccessToken
}
