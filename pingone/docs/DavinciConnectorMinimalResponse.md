# DavinciConnectorMinimalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DavinciConnectorCollectionMinimalResponseLinks**](DavinciConnectorCollectionMinimalResponseLinks.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Environment** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**Id** | **string** |  | 
**Metadata** | [**DavinciConnectorMinimalResponseMetadata**](DavinciConnectorMinimalResponseMetadata.md) |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Version** | **string** |  | 

## Methods

### NewDavinciConnectorMinimalResponse

`func NewDavinciConnectorMinimalResponse(links DavinciConnectorCollectionMinimalResponseLinks, createdAt time.Time, description string, environment ResourceRelationshipPingOne, id string, metadata DavinciConnectorMinimalResponseMetadata, name string, updatedAt time.Time, version string, ) *DavinciConnectorMinimalResponse`

NewDavinciConnectorMinimalResponse instantiates a new DavinciConnectorMinimalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDavinciConnectorMinimalResponseWithDefaults

`func NewDavinciConnectorMinimalResponseWithDefaults() *DavinciConnectorMinimalResponse`

NewDavinciConnectorMinimalResponseWithDefaults instantiates a new DavinciConnectorMinimalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DavinciConnectorMinimalResponse) GetLinks() DavinciConnectorCollectionMinimalResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DavinciConnectorMinimalResponse) GetLinksOk() (*DavinciConnectorCollectionMinimalResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DavinciConnectorMinimalResponse) SetLinks(v DavinciConnectorCollectionMinimalResponseLinks)`

SetLinks sets Links field to given value.


### GetCreatedAt

`func (o *DavinciConnectorMinimalResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DavinciConnectorMinimalResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DavinciConnectorMinimalResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *DavinciConnectorMinimalResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DavinciConnectorMinimalResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DavinciConnectorMinimalResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnvironment

`func (o *DavinciConnectorMinimalResponse) GetEnvironment() ResourceRelationshipPingOne`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DavinciConnectorMinimalResponse) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DavinciConnectorMinimalResponse) SetEnvironment(v ResourceRelationshipPingOne)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *DavinciConnectorMinimalResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DavinciConnectorMinimalResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DavinciConnectorMinimalResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *DavinciConnectorMinimalResponse) GetMetadata() DavinciConnectorMinimalResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DavinciConnectorMinimalResponse) GetMetadataOk() (*DavinciConnectorMinimalResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DavinciConnectorMinimalResponse) SetMetadata(v DavinciConnectorMinimalResponseMetadata)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *DavinciConnectorMinimalResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DavinciConnectorMinimalResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DavinciConnectorMinimalResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *DavinciConnectorMinimalResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DavinciConnectorMinimalResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DavinciConnectorMinimalResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *DavinciConnectorMinimalResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DavinciConnectorMinimalResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DavinciConnectorMinimalResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


