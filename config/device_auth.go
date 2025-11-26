// Copyright © 2025 Ping Identity Corporation
package config

import (
	"context"
	"fmt"

	"github.com/pingidentity/pingone-go-client/utils/browser"
	"golang.org/x/oauth2"
)

func (d *DeviceCode) DeviceAuthTokenSource(ctx context.Context, endpoints oauth2.Endpoint) (oauth2.TokenSource, error) {
	if d.DeviceCodeClientID == nil || *d.DeviceCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for device code grant type")
	}

	config := &oauth2.Config{
		ClientID: *d.DeviceCodeClientID,
		Scopes:   *d.DeviceCodeScopes,
		Endpoint: endpoints,
	}

	response, err := config.DeviceAuth(ctx)
	if err != nil {
		return nil, fmt.Errorf("device auth request failed: %w", err)
	}

	// Check if custom prompt handler is provided
	if d.OnDisplayPrompt != nil {
		// Use custom handler - delegate UX to consumer
		if err := d.OnDisplayPrompt(response.VerificationURI, response.UserCode); err != nil {
			return nil, fmt.Errorf("custom prompt handler failed: %w", err)
		}
	} else {
		// Fall back to default console output behavior
		fmt.Printf("\nDevice Authorization Required:\n")

		// Determine which URL to use and whether to auto-open browser
		browserAvailable := browser.CanOpen()
		hasCompleteURL := response.VerificationURIComplete != ""

		// Auto-open browser if available
		if browserAvailable {
			if hasCompleteURL {
				fmt.Printf("Opening browser to complete authorization...\n")
				fmt.Printf("URL: %s\n", response.VerificationURIComplete)
				if err := browser.Open(response.VerificationURIComplete); err != nil {
					fmt.Printf("Warning: Failed to open browser: %v\n", err)
				}
				fmt.Printf("\nIf the browser didn't open automatically:\n")
				fmt.Printf("   - Visit: %s\n", response.VerificationURI)
				fmt.Printf("   - Enter code: %s\n", response.UserCode)
			} else {
				fmt.Printf("Opening browser for authorization...\n")
				fmt.Printf("URL: %s\n", response.VerificationURI)
				fmt.Printf("Enter this code when prompted: %s\n", response.UserCode)
				if err := browser.Open(response.VerificationURI); err != nil {
					fmt.Printf("Warning: Failed to open browser: %v\n", err)
				}
			}
		} else {
			// No browser available - show manual instructions
			if hasCompleteURL {
				fmt.Printf("Open this URL in your browser to complete login: %s\n", response.VerificationURIComplete)
				fmt.Printf("\n   Alternatively, you can:\n")
				fmt.Printf("   - Visit: %s\n", response.VerificationURI)
				fmt.Printf("   - Enter code: %s\n", response.UserCode)
			} else {
				fmt.Printf("1. Open this URL in your browser: %s\n", response.VerificationURI)
				fmt.Printf("2. Enter this code: %s\n", response.UserCode)
			}
		}

		fmt.Printf("\nWaiting for authorization...\n")
	}

	deviceCodeToken, err := config.DeviceAccessToken(ctx, response)
	if err != nil {
		return nil, fmt.Errorf("device access token request failed: %w", err)
	}

	ts := config.TokenSource(ctx, deviceCodeToken)

	return ts, nil
}
