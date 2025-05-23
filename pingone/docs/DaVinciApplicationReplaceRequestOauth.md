# DaVinciApplicationReplaceRequestOauth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforceSignedRequestOpenid** | Pointer to **bool** |  | [optional] 
**GrantTypes** | Pointer to [**[]DaVinciApplicationReplaceRequestOauthGrantTypes**](DaVinciApplicationReplaceRequestOauthGrantTypes.md) |  | [optional] 
**LogoutUris** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to [**[]DaVinciApplicationReplaceRequestOauthScopes**](DaVinciApplicationReplaceRequestOauthScopes.md) |  | [optional] 
**SpJwksOpenid** | Pointer to **string** |  | [optional] 
**SpjwksUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewDaVinciApplicationReplaceRequestOauth

`func NewDaVinciApplicationReplaceRequestOauth() *DaVinciApplicationReplaceRequestOauth`

NewDaVinciApplicationReplaceRequestOauth instantiates a new DaVinciApplicationReplaceRequestOauth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciApplicationReplaceRequestOauthWithDefaults

`func NewDaVinciApplicationReplaceRequestOauthWithDefaults() *DaVinciApplicationReplaceRequestOauth`

NewDaVinciApplicationReplaceRequestOauthWithDefaults instantiates a new DaVinciApplicationReplaceRequestOauth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationReplaceRequestOauth) GetEnforceSignedRequestOpenid() bool`

GetEnforceSignedRequestOpenid returns the EnforceSignedRequestOpenid field if non-nil, zero value otherwise.

### GetEnforceSignedRequestOpenidOk

`func (o *DaVinciApplicationReplaceRequestOauth) GetEnforceSignedRequestOpenidOk() (*bool, bool)`

GetEnforceSignedRequestOpenidOk returns a tuple with the EnforceSignedRequestOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationReplaceRequestOauth) SetEnforceSignedRequestOpenid(v bool)`

SetEnforceSignedRequestOpenid sets EnforceSignedRequestOpenid field to given value.

### HasEnforceSignedRequestOpenid

`func (o *DaVinciApplicationReplaceRequestOauth) HasEnforceSignedRequestOpenid() bool`

HasEnforceSignedRequestOpenid returns a boolean if a field has been set.

### GetGrantTypes

`func (o *DaVinciApplicationReplaceRequestOauth) GetGrantTypes() []DaVinciApplicationReplaceRequestOauthGrantTypes`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *DaVinciApplicationReplaceRequestOauth) GetGrantTypesOk() (*[]DaVinciApplicationReplaceRequestOauthGrantTypes, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *DaVinciApplicationReplaceRequestOauth) SetGrantTypes(v []DaVinciApplicationReplaceRequestOauthGrantTypes)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *DaVinciApplicationReplaceRequestOauth) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetLogoutUris

`func (o *DaVinciApplicationReplaceRequestOauth) GetLogoutUris() []string`

GetLogoutUris returns the LogoutUris field if non-nil, zero value otherwise.

### GetLogoutUrisOk

`func (o *DaVinciApplicationReplaceRequestOauth) GetLogoutUrisOk() (*[]string, bool)`

GetLogoutUrisOk returns a tuple with the LogoutUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUris

`func (o *DaVinciApplicationReplaceRequestOauth) SetLogoutUris(v []string)`

SetLogoutUris sets LogoutUris field to given value.

### HasLogoutUris

`func (o *DaVinciApplicationReplaceRequestOauth) HasLogoutUris() bool`

HasLogoutUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *DaVinciApplicationReplaceRequestOauth) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *DaVinciApplicationReplaceRequestOauth) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *DaVinciApplicationReplaceRequestOauth) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *DaVinciApplicationReplaceRequestOauth) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *DaVinciApplicationReplaceRequestOauth) GetScopes() []DaVinciApplicationReplaceRequestOauthScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DaVinciApplicationReplaceRequestOauth) GetScopesOk() (*[]DaVinciApplicationReplaceRequestOauthScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DaVinciApplicationReplaceRequestOauth) SetScopes(v []DaVinciApplicationReplaceRequestOauthScopes)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DaVinciApplicationReplaceRequestOauth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSpJwksOpenid

`func (o *DaVinciApplicationReplaceRequestOauth) GetSpJwksOpenid() string`

GetSpJwksOpenid returns the SpJwksOpenid field if non-nil, zero value otherwise.

### GetSpJwksOpenidOk

`func (o *DaVinciApplicationReplaceRequestOauth) GetSpJwksOpenidOk() (*string, bool)`

GetSpJwksOpenidOk returns a tuple with the SpJwksOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpJwksOpenid

`func (o *DaVinciApplicationReplaceRequestOauth) SetSpJwksOpenid(v string)`

SetSpJwksOpenid sets SpJwksOpenid field to given value.

### HasSpJwksOpenid

`func (o *DaVinciApplicationReplaceRequestOauth) HasSpJwksOpenid() bool`

HasSpJwksOpenid returns a boolean if a field has been set.

### GetSpjwksUrl

`func (o *DaVinciApplicationReplaceRequestOauth) GetSpjwksUrl() string`

GetSpjwksUrl returns the SpjwksUrl field if non-nil, zero value otherwise.

### GetSpjwksUrlOk

`func (o *DaVinciApplicationReplaceRequestOauth) GetSpjwksUrlOk() (*string, bool)`

GetSpjwksUrlOk returns a tuple with the SpjwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpjwksUrl

`func (o *DaVinciApplicationReplaceRequestOauth) SetSpjwksUrl(v string)`

SetSpjwksUrl sets SpjwksUrl field to given value.

### HasSpjwksUrl

`func (o *DaVinciApplicationReplaceRequestOauth) HasSpjwksUrl() bool`

HasSpjwksUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


