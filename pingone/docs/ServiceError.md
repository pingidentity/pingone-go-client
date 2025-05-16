# ServiceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]ErrorResponseCommonDetails**](ErrorResponseCommonDetails.md) |  | [optional] 
**Id** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**Message** | **string** |  | 
**Code** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceError

`func NewServiceError(message string, ) *ServiceError`

NewServiceError instantiates a new ServiceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceErrorWithDefaults

`func NewServiceErrorWithDefaults() *ServiceError`

NewServiceErrorWithDefaults instantiates a new ServiceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *ServiceError) GetDetails() []ErrorResponseCommonDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ServiceError) GetDetailsOk() (*[]ErrorResponseCommonDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ServiceError) SetDetails(v []ErrorResponseCommonDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ServiceError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *ServiceError) GetId() uuid.UUID`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceError) GetIdOk() (*uuid.UUID, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceError) SetId(v uuid.UUID)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ServiceError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetCode

`func (o *ServiceError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServiceError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServiceError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServiceError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


