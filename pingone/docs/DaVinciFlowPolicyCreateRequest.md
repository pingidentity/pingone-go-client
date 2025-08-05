# DaVinciFlowPolicyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowDistributions** | [**[]DaVinciFlowPolicyCreateRequestFlowDistribution**](DaVinciFlowPolicyCreateRequestFlowDistribution.md) |  | 
**Name** | Pointer to **string** |  | [optional] [default to "New Policy"]
**Status** | Pointer to [**DaVinciFlowPolicyCreateRequestStatus**](DaVinciFlowPolicyCreateRequestStatus.md) |  | [optional] [default to DAVINCIFLOWPOLICYCREATEREQUESTSTATUS_ENABLED]
**Trigger** | Pointer to [**DaVinciFlowPolicyCreateRequestTrigger**](DaVinciFlowPolicyCreateRequestTrigger.md) |  | [optional] 

## Methods

### NewDaVinciFlowPolicyCreateRequest

`func NewDaVinciFlowPolicyCreateRequest(flowDistributions []DaVinciFlowPolicyCreateRequestFlowDistribution, ) *DaVinciFlowPolicyCreateRequest`

NewDaVinciFlowPolicyCreateRequest instantiates a new DaVinciFlowPolicyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowPolicyCreateRequestWithDefaults

`func NewDaVinciFlowPolicyCreateRequestWithDefaults() *DaVinciFlowPolicyCreateRequest`

NewDaVinciFlowPolicyCreateRequestWithDefaults instantiates a new DaVinciFlowPolicyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowDistributions

`func (o *DaVinciFlowPolicyCreateRequest) GetFlowDistributions() []DaVinciFlowPolicyCreateRequestFlowDistribution`

GetFlowDistributions returns the FlowDistributions field if non-nil, zero value otherwise.

### GetFlowDistributionsOk

`func (o *DaVinciFlowPolicyCreateRequest) GetFlowDistributionsOk() (*[]DaVinciFlowPolicyCreateRequestFlowDistribution, bool)`

GetFlowDistributionsOk returns a tuple with the FlowDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDistributions

`func (o *DaVinciFlowPolicyCreateRequest) SetFlowDistributions(v []DaVinciFlowPolicyCreateRequestFlowDistribution)`

SetFlowDistributions sets FlowDistributions field to given value.


### GetName

`func (o *DaVinciFlowPolicyCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciFlowPolicyCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciFlowPolicyCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DaVinciFlowPolicyCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DaVinciFlowPolicyCreateRequest) GetStatus() DaVinciFlowPolicyCreateRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DaVinciFlowPolicyCreateRequest) GetStatusOk() (*DaVinciFlowPolicyCreateRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DaVinciFlowPolicyCreateRequest) SetStatus(v DaVinciFlowPolicyCreateRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DaVinciFlowPolicyCreateRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciFlowPolicyCreateRequest) GetTrigger() DaVinciFlowPolicyCreateRequestTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciFlowPolicyCreateRequest) GetTriggerOk() (*DaVinciFlowPolicyCreateRequestTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciFlowPolicyCreateRequest) SetTrigger(v DaVinciFlowPolicyCreateRequestTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciFlowPolicyCreateRequest) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


