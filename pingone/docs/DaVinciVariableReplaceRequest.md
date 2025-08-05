# DaVinciVariableReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Context** | [**DaVinciVariableReplaceRequestContext**](DaVinciVariableReplaceRequestContext.md) |  | 
**DataType** | [**DaVinciVariableReplaceRequestDataType**](DaVinciVariableReplaceRequestDataType.md) |  | 
**Mutable** | **bool** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Flow** | Pointer to [**ResourceRelationshipDaVinci**](ResourceRelationshipDaVinci.md) |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] [default to 2000]
**Min** | Pointer to **int32** |  | [optional] [default to 0]
**Value** | Pointer to [**DaVinciVariableReplaceRequestValue**](DaVinciVariableReplaceRequestValue.md) |  | [optional] 

## Methods

### NewDaVinciVariableReplaceRequest

`func NewDaVinciVariableReplaceRequest(name string, context DaVinciVariableReplaceRequestContext, dataType DaVinciVariableReplaceRequestDataType, mutable bool, ) *DaVinciVariableReplaceRequest`

NewDaVinciVariableReplaceRequest instantiates a new DaVinciVariableReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciVariableReplaceRequestWithDefaults

`func NewDaVinciVariableReplaceRequestWithDefaults() *DaVinciVariableReplaceRequest`

NewDaVinciVariableReplaceRequestWithDefaults instantiates a new DaVinciVariableReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DaVinciVariableReplaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciVariableReplaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciVariableReplaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContext

`func (o *DaVinciVariableReplaceRequest) GetContext() DaVinciVariableReplaceRequestContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DaVinciVariableReplaceRequest) GetContextOk() (*DaVinciVariableReplaceRequestContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DaVinciVariableReplaceRequest) SetContext(v DaVinciVariableReplaceRequestContext)`

SetContext sets Context field to given value.


### GetDataType

`func (o *DaVinciVariableReplaceRequest) GetDataType() DaVinciVariableReplaceRequestDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DaVinciVariableReplaceRequest) GetDataTypeOk() (*DaVinciVariableReplaceRequestDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DaVinciVariableReplaceRequest) SetDataType(v DaVinciVariableReplaceRequestDataType)`

SetDataType sets DataType field to given value.


### GetMutable

`func (o *DaVinciVariableReplaceRequest) GetMutable() bool`

GetMutable returns the Mutable field if non-nil, zero value otherwise.

### GetMutableOk

`func (o *DaVinciVariableReplaceRequest) GetMutableOk() (*bool, bool)`

GetMutableOk returns a tuple with the Mutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutable

`func (o *DaVinciVariableReplaceRequest) SetMutable(v bool)`

SetMutable sets Mutable field to given value.


### GetDisplayName

`func (o *DaVinciVariableReplaceRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DaVinciVariableReplaceRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DaVinciVariableReplaceRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DaVinciVariableReplaceRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFlow

`func (o *DaVinciVariableReplaceRequest) GetFlow() ResourceRelationshipDaVinci`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciVariableReplaceRequest) GetFlowOk() (*ResourceRelationshipDaVinci, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciVariableReplaceRequest) SetFlow(v ResourceRelationshipDaVinci)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *DaVinciVariableReplaceRequest) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetMax

`func (o *DaVinciVariableReplaceRequest) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DaVinciVariableReplaceRequest) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DaVinciVariableReplaceRequest) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *DaVinciVariableReplaceRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *DaVinciVariableReplaceRequest) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DaVinciVariableReplaceRequest) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DaVinciVariableReplaceRequest) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *DaVinciVariableReplaceRequest) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetValue

`func (o *DaVinciVariableReplaceRequest) GetValue() DaVinciVariableReplaceRequestValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DaVinciVariableReplaceRequest) GetValueOk() (*DaVinciVariableReplaceRequestValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DaVinciVariableReplaceRequest) SetValue(v DaVinciVariableReplaceRequestValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *DaVinciVariableReplaceRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


