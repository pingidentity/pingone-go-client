# DaVinciFlowPolicyReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [default to "New Policy"]
**FlowDistributions** | Pointer to [**[]DaVinciFlowPolicyReplaceRequestFlowDistribution**](DaVinciFlowPolicyReplaceRequestFlowDistribution.md) |  | [optional] 
**Status** | Pointer to [**DaVinciFlowPolicyReplaceRequestStatus**](DaVinciFlowPolicyReplaceRequestStatus.md) |  | [optional] [default to DAVINCIFLOWPOLICYREPLACEREQUESTSTATUS_ENABLED]
**Trigger** | Pointer to [**DaVinciFlowPolicyReplaceRequestTrigger**](DaVinciFlowPolicyReplaceRequestTrigger.md) |  | [optional] 

## Methods

### NewDaVinciFlowPolicyReplaceRequest

`func NewDaVinciFlowPolicyReplaceRequest(name string, ) *DaVinciFlowPolicyReplaceRequest`

NewDaVinciFlowPolicyReplaceRequest instantiates a new DaVinciFlowPolicyReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowPolicyReplaceRequestWithDefaults

`func NewDaVinciFlowPolicyReplaceRequestWithDefaults() *DaVinciFlowPolicyReplaceRequest`

NewDaVinciFlowPolicyReplaceRequestWithDefaults instantiates a new DaVinciFlowPolicyReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DaVinciFlowPolicyReplaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciFlowPolicyReplaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciFlowPolicyReplaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFlowDistributions

`func (o *DaVinciFlowPolicyReplaceRequest) GetFlowDistributions() []DaVinciFlowPolicyReplaceRequestFlowDistribution`

GetFlowDistributions returns the FlowDistributions field if non-nil, zero value otherwise.

### GetFlowDistributionsOk

`func (o *DaVinciFlowPolicyReplaceRequest) GetFlowDistributionsOk() (*[]DaVinciFlowPolicyReplaceRequestFlowDistribution, bool)`

GetFlowDistributionsOk returns a tuple with the FlowDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDistributions

`func (o *DaVinciFlowPolicyReplaceRequest) SetFlowDistributions(v []DaVinciFlowPolicyReplaceRequestFlowDistribution)`

SetFlowDistributions sets FlowDistributions field to given value.

### HasFlowDistributions

`func (o *DaVinciFlowPolicyReplaceRequest) HasFlowDistributions() bool`

HasFlowDistributions returns a boolean if a field has been set.

### GetStatus

`func (o *DaVinciFlowPolicyReplaceRequest) GetStatus() DaVinciFlowPolicyReplaceRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DaVinciFlowPolicyReplaceRequest) GetStatusOk() (*DaVinciFlowPolicyReplaceRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DaVinciFlowPolicyReplaceRequest) SetStatus(v DaVinciFlowPolicyReplaceRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DaVinciFlowPolicyReplaceRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrigger

`func (o *DaVinciFlowPolicyReplaceRequest) GetTrigger() DaVinciFlowPolicyReplaceRequestTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciFlowPolicyReplaceRequest) GetTriggerOk() (*DaVinciFlowPolicyReplaceRequestTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciFlowPolicyReplaceRequest) SetTrigger(v DaVinciFlowPolicyReplaceRequestTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciFlowPolicyReplaceRequest) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


