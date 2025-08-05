# DaVinciApplicationOAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **string** |  | 
**EnforceSignedRequestOpenid** | Pointer to **bool** |  | [optional] 
**GrantTypes** | Pointer to [**[]DaVinciApplicationOAuthGrantType**](DaVinciApplicationOAuthGrantType.md) |  | [optional] [default to [authorizationCode]]
**LogoutUris** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to [**[]DaVinciApplicationOAuthScope**](DaVinciApplicationOAuthScope.md) |  | [optional] [default to [openid, profile]]
**SpJwksOpenid** | Pointer to **string** |  | [optional] 
**SpjwksUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewDaVinciApplicationOAuth

`func NewDaVinciApplicationOAuth(clientSecret string, ) *DaVinciApplicationOAuth`

NewDaVinciApplicationOAuth instantiates a new DaVinciApplicationOAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciApplicationOAuthWithDefaults

`func NewDaVinciApplicationOAuthWithDefaults() *DaVinciApplicationOAuth`

NewDaVinciApplicationOAuthWithDefaults instantiates a new DaVinciApplicationOAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *DaVinciApplicationOAuth) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *DaVinciApplicationOAuth) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *DaVinciApplicationOAuth) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationOAuth) GetEnforceSignedRequestOpenid() bool`

GetEnforceSignedRequestOpenid returns the EnforceSignedRequestOpenid field if non-nil, zero value otherwise.

### GetEnforceSignedRequestOpenidOk

`func (o *DaVinciApplicationOAuth) GetEnforceSignedRequestOpenidOk() (*bool, bool)`

GetEnforceSignedRequestOpenidOk returns a tuple with the EnforceSignedRequestOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationOAuth) SetEnforceSignedRequestOpenid(v bool)`

SetEnforceSignedRequestOpenid sets EnforceSignedRequestOpenid field to given value.

### HasEnforceSignedRequestOpenid

`func (o *DaVinciApplicationOAuth) HasEnforceSignedRequestOpenid() bool`

HasEnforceSignedRequestOpenid returns a boolean if a field has been set.

### GetGrantTypes

`func (o *DaVinciApplicationOAuth) GetGrantTypes() []DaVinciApplicationOAuthGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *DaVinciApplicationOAuth) GetGrantTypesOk() (*[]DaVinciApplicationOAuthGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *DaVinciApplicationOAuth) SetGrantTypes(v []DaVinciApplicationOAuthGrantType)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *DaVinciApplicationOAuth) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetLogoutUris

`func (o *DaVinciApplicationOAuth) GetLogoutUris() []string`

GetLogoutUris returns the LogoutUris field if non-nil, zero value otherwise.

### GetLogoutUrisOk

`func (o *DaVinciApplicationOAuth) GetLogoutUrisOk() (*[]string, bool)`

GetLogoutUrisOk returns a tuple with the LogoutUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUris

`func (o *DaVinciApplicationOAuth) SetLogoutUris(v []string)`

SetLogoutUris sets LogoutUris field to given value.

### HasLogoutUris

`func (o *DaVinciApplicationOAuth) HasLogoutUris() bool`

HasLogoutUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *DaVinciApplicationOAuth) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *DaVinciApplicationOAuth) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *DaVinciApplicationOAuth) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *DaVinciApplicationOAuth) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *DaVinciApplicationOAuth) GetScopes() []DaVinciApplicationOAuthScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DaVinciApplicationOAuth) GetScopesOk() (*[]DaVinciApplicationOAuthScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DaVinciApplicationOAuth) SetScopes(v []DaVinciApplicationOAuthScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DaVinciApplicationOAuth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSpJwksOpenid

`func (o *DaVinciApplicationOAuth) GetSpJwksOpenid() string`

GetSpJwksOpenid returns the SpJwksOpenid field if non-nil, zero value otherwise.

### GetSpJwksOpenidOk

`func (o *DaVinciApplicationOAuth) GetSpJwksOpenidOk() (*string, bool)`

GetSpJwksOpenidOk returns a tuple with the SpJwksOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpJwksOpenid

`func (o *DaVinciApplicationOAuth) SetSpJwksOpenid(v string)`

SetSpJwksOpenid sets SpJwksOpenid field to given value.

### HasSpJwksOpenid

`func (o *DaVinciApplicationOAuth) HasSpJwksOpenid() bool`

HasSpJwksOpenid returns a boolean if a field has been set.

### GetSpjwksUrl

`func (o *DaVinciApplicationOAuth) GetSpjwksUrl() string`

GetSpjwksUrl returns the SpjwksUrl field if non-nil, zero value otherwise.

### GetSpjwksUrlOk

`func (o *DaVinciApplicationOAuth) GetSpjwksUrlOk() (*string, bool)`

GetSpjwksUrlOk returns a tuple with the SpjwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpjwksUrl

`func (o *DaVinciApplicationOAuth) SetSpjwksUrl(v string)`

SetSpjwksUrl sets SpjwksUrl field to given value.

### HasSpjwksUrl

`func (o *DaVinciApplicationOAuth) HasSpjwksUrl() bool`

HasSpjwksUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


