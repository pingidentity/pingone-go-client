# DaVinciApplicationReplaceRequestOAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnforceSignedRequestOpenid** | Pointer to **bool** |  | [optional] 
**GrantTypes** | Pointer to [**[]DaVinciApplicationReplaceRequestOAuthGrantType**](DaVinciApplicationReplaceRequestOAuthGrantType.md) |  | [optional] 
**LogoutUris** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to [**[]DaVinciApplicationReplaceRequestOAuthScope**](DaVinciApplicationReplaceRequestOAuthScope.md) |  | [optional] 
**SpJwksOpenid** | Pointer to **string** |  | [optional] 
**SpjwksUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewDaVinciApplicationReplaceRequestOAuth

`func NewDaVinciApplicationReplaceRequestOAuth() *DaVinciApplicationReplaceRequestOAuth`

NewDaVinciApplicationReplaceRequestOAuth instantiates a new DaVinciApplicationReplaceRequestOAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciApplicationReplaceRequestOAuthWithDefaults

`func NewDaVinciApplicationReplaceRequestOAuthWithDefaults() *DaVinciApplicationReplaceRequestOAuth`

NewDaVinciApplicationReplaceRequestOAuthWithDefaults instantiates a new DaVinciApplicationReplaceRequestOAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationReplaceRequestOAuth) GetEnforceSignedRequestOpenid() bool`

GetEnforceSignedRequestOpenid returns the EnforceSignedRequestOpenid field if non-nil, zero value otherwise.

### GetEnforceSignedRequestOpenidOk

`func (o *DaVinciApplicationReplaceRequestOAuth) GetEnforceSignedRequestOpenidOk() (*bool, bool)`

GetEnforceSignedRequestOpenidOk returns a tuple with the EnforceSignedRequestOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationReplaceRequestOAuth) SetEnforceSignedRequestOpenid(v bool)`

SetEnforceSignedRequestOpenid sets EnforceSignedRequestOpenid field to given value.

### HasEnforceSignedRequestOpenid

`func (o *DaVinciApplicationReplaceRequestOAuth) HasEnforceSignedRequestOpenid() bool`

HasEnforceSignedRequestOpenid returns a boolean if a field has been set.

### GetGrantTypes

`func (o *DaVinciApplicationReplaceRequestOAuth) GetGrantTypes() []DaVinciApplicationReplaceRequestOAuthGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *DaVinciApplicationReplaceRequestOAuth) GetGrantTypesOk() (*[]DaVinciApplicationReplaceRequestOAuthGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *DaVinciApplicationReplaceRequestOAuth) SetGrantTypes(v []DaVinciApplicationReplaceRequestOAuthGrantType)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *DaVinciApplicationReplaceRequestOAuth) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetLogoutUris

`func (o *DaVinciApplicationReplaceRequestOAuth) GetLogoutUris() []string`

GetLogoutUris returns the LogoutUris field if non-nil, zero value otherwise.

### GetLogoutUrisOk

`func (o *DaVinciApplicationReplaceRequestOAuth) GetLogoutUrisOk() (*[]string, bool)`

GetLogoutUrisOk returns a tuple with the LogoutUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUris

`func (o *DaVinciApplicationReplaceRequestOAuth) SetLogoutUris(v []string)`

SetLogoutUris sets LogoutUris field to given value.

### HasLogoutUris

`func (o *DaVinciApplicationReplaceRequestOAuth) HasLogoutUris() bool`

HasLogoutUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *DaVinciApplicationReplaceRequestOAuth) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *DaVinciApplicationReplaceRequestOAuth) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *DaVinciApplicationReplaceRequestOAuth) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *DaVinciApplicationReplaceRequestOAuth) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *DaVinciApplicationReplaceRequestOAuth) GetScopes() []DaVinciApplicationReplaceRequestOAuthScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DaVinciApplicationReplaceRequestOAuth) GetScopesOk() (*[]DaVinciApplicationReplaceRequestOAuthScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DaVinciApplicationReplaceRequestOAuth) SetScopes(v []DaVinciApplicationReplaceRequestOAuthScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DaVinciApplicationReplaceRequestOAuth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSpJwksOpenid

`func (o *DaVinciApplicationReplaceRequestOAuth) GetSpJwksOpenid() string`

GetSpJwksOpenid returns the SpJwksOpenid field if non-nil, zero value otherwise.

### GetSpJwksOpenidOk

`func (o *DaVinciApplicationReplaceRequestOAuth) GetSpJwksOpenidOk() (*string, bool)`

GetSpJwksOpenidOk returns a tuple with the SpJwksOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpJwksOpenid

`func (o *DaVinciApplicationReplaceRequestOAuth) SetSpJwksOpenid(v string)`

SetSpJwksOpenid sets SpJwksOpenid field to given value.

### HasSpJwksOpenid

`func (o *DaVinciApplicationReplaceRequestOAuth) HasSpJwksOpenid() bool`

HasSpJwksOpenid returns a boolean if a field has been set.

### GetSpjwksUrl

`func (o *DaVinciApplicationReplaceRequestOAuth) GetSpjwksUrl() string`

GetSpjwksUrl returns the SpjwksUrl field if non-nil, zero value otherwise.

### GetSpjwksUrlOk

`func (o *DaVinciApplicationReplaceRequestOAuth) GetSpjwksUrlOk() (*string, bool)`

GetSpjwksUrlOk returns a tuple with the SpjwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpjwksUrl

`func (o *DaVinciApplicationReplaceRequestOAuth) SetSpjwksUrl(v string)`

SetSpjwksUrl sets SpjwksUrl field to given value.

### HasSpjwksUrl

`func (o *DaVinciApplicationReplaceRequestOAuth) HasSpjwksUrl() bool`

HasSpjwksUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


