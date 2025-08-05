# DaVinciVariableCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Context** | [**DaVinciVariableCreateRequestContext**](DaVinciVariableCreateRequestContext.md) |  | 
**DataType** | [**DaVinciVariableCreateRequestDataType**](DaVinciVariableCreateRequestDataType.md) |  | 
**Mutable** | **bool** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Flow** | Pointer to [**ResourceRelationshipDaVinci**](ResourceRelationshipDaVinci.md) |  | [optional] 
**Max** | Pointer to **int32** |  | [optional] [default to 2000]
**Min** | Pointer to **int32** |  | [optional] [default to 0]
**Value** | Pointer to [**DaVinciVariableCreateRequestValue**](DaVinciVariableCreateRequestValue.md) |  | [optional] 

## Methods

### NewDaVinciVariableCreateRequest

`func NewDaVinciVariableCreateRequest(name string, context DaVinciVariableCreateRequestContext, dataType DaVinciVariableCreateRequestDataType, mutable bool, ) *DaVinciVariableCreateRequest`

NewDaVinciVariableCreateRequest instantiates a new DaVinciVariableCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciVariableCreateRequestWithDefaults

`func NewDaVinciVariableCreateRequestWithDefaults() *DaVinciVariableCreateRequest`

NewDaVinciVariableCreateRequestWithDefaults instantiates a new DaVinciVariableCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DaVinciVariableCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciVariableCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciVariableCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContext

`func (o *DaVinciVariableCreateRequest) GetContext() DaVinciVariableCreateRequestContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DaVinciVariableCreateRequest) GetContextOk() (*DaVinciVariableCreateRequestContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DaVinciVariableCreateRequest) SetContext(v DaVinciVariableCreateRequestContext)`

SetContext sets Context field to given value.


### GetDataType

`func (o *DaVinciVariableCreateRequest) GetDataType() DaVinciVariableCreateRequestDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DaVinciVariableCreateRequest) GetDataTypeOk() (*DaVinciVariableCreateRequestDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DaVinciVariableCreateRequest) SetDataType(v DaVinciVariableCreateRequestDataType)`

SetDataType sets DataType field to given value.


### GetMutable

`func (o *DaVinciVariableCreateRequest) GetMutable() bool`

GetMutable returns the Mutable field if non-nil, zero value otherwise.

### GetMutableOk

`func (o *DaVinciVariableCreateRequest) GetMutableOk() (*bool, bool)`

GetMutableOk returns a tuple with the Mutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutable

`func (o *DaVinciVariableCreateRequest) SetMutable(v bool)`

SetMutable sets Mutable field to given value.


### GetDisplayName

`func (o *DaVinciVariableCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DaVinciVariableCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DaVinciVariableCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DaVinciVariableCreateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFlow

`func (o *DaVinciVariableCreateRequest) GetFlow() ResourceRelationshipDaVinci`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciVariableCreateRequest) GetFlowOk() (*ResourceRelationshipDaVinci, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciVariableCreateRequest) SetFlow(v ResourceRelationshipDaVinci)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *DaVinciVariableCreateRequest) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetMax

`func (o *DaVinciVariableCreateRequest) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DaVinciVariableCreateRequest) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DaVinciVariableCreateRequest) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *DaVinciVariableCreateRequest) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *DaVinciVariableCreateRequest) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DaVinciVariableCreateRequest) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DaVinciVariableCreateRequest) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *DaVinciVariableCreateRequest) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetValue

`func (o *DaVinciVariableCreateRequest) GetValue() DaVinciVariableCreateRequestValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DaVinciVariableCreateRequest) GetValueOk() (*DaVinciVariableCreateRequestValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DaVinciVariableCreateRequest) SetValue(v DaVinciVariableCreateRequestValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *DaVinciVariableCreateRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


