# Testing Strategy

This document outlines the comprehensive testing strategy for the pingone-go-client SDK, covering both integration tests for the generated API code and unit tests for the core SDK packages. The testing approach differs based on the code location and purpose.

## Overview

The SDK uses a dual testing approach:

1. **Integration Testing**: For `pingone/` directory (generated API code) using the `testframework` package with real PingOne environments
2. **Unit Testing**: For core SDK packages (`config`, `oauth2`, `oidc`, `testframework`) using standard Go testing patterns

## Testing Architecture

### Integration Tests (`pingone/` directory)
- **Purpose**: Test generated API client code against live PingOne services
- **Framework**: Uses `testframework` package with testify test suites
- **Environment**: Real PingOne environments with proper credentials
- **Location**: `pingone/test/` directory

### Unit Tests (Core packages)
- **Purpose**: Test SDK utility functions, configuration, and authentication logic
- **Framework**: Standard Go testing with testify assertions
- **Environment**: Mocked dependencies, no external service calls
- **Location**: Individual package directories (`config/`, `oauth2/`, etc.)

## Test Organization

### File Structure

**Integration Tests (pingone/ directory):**
- **API service tests**: `api_<service>_test.go` in `pingone/test/` directory
- **Generated code tests**: Test auto-generated API client functionality

**Unit Tests (Core packages):**
- **Package tests**: `<package>_test.go` in respective package directories
- **Utility tests**: Test configuration, authentication, and framework utilities

### Package Structure

**Integration Tests:**
- Integration tests are placed in `pingone_test` package for generated API code
- Import testframework utilities from `github.com/pingidentity/pingone-go-client/testframework`
- Use shared test suites and environment configurations from testframework

**Unit Tests:**
- Unit tests are placed in individual package test files (e.g., `config_test`, `oauth2_test`)
- Use standard Go testing patterns with testify assertions
- Mock external dependencies and avoid real API calls

---

## Part 1: Integration Testing Strategy (pingone/ directory)

*This section covers testing for the auto-generated API client code in the `pingone/` directory.*

### 1. Test Suite Structure

The SDK uses testify test suites for organized, repeatable testing. There are three main test suite types:

#### **Base Test Suite (PingOneTestSuite)**
For basic API client functionality:
```go
type MyApiServiceTestSuite struct {
    testframework.PingOneTestSuite
}

func (s *MyApiServiceTestSuite) TestBasicOperation() {
    // Use s.ApiClient for API calls
    // Test basic CRUD operations
}
```

#### **Shared Environment Test Suite**
For tests that can share a single environment across multiple test cases:
```go
type MyApiServiceTestSuite struct {
    testframework.SharedEnvironmentTestSuite
}

func (s *MyApiServiceTestSuite) TestWithSharedEnvironment() {
    // Use s.TestEnvironment.Environment for environment operations
    // Environment persists across all tests in the suite
}
```

#### **New Environment Test Suite**
For tests that need a fresh environment for each test case:
```go
type MyApiServiceTestSuite struct {
    testframework.NewEnvironmentTestSuite
}

func (s *MyApiServiceTestSuite) TestWithFreshEnvironment() {
    // Use s.TestEnvironment.Environment for environment operations
    // New environment created for each test
}
```

### 2. API Testing Strategy

#### **CRUD Lifecycle Testing**
Test complete Create, Read, Update, Delete operations for API resources:

```go
func (s *MyApiServiceTestSuite) TestResourceFullLifecycle() {
    // Create resource with minimal schema
    resourceID := s.test_API_Create(s.T(), s.MinimalCreateRequest)
    
    // Read and verify resource 
    s.test_API_Get(s.T(), resourceID, s.MinimalCreateRequest)
    
    // Update resource to maximal schema
    s.test_API_Replace(s.T(), resourceID, s.MaximalReplaceRequest)
    
    // Delete resource
    s.test_API_Delete(s.T(), resourceID)
}
```

#### **Schema Validation Testing**
Test API request/response validation with different data combinations:

- **Minimal Schema**: Test with only required fields
- **Maximal Schema**: Test with all optional fields populated
- **Invalid Data**: Test API error responses for invalid inputs
- **Edge Cases**: Test boundary conditions and special values

#### **Error Condition Testing**
Test expected API error scenarios:

```go
func (s *MyApiServiceTestSuite) TestResourceNotFound() {
    nonExistentID := uuid.New()
    
    resp, httpRes, err := s.ApiClient.MyApi.GetResourceById(s.T().Context(), nonExistentID).Execute()
    testframework.CheckNotFound(s.T(), resp, httpRes, err)
    testframework.CheckPingOneAPIErrorResponse(s.T(), err, pingone.NotFoundError{}, regexp.MustCompile("Expected error pattern"))
}
```

