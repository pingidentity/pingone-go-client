# SDK Design

This document provides an architectural and design overview of the PingOne Go Client SDK, outlining its structure, design principles, and development patterns.

## Overview

The PingOne Go Client SDK is a comprehensive Go library for interacting with PingOne identity and access management APIs. Built using OpenAPI Generator and modern Go patterns, it provides type-safe, well-documented access to PingOne services through a well-structured, maintainable codebase that follows Go best practices and includes automatic code generation from OpenAPI specifications.

## Architecture

### High-Level Architecture

The SDK follows a layered architecture pattern with clear separation of concerns:

```
┌─────────────────────────────────────────────────────────────┐
│                    Client Applications                       │
│  - Go applications using the PingOne SDK                    │
│  - Custom integrations and automations                      │
│  - User management applications                             │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│               SDK Public API (pingone package)              │
│  - APIClient initialization and configuration               │
│  - Service-specific API interfaces                          │
│  - Generated models and request/response types              │
│  - Error handling and response processing                   │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│            Configuration Layer (config package)             │
│  - Service configuration management                         │
│  - Environment variable parsing                             │
│  - OAuth2 credential configuration                          │
│  - Domain and endpoint configuration                        │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│          Authentication Layer (oauth2 package)              │
│  - OAuth2 token management                                  │
│  - Client credentials flow                                  │
│  - Token refresh and lifecycle                              │
│  - Authentication method abstractions                       │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│             OIDC Utilities (oidc package)                   │
│  - OpenID Connect helpers                                   │
│  - Token validation utilities                               │
│  - Endpoint discovery helpers                               │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│           Testing Framework (testframework package)          │
│  - Test environment utilities                               │
│  - Integration test helpers                                 │
│  - Mock and fixture management                              │
│  - Acceptance test patterns                                 │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                 HTTP Transport Layer                        │
│  - HTTP client configuration                                │
│  - Request/response middleware                               │
│  - Retry logic and circuit breakers                         │
│  - Logging and telemetry                                    │
└─────────────────────────────────────────────────────────────┘
                                │
                                ▼
┌─────────────────────────────────────────────────────────────┐
│                    PingOne APIs                             │
│  - Management API                                           │
│  - Service-specific APIs (DaVinci, Authorize, etc.)         │
│  - Multi-region endpoint support                            │
└─────────────────────────────────────────────────────────────┘
```

### Core Components

#### 1. PingOne Package (`pingone/`)

**Purpose**: Main SDK client and API models (generated from OpenAPI specifications)

**Key Features**:
- Auto-generated client code from OpenAPI specifications
- Type-safe request and response models
- Service-specific API implementations
- Built-in error handling and response types
- Support for pagination using Go 1.23+ iter package

#### 2. Configuration Package (`config/`)

**Purpose**: Service configuration management and environment setup

**Key Features**:
- Environment variable parsing with `envconfig`
- OAuth2 credential configuration
- Domain and endpoint configuration
- Multi-region support (pingone.com, pingone.eu, pingone.asia)
- Custom domain support

#### 3. OAuth2 Package (`oauth2/`)

**Purpose**: OAuth2 authentication and token management

**Key Features**:
- Client credentials grant type implementation
- Automatic token renewal and lifecycle management
- Authentication method abstractions
- Support for custom authentication flows

#### 4. OIDC Package (`oidc/`)

**Purpose**: OpenID Connect utilities and helpers

**Key Features**:
- Token validation utilities
- Endpoint discovery helpers
- OIDC-specific functionality

#### 5. Test Framework Package (`testframework/`)

**Purpose**: Testing utilities and integration test support

**Key Features**:
- Environment setup for integration tests
- Test client configuration helpers
- Acceptance test patterns
- Mock and fixture management utilities

**Note**: This document describes the patterns for the Go SDK implementation. The `pingone` package contains auto-generated code from OpenAPI specifications and should not be manually modified (as noted in the contributing guidelines).

### Design Principles

#### 1. **Separation of Concerns**

- **Package Isolation**: Each package has a clear, single responsibility
- **Layer Separation**: Distinct layers for configuration, authentication, API interaction, and testing
- **Interface Abstraction**: Well-defined interfaces between components

#### 2. **Code Generation Strategy**

- **OpenAPI-Driven**: Primary API client code is generated from OpenAPI specifications
- **Immutable Generated Code**: Generated files in the `pingone` package are not manually edited
- **Spec-First Development**: API changes drive SDK updates through specification updates
- **Consistency**: Generated code ensures uniform patterns across all API endpoints

#### 3. **Type Safety**

