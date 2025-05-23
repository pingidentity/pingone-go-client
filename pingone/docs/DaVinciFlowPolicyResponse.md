# DaVinciFlowPolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciFlowPolicyResponseLinks**](DaVinciFlowPolicyResponseLinks.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**FlowDistributions** | [**[]DaVinciFlowPolicyResponseFlowDistribution**](DaVinciFlowPolicyResponseFlowDistribution.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Status** | [**DaVinciFlowPolicyCreateRequestStatus**](DaVinciFlowPolicyCreateRequestStatus.md) |  | 
**Trigger** | Pointer to [**NullableDaVinciFlowPolicyResponseTrigger**](DaVinciFlowPolicyResponseTrigger.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDaVinciFlowPolicyResponse

`func NewDaVinciFlowPolicyResponse(links DaVinciFlowPolicyResponseLinks, environment ResourceRelationshipPingOne, flowDistributions []DaVinciFlowPolicyResponseFlowDistribution, id string, name string, status DaVinciFlowPolicyCreateRequestStatus, ) *DaVinciFlowPolicyResponse`

NewDaVinciFlowPolicyResponse instantiates a new DaVinciFlowPolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowPolicyResponseWithDefaults

`func NewDaVinciFlowPolicyResponseWithDefaults() *DaVinciFlowPolicyResponse`

NewDaVinciFlowPolicyResponseWithDefaults instantiates a new DaVinciFlowPolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciFlowPolicyResponse) GetLinks() DaVinciFlowPolicyResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciFlowPolicyResponse) GetLinksOk() (*DaVinciFlowPolicyResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciFlowPolicyResponse) SetLinks(v DaVinciFlowPolicyResponseLinks)`

SetLinks sets Links field to given value.


### GetCreatedAt

`func (o *DaVinciFlowPolicyResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciFlowPolicyResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciFlowPolicyResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciFlowPolicyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *DaVinciFlowPolicyResponse) GetEnvironment() ResourceRelationshipPingOne`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciFlowPolicyResponse) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciFlowPolicyResponse) SetEnvironment(v ResourceRelationshipPingOne)`

SetEnvironment sets Environment field to given value.


### GetFlowDistributions

`func (o *DaVinciFlowPolicyResponse) GetFlowDistributions() []DaVinciFlowPolicyResponseFlowDistribution`

GetFlowDistributions returns the FlowDistributions field if non-nil, zero value otherwise.

### GetFlowDistributionsOk

`func (o *DaVinciFlowPolicyResponse) GetFlowDistributionsOk() (*[]DaVinciFlowPolicyResponseFlowDistribution, bool)`

GetFlowDistributionsOk returns a tuple with the FlowDistributions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowDistributions

`func (o *DaVinciFlowPolicyResponse) SetFlowDistributions(v []DaVinciFlowPolicyResponseFlowDistribution)`

SetFlowDistributions sets FlowDistributions field to given value.


### GetId

`func (o *DaVinciFlowPolicyResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciFlowPolicyResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciFlowPolicyResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DaVinciFlowPolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciFlowPolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciFlowPolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DaVinciFlowPolicyResponse) GetStatus() DaVinciFlowPolicyCreateRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DaVinciFlowPolicyResponse) GetStatusOk() (*DaVinciFlowPolicyCreateRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DaVinciFlowPolicyResponse) SetStatus(v DaVinciFlowPolicyCreateRequestStatus)`

SetStatus sets Status field to given value.


### GetTrigger

`func (o *DaVinciFlowPolicyResponse) GetTrigger() DaVinciFlowPolicyResponseTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *DaVinciFlowPolicyResponse) GetTriggerOk() (*DaVinciFlowPolicyResponseTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *DaVinciFlowPolicyResponse) SetTrigger(v DaVinciFlowPolicyResponseTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *DaVinciFlowPolicyResponse) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### SetTriggerNil

`func (o *DaVinciFlowPolicyResponse) SetTriggerNil(b bool)`

SetTriggerNil sets the value for Trigger to be an explicit nil

### UnsetTrigger
`func (o *DaVinciFlowPolicyResponse) UnsetTrigger()`

UnsetTrigger ensures that no value is present for Trigger, not even an explicit nil
### GetUpdatedAt

`func (o *DaVinciFlowPolicyResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciFlowPolicyResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciFlowPolicyResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciFlowPolicyResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


