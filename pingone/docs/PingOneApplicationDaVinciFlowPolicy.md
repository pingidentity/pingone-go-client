# PingOneApplicationDaVinciFlowPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**PingOneApplicationDaVinciFlowPolicyLinks**](PingOneApplicationDaVinciFlowPolicyLinks.md) |  | [optional] 
**Application** | Pointer to [**PingOneApplicationFlowPolicyAssignmentApplication**](PingOneApplicationFlowPolicyAssignmentApplication.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Environment** | Pointer to [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to [**DaVinciFlowPolicyTrigger**](DaVinciFlowPolicyTrigger.md) |  | [optional] 

## Methods

### NewPingOneApplicationDaVinciFlowPolicy

`func NewPingOneApplicationDaVinciFlowPolicy(id string, ) *PingOneApplicationDaVinciFlowPolicy`

NewPingOneApplicationDaVinciFlowPolicy instantiates a new PingOneApplicationDaVinciFlowPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPingOneApplicationDaVinciFlowPolicyWithDefaults

`func NewPingOneApplicationDaVinciFlowPolicyWithDefaults() *PingOneApplicationDaVinciFlowPolicy`

NewPingOneApplicationDaVinciFlowPolicyWithDefaults instantiates a new PingOneApplicationDaVinciFlowPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PingOneApplicationDaVinciFlowPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PingOneApplicationDaVinciFlowPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PingOneApplicationDaVinciFlowPolicy) SetId(v string)`

SetId sets Id field to given value.


### GetEmbedded

`func (o *PingOneApplicationDaVinciFlowPolicy) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *PingOneApplicationDaVinciFlowPolicy) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *PingOneApplicationDaVinciFlowPolicy) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *PingOneApplicationDaVinciFlowPolicy) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *PingOneApplicationDaVinciFlowPolicy) GetLinks() PingOneApplicationDaVinciFlowPolicyLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PingOneApplicationDaVinciFlowPolicy) GetLinksOk() (*PingOneApplicationDaVinciFlowPolicyLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PingOneApplicationDaVinciFlowPolicy) SetLinks(v PingOneApplicationDaVinciFlowPolicyLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PingOneApplicationDaVinciFlowPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetApplication

`func (o *PingOneApplicationDaVinciFlowPolicy) GetApplication() PingOneApplicationFlowPolicyAssignmentApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *PingOneApplicationDaVinciFlowPolicy) GetApplicationOk() (*PingOneApplicationFlowPolicyAssignmentApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *PingOneApplicationDaVinciFlowPolicy) SetApplication(v PingOneApplicationFlowPolicyAssignmentApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *PingOneApplicationDaVinciFlowPolicy) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetEnabled

`func (o *PingOneApplicationDaVinciFlowPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PingOneApplicationDaVinciFlowPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PingOneApplicationDaVinciFlowPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PingOneApplicationDaVinciFlowPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetEnvironment

`func (o *PingOneApplicationDaVinciFlowPolicy) GetEnvironment() ResourceRelationshipReadOnly`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PingOneApplicationDaVinciFlowPolicy) GetEnvironmentOk() (*ResourceRelationshipReadOnly, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PingOneApplicationDaVinciFlowPolicy) SetEnvironment(v ResourceRelationshipReadOnly)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PingOneApplicationDaVinciFlowPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *PingOneApplicationDaVinciFlowPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PingOneApplicationDaVinciFlowPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PingOneApplicationDaVinciFlowPolicy) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PingOneApplicationDaVinciFlowPolicy) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrigger

`func (o *PingOneApplicationDaVinciFlowPolicy) GetTrigger() DaVinciFlowPolicyTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PingOneApplicationDaVinciFlowPolicy) GetTriggerOk() (*DaVinciFlowPolicyTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PingOneApplicationDaVinciFlowPolicy) SetTrigger(v DaVinciFlowPolicyTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *PingOneApplicationDaVinciFlowPolicy) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


