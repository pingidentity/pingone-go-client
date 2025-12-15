// Copyright © 2025 Ping Identity Corporation

package config_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/stretchr/testify/require"
)

func TestNewConfiguration(t *testing.T) {
	cfg := config.NewConfiguration()
	if cfg == nil {
		t.Fatal("Configuration should not be nil")
	}
}

func TestConfigurationWithMethods(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *config.Configuration
		validate func(t *testing.T, cfg *config.Configuration)
	}{
		{
			name: "WithEnvironmentID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithEnvironmentID("test-environment-id")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Endpoint.EnvironmentID == nil {
					t.Fatal("EnvironmentID should not be nil")
				}
				if *cfg.Endpoint.EnvironmentID != "test-environment-id" {
					t.Errorf("Expected EnvironmentID to be 'test-environment-id', got %q", *cfg.Endpoint.EnvironmentID)
				}
			},
		},
		{
			name: "WithClientID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithClientID("test-client-id")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Auth.ClientCredentials.ClientCredentialsClientID == nil {
					t.Fatal("ClientCredentialsClientID should not be nil")
				}
				if *cfg.Auth.ClientCredentials.ClientCredentialsClientID != "test-client-id" {
					t.Errorf("Expected ClientID to be 'test-client-id', got %q", *cfg.Auth.ClientCredentials.ClientCredentialsClientID)
				}
			},
		},
		{
			name: "WithGrantType",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithGrantType(oauth2.GrantTypeClientCredentials)
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Auth.GrantType == nil {
					t.Fatal("GrantType should not be nil")
				}
				if *cfg.Auth.GrantType != oauth2.GrantTypeClientCredentials {
					t.Errorf("Expected GrantType to be %q, got %q", oauth2.GrantTypeClientCredentials, *cfg.Auth.GrantType)
				}
			},
		},
		{
			name: "WithClientSecret",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithClientSecret("test-client-secret")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Auth.ClientCredentials.ClientCredentialsClientSecret == nil {
					t.Fatal("ClientCredentialsClientSecret should not be nil")
				}
				if *cfg.Auth.ClientCredentials.ClientCredentialsClientSecret != "test-client-secret" {
					t.Errorf("Expected ClientSecret to be 'test-client-secret', got %q", *cfg.Auth.ClientCredentials.ClientCredentialsClientSecret)
				}
			},
		},
		{
			name: "WithAccessToken",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAccessToken("test-access-token")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Auth.AccessToken == nil {
					t.Fatal("AccessToken should not be nil")
				}
				if *cfg.Auth.AccessToken != "test-access-token" {
					t.Errorf("Expected AccessToken to be 'test-access-token', got %q", *cfg.Auth.AccessToken)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			tt.validate(t, cfg)
		})
	}
}

func TestConfigurationDomainMethods(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *config.Configuration
		validate func(t *testing.T, cfg *config.Configuration)
	}{
		{
			name: "WithTopLevelDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithTopLevelDomain(config.TopLevelDomain("com"))
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				expected := config.TopLevelDomain("com")
				if cfg.Endpoint.TopLevelDomain == nil {
					t.Fatal("TopLevelDomain should not be nil")
				}
				if *cfg.Endpoint.TopLevelDomain != expected {
					t.Errorf("Expected TopLevelDomain to be %q, got %q", expected, *cfg.Endpoint.TopLevelDomain)
				}
			},
		},
		{
			name: "WithRootDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithRootDomain("pingone.com")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Endpoint.RootDomain == nil {
					t.Fatal("RootDomain should not be nil")
				}
				if *cfg.Endpoint.RootDomain != "pingone.com" {
					t.Errorf("Expected RootDomain to be 'pingone.com', got %q", *cfg.Endpoint.RootDomain)
				}
			},
		},
		{
			name: "WithAPIDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAPIDomain("api.pingone.com")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Endpoint.APIDomain == nil {
					t.Fatal("APIDomain should not be nil")
				}
				if *cfg.Endpoint.APIDomain != "api.pingone.com" {
					t.Errorf("Expected APIDomain to be 'api.pingone.com', got %q", *cfg.Endpoint.APIDomain)
				}
			},
		},
		{
			name: "WithCustomDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithCustomDomain("custom.pingone.com")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Endpoint.CustomDomain == nil {
					t.Fatal("CustomDomain should not be nil")
				}
				if *cfg.Endpoint.CustomDomain != "custom.pingone.com" {
					t.Errorf("Expected CustomDomain to be 'custom.pingone.com', got %q", *cfg.Endpoint.CustomDomain)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			tt.validate(t, cfg)
		})
	}
}

