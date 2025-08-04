# UnauthorizedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**uuid.UUID**](uuid.UUID.md) |  | 
**Code** | [**UnauthorizedErrorCode**](UnauthorizedErrorCode.md) |  | 
**Message** | **string** |  | 
**Details** | Pointer to [**[]UnauthorizedErrorDetail**](UnauthorizedErrorDetail.md) |  | [optional] 

## Methods

### NewUnauthorizedError

`func NewUnauthorizedError(id uuid.UUID, code UnauthorizedErrorCode, message string, ) *UnauthorizedError`

NewUnauthorizedError instantiates a new UnauthorizedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnauthorizedErrorWithDefaults

`func NewUnauthorizedErrorWithDefaults() *UnauthorizedError`

NewUnauthorizedErrorWithDefaults instantiates a new UnauthorizedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UnauthorizedError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnauthorizedError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnauthorizedError) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCode

`func (o *UnauthorizedError) GetCode() UnauthorizedErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnauthorizedError) GetCodeOk() (*UnauthorizedErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnauthorizedError) SetCode(v UnauthorizedErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *UnauthorizedError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnauthorizedError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnauthorizedError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *UnauthorizedError) GetDetails() []UnauthorizedErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *UnauthorizedError) GetDetailsOk() (*[]UnauthorizedErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *UnauthorizedError) SetDetails(v []UnauthorizedErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *UnauthorizedError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


