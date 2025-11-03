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

// UserInfoAddress represents the address claim in the UserInfo response.
type UserInfoAddress struct {
	StreetAddress *string `json:"street_address,omitempty"`
	Locality      *string `json:"locality,omitempty"`
	Region        *string `json:"region,omitempty"`
	PostalCode    *string `json:"postal_code,omitempty"`
	Country       *string `json:"country,omitempty"`
}

// UserInfoResponse represents the response from the OIDC userinfo endpoint.
// This structure includes typical claims returned by PingOne's userinfo endpoint.
//
// Only the sub claim is guaranteed to be present in all responses.
// All other claims are optional and may be omitted depending on the scopes requested.
//
// Any additional claims not explicitly defined in this struct will be captured in
// the AdditionalClaims field.
type UserInfoResponse struct {
	Sub               string           `json:"sub"`
	Name              *string          `json:"name,omitempty"`
	GivenName         *string          `json:"given_name,omitempty"`
	FamilyName        *string          `json:"family_name,omitempty"`
	MiddleName        *string          `json:"middle_name,omitempty"`
	PreferredUsername *string          `json:"preferred_username,omitempty"`
	Email             *string          `json:"email,omitempty"`
	Address           *UserInfoAddress `json:"address,omitempty"`
	UpdatedAt         *int64           `json:"updated_at,omitempty"`
	// AdditionalClaims contains any claims returned by the userinfo endpoint that are not
	// explicitly defined as fields in this struct.
	AdditionalClaims map[string]interface{} `json:"-"`
}

type _UserInfoResponse UserInfoResponse

// UnmarshalJSON implements custom JSON unmarshaling for UserInfoResponse.
// Any additional claims are added to the AdditionalClaims map.
func (u *UserInfoResponse) UnmarshalJSON(data []byte) error {
	var allClaims map[string]interface{}
	if err := json.Unmarshal(data, &allClaims); err != nil {
		return err
	}

	// Unmarshal the defined fields
	var response _UserInfoResponse
	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}
	*u = UserInfoResponse(response)

	// Populate AdditionalClaims with any claims not in the known set
	knownFields := map[string]bool{
		"sub":                true,
		"name":               true,
		"given_name":         true,
		"family_name":        true,
		"middle_name":        true,
		"preferred_username": true,
		"email":              true,
		"address":            true,
		"updated_at":         true,
	}
	u.AdditionalClaims = make(map[string]interface{})
	for key, value := range allClaims {
		if !knownFields[key] {
			u.AdditionalClaims[key] = value
		}
	}

	return nil
}

// UserInfo retrieves user information from the OIDC userinfo endpoint.
//
// This function makes an HTTP GET request to the userinfo endpoint with the provided
// access token and returns the user profile information as a structured UserInfoResponse.
//
// Only the sub claim is guaranteed to be present in the response.
// All other claims are optional and may be omitted depending on the scopes requested.
func (endpoint OIDCEndpoint) UserInfo(ctx context.Context, accessToken oauth2.Token) (*UserInfoResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.UserInfoURLPath, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create userinfo request: %w", err)
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken.AccessToken))

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
	var userInfo UserInfoResponse
	if err := json.Unmarshal(body, &userInfo); err != nil {
		return nil, fmt.Errorf("failed to parse userinfo response: %w", err)
	}

	return &userInfo, nil
}
