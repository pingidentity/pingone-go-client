// Copyright © 2025 Ping Identity Corporation

// Package config provides configuration management for the PingOne Go Client SDK.
// It handles service configuration, environment variable parsing, OAuth2 credential management,
// and endpoint configuration for connecting to PingOne services across different regions.
package config

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"

	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
)

// Configuration represents the complete configuration for the PingOne Go Client SDK.
// It contains authentication settings and endpoint configuration for connecting to PingOne services.
// The configuration can be populated from environment variables or set explicitly through the builder methods.
type Configuration struct {
	// Auth contains authentication-related configuration including client credentials and grant types.
	Auth struct {
		ClientID     *string              `envconfig:"PINGONE_CLIENT_ID" json:"clientId,omitempty"`
		ClientSecret *string              `envconfig:"PINGONE_CLIENT_SECRET" json:"clientSecret,omitempty"`
		AccessToken  *string              `envconfig:"PINGONE_API_ACCESS_TOKEN" json:"accessToken,omitempty"`
		GrantType    *svcOAuth2.GrantType `envconfig:"PINGONE_AUTH_GRANT_TYPE" json:"grantType,omitempty"`
	} `json:"auth"`
	// Endpoint contains endpoint configuration for API and authentication domains.
	Endpoint struct {
		AuthEnvironmentID *string `envconfig:"PINGONE_ENVIRONMENT_ID" json:"environmentId,omitempty"`
		TopLevelDomain    *string `envconfig:"PINGONE_TOP_LEVEL_DOMAIN" json:"topLevelDomain,omitempty"`
		RootDomain        *string `envconfig:"PINGONE_ROOT_DOMAIN" json:"rootDomain,omitempty"`
		APIDomain         *string `envconfig:"PINGONE_API_DOMAIN" json:"apiDomain,omitempty"`
		CustomDomain      *string `envconfig:"PINGONE_CUSTOM_DOMAIN" json:"customDomain,omitempty"`
	} `json:"endpoint"`
}

const (
	// defaultSLD represents the default second-level domain used for PingOne services.
	defaultSLD = "pingone"
	// defaultAPISubDomain represents the default subdomain used for PingOne API endpoints.
	defaultAPISubDomain = "api"
)

// NewConfiguration creates a new Configuration instance with default values.
// This configuration must be populated with authentication credentials and endpoint information
// before it can be used to create API clients. Use the builder methods (With...) to set
// the required configuration values.
func NewConfiguration() *Configuration {

	cfg := &Configuration{}

	return cfg
}

// WithAuthEnvironmentID sets the PingOne environment ID for authentication.
// The environmentID parameter must be a valid PingOne environment UUID string.
// This environment ID is used to construct the OAuth2 authentication endpoints.
func (c *Configuration) WithAuthEnvironmentID(environmentID string) *Configuration {
	c.Endpoint.AuthEnvironmentID = &environmentID
	return c
}

// WithClientID sets the OAuth2 client ID for authentication.
// The clientID parameter must be a valid client ID from a PingOne application.
// This is required when using client credentials grant type authentication.
func (c *Configuration) WithClientID(clientID string) *Configuration {
	c.Auth.ClientID = &clientID
	return c
}

// WithGrantType sets the OAuth2 grant type for authentication.
// The grantType parameter specifies which OAuth2 flow to use for token acquisition.
// Currently supported grant types are defined in the oauth2 package.
func (c *Configuration) WithGrantType(grantType svcOAuth2.GrantType) *Configuration {
	c.Auth.GrantType = &grantType
	return c
}

// WithClientSecret sets the OAuth2 client secret for authentication.
// The clientSecret parameter must be the secret associated with the configured client ID.
// This is required when using client credentials grant type with client secret authentication.
func (c *Configuration) WithClientSecret(clientSecret string) *Configuration {
	c.Auth.ClientSecret = &clientSecret
	return c
}

// WithAccessToken sets a static access token for API authentication.
// The accessToken parameter must be a valid PingOne API access token.
// When set, this bypasses OAuth2 flows and uses the provided token directly.
func (c *Configuration) WithAccessToken(accessToken string) *Configuration {
	c.Auth.AccessToken = &accessToken
	return c
}

