// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

// NewEnvironmentTestSuite extends PingOneTestSuite to provide automatic test environment creation.
// This suite creates a fresh PingOne environment for each test and cleans it up afterward,
// ensuring test isolation and preventing test interference.
type NewEnvironmentTestSuite struct {
	PingOneTestSuite
	// EnvironmentNamePrefix is prepended to generated environment names.
	EnvironmentNamePrefix string
	// EnvironmentNameSuffix is appended to generated environment names.
	EnvironmentNameSuffix string
	// TestEnvironment contains the created test environment for the current test.
	TestEnvironment *TestEnvironment
	// WithBootstrap determines whether the environment includes full DaVinci features.
	WithBootstrap bool
}

// SetupSuite initializes the test suite by calling the parent suite setup.
// This ensures the PingOne API client is properly configured before environment operations.
func (s *NewEnvironmentTestSuite) SetupSuite() {
	s.PingOneTestSuite.SetupSuite()
}

// SetupTest creates a new PingOne environment for each test method.
// This method requires the PINGONE_ENVIRONMENT_REGION and PINGONE_LICENSE_ID environment
// variables to be set. It generates a unique environment name using the configured prefix,
// random string, and suffix. The environment includes all PingOne services configured
// according to the WithBootstrap setting.
func (s *NewEnvironmentTestSuite) SetupTest() {
	s.PingOneTestSuite.SetupTest()

	if s.EnvironmentNamePrefix == "" {
		s.EnvironmentNamePrefix = "go-sdk-testacc-"
	}

	regionCode := os.Getenv("PINGONE_ENVIRONMENT_REGION")

	licenseID, err := uuid.Parse(os.Getenv("PINGONE_LICENSE_ID"))
	if err != nil {
		s.FailNow("Failed to parse license ID as a valid UUID (PingOne resource ID)", err)
	}

	testEnvironment := NewTestEnvironment(
		DefaultEnvironmentDefinition(
			fmt.Sprintf("%s%s%s", s.EnvironmentNamePrefix, randomString(10), s.EnvironmentNameSuffix),
			regionCode,
			licenseID,
			s.WithBootstrap,
		),
	)
	err = testEnvironment.Create(
		*CreateEnvironment(
			s.T().Context(),
			s.APIClient,
		))
	if err != nil {
		s.FailNow("Failed to create test environment", err)
	}

	s.TestEnvironment = testEnvironment
}

// TearDownTest is called after each test and cleans up the environment created for that test.
func (s *NewEnvironmentTestSuite) TearDownTest() {
	s.PingOneTestSuite.TearDownTest()

	if s.TestEnvironment != nil {
		err := s.TestEnvironment.Delete(
			*DeleteEnvironment(
				s.T().Context(),
				s.APIClient,
			).IfExists())
		if err != nil {
			s.FailNow("Failed to delete test environment", err)
		}
	}
}

// TearDownSuite is called after all tests in the suite have completed.
func (s *NewEnvironmentTestSuite) TearDownSuite() {
	s.PingOneTestSuite.TearDownSuite()
}
