# UnexpectedServiceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ErrorResponseCommonDetails**](ErrorResponseCommonDetails.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Message** | **string** |  | 
**Code** | [**UnexpectedServiceErrorCode**](UnexpectedServiceErrorCode.md) |  | 

## Methods

### NewUnexpectedServiceError

`func NewUnexpectedServiceError(message string, code UnexpectedServiceErrorCode, ) *UnexpectedServiceError`

NewUnexpectedServiceError instantiates a new UnexpectedServiceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnexpectedServiceErrorWithDefaults

`func NewUnexpectedServiceErrorWithDefaults() *UnexpectedServiceError`

NewUnexpectedServiceErrorWithDefaults instantiates a new UnexpectedServiceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *UnexpectedServiceError) GetDetails() []ErrorResponseCommonDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *UnexpectedServiceError) GetDetailsOk() (*[]ErrorResponseCommonDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *UnexpectedServiceError) SetDetails(v []ErrorResponseCommonDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *UnexpectedServiceError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *UnexpectedServiceError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UnexpectedServiceError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UnexpectedServiceError) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *UnexpectedServiceError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *UnexpectedServiceError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UnexpectedServiceError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UnexpectedServiceError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *UnexpectedServiceError) GetCode() UnexpectedServiceErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UnexpectedServiceError) GetCodeOk() (*UnexpectedServiceErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UnexpectedServiceError) SetCode(v UnexpectedServiceErrorCode)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


