# EnvironmentLicenseQuantities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NumberOfAadhaar** | Pointer to **int64** |  | [optional] 
**NumberOfAccountOwnership** | Pointer to **int64** |  | [optional] 
**NumberOfDataVerifications** | Pointer to **int64** |  | [optional] 
**RlgAnalyticsRpm** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgAppRegRpm** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgAuditRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgAuthnRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgAuthzRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgCatchAllRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgConfigRpm** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgCredRpm** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgDirAccessRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgDirBulkRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgDirFixedRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgDirWriteRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgDvFlowCallbackRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgDvFlowExecRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgDvFlowStartRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgMfaFixedRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgMfaPollingRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgMfaRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgRiskEvalsRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgTokenCheckRps** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 
**RlgVerifyRpm** | Pointer to [**EnvironmentAPILimit**](EnvironmentAPILimit.md) |  | [optional] 

## Methods

### NewEnvironmentLicenseQuantities

`func NewEnvironmentLicenseQuantities() *EnvironmentLicenseQuantities`

NewEnvironmentLicenseQuantities instantiates a new EnvironmentLicenseQuantities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentLicenseQuantitiesWithDefaults

`func NewEnvironmentLicenseQuantitiesWithDefaults() *EnvironmentLicenseQuantities`

NewEnvironmentLicenseQuantitiesWithDefaults instantiates a new EnvironmentLicenseQuantities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumberOfAadhaar

`func (o *EnvironmentLicenseQuantities) GetNumberOfAadhaar() int64`

GetNumberOfAadhaar returns the NumberOfAadhaar field if non-nil, zero value otherwise.

### GetNumberOfAadhaarOk

`func (o *EnvironmentLicenseQuantities) GetNumberOfAadhaarOk() (*int64, bool)`

GetNumberOfAadhaarOk returns a tuple with the NumberOfAadhaar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAadhaar

`func (o *EnvironmentLicenseQuantities) SetNumberOfAadhaar(v int64)`

SetNumberOfAadhaar sets NumberOfAadhaar field to given value.

### HasNumberOfAadhaar

`func (o *EnvironmentLicenseQuantities) HasNumberOfAadhaar() bool`

HasNumberOfAadhaar returns a boolean if a field has been set.

### GetNumberOfAccountOwnership

`func (o *EnvironmentLicenseQuantities) GetNumberOfAccountOwnership() int64`

GetNumberOfAccountOwnership returns the NumberOfAccountOwnership field if non-nil, zero value otherwise.

### GetNumberOfAccountOwnershipOk

`func (o *EnvironmentLicenseQuantities) GetNumberOfAccountOwnershipOk() (*int64, bool)`

GetNumberOfAccountOwnershipOk returns a tuple with the NumberOfAccountOwnership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAccountOwnership

`func (o *EnvironmentLicenseQuantities) SetNumberOfAccountOwnership(v int64)`

SetNumberOfAccountOwnership sets NumberOfAccountOwnership field to given value.

### HasNumberOfAccountOwnership

`func (o *EnvironmentLicenseQuantities) HasNumberOfAccountOwnership() bool`

HasNumberOfAccountOwnership returns a boolean if a field has been set.

### GetNumberOfDataVerifications

`func (o *EnvironmentLicenseQuantities) GetNumberOfDataVerifications() int64`

GetNumberOfDataVerifications returns the NumberOfDataVerifications field if non-nil, zero value otherwise.

### GetNumberOfDataVerificationsOk

`func (o *EnvironmentLicenseQuantities) GetNumberOfDataVerificationsOk() (*int64, bool)`

GetNumberOfDataVerificationsOk returns a tuple with the NumberOfDataVerifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfDataVerifications

`func (o *EnvironmentLicenseQuantities) SetNumberOfDataVerifications(v int64)`

SetNumberOfDataVerifications sets NumberOfDataVerifications field to given value.

### HasNumberOfDataVerifications

`func (o *EnvironmentLicenseQuantities) HasNumberOfDataVerifications() bool`

HasNumberOfDataVerifications returns a boolean if a field has been set.

### GetRlgAnalyticsRpm

`func (o *EnvironmentLicenseQuantities) GetRlgAnalyticsRpm() EnvironmentAPILimit`

GetRlgAnalyticsRpm returns the RlgAnalyticsRpm field if non-nil, zero value otherwise.