#### **Authorization Testing**
Test API security and access control:

```go
func (s *MyApiServiceNoAuthzTestSuite) TestResourceNoAuthorization() {
    resourceID := uuid.New()
    
    resp, httpRes, err := s.BadApiClient.MyApi.GetResourceById(s.T().Context(), resourceID).Execute()
    
    require.NotNil(s.T(), err)
    require.Nil(s.T(), resp)
    require.NotNil(s.T(), httpRes)
    assert.Equal(s.T(), 401, httpRes.StatusCode)
    require.IsType(s.T(), &pingone.UnauthorizedError{}, err)
}
```

### 3. Environment Setup and Prerequisites

#### **Required Environment Variables**
Integration tests require specific environment variables to be set:

**Core Authentication:**
- `PINGONE_CLIENT_ID`: OAuth2 client ID for API access
- `PINGONE_CLIENT_SECRET`: OAuth2 client secret for API access  
- `PINGONE_ENVIRONMENT_ID`: Target PingOne environment ID
- `PINGONE_ENVIRONMENT_REGION`: PingOne region (e.g., "NA", "EU", "AP")

**Optional Configuration:**
- `PINGONE_LICENSE_ID`: Required for tests that create new environments
- `PINGONE_ROOT_DOMAIN`: PingOne domain (defaults to "pingone.com")
- `PINGONE_CUSTOM_DOMAIN`: Custom domain configuration (if applicable)

#### **Test Environment Strategy**

**Shared Environment Usage:**
Most tests should use a shared environment for efficiency. The testframework provides `SharedEnvironmentTestSuite` that creates one environment for the entire test suite:

```go
type MyApiServiceTestSuite struct {
    testframework.SharedEnvironmentTestSuite
}
```

**New Environment Usage:**
Use dedicated environments for tests that:
- Modify global environment configuration
- Test environment-level operations 
- Require specific environment settings

```go
type MyApiServiceTestSuite struct {
    testframework.NewEnvironmentTestSuite
}
```

### 4. Testframework Utilities

#### **Validation Functions**
The testframework provides helper functions for common test assertions:

```go
// Success assertions
testframework.CheckCreated(t, responseData, expectedType, httpResponse, error)
testframework.CheckFound(t, responseData, expectedType, httpResponse, error) 
testframework.CheckReplaced(t, responseData, expectedType, httpResponse, error)
testframework.CheckDeleted(t, httpResponse, error)

// Error assertions
testframework.CheckNotFound(t, responseData, httpResponse, error)
testframework.CheckPingOneAPIErrorResponse(t, error, expectedErrorType, errorMessageRegex)
```

#### **Environment Management**
The testframework provides utilities for managing test environments:

```go
// Create test environment
testEnv := testframework.NewTestEnvironment(environmentCreateRequest)
err := testEnv.Create(*testframework.CreateEnvironment(ctx, apiClient).IfNotExists())

// Delete test environment  
err := testEnv.Delete(*testframework.DeleteEnvironment(ctx, apiClient).IfExists())

// Default environment definition
envDef := testframework.DefaultEnvironmentDefinition(name, regionCode, licenseId, withBootstrap)
```

---

## Part 2: Unit Testing Strategy (Core packages)

*This section covers testing for core SDK packages: `config`, `oauth2`, `oidc`, and `testframework`.*

### 1. Configuration Package Testing (`config/`)

The config package handles service configuration, environment variable parsing, and client setup.

#### **Core Test Areas:**
- **Configuration Creation**: Test `NewConfiguration()` with various inputs
- **Environment Variable Parsing**: Test envconfig integration and validation
- **OAuth2 Configuration**: Test client credentials and token configuration  
- **Domain Configuration**: Test custom domains and regional endpoints
- **Validation**: Test configuration validation and error conditions

#### **Test Patterns:**
```go
func TestNewConfiguration(t *testing.T) {
    // Test default configuration
    cfg := config.NewConfiguration()
    assert.NotNil(t, cfg)
    // Validate default values
}

func TestConfigurationWithOptions(t *testing.T) {
    cfg := config.NewConfiguration()
    cfg.WithClientID("test-client-id")
    cfg.WithClientSecret("test-secret")
    cfg.WithAuthEnvironmentID("test-env-id")
    
    // Validate configuration was set correctly
    assert.Equal(t, "test-client-id", cfg.GetClientID())
}

func TestEnvironmentVariableParsing(t *testing.T) {
    // Set test environment variables
    os.Setenv("PINGONE_CLIENT_ID", "env-client-id")
    defer os.Unsetenv("PINGONE_CLIENT_ID")
    
    // Test configuration from environment
    // Validate proper parsing
}
```

