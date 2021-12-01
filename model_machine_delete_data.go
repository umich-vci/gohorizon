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

// MachineDeleteData Specification applicable when deleting machines.
type MachineDeleteData struct {
	// Determines whether the machines from different desktop pools can be deleted. This defaults to false in which case only machines belonging to single desktop pool can be deleted. If true, machines from different desktop pools can be deleted.
	AllowDeleteFromMultiDesktopPools *bool `json:"allow_delete_from_multi_desktop_pools,omitempty"`
	// Determines the datastore where the persistent user disk will be saved for future use. Both this as well as the archiveDatastorePathId need to be set. If this is unset and archivePersistentDisk is specified, the persistent disk is archived in place.
	ArchiveDatastoreId *string `json:"archive_datastore_id,omitempty"`
	// Determines the location in the datastore where the persistent user disk will be saved for future use. If this is set, then archiveDatastoreId also needs to be specified.If this is unset and archivePersistentDisk is specified, the persistent disk is archived in place.
	ArchiveDatastorePathId *string `json:"archive_datastore_path_id,omitempty"`
	// Determines whether to detach the persistent user disk and save it for future use. This can only be specified for linked-clone desktops with redirectWindowsProfile enabled, in which case it defaults to true.
	ArchivePersistentDisk *bool `json:"archive_persistent_disk,omitempty"`
	// Determines whether the Machine VM should be deleted from vCenter Server. This is only applicable for managed machines. This must always be true for machines in linked and instant clone desktops. This defaults to true for linked and instant clone machines and false for all other types. If this is true, then machine being deleted must not have any active user session, otherwise delete operation would fail.
	DeleteFromDisk *bool `json:"delete_from_disk,omitempty"`
	// Determines whether active session on the machine to be logged off before deletion. This is only applicable for managed machines. If true, active session on the machine will be logged off before Machine delete. Otherwise,it will result in an exception.
	ForceLogoffSession *bool `json:"force_logoff_session,omitempty"`
}

// NewMachineDeleteData instantiates a new MachineDeleteData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMachineDeleteData() *MachineDeleteData {
	this := MachineDeleteData{}
	return &this
}

// NewMachineDeleteDataWithDefaults instantiates a new MachineDeleteData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMachineDeleteDataWithDefaults() *MachineDeleteData {
	this := MachineDeleteData{}
	return &this
}

// GetAllowDeleteFromMultiDesktopPools returns the AllowDeleteFromMultiDesktopPools field value if set, zero value otherwise.
func (o *MachineDeleteData) GetAllowDeleteFromMultiDesktopPools() bool {
	if o == nil || o.AllowDeleteFromMultiDesktopPools == nil {
		var ret bool
		return ret
	}
	return *o.AllowDeleteFromMultiDesktopPools
}

// GetAllowDeleteFromMultiDesktopPoolsOk returns a tuple with the AllowDeleteFromMultiDesktopPools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDeleteData) GetAllowDeleteFromMultiDesktopPoolsOk() (*bool, bool) {
	if o == nil || o.AllowDeleteFromMultiDesktopPools == nil {
		return nil, false
	}
	return o.AllowDeleteFromMultiDesktopPools, true
}

// HasAllowDeleteFromMultiDesktopPools returns a boolean if a field has been set.
func (o *MachineDeleteData) HasAllowDeleteFromMultiDesktopPools() bool {
	if o != nil && o.AllowDeleteFromMultiDesktopPools != nil {
		return true
	}

	return false
}

// SetAllowDeleteFromMultiDesktopPools gets a reference to the given bool and assigns it to the AllowDeleteFromMultiDesktopPools field.
func (o *MachineDeleteData) SetAllowDeleteFromMultiDesktopPools(v bool) {
	o.AllowDeleteFromMultiDesktopPools = &v
}

// GetArchiveDatastoreId returns the ArchiveDatastoreId field value if set, zero value otherwise.
func (o *MachineDeleteData) GetArchiveDatastoreId() string {
	if o == nil || o.ArchiveDatastoreId == nil {
		var ret string
		return ret
	}
	return *o.ArchiveDatastoreId
}

// GetArchiveDatastoreIdOk returns a tuple with the ArchiveDatastoreId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDeleteData) GetArchiveDatastoreIdOk() (*string, bool) {
	if o == nil || o.ArchiveDatastoreId == nil {
		return nil, false
	}
	return o.ArchiveDatastoreId, true
}

// HasArchiveDatastoreId returns a boolean if a field has been set.
func (o *MachineDeleteData) HasArchiveDatastoreId() bool {
	if o != nil && o.ArchiveDatastoreId != nil {
		return true
	}

	return false
}

// SetArchiveDatastoreId gets a reference to the given string and assigns it to the ArchiveDatastoreId field.
func (o *MachineDeleteData) SetArchiveDatastoreId(v string) {
	o.ArchiveDatastoreId = &v
}

// GetArchiveDatastorePathId returns the ArchiveDatastorePathId field value if set, zero value otherwise.
func (o *MachineDeleteData) GetArchiveDatastorePathId() string {
	if o == nil || o.ArchiveDatastorePathId == nil {
		var ret string
		return ret
	}
	return *o.ArchiveDatastorePathId
}

// GetArchiveDatastorePathIdOk returns a tuple with the ArchiveDatastorePathId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDeleteData) GetArchiveDatastorePathIdOk() (*string, bool) {
	if o == nil || o.ArchiveDatastorePathId == nil {
		return nil, false
	}
	return o.ArchiveDatastorePathId, true
}

