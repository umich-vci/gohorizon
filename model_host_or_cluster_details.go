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

// HostOrClusterDetails Details of the host or cluster.
type HostOrClusterDetails struct {
	// Whether or not this is a cluster or a host.
	Cluster *bool `json:"cluster,omitempty"`
	// Datacenter id for this host or cluster.
	DatacenterId *string `json:"datacenter_id,omitempty"`
	// Reasons that may preclude this Host Or Cluster from being used in desktop pool creation.
	IncompatibleReasons *[]string `json:"incompatible_reasons,omitempty"`
	// Host or cluster display name.
	Name *string `json:"name,omitempty"`
	// Host or cluster path.
	Path *string `json:"path,omitempty"`
	// Virtual Center id for this host or cluster.
	VcenterId *string `json:"vcenter_id,omitempty"`
	// Types of NVIDIA GRID vGPUs supported by this host or at least one host on this cluster. If unset, this host or cluster does not support NVIDIA GRID vGPUs and cannot be used for desktop creation with NVIDIA GRID vGPU support enabled.
	VgpuTypes *[]string `json:"vgpu_types,omitempty"`
}

// NewHostOrClusterDetails instantiates a new HostOrClusterDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostOrClusterDetails() *HostOrClusterDetails {
	this := HostOrClusterDetails{}
	return &this
}

// NewHostOrClusterDetailsWithDefaults instantiates a new HostOrClusterDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostOrClusterDetailsWithDefaults() *HostOrClusterDetails {
	this := HostOrClusterDetails{}
	return &this
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *HostOrClusterDetails) GetCluster() bool {
	if o == nil || o.Cluster == nil {
		var ret bool
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOrClusterDetails) GetClusterOk() (*bool, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *HostOrClusterDetails) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given bool and assigns it to the Cluster field.
func (o *HostOrClusterDetails) SetCluster(v bool) {
	o.Cluster = &v
}

// GetDatacenterId returns the DatacenterId field value if set, zero value otherwise.
func (o *HostOrClusterDetails) GetDatacenterId() string {
	if o == nil || o.DatacenterId == nil {
		var ret string
		return ret
	}
	return *o.DatacenterId
}

// GetDatacenterIdOk returns a tuple with the DatacenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOrClusterDetails) GetDatacenterIdOk() (*string, bool) {
	if o == nil || o.DatacenterId == nil {
		return nil, false
	}
	return o.DatacenterId, true
}

// HasDatacenterId returns a boolean if a field has been set.
func (o *HostOrClusterDetails) HasDatacenterId() bool {
	if o != nil && o.DatacenterId != nil {
		return true
	}

	return false
}

// SetDatacenterId gets a reference to the given string and assigns it to the DatacenterId field.
func (o *HostOrClusterDetails) SetDatacenterId(v string) {
	o.DatacenterId = &v
}

// GetIncompatibleReasons returns the IncompatibleReasons field value if set, zero value otherwise.
func (o *HostOrClusterDetails) GetIncompatibleReasons() []string {
	if o == nil || o.IncompatibleReasons == nil {
		var ret []string
		return ret
	}
	return *o.IncompatibleReasons
}

// GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOrClusterDetails) GetIncompatibleReasonsOk() (*[]string, bool) {
	if o == nil || o.IncompatibleReasons == nil {
		return nil, false
	}
	return o.IncompatibleReasons, true
}

// HasIncompatibleReasons returns a boolean if a field has been set.
func (o *HostOrClusterDetails) HasIncompatibleReasons() bool {
	if o != nil && o.IncompatibleReasons != nil {
		return true
	}

	return false
}

// SetIncompatibleReasons gets a reference to the given []string and assigns it to the IncompatibleReasons field.
func (o *HostOrClusterDetails) SetIncompatibleReasons(v []string) {
	o.IncompatibleReasons = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostOrClusterDetails) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOrClusterDetails) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostOrClusterDetails) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HostOrClusterDetails) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HostOrClusterDetails) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOrClusterDetails) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HostOrClusterDetails) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HostOrClusterDetails) SetPath(v string) {
	o.Path = &v
}

// GetVcenterId returns the VcenterId field value if set, zero value otherwise.
func (o *HostOrClusterDetails) GetVcenterId() string {
	if o == nil || o.VcenterId == nil {
		var ret string
		return ret
	}
	return *o.VcenterId
}

// GetVcenterIdOk returns a tuple with the VcenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOrClusterDetails) GetVcenterIdOk() (*string, bool) {
	if o == nil || o.VcenterId == nil {
		return nil, false
	}
	return o.VcenterId, true
}

// HasVcenterId returns a boolean if a field has been set.
func (o *HostOrClusterDetails) HasVcenterId() bool {
	if o != nil && o.VcenterId != nil {
		return true
	}

	return false
}

// SetVcenterId gets a reference to the given string and assigns it to the VcenterId field.
func (o *HostOrClusterDetails) SetVcenterId(v string) {
	o.VcenterId = &v
}

// GetVgpuTypes returns the VgpuTypes field value if set, zero value otherwise.
func (o *HostOrClusterDetails) GetVgpuTypes() []string {
	if o == nil || o.VgpuTypes == nil {
		var ret []string
		return ret
	}
	return *o.VgpuTypes
}

// GetVgpuTypesOk returns a tuple with the VgpuTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOrClusterDetails) GetVgpuTypesOk() (*[]string, bool) {
	if o == nil || o.VgpuTypes == nil {
		return nil, false
	}
	return o.VgpuTypes, true
}

// HasVgpuTypes returns a boolean if a field has been set.
func (o *HostOrClusterDetails) HasVgpuTypes() bool {
	if o != nil && o.VgpuTypes != nil {
		return true
	}

	return false
}

// SetVgpuTypes gets a reference to the given []string and assigns it to the VgpuTypes field.
func (o *HostOrClusterDetails) SetVgpuTypes(v []string) {
	o.VgpuTypes = &v
}

func (o HostOrClusterDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	if o.DatacenterId != nil {
		toSerialize["datacenter_id"] = o.DatacenterId
	}
	if o.IncompatibleReasons != nil {
		toSerialize["incompatible_reasons"] = o.IncompatibleReasons
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
	if o.VgpuTypes != nil {
		toSerialize["vgpu_types"] = o.VgpuTypes
	}
	return json.Marshal(toSerialize)
}

type NullableHostOrClusterDetails struct {
	value *HostOrClusterDetails
	isSet bool
}

func (v NullableHostOrClusterDetails) Get() *HostOrClusterDetails {
	return v.value
}

func (v *NullableHostOrClusterDetails) Set(val *HostOrClusterDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableHostOrClusterDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableHostOrClusterDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostOrClusterDetails(val *HostOrClusterDetails) *NullableHostOrClusterDetails {
	return &NullableHostOrClusterDetails{value: val, isSet: true}
}

func (v NullableHostOrClusterDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostOrClusterDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