#### **Test Coverage Requirements:**
- [ ] Configuration creation with all supported options
- [ ] Environment variable parsing for all supported variables
- [ ] Configuration validation and error handling
- [ ] Domain configuration (custom domains, regions)
- [ ] OAuth2 credential configuration
- [ ] Edge cases and invalid configurations

### 2. OAuth2 Package Testing (`oauth2/`)

The oauth2 package handles OAuth2 authentication flows and token management.

#### **Core Test Areas:**
- **Grant Types**: Test different OAuth2 grant type implementations
- **Token Authentication**: Test token acquisition and refresh logic
- **Endpoint Discovery**: Test OAuth2 endpoint configuration
- **Error Handling**: Test authentication failure scenarios

#### **Test Patterns:**
```go
func TestGrantTypeClientCredentials(t *testing.T) {
    // Test client credentials grant type
    grantType := oauth2.GrantTypeClientCredentials
    assert.Equal(t, "client_credentials", string(grantType))
}

func TestTokenAuthMethods(t *testing.T) {
    // Test token authentication method configuration
    // Mock HTTP client and test token requests
}

func TestEndpointConfiguration(t *testing.T) {
    // Test OAuth2 endpoint configuration for different regions
    // Validate endpoint URL construction
}
```

#### **Test Coverage Requirements:**
- [ ] All grant type constants and implementations
- [ ] Token authentication method configuration
- [ ] Endpoint URL construction for different regions
- [ ] Error handling for authentication failures
- [ ] Integration with standard golang.org/x/oauth2 package

### 3. OIDC Package Testing (`oidc/`)

The oidc package provides OpenID Connect utilities and endpoint discovery.

#### **Core Test Areas:**
- **Endpoint Discovery**: Test OIDC endpoint discovery mechanisms
- **Token Validation**: Test OIDC token validation utilities
- **Configuration**: Test OIDC-specific configuration options

#### **Test Patterns:**
```go
func TestOIDCEndpointDiscovery(t *testing.T) {
    // Test OIDC endpoint discovery
    // Mock HTTP responses for well-known endpoints
    // Validate endpoint parsing
}

func TestOIDCConfiguration(t *testing.T) {
    // Test OIDC configuration options
    // Validate configuration structures
}
```

#### **Test Coverage Requirements:**
- [ ] OIDC endpoint discovery functionality
- [ ] Token validation utilities
- [ ] Configuration structure validation
- [ ] Error handling for discovery failures
- [ ] Integration with standard OIDC libraries

### 4. Test Framework Package Testing (`testframework/`)

The testframework package provides testing utilities for integration tests.

#### **Core Test Areas:**
- **Test Client Creation**: Test API client creation for testing
- **Environment Management**: Test environment creation and cleanup
- **Test Suites**: Test suite initialization and teardown
- **Utility Functions**: Test helper functions and random data generation

#### **Test Patterns:**
```go
func TestTestClient(t *testing.T) {
    // Test API client creation for testing
    svcConfig := &config.Configuration{
        ClientID: "test-client",
    }
    
    client, err := testframework.TestClient(svcConfig)
    assert.NoError(t, err)
    assert.NotNil(t, client)
}

func TestEnvironmentCreation(t *testing.T) {
    // Test test environment creation (mocked)
    // Validate environment configuration
    // Test cleanup functionality
}

func TestRandomStringGeneration(t *testing.T) {
    // Test random string generation
    str1 := testframework.RandomResourceName()
    str2 := testframework.RandomResourceName()
    
    assert.NotEqual(t, str1, str2)
    assert.Len(t, str1, 10)
}
```

#### **Test Coverage Requirements:**
- [ ] Test client creation with various configurations
- [ ] Environment management utilities (mocked)
- [ ] Test suite lifecycle methods
- [ ] Random data generation utilities
- [ ] Validation and assertion helper functions

### Unit Testing Best Practices

#### **Mocking and Test Doubles:**
- Mock external dependencies (HTTP clients, API calls)
- Use test doubles for complex dependencies
- Avoid real network calls in unit tests
- Test both success and failure scenarios

#### **Test Organization:**
- Group related tests using subtests (`t.Run()`)
- Use table-driven tests for multiple input scenarios
- Separate unit tests from integration tests clearly
- Follow Go testing conventions

#### **Test Data Management:**
- Use consistent test data across related tests
- Generate random test data where appropriate
- Clean up test state between tests
- Use temporary files/directories for file operations

#### **Example Unit Test Structure:**
```go
func TestPackageFunction(t *testing.T) {
    tests := []struct {
        name     string
        input    InputType
        expected ExpectedType
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    validInput,
            expected: expectedOutput,
            wantErr:  false,
        },
        {
            name:    "invalid input",
            input:   invalidInput,
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := PackageFunction(tt.input)
            
            if tt.wantErr {
                assert.Error(t, err)
                return
            }
            
            assert.NoError(t, err)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

---

## Part 3: Integration Testing Strategy (pingone/ directory)

*This section covers testing for the auto-generated API client code in the `pingone/` directory.*

### Test Execution Commands

#### Run Integration Tests Only:
```bash
# Run all integration tests using testframework
make acc

