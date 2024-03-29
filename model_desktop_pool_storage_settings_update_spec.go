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

// DesktopPoolStorageSettingsUpdateSpec Applicable To: Automated desktop pool. <br>Virtual Center storage settings for Automated desktop pool.
type DesktopPoolStorageSettingsUpdateSpec struct {
	// Datastores to store the machine (or the OS disk using other options for instant clone machine storage)
	Datastores []DesktopPoolDatastoreSettingsUpdateSpec `json:"datastores"`
	// With vSphere 5.x, virtual machines can be configured to use a space efficient disk format that supports reclamation of unused disk space (such as deleted files). This option reclaims unused disk space on each virtual machine. The operation is initiated when an estimate of used disk space exceeds the specified threshold.
	ReclaimVmDiskSpace *bool `json:"reclaim_vm_disk_space,omitempty"`
	// Initiate reclamation when unused space on virtual machine exceeds the threshold in MB. Default value is 1.
	ReclamationThresholdMb *int64 `json:"reclamation_threshold_mb,omitempty"`
	// Datastore to store replica disks for instant clone machines.
	ReplicaDiskDatastoreId *string `json:"replica_disk_datastore_id,omitempty"`
	// Indicates whether to use separate datastores for replica and OS disks.
	UseSeparateDatastoresReplicaAndOsDisks *bool `json:"use_separate_datastores_replica_and_os_disks,omitempty"`
	// Indicates whether to use vSphere vSAN.
	UseVsan bool `json:"use_vsan"`
}

// NewDesktopPoolStorageSettingsUpdateSpec instantiates a new DesktopPoolStorageSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolStorageSettingsUpdateSpec(datastores []DesktopPoolDatastoreSettingsUpdateSpec, useVsan bool) *DesktopPoolStorageSettingsUpdateSpec {
	this := DesktopPoolStorageSettingsUpdateSpec{}
	this.Datastores = datastores
	this.UseVsan = useVsan
	return &this
}

// NewDesktopPoolStorageSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolStorageSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolStorageSettingsUpdateSpecWithDefaults() *DesktopPoolStorageSettingsUpdateSpec {
	this := DesktopPoolStorageSettingsUpdateSpec{}
	return &this
}

// GetDatastores returns the Datastores field value
func (o *DesktopPoolStorageSettingsUpdateSpec) GetDatastores() []DesktopPoolDatastoreSettingsUpdateSpec {
	if o == nil {
		var ret []DesktopPoolDatastoreSettingsUpdateSpec
		return ret
	}

	return o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetDatastoresOk() (*[]DesktopPoolDatastoreSettingsUpdateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datastores, true
}

// SetDatastores sets field value
func (o *DesktopPoolStorageSettingsUpdateSpec) SetDatastores(v []DesktopPoolDatastoreSettingsUpdateSpec) {
	o.Datastores = v
}

// GetReclaimVmDiskSpace returns the ReclaimVmDiskSpace field value if set, zero value otherwise.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetReclaimVmDiskSpace() bool {
	if o == nil || o.ReclaimVmDiskSpace == nil {
		var ret bool
		return ret
	}
	return *o.ReclaimVmDiskSpace
}

// GetReclaimVmDiskSpaceOk returns a tuple with the ReclaimVmDiskSpace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetReclaimVmDiskSpaceOk() (*bool, bool) {
	if o == nil || o.ReclaimVmDiskSpace == nil {
		return nil, false
	}
	return o.ReclaimVmDiskSpace, true
}

// HasReclaimVmDiskSpace returns a boolean if a field has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) HasReclaimVmDiskSpace() bool {
	if o != nil && o.ReclaimVmDiskSpace != nil {
		return true
	}

	return false
}

// SetReclaimVmDiskSpace gets a reference to the given bool and assigns it to the ReclaimVmDiskSpace field.
func (o *DesktopPoolStorageSettingsUpdateSpec) SetReclaimVmDiskSpace(v bool) {
	o.ReclaimVmDiskSpace = &v
}

// GetReclamationThresholdMb returns the ReclamationThresholdMb field value if set, zero value otherwise.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetReclamationThresholdMb() int64 {
	if o == nil || o.ReclamationThresholdMb == nil {
		var ret int64
		return ret
	}
	return *o.ReclamationThresholdMb
}

// GetReclamationThresholdMbOk returns a tuple with the ReclamationThresholdMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetReclamationThresholdMbOk() (*int64, bool) {
	if o == nil || o.ReclamationThresholdMb == nil {
		return nil, false
	}
	return o.ReclamationThresholdMb, true
}

// HasReclamationThresholdMb returns a boolean if a field has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) HasReclamationThresholdMb() bool {
	if o != nil && o.ReclamationThresholdMb != nil {
		return true
	}

	return false
}

