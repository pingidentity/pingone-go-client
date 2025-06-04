# pingone-go-client GO SDK

[![Go Tests](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-test.yml/badge.svg)](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-test.yml)
[![Go Linting](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-lint.yml/badge.svg)](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-lint.yml)
[![Go Security Scan](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-security.yml/badge.svg)](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-security.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pingidentity/pingone-go-client)](https://goreportcard.com/report/github.com/pingidentity/pingone-go-client)
[![GoDoc](https://godoc.org/github.com/pingidentity/pingone-go-client?status.svg)](https://godoc.org/github.com/pingidentity/pingone-go-client)

The pingone-go-client GO SDK provides a set of functions and stucts that help with interacting with the PingOne management API.

> [!IMPORTANT]  
> Content in this repository is under development and isn't yet functional.

## Getting started

Use the standard Go commands to install pingone-go-client to the project:

```shell
go get github.com/pingidentity/pingone-go-client
```

### Client initialisation

The client may be initialised by explicitly declared service configuration.  In this case, it's expected that service configuration values are provided by using environment variables according to the [service configuration parameters](#service-configuration-parameters) table below.

For example, using a worker application without a custom domain:

```go
package main

import (
    "log/slog"

    "github.com/pingidentity/pingone-go-client/config"
    "github.com/pingidentity/pingone-go-client/oauth2"
    "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
    serviceCfg := config.NewConfiguration()
    serviceCfg.WithGrantType(oauth2.GrantTypeClientCredentials)
    serviceCfg.WithClientID("595bdcd8-4e58-4094-ad19-132ee12d43f8")
    serviceCfg.WithClientSecret("***************")
    serviceCfg.WithAuthEnvironmentID("5dfe1e9f-95ce-43ed-b3c6-6f2f0a35ede5")
    serviceCfg.WithRootDomain("pingone.com")

    configuration := pingone.NewConfiguration(serviceCfg)
    client, err := pingone.NewAPIClient(cfg)
    if err != nil {
        slog.Error("Failed to create client", "error", err)
    }
}
```

If a custom domain has been configured in the environment, the custom domain can be applied during client initialisation instead of the environment ID and root domain (or top level domain).

For example:

```go
package main

import (
    "log/slog"

    "github.com/pingidentity/pingone-go-client/config"
    "github.com/pingidentity/pingone-go-client/oauth2"
    "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
    serviceCfg := config.NewConfiguration()
    serviceCfg.WithGrantType(oauth2.GrantTypeClientCredentials)
    serviceCfg.WithClientID("595bdcd8-4e58-4094-ad19-132ee12d43f8")
    serviceCfg.WithClientSecret("***************")
    serviceCfg.WithCustomDomain("auth.bxretail.org")

    configuration := pingone.NewConfiguration(serviceCfg)
    client, err := pingone.NewAPIClient(cfg)
    if err != nil {
        slog.Error("Failed to create client", "error", err)
    }
}
```

The client may be initialised without any declared service configuration.  In this case, it's expected that service configuration values are provided at runtime with environment variables according to the [PingOne service configuration parameters](#pingone-service-configuration) table below.  Client code projects don't need to define these environment variables explicitly in project code if the end user's environment has been set up with the appropriate environment variables exported in their shell.

Example when relying on environment variable values:

```go
package main

import (
    "log/slog"

    "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
    cfg := pingone.NewConfiguration(nil)
    client, err := pingone.NewAPIClient(cfg)
    if err != nil {
        slog.Error("Failed to create client", "error", err)
    }
}
```

## Client Configuration

The client configuration provides settings for how the SDK interacts with PingOne APIs. You can customize the configuration to suit your specific needs:

```go
cfg := pingone.NewConfiguration(serviceCfg)

// Set custom timeout
cfg.HTTPClient.Timeout = 30 * time.Second

// Set proxy
proxyURL, _ := url.Parse("http://proxy.example.com:8080")
cfg.HTTPClient.Transport = &http.Transport{
    Proxy: http.ProxyURL(proxyURL),
}

// Add custom headers
cfg.DefaultHeader["Custom-Header"] = "custom-value"

// Customize User-Agent
cfg.AppendUserAgent("my-application/1.0.0")
```

### User Agent Suffix

Strings can be applied as a suffix to the default User Agent string.  This can be achieved using the `AppendUserAgent()` function in the SDK client configuration.

For example:

```go
cfg := pingone.NewConfiguration(nil)
cfg.AppendUserAgent("my-custom-string")
client, err := pingone.NewAPIClient(cfg)
if err != nil {
    slog.Error("Failed to create client", "error", err)
}
```

## PingOne Service Configuration

The PingOne service configuration contains all the necessary settings to connect to PingOne services. You can provide these settings programmatically or through environment variables.

### Environment Variables

The following environment variables can be used to configure the client:

| Environment Variable | Description | Required |
|----------------------|-------------|----------|
| `PINGONE_CLIENT_ID` | OAuth 2.0 client ID for authentication | Yes |
| `PINGONE_CLIENT_SECRET` | OAuth 2.0 client secret for authentication | Yes |
| `PINGONE_ENVIRONMENT_ID` | PingOne environment ID | Yes* |
| `PINGONE_REGION` | PingOne region (e.g., `NA`, `EU`, `ASIA`) | No |
| `PINGONE_CUSTOM_DOMAIN` | Custom domain if configured | Yes* |

*Either `PINGONE_ENVIRONMENT_ID` or `PINGONE_CUSTOM_DOMAIN` must be provided.

## Logging

This Go client uses the standard [`log/slog` package](https://pkg.go.dev/log/slog) introduced in Go 1.21 to provide log output and can be configured according to slog documentation.

Using slog is recommended, as request and response models implement the [`LogValuer` interface](https://pkg.go.dev/log/slog#LogValuer) to obfuscate sensitive or password values in log output.

Use of slog is shown in this example:

```go
opts := &slog.HandlerOptions{
    Level:     slog.LevelDebug,
    AddSource: true,
}

logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
slog.SetDefault(logger)

cfg := pingone.NewConfiguration(nil)
client, err := pingone.NewAPIClient(cfg)
if err != nil {
    slog.Error("Failed to create client", "error", err)
}

variable, httpResp, err := client.DaVinciVariableApi.GetVariableById(context.Background(), environmentId, variableId).Execute()
if err != nil {
    slog.Error("Failed to read variable", "error", err)
    os.Exit(1)
}

if httpResp.StatusCode != 200 {
    slog.Error("Failed to read variable", "http response status", httpResp.Status)
    os.Exit(1)
}

slog.Debug("Variable successfully read", "variable", variable)
```

## Pagination

This Go client uses the standard [`iter` package](https://pkg.go.dev/iter) introduced in Go 1.23 to allow clients to easily process paginated results from supported APIs.

For example:

```go
pagedIterator := client.DaVinciVariableApi.GetVariables(context.Background(), environmentId).Execute()

// Set counters for the purpose of demonstration in log output
pagesProcessed := 0
variablesRead := 0

// Iterate over the pages
for pageCursor, err := range pagedIterator {
    pageNumber := pagesProcessed+1
    if err != nil {
        slog.Error("Failed to read variables", "error", err, "page number", pageNumber)
        break
    }

    if pageCursor.HTTPResponse.StatusCode != 200 {
        slog.Error("Failed to read variables", "http response status", pageCursor.HTTPResponse.Status, "page number", pageNumber)
        break
    }
        
    slog.Debug("Processing page of results", "page number", pageNumber)

    // Iterate over the variables in the page's collection
    if variables, ok := pageCursor.Data.Embedded.GetVariablesOk(); ok {
        for _, variable := range variables {
            slog.Debug("Variable successfully read", "variable", variable, "page number", pageNumber)
            variablesRead++
        }
    }

    pagesProcessed++
}
    
slog.Info("All variables over all pages have been read", "pages processed", pagesProcessed, "variables read", variablesRead)
```

## Error Handling

When service errors are returned from the Go SDK, custom error types are returned depending on the nature of the error. These errors can be handled using the standard `errors` package, using `errors.As` introduced in Go 1.13.

For example:

```go
variable, httpResp, err := client.DaVinciVariableApi.GetVariableById(context.Background(), environmentId, variableId).Execute()
if err != nil {
    var notFoundErr pingone.NotFoundError
    if errors.As(err, &notFoundErr) {
        slog.Error("The variable was not found", "error", notFoundErr)
            
        // Enter "not found" specific logic using detail attributes from the `pingone.NotFoundError` model
    }

    var invalidRequestErr pingone.InvalidRequestError
    if errors.As(err, &invalidRequestErr) {
        slog.Error("The API request was invalid", "error", invalidRequestErr)

        // Enter "invalid request" specific logic using detail attributes from the `pingone.InvalidRequestError` model
    }

    slog.Error("An error has occurred", "error", err)
    // Enter general error handling logic
}
```

## Packages

* **config** - Methods used when defining service configuration when connecting to PingOne tenants. [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/config)
* **pingone** - The main client SDK. [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/pingone)
* **oauth2** - Methods used in gaining OAuth2 access tokens. [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/oauth2)
* **oidc** - Methods used in gaining OpenID Connect tokens. [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/oidc)
* **testframework** - Methods used to structure project testing with PingOne services. [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/testframework)
