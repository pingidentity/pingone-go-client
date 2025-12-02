// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"github.com/stretchr/testify/suite"

	"github.com/pingidentity/pingone-go-client/pingone"
)

// PingOneTestSuite provides a base test suite for PingOne API integration tests.
// It extends testify's Suite with a pre-configured PingOne API client that is
// automatically set up before tests run and cleaned up afterward.
type PingOneTestSuite struct {
	suite.Suite
	// APIClient is the configured PingOne API client for test operations.
	APIClient *pingone.APIClient
}

// SetupSuite initializes the test suite by creating a configured API client.
// This method is called once before all tests in the suite run. If the APIClient
// is not already set, it creates a new test client using the default configuration.
// The test suite will fail if client creation fails.
func (s *PingOneTestSuite) SetupSuite() {
	if s.APIClient == nil {
		var err error
		s.APIClient, err = TestClient(nil)
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
	s.APIClient = nil
}
