/*
PingOne User and Configuration Management API

The PingOne User and Configuration Management API provides the interface to configure and manage users in the PingOne directory and the administration configuration of your PingOne organization.

Contact: developerexperiences@pingidentity.com
*/

package pingone

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

// DaVinciVariableCreateRequestValue - struct for DaVinciVariableCreateRequestValue
type DaVinciVariableCreateRequestValue struct {
	ValueBoolean *bool
	ValueNumber  *float32
	ValueObject  *any
	ValueSecret  *string
	ValueString  *string
}

// ValueBooleanAsDaVinciVariableCreateRequestValue is a convenience function that returns ValueBoolean wrapped in DaVinciVariableCreateRequestValue
func ValueBooleanAsDaVinciVariableCreateRequestValue(v *bool) DaVinciVariableCreateRequestValue {
	return DaVinciVariableCreateRequestValue{
		ValueBoolean: v,
	}
}

// ValueNumberAsDaVinciVariableCreateRequestValue is a convenience function that returns ValueNumber wrapped in DaVinciVariableCreateRequestValue
func ValueNumberAsDaVinciVariableCreateRequestValue(v *float32) DaVinciVariableCreateRequestValue {
	return DaVinciVariableCreateRequestValue{
		ValueNumber: v,
	}
}

// ValueObjectAsDaVinciVariableCreateRequestValue is a convenience function that returns ValueObject wrapped in DaVinciVariableCreateRequestValue
func ValueObjectAsDaVinciVariableCreateRequestValue(v *any) DaVinciVariableCreateRequestValue {
	return DaVinciVariableCreateRequestValue{
		ValueObject: v,
	}
}

// ValueSecretAsDaVinciVariableCreateRequestValue is a convenience function that returns ValueSecret wrapped in DaVinciVariableCreateRequestValue
func ValueSecretAsDaVinciVariableCreateRequestValue(v *string) DaVinciVariableCreateRequestValue {
	return DaVinciVariableCreateRequestValue{
		ValueSecret: v,
	}
}

