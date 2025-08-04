# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Region** | [**EnvironmentRegion**](EnvironmentRegion.md) |  | 
**Type** | [**EnvironmentType**](EnvironmentType.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Id** | [**uuid.UUID**](uuid.UUID.md) |  | 
**Organization** | [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**EnvironmentLinks**](EnvironmentLinks.md) |  | [optional] 
**BillOfMaterials** | Pointer to [**EnvironmentBillOfMaterials**](EnvironmentBillOfMaterials.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentCapabilities** | Pointer to [**EnvironmentCapabilities**](EnvironmentCapabilities.md) |  | [optional] 
**HardDeleteAllowedAt** | Pointer to **time.Time** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**License** | Pointer to [**EnvironmentLicense**](EnvironmentLicense.md) |  | [optional] 
**PingoneAccountId** | Pointer to **string** |  | [optional] 
**SoftDeletedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**EnvironmentStatus**](EnvironmentStatus.md) |  | [optional] 

## Methods

### NewEnvironment

`func NewEnvironment(name string, region EnvironmentRegion, type_ EnvironmentType, createdAt time.Time, updatedAt time.Time, id uuid.UUID, organization ResourceRelationshipReadOnly, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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


### GetOrganization

`func (o *Environment) GetOrganization() ResourceRelationshipReadOnly`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Environment) GetOrganizationOk() (*ResourceRelationshipReadOnly, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Environment) SetOrganization(v ResourceRelationshipReadOnly)`

SetOrganization sets Organization field to given value.


### GetEmbedded

`func (o *Environment) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *Environment) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *Environment) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *Environment) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

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

### GetEnvironmentCapabilities

`func (o *Environment) GetEnvironmentCapabilities() EnvironmentCapabilities`

GetEnvironmentCapabilities returns the EnvironmentCapabilities field if non-nil, zero value otherwise.

### GetEnvironmentCapabilitiesOk

`func (o *Environment) GetEnvironmentCapabilitiesOk() (*EnvironmentCapabilities, bool)`

GetEnvironmentCapabilitiesOk returns a tuple with the EnvironmentCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCapabilities

`func (o *Environment) SetEnvironmentCapabilities(v EnvironmentCapabilities)`

SetEnvironmentCapabilities sets EnvironmentCapabilities field to given value.

### HasEnvironmentCapabilities

`func (o *Environment) HasEnvironmentCapabilities() bool`

HasEnvironmentCapabilities returns a boolean if a field has been set.

### GetHardDeleteAllowedAt

`func (o *Environment) GetHardDeleteAllowedAt() time.Time`

GetHardDeleteAllowedAt returns the HardDeleteAllowedAt field if non-nil, zero value otherwise.

### GetHardDeleteAllowedAtOk

`func (o *Environment) GetHardDeleteAllowedAtOk() (*time.Time, bool)`

GetHardDeleteAllowedAtOk returns a tuple with the HardDeleteAllowedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDeleteAllowedAt

`func (o *Environment) SetHardDeleteAllowedAt(v time.Time)`

SetHardDeleteAllowedAt sets HardDeleteAllowedAt field to given value.

### HasHardDeleteAllowedAt

`func (o *Environment) HasHardDeleteAllowedAt() bool`

HasHardDeleteAllowedAt returns a boolean if a field has been set.

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

`func (o *Environment) GetLicense() EnvironmentLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Environment) GetLicenseOk() (*EnvironmentLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Environment) SetLicense(v EnvironmentLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Environment) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetPingoneAccountId

`func (o *Environment) GetPingoneAccountId() string`

GetPingoneAccountId returns the PingoneAccountId field if non-nil, zero value otherwise.

### GetPingoneAccountIdOk

`func (o *Environment) GetPingoneAccountIdOk() (*string, bool)`

GetPingoneAccountIdOk returns a tuple with the PingoneAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingoneAccountId

`func (o *Environment) SetPingoneAccountId(v string)`

SetPingoneAccountId sets PingoneAccountId field to given value.

### HasPingoneAccountId

`func (o *Environment) HasPingoneAccountId() bool`

HasPingoneAccountId returns a boolean if a field has been set.

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

### GetStatus

`func (o *Environment) GetStatus() EnvironmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Environment) GetStatusOk() (*EnvironmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Environment) SetStatus(v EnvironmentStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Environment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


