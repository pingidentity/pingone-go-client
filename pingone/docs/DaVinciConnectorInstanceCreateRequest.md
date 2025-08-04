# DaVinciConnectorInstanceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Connector** | [**ResourceRelationshipDaVinci**](ResourceRelationshipDaVinci.md) |  | 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDaVinciConnectorInstanceCreateRequest

`func NewDaVinciConnectorInstanceCreateRequest(name string, connector ResourceRelationshipDaVinci, ) *DaVinciConnectorInstanceCreateRequest`

NewDaVinciConnectorInstanceCreateRequest instantiates a new DaVinciConnectorInstanceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciConnectorInstanceCreateRequestWithDefaults

`func NewDaVinciConnectorInstanceCreateRequestWithDefaults() *DaVinciConnectorInstanceCreateRequest`

NewDaVinciConnectorInstanceCreateRequestWithDefaults instantiates a new DaVinciConnectorInstanceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DaVinciConnectorInstanceCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciConnectorInstanceCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciConnectorInstanceCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConnector

`func (o *DaVinciConnectorInstanceCreateRequest) GetConnector() ResourceRelationshipDaVinci`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DaVinciConnectorInstanceCreateRequest) GetConnectorOk() (*ResourceRelationshipDaVinci, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DaVinciConnectorInstanceCreateRequest) SetConnector(v ResourceRelationshipDaVinci)`

SetConnector sets Connector field to given value.


### GetProperties

`func (o *DaVinciConnectorInstanceCreateRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DaVinciConnectorInstanceCreateRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DaVinciConnectorInstanceCreateRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DaVinciConnectorInstanceCreateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


