# DaVinciApplicationResponseOAuth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **string** |  | 
**EnforceSignedRequestOpenid** | Pointer to **bool** |  | [optional] 
**GrantTypes** | Pointer to [**[]DaVinciApplicationResponseOAuthGrantType**](DaVinciApplicationResponseOAuthGrantType.md) |  | [optional] [default to [authorizationCode]]
**LogoutUris** | Pointer to **[]string** |  | [optional] 
**RedirectUris** | Pointer to **[]string** |  | [optional] 
**Scopes** | Pointer to [**[]DaVinciApplicationResponseOAuthScope**](DaVinciApplicationResponseOAuthScope.md) |  | [optional] [default to [openid, profile]]
**SpJwksOpenid** | Pointer to **string** |  | [optional] 
**SpjwksUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewDaVinciApplicationResponseOAuth

`func NewDaVinciApplicationResponseOAuth(clientSecret string, ) *DaVinciApplicationResponseOAuth`

NewDaVinciApplicationResponseOAuth instantiates a new DaVinciApplicationResponseOAuth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciApplicationResponseOAuthWithDefaults

`func NewDaVinciApplicationResponseOAuthWithDefaults() *DaVinciApplicationResponseOAuth`

NewDaVinciApplicationResponseOAuthWithDefaults instantiates a new DaVinciApplicationResponseOAuth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *DaVinciApplicationResponseOAuth) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *DaVinciApplicationResponseOAuth) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *DaVinciApplicationResponseOAuth) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationResponseOAuth) GetEnforceSignedRequestOpenid() bool`

GetEnforceSignedRequestOpenid returns the EnforceSignedRequestOpenid field if non-nil, zero value otherwise.

### GetEnforceSignedRequestOpenidOk

`func (o *DaVinciApplicationResponseOAuth) GetEnforceSignedRequestOpenidOk() (*bool, bool)`

GetEnforceSignedRequestOpenidOk returns a tuple with the EnforceSignedRequestOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationResponseOAuth) SetEnforceSignedRequestOpenid(v bool)`

SetEnforceSignedRequestOpenid sets EnforceSignedRequestOpenid field to given value.

### HasEnforceSignedRequestOpenid

`func (o *DaVinciApplicationResponseOAuth) HasEnforceSignedRequestOpenid() bool`

HasEnforceSignedRequestOpenid returns a boolean if a field has been set.

### GetGrantTypes

`func (o *DaVinciApplicationResponseOAuth) GetGrantTypes() []DaVinciApplicationResponseOAuthGrantType`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *DaVinciApplicationResponseOAuth) GetGrantTypesOk() (*[]DaVinciApplicationResponseOAuthGrantType, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *DaVinciApplicationResponseOAuth) SetGrantTypes(v []DaVinciApplicationResponseOAuthGrantType)`

SetGrantTypes sets GrantTypes field to given value.

### HasGrantTypes

`func (o *DaVinciApplicationResponseOAuth) HasGrantTypes() bool`

HasGrantTypes returns a boolean if a field has been set.

### GetLogoutUris

`func (o *DaVinciApplicationResponseOAuth) GetLogoutUris() []string`

GetLogoutUris returns the LogoutUris field if non-nil, zero value otherwise.

### GetLogoutUrisOk

`func (o *DaVinciApplicationResponseOAuth) GetLogoutUrisOk() (*[]string, bool)`

GetLogoutUrisOk returns a tuple with the LogoutUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUris

`func (o *DaVinciApplicationResponseOAuth) SetLogoutUris(v []string)`

SetLogoutUris sets LogoutUris field to given value.

### HasLogoutUris

`func (o *DaVinciApplicationResponseOAuth) HasLogoutUris() bool`

HasLogoutUris returns a boolean if a field has been set.

### GetRedirectUris

`func (o *DaVinciApplicationResponseOAuth) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *DaVinciApplicationResponseOAuth) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *DaVinciApplicationResponseOAuth) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.

### HasRedirectUris

`func (o *DaVinciApplicationResponseOAuth) HasRedirectUris() bool`

HasRedirectUris returns a boolean if a field has been set.

### GetScopes

`func (o *DaVinciApplicationResponseOAuth) GetScopes() []DaVinciApplicationResponseOAuthScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DaVinciApplicationResponseOAuth) GetScopesOk() (*[]DaVinciApplicationResponseOAuthScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DaVinciApplicationResponseOAuth) SetScopes(v []DaVinciApplicationResponseOAuthScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *DaVinciApplicationResponseOAuth) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetSpJwksOpenid

`func (o *DaVinciApplicationResponseOAuth) GetSpJwksOpenid() string`

GetSpJwksOpenid returns the SpJwksOpenid field if non-nil, zero value otherwise.

### GetSpJwksOpenidOk

`func (o *DaVinciApplicationResponseOAuth) GetSpJwksOpenidOk() (*string, bool)`

GetSpJwksOpenidOk returns a tuple with the SpJwksOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpJwksOpenid

`func (o *DaVinciApplicationResponseOAuth) SetSpJwksOpenid(v string)`

SetSpJwksOpenid sets SpJwksOpenid field to given value.

### HasSpJwksOpenid

`func (o *DaVinciApplicationResponseOAuth) HasSpJwksOpenid() bool`

HasSpJwksOpenid returns a boolean if a field has been set.

### GetSpjwksUrl

`func (o *DaVinciApplicationResponseOAuth) GetSpjwksUrl() string`

GetSpjwksUrl returns the SpjwksUrl field if non-nil, zero value otherwise.

### GetSpjwksUrlOk

`func (o *DaVinciApplicationResponseOAuth) GetSpjwksUrlOk() (*string, bool)`

GetSpjwksUrlOk returns a tuple with the SpjwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpjwksUrl

`func (o *DaVinciApplicationResponseOAuth) SetSpjwksUrl(v string)`

SetSpjwksUrl sets SpjwksUrl field to given value.

### HasSpjwksUrl

`func (o *DaVinciApplicationResponseOAuth) HasSpjwksUrl() bool`

HasSpjwksUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


