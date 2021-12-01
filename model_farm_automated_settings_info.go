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

// FarmAutomatedSettingsInfo Settings for Automated farm.
type FarmAutomatedSettingsInfo struct {
	CustomizationSettings *FarmCustomizationSettingsInfo `json:"customization_settings,omitempty"`
	// Indicates whether to enable provisioning immediately.<br>Supported Filters: 'Equals'.
	EnableProvisioning *bool `json:"enable_provisioning,omitempty"`
	// Source of image used in the farm.<br>Supported Filters: 'Equals'. * VIRTUAL_CENTER: Image was created in virtual center. * IMAGE_CATALOG: Image was created in image catalog.
	ImageSource *string `json:"image_source,omitempty"`
	// RDS Server type for max sessions. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions.
	MaxSessionType *string `json:"max_session_type,omitempty"`
	// Maximum number of sessions allowed for RDS Server. This is set when max_session_type is LIMITED.
	MaxSessions *int32 `json:"max_sessions,omitempty"`
	// Minimum number of ready (provisioned) RDS Servers during Instant clone maintenance operations. Use this setting to perform machine maintenance operations in a rolling fashion. Increasing this count may decrease the concurrency for Instant clone maintenance operations for the automated farm.
	MinReadyVms *int32 `json:"min_ready_vms,omitempty"`
	// Network interface card settings for RDS Servers provisioned for this farm. A NIC may appear at most once in these settings and must be present on this RDS Server's parent's snapshot. Not all NICs need be configured. Any that are not will use default settings.
	Nics *[]FarmNetworkInterfaceCardSettingsInfo `json:"nics,omitempty"`
	// The guest operating system. * UNKNOWN: Unknown * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_SERVER_OTHER: Linux Server (other)
	OperatingSystem *string `json:"operating_system,omitempty"`
	// The guest operating system architecture. * UNKNOWN: Operating System cannot be determined. * BIT_32: 32 bit Operating System Architecture. * BIT_64: 64 bit Operating System Architecture.
	OperatingSystemArchitecture *string                                  `json:"operating_system_architecture,omitempty"`
	PatternNamingSettings       *FarmRDSServersPatternNamingSettingsInfo `json:"pattern_naming_settings,omitempty"`
	ProvisioningSettings        *FarmProvisioningSettingsInfo            `json:"provisioning_settings,omitempty"`
	ProvisioningStatusData      *FarmProvisioningStatusInfo              `json:"provisioning_status_data,omitempty"`
	// Indicates whether provisioning on all VMs stops on error.
	StopProvisioningOnError *bool                    `json:"stop_provisioning_on_error,omitempty"`
	StorageSettings         *FarmStorageSettingsInfo `json:"storage_settings,omitempty"`
	// Transparent page sharing scope for the farm. * VM: Inter-VM page sharing is not permitted. * FARM: Inter-VM page sharing among VMs belonging to the same automated farm is permitted. * POD: Inter-VM page sharing among VMs belonging to the same Pod is permitted. * GLOBAL: Inter-VM page sharing among all VMs on the same host is permitted.
	TransparentPageSharingScope *string `json:"transparent_page_sharing_scope,omitempty"`
	// ID of the virtual center server.<br>Supported Filters: 'Equals'.
	VcenterId *string `json:"vcenter_id,omitempty"`
}

// NewFarmAutomatedSettingsInfo instantiates a new FarmAutomatedSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmAutomatedSettingsInfo() *FarmAutomatedSettingsInfo {
	this := FarmAutomatedSettingsInfo{}
	return &this
}

// NewFarmAutomatedSettingsInfoWithDefaults instantiates a new FarmAutomatedSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmAutomatedSettingsInfoWithDefaults() *FarmAutomatedSettingsInfo {
	this := FarmAutomatedSettingsInfo{}
	return &this
}

// GetCustomizationSettings returns the CustomizationSettings field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetCustomizationSettings() FarmCustomizationSettingsInfo {
	if o == nil || o.CustomizationSettings == nil {
		var ret FarmCustomizationSettingsInfo
		return ret
	}
	return *o.CustomizationSettings
}

