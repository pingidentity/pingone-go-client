/*
PingOne User and Configuration Management API

The PingOne User and Configuration Management API provides the interface to configure and manage users in the PingOne directory and the administration configuration of your PingOne organization.

Contact: developerexperiences@pingidentity.com
*/

package pingone

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
	"strings"

	"github.com/kelseyhightower/envconfig"
)

// contextKeys are used to identify the type of value in the context.
// Since these are string, it is possible to get a short description of the
// context key for logging and debugging using key.String().

type contextKey string

func (c contextKey) String() string {
	return "auth " + string(c)
}

var (
	// ContextOAuth2 takes an oauth2.TokenSource as authentication for the request.
	ContextOAuth2 = contextKey("token")

	// ContextAccessToken takes a string oauth2 access token as authentication for the request.
	ContextAccessToken = contextKey("accesstoken")

	// ContextServerIndex uses a server configuration from the index.
	ContextServerIndex = contextKey("serverIndex")

	// ContextOperationServerIndices uses a server configuration from the index mapping.
	ContextOperationServerIndices = contextKey("serverOperationIndices")

	// ContextServerVariables overrides a server configuration variables.
	ContextServerVariables = contextKey("serverVariables")

	// ContextOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextOperationServerVariables = contextKey("serverOperationVariables")

	// ContextAuthServerIndex uses a server configuration from the index.
	ContextAuthServerIndex = contextKey("authServerIndex")

	// ContextAuthOperationServerIndices uses a server configuration from the index mapping.
	ContextAuthOperationServerIndices = contextKey("serverAuthOperationIndices")

	// ContextAuthServerVariables overrides a server configuration variables.
	ContextAuthServerVariables = contextKey("authServerVariables")

	// ContextAuthOperationServerVariables overrides a server configuration variables using operation specific values.
	ContextAuthOperationServerVariables = contextKey("authServerOperationVariables")
)

// BasicAuth provides basic http authentication to a request passed via context using ContextBasicAuth
type BasicAuth struct {
	UserName string `json:"userName,omitempty"`
	Password string `json:"password,omitempty"`
}

// APIKey provides API key based authentication to a request passed via context using ContextAPIKey
type APIKey struct {
	Key    string
	Prefix string
}

// ServerVariable stores the information about a server variable
type ServerVariable struct {
	Description  string
	DefaultValue string
	EnumValues   []string
}

// ServerConfiguration stores the information about a server
type ServerConfiguration struct {
	URL         string
	Description string
	Variables   map[string]ServerVariable
}

// ServerConfigurations stores multiple ServerConfiguration items
type ServerConfigurations []ServerConfiguration

type ServiceConfiguration struct {
	PingOne *ServiceConfigurationPingOne `json:"pingone,omitempty"`
}

type ServiceConfigurationPingOne struct {
	Auth struct {
		ClientID     *string `envconfig:"PINGONE_CLIENT_ID" json:"clientId,omitempty"`
		ClientSecret *string `envconfig:"PINGONE_CLIENT_SECRET" json:"clientSecret,omitempty"`
		AccessToken  *string `envconfig:"PINGONE_API_ACCESS_TOKEN" json:"accessToken,omitempty"`
	} `json:"auth"`
	Endpoint struct {
		AuthEnvironmentID *string `envconfig:"PINGONE_ENVIRONMENT_ID" json:"environmentId,omitempty"`
		TopLevelDomain    *string `envconfig:"PINGONE_TOP_LEVEL_DOMAIN" json:"topLevelDomain,omitempty"`
		RootDomain        *string `envconfig:"PINGONE_ROOT_DOMAIN" json:"rootDomain,omitempty"`
		APIDomain         *string `envconfig:"PINGONE_API_DOMAIN" json:"apiDomain,omitempty"`
		AuthDomain        *string `envconfig:"PINGONE_AUTH_DOMAIN" json:"authDomain,omitempty"`
		CustomDomain      *string `envconfig:"PINGONE_CUSTOM_DOMAIN" json:"customDomain,omitempty"`
	} `json:"endpoint"`
}

