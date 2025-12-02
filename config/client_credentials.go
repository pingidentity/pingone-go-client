// Copyright © 2025 Ping Identity Corporation

package config

import (
	"context"
	"fmt"

	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func (c *ClientCredentials) ClientCredentialsTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (oauth2.TokenSource, error) {
	if c.ClientCredentialsClientID == nil || *c.ClientCredentialsClientID == "" {
		return nil, fmt.Errorf("client ID is required for client credentials grant type")
	}

	if c.ClientCredentialsClientSecret == nil || *c.ClientCredentialsClientSecret == "" {
		return nil, fmt.Errorf("client secret is required for client credentials grant type")
	}

	config := &clientcredentials.Config{
		ClientID:     *c.ClientCredentialsClientID,
		ClientSecret: *c.ClientCredentialsClientSecret,
		TokenURL:     endpoints.TokenURL,
		Scopes:       *c.ClientCredentialsScopes,
		AuthStyle:    oauth2.AuthStyleInHeader,
	}

	ts := config.TokenSource(ctx)
	return ts, nil
}
