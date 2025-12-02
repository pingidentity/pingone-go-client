// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

// SharedEnvironmentTestSuite extends PingOneTestSuite to provide a shared test environment.
// This suite creates a single PingOne environment that is shared across all test methods
// in the suite, improving test performance by reducing environment creation overhead.
type SharedEnvironmentTestSuite struct {
	PingOneTestSuite
	// EnvironmentNamePrefix is prepended to generated environment names.
	EnvironmentNamePrefix string
	// EnvironmentNameSuffix is appended to generated environment names.
	EnvironmentNameSuffix string
	// TestEnvironment contains the shared test environment for all tests in the suite.
	TestEnvironment *TestEnvironment
	// WithBootstrap determines whether the environment includes full DaVinci features.
	WithBootstrap bool
}

// SetupSuite creates a shared PingOne environment for all test methods in the suite.
// This method requires the PINGONE_ENVIRONMENT_REGION and PINGONE_LICENSE_ID environment
// variables to be set. It generates a unique environment name and creates an environment
// with all PingOne services. The environment is shared across all tests to improve performance.
func (s *SharedEnvironmentTestSuite) SetupSuite() {
	s.PingOneTestSuite.SetupSuite()

	if s.EnvironmentNamePrefix == "" {
		s.EnvironmentNamePrefix = "go-sdk-testacc-"
	}

	if s.EnvironmentNameSuffix == "" {
		s.EnvironmentNameSuffix = "-shared"
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
		).IfNotExists())
	if err != nil {
		s.FailNow("Failed to create test environment", err)
	}

	s.TestEnvironment = testEnvironment
}

// SetupTest prepares the test with the shared environment.
func (s *SharedEnvironmentTestSuite) SetupTest() {
	s.PingOneTestSuite.SetupTest()

	// TODO: Check environment is still present
}

// TearDownTest is called after each test to verify the shared environment.
func (s *SharedEnvironmentTestSuite) TearDownTest() {
	s.PingOneTestSuite.TearDownTest()

	// TODO: Check environment is still present (does the test force removal of the environment by mistake?)
}

// TearDownSuite cleans up the shared environment after all tests have completed.
func (s *SharedEnvironmentTestSuite) TearDownSuite() {

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

	s.PingOneTestSuite.TearDownSuite()
}
