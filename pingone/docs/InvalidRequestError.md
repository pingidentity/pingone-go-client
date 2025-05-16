# InvalidRequestError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ErrorResponseCommonDetails**](ErrorResponseCommonDetails.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Message** | **string** |  | 
**Code** | [**InvalidRequestErrorCode**](InvalidRequestErrorCode.md) |  | 

## Methods

### NewInvalidRequestError

`func NewInvalidRequestError(message string, code InvalidRequestErrorCode, ) *InvalidRequestError`

NewInvalidRequestError instantiates a new InvalidRequestError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvalidRequestErrorWithDefaults

`func NewInvalidRequestErrorWithDefaults() *InvalidRequestError`

NewInvalidRequestErrorWithDefaults instantiates a new InvalidRequestError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *InvalidRequestError) GetDetails() []ErrorResponseCommonDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *InvalidRequestError) GetDetailsOk() (*[]ErrorResponseCommonDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *InvalidRequestError) SetDetails(v []ErrorResponseCommonDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *InvalidRequestError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *InvalidRequestError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvalidRequestError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvalidRequestError) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *InvalidRequestError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *InvalidRequestError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *InvalidRequestError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *InvalidRequestError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *InvalidRequestError) GetCode() InvalidRequestErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InvalidRequestError) GetCodeOk() (*InvalidRequestErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InvalidRequestError) SetCode(v InvalidRequestErrorCode)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