// GetCustomizationSettingsOk returns a tuple with the CustomizationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetCustomizationSettingsOk() (*FarmCustomizationSettingsInfo, bool) {
	if o == nil || o.CustomizationSettings == nil {
		return nil, false
	}
	return o.CustomizationSettings, true
}

// HasCustomizationSettings returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasCustomizationSettings() bool {
	if o != nil && o.CustomizationSettings != nil {
		return true
	}

	return false
}

// SetCustomizationSettings gets a reference to the given FarmCustomizationSettingsInfo and assigns it to the CustomizationSettings field.
func (o *FarmAutomatedSettingsInfo) SetCustomizationSettings(v FarmCustomizationSettingsInfo) {
	o.CustomizationSettings = &v
}

// GetEnableProvisioning returns the EnableProvisioning field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetEnableProvisioning() bool {
	if o == nil || o.EnableProvisioning == nil {
		var ret bool
		return ret
	}
	return *o.EnableProvisioning
}

// GetEnableProvisioningOk returns a tuple with the EnableProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetEnableProvisioningOk() (*bool, bool) {
	if o == nil || o.EnableProvisioning == nil {
		return nil, false
	}
	return o.EnableProvisioning, true
}

// HasEnableProvisioning returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasEnableProvisioning() bool {
	if o != nil && o.EnableProvisioning != nil {
		return true
	}

	return false
}

// SetEnableProvisioning gets a reference to the given bool and assigns it to the EnableProvisioning field.
func (o *FarmAutomatedSettingsInfo) SetEnableProvisioning(v bool) {
	o.EnableProvisioning = &v
}

// GetImageSource returns the ImageSource field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetImageSource() string {
	if o == nil || o.ImageSource == nil {
		var ret string
		return ret
	}
	return *o.ImageSource
}

// GetImageSourceOk returns a tuple with the ImageSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetImageSourceOk() (*string, bool) {
	if o == nil || o.ImageSource == nil {
		return nil, false
	}
	return o.ImageSource, true
}

// HasImageSource returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasImageSource() bool {
	if o != nil && o.ImageSource != nil {
		return true
	}

	return false
}

// SetImageSource gets a reference to the given string and assigns it to the ImageSource field.
func (o *FarmAutomatedSettingsInfo) SetImageSource(v string) {
	o.ImageSource = &v
}

// GetMaxSessionType returns the MaxSessionType field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetMaxSessionType() string {
	if o == nil || o.MaxSessionType == nil {
		var ret string
		return ret
	}
	return *o.MaxSessionType
}

// GetMaxSessionTypeOk returns a tuple with the MaxSessionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetMaxSessionTypeOk() (*string, bool) {
	if o == nil || o.MaxSessionType == nil {
		return nil, false
	}
	return o.MaxSessionType, true
}

// HasMaxSessionType returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasMaxSessionType() bool {
	if o != nil && o.MaxSessionType != nil {
		return true
	}

	return false
}

// SetMaxSessionType gets a reference to the given string and assigns it to the MaxSessionType field.
func (o *FarmAutomatedSettingsInfo) SetMaxSessionType(v string) {
	o.MaxSessionType = &v
}

// GetMaxSessions returns the MaxSessions field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetMaxSessions() int32 {
	if o == nil || o.MaxSessions == nil {
		var ret int32
		return ret
	}
	return *o.MaxSessions
}

// GetMaxSessionsOk returns a tuple with the MaxSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetMaxSessionsOk() (*int32, bool) {
	if o == nil || o.MaxSessions == nil {
		return nil, false
	}
	return o.MaxSessions, true
}

// HasMaxSessions returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasMaxSessions() bool {
	if o != nil && o.MaxSessions != nil {
		return true
	}

	return false
}

// SetMaxSessions gets a reference to the given int32 and assigns it to the MaxSessions field.
func (o *FarmAutomatedSettingsInfo) SetMaxSessions(v int32) {
	o.MaxSessions = &v
}

// GetMinReadyVms returns the MinReadyVms field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetMinReadyVms() int32 {
	if o == nil || o.MinReadyVms == nil {
		var ret int32
		return ret
	}
	return *o.MinReadyVms
}

// GetMinReadyVmsOk returns a tuple with the MinReadyVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetMinReadyVmsOk() (*int32, bool) {
	if o == nil || o.MinReadyVms == nil {
		return nil, false
	}
	return o.MinReadyVms, true
}

