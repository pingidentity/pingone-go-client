# DaVinciFlowPolicyEventsCollectionEmbeddedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | [**DaVinciFlowPolicyEventsCollectionEmbeddedEventEnvironment**](DaVinciFlowPolicyEventsCollectionEmbeddedEventEnvironment.md) |  | 
**Flow** | [**DaVinciFlowPolicyEventsCollectionEmbeddedEventFlow**](DaVinciFlowPolicyEventsCollectionEmbeddedEventFlow.md) |  | 
**Events** | Pointer to [**[]DaVinciFlowPolicyEventsCollectionEmbeddedEventEvent**](DaVinciFlowPolicyEventsCollectionEmbeddedEventEvent.md) |  | [optional] 
**SuccessCount** | Pointer to **float32** |  | [optional] 
**TotalCount** | Pointer to **float32** |  | [optional] 

## Methods

### NewDaVinciFlowPolicyEventsCollectionEmbeddedEvent

`func NewDaVinciFlowPolicyEventsCollectionEmbeddedEvent(environment DaVinciFlowPolicyEventsCollectionEmbeddedEventEnvironment, flow DaVinciFlowPolicyEventsCollectionEmbeddedEventFlow, ) *DaVinciFlowPolicyEventsCollectionEmbeddedEvent`

NewDaVinciFlowPolicyEventsCollectionEmbeddedEvent instantiates a new DaVinciFlowPolicyEventsCollectionEmbeddedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowPolicyEventsCollectionEmbeddedEventWithDefaults

`func NewDaVinciFlowPolicyEventsCollectionEmbeddedEventWithDefaults() *DaVinciFlowPolicyEventsCollectionEmbeddedEvent`

NewDaVinciFlowPolicyEventsCollectionEmbeddedEventWithDefaults instantiates a new DaVinciFlowPolicyEventsCollectionEmbeddedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetEnvironment() DaVinciFlowPolicyEventsCollectionEmbeddedEventEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetEnvironmentOk() (*DaVinciFlowPolicyEventsCollectionEmbeddedEventEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) SetEnvironment(v DaVinciFlowPolicyEventsCollectionEmbeddedEventEnvironment)`

SetEnvironment sets Environment field to given value.


### GetFlow

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetFlow() DaVinciFlowPolicyEventsCollectionEmbeddedEventFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetFlowOk() (*DaVinciFlowPolicyEventsCollectionEmbeddedEventFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) SetFlow(v DaVinciFlowPolicyEventsCollectionEmbeddedEventFlow)`

SetFlow sets Flow field to given value.


### GetEvents

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetEvents() []DaVinciFlowPolicyEventsCollectionEmbeddedEventEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetEventsOk() (*[]DaVinciFlowPolicyEventsCollectionEmbeddedEventEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) SetEvents(v []DaVinciFlowPolicyEventsCollectionEmbeddedEventEvent)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetSuccessCount

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetSuccessCount() float32`

GetSuccessCount returns the SuccessCount field if non-nil, zero value otherwise.

### GetSuccessCountOk

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetSuccessCountOk() (*float32, bool)`

GetSuccessCountOk returns a tuple with the SuccessCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessCount

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) SetSuccessCount(v float32)`

SetSuccessCount sets SuccessCount field to given value.

### HasSuccessCount

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) HasSuccessCount() bool`

HasSuccessCount returns a boolean if a field has been set.

### GetTotalCount

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.

### HasTotalCount

`func (o *DaVinciFlowPolicyEventsCollectionEmbeddedEvent) HasTotalCount() bool`

HasTotalCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


