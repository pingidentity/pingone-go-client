// Copyright © 2025 Ping Identity Corporation
/*
PingOne User and Configuration Management API

Testing EnvironmentApiService

*/

package pingone_test

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/google/uuid"
	"github.com/pingidentity/pingone-go-client/pingone"
	"github.com/pingidentity/pingone-go-client/testframework"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type EnvironmentApiServiceNoAuthzTestSuite struct {
	suite.Suite
	BadApiClient *pingone.APIClient
}

func (s *EnvironmentApiServiceNoAuthzTestSuite) SetupSuite() {
	svcConfig := pingone.NewServiceConfiguration().WithAccessToken("bad-token")
	configuration := pingone.NewConfiguration(svcConfig)
	configuration.AppendUserAgent("testing")

	var err error
	s.BadApiClient, err = pingone.NewAPIClient(configuration)
	if err != nil {
		s.FailNow("Failed to create API client", err)
	}
}

type EnvironmentApiServiceTestSuite struct {
	testframework.PingOneTestSuite

	DefaultRegionCode string
	DefaultLicenseId  uuid.UUID

	EnvironmentNamePrefix string
	EnvironmentNameSuffix string

	DefaultEnvironmentMaxSchemaCreate  pingone.EnvironmentCreateRequest
	DefaultEnvironmentMaxSchemaReplace pingone.EnvironmentReplaceRequest
	DefaultEnvironmentMinSchemaCreate  pingone.EnvironmentCreateRequest
	DefaultEnvironmentMinSchemaReplace pingone.EnvironmentReplaceRequest

	DefaultEnvironmentBillOfMaterialsMaxSchemaCreate  pingone.EnvironmentBillOfMaterials
	DefaultEnvironmentBillOfMaterialsMaxSchemaReplace pingone.EnvironmentBillOfMaterialsReplaceRequest
	DefaultEnvironmentBillOfMaterialsMinSchemaCreate  pingone.EnvironmentBillOfMaterials
	DefaultEnvironmentBillOfMaterialsMinSchemaReplace pingone.EnvironmentBillOfMaterialsReplaceRequest
}

func (s *EnvironmentApiServiceTestSuite) SetupSuite() {
	s.PingOneTestSuite.SetupSuite()

	s.DefaultRegionCode = os.Getenv("PINGONE_ENVIRONMENT_REGION")

	s.EnvironmentNamePrefix = "go-sdk-testacc-"
	s.EnvironmentNameSuffix = "-api-service"

	var err error
	s.DefaultLicenseId, err = uuid.Parse(os.Getenv("PINGONE_LICENSE_ID"))
	if err != nil {
		s.FailNow("Failed to parse license ID as a valid UUID (PingOne resource ID)", err)
	}
}

