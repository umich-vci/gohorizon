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

// DesktopPoolViewStorageAcceleratorSettingsCreateSpec Applicable To: Managed Desktop Pool. <br>View Storage Accelerator settings for Managed desktop pool.
type DesktopPoolViewStorageAcceleratorSettingsCreateSpec struct {
	// A list of blackout times.
	BlackoutTimes *[]ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec `json:"blackout_times,omitempty"`
	// How often to regenerate the View Storage Accelerator cache. Measured in Days. <br> This property is required if useViewStorageAccelerator is set to true. <br> Default value is 7.
	RegenerateViewStorageAcceleratorDays *int32 `json:"regenerate_view_storage_accelerator_days,omitempty"`
	// Indicates whether to use View Storage Accelerator. <br> Default value is false.
	UseViewStorageAccelerator *bool `json:"use_view_storage_accelerator,omitempty"`
}

// NewDesktopPoolViewStorageAcceleratorSettingsCreateSpec instantiates a new DesktopPoolViewStorageAcceleratorSettingsCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolViewStorageAcceleratorSettingsCreateSpec() *DesktopPoolViewStorageAcceleratorSettingsCreateSpec {
	this := DesktopPoolViewStorageAcceleratorSettingsCreateSpec{}
	return &this
}

// NewDesktopPoolViewStorageAcceleratorSettingsCreateSpecWithDefaults instantiates a new DesktopPoolViewStorageAcceleratorSettingsCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolViewStorageAcceleratorSettingsCreateSpecWithDefaults() *DesktopPoolViewStorageAcceleratorSettingsCreateSpec {
	this := DesktopPoolViewStorageAcceleratorSettingsCreateSpec{}
	return &this
}

// GetBlackoutTimes returns the BlackoutTimes field value if set, zero value otherwise.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetBlackoutTimes() []ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec {
	if o == nil || o.BlackoutTimes == nil {
		var ret []ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec
		return ret
	}
	return *o.BlackoutTimes
}

// GetBlackoutTimesOk returns a tuple with the BlackoutTimes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetBlackoutTimesOk() (*[]ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec, bool) {
	if o == nil || o.BlackoutTimes == nil {
		return nil, false
	}
	return o.BlackoutTimes, true
}

// HasBlackoutTimes returns a boolean if a field has been set.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) HasBlackoutTimes() bool {
	if o != nil && o.BlackoutTimes != nil {
		return true
	}

	return false
}

// SetBlackoutTimes gets a reference to the given []ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec and assigns it to the BlackoutTimes field.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) SetBlackoutTimes(v []ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) {
	o.BlackoutTimes = &v
}

// GetRegenerateViewStorageAcceleratorDays returns the RegenerateViewStorageAcceleratorDays field value if set, zero value otherwise.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetRegenerateViewStorageAcceleratorDays() int32 {
	if o == nil || o.RegenerateViewStorageAcceleratorDays == nil {
		var ret int32
		return ret
	}
	return *o.RegenerateViewStorageAcceleratorDays
}

// GetRegenerateViewStorageAcceleratorDaysOk returns a tuple with the RegenerateViewStorageAcceleratorDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetRegenerateViewStorageAcceleratorDaysOk() (*int32, bool) {
	if o == nil || o.RegenerateViewStorageAcceleratorDays == nil {
		return nil, false
	}
	return o.RegenerateViewStorageAcceleratorDays, true
}

// HasRegenerateViewStorageAcceleratorDays returns a boolean if a field has been set.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) HasRegenerateViewStorageAcceleratorDays() bool {
	if o != nil && o.RegenerateViewStorageAcceleratorDays != nil {
		return true
	}

	return false
}

// SetRegenerateViewStorageAcceleratorDays gets a reference to the given int32 and assigns it to the RegenerateViewStorageAcceleratorDays field.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) SetRegenerateViewStorageAcceleratorDays(v int32) {
	o.RegenerateViewStorageAcceleratorDays = &v
}

// GetUseViewStorageAccelerator returns the UseViewStorageAccelerator field value if set, zero value otherwise.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetUseViewStorageAccelerator() bool {
	if o == nil || o.UseViewStorageAccelerator == nil {
		var ret bool
		return ret
	}
	return *o.UseViewStorageAccelerator
}

// GetUseViewStorageAcceleratorOk returns a tuple with the UseViewStorageAccelerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) GetUseViewStorageAcceleratorOk() (*bool, bool) {
	if o == nil || o.UseViewStorageAccelerator == nil {
		return nil, false
	}
	return o.UseViewStorageAccelerator, true
}

// HasUseViewStorageAccelerator returns a boolean if a field has been set.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) HasUseViewStorageAccelerator() bool {
	if o != nil && o.UseViewStorageAccelerator != nil {
		return true
	}

	return false
}

// SetUseViewStorageAccelerator gets a reference to the given bool and assigns it to the UseViewStorageAccelerator field.
func (o *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) SetUseViewStorageAccelerator(v bool) {
	o.UseViewStorageAccelerator = &v
}

func (o DesktopPoolViewStorageAcceleratorSettingsCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BlackoutTimes != nil {
		toSerialize["blackout_times"] = o.BlackoutTimes
	}
	if o.RegenerateViewStorageAcceleratorDays != nil {
		toSerialize["regenerate_view_storage_accelerator_days"] = o.RegenerateViewStorageAcceleratorDays
	}
	if o.UseViewStorageAccelerator != nil {
		toSerialize["use_view_storage_accelerator"] = o.UseViewStorageAccelerator
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec struct {
	value *DesktopPoolViewStorageAcceleratorSettingsCreateSpec
	isSet bool
}

func (v NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec) Get() *DesktopPoolViewStorageAcceleratorSettingsCreateSpec {
	return v.value
}

func (v *NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec) Set(val *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec(val *DesktopPoolViewStorageAcceleratorSettingsCreateSpec) *NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec {
	return &NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec{value: val, isSet: true}
}

func (v NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolViewStorageAcceleratorSettingsCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
