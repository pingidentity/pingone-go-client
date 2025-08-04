# GeneralErrorDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**InnerError** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 

## Methods

### NewGeneralErrorDetail

`func NewGeneralErrorDetail() *GeneralErrorDetail`

NewGeneralErrorDetail instantiates a new GeneralErrorDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralErrorDetailWithDefaults

`func NewGeneralErrorDetailWithDefaults() *GeneralErrorDetail`

NewGeneralErrorDetailWithDefaults instantiates a new GeneralErrorDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GeneralErrorDetail) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GeneralErrorDetail) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GeneralErrorDetail) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GeneralErrorDetail) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetInnerError

`func (o *GeneralErrorDetail) GetInnerError() map[string]map[string]interface{}`

GetInnerError returns the InnerError field if non-nil, zero value otherwise.

### GetInnerErrorOk

`func (o *GeneralErrorDetail) GetInnerErrorOk() (*map[string]map[string]interface{}, bool)`

GetInnerErrorOk returns a tuple with the InnerError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInnerError

`func (o *GeneralErrorDetail) SetInnerError(v map[string]map[string]interface{})`

SetInnerError sets InnerError field to given value.

### HasInnerError

`func (o *GeneralErrorDetail) HasInnerError() bool`

HasInnerError returns a boolean if a field has been set.

### GetMessage

`func (o *GeneralErrorDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GeneralErrorDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GeneralErrorDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GeneralErrorDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTarget

`func (o *GeneralErrorDetail) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *GeneralErrorDetail) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *GeneralErrorDetail) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *GeneralErrorDetail) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


