# DaVinciFlowVersionDetailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciFlowVersionDetailResponseLinks**](DaVinciFlowVersionDetailResponseLinks.md) |  | 
**Alias** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Connectors** | Pointer to [**[]DaVinciFlowVersionDetailResponseConnector**](DaVinciFlowVersionDetailResponseConnector.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeployedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Flow** | Pointer to [**DaVinciFlowVersionDetailResponseFlow**](DaVinciFlowVersionDetailResponseFlow.md) |  | [optional] 
**GraphData** | Pointer to [**DaVinciFlowGraphDataResponse**](DaVinciFlowGraphDataResponse.md) |  | [optional] 
**InputSchema** | Pointer to [**[]DaVinciFlowInputSchemaResponseItem**](DaVinciFlowInputSchemaResponseItem.md) |  | [optional] 
**OutputSchema** | Pointer to [**DaVinciFlowOutputSchemaResponse**](DaVinciFlowOutputSchemaResponse.md) |  | [optional] 
**Settings** | Pointer to [**DaVinciFlowSettingsResponse**](DaVinciFlowSettingsResponse.md) |  | [optional] 
**Skcomponents** | Pointer to [**[]DaVinciFlowVersionDetailResponseSkcomponent**](DaVinciFlowVersionDetailResponseSkcomponent.md) |  | [optional] 
**Trigger** | Pointer to [**DaVinciFlowTriggerResponse**](DaVinciFlowTriggerResponse.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Updates** | Pointer to **[]string** |  | [optional] 
**Version** | Pointer to **float32** |  | [optional] 

## Methods

### NewDaVinciFlowVersionDetailResponse

`func NewDaVinciFlowVersionDetailResponse(links DaVinciFlowVersionDetailResponseLinks, ) *DaVinciFlowVersionDetailResponse`

NewDaVinciFlowVersionDetailResponse instantiates a new DaVinciFlowVersionDetailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowVersionDetailResponseWithDefaults

`func NewDaVinciFlowVersionDetailResponseWithDefaults() *DaVinciFlowVersionDetailResponse`

NewDaVinciFlowVersionDetailResponseWithDefaults instantiates a new DaVinciFlowVersionDetailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciFlowVersionDetailResponse) GetLinks() DaVinciFlowVersionDetailResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciFlowVersionDetailResponse) GetLinksOk() (*DaVinciFlowVersionDetailResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciFlowVersionDetailResponse) SetLinks(v DaVinciFlowVersionDetailResponseLinks)`

SetLinks sets Links field to given value.


### GetAlias

`func (o *DaVinciFlowVersionDetailResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DaVinciFlowVersionDetailResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DaVinciFlowVersionDetailResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DaVinciFlowVersionDetailResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetColor

`func (o *DaVinciFlowVersionDetailResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DaVinciFlowVersionDetailResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DaVinciFlowVersionDetailResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DaVinciFlowVersionDetailResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetConnectors

`func (o *DaVinciFlowVersionDetailResponse) GetConnectors() []DaVinciFlowVersionDetailResponseConnector`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DaVinciFlowVersionDetailResponse) GetConnectorsOk() (*[]DaVinciFlowVersionDetailResponseConnector, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DaVinciFlowVersionDetailResponse) SetConnectors(v []DaVinciFlowVersionDetailResponseConnector)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *DaVinciFlowVersionDetailResponse) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciFlowVersionDetailResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciFlowVersionDetailResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciFlowVersionDetailResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciFlowVersionDetailResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *DaVinciFlowVersionDetailResponse) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *DaVinciFlowVersionDetailResponse) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *DaVinciFlowVersionDetailResponse) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *DaVinciFlowVersionDetailResponse) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DaVinciFlowVersionDetailResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciFlowVersionDetailResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciFlowVersionDetailResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaVinciFlowVersionDetailResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *DaVinciFlowVersionDetailResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DaVinciFlowVersionDetailResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DaVinciFlowVersionDetailResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DaVinciFlowVersionDetailResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFlow

`func (o *DaVinciFlowVersionDetailResponse) GetFlow() DaVinciFlowVersionDetailResponseFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciFlowVersionDetailResponse) GetFlowOk() (*DaVinciFlowVersionDetailResponseFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciFlowVersionDetailResponse) SetFlow(v DaVinciFlowVersionDetailResponseFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *DaVinciFlowVersionDetailResponse) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetGraphData

`func (o *DaVinciFlowVersionDetailResponse) GetGraphData() DaVinciFlowGraphDataResponse`

GetGraphData returns the GraphData field if non-nil, zero value otherwise.

### GetGraphDataOk

`func (o *DaVinciFlowVersionDetailResponse) GetGraphDataOk() (*DaVinciFlowGraphDataResponse, bool)`

GetGraphDataOk returns a tuple with the GraphData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphData

`func (o *DaVinciFlowVersionDetailResponse) SetGraphData(v DaVinciFlowGraphDataResponse)`

SetGraphData sets GraphData field to given value.

### HasGraphData

`func (o *DaVinciFlowVersionDetailResponse) HasGraphData() bool`

HasGraphData returns a boolean if a field has been set.

### GetInputSchema

`func (o *DaVinciFlowVersionDetailResponse) GetInputSchema() []DaVinciFlowInputSchemaResponseItem`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *DaVinciFlowVersionDetailResponse) GetInputSchemaOk() (*[]DaVinciFlowInputSchemaResponseItem, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *DaVinciFlowVersionDetailResponse) SetInputSchema(v []DaVinciFlowInputSchemaResponseItem)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *DaVinciFlowVersionDetailResponse) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *DaVinciFlowVersionDetailResponse) GetOutputSchema() DaVinciFlowOutputSchemaResponse`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *DaVinciFlowVersionDetailResponse) GetOutputSchemaOk() (*DaVinciFlowOutputSchemaResponse, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *DaVinciFlowVersionDetailResponse) SetOutputSchema(v DaVinciFlowOutputSchemaResponse)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *DaVinciFlowVersionDetailResponse) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetSettings

`func (o *DaVinciFlowVersionDetailResponse) GetSettings() DaVinciFlowSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DaVinciFlowVersionDetailResponse) GetSettingsOk() (*DaVinciFlowSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DaVinciFlowVersionDetailResponse) SetSettings(v DaVinciFlowSettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DaVinciFlowVersionDetailResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetSkcomponents

`func (o *DaVinciFlowVersionDetailResponse) GetSkcomponents() []DaVinciFlowVersionDetailResponseSkcomponent`

GetSkcomponents returns the Skcomponents field if non-nil, zero value otherwise.

### GetSkcomponentsOk

`func (o *DaVinciFlowVersionDetailResponse) GetSkcomponentsOk() (*[]DaVinciFlowVersionDetailResponseSkcomponent, bool)`

GetSkcomponentsOk returns a tuple with the Skcomponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkcomponents

`func (o *DaVinciFlowVersionDetailResponse) SetSkcomponents(v []DaVinciFlowVersionDetailResponseSkcomponent)`

SetSkcomponents sets Skcomponents field to given value.

### HasSkcomponents

`func (o *DaVinciFlowVersionDetailResponse) HasSkcomponents() bool`

HasSkcomponents returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciFlowVersionDetailResponse) GetTrigger() DaVinciFlowTriggerResponse`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciFlowVersionDetailResponse) GetTriggerOk() (*DaVinciFlowTriggerResponse, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciFlowVersionDetailResponse) SetTrigger(v DaVinciFlowTriggerResponse)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciFlowVersionDetailResponse) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciFlowVersionDetailResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciFlowVersionDetailResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciFlowVersionDetailResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciFlowVersionDetailResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetUpdates

`func (o *DaVinciFlowVersionDetailResponse) GetUpdates() []string`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *DaVinciFlowVersionDetailResponse) GetUpdatesOk() (*[]string, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *DaVinciFlowVersionDetailResponse) SetUpdates(v []string)`

SetUpdates sets Updates field to given value.

### HasUpdates

`func (o *DaVinciFlowVersionDetailResponse) HasUpdates() bool`

HasUpdates returns a boolean if a field has been set.

### GetVersion

`func (o *DaVinciFlowVersionDetailResponse) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciFlowVersionDetailResponse) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciFlowVersionDetailResponse) SetVersion(v float32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DaVinciFlowVersionDetailResponse) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


