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

// DesktopPoolCustomizationSettingsUpdateSpec Applicable To: Automated Desktop pool. <br> Customization settings for Automated desktop pool.
type DesktopPoolCustomizationSettingsUpdateSpec struct {
	// This is required for instant clone desktop pools.<br>Instant Clone Engine Active Directory container for ClonePrep.
	AdContainerRdn                 *string                                              `json:"ad_container_rdn,omitempty"`
	CloneprepCustomizationSettings *DesktopPoolCloneprepCustomizationSettingsUpdateSpec `json:"cloneprep_customization_settings,omitempty"`
	// Type of customization to use. * NONE: Applicable To: Full clone desktop pools.<br>No customization. * SYS_PREP: Applicable To: Full clone desktop pools.<br>Microsoft Sysprep is a tool to deploy the configured operating system installation from a base image. The machine can then be customized based on an answer script. Sysprep can modify a larger number of configurable parameters than QuickPrep. * CLONE_PREP: Applicable To: Instant clone desktop pools.<br>ClonePrep is a VMware system tool executed by Instant Clone Engine during a instant clone machine deployment. ClonePrep personalizes each machine created from the Master image.
	CustomizationType string `json:"customization_type"`
	// Indicates whether to power on VMs after creation. This is the settings when customization will be done manually. This property is required if customization_type is set NONE.
	DoNotPowerOnVmsAfterCreation *bool `json:"do_not_power_on_vms_after_creation,omitempty"`
	// This is required for instant clone desktop pools.<br>This is the administrator which will add the machines to its domain upon creation.
	InstantCloneDomainAccountId *string `json:"instant_clone_domain_account_id,omitempty"`
	// Applicable To: Automated desktop pools.<br>Indicates whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names.
	ReusePreExistingAccounts *bool `json:"reuse_pre_existing_accounts,omitempty"`
	// This is required when customization_type is set as SYS_PREP.<br>Customization specification to use when Sysprep customization is requested.
	SysprepCustomizationSpecId *string `json:"sysprep_customization_spec_id,omitempty"`
}

// NewDesktopPoolCustomizationSettingsUpdateSpec instantiates a new DesktopPoolCustomizationSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolCustomizationSettingsUpdateSpec(customizationType string) *DesktopPoolCustomizationSettingsUpdateSpec {
	this := DesktopPoolCustomizationSettingsUpdateSpec{}
	this.CustomizationType = customizationType
	return &this
}

// NewDesktopPoolCustomizationSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolCustomizationSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolCustomizationSettingsUpdateSpecWithDefaults() *DesktopPoolCustomizationSettingsUpdateSpec {
	this := DesktopPoolCustomizationSettingsUpdateSpec{}
	return &this
}

// GetAdContainerRdn returns the AdContainerRdn field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetAdContainerRdn() string {
	if o == nil || o.AdContainerRdn == nil {
		var ret string
		return ret
	}
	return *o.AdContainerRdn
}

// GetAdContainerRdnOk returns a tuple with the AdContainerRdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetAdContainerRdnOk() (*string, bool) {
	if o == nil || o.AdContainerRdn == nil {
		return nil, false
	}
	return o.AdContainerRdn, true
}

// HasAdContainerRdn returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasAdContainerRdn() bool {
	if o != nil && o.AdContainerRdn != nil {
		return true
	}

	return false
}

// SetAdContainerRdn gets a reference to the given string and assigns it to the AdContainerRdn field.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetAdContainerRdn(v string) {
	o.AdContainerRdn = &v
}

// GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetCloneprepCustomizationSettings() DesktopPoolCloneprepCustomizationSettingsUpdateSpec {
	if o == nil || o.CloneprepCustomizationSettings == nil {
		var ret DesktopPoolCloneprepCustomizationSettingsUpdateSpec
		return ret
	}
	return *o.CloneprepCustomizationSettings
}

// GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetCloneprepCustomizationSettingsOk() (*DesktopPoolCloneprepCustomizationSettingsUpdateSpec, bool) {
	if o == nil || o.CloneprepCustomizationSettings == nil {
		return nil, false
	}
	return o.CloneprepCustomizationSettings, true
}

// HasCloneprepCustomizationSettings returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasCloneprepCustomizationSettings() bool {
	if o != nil && o.CloneprepCustomizationSettings != nil {
		return true
	}

	return false
}

// SetCloneprepCustomizationSettings gets a reference to the given DesktopPoolCloneprepCustomizationSettingsUpdateSpec and assigns it to the CloneprepCustomizationSettings field.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetCloneprepCustomizationSettings(v DesktopPoolCloneprepCustomizationSettingsUpdateSpec) {
	o.CloneprepCustomizationSettings = &v
}

// GetCustomizationType returns the CustomizationType field value
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetCustomizationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CustomizationType
}

// GetCustomizationTypeOk returns a tuple with the CustomizationType field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetCustomizationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomizationType, true
}

// SetCustomizationType sets field value
func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetCustomizationType(v string) {
	o.CustomizationType = v
}

