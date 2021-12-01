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

// DesktopPoolInstantClonePushImageSettings Settings for the push image operation.
type DesktopPoolInstantClonePushImageSettings struct {
	// Whether to add Virtual TPM device.
	AddVirtualTpm *bool `json:"add_virtual_tpm,omitempty"`
	// Determines when to perform the operation on machines which have an active session. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions.
	LogoffPolicy *string `json:"logoff_policy,omitempty"`
	// When to start the operation. If unset or the time is in the past, the operation will begin immediately. Measured as epoch time.
	StartTime *int64 `json:"start_time,omitempty"`
	// Indicates that the operation should stop on first error.
	StopOnFirstError *bool `json:"stop_on_first_error,omitempty"`
}

// NewDesktopPoolInstantClonePushImageSettings instantiates a new DesktopPoolInstantClonePushImageSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolInstantClonePushImageSettings() *DesktopPoolInstantClonePushImageSettings {
	this := DesktopPoolInstantClonePushImageSettings{}
	return &this
}

// NewDesktopPoolInstantClonePushImageSettingsWithDefaults instantiates a new DesktopPoolInstantClonePushImageSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolInstantClonePushImageSettingsWithDefaults() *DesktopPoolInstantClonePushImageSettings {
	this := DesktopPoolInstantClonePushImageSettings{}
	return &this
}

// GetAddVirtualTpm returns the AddVirtualTpm field value if set, zero value otherwise.
func (o *DesktopPoolInstantClonePushImageSettings) GetAddVirtualTpm() bool {
	if o == nil || o.AddVirtualTpm == nil {
		var ret bool
		return ret
	}
	return *o.AddVirtualTpm
}

// GetAddVirtualTpmOk returns a tuple with the AddVirtualTpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolInstantClonePushImageSettings) GetAddVirtualTpmOk() (*bool, bool) {
	if o == nil || o.AddVirtualTpm == nil {
		return nil, false
	}
	return o.AddVirtualTpm, true
}

// HasAddVirtualTpm returns a boolean if a field has been set.
func (o *DesktopPoolInstantClonePushImageSettings) HasAddVirtualTpm() bool {
	if o != nil && o.AddVirtualTpm != nil {
		return true
	}

	return false
}

// SetAddVirtualTpm gets a reference to the given bool and assigns it to the AddVirtualTpm field.
func (o *DesktopPoolInstantClonePushImageSettings) SetAddVirtualTpm(v bool) {
	o.AddVirtualTpm = &v
}

// GetLogoffPolicy returns the LogoffPolicy field value if set, zero value otherwise.
func (o *DesktopPoolInstantClonePushImageSettings) GetLogoffPolicy() string {
	if o == nil || o.LogoffPolicy == nil {
		var ret string
		return ret
	}
	return *o.LogoffPolicy
}

// GetLogoffPolicyOk returns a tuple with the LogoffPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolInstantClonePushImageSettings) GetLogoffPolicyOk() (*string, bool) {
	if o == nil || o.LogoffPolicy == nil {
		return nil, false
	}
	return o.LogoffPolicy, true
}

// HasLogoffPolicy returns a boolean if a field has been set.
func (o *DesktopPoolInstantClonePushImageSettings) HasLogoffPolicy() bool {
	if o != nil && o.LogoffPolicy != nil {
		return true
	}

	return false
}

// SetLogoffPolicy gets a reference to the given string and assigns it to the LogoffPolicy field.
func (o *DesktopPoolInstantClonePushImageSettings) SetLogoffPolicy(v string) {
	o.LogoffPolicy = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DesktopPoolInstantClonePushImageSettings) GetStartTime() int64 {
	if o == nil || o.StartTime == nil {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolInstantClonePushImageSettings) GetStartTimeOk() (*int64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DesktopPoolInstantClonePushImageSettings) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *DesktopPoolInstantClonePushImageSettings) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStopOnFirstError returns the StopOnFirstError field value if set, zero value otherwise.
func (o *DesktopPoolInstantClonePushImageSettings) GetStopOnFirstError() bool {
	if o == nil || o.StopOnFirstError == nil {
		var ret bool
		return ret
	}
	return *o.StopOnFirstError
}

// GetStopOnFirstErrorOk returns a tuple with the StopOnFirstError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolInstantClonePushImageSettings) GetStopOnFirstErrorOk() (*bool, bool) {
	if o == nil || o.StopOnFirstError == nil {
		return nil, false
	}
	return o.StopOnFirstError, true
}

// HasStopOnFirstError returns a boolean if a field has been set.
func (o *DesktopPoolInstantClonePushImageSettings) HasStopOnFirstError() bool {
	if o != nil && o.StopOnFirstError != nil {
		return true
	}

	return false
}

// SetStopOnFirstError gets a reference to the given bool and assigns it to the StopOnFirstError field.
func (o *DesktopPoolInstantClonePushImageSettings) SetStopOnFirstError(v bool) {
	o.StopOnFirstError = &v
}

func (o DesktopPoolInstantClonePushImageSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddVirtualTpm != nil {
		toSerialize["add_virtual_tpm"] = o.AddVirtualTpm
	}
	if o.LogoffPolicy != nil {
		toSerialize["logoff_policy"] = o.LogoffPolicy
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.StopOnFirstError != nil {
		toSerialize["stop_on_first_error"] = o.StopOnFirstError
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolInstantClonePushImageSettings struct {
	value *DesktopPoolInstantClonePushImageSettings
	isSet bool
}

func (v NullableDesktopPoolInstantClonePushImageSettings) Get() *DesktopPoolInstantClonePushImageSettings {
	return v.value
}

func (v *NullableDesktopPoolInstantClonePushImageSettings) Set(val *DesktopPoolInstantClonePushImageSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolInstantClonePushImageSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolInstantClonePushImageSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolInstantClonePushImageSettings(val *DesktopPoolInstantClonePushImageSettings) *NullableDesktopPoolInstantClonePushImageSettings {
	return &NullableDesktopPoolInstantClonePushImageSettings{value: val, isSet: true}
}

func (v NullableDesktopPoolInstantClonePushImageSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolInstantClonePushImageSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
