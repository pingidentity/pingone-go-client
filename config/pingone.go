// Copyright © 2025 Ping Identity Corporation

package config

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"golang.org/x/oauth2"
)

type AuthCode struct {
	AuthCodeClientID      *string   `envconfig:"PINGONE_AUTH_CODE_CLIENT_ID" json:"authCodeClientId,omitempty"`
	AuthCodeEnvironmentID *string   `envconfig:"PINGONE_AUTH_CODE_ENVIRONMENT_ID" json:"authCodeEnvironmentId,omitempty"`
	AuthCodePort          *string   `envconfig:"PINGONE_AUTH_CODE_PORT" json:"authCodePort,omitempty"`
	AuthCodeScopes        *[]string `envconfig:"PINGONE_AUTH_CODE_SCOPES" json:"authCodeScopes,omitempty"`
}

type ClientCredentials struct {
	ClientCredentialsClientID      *string `envconfig:"PINGONE_CLIENT_CREDENTIALS_CLIENT_ID" json:"clientCredentialsClientId,omitempty"`
	ClientCredentialsClientSecret  *string `envconfig:"PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET" json:"clientCredentialsClientSecret,omitempty"`
	ClientCredentialsEnvironmentID *string `envconfig:"PINGONE_CLIENT_CREDENTIALS_ENVIRONMENT_ID" json:"clientCredentialsEnvironmentId,omitempty"`
}

type DeviceCode struct {
	DeviceCodeClientID      *string   `envconfig:"PINGONE_DEVICE_CODE_CLIENT_ID" json:"deviceCodeClientId,omitempty"`
	DeviceCodeEnvironmentID *string   `envconfig:"PINGONE_DEVICE_CODE_ENVIRONMENT_ID" json:"deviceCodeEnvironmentId,omitempty"`
	DeviceCodeScopes        *[]string `envconfig:"PINGONE_DEVICE_CODE_SCOPES" json:"deviceCodeScopes,omitempty"`
}
type Configuration struct {
	Auth struct {
		AccessToken       *string              `envconfig:"PINGONE_API_ACCESS_TOKEN" json:"accessToken,omitempty"`
		AccessTokenExpiry *int                 `envconfig:"PINGONE_API_ACCESS_TOKEN_EXPIRY" json:"accessTokenExpiry,omitempty"`
		AuthEnvironmentID *string              `envconfig:"PINGONE_AUTH_ENVIRONMENT_ID" json:"authEnvironmentId,omitempty"`
		AuthCode          *AuthCode            `envconfig:"PINGONE_AUTH_CODE" json:"authCode,omitempty"`
		ClientCredentials *ClientCredentials   `envconfig:"PINGONE_CLIENT_CREDENTIALS" json:"clientCredentials,omitempty"`
		DeviceCode        *DeviceCode          `envconfig:"PINGONE_DEVICE_CODE" json:"deviceCode,omitempty"`
		GrantType         *svcOAuth2.GrantType `envconfig:"PINGONE_AUTH_GRANT_TYPE" json:"grantType,omitempty"`
	} `json:"auth"`
	Endpoint struct {
		EnvironmentID  *string `envconfig:"PINGONE_ENVIRONMENT_ID" json:"environmentId,omitempty"`
		TopLevelDomain *string `envconfig:"PINGONE_TOP_LEVEL_DOMAIN" json:"topLevelDomain,omitempty"`
		RootDomain     *string `envconfig:"PINGONE_ROOT_DOMAIN" json:"rootDomain,omitempty"`
		APIDomain      *string `envconfig:"PINGONE_API_DOMAIN" json:"apiDomain,omitempty"`
		CustomDomain   *string `envconfig:"PINGONE_CUSTOM_DOMAIN" json:"customDomain,omitempty"`
	} `json:"endpoint"`
}

const (
	defaultSLD          = "pingone"
	defaultAPISubDomain = "api"
)

func NewConfiguration() *Configuration {

	cfg := &Configuration{}

	return cfg
}

func (c *Configuration) GetConfiguration() *Configuration {
	return c
}

func (c *Configuration) GetAccessToken() string {
	if c.Auth.AccessToken != nil {
		return *c.Auth.AccessToken
	}
	return ""
}

func (c *Configuration) GetAccessTokenExpiry() int {
	if c.Auth.AccessTokenExpiry != nil {
		return *c.Auth.AccessTokenExpiry
	}
	return 0
}

