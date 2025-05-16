# ErrorResponseCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ErrorResponseCommonDetails**](ErrorResponseCommonDetails.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Message** | **string** |  | 

## Methods

### NewErrorResponseCommon

`func NewErrorResponseCommon(message string, ) *ErrorResponseCommon`

NewErrorResponseCommon instantiates a new ErrorResponseCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseCommonWithDefaults

`func NewErrorResponseCommonWithDefaults() *ErrorResponseCommon`

NewErrorResponseCommonWithDefaults instantiates a new ErrorResponseCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ErrorResponseCommon) GetDetails() []ErrorResponseCommonDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorResponseCommon) GetDetailsOk() (*[]ErrorResponseCommonDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorResponseCommon) SetDetails(v []ErrorResponseCommonDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorResponseCommon) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *ErrorResponseCommon) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorResponseCommon) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorResponseCommon) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *ErrorResponseCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ErrorResponseCommon) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseCommon) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseCommon) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


