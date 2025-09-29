package oauth2

import (
	"net/http"
)

type commonClient struct {
	HTTPClient *http.Client
	ClientID   string
	EnvID      string
}

func (c *commonClient) Do(req *http.Request) (*http.Response, error) {
	return c.HTTPClient.Do(req)
}

// func parseTokenResponse(body io.Reader) (*TokenResponse, error) {
// 	var token TokenResponse
// 	decoder := json.NewDecoder(body)
// 	if err := decoder.Decode(&token); err != nil {
// 		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
// 	}
// 	return &token, nil
// }

// type Client struct {
// 	AuthURL    string
// 	Device     *DeviceAuthClient
// 	HTTPClient *http.Client
// 	PKCE       *PKCEAuthClient
// }