// HasMinReadyVms returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasMinReadyVms() bool {
	if o != nil && o.MinReadyVms != nil {
		return true
	}

	return false
}

// SetMinReadyVms gets a reference to the given int32 and assigns it to the MinReadyVms field.
func (o *FarmAutomatedSettingsInfo) SetMinReadyVms(v int32) {
	o.MinReadyVms = &v
}

// GetNics returns the Nics field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetNics() []FarmNetworkInterfaceCardSettingsInfo {
	if o == nil || o.Nics == nil {
		var ret []FarmNetworkInterfaceCardSettingsInfo
		return ret
	}
	return *o.Nics
}

// GetNicsOk returns a tuple with the Nics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetNicsOk() (*[]FarmNetworkInterfaceCardSettingsInfo, bool) {
	if o == nil || o.Nics == nil {
		return nil, false
	}
	return o.Nics, true
}

// HasNics returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasNics() bool {
	if o != nil && o.Nics != nil {
		return true
	}

	return false
}

// SetNics gets a reference to the given []FarmNetworkInterfaceCardSettingsInfo and assigns it to the Nics field.
func (o *FarmAutomatedSettingsInfo) SetNics(v []FarmNetworkInterfaceCardSettingsInfo) {
	o.Nics = &v
}

// GetOperatingSystem returns the OperatingSystem field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetOperatingSystem() string {
	if o == nil || o.OperatingSystem == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetOperatingSystemOk() (*string, bool) {
	if o == nil || o.OperatingSystem == nil {
		return nil, false
	}
	return o.OperatingSystem, true
}

// HasOperatingSystem returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasOperatingSystem() bool {
	if o != nil && o.OperatingSystem != nil {
		return true
	}

	return false
}

// SetOperatingSystem gets a reference to the given string and assigns it to the OperatingSystem field.
func (o *FarmAutomatedSettingsInfo) SetOperatingSystem(v string) {
	o.OperatingSystem = &v
}

// GetOperatingSystemArchitecture returns the OperatingSystemArchitecture field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetOperatingSystemArchitecture() string {
	if o == nil || o.OperatingSystemArchitecture == nil {
		var ret string
		return ret
	}
	return *o.OperatingSystemArchitecture
}

// GetOperatingSystemArchitectureOk returns a tuple with the OperatingSystemArchitecture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetOperatingSystemArchitectureOk() (*string, bool) {
	if o == nil || o.OperatingSystemArchitecture == nil {
		return nil, false
	}
	return o.OperatingSystemArchitecture, true
}

// HasOperatingSystemArchitecture returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasOperatingSystemArchitecture() bool {
	if o != nil && o.OperatingSystemArchitecture != nil {
		return true
	}

	return false
}

// SetOperatingSystemArchitecture gets a reference to the given string and assigns it to the OperatingSystemArchitecture field.
func (o *FarmAutomatedSettingsInfo) SetOperatingSystemArchitecture(v string) {
	o.OperatingSystemArchitecture = &v
}

// GetPatternNamingSettings returns the PatternNamingSettings field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetPatternNamingSettings() FarmRDSServersPatternNamingSettingsInfo {
	if o == nil || o.PatternNamingSettings == nil {
		var ret FarmRDSServersPatternNamingSettingsInfo
		return ret
	}
	return *o.PatternNamingSettings
}

// GetPatternNamingSettingsOk returns a tuple with the PatternNamingSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetPatternNamingSettingsOk() (*FarmRDSServersPatternNamingSettingsInfo, bool) {
	if o == nil || o.PatternNamingSettings == nil {
		return nil, false
	}
	return o.PatternNamingSettings, true
}

// HasPatternNamingSettings returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasPatternNamingSettings() bool {
	if o != nil && o.PatternNamingSettings != nil {
		return true
	}

	return false
}

// SetPatternNamingSettings gets a reference to the given FarmRDSServersPatternNamingSettingsInfo and assigns it to the PatternNamingSettings field.
func (o *FarmAutomatedSettingsInfo) SetPatternNamingSettings(v FarmRDSServersPatternNamingSettingsInfo) {
	o.PatternNamingSettings = &v
}