func (s *EnvironmentApiServiceTestSuite) SetupTest() {
	s.PingOneTestSuite.SetupTest()

	regionCode := pingone.EnvironmentRegionCode(s.DefaultRegionCode)

	s.DefaultEnvironmentBillOfMaterialsMaxSchemaCreate = pingone.EnvironmentBillOfMaterials{
		Products: []pingone.EnvironmentBillOfMaterialsProduct{
			{
				Type: "PING_ONE_AUTHORIZE",
			},
			{
				Type: "PING_ONE_BASE",
			},
			{
				Type: "PING_ONE_CREDENTIALS",
			},
			{
				Type: "PING_ONE_DAVINCI",
				Tags: []string{
					"DAVINCI_MINIMAL",
				},
			},
			{
				Type: "PING_ONE_MFA",
			},
			{
				Type: "PING_ONE_RISK",
			},
			{
				Type: "PING_ONE_VERIFY",
			},
		},
	}

	s.DefaultEnvironmentBillOfMaterialsMaxSchemaReplace = pingone.EnvironmentBillOfMaterialsReplaceRequest{
		Products: s.DefaultEnvironmentBillOfMaterialsMaxSchemaCreate.Products,
	}

	s.DefaultEnvironmentBillOfMaterialsMinSchemaCreate = pingone.EnvironmentBillOfMaterials{
		Products: []pingone.EnvironmentBillOfMaterialsProduct{
			{
				Type: "PING_ONE_BASE",
			},
		},
	}

	s.DefaultEnvironmentBillOfMaterialsMinSchemaReplace = pingone.EnvironmentBillOfMaterialsReplaceRequest{
		Products: s.DefaultEnvironmentBillOfMaterialsMaxSchemaReplace.Products,
	}

	s.DefaultEnvironmentMaxSchemaCreate = pingone.EnvironmentCreateRequest{
		BillOfMaterials: &s.DefaultEnvironmentBillOfMaterialsMaxSchemaCreate,
		Description:     pingone.PtrString("Test environment created by the PingOne Go SDK"),
		Icon:            pingone.PtrString("https://bxretail.org/icon.png"),
		License: &pingone.EnvironmentLicense{
			Id: s.DefaultLicenseId,
		},
		Name: fmt.Sprintf("%s%s%s", s.EnvironmentNamePrefix, testframework.RandomResourceName(), s.EnvironmentNameSuffix),
		Region: pingone.EnvironmentCreateRequestRegion{
			EnvironmentRegionCode: &regionCode,
		},
		Type: "SANDBOX",
	}

	s.DefaultEnvironmentMinSchemaCreate = pingone.EnvironmentCreateRequest{
		License: &pingone.EnvironmentLicense{
			Id: s.DefaultLicenseId,
		},
		Name: fmt.Sprintf("%s%s%s", s.EnvironmentNamePrefix, testframework.RandomResourceName(), s.EnvironmentNameSuffix),
		Region: pingone.EnvironmentCreateRequestRegion{
			EnvironmentRegionCode: &regionCode,
		},
		Type: "SANDBOX",
	}

	environmentStatus := pingone.ENVIRONMENTSTATUS_DELETE_PENDING

	s.DefaultEnvironmentMaxSchemaReplace = pingone.EnvironmentReplaceRequest{
		BillOfMaterials: &s.DefaultEnvironmentBillOfMaterialsMinSchemaReplace,
		Description:     pingone.PtrString("Test environment updated by the PingOne Go SDK"),
		Icon:            pingone.PtrString("https://bxretail.org/icon1.png"),
		License: &pingone.EnvironmentLicense{
			Id: s.DefaultLicenseId,
		},
		Name: fmt.Sprintf("%s%s%s", s.EnvironmentNamePrefix, testframework.RandomResourceName(), s.EnvironmentNameSuffix),
		Type: "SANDBOX",
		Region: pingone.EnvironmentReplaceRequestRegion{
			EnvironmentRegionCode: &regionCode,
		},
		Status: &environmentStatus,
	}

	s.DefaultEnvironmentMinSchemaReplace = pingone.EnvironmentReplaceRequest{
		License: &pingone.EnvironmentLicense{
			Id: s.DefaultLicenseId,
		},
		Name: fmt.Sprintf("%s%s%s", s.EnvironmentNamePrefix, testframework.RandomResourceName(), s.EnvironmentNameSuffix),
		Region: pingone.EnvironmentReplaceRequestRegion{
			EnvironmentRegionCode: &regionCode,
		},
		Type: "SANDBOX",
	}
}

func (s *EnvironmentApiServiceTestSuite) TearDownTest() {
	s.PingOneTestSuite.TearDownTest()
}

func (s *EnvironmentApiServiceTestSuite) TearDownSuite() {
	s.PingOneTestSuite.TearDownSuite()
}

func (s *EnvironmentApiServiceNoAuthzTestSuite) TestEnvironmentNoAuthorization() {

	environmentID := uuid.New()

	resp, httpRes, err := s.BadApiClient.EnvironmentApi.GetEnvironmentById(s.T().Context(), environmentID).Execute()

	require.NotNil(s.T(), err)
	require.Nil(s.T(), resp)
	require.NotNil(s.T(), httpRes)
	assert.Equal(s.T(), 401, httpRes.StatusCode)
	assert.EqualError(s.T(), err, "401 Unauthorized")

	require.IsType(s.T(), &pingone.UnauthorizedError{}, err)
}

func (s *EnvironmentApiServiceTestSuite) TestEnvironmentNeverFound() {

	environmentID := uuid.New()

	resp, httpRes, err := s.ApiClient.EnvironmentApi.GetEnvironmentById(s.T().Context(), environmentID).Execute()
	testframework.CheckNotFound(s.T(), resp, httpRes, err)
	testframework.CheckPingOneAPIErrorResponse(s.T(), err, pingone.NotFoundError{}, regexp.MustCompile("Unable to find environment"))
}

