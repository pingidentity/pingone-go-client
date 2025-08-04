# DaVinciApplicationReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ApiKeyEnabled** | Pointer to **bool** |  | [optional] 
**Oauth** | Pointer to [**DaVinciApplicationReplaceRequestOAuth**](DaVinciApplicationReplaceRequestOAuth.md) |  | [optional] 

## Methods

### NewDaVinciApplicationReplaceRequest

`func NewDaVinciApplicationReplaceRequest(name string, ) *DaVinciApplicationReplaceRequest`

NewDaVinciApplicationReplaceRequest instantiates a new DaVinciApplicationReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciApplicationReplaceRequestWithDefaults

`func NewDaVinciApplicationReplaceRequestWithDefaults() *DaVinciApplicationReplaceRequest`

NewDaVinciApplicationReplaceRequestWithDefaults instantiates a new DaVinciApplicationReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DaVinciApplicationReplaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciApplicationReplaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciApplicationReplaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetApiKeyEnabled

`func (o *DaVinciApplicationReplaceRequest) GetApiKeyEnabled() bool`

GetApiKeyEnabled returns the ApiKeyEnabled field if non-nil, zero value otherwise.

### GetApiKeyEnabledOk

`func (o *DaVinciApplicationReplaceRequest) GetApiKeyEnabledOk() (*bool, bool)`

GetApiKeyEnabledOk returns a tuple with the ApiKeyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyEnabled

`func (o *DaVinciApplicationReplaceRequest) SetApiKeyEnabled(v bool)`

SetApiKeyEnabled sets ApiKeyEnabled field to given value.

### HasApiKeyEnabled

`func (o *DaVinciApplicationReplaceRequest) HasApiKeyEnabled() bool`

HasApiKeyEnabled returns a boolean if a field has been set.

### GetOauth

`func (o *DaVinciApplicationReplaceRequest) GetOauth() DaVinciApplicationReplaceRequestOAuth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *DaVinciApplicationReplaceRequest) GetOauthOk() (*DaVinciApplicationReplaceRequestOAuth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *DaVinciApplicationReplaceRequest) SetOauth(v DaVinciApplicationReplaceRequestOAuth)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *DaVinciApplicationReplaceRequest) HasOauth() bool`

HasOauth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


