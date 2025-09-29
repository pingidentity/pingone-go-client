package config

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *ClientCredentials) ClientCredentialsTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (*oauth2.TokenSource, error) {
	if c.ClientCredentialsClientID == nil || *c.ClientCredentialsClientID == "" {
		return nil, fmt.Errorf("client ID is required for client credentials grant type")
	}

	slog.Debug("Using client credentials token source with provided client ID", "client ID", *c.ClientCredentialsClientID)
	if c.ClientCredentialsClientSecret != nil && *c.ClientCredentialsClientSecret != "" {
		config := &clientcredentials.Config{
			ClientID:     *c.ClientCredentialsClientID,
			ClientSecret: *c.ClientCredentialsClientSecret,
			TokenURL:     endpoints.TokenURL,
		}
		ts := config.TokenSource(ctx)
		slog.Debug("Using standard client credentials token source as client secret has been provided")
		return &ts, nil
	}

	// tmp until private keys are supported
	return nil, fmt.Errorf("client secret is required for client credentials grant type")

}
