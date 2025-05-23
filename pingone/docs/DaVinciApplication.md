# DaVinciApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**DaVinciApplicationLinks**](DaVinciApplicationLinks.md) |  | 
**ApiKey** | [**DaVinciApplicationApiKey**](DaVinciApplicationApiKey.md) |  | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Oauth** | [**DaVinciApplicationOauth**](DaVinciApplicationOauth.md) |  | 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDaVinciApplication

`func NewDaVinciApplication(links DaVinciApplicationLinks, apiKey DaVinciApplicationApiKey, environment ResourceRelationshipPingOne, id string, name string, oauth DaVinciApplicationOauth, ) *DaVinciApplication`

NewDaVinciApplication instantiates a new DaVinciApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDaVinciApplicationWithDefaults

`func NewDaVinciApplicationWithDefaults() *DaVinciApplication`

NewDaVinciApplicationWithDefaults instantiates a new DaVinciApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *DaVinciApplication) GetLinks() DaVinciApplicationLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DaVinciApplication) GetLinksOk() (*DaVinciApplicationLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DaVinciApplication) SetLinks(v DaVinciApplicationLinks)`

SetLinks sets Links field to given value.


### GetApiKey

`func (o *DaVinciApplication) GetApiKey() DaVinciApplicationApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *DaVinciApplication) GetApiKeyOk() (*DaVinciApplicationApiKey, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *DaVinciApplication) SetApiKey(v DaVinciApplicationApiKey)`

SetApiKey sets ApiKey field to given value.


### GetCreatedAt

`func (o *DaVinciApplication) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DaVinciApplication) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DaVinciApplication) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DaVinciApplication) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *DaVinciApplication) GetEnvironment() ResourceRelationshipPingOne`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DaVinciApplication) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DaVinciApplication) SetEnvironment(v ResourceRelationshipPingOne)`

SetEnvironment sets Environment field to given value.


### GetId

`func (o *DaVinciApplication) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DaVinciApplication) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DaVinciApplication) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DaVinciApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DaVinciApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DaVinciApplication) SetName(v string)`

SetName sets Name field to given value.


### GetOauth

`func (o *DaVinciApplication) GetOauth() DaVinciApplicationOauth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *DaVinciApplication) GetOauthOk() (*DaVinciApplicationOauth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *DaVinciApplication) SetOauth(v DaVinciApplicationOauth)`

SetOauth sets Oauth field to given value.


### GetUpdatedAt

`func (o *DaVinciApplication) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DaVinciApplication) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DaVinciApplication) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DaVinciApplication) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


