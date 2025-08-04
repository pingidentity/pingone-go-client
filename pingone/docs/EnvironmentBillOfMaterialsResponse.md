# EnvironmentBillOfMaterialsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to **map[string]map[string]interface{}** |  | [optional] [readonly] 
**Links** | Pointer to [**map[string]JSONHALLink**](JSONHALLink.md) |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**Products** | Pointer to [**[]EnvironmentBillOfMaterialsProduct**](EnvironmentBillOfMaterialsProduct.md) |  | [optional] 
**SolutionType** | Pointer to **string** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewEnvironmentBillOfMaterialsResponse

`func NewEnvironmentBillOfMaterialsResponse() *EnvironmentBillOfMaterialsResponse`

NewEnvironmentBillOfMaterialsResponse instantiates a new EnvironmentBillOfMaterialsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentBillOfMaterialsResponseWithDefaults

`func NewEnvironmentBillOfMaterialsResponseWithDefaults() *EnvironmentBillOfMaterialsResponse`

NewEnvironmentBillOfMaterialsResponseWithDefaults instantiates a new EnvironmentBillOfMaterialsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *EnvironmentBillOfMaterialsResponse) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *EnvironmentBillOfMaterialsResponse) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *EnvironmentBillOfMaterialsResponse) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *EnvironmentBillOfMaterialsResponse) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetLinks

`func (o *EnvironmentBillOfMaterialsResponse) GetLinks() map[string]JSONHALLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EnvironmentBillOfMaterialsResponse) GetLinksOk() (*map[string]JSONHALLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EnvironmentBillOfMaterialsResponse) SetLinks(v map[string]JSONHALLink)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EnvironmentBillOfMaterialsResponse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EnvironmentBillOfMaterialsResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EnvironmentBillOfMaterialsResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EnvironmentBillOfMaterialsResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EnvironmentBillOfMaterialsResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetProducts

`func (o *EnvironmentBillOfMaterialsResponse) GetProducts() []EnvironmentBillOfMaterialsProduct`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *EnvironmentBillOfMaterialsResponse) GetProductsOk() (*[]EnvironmentBillOfMaterialsProduct, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *EnvironmentBillOfMaterialsResponse) SetProducts(v []EnvironmentBillOfMaterialsProduct)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *EnvironmentBillOfMaterialsResponse) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetSolutionType

`func (o *EnvironmentBillOfMaterialsResponse) GetSolutionType() string`

GetSolutionType returns the SolutionType field if non-nil, zero value otherwise.

### GetSolutionTypeOk

`func (o *EnvironmentBillOfMaterialsResponse) GetSolutionTypeOk() (*string, bool)`

GetSolutionTypeOk returns a tuple with the SolutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolutionType

`func (o *EnvironmentBillOfMaterialsResponse) SetSolutionType(v string)`

SetSolutionType sets SolutionType field to given value.

### HasSolutionType

`func (o *EnvironmentBillOfMaterialsResponse) HasSolutionType() bool`

HasSolutionType returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *EnvironmentBillOfMaterialsResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EnvironmentBillOfMaterialsResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EnvironmentBillOfMaterialsResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EnvironmentBillOfMaterialsResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


