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

// FarmCustomizationSettingsUpdateSpec Customization settings for the automated farm.
type FarmCustomizationSettingsUpdateSpec struct {
	// Instant Clone Engine Active Directory container for clone prep.
	AdContainerRdn                 string                                       `json:"ad_container_rdn"`
	CloneprepCustomizationSettings FarmCloneprepCustomizationSettingsUpdateSpec `json:"cloneprep_customization_settings"`
	// Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names.
	ReusePreExistingAccounts bool `json:"reuse_pre_existing_accounts"`
}

// NewFarmCustomizationSettingsUpdateSpec instantiates a new FarmCustomizationSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmCustomizationSettingsUpdateSpec(adContainerRdn string, cloneprepCustomizationSettings FarmCloneprepCustomizationSettingsUpdateSpec, reusePreExistingAccounts bool) *FarmCustomizationSettingsUpdateSpec {
	this := FarmCustomizationSettingsUpdateSpec{}
	this.AdContainerRdn = adContainerRdn
	this.CloneprepCustomizationSettings = cloneprepCustomizationSettings
	this.ReusePreExistingAccounts = reusePreExistingAccounts
	return &this
}

// NewFarmCustomizationSettingsUpdateSpecWithDefaults instantiates a new FarmCustomizationSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmCustomizationSettingsUpdateSpecWithDefaults() *FarmCustomizationSettingsUpdateSpec {
	this := FarmCustomizationSettingsUpdateSpec{}
	return &this
}

// GetAdContainerRdn returns the AdContainerRdn field value
func (o *FarmCustomizationSettingsUpdateSpec) GetAdContainerRdn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdContainerRdn
}

// GetAdContainerRdnOk returns a tuple with the AdContainerRdn field value
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsUpdateSpec) GetAdContainerRdnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdContainerRdn, true
}

// SetAdContainerRdn sets field value
func (o *FarmCustomizationSettingsUpdateSpec) SetAdContainerRdn(v string) {
	o.AdContainerRdn = v
}

// GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field value
func (o *FarmCustomizationSettingsUpdateSpec) GetCloneprepCustomizationSettings() FarmCloneprepCustomizationSettingsUpdateSpec {
	if o == nil {
		var ret FarmCloneprepCustomizationSettingsUpdateSpec
		return ret
	}

	return o.CloneprepCustomizationSettings
}

// GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field value
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsUpdateSpec) GetCloneprepCustomizationSettingsOk() (*FarmCloneprepCustomizationSettingsUpdateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloneprepCustomizationSettings, true
}

// SetCloneprepCustomizationSettings sets field value
func (o *FarmCustomizationSettingsUpdateSpec) SetCloneprepCustomizationSettings(v FarmCloneprepCustomizationSettingsUpdateSpec) {
	o.CloneprepCustomizationSettings = v
}

// GetReusePreExistingAccounts returns the ReusePreExistingAccounts field value
func (o *FarmCustomizationSettingsUpdateSpec) GetReusePreExistingAccounts() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReusePreExistingAccounts
}

// GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field value
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsUpdateSpec) GetReusePreExistingAccountsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReusePreExistingAccounts, true
}

// SetReusePreExistingAccounts sets field value
func (o *FarmCustomizationSettingsUpdateSpec) SetReusePreExistingAccounts(v bool) {
	o.ReusePreExistingAccounts = v
}

func (o FarmCustomizationSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ad_container_rdn"] = o.AdContainerRdn
	}
	if true {
		toSerialize["cloneprep_customization_settings"] = o.CloneprepCustomizationSettings
	}
	if true {
		toSerialize["reuse_pre_existing_accounts"] = o.ReusePreExistingAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableFarmCustomizationSettingsUpdateSpec struct {
	value *FarmCustomizationSettingsUpdateSpec
	isSet bool
}

func (v NullableFarmCustomizationSettingsUpdateSpec) Get() *FarmCustomizationSettingsUpdateSpec {
	return v.value
}

func (v *NullableFarmCustomizationSettingsUpdateSpec) Set(val *FarmCustomizationSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmCustomizationSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmCustomizationSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmCustomizationSettingsUpdateSpec(val *FarmCustomizationSettingsUpdateSpec) *NullableFarmCustomizationSettingsUpdateSpec {
	return &NullableFarmCustomizationSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableFarmCustomizationSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmCustomizationSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
