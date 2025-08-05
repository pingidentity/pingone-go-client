# DaVinciApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciApplicationResponseLinks**](DaVinciApplicationResponseLinks.md) |  | 
**ApiKey** | [**DaVinciApplicationResponseApiKey**](DaVinciApplicationResponseApiKey.md) |  | 
**Environment** | [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Oauth** | [**DaVinciApplicationResponseOAuth**](DaVinciApplicationResponseOAuth.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDaVinciApplicationResponse

`func NewDaVinciApplicationResponse(links DaVinciApplicationResponseLinks, apiKey DaVinciApplicationResponseApiKey, environment ResourceRelationshipReadOnly, id string, name string, oauth DaVinciApplicationResponseOAuth, ) *DaVinciApplicationResponse`

NewDaVinciApplicationResponse instantiates a new DaVinciApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciApplicationResponseWithDefaults

`func NewDaVinciApplicationResponseWithDefaults() *DaVinciApplicationResponse`

NewDaVinciApplicationResponseWithDefaults instantiates a new DaVinciApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciApplicationResponse) GetLinks() DaVinciApplicationResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciApplicationResponse) GetLinksOk() (*DaVinciApplicationResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciApplicationResponse) SetLinks(v DaVinciApplicationResponseLinks)`

SetLinks sets Links field to given value.


### GetApiKey

`func (o *DaVinciApplicationResponse) GetApiKey() DaVinciApplicationResponseApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *DaVinciApplicationResponse) GetApiKeyOk() (*DaVinciApplicationResponseApiKey, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *DaVinciApplicationResponse) SetApiKey(v DaVinciApplicationResponseApiKey)`

SetApiKey sets ApiKey field to given value.


### GetEnvironment

`func (o *DaVinciApplicationResponse) GetEnvironment() ResourceRelationshipReadOnly`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciApplicationResponse) GetEnvironmentOk() (*ResourceRelationshipReadOnly, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciApplicationResponse) SetEnvironment(v ResourceRelationshipReadOnly)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *DaVinciApplicationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciApplicationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciApplicationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DaVinciApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciApplicationResponse) SetName(v string)`

SetName sets Name field to given value.


### GetOauth

`func (o *DaVinciApplicationResponse) GetOauth() DaVinciApplicationResponseOAuth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *DaVinciApplicationResponse) GetOauthOk() (*DaVinciApplicationResponseOAuth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *DaVinciApplicationResponse) SetOauth(v DaVinciApplicationResponseOAuth)`

SetOauth sets Oauth field to given value.


### GetCreatedAt

`func (o *DaVinciApplicationResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciApplicationResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciApplicationResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciApplicationResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DaVinciApplicationResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciApplicationResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciApplicationResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciApplicationResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