// WithTopLevelDomain sets the top-level domain for PingOne services.
// The tld parameter should be a domain suffix like "com", "eu", or "asia".
// This is used to construct service endpoints for different PingOne regions.
func (c *Configuration) WithTopLevelDomain(tld string) *Configuration {
	c.Endpoint.TopLevelDomain = &tld
	return c
}

// WithRootDomain sets the root domain for PingOne services.
// The rootDomain parameter should be a complete domain like "pingone.com" or "pingone.eu".
// This takes precedence over top-level domain configuration when constructing endpoints.
func (c *Configuration) WithRootDomain(rootDomain string) *Configuration {
	c.Endpoint.RootDomain = &rootDomain
	return c
}

// WithAPIDomain sets a custom API domain for PingOne services.
// The apiDomain parameter should be a complete API domain like "api.pingone.com".
// When set, this overrides automatic domain construction from other domain settings.
func (c *Configuration) WithAPIDomain(apiDomain string) *Configuration {
	c.Endpoint.APIDomain = &apiDomain
	return c
}

// WithCustomDomain sets a custom domain for PingOne authentication services.
// The customDomain parameter should be a complete custom domain configured in PingOne.
// When set, this takes precedence over all other domain configuration methods.
func (c *Configuration) WithCustomDomain(customDomain string) *Configuration {
	c.Endpoint.CustomDomain = &customDomain
	return c
}

// HasBearerToken checks if a static access token is configured.
// It returns true if an access token has been set and is not empty,
// false otherwise. This is used to determine whether to use static token
// authentication or OAuth2 flows.
func (c *Configuration) HasBearerToken() bool {
	return c.Auth.AccessToken != nil && *c.Auth.AccessToken != ""
}

// AddBearerTokenToContext adds the configured access token to a context.
// The parent parameter is the base context to extend, and key is the context key
// to use for storing the token value. It returns a new context with the token
// value if a token is configured, or the original context if no token is set.
func (c *Configuration) AddBearerTokenToContext(parent context.Context, key any) context.Context {
	if c.HasBearerToken() {
		return context.WithValue(parent, key, *c.Auth.AccessToken)
	}
	return parent
}

// BearerToken creates an OAuth2 token from the configured access token.
// It returns a pointer to an OAuth2 token with the access token and "Bearer" token type.
// This method assumes an access token has been configured; call HasBearerToken first
// to verify a token is available.
func (c *Configuration) BearerToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: *c.Auth.AccessToken,
		TokenType:   "Bearer",
	}
}

// Client creates an HTTP client configured for PingOne API authentication.
// The ctx parameter provides the context for the OAuth2 token source operations.
// The httpClient parameter is the base HTTP client to use for requests; if nil,
// a default client will be used. It returns an HTTP client that automatically
// handles authentication token management, or an error if token source creation fails.
func (c *Configuration) Client(ctx context.Context, httpClient *http.Client) (*http.Client, error) {
	ts, err := c.TokenSource(ctx)
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, oauth2.HTTPClient, httpClient)
	client := oauth2.NewClient(ctx, *ts)

	return client, nil
}

// TokenSource creates an OAuth2 token source based on the configured authentication method.
// The ctx parameter provides the context for OAuth2 operations. It returns a token source
// that can be used to obtain access tokens, or an error if the configuration is invalid
// or incomplete. If a static access token is configured, it returns a static token source.
// Otherwise, it creates a token source appropriate for the configured grant type.
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

// AuthEndpoints constructs the OAuth2 and OIDC endpoints for authentication.
// It returns an OIDCEndpoint struct containing the authorization, token, and userinfo URLs
// based on the configured domain settings. The method prioritizes custom domain over
// environment-specific domains. It returns an error if no valid endpoint configuration
// can be determined from the current settings.
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

// APIDomain constructs the API domain for PingOne service calls.
// It returns a string containing the complete API domain based on the configured
// domain settings. The method prioritizes explicit API domain configuration over
// root domain and top-level domain settings. It returns an error if no valid
// domain configuration can be determined from the current settings.
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
