// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"github.com/pingidentity/pingone-go-client/pingone"
	"github.com/stretchr/testify/suite"
)

// PingOneTestSuite provides a base test suite for PingOne API integration tests.
// It extends testify's Suite with a pre-configured PingOne API client that is
// automatically set up before tests run and cleaned up afterward.
type PingOneTestSuite struct {
	suite.Suite
	// ApiClient is the configured PingOne API client for test operations.
	ApiClient *pingone.APIClient
}

// SetupSuite initializes the test suite by creating a configured API client.
// This method is called once before all tests in the suite run. If the ApiClient
// is not already set, it creates a new test client using the default configuration.
// The test suite will fail if client creation fails.
func (s *PingOneTestSuite) SetupSuite() {
	if s.ApiClient == nil {
		var err error
		s.ApiClient, err = TestClient(nil)
		if err != nil {
			s.T().Fatalf("Failed to create API client: %v", err)
		}
	}
}

// SetupTest is called before each individual test method.
// This method is currently empty but can be overridden by test suites
// that need per-test setup logic.
func (s *PingOneTestSuite) SetupTest() {}

// TearDownTest is called after each individual test method.
// This method is currently empty but can be overridden by test suites
// that need per-test cleanup logic.
func (s *PingOneTestSuite) TearDownTest() {}

// TearDownSuite cleans up the test suite by clearing the API client.
// This method is called once after all tests in the suite have completed.
func (s *PingOneTestSuite) TearDownSuite() {
	s.ApiClient = nil
}
