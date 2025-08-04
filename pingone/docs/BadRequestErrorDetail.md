# BadRequestErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**BadRequestErrorDetailCode**](BadRequestErrorDetailCode.md) |  | 
**Message** | **string** |  | 
**InnerError** | Pointer to [**BadRequestErrorDetailInnerError**](BadRequestErrorDetailInnerError.md) |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewBadRequestErrorDetail

`func NewBadRequestErrorDetail(code BadRequestErrorDetailCode, message string, ) *BadRequestErrorDetail`

NewBadRequestErrorDetail instantiates a new BadRequestErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestErrorDetailWithDefaults

`func NewBadRequestErrorDetailWithDefaults() *BadRequestErrorDetail`

NewBadRequestErrorDetailWithDefaults instantiates a new BadRequestErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *BadRequestErrorDetail) GetCode() BadRequestErrorDetailCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadRequestErrorDetail) GetCodeOk() (*BadRequestErrorDetailCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadRequestErrorDetail) SetCode(v BadRequestErrorDetailCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *BadRequestErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequestErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequestErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInnerError

`func (o *BadRequestErrorDetail) GetInnerError() BadRequestErrorDetailInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *BadRequestErrorDetail) GetInnerErrorOk() (*BadRequestErrorDetailInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *BadRequestErrorDetail) SetInnerError(v BadRequestErrorDetailInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *BadRequestErrorDetail) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetTarget

`func (o *BadRequestErrorDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *BadRequestErrorDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *BadRequestErrorDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *BadRequestErrorDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