// HasArchiveDatastorePathId returns a boolean if a field has been set.
func (o *MachineDeleteData) HasArchiveDatastorePathId() bool {
	if o != nil && o.ArchiveDatastorePathId != nil {
		return true
	}

	return false
}

// SetArchiveDatastorePathId gets a reference to the given string and assigns it to the ArchiveDatastorePathId field.
func (o *MachineDeleteData) SetArchiveDatastorePathId(v string) {
	o.ArchiveDatastorePathId = &v
}

// GetArchivePersistentDisk returns the ArchivePersistentDisk field value if set, zero value otherwise.
func (o *MachineDeleteData) GetArchivePersistentDisk() bool {
	if o == nil || o.ArchivePersistentDisk == nil {
		var ret bool
		return ret
	}
	return *o.ArchivePersistentDisk
}

// GetArchivePersistentDiskOk returns a tuple with the ArchivePersistentDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDeleteData) GetArchivePersistentDiskOk() (*bool, bool) {
	if o == nil || o.ArchivePersistentDisk == nil {
		return nil, false
	}
	return o.ArchivePersistentDisk, true
}

// HasArchivePersistentDisk returns a boolean if a field has been set.
func (o *MachineDeleteData) HasArchivePersistentDisk() bool {
	if o != nil && o.ArchivePersistentDisk != nil {
		return true
	}

	return false
}

// SetArchivePersistentDisk gets a reference to the given bool and assigns it to the ArchivePersistentDisk field.
func (o *MachineDeleteData) SetArchivePersistentDisk(v bool) {
	o.ArchivePersistentDisk = &v
}

// GetDeleteFromDisk returns the DeleteFromDisk field value if set, zero value otherwise.
func (o *MachineDeleteData) GetDeleteFromDisk() bool {
	if o == nil || o.DeleteFromDisk == nil {
		var ret bool
		return ret
	}
	return *o.DeleteFromDisk
}

// GetDeleteFromDiskOk returns a tuple with the DeleteFromDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDeleteData) GetDeleteFromDiskOk() (*bool, bool) {
	if o == nil || o.DeleteFromDisk == nil {
		return nil, false
	}
	return o.DeleteFromDisk, true
}

// HasDeleteFromDisk returns a boolean if a field has been set.
func (o *MachineDeleteData) HasDeleteFromDisk() bool {
	if o != nil && o.DeleteFromDisk != nil {
		return true
	}

	return false
}

// SetDeleteFromDisk gets a reference to the given bool and assigns it to the DeleteFromDisk field.
func (o *MachineDeleteData) SetDeleteFromDisk(v bool) {
	o.DeleteFromDisk = &v
}

// GetForceLogoffSession returns the ForceLogoffSession field value if set, zero value otherwise.
func (o *MachineDeleteData) GetForceLogoffSession() bool {
	if o == nil || o.ForceLogoffSession == nil {
		var ret bool
		return ret
	}
	return *o.ForceLogoffSession
}

// GetForceLogoffSessionOk returns a tuple with the ForceLogoffSession field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MachineDeleteData) GetForceLogoffSessionOk() (*bool, bool) {
	if o == nil || o.ForceLogoffSession == nil {
		return nil, false
	}
	return o.ForceLogoffSession, true
}

// HasForceLogoffSession returns a boolean if a field has been set.
func (o *MachineDeleteData) HasForceLogoffSession() bool {
	if o != nil && o.ForceLogoffSession != nil {
		return true
	}

	return false
}

// SetForceLogoffSession gets a reference to the given bool and assigns it to the ForceLogoffSession field.
func (o *MachineDeleteData) SetForceLogoffSession(v bool) {
	o.ForceLogoffSession = &v
}

func (o MachineDeleteData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowDeleteFromMultiDesktopPools != nil {
		toSerialize["allow_delete_from_multi_desktop_pools"] = o.AllowDeleteFromMultiDesktopPools
	}
	if o.ArchiveDatastoreId != nil {
		toSerialize["archive_datastore_id"] = o.ArchiveDatastoreId
	}
	if o.ArchiveDatastorePathId != nil {
		toSerialize["archive_datastore_path_id"] = o.ArchiveDatastorePathId
	}
	if o.ArchivePersistentDisk != nil {
		toSerialize["archive_persistent_disk"] = o.ArchivePersistentDisk
	}
	if o.DeleteFromDisk != nil {
		toSerialize["delete_from_disk"] = o.DeleteFromDisk
	}
	if o.ForceLogoffSession != nil {
		toSerialize["force_logoff_session"] = o.ForceLogoffSession
	}
	return json.Marshal(toSerialize)
}

type NullableMachineDeleteData struct {
	value *MachineDeleteData
	isSet bool
}

func (v NullableMachineDeleteData) Get() *MachineDeleteData {
	return v.value
}

func (v *NullableMachineDeleteData) Set(val *MachineDeleteData) {
	v.value = val
	v.isSet = true
}

func (v NullableMachineDeleteData) IsSet() bool {
	return v.isSet
}

func (v *NullableMachineDeleteData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMachineDeleteData(val *MachineDeleteData) *NullableMachineDeleteData {
	return &NullableMachineDeleteData{value: val, isSet: true}
}

func (v NullableMachineDeleteData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMachineDeleteData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
