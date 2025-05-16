# ErrorResponseBadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ErrorResponseCommonDetails**](ErrorResponseCommonDetails.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Message** | **string** |  | 
**Code** | [**RequestFailedErrorCode**](RequestFailedErrorCode.md) |  | 

## Methods

### NewErrorResponseBadRequest

`func NewErrorResponseBadRequest(message string, code RequestFailedErrorCode, ) *ErrorResponseBadRequest`

NewErrorResponseBadRequest instantiates a new ErrorResponseBadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseBadRequestWithDefaults

`func NewErrorResponseBadRequestWithDefaults() *ErrorResponseBadRequest`

NewErrorResponseBadRequestWithDefaults instantiates a new ErrorResponseBadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ErrorResponseBadRequest) GetDetails() []ErrorResponseCommonDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorResponseBadRequest) GetDetailsOk() (*[]ErrorResponseCommonDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorResponseBadRequest) SetDetails(v []ErrorResponseCommonDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorResponseBadRequest) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *ErrorResponseBadRequest) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorResponseBadRequest) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorResponseBadRequest) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorResponseBadRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseBadRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseBadRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseBadRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ErrorResponseBadRequest) GetCode() RequestFailedErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseBadRequest) GetCodeOk() (*RequestFailedErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseBadRequest) SetCode(v RequestFailedErrorCode)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