// Configuration stores the configuration of the API client
type Configuration struct {
	Host                   string            `json:"host,omitempty"`
	Scheme                 string            `json:"scheme,omitempty"`
	DefaultHeader          map[string]string `json:"defaultHeader,omitempty"`
	UserAgent              string            `json:"userAgent,omitempty"`
	DefaultServerIndex     int               `json:"defaultServerIndex,omitempty"`
	DefaultAuthServerIndex int               `json:"defaultAuthServerIndex,omitempty"`
	ProxyURL               *string           `json:"proxyURL,omitempty"`
	Debug                  bool              `json:"debug,omitempty"`
	Servers                ServerConfigurations
	AuthServers            ServerConfigurations
	OperationServers       map[string]ServerConfigurations
	AuthOperationServers   map[string]ServerConfigurations
	HTTPClient             *http.Client
	Service                *ServiceConfiguration `json:"service,omitempty"`
}

// NewConfiguration returns a new Configuration object with builder pattern as param
func NewConfiguration() *Configuration {
	cfg := &Configuration{
		DefaultHeader:          make(map[string]string),
		UserAgent:              fmt.Sprintf("pingtools pingone-go-client/v0.0.1 (Go/%s; %s/%s)", runtime.Version(), runtime.GOOS, runtime.GOARCH),
		DefaultServerIndex:     0,
		DefaultAuthServerIndex: 0,
		Servers: ServerConfigurations{
			{
				URL:         "https://api.{sld}.{tld}/{basePath}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"basePath": {
						Description:  "No description provided",
						DefaultValue: "v1",
					},
					"sld": {
						Description:  "No description provided",
						DefaultValue: "pingone",
					},
					"tld": {
						Description:  "No description provided",
						DefaultValue: "com",
						EnumValues: []string{
							"eu",
							"com",
							"asia",
							"com.au",
							"ca",
						},
					},
				},
			},
			{
				URL:         "https://{domain}/{basePath}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"basePath": {
						Description:  "No description provided",
						DefaultValue: "v1",
					},
					"domain": {
						Description:  "No description provided",
						DefaultValue: "api.pingone.com",
					},
				},
			},
		},
		AuthServers: ServerConfigurations{
			{
				URL:         "https://auth.{sld}.{tld}/{environmentId}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"environmentId": {
						Description: "No description provided",
					},
					"sld": {
						Description:  "No description provided",
						DefaultValue: "pingone",
					},
					"tld": {
						Description:  "No description provided",
						DefaultValue: "com",
						EnumValues: []string{
							"eu",
							"com",
							"asia",
							"com.au",
							"ca",
						},
					},
				},
			},
			{
				URL:         "https://{domain}/{environmentId}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"environmentId": {
						Description: "No description provided",
					},
					"domain": {
						Description:  "No description provided",
						DefaultValue: "auth.pingone.com",
					},
				},
			},
			{
				URL:         "https://{domain}",
				Description: "No description provided",
				Variables: map[string]ServerVariable{
					"domain": {
						Description:  "No description provided",
						DefaultValue: "auth.bxretail.org",
					},
				},
			},
		},
		OperationServers:     map[string]ServerConfigurations{},
		AuthOperationServers: map[string]ServerConfigurations{},
	}

	// Load environment variables
	if err := envconfig.Process("", cfg); err != nil {
		slog.Error("Failed to process environment variables", "error", err)
	}

	return cfg
}

func (c *Configuration) AppendUserAgent(userAgent string) {
	c.UserAgent += fmt.Sprintf(" %s", userAgent)
}

func (c *Configuration) SetDefaultServerIndex(defaultServerIndex int) {
	c.DefaultServerIndex = defaultServerIndex
}

func (c *Configuration) SetDefaultServerVariableDefaultValue(variable string, value string) error {
	return c.SetServerVariableDefaultValue(c.DefaultServerIndex, variable, value)
}

