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
	"golang.org/x/oauth2/clientcredentials"
)

type Configuration struct {
	Auth struct {
		ClientID     *string              `envconfig:"PINGONE_CLIENT_ID" json:"clientId,omitempty"`
		ClientSecret *string              `envconfig:"PINGONE_CLIENT_SECRET" json:"clientSecret,omitempty"`
		AccessToken  *string              `envconfig:"PINGONE_API_ACCESS_TOKEN" json:"accessToken,omitempty"`
		GrantType    *svcOAuth2.GrantType `envconfig:"PINGONE_AUTH_GRANT_TYPE" json:"grantType,omitempty"`
	} `json:"auth"`
	Endpoint struct {
		AuthEnvironmentID *string `envconfig:"PINGONE_ENVIRONMENT_ID" json:"environmentId,omitempty"`
		TopLevelDomain    *string `envconfig:"PINGONE_TOP_LEVEL_DOMAIN" json:"topLevelDomain,omitempty"`
		RootDomain        *string `envconfig:"PINGONE_ROOT_DOMAIN" json:"rootDomain,omitempty"`
		APIDomain         *string `envconfig:"PINGONE_API_DOMAIN" json:"apiDomain,omitempty"`
		CustomDomain      *string `envconfig:"PINGONE_CUSTOM_DOMAIN" json:"customDomain,omitempty"`
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

func (c *Configuration) WithAuthEnvironmentID(environmentID string) *Configuration {
	c.Endpoint.AuthEnvironmentID = &environmentID
	return c
}

func (c *Configuration) WithClientID(clientID string) *Configuration {
	c.Auth.ClientID = &clientID
	return c
}

func (c *Configuration) WithGrantType(grantType svcOAuth2.GrantType) *Configuration {
	c.Auth.GrantType = &grantType
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

	if c.Auth.GrantType != nil && *c.Auth.GrantType == svcOAuth2.GrantTypeClientCredentials {
		if c.Auth.ClientID == nil || *c.Auth.ClientID == "" {
			return nil, fmt.Errorf("client ID is required for client credentials grant type")
		}

		slog.Debug("Using client credentials token source with provided client ID", "client ID", *c.Auth.ClientID)
		if c.Auth.ClientSecret != nil && *c.Auth.ClientSecret != "" {
			config := &clientcredentials.Config{
				ClientID:     *c.Auth.ClientID,
				ClientSecret: *c.Auth.ClientSecret,
				TokenURL:     endpoints.TokenURL,
			}
			ts := config.TokenSource(ctx)
			slog.Debug("Using standard client credentials token source as client secret has been provided")
			return &ts, nil
		}

		// tmp until private keys are supported
		return nil, fmt.Errorf("client secret is required for client credentials grant type")
	}

	return nil, fmt.Errorf("unsupported grant type: %s", *c.Auth.GrantType)
}

func (c *Configuration) AuthEndpoints() (endpoints.OIDCEndpoint, error) {
	// Custom domain takes precedence over environment ID and root domain
	if v := c.Endpoint.CustomDomain; v != nil && *v != "" {
		return endpoints.PingOneOIDCEndpoint(*v), nil
	}

	if e := c.Endpoint.AuthEnvironmentID; e != nil && *e != "" {
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
