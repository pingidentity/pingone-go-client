// Copyright © 2025 Ping Identity Corporation
package config

import (
	"context"
	_ "embed"
	"fmt"
	"html/template"
	"log/slog"
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
	// DefaultAuthCodeRedirectURIPort is the default port for the authorization code redirect URI
	DefaultAuthCodeRedirectURIPort = "8080"

	// DefaultAuthCodeRedirectURIPath is the default path for the authorization code redirect URI
	DefaultAuthCodeRedirectURIPath = "/callback"

	// DefaultAuthCodeRedirectURIPrefix is the default redirect URI for the authorization code
	DefaultAuthCodeRedirectURIPrefix = "http://localhost:"

	// DefaultAuthCodeRedirectURI
	DefaultAuthCodeRedirectURI = DefaultAuthCodeRedirectURIPrefix + DefaultAuthCodeRedirectURIPort + DefaultAuthCodeRedirectURIPath
)

func (a *AuthCode) AuthCodeTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (oauth2.TokenSource, error) {
	if a.AuthCodeClientID == nil || *a.AuthCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for authorization code grant type")
	}

	slog.Debug("Using authorization code token source with provided client ID", "client ID", *a.AuthCodeClientID)

	redirectURI := DefaultAuthCodeRedirectURI
	redirectURIPath := DefaultAuthCodeRedirectURIPath
	redirectURIPort := DefaultAuthCodeRedirectURIPort

	if a.AuthCodeRedirectURI.Port != "" && a.AuthCodeRedirectURI.Path != "" {
		redirectURIPath = a.AuthCodeRedirectURI.Path
		redirectURIPort = a.AuthCodeRedirectURI.Port
	} else if a.AuthCodeRedirectURI.Port != "" {
		redirectURIPort = a.AuthCodeRedirectURI.Port
	} else if a.AuthCodeRedirectURI.Path != "" {
		redirectURIPath = a.AuthCodeRedirectURI.Path
	}

	redirectURI = fmt.Sprintf("%s%s%s", DefaultAuthCodeRedirectURIPrefix, redirectURIPort, redirectURIPath)

	var scopes []string
	if a.AuthCodeScopes != nil {
		scopes = *a.AuthCodeScopes
	}

	config := &oauth2.Config{
		ClientID:    *a.AuthCodeClientID,
		Endpoint:    endpoints.Endpoint,
		RedirectURL: redirectURI,
		Scopes:      scopes,
	}

	codeVerifier := oauth2.GenerateVerifier()

	// Start local HTTP server to capture callback
	codeChan := make(chan string, 1)
	errChan := make(chan error, 1)

	server, err := startCallbackServer(redirectURI, codeChan, errChan)
	if err != nil {
		return nil, fmt.Errorf("failed to start callback server: %w", err)
	}
	defer server.Close()

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
		return nil, fmt.Errorf("failed to exchange code for token: %w", err)
	}

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
		Name:  "Authorization Code OAuth2 Flow Failed",
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
		Name:  "Authorization Code OAuth2 Flow Success",
	}

	tmpl, err := template.New("success").Parse(authResultHTML)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	return tmpl.Execute(w, successData)
}

// startCallbackServer starts a local HTTP server to handle OAuth2 callbacks
func startCallbackServer(redirectURI string, codeChan chan<- string, errChan chan<- error) (*http.Server, error) {
	// Parse the redirect URI to get the port
	parsedURI, err := url.Parse(redirectURI)
	if err != nil {
		return nil, fmt.Errorf("invalid redirect URI: %w", err)
	}

	// Extract port from URI or use default
	port := parsedURI.Port()
	if port == "" {
		port = "8080"
	}

	// Extract path and ensure it's valid for HTTP mux
	path := parsedURI.Path
	if path == "" {
		path = "/callback"
	}
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	// Create HTTP server
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
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

		// Send success response
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)

		err := returnSuccessPage(w)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading success page. Authentication was successful.\n%s", err)
		}
	})

	// Start server in background
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errChan <- fmt.Errorf("callback server error: %w", err)
		}
	}()

	// Give server time to start
	time.Sleep(100 * time.Millisecond)

	return server, nil
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
