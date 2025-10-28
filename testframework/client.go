// Copyright © 2025 Ping Identity Corporation

// Package testframework provides testing utilities and helpers for the PingOne Go Client SDK.
// It includes client configuration helpers, environment management utilities, and common
// test patterns for integration testing with PingOne services.
package testframework

import (
	"github.com/pingidentity/pingone-go-client/config"
	"github.com/pingidentity/pingone-go-client/pingone"
)

// TestClient creates a new PingOne API client configured for testing purposes.
// The svcConfig parameter must be a valid Configuration with authentication credentials
// and endpoint settings. The client is automatically configured with a "testing" user agent
// to help identify test traffic in logs. It returns a configured API client or an error
// if client creation fails.
func TestClient(svcConfig *config.Configuration) (*pingone.APIClient, error) {
	configuration := pingone.NewConfiguration(svcConfig)
	configuration.AppendUserAgent("testing")
	return pingone.NewAPIClient(configuration)
}