func TestConfigurationAuthMethods(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *config.Configuration
		validate func(t *testing.T, cfg *config.Configuration)
	}{
		{
			name: "WithClientCredentialsScopes",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithClientCredentialsScopes([]string{"scope1", "scope2"})
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				expected := []string{"scope1", "scope2"}
				if cfg.Auth.ClientCredentials.ClientCredentialsScopes == nil {
					t.Fatal("ClientCredentialsScopes should not be nil")
				}
				actual := *cfg.Auth.ClientCredentials.ClientCredentialsScopes
				if len(actual) != len(expected) {
					t.Errorf("Expected scopes length %d, got %d", len(expected), len(actual))
				}
				for i, scope := range expected {
					if i >= len(actual) || actual[i] != scope {
						t.Errorf("Expected scope[%d] to be %q, got %q", i, scope, actual[i])
					}
				}
			},
		},
		{
			name: "WithAuthorizationCodeClientID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAuthorizationCodeClientID("auth-code-client-id")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Auth.AuthorizationCode.AuthorizationCodeClientID == nil {
					t.Fatal("AuthorizationCodeClientID should not be nil")
				}
				if *cfg.Auth.AuthorizationCode.AuthorizationCodeClientID != "auth-code-client-id" {
					t.Errorf("Expected AuthorizationCodeClientID to be 'auth-code-client-id', got %q", *cfg.Auth.AuthorizationCode.AuthorizationCodeClientID)
				}
			},
		},
		{
			name: "WithAuthorizationCodeScopes",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAuthorizationCodeScopes([]string{"openid", "profile"})
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				expected := []string{"openid", "profile"}
				if cfg.Auth.AuthorizationCode.AuthorizationCodeScopes == nil {
					t.Fatal("AuthorizationCodeScopes should not be nil")
				}
				actual := *cfg.Auth.AuthorizationCode.AuthorizationCodeScopes
				if len(actual) != len(expected) {
					t.Errorf("Expected scopes length %d, got %d", len(expected), len(actual))
				}
				for i, scope := range expected {
					if i >= len(actual) || actual[i] != scope {
						t.Errorf("Expected scope[%d] to be %q, got %q", i, scope, actual[i])
					}
				}
			},
		},
		{
			name: "WithAuthorizationCodeRedirectURI",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAuthorizationCodeRedirectURI(config.AuthorizationCodeRedirectURI{
					Port: "9090",
					Path: "/callback",
				})
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Auth.AuthorizationCode.AuthorizationCodeRedirectURI.Port != "9090" {
					t.Errorf("Expected AuthorizationCodeRedirectURI.Port to be 9090, got %s", cfg.Auth.AuthorizationCode.AuthorizationCodeRedirectURI.Port)
				}
				if cfg.Auth.AuthorizationCode.AuthorizationCodeRedirectURI.Path != "/callback" {
					t.Errorf("Expected AuthorizationCodeRedirectURI.Path to be '/callback', got %q", cfg.Auth.AuthorizationCode.AuthorizationCodeRedirectURI.Path)
				}
			},
		},
		{
			name: "WithDeviceCodeClientID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithDeviceCodeClientID("device-code-client-id")
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				if cfg.Auth.DeviceCode.DeviceCodeClientID == nil {
					t.Fatal("DeviceCodeClientID should not be nil")
				}
				if *cfg.Auth.DeviceCode.DeviceCodeClientID != "device-code-client-id" {
					t.Errorf("Expected DeviceCodeClientID to be 'device-code-client-id', got %q", *cfg.Auth.DeviceCode.DeviceCodeClientID)
				}
			},
		},
		{
			name: "WithDeviceCodeScopes",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithDeviceCodeScopes([]string{"openid"})
			},
			validate: func(t *testing.T, cfg *config.Configuration) {
				expected := []string{"openid"}
				if cfg.Auth.DeviceCode.DeviceCodeScopes == nil {
					t.Fatal("DeviceCodeScopes should not be nil")
				}
				actual := *cfg.Auth.DeviceCode.DeviceCodeScopes
				if len(actual) != len(expected) {
					t.Errorf("Expected scopes length %d, got %d", len(expected), len(actual))
				}
				for i, scope := range expected {
					if i >= len(actual) || actual[i] != scope {
						t.Errorf("Expected scope[%d] to be %q, got %q", i, scope, actual[i])
					}
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			tt.validate(t, cfg)
		})
	}
}

