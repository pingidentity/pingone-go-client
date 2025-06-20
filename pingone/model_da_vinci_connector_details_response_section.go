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

// checks if the DaVinciConnectorDetailsResponseSection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaVinciConnectorDetailsResponseSection{}

// checks if the DaVinciConnectorDetailsResponseSection type satisfies the LogValuer interface at compile time
var _ slog.LogValuer = &DaVinciConnectorDetailsResponseSection{}

// DaVinciConnectorDetailsResponseSection struct for DaVinciConnectorDetailsResponseSection
type DaVinciConnectorDetailsResponseSection struct {
	Default              *bool  `json:"default,omitempty"`
	Name                 string `json:"name"`
	Value                string `json:"value"`
	AdditionalProperties map[string]interface{}
}

type _DaVinciConnectorDetailsResponseSection DaVinciConnectorDetailsResponseSection

// NewDaVinciConnectorDetailsResponseSection instantiates a new DaVinciConnectorDetailsResponseSection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaVinciConnectorDetailsResponseSection(name string, value string) *DaVinciConnectorDetailsResponseSection {
	this := DaVinciConnectorDetailsResponseSection{}
	this.Name = name
	this.Value = value
	return &this
}

// NewDaVinciConnectorDetailsResponseSectionWithDefaults instantiates a new DaVinciConnectorDetailsResponseSection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaVinciConnectorDetailsResponseSectionWithDefaults() *DaVinciConnectorDetailsResponseSection {
	this := DaVinciConnectorDetailsResponseSection{}
	return &this
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *DaVinciConnectorDetailsResponseSection) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorDetailsResponseSection) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *DaVinciConnectorDetailsResponseSection) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *DaVinciConnectorDetailsResponseSection) SetDefault(v bool) {
	o.Default = &v
}

// GetName returns the Name field value
func (o *DaVinciConnectorDetailsResponseSection) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorDetailsResponseSection) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DaVinciConnectorDetailsResponseSection) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *DaVinciConnectorDetailsResponseSection) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorDetailsResponseSection) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *DaVinciConnectorDetailsResponseSection) SetValue(v string) {
	o.Value = v
}

func (o DaVinciConnectorDetailsResponseSection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaVinciConnectorDetailsResponseSection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DaVinciConnectorDetailsResponseSection) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"value",
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

	varDaVinciConnectorDetailsResponseSection := _DaVinciConnectorDetailsResponseSection{}

	err = json.Unmarshal(data, &varDaVinciConnectorDetailsResponseSection)

	if err != nil {
		return err
	}

	*o = DaVinciConnectorDetailsResponseSection(varDaVinciConnectorDetailsResponseSection)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "default")
		delete(additionalProperties, "name")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o DaVinciConnectorDetailsResponseSection) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	if !IsNil(o.Default) {
		logAttrs = append(logAttrs, slog.Any("default", *o.Default))
	}
	logAttrs = append(logAttrs, slog.Any("name", o.Name))
	logAttrs = append(logAttrs, slog.Any("value", o.Value))
	logAttrs = append(logAttrs, slog.Any("additionalProperties", o.AdditionalProperties))

	return slog.GroupValue(logAttrs...)
}

type NullableDaVinciConnectorDetailsResponseSection struct {
	value *DaVinciConnectorDetailsResponseSection
	isSet bool
}

func (v NullableDaVinciConnectorDetailsResponseSection) Get() *DaVinciConnectorDetailsResponseSection {
	return v.value
}

func (v *NullableDaVinciConnectorDetailsResponseSection) Set(val *DaVinciConnectorDetailsResponseSection) {
	v.value = val
	v.isSet = true
}

func (v NullableDaVinciConnectorDetailsResponseSection) IsSet() bool {
	return v.isSet
}

func (v *NullableDaVinciConnectorDetailsResponseSection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaVinciConnectorDetailsResponseSection(val *DaVinciConnectorDetailsResponseSection) *NullableDaVinciConnectorDetailsResponseSection {
	return &NullableDaVinciConnectorDetailsResponseSection{value: val, isSet: true}
}

func (v NullableDaVinciConnectorDetailsResponseSection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaVinciConnectorDetailsResponseSection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
