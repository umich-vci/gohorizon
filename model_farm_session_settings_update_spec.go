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

// FarmSessionSettingsUpdateSpec Session settings related to the Farm.
type FarmSessionSettingsUpdateSpec struct {
	// Disconnected sessions timeout (in minutes). An empty disconnected session is logged off after the timeout. This is required if the disconnect_session_timeout_policy is set to AFTER.
	DisconnectedSessionTimeoutMinutes *int32 `json:"disconnected_session_timeout_minutes,omitempty"`
	// Log-off policy after disconnected session. * IMMEDIATELY: Immediately Logoff after user disconnect. * AFTER: Logoff after the specified number of minutes after user disconnect. * NEVER: Do not logoff after user disconnect.
	DisconnectedSessionTimeoutPolicy string `json:"disconnected_session_timeout_policy"`
	// Application empty session timeout (in minutes). An empty session (that has no remote-able window) is disconnected after the timeout. This is required if the empty_session_timeout_policy is set to AFTER.
	EmptySessionTimeoutMinutes *int32 `json:"empty_session_timeout_minutes,omitempty"`
	// Application empty session timeout policy. * IMMEDIATE: Empty session will be disconnected immediately. * NEVER: Empty session will never disconnected. * AFTER: Empty session will be disconnected after specified number of minutes.
	EmptySessionTimeoutPolicy string `json:"empty_session_timeout_policy"`
	// Indicates whether the empty application sessions are logged off (true) or disconnected (false) after timeout. Default value is false if emptySessionTimeoutPolicy is set to AFTER or IMMEDIATE
	LogoffAfterTimeout *bool `json:"logoff_after_timeout,omitempty"`
	// Application pre-launch session timeout (in minutes). A pre-launch session is disconnected after the timeout.This is required if pre_launch_session_timeout_policy is set to AFTER.
	PreLaunchSessionTimeoutMinutes *int32 `json:"pre_launch_session_timeout_minutes,omitempty"`
	// Application pre-launch session timeout policy. * AFTER: Pre-launched session is disconnected after specified number of minutes. * NEVER: Pre-launched session is never disconnected.
	PreLaunchSessionTimeoutPolicy *string `json:"pre_launch_session_timeout_policy,omitempty"`
}

// NewFarmSessionSettingsUpdateSpec instantiates a new FarmSessionSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmSessionSettingsUpdateSpec(disconnectedSessionTimeoutPolicy string, emptySessionTimeoutPolicy string) *FarmSessionSettingsUpdateSpec {
	this := FarmSessionSettingsUpdateSpec{}
	this.DisconnectedSessionTimeoutPolicy = disconnectedSessionTimeoutPolicy
	this.EmptySessionTimeoutPolicy = emptySessionTimeoutPolicy
	return &this
}

// NewFarmSessionSettingsUpdateSpecWithDefaults instantiates a new FarmSessionSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmSessionSettingsUpdateSpecWithDefaults() *FarmSessionSettingsUpdateSpec {
	this := FarmSessionSettingsUpdateSpec{}
	return &this
}

// GetDisconnectedSessionTimeoutMinutes returns the DisconnectedSessionTimeoutMinutes field value if set, zero value otherwise.
func (o *FarmSessionSettingsUpdateSpec) GetDisconnectedSessionTimeoutMinutes() int32 {
	if o == nil || o.DisconnectedSessionTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.DisconnectedSessionTimeoutMinutes
}

// GetDisconnectedSessionTimeoutMinutesOk returns a tuple with the DisconnectedSessionTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettingsUpdateSpec) GetDisconnectedSessionTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.DisconnectedSessionTimeoutMinutes == nil {
		return nil, false
	}
	return o.DisconnectedSessionTimeoutMinutes, true
}

// HasDisconnectedSessionTimeoutMinutes returns a boolean if a field has been set.
func (o *FarmSessionSettingsUpdateSpec) HasDisconnectedSessionTimeoutMinutes() bool {
	if o != nil && o.DisconnectedSessionTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetDisconnectedSessionTimeoutMinutes gets a reference to the given int32 and assigns it to the DisconnectedSessionTimeoutMinutes field.
func (o *FarmSessionSettingsUpdateSpec) SetDisconnectedSessionTimeoutMinutes(v int32) {
	o.DisconnectedSessionTimeoutMinutes = &v
}

// GetDisconnectedSessionTimeoutPolicy returns the DisconnectedSessionTimeoutPolicy field value
func (o *FarmSessionSettingsUpdateSpec) GetDisconnectedSessionTimeoutPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisconnectedSessionTimeoutPolicy
}

