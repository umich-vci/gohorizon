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

// DesktopPoolCloneprepCustomizationSettings Settings for ClonePrep customization. This setting is only applicable to instant clone desktop pools.
type DesktopPoolCloneprepCustomizationSettings struct {
	// Post synchronization script. ClonePrep can run a customization script on instant-clone machines after they are created or recovered or a new image is pushed. Provide the path to the script on the parent virtual machine.
	PostSynchronizationScriptName *string `json:"post_synchronization_script_name,omitempty"`
	// Post synchronization script parameters.
	PostSynchronizationScriptParameters *string `json:"post_synchronization_script_parameters,omitempty"`
	// Power off script. ClonePrep can run a customization script on instant-clone machines before they are powered off. Provide the path to the script on the parent virtual machine.
	PowerOffScriptName *string `json:"power_off_script_name,omitempty"`
	// Power off script parameters.
	PowerOffScriptParameters *string `json:"power_off_script_parameters,omitempty"`
	// Instant Clone publishing needs an additional computer account in the same AD domain as the clones. This field accepts the pre-created computer accounts.
	PrimingComputerAccount *string `json:"priming_computer_account,omitempty"`
}

// NewDesktopPoolCloneprepCustomizationSettings instantiates a new DesktopPoolCloneprepCustomizationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolCloneprepCustomizationSettings() *DesktopPoolCloneprepCustomizationSettings {
	this := DesktopPoolCloneprepCustomizationSettings{}
	return &this
}

// NewDesktopPoolCloneprepCustomizationSettingsWithDefaults instantiates a new DesktopPoolCloneprepCustomizationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolCloneprepCustomizationSettingsWithDefaults() *DesktopPoolCloneprepCustomizationSettings {
	this := DesktopPoolCloneprepCustomizationSettings{}
	return &this
}

// GetPostSynchronizationScriptName returns the PostSynchronizationScriptName field value if set, zero value otherwise.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPostSynchronizationScriptName() string {
	if o == nil || o.PostSynchronizationScriptName == nil {
		var ret string
		return ret
	}
	return *o.PostSynchronizationScriptName
}

// GetPostSynchronizationScriptNameOk returns a tuple with the PostSynchronizationScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPostSynchronizationScriptNameOk() (*string, bool) {
	if o == nil || o.PostSynchronizationScriptName == nil {
		return nil, false
	}
	return o.PostSynchronizationScriptName, true
}

// HasPostSynchronizationScriptName returns a boolean if a field has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) HasPostSynchronizationScriptName() bool {
	if o != nil && o.PostSynchronizationScriptName != nil {
		return true
	}

	return false
}

// SetPostSynchronizationScriptName gets a reference to the given string and assigns it to the PostSynchronizationScriptName field.
func (o *DesktopPoolCloneprepCustomizationSettings) SetPostSynchronizationScriptName(v string) {
	o.PostSynchronizationScriptName = &v
}

// GetPostSynchronizationScriptParameters returns the PostSynchronizationScriptParameters field value if set, zero value otherwise.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPostSynchronizationScriptParameters() string {
	if o == nil || o.PostSynchronizationScriptParameters == nil {
		var ret string
		return ret
	}
	return *o.PostSynchronizationScriptParameters
}

// GetPostSynchronizationScriptParametersOk returns a tuple with the PostSynchronizationScriptParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPostSynchronizationScriptParametersOk() (*string, bool) {
	if o == nil || o.PostSynchronizationScriptParameters == nil {
		return nil, false
	}
	return o.PostSynchronizationScriptParameters, true
}

// HasPostSynchronizationScriptParameters returns a boolean if a field has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) HasPostSynchronizationScriptParameters() bool {
	if o != nil && o.PostSynchronizationScriptParameters != nil {
		return true
	}

	return false
}

// SetPostSynchronizationScriptParameters gets a reference to the given string and assigns it to the PostSynchronizationScriptParameters field.
func (o *DesktopPoolCloneprepCustomizationSettings) SetPostSynchronizationScriptParameters(v string) {
	o.PostSynchronizationScriptParameters = &v
}

// GetPowerOffScriptName returns the PowerOffScriptName field value if set, zero value otherwise.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPowerOffScriptName() string {
	if o == nil || o.PowerOffScriptName == nil {
		var ret string
		return ret
	}
	return *o.PowerOffScriptName
}

