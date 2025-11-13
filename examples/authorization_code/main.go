// Copyright © 2025 Ping Identity Corporation

// Package main demonstrates authorization code flow authentication using the PingOne Go Client SDK.
// This example shows how to configure interactive user authentication via browser for web
// and native applications.
package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/google/uuid"
	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/oauth2"
	"github.com/pingidentity/pingone-go-client/pingone"
)

// main demonstrates authorization code flow for interactive user authentication.
// This example requires the following environment variables to be set:
// - PINGONE_CLIENT_ID: OAuth2 client ID from your PingOne Web/Native application
// - PINGONE_ENVIRONMENT_ID: Environment ID for authentication and API calls
// - PINGONE_ROOT_DOMAIN: (optional) PingOne domain, defaults to "pingone.com"
// - PINGONE_REDIRECT_URI_PORT: (optional) Redirect URI port, defaults to "8080"
// - PINGONE_REDIRECT_URI_PATH: (optional) Redirect URI path, defaults to "/callback"
//
// IMPORTANT: The redirect URI (http://localhost:PORT/PATH) must be configured in your PingOne application settings.
func main() {
	// Configure logging
	opts := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	// Initialize service configuration for authorization code flow
	serviceCfg := config.NewConfiguration()
	serviceCfg.WithGrantType(oauth2.GrantTypeAuthorizationCode)

	// Get configuration from environment variables
	clientID := os.Getenv("PINGONE_CLIENT_ID")
	authEnvironmentID := os.Getenv("PINGONE_ENVIRONMENT_ID")
	rootDomain := os.Getenv("PINGONE_ROOT_DOMAIN")
	redirectURIPort := os.Getenv("PINGONE_REDIRECT_URI_PORT")
	redirectURIPath := os.Getenv("PINGONE_REDIRECT_URI_PATH")

	if rootDomain == "" {
		rootDomain = "pingone.com"
	}
	if redirectURIPort == "" {
		redirectURIPort = config.GetDefaultAuthorizationCodeRedirectURIPort() // "8080"
	}
	if redirectURIPath == "" {
		redirectURIPath = config.GetDefaultAuthorizationCodeRedirectURIPath() // "/callback"
	}

	if clientID == "" || authEnvironmentID == "" {
		slog.Error("Missing required environment variables",
			"PINGONE_CLIENT_ID", clientID != "",
			"PINGONE_ENVIRONMENT_ID", authEnvironmentID != "")
		os.Exit(1)
	}

	serviceCfg.WithAuthorizationCodeClientID(clientID)
	serviceCfg.WithEnvironmentID(authEnvironmentID)
	serviceCfg.WithRootDomain(rootDomain)

	// Configure authorization code redirect URI
	redirectURI := config.AuthorizationCodeRedirectURI{
		Port: redirectURIPort,
		Path: redirectURIPath,
	}
	serviceCfg.WithAuthorizationCodeRedirectURI(redirectURI)
	serviceCfg.WithAuthorizationCodeScopes([]string{"openid", "profile"})

	// Create client configuration
	configuration := pingone.NewConfiguration(serviceCfg)

	// Add a custom user agent suffix (optional)
	configuration.AppendUserAgent("authorization-code-example")

	// Create the API client
	client, err := pingone.NewAPIClient(configuration)
	if err != nil {
		slog.Error("Failed to create client", "error", err)
		os.Exit(1)
	}

	slog.Info("PingOne API Client initialized successfully with authorization code flow")
	slog.Info("Note: The SDK will open a browser for user authentication")

	// Use the client to access PingOne APIs
	// For example, to get environment details:
	environmentID, err := uuid.Parse(authEnvironmentID)
	if err != nil {
		slog.Error("Invalid environment ID format", "error", err)
		os.Exit(1)
	}

	environment, httpResp, err := client.EnvironmentsApi.GetEnvironmentById(
		context.Background(),
		environmentID,
	).Execute()

	if err != nil {
		slog.Error("Failed to read environment", "error", err)
		os.Exit(1)
	}

	if httpResp.StatusCode != 200 {
		slog.Error("Failed to read environment", "http response status", httpResp.Status)
		os.Exit(1)
	}

	slog.Info("Successfully read environment",
		"id", environment.Id.String(),
		"name", environment.Name,
		"type", environment.Type,
		"region", environment.Region,
	)
}
