package config

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"golang.org/x/oauth2"
)

func (d *DeviceCode) DeviceAuthTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (*oauth2.TokenSource, error) {
	if d.DeviceCodeClientID == nil || *d.DeviceCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for device code grant type")
	}

	slog.Debug("Using device code token source with provided client ID", "client ID", *d.DeviceCodeClientID)
	config := &oauth2.Config{
		ClientID: *d.DeviceCodeClientID,
		Scopes:   *d.DeviceCodeScopes,
	}

	response, err := config.DeviceAuth(ctx)
	if err != nil {
		return nil, fmt.Errorf("device auth request failed: %w", err)
	}

	fmt.Printf("Visit %s and enter the code: %s\n", response.VerificationURIComplete, response.UserCode)
	deviceCodeToken, err := config.DeviceAccessToken(ctx, response)
	if err != nil {
		return nil, fmt.Errorf("device access token request failed: %w", err)
	}

	ts := config.TokenSource(ctx, deviceCodeToken)
	slog.Debug("Using standard device code token source as client secret has been provided")
	return &ts, nil
}
