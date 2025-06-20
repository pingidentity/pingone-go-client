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
	"log/slog"
)

// checks if the EnvironmentBillOfMaterialsProductDeployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentBillOfMaterialsProductDeployment{}

// checks if the EnvironmentBillOfMaterialsProductDeployment type satisfies the LogValuer interface at compile time
var _ slog.LogValuer = &EnvironmentBillOfMaterialsProductDeployment{}

// EnvironmentBillOfMaterialsProductDeployment struct for EnvironmentBillOfMaterialsProductDeployment
type EnvironmentBillOfMaterialsProductDeployment struct {
	Id                   *string `json:"id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentBillOfMaterialsProductDeployment EnvironmentBillOfMaterialsProductDeployment

// NewEnvironmentBillOfMaterialsProductDeployment instantiates a new EnvironmentBillOfMaterialsProductDeployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentBillOfMaterialsProductDeployment() *EnvironmentBillOfMaterialsProductDeployment {
	this := EnvironmentBillOfMaterialsProductDeployment{}
	return &this
}

// NewEnvironmentBillOfMaterialsProductDeploymentWithDefaults instantiates a new EnvironmentBillOfMaterialsProductDeployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentBillOfMaterialsProductDeploymentWithDefaults() *EnvironmentBillOfMaterialsProductDeployment {
	this := EnvironmentBillOfMaterialsProductDeployment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EnvironmentBillOfMaterialsProductDeployment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentBillOfMaterialsProductDeployment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EnvironmentBillOfMaterialsProductDeployment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EnvironmentBillOfMaterialsProductDeployment) SetId(v string) {
	o.Id = &v
}

func (o EnvironmentBillOfMaterialsProductDeployment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentBillOfMaterialsProductDeployment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentBillOfMaterialsProductDeployment) UnmarshalJSON(data []byte) (err error) {
	varEnvironmentBillOfMaterialsProductDeployment := _EnvironmentBillOfMaterialsProductDeployment{}

	err = json.Unmarshal(data, &varEnvironmentBillOfMaterialsProductDeployment)

	if err != nil {
		return err
	}

	*o = EnvironmentBillOfMaterialsProductDeployment(varEnvironmentBillOfMaterialsProductDeployment)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o EnvironmentBillOfMaterialsProductDeployment) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	if !IsNil(o.Id) {
		logAttrs = append(logAttrs, slog.Any("id", *o.Id))
	}
	logAttrs = append(logAttrs, slog.Any("additionalProperties", o.AdditionalProperties))

	return slog.GroupValue(logAttrs...)
}

type NullableEnvironmentBillOfMaterialsProductDeployment struct {
	value *EnvironmentBillOfMaterialsProductDeployment
	isSet bool
}

func (v NullableEnvironmentBillOfMaterialsProductDeployment) Get() *EnvironmentBillOfMaterialsProductDeployment {
	return v.value
}

func (v *NullableEnvironmentBillOfMaterialsProductDeployment) Set(val *EnvironmentBillOfMaterialsProductDeployment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentBillOfMaterialsProductDeployment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentBillOfMaterialsProductDeployment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentBillOfMaterialsProductDeployment(val *EnvironmentBillOfMaterialsProductDeployment) *NullableEnvironmentBillOfMaterialsProductDeployment {
	return &NullableEnvironmentBillOfMaterialsProductDeployment{value: val, isSet: true}
}

func (v NullableEnvironmentBillOfMaterialsProductDeployment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentBillOfMaterialsProductDeployment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
