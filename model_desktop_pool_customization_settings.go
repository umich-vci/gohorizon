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

// DesktopPoolCustomizationSettings Customization settings.
type DesktopPoolCustomizationSettings struct {
	// Applicable To: Linked/instant clone automated desktop pools.<br>View Composer and Instant Clone Engine Active Directory container for QuickPrep and ClonePrep.
	AdContainerRdn                 *string                                    `json:"ad_container_rdn,omitempty"`
	CloneprepCustomizationSettings *DesktopPoolCloneprepCustomizationSettings `json:"cloneprep_customization_settings,omitempty"`
	// Type of customization to use. * NONE: Applicable To: Full clone automated desktop pools.<br>No customization. * QUICK_PREP: Applicable To: Linked clone automated desktop pools.<br>QuickPrep is a VMware system tool executed by View Composer during a linked-clone machine deployment. QuickPrep personalizes each machine created from the Master image. * SYS_PREP: Applicable To: Full clone and Linked clone automated desktop pools.<br>Microsoft Sysprep is a tool to deploy the configured operating system installation from a base image. The machine can then be customized based on an answer script. Sysprep can modify a larger number of configurable parameters than QuickPrep. * CLONE_PREP: Applicable To: Instant clone automated desktop pools.<br>ClonePrep is a VMware system tool executed by Instant Clone Engine during a instant clone machine deployment. ClonePrep personalizes each machine created from the Master image.
	CustomizationType *string `json:"customization_type,omitempty"`
	// Whether to power on VMs after creation. This is the settings when customization will be done manually.
	DoNotPowerOnVmsAfterCreation *bool `json:"do_not_power_on_vms_after_creation,omitempty"`
	// Applicable To: Instant clone automated desktop pools.<br>Instant clone domain account. This is the administrator which will add the machines to its domain upon creation.
	InstantCloneDomainAccountId    *string                                    `json:"instant_clone_domain_account_id,omitempty"`
	QuickprepCustomizationSettings *DesktopPoolQuickprepCustomizationSettings `json:"quickprep_customization_settings,omitempty"`
	// Applicable To: Manual and automated desktop pools.<br>Whether to allow the use of existing AD computer accounts when the VM names of newly created clones match the existing computer account names. This is applicable only for automated desktop pools.
	ReusePreExistingAccounts *bool `json:"reuse_pre_existing_accounts,omitempty"`
	// Customization specification to use when Sysprep customization is requested.
	SysprepCustomizationSpecId *string `json:"sysprep_customization_spec_id,omitempty"`
	// Applicable To: Linked clone automated desktop pools.<br>View Composer domain account. This is the administrator which will add the machines to its domain upon creation. This must be set for linked-clone automated desktop pools.
	ViewComposerDomainAccountId *string `json:"view_composer_domain_account_id,omitempty"`
}

// NewDesktopPoolCustomizationSettings instantiates a new DesktopPoolCustomizationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolCustomizationSettings() *DesktopPoolCustomizationSettings {
	this := DesktopPoolCustomizationSettings{}
	return &this
}

// NewDesktopPoolCustomizationSettingsWithDefaults instantiates a new DesktopPoolCustomizationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolCustomizationSettingsWithDefaults() *DesktopPoolCustomizationSettings {
	this := DesktopPoolCustomizationSettings{}
	return &this
}

// GetAdContainerRdn returns the AdContainerRdn field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetAdContainerRdn() string {
	if o == nil || o.AdContainerRdn == nil {
		var ret string
		return ret
	}
	return *o.AdContainerRdn
}

// GetAdContainerRdnOk returns a tuple with the AdContainerRdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetAdContainerRdnOk() (*string, bool) {
	if o == nil || o.AdContainerRdn == nil {
		return nil, false
	}
	return o.AdContainerRdn, true
}

// HasAdContainerRdn returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasAdContainerRdn() bool {
	if o != nil && o.AdContainerRdn != nil {
		return true
	}

	return false
}

// SetAdContainerRdn gets a reference to the given string and assigns it to the AdContainerRdn field.
func (o *DesktopPoolCustomizationSettings) SetAdContainerRdn(v string) {
	o.AdContainerRdn = &v
}

// GetCloneprepCustomizationSettings returns the CloneprepCustomizationSettings field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetCloneprepCustomizationSettings() DesktopPoolCloneprepCustomizationSettings {
	if o == nil || o.CloneprepCustomizationSettings == nil {
		var ret DesktopPoolCloneprepCustomizationSettings
		return ret
	}
	return *o.CloneprepCustomizationSettings
}

// GetCloneprepCustomizationSettingsOk returns a tuple with the CloneprepCustomizationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetCloneprepCustomizationSettingsOk() (*DesktopPoolCloneprepCustomizationSettings, bool) {
	if o == nil || o.CloneprepCustomizationSettings == nil {
		return nil, false
	}
	return o.CloneprepCustomizationSettings, true
}

// HasCloneprepCustomizationSettings returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasCloneprepCustomizationSettings() bool {
	if o != nil && o.CloneprepCustomizationSettings != nil {
		return true
	}

	return false
}

