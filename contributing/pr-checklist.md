# Pull Request Checklist

The following provides the steps to check/run to prepare for creating a PR to the `main` branch. PRs that follow these checklists will merge faster than PRs that do not.

*Note: This checklist is designed to support both human contributors and automated code review tools.*

## For Automated Code Review

This checklist includes specific verification criteria marked with *Verification* that can be programmatically checked to support both manual and automated review processes

## PR Planning & Structure

- [ ] **PR Scope**. To ensure maintainer reviews are as quick and efficient as possible, please separate functionality into logical PRs. For example, support for a new API endpoint and related utilities can go in the same PR, however support for multiple unrelated API endpoints should be separated. It's acceptable to merge related changes into the same PR where structural improvements are being made.
  - *Verification*: Check that files modified are logically related (same package directory, related functionality)

- [ ] **PR Title**. To assist the maintainers in assessing PRs for priority, please provide a descriptive title of the functionality being supported. For example: `Add support for DaVinci Flow Policies API` or `Fix OAuth2 token refresh handling`
  - *Verification*: Title should be descriptive and match the type of changes (Add/Update/Fix/Remove)

- [ ] **PR Description**. Please follow the provided PR description template and check relevant boxes. Include a clear description of:
  - What functionality is being added/changed
  - Why the change is needed (e.g., to fix an issue - include the issue number as reference)
  - Any breaking changes
  - *Verification*: Check that PR description template boxes are completed and description sections are filled

## Code Development

### Architecture & Design

- [ ] **Code implementation**. New code should adhere to the [SDK Design](sdk-design.md) guide for implementation patterns.
  - *Verification*: 
    - New packages follow the established SDK architecture (config, oauth2, oidc, testframework)
    - Code follows Go best practices and package organization standards
    - Generated code in `pingone/` directory is not manually modified
    - Changes to non-generated packages follow consistent patterns

- [ ] **Generated Code Restrictions**. **Critical:** The files in the `pingone` directory are automatically generated from OpenAPI specifications and must not be manually modified.
  - *Verification*: 
    - No manual changes to files in `pingone/` directory
    - If generated code needs changes, issue should be created to address in OpenAPI generation process
    - All modifications are in `config/`, `oauth2/`, `oidc/`, `testframework/`, or other non-generated packages

### Code Quality

- [ ] **Dependencies Check**. Ensure go.mod and go.sum are properly maintained:

```shell
go mod tidy
```
*Verification*: Run command and verify no changes to go.mod/go.sum files

- [ ] **Build**. Verify the SDK builds successfully with your changes:

```shell
make build
```
*Verification*: Run command and verify exit code 0

- [ ] **Code Formatting**. Ensure code is properly formatted:

```shell
make fmt
```
*Verification*: Run command and verify no files are modified (clean git status)

- [ ] **Code Linting**. Run all linting checks to ensure code quality and consistency:

```shell
make lint
```
*Verification*: Command must exit with code 0

This includes:
- Go vet checks
- golangci-lint
- Import organization checks
- Go code formatting

## Testing

### Unit Tests

- [ ] **Unit Tests**. Where a code function performs work internally to a package, but has an external scope (i.e., a function with an initial capital letter `func MyFunction`), unit tests should ideally be created. Not all functions require a unit test, if in doubt please ask:

```shell
make test
```
*Verification*: Run command and verify exit code 0

### Integration Tests

- [ ] **Integration Tests**. Where new API functionality is being added or existing API interactions are being updated, integration tests should be created or modified to verify the SDK works correctly with live PingOne services:
  - *Verification*:
    - New API functionality has corresponding integration tests
    - Tests are in appropriate `*_test.go` files in the same package
    - Tests follow Go testing conventions and use `testing.T`
    - Tests include proper setup and cleanup

Example: To run integration tests (requires PingOne credentials):
```shell
make test-integration
```

To run all tests with coverage:
```shell
make test-coverage
```

- [ ] **Test Environment**. Ensure you have access to a PingOne trial or licensed environment for integration testing. *Integration tests create real API calls to PingOne services. Please ensure you have appropriate test environment access before running these tests.*

- [ ] **Test Credentials**. Ensure test credentials are properly configured via environment variables and not hardcoded in test files

## Documentation

### Code Documentation

