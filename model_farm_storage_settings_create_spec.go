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

// FarmStorageSettingsCreateSpec Virtual center storage settings for the farm.
type FarmStorageSettingsCreateSpec struct {
	// List of IDs of the datastore used to store the RDS Server.
	Datastores []FarmDatastoreSettingsCreateSpec `json:"datastores"`
	// Datastore to store replica disks for instant clone machines. This is required if use_separate_datastores_replica_and_os_disks is set to true.
	ReplicaDiskDatastoreId *string `json:"replica_disk_datastore_id,omitempty"`
	// Indicates whether to use separate datastores for replica and OS disks. Default value is false.
	UseSeparateDatastoresReplicaAndOsDisks *bool `json:"use_separate_datastores_replica_and_os_disks,omitempty"`
	// Indicates whether to use view storage accelerator. Default value is false.
	UseViewStorageAccelerator *bool `json:"use_view_storage_accelerator,omitempty"`
	// Indicates whether to use vSphere VSAN. Default value is false.
	UseVsan *bool `json:"use_vsan,omitempty"`
}

// NewFarmStorageSettingsCreateSpec instantiates a new FarmStorageSettingsCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmStorageSettingsCreateSpec(datastores []FarmDatastoreSettingsCreateSpec) *FarmStorageSettingsCreateSpec {
	this := FarmStorageSettingsCreateSpec{}
	this.Datastores = datastores
	return &this
}

// NewFarmStorageSettingsCreateSpecWithDefaults instantiates a new FarmStorageSettingsCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmStorageSettingsCreateSpecWithDefaults() *FarmStorageSettingsCreateSpec {
	this := FarmStorageSettingsCreateSpec{}
	return &this
}

// GetDatastores returns the Datastores field value
func (o *FarmStorageSettingsCreateSpec) GetDatastores() []FarmDatastoreSettingsCreateSpec {
	if o == nil {
		var ret []FarmDatastoreSettingsCreateSpec
		return ret
	}

	return o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value
// and a boolean to check if the value has been set.
func (o *FarmStorageSettingsCreateSpec) GetDatastoresOk() (*[]FarmDatastoreSettingsCreateSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Datastores, true
}

// SetDatastores sets field value
func (o *FarmStorageSettingsCreateSpec) SetDatastores(v []FarmDatastoreSettingsCreateSpec) {
	o.Datastores = v
}

// GetReplicaDiskDatastoreId returns the ReplicaDiskDatastoreId field value if set, zero value otherwise.
func (o *FarmStorageSettingsCreateSpec) GetReplicaDiskDatastoreId() string {
	if o == nil || o.ReplicaDiskDatastoreId == nil {
		var ret string
		return ret
	}
	return *o.ReplicaDiskDatastoreId
}

// GetReplicaDiskDatastoreIdOk returns a tuple with the ReplicaDiskDatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmStorageSettingsCreateSpec) GetReplicaDiskDatastoreIdOk() (*string, bool) {
	if o == nil || o.ReplicaDiskDatastoreId == nil {
		return nil, false
	}
	return o.ReplicaDiskDatastoreId, true
}

// HasReplicaDiskDatastoreId returns a boolean if a field has been set.
func (o *FarmStorageSettingsCreateSpec) HasReplicaDiskDatastoreId() bool {
	if o != nil && o.ReplicaDiskDatastoreId != nil {
		return true
	}

	return false
}

// SetReplicaDiskDatastoreId gets a reference to the given string and assigns it to the ReplicaDiskDatastoreId field.
func (o *FarmStorageSettingsCreateSpec) SetReplicaDiskDatastoreId(v string) {
	o.ReplicaDiskDatastoreId = &v
}

// GetUseSeparateDatastoresReplicaAndOsDisks returns the UseSeparateDatastoresReplicaAndOsDisks field value if set, zero value otherwise.
func (o *FarmStorageSettingsCreateSpec) GetUseSeparateDatastoresReplicaAndOsDisks() bool {
	if o == nil || o.UseSeparateDatastoresReplicaAndOsDisks == nil {
		var ret bool
		return ret
	}
	return *o.UseSeparateDatastoresReplicaAndOsDisks
}

// GetUseSeparateDatastoresReplicaAndOsDisksOk returns a tuple with the UseSeparateDatastoresReplicaAndOsDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmStorageSettingsCreateSpec) GetUseSeparateDatastoresReplicaAndOsDisksOk() (*bool, bool) {
	if o == nil || o.UseSeparateDatastoresReplicaAndOsDisks == nil {
		return nil, false
	}
	return o.UseSeparateDatastoresReplicaAndOsDisks, true
}

