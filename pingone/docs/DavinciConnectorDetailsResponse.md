# DaVinciConnectorDetailsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**DaVinciConnectorDetailsResponseLinks**](DaVinciConnectorDetailsResponseLinks.md) |  | [optional] 
**AccountConfigView** | Pointer to [**DaVinciConnectorDetailsResponseAccountConfigView**](DaVinciConnectorDetailsResponseAccountConfigView.md) |  | [optional] 
**Capabilities** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CredentialsView** | Pointer to [**DaVinciConnectorDetailsResponseCredentialsView**](DaVinciConnectorDetailsResponseCredentialsView.md) |  | [optional] 
**Environment** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**FlowSections** | Pointer to [**[]DaVinciConnectorDetailsResponseFlowSection**](DaVinciConnectorDetailsResponseFlowSection.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Sections** | Pointer to [**[]DaVinciConnectorDetailsResponseSection**](DaVinciConnectorDetailsResponseSection.md) |  | [optional] 

## Methods

### NewDaVinciConnectorDetailsResponse

`func NewDaVinciConnectorDetailsResponse(environment ResourceRelationshipPingOne, ) *DaVinciConnectorDetailsResponse`

NewDaVinciConnectorDetailsResponse instantiates a new DaVinciConnectorDetailsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciConnectorDetailsResponseWithDefaults

`func NewDaVinciConnectorDetailsResponseWithDefaults() *DaVinciConnectorDetailsResponse`

NewDaVinciConnectorDetailsResponseWithDefaults instantiates a new DaVinciConnectorDetailsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciConnectorDetailsResponse) GetLinks() DaVinciConnectorDetailsResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciConnectorDetailsResponse) GetLinksOk() (*DaVinciConnectorDetailsResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciConnectorDetailsResponse) SetLinks(v DaVinciConnectorDetailsResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DaVinciConnectorDetailsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAccountConfigView

`func (o *DaVinciConnectorDetailsResponse) GetAccountConfigView() DaVinciConnectorDetailsResponseAccountConfigView`

GetAccountConfigView returns the AccountConfigView field if non-nil, zero value otherwise.

### GetAccountConfigViewOk

`func (o *DaVinciConnectorDetailsResponse) GetAccountConfigViewOk() (*DaVinciConnectorDetailsResponseAccountConfigView, bool)`

GetAccountConfigViewOk returns a tuple with the AccountConfigView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountConfigView

`func (o *DaVinciConnectorDetailsResponse) SetAccountConfigView(v DaVinciConnectorDetailsResponseAccountConfigView)`

SetAccountConfigView sets AccountConfigView field to given value.

### HasAccountConfigView

`func (o *DaVinciConnectorDetailsResponse) HasAccountConfigView() bool`

HasAccountConfigView returns a boolean if a field has been set.

### GetCapabilities

`func (o *DaVinciConnectorDetailsResponse) GetCapabilities() map[string]map[string]interface{}`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DaVinciConnectorDetailsResponse) GetCapabilitiesOk() (*map[string]map[string]interface{}, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DaVinciConnectorDetailsResponse) SetCapabilities(v map[string]map[string]interface{})`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DaVinciConnectorDetailsResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetCredentialsView

`func (o *DaVinciConnectorDetailsResponse) GetCredentialsView() DaVinciConnectorDetailsResponseCredentialsView`

GetCredentialsView returns the CredentialsView field if non-nil, zero value otherwise.

### GetCredentialsViewOk

`func (o *DaVinciConnectorDetailsResponse) GetCredentialsViewOk() (*DaVinciConnectorDetailsResponseCredentialsView, bool)`

GetCredentialsViewOk returns a tuple with the CredentialsView field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialsView

`func (o *DaVinciConnectorDetailsResponse) SetCredentialsView(v DaVinciConnectorDetailsResponseCredentialsView)`

SetCredentialsView sets CredentialsView field to given value.

### HasCredentialsView

`func (o *DaVinciConnectorDetailsResponse) HasCredentialsView() bool`

HasCredentialsView returns a boolean if a field has been set.

### GetEnvironment

`func (o *DaVinciConnectorDetailsResponse) GetEnvironment() ResourceRelationshipPingOne`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciConnectorDetailsResponse) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciConnectorDetailsResponse) SetEnvironment(v ResourceRelationshipPingOne)`

SetEnvironment sets Environment field to given value.


### GetFlowSections

`func (o *DaVinciConnectorDetailsResponse) GetFlowSections() []DaVinciConnectorDetailsResponseFlowSection`

GetFlowSections returns the FlowSections field if non-nil, zero value otherwise.

### GetFlowSectionsOk

`func (o *DaVinciConnectorDetailsResponse) GetFlowSectionsOk() (*[]DaVinciConnectorDetailsResponseFlowSection, bool)`

GetFlowSectionsOk returns a tuple with the FlowSections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowSections

`func (o *DaVinciConnectorDetailsResponse) SetFlowSections(v []DaVinciConnectorDetailsResponseFlowSection)`

SetFlowSections sets FlowSections field to given value.

### HasFlowSections

`func (o *DaVinciConnectorDetailsResponse) HasFlowSections() bool`

HasFlowSections returns a boolean if a field has been set.

### GetProperties

`func (o *DaVinciConnectorDetailsResponse) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DaVinciConnectorDetailsResponse) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DaVinciConnectorDetailsResponse) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DaVinciConnectorDetailsResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSections

`func (o *DaVinciConnectorDetailsResponse) GetSections() []DaVinciConnectorDetailsResponseSection`

GetSections returns the Sections field if non-nil, zero value otherwise.

### GetSectionsOk

`func (o *DaVinciConnectorDetailsResponse) GetSectionsOk() (*[]DaVinciConnectorDetailsResponseSection, bool)`

GetSectionsOk returns a tuple with the Sections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSections

`func (o *DaVinciConnectorDetailsResponse) SetSections(v []DaVinciConnectorDetailsResponseSection)`

SetSections sets Sections field to given value.

### HasSections

`func (o *DaVinciConnectorDetailsResponse) HasSections() bool`

HasSections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


