# DaVinciConnectorInstanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciConnectorInstanceResponseLinks**](DaVinciConnectorInstanceResponseLinks.md) |  | 
**Connector** | [**ResourceRelationshipDaVinciReadOnly**](ResourceRelationshipDaVinciReadOnly.md) |  | 
**Environment** | [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDaVinciConnectorInstanceResponse

`func NewDaVinciConnectorInstanceResponse(links DaVinciConnectorInstanceResponseLinks, connector ResourceRelationshipDaVinciReadOnly, environment ResourceRelationshipReadOnly, id string, name string, ) *DaVinciConnectorInstanceResponse`

NewDaVinciConnectorInstanceResponse instantiates a new DaVinciConnectorInstanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciConnectorInstanceResponseWithDefaults

`func NewDaVinciConnectorInstanceResponseWithDefaults() *DaVinciConnectorInstanceResponse`

NewDaVinciConnectorInstanceResponseWithDefaults instantiates a new DaVinciConnectorInstanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciConnectorInstanceResponse) GetLinks() DaVinciConnectorInstanceResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciConnectorInstanceResponse) GetLinksOk() (*DaVinciConnectorInstanceResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciConnectorInstanceResponse) SetLinks(v DaVinciConnectorInstanceResponseLinks)`

SetLinks sets Links field to given value.


### GetConnector

`func (o *DaVinciConnectorInstanceResponse) GetConnector() ResourceRelationshipDaVinciReadOnly`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DaVinciConnectorInstanceResponse) GetConnectorOk() (*ResourceRelationshipDaVinciReadOnly, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DaVinciConnectorInstanceResponse) SetConnector(v ResourceRelationshipDaVinciReadOnly)`

SetConnector sets Connector field to given value.


### GetEnvironment

`func (o *DaVinciConnectorInstanceResponse) GetEnvironment() ResourceRelationshipReadOnly`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciConnectorInstanceResponse) GetEnvironmentOk() (*ResourceRelationshipReadOnly, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciConnectorInstanceResponse) SetEnvironment(v ResourceRelationshipReadOnly)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *DaVinciConnectorInstanceResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciConnectorInstanceResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciConnectorInstanceResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DaVinciConnectorInstanceResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciConnectorInstanceResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciConnectorInstanceResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *DaVinciConnectorInstanceResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciConnectorInstanceResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciConnectorInstanceResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciConnectorInstanceResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProperties

`func (o *DaVinciConnectorInstanceResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DaVinciConnectorInstanceResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DaVinciConnectorInstanceResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DaVinciConnectorInstanceResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciConnectorInstanceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciConnectorInstanceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciConnectorInstanceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciConnectorInstanceResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


