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

// DesktopPoolInfo Information related to Desktop Pool.
type DesktopPoolInfo struct {
	// Description of the Desktop Pool. The maximum length is 1024 characters.
	Description *string `json:"description,omitempty"`
	// Display name of the Desktop Pool. The maximum length is 256 characters.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether the Desktop Pool is enabled for brokering.
	Enabled bool `json:"enabled"`
	// Unique ID representing Desktop Pool.
	Id string `json:"id"`
	// Name of the Desktop Pool. The maximum length is 64 characters.
	Name     string              `json:"name"`
	Settings DesktopPoolSettings `json:"settings"`
	// Source of the Machines in this Desktop Pool. * INSTANT_CLONE: The Desktop Pool uses instant clone technology for provisioning the machines.Applicable for AUTOMATED type desktop pools. * LINKED_CLONE: The Desktop Pool uses linked clone technology for provisioning the machines.Applicable for AUTOMATED type desktop pools. * VIRTUAL_CENTER: The Desktop Pool uses Virtual Center as source for provisioning the machines.Applicable for AUTOMATED and MANUAL type desktop pools. * RDS: The Desktop Pool is backed by Farm. The Farm used in this Desktop Pool can be of any Source. * UNMANAGED: The Desktop Pool holds the non-vCenter source machines that includes physical computers,blade PCs and non-vCenter servers. Applicable for MANUAL type desktop pools.
	Source string `json:"source"`
	// Type of the Desktop Pool. * AUTOMATED: Automated Desktop Pool. * MANUAL: Manual Desktop Pool. * RDS: RDS Desktop Pool.
	Type string `json:"type"`
}

// NewDesktopPoolInfo instantiates a new DesktopPoolInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolInfo(enabled bool, id string, name string, settings DesktopPoolSettings, source string, type_ string) *DesktopPoolInfo {
	this := DesktopPoolInfo{}
	this.Enabled = enabled
	this.Id = id
	this.Name = name
	this.Settings = settings
	this.Source = source
	this.Type = type_
	return &this
}

// NewDesktopPoolInfoWithDefaults instantiates a new DesktopPoolInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolInfoWithDefaults() *DesktopPoolInfo {
	this := DesktopPoolInfo{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *DesktopPoolInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *DesktopPoolInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *DesktopPoolInfo) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *DesktopPoolInfo) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolInfo) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *DesktopPoolInfo) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *DesktopPoolInfo) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value
func (o *DesktopPoolInfo) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolInfo) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *DesktopPoolInfo) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value
func (o *DesktopPoolInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DesktopPoolInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DesktopPoolInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DesktopPoolInfo) SetName(v string) {
	o.Name = v
}

// GetSettings returns the Settings field value
func (o *DesktopPoolInfo) GetSettings() DesktopPoolSettings {
	if o == nil {
		var ret DesktopPoolSettings
		return ret
	}

	return o.Settings
}

// GetSettingsOk returns a tuple with the Settings field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolInfo) GetSettingsOk() (*DesktopPoolSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Settings, true
}

// SetSettings sets field value
func (o *DesktopPoolInfo) SetSettings(v DesktopPoolSettings) {
	o.Settings = v
}

// GetSource returns the Source field value
func (o *DesktopPoolInfo) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolInfo) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DesktopPoolInfo) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value
func (o *DesktopPoolInfo) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolInfo) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DesktopPoolInfo) SetType(v string) {
	o.Type = v
}

func (o DesktopPoolInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["settings"] = o.Settings
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolInfo struct {
	value *DesktopPoolInfo
	isSet bool
}

func (v NullableDesktopPoolInfo) Get() *DesktopPoolInfo {
	return v.value
}

func (v *NullableDesktopPoolInfo) Set(val *DesktopPoolInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolInfo(val *DesktopPoolInfo) *NullableDesktopPoolInfo {
	return &NullableDesktopPoolInfo{value: val, isSet: true}
}

func (v NullableDesktopPoolInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
