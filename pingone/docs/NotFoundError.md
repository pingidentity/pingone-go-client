# NotFoundError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ErrorResponseCommonDetails**](ErrorResponseCommonDetails.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Message** | **string** |  | 
**Code** | [**NotFoundErrorCode**](NotFoundErrorCode.md) |  | 

## Methods

### NewNotFoundError

`func NewNotFoundError(message string, code NotFoundErrorCode, ) *NotFoundError`

NewNotFoundError instantiates a new NotFoundError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundErrorWithDefaults

`func NewNotFoundErrorWithDefaults() *NotFoundError`

NewNotFoundErrorWithDefaults instantiates a new NotFoundError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *NotFoundError) GetDetails() []ErrorResponseCommonDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *NotFoundError) GetDetailsOk() (*[]ErrorResponseCommonDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *NotFoundError) SetDetails(v []ErrorResponseCommonDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *NotFoundError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *NotFoundError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotFoundError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotFoundError) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *NotFoundError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *NotFoundError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *NotFoundError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *NotFoundError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *NotFoundError) GetCode() NotFoundErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *NotFoundError) GetCodeOk() (*NotFoundErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *NotFoundError) SetCode(v NotFoundErrorCode)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