func (c *Configuration) WithEnvironmentID(environmentID string) *Configuration {
	c.Endpoint.EnvironmentID = &environmentID
	return c
}

func (c *Configuration) WithClientID(clientID string) *Configuration {
	c.Auth.ClientCredentials.ClientCredentialsClientID = &clientID
	return c
}

func (c *Configuration) WithGrantType(grantType svcOAuth2.GrantType) *Configuration {
	c.Auth.GrantType = &grantType
	return c
}

func (c *Configuration) WithClientSecret(clientSecret string) *Configuration {
	c.Auth.ClientCredentials.ClientCredentialsClientSecret = &clientSecret
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

func (c *Configuration) HasBearerToken() bool {
	return c.Auth.AccessToken != nil && *c.Auth.AccessToken != ""
}

func (c *Configuration) AddBearerTokenToContext(parent context.Context, key any) context.Context {
	if c.HasBearerToken() {
		return context.WithValue(parent, key, *c.Auth.AccessToken)
	}
	return parent
}

func (c *Configuration) BearerToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: *c.Auth.AccessToken,
		TokenType:   "Bearer",
	}
}

func (c *Configuration) Client(ctx context.Context, httpClient *http.Client) (*http.Client, error) {
	ts, err := c.TokenSource(ctx)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	client := oauth2.NewClient(ctx, *ts)

	return client, nil
}

func (c *Configuration) TokenSource(ctx context.Context) (*oauth2.TokenSource, error) {
	if at := c.Auth.AccessToken; at != nil && *at != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{
			AccessToken: *at,
			TokenType:   "Bearer",
		})
		slog.Debug("Using static token source as access token has been provided")
		return &ts, nil
	}

	endpoints, err := c.AuthEndpoints()
	if err != nil {
		return nil, err
	}

	if c.Auth.GrantType != nil {
		switch *c.Auth.GrantType {
		case svcOAuth2.GrantTypeAuthCode:
			return c.Auth.AuthCode.AuthCodeTokenSource(ctx, endpoints)
		case svcOAuth2.GrantTypeClientCredentials:
			return c.Auth.ClientCredentials.ClientCredentialsTokenSource(ctx, endpoints)
		case svcOAuth2.GrantTypeDeviceCode:
			return c.Auth.DeviceCode.DeviceAuthTokenSource(ctx, endpoints)
		}
	}

	return nil, fmt.Errorf("unsupported grant type")
}

func (c *Configuration) AuthEndpoints() (endpoints.OIDCEndpoint, error) {
	// Custom domain takes precedence over environment ID and root domain
	if v := c.Endpoint.CustomDomain; v != nil && *v != "" {
		return endpoints.PingOneOIDCEndpoint(*v), nil
	}

	if e := c.Endpoint.EnvironmentID; e != nil && *e != "" {
		if v := c.Endpoint.RootDomain; v != nil && *v != "" {
			return endpoints.PingOneEnvironmentOIDCEndpoint(*v, *e), nil
		}

		if v := c.Endpoint.TopLevelDomain; v != nil && strings.TrimPrefix(*v, ".") != "" {
			return endpoints.PingOneEnvironmentOIDCEndpoint(fmt.Sprintf("%s.%s", defaultSLD, strings.TrimPrefix(*v, ".")), *e), nil
		}
	}

	return endpoints.OIDCEndpoint{}, fmt.Errorf("no valid endpoint configuration found. Must provide either a custom domain or root domain/top level domain and auth environment ID")
}

func (c *Configuration) APIDomain() (string, error) {

	if v := c.Endpoint.APIDomain; v != nil && *v != "" {
		return *v, nil
	}

	if v := c.Endpoint.RootDomain; v != nil && strings.TrimPrefix(*v, ".") != "" {
		return fmt.Sprintf("%s.%s", defaultAPISubDomain, strings.TrimPrefix(*v, ".")), nil
	}

	if v := c.Endpoint.TopLevelDomain; v != nil && strings.TrimPrefix(*v, ".") != "" {
		return fmt.Sprintf("%s.%s.%s", defaultAPISubDomain, defaultSLD, strings.TrimPrefix(*v, ".")), nil
	}

	return "", fmt.Errorf("no valid endpoint configuration found. Must provide either a full API domain, a root domain, or a top level domain")
}
