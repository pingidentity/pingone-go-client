// Copyright © 2025 Ping Identity Corporation

package endpoints

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
)

// UserInfo retrieves user information from the OIDC userinfo endpoint.
//
// This function makes an HTTP GET request to the userinfo endpoint with the provided
// access token and returns the user profile information as a map of claims.
//
// It returns a map containing the user's profile claims (such as sub, name, email, etc.)
// as defined by the OpenID Connect specification, and an error if the request fails or
// the response cannot be parsed.
func (endpoint OIDCEndpoint) UserInfo(ctx context.Context, accessToken oauth2.Token) (map[string]interface{}, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.UserInfoURLPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create userinfo request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken.AccessToken))
	req.Header.Set("Accept", "application/json")

	// Execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to execute userinfo request: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read userinfo response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("userinfo request failed with status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the JSON response
	//TODO define p1-specific struct for relevant user info
	var userInfo map[string]interface{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to parse userinfo response: %w", err)
	}

	return userInfo, nil
}
