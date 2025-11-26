// Copyright © 2025 Ping Identity Corporation
package config

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
	"github.com/pingidentity/pingone-go-client/utils/browser"
	"golang.org/x/oauth2"
)

//go:embed html/auth_result.html
var authResultHTML string

const (
	// defaultAuthorizationCodeRedirectURIPort is the default port for the authorization code redirect URI
	defaultAuthorizationCodeRedirectURIPort = "7464"

	// defaultAuthorizationCodeRedirectURIPath is the default path for the authorization code redirect URI
	defaultAuthorizationCodeRedirectURIPath = "/callback"

	// defaultAuthorizationCodeRedirectURIPrefix is the default redirect URI for the authorization code
	defaultAuthorizationCodeRedirectURIPrefix = "http://127.0.0.1:"

	// defaultAuthorizationCodeRedirectURI is the default redirect URI for the authorization code
	defaultAuthorizationCodeRedirectURI = defaultAuthorizationCodeRedirectURIPrefix + defaultAuthorizationCodeRedirectURIPort + defaultAuthorizationCodeRedirectURIPath

	// defaultProjectName is the default project name displayed on auth result pages
	defaultProjectName = "Ping Identity Developer Tools"

	// defaultAuthFailedHeading is the default heading displayed on the authentication failure page
	defaultAuthFailedHeading = "Authorization Failed"

	// defaultAuthFailedMessage is the default message displayed on the authentication failure page
	defaultAuthFailedMessage = "An error has occurred and authorization was not successful."

	// defaultAuthSuccessHeading is the default heading displayed on the authentication success page
	defaultAuthSuccessHeading = "Authorization Success"

	// defaultAuthSuccessMessage is the default message displayed on the authentication success page
	defaultAuthSuccessMessage = "You have successfully authenticated to your PingOne environment and have authorized API access."

	// templateNameFailed is the template name for failed authentication page
	templateNameFailed = "failed"

	// templateNameSuccess is the template name for successful authentication page
	templateNameSuccess = "success"

	// contentTypeHTML is the content type for HTML responses
	contentTypeHTML = "text/html; charset=utf-8"

	// callbackServerReadHeaderTimeout is the timeout for reading HTTP headers on the callback server
	callbackServerReadHeaderTimeout = 10 * time.Second

	// tokenExchangeTimeout is the timeout for waiting for token exchange to complete
	tokenExchangeTimeout = 30 * time.Second

	// authSuccessWaitTime is the time to wait after successful auth before returning (allows HTTP response to be sent)
	authSuccessWaitTime = 1 * time.Second

	// serverVerificationRetryDelay is the delay between retries when verifying server startup
	serverVerificationRetryDelay = 10 * time.Millisecond

	// serverVerificationMaxRetries is the maximum number of retries when verifying server startup
	serverVerificationMaxRetries = 10

	// serverVerificationDialTimeout is the timeout for dial attempts when verifying server startup
	serverVerificationDialTimeout = 50 * time.Millisecond

	// httpChannelBufferSize is the buffer size for HTTP callback channels
	httpChannelBufferSize = 1

	// httpStatusUnauthorized is the HTTP status code for unauthorized requests
	httpStatusUnauthorized = http.StatusUnauthorized

	// urlQueryParamError is the URL query parameter name for error code
	urlQueryParamError = "error"

	// urlQueryParamErrorDescription is the URL query parameter name for error description
	urlQueryParamErrorDescription = "error_description"

	// urlQueryParamCode is the URL query parameter name for authorization code
	urlQueryParamCode = "code"

	// urlPathPrefix is the prefix character for URL paths
	urlPathPrefix = "/"

	// networkProtocolTCP is the network protocol for TCP connections
	networkProtocolTCP = "tcp"

	// networkPortPrefix is the prefix character for port numbers
	networkPortPrefix = ":"
)

// Get default authorization code redirect URI port
func GetDefaultAuthorizationCodeRedirectURIPort() string {
	return defaultAuthorizationCodeRedirectURIPort
}

// Get default authorization code redirect URI path
func GetDefaultAuthorizationCodeRedirectURIPath() string {
	return defaultAuthorizationCodeRedirectURIPath
}

// Get default authorization code redirect URI
func GetDefaultAuthorizationCodeRedirectURI() string {
	return defaultAuthorizationCodeRedirectURI
}