// HasUseSeparateDatastoresReplicaAndOsDisks returns a boolean if a field has been set.
func (o *FarmStorageSettingsCreateSpec) HasUseSeparateDatastoresReplicaAndOsDisks() bool {
	if o != nil && o.UseSeparateDatastoresReplicaAndOsDisks != nil {
		return true
	}

	return false
}

// SetUseSeparateDatastoresReplicaAndOsDisks gets a reference to the given bool and assigns it to the UseSeparateDatastoresReplicaAndOsDisks field.
func (o *FarmStorageSettingsCreateSpec) SetUseSeparateDatastoresReplicaAndOsDisks(v bool) {
	o.UseSeparateDatastoresReplicaAndOsDisks = &v
}

// GetUseViewStorageAccelerator returns the UseViewStorageAccelerator field value if set, zero value otherwise.
func (o *FarmStorageSettingsCreateSpec) GetUseViewStorageAccelerator() bool {
	if o == nil || o.UseViewStorageAccelerator == nil {
		var ret bool
		return ret
	}
	return *o.UseViewStorageAccelerator
}

// GetUseViewStorageAcceleratorOk returns a tuple with the UseViewStorageAccelerator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmStorageSettingsCreateSpec) GetUseViewStorageAcceleratorOk() (*bool, bool) {
	if o == nil || o.UseViewStorageAccelerator == nil {
		return nil, false
	}
	return o.UseViewStorageAccelerator, true
}

// HasUseViewStorageAccelerator returns a boolean if a field has been set.
func (o *FarmStorageSettingsCreateSpec) HasUseViewStorageAccelerator() bool {
	if o != nil && o.UseViewStorageAccelerator != nil {
		return true
	}

	return false
}

// SetUseViewStorageAccelerator gets a reference to the given bool and assigns it to the UseViewStorageAccelerator field.
func (o *FarmStorageSettingsCreateSpec) SetUseViewStorageAccelerator(v bool) {
	o.UseViewStorageAccelerator = &v
}

// GetUseVsan returns the UseVsan field value if set, zero value otherwise.
func (o *FarmStorageSettingsCreateSpec) GetUseVsan() bool {
	if o == nil || o.UseVsan == nil {
		var ret bool
		return ret
	}
	return *o.UseVsan
}

// GetUseVsanOk returns a tuple with the UseVsan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FarmStorageSettingsCreateSpec) GetUseVsanOk() (*bool, bool) {
	if o == nil || o.UseVsan == nil {
		return nil, false
	}
	return o.UseVsan, true
}

// HasUseVsan returns a boolean if a field has been set.
func (o *FarmStorageSettingsCreateSpec) HasUseVsan() bool {
	if o != nil && o.UseVsan != nil {
		return true
	}

	return false
}

// SetUseVsan gets a reference to the given bool and assigns it to the UseVsan field.
func (o *FarmStorageSettingsCreateSpec) SetUseVsan(v bool) {
	o.UseVsan = &v
}

func (o FarmStorageSettingsCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["datastores"] = o.Datastores
	}
	if o.ReplicaDiskDatastoreId != nil {
		toSerialize["replica_disk_datastore_id"] = o.ReplicaDiskDatastoreId
	}
	if o.UseSeparateDatastoresReplicaAndOsDisks != nil {
		toSerialize["use_separate_datastores_replica_and_os_disks"] = o.UseSeparateDatastoresReplicaAndOsDisks
	}
	if o.UseViewStorageAccelerator != nil {
		toSerialize["use_view_storage_accelerator"] = o.UseViewStorageAccelerator
	}
	if o.UseVsan != nil {
		toSerialize["use_vsan"] = o.UseVsan
	}
	return json.Marshal(toSerialize)
}

type NullableFarmStorageSettingsCreateSpec struct {
	value *FarmStorageSettingsCreateSpec
	isSet bool
}

func (v NullableFarmStorageSettingsCreateSpec) Get() *FarmStorageSettingsCreateSpec {
	return v.value
}

func (v *NullableFarmStorageSettingsCreateSpec) Set(val *FarmStorageSettingsCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmStorageSettingsCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmStorageSettingsCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmStorageSettingsCreateSpec(val *FarmStorageSettingsCreateSpec) *NullableFarmStorageSettingsCreateSpec {
	return &NullableFarmStorageSettingsCreateSpec{value: val, isSet: true}
}

func (v NullableFarmStorageSettingsCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmStorageSettingsCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
