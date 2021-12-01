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

// ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec Fields for specifying blackout time for View Storage Accelerator. Storage accelerator regeneration and VM disk space reclamation do not occur during blackout times. The same blackout policy applies to both operations.
type ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec struct {
	// List of days for a given range of time.
	Days []string `json:"days"`
	// Ending time for the blackout in 24-hour format.
	EndTime string `json:"end_time"`
	// Starting time for the blackout in 24-hour format.
	StartTime string `json:"start_time"`
}

// NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpec instantiates a new ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpec(days []string, endTime string, startTime string) *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec {
	this := ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec{}
	this.Days = days
	this.EndTime = endTime
	this.StartTime = startTime
	return &this
}

// NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpecWithDefaults instantiates a new ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewStorageAcceleratorBlackoutTimeSettingsCreateSpecWithDefaults() *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec {
	this := ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec{}
	return &this
}

// GetDays returns the Days field value
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetDays() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Days
}

// GetDaysOk returns a tuple with the Days field value
// and a boolean to check if the value has been set.
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetDaysOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Days, true
}

// SetDays sets field value
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) SetDays(v []string) {
	o.Days = v
}

// GetEndTime returns the EndTime field value
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetEndTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetEndTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) SetEndTime(v string) {
	o.EndTime = v
}

// GetStartTime returns the StartTime field value
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) SetStartTime(v string) {
	o.StartTime = v
}

func (o ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["days"] = o.Days
	}
	if true {
		toSerialize["end_time"] = o.EndTime
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec struct {
	value *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec
	isSet bool
}

func (v NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) Get() *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec {
	return v.value
}

func (v *NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) Set(val *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec(val *ViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) *NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec {
	return &NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec{value: val, isSet: true}
}

func (v NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewStorageAcceleratorBlackoutTimeSettingsCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
