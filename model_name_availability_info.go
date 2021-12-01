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

// NameAvailabilityInfo Information about name availability of an inventory resource.
type NameAvailabilityInfo struct {
	// Indicates whether name is available for resource creation.
	Available *bool `json:"available,omitempty"`
}

// NewNameAvailabilityInfo instantiates a new NameAvailabilityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNameAvailabilityInfo() *NameAvailabilityInfo {
	this := NameAvailabilityInfo{}
	return &this
}

// NewNameAvailabilityInfoWithDefaults instantiates a new NameAvailabilityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNameAvailabilityInfoWithDefaults() *NameAvailabilityInfo {
	this := NameAvailabilityInfo{}
	return &this
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *NameAvailabilityInfo) GetAvailable() bool {
	if o == nil || o.Available == nil {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NameAvailabilityInfo) GetAvailableOk() (*bool, bool) {
	if o == nil || o.Available == nil {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *NameAvailabilityInfo) HasAvailable() bool {
	if o != nil && o.Available != nil {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *NameAvailabilityInfo) SetAvailable(v bool) {
	o.Available = &v
}

func (o NameAvailabilityInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Available != nil {
		toSerialize["available"] = o.Available
	}
	return json.Marshal(toSerialize)
}

type NullableNameAvailabilityInfo struct {
	value *NameAvailabilityInfo
	isSet bool
}

func (v NullableNameAvailabilityInfo) Get() *NameAvailabilityInfo {
	return v.value
}

func (v *NullableNameAvailabilityInfo) Set(val *NameAvailabilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNameAvailabilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNameAvailabilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNameAvailabilityInfo(val *NameAvailabilityInfo) *NullableNameAvailabilityInfo {
	return &NullableNameAvailabilityInfo{value: val, isSet: true}
}

func (v NullableNameAvailabilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNameAvailabilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
