# TooManyRequestsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | [**TooManyRequestsErrorCode**](TooManyRequestsErrorCode.md) |  | 
**Message** | **string** |  | 
**Details** | Pointer to [**[]TooManyRequestsErrorDetail**](TooManyRequestsErrorDetail.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 

## Methods

### NewTooManyRequestsError

`func NewTooManyRequestsError(code TooManyRequestsErrorCode, message string, ) *TooManyRequestsError`

NewTooManyRequestsError instantiates a new TooManyRequestsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyRequestsErrorWithDefaults

`func NewTooManyRequestsErrorWithDefaults() *TooManyRequestsError`

NewTooManyRequestsErrorWithDefaults instantiates a new TooManyRequestsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TooManyRequestsError) GetCode() TooManyRequestsErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TooManyRequestsError) GetCodeOk() (*TooManyRequestsErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TooManyRequestsError) SetCode(v TooManyRequestsErrorCode)`

SetCode sets Code field to given value.


### GetMessage

`func (o *TooManyRequestsError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TooManyRequestsError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TooManyRequestsError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDetails

`func (o *TooManyRequestsError) GetDetails() []TooManyRequestsErrorDetail`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *TooManyRequestsError) GetDetailsOk() (*[]TooManyRequestsErrorDetail, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *TooManyRequestsError) SetDetails(v []TooManyRequestsErrorDetail)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *TooManyRequestsError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *TooManyRequestsError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TooManyRequestsError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TooManyRequestsError) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *TooManyRequestsError) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


