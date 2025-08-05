# DaVinciVariableResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciVariableResponseLinks**](DaVinciVariableResponseLinks.md) |  | 
**DataType** | [**DaVinciVariableResponseDataType**](DaVinciVariableResponseDataType.md) |  | 
**Environment** | [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | 
**Id** | [**uuid.UUID**](uuid.UUID.md) |  | 
**Name** | **string** |  | 
**Context** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Flow** | Pointer to [**ResourceRelationshipDaVinciReadOnly**](ResourceRelationshipDaVinciReadOnly.md) |  | [optional] 
**Max** | Pointer to **float32** |  | [optional] 
**Min** | Pointer to **float32** |  | [optional] 
**Mutable** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Value** | Pointer to [**DaVinciVariableResponseValue**](DaVinciVariableResponseValue.md) |  | [optional] 

## Methods

### NewDaVinciVariableResponse

`func NewDaVinciVariableResponse(links DaVinciVariableResponseLinks, dataType DaVinciVariableResponseDataType, environment ResourceRelationshipReadOnly, id uuid.UUID, name string, ) *DaVinciVariableResponse`

NewDaVinciVariableResponse instantiates a new DaVinciVariableResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciVariableResponseWithDefaults

`func NewDaVinciVariableResponseWithDefaults() *DaVinciVariableResponse`

NewDaVinciVariableResponseWithDefaults instantiates a new DaVinciVariableResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciVariableResponse) GetLinks() DaVinciVariableResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciVariableResponse) GetLinksOk() (*DaVinciVariableResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciVariableResponse) SetLinks(v DaVinciVariableResponseLinks)`

SetLinks sets Links field to given value.


### GetDataType

`func (o *DaVinciVariableResponse) GetDataType() DaVinciVariableResponseDataType`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *DaVinciVariableResponse) GetDataTypeOk() (*DaVinciVariableResponseDataType, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *DaVinciVariableResponse) SetDataType(v DaVinciVariableResponseDataType)`

SetDataType sets DataType field to given value.


### GetEnvironment

`func (o *DaVinciVariableResponse) GetEnvironment() ResourceRelationshipReadOnly`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciVariableResponse) GetEnvironmentOk() (*ResourceRelationshipReadOnly, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciVariableResponse) SetEnvironment(v ResourceRelationshipReadOnly)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *DaVinciVariableResponse) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciVariableResponse) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciVariableResponse) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetName

`func (o *DaVinciVariableResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciVariableResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciVariableResponse) SetName(v string)`

SetName sets Name field to given value.


### GetContext

`func (o *DaVinciVariableResponse) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DaVinciVariableResponse) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DaVinciVariableResponse) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *DaVinciVariableResponse) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciVariableResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciVariableResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciVariableResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciVariableResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisplayName

`func (o *DaVinciVariableResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DaVinciVariableResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DaVinciVariableResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DaVinciVariableResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetFlow

`func (o *DaVinciVariableResponse) GetFlow() ResourceRelationshipDaVinciReadOnly`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciVariableResponse) GetFlowOk() (*ResourceRelationshipDaVinciReadOnly, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciVariableResponse) SetFlow(v ResourceRelationshipDaVinciReadOnly)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *DaVinciVariableResponse) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetMax

`func (o *DaVinciVariableResponse) GetMax() float32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DaVinciVariableResponse) GetMaxOk() (*float32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DaVinciVariableResponse) SetMax(v float32)`

SetMax sets Max field to given value.

### HasMax

`func (o *DaVinciVariableResponse) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *DaVinciVariableResponse) GetMin() float32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *DaVinciVariableResponse) GetMinOk() (*float32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *DaVinciVariableResponse) SetMin(v float32)`

SetMin sets Min field to given value.

### HasMin

`func (o *DaVinciVariableResponse) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetMutable

`func (o *DaVinciVariableResponse) GetMutable() bool`

GetMutable returns the Mutable field if non-nil, zero value otherwise.

### GetMutableOk

`func (o *DaVinciVariableResponse) GetMutableOk() (*bool, bool)`

GetMutableOk returns a tuple with the Mutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutable

`func (o *DaVinciVariableResponse) SetMutable(v bool)`

SetMutable sets Mutable field to given value.

### HasMutable

`func (o *DaVinciVariableResponse) HasMutable() bool`

HasMutable returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciVariableResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciVariableResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciVariableResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciVariableResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *DaVinciVariableResponse) GetValue() DaVinciVariableResponseValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DaVinciVariableResponse) GetValueOk() (*DaVinciVariableResponseValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DaVinciVariableResponse) SetValue(v DaVinciVariableResponseValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *DaVinciVariableResponse) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


