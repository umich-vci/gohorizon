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

// DatastoreSpaceRequirementSpec Information required to compute Datastore Space Requirement.
type DatastoreSpaceRequirementSpec struct {
	// Parent VM snapshot ID. Must be set if source is INSTANT_CLONE.
	BaseSnapshotId *string `json:"base_snapshot_id,omitempty"`
	// Parent VM ID. Must be set if source is INSTANT_CLONE.
	BaseVmId *string `json:"base_vm_id,omitempty"`
	// Id of inventory resource for which space requirement is to be found. Can be desktop pool or farm id.
	Id *string `json:"id,omitempty"`
	// Desired size of the desktop pool or farm.
	PoolSize int32 `json:"pool_size"`
	// Source or provisioning type of machines. * FULL_CLONE: Virtual Machines created from a vCenter Server template. * INSTANT_CLONE: Virtual Machines created by instant clone engine.
	Source string `json:"source"`
	// Type of inventory resource for which space requirement is to be found. * DESKTOP_POOL: Desktop pool inventory resource. * FARM: Farm inventory resource.
	Type string `json:"type"`
	// Indicates whether separate datastores are to be used for OS and replica disks. Will be ignored if source is FULL_CLONE or vSAN is to be configured. Default value is false.
	UseSeparateReplicaAndOsDisk *bool `json:"use_separate_replica_and_os_disk,omitempty"`
	// Indicates whether vSAN is to be configured for the desktop pool or farm. Default value is false. vSAN should be configured if set to true.
	UseVsan *bool `json:"use_vsan,omitempty"`
	// User assignment of the desktop pool. Will be ignored if type is FARM. Default value is FLOATING. * DEDICATED: Dedicated user assignment. * FLOATING: Floating user assignment.
	UserAssignment *string `json:"user_assignment,omitempty"`
	// ID of virtual center where parent VM or master image is present.
	VcenterId string `json:"vcenter_id"`
	// VM template ID. Must be set if source is FULL_CLONE.
	VmTemplateId *string `json:"vm_template_id,omitempty"`
}

// NewDatastoreSpaceRequirementSpec instantiates a new DatastoreSpaceRequirementSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatastoreSpaceRequirementSpec(poolSize int32, source string, type_ string, vcenterId string) *DatastoreSpaceRequirementSpec {
	this := DatastoreSpaceRequirementSpec{}
	this.PoolSize = poolSize
	this.Source = source
	this.Type = type_
	this.VcenterId = vcenterId
	return &this
}

// NewDatastoreSpaceRequirementSpecWithDefaults instantiates a new DatastoreSpaceRequirementSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatastoreSpaceRequirementSpecWithDefaults() *DatastoreSpaceRequirementSpec {
	this := DatastoreSpaceRequirementSpec{}
	return &this
}

// GetBaseSnapshotId returns the BaseSnapshotId field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementSpec) GetBaseSnapshotId() string {
	if o == nil || o.BaseSnapshotId == nil {
		var ret string
		return ret
	}
	return *o.BaseSnapshotId
}

// GetBaseSnapshotIdOk returns a tuple with the BaseSnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetBaseSnapshotIdOk() (*string, bool) {
	if o == nil || o.BaseSnapshotId == nil {
		return nil, false
	}
	return o.BaseSnapshotId, true
}

// HasBaseSnapshotId returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementSpec) HasBaseSnapshotId() bool {
	if o != nil && o.BaseSnapshotId != nil {
		return true
	}

	return false
}

// SetBaseSnapshotId gets a reference to the given string and assigns it to the BaseSnapshotId field.
func (o *DatastoreSpaceRequirementSpec) SetBaseSnapshotId(v string) {
	o.BaseSnapshotId = &v
}

// GetBaseVmId returns the BaseVmId field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementSpec) GetBaseVmId() string {
	if o == nil || o.BaseVmId == nil {
		var ret string
		return ret
	}
	return *o.BaseVmId
}

// GetBaseVmIdOk returns a tuple with the BaseVmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetBaseVmIdOk() (*string, bool) {
	if o == nil || o.BaseVmId == nil {
		return nil, false
	}
	return o.BaseVmId, true
}

// HasBaseVmId returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementSpec) HasBaseVmId() bool {
	if o != nil && o.BaseVmId != nil {
		return true
	}

	return false
}

// SetBaseVmId gets a reference to the given string and assigns it to the BaseVmId field.
func (o *DatastoreSpaceRequirementSpec) SetBaseVmId(v string) {
	o.BaseVmId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementSpec) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementSpec) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DatastoreSpaceRequirementSpec) SetId(v string) {
	o.Id = &v
}

// GetPoolSize returns the PoolSize field value
func (o *DatastoreSpaceRequirementSpec) GetPoolSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PoolSize
}

// GetPoolSizeOk returns a tuple with the PoolSize field value
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetPoolSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PoolSize, true
}

// SetPoolSize sets field value
func (o *DatastoreSpaceRequirementSpec) SetPoolSize(v int32) {
	o.PoolSize = v
}

// GetSource returns the Source field value
func (o *DatastoreSpaceRequirementSpec) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DatastoreSpaceRequirementSpec) SetSource(v string) {
	o.Source = v
}

// GetType returns the Type field value
func (o *DatastoreSpaceRequirementSpec) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DatastoreSpaceRequirementSpec) SetType(v string) {
	o.Type = v
}

