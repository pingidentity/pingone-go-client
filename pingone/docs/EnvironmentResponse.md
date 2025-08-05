# EnvironmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Region** | [**EnvironmentRegionCode**](EnvironmentRegionCode.md) |  | 
**Type** | [**EnvironmentTypeValue**](EnvironmentTypeValue.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 
**Id** | [**uuid.UUID**](uuid.UUID.md) |  | 
**Organization** | [**ResourceRelationshipReadOnly**](ResourceRelationshipReadOnly.md) |  | 
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**EnvironmentResponseLinks**](EnvironmentResponseLinks.md) |  | [optional] 
**BillOfMaterials** | Pointer to [**EnvironmentBillOfMaterials**](EnvironmentBillOfMaterials.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**EnvironmentCapabilities** | Pointer to [**EnvironmentCapabilities**](EnvironmentCapabilities.md) |  | [optional] 
**HardDeleteAllowedAt** | Pointer to **time.Time** |  | [optional] 
**Icon** | Pointer to **string** |  | [optional] 
**License** | Pointer to [**EnvironmentLicense**](EnvironmentLicense.md) |  | [optional] 
**PingoneAccountId** | Pointer to **string** |  | [optional] 
**SoftDeletedAt** | Pointer to **time.Time** |  | [optional] 
**Status** | Pointer to [**EnvironmentStatusValue**](EnvironmentStatusValue.md) |  | [optional] 

## Methods

### NewEnvironmentResponse

`func NewEnvironmentResponse(name string, region EnvironmentRegionCode, type_ EnvironmentTypeValue, createdAt time.Time, updatedAt time.Time, id uuid.UUID, organization ResourceRelationshipReadOnly, ) *EnvironmentResponse`

NewEnvironmentResponse instantiates a new EnvironmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentResponseWithDefaults

`func NewEnvironmentResponseWithDefaults() *EnvironmentResponse`

NewEnvironmentResponseWithDefaults instantiates a new EnvironmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EnvironmentResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EnvironmentResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EnvironmentResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRegion

`func (o *EnvironmentResponse) GetRegion() EnvironmentRegionCode`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *EnvironmentResponse) GetRegionOk() (*EnvironmentRegionCode, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *EnvironmentResponse) SetRegion(v EnvironmentRegionCode)`

SetRegion sets Region field to given value.


### GetType

`func (o *EnvironmentResponse) GetType() EnvironmentTypeValue`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EnvironmentResponse) GetTypeOk() (*EnvironmentTypeValue, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EnvironmentResponse) SetType(v EnvironmentTypeValue)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *EnvironmentResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnvironmentResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetId

`func (o *EnvironmentResponse) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentResponse) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentResponse) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetOrganization

`func (o *EnvironmentResponse) GetOrganization() ResourceRelationshipReadOnly`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *EnvironmentResponse) GetOrganizationOk() (*ResourceRelationshipReadOnly, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *EnvironmentResponse) SetOrganization(v ResourceRelationshipReadOnly)`

SetOrganization sets Organization field to given value.


### GetEmbedded

`func (o *EnvironmentResponse) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *EnvironmentResponse) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *EnvironmentResponse) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *EnvironmentResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *EnvironmentResponse) GetLinks() EnvironmentResponseLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentResponse) GetLinksOk() (*EnvironmentResponseLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentResponse) SetLinks(v EnvironmentResponseLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetBillOfMaterials

`func (o *EnvironmentResponse) GetBillOfMaterials() EnvironmentBillOfMaterials`

GetBillOfMaterials returns the BillOfMaterials field if non-nil, zero value otherwise.

### GetBillOfMaterialsOk

`func (o *EnvironmentResponse) GetBillOfMaterialsOk() (*EnvironmentBillOfMaterials, bool)`

GetBillOfMaterialsOk returns a tuple with the BillOfMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfMaterials

`func (o *EnvironmentResponse) SetBillOfMaterials(v EnvironmentBillOfMaterials)`

SetBillOfMaterials sets BillOfMaterials field to given value.

### HasBillOfMaterials

`func (o *EnvironmentResponse) HasBillOfMaterials() bool`

HasBillOfMaterials returns a boolean if a field has been set.

### GetDescription

`func (o *EnvironmentResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EnvironmentResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EnvironmentResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EnvironmentResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironmentCapabilities

`func (o *EnvironmentResponse) GetEnvironmentCapabilities() EnvironmentCapabilities`

GetEnvironmentCapabilities returns the EnvironmentCapabilities field if non-nil, zero value otherwise.

### GetEnvironmentCapabilitiesOk

`func (o *EnvironmentResponse) GetEnvironmentCapabilitiesOk() (*EnvironmentCapabilities, bool)`

GetEnvironmentCapabilitiesOk returns a tuple with the EnvironmentCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentCapabilities

`func (o *EnvironmentResponse) SetEnvironmentCapabilities(v EnvironmentCapabilities)`

SetEnvironmentCapabilities sets EnvironmentCapabilities field to given value.

### HasEnvironmentCapabilities

`func (o *EnvironmentResponse) HasEnvironmentCapabilities() bool`

HasEnvironmentCapabilities returns a boolean if a field has been set.

### GetHardDeleteAllowedAt

`func (o *EnvironmentResponse) GetHardDeleteAllowedAt() time.Time`

GetHardDeleteAllowedAt returns the HardDeleteAllowedAt field if non-nil, zero value otherwise.

### GetHardDeleteAllowedAtOk

`func (o *EnvironmentResponse) GetHardDeleteAllowedAtOk() (*time.Time, bool)`

GetHardDeleteAllowedAtOk returns a tuple with the HardDeleteAllowedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDeleteAllowedAt

`func (o *EnvironmentResponse) SetHardDeleteAllowedAt(v time.Time)`

SetHardDeleteAllowedAt sets HardDeleteAllowedAt field to given value.

### HasHardDeleteAllowedAt

`func (o *EnvironmentResponse) HasHardDeleteAllowedAt() bool`

HasHardDeleteAllowedAt returns a boolean if a field has been set.

### GetIcon

`func (o *EnvironmentResponse) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EnvironmentResponse) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EnvironmentResponse) SetIcon(v string)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *EnvironmentResponse) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetLicense

`func (o *EnvironmentResponse) GetLicense() EnvironmentLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *EnvironmentResponse) GetLicenseOk() (*EnvironmentLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *EnvironmentResponse) SetLicense(v EnvironmentLicense)`

SetLicense sets License field to given value.

### HasLicense

`func (o *EnvironmentResponse) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetPingoneAccountId

`func (o *EnvironmentResponse) GetPingoneAccountId() string`

GetPingoneAccountId returns the PingoneAccountId field if non-nil, zero value otherwise.

### GetPingoneAccountIdOk

`func (o *EnvironmentResponse) GetPingoneAccountIdOk() (*string, bool)`

GetPingoneAccountIdOk returns a tuple with the PingoneAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPingoneAccountId

`func (o *EnvironmentResponse) SetPingoneAccountId(v string)`

SetPingoneAccountId sets PingoneAccountId field to given value.

### HasPingoneAccountId

`func (o *EnvironmentResponse) HasPingoneAccountId() bool`

HasPingoneAccountId returns a boolean if a field has been set.

### GetSoftDeletedAt

`func (o *EnvironmentResponse) GetSoftDeletedAt() time.Time`

GetSoftDeletedAt returns the SoftDeletedAt field if non-nil, zero value otherwise.

### GetSoftDeletedAtOk

`func (o *EnvironmentResponse) GetSoftDeletedAtOk() (*time.Time, bool)`

GetSoftDeletedAtOk returns a tuple with the SoftDeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftDeletedAt

`func (o *EnvironmentResponse) SetSoftDeletedAt(v time.Time)`

SetSoftDeletedAt sets SoftDeletedAt field to given value.

### HasSoftDeletedAt

`func (o *EnvironmentResponse) HasSoftDeletedAt() bool`

HasSoftDeletedAt returns a boolean if a field has been set.

### GetStatus

`func (o *EnvironmentResponse) GetStatus() EnvironmentStatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnvironmentResponse) GetStatusOk() (*EnvironmentStatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnvironmentResponse) SetStatus(v EnvironmentStatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EnvironmentResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