// AuthorizationCodeTokenSource returns an oauth2.TokenSource using the authorization code grant type
func (a *AuthorizationCode) AuthorizationCodeTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (oauth2.TokenSource, error) {
	if a.AuthorizationCodeClientID == nil || *a.AuthorizationCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for authorization code grant type")
	}

	slog.Debug("Using authorization code token source with provided client ID", "client ID", *a.AuthorizationCodeClientID)
	redirectURIPath := GetDefaultAuthorizationCodeRedirectURIPath()
	redirectURIPort := GetDefaultAuthorizationCodeRedirectURIPort()

	if a.AuthorizationCodeRedirectURI.Port != "" {
		redirectURIPort = a.AuthorizationCodeRedirectURI.Port
	}

	if a.AuthorizationCodeRedirectURI.Path != "" {
		redirectURIPath = a.AuthorizationCodeRedirectURI.Path
	}

	redirectURI := fmt.Sprintf("%s%s%s", defaultAuthorizationCodeRedirectURIPrefix, redirectURIPort, redirectURIPath)

	var scopes []string
	if a.AuthorizationCodeScopes != nil {
		scopes = *a.AuthorizationCodeScopes
	}

	config := &oauth2.Config{
		ClientID:    *a.AuthorizationCodeClientID,
		Endpoint:    endpoints.Endpoint,
		RedirectURL: redirectURI,
		Scopes:      scopes,
	}

	codeVerifier := oauth2.GenerateVerifier()

	// Start local HTTP server to capture callback
	codeChan := make(chan string, httpChannelBufferSize)
	errChan := make(chan error, httpChannelBufferSize)
	tokenResultChan := make(chan error, httpChannelBufferSize) // nil for success, error for failure

	server, err := startCallbackServer(redirectURI, codeChan, errChan, tokenResultChan, a.CustomPageDataSuccess, a.CustomPageDataError)
	if err != nil {
		return nil, fmt.Errorf("failed to start callback server: %w", err)
	}
	defer func() {
		if closeErr := server.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close server: %v\n", closeErr)
		}
	}()

	// Generate authorization URL and handle browser opening
	authURL := config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(codeVerifier))

	// Check if custom browser handler is provided
	if a.OnOpenBrowser != nil {
		// Use custom handler - delegate UX to consumer
		if err := a.OnOpenBrowser(authURL); err != nil {
			return nil, fmt.Errorf("custom browser handler failed: %w", err)
		}
	} else {
		// Fall back to default browser opening behavior
		fmt.Printf("Opening browser for authorization: %s\n", authURL)
		if err := browser.Open(authURL); err != nil {
			fmt.Printf("Warning: Failed to open browser automatically: %v\n", err)
			fmt.Printf("Please open this URL in your browser manually: %s\n", authURL)
		}
		fmt.Println("Waiting for authorization callback...")
	}

	// Wait for authorization code or error
	var code string
	select {
	case code = <-codeChan:
		fmt.Println("Authorization code received")
	case err := <-errChan:
		// Wait a moment for the HTTP response to be sent before returning
		time.Sleep(authSuccessWaitTime)
		return nil, fmt.Errorf("authorization failed: %w", err)
	case <-ctx.Done():
		return nil, fmt.Errorf("authorization cancelled: %w", ctx.Err())
	}

	// Exchange authorization code for token
	tok, err := config.Exchange(ctx, code, oauth2.VerifierOption(codeVerifier))

	if err != nil {
		// Signal failure to show error page
		tokenResultChan <- err

		// Wait a moment for the HTTP response to be sent before returning
		time.Sleep(authSuccessWaitTime)

		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	// Signal success to show success page
	tokenResultChan <- nil

	// Wait a moment for the HTTP response to be sent before returning
	time.Sleep(authSuccessWaitTime)

	slog.Debug("Successfully obtained access token via authorization code flow")

	// Return a static token source with the token we just obtained
	return oauth2.StaticTokenSource(tok), nil
}

func returnFailedPage(w http.ResponseWriter, errorDetails string, customPageData *AuthResultPageData) error {
	// Create template data
	templateData := struct {
		ProjectName  string
		Heading      string
		Message      string
		ErrorDetails string
		IsSuccess    bool
	}{
		ProjectName:  defaultProjectName,
		Heading:      defaultAuthFailedHeading,
		Message:      defaultAuthFailedMessage,
		ErrorDetails: errorDetails,
		IsSuccess:    false,
	}

	// Override with custom data if provided
	if customPageData != nil {
		if customPageData.ProjectName != "" {
			templateData.ProjectName = customPageData.ProjectName
		}
		if customPageData.Heading != "" {
			templateData.Heading = customPageData.Heading
		}
		if customPageData.Message != "" {
			templateData.Message = customPageData.Message
		}
	}

	tmpl, err := template.New(templateNameFailed).Parse(authResultHTML)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}
	return tmpl.Execute(w, templateData)
}

func returnSuccessPage(w http.ResponseWriter, customPageData *AuthResultPageData) error {
	// Create template data
	templateData := struct {
		ProjectName  string
		Heading      string
		Message      string
		ErrorDetails string
		IsSuccess    bool
	}{
		ProjectName:  defaultProjectName,
		Heading:      defaultAuthSuccessHeading,
		Message:      defaultAuthSuccessMessage,
		ErrorDetails: "", // Empty for success
		IsSuccess:    true,
	}

	// Override with custom data if provided
	if customPageData != nil {
		if customPageData.ProjectName != "" {
			templateData.ProjectName = customPageData.ProjectName
		}
		if customPageData.Heading != "" {
			templateData.Heading = customPageData.Heading
		}
		if customPageData.Message != "" {
			templateData.Message = customPageData.Message
		}
	}

	tmpl, err := template.New(templateNameSuccess).Parse(authResultHTML)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	return tmpl.Execute(w, templateData)
}