- [ ] **Package Documentation**. Each package should have proper package-level documentation following Go conventions:
  - *Verification*: 
    - All packages have package comments that describe their purpose
    - Public functions and types have appropriate doc comments
    - Doc comments follow Go documentation standards (start with function/type name)
    - Documentation is clear and includes usage examples where helpful

- [ ] **Function Documentation**. Public functions should have clear documentation describing their purpose, parameters, return values, and any important behavior:
  - *Verification*: 
    - All exported functions have doc comments
    - Comments describe what the function does, not how it does it
    - Parameter and return value descriptions are included where non-obvious
    - Error conditions are documented where applicable

- [ ] **Error Documentation**. Custom error types and error handling should be well documented with actionable guidance for users:
  - *Verification*: Error types include clear descriptions and suggested resolution steps

### API Documentation

- [ ] **Generated Documentation**. The `pingone` package contains auto-generated documentation from OpenAPI specifications. Ensure this documentation remains intact and is not manually modified:
  - *Verification*: No manual changes to generated documentation in `pingone/` package

### Examples

- [ ] **Go Code Examples**. New/modified SDK functionality should have appropriate Go code examples created/altered and stored in the `examples` directory. These demonstrate practical usage patterns for SDK users:
  - *Verification*:
    - Examples exist in `examples/` directory with descriptive subdirectory names
    - Example files have `.go` extension and valid Go syntax
    - Examples demonstrate both basic and comprehensive usage where applicable
    - Examples include proper error handling and follow Go best practices
    - Each example includes a README.md explaining its purpose and usage

- [ ] **Documentation Examples**. Code examples in package documentation should be accurate and tested:
  - *Verification*: 
    - Doc examples compile and run correctly
    - Examples demonstrate common use cases
    - Examples are kept up-to-date with API changes

## Security & Compliance

- [ ] **Security Scan**. Ensure your code passes security scanning (this will be automatically checked in CI, but you can run locally):

```shell
make security
```
*Verification*: No high-severity security issues reported

- [ ] **Sensitive Data**. Ensure no sensitive data (API keys, tokens, etc.) are committed to the repository
  - *Verification*: 
    - No API keys, passwords, or tokens in code or test files
    - Sensitive test data uses environment variables
    - No `.env` files or similar containing credentials
    - Test examples use placeholder credentials or environment variable references

- [ ] **Input Validation**. Implement appropriate input validation for all user-provided data in SDK functions
  - *Verification*: 
    - Public functions validate input parameters
    - Error handling provides clear feedback for invalid inputs
    - No potential for injection attacks or unsafe operations

- [ ] **Dependency Security**. Ensure all dependencies are secure and up-to-date
  - *Verification*:
    - `go mod tidy` produces no warnings about vulnerabilities
    - Dependencies are from trusted sources
    - Minimal dependency footprint is maintained

## Final Checks

- [ ] **All Development Checks**. Run the comprehensive development check:

```shell
make test
make lint  
make security
```
*Verification*: All commands exit with code 0

- [ ] **CI Compatibility**. Verify your changes will pass automated CI checks by ensuring all the above steps pass locally
  - *Verification*: All previous verification steps completed successfully

- [ ] **Breaking Changes**. If your PR introduces breaking changes, ensure they are:
  - Clearly documented in the PR description
  - Included in the changelog entry  
  - Follow the project's versioning strategy
  - *Verification*: 
    - Breaking changes are documented in PR description
    - Changelog entry uses appropriate format for breaking changes
    - Backward compatibility considerations are addressed where possible

## Additional Notes

- The maintainers may run additional tests in different PingOne regions
- Large PRs may take longer to review - consider breaking them into smaller, focused changes
- If you're unsure about any step, please ask questions in your PR or create an issue for discussion

---

## Documentation-Only Changes

If you are making documentation-only changes (guides, examples, or README updates), you can use this simplified checklist:

- [ ] **Documentation Updates**. New or updated documentation should be clear, well-structured, and include practical examples

- [ ] **Example Updates**. Ensure any Go examples are syntactically correct and follow best practices:

```shell
go build ./examples/...
```

- [ ] **Formatting**. Ensure documentation follows proper formatting:

```shell
make fmt
```

Documentation changes are generally merged quicker than code changes as there is less to review.