// GetDisconnectedSessionTimeoutPolicyOk returns a tuple with the DisconnectedSessionTimeoutPolicy field value
// and a boolean to check if the value has been set.
func (o *FarmSessionSettingsUpdateSpec) GetDisconnectedSessionTimeoutPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisconnectedSessionTimeoutPolicy, true
}

// SetDisconnectedSessionTimeoutPolicy sets field value
func (o *FarmSessionSettingsUpdateSpec) SetDisconnectedSessionTimeoutPolicy(v string) {
	o.DisconnectedSessionTimeoutPolicy = v
}

// GetEmptySessionTimeoutMinutes returns the EmptySessionTimeoutMinutes field value if set, zero value otherwise.
func (o *FarmSessionSettingsUpdateSpec) GetEmptySessionTimeoutMinutes() int32 {
	if o == nil || o.EmptySessionTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.EmptySessionTimeoutMinutes
}

// GetEmptySessionTimeoutMinutesOk returns a tuple with the EmptySessionTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettingsUpdateSpec) GetEmptySessionTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.EmptySessionTimeoutMinutes == nil {
		return nil, false
	}
	return o.EmptySessionTimeoutMinutes, true
}

// HasEmptySessionTimeoutMinutes returns a boolean if a field has been set.
func (o *FarmSessionSettingsUpdateSpec) HasEmptySessionTimeoutMinutes() bool {
	if o != nil && o.EmptySessionTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetEmptySessionTimeoutMinutes gets a reference to the given int32 and assigns it to the EmptySessionTimeoutMinutes field.
func (o *FarmSessionSettingsUpdateSpec) SetEmptySessionTimeoutMinutes(v int32) {
	o.EmptySessionTimeoutMinutes = &v
}

// GetEmptySessionTimeoutPolicy returns the EmptySessionTimeoutPolicy field value
func (o *FarmSessionSettingsUpdateSpec) GetEmptySessionTimeoutPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmptySessionTimeoutPolicy
}

// GetEmptySessionTimeoutPolicyOk returns a tuple with the EmptySessionTimeoutPolicy field value
// and a boolean to check if the value has been set.
func (o *FarmSessionSettingsUpdateSpec) GetEmptySessionTimeoutPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmptySessionTimeoutPolicy, true
}

// SetEmptySessionTimeoutPolicy sets field value
func (o *FarmSessionSettingsUpdateSpec) SetEmptySessionTimeoutPolicy(v string) {
	o.EmptySessionTimeoutPolicy = v
}

// GetLogoffAfterTimeout returns the LogoffAfterTimeout field value if set, zero value otherwise.
func (o *FarmSessionSettingsUpdateSpec) GetLogoffAfterTimeout() bool {
	if o == nil || o.LogoffAfterTimeout == nil {
		var ret bool
		return ret
	}
	return *o.LogoffAfterTimeout
}

// GetLogoffAfterTimeoutOk returns a tuple with the LogoffAfterTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettingsUpdateSpec) GetLogoffAfterTimeoutOk() (*bool, bool) {
	if o == nil || o.LogoffAfterTimeout == nil {
		return nil, false
	}
	return o.LogoffAfterTimeout, true
}

// HasLogoffAfterTimeout returns a boolean if a field has been set.
func (o *FarmSessionSettingsUpdateSpec) HasLogoffAfterTimeout() bool {
	if o != nil && o.LogoffAfterTimeout != nil {
		return true
	}

	return false
}

// SetLogoffAfterTimeout gets a reference to the given bool and assigns it to the LogoffAfterTimeout field.
func (o *FarmSessionSettingsUpdateSpec) SetLogoffAfterTimeout(v bool) {
	o.LogoffAfterTimeout = &v
}

// GetPreLaunchSessionTimeoutMinutes returns the PreLaunchSessionTimeoutMinutes field value if set, zero value otherwise.
func (o *FarmSessionSettingsUpdateSpec) GetPreLaunchSessionTimeoutMinutes() int32 {
	if o == nil || o.PreLaunchSessionTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.PreLaunchSessionTimeoutMinutes
}

