# DaVinciFlowCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GraphData** | Pointer to [**DaVinciFlowGraphDataRequest**](DaVinciFlowGraphDataRequest.md) |  | [optional] 
**InputSchema** | Pointer to [**[]DaVinciFlowInputSchemaRequestItem**](DaVinciFlowInputSchemaRequestItem.md) |  | [optional] 
**OutputSchema** | Pointer to [**DaVinciFlowCreateRequestOutputSchema**](DaVinciFlowCreateRequestOutputSchema.md) |  | [optional] 
**Settings** | Pointer to [**DaVinciFlowSettingsRequest**](DaVinciFlowSettingsRequest.md) |  | [optional] 
**Trigger** | Pointer to [**DaVinciFlowCreateRequestTrigger**](DaVinciFlowCreateRequestTrigger.md) |  | [optional] 

## Methods

### NewDaVinciFlowCreateRequest

`func NewDaVinciFlowCreateRequest(name string, ) *DaVinciFlowCreateRequest`

NewDaVinciFlowCreateRequest instantiates a new DaVinciFlowCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowCreateRequestWithDefaults

`func NewDaVinciFlowCreateRequestWithDefaults() *DaVinciFlowCreateRequest`

NewDaVinciFlowCreateRequestWithDefaults instantiates a new DaVinciFlowCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DaVinciFlowCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciFlowCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciFlowCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *DaVinciFlowCreateRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DaVinciFlowCreateRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DaVinciFlowCreateRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DaVinciFlowCreateRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *DaVinciFlowCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciFlowCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciFlowCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaVinciFlowCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGraphData

`func (o *DaVinciFlowCreateRequest) GetGraphData() DaVinciFlowGraphDataRequest`

GetGraphData returns the GraphData field if non-nil, zero value otherwise.

### GetGraphDataOk

`func (o *DaVinciFlowCreateRequest) GetGraphDataOk() (*DaVinciFlowGraphDataRequest, bool)`

GetGraphDataOk returns a tuple with the GraphData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphData

`func (o *DaVinciFlowCreateRequest) SetGraphData(v DaVinciFlowGraphDataRequest)`

SetGraphData sets GraphData field to given value.

### HasGraphData

`func (o *DaVinciFlowCreateRequest) HasGraphData() bool`

HasGraphData returns a boolean if a field has been set.

### GetInputSchema

`func (o *DaVinciFlowCreateRequest) GetInputSchema() []DaVinciFlowInputSchemaRequestItem`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *DaVinciFlowCreateRequest) GetInputSchemaOk() (*[]DaVinciFlowInputSchemaRequestItem, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *DaVinciFlowCreateRequest) SetInputSchema(v []DaVinciFlowInputSchemaRequestItem)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *DaVinciFlowCreateRequest) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *DaVinciFlowCreateRequest) GetOutputSchema() DaVinciFlowCreateRequestOutputSchema`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *DaVinciFlowCreateRequest) GetOutputSchemaOk() (*DaVinciFlowCreateRequestOutputSchema, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *DaVinciFlowCreateRequest) SetOutputSchema(v DaVinciFlowCreateRequestOutputSchema)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *DaVinciFlowCreateRequest) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetSettings

`func (o *DaVinciFlowCreateRequest) GetSettings() DaVinciFlowSettingsRequest`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DaVinciFlowCreateRequest) GetSettingsOk() (*DaVinciFlowSettingsRequest, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DaVinciFlowCreateRequest) SetSettings(v DaVinciFlowSettingsRequest)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DaVinciFlowCreateRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciFlowCreateRequest) GetTrigger() DaVinciFlowCreateRequestTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciFlowCreateRequest) GetTriggerOk() (*DaVinciFlowCreateRequestTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciFlowCreateRequest) SetTrigger(v DaVinciFlowCreateRequestTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciFlowCreateRequest) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