func (s *EnvironmentApiServiceTestSuite) TestEnvironmentFullLifecycle() {

	// Min Schema
	environmentID := s.test_pingone_EnvironmentApiService_Create(s.T(), s.DefaultEnvironmentMinSchemaCreate)

	s.test_pingone_EnvironmentApiService_Delete(s.T(), environmentID)

	// Max Schema
	environmentID = s.test_pingone_EnvironmentApiService_Create(s.T(), s.DefaultEnvironmentMaxSchemaCreate)

	// Replace Max with Min
	s.test_pingone_EnvironmentApiService_Replace(s.T(), environmentID, s.DefaultEnvironmentMinSchemaReplace)

	// Replace Min with Max
	s.test_pingone_EnvironmentApiService_Replace(s.T(), environmentID, s.DefaultEnvironmentMaxSchemaReplace)

	// Finally Delete
	s.test_pingone_EnvironmentApiService_Delete(s.T(), environmentID)
}

func (s *EnvironmentApiServiceTestSuite) test_pingone_EnvironmentApiService_Create(t *testing.T, payload pingone.EnvironmentCreateRequest) (environmentID uuid.UUID) {
	resp, httpRes, err := s.ApiClient.EnvironmentApi.CreateEnvironment(s.T().Context()).EnvironmentCreateRequest(payload).Execute()
	testframework.CheckCreated(t, resp, &pingone.Environment{}, httpRes, err)

	require.NotNil(t, resp.Id)
	assert.NotNil(t, resp.Links)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)
	// TODO: Check data

	s.test_pingone_EnvironmentApiService_Get(t, resp.Id, payload)

	return resp.GetId()
}

func (s *EnvironmentApiServiceTestSuite) test_pingone_EnvironmentApiService_Get(t *testing.T, environmentID uuid.UUID, payload any) {
	resp, httpRes, err := s.ApiClient.EnvironmentApi.GetEnvironmentById(s.T().Context(), environmentID).Execute()
	testframework.CheckFound(t, resp, &pingone.Environment{}, httpRes, err)

	require.NotNil(t, resp.Id)
	assert.NotNil(t, resp.Links)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)

	// TODO: Check data
	switch obj := payload.(type) {
	case pingone.EnvironmentCreateRequest:
		assert.Equal(t, resp.Description, obj.Description)
	case pingone.EnvironmentReplaceRequest:
		assert.Equal(t, resp.Description, obj.Description)
	case pingone.Environment:
		assert.Equal(t, resp.Description, obj.Description)
	default:
		assert.Fail(t, "Unknown payload type")
	}
}

func (s *EnvironmentApiServiceTestSuite) test_pingone_EnvironmentApiService_Replace(t *testing.T, environmentID uuid.UUID, payload pingone.EnvironmentReplaceRequest) {
	resp, httpRes, err := s.ApiClient.EnvironmentApi.ReplaceEnvironmentById(s.T().Context(), environmentID).EnvironmentReplaceRequest(payload).Execute()
	testframework.CheckReplaced(t, resp, &pingone.Environment{}, httpRes, err)

	require.Equal(t, resp.Id, environmentID)
	assert.NotNil(t, resp.Links)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)
	// TODO: Check data

	s.test_pingone_EnvironmentApiService_Get(t, resp.Id, payload)
}

func (s *EnvironmentApiServiceTestSuite) test_pingone_EnvironmentApiService_Delete(t *testing.T, environmentID uuid.UUID) {
	httpRes, err := s.ApiClient.EnvironmentApi.DeleteEnvironmentById(s.T().Context(), environmentID).Execute()
	testframework.CheckDeleted(t, httpRes, err)

	resp, httpRes, err := s.ApiClient.EnvironmentApi.GetEnvironmentById(s.T().Context(), environmentID).Execute()
	testframework.CheckNotFound(t, resp, httpRes, err)
}

// func (s *EnvironmentApiServiceTestSuite) TestEnvironmentChangeEnvironmentType() {
// }

// func (s *EnvironmentApiServiceTestSuite) TestEnvironmentChangeStatus() {
// }

type EnvironmentApiServiceModifyTestSuite struct {
	testframework.NewEnvironmentTestSuite

	DefaultEnvironmentBillOfMaterialsMaxSchemaReplace pingone.EnvironmentBillOfMaterialsReplaceRequest
	DefaultEnvironmentBillOfMaterialsMinSchemaReplace pingone.EnvironmentBillOfMaterialsReplaceRequest
}

func (s *EnvironmentApiServiceModifyTestSuite) SetupSuite() {
	s.NewEnvironmentTestSuite.SetupSuite()
}

