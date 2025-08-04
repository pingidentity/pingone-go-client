# BadRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**uuid.UUID**](uuid.UUID.md) |  | 
**Code** | [**BadRequestErrorCode**](BadRequestErrorCode.md) |  | 
**Message** | **string** |  | 
**Details** | Pointer to [**[]BadRequestErrorDetail**](BadRequestErrorDetail.md) |  | [optional] 

## Methods

### NewBadRequestError

`func NewBadRequestError(id uuid.UUID, code BadRequestErrorCode, message string, ) *BadRequestError`

NewBadRequestError instantiates a new BadRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestErrorWithDefaults

`func NewBadRequestErrorWithDefaults() *BadRequestError`

NewBadRequestErrorWithDefaults instantiates a new BadRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BadRequestError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BadRequestError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BadRequestError) SetId(v uuid.UUID)`

SetId sets Id field to given value.


### GetCode

`func (o *BadRequestError) GetCode() BadRequestErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *BadRequestError) GetCodeOk() (*BadRequestErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *BadRequestError) SetCode(v BadRequestErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *BadRequestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *BadRequestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *BadRequestError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *BadRequestError) GetDetails() []BadRequestErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BadRequestError) GetDetailsOk() (*[]BadRequestErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BadRequestError) SetDetails(v []BadRequestErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BadRequestError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


