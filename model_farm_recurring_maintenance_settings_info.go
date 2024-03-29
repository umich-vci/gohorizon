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

// FarmRecurringMaintenanceSettingsInfo Settings for recurring maintenance operation for the farm.
type FarmRecurringMaintenanceSettingsInfo struct {
	// Settings for recurring maintenance operations. * DAILY: Daily recurring maintenance * WEEKLY: Weekly recurring maintenance * MONTHLY: Monthly recurring maintenance
	MaintenancePeriod *string `json:"maintenance_period,omitempty"`
	// Indicates frequency of repeating maintenance and is expressed as a multiple of the maintenance_period.
	MaintenancePeriodFrequency *int32 `json:"maintenance_period_frequency,omitempty"`
	// Start index for weekly or monthly maintenance. Weekly: 1-7 (Sun-Sat), Monthly: 1-31. This is set when maintenance_period is WEEKLY or MONTHLY.
	StartIndex *int32 `json:"start_index,omitempty"`
	// Start time configured for the recurring maintenance. This is in the form hh:mm in 24 hours format
	StartTime *string `json:"start_time,omitempty"`
}

// NewFarmRecurringMaintenanceSettingsInfo instantiates a new FarmRecurringMaintenanceSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmRecurringMaintenanceSettingsInfo() *FarmRecurringMaintenanceSettingsInfo {
	this := FarmRecurringMaintenanceSettingsInfo{}
	return &this
}

// NewFarmRecurringMaintenanceSettingsInfoWithDefaults instantiates a new FarmRecurringMaintenanceSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmRecurringMaintenanceSettingsInfoWithDefaults() *FarmRecurringMaintenanceSettingsInfo {
	this := FarmRecurringMaintenanceSettingsInfo{}
	return &this
}

// GetMaintenancePeriod returns the MaintenancePeriod field value if set, zero value otherwise.
func (o *FarmRecurringMaintenanceSettingsInfo) GetMaintenancePeriod() string {
	if o == nil || o.MaintenancePeriod == nil {
		var ret string
		return ret
	}
	return *o.MaintenancePeriod
}

// GetMaintenancePeriodOk returns a tuple with the MaintenancePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmRecurringMaintenanceSettingsInfo) GetMaintenancePeriodOk() (*string, bool) {
	if o == nil || o.MaintenancePeriod == nil {
		return nil, false
	}
	return o.MaintenancePeriod, true
}

// HasMaintenancePeriod returns a boolean if a field has been set.
func (o *FarmRecurringMaintenanceSettingsInfo) HasMaintenancePeriod() bool {
	if o != nil && o.MaintenancePeriod != nil {
		return true
	}

	return false
}

// SetMaintenancePeriod gets a reference to the given string and assigns it to the MaintenancePeriod field.
func (o *FarmRecurringMaintenanceSettingsInfo) SetMaintenancePeriod(v string) {
	o.MaintenancePeriod = &v
}

// GetMaintenancePeriodFrequency returns the MaintenancePeriodFrequency field value if set, zero value otherwise.
func (o *FarmRecurringMaintenanceSettingsInfo) GetMaintenancePeriodFrequency() int32 {
	if o == nil || o.MaintenancePeriodFrequency == nil {
		var ret int32
		return ret
	}
	return *o.MaintenancePeriodFrequency
}

// GetMaintenancePeriodFrequencyOk returns a tuple with the MaintenancePeriodFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmRecurringMaintenanceSettingsInfo) GetMaintenancePeriodFrequencyOk() (*int32, bool) {
	if o == nil || o.MaintenancePeriodFrequency == nil {
		return nil, false
	}
	return o.MaintenancePeriodFrequency, true
}

// HasMaintenancePeriodFrequency returns a boolean if a field has been set.
func (o *FarmRecurringMaintenanceSettingsInfo) HasMaintenancePeriodFrequency() bool {
	if o != nil && o.MaintenancePeriodFrequency != nil {
		return true
	}

	return false
}

// SetMaintenancePeriodFrequency gets a reference to the given int32 and assigns it to the MaintenancePeriodFrequency field.
func (o *FarmRecurringMaintenanceSettingsInfo) SetMaintenancePeriodFrequency(v int32) {
	o.MaintenancePeriodFrequency = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *FarmRecurringMaintenanceSettingsInfo) GetStartIndex() int32 {
	if o == nil || o.StartIndex == nil {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmRecurringMaintenanceSettingsInfo) GetStartIndexOk() (*int32, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *FarmRecurringMaintenanceSettingsInfo) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *FarmRecurringMaintenanceSettingsInfo) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *FarmRecurringMaintenanceSettingsInfo) GetStartTime() string {
	if o == nil || o.StartTime == nil {
		var ret string
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmRecurringMaintenanceSettingsInfo) GetStartTimeOk() (*string, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *FarmRecurringMaintenanceSettingsInfo) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given string and assigns it to the StartTime field.
func (o *FarmRecurringMaintenanceSettingsInfo) SetStartTime(v string) {
	o.StartTime = &v
}

func (o FarmRecurringMaintenanceSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaintenancePeriod != nil {
		toSerialize["maintenance_period"] = o.MaintenancePeriod
	}
	if o.MaintenancePeriodFrequency != nil {
		toSerialize["maintenance_period_frequency"] = o.MaintenancePeriodFrequency
	}
	if o.StartIndex != nil {
		toSerialize["start_index"] = o.StartIndex
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableFarmRecurringMaintenanceSettingsInfo struct {
	value *FarmRecurringMaintenanceSettingsInfo
	isSet bool
}

func (v NullableFarmRecurringMaintenanceSettingsInfo) Get() *FarmRecurringMaintenanceSettingsInfo {
	return v.value
}

func (v *NullableFarmRecurringMaintenanceSettingsInfo) Set(val *FarmRecurringMaintenanceSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmRecurringMaintenanceSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmRecurringMaintenanceSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmRecurringMaintenanceSettingsInfo(val *FarmRecurringMaintenanceSettingsInfo) *NullableFarmRecurringMaintenanceSettingsInfo {
	return &NullableFarmRecurringMaintenanceSettingsInfo{value: val, isSet: true}
}

func (v NullableFarmRecurringMaintenanceSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmRecurringMaintenanceSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