// SetCloneprepCustomizationSettings gets a reference to the given DesktopPoolCloneprepCustomizationSettings and assigns it to the CloneprepCustomizationSettings field.
func (o *DesktopPoolCustomizationSettings) SetCloneprepCustomizationSettings(v DesktopPoolCloneprepCustomizationSettings) {
	o.CloneprepCustomizationSettings = &v
}

// GetCustomizationType returns the CustomizationType field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetCustomizationType() string {
	if o == nil || o.CustomizationType == nil {
		var ret string
		return ret
	}
	return *o.CustomizationType
}

// GetCustomizationTypeOk returns a tuple with the CustomizationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetCustomizationTypeOk() (*string, bool) {
	if o == nil || o.CustomizationType == nil {
		return nil, false
	}
	return o.CustomizationType, true
}

// HasCustomizationType returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasCustomizationType() bool {
	if o != nil && o.CustomizationType != nil {
		return true
	}

	return false
}

// SetCustomizationType gets a reference to the given string and assigns it to the CustomizationType field.
func (o *DesktopPoolCustomizationSettings) SetCustomizationType(v string) {
	o.CustomizationType = &v
}

// GetDoNotPowerOnVmsAfterCreation returns the DoNotPowerOnVmsAfterCreation field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetDoNotPowerOnVmsAfterCreation() bool {
	if o == nil || o.DoNotPowerOnVmsAfterCreation == nil {
		var ret bool
		return ret
	}
	return *o.DoNotPowerOnVmsAfterCreation
}

// GetDoNotPowerOnVmsAfterCreationOk returns a tuple with the DoNotPowerOnVmsAfterCreation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetDoNotPowerOnVmsAfterCreationOk() (*bool, bool) {
	if o == nil || o.DoNotPowerOnVmsAfterCreation == nil {
		return nil, false
	}
	return o.DoNotPowerOnVmsAfterCreation, true
}

// HasDoNotPowerOnVmsAfterCreation returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasDoNotPowerOnVmsAfterCreation() bool {
	if o != nil && o.DoNotPowerOnVmsAfterCreation != nil {
		return true
	}

	return false
}

// SetDoNotPowerOnVmsAfterCreation gets a reference to the given bool and assigns it to the DoNotPowerOnVmsAfterCreation field.
func (o *DesktopPoolCustomizationSettings) SetDoNotPowerOnVmsAfterCreation(v bool) {
	o.DoNotPowerOnVmsAfterCreation = &v
}

// GetInstantCloneDomainAccountId returns the InstantCloneDomainAccountId field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetInstantCloneDomainAccountId() string {
	if o == nil || o.InstantCloneDomainAccountId == nil {
		var ret string
		return ret
	}
	return *o.InstantCloneDomainAccountId
}

// GetInstantCloneDomainAccountIdOk returns a tuple with the InstantCloneDomainAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetInstantCloneDomainAccountIdOk() (*string, bool) {
	if o == nil || o.InstantCloneDomainAccountId == nil {
		return nil, false
	}
	return o.InstantCloneDomainAccountId, true
}

// HasInstantCloneDomainAccountId returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasInstantCloneDomainAccountId() bool {
	if o != nil && o.InstantCloneDomainAccountId != nil {
		return true
	}

	return false
}

// SetInstantCloneDomainAccountId gets a reference to the given string and assigns it to the InstantCloneDomainAccountId field.
func (o *DesktopPoolCustomizationSettings) SetInstantCloneDomainAccountId(v string) {
	o.InstantCloneDomainAccountId = &v
}

// GetQuickprepCustomizationSettings returns the QuickprepCustomizationSettings field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetQuickprepCustomizationSettings() DesktopPoolQuickprepCustomizationSettings {
	if o == nil || o.QuickprepCustomizationSettings == nil {
		var ret DesktopPoolQuickprepCustomizationSettings
		return ret
	}
	return *o.QuickprepCustomizationSettings
}

// GetQuickprepCustomizationSettingsOk returns a tuple with the QuickprepCustomizationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetQuickprepCustomizationSettingsOk() (*DesktopPoolQuickprepCustomizationSettings, bool) {
	if o == nil || o.QuickprepCustomizationSettings == nil {
		return nil, false
	}
	return o.QuickprepCustomizationSettings, true
}

// HasQuickprepCustomizationSettings returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasQuickprepCustomizationSettings() bool {
	if o != nil && o.QuickprepCustomizationSettings != nil {
		return true
	}

	return false
}

// SetQuickprepCustomizationSettings gets a reference to the given DesktopPoolQuickprepCustomizationSettings and assigns it to the QuickprepCustomizationSettings field.
func (o *DesktopPoolCustomizationSettings) SetQuickprepCustomizationSettings(v DesktopPoolQuickprepCustomizationSettings) {
	o.QuickprepCustomizationSettings = &v
}

// GetReusePreExistingAccounts returns the ReusePreExistingAccounts field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetReusePreExistingAccounts() bool {
	if o == nil || o.ReusePreExistingAccounts == nil {
		var ret bool
		return ret
	}
	return *o.ReusePreExistingAccounts
}