// GetPowerOffScriptNameOk returns a tuple with the PowerOffScriptName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPowerOffScriptNameOk() (*string, bool) {
	if o == nil || o.PowerOffScriptName == nil {
		return nil, false
	}
	return o.PowerOffScriptName, true
}

// HasPowerOffScriptName returns a boolean if a field has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) HasPowerOffScriptName() bool {
	if o != nil && o.PowerOffScriptName != nil {
		return true
	}

	return false
}

// SetPowerOffScriptName gets a reference to the given string and assigns it to the PowerOffScriptName field.
func (o *DesktopPoolCloneprepCustomizationSettings) SetPowerOffScriptName(v string) {
	o.PowerOffScriptName = &v
}

// GetPowerOffScriptParameters returns the PowerOffScriptParameters field value if set, zero value otherwise.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPowerOffScriptParameters() string {
	if o == nil || o.PowerOffScriptParameters == nil {
		var ret string
		return ret
	}
	return *o.PowerOffScriptParameters
}

// GetPowerOffScriptParametersOk returns a tuple with the PowerOffScriptParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPowerOffScriptParametersOk() (*string, bool) {
	if o == nil || o.PowerOffScriptParameters == nil {
		return nil, false
	}
	return o.PowerOffScriptParameters, true
}

// HasPowerOffScriptParameters returns a boolean if a field has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) HasPowerOffScriptParameters() bool {
	if o != nil && o.PowerOffScriptParameters != nil {
		return true
	}

	return false
}

// SetPowerOffScriptParameters gets a reference to the given string and assigns it to the PowerOffScriptParameters field.
func (o *DesktopPoolCloneprepCustomizationSettings) SetPowerOffScriptParameters(v string) {
	o.PowerOffScriptParameters = &v
}

// GetPrimingComputerAccount returns the PrimingComputerAccount field value if set, zero value otherwise.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPrimingComputerAccount() string {
	if o == nil || o.PrimingComputerAccount == nil {
		var ret string
		return ret
	}
	return *o.PrimingComputerAccount
}

// GetPrimingComputerAccountOk returns a tuple with the PrimingComputerAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) GetPrimingComputerAccountOk() (*string, bool) {
	if o == nil || o.PrimingComputerAccount == nil {
		return nil, false
	}
	return o.PrimingComputerAccount, true
}

// HasPrimingComputerAccount returns a boolean if a field has been set.
func (o *DesktopPoolCloneprepCustomizationSettings) HasPrimingComputerAccount() bool {
	if o != nil && o.PrimingComputerAccount != nil {
		return true
	}

	return false
}

// SetPrimingComputerAccount gets a reference to the given string and assigns it to the PrimingComputerAccount field.
func (o *DesktopPoolCloneprepCustomizationSettings) SetPrimingComputerAccount(v string) {
	o.PrimingComputerAccount = &v
}

func (o DesktopPoolCloneprepCustomizationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PostSynchronizationScriptName != nil {
		toSerialize["post_synchronization_script_name"] = o.PostSynchronizationScriptName
	}
	if o.PostSynchronizationScriptParameters != nil {
		toSerialize["post_synchronization_script_parameters"] = o.PostSynchronizationScriptParameters
	}
	if o.PowerOffScriptName != nil {
		toSerialize["power_off_script_name"] = o.PowerOffScriptName
	}
	if o.PowerOffScriptParameters != nil {
		toSerialize["power_off_script_parameters"] = o.PowerOffScriptParameters
	}
	if o.PrimingComputerAccount != nil {
		toSerialize["priming_computer_account"] = o.PrimingComputerAccount
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolCloneprepCustomizationSettings struct {
	value *DesktopPoolCloneprepCustomizationSettings
	isSet bool
}

func (v NullableDesktopPoolCloneprepCustomizationSettings) Get() *DesktopPoolCloneprepCustomizationSettings {
	return v.value
}

func (v *NullableDesktopPoolCloneprepCustomizationSettings) Set(val *DesktopPoolCloneprepCustomizationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolCloneprepCustomizationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolCloneprepCustomizationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolCloneprepCustomizationSettings(val *DesktopPoolCloneprepCustomizationSettings) *NullableDesktopPoolCloneprepCustomizationSettings {
	return &NullableDesktopPoolCloneprepCustomizationSettings{value: val, isSet: true}
}

func (v NullableDesktopPoolCloneprepCustomizationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolCloneprepCustomizationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
