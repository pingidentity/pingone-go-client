// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oauth2"
)

func TestNewConfiguration(t *testing.T) {
	cfg := config.NewConfiguration()
	assert.NotNil(t, cfg, "Configuration should not be nil")
}

func TestWithAuthEnvironmentID(t *testing.T) {
	cfg := config.NewConfiguration()
	envID := "test-environment-id"
	cfg = cfg.WithAuthEnvironmentID(envID)

	assert.NotNil(t, cfg.Endpoint.AuthEnvironmentID)
	assert.Equal(t, envID, *cfg.Endpoint.AuthEnvironmentID)
}

func TestWithClientID(t *testing.T) {
	cfg := config.NewConfiguration()
	clientID := "test-client-id"
	cfg = cfg.WithClientID(clientID)

	assert.NotNil(t, cfg.Auth.ClientID)
	assert.Equal(t, clientID, *cfg.Auth.ClientID)
}

func TestWithGrantType(t *testing.T) {
	cfg := config.NewConfiguration()
	grantType := oauth2.GrantTypeClientCredentials
	cfg = cfg.WithGrantType(grantType)

	assert.NotNil(t, cfg.Auth.GrantType)
	assert.Equal(t, grantType, *cfg.Auth.GrantType)
}

func TestWithClientSecret(t *testing.T) {
	cfg := config.NewConfiguration()
	clientSecret := "test-client-secret"
	cfg = cfg.WithClientSecret(clientSecret)

	assert.NotNil(t, cfg.Auth.ClientSecret)
	assert.Equal(t, clientSecret, *cfg.Auth.ClientSecret)
}

func TestWithAccessToken(t *testing.T) {
	cfg := config.NewConfiguration()
	accessToken := "test-access-token"
	cfg = cfg.WithAccessToken(accessToken)

	assert.NotNil(t, cfg.Auth.AccessToken)
	assert.Equal(t, accessToken, *cfg.Auth.AccessToken)
}

func TestWithTopLevelDomain(t *testing.T) {
	cfg := config.NewConfiguration()
	tld := "com"
	cfg = cfg.WithTopLevelDomain(tld)

	assert.NotNil(t, cfg.Endpoint.TopLevelDomain)
	assert.Equal(t, tld, *cfg.Endpoint.TopLevelDomain)
}

func TestWithRootDomain(t *testing.T) {
	cfg := config.NewConfiguration()
	rootDomain := "pingone.com"
	cfg = cfg.WithRootDomain(rootDomain)

	assert.NotNil(t, cfg.Endpoint.RootDomain)
	assert.Equal(t, rootDomain, *cfg.Endpoint.RootDomain)
}

func TestWithAPIDomain(t *testing.T) {
	cfg := config.NewConfiguration()
	apiDomain := "api.pingone.com"
	cfg = cfg.WithAPIDomain(apiDomain)

	assert.NotNil(t, cfg.Endpoint.APIDomain)
	assert.Equal(t, apiDomain, *cfg.Endpoint.APIDomain)
}

func TestWithCustomDomain(t *testing.T) {
	cfg := config.NewConfiguration()
	customDomain := "custom.pingone.com"
	cfg = cfg.WithCustomDomain(customDomain)

	assert.NotNil(t, cfg.Endpoint.CustomDomain)
	assert.Equal(t, customDomain, *cfg.Endpoint.CustomDomain)
}

func TestHasBearerToken(t *testing.T) {
	tests := []struct {
		name        string
		accessToken *string
		expected    bool
	}{
		{
			name:        "Has token",
			accessToken: stringPtr("test-token"),
			expected:    true,
		},
		{
			name:        "Empty token",
			accessToken: stringPtr(""),
			expected:    false,
		},
		{
			name:        "Nil token",
			accessToken: nil,
			expected:    false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := config.NewConfiguration()
			if tc.accessToken != nil {
				cfg = cfg.WithAccessToken(*tc.accessToken)
			}
			assert.Equal(t, tc.expected, cfg.HasBearerToken())
		})
	}
}

func TestAddBearerTokenToContext(t *testing.T) {
	tests := []struct {
		name        string
		accessToken *string
		shouldAdd   bool
	}{
		{
			name:        "Has token",
			accessToken: stringPtr("test-token"),
			shouldAdd:   true,
		},
		{
			name:        "Empty token",
			accessToken: stringPtr(""),
			shouldAdd:   false,
		},
		{
			name:        "Nil token",
			accessToken: nil,
			shouldAdd:   false,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := config.NewConfiguration()
			if tc.accessToken != nil {
				cfg = cfg.WithAccessToken(*tc.accessToken)
			}

			key := "token-key"
			ctx := context.Background()
			newCtx := cfg.AddBearerTokenToContext(ctx, key)

			if tc.shouldAdd {
				token, ok := newCtx.Value(key).(string)
				assert.True(t, ok, "Context should contain token")
				assert.Equal(t, *tc.accessToken, token)
			} else {
				assert.Equal(t, ctx, newCtx, "Context should not be changed")
			}
		})
	}
}

