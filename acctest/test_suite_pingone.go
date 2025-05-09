package acctest

import (
	"github.com/pingidentity/pingone-go-client/pingone"
	"github.com/stretchr/testify/suite"
)

type PingOneTestSuite struct {
	suite.Suite
	ApiClient *pingone.APIClient
}

func (s *PingOneTestSuite) SetupSuite() {
	if s.ApiClient == nil {
		s.ApiClient = TestClient(s.T().Context())
	}
}

func (s *PingOneTestSuite) SetupTest() {}

func (s *PingOneTestSuite) TearDownTest() {}

func (s *PingOneTestSuite) TearDownSuite() {
	s.ApiClient = nil
}
