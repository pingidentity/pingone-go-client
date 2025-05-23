# DavinciConnectorInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DavinciConnectorInstanceLinks**](DavinciConnectorInstanceLinks.md) |  | 
**Connector** | [**ResourceRelationshipDaVinci**](ResourceRelationshipDaVinci.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDavinciConnectorInstance

`func NewDavinciConnectorInstance(links DavinciConnectorInstanceLinks, connector ResourceRelationshipDaVinci, environment ResourceRelationshipPingOne, id string, name string, ) *DavinciConnectorInstance`

NewDavinciConnectorInstance instantiates a new DavinciConnectorInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDavinciConnectorInstanceWithDefaults

`func NewDavinciConnectorInstanceWithDefaults() *DavinciConnectorInstance`

NewDavinciConnectorInstanceWithDefaults instantiates a new DavinciConnectorInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DavinciConnectorInstance) GetLinks() DavinciConnectorInstanceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DavinciConnectorInstance) GetLinksOk() (*DavinciConnectorInstanceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DavinciConnectorInstance) SetLinks(v DavinciConnectorInstanceLinks)`

SetLinks sets Links field to given value.


### GetConnector

`func (o *DavinciConnectorInstance) GetConnector() ResourceRelationshipDaVinci`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DavinciConnectorInstance) GetConnectorOk() (*ResourceRelationshipDaVinci, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DavinciConnectorInstance) SetConnector(v ResourceRelationshipDaVinci)`

SetConnector sets Connector field to given value.


### GetCreatedAt

`func (o *DavinciConnectorInstance) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DavinciConnectorInstance) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DavinciConnectorInstance) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DavinciConnectorInstance) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *DavinciConnectorInstance) GetEnvironment() ResourceRelationshipPingOne`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DavinciConnectorInstance) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DavinciConnectorInstance) SetEnvironment(v ResourceRelationshipPingOne)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *DavinciConnectorInstance) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DavinciConnectorInstance) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DavinciConnectorInstance) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DavinciConnectorInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DavinciConnectorInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DavinciConnectorInstance) SetName(v string)`

SetName sets Name field to given value.


### GetProperties

`func (o *DavinciConnectorInstance) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DavinciConnectorInstance) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DavinciConnectorInstance) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DavinciConnectorInstance) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DavinciConnectorInstance) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DavinciConnectorInstance) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DavinciConnectorInstance) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DavinciConnectorInstance) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