func (c *Configuration) SetServerVariableDefaultValue(serverIndex int, variable string, value string) error {
	if serverIndex >= 0 && serverIndex < len(c.Servers) {
		if entry, ok := c.Servers[serverIndex].Variables[variable]; ok {
			entry.DefaultValue = value
			c.Servers[serverIndex].Variables[variable] = entry
			return nil
		} else {
			return fmt.Errorf("variable %v not defined in server %v", variable, serverIndex)
		}
	} else {
		return fmt.Errorf("server index %v out of range %v", serverIndex, len(c.Servers)-1)
	}
}

func (c *Configuration) SetDefaultAuthServerIndex(defaultAuthServerIndex int) {
	c.DefaultAuthServerIndex = defaultAuthServerIndex
}

func (c *Configuration) SetDefaultAuthServerVariableDefaultValue(variable string, value string) error {
	return c.SetAuthServerVariableDefaultValue(c.DefaultAuthServerIndex, variable, value)
}

func (c *Configuration) SetAuthServerVariableDefaultValue(serverIndex int, variable string, value string) error {
	if serverIndex >= 0 && serverIndex < len(c.AuthServers) {
		if entry, ok := c.AuthServers[serverIndex].Variables[variable]; ok {
			entry.DefaultValue = value
			c.AuthServers[serverIndex].Variables[variable] = entry
			return nil
		} else {
			return fmt.Errorf("variable %v not defined in server %v", variable, serverIndex)
		}
	} else {
		return fmt.Errorf("server index %v out of range %v", serverIndex, len(c.AuthServers)-1)
	}
}

// AddDefaultHeader adds a new HTTP header to the default header in the request
func (c *Configuration) AddDefaultHeader(key string, value string) {
	c.DefaultHeader[key] = value
}

func (c *Configuration) WithAuthEnvironmentID(environmentID string) *Configuration {
	c.Service.PingOne.Endpoint.AuthEnvironmentID = &environmentID
	return c
}

func (c *Configuration) WithClientID(clientID string) *Configuration {
	c.Service.PingOne.Auth.ClientID = &clientID
	return c
}

func (c *Configuration) WithClientSecret(clientSecret string) *Configuration {
	c.Service.PingOne.Auth.ClientSecret = &clientSecret
	return c
}

func (c *Configuration) WithAccessToken(accessToken string) *Configuration {
	c.Service.PingOne.Auth.AccessToken = &accessToken
	return c
}

func (c *Configuration) WithTopLevelDomain(tld string) *Configuration {
	c.Service.PingOne.Endpoint.TopLevelDomain = &tld
	return c
}

func (c *Configuration) WithRootDomain(rootDomain string) *Configuration {
	c.Service.PingOne.Endpoint.RootDomain = &rootDomain
	return c
}

func (c *Configuration) WithAPIDomain(apiDomain string) *Configuration {
	c.Service.PingOne.Endpoint.APIDomain = &apiDomain
	return c
}

func (c *Configuration) WithAuthDomain(authDomain string) *Configuration {
	c.Service.PingOne.Endpoint.AuthDomain = &authDomain
	return c
}

func (c *Configuration) WithCustomDomain(customDomain string) *Configuration {
	c.Service.PingOne.Endpoint.CustomDomain = &customDomain
	return c
}

// URL formats template on a index using given variables
func (sc ServerConfigurations) URL(index int, variables map[string]string) (string, error) {
	if index < 0 || len(sc) <= index {
		return "", fmt.Errorf("index %v out of range %v", index, len(sc)-1)
	}
	server := sc[index]
	url := server.URL

	// go through variables and replace placeholders
	for name, variable := range server.Variables {
		if value, ok := variables[name]; ok {
			found := bool(len(variable.EnumValues) == 0)
			for _, enumValue := range variable.EnumValues {
				if value == enumValue {
					found = true
				}
			}
			if !found {
				return "", fmt.Errorf("the variable %s in the server URL has invalid value %v. Must be %v", name, value, variable.EnumValues)
			}
			url = strings.Replace(url, "{"+name+"}", value, -1)
		} else {
			url = strings.Replace(url, "{"+name+"}", variable.DefaultValue, -1)
		}
	}
	return url, nil
}