// GetProvisioningSettings returns the ProvisioningSettings field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetProvisioningSettings() FarmProvisioningSettingsInfo {
	if o == nil || o.ProvisioningSettings == nil {
		var ret FarmProvisioningSettingsInfo
		return ret
	}
	return *o.ProvisioningSettings
}

// GetProvisioningSettingsOk returns a tuple with the ProvisioningSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetProvisioningSettingsOk() (*FarmProvisioningSettingsInfo, bool) {
	if o == nil || o.ProvisioningSettings == nil {
		return nil, false
	}
	return o.ProvisioningSettings, true
}

// HasProvisioningSettings returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasProvisioningSettings() bool {
	if o != nil && o.ProvisioningSettings != nil {
		return true
	}

	return false
}

// SetProvisioningSettings gets a reference to the given FarmProvisioningSettingsInfo and assigns it to the ProvisioningSettings field.
func (o *FarmAutomatedSettingsInfo) SetProvisioningSettings(v FarmProvisioningSettingsInfo) {
	o.ProvisioningSettings = &v
}

// GetProvisioningStatusData returns the ProvisioningStatusData field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetProvisioningStatusData() FarmProvisioningStatusInfo {
	if o == nil || o.ProvisioningStatusData == nil {
		var ret FarmProvisioningStatusInfo
		return ret
	}
	return *o.ProvisioningStatusData
}

// GetProvisioningStatusDataOk returns a tuple with the ProvisioningStatusData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetProvisioningStatusDataOk() (*FarmProvisioningStatusInfo, bool) {
	if o == nil || o.ProvisioningStatusData == nil {
		return nil, false
	}
	return o.ProvisioningStatusData, true
}

// HasProvisioningStatusData returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasProvisioningStatusData() bool {
	if o != nil && o.ProvisioningStatusData != nil {
		return true
	}

	return false
}

// SetProvisioningStatusData gets a reference to the given FarmProvisioningStatusInfo and assigns it to the ProvisioningStatusData field.
func (o *FarmAutomatedSettingsInfo) SetProvisioningStatusData(v FarmProvisioningStatusInfo) {
	o.ProvisioningStatusData = &v
}

// GetStopProvisioningOnError returns the StopProvisioningOnError field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetStopProvisioningOnError() bool {
	if o == nil || o.StopProvisioningOnError == nil {
		var ret bool
		return ret
	}
	return *o.StopProvisioningOnError
}

// GetStopProvisioningOnErrorOk returns a tuple with the StopProvisioningOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetStopProvisioningOnErrorOk() (*bool, bool) {
	if o == nil || o.StopProvisioningOnError == nil {
		return nil, false
	}
	return o.StopProvisioningOnError, true
}

// HasStopProvisioningOnError returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasStopProvisioningOnError() bool {
	if o != nil && o.StopProvisioningOnError != nil {
		return true
	}

	return false
}

// SetStopProvisioningOnError gets a reference to the given bool and assigns it to the StopProvisioningOnError field.
func (o *FarmAutomatedSettingsInfo) SetStopProvisioningOnError(v bool) {
	o.StopProvisioningOnError = &v
}

// GetStorageSettings returns the StorageSettings field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetStorageSettings() FarmStorageSettingsInfo {
	if o == nil || o.StorageSettings == nil {
		var ret FarmStorageSettingsInfo
		return ret
	}
	return *o.StorageSettings
}

// GetStorageSettingsOk returns a tuple with the StorageSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetStorageSettingsOk() (*FarmStorageSettingsInfo, bool) {
	if o == nil || o.StorageSettings == nil {
		return nil, false
	}
	return o.StorageSettings, true
}

// HasStorageSettings returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasStorageSettings() bool {
	if o != nil && o.StorageSettings != nil {
		return true
	}

	return false
}

// SetStorageSettings gets a reference to the given FarmStorageSettingsInfo and assigns it to the StorageSettings field.
func (o *FarmAutomatedSettingsInfo) SetStorageSettings(v FarmStorageSettingsInfo) {
	o.StorageSettings = &v
}

// GetTransparentPageSharingScope returns the TransparentPageSharingScope field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetTransparentPageSharingScope() string {
	if o == nil || o.TransparentPageSharingScope == nil {
		var ret string
		return ret
	}
	return *o.TransparentPageSharingScope
}

