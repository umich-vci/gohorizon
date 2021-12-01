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

// DesktopPoolProvisioningSettings Virtual Center provisioning settings for the automated desktop pool.
type DesktopPoolProvisioningSettings struct {
	// Whether to add Virtual TPM device.
	AddVirtualTpm *bool `json:"add_virtual_tpm,omitempty"`
	// Applicable To: Linked/instant clone automated desktop pools.<br>Base image snapshot for linked clone desktop pool and current image snapshot for instant clone desktop pool.
	BaseSnapshotId *string `json:"base_snapshot_id,omitempty"`
	// Datacenter within which the desktop pool is configured.
	DatacenterId *string `json:"datacenter_id,omitempty"`
	// Host or cluster where the machines are deployed in.
	HostOrClusterId *string `json:"host_or_cluster_id,omitempty"`
	// Applicable To: Full/instant clone automated desktop pools.<br>Image management stream used in desktop pool when Image Management feature is enabled.<br>Supported Filters: 'Equals'.
	ImStreamId *string `json:"im_stream_id,omitempty"`
	// Applicable To: Full/instant clone automated desktop pools.<br>Image management tag associated with the selected image management stream which is used in desktop pool when Image Management feature is enabled.<br>Supported Filters: 'Equals'.
	ImTagId *string `json:"im_tag_id,omitempty"`
	// Applicable To: Linked clone automated desktop pools.<br>Minimum number of ready (provisioned) machines during View Composer maintenance operations.
	MinReadyVmsOnVcomposerMaintenance *int32 `json:"min_ready_vms_on_vcomposer_maintenance,omitempty"`
	// Applicable To: Linked/instant clone automated desktop pools.<br>Base image VM for linked clone desktop pool and current image for instant clone desktop pool.
	ParentVmId *string `json:"parent_vm_id,omitempty"`
	// Resource pool to deploy the machines.
	ResourcePoolId *string `json:"resource_pool_id,omitempty"`
	// VM folder where the machines are deployed to.
	VmFolderId *string `json:"vm_folder_id,omitempty"`
	// Applicable To: Full clone automated desktop pools.<br>Template from which full clone machines are deployed.
	VmTemplateId *string `json:"vm_template_id,omitempty"`
}

// NewDesktopPoolProvisioningSettings instantiates a new DesktopPoolProvisioningSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolProvisioningSettings() *DesktopPoolProvisioningSettings {
	this := DesktopPoolProvisioningSettings{}
	return &this
}

// NewDesktopPoolProvisioningSettingsWithDefaults instantiates a new DesktopPoolProvisioningSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolProvisioningSettingsWithDefaults() *DesktopPoolProvisioningSettings {
	this := DesktopPoolProvisioningSettings{}
	return &this
}

// GetAddVirtualTpm returns the AddVirtualTpm field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetAddVirtualTpm() bool {
	if o == nil || o.AddVirtualTpm == nil {
		var ret bool
		return ret
	}
	return *o.AddVirtualTpm
}

// GetAddVirtualTpmOk returns a tuple with the AddVirtualTpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetAddVirtualTpmOk() (*bool, bool) {
	if o == nil || o.AddVirtualTpm == nil {
		return nil, false
	}
	return o.AddVirtualTpm, true
}

// HasAddVirtualTpm returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasAddVirtualTpm() bool {
	if o != nil && o.AddVirtualTpm != nil {
		return true
	}

	return false
}

// SetAddVirtualTpm gets a reference to the given bool and assigns it to the AddVirtualTpm field.
func (o *DesktopPoolProvisioningSettings) SetAddVirtualTpm(v bool) {
	o.AddVirtualTpm = &v
}

// GetBaseSnapshotId returns the BaseSnapshotId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetBaseSnapshotId() string {
	if o == nil || o.BaseSnapshotId == nil {
		var ret string
		return ret
	}
	return *o.BaseSnapshotId
}

// GetBaseSnapshotIdOk returns a tuple with the BaseSnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetBaseSnapshotIdOk() (*string, bool) {
	if o == nil || o.BaseSnapshotId == nil {
		return nil, false
	}
	return o.BaseSnapshotId, true
}

// HasBaseSnapshotId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasBaseSnapshotId() bool {
	if o != nil && o.BaseSnapshotId != nil {
		return true
	}

	return false
}

