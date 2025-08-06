# DaVinciImportSubflowResponse

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

### NewDaVinciImportSubflowResponse

`func NewDaVinciImportSubflowResponse(links DaVinciImportFlowResponseLinks, environment ResourceRelationshipReadOnly, flow DaVinciImportFlowResponseFlow, publishedVersion float32, version float32, ) *DaVinciImportSubflowResponse`

NewDaVinciImportSubflowResponse instantiates a new DaVinciImportSubflowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciImportSubflowResponseWithDefaults

`func NewDaVinciImportSubflowResponseWithDefaults() *DaVinciImportSubflowResponse`

NewDaVinciImportSubflowResponseWithDefaults instantiates a new DaVinciImportSubflowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciImportSubflowResponse) GetLinks() DaVinciImportFlowResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciImportSubflowResponse) GetLinksOk() (*DaVinciImportFlowResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciImportSubflowResponse) SetLinks(v DaVinciImportFlowResponseLinks)`

SetLinks sets Links field to given value.


### GetEnvironment

`func (o *DaVinciImportSubflowResponse) GetEnvironment() ResourceRelationshipReadOnly`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciImportSubflowResponse) GetEnvironmentOk() (*ResourceRelationshipReadOnly, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciImportSubflowResponse) SetEnvironment(v ResourceRelationshipReadOnly)`

SetEnvironment sets Environment field to given value.


### GetFlow

`func (o *DaVinciImportSubflowResponse) GetFlow() DaVinciImportFlowResponseFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciImportSubflowResponse) GetFlowOk() (*DaVinciImportFlowResponseFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciImportSubflowResponse) SetFlow(v DaVinciImportFlowResponseFlow)`

SetFlow sets Flow field to given value.


### GetPublishedVersion

`func (o *DaVinciImportSubflowResponse) GetPublishedVersion() float32`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *DaVinciImportSubflowResponse) GetPublishedVersionOk() (*float32, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *DaVinciImportSubflowResponse) SetPublishedVersion(v float32)`

SetPublishedVersion sets PublishedVersion field to given value.


### GetVersion

`func (o *DaVinciImportSubflowResponse) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciImportSubflowResponse) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciImportSubflowResponse) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetAlias

`func (o *DaVinciImportSubflowResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DaVinciImportSubflowResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DaVinciImportSubflowResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DaVinciImportSubflowResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetClonedFrom

`func (o *DaVinciImportSubflowResponse) GetClonedFrom() float32`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *DaVinciImportSubflowResponse) GetClonedFromOk() (*float32, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *DaVinciImportSubflowResponse) SetClonedFrom(v float32)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *DaVinciImportSubflowResponse) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### GetColor

`func (o *DaVinciImportSubflowResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DaVinciImportSubflowResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DaVinciImportSubflowResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DaVinciImportSubflowResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetConnectors

`func (o *DaVinciImportSubflowResponse) GetConnectors() []ResourceRelationshipDaVinciReadOnly`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DaVinciImportSubflowResponse) GetConnectorsOk() (*[]ResourceRelationshipDaVinciReadOnly, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DaVinciImportSubflowResponse) SetConnectors(v []ResourceRelationshipDaVinciReadOnly)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *DaVinciImportSubflowResponse) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciImportSubflowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciImportSubflowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciImportSubflowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciImportSubflowResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *DaVinciImportSubflowResponse) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *DaVinciImportSubflowResponse) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *DaVinciImportSubflowResponse) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *DaVinciImportSubflowResponse) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DaVinciImportSubflowResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciImportSubflowResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciImportSubflowResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaVinciImportSubflowResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DaVinciImportSubflowResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DaVinciImportSubflowResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DaVinciImportSubflowResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DaVinciImportSubflowResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGraphData

`func (o *DaVinciImportSubflowResponse) GetGraphData() DaVinciFlowGraphDataResponse`

GetGraphData returns the GraphData field if non-nil, zero value otherwise.

### GetGraphDataOk

`func (o *DaVinciImportSubflowResponse) GetGraphDataOk() (*DaVinciFlowGraphDataResponse, bool)`

GetGraphDataOk returns a tuple with the GraphData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphData

`func (o *DaVinciImportSubflowResponse) SetGraphData(v DaVinciFlowGraphDataResponse)`

SetGraphData sets GraphData field to given value.

### HasGraphData

`func (o *DaVinciImportSubflowResponse) HasGraphData() bool`

HasGraphData returns a boolean if a field has been set.

### GetInputSchema

`func (o *DaVinciImportSubflowResponse) GetInputSchema() []DaVinciFlowInputSchemaResponseItem`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *DaVinciImportSubflowResponse) GetInputSchemaOk() (*[]DaVinciFlowInputSchemaResponseItem, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *DaVinciImportSubflowResponse) SetInputSchema(v []DaVinciFlowInputSchemaResponseItem)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *DaVinciImportSubflowResponse) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *DaVinciImportSubflowResponse) GetOutputSchema() DaVinciImportFlowResponseOutputSchema`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *DaVinciImportSubflowResponse) GetOutputSchemaOk() (*DaVinciImportFlowResponseOutputSchema, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *DaVinciImportSubflowResponse) SetOutputSchema(v DaVinciImportFlowResponseOutputSchema)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *DaVinciImportSubflowResponse) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetSettings

`func (o *DaVinciImportSubflowResponse) GetSettings() DaVinciFlowSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DaVinciImportSubflowResponse) GetSettingsOk() (*DaVinciFlowSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DaVinciImportSubflowResponse) SetSettings(v DaVinciFlowSettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DaVinciImportSubflowResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciImportSubflowResponse) GetTrigger() DaVinciImportFlowResponseTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciImportSubflowResponse) GetTriggerOk() (*DaVinciImportFlowResponseTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciImportSubflowResponse) SetTrigger(v DaVinciImportFlowResponseTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciImportSubflowResponse) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciImportSubflowResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciImportSubflowResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciImportSubflowResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciImportSubflowResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdates

`func (o *DaVinciImportSubflowResponse) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *DaVinciImportSubflowResponse) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *DaVinciImportSubflowResponse) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *DaVinciImportSubflowResponse) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVariables

`func (o *DaVinciImportSubflowResponse) GetVariables() []map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DaVinciImportSubflowResponse) GetVariablesOk() (*[]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DaVinciImportSubflowResponse) SetVariables(v []map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DaVinciImportSubflowResponse) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetEmbedded

`func (o *DaVinciImportSubflowResponse) GetEmbedded() []DaVinciImportFlowResponse`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DaVinciImportSubflowResponse) GetEmbeddedOk() (*[]DaVinciImportFlowResponse, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DaVinciImportSubflowResponse) SetEmbedded(v []DaVinciImportFlowResponse)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DaVinciImportSubflowResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


