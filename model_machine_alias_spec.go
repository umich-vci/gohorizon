/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2106
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// MachineAliasSpec The specification for updating machine aliases of assigned users.
type MachineAliasSpec struct {
	// Sid of the user
	AdUserId *string `json:"ad_user_id,omitempty"`
	// Alias name of the user.
	AliasName *string `json:"alias_name,omitempty"`
}

// NewMachineAliasSpec instantiates a new MachineAliasSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineAliasSpec() *MachineAliasSpec {
	this := MachineAliasSpec{}
	return &this
}

// NewMachineAliasSpecWithDefaults instantiates a new MachineAliasSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineAliasSpecWithDefaults() *MachineAliasSpec {
	this := MachineAliasSpec{}
	return &this
}

// GetAdUserId returns the AdUserId field value if set, zero value otherwise.
func (o *MachineAliasSpec) GetAdUserId() string {
	if o == nil || o.AdUserId == nil {
		var ret string
		return ret
	}
	return *o.AdUserId
}

// GetAdUserIdOk returns a tuple with the AdUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAliasSpec) GetAdUserIdOk() (*string, bool) {
	if o == nil || o.AdUserId == nil {
		return nil, false
	}
	return o.AdUserId, true
}

// HasAdUserId returns a boolean if a field has been set.
func (o *MachineAliasSpec) HasAdUserId() bool {
	if o != nil && o.AdUserId != nil {
		return true
	}

	return false
}

// SetAdUserId gets a reference to the given string and assigns it to the AdUserId field.
func (o *MachineAliasSpec) SetAdUserId(v string) {
	o.AdUserId = &v
}

// GetAliasName returns the AliasName field value if set, zero value otherwise.
func (o *MachineAliasSpec) GetAliasName() string {
	if o == nil || o.AliasName == nil {
		var ret string
		return ret
	}
	return *o.AliasName
}

// GetAliasNameOk returns a tuple with the AliasName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineAliasSpec) GetAliasNameOk() (*string, bool) {
	if o == nil || o.AliasName == nil {
		return nil, false
	}
	return o.AliasName, true
}

// HasAliasName returns a boolean if a field has been set.
func (o *MachineAliasSpec) HasAliasName() bool {
	if o != nil && o.AliasName != nil {
		return true
	}

	return false
}

// SetAliasName gets a reference to the given string and assigns it to the AliasName field.
func (o *MachineAliasSpec) SetAliasName(v string) {
	o.AliasName = &v
}

func (o MachineAliasSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdUserId != nil {
		toSerialize["ad_user_id"] = o.AdUserId
	}
	if o.AliasName != nil {
		toSerialize["alias_name"] = o.AliasName
	}
	return json.Marshal(toSerialize)
}

type NullableMachineAliasSpec struct {
	value *MachineAliasSpec
	isSet bool
}

func (v NullableMachineAliasSpec) Get() *MachineAliasSpec {
	return v.value
}

func (v *NullableMachineAliasSpec) Set(val *MachineAliasSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineAliasSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineAliasSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineAliasSpec(val *MachineAliasSpec) *NullableMachineAliasSpec {
	return &NullableMachineAliasSpec{value: val, isSet: true}
}

func (v NullableMachineAliasSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineAliasSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}