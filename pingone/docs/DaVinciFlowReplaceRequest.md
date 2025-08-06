# DaVinciFlowReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Color** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**GraphData** | Pointer to [**DaVinciFlowGraphDataRequest**](DaVinciFlowGraphDataRequest.md) |  | [optional] 
**InputSchema** | Pointer to [**[]DaVinciFlowInputSchemaRequestItem**](DaVinciFlowInputSchemaRequestItem.md) |  | [optional] 
**OutputSchema** | Pointer to [**DaVinciFlowReplaceRequestOutputSchema**](DaVinciFlowReplaceRequestOutputSchema.md) |  | [optional] 
**Settings** | Pointer to [**DaVinciFlowSettingsRequest**](DaVinciFlowSettingsRequest.md) |  | [optional] 
**Trigger** | Pointer to [**DaVinciFlowReplaceRequestTrigger**](DaVinciFlowReplaceRequestTrigger.md) |  | [optional] 

## Methods

### NewDaVinciFlowReplaceRequest

`func NewDaVinciFlowReplaceRequest(name string, ) *DaVinciFlowReplaceRequest`

NewDaVinciFlowReplaceRequest instantiates a new DaVinciFlowReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowReplaceRequestWithDefaults

`func NewDaVinciFlowReplaceRequestWithDefaults() *DaVinciFlowReplaceRequest`

NewDaVinciFlowReplaceRequestWithDefaults instantiates a new DaVinciFlowReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DaVinciFlowReplaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciFlowReplaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciFlowReplaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetColor

`func (o *DaVinciFlowReplaceRequest) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *DaVinciFlowReplaceRequest) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *DaVinciFlowReplaceRequest) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *DaVinciFlowReplaceRequest) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetDescription

`func (o *DaVinciFlowReplaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciFlowReplaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciFlowReplaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DaVinciFlowReplaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGraphData

`func (o *DaVinciFlowReplaceRequest) GetGraphData() DaVinciFlowGraphDataRequest`

GetGraphData returns the GraphData field if non-nil, zero value otherwise.

### GetGraphDataOk

`func (o *DaVinciFlowReplaceRequest) GetGraphDataOk() (*DaVinciFlowGraphDataRequest, bool)`

GetGraphDataOk returns a tuple with the GraphData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraphData

`func (o *DaVinciFlowReplaceRequest) SetGraphData(v DaVinciFlowGraphDataRequest)`

SetGraphData sets GraphData field to given value.

### HasGraphData

`func (o *DaVinciFlowReplaceRequest) HasGraphData() bool`

HasGraphData returns a boolean if a field has been set.

### GetInputSchema

`func (o *DaVinciFlowReplaceRequest) GetInputSchema() []DaVinciFlowInputSchemaRequestItem`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *DaVinciFlowReplaceRequest) GetInputSchemaOk() (*[]DaVinciFlowInputSchemaRequestItem, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *DaVinciFlowReplaceRequest) SetInputSchema(v []DaVinciFlowInputSchemaRequestItem)`

SetInputSchema sets InputSchema field to given value.

### HasInputSchema

`func (o *DaVinciFlowReplaceRequest) HasInputSchema() bool`

HasInputSchema returns a boolean if a field has been set.

### GetOutputSchema

`func (o *DaVinciFlowReplaceRequest) GetOutputSchema() DaVinciFlowReplaceRequestOutputSchema`

GetOutputSchema returns the OutputSchema field if non-nil, zero value otherwise.

### GetOutputSchemaOk

`func (o *DaVinciFlowReplaceRequest) GetOutputSchemaOk() (*DaVinciFlowReplaceRequestOutputSchema, bool)`

GetOutputSchemaOk returns a tuple with the OutputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputSchema

`func (o *DaVinciFlowReplaceRequest) SetOutputSchema(v DaVinciFlowReplaceRequestOutputSchema)`

SetOutputSchema sets OutputSchema field to given value.

### HasOutputSchema

`func (o *DaVinciFlowReplaceRequest) HasOutputSchema() bool`

HasOutputSchema returns a boolean if a field has been set.

### GetSettings

`func (o *DaVinciFlowReplaceRequest) GetSettings() DaVinciFlowSettingsRequest`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *DaVinciFlowReplaceRequest) GetSettingsOk() (*DaVinciFlowSettingsRequest, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *DaVinciFlowReplaceRequest) SetSettings(v DaVinciFlowSettingsRequest)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *DaVinciFlowReplaceRequest) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciFlowReplaceRequest) GetTrigger() DaVinciFlowReplaceRequestTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciFlowReplaceRequest) GetTriggerOk() (*DaVinciFlowReplaceRequestTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciFlowReplaceRequest) SetTrigger(v DaVinciFlowReplaceRequestTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciFlowReplaceRequest) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


