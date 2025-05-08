# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**License** | [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | 
**Name** | **string** |  | 
**Type** | [**EnvironmentType**](EnvironmentType.md) |  | 
**Status** | Pointer to [**EnvironmentStatusValue**](EnvironmentStatusValue.md) |  | [optional] 
**Links** | Pointer to [**EnvironmentLinks**](EnvironmentLinks.md) |  | [optional] 
**BillOfMaterials** | Pointer to [**EnvironmentBillOfMaterials**](EnvironmentBillOfMaterials.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**HardDeletedAllowedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] [readonly] 
**Organization** | Pointer to [**ResourceRelationshipPingOne**](ResourceRelationshipPingOne.md) |  | [optional] 
**Region** | Pointer to [**EnvironmentRegion**](EnvironmentRegion.md) |  | [optional] 
**SoftDeletedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewEnvironment

`func NewEnvironment(license ResourceRelationshipPingOne, name string, type_ EnvironmentType, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *Environment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Environment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Environment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Environment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIcon

`func (o *Environment) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *Environment) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *Environment) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *Environment) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLicense

`func (o *Environment) GetLicense() ResourceRelationshipPingOne`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Environment) GetLicenseOk() (*ResourceRelationshipPingOne, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Environment) SetLicense(v ResourceRelationshipPingOne)`

SetLicense sets License field to given value.


### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *Environment) GetType() EnvironmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Environment) GetTypeOk() (*EnvironmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Environment) SetType(v EnvironmentType)`

SetType sets Type field to given value.


### GetStatus

`func (o *Environment) GetStatus() EnvironmentStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Environment) GetStatusOk() (*EnvironmentStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Environment) SetStatus(v EnvironmentStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Environment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLinks

`func (o *Environment) GetLinks() EnvironmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Environment) GetLinksOk() (*EnvironmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Environment) SetLinks(v EnvironmentLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Environment) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetBillOfMaterials

`func (o *Environment) GetBillOfMaterials() EnvironmentBillOfMaterials`

GetBillOfMaterials returns the BillOfMaterials field if non-nil, zero value otherwise.

### GetBillOfMaterialsOk

`func (o *Environment) GetBillOfMaterialsOk() (*EnvironmentBillOfMaterials, bool)`

GetBillOfMaterialsOk returns a tuple with the BillOfMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfMaterials

`func (o *Environment) SetBillOfMaterials(v EnvironmentBillOfMaterials)`

SetBillOfMaterials sets BillOfMaterials field to given value.

### HasBillOfMaterials

`func (o *Environment) HasBillOfMaterials() bool`

HasBillOfMaterials returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Environment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Environment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Environment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Environment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHardDeletedAllowedAt

`func (o *Environment) GetHardDeletedAllowedAt() time.Time`

GetHardDeletedAllowedAt returns the HardDeletedAllowedAt field if non-nil, zero value otherwise.

### GetHardDeletedAllowedAtOk

`func (o *Environment) GetHardDeletedAllowedAtOk() (*time.Time, bool)`

GetHardDeletedAllowedAtOk returns a tuple with the HardDeletedAllowedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDeletedAllowedAt

`func (o *Environment) SetHardDeletedAllowedAt(v time.Time)`

SetHardDeletedAllowedAt sets HardDeletedAllowedAt field to given value.

### HasHardDeletedAllowedAt

`func (o *Environment) HasHardDeletedAllowedAt() bool`

HasHardDeletedAllowedAt returns a boolean if a field has been set.

### GetId

`func (o *Environment) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *Environment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrganization

`func (o *Environment) GetOrganization() ResourceRelationshipPingOne`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Environment) GetOrganizationOk() (*ResourceRelationshipPingOne, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Environment) SetOrganization(v ResourceRelationshipPingOne)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Environment) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRegion

`func (o *Environment) GetRegion() EnvironmentRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Environment) GetRegionOk() (*EnvironmentRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Environment) SetRegion(v EnvironmentRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Environment) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSoftDeletedAt

`func (o *Environment) GetSoftDeletedAt() time.Time`

GetSoftDeletedAt returns the SoftDeletedAt field if non-nil, zero value otherwise.

### GetSoftDeletedAtOk

`func (o *Environment) GetSoftDeletedAtOk() (*time.Time, bool)`

GetSoftDeletedAtOk returns a tuple with the SoftDeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeletedAt

`func (o *Environment) SetSoftDeletedAt(v time.Time)`

SetSoftDeletedAt sets SoftDeletedAt field to given value.

### HasSoftDeletedAt

`func (o *Environment) HasSoftDeletedAt() bool`

HasSoftDeletedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Environment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Environment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Environment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Environment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


