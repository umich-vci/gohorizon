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

// DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec Applicable To: Automated desktop pool. <br>Specified naming setting for Automated desktop pool.
type DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec struct {
	// Number of unassigned machines kept powered on. This value must be lesser than or equal to the number of specified names.
	NumUnassignedMachinesKeptPoweredOn int32 `json:"num_unassigned_machines_kept_powered_on"`
	// Initial specified names of machines in the desktop pool.
	SpecifiedNames *[]MachineSpecifiedName `json:"specified_names,omitempty"`
	// Allows virtual machines to be customized manually before users can log in and access them. This mode must be exited manually.
	StartMachinesInMaintenanceMode bool `json:"start_machines_in_maintenance_mode"`
}

// NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec instantiates a new DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec(numUnassignedMachinesKeptPoweredOn int32, startMachinesInMaintenanceMode bool) *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec {
	this := DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec{}
	this.NumUnassignedMachinesKeptPoweredOn = numUnassignedMachinesKeptPoweredOn
	this.StartMachinesInMaintenanceMode = startMachinesInMaintenanceMode
	return &this
}

// NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpecWithDefaults() *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec {
	this := DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec{}
	return &this
}

// GetNumUnassignedMachinesKeptPoweredOn returns the NumUnassignedMachinesKeptPoweredOn field value
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetNumUnassignedMachinesKeptPoweredOn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumUnassignedMachinesKeptPoweredOn
}

// GetNumUnassignedMachinesKeptPoweredOnOk returns a tuple with the NumUnassignedMachinesKeptPoweredOn field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetNumUnassignedMachinesKeptPoweredOnOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumUnassignedMachinesKeptPoweredOn, true
}

// SetNumUnassignedMachinesKeptPoweredOn sets field value
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) SetNumUnassignedMachinesKeptPoweredOn(v int32) {
	o.NumUnassignedMachinesKeptPoweredOn = v
}

// GetSpecifiedNames returns the SpecifiedNames field value if set, zero value otherwise.
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetSpecifiedNames() []MachineSpecifiedName {
	if o == nil || o.SpecifiedNames == nil {
		var ret []MachineSpecifiedName
		return ret
	}
	return *o.SpecifiedNames
}

// GetSpecifiedNamesOk returns a tuple with the SpecifiedNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetSpecifiedNamesOk() (*[]MachineSpecifiedName, bool) {
	if o == nil || o.SpecifiedNames == nil {
		return nil, false
	}
	return o.SpecifiedNames, true
}

// HasSpecifiedNames returns a boolean if a field has been set.
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) HasSpecifiedNames() bool {
	if o != nil && o.SpecifiedNames != nil {
		return true
	}

	return false
}

// SetSpecifiedNames gets a reference to the given []MachineSpecifiedName and assigns it to the SpecifiedNames field.
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) SetSpecifiedNames(v []MachineSpecifiedName) {
	o.SpecifiedNames = &v
}

// GetStartMachinesInMaintenanceMode returns the StartMachinesInMaintenanceMode field value
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetStartMachinesInMaintenanceMode() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StartMachinesInMaintenanceMode
}

// GetStartMachinesInMaintenanceModeOk returns a tuple with the StartMachinesInMaintenanceMode field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) GetStartMachinesInMaintenanceModeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartMachinesInMaintenanceMode, true
}

// SetStartMachinesInMaintenanceMode sets field value
func (o *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) SetStartMachinesInMaintenanceMode(v bool) {
	o.StartMachinesInMaintenanceMode = v
}

func (o DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["num_unassigned_machines_kept_powered_on"] = o.NumUnassignedMachinesKeptPoweredOn
	}
	if o.SpecifiedNames != nil {
		toSerialize["specified_names"] = o.SpecifiedNames
	}
	if true {
		toSerialize["start_machines_in_maintenance_mode"] = o.StartMachinesInMaintenanceMode
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec struct {
	value *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec
	isSet bool
}

func (v NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) Get() *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec {
	return v.value
}

func (v *NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) Set(val *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec(val *DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) *NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec {
	return &NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