// GetTransparentPageSharingScopeOk returns a tuple with the TransparentPageSharingScope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetTransparentPageSharingScopeOk() (*string, bool) {
	if o == nil || o.TransparentPageSharingScope == nil {
		return nil, false
	}
	return o.TransparentPageSharingScope, true
}

// HasTransparentPageSharingScope returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasTransparentPageSharingScope() bool {
	if o != nil && o.TransparentPageSharingScope != nil {
		return true
	}

	return false
}

// SetTransparentPageSharingScope gets a reference to the given string and assigns it to the TransparentPageSharingScope field.
func (o *FarmAutomatedSettingsInfo) SetTransparentPageSharingScope(v string) {
	o.TransparentPageSharingScope = &v
}

// GetVcenterId returns the VcenterId field value if set, zero value otherwise.
func (o *FarmAutomatedSettingsInfo) GetVcenterId() string {
	if o == nil || o.VcenterId == nil {
		var ret string
		return ret
	}
	return *o.VcenterId
}

// GetVcenterIdOk returns a tuple with the VcenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmAutomatedSettingsInfo) GetVcenterIdOk() (*string, bool) {
	if o == nil || o.VcenterId == nil {
		return nil, false
	}
	return o.VcenterId, true
}

// HasVcenterId returns a boolean if a field has been set.
func (o *FarmAutomatedSettingsInfo) HasVcenterId() bool {
	if o != nil && o.VcenterId != nil {
		return true
	}

	return false
}

// SetVcenterId gets a reference to the given string and assigns it to the VcenterId field.
func (o *FarmAutomatedSettingsInfo) SetVcenterId(v string) {
	o.VcenterId = &v
}

func (o FarmAutomatedSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomizationSettings != nil {
		toSerialize["customization_settings"] = o.CustomizationSettings
	}
	if o.EnableProvisioning != nil {
		toSerialize["enable_provisioning"] = o.EnableProvisioning
	}
	if o.ImageSource != nil {
		toSerialize["image_source"] = o.ImageSource
	}
	if o.MaxSessionType != nil {
		toSerialize["max_session_type"] = o.MaxSessionType
	}
	if o.MaxSessions != nil {
		toSerialize["max_sessions"] = o.MaxSessions
	}
	if o.MinReadyVms != nil {
		toSerialize["min_ready_vms"] = o.MinReadyVms
	}
	if o.Nics != nil {
		toSerialize["nics"] = o.Nics
	}
	if o.OperatingSystem != nil {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if o.OperatingSystemArchitecture != nil {
		toSerialize["operating_system_architecture"] = o.OperatingSystemArchitecture
	}
	if o.PatternNamingSettings != nil {
		toSerialize["pattern_naming_settings"] = o.PatternNamingSettings
	}
	if o.ProvisioningSettings != nil {
		toSerialize["provisioning_settings"] = o.ProvisioningSettings
	}
	if o.ProvisioningStatusData != nil {
		toSerialize["provisioning_status_data"] = o.ProvisioningStatusData
	}
	if o.StopProvisioningOnError != nil {
		toSerialize["stop_provisioning_on_error"] = o.StopProvisioningOnError
	}
	if o.StorageSettings != nil {
		toSerialize["storage_settings"] = o.StorageSettings
	}
	if o.TransparentPageSharingScope != nil {
		toSerialize["transparent_page_sharing_scope"] = o.TransparentPageSharingScope
	}
	if o.VcenterId != nil {
		toSerialize["vcenter_id"] = o.VcenterId
	}
	return json.Marshal(toSerialize)
}

type NullableFarmAutomatedSettingsInfo struct {
	value *FarmAutomatedSettingsInfo
	isSet bool
}

func (v NullableFarmAutomatedSettingsInfo) Get() *FarmAutomatedSettingsInfo {
	return v.value
}

func (v *NullableFarmAutomatedSettingsInfo) Set(val *FarmAutomatedSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmAutomatedSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmAutomatedSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmAutomatedSettingsInfo(val *FarmAutomatedSettingsInfo) *NullableFarmAutomatedSettingsInfo {
	return &NullableFarmAutomatedSettingsInfo{value: val, isSet: true}
}

func (v NullableFarmAutomatedSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmAutomatedSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
