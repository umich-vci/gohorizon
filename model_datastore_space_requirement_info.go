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

// DatastoreSpaceRequirementInfo Information about Datastore Space Requirement.
type DatastoreSpaceRequirementInfo struct {
	// Indicates the type of disk used for storage. * OS: Disk to store operating system related data. * REPLICA: Disk for placement of replica VMs created by instant clone engine.
	DiskType *string `json:"disk_type,omitempty"`
	// Indicates maximum recommended disk space, in GB.
	MaxSizeDiskGb *float64 `json:"max_size_disk_gb,omitempty"`
	// Indicates recommended disk space with 50% utilization, in GB.
	MidSizeDiskGb *float64 `json:"mid_size_disk_gb,omitempty"`
	// Indicates minimum recommended disk space, in GB.
	MinSizeDiskGb *float64 `json:"min_size_disk_gb,omitempty"`
}

// NewDatastoreSpaceRequirementInfo instantiates a new DatastoreSpaceRequirementInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatastoreSpaceRequirementInfo() *DatastoreSpaceRequirementInfo {
	this := DatastoreSpaceRequirementInfo{}
	return &this
}

// NewDatastoreSpaceRequirementInfoWithDefaults instantiates a new DatastoreSpaceRequirementInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatastoreSpaceRequirementInfoWithDefaults() *DatastoreSpaceRequirementInfo {
	this := DatastoreSpaceRequirementInfo{}
	return &this
}

// GetDiskType returns the DiskType field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementInfo) GetDiskType() string {
	if o == nil || o.DiskType == nil {
		var ret string
		return ret
	}
	return *o.DiskType
}

// GetDiskTypeOk returns a tuple with the DiskType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementInfo) GetDiskTypeOk() (*string, bool) {
	if o == nil || o.DiskType == nil {
		return nil, false
	}
	return o.DiskType, true
}

// HasDiskType returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementInfo) HasDiskType() bool {
	if o != nil && o.DiskType != nil {
		return true
	}

	return false
}

// SetDiskType gets a reference to the given string and assigns it to the DiskType field.
func (o *DatastoreSpaceRequirementInfo) SetDiskType(v string) {
	o.DiskType = &v
}

// GetMaxSizeDiskGb returns the MaxSizeDiskGb field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementInfo) GetMaxSizeDiskGb() float64 {
	if o == nil || o.MaxSizeDiskGb == nil {
		var ret float64
		return ret
	}
	return *o.MaxSizeDiskGb
}

// GetMaxSizeDiskGbOk returns a tuple with the MaxSizeDiskGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementInfo) GetMaxSizeDiskGbOk() (*float64, bool) {
	if o == nil || o.MaxSizeDiskGb == nil {
		return nil, false
	}
	return o.MaxSizeDiskGb, true
}

// HasMaxSizeDiskGb returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementInfo) HasMaxSizeDiskGb() bool {
	if o != nil && o.MaxSizeDiskGb != nil {
		return true
	}

	return false
}

// SetMaxSizeDiskGb gets a reference to the given float64 and assigns it to the MaxSizeDiskGb field.
func (o *DatastoreSpaceRequirementInfo) SetMaxSizeDiskGb(v float64) {
	o.MaxSizeDiskGb = &v
}

// GetMidSizeDiskGb returns the MidSizeDiskGb field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementInfo) GetMidSizeDiskGb() float64 {
	if o == nil || o.MidSizeDiskGb == nil {
		var ret float64
		return ret
	}
	return *o.MidSizeDiskGb
}

// GetMidSizeDiskGbOk returns a tuple with the MidSizeDiskGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementInfo) GetMidSizeDiskGbOk() (*float64, bool) {
	if o == nil || o.MidSizeDiskGb == nil {
		return nil, false
	}
	return o.MidSizeDiskGb, true
}

// HasMidSizeDiskGb returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementInfo) HasMidSizeDiskGb() bool {
	if o != nil && o.MidSizeDiskGb != nil {
		return true
	}

	return false
}

// SetMidSizeDiskGb gets a reference to the given float64 and assigns it to the MidSizeDiskGb field.
func (o *DatastoreSpaceRequirementInfo) SetMidSizeDiskGb(v float64) {
	o.MidSizeDiskGb = &v
}

// GetMinSizeDiskGb returns the MinSizeDiskGb field value if set, zero value otherwise.
func (o *DatastoreSpaceRequirementInfo) GetMinSizeDiskGb() float64 {
	if o == nil || o.MinSizeDiskGb == nil {
		var ret float64
		return ret
	}
	return *o.MinSizeDiskGb
}

// GetMinSizeDiskGbOk returns a tuple with the MinSizeDiskGb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreSpaceRequirementInfo) GetMinSizeDiskGbOk() (*float64, bool) {
	if o == nil || o.MinSizeDiskGb == nil {
		return nil, false
	}
	return o.MinSizeDiskGb, true
}

// HasMinSizeDiskGb returns a boolean if a field has been set.
func (o *DatastoreSpaceRequirementInfo) HasMinSizeDiskGb() bool {
	if o != nil && o.MinSizeDiskGb != nil {
		return true
	}

	return false
}

// SetMinSizeDiskGb gets a reference to the given float64 and assigns it to the MinSizeDiskGb field.
func (o *DatastoreSpaceRequirementInfo) SetMinSizeDiskGb(v float64) {
	o.MinSizeDiskGb = &v
}

func (o DatastoreSpaceRequirementInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DiskType != nil {
		toSerialize["disk_type"] = o.DiskType
	}
	if o.MaxSizeDiskGb != nil {
		toSerialize["max_size_disk_gb"] = o.MaxSizeDiskGb
	}
	if o.MidSizeDiskGb != nil {
		toSerialize["mid_size_disk_gb"] = o.MidSizeDiskGb
	}
	if o.MinSizeDiskGb != nil {
		toSerialize["min_size_disk_gb"] = o.MinSizeDiskGb
	}
	return json.Marshal(toSerialize)
}

type NullableDatastoreSpaceRequirementInfo struct {
	value *DatastoreSpaceRequirementInfo
	isSet bool
}

func (v NullableDatastoreSpaceRequirementInfo) Get() *DatastoreSpaceRequirementInfo {
	return v.value
}

func (v *NullableDatastoreSpaceRequirementInfo) Set(val *DatastoreSpaceRequirementInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreSpaceRequirementInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreSpaceRequirementInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreSpaceRequirementInfo(val *DatastoreSpaceRequirementInfo) *NullableDatastoreSpaceRequirementInfo {
	return &NullableDatastoreSpaceRequirementInfo{value: val, isSet: true}
}

func (v NullableDatastoreSpaceRequirementInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreSpaceRequirementInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
