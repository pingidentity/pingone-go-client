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

// checks if the EnvironmentLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnvironmentLinks{}

// checks if the EnvironmentLinks type satisfies the LogValuer interface at compile time
var _ slog.LogValuer = &EnvironmentLinks{}

// EnvironmentLinks struct for EnvironmentLinks
type EnvironmentLinks struct {
	Self                 *JSONHALLink `json:"self,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EnvironmentLinks EnvironmentLinks

// NewEnvironmentLinks instantiates a new EnvironmentLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentLinks() *EnvironmentLinks {
	this := EnvironmentLinks{}
	return &this
}

// NewEnvironmentLinksWithDefaults instantiates a new EnvironmentLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentLinksWithDefaults() *EnvironmentLinks {
	this := EnvironmentLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *EnvironmentLinks) GetSelf() JSONHALLink {
	if o == nil || IsNil(o.Self) {
		var ret JSONHALLink
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentLinks) GetSelfOk() (*JSONHALLink, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *EnvironmentLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given JSONHALLink and assigns it to the Self field.
func (o *EnvironmentLinks) SetSelf(v JSONHALLink) {
	o.Self = &v
}

func (o EnvironmentLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnvironmentLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EnvironmentLinks) UnmarshalJSON(data []byte) (err error) {
	varEnvironmentLinks := _EnvironmentLinks{}

	err = json.Unmarshal(data, &varEnvironmentLinks)

	if err != nil {
		return err
	}

	*o = EnvironmentLinks(varEnvironmentLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o EnvironmentLinks) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	if !IsNil(o.Self) {
		logAttrs = append(logAttrs, slog.Any("self", *o.Self))
	}
	logAttrs = append(logAttrs, slog.Any("additionalProperties", o.AdditionalProperties))

	return slog.GroupValue(logAttrs...)
}

type NullableEnvironmentLinks struct {
	value *EnvironmentLinks
	isSet bool
}

func (v NullableEnvironmentLinks) Get() *EnvironmentLinks {
	return v.value
}

func (v *NullableEnvironmentLinks) Set(val *EnvironmentLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentLinks(val *EnvironmentLinks) *NullableEnvironmentLinks {
	return &NullableEnvironmentLinks{value: val, isSet: true}
}

func (v NullableEnvironmentLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
