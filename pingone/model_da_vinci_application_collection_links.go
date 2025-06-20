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

// checks if the DaVinciApplicationCollectionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaVinciApplicationCollectionLinks{}

// checks if the DaVinciApplicationCollectionLinks type satisfies the LogValuer interface at compile time
var _ slog.LogValuer = &DaVinciApplicationCollectionLinks{}

// DaVinciApplicationCollectionLinks struct for DaVinciApplicationCollectionLinks
type DaVinciApplicationCollectionLinks struct {
	Environment          JSONHALLink `json:"environment"`
	Self                 JSONHALLink `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _DaVinciApplicationCollectionLinks DaVinciApplicationCollectionLinks

// NewDaVinciApplicationCollectionLinks instantiates a new DaVinciApplicationCollectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaVinciApplicationCollectionLinks(environment JSONHALLink, self JSONHALLink) *DaVinciApplicationCollectionLinks {
	this := DaVinciApplicationCollectionLinks{}
	this.Environment = environment
	this.Self = self
	return &this
}

// NewDaVinciApplicationCollectionLinksWithDefaults instantiates a new DaVinciApplicationCollectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaVinciApplicationCollectionLinksWithDefaults() *DaVinciApplicationCollectionLinks {
	this := DaVinciApplicationCollectionLinks{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *DaVinciApplicationCollectionLinks) GetEnvironment() JSONHALLink {
	if o == nil {
		var ret JSONHALLink
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *DaVinciApplicationCollectionLinks) GetEnvironmentOk() (*JSONHALLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *DaVinciApplicationCollectionLinks) SetEnvironment(v JSONHALLink) {
	o.Environment = v
}

// GetSelf returns the Self field value
func (o *DaVinciApplicationCollectionLinks) GetSelf() JSONHALLink {
	if o == nil {
		var ret JSONHALLink
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *DaVinciApplicationCollectionLinks) GetSelfOk() (*JSONHALLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *DaVinciApplicationCollectionLinks) SetSelf(v JSONHALLink) {
	o.Self = v
}

func (o DaVinciApplicationCollectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaVinciApplicationCollectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment"] = o.Environment
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DaVinciApplicationCollectionLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"environment",
		"self",
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

	varDaVinciApplicationCollectionLinks := _DaVinciApplicationCollectionLinks{}

	err = json.Unmarshal(data, &varDaVinciApplicationCollectionLinks)

	if err != nil {
		return err
	}

	*o = DaVinciApplicationCollectionLinks(varDaVinciApplicationCollectionLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environment")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o DaVinciApplicationCollectionLinks) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	logAttrs = append(logAttrs, slog.Any("environment", o.Environment))
	logAttrs = append(logAttrs, slog.Any("self", o.Self))
	logAttrs = append(logAttrs, slog.Any("additionalProperties", o.AdditionalProperties))

	return slog.GroupValue(logAttrs...)
}

type NullableDaVinciApplicationCollectionLinks struct {
	value *DaVinciApplicationCollectionLinks
	isSet bool
}

func (v NullableDaVinciApplicationCollectionLinks) Get() *DaVinciApplicationCollectionLinks {
	return v.value
}

func (v *NullableDaVinciApplicationCollectionLinks) Set(val *DaVinciApplicationCollectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDaVinciApplicationCollectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDaVinciApplicationCollectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaVinciApplicationCollectionLinks(val *DaVinciApplicationCollectionLinks) *NullableDaVinciApplicationCollectionLinks {
	return &NullableDaVinciApplicationCollectionLinks{value: val, isSet: true}
}

func (v NullableDaVinciApplicationCollectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaVinciApplicationCollectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