// SetReclamationThresholdMb gets a reference to the given int64 and assigns it to the ReclamationThresholdMb field.
func (o *DesktopPoolStorageSettingsUpdateSpec) SetReclamationThresholdMb(v int64) {
	o.ReclamationThresholdMb = &v
}

// GetReplicaDiskDatastoreId returns the ReplicaDiskDatastoreId field value if set, zero value otherwise.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetReplicaDiskDatastoreId() string {
	if o == nil || o.ReplicaDiskDatastoreId == nil {
		var ret string
		return ret
	}
	return *o.ReplicaDiskDatastoreId
}

// GetReplicaDiskDatastoreIdOk returns a tuple with the ReplicaDiskDatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetReplicaDiskDatastoreIdOk() (*string, bool) {
	if o == nil || o.ReplicaDiskDatastoreId == nil {
		return nil, false
	}
	return o.ReplicaDiskDatastoreId, true
}

// HasReplicaDiskDatastoreId returns a boolean if a field has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) HasReplicaDiskDatastoreId() bool {
	if o != nil && o.ReplicaDiskDatastoreId != nil {
		return true
	}

	return false
}

// SetReplicaDiskDatastoreId gets a reference to the given string and assigns it to the ReplicaDiskDatastoreId field.
func (o *DesktopPoolStorageSettingsUpdateSpec) SetReplicaDiskDatastoreId(v string) {
	o.ReplicaDiskDatastoreId = &v
}

// GetUseSeparateDatastoresReplicaAndOsDisks returns the UseSeparateDatastoresReplicaAndOsDisks field value if set, zero value otherwise.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetUseSeparateDatastoresReplicaAndOsDisks() bool {
	if o == nil || o.UseSeparateDatastoresReplicaAndOsDisks == nil {
		var ret bool
		return ret
	}
	return *o.UseSeparateDatastoresReplicaAndOsDisks
}

// GetUseSeparateDatastoresReplicaAndOsDisksOk returns a tuple with the UseSeparateDatastoresReplicaAndOsDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetUseSeparateDatastoresReplicaAndOsDisksOk() (*bool, bool) {
	if o == nil || o.UseSeparateDatastoresReplicaAndOsDisks == nil {
		return nil, false
	}
	return o.UseSeparateDatastoresReplicaAndOsDisks, true
}

// HasUseSeparateDatastoresReplicaAndOsDisks returns a boolean if a field has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) HasUseSeparateDatastoresReplicaAndOsDisks() bool {
	if o != nil && o.UseSeparateDatastoresReplicaAndOsDisks != nil {
		return true
	}

	return false
}

// SetUseSeparateDatastoresReplicaAndOsDisks gets a reference to the given bool and assigns it to the UseSeparateDatastoresReplicaAndOsDisks field.
func (o *DesktopPoolStorageSettingsUpdateSpec) SetUseSeparateDatastoresReplicaAndOsDisks(v bool) {
	o.UseSeparateDatastoresReplicaAndOsDisks = &v
}

// GetUseVsan returns the UseVsan field value
func (o *DesktopPoolStorageSettingsUpdateSpec) GetUseVsan() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseVsan
}

// GetUseVsanOk returns a tuple with the UseVsan field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolStorageSettingsUpdateSpec) GetUseVsanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseVsan, true
}

// SetUseVsan sets field value
func (o *DesktopPoolStorageSettingsUpdateSpec) SetUseVsan(v bool) {
	o.UseVsan = v
}

func (o DesktopPoolStorageSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["datastores"] = o.Datastores
	}
	if o.ReclaimVmDiskSpace != nil {
		toSerialize["reclaim_vm_disk_space"] = o.ReclaimVmDiskSpace
	}
	if o.ReclamationThresholdMb != nil {
		toSerialize["reclamation_threshold_mb"] = o.ReclamationThresholdMb
	}
	if o.ReplicaDiskDatastoreId != nil {
		toSerialize["replica_disk_datastore_id"] = o.ReplicaDiskDatastoreId
	}
	if o.UseSeparateDatastoresReplicaAndOsDisks != nil {
		toSerialize["use_separate_datastores_replica_and_os_disks"] = o.UseSeparateDatastoresReplicaAndOsDisks
	}
	if true {
		toSerialize["use_vsan"] = o.UseVsan
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolStorageSettingsUpdateSpec struct {
	value *DesktopPoolStorageSettingsUpdateSpec
	isSet bool
}

func (v NullableDesktopPoolStorageSettingsUpdateSpec) Get() *DesktopPoolStorageSettingsUpdateSpec {
	return v.value
}

func (v *NullableDesktopPoolStorageSettingsUpdateSpec) Set(val *DesktopPoolStorageSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolStorageSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolStorageSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolStorageSettingsUpdateSpec(val *DesktopPoolStorageSettingsUpdateSpec) *NullableDesktopPoolStorageSettingsUpdateSpec {
	return &NullableDesktopPoolStorageSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableDesktopPoolStorageSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolStorageSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
