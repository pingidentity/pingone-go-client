# DaVinciFlowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciFlowResponseLinks**](DaVinciFlowResponseLinks.md) |  | 
**Environment** | [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**Connectors** | Pointer to [**[]ResourceRelationshipDaVinciReadOnly**](ResourceRelationshipDaVinciReadOnly.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**CurrentVersion** | Pointer to **float32** |  | [optional] 
**DeployedAt** | Pointer to **time.Time** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**DvlinterErrorCount** | Pointer to **float32** |  | [optional] 
**DvlinterWarningCount** | Pointer to **float32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**GraphData** | Pointer to [**DaVinciFlowGraphDataResponse**](DaVinciFlowGraphDataResponse.md) |  | [optional] 
**InputSchema** | Pointer to [**[]DaVinciFlowInputSchemaResponseItem**](DaVinciFlowInputSchemaResponseItem.md) |  | [optional] 
**OutputSchema** | Pointer to [**DaVinciFlowResponseOutputSchema**](DaVinciFlowResponseOutputSchema.md) |  | [optional] 
**PublishedVersion** | Pointer to **float32** |  | [optional] 
**Settings** | Pointer to [**DaVinciFlowSettingsResponse**](DaVinciFlowSettingsResponse.md) |  | [optional] 
**Trigger** | Pointer to [**DaVinciFlowTriggerResponse**](DaVinciFlowTriggerResponse.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDaVinciFlowResponse

`func NewDaVinciFlowResponse(links DaVinciFlowResponseLinks, environment ResourceRelationshipReadOnly, id string, name string, ) *DaVinciFlowResponse`

NewDaVinciFlowResponse instantiates a new DaVinciFlowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowResponseWithDefaults

`func NewDaVinciFlowResponseWithDefaults() *DaVinciFlowResponse`

NewDaVinciFlowResponseWithDefaults instantiates a new DaVinciFlowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciFlowResponse) GetLinks() DaVinciFlowResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciFlowResponse) GetLinksOk() (*DaVinciFlowResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciFlowResponse) SetLinks(v DaVinciFlowResponseLinks)`

SetLinks sets Links field to given value.


### GetEnvironment

`func (o *DaVinciFlowResponse) GetEnvironment() ResourceRelationshipReadOnly`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciFlowResponse) GetEnvironmentOk() (*ResourceRelationshipReadOnly, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciFlowResponse) SetEnvironment(v ResourceRelationshipReadOnly)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *DaVinciFlowResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciFlowResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciFlowResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DaVinciFlowResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciFlowResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciFlowResponse) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *DaVinciFlowResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DaVinciFlowResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DaVinciFlowResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DaVinciFlowResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetConnectors

`func (o *DaVinciFlowResponse) GetConnectors() []ResourceRelationshipDaVinciReadOnly`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DaVinciFlowResponse) GetConnectorsOk() (*[]ResourceRelationshipDaVinciReadOnly, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DaVinciFlowResponse) SetConnectors(v []ResourceRelationshipDaVinciReadOnly)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *DaVinciFlowResponse) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciFlowResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciFlowResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciFlowResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciFlowResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *DaVinciFlowResponse) GetCurrentVersion() float32`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *DaVinciFlowResponse) GetCurrentVersionOk() (*float32, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *DaVinciFlowResponse) SetCurrentVersion(v float32)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *DaVinciFlowResponse) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetDeployedAt

`func (o *DaVinciFlowResponse) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *DaVinciFlowResponse) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *DaVinciFlowResponse) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *DaVinciFlowResponse) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetDescription

`func (o *DaVinciFlowResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciFlowResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciFlowResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaVinciFlowResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDvlinterErrorCount

`func (o *DaVinciFlowResponse) GetDvlinterErrorCount() float32`

GetDvlinterErrorCount returns the DvlinterErrorCount field if non-nil, zero value otherwise.

### GetDvlinterErrorCountOk

`func (o *DaVinciFlowResponse) GetDvlinterErrorCountOk() (*float32, bool)`

GetDvlinterErrorCountOk returns a tuple with the DvlinterErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvlinterErrorCount

`func (o *DaVinciFlowResponse) SetDvlinterErrorCount(v float32)`

SetDvlinterErrorCount sets DvlinterErrorCount field to given value.

### HasDvlinterErrorCount

`func (o *DaVinciFlowResponse) HasDvlinterErrorCount() bool`

HasDvlinterErrorCount returns a boolean if a field has been set.

### GetDvlinterWarningCount

`func (o *DaVinciFlowResponse) GetDvlinterWarningCount() float32`

GetDvlinterWarningCount returns the DvlinterWarningCount field if non-nil, zero value otherwise.

### GetDvlinterWarningCountOk

`func (o *DaVinciFlowResponse) GetDvlinterWarningCountOk() (*float32, bool)`

GetDvlinterWarningCountOk returns a tuple with the DvlinterWarningCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDvlinterWarningCount

`func (o *DaVinciFlowResponse) SetDvlinterWarningCount(v float32)`

SetDvlinterWarningCount sets DvlinterWarningCount field to given value.

### HasDvlinterWarningCount

`func (o *DaVinciFlowResponse) HasDvlinterWarningCount() bool`

HasDvlinterWarningCount returns a boolean if a field has been set.

### GetEnabled

`func (o *DaVinciFlowResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DaVinciFlowResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DaVinciFlowResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DaVinciFlowResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGraphData

`func (o *DaVinciFlowResponse) GetGraphData() DaVinciFlowGraphDataResponse`

GetGraphData returns the GraphData field if non-nil, zero value otherwise.

### GetGraphDataOk

`func (o *DaVinciFlowResponse) GetGraphDataOk() (*DaVinciFlowGraphDataResponse, bool)`

GetGraphDataOk returns a tuple with the GraphData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphData

`func (o *DaVinciFlowResponse) SetGraphData(v DaVinciFlowGraphDataResponse)`

SetGraphData sets GraphData field to given value.

### HasGraphData

`func (o *DaVinciFlowResponse) HasGraphData() bool`

HasGraphData returns a boolean if a field has been set.

### GetInputSchema

`func (o *DaVinciFlowResponse) GetInputSchema() []DaVinciFlowInputSchemaResponseItem`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *DaVinciFlowResponse) GetInputSchemaOk() (*[]DaVinciFlowInputSchemaResponseItem, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *DaVinciFlowResponse) SetInputSchema(v []DaVinciFlowInputSchemaResponseItem)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *DaVinciFlowResponse) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *DaVinciFlowResponse) GetOutputSchema() DaVinciFlowResponseOutputSchema`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *DaVinciFlowResponse) GetOutputSchemaOk() (*DaVinciFlowResponseOutputSchema, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *DaVinciFlowResponse) SetOutputSchema(v DaVinciFlowResponseOutputSchema)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *DaVinciFlowResponse) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetPublishedVersion

`func (o *DaVinciFlowResponse) GetPublishedVersion() float32`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *DaVinciFlowResponse) GetPublishedVersionOk() (*float32, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *DaVinciFlowResponse) SetPublishedVersion(v float32)`

SetPublishedVersion sets PublishedVersion field to given value.

### HasPublishedVersion

`func (o *DaVinciFlowResponse) HasPublishedVersion() bool`

HasPublishedVersion returns a boolean if a field has been set.

### GetSettings

`func (o *DaVinciFlowResponse) GetSettings() DaVinciFlowSettingsResponse`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DaVinciFlowResponse) GetSettingsOk() (*DaVinciFlowSettingsResponse, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DaVinciFlowResponse) SetSettings(v DaVinciFlowSettingsResponse)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DaVinciFlowResponse) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciFlowResponse) GetTrigger() DaVinciFlowTriggerResponse`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciFlowResponse) GetTriggerOk() (*DaVinciFlowTriggerResponse, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciFlowResponse) SetTrigger(v DaVinciFlowTriggerResponse)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciFlowResponse) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciFlowResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciFlowResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciFlowResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciFlowResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


