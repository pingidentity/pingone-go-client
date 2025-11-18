# EnvironmentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Region** | [**EnvironmentRegionCode**](EnvironmentRegionCode.md) |  | 
**Type** | [**EnvironmentTypeValue**](EnvironmentTypeValue.md) |  | 
**License** | [**EnvironmentLicense**](EnvironmentLicense.md) |  | 
**BillOfMaterials** | Pointer to [**EnvironmentBillOfMaterials**](EnvironmentBillOfMaterials.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**Subtype** | Pointer to [**EnvironmentCreateRequestSubtype**](EnvironmentCreateRequestSubtype.md) |  | [optional] 

## Methods

### NewEnvironmentCreateRequest

`func NewEnvironmentCreateRequest(name string, region EnvironmentRegionCode, type_ EnvironmentTypeValue, license EnvironmentLicense, ) *EnvironmentCreateRequest`

NewEnvironmentCreateRequest instantiates a new EnvironmentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentCreateRequestWithDefaults

`func NewEnvironmentCreateRequestWithDefaults() *EnvironmentCreateRequest`

NewEnvironmentCreateRequestWithDefaults instantiates a new EnvironmentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *EnvironmentCreateRequest) GetRegion() EnvironmentRegionCode`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnvironmentCreateRequest) GetRegionOk() (*EnvironmentRegionCode, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnvironmentCreateRequest) SetRegion(v EnvironmentRegionCode)`

SetRegion sets Region field to given value.


### GetType

`func (o *EnvironmentCreateRequest) GetType() EnvironmentTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentCreateRequest) GetTypeOk() (*EnvironmentTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentCreateRequest) SetType(v EnvironmentTypeValue)`

SetType sets Type field to given value.


### GetLicense

`func (o *EnvironmentCreateRequest) GetLicense() EnvironmentLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EnvironmentCreateRequest) GetLicenseOk() (*EnvironmentLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *EnvironmentCreateRequest) SetLicense(v EnvironmentLicense)`

SetLicense sets License field to given value.


### GetBillOfMaterials

`func (o *EnvironmentCreateRequest) GetBillOfMaterials() EnvironmentBillOfMaterials`

GetBillOfMaterials returns the BillOfMaterials field if non-nil, zero value otherwise.

### GetBillOfMaterialsOk

`func (o *EnvironmentCreateRequest) GetBillOfMaterialsOk() (*EnvironmentBillOfMaterials, bool)`

GetBillOfMaterialsOk returns a tuple with the BillOfMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfMaterials

`func (o *EnvironmentCreateRequest) SetBillOfMaterials(v EnvironmentBillOfMaterials)`

SetBillOfMaterials sets BillOfMaterials field to given value.

### HasBillOfMaterials

`func (o *EnvironmentCreateRequest) HasBillOfMaterials() bool`

HasBillOfMaterials returns a boolean if a field has been set.

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

### GetSubtype

`func (o *EnvironmentCreateRequest) GetSubtype() EnvironmentCreateRequestSubtype`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *EnvironmentCreateRequest) GetSubtypeOk() (*EnvironmentCreateRequestSubtype, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *EnvironmentCreateRequest) SetSubtype(v EnvironmentCreateRequestSubtype)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *EnvironmentCreateRequest) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


