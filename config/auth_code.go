package config

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"golang.org/x/oauth2"
)

func (a *AuthCode) AuthCodeTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (*oauth2.TokenSource, error) {
	if a.AuthCodeClientID == nil || *a.AuthCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for client credentials grant type")
	}

	slog.Debug("Using client credentials token source with provided client ID", "client ID", *a.AuthCodeClientID)

	tokenEndpoint := &oauth2.Endpoint{
		AuthURL:  endpoints.AuthURL,
		TokenURL: endpoints.TokenURL,
	}

	config := &oauth2.Config{
		ClientID: *a.AuthCodeClientID,
		Endpoint: *tokenEndpoint,
		Scopes:   *a.AuthCodeScopes,
	}

	codeVerifier := oauth2.GenerateVerifier()

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(codeVerifier))
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect
	// URL. Exchange will do the handshake to retrieve the
	// initial access token. The HTTP Client returned by
	// conf.Client will refresh the token as necessary.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := config.Exchange(ctx, code, oauth2.VerifierOption(codeVerifier))
	if err != nil {
		log.Fatal(err)
	}

	ts := config.TokenSource(ctx, tok)
	slog.Debug("Using standard client credentials token source as client secret has been provided")
	return &ts, nil

}
