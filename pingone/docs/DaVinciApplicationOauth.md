# DaVinciApplicationOauth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **string** |  | 
**EnforceSignedRequestOpenid** | Pointer to **bool** |  | [optional] 
**GrantTypes** | [**[]DaVinciApplicationOauthGrantTypes**](DaVinciApplicationOauthGrantTypes.md) |  | 
**LogoutUris** | **[]string** |  | 
**RedirectUris** | **[]string** |  | 
**Scopes** | [**[]DaVinciApplicationOauthScopes**](DaVinciApplicationOauthScopes.md) |  | 
**SpJwksOpenid** | Pointer to **string** |  | [optional] 
**SpjwksUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewDaVinciApplicationOauth

`func NewDaVinciApplicationOauth(clientSecret string, grantTypes []DaVinciApplicationOauthGrantTypes, logoutUris []string, redirectUris []string, scopes []DaVinciApplicationOauthScopes, ) *DaVinciApplicationOauth`

NewDaVinciApplicationOauth instantiates a new DaVinciApplicationOauth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciApplicationOauthWithDefaults

`func NewDaVinciApplicationOauthWithDefaults() *DaVinciApplicationOauth`

NewDaVinciApplicationOauthWithDefaults instantiates a new DaVinciApplicationOauth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *DaVinciApplicationOauth) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *DaVinciApplicationOauth) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *DaVinciApplicationOauth) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationOauth) GetEnforceSignedRequestOpenid() bool`

GetEnforceSignedRequestOpenid returns the EnforceSignedRequestOpenid field if non-nil, zero value otherwise.

### GetEnforceSignedRequestOpenidOk

`func (o *DaVinciApplicationOauth) GetEnforceSignedRequestOpenidOk() (*bool, bool)`

GetEnforceSignedRequestOpenidOk returns a tuple with the EnforceSignedRequestOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceSignedRequestOpenid

`func (o *DaVinciApplicationOauth) SetEnforceSignedRequestOpenid(v bool)`

SetEnforceSignedRequestOpenid sets EnforceSignedRequestOpenid field to given value.

### HasEnforceSignedRequestOpenid

`func (o *DaVinciApplicationOauth) HasEnforceSignedRequestOpenid() bool`

HasEnforceSignedRequestOpenid returns a boolean if a field has been set.

### GetGrantTypes

`func (o *DaVinciApplicationOauth) GetGrantTypes() []DaVinciApplicationOauthGrantTypes`

GetGrantTypes returns the GrantTypes field if non-nil, zero value otherwise.

### GetGrantTypesOk

`func (o *DaVinciApplicationOauth) GetGrantTypesOk() (*[]DaVinciApplicationOauthGrantTypes, bool)`

GetGrantTypesOk returns a tuple with the GrantTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantTypes

`func (o *DaVinciApplicationOauth) SetGrantTypes(v []DaVinciApplicationOauthGrantTypes)`

SetGrantTypes sets GrantTypes field to given value.


### GetLogoutUris

`func (o *DaVinciApplicationOauth) GetLogoutUris() []string`

GetLogoutUris returns the LogoutUris field if non-nil, zero value otherwise.

### GetLogoutUrisOk

`func (o *DaVinciApplicationOauth) GetLogoutUrisOk() (*[]string, bool)`

GetLogoutUrisOk returns a tuple with the LogoutUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUris

`func (o *DaVinciApplicationOauth) SetLogoutUris(v []string)`

SetLogoutUris sets LogoutUris field to given value.


### GetRedirectUris

`func (o *DaVinciApplicationOauth) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *DaVinciApplicationOauth) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *DaVinciApplicationOauth) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.


### GetScopes

`func (o *DaVinciApplicationOauth) GetScopes() []DaVinciApplicationOauthScopes`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *DaVinciApplicationOauth) GetScopesOk() (*[]DaVinciApplicationOauthScopes, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *DaVinciApplicationOauth) SetScopes(v []DaVinciApplicationOauthScopes)`

SetScopes sets Scopes field to given value.


### GetSpJwksOpenid

`func (o *DaVinciApplicationOauth) GetSpJwksOpenid() string`

GetSpJwksOpenid returns the SpJwksOpenid field if non-nil, zero value otherwise.

### GetSpJwksOpenidOk

`func (o *DaVinciApplicationOauth) GetSpJwksOpenidOk() (*string, bool)`

GetSpJwksOpenidOk returns a tuple with the SpJwksOpenid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpJwksOpenid

`func (o *DaVinciApplicationOauth) SetSpJwksOpenid(v string)`

SetSpJwksOpenid sets SpJwksOpenid field to given value.

### HasSpJwksOpenid

`func (o *DaVinciApplicationOauth) HasSpJwksOpenid() bool`

HasSpJwksOpenid returns a boolean if a field has been set.

### GetSpjwksUrl

`func (o *DaVinciApplicationOauth) GetSpjwksUrl() string`

GetSpjwksUrl returns the SpjwksUrl field if non-nil, zero value otherwise.

### GetSpjwksUrlOk

`func (o *DaVinciApplicationOauth) GetSpjwksUrlOk() (*string, bool)`

GetSpjwksUrlOk returns a tuple with the SpjwksUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpjwksUrl

`func (o *DaVinciApplicationOauth) SetSpjwksUrl(v string)`

SetSpjwksUrl sets SpjwksUrl field to given value.

### HasSpjwksUrl

`func (o *DaVinciApplicationOauth) HasSpjwksUrl() bool`

HasSpjwksUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