### GetRlgAnalyticsRpmOk

`func (o *EnvironmentLicenseQuantities) GetRlgAnalyticsRpmOk() (*EnvironmentAPILimit, bool)`

GetRlgAnalyticsRpmOk returns a tuple with the RlgAnalyticsRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgAnalyticsRpm

`func (o *EnvironmentLicenseQuantities) SetRlgAnalyticsRpm(v EnvironmentAPILimit)`

SetRlgAnalyticsRpm sets RlgAnalyticsRpm field to given value.

### HasRlgAnalyticsRpm

`func (o *EnvironmentLicenseQuantities) HasRlgAnalyticsRpm() bool`

HasRlgAnalyticsRpm returns a boolean if a field has been set.

### GetRlgAppRegRpm

`func (o *EnvironmentLicenseQuantities) GetRlgAppRegRpm() EnvironmentAPILimit`

GetRlgAppRegRpm returns the RlgAppRegRpm field if non-nil, zero value otherwise.

### GetRlgAppRegRpmOk

`func (o *EnvironmentLicenseQuantities) GetRlgAppRegRpmOk() (*EnvironmentAPILimit, bool)`

GetRlgAppRegRpmOk returns a tuple with the RlgAppRegRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgAppRegRpm

`func (o *EnvironmentLicenseQuantities) SetRlgAppRegRpm(v EnvironmentAPILimit)`

SetRlgAppRegRpm sets RlgAppRegRpm field to given value.

### HasRlgAppRegRpm

`func (o *EnvironmentLicenseQuantities) HasRlgAppRegRpm() bool`

HasRlgAppRegRpm returns a boolean if a field has been set.

### GetRlgAuditRps

`func (o *EnvironmentLicenseQuantities) GetRlgAuditRps() EnvironmentAPILimit`

GetRlgAuditRps returns the RlgAuditRps field if non-nil, zero value otherwise.

### GetRlgAuditRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgAuditRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgAuditRpsOk returns a tuple with the RlgAuditRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgAuditRps

`func (o *EnvironmentLicenseQuantities) SetRlgAuditRps(v EnvironmentAPILimit)`

SetRlgAuditRps sets RlgAuditRps field to given value.

### HasRlgAuditRps

`func (o *EnvironmentLicenseQuantities) HasRlgAuditRps() bool`

HasRlgAuditRps returns a boolean if a field has been set.

### GetRlgAuthnRps

`func (o *EnvironmentLicenseQuantities) GetRlgAuthnRps() EnvironmentAPILimit`

GetRlgAuthnRps returns the RlgAuthnRps field if non-nil, zero value otherwise.

### GetRlgAuthnRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgAuthnRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgAuthnRpsOk returns a tuple with the RlgAuthnRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgAuthnRps

`func (o *EnvironmentLicenseQuantities) SetRlgAuthnRps(v EnvironmentAPILimit)`

SetRlgAuthnRps sets RlgAuthnRps field to given value.

### HasRlgAuthnRps

`func (o *EnvironmentLicenseQuantities) HasRlgAuthnRps() bool`

HasRlgAuthnRps returns a boolean if a field has been set.

### GetRlgAuthzRps

`func (o *EnvironmentLicenseQuantities) GetRlgAuthzRps() EnvironmentAPILimit`

GetRlgAuthzRps returns the RlgAuthzRps field if non-nil, zero value otherwise.

### GetRlgAuthzRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgAuthzRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgAuthzRpsOk returns a tuple with the RlgAuthzRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgAuthzRps

`func (o *EnvironmentLicenseQuantities) SetRlgAuthzRps(v EnvironmentAPILimit)`

SetRlgAuthzRps sets RlgAuthzRps field to given value.

### HasRlgAuthzRps

`func (o *EnvironmentLicenseQuantities) HasRlgAuthzRps() bool`

HasRlgAuthzRps returns a boolean if a field has been set.

### GetRlgCatchAllRps

`func (o *EnvironmentLicenseQuantities) GetRlgCatchAllRps() EnvironmentAPILimit`

GetRlgCatchAllRps returns the RlgCatchAllRps field if non-nil, zero value otherwise.

### GetRlgCatchAllRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgCatchAllRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgCatchAllRpsOk returns a tuple with the RlgCatchAllRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgCatchAllRps

`func (o *EnvironmentLicenseQuantities) SetRlgCatchAllRps(v EnvironmentAPILimit)`

