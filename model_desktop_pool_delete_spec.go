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

// DesktopPoolDeleteSpec Information required to delete a desktop pool.
type DesktopPoolDeleteSpec struct {
	// Determines whether the machine VMs should be deleted from vCenter Server.<br> This must be false for RDS and unmanaged desktop pools and true for Instant Clone desktop pools.<br>Default value is true for IC pools and false for all other types of desktop pools. <br>
	DeleteFromDisk *bool `json:"delete_from_disk,omitempty"`
}

// NewDesktopPoolDeleteSpec instantiates a new DesktopPoolDeleteSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolDeleteSpec() *DesktopPoolDeleteSpec {
	this := DesktopPoolDeleteSpec{}
	return &this
}

// NewDesktopPoolDeleteSpecWithDefaults instantiates a new DesktopPoolDeleteSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolDeleteSpecWithDefaults() *DesktopPoolDeleteSpec {
	this := DesktopPoolDeleteSpec{}
	return &this
}

// GetDeleteFromDisk returns the DeleteFromDisk field value if set, zero value otherwise.
func (o *DesktopPoolDeleteSpec) GetDeleteFromDisk() bool {
	if o == nil || o.DeleteFromDisk == nil {
		var ret bool
		return ret
	}
	return *o.DeleteFromDisk
}

// GetDeleteFromDiskOk returns a tuple with the DeleteFromDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDeleteSpec) GetDeleteFromDiskOk() (*bool, bool) {
	if o == nil || o.DeleteFromDisk == nil {
		return nil, false
	}
	return o.DeleteFromDisk, true
}

// HasDeleteFromDisk returns a boolean if a field has been set.
func (o *DesktopPoolDeleteSpec) HasDeleteFromDisk() bool {
	if o != nil && o.DeleteFromDisk != nil {
		return true
	}

	return false
}

// SetDeleteFromDisk gets a reference to the given bool and assigns it to the DeleteFromDisk field.
func (o *DesktopPoolDeleteSpec) SetDeleteFromDisk(v bool) {
	o.DeleteFromDisk = &v
}

func (o DesktopPoolDeleteSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeleteFromDisk != nil {
		toSerialize["delete_from_disk"] = o.DeleteFromDisk
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolDeleteSpec struct {
	value *DesktopPoolDeleteSpec
	isSet bool
}

func (v NullableDesktopPoolDeleteSpec) Get() *DesktopPoolDeleteSpec {
	return v.value
}

func (v *NullableDesktopPoolDeleteSpec) Set(val *DesktopPoolDeleteSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolDeleteSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolDeleteSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolDeleteSpec(val *DesktopPoolDeleteSpec) *NullableDesktopPoolDeleteSpec {
	return &NullableDesktopPoolDeleteSpec{value: val, isSet: true}
}

func (v NullableDesktopPoolDeleteSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolDeleteSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