// startCallbackServer starts a local HTTP server to handle OAuth2 callbacks
func startCallbackServer(redirectURI string, codeChan chan<- string, errChan chan<- error, tokenResultChan <-chan error, customPageDataSuccess *AuthResultPageData, customPageDataError *AuthResultPageData) (*http.Server, error) {
	// Parse the redirect URI to get the port
	parsedURI, err := url.Parse(redirectURI)
	if err != nil {
		return nil, fmt.Errorf("invalid redirect URI: %w", err)
	}

	// Extract port from URI or use default
	port := parsedURI.Port()
	if port == "" {
		port = defaultAuthorizationCodeRedirectURIPort
	}

	// Extract path and ensure it's valid for HTTP mux
	path := parsedURI.Path
	if path == "" {
		path = defaultAuthorizationCodeRedirectURIPath
	}
	if !strings.HasPrefix(path, urlPathPrefix) {
		path = urlPathPrefix + path
	}

	// Test if port is available and keep the listener
	listener, err := net.Listen(networkProtocolTCP, networkPortPrefix+port)
	if err != nil {
		return nil, fmt.Errorf("port %s is not available: %w", port, err)
	}

	// Create HTTP server
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:              networkPortPrefix + port,
		Handler:           mux,
		ReadHeaderTimeout: callbackServerReadHeaderTimeout,
	}

	// Handle callback endpoint
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		// Check for error in callback
		if errCode := query.Get(urlQueryParamError); errCode != "" {
			errDesc := query.Get(urlQueryParamErrorDescription)
			if errDesc == "" {
				errDesc = errCode
			}
			errChan <- fmt.Errorf("authorization error: %s", errDesc)

			w.Header().Set("Content-Type", contentTypeHTML)
			w.WriteHeader(http.StatusBadRequest)

			err := returnFailedPage(w, errDesc, customPageDataError)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Authentication failed.")
			}

			return
		}

		// Get authorization code
		code := query.Get(urlQueryParamCode)
		if code == "" {
			errChan <- fmt.Errorf("no authorization code received")

			w.Header().Set("Content-Type", contentTypeHTML)
			w.WriteHeader(http.StatusBadRequest)
			err := returnFailedPage(w, "No authorization code received", customPageDataError)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Authentication failed.")
			}
			return
		}

		// Send code to channel
		codeChan <- code

		// Wait for token exchange to complete before showing page
		select {
		case tokenErr := <-tokenResultChan:
			if tokenErr == nil {
				// Token exchange succeeded
				w.Header().Set("Content-Type", contentTypeHTML)
				w.WriteHeader(http.StatusOK)

				err := returnSuccessPage(w, customPageDataSuccess)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error loading success page. Authentication was successful.\n%s", err)
				}
			} else {
				// Token exchange failed
				w.Header().Set("Content-Type", contentTypeHTML)
				w.WriteHeader(httpStatusUnauthorized)
				err := returnFailedPage(w, fmt.Sprintf("Token exchange failed: %v", tokenErr), customPageDataError)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error loading failed page. Token exchange failed: %v", tokenErr)
				}
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
			}
		case <-time.After(tokenExchangeTimeout):
			// Token exchange timed out
			w.Header().Set("Content-Type", contentTypeHTML)
			w.WriteHeader(http.StatusInternalServerError)
			err := returnFailedPage(w, "Token exchange timed out", customPageDataError)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Token exchange timed out.")
			}
		}
	})

	// Start server in background using the existing listener
	serverStarted := make(chan error, httpChannelBufferSize)
	go func() {
		// Use Serve() with the existing listener instead of ListenAndServe()
		if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
			select {
			case serverStarted <- err:
				// Server failed to start
			default:
				// Server was already marked as started, so this is a runtime error
				errChan <- fmt.Errorf("callback server error: %w", err)
			}
		}
	}()

	// Verify server is actually accepting connections before returning
	maxRetries := serverVerificationMaxRetries
	for i := 0; i < maxRetries; i++ {
		// Check if server failed
		select {
		case err := <-serverStarted:
			return nil, fmt.Errorf("failed to start server: %w", err)
		default:
		}

		// Try to connect
		conn, err := net.DialTimeout(networkProtocolTCP, networkPortPrefix+port, serverVerificationDialTimeout)
		if err == nil {
			if closeErr := conn.Close(); closeErr != nil {
				slog.Warn("Failed to close connection during server verification", "error", closeErr)
			}
			// Server is accepting connections
			return server, nil
		}

		// Wait a bit before retrying
		time.Sleep(serverVerificationRetryDelay)
	}

	return nil, fmt.Errorf("server failed to start accepting connections in time")
}
