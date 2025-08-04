# DaVinciFlowPolicyCreateRequestFlowDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Version** | **float32** |  | 
**Ip** | Pointer to **[]string** |  | [optional] 
**SuccessNodes** | Pointer to [**[]DaVinciFlowPolicyCreateRequestFlowDistributionSuccessNode**](DaVinciFlowPolicyCreateRequestFlowDistributionSuccessNode.md) |  | [optional] 
**Weight** | Pointer to **float32** |  | [optional] 

## Methods

### NewDaVinciFlowPolicyCreateRequestFlowDistribution

`func NewDaVinciFlowPolicyCreateRequestFlowDistribution(id string, version float32, ) *DaVinciFlowPolicyCreateRequestFlowDistribution`

NewDaVinciFlowPolicyCreateRequestFlowDistribution instantiates a new DaVinciFlowPolicyCreateRequestFlowDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowPolicyCreateRequestFlowDistributionWithDefaults

`func NewDaVinciFlowPolicyCreateRequestFlowDistributionWithDefaults() *DaVinciFlowPolicyCreateRequestFlowDistribution`

NewDaVinciFlowPolicyCreateRequestFlowDistributionWithDefaults instantiates a new DaVinciFlowPolicyCreateRequestFlowDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) SetId(v string)`

SetId sets Id field to given value.


### GetVersion

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetIp

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetIp() []string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetIpOk() (*[]string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) SetIp(v []string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSuccessNodes

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetSuccessNodes() []DaVinciFlowPolicyCreateRequestFlowDistributionSuccessNode`

GetSuccessNodes returns the SuccessNodes field if non-nil, zero value otherwise.

### GetSuccessNodesOk

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetSuccessNodesOk() (*[]DaVinciFlowPolicyCreateRequestFlowDistributionSuccessNode, bool)`

GetSuccessNodesOk returns a tuple with the SuccessNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessNodes

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) SetSuccessNodes(v []DaVinciFlowPolicyCreateRequestFlowDistributionSuccessNode)`

SetSuccessNodes sets SuccessNodes field to given value.

### HasSuccessNodes

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) HasSuccessNodes() bool`

HasSuccessNodes returns a boolean if a field has been set.

### GetWeight

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetWeight() float32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) GetWeightOk() (*float32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) SetWeight(v float32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *DaVinciFlowPolicyCreateRequestFlowDistribution) HasWeight() bool`

HasWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


