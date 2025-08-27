# DaVinciExportFlowVersionResponse

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
**Trigger** | Pointer to [**DaVinciExportFlowVersionResponseTrigger**](DaVinciExportFlowVersionResponseTrigger.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Updates** | Pointer to **[]string** |  | [optional] 
**Variables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewDaVinciExportFlowVersionResponse

`func NewDaVinciExportFlowVersionResponse(links DaVinciExportFlowVersionResponseLinks, environment DaVinciExportFlowVersionResponseEnvironment, flow DaVinciExportFlowVersionResponseFlow, publishedVersion float32, version float32, ) *DaVinciExportFlowVersionResponse`

NewDaVinciExportFlowVersionResponse instantiates a new DaVinciExportFlowVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciExportFlowVersionResponseWithDefaults

`func NewDaVinciExportFlowVersionResponseWithDefaults() *DaVinciExportFlowVersionResponse`

NewDaVinciExportFlowVersionResponseWithDefaults instantiates a new DaVinciExportFlowVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciExportFlowVersionResponse) GetLinks() DaVinciExportFlowVersionResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciExportFlowVersionResponse) GetLinksOk() (*DaVinciExportFlowVersionResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciExportFlowVersionResponse) SetLinks(v DaVinciExportFlowVersionResponseLinks)`

SetLinks sets Links field to given value.


### GetEnvironment

`func (o *DaVinciExportFlowVersionResponse) GetEnvironment() DaVinciExportFlowVersionResponseEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciExportFlowVersionResponse) GetEnvironmentOk() (*DaVinciExportFlowVersionResponseEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciExportFlowVersionResponse) SetEnvironment(v DaVinciExportFlowVersionResponseEnvironment)`

SetEnvironment sets Environment field to given value.


### GetFlow

`func (o *DaVinciExportFlowVersionResponse) GetFlow() DaVinciExportFlowVersionResponseFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciExportFlowVersionResponse) GetFlowOk() (*DaVinciExportFlowVersionResponseFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciExportFlowVersionResponse) SetFlow(v DaVinciExportFlowVersionResponseFlow)`

SetFlow sets Flow field to given value.


### GetPublishedVersion

`func (o *DaVinciExportFlowVersionResponse) GetPublishedVersion() float32`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *DaVinciExportFlowVersionResponse) GetPublishedVersionOk() (*float32, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *DaVinciExportFlowVersionResponse) SetPublishedVersion(v float32)`

SetPublishedVersion sets PublishedVersion field to given value.


### GetVersion

`func (o *DaVinciExportFlowVersionResponse) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciExportFlowVersionResponse) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciExportFlowVersionResponse) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetAlias

`func (o *DaVinciExportFlowVersionResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DaVinciExportFlowVersionResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DaVinciExportFlowVersionResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DaVinciExportFlowVersionResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetClonedFrom

`func (o *DaVinciExportFlowVersionResponse) GetClonedFrom() float32`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *DaVinciExportFlowVersionResponse) GetClonedFromOk() (*float32, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *DaVinciExportFlowVersionResponse) SetClonedFrom(v float32)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *DaVinciExportFlowVersionResponse) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### GetColor

`func (o *DaVinciExportFlowVersionResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DaVinciExportFlowVersionResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DaVinciExportFlowVersionResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DaVinciExportFlowVersionResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetConnectors

`func (o *DaVinciExportFlowVersionResponse) GetConnectors() []DaVinciExportFlowVersionResponseConnector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DaVinciExportFlowVersionResponse) GetConnectorsOk() (*[]DaVinciExportFlowVersionResponseConnector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DaVinciExportFlowVersionResponse) SetConnectors(v []DaVinciExportFlowVersionResponseConnector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *DaVinciExportFlowVersionResponse) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciExportFlowVersionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciExportFlowVersionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciExportFlowVersionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciExportFlowVersionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *DaVinciExportFlowVersionResponse) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *DaVinciExportFlowVersionResponse) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *DaVinciExportFlowVersionResponse) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *DaVinciExportFlowVersionResponse) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DaVinciExportFlowVersionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciExportFlowVersionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciExportFlowVersionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaVinciExportFlowVersionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DaVinciExportFlowVersionResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DaVinciExportFlowVersionResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DaVinciExportFlowVersionResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DaVinciExportFlowVersionResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGraphData

`func (o *DaVinciExportFlowVersionResponse) GetGraphData() DaVinciFlowGraphDataResponse`

GetGraphData returns the GraphData field if non-nil, zero value otherwise.

### GetGraphDataOk

`func (o *DaVinciExportFlowVersionResponse) GetGraphDataOk() (*DaVinciFlowGraphDataResponse, bool)`

GetGraphDataOk returns a tuple with the GraphData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphData

`func (o *DaVinciExportFlowVersionResponse) SetGraphData(v DaVinciFlowGraphDataResponse)`

SetGraphData sets GraphData field to given value.

### HasGraphData

`func (o *DaVinciExportFlowVersionResponse) HasGraphData() bool`

HasGraphData returns a boolean if a field has been set.

### GetInputSchema

`func (o *DaVinciExportFlowVersionResponse) GetInputSchema() []DaVinciFlowInputSchemaResponseItem`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *DaVinciExportFlowVersionResponse) GetInputSchemaOk() (*[]DaVinciFlowInputSchemaResponseItem, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *DaVinciExportFlowVersionResponse) SetInputSchema(v []DaVinciFlowInputSchemaResponseItem)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *DaVinciExportFlowVersionResponse) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *DaVinciExportFlowVersionResponse) GetOutputSchema() DaVinciExportFlowVersionResponseOutputSchema`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *DaVinciExportFlowVersionResponse) GetOutputSchemaOk() (*DaVinciExportFlowVersionResponseOutputSchema, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *DaVinciExportFlowVersionResponse) SetOutputSchema(v DaVinciExportFlowVersionResponseOutputSchema)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *DaVinciExportFlowVersionResponse) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetSettings

`func (o *DaVinciExportFlowVersionResponse) GetSettings() DaVinciFlowSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DaVinciExportFlowVersionResponse) GetSettingsOk() (*DaVinciFlowSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DaVinciExportFlowVersionResponse) SetSettings(v DaVinciFlowSettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DaVinciExportFlowVersionResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciExportFlowVersionResponse) GetTrigger() DaVinciExportFlowVersionResponseTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciExportFlowVersionResponse) GetTriggerOk() (*DaVinciExportFlowVersionResponseTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciExportFlowVersionResponse) SetTrigger(v DaVinciExportFlowVersionResponseTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciExportFlowVersionResponse) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciExportFlowVersionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciExportFlowVersionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciExportFlowVersionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciExportFlowVersionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdates

`func (o *DaVinciExportFlowVersionResponse) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *DaVinciExportFlowVersionResponse) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *DaVinciExportFlowVersionResponse) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *DaVinciExportFlowVersionResponse) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVariables

`func (o *DaVinciExportFlowVersionResponse) GetVariables() []map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *DaVinciExportFlowVersionResponse) GetVariablesOk() (*[]map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *DaVinciExportFlowVersionResponse) SetVariables(v []map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *DaVinciExportFlowVersionResponse) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