// SetBaseSnapshotId gets a reference to the given string and assigns it to the BaseSnapshotId field.
func (o *DesktopPoolProvisioningSettings) SetBaseSnapshotId(v string) {
	o.BaseSnapshotId = &v
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetDatacenterId() string {
	if o == nil || o.DatacenterId == nil {
		var ret string
		return ret
	}
	return *o.DatacenterId
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetDatacenterIdOk() (*string, bool) {
	if o == nil || o.DatacenterId == nil {
		return nil, false
	}
	return o.DatacenterId, true
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given string and assigns it to the DatacenterId field.
func (o *DesktopPoolProvisioningSettings) SetDatacenterId(v string) {
	o.DatacenterId = &v
}

// GetHostOrClusterId returns the HostOrClusterId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetHostOrClusterId() string {
	if o == nil || o.HostOrClusterId == nil {
		var ret string
		return ret
	}
	return *o.HostOrClusterId
}

// GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetHostOrClusterIdOk() (*string, bool) {
	if o == nil || o.HostOrClusterId == nil {
		return nil, false
	}
	return o.HostOrClusterId, true
}

// HasHostOrClusterId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasHostOrClusterId() bool {
	if o != nil && o.HostOrClusterId != nil {
		return true
	}

	return false
}

// SetHostOrClusterId gets a reference to the given string and assigns it to the HostOrClusterId field.
func (o *DesktopPoolProvisioningSettings) SetHostOrClusterId(v string) {
	o.HostOrClusterId = &v
}

// GetImStreamId returns the ImStreamId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetImStreamId() string {
	if o == nil || o.ImStreamId == nil {
		var ret string
		return ret
	}
	return *o.ImStreamId
}

// GetImStreamIdOk returns a tuple with the ImStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetImStreamIdOk() (*string, bool) {
	if o == nil || o.ImStreamId == nil {
		return nil, false
	}
	return o.ImStreamId, true
}

// HasImStreamId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasImStreamId() bool {
	if o != nil && o.ImStreamId != nil {
		return true
	}

	return false
}

// SetImStreamId gets a reference to the given string and assigns it to the ImStreamId field.
func (o *DesktopPoolProvisioningSettings) SetImStreamId(v string) {
	o.ImStreamId = &v
}

// GetImTagId returns the ImTagId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetImTagId() string {
	if o == nil || o.ImTagId == nil {
		var ret string
		return ret
	}
	return *o.ImTagId
}

// GetImTagIdOk returns a tuple with the ImTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetImTagIdOk() (*string, bool) {
	if o == nil || o.ImTagId == nil {
		return nil, false
	}
	return o.ImTagId, true
}

// HasImTagId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasImTagId() bool {
	if o != nil && o.ImTagId != nil {
		return true
	}

	return false
}

// SetImTagId gets a reference to the given string and assigns it to the ImTagId field.
func (o *DesktopPoolProvisioningSettings) SetImTagId(v string) {
	o.ImTagId = &v
}

// GetMinReadyVmsOnVcomposerMaintenance returns the MinReadyVmsOnVcomposerMaintenance field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetMinReadyVmsOnVcomposerMaintenance() int32 {
	if o == nil || o.MinReadyVmsOnVcomposerMaintenance == nil {
		var ret int32
		return ret
	}
	return *o.MinReadyVmsOnVcomposerMaintenance
}

// GetMinReadyVmsOnVcomposerMaintenanceOk returns a tuple with the MinReadyVmsOnVcomposerMaintenance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetMinReadyVmsOnVcomposerMaintenanceOk() (*int32, bool) {
	if o == nil || o.MinReadyVmsOnVcomposerMaintenance == nil {
		return nil, false
	}
	return o.MinReadyVmsOnVcomposerMaintenance, true
}

// HasMinReadyVmsOnVcomposerMaintenance returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasMinReadyVmsOnVcomposerMaintenance() bool {
	if o != nil && o.MinReadyVmsOnVcomposerMaintenance != nil {
		return true
	}

	return false
}

// SetMinReadyVmsOnVcomposerMaintenance gets a reference to the given int32 and assigns it to the MinReadyVmsOnVcomposerMaintenance field.
func (o *DesktopPoolProvisioningSettings) SetMinReadyVmsOnVcomposerMaintenance(v int32) {
	o.MinReadyVmsOnVcomposerMaintenance = &v
}

// GetParentVmId returns the ParentVmId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetParentVmId() string {
	if o == nil || o.ParentVmId == nil {
		var ret string
		return ret
	}
	return *o.ParentVmId
}

// GetParentVmIdOk returns a tuple with the ParentVmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetParentVmIdOk() (*string, bool) {
	if o == nil || o.ParentVmId == nil {
		return nil, false
	}
	return o.ParentVmId, true
}

// HasParentVmId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasParentVmId() bool {
	if o != nil && o.ParentVmId != nil {
		return true
	}

	return false
}

