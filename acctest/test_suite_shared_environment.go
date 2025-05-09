package acctest

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

type SharedEnvironmentTestSuite struct {
	PingOneTestSuite
	EnvironmentNamePrefix string
	EnvironmentNameSuffix string
	TestEnvironment       *TestEnvironment
	WithBootstrap         bool
}

// Set up the entire suite with a shared environment
func (s *SharedEnvironmentTestSuite) SetupSuite() {
	s.PingOneTestSuite.SetupSuite()

	if s.EnvironmentNamePrefix == "" {
		s.EnvironmentNamePrefix = "go-sdk-testacc-"
	}

	if s.EnvironmentNameSuffix == "" {
		s.EnvironmentNameSuffix = "-shared"
	}

	regionCode := os.Getenv("PINGONE_ENVIRONMENT_REGION")

	licenseId, err := uuid.Parse(os.Getenv("PINGONE_LICENSE_ID"))
	if err != nil {
		s.FailNow("Failed to parse license ID as a valid UUID (PingOne resource ID)", err)
	}

	testEnvironment := NewTestEnvironment(
		DefaultEnvironmentDefinition(
			fmt.Sprintf("%s%s%s", s.EnvironmentNamePrefix, randomString(10), s.EnvironmentNamePrefix),
			regionCode,
			licenseId,
			s.WithBootstrap,
		),
	)
	err = testEnvironment.Create(
		*CreateEnvironment(
			s.T().Context(),
			s.PingOneTestSuite.ApiClient,
		).IfNotExists())
	if err != nil {
		s.FailNow("Failed to create test environment", err)
	}

	s.TestEnvironment = testEnvironment
}

// Set up the test with a new environment
func (s *SharedEnvironmentTestSuite) SetupTest() {
	s.PingOneTestSuite.SetupTest()

	// TODO: Check environment is still present
}

func (s *SharedEnvironmentTestSuite) TearDownTest() {
	s.PingOneTestSuite.TearDownTest()

	// TODO: Check environment is still present (does the test force removal of the environment by mistake?)
}

func (s *SharedEnvironmentTestSuite) TearDownSuite() {
	s.PingOneTestSuite.TearDownSuite()

	if s.TestEnvironment != nil {
		err := s.TestEnvironment.Delete(
			*DeleteEnvironment(
				s.T().Context(),
				s.PingOneTestSuite.ApiClient,
			).IfExists())
		if err != nil {
			s.FailNow("Failed to delete test environment", err)
		}
	}
}
