# EnvironmentLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Package** | Pointer to **string** |  | [optional] 

## Methods

### NewEnvironmentLicense

`func NewEnvironmentLicense() *EnvironmentLicense`

NewEnvironmentLicense instantiates a new EnvironmentLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentLicenseWithDefaults

`func NewEnvironmentLicenseWithDefaults() *EnvironmentLicense`

NewEnvironmentLicenseWithDefaults instantiates a new EnvironmentLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EnvironmentLicense) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnvironmentLicense) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnvironmentLicense) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *EnvironmentLicense) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPackage

`func (o *EnvironmentLicense) GetPackage() string`

GetPackage returns the Package field if non-nil, zero value otherwise.

### GetPackageOk

`func (o *EnvironmentLicense) GetPackageOk() (*string, bool)`

GetPackageOk returns a tuple with the Package field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackage

`func (o *EnvironmentLicense) SetPackage(v string)`

SetPackage sets Package field to given value.

### HasPackage

`func (o *EnvironmentLicense) HasPackage() bool`

HasPackage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


