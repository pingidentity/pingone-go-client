# DaVinciConnectorInstanceResponseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**JSONHALLink**](JSONHALLink.md) |  | 
**Self** | [**JSONHALLink**](JSONHALLink.md) |  | 
**ConnectorInstanceClone** | [**JSONHALLink**](JSONHALLink.md) |  | 
**Applications** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 
**DeviceAuthenticationPolicies** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 
**Gateways** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 
**NotificationsPolicies** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 

## Methods

### NewDaVinciConnectorInstanceResponseLinks

`func NewDaVinciConnectorInstanceResponseLinks(environment JSONHALLink, self JSONHALLink, connectorInstanceClone JSONHALLink, ) *DaVinciConnectorInstanceResponseLinks`

NewDaVinciConnectorInstanceResponseLinks instantiates a new DaVinciConnectorInstanceResponseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciConnectorInstanceResponseLinksWithDefaults

`func NewDaVinciConnectorInstanceResponseLinksWithDefaults() *DaVinciConnectorInstanceResponseLinks`

NewDaVinciConnectorInstanceResponseLinksWithDefaults instantiates a new DaVinciConnectorInstanceResponseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *DaVinciConnectorInstanceResponseLinks) GetEnvironment() JSONHALLink`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciConnectorInstanceResponseLinks) GetEnvironmentOk() (*JSONHALLink, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciConnectorInstanceResponseLinks) SetEnvironment(v JSONHALLink)`

SetEnvironment sets Environment field to given value.


### GetSelf

`func (o *DaVinciConnectorInstanceResponseLinks) GetSelf() JSONHALLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *DaVinciConnectorInstanceResponseLinks) GetSelfOk() (*JSONHALLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *DaVinciConnectorInstanceResponseLinks) SetSelf(v JSONHALLink)`

SetSelf sets Self field to given value.


### GetConnectorInstanceClone

`func (o *DaVinciConnectorInstanceResponseLinks) GetConnectorInstanceClone() JSONHALLink`

GetConnectorInstanceClone returns the ConnectorInstanceClone field if non-nil, zero value otherwise.

### GetConnectorInstanceCloneOk

`func (o *DaVinciConnectorInstanceResponseLinks) GetConnectorInstanceCloneOk() (*JSONHALLink, bool)`

GetConnectorInstanceCloneOk returns a tuple with the ConnectorInstanceClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorInstanceClone

`func (o *DaVinciConnectorInstanceResponseLinks) SetConnectorInstanceClone(v JSONHALLink)`

SetConnectorInstanceClone sets ConnectorInstanceClone field to given value.


### GetApplications

`func (o *DaVinciConnectorInstanceResponseLinks) GetApplications() JSONHALLink`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DaVinciConnectorInstanceResponseLinks) GetApplicationsOk() (*JSONHALLink, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DaVinciConnectorInstanceResponseLinks) SetApplications(v JSONHALLink)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DaVinciConnectorInstanceResponseLinks) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetDeviceAuthenticationPolicies

`func (o *DaVinciConnectorInstanceResponseLinks) GetDeviceAuthenticationPolicies() JSONHALLink`

GetDeviceAuthenticationPolicies returns the DeviceAuthenticationPolicies field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPoliciesOk

`func (o *DaVinciConnectorInstanceResponseLinks) GetDeviceAuthenticationPoliciesOk() (*JSONHALLink, bool)`

GetDeviceAuthenticationPoliciesOk returns a tuple with the DeviceAuthenticationPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicies

`func (o *DaVinciConnectorInstanceResponseLinks) SetDeviceAuthenticationPolicies(v JSONHALLink)`

SetDeviceAuthenticationPolicies sets DeviceAuthenticationPolicies field to given value.

### HasDeviceAuthenticationPolicies

`func (o *DaVinciConnectorInstanceResponseLinks) HasDeviceAuthenticationPolicies() bool`

HasDeviceAuthenticationPolicies returns a boolean if a field has been set.

### GetGateways

`func (o *DaVinciConnectorInstanceResponseLinks) GetGateways() JSONHALLink`

GetGateways returns the Gateways field if non-nil, zero value otherwise.

### GetGatewaysOk

`func (o *DaVinciConnectorInstanceResponseLinks) GetGatewaysOk() (*JSONHALLink, bool)`

GetGatewaysOk returns a tuple with the Gateways field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateways

`func (o *DaVinciConnectorInstanceResponseLinks) SetGateways(v JSONHALLink)`

SetGateways sets Gateways field to given value.

### HasGateways

`func (o *DaVinciConnectorInstanceResponseLinks) HasGateways() bool`

HasGateways returns a boolean if a field has been set.

### GetNotificationsPolicies

`func (o *DaVinciConnectorInstanceResponseLinks) GetNotificationsPolicies() JSONHALLink`

GetNotificationsPolicies returns the NotificationsPolicies field if non-nil, zero value otherwise.

### GetNotificationsPoliciesOk

`func (o *DaVinciConnectorInstanceResponseLinks) GetNotificationsPoliciesOk() (*JSONHALLink, bool)`

GetNotificationsPoliciesOk returns a tuple with the NotificationsPolicies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationsPolicies

`func (o *DaVinciConnectorInstanceResponseLinks) SetNotificationsPolicies(v JSONHALLink)`

SetNotificationsPolicies sets NotificationsPolicies field to given value.

### HasNotificationsPolicies

`func (o *DaVinciConnectorInstanceResponseLinks) HasNotificationsPolicies() bool`

HasNotificationsPolicies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


