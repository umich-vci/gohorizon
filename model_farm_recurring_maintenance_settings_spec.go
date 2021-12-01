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

// FarmRecurringMaintenanceSettingsSpec Settings for recurring maintenance operation for the farm. This is Required only if maintenance_mode is set to RECURRING
type FarmRecurringMaintenanceSettingsSpec struct {
	// This represents the frequency at which to perform recurring maintenance. * DAILY: Daily recurring maintenance * WEEKLY: Weekly recurring maintenance * MONTHLY: Monthly recurring maintenance
	MaintenancePeriod string `json:"maintenance_period"`
	// Indicates how frequently to repeat maintenance, expressed as multiple of the maintenance period. e.g. Every 2 weeks. Default value is 1.
	MaintenancePeriodFrequency *int32 `json:"maintenance_period_frequency,omitempty"`
	// Start index for weekly or monthly maintenance.This required if maintenance_period is set to either WEEKLY or MONTHLY. Weekly: 1-7 (Sun-Sat), Monthly: 1-31
	StartIndex *int32 `json:"start_index,omitempty"`
	// Start time configured for the recurring maintenance. This should be in the form hh:mm in 24 hours format.
	StartTime string `json:"start_time"`
}

// NewFarmRecurringMaintenanceSettingsSpec instantiates a new FarmRecurringMaintenanceSettingsSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmRecurringMaintenanceSettingsSpec(maintenancePeriod string, startTime string) *FarmRecurringMaintenanceSettingsSpec {
	this := FarmRecurringMaintenanceSettingsSpec{}
	this.MaintenancePeriod = maintenancePeriod
	this.StartTime = startTime
	return &this
}

// NewFarmRecurringMaintenanceSettingsSpecWithDefaults instantiates a new FarmRecurringMaintenanceSettingsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmRecurringMaintenanceSettingsSpecWithDefaults() *FarmRecurringMaintenanceSettingsSpec {
	this := FarmRecurringMaintenanceSettingsSpec{}
	return &this
}

// GetMaintenancePeriod returns the MaintenancePeriod field value
func (o *FarmRecurringMaintenanceSettingsSpec) GetMaintenancePeriod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaintenancePeriod
}

// GetMaintenancePeriodOk returns a tuple with the MaintenancePeriod field value
// and a boolean to check if the value has been set.
func (o *FarmRecurringMaintenanceSettingsSpec) GetMaintenancePeriodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaintenancePeriod, true
}

// SetMaintenancePeriod sets field value
func (o *FarmRecurringMaintenanceSettingsSpec) SetMaintenancePeriod(v string) {
	o.MaintenancePeriod = v
}

// GetMaintenancePeriodFrequency returns the MaintenancePeriodFrequency field value if set, zero value otherwise.
func (o *FarmRecurringMaintenanceSettingsSpec) GetMaintenancePeriodFrequency() int32 {
	if o == nil || o.MaintenancePeriodFrequency == nil {
		var ret int32
		return ret
	}
	return *o.MaintenancePeriodFrequency
}

// GetMaintenancePeriodFrequencyOk returns a tuple with the MaintenancePeriodFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmRecurringMaintenanceSettingsSpec) GetMaintenancePeriodFrequencyOk() (*int32, bool) {
	if o == nil || o.MaintenancePeriodFrequency == nil {
		return nil, false
	}
	return o.MaintenancePeriodFrequency, true
}

// HasMaintenancePeriodFrequency returns a boolean if a field has been set.
func (o *FarmRecurringMaintenanceSettingsSpec) HasMaintenancePeriodFrequency() bool {
	if o != nil && o.MaintenancePeriodFrequency != nil {
		return true
	}

	return false
}

// SetMaintenancePeriodFrequency gets a reference to the given int32 and assigns it to the MaintenancePeriodFrequency field.
func (o *FarmRecurringMaintenanceSettingsSpec) SetMaintenancePeriodFrequency(v int32) {
	o.MaintenancePeriodFrequency = &v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *FarmRecurringMaintenanceSettingsSpec) GetStartIndex() int32 {
	if o == nil || o.StartIndex == nil {
		var ret int32
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmRecurringMaintenanceSettingsSpec) GetStartIndexOk() (*int32, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *FarmRecurringMaintenanceSettingsSpec) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int32 and assigns it to the StartIndex field.
func (o *FarmRecurringMaintenanceSettingsSpec) SetStartIndex(v int32) {
	o.StartIndex = &v
}

// GetStartTime returns the StartTime field value
func (o *FarmRecurringMaintenanceSettingsSpec) GetStartTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *FarmRecurringMaintenanceSettingsSpec) GetStartTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *FarmRecurringMaintenanceSettingsSpec) SetStartTime(v string) {
	o.StartTime = v
}

func (o FarmRecurringMaintenanceSettingsSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["maintenance_period"] = o.MaintenancePeriod
	}
	if o.MaintenancePeriodFrequency != nil {
		toSerialize["maintenance_period_frequency"] = o.MaintenancePeriodFrequency
	}
	if o.StartIndex != nil {
		toSerialize["start_index"] = o.StartIndex
	}
	if true {
		toSerialize["start_time"] = o.StartTime
	}
	return json.Marshal(toSerialize)
}

type NullableFarmRecurringMaintenanceSettingsSpec struct {
	value *FarmRecurringMaintenanceSettingsSpec
	isSet bool
}

func (v NullableFarmRecurringMaintenanceSettingsSpec) Get() *FarmRecurringMaintenanceSettingsSpec {
	return v.value
}

func (v *NullableFarmRecurringMaintenanceSettingsSpec) Set(val *FarmRecurringMaintenanceSettingsSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmRecurringMaintenanceSettingsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmRecurringMaintenanceSettingsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmRecurringMaintenanceSettingsSpec(val *FarmRecurringMaintenanceSettingsSpec) *NullableFarmRecurringMaintenanceSettingsSpec {
	return &NullableFarmRecurringMaintenanceSettingsSpec{value: val, isSet: true}
}

func (v NullableFarmRecurringMaintenanceSettingsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmRecurringMaintenanceSettingsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}