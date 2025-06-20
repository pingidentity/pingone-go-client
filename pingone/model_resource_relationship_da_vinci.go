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
)

// checks if the ResourceRelationshipDaVinci type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceRelationshipDaVinci{}

// checks if the ResourceRelationshipDaVinci type satisfies the LogValuer interface at compile time
var _ slog.LogValuer = &ResourceRelationshipDaVinci{}

// ResourceRelationshipDaVinci struct for ResourceRelationshipDaVinci
type ResourceRelationshipDaVinci struct {
	Id                   string `json:"id" validate:"regexp=^(?=\\\\S)[\\\\p{L}\\\\p{M}\\\\p{N}\\\\p{So}\\/.'_ -]*(?!.*((<)|(\\\\$\\\\{)))"`
	AdditionalProperties map[string]interface{}
}

type _ResourceRelationshipDaVinci ResourceRelationshipDaVinci

// NewResourceRelationshipDaVinci instantiates a new ResourceRelationshipDaVinci object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceRelationshipDaVinci(id string) *ResourceRelationshipDaVinci {
	this := ResourceRelationshipDaVinci{}
	this.Id = id
	return &this
}

// NewResourceRelationshipDaVinciWithDefaults instantiates a new ResourceRelationshipDaVinci object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceRelationshipDaVinciWithDefaults() *ResourceRelationshipDaVinci {
	this := ResourceRelationshipDaVinci{}
	return &this
}

// GetId returns the Id field value
func (o *ResourceRelationshipDaVinci) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ResourceRelationshipDaVinci) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ResourceRelationshipDaVinci) SetId(v string) {
	o.Id = v
}

func (o ResourceRelationshipDaVinci) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceRelationshipDaVinci) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceRelationshipDaVinci) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varResourceRelationshipDaVinci := _ResourceRelationshipDaVinci{}

	err = json.Unmarshal(data, &varResourceRelationshipDaVinci)

	if err != nil {
		return err
	}

	*o = ResourceRelationshipDaVinci(varResourceRelationshipDaVinci)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o ResourceRelationshipDaVinci) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	logAttrs = append(logAttrs, slog.Any("id", o.Id))
	logAttrs = append(logAttrs, slog.Any("additionalProperties", o.AdditionalProperties))

	return slog.GroupValue(logAttrs...)
}

type NullableResourceRelationshipDaVinci struct {
	value *ResourceRelationshipDaVinci
	isSet bool
}

func (v NullableResourceRelationshipDaVinci) Get() *ResourceRelationshipDaVinci {
	return v.value
}

func (v *NullableResourceRelationshipDaVinci) Set(val *ResourceRelationshipDaVinci) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceRelationshipDaVinci) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceRelationshipDaVinci) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceRelationshipDaVinci(val *ResourceRelationshipDaVinci) *NullableResourceRelationshipDaVinci {
	return &NullableResourceRelationshipDaVinci{value: val, isSet: true}
}

func (v NullableResourceRelationshipDaVinci) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceRelationshipDaVinci) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
