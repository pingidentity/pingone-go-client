// Copyright © 2025 Ping Identity Corporation
/*
PingOne User and Configuration Management API

Testing DaVinciConnectorApiService

*/

package pingone_test

import (
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/pingidentity/pingone-go-client/pingone"
	"github.com/pingidentity/pingone-go-client/testframework"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type DaVinciConnectorAPIServiceSharedEnvTestSuite struct {
	testframework.SharedEnvironmentTestSuite

	DefaultDaVinciConnectorInstanceMaxSchemaCreate pingone.DaVinciConnectorInstanceCreateRequest
	DefaultDaVinciConnectorInstanceMinSchemaCreate pingone.DaVinciConnectorInstanceCreateRequest
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) SetupSuite() {
	s.SharedEnvironmentTestSuite.SetupSuite()
}

// Set up the test with a new environment
func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) SetupTest() {
	s.SharedEnvironmentTestSuite.SetupTest()

	s.DefaultDaVinciConnectorInstanceMaxSchemaCreate = pingone.DaVinciConnectorInstanceCreateRequest{
		Name:          pingone.PtrString(testframework.RandomResourceName()),
		Configuration: map[string]interface{}{"key1": "value1", "key2": 123},
		ConnectorId:   uuid.New().String(),
		CompanyId:     pingone.PtrString(testframework.RandomResourceName()),
	}

	s.DefaultDaVinciConnectorInstanceMinSchemaCreate = pingone.DaVinciConnectorInstanceCreateRequest{
		Name:        pingone.PtrString(testframework.RandomResourceName()),
		ConnectorId: uuid.New().String(),
	}
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) TearDownTest() {
	s.SharedEnvironmentTestSuite.TearDownTest()
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) TearDownSuite() {
	s.SharedEnvironmentTestSuite.TearDownSuite()
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) TestDaVinciConnectorInstanceNeverFound() {
	davinciConnectorInstanceID := uuid.New()

	resp, httpRes, err := s.ApiClient.DaVinciConnectorApi.GetConnectorInstanceById(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId(), davinciConnectorInstanceID).Execute()
	testframework.CheckNotFound(s.T(), resp, httpRes, err)
	testframework.CheckPingOneAPIErrorResponse(s.T(), err, pingone.NotFoundError{}, regexp.MustCompile("The requested resource was not found"))
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) TestDaVinciConnectorFullLifecycle() {
	// Min Schema
	connectorInstanceID := s.test_pingone_DaVinciConnectorApiService_CreateInstance(s.T(), s.DefaultDaVinciConnectorInstanceMinSchemaCreate)

	s.test_pingone_DaVinciConnectorApiService_DeleteInstance(s.T(), connectorInstanceID)

	// Max Schema
	connectorInstanceID = s.test_pingone_DaVinciConnectorApiService_CreateInstance(s.T(), s.DefaultDaVinciConnectorInstanceMaxSchemaCreate)

	// Finally Delete
	s.test_pingone_DaVinciConnectorApiService_DeleteInstance(s.T(), connectorInstanceID)
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) test_pingone_DaVinciConnectorApiService_CreateInstance(t *testing.T, payload pingone.DaVinciConnectorInstanceCreateRequest) (connectorInstanceID string) {
	resp, httpRes, err := s.ApiClient.DaVinciConnectorApi.CreateConnectorInstance(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId()).DaVinciConnectorInstanceCreateRequest(payload).Execute()
	testframework.CheckCreated(t, resp, &pingone.DaVinciConnectorInstance{}, httpRes, err)

	require.NotNil(t, resp.Id)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)
	assert.Equal(t, payload.Name, resp.Name)

	s.test_pingone_DaVinciConnectorApiService_GetInstance(t, resp.Id, payload)

	return resp.Id
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) test_pingone_DaVinciConnectorApiService_GetInstance(t *testing.T, connectorInstanceID string, payload any) {
	resp, httpRes, err := s.ApiClient.DaVinciConnectorApi.GetConnectorInstanceById(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId(), connectorInstanceID).Execute()
	testframework.CheckFound(t, resp, &pingone.DaVinciConnectorInstance{}, httpRes, err)

	require.NotNil(t, resp.Id)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)

	// switch obj := payload.(type) {
	// case pingone.DaVinciConnectorInstanceCreateRequest:
	// 	assert.Equal(t, *obj.Name, *resp.Name)
	// 	assert.Equal(t, obj.ConnectorId, resp.ConnectorId)
	// 	if obj.CompanyId != nil {
	// 		assert.Equal(t, *obj.CompanyId, *resp.CompanyId)
	// 	}
	// case pingone.DaVinciConnectorInstance:
	// 	assert.Equal(t, *obj.Name, *resp.Name)
	// 	assert.Equal(t, obj.ConnectorId, resp.ConnectorId)
	// 	if obj.CompanyId != nil {
	// 		assert.Equal(t, *obj.CompanyId, *resp.CompanyId)
	// 	}
	// default:
	// 	assert.Fail(t, "Unknown payload type")
	// }
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) test_pingone_DaVinciConnectorApiService_DeleteInstance(t *testing.T, connectorInstanceID string) {
	httpRes, err := s.ApiClient.DaVinciConnectorApi.DeleteConnectorInstanceById(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId(), connectorInstanceID).Execute()
	testframework.CheckDeleted(t, httpRes, err)

	resp, httpRes, err := s.ApiClient.DaVinciConnectorApi.GetConnectorInstanceById(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId(), connectorInstanceID).Execute()
	testframework.CheckNotFound(t, resp, httpRes, err)
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) TestDaVinciConnectorNeverFound() {
	davinciConnectorID := "neverFound"

	resp, httpRes, err := s.ApiClient.DaVinciConnectorApi.GetConnectorById(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId(), davinciConnectorID).Execute()
	testframework.CheckNotFound(s.T(), resp, httpRes, err)
	testframework.CheckPingOneAPIErrorResponse(s.T(), err, pingone.NotFoundError{}, regexp.MustCompile("The requested resource was not found"))
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) TestDaVinciConnectorGetById() {
	// This test assumes there's at least one connector available
	iterator := s.ApiClient.DaVinciConnectorApi.GetConnectors(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId()).Execute()

	if !iterator.HasNext() {
		s.T().Skip("No connectors found to test GetConnectorById, skipping test")
		return
	}

	connector, err := iterator.Next()
	require.NoError(s.T(), err)

	resp, httpRes, err := s.ApiClient.DaVinciConnectorApi.GetConnectorById(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId(), connector.Id).Execute()
	testframework.CheckFound(s.T(), resp, &pingone.DaVinciConnectorMinimalResponse{}, httpRes, err)

	assert.Equal(s.T(), connector.Id, resp.Id)
	assert.NotEmpty(s.T(), resp.Name)
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) TestDaVinciConnectorInstancesList() {
	// Create a few connector instances to ensure we have data to list
	connectorInstanceID1 := s.test_pingone_DaVinciConnectorApiService_CreateInstance(s.T(), s.DefaultDaVinciConnectorInstanceMinSchemaCreate)
	connectorInstanceID2 := s.test_pingone_DaVinciConnectorApiService_CreateInstance(s.T(), s.DefaultDaVinciConnectorInstanceMaxSchemaCreate)

	defer func() {
		s.test_pingone_DaVinciConnectorApiService_DeleteInstance(s.T(), connectorInstanceID1)
		s.test_pingone_DaVinciConnectorApiService_DeleteInstance(s.T(), connectorInstanceID2)
	}()

	// Test list all connector instances
	resp, httpRes, err := s.ApiClient.DaVinciConnectorApi.GetConnectorInstances(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId()).Execute()
	testframework.CheckFound(s.T(), resp, &pingone.DaVinciConnectorInstanceCollection{}, httpRes, err)

	require.NotNil(s.T(), resp.Embedded)
	instances := resp.Embedded.GetConnections()

	foundIDs := make(map[uuid.UUID]bool)
	for _, instance := range instances {
		foundIDs[instance.Id] = true
	}

	assert.True(s.T(), foundIDs[connectorInstanceID1], "Should find the first created connector instance")
	assert.True(s.T(), foundIDs[connectorInstanceID2], "Should find the second created connector instance")
}

func (s *DaVinciConnectorAPIServiceSharedEnvTestSuite) TestGetConnectors() {
	resp, httpRes, err := s.ApiClient.DaVinciConnectorApi.GetConnectors(s.T().Context(), s.SharedEnvironmentTestSuite.TestEnvironment.Environment.GetId()).Execute()
	testframework.CheckFound(s.T(), resp, &pingone.DaVinciConnectorCollectionMinimalResponse{}, httpRes, err)

	require.NotNil(s.T(), resp.Embedded)
	connectors := resp.Embedded.GetConnectors()

	// Just verify we get a response with connectors
	assert.NotEmpty(s.T(), connectors, "Should find some connectors")

	for _, connector := range connectors {
		assert.NotEmpty(s.T(), connector.Id)
		assert.NotEmpty(s.T(), connector.Name)
	}
}

func Test_pingone_DaVinciConnectorAPIService(t *testing.T) {
	suite.Run(t, &DaVinciConnectorAPIServiceSharedEnvTestSuite{})
}
