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

// DesktopPoolVirtualMachinePatternNamingSettings Settings related to specification of VMs with a naming pattern.
type DesktopPoolVirtualMachinePatternNamingSettings struct {
	// Maximum number of machines in the desktop pool.
	MaxNumberOfMachines *int32 `json:"max_number_of_machines,omitempty"`
	// The minimum number of machines to have provisioned if on demand provisioning is selected.
	MinNumberOfMachines *int32 `json:"min_number_of_machines,omitempty"`
	// Virtual machines will be named according to the specified naming pattern.<br>Supported Filters: 'Equals'.
	NamingPattern *string `json:"naming_pattern,omitempty"`
	// Number of spare powered on machines.
	NumberOfSpareMachines *int32 `json:"number_of_spare_machines,omitempty"`
	// Determines when the machines are provisioned. * ON_DEMAND: Provision machines on demand. * UP_FRONT: Provision all machines up-front.
	ProvisioningTime *string `json:"provisioning_time,omitempty"`
}

// NewDesktopPoolVirtualMachinePatternNamingSettings instantiates a new DesktopPoolVirtualMachinePatternNamingSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolVirtualMachinePatternNamingSettings() *DesktopPoolVirtualMachinePatternNamingSettings {
	this := DesktopPoolVirtualMachinePatternNamingSettings{}
	return &this
}

// NewDesktopPoolVirtualMachinePatternNamingSettingsWithDefaults instantiates a new DesktopPoolVirtualMachinePatternNamingSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolVirtualMachinePatternNamingSettingsWithDefaults() *DesktopPoolVirtualMachinePatternNamingSettings {
	this := DesktopPoolVirtualMachinePatternNamingSettings{}
	return &this
}

// GetMaxNumberOfMachines returns the MaxNumberOfMachines field value if set, zero value otherwise.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetMaxNumberOfMachines() int32 {
	if o == nil || o.MaxNumberOfMachines == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfMachines
}

// GetMaxNumberOfMachinesOk returns a tuple with the MaxNumberOfMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetMaxNumberOfMachinesOk() (*int32, bool) {
	if o == nil || o.MaxNumberOfMachines == nil {
		return nil, false
	}
	return o.MaxNumberOfMachines, true
}

// HasMaxNumberOfMachines returns a boolean if a field has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasMaxNumberOfMachines() bool {
	if o != nil && o.MaxNumberOfMachines != nil {
		return true
	}

	return false
}

// SetMaxNumberOfMachines gets a reference to the given int32 and assigns it to the MaxNumberOfMachines field.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetMaxNumberOfMachines(v int32) {
	o.MaxNumberOfMachines = &v
}

// GetMinNumberOfMachines returns the MinNumberOfMachines field value if set, zero value otherwise.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetMinNumberOfMachines() int32 {
	if o == nil || o.MinNumberOfMachines == nil {
		var ret int32
		return ret
	}
	return *o.MinNumberOfMachines
}

// GetMinNumberOfMachinesOk returns a tuple with the MinNumberOfMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetMinNumberOfMachinesOk() (*int32, bool) {
	if o == nil || o.MinNumberOfMachines == nil {
		return nil, false
	}
	return o.MinNumberOfMachines, true
}

// HasMinNumberOfMachines returns a boolean if a field has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasMinNumberOfMachines() bool {
	if o != nil && o.MinNumberOfMachines != nil {
		return true
	}

	return false
}

// SetMinNumberOfMachines gets a reference to the given int32 and assigns it to the MinNumberOfMachines field.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetMinNumberOfMachines(v int32) {
	o.MinNumberOfMachines = &v
}

// GetNamingPattern returns the NamingPattern field value if set, zero value otherwise.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetNamingPattern() string {
	if o == nil || o.NamingPattern == nil {
		var ret string
		return ret
	}
	return *o.NamingPattern
}

// GetNamingPatternOk returns a tuple with the NamingPattern field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetNamingPatternOk() (*string, bool) {
	if o == nil || o.NamingPattern == nil {
		return nil, false
	}
	return o.NamingPattern, true
}