func TestHasBearerToken(t *testing.T) {
	tests := []struct {
		name     string
		setup    func() *config.Configuration
		expected bool
	}{
		{
			name: "HasToken",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAccessToken("test-token")
			},
			expected: true,
		},
		{
			name: "EmptyToken",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAccessToken("")
			},
			expected: false,
		},
		{
			name: "NilToken",
			setup: func() *config.Configuration {
				return config.NewConfiguration()
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			result := cfg.HasBearerToken()
			if result != tt.expected {
				t.Errorf("Expected HasBearerToken() to return %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestAddBearerTokenToContext(t *testing.T) {
	tests := []struct {
		name           string
		setup          func() *config.Configuration
		shouldAddToken bool
		expectedToken  string
	}{
		{
			name: "HasToken",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAccessToken("test-token")
			},
			shouldAddToken: true,
			expectedToken:  "test-token",
		},
		{
			name: "EmptyToken",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAccessToken("")
			},
			shouldAddToken: false,
		},
		{
			name: "NilToken",
			setup: func() *config.Configuration {
				return config.NewConfiguration()
			},
			shouldAddToken: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			key := "token-key"
			ctx := context.Background()
			newCtx := cfg.AddBearerTokenToContext(ctx, key)

			if tt.shouldAddToken {
				token, ok := newCtx.Value(key).(string)
				if !ok {
					t.Fatal("Context should contain token")
				}
				if token != tt.expectedToken {
					t.Errorf("Expected token to be %q, got %q", tt.expectedToken, token)
				}
			} else {
				if newCtx != ctx {
					t.Error("Context should not be changed when no token is present")
				}
			}
		})
	}
}

func TestBearerToken(t *testing.T) {
	accessToken := "test-token"
	cfg := config.NewConfiguration().WithAccessToken(accessToken)

	token := cfg.BearerToken()
	require.NotNil(t, token, "Token should not be nil")
	if token.AccessToken != accessToken {
		t.Errorf("Expected AccessToken to be %q, got %q", accessToken, token.AccessToken)
	}
	if token.TokenType != "Bearer" {
		t.Errorf("Expected TokenType to be 'Bearer', got %q", token.TokenType)
	}
}

func TestAuthEndpoints(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *config.Configuration
		expectError bool
		urlContains string
	}{
		{
			name: "WithCustomDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithCustomDomain("custom.pingone.com")
			},
			expectError: false,
			urlContains: "https://custom.pingone.com",
		},
		{
			name: "WithEnvironmentIDAndRootDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithEnvironmentID("env-id").WithRootDomain("pingone.com")
			},
			expectError: false,
			urlContains: "https://auth.pingone.com/env-id",
		},
		{
			name: "WithEnvironmentIDAndTopLevelDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithEnvironmentID("env-id").WithTopLevelDomain("com")
			},
			expectError: false,
			urlContains: "https://auth.pingone.com/env-id",
		},
		{
			name: "MissingConfiguration",
			setup: func() *config.Configuration {
				return config.NewConfiguration()
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			endpoints, err := cfg.AuthEndpoints()

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if tt.urlContains != "" {
				if !containsString(endpoints.TokenURL, tt.urlContains) {
					t.Errorf("Expected TokenURL to contain %q, got %q", tt.urlContains, endpoints.TokenURL)
				}
			}
		})
	}
}

// Helper function to check if a string contains a substring
func containsString(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || (len(substr) > 0 && stringContains(s, substr)))
}

func stringContains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

func TestAuthEndpointsGrantTypeSpecific(t *testing.T) {
	tests := []struct {
		name  string
		setup func() *config.Configuration
	}{
		{
			name: "WithEnvironmentID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().
					WithGrantType(oauth2.GrantTypeClientCredentials).
					WithEnvironmentID("test-env-id").
					WithTopLevelDomain("com")
			},
		},
		{
			name: "AuthorizationCodeWithEnvironmentID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().
					WithGrantType(oauth2.GrantTypeAuthorizationCode).
					WithEnvironmentID("test-env-id").
					WithTopLevelDomain("com")
			},
		},
		{
			name: "DeviceCodeWithEnvironmentID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().
					WithGrantType(oauth2.GrantTypeDeviceCode).
					WithEnvironmentID("test-env-id").
					WithTopLevelDomain("com")
			},
		},
		{
			name: "SharedEnvironmentID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().
					WithGrantType(oauth2.GrantTypeClientCredentials).
					WithEnvironmentID("shared-env-id").
					WithTopLevelDomain("com")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			endpoints, err := cfg.AuthEndpoints()

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if endpoints.TokenURL == "" {
				t.Error("Expected TokenURL to not be empty")
			}
		})
	}
}

