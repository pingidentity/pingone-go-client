# EnvironmentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**License** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**Name** | **string** |  | 
**Region** | [**EnvironmentRegion**](EnvironmentRegion.md) |  | 
**Type** | [**EnvironmentType**](EnvironmentType.md) |  | 
**BillOfMaterials** | Pointer to [**EnvironmentBillOfMaterialsCreateRequest**](EnvironmentBillOfMaterialsCreateRequest.md) |  | [optional] 

## Methods

### NewEnvironmentCreateRequest

`func NewEnvironmentCreateRequest(license ResourceRelationshipPingOne, name string, region EnvironmentRegion, type_ EnvironmentType, ) *EnvironmentCreateRequest`

NewEnvironmentCreateRequest instantiates a new EnvironmentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCreateRequestWithDefaults

`func NewEnvironmentCreateRequestWithDefaults() *EnvironmentCreateRequest`

NewEnvironmentCreateRequestWithDefaults instantiates a new EnvironmentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EnvironmentCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *EnvironmentCreateRequest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EnvironmentCreateRequest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EnvironmentCreateRequest) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EnvironmentCreateRequest) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLicense

`func (o *EnvironmentCreateRequest) GetLicense() ResourceRelationshipPingOne`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EnvironmentCreateRequest) GetLicenseOk() (*ResourceRelationshipPingOne, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *EnvironmentCreateRequest) SetLicense(v ResourceRelationshipPingOne)`

SetLicense sets License field to given value.


### GetName

`func (o *EnvironmentCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *EnvironmentCreateRequest) GetRegion() EnvironmentRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnvironmentCreateRequest) GetRegionOk() (*EnvironmentRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnvironmentCreateRequest) SetRegion(v EnvironmentRegion)`

SetRegion sets Region field to given value.


### GetType

`func (o *EnvironmentCreateRequest) GetType() EnvironmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentCreateRequest) GetTypeOk() (*EnvironmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentCreateRequest) SetType(v EnvironmentType)`

SetType sets Type field to given value.


### GetBillOfMaterials

`func (o *EnvironmentCreateRequest) GetBillOfMaterials() EnvironmentBillOfMaterialsCreateRequest`

GetBillOfMaterials returns the BillOfMaterials field if non-nil, zero value otherwise.

### GetBillOfMaterialsOk

`func (o *EnvironmentCreateRequest) GetBillOfMaterialsOk() (*EnvironmentBillOfMaterialsCreateRequest, bool)`

GetBillOfMaterialsOk returns a tuple with the BillOfMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfMaterials

`func (o *EnvironmentCreateRequest) SetBillOfMaterials(v EnvironmentBillOfMaterialsCreateRequest)`

SetBillOfMaterials sets BillOfMaterials field to given value.

### HasBillOfMaterials

`func (o *EnvironmentCreateRequest) HasBillOfMaterials() bool`

HasBillOfMaterials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


