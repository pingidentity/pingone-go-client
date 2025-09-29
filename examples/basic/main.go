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

func main() {
	// Configure logging
	opts := &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
	slog.SetDefault(logger)

	// Initialize service configuration
	serviceCfg := config.NewConfiguration()
	serviceCfg.WithGrantType(oauth2.GrantTypeClientCredentials)

	// These values should be set in environment variables or passed directly
	// PINGONE_CLIENT_ID=your-client-id
	// PINGONE_CLIENT_SECRET=your-client-secret
	// PINGONE_ENVIRONMENT_ID=your-environment-id
	clientID := os.Getenv("PINGONE_CLIENT_ID")
	clientSecret := os.Getenv("PINGONE_CLIENT_SECRET")
	authEnvironmentID := os.Getenv("PINGONE_ENVIRONMENT_ID")

	if clientID == "" || clientSecret == "" || authEnvironmentID == "" {
		slog.Error("Missing required environment variables",
			"PINGONE_CLIENT_ID", clientID != "",
			"PINGONE_CLIENT_SECRET", clientSecret != "",
			"PINGONE_ENVIRONMENT_ID", authEnvironmentID != "")
		os.Exit(1)
	}

	serviceCfg.WithClientID(clientID)
	serviceCfg.WithClientSecret(clientSecret)
	serviceCfg.WithEnvironmentID(authEnvironmentID)
	serviceCfg.WithRootDomain("pingone.com")

	// Create client configuration
	configuration := pingone.NewConfiguration(serviceCfg)

	// Add a custom user agent suffix (optional)
	configuration.AppendUserAgent("example-application")

	// Create the API client
	client, err := pingone.NewAPIClient(configuration)
	if err != nil {
		slog.Error("Failed to create client", "error", err)
		os.Exit(1)
	}

	// Use the client to access PingOne APIs
	// For example, to get environment details:
	environmentID, err := uuid.Parse(os.Getenv("PINGONE_MY_USER_ENVIRONMENT_ID"))
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