// GetUseSeparateReplicaAndOsDisk returns the UseSeparateReplicaAndOsDisk field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementSpec) GetUseSeparateReplicaAndOsDisk() bool {
	if o == nil || o.UseSeparateReplicaAndOsDisk == nil {
		var ret bool
		return ret
	}
	return *o.UseSeparateReplicaAndOsDisk
}

// GetUseSeparateReplicaAndOsDiskOk returns a tuple with the UseSeparateReplicaAndOsDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetUseSeparateReplicaAndOsDiskOk() (*bool, bool) {
	if o == nil || o.UseSeparateReplicaAndOsDisk == nil {
		return nil, false
	}
	return o.UseSeparateReplicaAndOsDisk, true
}

// HasUseSeparateReplicaAndOsDisk returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementSpec) HasUseSeparateReplicaAndOsDisk() bool {
	if o != nil && o.UseSeparateReplicaAndOsDisk != nil {
		return true
	}

	return false
}

// SetUseSeparateReplicaAndOsDisk gets a reference to the given bool and assigns it to the UseSeparateReplicaAndOsDisk field.
func (o *DatastoreSpaceRequirementSpec) SetUseSeparateReplicaAndOsDisk(v bool) {
	o.UseSeparateReplicaAndOsDisk = &v
}

// GetUseVsan returns the UseVsan field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementSpec) GetUseVsan() bool {
	if o == nil || o.UseVsan == nil {
		var ret bool
		return ret
	}
	return *o.UseVsan
}

// GetUseVsanOk returns a tuple with the UseVsan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetUseVsanOk() (*bool, bool) {
	if o == nil || o.UseVsan == nil {
		return nil, false
	}
	return o.UseVsan, true
}

// HasUseVsan returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementSpec) HasUseVsan() bool {
	if o != nil && o.UseVsan != nil {
		return true
	}

	return false
}

// SetUseVsan gets a reference to the given bool and assigns it to the UseVsan field.
func (o *DatastoreSpaceRequirementSpec) SetUseVsan(v bool) {
	o.UseVsan = &v
}

// GetUserAssignment returns the UserAssignment field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementSpec) GetUserAssignment() string {
	if o == nil || o.UserAssignment == nil {
		var ret string
		return ret
	}
	return *o.UserAssignment
}

// GetUserAssignmentOk returns a tuple with the UserAssignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetUserAssignmentOk() (*string, bool) {
	if o == nil || o.UserAssignment == nil {
		return nil, false
	}
	return o.UserAssignment, true
}

// HasUserAssignment returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementSpec) HasUserAssignment() bool {
	if o != nil && o.UserAssignment != nil {
		return true
	}

	return false
}

// SetUserAssignment gets a reference to the given string and assigns it to the UserAssignment field.
func (o *DatastoreSpaceRequirementSpec) SetUserAssignment(v string) {
	o.UserAssignment = &v
}

// GetVcenterId returns the VcenterId field value
func (o *DatastoreSpaceRequirementSpec) GetVcenterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VcenterId
}

// GetVcenterIdOk returns a tuple with the VcenterId field value
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetVcenterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VcenterId, true
}

// SetVcenterId sets field value
func (o *DatastoreSpaceRequirementSpec) SetVcenterId(v string) {
	o.VcenterId = v
}

// GetVmTemplateId returns the VmTemplateId field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementSpec) GetVmTemplateId() string {
	if o == nil || o.VmTemplateId == nil {
		var ret string
		return ret
	}
	return *o.VmTemplateId
}

// GetVmTemplateIdOk returns a tuple with the VmTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementSpec) GetVmTemplateIdOk() (*string, bool) {
	if o == nil || o.VmTemplateId == nil {
		return nil, false
	}
	return o.VmTemplateId, true
}

// HasVmTemplateId returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementSpec) HasVmTemplateId() bool {
	if o != nil && o.VmTemplateId != nil {
		return true
	}

	return false
}

// SetVmTemplateId gets a reference to the given string and assigns it to the VmTemplateId field.
func (o *DatastoreSpaceRequirementSpec) SetVmTemplateId(v string) {
	o.VmTemplateId = &v
}

func (o DatastoreSpaceRequirementSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseSnapshotId != nil {
		toSerialize["base_snapshot_id"] = o.BaseSnapshotId
	}
	if o.BaseVmId != nil {
		toSerialize["base_vm_id"] = o.BaseVmId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["pool_size"] = o.PoolSize
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UseSeparateReplicaAndOsDisk != nil {
		toSerialize["use_separate_replica_and_os_disk"] = o.UseSeparateReplicaAndOsDisk
	}
	if o.UseVsan != nil {
		toSerialize["use_vsan"] = o.UseVsan
	}
	if o.UserAssignment != nil {
		toSerialize["user_assignment"] = o.UserAssignment
	}
	if true {
		toSerialize["vcenter_id"] = o.VcenterId
	}
	if o.VmTemplateId != nil {
		toSerialize["vm_template_id"] = o.VmTemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableDatastoreSpaceRequirementSpec struct {
	value *DatastoreSpaceRequirementSpec
	isSet bool
}

func (v NullableDatastoreSpaceRequirementSpec) Get() *DatastoreSpaceRequirementSpec {
	return v.value
}

func (v *NullableDatastoreSpaceRequirementSpec) Set(val *DatastoreSpaceRequirementSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreSpaceRequirementSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreSpaceRequirementSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreSpaceRequirementSpec(val *DatastoreSpaceRequirementSpec) *NullableDatastoreSpaceRequirementSpec {
	return &NullableDatastoreSpaceRequirementSpec{value: val, isSet: true}
}

func (v NullableDatastoreSpaceRequirementSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreSpaceRequirementSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
