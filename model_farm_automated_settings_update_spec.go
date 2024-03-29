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

// FarmAutomatedSettingsUpdateSpec Settings for Automated farm. This is required if type is set to AUTOMATED.
type FarmAutomatedSettingsUpdateSpec struct {
	CustomizationSettings FarmCustomizationSettingsUpdateSpec `json:"customization_settings"`
	// Indicates whether to enable provisioning immediately.
	EnableProvisioning bool `json:"enable_provisioning"`
	// RDS Server type for max sessions. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions.
	MaxSessionType string `json:"max_session_type"`
	// Maximum number of sessions allowed for RDS Server. This is required if max_session_type is set to LIMITED.
	MaxSessions *int32 `json:"max_sessions,omitempty"`
	// Minimum number of ready (provisioned) RDS Servers during Instant clone maintenance operations. Use this setting to perform machine maintenance operations in a rolling fashion. Increasing this count may decrease the concurrency for Instant clone maintenance operations for the automated farm.
	MinReadyVms int32 `json:"min_ready_vms"`
	// Network interface card settings for RDS Servers provisioned for this farm. A NIC may appear at most once in these settings and must be present on this RDS Server's parent's snapshot. Not all NICs need be configured. Any that are not will use default settings.
	Nics                  *[]FarmNetworkInterfaceCardSettingsUpdateSpec `json:"nics,omitempty"`
	PatternNamingSettings FarmRDSServersPatternNamingSettingsUpdateSpec `json:"pattern_naming_settings"`
	ProvisioningSettings  FarmProvisioningSettingsUpdateSpec            `json:"provisioning_settings"`
	// Indicates whether provisioning on all VMs stops on error.
	StopProvisioningOnError bool                          `json:"stop_provisioning_on_error"`
	StorageSettings         FarmStorageSettingsUpdateSpec `json:"storage_settings"`
	// Transparent page sharing scope for the farm. * VM: Inter-VM page sharing is not permitted. * FARM: Inter-VM page sharing among VMs belonging to the same automated farm is permitted. * POD: Inter-VM page sharing among VMs belonging to the same Pod is permitted. * GLOBAL: Inter-VM page sharing among all VMs on the same host is permitted.
	TransparentPageSharingScope string `json:"transparent_page_sharing_scope"`
}

// NewFarmAutomatedSettingsUpdateSpec instantiates a new FarmAutomatedSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmAutomatedSettingsUpdateSpec(customizationSettings FarmCustomizationSettingsUpdateSpec, enableProvisioning bool, maxSessionType string, minReadyVms int32, patternNamingSettings FarmRDSServersPatternNamingSettingsUpdateSpec, provisioningSettings FarmProvisioningSettingsUpdateSpec, stopProvisioningOnError bool, storageSettings FarmStorageSettingsUpdateSpec, transparentPageSharingScope string) *FarmAutomatedSettingsUpdateSpec {
	this := FarmAutomatedSettingsUpdateSpec{}
	this.CustomizationSettings = customizationSettings
	this.EnableProvisioning = enableProvisioning
	this.MaxSessionType = maxSessionType
	this.MinReadyVms = minReadyVms
	this.PatternNamingSettings = patternNamingSettings
	this.ProvisioningSettings = provisioningSettings
	this.StopProvisioningOnError = stopProvisioningOnError
	this.StorageSettings = storageSettings
	this.TransparentPageSharingScope = transparentPageSharingScope
	return &this
}

// NewFarmAutomatedSettingsUpdateSpecWithDefaults instantiates a new FarmAutomatedSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmAutomatedSettingsUpdateSpecWithDefaults() *FarmAutomatedSettingsUpdateSpec {
	this := FarmAutomatedSettingsUpdateSpec{}
	return &this
}

// GetCustomizationSettings returns the CustomizationSettings field value
func (o *FarmAutomatedSettingsUpdateSpec) GetCustomizationSettings() FarmCustomizationSettingsUpdateSpec {
	if o == nil {
		var ret FarmCustomizationSettingsUpdateSpec
		return ret
	}

	return o.CustomizationSettings
}

