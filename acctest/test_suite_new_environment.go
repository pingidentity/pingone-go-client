package acctest

import (
	"fmt"
	"os"

	"github.com/google/uuid"
)

type NewEnvironmentTestSuite struct {
	PingOneTestSuite
	EnvironmentNamePrefix string
	EnvironmentNameSuffix string
	TestEnvironment       *TestEnvironment
	WithBootstrap         bool
}

func (s *NewEnvironmentTestSuite) SetupSuite() {
	s.PingOneTestSuite.SetupSuite()
}

// Set up the test with a new environment
func (s *NewEnvironmentTestSuite) SetupTest() {
	s.PingOneTestSuite.SetupTest()

	if s.EnvironmentNamePrefix == "" {
		s.EnvironmentNamePrefix = "go-sdk-testacc-"
	}

	regionCode := os.Getenv("PINGONE_ENVIRONMENT_REGION")

	licenseId, err := uuid.Parse(os.Getenv("PINGONE_LICENSE_ID"))
	if err != nil {
		s.FailNow("Failed to parse license ID as a valid UUID (PingOne resource ID)", err)
	}

	testEnvironment := NewTestEnvironment(
		DefaultEnvironmentDefinition(
			fmt.Sprintf("%s%s%s", s.EnvironmentNamePrefix, randomString(10), s.EnvironmentNameSuffix),
			regionCode,
			licenseId,
			s.WithBootstrap,
		),
	)
	err = testEnvironment.Create(
		*CreateEnvironment(
			s.T().Context(),
			s.PingOneTestSuite.ApiClient,
		))
	if err != nil {
		s.FailNow("Failed to create test environment", err)
	}

	s.TestEnvironment = testEnvironment
}

func (s *NewEnvironmentTestSuite) TearDownTest() {
	s.PingOneTestSuite.TearDownTest()

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

func (s *NewEnvironmentTestSuite) TearDownSuite() {
	s.PingOneTestSuite.TearDownSuite()
}