func TestAPIDomain(t *testing.T) {
	tests := []struct {
		name           string
		setup          func() *config.Configuration
		expectedDomain string
		expectError    bool
	}{
		{
			name: "WithAPIDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithAPIDomain("api.pingone.com")
			},
			expectedDomain: "api.pingone.com",
			expectError:    false,
		},
		{
			name: "WithRootDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithRootDomain("pingone.com")
			},
			expectedDomain: "api.pingone.com",
			expectError:    false,
		},
		{
			name: "WithTopLevelDomain",
			setup: func() *config.Configuration {
				return config.NewConfiguration().WithTopLevelDomain("com")
			},
			expectedDomain: "api.pingone.com",
			expectError:    false,
		},
		{
			name: "MissingConfiguration",
			setup: func() *config.Configuration {
				return config.NewConfiguration()
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			domain, err := cfg.APIDomain()

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				return
			}

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
				return
			}

			if domain != tt.expectedDomain {
				t.Errorf("Expected domain to be %q, got %q", tt.expectedDomain, domain)
			}
		})
	}
}

func TestTokenSource(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *config.Configuration
		expectError bool
	}{
		{
			name: "WithAccessToken",
			setup: func() *config.Configuration {
				return config.NewConfiguration().
					WithAccessToken("test-token").
					WithCustomDomain("custom.pingone.com")
			},
			expectError: false,
		},
		{
			name: "WithClientCredentials",
			setup: func() *config.Configuration {
				return config.NewConfiguration().
					WithGrantType(oauth2.GrantTypeClientCredentials).
					WithEnvironmentID("test-env-id").
					WithClientID("client-id").
					WithClientSecret("client-secret").
					WithClientCredentialsScopes([]string{"test"}).
					WithTopLevelDomain("com")
			},
			expectError: true, // This will fail because it tries to make a real network request
		},
		{
			name: "MissingClientID",
			setup: func() *config.Configuration {
				return config.NewConfiguration().
					WithGrantType(oauth2.GrantTypeClientCredentials).
					WithCustomDomain("custom.pingone.com")
			},
			expectError: true,
		},
		{
			name: "InvalidConfiguration",
			setup: func() *config.Configuration {
				return config.NewConfiguration()
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			ts, err := cfg.TokenSource(context.Background())

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if ts != nil {
					t.Error("Expected TokenSource to be nil when error occurs")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if ts == nil {
					t.Error("Expected TokenSource to not be nil")
				}
			}
		})
	}
}

func TestClient(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *config.Configuration
		expectError bool
	}{
		{
			name: "ValidConfiguration",
			setup: func() *config.Configuration {
				return config.NewConfiguration().
					WithAccessToken("test-token").
					WithCustomDomain("custom.pingone.com")
			},
			expectError: false,
		},
		{
			name: "InvalidConfiguration",
			setup: func() *config.Configuration {
				return config.NewConfiguration()
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.setup()
			client, err := cfg.Client(context.Background(), http.DefaultClient)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error but got none")
				}
				if client != nil {
					t.Error("Expected client to be nil when error occurs")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if client == nil {
					t.Error("Expected client to not be nil")
				}
			}
		})
	}
}

func TestWithStorageType(t *testing.T) {
	tests := []struct {
		name        string
		storageType config.StorageType
	}{
		{
			name:        "StorageTypeSecureLocal",
			storageType: config.StorageTypeSecureLocal,
		},
		{
			name:        "StorageTypeFileSystem",
			storageType: config.StorageTypeFileSystem,
		},
		{
			name:        "StorageTypeSecureRemote",
			storageType: config.StorageTypeSecureRemote,
		},
		{
			name:        "StorageTypeNone",
			storageType: config.StorageTypeNone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewConfiguration().WithStorageType(tt.storageType)

			if cfg.Auth.Storage == nil {
				t.Fatal("Storage should not be nil")
			}

			if cfg.Auth.Storage.Type != tt.storageType {
				t.Errorf("Expected StorageType to be %q, got %q", tt.storageType, cfg.Auth.Storage.Type)
			}
		})
	}
}

func TestWithUseKeychain(t *testing.T) {
	tests := []struct {
		name              string
		useKeychain       bool
		expectStorageType config.StorageType
	}{
		{
			name:              "UseKeychain true sets StorageTypeSecureLocal",
			useKeychain:       true,
			expectStorageType: config.StorageTypeSecureLocal,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := config.NewConfiguration().WithUseKeychain(tt.useKeychain)

			if cfg.Auth.Storage == nil {
				t.Fatal("Storage should not be nil")
			}

			if cfg.Auth.Storage.Type != tt.expectStorageType {
				t.Errorf("Expected StorageType to be %q, got %q", tt.expectStorageType, cfg.Auth.Storage.Type)
			}
		})
	}
}

func TestWithStorageName(t *testing.T) {
	cfg := config.NewConfiguration().WithStorageName("test-storage-name")

	if cfg.Auth.Storage == nil {
		t.Fatal("Storage should not be nil")
	}

	if cfg.Auth.Storage.KeychainName != "test-storage-name" {
		t.Errorf("Expected Storage.KeychainName to be 'test-storage-name', got %q", cfg.Auth.Storage.KeychainName)
	}
}
