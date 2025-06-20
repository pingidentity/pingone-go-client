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

// EnvironmentRegion struct for EnvironmentRegion
type EnvironmentRegion struct {
	EnvironmentRegionCode *EnvironmentRegionCode
	String                *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EnvironmentRegion) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EnvironmentRegionCode
	err = json.Unmarshal(data, &dst.EnvironmentRegionCode)
	if err == nil {
		jsonEnvironmentRegionCode, _ := json.Marshal(dst.EnvironmentRegionCode)
		if string(jsonEnvironmentRegionCode) == "{}" { // empty struct
			dst.EnvironmentRegionCode = nil
		} else {
			return nil // data stored in dst.EnvironmentRegionCode, return on the first match
		}
	} else {
		dst.EnvironmentRegionCode = nil
	}

	// try to unmarshal JSON data into String
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EnvironmentRegion)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EnvironmentRegion) MarshalJSON() ([]byte, error) {
	if src.EnvironmentRegionCode != nil {
		return json.Marshal(&src.EnvironmentRegionCode)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

func (o EnvironmentRegion) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	if !IsNil(o.EnvironmentRegionCode) {
		logAttrs = append(logAttrs, slog.Any("EnvironmentRegionCode", *o.EnvironmentRegionCode))
	}
	if !IsNil(o.String) {
		logAttrs = append(logAttrs, slog.Any("String", *o.String))
	}

	return slog.GroupValue(logAttrs...)
}

type NullableEnvironmentRegion struct {
	value *EnvironmentRegion
	isSet bool
}

func (v NullableEnvironmentRegion) Get() *EnvironmentRegion {
	return v.value
}

func (v *NullableEnvironmentRegion) Set(val *EnvironmentRegion) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentRegion) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentRegion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentRegion(val *EnvironmentRegion) *NullableEnvironmentRegion {
	return &NullableEnvironmentRegion{value: val, isSet: true}
}

func (v NullableEnvironmentRegion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentRegion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
