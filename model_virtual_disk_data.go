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

// VirtualDiskData Information related to a virtual disk.
type VirtualDiskData struct {
	// The disk capacity in MB.
	CapacityMb *int64 `json:"capacity_mb,omitempty"`
	// The virtual disk's datastore.
	DatastorePath *string `json:"datastore_path,omitempty"`
	// The virtual disk path.
	Path *string `json:"path,omitempty"`
}

// NewVirtualDiskData instantiates a new VirtualDiskData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualDiskData() *VirtualDiskData {
	this := VirtualDiskData{}
	return &this
}

// NewVirtualDiskDataWithDefaults instantiates a new VirtualDiskData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualDiskDataWithDefaults() *VirtualDiskData {
	this := VirtualDiskData{}
	return &this
}

// GetCapacityMb returns the CapacityMb field value if set, zero value otherwise.
func (o *VirtualDiskData) GetCapacityMb() int64 {
	if o == nil || o.CapacityMb == nil {
		var ret int64
		return ret
	}
	return *o.CapacityMb
}

// GetCapacityMbOk returns a tuple with the CapacityMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDiskData) GetCapacityMbOk() (*int64, bool) {
	if o == nil || o.CapacityMb == nil {
		return nil, false
	}
	return o.CapacityMb, true
}

// HasCapacityMb returns a boolean if a field has been set.
func (o *VirtualDiskData) HasCapacityMb() bool {
	if o != nil && o.CapacityMb != nil {
		return true
	}

	return false
}

// SetCapacityMb gets a reference to the given int64 and assigns it to the CapacityMb field.
func (o *VirtualDiskData) SetCapacityMb(v int64) {
	o.CapacityMb = &v
}

// GetDatastorePath returns the DatastorePath field value if set, zero value otherwise.
func (o *VirtualDiskData) GetDatastorePath() string {
	if o == nil || o.DatastorePath == nil {
		var ret string
		return ret
	}
	return *o.DatastorePath
}

// GetDatastorePathOk returns a tuple with the DatastorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDiskData) GetDatastorePathOk() (*string, bool) {
	if o == nil || o.DatastorePath == nil {
		return nil, false
	}
	return o.DatastorePath, true
}

// HasDatastorePath returns a boolean if a field has been set.
func (o *VirtualDiskData) HasDatastorePath() bool {
	if o != nil && o.DatastorePath != nil {
		return true
	}

	return false
}

// SetDatastorePath gets a reference to the given string and assigns it to the DatastorePath field.
func (o *VirtualDiskData) SetDatastorePath(v string) {
	o.DatastorePath = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *VirtualDiskData) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualDiskData) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *VirtualDiskData) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *VirtualDiskData) SetPath(v string) {
	o.Path = &v
}

func (o VirtualDiskData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapacityMb != nil {
		toSerialize["capacity_mb"] = o.CapacityMb
	}
	if o.DatastorePath != nil {
		toSerialize["datastore_path"] = o.DatastorePath
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualDiskData struct {
	value *VirtualDiskData
	isSet bool
}

func (v NullableVirtualDiskData) Get() *VirtualDiskData {
	return v.value
}

func (v *NullableVirtualDiskData) Set(val *VirtualDiskData) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualDiskData) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualDiskData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualDiskData(val *VirtualDiskData) *NullableVirtualDiskData {
	return &NullableVirtualDiskData{value: val, isSet: true}
}

func (v NullableVirtualDiskData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualDiskData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
