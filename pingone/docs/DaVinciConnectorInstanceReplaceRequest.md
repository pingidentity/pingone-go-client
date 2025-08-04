# DaVinciConnectorInstanceReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Connector** | [**ResourceRelationshipDaVinci**](ResourceRelationshipDaVinci.md) |  | 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDaVinciConnectorInstanceReplaceRequest

`func NewDaVinciConnectorInstanceReplaceRequest(name string, connector ResourceRelationshipDaVinci, ) *DaVinciConnectorInstanceReplaceRequest`

NewDaVinciConnectorInstanceReplaceRequest instantiates a new DaVinciConnectorInstanceReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciConnectorInstanceReplaceRequestWithDefaults

`func NewDaVinciConnectorInstanceReplaceRequestWithDefaults() *DaVinciConnectorInstanceReplaceRequest`

NewDaVinciConnectorInstanceReplaceRequestWithDefaults instantiates a new DaVinciConnectorInstanceReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DaVinciConnectorInstanceReplaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciConnectorInstanceReplaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciConnectorInstanceReplaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetConnector

`func (o *DaVinciConnectorInstanceReplaceRequest) GetConnector() ResourceRelationshipDaVinci`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DaVinciConnectorInstanceReplaceRequest) GetConnectorOk() (*ResourceRelationshipDaVinci, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DaVinciConnectorInstanceReplaceRequest) SetConnector(v ResourceRelationshipDaVinci)`

SetConnector sets Connector field to given value.


### GetProperties

`func (o *DaVinciConnectorInstanceReplaceRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DaVinciConnectorInstanceReplaceRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DaVinciConnectorInstanceReplaceRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DaVinciConnectorInstanceReplaceRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


