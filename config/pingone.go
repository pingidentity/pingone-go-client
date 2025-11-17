// Copyright © 2025 Ping Identity Corporation

// Package config provides configuration management for the PingOne Go Client SDK.
// It handles service configuration, environment variable parsing, OAuth2 credential management,
// and endpoint configuration for connecting to PingOne services across different regions.
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

// keychainTokenSource wraps an oauth2.TokenSource and saves tokens to keychain after refresh
type keychainTokenSource struct {
	base            oauth2.TokenSource
	keychainStorage *svcOAuth2.KeychainStorage
	tokenKey        string
}

// Token implements oauth2.TokenSource interface
// It gets a token from the base source and saves it to keychain
func (k *keychainTokenSource) Token() (*oauth2.Token, error) {
	token, err := k.base.Token()
	if err != nil {
		return nil, err
	}

	// Save the token to keychain for future use
	if saveErr := k.keychainStorage.SaveToken(token); saveErr != nil {
		slog.Warn("Failed to save refreshed token to keychain", "error", saveErr, "tokenKey", k.tokenKey)
		// Don't return error - token is still valid even if caching failed
	} else {
		slog.Debug("Refreshed token saved to keychain", "tokenKey", k.tokenKey, "expires", token.Expiry)
	}

	return token, nil
}

