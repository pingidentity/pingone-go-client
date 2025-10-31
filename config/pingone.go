// Copyright © 2025 Ping Identity Corporation

package config

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log/slog"
	"net/http"
	"strings"

	svcOAuth2 "github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"golang.org/x/oauth2"
)

func (c *Configuration) generateTokenKey(grantType svcOAuth2.GrantType) (string, error) {
	var environmentID, clientID string

	switch grantType {
	case svcOAuth2.GrantTypeDeviceCode:
		if c.Auth.DeviceCode != nil {
			if c.Auth.DeviceCode.DeviceCodeEnvironmentID != nil {
				environmentID = *c.Auth.DeviceCode.DeviceCodeEnvironmentID
			}
			if c.Auth.DeviceCode.DeviceCodeClientID != nil {
				clientID = *c.Auth.DeviceCode.DeviceCodeClientID
			}
		}
	case svcOAuth2.GrantTypeAuthCode:
		if c.Auth.AuthCode != nil {
			if c.Auth.AuthCode.AuthCodeEnvironmentID != nil {
				environmentID = *c.Auth.AuthCode.AuthCodeEnvironmentID
			}
			if c.Auth.AuthCode.AuthCodeClientID != nil {
				clientID = *c.Auth.AuthCode.AuthCodeClientID
			}
		}
	case svcOAuth2.GrantTypeClientCredentials:
		if c.Auth.ClientCredentials != nil {
			if c.Auth.ClientCredentials.ClientCredentialsEnvironmentID != nil {
				environmentID = *c.Auth.ClientCredentials.ClientCredentialsEnvironmentID
			}
			if c.Auth.ClientCredentials.ClientCredentialsClientID != nil {
				clientID = *c.Auth.ClientCredentials.ClientCredentialsClientID
			}
		}
	default:
		return "", fmt.Errorf("unsupported grant type: %s", grantType)
	}

	// Fallback to shared environment ID if grant-type-specific one is not available
	if environmentID == "" && c.Endpoint.EnvironmentID != nil {
		environmentID = *c.Endpoint.EnvironmentID
	}

	if environmentID == "" || clientID == "" {
		return "", fmt.Errorf("environment ID and client ID are required for token key generation")
	}

	hash := sha256.Sum256([]byte(fmt.Sprintf("%s:%s:%s", environmentID, clientID, grantType)))
	tokenKey := fmt.Sprintf("token-%x", hash[:8])

	slog.Debug("Generated token key", "environmentID", environmentID, "clientID", clientID, "grantType", grantType, "tokenKey", tokenKey)

	return tokenKey, nil
}

type AuthCodeRedirectURI struct {
	Port string `envconfig:"PINGONE_AUTH_CODE_REDIRECT_URI_PORT" json:"port,omitempty"`
	Path string `envconfig:"PINGONE_AUTH_CODE_REDIRECT_URI_PATH" json:"path,omitempty"`
}

type AuthCode struct {
	AuthCodeClientID      *string             `envconfig:"PINGONE_AUTH_CODE_CLIENT_ID" json:"authCodeClientId,omitempty"`
	AuthCodeEnvironmentID *string             `envconfig:"PINGONE_AUTH_CODE_ENVIRONMENT_ID" json:"authCodeEnvironmentId,omitempty"`
	AuthCodeRedirectURI   AuthCodeRedirectURI `envconfig:"PINGONE_AUTH_CODE_REDIRECT_URI" json:"redirectUri,omitempty"`
	AuthCodeScopes        *[]string           `envconfig:"PINGONE_AUTH_CODE_SCOPES" json:"authCodeScopes,omitempty"`
}

