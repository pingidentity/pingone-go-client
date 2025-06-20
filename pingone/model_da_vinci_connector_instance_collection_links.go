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

// checks if the DaVinciConnectorInstanceCollectionLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaVinciConnectorInstanceCollectionLinks{}

// checks if the DaVinciConnectorInstanceCollectionLinks type satisfies the LogValuer interface at compile time
var _ slog.LogValuer = &DaVinciConnectorInstanceCollectionLinks{}

// DaVinciConnectorInstanceCollectionLinks struct for DaVinciConnectorInstanceCollectionLinks
type DaVinciConnectorInstanceCollectionLinks struct {
	Environment          JSONHALLink `json:"environment"`
	Self                 JSONHALLink `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _DaVinciConnectorInstanceCollectionLinks DaVinciConnectorInstanceCollectionLinks

// NewDaVinciConnectorInstanceCollectionLinks instantiates a new DaVinciConnectorInstanceCollectionLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaVinciConnectorInstanceCollectionLinks(environment JSONHALLink, self JSONHALLink) *DaVinciConnectorInstanceCollectionLinks {
	this := DaVinciConnectorInstanceCollectionLinks{}
	this.Environment = environment
	this.Self = self
	return &this
}

// NewDaVinciConnectorInstanceCollectionLinksWithDefaults instantiates a new DaVinciConnectorInstanceCollectionLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaVinciConnectorInstanceCollectionLinksWithDefaults() *DaVinciConnectorInstanceCollectionLinks {
	this := DaVinciConnectorInstanceCollectionLinks{}
	return &this
}

// GetEnvironment returns the Environment field value
func (o *DaVinciConnectorInstanceCollectionLinks) GetEnvironment() JSONHALLink {
	if o == nil {
		var ret JSONHALLink
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceCollectionLinks) GetEnvironmentOk() (*JSONHALLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *DaVinciConnectorInstanceCollectionLinks) SetEnvironment(v JSONHALLink) {
	o.Environment = v
}

// GetSelf returns the Self field value
func (o *DaVinciConnectorInstanceCollectionLinks) GetSelf() JSONHALLink {
	if o == nil {
		var ret JSONHALLink
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceCollectionLinks) GetSelfOk() (*JSONHALLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *DaVinciConnectorInstanceCollectionLinks) SetSelf(v JSONHALLink) {
	o.Self = v
}

func (o DaVinciConnectorInstanceCollectionLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaVinciConnectorInstanceCollectionLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["environment"] = o.Environment
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DaVinciConnectorInstanceCollectionLinks) UnmarshalJSON(data []byte) (err error) {
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

	varDaVinciConnectorInstanceCollectionLinks := _DaVinciConnectorInstanceCollectionLinks{}

	err = json.Unmarshal(data, &varDaVinciConnectorInstanceCollectionLinks)

	if err != nil {
		return err
	}

	*o = DaVinciConnectorInstanceCollectionLinks(varDaVinciConnectorInstanceCollectionLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "environment")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o DaVinciConnectorInstanceCollectionLinks) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	logAttrs = append(logAttrs, slog.Any("environment", o.Environment))
	logAttrs = append(logAttrs, slog.Any("self", o.Self))
	logAttrs = append(logAttrs, slog.Any("additionalProperties", o.AdditionalProperties))

	return slog.GroupValue(logAttrs...)
}

type NullableDaVinciConnectorInstanceCollectionLinks struct {
	value *DaVinciConnectorInstanceCollectionLinks
	isSet bool
}

func (v NullableDaVinciConnectorInstanceCollectionLinks) Get() *DaVinciConnectorInstanceCollectionLinks {
	return v.value
}

func (v *NullableDaVinciConnectorInstanceCollectionLinks) Set(val *DaVinciConnectorInstanceCollectionLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDaVinciConnectorInstanceCollectionLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDaVinciConnectorInstanceCollectionLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaVinciConnectorInstanceCollectionLinks(val *DaVinciConnectorInstanceCollectionLinks) *NullableDaVinciConnectorInstanceCollectionLinks {
	return &NullableDaVinciConnectorInstanceCollectionLinks{value: val, isSet: true}
}

func (v NullableDaVinciConnectorInstanceCollectionLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaVinciConnectorInstanceCollectionLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
