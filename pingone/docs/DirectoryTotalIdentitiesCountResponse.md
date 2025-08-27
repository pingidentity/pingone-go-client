# DirectoryTotalIdentitiesCountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **time.Time** |  | [optional] 
**TotalIdentities** | Pointer to **int32** |  | [optional] 

## Methods

### NewDirectoryTotalIdentitiesCountResponse

`func NewDirectoryTotalIdentitiesCountResponse() *DirectoryTotalIdentitiesCountResponse`

NewDirectoryTotalIdentitiesCountResponse instantiates a new DirectoryTotalIdentitiesCountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDirectoryTotalIdentitiesCountResponseWithDefaults

`func NewDirectoryTotalIdentitiesCountResponseWithDefaults() *DirectoryTotalIdentitiesCountResponse`

NewDirectoryTotalIdentitiesCountResponseWithDefaults instantiates a new DirectoryTotalIdentitiesCountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DirectoryTotalIdentitiesCountResponse) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DirectoryTotalIdentitiesCountResponse) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DirectoryTotalIdentitiesCountResponse) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *DirectoryTotalIdentitiesCountResponse) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetTotalIdentities

`func (o *DirectoryTotalIdentitiesCountResponse) GetTotalIdentities() int32`

GetTotalIdentities returns the TotalIdentities field if non-nil, zero value otherwise.

### GetTotalIdentitiesOk

`func (o *DirectoryTotalIdentitiesCountResponse) GetTotalIdentitiesOk() (*int32, bool)`

GetTotalIdentitiesOk returns a tuple with the TotalIdentities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalIdentities

`func (o *DirectoryTotalIdentitiesCountResponse) SetTotalIdentities(v int32)`

SetTotalIdentities sets TotalIdentities field to given value.

### HasTotalIdentities

`func (o *DirectoryTotalIdentitiesCountResponse) HasTotalIdentities() bool`

HasTotalIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


