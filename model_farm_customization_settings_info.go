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

// FarmCustomizationSettingsInfo Customization settings for the automated farm.
type FarmCustomizationSettingsInfo struct {
	// Instant Clone Engine Active Directory container for clone prep.
	AdContainerRdn                 *string                                 `json:"ad_container_rdn,omitempty"`
	CloneprepCustomizationSettings *FarmCloneprepCustomizationSettingsInfo `json:"cloneprep_customization_settings,omitempty"`
	// Instant clone domain account. This is the administrator which will add the machines to its domain upon creation.
	InstantCloneDomainAccountId *string `json:"instant_clone_domain_account_id,omitempty"`
	// Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names.
	ReusePreExistingAccounts *bool `json:"reuse_pre_existing_accounts,omitempty"`
}

// NewFarmCustomizationSettingsInfo instantiates a new FarmCustomizationSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmCustomizationSettingsInfo() *FarmCustomizationSettingsInfo {
	this := FarmCustomizationSettingsInfo{}
	return &this
}

// NewFarmCustomizationSettingsInfoWithDefaults instantiates a new FarmCustomizationSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmCustomizationSettingsInfoWithDefaults() *FarmCustomizationSettingsInfo {
	this := FarmCustomizationSettingsInfo{}
	return &this
}

// GetAdContainerRdn returns the AdContainerRdn field value if set, zero value otherwise.
func (o *FarmCustomizationSettingsInfo) GetAdContainerRdn() string {
	if o == nil || o.AdContainerRdn == nil {
		var ret string
		return ret
	}
	return *o.AdContainerRdn
}

// GetAdContainerRdnOk returns a tuple with the AdContainerRdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsInfo) GetAdContainerRdnOk() (*string, bool) {
	if o == nil || o.AdContainerRdn == nil {
		return nil, false
	}
	return o.AdContainerRdn, true
}

// HasAdContainerRdn returns a boolean if a field has been set.
func (o *FarmCustomizationSettingsInfo) HasAdContainerRdn() bool {
	if o != nil && o.AdContainerRdn != nil {
		return true
	}

	return false
}

// SetAdContainerRdn gets a reference to the given string and assigns it to the AdContainerRdn field.
func (o *FarmCustomizationSettingsInfo) SetAdContainerRdn(v string) {
	o.AdContainerRdn = &v
}

// GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field value if set, zero value otherwise.
func (o *FarmCustomizationSettingsInfo) GetCloneprepCustomizationSettings() FarmCloneprepCustomizationSettingsInfo {
	if o == nil || o.CloneprepCustomizationSettings == nil {
		var ret FarmCloneprepCustomizationSettingsInfo
		return ret
	}
	return *o.CloneprepCustomizationSettings
}

// GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsInfo) GetCloneprepCustomizationSettingsOk() (*FarmCloneprepCustomizationSettingsInfo, bool) {
	if o == nil || o.CloneprepCustomizationSettings == nil {
		return nil, false
	}
	return o.CloneprepCustomizationSettings, true
}

// HasCloneprepCustomizationSettings returns a boolean if a field has been set.
func (o *FarmCustomizationSettingsInfo) HasCloneprepCustomizationSettings() bool {
	if o != nil && o.CloneprepCustomizationSettings != nil {
		return true
	}

	return false
}

// SetCloneprepCustomizationSettings gets a reference to the given FarmCloneprepCustomizationSettingsInfo and assigns it to the CloneprepCustomizationSettings field.
func (o *FarmCustomizationSettingsInfo) SetCloneprepCustomizationSettings(v FarmCloneprepCustomizationSettingsInfo) {
	o.CloneprepCustomizationSettings = &v
}

// GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field value if set, zero value otherwise.
func (o *FarmCustomizationSettingsInfo) GetInstantCloneDomainAccountId() string {
	if o == nil || o.InstantCloneDomainAccountId == nil {
		var ret string
		return ret
	}
	return *o.InstantCloneDomainAccountId
}

// GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsInfo) GetInstantCloneDomainAccountIdOk() (*string, bool) {
	if o == nil || o.InstantCloneDomainAccountId == nil {
		return nil, false
	}
	return o.InstantCloneDomainAccountId, true
}

// HasInstantCloneDomainAccountId returns a boolean if a field has been set.
func (o *FarmCustomizationSettingsInfo) HasInstantCloneDomainAccountId() bool {
	if o != nil && o.InstantCloneDomainAccountId != nil {
		return true
	}

	return false
}

// SetInstantCloneDomainAccountId gets a reference to the given string and assigns it to the InstantCloneDomainAccountId field.
func (o *FarmCustomizationSettingsInfo) SetInstantCloneDomainAccountId(v string) {
	o.InstantCloneDomainAccountId = &v
}

// GetReusePreExistingAccounts returns the ReusePreExistingAccounts field value if set, zero value otherwise.
func (o *FarmCustomizationSettingsInfo) GetReusePreExistingAccounts() bool {
	if o == nil || o.ReusePreExistingAccounts == nil {
		var ret bool
		return ret
	}
	return *o.ReusePreExistingAccounts
}

// GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsInfo) GetReusePreExistingAccountsOk() (*bool, bool) {
	if o == nil || o.ReusePreExistingAccounts == nil {
		return nil, false
	}
	return o.ReusePreExistingAccounts, true
}

// HasReusePreExistingAccounts returns a boolean if a field has been set.
func (o *FarmCustomizationSettingsInfo) HasReusePreExistingAccounts() bool {
	if o != nil && o.ReusePreExistingAccounts != nil {
		return true
	}

	return false
}

// SetReusePreExistingAccounts gets a reference to the given bool and assigns it to the ReusePreExistingAccounts field.
func (o *FarmCustomizationSettingsInfo) SetReusePreExistingAccounts(v bool) {
	o.ReusePreExistingAccounts = &v
}

func (o FarmCustomizationSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdContainerRdn != nil {
		toSerialize["ad_container_rdn"] = o.AdContainerRdn
	}
	if o.CloneprepCustomizationSettings != nil {
		toSerialize["cloneprep_customization_settings"] = o.CloneprepCustomizationSettings
	}
	if o.InstantCloneDomainAccountId != nil {
		toSerialize["instant_clone_domain_account_id"] = o.InstantCloneDomainAccountId
	}
	if o.ReusePreExistingAccounts != nil {
		toSerialize["reuse_pre_existing_accounts"] = o.ReusePreExistingAccounts
	}
	return json.Marshal(toSerialize)
}

type NullableFarmCustomizationSettingsInfo struct {
	value *FarmCustomizationSettingsInfo
	isSet bool
}

func (v NullableFarmCustomizationSettingsInfo) Get() *FarmCustomizationSettingsInfo {
	return v.value
}

func (v *NullableFarmCustomizationSettingsInfo) Set(val *FarmCustomizationSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmCustomizationSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmCustomizationSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmCustomizationSettingsInfo(val *FarmCustomizationSettingsInfo) *NullableFarmCustomizationSettingsInfo {
	return &NullableFarmCustomizationSettingsInfo{value: val, isSet: true}
}

func (v NullableFarmCustomizationSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmCustomizationSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
