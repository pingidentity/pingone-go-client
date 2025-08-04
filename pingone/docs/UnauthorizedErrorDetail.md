# UnauthorizedErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**UnauthorizedErrorDetailCode**](UnauthorizedErrorDetailCode.md) |  | 
**Message** | **string** |  | 
**InnerError** | Pointer to **map[string]interface{}** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewUnauthorizedErrorDetail

`func NewUnauthorizedErrorDetail(code UnauthorizedErrorDetailCode, message string, ) *UnauthorizedErrorDetail`

NewUnauthorizedErrorDetail instantiates a new UnauthorizedErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorDetailWithDefaults

`func NewUnauthorizedErrorDetailWithDefaults() *UnauthorizedErrorDetail`

NewUnauthorizedErrorDetailWithDefaults instantiates a new UnauthorizedErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UnauthorizedErrorDetail) GetCode() UnauthorizedErrorDetailCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnauthorizedErrorDetail) GetCodeOk() (*UnauthorizedErrorDetailCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnauthorizedErrorDetail) SetCode(v UnauthorizedErrorDetailCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *UnauthorizedErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthorizedErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthorizedErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInnerError

`func (o *UnauthorizedErrorDetail) GetInnerError() map[string]interface{}`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *UnauthorizedErrorDetail) GetInnerErrorOk() (*map[string]interface{}, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *UnauthorizedErrorDetail) SetInnerError(v map[string]interface{})`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *UnauthorizedErrorDetail) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetTarget

`func (o *UnauthorizedErrorDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *UnauthorizedErrorDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *UnauthorizedErrorDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *UnauthorizedErrorDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


