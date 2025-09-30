# GitHub Copilot Instructions for PingOne Go Client SDK

## Code Documentation Standards

**Important Note:** The documentation standards outlined in this section apply only to manually written code in the core SDK packages (`config`, `oauth2`, `oidc`, `testframework`, etc.). The files in the `pingone/` directory are automatically generated from OpenAPI specifications and should not be manually modified or documented, as any changes will be overwritten by the code generator.

### Go Documentation (godoc) Best Practices

When writing or modifying Go code in the core SDK packages, ensure all documentation follows these godoc standards:

#### Package Documentation
- Add a package comment before the `package` declaration
- Start with "Package packagename" followed by a description
- Use complete sentences with proper punctuation
- Example:
  ```go
  // Package config provides configuration management for the PingOne Go Client SDK.
  package config
  ```

#### Function/Method Documentation
- Add comments directly above the function/method declaration
- Start the comment with the function name
- Use complete sentences in present tense
- Document comprehensively including:
  - The overall purpose of the function
  - What is returned from the function in human terms (e.g., "a string that represents the user's email address")
  - What is required from each of the inputs in human terms, providing context as to what the various parameters do for the function
  - External dependencies that can modify the logic (e.g., environment variables that must be set by the developer)
- Example:
  ```go
  // FullIsoList returns a slice of all available ISO language codes.
  // The returned slice contains string values representing standardized ISO language codes
  // such as "en" for English and "fr-CA" for French (Canada). No parameters are required.
  func FullIsoList() []string {
  ```
  
  More comprehensive example:
  ```go
  // ValidateEmailAddress checks if the provided email address is valid and accessible.
  // It returns a boolean indicating validity and an error if validation fails.
  // The email parameter must be a non-empty string containing a properly formatted email address.
  // The timeout parameter specifies the maximum duration in seconds to wait for validation.
  // This function requires the SMTP_SERVER environment variable to be set for external validation.
  func ValidateEmailAddress(email string, timeout int) (bool, error) {
  ```

#### Type Documentation
- Document all exported types (those starting with capital letters)
- Start the comment with the type name
- Explain the purpose and usage of the type
- Example:
  ```go
  // IsoCountry represents an ISO language code with its corresponding human-readable name.
  type IsoCountry struct {
  ```

#### Field Documentation
- Document exported struct fields
- Use concise but descriptive comments
- Example:
  ```go
  type IsoCountry struct {
      // Code is the ISO language code (e.g., "en", "fr-CA")
      Code string
      // Name is the human-readable language name (e.g., "English", "French (Canada)")
      Name string
  }
  ```

#### Variable/Constant Documentation
- Document exported variables and constants
- Group related constants with a single comment block when appropriate
- Example:
  ```go
  // reservedLanguageCodes contains ISO language codes that are reserved and cannot be used for custom languages.
  var reservedLanguageCodes = []string{
  ```

## Go SDK Specific Guidelines

### API Client Implementation
- Use the auto-generated API client code in the `pingone/` directory (do not modify these files directly)
- Implement SDK functionality in the core packages: `config`, `oauth2`, `oidc`, and `testframework`
- Use proper configuration patterns for environment variables and client setup
- Follow established patterns for OAuth2 authentication and token management

### Error Handling
- Use proper Go error handling patterns with explicit error returns
- Implement custom error types when appropriate for SDK-specific errors
- Provide clear error messages that help SDK users understand and resolve issues
- Use proper validation for input parameters

### Code Organization
- Place configuration logic in the `config/` package
- Place OAuth2 authentication logic in the `oauth2/` package  
- Place OIDC utilities in the `oidc/` package
- Place testing utilities in the `testframework/` package
- Do not modify auto-generated files in the `pingone/` directory

### Testing
- Write unit tests for all SDK functionality in core packages
- Use integration tests for testing against live PingOne APIs (requires credentials)
- Use the `testframework` package for integration test utilities
- Run `make test` for unit tests and `make test-integration` for integration tests
- Include `//go:build beta` tags for beta functionality testing

## General Go Best Practices

### Code Style
- Follow standard Go formatting (use `gofmt`)
- Use meaningful variable and function names
- Keep functions focused and concise
- Use proper error handling patterns
- Leverage Go modules for dependency management

### Comments and Documentation
- Use complete sentences with proper capitalization and punctuation
- Be comprehensive and descriptive, not just concise
- Explain the "why" not just the "what" when appropriate
- Document any non-obvious behavior or edge cases
- Use present tense in function descriptions ("returns", not "will return")
- For functions, always document:
  - Overall purpose and behavior
  - Return values in human-readable terms
  - Parameter requirements and their purpose
  - External dependencies (environment variables, configuration files, etc.)

### Import Organization
- Group imports logically (standard library, third-party, local)
- Use blank lines to separate import groups
- Avoid unnecessary imports

## Code Review Considerations

When suggesting code changes or improvements:
1. Ensure all exported items are properly documented
2. Verify adherence to the established SDK patterns
3. Check that error handling follows Go best practices
4. Confirm new code includes appropriate tests
5. Validate that API usage follows the established patterns
6. Ensure code follows Go best practices and idioms
7. Respect the separation between generated code (`pingone/`) and SDK code

## Additional Resources

- [Contributing Guidelines](../CONTRIBUTING.md)
- [Development Environment Setup](../contributing/development-environment.md)
- [SDK Design Guide](../contributing/sdk-design.md)
- [Beta Development Guide](../contributing/beta.md)
- [Testing Strategy](../contributing/acceptance-test-strategy.md)
- [PR Checklist](../contributing/pr-checklist.md)