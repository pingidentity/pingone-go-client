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

	// deviceAuthURLLabel is the label for displaying the URL
	deviceAuthURLLabel = "URL: %s\n"

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

	// deviceAuthWaitingMessage is the message shown while waiting for authorization
	deviceAuthWaitingMessage = "\nWaiting for authorization...\n"
)

// DeviceAuthTokenSource returns an oauth2.TokenSource using the device code grant type.
// This function implements the OAuth2 Device Authorization Grant (RFC 8628) with PKCE (RFC 7636)
// for enhanced security.
//
// The device code flow is designed for devices with limited input capabilities (e.g., smart TVs,
// CLI tools, IoT devices). The user completes authentication on a separate device with better
// input/display capabilities.
//
// PKCE (Proof Key for Code Exchange) is automatically enabled for this flow to prevent
// authorization code interception attacks, following OAuth2 security best practices.
//
// The function requires a valid client ID to be configured. The device code flow does not
// require a client secret, making it suitable for public clients.
//
// Security features:
//   - PKCE with S256 challenge method (SHA256 hashing)
//   - Device code is single-use and time-limited
//   - User must authenticate on a trusted device
//   - Automatic polling with server-controlled intervals
//
// Returns an oauth2.TokenSource that can be used for authenticated API calls, or an error
// if the device authorization fails.
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

	// Generate PKCE verifier for enhanced security
	// PKCE prevents authorization code interception attacks and is recommended for all OAuth2 flows
	codeVerifier := oauth2.GenerateVerifier()

	// Request device authorization with PKCE challenge
	response, err := config.DeviceAuth(ctx, oauth2.S256ChallengeOption(codeVerifier), oauth2.AccessTypeOffline)
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

	// Exchange device code for access token with PKCE verifier
	// The verifier proves that this client is the same one that initiated the device authorization
	deviceCodeToken, err := config.DeviceAccessToken(ctx, response, oauth2.VerifierOption(codeVerifier))
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