- **Strong Typing**: Leverages Go's type system for compile-time safety
- **Generated Models**: Type-safe request and response models from OpenAPI specs
- **Interface Contracts**: Clear interfaces for all major components
- **Error Types**: Structured error types for different failure scenarios

#### 4. **Go Best Practices**

- **Standard Library**: Utilizes Go standard library patterns and idioms
- **Module Structure**: Follows Go module best practices and semantic versioning
- **Documentation**: Comprehensive documentation following Go documentation standards
- **Testing**: Idiomatic Go testing patterns with integration and unit tests

#### 5. **Developer Experience**

- **Simple Configuration**: Easy setup with environment variables or explicit configuration
- **Sensible Defaults**: Works out-of-the-box with minimal configuration
- **Comprehensive Examples**: Clear examples for common use cases
- **Error Messages**: Helpful error messages with actionable guidance

#### 6. **Authentication & Security**

- **OAuth2 Standards**: Full OAuth2 client credentials flow implementation
- **Token Management**: Automatic token refresh and lifecycle management
- **Secure Defaults**: Secure configuration defaults and best practices
- **Credential Protection**: Safe handling of sensitive authentication data

## SDK Package Reference

The PingOne Go SDK provides a comprehensive set of packages that simplify interaction with PingOne APIs. Each package serves a specific purpose in the overall SDK architecture and follows Go best practices for usability and maintainability.

### Configuration Package (`config`)

The `config` package provides utilities for configuring the SDK connection to PingOne services. It includes environment variable parsing, OAuth2 credential management, and service endpoint configuration.

**Key Features:**
- Environment variable parsing using `envconfig`
- Support for multiple PingOne regions (pingone.com, pingone.eu, pingone.asia)
- Custom domain configuration
- OAuth2 client credentials setup
- Validation of required configuration parameters

**Common Usage Patterns:**
- Initialize configuration from environment variables
- Set up explicit configuration with credentials
- Configure custom domains and regional endpoints
- Validate configuration completeness before client creation

### Main SDK Package (`pingone`)

The `pingone` package contains the auto-generated API client and models from OpenAPI specifications. This package provides type-safe access to all PingOne APIs and should not be manually modified.

**Key Features:**
- Auto-generated from OpenAPI specifications
- Type-safe request and response models
- Service-specific API client interfaces
- Built-in error handling with structured error types
- Support for pagination using Go 1.23+ iter package
- Comprehensive API coverage for all PingOne services

**Common Usage Patterns:**
- Initialize API client with configuration
- Make type-safe API calls with proper error handling
- Process paginated results using iterator patterns
- Handle service-specific error types appropriately

### OAuth2 Package (`oauth2`)

The `oauth2` package provides OAuth2 authentication utilities specifically designed for PingOne API access. It handles token acquisition, renewal, and management.

**Key Features:**
- Client credentials grant type implementation
- Automatic token refresh and lifecycle management
- Integration with Go's standard `golang.org/x/oauth2` package
- Support for custom authentication flows
- Token caching and reuse

**Common Usage Patterns:**
- Configure OAuth2 client credentials
- Handle automatic token renewal
- Implement custom authentication flows when needed
- Integrate with HTTP client for authenticated requests

### OIDC Package (`oidc`)

The `oidc` package provides OpenID Connect utilities for working with OIDC tokens and endpoints in the PingOne ecosystem.

**Key Features:**
- Token validation utilities
- Endpoint discovery helpers
- OIDC-specific functionality for PingOne
- Integration with standard OIDC libraries

**Common Usage Patterns:**
- Validate OIDC tokens from PingOne
- Discover OIDC endpoints for configuration
- Work with OIDC-specific PingOne features

### Test Framework Package (`testframework`)

The `testframework` package provides utilities for testing applications that use the PingOne SDK. It includes helpers for setting up test environments and managing test data.

**Key Features:**
- Test environment setup and configuration
- Integration test helpers and utilities
- Test client configuration management
- Mock and fixture support for testing
- Acceptance test patterns for PingOne integration

**Common Usage Patterns:**
- Set up test environments for integration testing
- Configure test clients with appropriate credentials
- Create and manage test data for acceptance tests
- Implement testing patterns for PingOne integrations

## Error Handling Patterns

The PingOne Go SDK provides structured error handling that allows applications to respond appropriately to different types of failures. Understanding these patterns is essential for building robust applications.

### Error Types

The SDK defines specific error types for common API failure scenarios:

**Key Error Types:**
- `BadRequestError`: Invalid request parameters or malformed data
- `UnauthorizedError`: Authentication failures or invalid credentials
- `ForbiddenError`: Authorization failures or insufficient permissions
- `NotFoundError`: Requested resources that don't exist
- `TooManyRequestsError`: Rate limiting and throttling errors
- `InternalServerError`: Service-side errors and outages

