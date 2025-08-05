# ForbiddenErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**ForbiddenErrorDetailCode**](ForbiddenErrorDetailCode.md) |  | 
**Message** | **string** |  | 
**InnerError** | Pointer to [**ForbiddenErrorDetailInnerError**](ForbiddenErrorDetailInnerError.md) |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewForbiddenErrorDetail

`func NewForbiddenErrorDetail(code ForbiddenErrorDetailCode, message string, ) *ForbiddenErrorDetail`

NewForbiddenErrorDetail instantiates a new ForbiddenErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForbiddenErrorDetailWithDefaults

`func NewForbiddenErrorDetailWithDefaults() *ForbiddenErrorDetail`

NewForbiddenErrorDetailWithDefaults instantiates a new ForbiddenErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ForbiddenErrorDetail) GetCode() ForbiddenErrorDetailCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ForbiddenErrorDetail) GetCodeOk() (*ForbiddenErrorDetailCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ForbiddenErrorDetail) SetCode(v ForbiddenErrorDetailCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *ForbiddenErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ForbiddenErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ForbiddenErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInnerError

`func (o *ForbiddenErrorDetail) GetInnerError() ForbiddenErrorDetailInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *ForbiddenErrorDetail) GetInnerErrorOk() (*ForbiddenErrorDetailInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *ForbiddenErrorDetail) SetInnerError(v ForbiddenErrorDetailInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *ForbiddenErrorDetail) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetTarget

`func (o *ForbiddenErrorDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ForbiddenErrorDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ForbiddenErrorDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ForbiddenErrorDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


