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

// FarmSessionSettings Session related settings for the Farm.
type FarmSessionSettings struct {
	// Disconnected sessions timeout (in minutes).Will be set when disconnected_session_timeout_policy is set to AFTER.
	DisconnectedSessionTimeoutMinutes *int32 `json:"disconnected_session_timeout_minutes,omitempty"`
	// Log-off policy after disconnected session. Default value is NEVER. * IMMEDIATELY: Immediately Logoff after user disconnect. * AFTER: Logoff after the specified number of minutes after user disconnect. * NEVER: Do not logoff after user disconnect.
	DisconnectedSessionTimeoutPolicy *string `json:"disconnected_session_timeout_policy,omitempty"`
	// Application empty session timeout in minutes. An empty session that has no remote-ablewindow is disconnected after the timeout. Default value is 1.Will be set when the empty_session_timeout_policy set to AFTER.
	EmptySessionTimeoutMinutes *int32 `json:"empty_session_timeout_minutes,omitempty"`
	// Application empty session timeout policy. Default value is AFTER. * IMMEDIATE: Empty session will be disconnected immediately. * NEVER: Empty session will never disconnected. * AFTER: Empty session will be disconnected after specified number of minutes.
	EmptySessionTimeoutPolicy *string `json:"empty_session_timeout_policy,omitempty"`
	// After timeout, empty application sessions are logged off when set to true. Otherwise sessions are disconnected.Default value is false.
	LogoffAfterTimeout *bool `json:"logoff_after_timeout,omitempty"`
	// Application pre-launch session timeout in minutes. A pre-launch session is disconnected after the timeout.Will be set only when pre_launch_timeout_policy is set to AFTER.
	PreLaunchSessionTimeoutMinutes *int32 `json:"pre_launch_session_timeout_minutes,omitempty"`
	// Pre-launch session timeout policy for the application sessions on this Farm. * AFTER: Pre-launched session is disconnected after specified number of minutes. * NEVER: Pre-launched session is never disconnected.
	PreLaunchSessionTimeoutPolicy *string `json:"pre_launch_session_timeout_policy,omitempty"`
}

// NewFarmSessionSettings instantiates a new FarmSessionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmSessionSettings() *FarmSessionSettings {
	this := FarmSessionSettings{}
	return &this
}

// NewFarmSessionSettingsWithDefaults instantiates a new FarmSessionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmSessionSettingsWithDefaults() *FarmSessionSettings {
	this := FarmSessionSettings{}
	return &this
}

// GetDisconnectedSessionTimeoutMinutes returns the DisconnectedSessionTimeoutMinutes field value if set, zero value otherwise.
func (o *FarmSessionSettings) GetDisconnectedSessionTimeoutMinutes() int32 {
	if o == nil || o.DisconnectedSessionTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.DisconnectedSessionTimeoutMinutes
}

// GetDisconnectedSessionTimeoutMinutesOk returns a tuple with the DisconnectedSessionTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettings) GetDisconnectedSessionTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.DisconnectedSessionTimeoutMinutes == nil {
		return nil, false
	}
	return o.DisconnectedSessionTimeoutMinutes, true
}

// HasDisconnectedSessionTimeoutMinutes returns a boolean if a field has been set.
func (o *FarmSessionSettings) HasDisconnectedSessionTimeoutMinutes() bool {
	if o != nil && o.DisconnectedSessionTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetDisconnectedSessionTimeoutMinutes gets a reference to the given int32 and assigns it to the DisconnectedSessionTimeoutMinutes field.
func (o *FarmSessionSettings) SetDisconnectedSessionTimeoutMinutes(v int32) {
	o.DisconnectedSessionTimeoutMinutes = &v
}

// GetDisconnectedSessionTimeoutPolicy returns the DisconnectedSessionTimeoutPolicy field value if set, zero value otherwise.
func (o *FarmSessionSettings) GetDisconnectedSessionTimeoutPolicy() string {
	if o == nil || o.DisconnectedSessionTimeoutPolicy == nil {
		var ret string
		return ret
	}
	return *o.DisconnectedSessionTimeoutPolicy
}

// GetDisconnectedSessionTimeoutPolicyOk returns a tuple with the DisconnectedSessionTimeoutPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettings) GetDisconnectedSessionTimeoutPolicyOk() (*string, bool) {
	if o == nil || o.DisconnectedSessionTimeoutPolicy == nil {
		return nil, false
	}
	return o.DisconnectedSessionTimeoutPolicy, true
}

