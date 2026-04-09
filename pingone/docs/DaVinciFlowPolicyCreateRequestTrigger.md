# DaVinciFlowPolicyCreateRequestTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**DaVinciFlowPolicyCreateRequestTriggerConfiguration**](DaVinciFlowPolicyCreateRequestTriggerConfiguration.md) |  | [optional] 
**Subtype** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**DaVinciFlowPolicyCreateRequestTriggerType**](DaVinciFlowPolicyCreateRequestTriggerType.md) |  | [optional] [default to DAVINCIFLOWPOLICYCREATEREQUESTTRIGGERTYPE_AUTHENTICATION]

## Methods

### NewDaVinciFlowPolicyCreateRequestTrigger

`func NewDaVinciFlowPolicyCreateRequestTrigger() *DaVinciFlowPolicyCreateRequestTrigger`

NewDaVinciFlowPolicyCreateRequestTrigger instantiates a new DaVinciFlowPolicyCreateRequestTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowPolicyCreateRequestTriggerWithDefaults

`func NewDaVinciFlowPolicyCreateRequestTriggerWithDefaults() *DaVinciFlowPolicyCreateRequestTrigger`

NewDaVinciFlowPolicyCreateRequestTriggerWithDefaults instantiates a new DaVinciFlowPolicyCreateRequestTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *DaVinciFlowPolicyCreateRequestTrigger) GetConfiguration() DaVinciFlowPolicyCreateRequestTriggerConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *DaVinciFlowPolicyCreateRequestTrigger) GetConfigurationOk() (*DaVinciFlowPolicyCreateRequestTriggerConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *DaVinciFlowPolicyCreateRequestTrigger) SetConfiguration(v DaVinciFlowPolicyCreateRequestTriggerConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *DaVinciFlowPolicyCreateRequestTrigger) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetSubtype

`func (o *DaVinciFlowPolicyCreateRequestTrigger) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *DaVinciFlowPolicyCreateRequestTrigger) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *DaVinciFlowPolicyCreateRequestTrigger) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *DaVinciFlowPolicyCreateRequestTrigger) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetType

`func (o *DaVinciFlowPolicyCreateRequestTrigger) GetType() DaVinciFlowPolicyCreateRequestTriggerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DaVinciFlowPolicyCreateRequestTrigger) GetTypeOk() (*DaVinciFlowPolicyCreateRequestTriggerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DaVinciFlowPolicyCreateRequestTrigger) SetType(v DaVinciFlowPolicyCreateRequestTriggerType)`

SetType sets Type field to given value.

### HasType

`func (o *DaVinciFlowPolicyCreateRequestTrigger) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