// GetReusePreExistingAccountsOk returns a tuple with the ReusePreExistingAccounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetReusePreExistingAccountsOk() (*bool, bool) {
	if o == nil || o.ReusePreExistingAccounts == nil {
		return nil, false
	}
	return o.ReusePreExistingAccounts, true
}

// HasReusePreExistingAccounts returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasReusePreExistingAccounts() bool {
	if o != nil && o.ReusePreExistingAccounts != nil {
		return true
	}

	return false
}

// SetReusePreExistingAccounts gets a reference to the given bool and assigns it to the ReusePreExistingAccounts field.
func (o *DesktopPoolCustomizationSettings) SetReusePreExistingAccounts(v bool) {
	o.ReusePreExistingAccounts = &v
}

// GetSysprepCustomizationSpecId returns the SysprepCustomizationSpecId field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetSysprepCustomizationSpecId() string {
	if o == nil || o.SysprepCustomizationSpecId == nil {
		var ret string
		return ret
	}
	return *o.SysprepCustomizationSpecId
}

// GetSysprepCustomizationSpecIdOk returns a tuple with the SysprepCustomizationSpecId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetSysprepCustomizationSpecIdOk() (*string, bool) {
	if o == nil || o.SysprepCustomizationSpecId == nil {
		return nil, false
	}
	return o.SysprepCustomizationSpecId, true
}

// HasSysprepCustomizationSpecId returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasSysprepCustomizationSpecId() bool {
	if o != nil && o.SysprepCustomizationSpecId != nil {
		return true
	}

	return false
}

// SetSysprepCustomizationSpecId gets a reference to the given string and assigns it to the SysprepCustomizationSpecId field.
func (o *DesktopPoolCustomizationSettings) SetSysprepCustomizationSpecId(v string) {
	o.SysprepCustomizationSpecId = &v
}

// GetViewComposerDomainAccountId returns the ViewComposerDomainAccountId field value if set, zero value otherwise.
func (o *DesktopPoolCustomizationSettings) GetViewComposerDomainAccountId() string {
	if o == nil || o.ViewComposerDomainAccountId == nil {
		var ret string
		return ret
	}
	return *o.ViewComposerDomainAccountId
}

// GetViewComposerDomainAccountIdOk returns a tuple with the ViewComposerDomainAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCustomizationSettings) GetViewComposerDomainAccountIdOk() (*string, bool) {
	if o == nil || o.ViewComposerDomainAccountId == nil {
		return nil, false
	}
	return o.ViewComposerDomainAccountId, true
}

// HasViewComposerDomainAccountId returns a boolean if a field has been set.
func (o *DesktopPoolCustomizationSettings) HasViewComposerDomainAccountId() bool {
	if o != nil && o.ViewComposerDomainAccountId != nil {
		return true
	}

	return false
}

// SetViewComposerDomainAccountId gets a reference to the given string and assigns it to the ViewComposerDomainAccountId field.
func (o *DesktopPoolCustomizationSettings) SetViewComposerDomainAccountId(v string) {
	o.ViewComposerDomainAccountId = &v
}

func (o DesktopPoolCustomizationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdContainerRdn != nil {
		toSerialize["ad_container_rdn"] = o.AdContainerRdn
	}
	if o.CloneprepCustomizationSettings != nil {
		toSerialize["cloneprep_customization_settings"] = o.CloneprepCustomizationSettings
	}
	if o.CustomizationType != nil {
		toSerialize["customization_type"] = o.CustomizationType
	}
	if o.DoNotPowerOnVmsAfterCreation != nil {
		toSerialize["do_not_power_on_vms_after_creation"] = o.DoNotPowerOnVmsAfterCreation
	}
	if o.InstantCloneDomainAccountId != nil {
		toSerialize["instant_clone_domain_account_id"] = o.InstantCloneDomainAccountId
	}
	if o.QuickprepCustomizationSettings != nil {
		toSerialize["quickprep_customization_settings"] = o.QuickprepCustomizationSettings
	}
	if o.ReusePreExistingAccounts != nil {
		toSerialize["reuse_pre_existing_accounts"] = o.ReusePreExistingAccounts
	}
	if o.SysprepCustomizationSpecId != nil {
		toSerialize["sysprep_customization_spec_id"] = o.SysprepCustomizationSpecId
	}
	if o.ViewComposerDomainAccountId != nil {
		toSerialize["view_composer_domain_account_id"] = o.ViewComposerDomainAccountId
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolCustomizationSettings struct {
	value *DesktopPoolCustomizationSettings
	isSet bool
}

func (v NullableDesktopPoolCustomizationSettings) Get() *DesktopPoolCustomizationSettings {
	return v.value
}

func (v *NullableDesktopPoolCustomizationSettings) Set(val *DesktopPoolCustomizationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolCustomizationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolCustomizationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolCustomizationSettings(val *DesktopPoolCustomizationSettings) *NullableDesktopPoolCustomizationSettings {
	return &NullableDesktopPoolCustomizationSettings{value: val, isSet: true}
}

func (v NullableDesktopPoolCustomizationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolCustomizationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
