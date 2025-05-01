# ErrorResponseDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**InnerError** | Pointer to [**ErrorResponseDetailsInnerInnerError**](ErrorResponseDetailsInnerInnerError.md) |  | [optional] 
**Message** | **string** |  | 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorResponseDetailsInner

`func NewErrorResponseDetailsInner(code string, message string, ) *ErrorResponseDetailsInner`

NewErrorResponseDetailsInner instantiates a new ErrorResponseDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseDetailsInnerWithDefaults

`func NewErrorResponseDetailsInnerWithDefaults() *ErrorResponseDetailsInner`

NewErrorResponseDetailsInnerWithDefaults instantiates a new ErrorResponseDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseDetailsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseDetailsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseDetailsInner) SetCode(v string)`

SetCode sets Code field to given value.


### GetInnerError

`func (o *ErrorResponseDetailsInner) GetInnerError() ErrorResponseDetailsInnerInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *ErrorResponseDetailsInner) GetInnerErrorOk() (*ErrorResponseDetailsInnerInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *ErrorResponseDetailsInner) SetInnerError(v ErrorResponseDetailsInnerInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *ErrorResponseDetailsInner) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseDetailsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseDetailsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseDetailsInner) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTarget

`func (o *ErrorResponseDetailsInner) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ErrorResponseDetailsInner) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ErrorResponseDetailsInner) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ErrorResponseDetailsInner) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


