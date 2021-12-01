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

// FarmLoadBalancerSettings Settings for load balancing the session requests across the RDS servers in the farm.
type FarmLoadBalancerSettings struct {
	// Indicates whether to use custom scripts for Load Balancing.Default is false.
	CustomScriptInUse *bool                       `json:"custom_script_in_use,omitempty"`
	LbMetricSettings  *LoadBalancerMetricSettings `json:"lb_metric_settings,omitempty"`
}

// NewFarmLoadBalancerSettings instantiates a new FarmLoadBalancerSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmLoadBalancerSettings() *FarmLoadBalancerSettings {
	this := FarmLoadBalancerSettings{}
	return &this
}

// NewFarmLoadBalancerSettingsWithDefaults instantiates a new FarmLoadBalancerSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmLoadBalancerSettingsWithDefaults() *FarmLoadBalancerSettings {
	this := FarmLoadBalancerSettings{}
	return &this
}

// GetCustomScriptInUse returns the CustomScriptInUse field value if set, zero value otherwise.
func (o *FarmLoadBalancerSettings) GetCustomScriptInUse() bool {
	if o == nil || o.CustomScriptInUse == nil {
		var ret bool
		return ret
	}
	return *o.CustomScriptInUse
}

// GetCustomScriptInUseOk returns a tuple with the CustomScriptInUse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmLoadBalancerSettings) GetCustomScriptInUseOk() (*bool, bool) {
	if o == nil || o.CustomScriptInUse == nil {
		return nil, false
	}
	return o.CustomScriptInUse, true
}

// HasCustomScriptInUse returns a boolean if a field has been set.
func (o *FarmLoadBalancerSettings) HasCustomScriptInUse() bool {
	if o != nil && o.CustomScriptInUse != nil {
		return true
	}

	return false
}

// SetCustomScriptInUse gets a reference to the given bool and assigns it to the CustomScriptInUse field.
func (o *FarmLoadBalancerSettings) SetCustomScriptInUse(v bool) {
	o.CustomScriptInUse = &v
}

// GetLbMetricSettings returns the LbMetricSettings field value if set, zero value otherwise.
func (o *FarmLoadBalancerSettings) GetLbMetricSettings() LoadBalancerMetricSettings {
	if o == nil || o.LbMetricSettings == nil {
		var ret LoadBalancerMetricSettings
		return ret
	}
	return *o.LbMetricSettings
}

// GetLbMetricSettingsOk returns a tuple with the LbMetricSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmLoadBalancerSettings) GetLbMetricSettingsOk() (*LoadBalancerMetricSettings, bool) {
	if o == nil || o.LbMetricSettings == nil {
		return nil, false
	}
	return o.LbMetricSettings, true
}

// HasLbMetricSettings returns a boolean if a field has been set.
func (o *FarmLoadBalancerSettings) HasLbMetricSettings() bool {
	if o != nil && o.LbMetricSettings != nil {
		return true
	}

	return false
}

// SetLbMetricSettings gets a reference to the given LoadBalancerMetricSettings and assigns it to the LbMetricSettings field.
func (o *FarmLoadBalancerSettings) SetLbMetricSettings(v LoadBalancerMetricSettings) {
	o.LbMetricSettings = &v
}

func (o FarmLoadBalancerSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomScriptInUse != nil {
		toSerialize["custom_script_in_use"] = o.CustomScriptInUse
	}
	if o.LbMetricSettings != nil {
		toSerialize["lb_metric_settings"] = o.LbMetricSettings
	}
	return json.Marshal(toSerialize)
}

type NullableFarmLoadBalancerSettings struct {
	value *FarmLoadBalancerSettings
	isSet bool
}

func (v NullableFarmLoadBalancerSettings) Get() *FarmLoadBalancerSettings {
	return v.value
}

func (v *NullableFarmLoadBalancerSettings) Set(val *FarmLoadBalancerSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmLoadBalancerSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmLoadBalancerSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmLoadBalancerSettings(val *FarmLoadBalancerSettings) *NullableFarmLoadBalancerSettings {
	return &NullableFarmLoadBalancerSettings{value: val, isSet: true}
}

func (v NullableFarmLoadBalancerSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmLoadBalancerSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
