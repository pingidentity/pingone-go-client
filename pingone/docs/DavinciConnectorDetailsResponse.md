# DavinciConnectorDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**DavinciConnectorCollectionMinimalResponseLinks**](DavinciConnectorCollectionMinimalResponseLinks.md) |  | [optional] 
**AccountConfigView** | Pointer to [**DavinciConnectorDetailsResponseAccountConfigView**](DavinciConnectorDetailsResponseAccountConfigView.md) |  | [optional] 
**Capabilities** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CredentialsView** | Pointer to [**DavinciConnectorDetailsResponseCredentialsView**](DavinciConnectorDetailsResponseCredentialsView.md) |  | [optional] 
**Environment** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**FlowSections** | Pointer to [**[]DavinciConnectorDetailsResponseFlowSections**](DavinciConnectorDetailsResponseFlowSections.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Sections** | Pointer to [**[]DavinciConnectorDetailsResponseSections**](DavinciConnectorDetailsResponseSections.md) |  | [optional] 

## Methods

### NewDavinciConnectorDetailsResponse

`func NewDavinciConnectorDetailsResponse(environment ResourceRelationshipPingOne, ) *DavinciConnectorDetailsResponse`

NewDavinciConnectorDetailsResponse instantiates a new DavinciConnectorDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDavinciConnectorDetailsResponseWithDefaults

`func NewDavinciConnectorDetailsResponseWithDefaults() *DavinciConnectorDetailsResponse`

NewDavinciConnectorDetailsResponseWithDefaults instantiates a new DavinciConnectorDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DavinciConnectorDetailsResponse) GetLinks() DavinciConnectorCollectionMinimalResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DavinciConnectorDetailsResponse) GetLinksOk() (*DavinciConnectorCollectionMinimalResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DavinciConnectorDetailsResponse) SetLinks(v DavinciConnectorCollectionMinimalResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DavinciConnectorDetailsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccountConfigView

`func (o *DavinciConnectorDetailsResponse) GetAccountConfigView() DavinciConnectorDetailsResponseAccountConfigView`

GetAccountConfigView returns the AccountConfigView field if non-nil, zero value otherwise.

### GetAccountConfigViewOk

`func (o *DavinciConnectorDetailsResponse) GetAccountConfigViewOk() (*DavinciConnectorDetailsResponseAccountConfigView, bool)`

GetAccountConfigViewOk returns a tuple with the AccountConfigView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigView

`func (o *DavinciConnectorDetailsResponse) SetAccountConfigView(v DavinciConnectorDetailsResponseAccountConfigView)`

SetAccountConfigView sets AccountConfigView field to given value.

### HasAccountConfigView

`func (o *DavinciConnectorDetailsResponse) HasAccountConfigView() bool`

HasAccountConfigView returns a boolean if a field has been set.

### GetCapabilities

`func (o *DavinciConnectorDetailsResponse) GetCapabilities() map[string]map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DavinciConnectorDetailsResponse) GetCapabilitiesOk() (*map[string]map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DavinciConnectorDetailsResponse) SetCapabilities(v map[string]map[string]interface{})`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DavinciConnectorDetailsResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCredentialsView

`func (o *DavinciConnectorDetailsResponse) GetCredentialsView() DavinciConnectorDetailsResponseCredentialsView`

GetCredentialsView returns the CredentialsView field if non-nil, zero value otherwise.

### GetCredentialsViewOk

`func (o *DavinciConnectorDetailsResponse) GetCredentialsViewOk() (*DavinciConnectorDetailsResponseCredentialsView, bool)`

GetCredentialsViewOk returns a tuple with the CredentialsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsView

`func (o *DavinciConnectorDetailsResponse) SetCredentialsView(v DavinciConnectorDetailsResponseCredentialsView)`

SetCredentialsView sets CredentialsView field to given value.

### HasCredentialsView

`func (o *DavinciConnectorDetailsResponse) HasCredentialsView() bool`

HasCredentialsView returns a boolean if a field has been set.

### GetEnvironment

`func (o *DavinciConnectorDetailsResponse) GetEnvironment() ResourceRelationshipPingOne`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DavinciConnectorDetailsResponse) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DavinciConnectorDetailsResponse) SetEnvironment(v ResourceRelationshipPingOne)`

SetEnvironment sets Environment field to given value.


### GetFlowSections

`func (o *DavinciConnectorDetailsResponse) GetFlowSections() []DavinciConnectorDetailsResponseFlowSections`

GetFlowSections returns the FlowSections field if non-nil, zero value otherwise.

### GetFlowSectionsOk

`func (o *DavinciConnectorDetailsResponse) GetFlowSectionsOk() (*[]DavinciConnectorDetailsResponseFlowSections, bool)`

GetFlowSectionsOk returns a tuple with the FlowSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSections

`func (o *DavinciConnectorDetailsResponse) SetFlowSections(v []DavinciConnectorDetailsResponseFlowSections)`

SetFlowSections sets FlowSections field to given value.

### HasFlowSections

`func (o *DavinciConnectorDetailsResponse) HasFlowSections() bool`

HasFlowSections returns a boolean if a field has been set.

### GetProperties

`func (o *DavinciConnectorDetailsResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DavinciConnectorDetailsResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DavinciConnectorDetailsResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DavinciConnectorDetailsResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSections

`func (o *DavinciConnectorDetailsResponse) GetSections() []DavinciConnectorDetailsResponseSections`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *DavinciConnectorDetailsResponse) GetSectionsOk() (*[]DavinciConnectorDetailsResponseSections, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *DavinciConnectorDetailsResponse) SetSections(v []DavinciConnectorDetailsResponseSections)`

SetSections sets Sections field to given value.

### HasSections

`func (o *DavinciConnectorDetailsResponse) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