SetRlgCatchAllRps sets RlgCatchAllRps field to given value.

### HasRlgCatchAllRps

`func (o *EnvironmentLicenseQuantities) HasRlgCatchAllRps() bool`

HasRlgCatchAllRps returns a boolean if a field has been set.

### GetRlgConfigRpm

`func (o *EnvironmentLicenseQuantities) GetRlgConfigRpm() EnvironmentAPILimit`

GetRlgConfigRpm returns the RlgConfigRpm field if non-nil, zero value otherwise.

### GetRlgConfigRpmOk

`func (o *EnvironmentLicenseQuantities) GetRlgConfigRpmOk() (*EnvironmentAPILimit, bool)`

GetRlgConfigRpmOk returns a tuple with the RlgConfigRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgConfigRpm

`func (o *EnvironmentLicenseQuantities) SetRlgConfigRpm(v EnvironmentAPILimit)`

SetRlgConfigRpm sets RlgConfigRpm field to given value.

### HasRlgConfigRpm

`func (o *EnvironmentLicenseQuantities) HasRlgConfigRpm() bool`

HasRlgConfigRpm returns a boolean if a field has been set.

### GetRlgCredRpm

`func (o *EnvironmentLicenseQuantities) GetRlgCredRpm() EnvironmentAPILimit`

GetRlgCredRpm returns the RlgCredRpm field if non-nil, zero value otherwise.

### GetRlgCredRpmOk

`func (o *EnvironmentLicenseQuantities) GetRlgCredRpmOk() (*EnvironmentAPILimit, bool)`

GetRlgCredRpmOk returns a tuple with the RlgCredRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgCredRpm

`func (o *EnvironmentLicenseQuantities) SetRlgCredRpm(v EnvironmentAPILimit)`

SetRlgCredRpm sets RlgCredRpm field to given value.

### HasRlgCredRpm

`func (o *EnvironmentLicenseQuantities) HasRlgCredRpm() bool`

HasRlgCredRpm returns a boolean if a field has been set.

### GetRlgDirAccessRps

`func (o *EnvironmentLicenseQuantities) GetRlgDirAccessRps() EnvironmentAPILimit`

GetRlgDirAccessRps returns the RlgDirAccessRps field if non-nil, zero value otherwise.

### GetRlgDirAccessRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgDirAccessRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgDirAccessRpsOk returns a tuple with the RlgDirAccessRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgDirAccessRps

`func (o *EnvironmentLicenseQuantities) SetRlgDirAccessRps(v EnvironmentAPILimit)`

SetRlgDirAccessRps sets RlgDirAccessRps field to given value.

### HasRlgDirAccessRps

`func (o *EnvironmentLicenseQuantities) HasRlgDirAccessRps() bool`

HasRlgDirAccessRps returns a boolean if a field has been set.

### GetRlgDirBulkRps

`func (o *EnvironmentLicenseQuantities) GetRlgDirBulkRps() EnvironmentAPILimit`

GetRlgDirBulkRps returns the RlgDirBulkRps field if non-nil, zero value otherwise.

### GetRlgDirBulkRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgDirBulkRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgDirBulkRpsOk returns a tuple with the RlgDirBulkRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgDirBulkRps

`func (o *EnvironmentLicenseQuantities) SetRlgDirBulkRps(v EnvironmentAPILimit)`

SetRlgDirBulkRps sets RlgDirBulkRps field to given value.

### HasRlgDirBulkRps

`func (o *EnvironmentLicenseQuantities) HasRlgDirBulkRps() bool`

HasRlgDirBulkRps returns a boolean if a field has been set.

### GetRlgDirFixedRps

`func (o *EnvironmentLicenseQuantities) GetRlgDirFixedRps() EnvironmentAPILimit`

GetRlgDirFixedRps returns the RlgDirFixedRps field if non-nil, zero value otherwise.

### GetRlgDirFixedRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgDirFixedRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgDirFixedRpsOk returns a tuple with the RlgDirFixedRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgDirFixedRps

`func (o *EnvironmentLicenseQuantities) SetRlgDirFixedRps(v EnvironmentAPILimit)`

SetRlgDirFixedRps sets RlgDirFixedRps field to given value.

### HasRlgDirFixedRps

`func (o *EnvironmentLicenseQuantities) HasRlgDirFixedRps() bool`

