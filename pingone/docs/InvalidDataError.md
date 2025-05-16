# InvalidDataError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ErrorResponseCommonDetails**](ErrorResponseCommonDetails.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Message** | **string** |  | 
**Code** | [**InvalidDataErrorCode**](InvalidDataErrorCode.md) |  | 

## Methods

### NewInvalidDataError

`func NewInvalidDataError(message string, code InvalidDataErrorCode, ) *InvalidDataError`

NewInvalidDataError instantiates a new InvalidDataError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidDataErrorWithDefaults

`func NewInvalidDataErrorWithDefaults() *InvalidDataError`

NewInvalidDataErrorWithDefaults instantiates a new InvalidDataError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *InvalidDataError) GetDetails() []ErrorResponseCommonDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InvalidDataError) GetDetailsOk() (*[]ErrorResponseCommonDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InvalidDataError) SetDetails(v []ErrorResponseCommonDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InvalidDataError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *InvalidDataError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvalidDataError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvalidDataError) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *InvalidDataError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *InvalidDataError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvalidDataError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvalidDataError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *InvalidDataError) GetCode() InvalidDataErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InvalidDataError) GetCodeOk() (*InvalidDataErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InvalidDataError) SetCode(v InvalidDataErrorCode)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