// HasDisconnectedSessionTimeoutPolicy returns a boolean if a field has been set.
func (o *FarmSessionSettings) HasDisconnectedSessionTimeoutPolicy() bool {
	if o != nil && o.DisconnectedSessionTimeoutPolicy != nil {
		return true
	}

	return false
}

// SetDisconnectedSessionTimeoutPolicy gets a reference to the given string and assigns it to the DisconnectedSessionTimeoutPolicy field.
func (o *FarmSessionSettings) SetDisconnectedSessionTimeoutPolicy(v string) {
	o.DisconnectedSessionTimeoutPolicy = &v
}

// GetEmptySessionTimeoutMinutes returns the EmptySessionTimeoutMinutes field value if set, zero value otherwise.
func (o *FarmSessionSettings) GetEmptySessionTimeoutMinutes() int32 {
	if o == nil || o.EmptySessionTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.EmptySessionTimeoutMinutes
}

// GetEmptySessionTimeoutMinutesOk returns a tuple with the EmptySessionTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettings) GetEmptySessionTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.EmptySessionTimeoutMinutes == nil {
		return nil, false
	}
	return o.EmptySessionTimeoutMinutes, true
}

// HasEmptySessionTimeoutMinutes returns a boolean if a field has been set.
func (o *FarmSessionSettings) HasEmptySessionTimeoutMinutes() bool {
	if o != nil && o.EmptySessionTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetEmptySessionTimeoutMinutes gets a reference to the given int32 and assigns it to the EmptySessionTimeoutMinutes field.
func (o *FarmSessionSettings) SetEmptySessionTimeoutMinutes(v int32) {
	o.EmptySessionTimeoutMinutes = &v
}

// GetEmptySessionTimeoutPolicy returns the EmptySessionTimeoutPolicy field value if set, zero value otherwise.
func (o *FarmSessionSettings) GetEmptySessionTimeoutPolicy() string {
	if o == nil || o.EmptySessionTimeoutPolicy == nil {
		var ret string
		return ret
	}
	return *o.EmptySessionTimeoutPolicy
}

// GetEmptySessionTimeoutPolicyOk returns a tuple with the EmptySessionTimeoutPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettings) GetEmptySessionTimeoutPolicyOk() (*string, bool) {
	if o == nil || o.EmptySessionTimeoutPolicy == nil {
		return nil, false
	}
	return o.EmptySessionTimeoutPolicy, true
}

// HasEmptySessionTimeoutPolicy returns a boolean if a field has been set.
func (o *FarmSessionSettings) HasEmptySessionTimeoutPolicy() bool {
	if o != nil && o.EmptySessionTimeoutPolicy != nil {
		return true
	}

	return false
}

// SetEmptySessionTimeoutPolicy gets a reference to the given string and assigns it to the EmptySessionTimeoutPolicy field.
func (o *FarmSessionSettings) SetEmptySessionTimeoutPolicy(v string) {
	o.EmptySessionTimeoutPolicy = &v
}

// GetLogoffAfterTimeout returns the LogoffAfterTimeout field value if set, zero value otherwise.
func (o *FarmSessionSettings) GetLogoffAfterTimeout() bool {
	if o == nil || o.LogoffAfterTimeout == nil {
		var ret bool
		return ret
	}
	return *o.LogoffAfterTimeout
}

// GetLogoffAfterTimeoutOk returns a tuple with the LogoffAfterTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettings) GetLogoffAfterTimeoutOk() (*bool, bool) {
	if o == nil || o.LogoffAfterTimeout == nil {
		return nil, false
	}
	return o.LogoffAfterTimeout, true
}

// HasLogoffAfterTimeout returns a boolean if a field has been set.
func (o *FarmSessionSettings) HasLogoffAfterTimeout() bool {
	if o != nil && o.LogoffAfterTimeout != nil {
		return true
	}

	return false
}