// ValueStringAsDaVinciVariableCreateRequestValue is a convenience function that returns ValueString wrapped in DaVinciVariableCreateRequestValue
func ValueStringAsDaVinciVariableCreateRequestValue(v *string) DaVinciVariableCreateRequestValue {
	return DaVinciVariableCreateRequestValue{
		ValueString: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *DaVinciVariableCreateRequestValue) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = newStrictDecoder(data).Decode(&jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'boolean'
	if jsonDict["dataType"] == "boolean" {
		// try to unmarshal JSON data into ValueBoolean
		err = json.Unmarshal(data, &dst.ValueBoolean)
		if err == nil {
			return nil // data stored in dst.ValueBoolean, return on the first match
		} else {
			dst.ValueBoolean = nil
			return fmt.Errorf("failed to unmarshal DaVinciVariableCreateRequestValue as ValueBoolean: %s", err.Error())
		}
	}

	// check if the discriminator value is 'number'
	if jsonDict["dataType"] == "number" {
		// try to unmarshal JSON data into ValueNumber
		err = json.Unmarshal(data, &dst.ValueNumber)
		if err == nil {
			return nil // data stored in dst.ValueNumber, return on the first match
		} else {
			dst.ValueNumber = nil
			return fmt.Errorf("failed to unmarshal DaVinciVariableCreateRequestValue as ValueNumber: %s", err.Error())
		}
	}

	// check if the discriminator value is 'object'
	if jsonDict["dataType"] == "object" {
		// try to unmarshal JSON data into ValueObject
		err = json.Unmarshal(data, &dst.ValueObject)
		if err == nil {
			return nil // data stored in dst.ValueObject, return on the first match
		} else {
			dst.ValueObject = nil
			return fmt.Errorf("failed to unmarshal DaVinciVariableCreateRequestValue as ValueObject: %s", err.Error())
		}
	}

	// check if the discriminator value is 'secret'
	if jsonDict["dataType"] == "secret" {
		// try to unmarshal JSON data into ValueSecret
		err = json.Unmarshal(data, &dst.ValueSecret)
		if err == nil {
			return nil // data stored in dst.ValueSecret, return on the first match
		} else {
			dst.ValueSecret = nil
			return fmt.Errorf("failed to unmarshal DaVinciVariableCreateRequestValue as ValueSecret: %s", err.Error())
		}
	}

	// check if the discriminator value is 'string'
	if jsonDict["dataType"] == "string" {
		// try to unmarshal JSON data into ValueString
		err = json.Unmarshal(data, &dst.ValueString)
		if err == nil {
			return nil // data stored in dst.ValueString, return on the first match
		} else {
			dst.ValueString = nil
			return fmt.Errorf("failed to unmarshal DaVinciVariableCreateRequestValue as ValueString: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DaVinciVariableCreateRequestValue) MarshalJSON() ([]byte, error) {
	if src.ValueBoolean != nil {
		return json.Marshal(&src.ValueBoolean)
	}

	if src.ValueNumber != nil {
		return json.Marshal(&src.ValueNumber)
	}

	if src.ValueObject != nil {
		return json.Marshal(&src.ValueObject)
	}

	if src.ValueSecret != nil {
		return json.Marshal(&src.ValueSecret)
	}

	if src.ValueString != nil {
		return json.Marshal(&src.ValueString)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DaVinciVariableCreateRequestValue) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ValueBoolean != nil {
		return obj.ValueBoolean
	}

	if obj.ValueNumber != nil {
		return obj.ValueNumber
	}

	if obj.ValueObject != nil {
		return obj.ValueObject
	}

	if obj.ValueSecret != nil {
		return obj.ValueSecret
	}

	if obj.ValueString != nil {
		return obj.ValueString
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj DaVinciVariableCreateRequestValue) GetActualInstanceValue() interface{} {
	if obj.ValueBoolean != nil {
		return *obj.ValueBoolean
	}

	if obj.ValueNumber != nil {
		return *obj.ValueNumber
	}

	if obj.ValueObject != nil {
		return *obj.ValueObject
	}

	if obj.ValueSecret != nil {
		return *obj.ValueSecret
	}

	if obj.ValueString != nil {
		return *obj.ValueString
	}

	// all schemas are nil
	return nil
}

func (o DaVinciVariableCreateRequestValue) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	if !IsNil(o.ValueBoolean) {
		logAttrs = append(logAttrs, slog.Any("ValueBoolean", *o.ValueBoolean))
	}
	if !IsNil(o.ValueNumber) {
		logAttrs = append(logAttrs, slog.Any("ValueNumber", *o.ValueNumber))
	}
	if !IsNil(o.ValueObject) {
		logAttrs = append(logAttrs, slog.Any("ValueObject", *o.ValueObject))
	}
	if !IsNil(o.ValueSecret) {
		logAttrs = append(logAttrs, slog.String("ValueSecret", "****"))
	}
	if !IsNil(o.ValueString) {
		logAttrs = append(logAttrs, slog.Any("ValueString", *o.ValueString))
	}

	return slog.GroupValue(logAttrs...)
}

type NullableDaVinciVariableCreateRequestValue struct {
	value *DaVinciVariableCreateRequestValue
	isSet bool
}

func (v NullableDaVinciVariableCreateRequestValue) Get() *DaVinciVariableCreateRequestValue {
	return v.value
}

func (v *NullableDaVinciVariableCreateRequestValue) Set(val *DaVinciVariableCreateRequestValue) {
	v.value = val
	v.isSet = true
}

func (v NullableDaVinciVariableCreateRequestValue) IsSet() bool {
	return v.isSet
}

func (v *NullableDaVinciVariableCreateRequestValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaVinciVariableCreateRequestValue(val *DaVinciVariableCreateRequestValue) *NullableDaVinciVariableCreateRequestValue {
	return &NullableDaVinciVariableCreateRequestValue{value: val, isSet: true}
}

func (v NullableDaVinciVariableCreateRequestValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaVinciVariableCreateRequestValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
