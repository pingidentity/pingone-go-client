package oauth2

type DeviceAuthCode struct {
	RegionCode string `json:"regionCode"`
	Scope      string `json:"scope"`
}

type DeviceAuthResponse struct {
	DeviceCode      string `json:"device_code"`
	UserCode        string `json:"user_code"`
	VerificationURI string `json:"verification_uri"`
	ExpiresIn       int64  `json:"expires_in"`
	Interval        int64  `json:"interval"`
}

// func (c *Configuration) DeviceAuthTokenSource(ctx context.Context, endpoints endpoints.OIDCEndpoint) (*oauth2.TokenSource, error) {
// 	if c.Auth.ClientID == nil || *c.Auth.ClientID == "" {
// 		return nil, fmt.Errorf("client ID is required for client credentials grant type")
// 	}

// 	slog.Debug("Using client credentials token source with provided client ID", "client ID", *c.Auth.ClientID)
// 	if c.Auth.ClientSecret != nil && *c.Auth.ClientSecret != "" {
// 		config := &clientcredentials.Config{
// 			ClientID:     *c.Auth.ClientID,
// 			ClientSecret: *c.Auth.ClientSecret,
// 			TokenURL:     endpoints.TokenURL,
// 		}
// 		ts := config.TokenSource(ctx)
// 		slog.Debug("Using standard client credentials token source as client secret has been provided")
// 		return &ts, nil
// 	}

// 	// tmp until private keys are supported
// 	return nil, fmt.Errorf("client secret is required for client credentials grant type")

// }

// func (c *Client) RequestDeviceCode(ctx context.Context, clientID, scope, deviceAuthURL string) (*DeviceAuthResponse, error) {
// 	formData := url.Values{
// 		"grant_type": {"urn:ietf:params:oauth:grant-type:device_code"},
// 		"client_id":  {clientID},
// 		"scope":      {scope},
// 	}

// 	c.Device.RegionCode = regionCode
// 	c.Device.Scope = scope

// 	req, err := http.NewRequestWithContext(ctx, http.MethodPost, deviceAuthURL, strings.NewReader(formData.Encode()))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create request: %w", err)
// 	}
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	res, err := c.Do(req)
// 	if err != nil {
// 		return nil, fmt.Errorf("device auth request failed: %w", err)
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read response body: %w", err)
// 	}

// 	if res.StatusCode != http.StatusOK {
// 		return nil, fmt.Errorf("device auth request failed with status %d: %s", res.StatusCode, body)
// 	}

// 	var result DeviceAuthResponse
// 	if err := json.Unmarshal(body, &result); err != nil {
// 		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
// 	}
// 	return &result, nil
// }

// func (c *Client) PollForDeviceToken(ctx context.Context, c.Device.OIDCClientID, deviceCode, tokenURL string) (*TokenResponse, error) {
// 	formData := url.Values{
// 		"grant_type":  {"urn:ietf:params:oauth:grant-type:device_code"},
// 		"client_id":   {clientID},
// 		"device_code": {deviceCode},
// 	}
// 	req, err := http.NewRequestWithContext(ctx, http.MethodPost, tokenURL, strings.NewReader(formData.Encode()))
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create request: %w", err)
// 	}
// 	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 	res, err := c.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read response body: %w", err)
// 	}

// 	if res.StatusCode == http.StatusOK {
// 		var token TokenResponse
// 		if err := json.Unmarshal(body, &token); err != nil {
// 			return nil, fmt.Errorf("failed to unmarshal token: %w", err)
// 		}
// 		return &token, nil
// 	}

// 	if res.StatusCode == http.StatusForbidden {
// 		return nil, fmt.Errorf("res: %s body: %s", res.Status, string(body))
// 	}

// 	return nil, fmt.Errorf("unexpected polling error: %d, body: %s", res.StatusCode, body)
// }
