# ErrorResponseCommonDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**ErrorResponseCommonDetailsCode**](ErrorResponseCommonDetailsCode.md) |  | 
**InnerError** | Pointer to [**ErrorResponseCommonDetailsInnerError**](ErrorResponseCommonDetailsInnerError.md) |  | [optional] 
**Message** | **string** |  | 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewErrorResponseCommonDetails

`func NewErrorResponseCommonDetails(code ErrorResponseCommonDetailsCode, message string, ) *ErrorResponseCommonDetails`

NewErrorResponseCommonDetails instantiates a new ErrorResponseCommonDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseCommonDetailsWithDefaults

`func NewErrorResponseCommonDetailsWithDefaults() *ErrorResponseCommonDetails`

NewErrorResponseCommonDetailsWithDefaults instantiates a new ErrorResponseCommonDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseCommonDetails) GetCode() ErrorResponseCommonDetailsCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseCommonDetails) GetCodeOk() (*ErrorResponseCommonDetailsCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseCommonDetails) SetCode(v ErrorResponseCommonDetailsCode)`

SetCode sets Code field to given value.


### GetInnerError

`func (o *ErrorResponseCommonDetails) GetInnerError() ErrorResponseCommonDetailsInnerError`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *ErrorResponseCommonDetails) GetInnerErrorOk() (*ErrorResponseCommonDetailsInnerError, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *ErrorResponseCommonDetails) SetInnerError(v ErrorResponseCommonDetailsInnerError)`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *ErrorResponseCommonDetails) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseCommonDetails) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseCommonDetails) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseCommonDetails) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTarget

`func (o *ErrorResponseCommonDetails) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *ErrorResponseCommonDetails) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *ErrorResponseCommonDetails) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *ErrorResponseCommonDetails) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


