# DaVinciExportFlowVersionSubflowsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciExportFlowVersionResponseLinks**](DaVinciExportFlowVersionResponseLinks.md) |  | 
**Environment** | [**DaVinciExportFlowVersionResponseEnvironment**](DaVinciExportFlowVersionResponseEnvironment.md) |  | 
**Flow** | [**DaVinciExportFlowVersionResponseFlow**](DaVinciExportFlowVersionResponseFlow.md) |  | 
**PublishedVersion** | **float32** |  | 
**Version** | **float32** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**ClonedFrom** | Pointer to **float32** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Connectors** | Pointer to [**[]DaVinciExportFlowVersionResponseConnector**](DaVinciExportFlowVersionResponseConnector.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeployedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**GraphData** | Pointer to [**DaVinciFlowGraphDataResponse**](DaVinciFlowGraphDataResponse.md) |  | [optional] 
**InputSchema** | Pointer to [**[]DaVinciFlowInputSchemaResponseItem**](DaVinciFlowInputSchemaResponseItem.md) |  | [optional] 
**OutputSchema** | Pointer to [**DaVinciExportFlowVersionResponseOutputSchema**](DaVinciExportFlowVersionResponseOutputSchema.md) |  | [optional] 
**Settings** | Pointer to [**DaVinciFlowSettingsResponse**](DaVinciFlowSettingsResponse.md) |  | [optional] 
**Trigger** | Pointer to [**DaVinciFlowTriggerResponse**](DaVinciFlowTriggerResponse.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Updates** | Pointer to **[]string** |  | [optional] 
**Variables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Embedded** | Pointer to [**[]DaVinciExportFlowVersionResponse**](DaVinciExportFlowVersionResponse.md) |  | [optional] 

## Methods

### NewDaVinciExportFlowVersionSubflowsResponse

`func NewDaVinciExportFlowVersionSubflowsResponse(links DaVinciExportFlowVersionResponseLinks, environment DaVinciExportFlowVersionResponseEnvironment, flow DaVinciExportFlowVersionResponseFlow, publishedVersion float32, version float32, ) *DaVinciExportFlowVersionSubflowsResponse`

NewDaVinciExportFlowVersionSubflowsResponse instantiates a new DaVinciExportFlowVersionSubflowsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciExportFlowVersionSubflowsResponseWithDefaults

`func NewDaVinciExportFlowVersionSubflowsResponseWithDefaults() *DaVinciExportFlowVersionSubflowsResponse`

NewDaVinciExportFlowVersionSubflowsResponseWithDefaults instantiates a new DaVinciExportFlowVersionSubflowsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetLinks() DaVinciExportFlowVersionResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetLinksOk() (*DaVinciExportFlowVersionResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetLinks(v DaVinciExportFlowVersionResponseLinks)`

SetLinks sets Links field to given value.


### GetEnvironment

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetEnvironment() DaVinciExportFlowVersionResponseEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetEnvironmentOk() (*DaVinciExportFlowVersionResponseEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetEnvironment(v DaVinciExportFlowVersionResponseEnvironment)`

SetEnvironment sets Environment field to given value.


### GetFlow

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetFlow() DaVinciExportFlowVersionResponseFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetFlowOk() (*DaVinciExportFlowVersionResponseFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetFlow(v DaVinciExportFlowVersionResponseFlow)`

SetFlow sets Flow field to given value.


### GetPublishedVersion

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetPublishedVersion() float32`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetPublishedVersionOk() (*float32, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetPublishedVersion(v float32)`

SetPublishedVersion sets PublishedVersion field to given value.


### GetVersion

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetAlias

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetClonedFrom

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetClonedFrom() float32`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetClonedFromOk() (*float32, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetClonedFrom(v float32)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### GetColor

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetConnectors

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetConnectors() []DaVinciExportFlowVersionResponseConnector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetConnectorsOk() (*[]DaVinciExportFlowVersionResponseConnector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetConnectors(v []DaVinciExportFlowVersionResponseConnector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGraphData

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetGraphData() DaVinciFlowGraphDataResponse`

GetGraphData returns the GraphData field if non-nil, zero value otherwise.

### GetGraphDataOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetGraphDataOk() (*DaVinciFlowGraphDataResponse, bool)`

GetGraphDataOk returns a tuple with the GraphData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphData

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetGraphData(v DaVinciFlowGraphDataResponse)`

SetGraphData sets GraphData field to given value.

### HasGraphData

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasGraphData() bool`

HasGraphData returns a boolean if a field has been set.

### GetInputSchema

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetInputSchema() []DaVinciFlowInputSchemaResponseItem`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetInputSchemaOk() (*[]DaVinciFlowInputSchemaResponseItem, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetInputSchema(v []DaVinciFlowInputSchemaResponseItem)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetOutputSchema() DaVinciExportFlowVersionResponseOutputSchema`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetOutputSchemaOk() (*DaVinciExportFlowVersionResponseOutputSchema, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetOutputSchema(v DaVinciExportFlowVersionResponseOutputSchema)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetSettings

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetSettings() DaVinciFlowSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetSettingsOk() (*DaVinciFlowSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetSettings(v DaVinciFlowSettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetTrigger() DaVinciFlowTriggerResponse`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetTriggerOk() (*DaVinciFlowTriggerResponse, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetTrigger(v DaVinciFlowTriggerResponse)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdates

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVariables

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetVariables() []map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetVariablesOk() (*[]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetVariables(v []map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetEmbedded

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetEmbedded() []DaVinciExportFlowVersionResponse`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *DaVinciExportFlowVersionSubflowsResponse) GetEmbeddedOk() (*[]DaVinciExportFlowVersionResponse, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *DaVinciExportFlowVersionSubflowsResponse) SetEmbedded(v []DaVinciExportFlowVersionResponse)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *DaVinciExportFlowVersionSubflowsResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


