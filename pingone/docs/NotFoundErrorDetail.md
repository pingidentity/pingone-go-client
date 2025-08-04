# NotFoundErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** |  | 
**Message** | **string** |  | 
**InnerError** | Pointer to **map[string]interface{}** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewNotFoundErrorDetail

`func NewNotFoundErrorDetail(code string, message string, ) *NotFoundErrorDetail`

NewNotFoundErrorDetail instantiates a new NotFoundErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundErrorDetailWithDefaults

`func NewNotFoundErrorDetailWithDefaults() *NotFoundErrorDetail`

NewNotFoundErrorDetailWithDefaults instantiates a new NotFoundErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *NotFoundErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NotFoundErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NotFoundErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.


### GetMessage

`func (o *NotFoundErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotFoundErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotFoundErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetInnerError

`func (o *NotFoundErrorDetail) GetInnerError() map[string]interface{}`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *NotFoundErrorDetail) GetInnerErrorOk() (*map[string]interface{}, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *NotFoundErrorDetail) SetInnerError(v map[string]interface{})`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *NotFoundErrorDetail) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetTarget

`func (o *NotFoundErrorDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *NotFoundErrorDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *NotFoundErrorDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *NotFoundErrorDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