// Set up the test with a new environment
func (s *EnvironmentApiServiceModifyTestSuite) SetupTest() {
	s.NewEnvironmentTestSuite.SetupTest()

	s.DefaultEnvironmentBillOfMaterialsMaxSchemaReplace = pingone.EnvironmentBillOfMaterialsReplaceRequest{
		Products: []pingone.EnvironmentBillOfMaterialsProduct{
			{
				Type: "PING_ONE_AUTHORIZE",
			},
			{
				Type: "PING_ONE_BASE",
			},
			{
				Type: "PING_ONE_CREDENTIALS",
			},
			{
				Type: "PING_ONE_DAVINCI",
				Tags: []string{
					"DAVINCI_MINIMAL",
				},
			},
			{
				Type: "PING_ONE_MFA",
			},
			{
				Type: "PING_ONE_RISK",
			},
			{
				Type: "PING_ONE_VERIFY",
			},
		},
	}

	s.DefaultEnvironmentBillOfMaterialsMinSchemaReplace = pingone.EnvironmentBillOfMaterialsReplaceRequest{
		Products: []pingone.EnvironmentBillOfMaterialsProduct{
			{
				Type: "PING_ONE_BASE",
			},
		},
	}
}

func (s *EnvironmentApiServiceModifyTestSuite) TearDownTest() {
	s.NewEnvironmentTestSuite.TearDownTest()
}

func (s *EnvironmentApiServiceModifyTestSuite) TearDownSuite() {
	s.NewEnvironmentTestSuite.TearDownSuite()
}

func (s *EnvironmentApiServiceModifyTestSuite) TestEnvironmentBillOfMaterialsFullLifecycle() {
	// Replace Max with Min
	s.test_pingone_EnvironmentBillOfMaterialsApiService_Replace(s.T(), s.TestEnvironment.Environment.GetId(), s.DefaultEnvironmentBillOfMaterialsMinSchemaReplace)

	// Replace Min with Max
	s.test_pingone_EnvironmentBillOfMaterialsApiService_Replace(s.T(), s.TestEnvironment.Environment.GetId(), s.DefaultEnvironmentBillOfMaterialsMaxSchemaReplace)

	// Replace Max with Min again
	s.test_pingone_EnvironmentBillOfMaterialsApiService_Replace(s.T(), s.TestEnvironment.Environment.GetId(), s.DefaultEnvironmentBillOfMaterialsMaxSchemaReplace)

}

func (s *EnvironmentApiServiceModifyTestSuite) test_pingone_EnvironmentBillOfMaterialsApiService_Get(t *testing.T, environmentID uuid.UUID, payload any) {
	resp, httpRes, err := s.ApiClient.EnvironmentApi.GetBillOfMaterialsByEnvironmentId(s.T().Context(), environmentID).Execute()
	testframework.CheckFound(t, resp, &pingone.EnvironmentBillOfMaterials{}, httpRes, err)

	assert.NotNil(t, resp.Links)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)

	// TODO: Check data
	switch obj := payload.(type) {
	case *pingone.EnvironmentBillOfMaterials:
		assert.Equal(t, resp.SolutionType, obj.SolutionType)
		assert.Equal(t, resp.Products, obj.Products)
	case pingone.EnvironmentBillOfMaterialsReplaceRequest:
		assert.Equal(t, resp.Products, obj.Products)
	case pingone.EnvironmentBillOfMaterials:
		assert.Equal(t, resp.SolutionType, obj.SolutionType)
		assert.Equal(t, resp.Products, obj.Products)
	default:
		assert.Fail(t, "Unknown payload type")
	}
}

func (s *EnvironmentApiServiceModifyTestSuite) test_pingone_EnvironmentBillOfMaterialsApiService_Replace(t *testing.T, environmentID uuid.UUID, payload pingone.EnvironmentBillOfMaterialsReplaceRequest) {
	resp, httpRes, err := s.ApiClient.EnvironmentApi.ReplaceBillOfMaterialsByEnvironmentId(s.T().Context(), environmentID).EnvironmentBillOfMaterialsReplaceRequest(payload).Execute()
	testframework.CheckReplaced(t, resp, &pingone.EnvironmentBillOfMaterials{}, httpRes, err)

	assert.NotNil(t, resp.Links)
	assert.NotNil(t, resp.CreatedAt)
	assert.NotNil(t, resp.UpdatedAt)
	// TODO: Check data

	s.test_pingone_EnvironmentBillOfMaterialsApiService_Get(t, environmentID, payload)
}

func Test_pingone_EnvironmentApiService(t *testing.T) {
	// suite.Run(t, &EnvironmentApiServiceNoAuthzTestSuite{})
	suite.Run(t, &EnvironmentApiServiceTestSuite{})
	// suite.Run(t, &EnvironmentApiServiceModifyTestSuite{})
}
