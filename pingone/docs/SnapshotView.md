# SnapshotView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**SnapshotViewLinks**](SnapshotViewLinks.md) |  | [optional] 
**CompletedAt** | Pointer to **time.Time** |  | [optional] 
**Configuration** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Environment** | Pointer to [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] [readonly] 
**ImportMetadata** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**ReferencedResources** | Pointer to **[]string** |  | [optional] 
**Resource** | Pointer to [**SnapshotResourceInformation**](SnapshotResourceInformation.md) |  | [optional] 
**ResourceUrl** | Pointer to **string** |  | [optional] 
**SnapshotId** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**StartedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**SnapshotViewStatus**](SnapshotViewStatus.md) |  | [optional] 
**VersionedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewSnapshotView

`func NewSnapshotView() *SnapshotView`

NewSnapshotView instantiates a new SnapshotView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotViewWithDefaults

`func NewSnapshotViewWithDefaults() *SnapshotView`

NewSnapshotViewWithDefaults instantiates a new SnapshotView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *SnapshotView) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *SnapshotView) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *SnapshotView) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *SnapshotView) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *SnapshotView) GetLinks() SnapshotViewLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SnapshotView) GetLinksOk() (*SnapshotViewLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SnapshotView) SetLinks(v SnapshotViewLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SnapshotView) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCompletedAt

`func (o *SnapshotView) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *SnapshotView) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *SnapshotView) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *SnapshotView) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### GetConfiguration

`func (o *SnapshotView) GetConfiguration() string`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SnapshotView) GetConfigurationOk() (*string, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SnapshotView) SetConfiguration(v string)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SnapshotView) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SnapshotView) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SnapshotView) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SnapshotView) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SnapshotView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *SnapshotView) GetEnvironment() ResourceRelationshipReadOnly`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SnapshotView) GetEnvironmentOk() (*ResourceRelationshipReadOnly, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SnapshotView) SetEnvironment(v ResourceRelationshipReadOnly)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SnapshotView) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetError

`func (o *SnapshotView) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SnapshotView) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SnapshotView) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SnapshotView) HasError() bool`

HasError returns a boolean if a field has been set.

### GetId

`func (o *SnapshotView) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SnapshotView) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SnapshotView) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *SnapshotView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImportMetadata

`func (o *SnapshotView) GetImportMetadata() map[string]map[string]interface{}`

GetImportMetadata returns the ImportMetadata field if non-nil, zero value otherwise.

### GetImportMetadataOk

`func (o *SnapshotView) GetImportMetadataOk() (*map[string]map[string]interface{}, bool)`

GetImportMetadataOk returns a tuple with the ImportMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMetadata

`func (o *SnapshotView) SetImportMetadata(v map[string]map[string]interface{})`

SetImportMetadata sets ImportMetadata field to given value.

### HasImportMetadata

`func (o *SnapshotView) HasImportMetadata() bool`

HasImportMetadata returns a boolean if a field has been set.

### GetReferencedResources

`func (o *SnapshotView) GetReferencedResources() []string`

GetReferencedResources returns the ReferencedResources field if non-nil, zero value otherwise.

### GetReferencedResourcesOk

`func (o *SnapshotView) GetReferencedResourcesOk() (*[]string, bool)`

GetReferencedResourcesOk returns a tuple with the ReferencedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferencedResources

`func (o *SnapshotView) SetReferencedResources(v []string)`

SetReferencedResources sets ReferencedResources field to given value.

### HasReferencedResources

`func (o *SnapshotView) HasReferencedResources() bool`

HasReferencedResources returns a boolean if a field has been set.

### GetResource

`func (o *SnapshotView) GetResource() SnapshotResourceInformation`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SnapshotView) GetResourceOk() (*SnapshotResourceInformation, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SnapshotView) SetResource(v SnapshotResourceInformation)`

SetResource sets Resource field to given value.

### HasResource

`func (o *SnapshotView) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceUrl

`func (o *SnapshotView) GetResourceUrl() string`

GetResourceUrl returns the ResourceUrl field if non-nil, zero value otherwise.

### GetResourceUrlOk

`func (o *SnapshotView) GetResourceUrlOk() (*string, bool)`

GetResourceUrlOk returns a tuple with the ResourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceUrl

`func (o *SnapshotView) SetResourceUrl(v string)`

SetResourceUrl sets ResourceUrl field to given value.

### HasResourceUrl

`func (o *SnapshotView) HasResourceUrl() bool`

HasResourceUrl returns a boolean if a field has been set.

### GetSnapshotId

`func (o *SnapshotView) GetSnapshotId() uuid.UUID`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *SnapshotView) GetSnapshotIdOk() (*uuid.UUID, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *SnapshotView) SetSnapshotId(v uuid.UUID)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *SnapshotView) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetStartedAt

`func (o *SnapshotView) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *SnapshotView) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *SnapshotView) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *SnapshotView) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetStatus

`func (o *SnapshotView) GetStatus() SnapshotViewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnapshotView) GetStatusOk() (*SnapshotViewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnapshotView) SetStatus(v SnapshotViewStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnapshotView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVersionedAt

`func (o *SnapshotView) GetVersionedAt() time.Time`

GetVersionedAt returns the VersionedAt field if non-nil, zero value otherwise.

### GetVersionedAtOk

`func (o *SnapshotView) GetVersionedAtOk() (*time.Time, bool)`

GetVersionedAtOk returns a tuple with the VersionedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionedAt

`func (o *SnapshotView) SetVersionedAt(v time.Time)`

SetVersionedAt sets VersionedAt field to given value.

### HasVersionedAt

`func (o *SnapshotView) HasVersionedAt() bool`

HasVersionedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


