# DaVinciConnectorInstanceLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applications** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 
**ConnectorInstanceClone** | [**JSONHALLink**](JSONHALLink.md) |  | 
**DeviceAuthenticationPolicies** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 
**Environment** | [**JSONHALLink**](JSONHALLink.md) |  | 
**Gateways** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 
**NotificationsPolicies** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 
**Self** | [**JSONHALLink**](JSONHALLink.md) |  | 

## Methods

### NewDaVinciConnectorInstanceLinks

`func NewDaVinciConnectorInstanceLinks(connectorInstanceClone JSONHALLink, environment JSONHALLink, self JSONHALLink, ) *DaVinciConnectorInstanceLinks`

NewDaVinciConnectorInstanceLinks instantiates a new DaVinciConnectorInstanceLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciConnectorInstanceLinksWithDefaults

`func NewDaVinciConnectorInstanceLinksWithDefaults() *DaVinciConnectorInstanceLinks`

NewDaVinciConnectorInstanceLinksWithDefaults instantiates a new DaVinciConnectorInstanceLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplications

`func (o *DaVinciConnectorInstanceLinks) GetApplications() JSONHALLink`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DaVinciConnectorInstanceLinks) GetApplicationsOk() (*JSONHALLink, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DaVinciConnectorInstanceLinks) SetApplications(v JSONHALLink)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DaVinciConnectorInstanceLinks) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetConnectorInstanceClone

`func (o *DaVinciConnectorInstanceLinks) GetConnectorInstanceClone() JSONHALLink`

GetConnectorInstanceClone returns the ConnectorInstanceClone field if non-nil, zero value otherwise.

### GetConnectorInstanceCloneOk

`func (o *DaVinciConnectorInstanceLinks) GetConnectorInstanceCloneOk() (*JSONHALLink, bool)`

GetConnectorInstanceCloneOk returns a tuple with the ConnectorInstanceClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorInstanceClone

`func (o *DaVinciConnectorInstanceLinks) SetConnectorInstanceClone(v JSONHALLink)`

SetConnectorInstanceClone sets ConnectorInstanceClone field to given value.


### GetDeviceAuthenticationPolicies

`func (o *DaVinciConnectorInstanceLinks) GetDeviceAuthenticationPolicies() JSONHALLink`

GetDeviceAuthenticationPolicies returns the DeviceAuthenticationPolicies field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPoliciesOk

`func (o *DaVinciConnectorInstanceLinks) GetDeviceAuthenticationPoliciesOk() (*JSONHALLink, bool)`

GetDeviceAuthenticationPoliciesOk returns a tuple with the DeviceAuthenticationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicies

`func (o *DaVinciConnectorInstanceLinks) SetDeviceAuthenticationPolicies(v JSONHALLink)`

SetDeviceAuthenticationPolicies sets DeviceAuthenticationPolicies field to given value.

### HasDeviceAuthenticationPolicies

`func (o *DaVinciConnectorInstanceLinks) HasDeviceAuthenticationPolicies() bool`

HasDeviceAuthenticationPolicies returns a boolean if a field has been set.

### GetEnvironment

`func (o *DaVinciConnectorInstanceLinks) GetEnvironment() JSONHALLink`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciConnectorInstanceLinks) GetEnvironmentOk() (*JSONHALLink, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciConnectorInstanceLinks) SetEnvironment(v JSONHALLink)`

SetEnvironment sets Environment field to given value.


### GetGateways

`func (o *DaVinciConnectorInstanceLinks) GetGateways() JSONHALLink`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *DaVinciConnectorInstanceLinks) GetGatewaysOk() (*JSONHALLink, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *DaVinciConnectorInstanceLinks) SetGateways(v JSONHALLink)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *DaVinciConnectorInstanceLinks) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetNotificationsPolicies

`func (o *DaVinciConnectorInstanceLinks) GetNotificationsPolicies() JSONHALLink`

GetNotificationsPolicies returns the NotificationsPolicies field if non-nil, zero value otherwise.

### GetNotificationsPoliciesOk

`func (o *DaVinciConnectorInstanceLinks) GetNotificationsPoliciesOk() (*JSONHALLink, bool)`

GetNotificationsPoliciesOk returns a tuple with the NotificationsPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsPolicies

`func (o *DaVinciConnectorInstanceLinks) SetNotificationsPolicies(v JSONHALLink)`

SetNotificationsPolicies sets NotificationsPolicies field to given value.

### HasNotificationsPolicies

`func (o *DaVinciConnectorInstanceLinks) HasNotificationsPolicies() bool`

HasNotificationsPolicies returns a boolean if a field has been set.

### GetSelf

`func (o *DaVinciConnectorInstanceLinks) GetSelf() JSONHALLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *DaVinciConnectorInstanceLinks) GetSelfOk() (*JSONHALLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *DaVinciConnectorInstanceLinks) SetSelf(v JSONHALLink)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


