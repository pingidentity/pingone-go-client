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

// checks if the DaVinciConnectorInstanceLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DaVinciConnectorInstanceLinks{}

// checks if the DaVinciConnectorInstanceLinks type satisfies the LogValuer interface at compile time
var _ slog.LogValuer = &DaVinciConnectorInstanceLinks{}

// DaVinciConnectorInstanceLinks struct for DaVinciConnectorInstanceLinks
type DaVinciConnectorInstanceLinks struct {
	Applications                 *JSONHALLink `json:"applications,omitempty"`
	ConnectorInstanceClone       JSONHALLink  `json:"connectorInstance.clone"`
	DeviceAuthenticationPolicies *JSONHALLink `json:"deviceAuthenticationPolicies,omitempty"`
	Environment                  JSONHALLink  `json:"environment"`
	Gateways                     *JSONHALLink `json:"gateways,omitempty"`
	NotificationsPolicies        *JSONHALLink `json:"notificationsPolicies,omitempty"`
	Self                         JSONHALLink  `json:"self"`
	AdditionalProperties         map[string]interface{}
}

type _DaVinciConnectorInstanceLinks DaVinciConnectorInstanceLinks

// NewDaVinciConnectorInstanceLinks instantiates a new DaVinciConnectorInstanceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDaVinciConnectorInstanceLinks(connectorInstanceClone JSONHALLink, environment JSONHALLink, self JSONHALLink) *DaVinciConnectorInstanceLinks {
	this := DaVinciConnectorInstanceLinks{}
	this.ConnectorInstanceClone = connectorInstanceClone
	this.Environment = environment
	this.Self = self
	return &this
}

// NewDaVinciConnectorInstanceLinksWithDefaults instantiates a new DaVinciConnectorInstanceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDaVinciConnectorInstanceLinksWithDefaults() *DaVinciConnectorInstanceLinks {
	this := DaVinciConnectorInstanceLinks{}
	return &this
}

// GetApplications returns the Applications field value if set, zero value otherwise.
func (o *DaVinciConnectorInstanceLinks) GetApplications() JSONHALLink {
	if o == nil || IsNil(o.Applications) {
		var ret JSONHALLink
		return ret
	}
	return *o.Applications
}

// GetApplicationsOk returns a tuple with the Applications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceLinks) GetApplicationsOk() (*JSONHALLink, bool) {
	if o == nil || IsNil(o.Applications) {
		return nil, false
	}
	return o.Applications, true
}

// HasApplications returns a boolean if a field has been set.
func (o *DaVinciConnectorInstanceLinks) HasApplications() bool {
	if o != nil && !IsNil(o.Applications) {
		return true
	}

	return false
}

// SetApplications gets a reference to the given JSONHALLink and assigns it to the Applications field.
func (o *DaVinciConnectorInstanceLinks) SetApplications(v JSONHALLink) {
	o.Applications = &v
}

// GetConnectorInstanceClone returns the ConnectorInstanceClone field value
func (o *DaVinciConnectorInstanceLinks) GetConnectorInstanceClone() JSONHALLink {
	if o == nil {
		var ret JSONHALLink
		return ret
	}

	return o.ConnectorInstanceClone
}

// GetConnectorInstanceCloneOk returns a tuple with the ConnectorInstanceClone field value
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceLinks) GetConnectorInstanceCloneOk() (*JSONHALLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectorInstanceClone, true
}

// SetConnectorInstanceClone sets field value
func (o *DaVinciConnectorInstanceLinks) SetConnectorInstanceClone(v JSONHALLink) {
	o.ConnectorInstanceClone = v
}

// GetDeviceAuthenticationPolicies returns the DeviceAuthenticationPolicies field value if set, zero value otherwise.
func (o *DaVinciConnectorInstanceLinks) GetDeviceAuthenticationPolicies() JSONHALLink {
	if o == nil || IsNil(o.DeviceAuthenticationPolicies) {
		var ret JSONHALLink
		return ret
	}
	return *o.DeviceAuthenticationPolicies
}

// GetDeviceAuthenticationPoliciesOk returns a tuple with the DeviceAuthenticationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceLinks) GetDeviceAuthenticationPoliciesOk() (*JSONHALLink, bool) {
	if o == nil || IsNil(o.DeviceAuthenticationPolicies) {
		return nil, false
	}
	return o.DeviceAuthenticationPolicies, true
}

// HasDeviceAuthenticationPolicies returns a boolean if a field has been set.
func (o *DaVinciConnectorInstanceLinks) HasDeviceAuthenticationPolicies() bool {
	if o != nil && !IsNil(o.DeviceAuthenticationPolicies) {
		return true
	}

	return false
}

// SetDeviceAuthenticationPolicies gets a reference to the given JSONHALLink and assigns it to the DeviceAuthenticationPolicies field.
func (o *DaVinciConnectorInstanceLinks) SetDeviceAuthenticationPolicies(v JSONHALLink) {
	o.DeviceAuthenticationPolicies = &v
}

// GetEnvironment returns the Environment field value
func (o *DaVinciConnectorInstanceLinks) GetEnvironment() JSONHALLink {
	if o == nil {
		var ret JSONHALLink
		return ret
	}

	return o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceLinks) GetEnvironmentOk() (*JSONHALLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Environment, true
}

// SetEnvironment sets field value
func (o *DaVinciConnectorInstanceLinks) SetEnvironment(v JSONHALLink) {
	o.Environment = v
}

// GetGateways returns the Gateways field value if set, zero value otherwise.
func (o *DaVinciConnectorInstanceLinks) GetGateways() JSONHALLink {
	if o == nil || IsNil(o.Gateways) {
		var ret JSONHALLink
		return ret
	}
	return *o.Gateways
}

