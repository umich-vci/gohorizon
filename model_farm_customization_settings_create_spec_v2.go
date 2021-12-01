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

// FarmCustomizationSettingsCreateSpecV2 Customization settings for the automated farm.
type FarmCustomizationSettingsCreateSpecV2 struct {
	// Instant Clone Engine Active Directory container for clone prep. Default value is CN=Computers
	AdContainerRdn                 *string                                       `json:"ad_container_rdn,omitempty"`
	CloneprepCustomizationSettings *FarmCloneprepCustomizationSettingsCreateSpec `json:"cloneprep_customization_settings,omitempty"`
	// Type of customization to use. * SYS_PREP: Applicable To: Instant clone automated Farms.<br>Microsoft Sysprep is a tool to deploy the configured operating system installation from a base image. The machine can then be customized based on an answer script. * CLONE_PREP: Applicable To: Instant clone automated Farms.<br>ClonePrep is a VMware system tool executed by Instant Clone Engine during a instant clone machine deployment. ClonePrep personalizes each machine created from the Master image.
	CustomizationType string `json:"customization_type"`
	// Instant clone domain account. This is the administrator which will add the machines to its domain upon creation.
	InstantCloneDomainAccountId string `json:"instant_clone_domain_account_id"`
	// Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. Default value is false.
	ReusePreExistingAccounts *bool `json:"reuse_pre_existing_accounts,omitempty"`
	// Customization specification to use when Sysprep customization is requested. This is required if customization_type is set to SYS_PREP
	SysprepCustomizationSpecId *string `json:"sysprep_customization_spec_id,omitempty"`
}

// NewFarmCustomizationSettingsCreateSpecV2 instantiates a new FarmCustomizationSettingsCreateSpecV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmCustomizationSettingsCreateSpecV2(customizationType string, instantCloneDomainAccountId string) *FarmCustomizationSettingsCreateSpecV2 {
	this := FarmCustomizationSettingsCreateSpecV2{}
	this.CustomizationType = customizationType
	this.InstantCloneDomainAccountId = instantCloneDomainAccountId
	return &this
}

// NewFarmCustomizationSettingsCreateSpecV2WithDefaults instantiates a new FarmCustomizationSettingsCreateSpecV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmCustomizationSettingsCreateSpecV2WithDefaults() *FarmCustomizationSettingsCreateSpecV2 {
	this := FarmCustomizationSettingsCreateSpecV2{}
	return &this
}

// GetAdContainerRdn returns the AdContainerRdn field value if set, zero value otherwise.
func (o *FarmCustomizationSettingsCreateSpecV2) GetAdContainerRdn() string {
	if o == nil || o.AdContainerRdn == nil {
		var ret string
		return ret
	}
	return *o.AdContainerRdn
}

// GetAdContainerRdnOk returns a tuple with the AdContainerRdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) GetAdContainerRdnOk() (*string, bool) {
	if o == nil || o.AdContainerRdn == nil {
		return nil, false
	}
	return o.AdContainerRdn, true
}

// HasAdContainerRdn returns a boolean if a field has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) HasAdContainerRdn() bool {
	if o != nil && o.AdContainerRdn != nil {
		return true
	}

	return false
}

// SetAdContainerRdn gets a reference to the given string and assigns it to the AdContainerRdn field.
func (o *FarmCustomizationSettingsCreateSpecV2) SetAdContainerRdn(v string) {
	o.AdContainerRdn = &v
}

// GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field value if set, zero value otherwise.
func (o *FarmCustomizationSettingsCreateSpecV2) GetCloneprepCustomizationSettings() FarmCloneprepCustomizationSettingsCreateSpec {
	if o == nil || o.CloneprepCustomizationSettings == nil {
		var ret FarmCloneprepCustomizationSettingsCreateSpec
		return ret
	}
	return *o.CloneprepCustomizationSettings
}

// GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) GetCloneprepCustomizationSettingsOk() (*FarmCloneprepCustomizationSettingsCreateSpec, bool) {
	if o == nil || o.CloneprepCustomizationSettings == nil {
		return nil, false
	}
	return o.CloneprepCustomizationSettings, true
}

// HasCloneprepCustomizationSettings returns a boolean if a field has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) HasCloneprepCustomizationSettings() bool {
	if o != nil && o.CloneprepCustomizationSettings != nil {
		return true
	}

	return false
}

// SetCloneprepCustomizationSettings gets a reference to the given FarmCloneprepCustomizationSettingsCreateSpec and assigns it to the CloneprepCustomizationSettings field.
func (o *FarmCustomizationSettingsCreateSpecV2) SetCloneprepCustomizationSettings(v FarmCloneprepCustomizationSettingsCreateSpec) {
	o.CloneprepCustomizationSettings = &v
}

// GetCustomizationType returns the CustomizationType field value
func (o *FarmCustomizationSettingsCreateSpecV2) GetCustomizationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomizationType
}

