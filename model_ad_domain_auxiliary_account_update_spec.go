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

// ADDomainAuxiliaryAccountUpdateSpec Specification to update auxiliary accounts of untrusted domain
type ADDomainAuxiliaryAccountUpdateSpec struct {
	// Auxiliary service account credentials.
	AuxiliaryAccounts []AuxiliaryAccountUpdateData `json:"auxiliary_accounts"`
}

// NewADDomainAuxiliaryAccountUpdateSpec instantiates a new ADDomainAuxiliaryAccountUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewADDomainAuxiliaryAccountUpdateSpec(auxiliaryAccounts []AuxiliaryAccountUpdateData) *ADDomainAuxiliaryAccountUpdateSpec {
	this := ADDomainAuxiliaryAccountUpdateSpec{}
	this.AuxiliaryAccounts = auxiliaryAccounts
	return &this
}

// NewADDomainAuxiliaryAccountUpdateSpecWithDefaults instantiates a new ADDomainAuxiliaryAccountUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewADDomainAuxiliaryAccountUpdateSpecWithDefaults() *ADDomainAuxiliaryAccountUpdateSpec {
	this := ADDomainAuxiliaryAccountUpdateSpec{}
	return &this
}

// GetAuxiliaryAccounts returns the AuxiliaryAccounts field value
func (o *ADDomainAuxiliaryAccountUpdateSpec) GetAuxiliaryAccounts() []AuxiliaryAccountUpdateData {
	if o == nil {
		var ret []AuxiliaryAccountUpdateData
		return ret
	}

	return o.AuxiliaryAccounts
}

// GetAuxiliaryAccountsOk returns a tuple with the AuxiliaryAccounts field value
// and a boolean to check if the value has been set.
func (o *ADDomainAuxiliaryAccountUpdateSpec) GetAuxiliaryAccountsOk() (*[]AuxiliaryAccountUpdateData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuxiliaryAccounts, true
}

// SetAuxiliaryAccounts sets field value
func (o *ADDomainAuxiliaryAccountUpdateSpec) SetAuxiliaryAccounts(v []AuxiliaryAccountUpdateData) {
	o.AuxiliaryAccounts = v
}

func (o ADDomainAuxiliaryAccountUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auxiliary_accounts"] = o.AuxiliaryAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableADDomainAuxiliaryAccountUpdateSpec struct {
	value *ADDomainAuxiliaryAccountUpdateSpec
	isSet bool
}

func (v NullableADDomainAuxiliaryAccountUpdateSpec) Get() *ADDomainAuxiliaryAccountUpdateSpec {
	return v.value
}

func (v *NullableADDomainAuxiliaryAccountUpdateSpec) Set(val *ADDomainAuxiliaryAccountUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableADDomainAuxiliaryAccountUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableADDomainAuxiliaryAccountUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableADDomainAuxiliaryAccountUpdateSpec(val *ADDomainAuxiliaryAccountUpdateSpec) *NullableADDomainAuxiliaryAccountUpdateSpec {
	return &NullableADDomainAuxiliaryAccountUpdateSpec{value: val, isSet: true}
}

func (v NullableADDomainAuxiliaryAccountUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableADDomainAuxiliaryAccountUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
