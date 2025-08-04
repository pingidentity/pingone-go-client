# TooManyRequestsErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**TooManyRequestsErrorDetailCode**](TooManyRequestsErrorDetailCode.md) |  | 
**Message** | **string** |  | 
**InnerError** | Pointer to [**TooManyRequestsErrorDetailInnerError**](TooManyRequestsErrorDetailInnerError.md) |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewTooManyRequestsErrorDetail

`func NewTooManyRequestsErrorDetail(code TooManyRequestsErrorDetailCode, message string, ) *TooManyRequestsErrorDetail`

NewTooManyRequestsErrorDetail instantiates a new TooManyRequestsErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyRequestsErrorDetailWithDefaults

`func NewTooManyRequestsErrorDetailWithDefaults() *TooManyRequestsErrorDetail`

NewTooManyRequestsErrorDetailWithDefaults instantiates a new TooManyRequestsErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TooManyRequestsErrorDetail) GetCode() TooManyRequestsErrorDetailCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TooManyRequestsErrorDetail) GetCodeOk() (*TooManyRequestsErrorDetailCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TooManyRequestsErrorDetail) SetCode(v TooManyRequestsErrorDetailCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *TooManyRequestsErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TooManyRequestsErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TooManyRequestsErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInnerError

`func (o *TooManyRequestsErrorDetail) GetInnerError() TooManyRequestsErrorDetailInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *TooManyRequestsErrorDetail) GetInnerErrorOk() (*TooManyRequestsErrorDetailInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *TooManyRequestsErrorDetail) SetInnerError(v TooManyRequestsErrorDetailInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *TooManyRequestsErrorDetail) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetTarget

`func (o *TooManyRequestsErrorDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *TooManyRequestsErrorDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *TooManyRequestsErrorDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *TooManyRequestsErrorDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


