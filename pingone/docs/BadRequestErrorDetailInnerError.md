# BadRequestErrorDetailInnerError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedPattern** | Pointer to **string** |  | [optional] 
**AllowedValues** | Pointer to **[]string** |  | [optional] 
**Claim** | Pointer to **string** |  | [optional] 
**ExistingId** | Pointer to [**uuid.UUID**](uuid.UUID.md) |  | [optional] 
**MaximumValue** | Pointer to **float32** |  | [optional] 
**QuotaLimit** | Pointer to **float32** |  | [optional] 
**QuotaResetTime** | Pointer to **time.Time** |  | [optional] 
**RangeMaximumValue** | Pointer to **float32** |  | [optional] 
**RangeMinimumValue** | Pointer to **float32** |  | [optional] 

## Methods

### NewBadRequestErrorDetailInnerError

`func NewBadRequestErrorDetailInnerError() *BadRequestErrorDetailInnerError`

NewBadRequestErrorDetailInnerError instantiates a new BadRequestErrorDetailInnerError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBadRequestErrorDetailInnerErrorWithDefaults

`func NewBadRequestErrorDetailInnerErrorWithDefaults() *BadRequestErrorDetailInnerError`

NewBadRequestErrorDetailInnerErrorWithDefaults instantiates a new BadRequestErrorDetailInnerError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedPattern

`func (o *BadRequestErrorDetailInnerError) GetAllowedPattern() string`

GetAllowedPattern returns the AllowedPattern field if non-nil, zero value otherwise.

### GetAllowedPatternOk

`func (o *BadRequestErrorDetailInnerError) GetAllowedPatternOk() (*string, bool)`

GetAllowedPatternOk returns a tuple with the AllowedPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedPattern

`func (o *BadRequestErrorDetailInnerError) SetAllowedPattern(v string)`

SetAllowedPattern sets AllowedPattern field to given value.

### HasAllowedPattern

`func (o *BadRequestErrorDetailInnerError) HasAllowedPattern() bool`

HasAllowedPattern returns a boolean if a field has been set.

### GetAllowedValues

`func (o *BadRequestErrorDetailInnerError) GetAllowedValues() []string`

GetAllowedValues returns the AllowedValues field if non-nil, zero value otherwise.

### GetAllowedValuesOk

`func (o *BadRequestErrorDetailInnerError) GetAllowedValuesOk() (*[]string, bool)`

GetAllowedValuesOk returns a tuple with the AllowedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedValues

`func (o *BadRequestErrorDetailInnerError) SetAllowedValues(v []string)`

SetAllowedValues sets AllowedValues field to given value.

### HasAllowedValues

`func (o *BadRequestErrorDetailInnerError) HasAllowedValues() bool`

HasAllowedValues returns a boolean if a field has been set.

### GetClaim

`func (o *BadRequestErrorDetailInnerError) GetClaim() string`

GetClaim returns the Claim field if non-nil, zero value otherwise.

### GetClaimOk

`func (o *BadRequestErrorDetailInnerError) GetClaimOk() (*string, bool)`

GetClaimOk returns a tuple with the Claim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaim

`func (o *BadRequestErrorDetailInnerError) SetClaim(v string)`

SetClaim sets Claim field to given value.

### HasClaim

`func (o *BadRequestErrorDetailInnerError) HasClaim() bool`

HasClaim returns a boolean if a field has been set.

### GetExistingId

`func (o *BadRequestErrorDetailInnerError) GetExistingId() uuid.UUID`

GetExistingId returns the ExistingId field if non-nil, zero value otherwise.

### GetExistingIdOk

`func (o *BadRequestErrorDetailInnerError) GetExistingIdOk() (*uuid.UUID, bool)`

GetExistingIdOk returns a tuple with the ExistingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistingId

`func (o *BadRequestErrorDetailInnerError) SetExistingId(v uuid.UUID)`

SetExistingId sets ExistingId field to given value.

### HasExistingId

`func (o *BadRequestErrorDetailInnerError) HasExistingId() bool`

HasExistingId returns a boolean if a field has been set.

### GetMaximumValue

`func (o *BadRequestErrorDetailInnerError) GetMaximumValue() float32`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *BadRequestErrorDetailInnerError) GetMaximumValueOk() (*float32, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *BadRequestErrorDetailInnerError) SetMaximumValue(v float32)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *BadRequestErrorDetailInnerError) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetQuotaLimit

`func (o *BadRequestErrorDetailInnerError) GetQuotaLimit() float32`

GetQuotaLimit returns the QuotaLimit field if non-nil, zero value otherwise.

### GetQuotaLimitOk

`func (o *BadRequestErrorDetailInnerError) GetQuotaLimitOk() (*float32, bool)`

GetQuotaLimitOk returns a tuple with the QuotaLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaLimit

`func (o *BadRequestErrorDetailInnerError) SetQuotaLimit(v float32)`

SetQuotaLimit sets QuotaLimit field to given value.

### HasQuotaLimit

`func (o *BadRequestErrorDetailInnerError) HasQuotaLimit() bool`

HasQuotaLimit returns a boolean if a field has been set.

### GetQuotaResetTime

`func (o *BadRequestErrorDetailInnerError) GetQuotaResetTime() time.Time`

GetQuotaResetTime returns the QuotaResetTime field if non-nil, zero value otherwise.

### GetQuotaResetTimeOk

`func (o *BadRequestErrorDetailInnerError) GetQuotaResetTimeOk() (*time.Time, bool)`

GetQuotaResetTimeOk returns a tuple with the QuotaResetTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotaResetTime

`func (o *BadRequestErrorDetailInnerError) SetQuotaResetTime(v time.Time)`

SetQuotaResetTime sets QuotaResetTime field to given value.

### HasQuotaResetTime

`func (o *BadRequestErrorDetailInnerError) HasQuotaResetTime() bool`

HasQuotaResetTime returns a boolean if a field has been set.

### GetRangeMaximumValue

`func (o *BadRequestErrorDetailInnerError) GetRangeMaximumValue() float32`

GetRangeMaximumValue returns the RangeMaximumValue field if non-nil, zero value otherwise.

### GetRangeMaximumValueOk

`func (o *BadRequestErrorDetailInnerError) GetRangeMaximumValueOk() (*float32, bool)`

GetRangeMaximumValueOk returns a tuple with the RangeMaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeMaximumValue

`func (o *BadRequestErrorDetailInnerError) SetRangeMaximumValue(v float32)`

SetRangeMaximumValue sets RangeMaximumValue field to given value.

### HasRangeMaximumValue

`func (o *BadRequestErrorDetailInnerError) HasRangeMaximumValue() bool`

HasRangeMaximumValue returns a boolean if a field has been set.

### GetRangeMinimumValue

`func (o *BadRequestErrorDetailInnerError) GetRangeMinimumValue() float32`

GetRangeMinimumValue returns the RangeMinimumValue field if non-nil, zero value otherwise.

### GetRangeMinimumValueOk

`func (o *BadRequestErrorDetailInnerError) GetRangeMinimumValueOk() (*float32, bool)`

GetRangeMinimumValueOk returns a tuple with the RangeMinimumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeMinimumValue

`func (o *BadRequestErrorDetailInnerError) SetRangeMinimumValue(v float32)`

SetRangeMinimumValue sets RangeMinimumValue field to given value.

### HasRangeMinimumValue

`func (o *BadRequestErrorDetailInnerError) HasRangeMinimumValue() bool`

HasRangeMinimumValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