// GetGatewaysOk returns a tuple with the Gateways field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceLinks) GetGatewaysOk() (*JSONHALLink, bool) {
	if o == nil || IsNil(o.Gateways) {
		return nil, false
	}
	return o.Gateways, true
}

// HasGateways returns a boolean if a field has been set.
func (o *DaVinciConnectorInstanceLinks) HasGateways() bool {
	if o != nil && !IsNil(o.Gateways) {
		return true
	}

	return false
}

// SetGateways gets a reference to the given JSONHALLink and assigns it to the Gateways field.
func (o *DaVinciConnectorInstanceLinks) SetGateways(v JSONHALLink) {
	o.Gateways = &v
}

// GetNotificationsPolicies returns the NotificationsPolicies field value if set, zero value otherwise.
func (o *DaVinciConnectorInstanceLinks) GetNotificationsPolicies() JSONHALLink {
	if o == nil || IsNil(o.NotificationsPolicies) {
		var ret JSONHALLink
		return ret
	}
	return *o.NotificationsPolicies
}

// GetNotificationsPoliciesOk returns a tuple with the NotificationsPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceLinks) GetNotificationsPoliciesOk() (*JSONHALLink, bool) {
	if o == nil || IsNil(o.NotificationsPolicies) {
		return nil, false
	}
	return o.NotificationsPolicies, true
}

// HasNotificationsPolicies returns a boolean if a field has been set.
func (o *DaVinciConnectorInstanceLinks) HasNotificationsPolicies() bool {
	if o != nil && !IsNil(o.NotificationsPolicies) {
		return true
	}

	return false
}

// SetNotificationsPolicies gets a reference to the given JSONHALLink and assigns it to the NotificationsPolicies field.
func (o *DaVinciConnectorInstanceLinks) SetNotificationsPolicies(v JSONHALLink) {
	o.NotificationsPolicies = &v
}

// GetSelf returns the Self field value
func (o *DaVinciConnectorInstanceLinks) GetSelf() JSONHALLink {
	if o == nil {
		var ret JSONHALLink
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *DaVinciConnectorInstanceLinks) GetSelfOk() (*JSONHALLink, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *DaVinciConnectorInstanceLinks) SetSelf(v JSONHALLink) {
	o.Self = v
}

func (o DaVinciConnectorInstanceLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DaVinciConnectorInstanceLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Applications) {
		toSerialize["applications"] = o.Applications
	}
	toSerialize["connectorInstance.clone"] = o.ConnectorInstanceClone
	if !IsNil(o.DeviceAuthenticationPolicies) {
		toSerialize["deviceAuthenticationPolicies"] = o.DeviceAuthenticationPolicies
	}
	toSerialize["environment"] = o.Environment
	if !IsNil(o.Gateways) {
		toSerialize["gateways"] = o.Gateways
	}
	if !IsNil(o.NotificationsPolicies) {
		toSerialize["notificationsPolicies"] = o.NotificationsPolicies
	}
	toSerialize["self"] = o.Self

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DaVinciConnectorInstanceLinks) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"connectorInstance.clone",
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

	varDaVinciConnectorInstanceLinks := _DaVinciConnectorInstanceLinks{}

	err = json.Unmarshal(data, &varDaVinciConnectorInstanceLinks)

	if err != nil {
		return err
	}

	*o = DaVinciConnectorInstanceLinks(varDaVinciConnectorInstanceLinks)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "applications")
		delete(additionalProperties, "connectorInstance.clone")
		delete(additionalProperties, "deviceAuthenticationPolicies")
		delete(additionalProperties, "environment")
		delete(additionalProperties, "gateways")
		delete(additionalProperties, "notificationsPolicies")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

func (o DaVinciConnectorInstanceLinks) LogValue() slog.Value {
	logAttrs := make([]slog.Attr, 0)

	if !IsNil(o.Applications) {
		logAttrs = append(logAttrs, slog.Any("applications", *o.Applications))
	}
	logAttrs = append(logAttrs, slog.Any("connectorInstance.clone", o.ConnectorInstanceClone))
	if !IsNil(o.DeviceAuthenticationPolicies) {
		logAttrs = append(logAttrs, slog.Any("deviceAuthenticationPolicies", *o.DeviceAuthenticationPolicies))
	}
	logAttrs = append(logAttrs, slog.Any("environment", o.Environment))
	if !IsNil(o.Gateways) {
		logAttrs = append(logAttrs, slog.Any("gateways", *o.Gateways))
	}
	if !IsNil(o.NotificationsPolicies) {
		logAttrs = append(logAttrs, slog.Any("notificationsPolicies", *o.NotificationsPolicies))
	}
	logAttrs = append(logAttrs, slog.Any("self", o.Self))
	logAttrs = append(logAttrs, slog.Any("additionalProperties", o.AdditionalProperties))

	return slog.GroupValue(logAttrs...)
}

type NullableDaVinciConnectorInstanceLinks struct {
	value *DaVinciConnectorInstanceLinks
	isSet bool
}

func (v NullableDaVinciConnectorInstanceLinks) Get() *DaVinciConnectorInstanceLinks {
	return v.value
}

func (v *NullableDaVinciConnectorInstanceLinks) Set(val *DaVinciConnectorInstanceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDaVinciConnectorInstanceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDaVinciConnectorInstanceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDaVinciConnectorInstanceLinks(val *DaVinciConnectorInstanceLinks) *NullableDaVinciConnectorInstanceLinks {
	return &NullableDaVinciConnectorInstanceLinks{value: val, isSet: true}
}

func (v NullableDaVinciConnectorInstanceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDaVinciConnectorInstanceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
