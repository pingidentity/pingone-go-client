.PHONY: test lint security build clean fmt help

# Default target
help:
	@echo "Available targets:"
	@echo "  test      - Run tests"
	@echo "  lint      - Run linting checks"
	@echo "  security  - Run security scan"
	@echo "  fmt       - Format code"
	@echo "  build     - Build the package"
	@echo "  clean     - Clean build artifacts"
	@echo "  help      - Show this help message"

# Run tests
test:
	go test ./config/... ./oauth2/... ./oidc/... -v

# Run integration tests (requires credentials)
test-integration:
	go test ./pingone/... -v

# Run all tests with coverage
test-coverage:
	go test -coverprofile=coverage.out ./config/... ./oauth2/... ./oidc/...
	go tool cover -html=coverage.out -o coverage.html

# Run linting checks
lint:
	golangci-lint run

# Run security scan
security:
	gosec -no-fail -fmt sarif -out results.sarif ./...

# Format code
fmt:
	gofmt -s -w .

# Build the package
build:
	go build -v ./...

# Clean build artifacts
clean:
	rm coverage.out coverage.html results.sarif
