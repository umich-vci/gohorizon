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

// ADDomainAuxiliaryAccountDeleteSpec Specification to delete auxiliary accounts from untrusted domain.
type ADDomainAuxiliaryAccountDeleteSpec struct {
	// Auxiliary account ID's.
	AuxiliaryAccountIds []string `json:"auxiliary_account_ids"`
}

// NewADDomainAuxiliaryAccountDeleteSpec instantiates a new ADDomainAuxiliaryAccountDeleteSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewADDomainAuxiliaryAccountDeleteSpec(auxiliaryAccountIds []string) *ADDomainAuxiliaryAccountDeleteSpec {
	this := ADDomainAuxiliaryAccountDeleteSpec{}
	this.AuxiliaryAccountIds = auxiliaryAccountIds
	return &this
}

// NewADDomainAuxiliaryAccountDeleteSpecWithDefaults instantiates a new ADDomainAuxiliaryAccountDeleteSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewADDomainAuxiliaryAccountDeleteSpecWithDefaults() *ADDomainAuxiliaryAccountDeleteSpec {
	this := ADDomainAuxiliaryAccountDeleteSpec{}
	return &this
}

// GetAuxiliaryAccountIds returns the AuxiliaryAccountIds field value
func (o *ADDomainAuxiliaryAccountDeleteSpec) GetAuxiliaryAccountIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AuxiliaryAccountIds
}

// GetAuxiliaryAccountIdsOk returns a tuple with the AuxiliaryAccountIds field value
// and a boolean to check if the value has been set.
func (o *ADDomainAuxiliaryAccountDeleteSpec) GetAuxiliaryAccountIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuxiliaryAccountIds, true
}

// SetAuxiliaryAccountIds sets field value
func (o *ADDomainAuxiliaryAccountDeleteSpec) SetAuxiliaryAccountIds(v []string) {
	o.AuxiliaryAccountIds = v
}

func (o ADDomainAuxiliaryAccountDeleteSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auxiliary_account_ids"] = o.AuxiliaryAccountIds
	}
	return json.Marshal(toSerialize)
}

type NullableADDomainAuxiliaryAccountDeleteSpec struct {
	value *ADDomainAuxiliaryAccountDeleteSpec
	isSet bool
}

func (v NullableADDomainAuxiliaryAccountDeleteSpec) Get() *ADDomainAuxiliaryAccountDeleteSpec {
	return v.value
}

func (v *NullableADDomainAuxiliaryAccountDeleteSpec) Set(val *ADDomainAuxiliaryAccountDeleteSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableADDomainAuxiliaryAccountDeleteSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableADDomainAuxiliaryAccountDeleteSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableADDomainAuxiliaryAccountDeleteSpec(val *ADDomainAuxiliaryAccountDeleteSpec) *NullableADDomainAuxiliaryAccountDeleteSpec {
	return &NullableADDomainAuxiliaryAccountDeleteSpec{value: val, isSet: true}
}

func (v NullableADDomainAuxiliaryAccountDeleteSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableADDomainAuxiliaryAccountDeleteSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
