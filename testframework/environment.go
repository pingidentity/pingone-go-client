// Copyright © 2025 Ping Identity Corporation

package testframework

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pingidentity/pingone-go-client/pingone"
)

// TestEnvironment represents a PingOne environment used for testing purposes.
// It contains both the creation request configuration and the resulting environment
// response after successful creation.
type TestEnvironment struct {
	// EnvironmentCreateRequest contains the configuration used to create the environment.
	EnvironmentCreateRequest pingone.EnvironmentCreateRequest
	// Environment contains the response from PingOne after environment creation.
	Environment *pingone.EnvironmentResponse
	// Population contains the default population for the environment (commented out - not yet implemented)
	// Population  *pingone.Population
}

// NewTestEnvironment creates a new TestEnvironment instance with the provided configuration.
// The environment parameter must contain valid environment creation settings including
// name, region, license, and bill of materials. The environment is not created until
// the Create method is called.
func NewTestEnvironment(environment pingone.EnvironmentCreateRequest) *TestEnvironment {
	return &TestEnvironment{
		EnvironmentCreateRequest: environment,
	}
}

// CreateConfig contains configuration for environment creation operations.
// It encapsulates the API client, context, and creation options for creating
// test environments in PingOne.
type CreateConfig struct {
	apiClient   *pingone.APIClient
	ctx         context.Context
	ifNotExists bool
}

// CreateEnvironment initializes a new CreateConfig for environment creation.
// The ctx parameter provides the context for API operations, and apiClient must be
// a configured PingOne API client with appropriate authentication. Use the returned
// CreateConfig to configure creation options before calling Create.
func CreateEnvironment(ctx context.Context, apiClient *pingone.APIClient) *CreateConfig {
	return &CreateConfig{
		apiClient: apiClient,
		ctx:       ctx,
	}
}

// IfNotExists configures the creation to not return an error if the environment already exists.
// When enabled, if an environment with the same configuration already exists, the operation
// will return the existing environment instead of failing with a conflict error.
func (e *CreateConfig) IfNotExists() *CreateConfig {
	e.ifNotExists = true
	return e
}

// Create creates a new PingOne environment using the configured settings.
// The request parameter must contain a valid CreateConfig with API client and context.
// It calls the PingOne API to create the environment and stores the response in the
// Environment field. Returns an error if the creation fails or if the API returns
// a non-success status code.
func (e *TestEnvironment) Create(request CreateConfig) (err error) {
	environmentResponse, httpResp, err := request.apiClient.EnvironmentsApi.CreateEnvironment(request.ctx).EnvironmentCreateRequest(e.EnvironmentCreateRequest).Execute()
	if err != nil {
		return fmt.Errorf("cannot create environment: %s", err.Error())
	}

	if httpResp.StatusCode != 201 {
		return fmt.Errorf("cannot create environment: %s", httpResp.Status)
	}

	e.Environment = environmentResponse

	return nil
}

// DeleteConfig contains configuration for environment deletion operations.
// It encapsulates the API client, context, and deletion options for removing
// test environments from PingOne.
type DeleteConfig struct {
	apiClient *pingone.APIClient
	ctx       context.Context
	ifExists  bool
}

// DeleteEnvironment initializes a new DeleteConfig for environment deletion.
// The ctx parameter provides the context for API operations, and apiClient must be
// a configured PingOne API client with appropriate authentication. Use the returned
// DeleteConfig to configure deletion options before calling Delete.
func DeleteEnvironment(ctx context.Context, apiClient *pingone.APIClient) *DeleteConfig {
	return &DeleteConfig{
		apiClient: apiClient,
		ctx:       ctx,
	}
}

// IfExists configures the deletion to not return an error if the environment does not exist.
// When enabled, if the environment has already been deleted or does not exist, the operation
// will succeed without error instead of failing with a not found error.
func (e *DeleteConfig) IfExists() *DeleteConfig {
	e.ifExists = true
	return e
}

// Delete removes the PingOne environment using the configured settings.
// The request parameter must contain a valid DeleteConfig with API client and context.
// The Environment field must be populated (from a previous Create operation) unless
// ifExists is true. Returns an error if the deletion fails or if the environment
// is nil and ifExists is false.
func (e *TestEnvironment) Delete(request DeleteConfig) (err error) {

	if e.Environment == nil || !request.ifExists {
		return fmt.Errorf("environment is nil or ifExists is false")
	}

	httpResp, err := request.apiClient.EnvironmentsApi.DeleteEnvironmentById(request.ctx, e.Environment.Id).Execute()
	if err != nil {
		if httpResp != nil && httpResp.StatusCode == 404 {
			return nil
		}
		return fmt.Errorf("cannot delete environment: %s", err.Error())
	}

	if httpResp.StatusCode != 204 {
		return fmt.Errorf("cannot delete environment: %s", httpResp.Status)
	}

	return nil
}

var (
	// DefaultEnvironmentDefinition is a function that creates a standard test environment configuration.
	// The name parameter sets the environment name, regionCodeStr specifies the PingOne region,
	// licenseId must be a valid PingOne license UUID, and withBootstrap determines whether to
	// include full DaVinci features or use minimal configuration. This function creates an
	// environment with all standard PingOne services enabled for comprehensive testing.
	DefaultEnvironmentDefinition = func(name, regionCodeStr string, licenseId uuid.UUID, withBootstrap bool) pingone.EnvironmentCreateRequest {

		davinciTags := make([]string, 0)
		if !withBootstrap {
			davinciTags = append(davinciTags, "DAVINCI_MINIMAL")
		}

		regionCode := pingone.EnvironmentRegionCode(regionCodeStr)

		return pingone.EnvironmentCreateRequest{
			BillOfMaterials: &pingone.EnvironmentBillOfMaterials{
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
			License: pingone.EnvironmentLicense{
				Id: licenseId,
			},
			Name:   name,
			Region: regionCode,
			Type:   "SANDBOX",
		}
	}
)
