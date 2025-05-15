# ErrorResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**ErrorResponseDetailsCode**](ErrorResponseDetailsCode.md) |  | 
**InnerError** | Pointer to [**ErrorResponseDetailsInnerError**](ErrorResponseDetailsInnerError.md) |  | [optional] 
**Message** | **string** |  | 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorResponseDetails

`func NewErrorResponseDetails(code ErrorResponseDetailsCode, message string, ) *ErrorResponseDetails`

NewErrorResponseDetails instantiates a new ErrorResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseDetailsWithDefaults

`func NewErrorResponseDetailsWithDefaults() *ErrorResponseDetails`

NewErrorResponseDetailsWithDefaults instantiates a new ErrorResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseDetails) GetCode() ErrorResponseDetailsCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseDetails) GetCodeOk() (*ErrorResponseDetailsCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseDetails) SetCode(v ErrorResponseDetailsCode)`

SetCode sets Code field to given value.


### GetInnerError

`func (o *ErrorResponseDetails) GetInnerError() ErrorResponseDetailsInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *ErrorResponseDetails) GetInnerErrorOk() (*ErrorResponseDetailsInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *ErrorResponseDetails) SetInnerError(v ErrorResponseDetailsInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *ErrorResponseDetails) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseDetails) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTarget

`func (o *ErrorResponseDetails) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ErrorResponseDetails) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ErrorResponseDetails) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ErrorResponseDetails) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