### Error Handling Best Practices

When working with the SDK, use Go's standard error handling patterns with type assertions:

**Recommended Patterns:**
- Use `errors.As()` to check for specific error types
- Handle authentication errors by refreshing tokens when appropriate
- Implement retry logic for transient errors like rate limiting
- Log error details appropriately while protecting sensitive information
- Provide meaningful error messages to end users

### Logging Integration

The SDK integrates with Go's standard `log/slog` package and implements `slog.LogValuer` interface for secure logging:

**Logging Features:**
- Automatic obfuscation of sensitive data in log output
- Structured logging support with request/response context
- Integration with standard Go logging frameworks
- Configurable log levels and output formats

## API Usage Patterns

The SDK follows consistent patterns across all PingOne APIs to provide a familiar developer experience.

### Client Initialization

All API interactions start with proper client initialization:

**Standard Pattern:**
1. Configure service settings (credentials, endpoints, etc.)
2. Create PingOne configuration object
3. Initialize API client with configuration
4. Use client for API operations

### Request/Response Handling

The SDK provides consistent request and response handling across all APIs:

**Common Patterns:**
- Use context for request lifecycle management
- Handle pagination using Go 1.23+ iterator patterns
- Process responses with proper error checking
- Extract data from response objects using getter methods

### Authentication Flow

The SDK handles OAuth2 authentication automatically:

**Authentication Process:**
- Configure client credentials during initialization
- SDK automatically acquires access tokens as needed
- Tokens are refreshed automatically before expiration
- Failed authentication triggers appropriate error responses

## SDK Documentation

The PingOne Go SDK uses standard Go documentation practices to provide comprehensive, accessible documentation for developers. The documentation approach follows Go conventions and integrates with the Go ecosystem's documentation tools.

### Documentation Architecture

The SDK documentation system consists of three main components that work together to provide comprehensive developer resources:

#### Go Package Documentation (`godoc`)

Standard Go documentation using doc comments and the `godoc` tool. This provides the primary API reference documentation that integrates with go.dev/pkg and other Go documentation tools.

**Features:**
- Comprehensive API reference with method signatures and descriptions
- Example code embedded in documentation comments
- Integration with go.dev/pkg for online documentation
- Searchable documentation through Go tooling
- Automatic cross-linking between packages and types

#### README and Usage Examples (`/examples`)

Practical examples and getting started guides that demonstrate real-world usage patterns and best practices for the SDK.

**Structure:** `/examples/{use-case}/`
**Content Types:**
- Basic authentication and client setup
- Common API operation examples
- Integration patterns and best practices
- Complete application examples

#### Generated Documentation (`/docs` in pingone package)

Auto-generated documentation from OpenAPI specifications that provides detailed API reference material.

**Features:**
- Comprehensive API endpoint documentation
- Request/response schema documentation
- Error code and response documentation
- Generated from authoritative OpenAPI specifications

### Documentation Generation Process

#### Package Documentation

Standard Go documentation is maintained through doc comments in the source code:

```go
// Package pingone provides a client for the PingOne Management API.
//
// The client supports all PingOne services including user management,
// application configuration, and identity governance.
package pingone
```

#### Example Documentation

Examples are maintained as complete, runnable Go programs that demonstrate specific use cases:

```go
// Package main demonstrates basic client authentication and setup.
package main

// Example code here...
```

### Go Documentation Best Practices

The SDK follows Go documentation conventions to ensure consistency and discoverability:

#### Package-Level Documentation
- Clear package purpose and scope
- Overview of key concepts and usage patterns
- Links to relevant examples and guides
- Import path and version information

#### Function and Method Documentation
- Clear description of purpose and behavior
- Parameter descriptions and constraints
- Return value descriptions including error conditions
- Usage examples for complex functions

#### Type Documentation
- Purpose and usage of each type
- Field descriptions for structs
- Interface contract descriptions
- Example usage patterns

### Integration with Go Ecosystem

The SDK documentation integrates seamlessly with the broader Go ecosystem:

**Go.dev Integration:**
- Automatic publication to pkg.go.dev
- Search integration with Go package discovery
- Cross-linking with standard library documentation
- Version-aware documentation for different SDK releases

**IDE Integration:**
- IntelliSense and code completion support
- Hover documentation in editors
- Go language server integration
- Integrated example code and snippets

**Tooling Integration:**
- `go doc` command-line access
- Integration with `godoc` web server
- Documentation testing with `go test`
- Example validation and testing
