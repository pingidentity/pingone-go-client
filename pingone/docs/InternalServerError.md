# InternalServerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**uuid.UUID**](uuid.UUID.md) |  | 
**Code** | [**InternalServerErrorCode**](InternalServerErrorCode.md) |  | 
**Message** | **string** |  | 
**Details** | Pointer to [**[]InternalServerErrorDetail**](InternalServerErrorDetail.md) |  | [optional] 

## Methods

### NewInternalServerError

`func NewInternalServerError(id uuid.UUID, code InternalServerErrorCode, message string, ) *InternalServerError`

NewInternalServerError instantiates a new InternalServerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternalServerErrorWithDefaults

`func NewInternalServerErrorWithDefaults() *InternalServerError`

NewInternalServerErrorWithDefaults instantiates a new InternalServerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InternalServerError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InternalServerError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InternalServerError) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCode

`func (o *InternalServerError) GetCode() InternalServerErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InternalServerError) GetCodeOk() (*InternalServerErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InternalServerError) SetCode(v InternalServerErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *InternalServerError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InternalServerError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InternalServerError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *InternalServerError) GetDetails() []InternalServerErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InternalServerError) GetDetailsOk() (*[]InternalServerErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InternalServerError) SetDetails(v []InternalServerErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InternalServerError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