// SetParentVmId gets a reference to the given string and assigns it to the ParentVmId field.
func (o *DesktopPoolProvisioningSettings) SetParentVmId(v string) {
	o.ParentVmId = &v
}

// GetResourcePoolId returns the ResourcePoolId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetResourcePoolId() string {
	if o == nil || o.ResourcePoolId == nil {
		var ret string
		return ret
	}
	return *o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetResourcePoolIdOk() (*string, bool) {
	if o == nil || o.ResourcePoolId == nil {
		return nil, false
	}
	return o.ResourcePoolId, true
}

// HasResourcePoolId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasResourcePoolId() bool {
	if o != nil && o.ResourcePoolId != nil {
		return true
	}

	return false
}

// SetResourcePoolId gets a reference to the given string and assigns it to the ResourcePoolId field.
func (o *DesktopPoolProvisioningSettings) SetResourcePoolId(v string) {
	o.ResourcePoolId = &v
}

// GetVmFolderId returns the VmFolderId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetVmFolderId() string {
	if o == nil || o.VmFolderId == nil {
		var ret string
		return ret
	}
	return *o.VmFolderId
}

// GetVmFolderIdOk returns a tuple with the VmFolderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetVmFolderIdOk() (*string, bool) {
	if o == nil || o.VmFolderId == nil {
		return nil, false
	}
	return o.VmFolderId, true
}

// HasVmFolderId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasVmFolderId() bool {
	if o != nil && o.VmFolderId != nil {
		return true
	}

	return false
}

// SetVmFolderId gets a reference to the given string and assigns it to the VmFolderId field.
func (o *DesktopPoolProvisioningSettings) SetVmFolderId(v string) {
	o.VmFolderId = &v
}

// GetVmTemplateId returns the VmTemplateId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningSettings) GetVmTemplateId() string {
	if o == nil || o.VmTemplateId == nil {
		var ret string
		return ret
	}
	return *o.VmTemplateId
}

// GetVmTemplateIdOk returns a tuple with the VmTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningSettings) GetVmTemplateIdOk() (*string, bool) {
	if o == nil || o.VmTemplateId == nil {
		return nil, false
	}
	return o.VmTemplateId, true
}

// HasVmTemplateId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningSettings) HasVmTemplateId() bool {
	if o != nil && o.VmTemplateId != nil {
		return true
	}

	return false
}

// SetVmTemplateId gets a reference to the given string and assigns it to the VmTemplateId field.
func (o *DesktopPoolProvisioningSettings) SetVmTemplateId(v string) {
	o.VmTemplateId = &v
}

func (o DesktopPoolProvisioningSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddVirtualTpm != nil {
		toSerialize["add_virtual_tpm"] = o.AddVirtualTpm
	}
	if o.BaseSnapshotId != nil {
		toSerialize["base_snapshot_id"] = o.BaseSnapshotId
	}
	if o.DatacenterId != nil {
		toSerialize["datacenter_id"] = o.DatacenterId
	}
	if o.HostOrClusterId != nil {
		toSerialize["host_or_cluster_id"] = o.HostOrClusterId
	}
	if o.ImStreamId != nil {
		toSerialize["im_stream_id"] = o.ImStreamId
	}
	if o.ImTagId != nil {
		toSerialize["im_tag_id"] = o.ImTagId
	}
	if o.MinReadyVmsOnVcomposerMaintenance != nil {
		toSerialize["min_ready_vms_on_vcomposer_maintenance"] = o.MinReadyVmsOnVcomposerMaintenance
	}
	if o.ParentVmId != nil {
		toSerialize["parent_vm_id"] = o.ParentVmId
	}
	if o.ResourcePoolId != nil {
		toSerialize["resource_pool_id"] = o.ResourcePoolId
	}
	if o.VmFolderId != nil {
		toSerialize["vm_folder_id"] = o.VmFolderId
	}
	if o.VmTemplateId != nil {
		toSerialize["vm_template_id"] = o.VmTemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolProvisioningSettings struct {
	value *DesktopPoolProvisioningSettings
	isSet bool
}

func (v NullableDesktopPoolProvisioningSettings) Get() *DesktopPoolProvisioningSettings {
	return v.value
}

func (v *NullableDesktopPoolProvisioningSettings) Set(val *DesktopPoolProvisioningSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolProvisioningSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolProvisioningSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolProvisioningSettings(val *DesktopPoolProvisioningSettings) *NullableDesktopPoolProvisioningSettings {
	return &NullableDesktopPoolProvisioningSettings{value: val, isSet: true}
}

func (v NullableDesktopPoolProvisioningSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolProvisioningSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