// GetCustomizationSettingsOk returns a tuple with the CustomizationSettings field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetCustomizationSettingsOk() (*FarmCustomizationSettingsUpdateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CustomizationSettings, true
}

// SetCustomizationSettings sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetCustomizationSettings(v FarmCustomizationSettingsUpdateSpec) {
	o.CustomizationSettings = v
}

// GetEnableProvisioning returns the EnableProvisioning field value
func (o *FarmAutomatedSettingsUpdateSpec) GetEnableProvisioning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableProvisioning
}

// GetEnableProvisioningOk returns a tuple with the EnableProvisioning field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetEnableProvisioningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableProvisioning, true
}

// SetEnableProvisioning sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetEnableProvisioning(v bool) {
	o.EnableProvisioning = v
}

// GetMaxSessionType returns the MaxSessionType field value
func (o *FarmAutomatedSettingsUpdateSpec) GetMaxSessionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaxSessionType
}

// GetMaxSessionTypeOk returns a tuple with the MaxSessionType field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetMaxSessionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSessionType, true
}

// SetMaxSessionType sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetMaxSessionType(v string) {
	o.MaxSessionType = v
}

// GetMaxSessions returns the MaxSessions field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsUpdateSpec) GetMaxSessions() int32 {
	if o == nil || o.MaxSessions == nil {
		var ret int32
		return ret
	}
	return *o.MaxSessions
}

// GetMaxSessionsOk returns a tuple with the MaxSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetMaxSessionsOk() (*int32, bool) {
	if o == nil || o.MaxSessions == nil {
		return nil, false
	}
	return o.MaxSessions, true
}

// HasMaxSessions returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsUpdateSpec) HasMaxSessions() bool {
	if o != nil && o.MaxSessions != nil {
		return true
	}

	return false
}

// SetMaxSessions gets a reference to the given int32 and assigns it to the MaxSessions field.
func (o *FarmAutomatedSettingsUpdateSpec) SetMaxSessions(v int32) {
	o.MaxSessions = &v
}

// GetMinReadyVms returns the MinReadyVms field value
func (o *FarmAutomatedSettingsUpdateSpec) GetMinReadyVms() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinReadyVms
}

// GetMinReadyVmsOk returns a tuple with the MinReadyVms field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetMinReadyVmsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinReadyVms, true
}

// SetMinReadyVms sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetMinReadyVms(v int32) {
	o.MinReadyVms = v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsUpdateSpec) GetNics() []FarmNetworkInterfaceCardSettingsUpdateSpec {
	if o == nil || o.Nics == nil {
		var ret []FarmNetworkInterfaceCardSettingsUpdateSpec
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetNicsOk() (*[]FarmNetworkInterfaceCardSettingsUpdateSpec, bool) {
	if o == nil || o.Nics == nil {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsUpdateSpec) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

// SetNics gets a reference to the given []FarmNetworkInterfaceCardSettingsUpdateSpec and assigns it to the Nics field.
func (o *FarmAutomatedSettingsUpdateSpec) SetNics(v []FarmNetworkInterfaceCardSettingsUpdateSpec) {
	o.Nics = &v
}

// GetPatternNamingSettings returns the PatternNamingSettings field value
func (o *FarmAutomatedSettingsUpdateSpec) GetPatternNamingSettings() FarmRDSServersPatternNamingSettingsUpdateSpec {
	if o == nil {
		var ret FarmRDSServersPatternNamingSettingsUpdateSpec
		return ret
	}

	return o.PatternNamingSettings
}

// GetPatternNamingSettingsOk returns a tuple with the PatternNamingSettings field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetPatternNamingSettingsOk() (*FarmRDSServersPatternNamingSettingsUpdateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PatternNamingSettings, true
}

// SetPatternNamingSettings sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetPatternNamingSettings(v FarmRDSServersPatternNamingSettingsUpdateSpec) {
	o.PatternNamingSettings = v
}

// GetProvisioningSettings returns the ProvisioningSettings field value
func (o *FarmAutomatedSettingsUpdateSpec) GetProvisioningSettings() FarmProvisioningSettingsUpdateSpec {
	if o == nil {
		var ret FarmProvisioningSettingsUpdateSpec
		return ret
	}

	return o.ProvisioningSettings
}

// GetProvisioningSettingsOk returns a tuple with the ProvisioningSettings field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetProvisioningSettingsOk() (*FarmProvisioningSettingsUpdateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningSettings, true
}