// HasNamingPattern returns a boolean if a field has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasNamingPattern() bool {
	if o != nil && o.NamingPattern != nil {
		return true
	}

	return false
}

// SetNamingPattern gets a reference to the given string and assigns it to the NamingPattern field.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetNamingPattern(v string) {
	o.NamingPattern = &v
}

// GetNumberOfSpareMachines returns the NumberOfSpareMachines field value if set, zero value otherwise.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetNumberOfSpareMachines() int32 {
	if o == nil || o.NumberOfSpareMachines == nil {
		var ret int32
		return ret
	}
	return *o.NumberOfSpareMachines
}

// GetNumberOfSpareMachinesOk returns a tuple with the NumberOfSpareMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetNumberOfSpareMachinesOk() (*int32, bool) {
	if o == nil || o.NumberOfSpareMachines == nil {
		return nil, false
	}
	return o.NumberOfSpareMachines, true
}

// HasNumberOfSpareMachines returns a boolean if a field has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasNumberOfSpareMachines() bool {
	if o != nil && o.NumberOfSpareMachines != nil {
		return true
	}

	return false
}

// SetNumberOfSpareMachines gets a reference to the given int32 and assigns it to the NumberOfSpareMachines field.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetNumberOfSpareMachines(v int32) {
	o.NumberOfSpareMachines = &v
}

// GetProvisioningTime returns the ProvisioningTime field value if set, zero value otherwise.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetProvisioningTime() string {
	if o == nil || o.ProvisioningTime == nil {
		var ret string
		return ret
	}
	return *o.ProvisioningTime
}

// GetProvisioningTimeOk returns a tuple with the ProvisioningTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) GetProvisioningTimeOk() (*string, bool) {
	if o == nil || o.ProvisioningTime == nil {
		return nil, false
	}
	return o.ProvisioningTime, true
}

// HasProvisioningTime returns a boolean if a field has been set.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) HasProvisioningTime() bool {
	if o != nil && o.ProvisioningTime != nil {
		return true
	}

	return false
}

// SetProvisioningTime gets a reference to the given string and assigns it to the ProvisioningTime field.
func (o *DesktopPoolVirtualMachinePatternNamingSettings) SetProvisioningTime(v string) {
	o.ProvisioningTime = &v
}

func (o DesktopPoolVirtualMachinePatternNamingSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxNumberOfMachines != nil {
		toSerialize["max_number_of_machines"] = o.MaxNumberOfMachines
	}
	if o.MinNumberOfMachines != nil {
		toSerialize["min_number_of_machines"] = o.MinNumberOfMachines
	}
	if o.NamingPattern != nil {
		toSerialize["naming_pattern"] = o.NamingPattern
	}
	if o.NumberOfSpareMachines != nil {
		toSerialize["number_of_spare_machines"] = o.NumberOfSpareMachines
	}
	if o.ProvisioningTime != nil {
		toSerialize["provisioning_time"] = o.ProvisioningTime
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolVirtualMachinePatternNamingSettings struct {
	value *DesktopPoolVirtualMachinePatternNamingSettings
	isSet bool
}

func (v NullableDesktopPoolVirtualMachinePatternNamingSettings) Get() *DesktopPoolVirtualMachinePatternNamingSettings {
	return v.value
}

func (v *NullableDesktopPoolVirtualMachinePatternNamingSettings) Set(val *DesktopPoolVirtualMachinePatternNamingSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolVirtualMachinePatternNamingSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolVirtualMachinePatternNamingSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolVirtualMachinePatternNamingSettings(val *DesktopPoolVirtualMachinePatternNamingSettings) *NullableDesktopPoolVirtualMachinePatternNamingSettings {
	return &NullableDesktopPoolVirtualMachinePatternNamingSettings{value: val, isSet: true}
}

func (v NullableDesktopPoolVirtualMachinePatternNamingSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolVirtualMachinePatternNamingSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