// GetDoNotPowerOnVmsAfterCreation returns the DoNotPowerOnVmsAfterCreation field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetDoNotPowerOnVmsAfterCreation() bool {
	if o == nil || o.DoNotPowerOnVmsAfterCreation == nil {
		var ret bool
		return ret
	}
	return *o.DoNotPowerOnVmsAfterCreation
}

// GetDoNotPowerOnVmsAfterCreationOk returns a tuple with the DoNotPowerOnVmsAfterCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetDoNotPowerOnVmsAfterCreationOk() (*bool, bool) {
	if o == nil || o.DoNotPowerOnVmsAfterCreation == nil {
		return nil, false
	}
	return o.DoNotPowerOnVmsAfterCreation, true
}

// HasDoNotPowerOnVmsAfterCreation returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasDoNotPowerOnVmsAfterCreation() bool {
	if o != nil && o.DoNotPowerOnVmsAfterCreation != nil {
		return true
	}

	return false
}

// SetDoNotPowerOnVmsAfterCreation gets a reference to the given bool and assigns it to the DoNotPowerOnVmsAfterCreation field.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetDoNotPowerOnVmsAfterCreation(v bool) {
	o.DoNotPowerOnVmsAfterCreation = &v
}

// GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetInstantCloneDomainAccountId() string {
	if o == nil || o.InstantCloneDomainAccountId == nil {
		var ret string
		return ret
	}
	return *o.InstantCloneDomainAccountId
}

// GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetInstantCloneDomainAccountIdOk() (*string, bool) {
	if o == nil || o.InstantCloneDomainAccountId == nil {
		return nil, false
	}
	return o.InstantCloneDomainAccountId, true
}

// HasInstantCloneDomainAccountId returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasInstantCloneDomainAccountId() bool {
	if o != nil && o.InstantCloneDomainAccountId != nil {
		return true
	}

	return false
}

// SetInstantCloneDomainAccountId gets a reference to the given string and assigns it to the InstantCloneDomainAccountId field.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetInstantCloneDomainAccountId(v string) {
	o.InstantCloneDomainAccountId = &v
}

// GetReusePreExistingAccounts returns the ReusePreExistingAccounts field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetReusePreExistingAccounts() bool {
	if o == nil || o.ReusePreExistingAccounts == nil {
		var ret bool
		return ret
	}
	return *o.ReusePreExistingAccounts
}

// GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetReusePreExistingAccountsOk() (*bool, bool) {
	if o == nil || o.ReusePreExistingAccounts == nil {
		return nil, false
	}
	return o.ReusePreExistingAccounts, true
}

// HasReusePreExistingAccounts returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasReusePreExistingAccounts() bool {
	if o != nil && o.ReusePreExistingAccounts != nil {
		return true
	}

	return false
}

// SetReusePreExistingAccounts gets a reference to the given bool and assigns it to the ReusePreExistingAccounts field.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetReusePreExistingAccounts(v bool) {
	o.ReusePreExistingAccounts = &v
}

// GetSysprepCustomizationSpecId returns the SysprepCustomizationSpecId field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetSysprepCustomizationSpecId() string {
	if o == nil || o.SysprepCustomizationSpecId == nil {
		var ret string
		return ret
	}
	return *o.SysprepCustomizationSpecId
}

// GetSysprepCustomizationSpecIdOk returns a tuple with the SysprepCustomizationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) GetSysprepCustomizationSpecIdOk() (*string, bool) {
	if o == nil || o.SysprepCustomizationSpecId == nil {
		return nil, false
	}
	return o.SysprepCustomizationSpecId, true
}

// HasSysprepCustomizationSpecId returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) HasSysprepCustomizationSpecId() bool {
	if o != nil && o.SysprepCustomizationSpecId != nil {
		return true
	}

	return false
}

// SetSysprepCustomizationSpecId gets a reference to the given string and assigns it to the SysprepCustomizationSpecId field.
func (o *DesktopPoolCustomizationSettingsUpdateSpec) SetSysprepCustomizationSpecId(v string) {
	o.SysprepCustomizationSpecId = &v
}

func (o DesktopPoolCustomizationSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
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
	if o.DoNotPowerOnVmsAfterCreation != nil {
		toSerialize["do_not_power_on_vms_after_creation"] = o.DoNotPowerOnVmsAfterCreation
	}
	if o.InstantCloneDomainAccountId != nil {
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

type NullableDesktopPoolCustomizationSettingsUpdateSpec struct {
	value *DesktopPoolCustomizationSettingsUpdateSpec
	isSet bool
}

func (v NullableDesktopPoolCustomizationSettingsUpdateSpec) Get() *DesktopPoolCustomizationSettingsUpdateSpec {
	return v.value
}

func (v *NullableDesktopPoolCustomizationSettingsUpdateSpec) Set(val *DesktopPoolCustomizationSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolCustomizationSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolCustomizationSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolCustomizationSettingsUpdateSpec(val *DesktopPoolCustomizationSettingsUpdateSpec) *NullableDesktopPoolCustomizationSettingsUpdateSpec {
	return &NullableDesktopPoolCustomizationSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableDesktopPoolCustomizationSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolCustomizationSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}