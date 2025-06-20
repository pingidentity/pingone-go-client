// Copyright © 2025 Ping Identity Corporation
/*
PingOne User and Configuration Management API

The PingOne User and Configuration Management API provides the interface to configure and manage users in the PingOne directory and the administration configuration of your PingOne organization.

Contact: developerexperiences@pingidentity.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pingone

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"
)

// checks if the DaVinciFlowPolicyResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaVinciFlowPolicyResponse{}

// checks if the DaVinciFlowPolicyResponse type satisfies the LogValuer interface at compile time
var _ slog.LogValuer = &DaVinciFlowPolicyResponse{}

// DaVinciFlowPolicyResponse struct for DaVinciFlowPolicyResponse
type DaVinciFlowPolicyResponse struct {
	Links                DaVinciFlowPolicyResponseLinks              `json:"_links"`
	CreatedAt            *time.Time                                  `json:"createdAt,omitempty"`
	Environment          ResourceRelationshipPingOne                 `json:"environment"`
	FlowDistributions    []DaVinciFlowPolicyResponseFlowDistribution `json:"flowDistributions"`
	Id                   string                                      `json:"id"`
	Name                 string                                      `json:"name"`
	Status               DaVinciFlowPolicyCreateRequestStatus        `json:"status"`
	Trigger              NullableDaVinciFlowPolicyResponseTrigger    `json:"trigger,omitempty"`
	UpdatedAt            *time.Time                                  `json:"updatedAt,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DaVinciFlowPolicyResponse DaVinciFlowPolicyResponse

// NewDaVinciFlowPolicyResponse instantiates a new DaVinciFlowPolicyResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaVinciFlowPolicyResponse(links DaVinciFlowPolicyResponseLinks, environment ResourceRelationshipPingOne, flowDistributions []DaVinciFlowPolicyResponseFlowDistribution, id string, name string, status DaVinciFlowPolicyCreateRequestStatus) *DaVinciFlowPolicyResponse {
	this := DaVinciFlowPolicyResponse{}
	this.Links = links
	this.Environment = environment
	this.FlowDistributions = flowDistributions
	this.Id = id
	this.Name = name
	this.Status = status
	return &this
}

// NewDaVinciFlowPolicyResponseWithDefaults instantiates a new DaVinciFlowPolicyResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaVinciFlowPolicyResponseWithDefaults() *DaVinciFlowPolicyResponse {
	this := DaVinciFlowPolicyResponse{}
	return &this
}

// GetLinks returns the Links field value
func (o *DaVinciFlowPolicyResponse) GetLinks() DaVinciFlowPolicyResponseLinks {
	if o == nil {
		var ret DaVinciFlowPolicyResponseLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *DaVinciFlowPolicyResponse) GetLinksOk() (*DaVinciFlowPolicyResponseLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *DaVinciFlowPolicyResponse) SetLinks(v DaVinciFlowPolicyResponseLinks) {
	o.Links = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DaVinciFlowPolicyResponse) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaVinciFlowPolicyResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DaVinciFlowPolicyResponse) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DaVinciFlowPolicyResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEnvironment returns the Environment field value
func (o *DaVinciFlowPolicyResponse) GetEnvironment() ResourceRelationshipPingOne {
	if o == nil {
		var ret ResourceRelationshipPingOne
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *DaVinciFlowPolicyResponse) GetEnvironmentOk() (*ResourceRelationshipPingOne, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *DaVinciFlowPolicyResponse) SetEnvironment(v ResourceRelationshipPingOne) {
	o.Environment = v
}

// GetFlowDistributions returns the FlowDistributions field value
func (o *DaVinciFlowPolicyResponse) GetFlowDistributions() []DaVinciFlowPolicyResponseFlowDistribution {
	if o == nil {
		var ret []DaVinciFlowPolicyResponseFlowDistribution
		return ret
	}

	return o.FlowDistributions
}

// GetFlowDistributionsOk returns a tuple with the FlowDistributions field value
// and a boolean to check if the value has been set.
func (o *DaVinciFlowPolicyResponse) GetFlowDistributionsOk() ([]DaVinciFlowPolicyResponseFlowDistribution, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowDistributions, true
}

// SetFlowDistributions sets field value
func (o *DaVinciFlowPolicyResponse) SetFlowDistributions(v []DaVinciFlowPolicyResponseFlowDistribution) {
	o.FlowDistributions = v
}

// GetId returns the Id field value
func (o *DaVinciFlowPolicyResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DaVinciFlowPolicyResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DaVinciFlowPolicyResponse) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DaVinciFlowPolicyResponse) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DaVinciFlowPolicyResponse) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DaVinciFlowPolicyResponse) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *DaVinciFlowPolicyResponse) GetStatus() DaVinciFlowPolicyCreateRequestStatus {
	if o == nil {
		var ret DaVinciFlowPolicyCreateRequestStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *DaVinciFlowPolicyResponse) GetStatusOk() (*DaVinciFlowPolicyCreateRequestStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *DaVinciFlowPolicyResponse) SetStatus(v DaVinciFlowPolicyCreateRequestStatus) {
	o.Status = v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DaVinciFlowPolicyResponse) GetTrigger() DaVinciFlowPolicyResponseTrigger {
	if o == nil || IsNil(o.Trigger.Get()) {
		var ret DaVinciFlowPolicyResponseTrigger
		return ret
	}
	return *o.Trigger.Get()
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DaVinciFlowPolicyResponse) GetTriggerOk() (*DaVinciFlowPolicyResponseTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return o.Trigger.Get(), o.Trigger.IsSet()
}

// HasTrigger returns a boolean if a field has been set.
func (o *DaVinciFlowPolicyResponse) HasTrigger() bool {
	if o != nil && o.Trigger.IsSet() {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given NullableDaVinciFlowPolicyResponseTrigger and assigns it to the Trigger field.
func (o *DaVinciFlowPolicyResponse) SetTrigger(v DaVinciFlowPolicyResponseTrigger) {
	o.Trigger.Set(&v)
}

// SetTriggerNil sets the value for Trigger to be an explicit nil
func (o *DaVinciFlowPolicyResponse) SetTriggerNil() {
	o.Trigger.Set(nil)
}

// UnsetTrigger ensures that no value is present for Trigger, not even an explicit nil
func (o *DaVinciFlowPolicyResponse) UnsetTrigger() {
	o.Trigger.Unset()
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DaVinciFlowPolicyResponse) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaVinciFlowPolicyResponse) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DaVinciFlowPolicyResponse) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DaVinciFlowPolicyResponse) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DaVinciFlowPolicyResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaVinciFlowPolicyResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_links"] = o.Links
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["environment"] = o.Environment
	toSerialize["flowDistributions"] = o.FlowDistributions
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if o.Trigger.IsSet() {
		toSerialize["trigger"] = o.Trigger.Get()
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DaVinciFlowPolicyResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"_links",
		"environment",
		"flowDistributions",
		"id",
		"name",
		"status",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDaVinciFlowPolicyResponse := _DaVinciFlowPolicyResponse{}

	err = json.Unmarshal(data, &varDaVinciFlowPolicyResponse)

	if err != nil {
		return err
	}

	*o = DaVinciFlowPolicyResponse(varDaVinciFlowPolicyResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "_links")
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "flowDistributions")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "status")
		delete(additionalProperties, "trigger")
		delete(additionalProperties, "updatedAt")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o DaVinciFlowPolicyResponse) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	logAttrs = append(logAttrs, slog.Any("_links", o.Links))
	if !IsNil(o.CreatedAt) {
		logAttrs = append(logAttrs, slog.Any("createdAt", *o.CreatedAt))
	}
	logAttrs = append(logAttrs, slog.Any("environment", o.Environment))
	logAttrs = append(logAttrs, slog.Any("flowDistributions", o.FlowDistributions))
	logAttrs = append(logAttrs, slog.Any("id", o.Id))
	logAttrs = append(logAttrs, slog.Any("name", o.Name))
	logAttrs = append(logAttrs, slog.Any("status", o.Status))
	if !IsNil(o.Trigger) {
		logAttrs = append(logAttrs, slog.Any("trigger", o.Trigger))
	}
	if !IsNil(o.UpdatedAt) {
		logAttrs = append(logAttrs, slog.Any("updatedAt", *o.UpdatedAt))
	}
	logAttrs = append(logAttrs, slog.Any("additionalProperties", o.AdditionalProperties))

	return slog.GroupValue(logAttrs...)
}

type NullableDaVinciFlowPolicyResponse struct {
	value *DaVinciFlowPolicyResponse
	isSet bool
}

func (v NullableDaVinciFlowPolicyResponse) Get() *DaVinciFlowPolicyResponse {
	return v.value
}

func (v *NullableDaVinciFlowPolicyResponse) Set(val *DaVinciFlowPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDaVinciFlowPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDaVinciFlowPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaVinciFlowPolicyResponse(val *DaVinciFlowPolicyResponse) *NullableDaVinciFlowPolicyResponse {
	return &NullableDaVinciFlowPolicyResponse{value: val, isSet: true}
}

func (v NullableDaVinciFlowPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaVinciFlowPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