func TestBearerToken(t *testing.T) {
	accessToken := "test-token"
	cfg := config.NewConfiguration().WithAccessToken(accessToken)

	token := cfg.BearerToken()
	assert.NotNil(t, token)
	assert.Equal(t, accessToken, token.AccessToken)
	assert.Equal(t, "Bearer", token.TokenType)
}

func TestAuthEndpoints(t *testing.T) {
	tests := []struct {
		name         string
		setupConfig  func(*config.Configuration) *config.Configuration
		expectError  bool
		validateURLs func(*testing.T, interface{})
	}{
		{
			name: "With custom domain",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithCustomDomain("custom.pingone.com")
			},
			expectError: false,
			validateURLs: func(t *testing.T, endpoints interface{}) {
				e, ok := endpoints.(struct{ TokenURL string })
				require.True(t, ok, "endpoints should be of expected type")
				assert.Contains(t, e.TokenURL, "https://custom.pingone.com")
			},
		},
		{
			name: "With environment ID and root domain",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithAuthEnvironmentID("env-id").WithRootDomain("pingone.com")
			},
			expectError: false,
			validateURLs: func(t *testing.T, endpoints interface{}) {
				e, ok := endpoints.(struct{ TokenURL string })
				require.True(t, ok, "endpoints should be of expected type")
				assert.Contains(t, e.TokenURL, "https://auth.pingone.com/env-id")
			},
		},
		{
			name: "With environment ID and top level domain",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithAuthEnvironmentID("env-id").WithTopLevelDomain("com")
			},
			expectError: false,
			validateURLs: func(t *testing.T, endpoints interface{}) {
				e, ok := endpoints.(struct{ TokenURL string })
				require.True(t, ok, "endpoints should be of expected type")
				assert.Contains(t, e.TokenURL, "https://auth.pingone.com/env-id")
			},
		},
		{
			name: "Missing configuration",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := tc.setupConfig(config.NewConfiguration())

			endpoints, err := cfg.AuthEndpoints()

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				tc.validateURLs(t, struct{ TokenURL string }{TokenURL: endpoints.TokenURL})
			}
		})
	}
}

func TestAPIDomain(t *testing.T) {
	tests := []struct {
		name        string
		setupConfig func(*config.Configuration) *config.Configuration
		expectError bool
		expected    string
	}{
		{
			name: "With API domain",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithAPIDomain("api.pingone.com")
			},
			expectError: false,
			expected:    "api.pingone.com",
		},
		{
			name: "With root domain",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithRootDomain("pingone.com")
			},
			expectError: false,
			expected:    "api.pingone.com",
		},
		{
			name: "With top level domain",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithTopLevelDomain("com")
			},
			expectError: false,
			expected:    "api.pingone.com",
		},
		{
			name: "Missing configuration",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg
			},
			expectError: true,
			expected:    "",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := tc.setupConfig(config.NewConfiguration())

			domain, err := cfg.APIDomain()

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expected, domain)
			}
		})
	}
}

func TestTokenSource(t *testing.T) {
	tests := []struct {
		name        string
		setupConfig func(*config.Configuration) *config.Configuration
		expectError bool
	}{
		{
			name: "With access token",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithAccessToken("test-token").
					WithCustomDomain("custom.pingone.com")
			},
			expectError: false,
		},
		{
			name: "With client credentials and client secret",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithGrantType(oauth2.GrantTypeClientCredentials).
					WithClientID("client-id").
					WithClientSecret("client-secret").
					WithCustomDomain("custom.pingone.com")
			},
			expectError: false,
		},
		{
			name: "With client credentials missing client ID",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithGrantType(oauth2.GrantTypeClientCredentials).
					WithCustomDomain("custom.pingone.com")
			},
			expectError: true,
		},
		{
			name: "With client credentials missing client secret",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithGrantType(oauth2.GrantTypeClientCredentials).
					WithClientID("client-id").
					WithCustomDomain("custom.pingone.com")
			},
			expectError: true,
		},
		{
			name: "Missing valid endpoints",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithGrantType(oauth2.GrantTypeClientCredentials).
					WithClientID("client-id").
					WithClientSecret("client-secret")
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := tc.setupConfig(config.NewConfiguration())

			ts, err := cfg.TokenSource(context.Background())

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, ts)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, ts)
			}
		})
	}
}

func TestClient(t *testing.T) {
	tests := []struct {
		name        string
		setupConfig func(*config.Configuration) *config.Configuration
		expectError bool
	}{
		{
			name: "Valid configuration with token",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg.WithAccessToken("test-token").
					WithCustomDomain("custom.pingone.com")
			},
			expectError: false,
		},
		{
			name: "Invalid configuration",
			setupConfig: func(cfg *config.Configuration) *config.Configuration {
				return cfg
			},
			expectError: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cfg := tc.setupConfig(config.NewConfiguration())

			client, err := cfg.Client(context.Background(), http.DefaultClient)

			if tc.expectError {
				assert.Error(t, err)
				assert.Nil(t, client)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, client)
			}
		})
	}
}

// Helper function to create a pointer to a string
func stringPtr(s string) *string {
	return &s
}
