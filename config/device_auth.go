// Copyright © 2025 Ping Identity Corporation
package config

import (
	"context"
	"fmt"

	"github.com/pingidentity/pingone-go-client/utils/browser"
	"golang.org/x/oauth2"
)

const (
	// deviceAuthPromptHeader is the header displayed for device authorization prompt
	deviceAuthPromptHeader = "\nDevice Authorization Required:\n"

	// deviceAuthBrowserOpeningMessage is the message shown when opening browser with complete URL
	deviceAuthBrowserOpeningMessage = "Opening browser to complete authorization...\n"

	// deviceAuthBrowserOpeningMessageBasic is the message shown when opening browser with basic URL
	deviceAuthBrowserOpeningMessageBasic = "Opening browser for authorization...\n"

	// deviceAuthURLLabel is the label for displaying the URL
	deviceAuthURLLabel = "URL: %s\n"

	// deviceAuthCodePrompt is the prompt for entering the user code
	deviceAuthCodePrompt = "Enter this code when prompted: %s\n"

	// deviceAuthBrowserFailWarning is the warning when browser fails to open
	deviceAuthBrowserFailWarning = "Warning: Failed to open browser: %v\n"

	// deviceAuthManualInstructionsHeader is the header for manual browser instructions
	deviceAuthManualInstructionsHeader = "\nIf the browser didn't open automatically:\n"

	// deviceAuthManualVisitPrompt is the prompt to visit the URL manually
	deviceAuthManualVisitPrompt = "   - Visit: %s\n"

	// deviceAuthManualCodePrompt is the prompt to enter the code manually
	deviceAuthManualCodePrompt = "   - Enter code: %s\n"

	// deviceAuthCompleteURLPrompt is the prompt when complete URL is available without browser
	deviceAuthCompleteURLPrompt = "Open this URL in your browser to complete login: %s\n"

	// deviceAuthAlternativeInstructionsHeader is the header for alternative instructions
	deviceAuthAlternativeInstructionsHeader = "\n   Alternatively, you can:\n"

	// deviceAuthManualStep1 is the first step for manual authentication
	deviceAuthManualStep1 = "1. Open this URL in your browser: %s\n"

	// deviceAuthManualStep2 is the second step for manual authentication
	deviceAuthManualStep2 = "2. Enter this code: %s\n"

	// deviceAuthWaitingMessage is the message shown while waiting for authorization
	deviceAuthWaitingMessage = "\nWaiting for authorization...\n"
)

func (d *DeviceCode) DeviceAuthTokenSource(ctx context.Context, endpoints oauth2.Endpoint) (oauth2.TokenSource, error) {
	if d.DeviceCodeClientID == nil || *d.DeviceCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for device code grant type")
	}

	var scopes []string
	if d.DeviceCodeScopes != nil {
		scopes = *d.DeviceCodeScopes
	}

	config := &oauth2.Config{
		ClientID: *d.DeviceCodeClientID,
		Scopes:   scopes,
		Endpoint: endpoints,
	}

	response, err := config.DeviceAuth(ctx)
	if err != nil {
		return nil, fmt.Errorf("device auth request failed: %w", err)
	}

	// Use custom handler if provided, otherwise use default with complete URL support
	handler := d.OnDisplayPrompt
	if handler == nil {
		// Use a closure to capture the complete URL for the default handler
		handler = DefaultDeviceCodePromptHandler
	}

	if err := handler(response.VerificationURI, response.UserCode); err != nil {
		return nil, fmt.Errorf("prompt handler failed: %w", err)
	}

	deviceCodeToken, err := config.DeviceAccessToken(ctx, response)
	if err != nil {
		return nil, fmt.Errorf("device access token request failed: %w", err)
	}

	ts := config.TokenSource(ctx, deviceCodeToken)

	return ts, nil
}

// DefaultDeviceCodePromptHandler is a simple handler that displays device code prompts.
// This function can be used by consumer projects as a reference implementation or directly.
// For better UX with complete URLs, use the closure pattern shown in DeviceAuthTokenSource.
// This function implements the DeviceCodePromptHandler interface pattern.
func DefaultDeviceCodePromptHandler(verificationURI, userCode string) error {
	fmt.Print(deviceAuthPromptHeader)

	// Determine which URL to use and whether to auto-open browser
	browserAvailable := browser.CanOpen()
	verificationURIComplete := fmt.Sprintf("%s?user_code=%s", verificationURI, userCode)

	// Auto-open browser if available
	if browserAvailable {
		fmt.Print(deviceAuthBrowserOpeningMessage)
		fmt.Printf(deviceAuthURLLabel, verificationURIComplete)
		if err := browser.Open(verificationURIComplete); err != nil {
			fmt.Printf(deviceAuthBrowserFailWarning, err)
		}
		fmt.Print(deviceAuthManualInstructionsHeader)
		fmt.Printf(deviceAuthManualVisitPrompt, verificationURI)
		fmt.Printf(deviceAuthManualCodePrompt, userCode)
	} else {
		// No browser available - show manual instructions
		fmt.Printf(deviceAuthCompleteURLPrompt, verificationURIComplete)
		fmt.Print(deviceAuthAlternativeInstructionsHeader)
		fmt.Printf(deviceAuthManualVisitPrompt, verificationURI)
		fmt.Printf(deviceAuthManualCodePrompt, userCode)
	}

	fmt.Print(deviceAuthWaitingMessage)
	return nil
}