HasRlgDirFixedRps returns a boolean if a field has been set.

### GetRlgDirWriteRps

`func (o *EnvironmentLicenseQuantities) GetRlgDirWriteRps() EnvironmentAPILimit`

GetRlgDirWriteRps returns the RlgDirWriteRps field if non-nil, zero value otherwise.

### GetRlgDirWriteRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgDirWriteRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgDirWriteRpsOk returns a tuple with the RlgDirWriteRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgDirWriteRps

`func (o *EnvironmentLicenseQuantities) SetRlgDirWriteRps(v EnvironmentAPILimit)`

SetRlgDirWriteRps sets RlgDirWriteRps field to given value.

### HasRlgDirWriteRps

`func (o *EnvironmentLicenseQuantities) HasRlgDirWriteRps() bool`

HasRlgDirWriteRps returns a boolean if a field has been set.

### GetRlgDvFlowCallbackRps

`func (o *EnvironmentLicenseQuantities) GetRlgDvFlowCallbackRps() EnvironmentAPILimit`

GetRlgDvFlowCallbackRps returns the RlgDvFlowCallbackRps field if non-nil, zero value otherwise.

### GetRlgDvFlowCallbackRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgDvFlowCallbackRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgDvFlowCallbackRpsOk returns a tuple with the RlgDvFlowCallbackRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgDvFlowCallbackRps

`func (o *EnvironmentLicenseQuantities) SetRlgDvFlowCallbackRps(v EnvironmentAPILimit)`

SetRlgDvFlowCallbackRps sets RlgDvFlowCallbackRps field to given value.

### HasRlgDvFlowCallbackRps

`func (o *EnvironmentLicenseQuantities) HasRlgDvFlowCallbackRps() bool`

HasRlgDvFlowCallbackRps returns a boolean if a field has been set.

### GetRlgDvFlowExecRps

`func (o *EnvironmentLicenseQuantities) GetRlgDvFlowExecRps() EnvironmentAPILimit`

GetRlgDvFlowExecRps returns the RlgDvFlowExecRps field if non-nil, zero value otherwise.

### GetRlgDvFlowExecRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgDvFlowExecRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgDvFlowExecRpsOk returns a tuple with the RlgDvFlowExecRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgDvFlowExecRps

`func (o *EnvironmentLicenseQuantities) SetRlgDvFlowExecRps(v EnvironmentAPILimit)`

SetRlgDvFlowExecRps sets RlgDvFlowExecRps field to given value.

### HasRlgDvFlowExecRps

`func (o *EnvironmentLicenseQuantities) HasRlgDvFlowExecRps() bool`

HasRlgDvFlowExecRps returns a boolean if a field has been set.

### GetRlgDvFlowStartRps

`func (o *EnvironmentLicenseQuantities) GetRlgDvFlowStartRps() EnvironmentAPILimit`

GetRlgDvFlowStartRps returns the RlgDvFlowStartRps field if non-nil, zero value otherwise.

### GetRlgDvFlowStartRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgDvFlowStartRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgDvFlowStartRpsOk returns a tuple with the RlgDvFlowStartRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgDvFlowStartRps

`func (o *EnvironmentLicenseQuantities) SetRlgDvFlowStartRps(v EnvironmentAPILimit)`

SetRlgDvFlowStartRps sets RlgDvFlowStartRps field to given value.

### HasRlgDvFlowStartRps

`func (o *EnvironmentLicenseQuantities) HasRlgDvFlowStartRps() bool`

HasRlgDvFlowStartRps returns a boolean if a field has been set.

### GetRlgMfaFixedRps

`func (o *EnvironmentLicenseQuantities) GetRlgMfaFixedRps() EnvironmentAPILimit`

GetRlgMfaFixedRps returns the RlgMfaFixedRps field if non-nil, zero value otherwise.

### GetRlgMfaFixedRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgMfaFixedRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgMfaFixedRpsOk returns a tuple with the RlgMfaFixedRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgMfaFixedRps

`func (o *EnvironmentLicenseQuantities) SetRlgMfaFixedRps(v EnvironmentAPILimit)`

SetRlgMfaFixedRps sets RlgMfaFixedRps field to given value.

### HasRlgMfaFixedRps

`func (o *EnvironmentLicenseQuantities) HasRlgMfaFixedRps() bool`