// ServerURL returns URL based on server settings
func (c *Configuration) ServerURL(index int, variables map[string]string) (string, error) {
	return c.Servers.URL(index, variables)
}

func getServerIndex(ctx context.Context, defaultServerIndex int) (int, error) {
	si := ctx.Value(ContextServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return defaultServerIndex, reportError("Invalid type %T should be int", si)
	}
	return defaultServerIndex, nil
}

func getAuthServerIndex(ctx context.Context, defaultServerIndex int) (int, error) {
	si := ctx.Value(ContextAuthServerIndex)
	if si != nil {
		if index, ok := si.(int); ok {
			return index, nil
		}
		return defaultServerIndex, reportError("Invalid type %T should be int", si)
	}
	return defaultServerIndex, nil
}

func getServerOperationIndex(ctx context.Context, endpoint string, defaultServerIndex int) (int, error) {
	osi := ctx.Value(ContextOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return defaultServerIndex, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getServerIndex(ctx, defaultServerIndex)
}
func getAuthServerOperationIndex(ctx context.Context, endpoint string, defaultServerIndex int) (int, error) {
	osi := ctx.Value(ContextAuthOperationServerIndices)
	if osi != nil {
		if operationIndices, ok := osi.(map[string]int); !ok {
			return defaultServerIndex, reportError("Invalid type %T should be map[string]int", osi)
		} else {
			index, ok := operationIndices[endpoint]
			if ok {
				return index, nil
			}
		}
	}
	return getAuthServerIndex(ctx, defaultServerIndex)
}

func getServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getAuthServerVariables(ctx context.Context) (map[string]string, error) {
	sv := ctx.Value(ContextAuthServerVariables)
	if sv != nil {
		if variables, ok := sv.(map[string]string); ok {
			return variables, nil
		}
		return nil, reportError("ctx value of ContextAuthServerVariables has invalid type %T should be map[string]string", sv)
	}
	return nil, nil
}

func getServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getServerVariables(ctx)
}

func getAuthServerOperationVariables(ctx context.Context, endpoint string) (map[string]string, error) {
	osv := ctx.Value(ContextAuthOperationServerVariables)
	if osv != nil {
		if operationVariables, ok := osv.(map[string]map[string]string); !ok {
			return nil, reportError("ctx value of ContextAuthOperationServerVariables has invalid type %T should be map[string]map[string]string", osv)
		} else {
			variables, ok := operationVariables[endpoint]
			if ok {
				return variables, nil
			}
		}
	}
	return getAuthServerVariables(ctx)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) ServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.OperationServers[endpoint]
	if !ok {
		sc = c.Servers
	}

	if ctx == nil {
		return sc.URL(c.DefaultServerIndex, nil)
	}

	index, err := getServerOperationIndex(ctx, endpoint, c.DefaultServerIndex)
	if err != nil {
		return "", err
	}

	variables, err := getServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}

// ServerURLWithContext returns a new server URL given an endpoint
func (c *Configuration) AuthServerURLWithContext(ctx context.Context, endpoint string) (string, error) {
	sc, ok := c.AuthOperationServers[endpoint]
	if !ok {
		sc = c.AuthServers
	}

	if ctx == nil {
		return sc.URL(c.DefaultAuthServerIndex, nil)
	}

	index, err := getAuthServerOperationIndex(ctx, endpoint, c.DefaultAuthServerIndex)
	if err != nil {
		return "", err
	}

	variables, err := getAuthServerOperationVariables(ctx, endpoint)
	if err != nil {
		return "", err
	}

	return sc.URL(index, variables)
}
