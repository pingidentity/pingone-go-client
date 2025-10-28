# Development Environment Setup

## Requirements

- [Go](https://golang.org/doc/install) 1.24+ (to build and test the SDK)
- [golangci-lint](https://golangci-lint.run/usage/install/) (for code linting)
- [gosec](https://github.com/securecodewarrior/gosec) (for security scanning, optional)

## Quick Start

If you wish to work on the SDK, you'll first need [Go](http://www.golang.org) installed on your machine (please check the [requirements](#requirements) before proceeding).

*Note:* This project uses [Go Modules](https://blog.golang.org/using-go-modules) making it safe to work with it outside of your existing [GOPATH](http://golang.org/doc/code.html#GOPATH). The instructions that follow assume a directory in your home directory outside of the standard GOPATH (i.e `$HOME/development/pingidentity/`).

Clone repository to: `$HOME/development/pingidentity/`

```sh
mkdir -p $HOME/development/pingidentity/; cd $HOME/development/pingidentity/
git clone git@github.com:pingidentity/pingone-go-client.git
cd pingone-go-client
```

To download dependencies, run:

```sh
go mod download
```

To build the SDK, run `make build`.

```sh
make build
```

## Testing the SDK

In order to test the SDK locally with no connection to PingOne, you can run `make test`.

```sh
make test
```

This will run unit tests for the core packages (config, oauth2, oidc) that don't require PingOne credentials.

For integration tests that connect to live PingOne services, you can run:

```sh
make test-integration
```

*Note:* Integration tests require valid PingOne credentials and create real API calls to PingOne services. Please ensure you have appropriate test environment access before running these tests.

To run tests with coverage reporting:

```sh
make test-coverage
```

## Code Quality Checks

### Linting

To run linting checks, first install golangci-lint:

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
```

Then run the linter:

```sh
make lint
```

### Security Scanning

To run security scans (optional):

```sh
make security
```

### Code Formatting

To format your code according to Go standards:

```sh
make fmt
```

## Using the SDK in Development

To use your locally built SDK in other Go projects, you have a few options:

### Using replace directive in go.mod

In your application's `go.mod` file, you can add a replace directive to use your local SDK version:

```go
module your-application

go 1.24

replace github.com/pingidentity/pingone-go-client => /path/to/your/local/pingone-go-client

require (
    github.com/pingidentity/pingone-go-client v0.x.x
    // ... other dependencies
)
```

### Testing SDK Changes

When developing the SDK, you can create test applications in the `examples/` directory to validate your changes work as expected.

## Important Notes on Generated Code

**Critical:** The files in the `pingone` directory are automatically generated using [OpenAPI Generator](https://openapi-generator.tech/) from specifications found in the [PingOne OpenAPI Specifications](https://github.com/pingidentity/pingone-openapi-specifications) repository.

- ✅ **Welcome contributions:** All files outside the `pingone` directory
- ❌ **Will not merge:** Direct changes to files in the `pingone` directory (they will be overwritten)

You can still modify files in the `pingone` directory to test ideas or highlight issues, but these changes won't be merged to main. Instead, use them to demonstrate potential improvements that can be incorporated into the OpenAPI generation process.

## Working with Dependencies

This SDK has minimal external dependencies to maintain simplicity and reliability. The main dependencies are:

- `github.com/google/uuid` - For UUID generation and validation
- `github.com/kelseyhightower/envconfig` - For environment variable configuration
- `golang.org/x/oauth2` - For OAuth2 authentication flows
- `github.com/stretchr/testify` - For testing utilities

### Updating Dependencies

To update dependencies, run:

```shell
go get -u ./...
go mod tidy
```

### Adding New Dependencies

When adding new dependencies, ensure they:

- Are actively maintained
- Have minimal security vulnerabilities
- Follow Go best practices
- Are necessary for the SDK's core functionality

After adding dependencies, run the full test suite to ensure compatibility:

```shell
make test
make lint
make security
```