func (c *Configuration) generateTokenKey(grantType svcOAuth2.GrantType) (string, error) {
	var environmentID, clientID string

	// Get client ID based on grant type
	switch grantType {
	case svcOAuth2.GrantTypeAuthorizationCode:
		if c.Auth.AuthorizationCode != nil && c.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
			clientID = *c.Auth.AuthorizationCode.AuthorizationCodeClientID
		}
	case svcOAuth2.GrantTypeClientCredentials:
		if c.Auth.ClientCredentials != nil && c.Auth.ClientCredentials.ClientCredentialsClientID != nil {
			clientID = *c.Auth.ClientCredentials.ClientCredentialsClientID
		}
	case svcOAuth2.GrantTypeDeviceCode:
		if c.Auth.DeviceCode != nil && c.Auth.DeviceCode.DeviceCodeClientID != nil {
			clientID = *c.Auth.DeviceCode.DeviceCodeClientID
		}
	default:
		return "", fmt.Errorf("unsupported grant type: %s", grantType)
	}

	// Use Endpoint.EnvironmentID for token key generation
	if c.Endpoint.EnvironmentID != nil {
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

// createOAuth2ConfigForRefresh creates a minimal OAuth2 config for token refresh
func (c *Configuration) createOAuth2ConfigForRefresh(endpoints endpoints.OIDCEndpoint) (*oauth2.Config, error) {
	if c.Auth.GrantType == nil {
		return nil, fmt.Errorf("grant type is required")
	}

	switch *c.Auth.GrantType {
	case svcOAuth2.GrantTypeDeviceCode:
		if c.Auth.DeviceCode != nil && c.Auth.DeviceCode.DeviceCodeClientID != nil {
			var scopes []string
			if c.Auth.DeviceCode.DeviceCodeScopes != nil {
				scopes = *c.Auth.DeviceCode.DeviceCodeScopes
			}
			return &oauth2.Config{
				ClientID: *c.Auth.DeviceCode.DeviceCodeClientID,
				Endpoint: endpoints.Endpoint,
				Scopes:   scopes,
			}, nil
		}
	case svcOAuth2.GrantTypeAuthorizationCode:
		if c.Auth.AuthorizationCode != nil && c.Auth.AuthorizationCode.AuthorizationCodeClientID != nil {
			var scopes []string
			if c.Auth.AuthorizationCode.AuthorizationCodeScopes != nil {
				scopes = *c.Auth.AuthorizationCode.AuthorizationCodeScopes
			}
			return &oauth2.Config{
				ClientID: *c.Auth.AuthorizationCode.AuthorizationCodeClientID,
				Endpoint: endpoints.Endpoint,
				Scopes:   scopes,
			}, nil
		}
	}

	return nil, fmt.Errorf("no valid configuration found for grant type: %s", *c.Auth.GrantType)
}

type AuthorizationCodeRedirectURI struct {
	Port string `envconfig:"PINGONE_AUTHORIZATION_CODE_REDIRECT_URI_PORT" json:"port,omitempty"`
	Path string `envconfig:"PINGONE_AUTHORIZATION_CODE_REDIRECT_URI_PATH" json:"path,omitempty"`
}

type AuthorizationCode struct {
	AuthorizationCodeClientID    *string                      `envconfig:"PINGONE_AUTHORIZATION_CODE_CLIENT_ID" json:"authorizationCodeClientId,omitempty"`
	AuthorizationCodeRedirectURI AuthorizationCodeRedirectURI `envconfig:"PINGONE_AUTHORIZATION_CODE_REDIRECT_URI" json:"authorizationCodeRedirectUri,omitempty"`
	AuthorizationCodeScopes      *[]string                    `envconfig:"PINGONE_AUTHORIZATION_CODE_SCOPES" json:"authorizationCodeScopes,omitempty"`
}

type ClientCredentials struct {
	ClientCredentialsClientID     *string   `envconfig:"PINGONE_CLIENT_CREDENTIALS_CLIENT_ID" json:"clientCredentialsClientId,omitempty"`
	ClientCredentialsClientSecret *string   `envconfig:"PINGONE_CLIENT_CREDENTIALS_CLIENT_SECRET" json:"clientCredentialsClientSecret,omitempty"`
	ClientCredentialsScopes       *[]string `envconfig:"PINGONE_CLIENT_CREDENTIALS_SCOPES" json:"clientCredentialsScopes,omitempty"`
}

type DeviceCode struct {
	DeviceCodeClientID *string   `envconfig:"PINGONE_DEVICE_CODE_CLIENT_ID" json:"deviceCodeClientId,omitempty"`
	DeviceCodeScopes   *[]string `envconfig:"PINGONE_DEVICE_CODE_SCOPES" json:"deviceCodeScopes,omitempty"`
}

type Storage struct {
	KeychainName string      `envconfig:"PINGONE_STORAGE_NAME" json:"name,omitempty"`
	Type         StorageType `envconfig:"PINGONE_STORAGE_TYPE" json:"type,omitempty"`
}

// Configuration represents the complete configuration for the PingOne Go Client SDK.
// It contains authentication settings and endpoint configuration for connecting to PingOne services.
// The configuration can be populated from environment variables or set explicitly through the builder methods.
type Configuration struct {
	// Auth contains authentication-related configuration including client credentials and grant types.
	Auth struct {
		AccessToken       *string              `envconfig:"PINGONE_API_ACCESS_TOKEN" json:"accessToken,omitempty"`
		AccessTokenExpiry *int                 `envconfig:"PINGONE_API_ACCESS_TOKEN_EXPIRY" json:"accessTokenExpiry,omitempty"`
		AuthorizationCode *AuthorizationCode   `envconfig:"PINGONE_AUTHORIZATION_CODE" json:"authorizationCode,omitempty"`
		ClientCredentials *ClientCredentials   `envconfig:"PINGONE_CLIENT_CREDENTIALS" json:"clientCredentials,omitempty"`
		DeviceCode        *DeviceCode          `envconfig:"PINGONE_DEVICE_CODE" json:"deviceCode,omitempty"`
		GrantType         *svcOAuth2.GrantType `envconfig:"PINGONE_AUTH_GRANT_TYPE" json:"grantType,omitempty"`
		Storage           *Storage             `envconfig:"PINGONE_STORAGE_OPTIONS" json:"storageOptions,omitempty"`
	} `json:"auth"`
	// Endpoint contains endpoint configuration for API and authentication domains.
	Endpoint struct {
		EnvironmentID  *string         `envconfig:"PINGONE_ENVIRONMENT_ID" json:"environmentId,omitempty"`
		TopLevelDomain *TopLevelDomain `envconfig:"PINGONE_TOP_LEVEL_DOMAIN" json:"topLevelDomain,omitempty"`
		RootDomain     *string         `envconfig:"PINGONE_ROOT_DOMAIN" json:"rootDomain,omitempty"`
		APIDomain      *string         `envconfig:"PINGONE_API_DOMAIN" json:"apiDomain,omitempty"`
		CustomDomain   *string         `envconfig:"PINGONE_CUSTOM_DOMAIN" json:"customDomain,omitempty"`
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
func (c *Configuration) GetConfiguration() *Configuration {
	return c
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

func (c *Configuration) WithClientCredentialsClientID(clientID string) *Configuration {
	if c.Auth.ClientCredentials == nil {
		c.Auth.ClientCredentials = &ClientCredentials{}
	}
	c.Auth.ClientCredentials.ClientCredentialsClientID = &clientID
	return c
}

// WithClientID sets the OAuth2 client ID for authentication.
// The clientID parameter must be a valid client ID from a PingOne application.
// This is required when using client credentials grant type authentication.
func (c *Configuration) WithClientID(clientID string) *Configuration {
	return c.WithClientCredentialsClientID(clientID)
}

// WithGrantType sets the OAuth2 grant type for authentication.
// The grantType parameter specifies which OAuth2 flow to use for token acquisition.
// Currently supported grant types are defined in the oauth2 package.
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

// WithClientSecret sets the OAuth2 client secret for authentication.
// The clientSecret parameter must be the secret associated with the configured client ID.
// This is required when using client credentials grant type with client secret authentication.
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

// WithAccessToken sets a static access token for API authentication.
// The accessToken parameter must be valid PingOne API access token.
// When set, this bypasses OAuth2 flows and uses the provided token directly.
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

// WithTopLevelDomain sets the top-level domain for PingOne services.
// The tld parameter should be a domain suffix like "com", "eu", or "asia".
// This is used to construct service endpoints for different PingOne regions.
func (c *Configuration) WithTopLevelDomain(tld TopLevelDomain) *Configuration {
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

func (c *Configuration) WithAuthorizationCodeClientID(authorizationCodeClientID string) *Configuration {
	if c.Auth.AuthorizationCode == nil {
		c.Auth.AuthorizationCode = &AuthorizationCode{}
	}
	c.Auth.AuthorizationCode.AuthorizationCodeClientID = &authorizationCodeClientID
	return c
}

func (c *Configuration) WithAuthorizationCodeScopes(authorizationCodeScopes []string) *Configuration {
	if c.Auth.AuthorizationCode == nil {
		c.Auth.AuthorizationCode = &AuthorizationCode{}
	}
	c.Auth.AuthorizationCode.AuthorizationCodeScopes = &authorizationCodeScopes
	return c
}

func (c *Configuration) WithAuthorizationCodeRedirectURI(authorizationCodeRedirectURI AuthorizationCodeRedirectURI) *Configuration {
	if c.Auth.AuthorizationCode == nil {
		c.Auth.AuthorizationCode = &AuthorizationCode{}
	}
	c.Auth.AuthorizationCode.AuthorizationCodeRedirectURI = AuthorizationCodeRedirectURI{
		Port: authorizationCodeRedirectURI.Port,
		Path: authorizationCodeRedirectURI.Path,
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

func (c *Configuration) WithDeviceCodeScopes(deviceCodeScopes []string) *Configuration {
	if c.Auth.DeviceCode == nil {
		c.Auth.DeviceCode = &DeviceCode{}
	}
	c.Auth.DeviceCode.DeviceCodeScopes = &deviceCodeScopes
	return c
}

func (c *Configuration) WithUseKeychain(useKeychain bool) *Configuration {
	if c.Auth.Storage == nil {
		c.Auth.Storage = &Storage{}
	}
	// Set the storage type based on useKeychain value
	if useKeychain {
		c.Auth.Storage.Type = StorageTypeKeychain
	}
	return c
}

func (c *Configuration) WithStorageType(storageType StorageType) *Configuration {
	if c.Auth.Storage == nil {
		c.Auth.Storage = &Storage{}
	}
	c.Auth.Storage.Type = storageType
	return c
}

func (c *Configuration) WithStorageName(name string) *Configuration {
	if c.Auth.Storage == nil {
		c.Auth.Storage = &Storage{}
	}
	c.Auth.Storage.KeychainName = name
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

// TokenSource creates an OAuth2 token source based on the configured authentication method.
// The ctx parameter provides the context for OAuth2 operations. It returns a token source
// that can be used to obtain access tokens, or an error if the configuration is invalid
// or incomplete. If a static access token is configured, it returns a static token source.
// Otherwise, it creates a token source appropriate for the configured grant type.
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
		// Default to storage type of keychain
		shouldCheckKeychain := c.Auth.Storage == nil || c.Auth.Storage.Type == "" || c.Auth.Storage.Type == StorageTypeKeychain

		if shouldCheckKeychain {
			// Validate storage name is set when using keychain
			if c.Auth.Storage == nil || c.Auth.Storage.KeychainName == "" {
				return nil, fmt.Errorf("storage name is required when using keychain storage. Use WithStorageName() to set it")
			}

			tokenKey, err := c.generateTokenKey(*c.Auth.GrantType)
			if err != nil {
				return nil, fmt.Errorf("could not generate token key for caching: %s", err)
			} else {
				keychainStorage, err := svcOAuth2.NewKeychainStorage(c.Auth.Storage.KeychainName, tokenKey)
				if err != nil {
					return nil, fmt.Errorf("failed to create keychain storage: %w", err)
				}
				existingToken, err := keychainStorage.LoadToken()
				if err == nil && existingToken != nil && existingToken.RefreshToken != "" {
					// Token has refresh token - set up automatic refresh
					slog.Debug("Found cached token with refresh capability", "tokenKey", tokenKey, "expires", existingToken.Expiry, "valid", existingToken.Valid())

					endpoints, err := c.AuthEndpoints()
					if err != nil {
						slog.Warn("Failed to get endpoints for token refresh", "error", err)
					} else {
						oauthConfig, err := c.createOAuth2ConfigForRefresh(endpoints)
						if err != nil {
							slog.Warn("Failed to create OAuth2 config for refresh", "error", err)
						} else {
							// oauth2.Config.TokenSource automatically handles refresh when needed
							baseTS := oauthConfig.TokenSource(ctx, existingToken)
							ts := oauth2.ReuseTokenSource(nil, &keychainTokenSource{
								base:            baseTS,
								keychainStorage: keychainStorage,
								tokenKey:        tokenKey,
							})

							// Verify token source works (will use cached token if valid, or refresh if expired)
							if _, err := ts.Token(); err == nil {
								slog.Debug("Token source ready", "tokenKey", tokenKey)
								return ts, nil
							}
							slog.Debug("Token source failed, will re-authenticate", "error", err)
						}
					}
				} else if err != nil {
					slog.Debug("Unable to load cached token from keychain:", "error", err)
				}
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
		case svcOAuth2.GrantTypeAuthorizationCode:
			if c.Auth.AuthorizationCode == nil {
				return nil, fmt.Errorf("authorization code configuration is required for auth code grant type")
			}
			ts, err := c.Auth.AuthorizationCode.AuthorizationCodeTokenSource(ctx, endpoints)
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
			ts, err := c.Auth.DeviceCode.DeviceAuthTokenSource(ctx, endpoints.Endpoint)
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

			// Save token to keychain for future use - only if storage type is keychain (or not set)
			shouldSaveToKeychain := c.Auth.Storage == nil || c.Auth.Storage.Type == "" || c.Auth.Storage.Type == StorageTypeKeychain

			if shouldSaveToKeychain {
				// Validate storage name is set when using keychain
				if c.Auth.Storage == nil || c.Auth.Storage.KeychainName == "" {
					return nil, fmt.Errorf("storage name is required when using keychain storage. Use WithStorageName() to set it")
				}

				keychainStorage, err := svcOAuth2.NewKeychainStorage(c.Auth.Storage.KeychainName, tokenKey)
				if err != nil {
					return nil, fmt.Errorf("failed to create keychain storage: %w", err)
				}
				if err := keychainStorage.SaveToken(token); err != nil {
					slog.Warn("Failed to save token to keychain", "error", err, "tokenKey", tokenKey)
					// Don't return error here - token is still valid even if caching failed
				} else {
					slog.Debug("Token saved to keychain", "tokenKey", tokenKey, "expires", token.Expiry)
				}

				// If token has refresh capability, wrap with automatic refresh support
				if token.RefreshToken != "" {
					endpoints, err := c.AuthEndpoints()
					if err != nil {
						slog.Warn("Failed to get endpoints for token refresh", "error", err)
					} else {
						oauthConfig, err := c.createOAuth2ConfigForRefresh(endpoints)
						if err != nil {
							slog.Warn("Failed to create OAuth2 config for refresh", "error", err)
						} else {
							// Use ReuseTokenSource with keychainTokenSource wrapper for automatic refresh
							baseTS := oauthConfig.TokenSource(ctx, token)
							return oauth2.ReuseTokenSource(nil, &keychainTokenSource{
								base:            baseTS,
								keychainStorage: keychainStorage,
								tokenKey:        tokenKey,
							}), nil
						}
					}
				}
			} else {
				slog.Debug("Skipping keychain storage (file storage mode)", "tokenKey", tokenKey)
			}

			return oauth2.StaticTokenSource(token), nil
		}

		return tokenSource, nil
	}

	return nil, fmt.Errorf("grant type is required")
}

// AuthEndpoints constructs the OAuth2 and OIDC endpoints for authentication.
// It returns an OIDCEndpoint struct containing the authorization, token, and userinfo URLs
// based on the configured domain settings. The method prioritizes custom domain over
// environment-specific domains. It returns an error if no valid endpoint configuration
// can be determined from the current settings.
func (c *Configuration) AuthEndpoints() (endpoints.OIDCEndpoint, error) {
	// Custom domain takes precedence over environment ID and root domain
	if customDomain := c.Endpoint.CustomDomain; customDomain != nil && *customDomain != "" {
		return endpoints.PingOneOIDCEndpoint(*customDomain), nil
	}

	// Use Endpoint.EnvironmentID for constructing auth endpoints
	var environmentID string
	if c.Endpoint.EnvironmentID != nil && *c.Endpoint.EnvironmentID != "" {
		environmentID = *c.Endpoint.EnvironmentID
	}

	if environmentID != "" {
		if v := c.Endpoint.RootDomain; v != nil && *v != "" {
			return endpoints.PingOneEnvironmentOIDCEndpoint(*v, environmentID), nil
		}

		if tld := c.Endpoint.TopLevelDomain; tld != nil && strings.TrimPrefix(string(*tld), ".") != "" {
			return endpoints.PingOneEnvironmentOIDCEndpoint(fmt.Sprintf("%s.%s", defaultSLD, strings.TrimPrefix(string(*tld), ".")), environmentID), nil
		}
	}

	return endpoints.OIDCEndpoint{}, fmt.Errorf("no valid endpoint configuration found. Must provide either a custom domain or root domain/top level domain and environment ID")
}

// APIDomain constructs the API domain for PingOne service calls.
// It returns a string containing the complete API domain based on the configured
// domain settings. The method prioritizes explicit API domain configuration over
// root domain and top-level domain settings. It returns an error if no valid
// domain configuration can be determined from the current settings.
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
