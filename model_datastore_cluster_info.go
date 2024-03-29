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

// DatastoreClusterInfo struct for DatastoreClusterInfo
type DatastoreClusterInfo struct {
	// Maximum capacity of this datastore cluster, in MB.
	CapacityMb *int64 `json:"capacity_mb,omitempty"`
	// ID of the datacenter for this datastore cluster.
	DatacenterId *string `json:"datacenter_id,omitempty"`
	// IDs of datastores which are a part of this datastore cluster.
	DatastoreIds *[]string `json:"datastore_ids,omitempty"`
	// Available capacity of this datastore cluster, in MB.
	FreeSpaceMb *int64 `json:"free_space_mb,omitempty"`
	// ID of the host or cluster for this datastore cluster.
	HostOrClusterId *string `json:"host_or_cluster_id,omitempty"`
	// Unique ID representing this datastore cluster.
	Id *string `json:"id,omitempty"`
	// Datastore cluster name.
	Name *string `json:"name,omitempty"`
	// Datastore cluster path.
	Path *string `json:"path,omitempty"`
	// ID of the virtual center for this datastore cluster.
	VcenterId *string `json:"vcenter_id,omitempty"`
}

// NewDatastoreClusterInfo instantiates a new DatastoreClusterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatastoreClusterInfo() *DatastoreClusterInfo {
	this := DatastoreClusterInfo{}
	return &this
}

// NewDatastoreClusterInfoWithDefaults instantiates a new DatastoreClusterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatastoreClusterInfoWithDefaults() *DatastoreClusterInfo {
	this := DatastoreClusterInfo{}
	return &this
}

// GetCapacityMb returns the CapacityMb field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetCapacityMb() int64 {
	if o == nil || o.CapacityMb == nil {
		var ret int64
		return ret
	}
	return *o.CapacityMb
}

// GetCapacityMbOk returns a tuple with the CapacityMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetCapacityMbOk() (*int64, bool) {
	if o == nil || o.CapacityMb == nil {
		return nil, false
	}
	return o.CapacityMb, true
}

// HasCapacityMb returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasCapacityMb() bool {
	if o != nil && o.CapacityMb != nil {
		return true
	}

	return false
}

// SetCapacityMb gets a reference to the given int64 and assigns it to the CapacityMb field.
func (o *DatastoreClusterInfo) SetCapacityMb(v int64) {
	o.CapacityMb = &v
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetDatacenterId() string {
	if o == nil || o.DatacenterId == nil {
		var ret string
		return ret
	}
	return *o.DatacenterId
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetDatacenterIdOk() (*string, bool) {
	if o == nil || o.DatacenterId == nil {
		return nil, false
	}
	return o.DatacenterId, true
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given string and assigns it to the DatacenterId field.
func (o *DatastoreClusterInfo) SetDatacenterId(v string) {
	o.DatacenterId = &v
}

// GetDatastoreIds returns the DatastoreIds field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetDatastoreIds() []string {
	if o == nil || o.DatastoreIds == nil {
		var ret []string
		return ret
	}
	return *o.DatastoreIds
}

// GetDatastoreIdsOk returns a tuple with the DatastoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetDatastoreIdsOk() (*[]string, bool) {
	if o == nil || o.DatastoreIds == nil {
		return nil, false
	}
	return o.DatastoreIds, true
}

// HasDatastoreIds returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasDatastoreIds() bool {
	if o != nil && o.DatastoreIds != nil {
		return true
	}

	return false
}

// SetDatastoreIds gets a reference to the given []string and assigns it to the DatastoreIds field.
func (o *DatastoreClusterInfo) SetDatastoreIds(v []string) {
	o.DatastoreIds = &v
}

// GetFreeSpaceMb returns the FreeSpaceMb field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetFreeSpaceMb() int64 {
	if o == nil || o.FreeSpaceMb == nil {
		var ret int64
		return ret
	}
	return *o.FreeSpaceMb
}

// GetFreeSpaceMbOk returns a tuple with the FreeSpaceMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetFreeSpaceMbOk() (*int64, bool) {
	if o == nil || o.FreeSpaceMb == nil {
		return nil, false
	}
	return o.FreeSpaceMb, true
}

// HasFreeSpaceMb returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasFreeSpaceMb() bool {
	if o != nil && o.FreeSpaceMb != nil {
		return true
	}

	return false
}

// SetFreeSpaceMb gets a reference to the given int64 and assigns it to the FreeSpaceMb field.
func (o *DatastoreClusterInfo) SetFreeSpaceMb(v int64) {
	o.FreeSpaceMb = &v
}

// GetHostOrClusterId returns the HostOrClusterId field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetHostOrClusterId() string {
	if o == nil || o.HostOrClusterId == nil {
		var ret string
		return ret
	}
	return *o.HostOrClusterId
}

// GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetHostOrClusterIdOk() (*string, bool) {
	if o == nil || o.HostOrClusterId == nil {
		return nil, false
	}
	return o.HostOrClusterId, true
}

// HasHostOrClusterId returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasHostOrClusterId() bool {
	if o != nil && o.HostOrClusterId != nil {
		return true
	}

	return false
}

// SetHostOrClusterId gets a reference to the given string and assigns it to the HostOrClusterId field.
func (o *DatastoreClusterInfo) SetHostOrClusterId(v string) {
	o.HostOrClusterId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DatastoreClusterInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *DatastoreClusterInfo) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *DatastoreClusterInfo) SetPath(v string) {
	o.Path = &v
}

// GetVcenterId returns the VcenterId field value if set, zero value otherwise.
func (o *DatastoreClusterInfo) GetVcenterId() string {
	if o == nil || o.VcenterId == nil {
		var ret string
		return ret
	}
	return *o.VcenterId
}

// GetVcenterIdOk returns a tuple with the VcenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatastoreClusterInfo) GetVcenterIdOk() (*string, bool) {
	if o == nil || o.VcenterId == nil {
		return nil, false
	}
	return o.VcenterId, true
}

// HasVcenterId returns a boolean if a field has been set.
func (o *DatastoreClusterInfo) HasVcenterId() bool {
	if o != nil && o.VcenterId != nil {
		return true
	}

	return false
}

// SetVcenterId gets a reference to the given string and assigns it to the VcenterId field.
func (o *DatastoreClusterInfo) SetVcenterId(v string) {
	o.VcenterId = &v
}

func (o DatastoreClusterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapacityMb != nil {
		toSerialize["capacity_mb"] = o.CapacityMb
	}
	if o.DatacenterId != nil {
		toSerialize["datacenter_id"] = o.DatacenterId
	}
	if o.DatastoreIds != nil {
		toSerialize["datastore_ids"] = o.DatastoreIds
	}
	if o.FreeSpaceMb != nil {
		toSerialize["free_space_mb"] = o.FreeSpaceMb
	}
	if o.HostOrClusterId != nil {
		toSerialize["host_or_cluster_id"] = o.HostOrClusterId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.VcenterId != nil {
		toSerialize["vcenter_id"] = o.VcenterId
	}
	return json.Marshal(toSerialize)
}

type NullableDatastoreClusterInfo struct {
	value *DatastoreClusterInfo
	isSet bool
}

func (v NullableDatastoreClusterInfo) Get() *DatastoreClusterInfo {
	return v.value
}

func (v *NullableDatastoreClusterInfo) Set(val *DatastoreClusterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDatastoreClusterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDatastoreClusterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatastoreClusterInfo(val *DatastoreClusterInfo) *NullableDatastoreClusterInfo {
	return &NullableDatastoreClusterInfo{value: val, isSet: true}
}

func (v NullableDatastoreClusterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatastoreClusterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
