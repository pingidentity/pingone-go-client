package config

import (
	"context"
	"fmt"

	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"golang.org/x/oauth2"
)

func (d *DeviceCode) DeviceAuthTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (oauth2.TokenSource, error) {
	if d.DeviceCodeClientID == nil || *d.DeviceCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for device code grant type")
	}

	config := &oauth2.Config{
		ClientID: *d.DeviceCodeClientID,
		Scopes:   *d.DeviceCodeScopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:       endpoints.AuthorizationURLPath,
			DeviceAuthURL: endpoints.DeviceAuthorizationURLPath,
			TokenURL:      endpoints.TokenURLPath,
		},
	}

	response, err := config.DeviceAuth(ctx)
	if err != nil {
		return nil, fmt.Errorf("device auth request failed: %w", err)
	}

	fmt.Printf("\nDevice Authorization Required:\n")
	fmt.Printf("1. Open this URL in your browser: %s\n", response.VerificationURI)
	fmt.Printf("2. Enter this code: %s\n", response.UserCode)
	if response.VerificationURIComplete != "" {
		fmt.Printf("   (Or visit this complete URL: %s)\n", response.VerificationURIComplete)
	}
	fmt.Printf("\nWaiting for authorization...\n")

	deviceCodeToken, err := config.DeviceAccessToken(ctx, response)
	if err != nil {
		return nil, fmt.Errorf("device access token request failed: %w", err)
	}

	ts := config.TokenSource(ctx, deviceCodeToken)

	return ts, nil
}