// SetLogoffAfterTimeout gets a reference to the given bool and assigns it to the LogoffAfterTimeout field.
func (o *FarmSessionSettings) SetLogoffAfterTimeout(v bool) {
	o.LogoffAfterTimeout = &v
}

// GetPreLaunchSessionTimeoutMinutes returns the PreLaunchSessionTimeoutMinutes field value if set, zero value otherwise.
func (o *FarmSessionSettings) GetPreLaunchSessionTimeoutMinutes() int32 {
	if o == nil || o.PreLaunchSessionTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.PreLaunchSessionTimeoutMinutes
}

// GetPreLaunchSessionTimeoutMinutesOk returns a tuple with the PreLaunchSessionTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettings) GetPreLaunchSessionTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.PreLaunchSessionTimeoutMinutes == nil {
		return nil, false
	}
	return o.PreLaunchSessionTimeoutMinutes, true
}

// HasPreLaunchSessionTimeoutMinutes returns a boolean if a field has been set.
func (o *FarmSessionSettings) HasPreLaunchSessionTimeoutMinutes() bool {
	if o != nil && o.PreLaunchSessionTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetPreLaunchSessionTimeoutMinutes gets a reference to the given int32 and assigns it to the PreLaunchSessionTimeoutMinutes field.
func (o *FarmSessionSettings) SetPreLaunchSessionTimeoutMinutes(v int32) {
	o.PreLaunchSessionTimeoutMinutes = &v
}

// GetPreLaunchSessionTimeoutPolicy returns the PreLaunchSessionTimeoutPolicy field value if set, zero value otherwise.
func (o *FarmSessionSettings) GetPreLaunchSessionTimeoutPolicy() string {
	if o == nil || o.PreLaunchSessionTimeoutPolicy == nil {
		var ret string
		return ret
	}
	return *o.PreLaunchSessionTimeoutPolicy
}

// GetPreLaunchSessionTimeoutPolicyOk returns a tuple with the PreLaunchSessionTimeoutPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmSessionSettings) GetPreLaunchSessionTimeoutPolicyOk() (*string, bool) {
	if o == nil || o.PreLaunchSessionTimeoutPolicy == nil {
		return nil, false
	}
	return o.PreLaunchSessionTimeoutPolicy, true
}

// HasPreLaunchSessionTimeoutPolicy returns a boolean if a field has been set.
func (o *FarmSessionSettings) HasPreLaunchSessionTimeoutPolicy() bool {
	if o != nil && o.PreLaunchSessionTimeoutPolicy != nil {
		return true
	}

	return false
}

// SetPreLaunchSessionTimeoutPolicy gets a reference to the given string and assigns it to the PreLaunchSessionTimeoutPolicy field.
func (o *FarmSessionSettings) SetPreLaunchSessionTimeoutPolicy(v string) {
	o.PreLaunchSessionTimeoutPolicy = &v
}

func (o FarmSessionSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisconnectedSessionTimeoutMinutes != nil {
		toSerialize["disconnected_session_timeout_minutes"] = o.DisconnectedSessionTimeoutMinutes
	}
	if o.DisconnectedSessionTimeoutPolicy != nil {
		toSerialize["disconnected_session_timeout_policy"] = o.DisconnectedSessionTimeoutPolicy
	}
	if o.EmptySessionTimeoutMinutes != nil {
		toSerialize["empty_session_timeout_minutes"] = o.EmptySessionTimeoutMinutes
	}
	if o.EmptySessionTimeoutPolicy != nil {
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

type NullableFarmSessionSettings struct {
	value *FarmSessionSettings
	isSet bool
}

func (v NullableFarmSessionSettings) Get() *FarmSessionSettings {
	return v.value
}

func (v *NullableFarmSessionSettings) Set(val *FarmSessionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmSessionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmSessionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmSessionSettings(val *FarmSessionSettings) *NullableFarmSessionSettings {
	return &NullableFarmSessionSettings{value: val, isSet: true}
}

func (v NullableFarmSessionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmSessionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
