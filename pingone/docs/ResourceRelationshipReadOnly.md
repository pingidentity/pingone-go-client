# ResourceRelationshipReadOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**uuid.UUID**](uuid.UUID.md) |  | [readonly] 

## Methods

### NewResourceRelationshipReadOnly

`func NewResourceRelationshipReadOnly(id uuid.UUID, ) *ResourceRelationshipReadOnly`

NewResourceRelationshipReadOnly instantiates a new ResourceRelationshipReadOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceRelationshipReadOnlyWithDefaults

`func NewResourceRelationshipReadOnlyWithDefaults() *ResourceRelationshipReadOnly`

NewResourceRelationshipReadOnlyWithDefaults instantiates a new ResourceRelationshipReadOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceRelationshipReadOnly) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceRelationshipReadOnly) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceRelationshipReadOnly) SetId(v uuid.UUID)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