// GetCustomizationTypeOk returns a tuple with the CustomizationType field value
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) GetCustomizationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomizationType, true
}

// SetCustomizationType sets field value
func (o *FarmCustomizationSettingsCreateSpecV2) SetCustomizationType(v string) {
	o.CustomizationType = v
}

// GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field value
func (o *FarmCustomizationSettingsCreateSpecV2) GetInstantCloneDomainAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstantCloneDomainAccountId
}

// GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field value
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) GetInstantCloneDomainAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InstantCloneDomainAccountId, true
}

// SetInstantCloneDomainAccountId sets field value
func (o *FarmCustomizationSettingsCreateSpecV2) SetInstantCloneDomainAccountId(v string) {
	o.InstantCloneDomainAccountId = v
}

// GetReusePreExistingAccounts returns the ReusePreExistingAccounts field value if set, zero value otherwise.
func (o *FarmCustomizationSettingsCreateSpecV2) GetReusePreExistingAccounts() bool {
	if o == nil || o.ReusePreExistingAccounts == nil {
		var ret bool
		return ret
	}
	return *o.ReusePreExistingAccounts
}

// GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) GetReusePreExistingAccountsOk() (*bool, bool) {
	if o == nil || o.ReusePreExistingAccounts == nil {
		return nil, false
	}
	return o.ReusePreExistingAccounts, true
}

// HasReusePreExistingAccounts returns a boolean if a field has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) HasReusePreExistingAccounts() bool {
	if o != nil && o.ReusePreExistingAccounts != nil {
		return true
	}

	return false
}

// SetReusePreExistingAccounts gets a reference to the given bool and assigns it to the ReusePreExistingAccounts field.
func (o *FarmCustomizationSettingsCreateSpecV2) SetReusePreExistingAccounts(v bool) {
	o.ReusePreExistingAccounts = &v
}

// GetSysprepCustomizationSpecId returns the SysprepCustomizationSpecId field value if set, zero value otherwise.
func (o *FarmCustomizationSettingsCreateSpecV2) GetSysprepCustomizationSpecId() string {
	if o == nil || o.SysprepCustomizationSpecId == nil {
		var ret string
		return ret
	}
	return *o.SysprepCustomizationSpecId
}

// GetSysprepCustomizationSpecIdOk returns a tuple with the SysprepCustomizationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) GetSysprepCustomizationSpecIdOk() (*string, bool) {
	if o == nil || o.SysprepCustomizationSpecId == nil {
		return nil, false
	}
	return o.SysprepCustomizationSpecId, true
}

// HasSysprepCustomizationSpecId returns a boolean if a field has been set.
func (o *FarmCustomizationSettingsCreateSpecV2) HasSysprepCustomizationSpecId() bool {
	if o != nil && o.SysprepCustomizationSpecId != nil {
		return true
	}

	return false
}

// SetSysprepCustomizationSpecId gets a reference to the given string and assigns it to the SysprepCustomizationSpecId field.
func (o *FarmCustomizationSettingsCreateSpecV2) SetSysprepCustomizationSpecId(v string) {
	o.SysprepCustomizationSpecId = &v
}

func (o FarmCustomizationSettingsCreateSpecV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdContainerRdn != nil {
		toSerialize["ad_container_rdn"] = o.AdContainerRdn
	}
	if o.CloneprepCustomizationSettings != nil {
		toSerialize["cloneprep_customization_settings"] = o.CloneprepCustomizationSettings
	}
	if true {
		toSerialize["customization_type"] = o.CustomizationType
	}
	if true {
		toSerialize["instant_clone_domain_account_id"] = o.InstantCloneDomainAccountId
	}
	if o.ReusePreExistingAccounts != nil {
		toSerialize["reuse_pre_existing_accounts"] = o.ReusePreExistingAccounts
	}
	if o.SysprepCustomizationSpecId != nil {
		toSerialize["sysprep_customization_spec_id"] = o.SysprepCustomizationSpecId
	}
	return json.Marshal(toSerialize)
}

type NullableFarmCustomizationSettingsCreateSpecV2 struct {
	value *FarmCustomizationSettingsCreateSpecV2
	isSet bool
}

func (v NullableFarmCustomizationSettingsCreateSpecV2) Get() *FarmCustomizationSettingsCreateSpecV2 {
	return v.value
}

func (v *NullableFarmCustomizationSettingsCreateSpecV2) Set(val *FarmCustomizationSettingsCreateSpecV2) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmCustomizationSettingsCreateSpecV2) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmCustomizationSettingsCreateSpecV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmCustomizationSettingsCreateSpecV2(val *FarmCustomizationSettingsCreateSpecV2) *NullableFarmCustomizationSettingsCreateSpecV2 {
	return &NullableFarmCustomizationSettingsCreateSpecV2{value: val, isSet: true}
}

func (v NullableFarmCustomizationSettingsCreateSpecV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmCustomizationSettingsCreateSpecV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}