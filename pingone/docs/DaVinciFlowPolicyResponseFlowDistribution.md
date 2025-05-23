# DaVinciFlowPolicyResponseFlowDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Ip** | Pointer to **[]string** |  | [optional] 
**SuccessNodes** | Pointer to [**[]DaVinciFlowPolicyResponseFlowDistributionSuccessNode**](DaVinciFlowPolicyResponseFlowDistributionSuccessNode.md) |  | [optional] 
**Version** | **float32** |  | 
**Weight** | Pointer to **float32** |  | [optional] 

## Methods

### NewDaVinciFlowPolicyResponseFlowDistribution

`func NewDaVinciFlowPolicyResponseFlowDistribution(id string, version float32, ) *DaVinciFlowPolicyResponseFlowDistribution`

NewDaVinciFlowPolicyResponseFlowDistribution instantiates a new DaVinciFlowPolicyResponseFlowDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowPolicyResponseFlowDistributionWithDefaults

`func NewDaVinciFlowPolicyResponseFlowDistributionWithDefaults() *DaVinciFlowPolicyResponseFlowDistribution`

NewDaVinciFlowPolicyResponseFlowDistributionWithDefaults instantiates a new DaVinciFlowPolicyResponseFlowDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciFlowPolicyResponseFlowDistribution) SetId(v string)`

SetId sets Id field to given value.


### GetIp

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DaVinciFlowPolicyResponseFlowDistribution) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DaVinciFlowPolicyResponseFlowDistribution) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSuccessNodes

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetSuccessNodes() []DaVinciFlowPolicyResponseFlowDistributionSuccessNode`

GetSuccessNodes returns the SuccessNodes field if non-nil, zero value otherwise.

### GetSuccessNodesOk

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetSuccessNodesOk() (*[]DaVinciFlowPolicyResponseFlowDistributionSuccessNode, bool)`

GetSuccessNodesOk returns a tuple with the SuccessNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessNodes

`func (o *DaVinciFlowPolicyResponseFlowDistribution) SetSuccessNodes(v []DaVinciFlowPolicyResponseFlowDistributionSuccessNode)`

SetSuccessNodes sets SuccessNodes field to given value.

### HasSuccessNodes

`func (o *DaVinciFlowPolicyResponseFlowDistribution) HasSuccessNodes() bool`

HasSuccessNodes returns a boolean if a field has been set.

### GetVersion

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciFlowPolicyResponseFlowDistribution) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetWeight

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DaVinciFlowPolicyResponseFlowDistribution) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DaVinciFlowPolicyResponseFlowDistribution) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DaVinciFlowPolicyResponseFlowDistribution) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