// GetPreLaunchSessionTimeoutMinutesOk returns a tuple with the PreLaunchSessionTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettingsUpdateSpec) GetPreLaunchSessionTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.PreLaunchSessionTimeoutMinutes == nil {
		return nil, false
	}
	return o.PreLaunchSessionTimeoutMinutes, true
}

// HasPreLaunchSessionTimeoutMinutes returns a boolean if a field has been set.
func (o *FarmSessionSettingsUpdateSpec) HasPreLaunchSessionTimeoutMinutes() bool {
	if o != nil && o.PreLaunchSessionTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetPreLaunchSessionTimeoutMinutes gets a reference to the given int32 and assigns it to the PreLaunchSessionTimeoutMinutes field.
func (o *FarmSessionSettingsUpdateSpec) SetPreLaunchSessionTimeoutMinutes(v int32) {
	o.PreLaunchSessionTimeoutMinutes = &v
}

// GetPreLaunchSessionTimeoutPolicy returns the PreLaunchSessionTimeoutPolicy field value if set, zero value otherwise.
func (o *FarmSessionSettingsUpdateSpec) GetPreLaunchSessionTimeoutPolicy() string {
	if o == nil || o.PreLaunchSessionTimeoutPolicy == nil {
		var ret string
		return ret
	}
	return *o.PreLaunchSessionTimeoutPolicy
}

// GetPreLaunchSessionTimeoutPolicyOk returns a tuple with the PreLaunchSessionTimeoutPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettingsUpdateSpec) GetPreLaunchSessionTimeoutPolicyOk() (*string, bool) {
	if o == nil || o.PreLaunchSessionTimeoutPolicy == nil {
		return nil, false
	}
	return o.PreLaunchSessionTimeoutPolicy, true
}

// HasPreLaunchSessionTimeoutPolicy returns a boolean if a field has been set.
func (o *FarmSessionSettingsUpdateSpec) HasPreLaunchSessionTimeoutPolicy() bool {
	if o != nil && o.PreLaunchSessionTimeoutPolicy != nil {
		return true
	}

	return false
}

// SetPreLaunchSessionTimeoutPolicy gets a reference to the given string and assigns it to the PreLaunchSessionTimeoutPolicy field.
func (o *FarmSessionSettingsUpdateSpec) SetPreLaunchSessionTimeoutPolicy(v string) {
	o.PreLaunchSessionTimeoutPolicy = &v
}

func (o FarmSessionSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisconnectedSessionTimeoutMinutes != nil {
		toSerialize["disconnected_session_timeout_minutes"] = o.DisconnectedSessionTimeoutMinutes
	}
	if true {
		toSerialize["disconnected_session_timeout_policy"] = o.DisconnectedSessionTimeoutPolicy
	}
	if o.EmptySessionTimeoutMinutes != nil {
		toSerialize["empty_session_timeout_minutes"] = o.EmptySessionTimeoutMinutes
	}
	if true {
		toSerialize["empty_session_timeout_policy"] = o.EmptySessionTimeoutPolicy
	}
	if o.LogoffAfterTimeout != nil {
		toSerialize["logoff_after_timeout"] = o.LogoffAfterTimeout
	}
	if o.PreLaunchSessionTimeoutMinutes != nil {
		toSerialize["pre_launch_session_timeout_minutes"] = o.PreLaunchSessionTimeoutMinutes
	}
	if o.PreLaunchSessionTimeoutPolicy != nil {
		toSerialize["pre_launch_session_timeout_policy"] = o.PreLaunchSessionTimeoutPolicy
	}
	return json.Marshal(toSerialize)
}

type NullableFarmSessionSettingsUpdateSpec struct {
	value *FarmSessionSettingsUpdateSpec
	isSet bool
}

func (v NullableFarmSessionSettingsUpdateSpec) Get() *FarmSessionSettingsUpdateSpec {
	return v.value
}

func (v *NullableFarmSessionSettingsUpdateSpec) Set(val *FarmSessionSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmSessionSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmSessionSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmSessionSettingsUpdateSpec(val *FarmSessionSettingsUpdateSpec) *NullableFarmSessionSettingsUpdateSpec {
	return &NullableFarmSessionSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableFarmSessionSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmSessionSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
