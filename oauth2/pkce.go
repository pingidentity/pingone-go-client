package oauth2

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type PKCEConfig struct {
	AuthzURL     string
	OIDCClientID string
	RedirectURI  string
	TokenURL     string
}

type PKCEClient struct {
	AccessToken string
	*commonClient
	PKCEConfig
}

// func (p *PKCEClient) APIClient() *PKCEClient {
// 	if p == nil {
// 		return &PKCEClient{
// 			commonClient: &commonClient{HTTPClient: &http.Client{}},
// 			PKCEConfig:   PKCEConfig{},
// 		}
// 	}
// 	p.AccessToken = requestPKCEToken()
// 	return p
// }

// func (p *PKCEClient) GetAccessToken() (string, bool) {
// 	if p.AccessToken != "" {
// 		return p.AccessToken, false
// 	}
// 	return p.AccessToken, true
// }

func (p *PKCEClient) exchangeCodeWithPKCE(ctx context.Context, clientID, codeVerifier, code, redirectURI, tokenURL string) (*http.Response, error) {
	formData := url.Values{
		"grant_type":    {"authorization_code"},
		"client_id":     {clientID},
		"redirect_uri":  {redirectURI},
		"code":          {code},
		"code_verifier": {codeVerifier},
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(formData.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := p.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("token exchange failed with status %d: %s", res.StatusCode, body)
	}

	return res, nil
}

func GenerateCodeVerifier() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func GenerateCodeChallenge(codeVerifier string) string {
	h := sha256.New()
	h.Write([]byte(codeVerifier))
	return base64.RawURLEncoding.EncodeToString(h.Sum(nil))
}

// func (p *PKCEClient) requestPKCEToken(authzURL, oidcClientID, port, redirectURI string) (*http.Response, error) {
// 	// --- PKCE setup ---
// 	codeVerifier, err := GenerateCodeVerifier()
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to generate code verifier: %w", err)
// 	}
// 	codeChallenge := GenerateCodeChallenge(codeVerifier)

// 	// The variable name is now consistent.
// 	authZRequestURL := fmt.Sprintf("%s?response_type=code&client_id=%s&redirect_uri=%s&scope=openid&code_challenge=%s&code_challenge_method=S256",
// 		authzURL, oidcClientID, redirectURI, codeChallenge)

// 	// --- Start HTTP Server and Register Handler ---
// 	authCodeChan := make(chan string)

// 	server := &http.Server{Addr: fmt.Sprintf(":%s", port)}

// 	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
// 		// ... (existing handler logic)
// 		code := r.URL.Query().Get("code")
// 		if code == "" {
// 			http.Error(w, "authorization code not found", http.StatusBadRequest)
// 			return
// 		}

// 		// Send the code and shut down the server.
// 		// This stops the ListenAndServe call in the background.
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte("Authentication successful! You can close this window."))

// 		authCodeChan <- code
// 		go func() {
// 			server.Shutdown(context.Background())
// 		}()
// 	})

// 	// --- Open Browser and Wait for Callback ---
// 	// Use the corrected variable name for the URL.
// 	openBrowser(authZRequestURL)

// 	// Start the server in a goroutine.
// 	go func() {
// 		if err := server.ListenAndServe(); err != http.ErrServerClosed {
// 			// Log any other unexpected errors
// 			log.Printf("server error: %v", err)
// 		}
// 	}()

// 	// Wait for the auth code
// 	authCode := <-authCodeChan

// 	// ... (rest of the token exchange logic)
// 	tokenResponse, err := p.exchangeCodeWithPKCE(context.Background(), p.OIDCClientID, codeVerifier, authCode, p.RedirectURI, p.TokenURL)
// 	if err != nil {
// 		return nil, fmt.Errorf("token exchange failed: %w", err)
// 	}
// 	return tokenResponse, nil
// }

// openBrowser opens the specified URL in the user's default browser.