# Run specific package integration tests
go test -v ./pingone -tags integration

# Run with specific environment
PINGONE_ENVIRONMENT_ID=your-env-id make acc
```

#### Run Unit Tests Only:
```bash
# Run all unit tests (excluding integration tests)
make test

# Run specific package unit tests
go test -v ./config
go test -v ./oauth2
go test -v ./oidc
go test -v ./testframework
```

### Integration Test Examples

### Complete API Service Test Structure

```go
package pingone_test

import (
    "testing"
    "github.com/pingidentity/pingone-go-client/testframework"
    "github.com/stretchr/testify/suite"
)

// Basic API tests without authorization
type MyApiServiceNoAuthzTestSuite struct {
    suite.Suite
    BadApiClient *pingone.APIClient
}

// Main API service tests
type MyApiServiceTestSuite struct {
    testframework.PingOneTestSuite
    
    // Test data setup
    MinimalCreateRequest pingone.ResourceCreateRequest
    MaximalCreateRequest pingone.ResourceCreateRequest
    MinimalReplaceRequest pingone.ResourceReplaceRequest
    MaximalReplaceRequest pingone.ResourceReplaceRequest
}

// Environment modification tests
type MyApiServiceModifyTestSuite struct {
    testframework.NewEnvironmentTestSuite
}

func (s *MyApiServiceTestSuite) TestResourceFullLifecycle() {
    // Test minimal schema lifecycle
    resourceID := s.test_API_Create(s.T(), s.MinimalCreateRequest)
    s.test_API_Delete(s.T(), resourceID)
    
    // Test maximal schema lifecycle 
    resourceID = s.test_API_Create(s.T(), s.MaximalCreateRequest)
    s.test_API_Replace(s.T(), resourceID, s.MinimalReplaceRequest)
    s.test_API_Replace(s.T(), resourceID, s.MaximalReplaceRequest)
    s.test_API_Delete(s.T(), resourceID)
}

func Test_MyApiService(t *testing.T) {
    suite.Run(t, &MyApiServiceTestSuite{})
    suite.Run(t, &MyApiServiceModifyTestSuite{})
}
```

## Quality Standards

### Test Coverage Requirements

All SDK functionality must have comprehensive integration test coverage:

- [ ] **API authorization testing** (unauthorized access scenarios)
- [ ] **CRUD lifecycle validation** with minimal and maximal data
- [ ] **Error condition testing** for expected API failures
- [ ] **Schema validation testing** with various data combinations
- [ ] **Environment management testing** (creation, modification, deletion)

#### **Additional Requirements for API Changes**
When modifying existing SDK functionality:
- [ ] **Backward compatibility testing** for any API behavior changes
- [ ] **Migration path validation** for deprecated functionality
- [ ] **Error message validation** for improved user experience

### Test Reliability

- Use `t.Parallel()` for parallel execution where appropriate
- Implement proper environment cleanup in test teardown
- Handle race conditions for environment-level operations
- Use appropriate timeouts for API operations
- Validate test stability across multiple runs

### Maintenance Considerations

- Keep test data DRY with shared setup methods
- Use descriptive test and variable names
- Document complex test scenarios and data dependencies
- Regular review and update of test patterns as APIs evolve
- Ensure tests remain valid as OpenAPI specifications change

## Examples

### Complete API Service Test Structure
```go
// Authorization failure tests
type MyApiServiceNoAuthzTestSuite struct { /* ... */ }

// Main API functionality tests  
type MyApiServiceTestSuite struct { /* ... */ }

// Environment-specific tests
type MyApiServiceModifyTestSuite struct { /* ... */ }

// Lifecycle tests
func (s *MyApiServiceTestSuite) TestResourceFullLifecycle() { /* ... */ }

// Error condition tests
func (s *MyApiServiceTestSuite) TestResourceNotFound() { /* ... */ }
```

### Integration Test Helper Methods
```go
func (s *TestSuite) test_API_Create(t *testing.T, payload CreateRequest) uuid.UUID { /* ... */ }
func (s *TestSuite) test_API_Get(t *testing.T, id uuid.UUID, expected interface{}) { /* ... */ }
func (s *TestSuite) test_API_Replace(t *testing.T, id uuid.UUID, payload ReplaceRequest) { /* ... */ }
func (s *TestSuite) test_API_Delete(t *testing.T, id uuid.UUID) { /* ... */ }
```

This comprehensive testing strategy ensures that all SDK functionality is thoroughly validated, providing confidence in the reliability and correctness of the pingone-go-client SDK.
