# DaVinciConnectorMinimalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciConnectorMinimalResponseLinks**](DaVinciConnectorMinimalResponseLinks.md) |  | 
**CreatedAt** | **time.Time** |  | 
**Description** | **string** |  | 
**Environment** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**Id** | **string** |  | 
**Metadata** | [**DaVinciConnectorMinimalResponseMetadata**](DaVinciConnectorMinimalResponseMetadata.md) |  | 
**Name** | **string** |  | 
**UpdatedAt** | **time.Time** |  | 
**Version** | **string** |  | 

## Methods

### NewDaVinciConnectorMinimalResponse

`func NewDaVinciConnectorMinimalResponse(links DaVinciConnectorMinimalResponseLinks, createdAt time.Time, description string, environment ResourceRelationshipPingOne, id string, metadata DaVinciConnectorMinimalResponseMetadata, name string, updatedAt time.Time, version string, ) *DaVinciConnectorMinimalResponse`

NewDaVinciConnectorMinimalResponse instantiates a new DaVinciConnectorMinimalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciConnectorMinimalResponseWithDefaults

`func NewDaVinciConnectorMinimalResponseWithDefaults() *DaVinciConnectorMinimalResponse`

NewDaVinciConnectorMinimalResponseWithDefaults instantiates a new DaVinciConnectorMinimalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciConnectorMinimalResponse) GetLinks() DaVinciConnectorMinimalResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciConnectorMinimalResponse) GetLinksOk() (*DaVinciConnectorMinimalResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciConnectorMinimalResponse) SetLinks(v DaVinciConnectorMinimalResponseLinks)`

SetLinks sets Links field to given value.


### GetCreatedAt

`func (o *DaVinciConnectorMinimalResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciConnectorMinimalResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciConnectorMinimalResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *DaVinciConnectorMinimalResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DaVinciConnectorMinimalResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DaVinciConnectorMinimalResponse) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnvironment

`func (o *DaVinciConnectorMinimalResponse) GetEnvironment() ResourceRelationshipPingOne`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciConnectorMinimalResponse) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciConnectorMinimalResponse) SetEnvironment(v ResourceRelationshipPingOne)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *DaVinciConnectorMinimalResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciConnectorMinimalResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciConnectorMinimalResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMetadata

`func (o *DaVinciConnectorMinimalResponse) GetMetadata() DaVinciConnectorMinimalResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DaVinciConnectorMinimalResponse) GetMetadataOk() (*DaVinciConnectorMinimalResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DaVinciConnectorMinimalResponse) SetMetadata(v DaVinciConnectorMinimalResponseMetadata)`

SetMetadata sets Metadata field to given value.


### GetName

`func (o *DaVinciConnectorMinimalResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciConnectorMinimalResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciConnectorMinimalResponse) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *DaVinciConnectorMinimalResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciConnectorMinimalResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciConnectorMinimalResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVersion

`func (o *DaVinciConnectorMinimalResponse) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciConnectorMinimalResponse) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciConnectorMinimalResponse) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


