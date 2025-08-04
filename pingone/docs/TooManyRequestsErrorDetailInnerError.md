# TooManyRequestsErrorDetailInnerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QuotaLimit** | Pointer to **float32** |  | [optional] 
**QuotaResetTime** | Pointer to **time.Time** |  | [optional] 
**RetryAfter** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewTooManyRequestsErrorDetailInnerError

`func NewTooManyRequestsErrorDetailInnerError() *TooManyRequestsErrorDetailInnerError`

NewTooManyRequestsErrorDetailInnerError instantiates a new TooManyRequestsErrorDetailInnerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTooManyRequestsErrorDetailInnerErrorWithDefaults

`func NewTooManyRequestsErrorDetailInnerErrorWithDefaults() *TooManyRequestsErrorDetailInnerError`

NewTooManyRequestsErrorDetailInnerErrorWithDefaults instantiates a new TooManyRequestsErrorDetailInnerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuotaLimit

`func (o *TooManyRequestsErrorDetailInnerError) GetQuotaLimit() float32`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *TooManyRequestsErrorDetailInnerError) GetQuotaLimitOk() (*float32, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *TooManyRequestsErrorDetailInnerError) SetQuotaLimit(v float32)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *TooManyRequestsErrorDetailInnerError) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### GetQuotaResetTime

`func (o *TooManyRequestsErrorDetailInnerError) GetQuotaResetTime() time.Time`

GetQuotaResetTime returns the QuotaResetTime field if non-nil, zero value otherwise.

### GetQuotaResetTimeOk

`func (o *TooManyRequestsErrorDetailInnerError) GetQuotaResetTimeOk() (*time.Time, bool)`

GetQuotaResetTimeOk returns a tuple with the QuotaResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaResetTime

`func (o *TooManyRequestsErrorDetailInnerError) SetQuotaResetTime(v time.Time)`

SetQuotaResetTime sets QuotaResetTime field to given value.

### HasQuotaResetTime

`func (o *TooManyRequestsErrorDetailInnerError) HasQuotaResetTime() bool`

HasQuotaResetTime returns a boolean if a field has been set.

### GetRetryAfter

`func (o *TooManyRequestsErrorDetailInnerError) GetRetryAfter() time.Time`

GetRetryAfter returns the RetryAfter field if non-nil, zero value otherwise.

### GetRetryAfterOk

`func (o *TooManyRequestsErrorDetailInnerError) GetRetryAfterOk() (*time.Time, bool)`

GetRetryAfterOk returns a tuple with the RetryAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfter

`func (o *TooManyRequestsErrorDetailInnerError) SetRetryAfter(v time.Time)`

SetRetryAfter sets RetryAfter field to given value.

### HasRetryAfter

`func (o *TooManyRequestsErrorDetailInnerError) HasRetryAfter() bool`

HasRetryAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


