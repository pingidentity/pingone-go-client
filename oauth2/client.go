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
