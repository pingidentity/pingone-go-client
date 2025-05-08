# DaVinciVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciVariableLinks**](DaVinciVariableLinks.md) |  | 
**Context** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DataType** | [**DaVinciVariableDataType**](DaVinciVariableDataType.md) |  | 
**DisplayName** | Pointer to **string** |  | [optional] 
**Environment** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**Flow** | Pointer to [**ResourceRelationshipDaVinci**](ResourceRelationshipDaVinci.md) |  | [optional] 
**Id** | **string** |  | 
**Max** | Pointer to **float32** |  | [optional] 
**Min** | Pointer to **float32** |  | [optional] 
**Mutable** | Pointer to **bool** |  | [optional] 
**Name** | **string** |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to [**DaVinciVariableValue**](DaVinciVariableValue.md) |  | [optional] 

## Methods

### NewDaVinciVariable

`func NewDaVinciVariable(links DaVinciVariableLinks, dataType DaVinciVariableDataType, environment ResourceRelationshipPingOne, id string, name string, ) *DaVinciVariable`

NewDaVinciVariable instantiates a new DaVinciVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciVariableWithDefaults

`func NewDaVinciVariableWithDefaults() *DaVinciVariable`

NewDaVinciVariableWithDefaults instantiates a new DaVinciVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciVariable) GetLinks() DaVinciVariableLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciVariable) GetLinksOk() (*DaVinciVariableLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciVariable) SetLinks(v DaVinciVariableLinks)`

SetLinks sets Links field to given value.


### GetContext

`func (o *DaVinciVariable) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DaVinciVariable) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DaVinciVariable) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *DaVinciVariable) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciVariable) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciVariable) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciVariable) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciVariable) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDataType

`func (o *DaVinciVariable) GetDataType() DaVinciVariableDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DaVinciVariable) GetDataTypeOk() (*DaVinciVariableDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DaVinciVariable) SetDataType(v DaVinciVariableDataType)`

SetDataType sets DataType field to given value.


### GetDisplayName

`func (o *DaVinciVariable) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DaVinciVariable) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DaVinciVariable) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DaVinciVariable) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnvironment

`func (o *DaVinciVariable) GetEnvironment() ResourceRelationshipPingOne`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciVariable) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciVariable) SetEnvironment(v ResourceRelationshipPingOne)`

SetEnvironment sets Environment field to given value.


### GetFlow

`func (o *DaVinciVariable) GetFlow() ResourceRelationshipDaVinci`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciVariable) GetFlowOk() (*ResourceRelationshipDaVinci, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciVariable) SetFlow(v ResourceRelationshipDaVinci)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *DaVinciVariable) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetId

`func (o *DaVinciVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciVariable) SetId(v string)`

SetId sets Id field to given value.


### GetMax

`func (o *DaVinciVariable) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DaVinciVariable) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DaVinciVariable) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *DaVinciVariable) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *DaVinciVariable) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DaVinciVariable) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DaVinciVariable) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *DaVinciVariable) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMutable

`func (o *DaVinciVariable) GetMutable() bool`

GetMutable returns the Mutable field if non-nil, zero value otherwise.

### GetMutableOk

`func (o *DaVinciVariable) GetMutableOk() (*bool, bool)`

GetMutableOk returns a tuple with the Mutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutable

`func (o *DaVinciVariable) SetMutable(v bool)`

SetMutable sets Mutable field to given value.

### HasMutable

`func (o *DaVinciVariable) HasMutable() bool`

HasMutable returns a boolean if a field has been set.

### GetName

`func (o *DaVinciVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciVariable) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *DaVinciVariable) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciVariable) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciVariable) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciVariable) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *DaVinciVariable) GetValue() DaVinciVariableValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DaVinciVariable) GetValueOk() (*DaVinciVariableValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DaVinciVariable) SetValue(v DaVinciVariableValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *DaVinciVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


