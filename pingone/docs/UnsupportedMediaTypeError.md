# UnsupportedMediaTypeError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**uuid.UUID**](uuid.UUID.md) |  | 
**Code** | [**UnsupportedMediaTypeErrorCode**](UnsupportedMediaTypeErrorCode.md) |  | 
**Message** | **string** |  | 
**Details** | Pointer to [**[]UnsupportedMediaTypeErrorDetail**](UnsupportedMediaTypeErrorDetail.md) |  | [optional] 

## Methods

### NewUnsupportedMediaTypeError

`func NewUnsupportedMediaTypeError(id uuid.UUID, code UnsupportedMediaTypeErrorCode, message string, ) *UnsupportedMediaTypeError`

NewUnsupportedMediaTypeError instantiates a new UnsupportedMediaTypeError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnsupportedMediaTypeErrorWithDefaults

`func NewUnsupportedMediaTypeErrorWithDefaults() *UnsupportedMediaTypeError`

NewUnsupportedMediaTypeErrorWithDefaults instantiates a new UnsupportedMediaTypeError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnsupportedMediaTypeError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnsupportedMediaTypeError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnsupportedMediaTypeError) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCode

`func (o *UnsupportedMediaTypeError) GetCode() UnsupportedMediaTypeErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnsupportedMediaTypeError) GetCodeOk() (*UnsupportedMediaTypeErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnsupportedMediaTypeError) SetCode(v UnsupportedMediaTypeErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *UnsupportedMediaTypeError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnsupportedMediaTypeError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnsupportedMediaTypeError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *UnsupportedMediaTypeError) GetDetails() []UnsupportedMediaTypeErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *UnsupportedMediaTypeError) GetDetailsOk() (*[]UnsupportedMediaTypeErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *UnsupportedMediaTypeError) SetDetails(v []UnsupportedMediaTypeErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *UnsupportedMediaTypeError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


