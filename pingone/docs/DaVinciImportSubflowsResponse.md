# DaVinciImportSubflowsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciImportFlowResponseLinks**](DaVinciImportFlowResponseLinks.md) |  | 
**Environment** | [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | 
**Flow** | [**DaVinciImportFlowResponseFlow**](DaVinciImportFlowResponseFlow.md) |  | 
**PublishedVersion** | **float32** |  | 
**Version** | **float32** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**ClonedFrom** | Pointer to **float32** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Connectors** | Pointer to [**[]ResourceRelationshipDaVinciReadOnly**](ResourceRelationshipDaVinciReadOnly.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeployedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**GraphData** | Pointer to [**DaVinciFlowGraphDataResponse**](DaVinciFlowGraphDataResponse.md) |  | [optional] 
**InputSchema** | Pointer to [**[]DaVinciFlowInputSchemaResponseItem**](DaVinciFlowInputSchemaResponseItem.md) |  | [optional] 
**OutputSchema** | Pointer to [**DaVinciImportFlowResponseOutputSchema**](DaVinciImportFlowResponseOutputSchema.md) |  | [optional] 
**Settings** | Pointer to [**DaVinciFlowSettingsResponse**](DaVinciFlowSettingsResponse.md) |  | [optional] 
**Trigger** | Pointer to [**DaVinciImportFlowResponseTrigger**](DaVinciImportFlowResponseTrigger.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Updates** | Pointer to **[]string** |  | [optional] 
**Variables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Embedded** | Pointer to [**[]DaVinciImportFlowResponse**](DaVinciImportFlowResponse.md) |  | [optional] 

## Methods

### NewDaVinciImportSubflowsResponse

`func NewDaVinciImportSubflowsResponse(links DaVinciImportFlowResponseLinks, environment ResourceRelationshipReadOnly, flow DaVinciImportFlowResponseFlow, publishedVersion float32, version float32, ) *DaVinciImportSubflowsResponse`

NewDaVinciImportSubflowsResponse instantiates a new DaVinciImportSubflowsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciImportSubflowsResponseWithDefaults

`func NewDaVinciImportSubflowsResponseWithDefaults() *DaVinciImportSubflowsResponse`

NewDaVinciImportSubflowsResponseWithDefaults instantiates a new DaVinciImportSubflowsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciImportSubflowsResponse) GetLinks() DaVinciImportFlowResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciImportSubflowsResponse) GetLinksOk() (*DaVinciImportFlowResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciImportSubflowsResponse) SetLinks(v DaVinciImportFlowResponseLinks)`

SetLinks sets Links field to given value.


### GetEnvironment

`func (o *DaVinciImportSubflowsResponse) GetEnvironment() ResourceRelationshipReadOnly`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciImportSubflowsResponse) GetEnvironmentOk() (*ResourceRelationshipReadOnly, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciImportSubflowsResponse) SetEnvironment(v ResourceRelationshipReadOnly)`

SetEnvironment sets Environment field to given value.


### GetFlow

`func (o *DaVinciImportSubflowsResponse) GetFlow() DaVinciImportFlowResponseFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciImportSubflowsResponse) GetFlowOk() (*DaVinciImportFlowResponseFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciImportSubflowsResponse) SetFlow(v DaVinciImportFlowResponseFlow)`

SetFlow sets Flow field to given value.


### GetPublishedVersion

`func (o *DaVinciImportSubflowsResponse) GetPublishedVersion() float32`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *DaVinciImportSubflowsResponse) GetPublishedVersionOk() (*float32, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *DaVinciImportSubflowsResponse) SetPublishedVersion(v float32)`

SetPublishedVersion sets PublishedVersion field to given value.


### GetVersion

`func (o *DaVinciImportSubflowsResponse) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciImportSubflowsResponse) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciImportSubflowsResponse) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetAlias

`func (o *DaVinciImportSubflowsResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DaVinciImportSubflowsResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DaVinciImportSubflowsResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DaVinciImportSubflowsResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetClonedFrom

`func (o *DaVinciImportSubflowsResponse) GetClonedFrom() float32`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *DaVinciImportSubflowsResponse) GetClonedFromOk() (*float32, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *DaVinciImportSubflowsResponse) SetClonedFrom(v float32)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *DaVinciImportSubflowsResponse) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### GetColor

`func (o *DaVinciImportSubflowsResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DaVinciImportSubflowsResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DaVinciImportSubflowsResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DaVinciImportSubflowsResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetConnectors

`func (o *DaVinciImportSubflowsResponse) GetConnectors() []ResourceRelationshipDaVinciReadOnly`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DaVinciImportSubflowsResponse) GetConnectorsOk() (*[]ResourceRelationshipDaVinciReadOnly, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DaVinciImportSubflowsResponse) SetConnectors(v []ResourceRelationshipDaVinciReadOnly)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *DaVinciImportSubflowsResponse) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciImportSubflowsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciImportSubflowsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciImportSubflowsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciImportSubflowsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *DaVinciImportSubflowsResponse) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *DaVinciImportSubflowsResponse) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *DaVinciImportSubflowsResponse) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *DaVinciImportSubflowsResponse) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DaVinciImportSubflowsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciImportSubflowsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciImportSubflowsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaVinciImportSubflowsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DaVinciImportSubflowsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DaVinciImportSubflowsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DaVinciImportSubflowsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DaVinciImportSubflowsResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGraphData

`func (o *DaVinciImportSubflowsResponse) GetGraphData() DaVinciFlowGraphDataResponse`

GetGraphData returns the GraphData field if non-nil, zero value otherwise.

### GetGraphDataOk

`func (o *DaVinciImportSubflowsResponse) GetGraphDataOk() (*DaVinciFlowGraphDataResponse, bool)`

GetGraphDataOk returns a tuple with the GraphData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphData

`func (o *DaVinciImportSubflowsResponse) SetGraphData(v DaVinciFlowGraphDataResponse)`

SetGraphData sets GraphData field to given value.

### HasGraphData

`func (o *DaVinciImportSubflowsResponse) HasGraphData() bool`

HasGraphData returns a boolean if a field has been set.

### GetInputSchema

`func (o *DaVinciImportSubflowsResponse) GetInputSchema() []DaVinciFlowInputSchemaResponseItem`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *DaVinciImportSubflowsResponse) GetInputSchemaOk() (*[]DaVinciFlowInputSchemaResponseItem, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *DaVinciImportSubflowsResponse) SetInputSchema(v []DaVinciFlowInputSchemaResponseItem)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *DaVinciImportSubflowsResponse) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *DaVinciImportSubflowsResponse) GetOutputSchema() DaVinciImportFlowResponseOutputSchema`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *DaVinciImportSubflowsResponse) GetOutputSchemaOk() (*DaVinciImportFlowResponseOutputSchema, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *DaVinciImportSubflowsResponse) SetOutputSchema(v DaVinciImportFlowResponseOutputSchema)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *DaVinciImportSubflowsResponse) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetSettings

`func (o *DaVinciImportSubflowsResponse) GetSettings() DaVinciFlowSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DaVinciImportSubflowsResponse) GetSettingsOk() (*DaVinciFlowSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DaVinciImportSubflowsResponse) SetSettings(v DaVinciFlowSettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DaVinciImportSubflowsResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciImportSubflowsResponse) GetTrigger() DaVinciImportFlowResponseTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciImportSubflowsResponse) GetTriggerOk() (*DaVinciImportFlowResponseTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciImportSubflowsResponse) SetTrigger(v DaVinciImportFlowResponseTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciImportSubflowsResponse) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciImportSubflowsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciImportSubflowsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciImportSubflowsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciImportSubflowsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdates

`func (o *DaVinciImportSubflowsResponse) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *DaVinciImportSubflowsResponse) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *DaVinciImportSubflowsResponse) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *DaVinciImportSubflowsResponse) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVariables

`func (o *DaVinciImportSubflowsResponse) GetVariables() []map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DaVinciImportSubflowsResponse) GetVariablesOk() (*[]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DaVinciImportSubflowsResponse) SetVariables(v []map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DaVinciImportSubflowsResponse) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetEmbedded

`func (o *DaVinciImportSubflowsResponse) GetEmbedded() []DaVinciImportFlowResponse`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DaVinciImportSubflowsResponse) GetEmbeddedOk() (*[]DaVinciImportFlowResponse, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DaVinciImportSubflowsResponse) SetEmbedded(v []DaVinciImportFlowResponse)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DaVinciImportSubflowsResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


