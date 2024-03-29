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

// SpecifiedNamesValidationSpec Information required to validate specified names.
type SpecifiedNamesValidationSpec struct {
	// Indicates whether desktop pool is dedicated or floating. Default value is false.
	Dedicated *bool `json:"dedicated,omitempty"`
	// ID of the desktop pool to which the manually defined virtual machines will belong. This is required only if virtual machines are being added to an existing pool.
	Id *string `json:"id,omitempty"`
	// List of manually defined virtual machines and users.
	NamesSpec []NamesSpec `json:"names_spec"`
}

// NewSpecifiedNamesValidationSpec instantiates a new SpecifiedNamesValidationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecifiedNamesValidationSpec(namesSpec []NamesSpec) *SpecifiedNamesValidationSpec {
	this := SpecifiedNamesValidationSpec{}
	this.NamesSpec = namesSpec
	return &this
}

// NewSpecifiedNamesValidationSpecWithDefaults instantiates a new SpecifiedNamesValidationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecifiedNamesValidationSpecWithDefaults() *SpecifiedNamesValidationSpec {
	this := SpecifiedNamesValidationSpec{}
	return &this
}

// GetDedicated returns the Dedicated field value if set, zero value otherwise.
func (o *SpecifiedNamesValidationSpec) GetDedicated() bool {
	if o == nil || o.Dedicated == nil {
		var ret bool
		return ret
	}
	return *o.Dedicated
}

// GetDedicatedOk returns a tuple with the Dedicated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecifiedNamesValidationSpec) GetDedicatedOk() (*bool, bool) {
	if o == nil || o.Dedicated == nil {
		return nil, false
	}
	return o.Dedicated, true
}

// HasDedicated returns a boolean if a field has been set.
func (o *SpecifiedNamesValidationSpec) HasDedicated() bool {
	if o != nil && o.Dedicated != nil {
		return true
	}

	return false
}

// SetDedicated gets a reference to the given bool and assigns it to the Dedicated field.
func (o *SpecifiedNamesValidationSpec) SetDedicated(v bool) {
	o.Dedicated = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SpecifiedNamesValidationSpec) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecifiedNamesValidationSpec) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SpecifiedNamesValidationSpec) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SpecifiedNamesValidationSpec) SetId(v string) {
	o.Id = &v
}

// GetNamesSpec returns the NamesSpec field value
func (o *SpecifiedNamesValidationSpec) GetNamesSpec() []NamesSpec {
	if o == nil {
		var ret []NamesSpec
		return ret
	}

	return o.NamesSpec
}

// GetNamesSpecOk returns a tuple with the NamesSpec field value
// and a boolean to check if the value has been set.
func (o *SpecifiedNamesValidationSpec) GetNamesSpecOk() (*[]NamesSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NamesSpec, true
}

// SetNamesSpec sets field value
func (o *SpecifiedNamesValidationSpec) SetNamesSpec(v []NamesSpec) {
	o.NamesSpec = v
}

func (o SpecifiedNamesValidationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dedicated != nil {
		toSerialize["dedicated"] = o.Dedicated
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["names_spec"] = o.NamesSpec
	}
	return json.Marshal(toSerialize)
}

type NullableSpecifiedNamesValidationSpec struct {
	value *SpecifiedNamesValidationSpec
	isSet bool
}

func (v NullableSpecifiedNamesValidationSpec) Get() *SpecifiedNamesValidationSpec {
	return v.value
}

func (v *NullableSpecifiedNamesValidationSpec) Set(val *SpecifiedNamesValidationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecifiedNamesValidationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecifiedNamesValidationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecifiedNamesValidationSpec(val *SpecifiedNamesValidationSpec) *NullableSpecifiedNamesValidationSpec {
	return &NullableSpecifiedNamesValidationSpec{value: val, isSet: true}
}

func (v NullableSpecifiedNamesValidationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecifiedNamesValidationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
