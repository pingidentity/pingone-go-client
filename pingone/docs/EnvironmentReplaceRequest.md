# EnvironmentReplaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillOfMaterials** | Pointer to [**BillOfMaterials**](BillOfMaterials.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**License** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**Name** | **string** |  | 
**Type** | [**EnvironmentType**](EnvironmentType.md) |  | 

## Methods

### NewEnvironmentReplaceRequest

`func NewEnvironmentReplaceRequest(license ResourceRelationshipPingOne, name string, type_ EnvironmentType, ) *EnvironmentReplaceRequest`

NewEnvironmentReplaceRequest instantiates a new EnvironmentReplaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentReplaceRequestWithDefaults

`func NewEnvironmentReplaceRequestWithDefaults() *EnvironmentReplaceRequest`

NewEnvironmentReplaceRequestWithDefaults instantiates a new EnvironmentReplaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillOfMaterials

`func (o *EnvironmentReplaceRequest) GetBillOfMaterials() BillOfMaterials`

GetBillOfMaterials returns the BillOfMaterials field if non-nil, zero value otherwise.

### GetBillOfMaterialsOk

`func (o *EnvironmentReplaceRequest) GetBillOfMaterialsOk() (*BillOfMaterials, bool)`

GetBillOfMaterialsOk returns a tuple with the BillOfMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfMaterials

`func (o *EnvironmentReplaceRequest) SetBillOfMaterials(v BillOfMaterials)`

SetBillOfMaterials sets BillOfMaterials field to given value.

### HasBillOfMaterials

`func (o *EnvironmentReplaceRequest) HasBillOfMaterials() bool`

HasBillOfMaterials returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentReplaceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentReplaceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentReplaceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentReplaceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *EnvironmentReplaceRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EnvironmentReplaceRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EnvironmentReplaceRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EnvironmentReplaceRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLicense

`func (o *EnvironmentReplaceRequest) GetLicense() ResourceRelationshipPingOne`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EnvironmentReplaceRequest) GetLicenseOk() (*ResourceRelationshipPingOne, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *EnvironmentReplaceRequest) SetLicense(v ResourceRelationshipPingOne)`

SetLicense sets License field to given value.


### GetName

`func (o *EnvironmentReplaceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentReplaceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentReplaceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *EnvironmentReplaceRequest) GetType() EnvironmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentReplaceRequest) GetTypeOk() (*EnvironmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentReplaceRequest) SetType(v EnvironmentType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