// SetProvisioningSettings sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetProvisioningSettings(v FarmProvisioningSettingsUpdateSpec) {
	o.ProvisioningSettings = v
}

// GetStopProvisioningOnError returns the StopProvisioningOnError field value
func (o *FarmAutomatedSettingsUpdateSpec) GetStopProvisioningOnError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.StopProvisioningOnError
}

// GetStopProvisioningOnErrorOk returns a tuple with the StopProvisioningOnError field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetStopProvisioningOnErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StopProvisioningOnError, true
}

// SetStopProvisioningOnError sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetStopProvisioningOnError(v bool) {
	o.StopProvisioningOnError = v
}

// GetStorageSettings returns the StorageSettings field value
func (o *FarmAutomatedSettingsUpdateSpec) GetStorageSettings() FarmStorageSettingsUpdateSpec {
	if o == nil {
		var ret FarmStorageSettingsUpdateSpec
		return ret
	}

	return o.StorageSettings
}

// GetStorageSettingsOk returns a tuple with the StorageSettings field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetStorageSettingsOk() (*FarmStorageSettingsUpdateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorageSettings, true
}

// SetStorageSettings sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetStorageSettings(v FarmStorageSettingsUpdateSpec) {
	o.StorageSettings = v
}

// GetTransparentPageSharingScope returns the TransparentPageSharingScope field value
func (o *FarmAutomatedSettingsUpdateSpec) GetTransparentPageSharingScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransparentPageSharingScope
}

// GetTransparentPageSharingScopeOk returns a tuple with the TransparentPageSharingScope field value
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsUpdateSpec) GetTransparentPageSharingScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransparentPageSharingScope, true
}

// SetTransparentPageSharingScope sets field value
func (o *FarmAutomatedSettingsUpdateSpec) SetTransparentPageSharingScope(v string) {
	o.TransparentPageSharingScope = v
}

func (o FarmAutomatedSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["customization_settings"] = o.CustomizationSettings
	}
	if true {
		toSerialize["enable_provisioning"] = o.EnableProvisioning
	}
	if true {
		toSerialize["max_session_type"] = o.MaxSessionType
	}
	if o.MaxSessions != nil {
		toSerialize["max_sessions"] = o.MaxSessions
	}
	if true {
		toSerialize["min_ready_vms"] = o.MinReadyVms
	}
	if o.Nics != nil {
		toSerialize["nics"] = o.Nics
	}
	if true {
		toSerialize["pattern_naming_settings"] = o.PatternNamingSettings
	}
	if true {
		toSerialize["provisioning_settings"] = o.ProvisioningSettings
	}
	if true {
		toSerialize["stop_provisioning_on_error"] = o.StopProvisioningOnError
	}
	if true {
		toSerialize["storage_settings"] = o.StorageSettings
	}
	if true {
		toSerialize["transparent_page_sharing_scope"] = o.TransparentPageSharingScope
	}
	return json.Marshal(toSerialize)
}

type NullableFarmAutomatedSettingsUpdateSpec struct {
	value *FarmAutomatedSettingsUpdateSpec
	isSet bool
}

func (v NullableFarmAutomatedSettingsUpdateSpec) Get() *FarmAutomatedSettingsUpdateSpec {
	return v.value
}

func (v *NullableFarmAutomatedSettingsUpdateSpec) Set(val *FarmAutomatedSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmAutomatedSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmAutomatedSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmAutomatedSettingsUpdateSpec(val *FarmAutomatedSettingsUpdateSpec) *NullableFarmAutomatedSettingsUpdateSpec {
	return &NullableFarmAutomatedSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableFarmAutomatedSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmAutomatedSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
