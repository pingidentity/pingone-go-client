# DaVinciFlowVersionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciFlowVersionResponseLinks**](DaVinciFlowVersionResponseLinks.md) |  | 
**Environment** | [**DaVinciFlowVersionResponseEnvironment**](DaVinciFlowVersionResponseEnvironment.md) |  | 
**Flow** | [**DaVinciFlowVersionResponseFlow**](DaVinciFlowVersionResponseFlow.md) |  | 
**Version** | **float32** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**ClonedFrom** | Pointer to **float32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**DeployedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDaVinciFlowVersionResponse

`func NewDaVinciFlowVersionResponse(links DaVinciFlowVersionResponseLinks, environment DaVinciFlowVersionResponseEnvironment, flow DaVinciFlowVersionResponseFlow, version float32, ) *DaVinciFlowVersionResponse`

NewDaVinciFlowVersionResponse instantiates a new DaVinciFlowVersionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciFlowVersionResponseWithDefaults

`func NewDaVinciFlowVersionResponseWithDefaults() *DaVinciFlowVersionResponse`

NewDaVinciFlowVersionResponseWithDefaults instantiates a new DaVinciFlowVersionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciFlowVersionResponse) GetLinks() DaVinciFlowVersionResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciFlowVersionResponse) GetLinksOk() (*DaVinciFlowVersionResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciFlowVersionResponse) SetLinks(v DaVinciFlowVersionResponseLinks)`

SetLinks sets Links field to given value.


### GetEnvironment

`func (o *DaVinciFlowVersionResponse) GetEnvironment() DaVinciFlowVersionResponseEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciFlowVersionResponse) GetEnvironmentOk() (*DaVinciFlowVersionResponseEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciFlowVersionResponse) SetEnvironment(v DaVinciFlowVersionResponseEnvironment)`

SetEnvironment sets Environment field to given value.


### GetFlow

`func (o *DaVinciFlowVersionResponse) GetFlow() DaVinciFlowVersionResponseFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *DaVinciFlowVersionResponse) GetFlowOk() (*DaVinciFlowVersionResponseFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *DaVinciFlowVersionResponse) SetFlow(v DaVinciFlowVersionResponseFlow)`

SetFlow sets Flow field to given value.


### GetVersion

`func (o *DaVinciFlowVersionResponse) GetVersion() float32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DaVinciFlowVersionResponse) GetVersionOk() (*float32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DaVinciFlowVersionResponse) SetVersion(v float32)`

SetVersion sets Version field to given value.


### GetAlias

`func (o *DaVinciFlowVersionResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *DaVinciFlowVersionResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *DaVinciFlowVersionResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *DaVinciFlowVersionResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetClonedFrom

`func (o *DaVinciFlowVersionResponse) GetClonedFrom() float32`

GetClonedFrom returns the ClonedFrom field if non-nil, zero value otherwise.

### GetClonedFromOk

`func (o *DaVinciFlowVersionResponse) GetClonedFromOk() (*float32, bool)`

GetClonedFromOk returns a tuple with the ClonedFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClonedFrom

`func (o *DaVinciFlowVersionResponse) SetClonedFrom(v float32)`

SetClonedFrom sets ClonedFrom field to given value.

### HasClonedFrom

`func (o *DaVinciFlowVersionResponse) HasClonedFrom() bool`

HasClonedFrom returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DaVinciFlowVersionResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciFlowVersionResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciFlowVersionResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciFlowVersionResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDeployedAt

`func (o *DaVinciFlowVersionResponse) GetDeployedAt() time.Time`

GetDeployedAt returns the DeployedAt field if non-nil, zero value otherwise.

### GetDeployedAtOk

`func (o *DaVinciFlowVersionResponse) GetDeployedAtOk() (*time.Time, bool)`

GetDeployedAtOk returns a tuple with the DeployedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployedAt

`func (o *DaVinciFlowVersionResponse) SetDeployedAt(v time.Time)`

SetDeployedAt sets DeployedAt field to given value.

### HasDeployedAt

`func (o *DaVinciFlowVersionResponse) HasDeployedAt() bool`

HasDeployedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciFlowVersionResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciFlowVersionResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciFlowVersionResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciFlowVersionResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


