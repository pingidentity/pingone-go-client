# EnvironmentBillOfMaterialsCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | [**[]EnvironmentBillOfMaterialsProduct**](EnvironmentBillOfMaterialsProduct.md) |  | 
**SolutionType** | Pointer to [**EnvironmentBillOfMaterialsSolutionType**](EnvironmentBillOfMaterialsSolutionType.md) |  | [optional] 

## Methods

### NewEnvironmentBillOfMaterialsCreateRequest

`func NewEnvironmentBillOfMaterialsCreateRequest(products []EnvironmentBillOfMaterialsProduct, ) *EnvironmentBillOfMaterialsCreateRequest`

NewEnvironmentBillOfMaterialsCreateRequest instantiates a new EnvironmentBillOfMaterialsCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentBillOfMaterialsCreateRequestWithDefaults

`func NewEnvironmentBillOfMaterialsCreateRequestWithDefaults() *EnvironmentBillOfMaterialsCreateRequest`

NewEnvironmentBillOfMaterialsCreateRequestWithDefaults instantiates a new EnvironmentBillOfMaterialsCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *EnvironmentBillOfMaterialsCreateRequest) GetProducts() []EnvironmentBillOfMaterialsProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EnvironmentBillOfMaterialsCreateRequest) GetProductsOk() (*[]EnvironmentBillOfMaterialsProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EnvironmentBillOfMaterialsCreateRequest) SetProducts(v []EnvironmentBillOfMaterialsProduct)`

SetProducts sets Products field to given value.


### GetSolutionType

`func (o *EnvironmentBillOfMaterialsCreateRequest) GetSolutionType() EnvironmentBillOfMaterialsSolutionType`

GetSolutionType returns the SolutionType field if non-nil, zero value otherwise.

### GetSolutionTypeOk

`func (o *EnvironmentBillOfMaterialsCreateRequest) GetSolutionTypeOk() (*EnvironmentBillOfMaterialsSolutionType, bool)`

GetSolutionTypeOk returns a tuple with the SolutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionType

`func (o *EnvironmentBillOfMaterialsCreateRequest) SetSolutionType(v EnvironmentBillOfMaterialsSolutionType)`

SetSolutionType sets SolutionType field to given value.

### HasSolutionType

`func (o *EnvironmentBillOfMaterialsCreateRequest) HasSolutionType() bool`

HasSolutionType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


