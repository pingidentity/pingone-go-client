# EnvironmentBillOfMaterials

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]EnvironmentBillOfMaterialsProduct**](EnvironmentBillOfMaterialsProduct.md) |  | 
**SolutionType** | Pointer to [**EnvironmentBillOfMaterialsSolutionType**](EnvironmentBillOfMaterialsSolutionType.md) |  | [optional] 
**Links** | [**EnvironmentBillOfMaterialsLinks**](EnvironmentBillOfMaterialsLinks.md) |  | 
**CreatedAt** | **time.Time** |  | [readonly] 
**UpdatedAt** | **time.Time** |  | [readonly] 

## Methods

### NewEnvironmentBillOfMaterials

`func NewEnvironmentBillOfMaterials(products []EnvironmentBillOfMaterialsProduct, links EnvironmentBillOfMaterialsLinks, createdAt time.Time, updatedAt time.Time, ) *EnvironmentBillOfMaterials`

NewEnvironmentBillOfMaterials instantiates a new EnvironmentBillOfMaterials object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentBillOfMaterialsWithDefaults

`func NewEnvironmentBillOfMaterialsWithDefaults() *EnvironmentBillOfMaterials`

NewEnvironmentBillOfMaterialsWithDefaults instantiates a new EnvironmentBillOfMaterials object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *EnvironmentBillOfMaterials) GetProducts() []EnvironmentBillOfMaterialsProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EnvironmentBillOfMaterials) GetProductsOk() (*[]EnvironmentBillOfMaterialsProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EnvironmentBillOfMaterials) SetProducts(v []EnvironmentBillOfMaterialsProduct)`

SetProducts sets Products field to given value.


### GetSolutionType

`func (o *EnvironmentBillOfMaterials) GetSolutionType() EnvironmentBillOfMaterialsSolutionType`

GetSolutionType returns the SolutionType field if non-nil, zero value otherwise.

### GetSolutionTypeOk

`func (o *EnvironmentBillOfMaterials) GetSolutionTypeOk() (*EnvironmentBillOfMaterialsSolutionType, bool)`

GetSolutionTypeOk returns a tuple with the SolutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionType

`func (o *EnvironmentBillOfMaterials) SetSolutionType(v EnvironmentBillOfMaterialsSolutionType)`

SetSolutionType sets SolutionType field to given value.

### HasSolutionType

`func (o *EnvironmentBillOfMaterials) HasSolutionType() bool`

HasSolutionType returns a boolean if a field has been set.

### GetLinks

`func (o *EnvironmentBillOfMaterials) GetLinks() EnvironmentBillOfMaterialsLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentBillOfMaterials) GetLinksOk() (*EnvironmentBillOfMaterialsLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentBillOfMaterials) SetLinks(v EnvironmentBillOfMaterialsLinks)`

SetLinks sets Links field to given value.


### GetCreatedAt

`func (o *EnvironmentBillOfMaterials) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentBillOfMaterials) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentBillOfMaterials) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *EnvironmentBillOfMaterials) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentBillOfMaterials) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentBillOfMaterials) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


