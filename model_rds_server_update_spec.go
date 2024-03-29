/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2111
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// RDSServerUpdateSpec Information required to update an RDS Server.
type RDSServerUpdateSpec struct {
	// Indicates if RDS server is enabled.
	Enabled bool `json:"enabled"`
	// Maximum number of sessions for RDS server as configured by administrator.
	MaxSessionsCountConfigured *int32 `json:"max_sessions_count_configured,omitempty"`
	// The configured RDS Server max sessions type. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions. * UNCONFIGURED: The max number of sessions has not yet been defined for the RDSServer.
	MaxSessionsTypeConfigured string `json:"max_sessions_type_configured"`
}

// NewRDSServerUpdateSpec instantiates a new RDSServerUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRDSServerUpdateSpec(enabled bool, maxSessionsTypeConfigured string) *RDSServerUpdateSpec {
	this := RDSServerUpdateSpec{}
	this.Enabled = enabled
	this.MaxSessionsTypeConfigured = maxSessionsTypeConfigured
	return &this
}

// NewRDSServerUpdateSpecWithDefaults instantiates a new RDSServerUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRDSServerUpdateSpecWithDefaults() *RDSServerUpdateSpec {
	this := RDSServerUpdateSpec{}
	return &this
}

// GetEnabled returns the Enabled field value
func (o *RDSServerUpdateSpec) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *RDSServerUpdateSpec) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *RDSServerUpdateSpec) SetEnabled(v bool) {
	o.Enabled = v
}

// GetMaxSessionsCountConfigured returns the MaxSessionsCountConfigured field value if set, zero value otherwise.
func (o *RDSServerUpdateSpec) GetMaxSessionsCountConfigured() int32 {
	if o == nil || o.MaxSessionsCountConfigured == nil {
		var ret int32
		return ret
	}
	return *o.MaxSessionsCountConfigured
}

// GetMaxSessionsCountConfiguredOk returns a tuple with the MaxSessionsCountConfigured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerUpdateSpec) GetMaxSessionsCountConfiguredOk() (*int32, bool) {
	if o == nil || o.MaxSessionsCountConfigured == nil {
		return nil, false
	}
	return o.MaxSessionsCountConfigured, true
}

// HasMaxSessionsCountConfigured returns a boolean if a field has been set.
func (o *RDSServerUpdateSpec) HasMaxSessionsCountConfigured() bool {
	if o != nil && o.MaxSessionsCountConfigured != nil {
		return true
	}

	return false
}

// SetMaxSessionsCountConfigured gets a reference to the given int32 and assigns it to the MaxSessionsCountConfigured field.
func (o *RDSServerUpdateSpec) SetMaxSessionsCountConfigured(v int32) {
	o.MaxSessionsCountConfigured = &v
}

// GetMaxSessionsTypeConfigured returns the MaxSessionsTypeConfigured field value
func (o *RDSServerUpdateSpec) GetMaxSessionsTypeConfigured() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxSessionsTypeConfigured
}

// GetMaxSessionsTypeConfiguredOk returns a tuple with the MaxSessionsTypeConfigured field value
// and a boolean to check if the value has been set.
func (o *RDSServerUpdateSpec) GetMaxSessionsTypeConfiguredOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSessionsTypeConfigured, true
}

// SetMaxSessionsTypeConfigured sets field value
func (o *RDSServerUpdateSpec) SetMaxSessionsTypeConfigured(v string) {
	o.MaxSessionsTypeConfigured = v
}

func (o RDSServerUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.MaxSessionsCountConfigured != nil {
		toSerialize["max_sessions_count_configured"] = o.MaxSessionsCountConfigured
	}
	if true {
		toSerialize["max_sessions_type_configured"] = o.MaxSessionsTypeConfigured
	}
	return json.Marshal(toSerialize)
}

type NullableRDSServerUpdateSpec struct {
	value *RDSServerUpdateSpec
	isSet bool
}

func (v NullableRDSServerUpdateSpec) Get() *RDSServerUpdateSpec {
	return v.value
}

func (v *NullableRDSServerUpdateSpec) Set(val *RDSServerUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableRDSServerUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableRDSServerUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRDSServerUpdateSpec(val *RDSServerUpdateSpec) *NullableRDSServerUpdateSpec {
	return &NullableRDSServerUpdateSpec{value: val, isSet: true}
}

func (v NullableRDSServerUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRDSServerUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
