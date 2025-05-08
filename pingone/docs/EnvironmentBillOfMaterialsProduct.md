# EnvironmentBillOfMaterialsProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bookmarks** | Pointer to [**[]EnvironmentBillOfMaterialsProductBookmark**](EnvironmentBillOfMaterialsProductBookmark.md) |  | [optional] 
**Console** | Pointer to [**EnvironmentBillOfMaterialsProductConsole**](EnvironmentBillOfMaterialsProductConsole.md) |  | [optional] 
**Deployment** | Pointer to [**EnvironmentBillOfMaterialsProductDeployment**](EnvironmentBillOfMaterialsProductDeployment.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] [readonly] 
**Tags** | Pointer to [**[]EnvironmentBillOfMaterialsProductTags**](EnvironmentBillOfMaterialsProductTags.md) |  | [optional] 
**Type** | [**EnvironmentBillOfMaterialsProductType**](EnvironmentBillOfMaterialsProductType.md) |  | 

## Methods

### NewEnvironmentBillOfMaterialsProduct

`func NewEnvironmentBillOfMaterialsProduct(type_ EnvironmentBillOfMaterialsProductType, ) *EnvironmentBillOfMaterialsProduct`

NewEnvironmentBillOfMaterialsProduct instantiates a new EnvironmentBillOfMaterialsProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentBillOfMaterialsProductWithDefaults

`func NewEnvironmentBillOfMaterialsProductWithDefaults() *EnvironmentBillOfMaterialsProduct`

NewEnvironmentBillOfMaterialsProductWithDefaults instantiates a new EnvironmentBillOfMaterialsProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBookmarks

`func (o *EnvironmentBillOfMaterialsProduct) GetBookmarks() []EnvironmentBillOfMaterialsProductBookmark`

GetBookmarks returns the Bookmarks field if non-nil, zero value otherwise.

### GetBookmarksOk

`func (o *EnvironmentBillOfMaterialsProduct) GetBookmarksOk() (*[]EnvironmentBillOfMaterialsProductBookmark, bool)`

GetBookmarksOk returns a tuple with the Bookmarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookmarks

`func (o *EnvironmentBillOfMaterialsProduct) SetBookmarks(v []EnvironmentBillOfMaterialsProductBookmark)`

SetBookmarks sets Bookmarks field to given value.

### HasBookmarks

`func (o *EnvironmentBillOfMaterialsProduct) HasBookmarks() bool`

HasBookmarks returns a boolean if a field has been set.

### GetConsole

`func (o *EnvironmentBillOfMaterialsProduct) GetConsole() EnvironmentBillOfMaterialsProductConsole`

GetConsole returns the Console field if non-nil, zero value otherwise.

### GetConsoleOk

`func (o *EnvironmentBillOfMaterialsProduct) GetConsoleOk() (*EnvironmentBillOfMaterialsProductConsole, bool)`

GetConsoleOk returns a tuple with the Console field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsole

`func (o *EnvironmentBillOfMaterialsProduct) SetConsole(v EnvironmentBillOfMaterialsProductConsole)`

SetConsole sets Console field to given value.

### HasConsole

`func (o *EnvironmentBillOfMaterialsProduct) HasConsole() bool`

HasConsole returns a boolean if a field has been set.

### GetDeployment

`func (o *EnvironmentBillOfMaterialsProduct) GetDeployment() EnvironmentBillOfMaterialsProductDeployment`

GetDeployment returns the Deployment field if non-nil, zero value otherwise.

### GetDeploymentOk

`func (o *EnvironmentBillOfMaterialsProduct) GetDeploymentOk() (*EnvironmentBillOfMaterialsProductDeployment, bool)`

GetDeploymentOk returns a tuple with the Deployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployment

`func (o *EnvironmentBillOfMaterialsProduct) SetDeployment(v EnvironmentBillOfMaterialsProductDeployment)`

SetDeployment sets Deployment field to given value.

### HasDeployment

`func (o *EnvironmentBillOfMaterialsProduct) HasDeployment() bool`

HasDeployment returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentBillOfMaterialsProduct) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentBillOfMaterialsProduct) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentBillOfMaterialsProduct) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentBillOfMaterialsProduct) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *EnvironmentBillOfMaterialsProduct) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentBillOfMaterialsProduct) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentBillOfMaterialsProduct) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentBillOfMaterialsProduct) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTags

`func (o *EnvironmentBillOfMaterialsProduct) GetTags() []EnvironmentBillOfMaterialsProductTags`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EnvironmentBillOfMaterialsProduct) GetTagsOk() (*[]EnvironmentBillOfMaterialsProductTags, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EnvironmentBillOfMaterialsProduct) SetTags(v []EnvironmentBillOfMaterialsProductTags)`

SetTags sets Tags field to given value.

### HasTags

`func (o *EnvironmentBillOfMaterialsProduct) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetType

`func (o *EnvironmentBillOfMaterialsProduct) GetType() EnvironmentBillOfMaterialsProductType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentBillOfMaterialsProduct) GetTypeOk() (*EnvironmentBillOfMaterialsProductType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentBillOfMaterialsProduct) SetType(v EnvironmentBillOfMaterialsProductType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


