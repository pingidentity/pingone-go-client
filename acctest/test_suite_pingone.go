// Copyright © 2025 Ping Identity Corporation

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
		var err error
		s.ApiClient, err = TestClient(nil)
		if err != nil {
			s.T().Fatalf("Failed to create API client: %v", err)
		}
	}
}

func (s *PingOneTestSuite) SetupTest() {}

func (s *PingOneTestSuite) TearDownTest() {}

func (s *PingOneTestSuite) TearDownSuite() {
	s.ApiClient = nil
}
