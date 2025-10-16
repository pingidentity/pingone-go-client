package config

import (
	"context"
	"fmt"
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

func (a *AuthCode) AuthCodeTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (oauth2.TokenSource, error) {
	if a.AuthCodeClientID == nil || *a.AuthCodeClientID == "" {
		return nil, fmt.Errorf("client ID is required for authorization code grant type")
	}

	slog.Debug("Using authorization code token source with provided client ID", "client ID", *a.AuthCodeClientID)

	// Use localhost callback if no redirect URI specified
	redirectURI := "http://localhost:8080/callback"
	if a.AuthCodeRedirectURI != nil && *a.AuthCodeRedirectURI != "" {
		redirectURI = *a.AuthCodeRedirectURI
	}

	var scopes []string
	if a.AuthCodeScopes != nil {
		scopes = *a.AuthCodeScopes
	}

	config := &oauth2.Config{
		ClientID: *a.AuthCodeClientID,
		Endpoint: oauth2.Endpoint{
			AuthURL:  endpoints.AuthorizationURLPath,
			TokenURL: endpoints.TokenURLPath,
		},
		RedirectURL: redirectURI,
		Scopes:      scopes,
	}

	codeVerifier := oauth2.GenerateVerifier()

	// Start local HTTP server to capture callback
	codeChan := make(chan string, 1)
	errChan := make(chan error, 1)

	server, err := startCallbackServer(ctx, redirectURI, codeChan, errChan)
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

// startCallbackServer starts a local HTTP server to handle OAuth2 callbacks
func startCallbackServer(ctx context.Context, redirectURI string, codeChan chan<- string, errChan chan<- error) (*http.Server, error) {
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

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "<html><body><h2>Authorization Failed</h2><p>%s</p><p>You can close this window.</p></body></html>", errDesc)
			return
		}

		// Get authorization code
		code := query.Get("code")
		if code == "" {
			errChan <- fmt.Errorf("no authorization code received")

			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "<html><body><h2>Authorization Failed</h2><p>No authorization code received</p><p>You can close this window.</p></body></html>")
			return
		}

		// Send code to channel
		codeChan <- code

		// Send success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<html><body><h2>Authorization Successful</h2><p>You can close this window and return to the CLI.</p></body></html>")
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
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening browser: %v\nPlease go to the URL manually: %s\n", err, url)
	}
}
