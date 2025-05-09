package acctest

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pingidentity/pingone-go-client/pingone"
)

type TestEnvironment struct {
	EnvironmentCreateRequest pingone.EnvironmentCreateRequest
	Environment              *pingone.Environment
	// Population  *pingone.Population
}

func NewTestEnvironment(environment pingone.EnvironmentCreateRequest) *TestEnvironment {
	return &TestEnvironment{
		EnvironmentCreateRequest: environment,
	}
}

type CreateConfig struct {
	apiClient   *pingone.APIClient
	ctx         context.Context
	ifNotExists bool
}

func CreateEnvironment(ctx context.Context, apiClient *pingone.APIClient) *CreateConfig {
	return &CreateConfig{
		apiClient: apiClient,
		ctx:       ctx,
	}
}

// IfNotExists will not return an error if the environment already exists, but instead will return the existing environment
func (e *CreateConfig) IfNotExists() *CreateConfig {
	e.ifNotExists = true
	return e
}

func (e *TestEnvironment) Create(request CreateConfig) (err error) {
	environmentResponse, httpResp, err := request.apiClient.EnvironmentApi.CreateEnvironment(request.ctx).EnvironmentCreateRequest(e.EnvironmentCreateRequest).Execute()
	if err != nil {
		return fmt.Errorf("Cannot create environment: %s", err.Error())
	}

	if httpResp.StatusCode != 201 {
		return fmt.Errorf("Cannot create environment: %s", httpResp.Status)
	}

	e.Environment = environmentResponse

	return nil
}

type DeleteConfig struct {
	apiClient *pingone.APIClient
	ctx       context.Context
	ifExists  bool
}

func DeleteEnvironment(ctx context.Context, apiClient *pingone.APIClient) *DeleteConfig {
	return &DeleteConfig{
		apiClient: apiClient,
		ctx:       ctx,
	}
}

// IfExists will not return an error if the environment does not exist
func (e *DeleteConfig) IfExists() *DeleteConfig {
	e.ifExists = true
	return e
}

func (e *TestEnvironment) Delete(request DeleteConfig) (err error) {

	if e.Environment == nil || !request.ifExists {
		return fmt.Errorf("Environment is nil or ifExists is false")
	}

	httpResp, err := request.apiClient.EnvironmentApi.DeleteEnvironmentById(request.ctx, *e.Environment.Id).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("Cannot delete environment: %s", err.Error())
	}

	if httpResp.StatusCode != 204 {
		return fmt.Errorf("Cannot delete environment: %s", httpResp.Status)
	}

	return nil
}

var (
	DefaultEnvironmentDefinition = func(name, regionCode string, licenseId uuid.UUID, withBootstrap bool) pingone.EnvironmentCreateRequest {

		davinciTags := make([]pingone.EnvironmentBillOfMaterialsProductTags, 0)
		if !withBootstrap {
			davinciTags = append(davinciTags, "DAVINCI_MINIMAL")
		}

		return pingone.EnvironmentCreateRequest{
			BillOfMaterials: &pingone.EnvironmentBillOfMaterialsCreateRequest{
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
						Tags: davinciTags,
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
			},
			Description: pingone.PtrString("Test environment created by the PingOne Go SDK"),
			License: pingone.ResourceRelationshipPingOne{
				Id: licenseId,
			},
			Name:   name,
			Region: pingone.StringAsEnvironmentRegion(&regionCode),
			Type:   "SANDBOX",
		}
	}
)
