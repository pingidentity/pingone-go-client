# AccessFailedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ErrorResponseCommonDetails**](ErrorResponseCommonDetails.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Message** | **string** |  | 
**Code** | [**AccessFailedErrorCode**](AccessFailedErrorCode.md) |  | 

## Methods

### NewAccessFailedError

`func NewAccessFailedError(message string, code AccessFailedErrorCode, ) *AccessFailedError`

NewAccessFailedError instantiates a new AccessFailedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessFailedErrorWithDefaults

`func NewAccessFailedErrorWithDefaults() *AccessFailedError`

NewAccessFailedErrorWithDefaults instantiates a new AccessFailedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *AccessFailedError) GetDetails() []ErrorResponseCommonDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *AccessFailedError) GetDetailsOk() (*[]ErrorResponseCommonDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *AccessFailedError) SetDetails(v []ErrorResponseCommonDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *AccessFailedError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *AccessFailedError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AccessFailedError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AccessFailedError) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *AccessFailedError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *AccessFailedError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AccessFailedError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AccessFailedError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *AccessFailedError) GetCode() AccessFailedErrorCode`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AccessFailedError) GetCodeOk() (*AccessFailedErrorCode, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AccessFailedError) SetCode(v AccessFailedErrorCode)`

SetCode sets Code field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


