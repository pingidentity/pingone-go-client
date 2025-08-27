# Contributing to pingone-go-client

Thank you for your interest in contributing to the pingone-go-client project! This document provides guidelines and instructions for contributing to this project.

## Code Generation and Contribution Scope

**Important:** The files in the `pingone` directory are automatically generated using [OpenAPI Generator](https://openapi-generator.tech/) from specifications found in the [PingOne OpenAPI Specifications](https://github.com/pingidentity/pingone-openapi-specifications) repository.

* ✅ **Welcome contributions:** All files outside the `pingone` directory
* ❌ **Will not merge:** Direct changes to files in the `pingone` directory (they will be overwritten)

You can still modify files in the `pingone` directory to test ideas or highlight issues, but these changes won't be merged to main. Instead, use them to demonstrate potential improvements that can be incorporated into the OpenAPI generation process.

## Development Workflow

### Setting Up Your Development Environment

1. Fork the repository on GitHub
2. Clone your fork locally
3. Install Go (version 1.24 or higher is required)
4. Run `go mod download` to install dependencies

### Testing Your Changes

Before submitting a pull request, make sure your changes pass all tests:

```shell
# Run unit tests
go test ./config/... ./oauth2/... ./oidc/... -v

# Run linting checks
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
golangci-lint run
```

### Continuous Integration

All pull requests are automatically tested using GitHub Actions. The following checks will run:

1. **Unit Tests**: Verifies that your code changes don't break existing functionality
2. **Linting**: Ensures code quality and consistency
3. **CodeQL Analysis**: Performs static code analysis to identify security vulnerabilities
4. **Security Scanning**: Uses gosec to detect potential security issues in the code

Make sure all CI checks pass before your pull request can be merged.

### Security Best Practices

When contributing code, please follow these security best practices:

- Avoid hardcoding sensitive information like credentials or tokens
- Use proper input validation for all external inputs
- Follow the principle of least privilege when implementing functionality
- Make sure all API endpoints are properly authenticated and authorized
- Document any security considerations for your code

## Pull Request Process

1. Create a branch for your changes
2. Make your changes with appropriate tests
3. Ensure all tests pass locally
4. Update documentation if needed
5. Submit a pull request
6. Address any feedback from reviewers

## Code Style

This project follows standard Go code style. Please ensure your code:

- Is properly formatted with `gofmt`
- Passes linting checks
- Includes appropriate comments and documentation
- Has meaningful test coverage

## Licensing

By contributing to this project, you agree that your contributions will be licensed under the project's license.
