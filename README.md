# PingOne Go Client SDK

[![Go Tests](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-test.yml/badge.svg)](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-test.yml)
[![Go Linting](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-lint.yml/badge.svg)](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-lint.yml)
[![Go Security Scan](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-security.yml/badge.svg)](https://github.com/pingidentity/pingone-go-client/actions/workflows/go-security.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/pingidentity/pingone-go-client)](https://goreportcard.com/report/github.com/pingidentity/pingone-go-client)
[![GoDoc](https://godoc.org/github.com/pingidentity/pingone-go-client?status.svg)](https://godoc.org/github.com/pingidentity/pingone-go-client)

The **PingOne Go Client SDK** provides functions and structures to facilitate interaction with the PingOne Management API.

> [!IMPORTANT]  
> Content in this repository is under active development and is being built out incrementally. If there are operations you need that aren't yet available, please raise an issue on this project.

## ✨ Why use the PingOne Go SDK?

This SDK simplifies working with PingOne's user and management APIs by offering:

* **Built-in Best Practices**: Implements automatic retries for transient errors and handles platform eventual consistency.
* **Telemetry Support**: Allows telemetry headers for improved monitoring and troubleshooting.
* **Deprecation Policy Compliance**: Adheres to Ping's deprecation policy for smoother API transitions.
* **Forward Compatibility**: Follows documented best practices to reduce breaking changes with new platform features.
* **Type Safety**: Provides strong typing for requests and responses, minimizing runtime errors.
* **Authentication Handling**: Manages OAuth2 token acquisition and renewal automatically.


## 🛠️ Prerequisites

  * Go 1.24 or later (current go.mod specifies 1.24.1)
  * Go 1.21+ required for `log/slog` functionality
  * Go 1.23+ required for `iter` package features used in pagination examples


## 🚀 Getting Started

### Installation

Install the SDK using the standard `go get` command:

```shell
go get github.com/pingidentity/pingone-go-client
```


### Client Initialization

You can initialize the API client in a few ways:

1.  **Explicit Configuration:**
    Provide service configuration details directly in your code.

    *Example (Worker App without Custom Domain):*

    ```go
    package main

    import (
        "log/slog"
        "os"

        "github.com/pingidentity/pingone-go-client/config"
        "github.com/pingidentity/pingone-go-client/oauth2"
        "github.com/pingidentity/pingone-go-client/pingone"
    )

    func main() {
        serviceCfg := config.NewConfiguration()
        serviceCfg.WithGrantType(oauth2.GrantTypeClientCredentials)
        serviceCfg.WithClientID("YOUR_CLIENT_ID") // Replace with your actual Client ID
        serviceCfg.WithClientSecret("YOUR_CLIENT_SECRET") // Replace with your actual Client Secret
        serviceCfg.WithAuthEnvironmentID("YOUR_AUTH_ENVIRONMENT_ID") // e.g., 5dfe1e9f-95ce-43ed-b3c6-6f2f0a35ede5
        serviceCfg.WithRootDomain("pingone.com") // or pingone.eu, pingone.asia

        configuration := pingone.NewConfiguration(serviceCfg)
        client, err := pingone.NewAPIClient(configuration)
        if err != nil {
            slog.Error("Failed to create client", "error", err)
            os.Exit(1)
        }

        slog.Info("PingOne API Client initialized successfully!")
        // Your client is ready to use
    }
    ```

    *Example (With Custom Domain):*
    If you have a custom domain configured in your PingOne environment:

    ```go
    // ... (imports similar to above)
    func main() {
        serviceCfg := config.NewConfiguration()
        serviceCfg.WithGrantType(oauth2.GrantTypeClientCredentials)
        serviceCfg.WithClientID("YOUR_CLIENT_ID")
        serviceCfg.WithClientSecret("YOUR_CLIENT_SECRET")
        serviceCfg.WithCustomDomain("auth.yourcustomdomain.com") // Replace with your custom domain

        configuration := pingone.NewConfiguration(serviceCfg)
        client, err := pingone.NewAPIClient(configuration)
        if err != nil {
            slog.Error("Failed to create client", "error", err)
            os.Exit(1)
        }
        slog.Info("PingOne API Client with custom domain initialized successfully!")
    }
    ```

2.  **Configuration via Environment Variables:**
    Initialize the client without explicit parameters. The SDK will look for configuration values in environment variables (see [Service Configuration](#-pingone-service-configuration) for details).

    *Example:*

    ```go
    package main

    import (
        "log/slog"
        "os"

        "github.com/pingidentity/pingone-go-client/pingone"
    )

    func main() {
        // Ensure PINGONE_CLIENT_ID, PINGONE_CLIENT_SECRET, etc., are set in your environment
        cfg := pingone.NewConfiguration(nil) // Pass nil to use environment variables
        client, err := pingone.NewAPIClient(cfg)
        if err != nil {
            slog.Error("Failed to create client from environment variables", "error", err)
            os.Exit(1)
        }
        slog.Info("PingOne API Client initialized successfully using environment variables!")
    }
    ```


## ⚙️ Client Configuration

Customize the SDK's behavior beyond the initial service setup:

```go
package main

import (
    "log/slog"
    "net/http"
    "net/url"
    "os"
    "time"

    "github.com/pingidentity/pingone-go-client/config"
    "github.com/pingidentity/pingone-go-client/oauth2"
    "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
    // Example: Initialize with some explicit service config
    serviceCfg := config.NewConfiguration()
    serviceCfg.WithGrantType(oauth2.GrantTypeClientCredentials)
    serviceCfg.WithClientID(os.Getenv("PINGONE_CLIENT_ID")) // Prefer loading sensitive data from env
    serviceCfg.WithClientSecret(os.Getenv("PINGONE_CLIENT_SECRET"))
    serviceCfg.WithAuthEnvironmentID(os.Getenv("PINGONE_ENVIRONMENT_ID"))
    serviceCfg.WithRootDomain("pingone.com")

    cfg := pingone.NewConfiguration(serviceCfg)

    // Set custom HTTP timeout
    cfg.HTTPClient.Timeout = 30 * time.Second

    // Set a proxy
    proxyURLString := "http://proxy.example.com:8080"
    proxyURL, err := url.Parse(proxyURLString)
    if err != nil {
        slog.Error("Failed to parse proxy URL", "url", proxyURLString, "error", err)
    } else {
        cfg.HTTPClient.Transport = &http.Transport{
            Proxy: http.ProxyURL(proxyURL),
        }
    }

    // Add custom default headers for all requests
    cfg.DefaultHeader["X-Custom-Header"] = "my-custom-value"

    // Append a string to the default User-Agent header
    cfg.AppendUserAgent("my-application/1.2.3")

    client, err := pingone.NewAPIClient(cfg)
    if err != nil {
        slog.Error("Failed to create customized client", "error", err)
        os.Exit(1)
    }

    slog.Info("Customized PingOne API Client initialized successfully!", "userAgent", cfg.UserAgent)
    // Client is ready with custom configurations
}
```


## 🔧 Configuration Methods

The SDK provides the following configuration methods to customize your client:

| Method                    | Description                                        | Required                                  |
| :------------------------ | :------------------------------------------------- | :---------------------------------------- |
| `WithClientID`            | Sets the OAuth 2.0 Client ID for authentication    | **Yes** (unless set via env var)          |
| `WithClientSecret`        | Sets the OAuth 2.0 Client Secret                   | **Yes** (unless set via env var)          |
| `WithAuthEnvironmentID`   | Sets the PingOne Environment ID                    | **Yes\*** (unless using custom domain)     |
| `WithRootDomain`          | Sets the PingOne root domain (e.g., `pingone.com`) | No (defaults to `pingone.com`)            |
| `WithTopLevelDomain`      | Sets the PingOne top-level domain                  | No                                        |
| `WithCustomDomain`        | Sets your configured custom domain                 | **Yes\*** (unless using environment ID)    |
| `WithGrantType`           | Sets the OAuth 2.0 Grant Type                      | No (defaults to `CLIENT_CREDENTIALS`)     |
| `WithAccessToken`         | Sets a pre-existing Access Token                   | No                                        |
| `WithAPIDomain`           | Overrides the API service domain                   | No                                        |

**\*Note:** You must provide either `WithAuthEnvironmentID` (along with `WithRootDomain` if not using the default domain) or `WithCustomDomain`.

*Example:*

```go
serviceCfg := config.NewConfiguration()
serviceCfg.WithGrantType(oauth2.GrantTypeClientCredentials)
serviceCfg.WithClientID("YOUR_CLIENT_ID")
serviceCfg.WithClientSecret("YOUR_CLIENT_SECRET")
serviceCfg.WithAuthEnvironmentID("YOUR_ENVIRONMENT_ID")
serviceCfg.WithRootDomain("pingone.com")
```


## 🌍 PingOne Service Configuration

Configure the connection to PingOne services either programmatically (as shown in [Client Initialization](#-getting-started)) or through environment variables.

### Environment Variables

The SDK can be configured using the following environment variables:

| Environment Variable        | Description                                  | Required                                     |
| :-------------------------- | :------------------------------------------- | :------------------------------------------- |
| `PINGONE_CLIENT_ID`         | OAuth 2.0 Client ID for authentication       | **Yes** |
| `PINGONE_CLIENT_SECRET`     | OAuth 2.0 Client Secret for authentication   | **Yes** |
| `PINGONE_ENVIRONMENT_ID`    | PingOne Environment ID                       | **Yes\*** (or `PINGONE_CUSTOM_DOMAIN`)       |
| `PINGONE_ROOT_DOMAIN`       | PingOne root domain (e.g., `pingone.com`, `pingone.eu`, `pingone.asia`) | No (defaults to `pingone.com`) |
| `PINGONE_TOP_LEVEL_DOMAIN`  | PingOne top-level domain                     | No |
| `PINGONE_CUSTOM_DOMAIN`     | Your configured custom domain (e.g., `auth.example.com`) | **Yes\*** (or `PINGONE_ENVIRONMENT_ID`) |
| `PINGONE_AUTH_GRANT_TYPE`   | OAuth 2.0 Grant Type (e.g. `CLIENT_CREDENTIALS`) | No (defaults to `CLIENT_CREDENTIALS`)        |
| `PINGONE_API_ACCESS_TOKEN`  | A pre-existing Access Token to use for authentication. The SDK will not perform auth if this is set. | No                                           |
| `PINGONE_API_DOMAIN`        | Overrides the API service domain.            | No                                           |

**\*Note:** You must provide either `PINGONE_ENVIRONMENT_ID` (along with `PINGONE_ROOT_DOMAIN` if not using the default domain) or `PINGONE_CUSTOM_DOMAIN`.


## 📝 Logging

This SDK uses the standard [`log/slog`](https://pkg.go.dev/log/slog) package (introduced in Go 1.21) for logging. You can configure `slog` as needed. Using `slog` is recommended because request and response models implement the [`slog.LogValuer`](https://pkg.go.dev/log/slog#LogValuer) interface, which obfuscates sensitive data (like passwords) in log output.

*Example `slog` Configuration and Usage:*

```go
package main

import (
    "context"
    "log/slog"
    "os"
    // ... other necessary pingone imports
    "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
    opts := &slog.HandlerOptions{
        Level:     slog.LevelDebug, // Set desired log level
        AddSource: true,          // Include source file and line number
    }

    logger := slog.New(slog.NewTextHandler(os.Stdout, opts))
    slog.SetDefault(logger)

    // Assume client is initialized as shown previously
    cfg := pingone.NewConfiguration(nil) // Or your explicit config
    client, err := pingone.NewAPIClient(cfg)
    if err != nil {
        slog.Error("Failed to create client", "error", err)
        os.Exit(1)
    }

    // Example API Call (replace with actual IDs and client setup)
    var client *pingone.APIClient // Initialize this properly
    var environmentId string = "YOUR_ENVIRONMENT_ID"
    var variableId string = "YOUR_VARIABLE_ID"

    variable, httpResp, err := client.DaVinciVariablesApi.GetVariableById(context.Background(), environmentId, variableId).Execute()
    if err != nil {
        slog.Error("Failed to read variable", "error", err)
        // Check httpResp for more details if available
        if httpResp != nil {
            slog.Error("HTTP Response details", "status", httpResp.Status, "statusCode", httpResp.StatusCode)
        }
        os.Exit(1)
    }

    if httpResp.StatusCode != 200 {
        slog.Error("Failed to read variable", "http_response_status", httpResp.Status)
        os.Exit(1)
    }

    slog.Debug("Variable successfully read", "variable", variable) // 'variable' would be the first return value
}
```

## 📄 Pagination

The SDK uses the [`iter`](https://pkg.go.dev/iter) package (Go 1.23+) for easy processing of paginated API results.

*Example:*

```go
// Assume client and environmentId are initialized
pagedIterator, _, err := client.DaVinciVariablesApi.GetVariables(context.Background(), environmentId).Execute()
if err != nil {
    slog.Error("Failed to initiate variable listing", "error", err)
    return
}

pagesProcessed := 0
variablesRead := 0

for pageCursor, err := range pagedIterator { // Go 1.23 range over func iterator
    pageNumber := pagesProcessed + 1
    if err != nil {
        slog.Error("Failed to read variables page", "error", err, "page_number", pageNumber)
        break
    }

    if pageCursor.HTTPResponse.StatusCode != 200 {
        slog.Error("API error on page", "http_response_status", pageCursor.HTTPResponse.Status, "page_number", pageNumber)
        break
    }
        
    slog.Debug("Processing page of results", "page_number", pageNumber)

    if variables, ok := pageCursor.Data.Embedded.GetVariablesOk(); ok {
        for _, variable := range variables {
            slog.Debug("Variable successfully read", "variable_name", variable.GetName(), "page_number", pageNumber) // Example: GetName()
            variablesRead++
        }
    }
    pagesProcessed++
}
    
slog.Info("Finished processing variables", "pages_processed", pagesProcessed, "variables_read", variablesRead)
```


## ⚠️ Error Handling

The SDK returns custom error types for service errors. Use the standard `errors` package with `errors.As` (Go 1.13+) to handle specific error types.

*Example:*

```go
// Assume client, environmentId, variableId are initialized
variable, httpResp, err := client.DaVinciVariablesApi.GetVariableById(context.Background(), environmentId, variableId).Execute()
if err != nil {
    var notFoundErr *pingone.NotFoundError // Use pointer for errors.As with struct types
    var invalidRequestErr *pingone.InvalidRequestError

    if errors.As(err, &notFoundErr) {
        slog.Error("The variable was not found", "error_details", notFoundErr.GetDetails(), "original_error", notFoundErr.Error())
        // Add "not found" specific logic using details from notFoundErr
    } else if errors.As(err, &invalidRequestErr) {
        slog.Error("The API request was invalid", "error_details", invalidRequestErr.GetDetails(), "original_error", invalidRequestErr.Error())
        // Add "invalid request" specific logic using details from invalidRequestErr
    } else {
        // Handle other generic errors
        slog.Error("An unexpected error occurred", "error", err)
        if httpResp != nil {
            slog.Error("HTTP Response", "status", httpResp.Status)
        }
    }
    // General error handling logic for any error
    return
}
// Process 'variable'
```


## 📦 Packages

This SDK is organized into the following packages:

* **`config`**: Utilities for defining service configuration to connect to PingOne.
  * [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/config)
* **`pingone`**: The main SDK client and API models.
  * [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/pingone)
* **`oauth2`**: Helpers for OAuth 2.0 token acquisition.
  * [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/oauth2)
* **`oidc`**: Helpers for OpenID Connect token processes.
  * [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/oidc)
* **`testframework`**: Utilities for structuring tests against PingOne services.
  * [Documentation](https://pkg.go.dev/github.com/pingidentity/pingone-go-client/testframework)