type ClientCredentials struct {
	ClientCredentialsClientID      *string   `envconfig:"PINGONE_CLIENT_CREDENTIALS_CLIENT_ID" json:"clientCredentialsClientId,omitempty"`
	ClientCredentialsClientSecret  *string   `envconfig:"PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET" json:"clientCredentialsClientSecret,omitempty"`
	ClientCredentialsEnvironmentID *string   `envconfig:"PINGONE_CLIENT_CREDENTIALS_ENVIRONMENT_ID" json:"clientCredentialsEnvironmentId,omitempty"`
	ClientCredentialsScopes        *[]string `envconfig:"PINGONE_CLIENT_CREDENTIALS_SCOPES" json:"clientCredentialsScopes,omitempty"`
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
		EnvironmentID  *string                   `envconfig:"PINGONE_ENVIRONMENT_ID" json:"environmentId,omitempty"`
		TopLevelDomain *svcOAuth2.TopLevelDomain `envconfig:"PINGONE_TOP_LEVEL_DOMAIN" json:"topLevelDomain,omitempty"`
		RootDomain     *string                   `envconfig:"PINGONE_ROOT_DOMAIN" json:"rootDomain,omitempty"`
		APIDomain      *string                   `envconfig:"PINGONE_API_DOMAIN" json:"apiDomain,omitempty"`
		CustomDomain   *string                   `envconfig:"PINGONE_CUSTOM_DOMAIN" json:"customDomain,omitempty"`
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

func (c *Configuration) GetAccessTokenExpiry() int {
	if c.Auth.AccessTokenExpiry != nil {
		return *c.Auth.AccessTokenExpiry
	}
	return 0
}

func (c *Configuration) WithClientCredentialsEnvironmentID(environmentID string) *Configuration {
	if c.Auth.ClientCredentials == nil {
		c.Auth.ClientCredentials = &ClientCredentials{}
	}
	c.Auth.ClientCredentials.ClientCredentialsEnvironmentID = &environmentID
	return c
}

func (c *Configuration) WithEnvironmentID(environmentID string) *Configuration {
	c.Endpoint.EnvironmentID = &environmentID
	return c
}

func (c *Configuration) WithClientCredentialsClientID(clientID string) *Configuration {
	if c.Auth.ClientCredentials == nil {
		c.Auth.ClientCredentials = &ClientCredentials{}
	}
	c.Auth.ClientCredentials.ClientCredentialsClientID = &clientID
	return c
}

func (c *Configuration) WithClientID(clientID string) *Configuration {
	return c.WithClientCredentialsClientID(clientID)
}

func (c *Configuration) WithGrantType(grantType svcOAuth2.GrantType) *Configuration {
	c.Auth.GrantType = &grantType
	return c
}

func (c *Configuration) WithClientCredentialsClientSecret(clientSecret string) *Configuration {
	if c.Auth.ClientCredentials == nil {
		c.Auth.ClientCredentials = &ClientCredentials{}
	}

	c.Auth.ClientCredentials.ClientCredentialsClientSecret = &clientSecret
	return c
}

func (c *Configuration) WithClientSecret(clientSecret string) *Configuration {
	return c.WithClientCredentialsClientSecret(clientSecret)
}

func (c *Configuration) WithClientCredentialsScopes(clientCredentialsScopes []string) *Configuration {
	if c.Auth.ClientCredentials == nil {
		c.Auth.ClientCredentials = &ClientCredentials{}
	}
	c.Auth.ClientCredentials.ClientCredentialsScopes = &clientCredentialsScopes
	return c
}

func (c *Configuration) WithAccessToken(accessToken string) *Configuration {
	c.Auth.AccessToken = &accessToken
	return c
}

func (c *Configuration) GetTopLevelDomain() string {
	if c.Endpoint.TopLevelDomain != nil {
		return string(*c.Endpoint.TopLevelDomain)
	}
	return ""
}

func (c *Configuration) WithTopLevelDomain(tld svcOAuth2.TopLevelDomain) *Configuration {
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

func (c *Configuration) WithAuthCodeClientID(authCodeClientID string) *Configuration {
	if c.Auth.AuthCode == nil {
		c.Auth.AuthCode = &AuthCode{}
	}
	c.Auth.AuthCode.AuthCodeClientID = &authCodeClientID
	return c
}

func (c *Configuration) WithAuthCodeEnvironmentID(authCodeEnvironmentID string) *Configuration {
	if c.Auth.AuthCode == nil {
		c.Auth.AuthCode = &AuthCode{}
	}
	c.Auth.AuthCode.AuthCodeEnvironmentID = &authCodeEnvironmentID
	return c
}

func (c *Configuration) WithAuthCodeScopes(authCodeScopes []string) *Configuration {
	if c.Auth.AuthCode == nil {
		c.Auth.AuthCode = &AuthCode{}
	}
	c.Auth.AuthCode.AuthCodeScopes = &authCodeScopes
	return c
}

func (c *Configuration) WithAuthCodeRedirectURI(authCodeRedirectURI AuthCodeRedirectURI) *Configuration {
	if c.Auth.AuthCode == nil {
		c.Auth.AuthCode = &AuthCode{}
	}
	c.Auth.AuthCode.AuthCodeRedirectURI = AuthCodeRedirectURI{
		Port: authCodeRedirectURI.Port,
		Path: authCodeRedirectURI.Path,
	}
	return c
}

func (c *Configuration) WithDeviceCodeClientID(deviceCodeClientID string) *Configuration {
	if c.Auth.DeviceCode == nil {
		c.Auth.DeviceCode = &DeviceCode{}
	}
	c.Auth.DeviceCode.DeviceCodeClientID = &deviceCodeClientID
	return c
}

func (c *Configuration) WithDeviceCodeEnvironmentID(deviceCodeEnvironmentID string) *Configuration {
	if c.Auth.DeviceCode == nil {
		c.Auth.DeviceCode = &DeviceCode{}
	}
	c.Auth.DeviceCode.DeviceCodeEnvironmentID = &deviceCodeEnvironmentID
	return c
}

func (c *Configuration) WithDeviceCodeScopes(deviceCodeScopes []string) *Configuration {
	if c.Auth.DeviceCode == nil {
		c.Auth.DeviceCode = &DeviceCode{}
	}
	c.Auth.DeviceCode.DeviceCodeScopes = &deviceCodeScopes
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
	client := oauth2.NewClient(ctx, ts)

	return client, nil
}

func (c *Configuration) GetAccessToken(ctx context.Context) (string, error) {
	ts, err := c.TokenSource(ctx)
	if err != nil {
		return "", err
	}

	token, err := ts.Token()
	if err != nil {
		return "", err
	}

	if token.AccessToken == "" {
		return "", fmt.Errorf("failed to retrieve access token")
	}

	return token.AccessToken, nil
}

func (c *Configuration) TokenSource(ctx context.Context) (oauth2.TokenSource, error) {
	// Client credentials flow should always use the OAuth2 token source for automatic refresh
	if at := c.Auth.AccessToken; at != nil && *at != "" {
		if c.Auth.GrantType != nil && *c.Auth.GrantType == svcOAuth2.GrantTypeClientCredentials {
			// Don't use static token source for client credentials - fall through to proper OAuth2 flow
			slog.Debug("Skipping static token source for client credentials flow to enable automatic token refresh")
		} else {
			ts := oauth2.StaticTokenSource(&oauth2.Token{
				AccessToken: *at,
				TokenType:   "Bearer",
			})
			slog.Debug("Using static token source as access token has been provided")
			return ts, nil
		}
	}

	// Check keychain for existing valid token before performing auth
	if c.Auth.GrantType != nil {
		tokenKey, err := c.generateTokenKey(*c.Auth.GrantType)
		if err != nil {
			slog.Debug("Could not generate token key for caching", "error", err)
		} else {
			keychainStorage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
			if existingToken, err := keychainStorage.LoadToken(); err == nil && existingToken != nil && existingToken.Valid() {
				slog.Debug("Using cached token from keychain", "tokenKey", tokenKey, "expires", existingToken.Expiry)
				return oauth2.StaticTokenSource(existingToken), nil
			}
		}
	}

	endpoints, err := c.AuthEndpoints()
	if err != nil {
		return nil, err
	}

	// Perform authentication and cache the result
	if c.Auth.GrantType != nil {
		var tokenSource oauth2.TokenSource

		// Generate unique token key for caching
		tokenKey, err := c.generateTokenKey(*c.Auth.GrantType)
		if err != nil {
			return nil, fmt.Errorf("failed to generate token key: %w", err)
		}

		// Get the appropriate token source
		switch *c.Auth.GrantType {
		case svcOAuth2.GrantTypeAuthCode:
			if c.Auth.AuthCode == nil {
				return nil, fmt.Errorf("auth code configuration is required for auth code grant type")
			}
			ts, err := c.Auth.AuthCode.AuthCodeTokenSource(ctx, endpoints)
			if err != nil {
				return nil, err
			}
			tokenSource = ts
		case svcOAuth2.GrantTypeClientCredentials:
			if c.Auth.ClientCredentials == nil {
				return nil, fmt.Errorf("client credentials configuration is required for client credentials grant type")
			}
			ts, err := c.Auth.ClientCredentials.ClientCredentialsTokenSource(ctx, endpoints)
			if err != nil {
				return nil, err
			}
			tokenSource = ts
		case svcOAuth2.GrantTypeDeviceCode:
			if c.Auth.DeviceCode == nil {
				return nil, fmt.Errorf("device code configuration is required for device code grant type")
			}
			ts, err := c.Auth.DeviceCode.DeviceAuthTokenSource(ctx, endpoints)
			if err != nil {
				return nil, err
			}
			tokenSource = ts
		default:
			return nil, fmt.Errorf("unsupported grant type: %s", *c.Auth.GrantType)
		}

		// Get token to validate and cache it
		if tokenSource != nil && tokenKey != "" {
			token, err := tokenSource.Token()
			if err != nil {
				return nil, fmt.Errorf("failed to get token: %w", err)
			}

			// Save token to keychain for future use
			keychainStorage := svcOAuth2.NewKeychainStorage("pingcli", tokenKey)
			if err := keychainStorage.SaveToken(token); err != nil {
				slog.Warn("Failed to save token to keychain", "error", err, "tokenKey", tokenKey)
				// Don't return error here - token is still valid even if caching failed
			} else {
				slog.Debug("Token saved to keychain", "tokenKey", tokenKey, "expires", token.Expiry)
			}

			return oauth2.StaticTokenSource(token), nil
		}

		return tokenSource, nil
	}

	return nil, fmt.Errorf("grant type is required")
}

func (c *Configuration) AuthEndpoints() (endpoints.OIDCEndpoint, error) {
	// Custom domain takes precedence over environment ID and root domain
	if customDomain := c.Endpoint.CustomDomain; customDomain != nil && *customDomain != "" {
		return endpoints.PingOneOIDCEndpoint(*customDomain), nil
	}

	// Use the appropriate environment ID based on the grant type
	var environmentID string
	if c.Auth.GrantType != nil {
		switch *c.Auth.GrantType {
		case svcOAuth2.GrantTypeAuthCode:
			if c.Auth.AuthCode != nil && c.Auth.AuthCode.AuthCodeEnvironmentID != nil && *c.Auth.AuthCode.AuthCodeEnvironmentID != "" {
				environmentID = *c.Auth.AuthCode.AuthCodeEnvironmentID
			}
		case svcOAuth2.GrantTypeClientCredentials:
			if c.Auth.ClientCredentials != nil && c.Auth.ClientCredentials.ClientCredentialsEnvironmentID != nil && *c.Auth.ClientCredentials.ClientCredentialsEnvironmentID != "" {
				environmentID = *c.Auth.ClientCredentials.ClientCredentialsEnvironmentID
			}
		case svcOAuth2.GrantTypeDeviceCode:
			if c.Auth.DeviceCode != nil && c.Auth.DeviceCode.DeviceCodeEnvironmentID != nil && *c.Auth.DeviceCode.DeviceCodeEnvironmentID != "" {
				environmentID = *c.Auth.DeviceCode.DeviceCodeEnvironmentID
			}
		}
	}

	// Fallback to the shared environment ID if grant-type-specific one is not available
	if environmentID == "" {
		if c.Endpoint.EnvironmentID != nil && *c.Endpoint.EnvironmentID != "" {
			environmentID = *c.Endpoint.EnvironmentID
		}
	}

	if environmentID != "" {
		if v := c.Endpoint.RootDomain; v != nil && *v != "" {
			return endpoints.PingOneEnvironmentOIDCEndpoint(*v, environmentID), nil
		}

		if tld := c.Endpoint.TopLevelDomain; tld != nil && strings.TrimPrefix(string(*tld), ".") != "" {
			return endpoints.PingOneEnvironmentOIDCEndpoint(fmt.Sprintf("%s.%s", defaultSLD, strings.TrimPrefix(string(*tld), ".")), environmentID), nil
		}
	}

	return endpoints.OIDCEndpoint{}, fmt.Errorf("no valid endpoint configuration found. Must provide either a custom domain or root domain/top level domain and auth environment ID")
}

func (c *Configuration) APIDomain() (string, error) {

	if apiDomain := c.Endpoint.APIDomain; apiDomain != nil && *apiDomain != "" {
		return *apiDomain, nil
	}

	if rootDomain := c.Endpoint.RootDomain; rootDomain != nil && strings.TrimPrefix(*rootDomain, ".") != "" {
		return fmt.Sprintf("%s.%s", defaultAPISubDomain, strings.TrimPrefix(*rootDomain, ".")), nil
	}

	if topLevelDomain := c.Endpoint.TopLevelDomain; topLevelDomain != nil && strings.TrimPrefix(string(*topLevelDomain), ".") != "" {
		return fmt.Sprintf("%s.%s.%s", defaultAPISubDomain, defaultSLD, strings.TrimPrefix(string(*topLevelDomain), ".")), nil
	}

	return "", fmt.Errorf("no valid endpoint configuration found. Must provide either a full API domain, a root domain, or a top level domain")
}