HasRlgMfaFixedRps returns a boolean if a field has been set.

### GetRlgMfaPollingRps

`func (o *EnvironmentLicenseQuantities) GetRlgMfaPollingRps() EnvironmentAPILimit`

GetRlgMfaPollingRps returns the RlgMfaPollingRps field if non-nil, zero value otherwise.

### GetRlgMfaPollingRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgMfaPollingRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgMfaPollingRpsOk returns a tuple with the RlgMfaPollingRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgMfaPollingRps

`func (o *EnvironmentLicenseQuantities) SetRlgMfaPollingRps(v EnvironmentAPILimit)`

SetRlgMfaPollingRps sets RlgMfaPollingRps field to given value.

### HasRlgMfaPollingRps

`func (o *EnvironmentLicenseQuantities) HasRlgMfaPollingRps() bool`

HasRlgMfaPollingRps returns a boolean if a field has been set.

### GetRlgMfaRps

`func (o *EnvironmentLicenseQuantities) GetRlgMfaRps() EnvironmentAPILimit`

GetRlgMfaRps returns the RlgMfaRps field if non-nil, zero value otherwise.

### GetRlgMfaRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgMfaRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgMfaRpsOk returns a tuple with the RlgMfaRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgMfaRps

`func (o *EnvironmentLicenseQuantities) SetRlgMfaRps(v EnvironmentAPILimit)`

SetRlgMfaRps sets RlgMfaRps field to given value.

### HasRlgMfaRps

`func (o *EnvironmentLicenseQuantities) HasRlgMfaRps() bool`

HasRlgMfaRps returns a boolean if a field has been set.

### GetRlgRiskEvalsRps

`func (o *EnvironmentLicenseQuantities) GetRlgRiskEvalsRps() EnvironmentAPILimit`

GetRlgRiskEvalsRps returns the RlgRiskEvalsRps field if non-nil, zero value otherwise.

### GetRlgRiskEvalsRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgRiskEvalsRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgRiskEvalsRpsOk returns a tuple with the RlgRiskEvalsRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgRiskEvalsRps

`func (o *EnvironmentLicenseQuantities) SetRlgRiskEvalsRps(v EnvironmentAPILimit)`

SetRlgRiskEvalsRps sets RlgRiskEvalsRps field to given value.

### HasRlgRiskEvalsRps

`func (o *EnvironmentLicenseQuantities) HasRlgRiskEvalsRps() bool`

HasRlgRiskEvalsRps returns a boolean if a field has been set.

### GetRlgTokenCheckRps

`func (o *EnvironmentLicenseQuantities) GetRlgTokenCheckRps() EnvironmentAPILimit`

GetRlgTokenCheckRps returns the RlgTokenCheckRps field if non-nil, zero value otherwise.

### GetRlgTokenCheckRpsOk

`func (o *EnvironmentLicenseQuantities) GetRlgTokenCheckRpsOk() (*EnvironmentAPILimit, bool)`

GetRlgTokenCheckRpsOk returns a tuple with the RlgTokenCheckRps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgTokenCheckRps

`func (o *EnvironmentLicenseQuantities) SetRlgTokenCheckRps(v EnvironmentAPILimit)`

SetRlgTokenCheckRps sets RlgTokenCheckRps field to given value.

### HasRlgTokenCheckRps

`func (o *EnvironmentLicenseQuantities) HasRlgTokenCheckRps() bool`

HasRlgTokenCheckRps returns a boolean if a field has been set.

### GetRlgVerifyRpm

`func (o *EnvironmentLicenseQuantities) GetRlgVerifyRpm() EnvironmentAPILimit`

GetRlgVerifyRpm returns the RlgVerifyRpm field if non-nil, zero value otherwise.

### GetRlgVerifyRpmOk

`func (o *EnvironmentLicenseQuantities) GetRlgVerifyRpmOk() (*EnvironmentAPILimit, bool)`

GetRlgVerifyRpmOk returns a tuple with the RlgVerifyRpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRlgVerifyRpm

`func (o *EnvironmentLicenseQuantities) SetRlgVerifyRpm(v EnvironmentAPILimit)`

SetRlgVerifyRpm sets RlgVerifyRpm field to given value.

### HasRlgVerifyRpm

`func (o *EnvironmentLicenseQuantities) HasRlgVerifyRpm() bool`

HasRlgVerifyRpm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


