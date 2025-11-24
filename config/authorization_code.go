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
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/pingidentity/pingone-go-client/oidc/endpoints"
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
	codeChan := make(chan string, 1)
	errChan := make(chan error, 1)
	tokenResultChan := make(chan error, 1) // nil for success, error for failure

	server, err := startCallbackServer(redirectURI, codeChan, errChan, tokenResultChan)
	if err != nil {
		return nil, fmt.Errorf("failed to start callback server: %w", err)
	}
	defer func() {
		if closeErr := server.Close(); closeErr != nil {
			fmt.Printf("Warning: failed to close server: %v\n", closeErr)
		}
	}()

	// Generate authorization URL and open browser
	authURL := config.AuthCodeURL("state", oauth2.AccessTypeOffline, oauth2.S256ChallengeOption(codeVerifier))
	fmt.Printf("Opening browser for authorization: %s\n", authURL)
	openBrowser(authURL)
	fmt.Println("Waiting for authorization callback...")

	// Wait for authorization code or error
	var code string
	select {
	case code = <-codeChan:
		fmt.Println("Authorization code received")
	case err := <-errChan:
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
		time.Sleep(500 * time.Millisecond)

		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

	// Signal success to show success page
	tokenResultChan <- nil

	// Wait a moment for the HTTP response to be sent before returning
	time.Sleep(500 * time.Millisecond)

	slog.Debug("Successfully obtained access token via authorization code flow")

	// Return a static token source with the token we just obtained
	return oauth2.StaticTokenSource(tok), nil
}

func returnFailedPage(w http.ResponseWriter) error {
	failedData := struct {
		Title string
		Name  string
	}{
		Title: "Authorization Failed",
		Name:  "An error has occurred and authorization was not successful.",
	}

	tmpl, err := template.New("failed").Parse(authResultHTML)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}
	return tmpl.Execute(w, failedData)
}

func returnSuccessPage(w http.ResponseWriter) error {
	successData := struct {
		Title string
		Name  string
	}{
		Title: "Authorization Success",
		Name:  "You have successfully authenticated to your PingOne environment and have authorized API access.",
	}

	tmpl, err := template.New("success").Parse(authResultHTML)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	return tmpl.Execute(w, successData)
}

// startCallbackServer starts a local HTTP server to handle OAuth2 callbacks
func startCallbackServer(redirectURI string, codeChan chan<- string, errChan chan<- error, tokenResultChan <-chan error) (*http.Server, error) {
	// Parse the redirect URI to get the port
	parsedURI, err := url.Parse(redirectURI)
	if err != nil {
		return nil, fmt.Errorf("invalid redirect URI: %w", err)
	}

	// Extract port from URI or use default
	port := parsedURI.Port()
	if port == "" {
		port = "7464"
	}

	// Extract path and ensure it's valid for HTTP mux
	path := parsedURI.Path
	if path == "" {
		path = "/callback"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// Test if port is available and keep the listener
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return nil, fmt.Errorf("port %s is not available: %w", port, err)
	}

	// Create HTTP server
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}

	// Handle callback endpoint
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		// Check for error in callback
		if errCode := query.Get("error"); errCode != "" {
			errDesc := query.Get("error_description")
			if errDesc == "" {
				errDesc = errCode
			}
			errChan <- fmt.Errorf("authorization error: %s", errDesc)

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)

			err := returnFailedPage(w)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Authentication failed.")
			}

			return
		}

		// Get authorization code
		code := query.Get("code")
		if code == "" {
			errChan <- fmt.Errorf("no authorization code received")

			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusBadRequest)
			err := returnFailedPage(w)
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
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(http.StatusOK)

				err := returnSuccessPage(w)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error loading success page. Authentication was successful.\n%s", err)
				}
			} else {
				// Token exchange failed
				w.Header().Set("Content-Type", "text/html; charset=utf-8")
				w.WriteHeader(http.StatusUnauthorized)
				err := returnFailedPage(w)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error loading failed page. Token exchange failed: %v", tokenErr)
				}
				if flusher, ok := w.(http.Flusher); ok {
					flusher.Flush()
				}
			}
		case <-time.After(30 * time.Second):
			// Token exchange timed out
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusInternalServerError)
			err := returnFailedPage(w)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error loading failed page. Token exchange timed out.")
			}
		}
	})

	// Start server in background using the existing listener
	serverStarted := make(chan error, 1)
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
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		// Check if server failed
		select {
		case err := <-serverStarted:
			return nil, fmt.Errorf("failed to start server: %w", err)
		default:
		}

		// Try to connect
		conn, err := net.DialTimeout("tcp", ":"+port, 50*time.Millisecond)
		if err == nil {
			if closeErr := conn.Close(); closeErr != nil {
				slog.Warn("Failed to close connection during server verification", "error", closeErr)
			}
			// Server is accepting connections
			return server, nil
		}

		// Wait a bit before retrying
		time.Sleep(10 * time.Millisecond)
	}

	return nil, fmt.Errorf("server failed to start accepting connections in time")
}

func openBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("explorer", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening browser: %v\nPlease go to the URL manually: %s\n", err, url)
	}
}
