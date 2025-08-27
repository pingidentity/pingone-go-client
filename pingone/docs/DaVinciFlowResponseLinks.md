# DaVinciFlowResponseLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**JSONHALLink**](JSONHALLink.md) |  | 
**Self** | [**JSONHALLink**](JSONHALLink.md) |  | 
**ConnectorInstances** | [**JSONHALLink**](JSONHALLink.md) |  | 
**Connectors** | [**JSONHALLink**](JSONHALLink.md) |  | 
**FlowDeploy** | [**JSONHALLink**](JSONHALLink.md) |  | 
**FlowClone** | [**JSONHALLink**](JSONHALLink.md) |  | 
**FlowEnabled** | [**JSONHALLink**](JSONHALLink.md) |  | 
**Version** | [**JSONHALLink**](JSONHALLink.md) |  | 
**FlowValidate** | Pointer to [**JSONHALLink**](JSONHALLink.md) |  | [optional] 

## Methods

### NewDaVinciFlowResponseLinks

`func NewDaVinciFlowResponseLinks(environment JSONHALLink, self JSONHALLink, connectorInstances JSONHALLink, connectors JSONHALLink, flowDeploy JSONHALLink, flowClone JSONHALLink, flowEnabled JSONHALLink, version JSONHALLink, ) *DaVinciFlowResponseLinks`

NewDaVinciFlowResponseLinks instantiates a new DaVinciFlowResponseLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowResponseLinksWithDefaults

`func NewDaVinciFlowResponseLinksWithDefaults() *DaVinciFlowResponseLinks`

NewDaVinciFlowResponseLinksWithDefaults instantiates a new DaVinciFlowResponseLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *DaVinciFlowResponseLinks) GetEnvironment() JSONHALLink`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciFlowResponseLinks) GetEnvironmentOk() (*JSONHALLink, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciFlowResponseLinks) SetEnvironment(v JSONHALLink)`

SetEnvironment sets Environment field to given value.


### GetSelf

`func (o *DaVinciFlowResponseLinks) GetSelf() JSONHALLink`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *DaVinciFlowResponseLinks) GetSelfOk() (*JSONHALLink, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *DaVinciFlowResponseLinks) SetSelf(v JSONHALLink)`

SetSelf sets Self field to given value.


### GetConnectorInstances

`func (o *DaVinciFlowResponseLinks) GetConnectorInstances() JSONHALLink`

GetConnectorInstances returns the ConnectorInstances field if non-nil, zero value otherwise.

### GetConnectorInstancesOk

`func (o *DaVinciFlowResponseLinks) GetConnectorInstancesOk() (*JSONHALLink, bool)`

GetConnectorInstancesOk returns a tuple with the ConnectorInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorInstances

`func (o *DaVinciFlowResponseLinks) SetConnectorInstances(v JSONHALLink)`

SetConnectorInstances sets ConnectorInstances field to given value.


### GetConnectors

`func (o *DaVinciFlowResponseLinks) GetConnectors() JSONHALLink`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DaVinciFlowResponseLinks) GetConnectorsOk() (*JSONHALLink, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DaVinciFlowResponseLinks) SetConnectors(v JSONHALLink)`

SetConnectors sets Connectors field to given value.


### GetFlowDeploy

`func (o *DaVinciFlowResponseLinks) GetFlowDeploy() JSONHALLink`

GetFlowDeploy returns the FlowDeploy field if non-nil, zero value otherwise.

### GetFlowDeployOk

`func (o *DaVinciFlowResponseLinks) GetFlowDeployOk() (*JSONHALLink, bool)`

GetFlowDeployOk returns a tuple with the FlowDeploy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDeploy

`func (o *DaVinciFlowResponseLinks) SetFlowDeploy(v JSONHALLink)`

SetFlowDeploy sets FlowDeploy field to given value.


### GetFlowClone

`func (o *DaVinciFlowResponseLinks) GetFlowClone() JSONHALLink`

GetFlowClone returns the FlowClone field if non-nil, zero value otherwise.

### GetFlowCloneOk

`func (o *DaVinciFlowResponseLinks) GetFlowCloneOk() (*JSONHALLink, bool)`

GetFlowCloneOk returns a tuple with the FlowClone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowClone

`func (o *DaVinciFlowResponseLinks) SetFlowClone(v JSONHALLink)`

SetFlowClone sets FlowClone field to given value.


### GetFlowEnabled

`func (o *DaVinciFlowResponseLinks) GetFlowEnabled() JSONHALLink`

GetFlowEnabled returns the FlowEnabled field if non-nil, zero value otherwise.

### GetFlowEnabledOk

`func (o *DaVinciFlowResponseLinks) GetFlowEnabledOk() (*JSONHALLink, bool)`

GetFlowEnabledOk returns a tuple with the FlowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowEnabled

`func (o *DaVinciFlowResponseLinks) SetFlowEnabled(v JSONHALLink)`

SetFlowEnabled sets FlowEnabled field to given value.


### GetVersion

`func (o *DaVinciFlowResponseLinks) GetVersion() JSONHALLink`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciFlowResponseLinks) GetVersionOk() (*JSONHALLink, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciFlowResponseLinks) SetVersion(v JSONHALLink)`

SetVersion sets Version field to given value.


### GetFlowValidate

`func (o *DaVinciFlowResponseLinks) GetFlowValidate() JSONHALLink`

GetFlowValidate returns the FlowValidate field if non-nil, zero value otherwise.

### GetFlowValidateOk

`func (o *DaVinciFlowResponseLinks) GetFlowValidateOk() (*JSONHALLink, bool)`

GetFlowValidateOk returns a tuple with the FlowValidate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowValidate

`func (o *DaVinciFlowResponseLinks) SetFlowValidate(v JSONHALLink)`

SetFlowValidate sets FlowValidate field to given value.

### HasFlowValidate

`func (o *DaVinciFlowResponseLinks) HasFlowValidate() bool`

HasFlowValidate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


