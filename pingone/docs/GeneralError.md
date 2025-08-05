# GeneralError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**[]GeneralErrorDetail**](GeneralErrorDetail.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewGeneralError

`func NewGeneralError() *GeneralError`

NewGeneralError instantiates a new GeneralError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralErrorWithDefaults

`func NewGeneralErrorWithDefaults() *GeneralError`

NewGeneralErrorWithDefaults instantiates a new GeneralError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GeneralError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GeneralError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GeneralError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GeneralError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *GeneralError) GetDetails() []GeneralErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GeneralError) GetDetailsOk() (*[]GeneralErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GeneralError) SetDetails(v []GeneralErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GeneralError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *GeneralError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeneralError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeneralError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GeneralError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *GeneralError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GeneralError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GeneralError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GeneralError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


