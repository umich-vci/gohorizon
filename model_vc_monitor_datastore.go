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

// VCMonitorDatastore Information about the datastore of the host(s) connected to the given Virtual Center.
type VCMonitorDatastore struct {
	// The capacity of the datastore in megabytes.
	CapacityMb *int32                     `json:"capacity_mb,omitempty"`
	Details    *VCMonitorDatastoreDetails `json:"details,omitempty"`
	// The free space on the datastore in megabytes.
	FreeSpaceMb *int32 `json:"free_space_mb,omitempty"`
	// Status of the datastore. * ACCESSIBLE: The datastore is accessible. * NOT_ACCESSIBLE: The datastore is not accessible.
	Status *string `json:"status,omitempty"`
	// Type of the datastore. * VMFS: VMFS datastore. * VSAN: VSAN datastore. * VVOL: VVOL datastore.
	Type *string `json:"type,omitempty"`
}

// NewVCMonitorDatastore instantiates a new VCMonitorDatastore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVCMonitorDatastore() *VCMonitorDatastore {
	this := VCMonitorDatastore{}
	return &this
}

// NewVCMonitorDatastoreWithDefaults instantiates a new VCMonitorDatastore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVCMonitorDatastoreWithDefaults() *VCMonitorDatastore {
	this := VCMonitorDatastore{}
	return &this
}

// GetCapacityMb returns the CapacityMb field value if set, zero value otherwise.
func (o *VCMonitorDatastore) GetCapacityMb() int32 {
	if o == nil || o.CapacityMb == nil {
		var ret int32
		return ret
	}
	return *o.CapacityMb
}

// GetCapacityMbOk returns a tuple with the CapacityMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorDatastore) GetCapacityMbOk() (*int32, bool) {
	if o == nil || o.CapacityMb == nil {
		return nil, false
	}
	return o.CapacityMb, true
}

// HasCapacityMb returns a boolean if a field has been set.
func (o *VCMonitorDatastore) HasCapacityMb() bool {
	if o != nil && o.CapacityMb != nil {
		return true
	}

	return false
}

// SetCapacityMb gets a reference to the given int32 and assigns it to the CapacityMb field.
func (o *VCMonitorDatastore) SetCapacityMb(v int32) {
	o.CapacityMb = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VCMonitorDatastore) GetDetails() VCMonitorDatastoreDetails {
	if o == nil || o.Details == nil {
		var ret VCMonitorDatastoreDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorDatastore) GetDetailsOk() (*VCMonitorDatastoreDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VCMonitorDatastore) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given VCMonitorDatastoreDetails and assigns it to the Details field.
func (o *VCMonitorDatastore) SetDetails(v VCMonitorDatastoreDetails) {
	o.Details = &v
}

// GetFreeSpaceMb returns the FreeSpaceMb field value if set, zero value otherwise.
func (o *VCMonitorDatastore) GetFreeSpaceMb() int32 {
	if o == nil || o.FreeSpaceMb == nil {
		var ret int32
		return ret
	}
	return *o.FreeSpaceMb
}

// GetFreeSpaceMbOk returns a tuple with the FreeSpaceMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorDatastore) GetFreeSpaceMbOk() (*int32, bool) {
	if o == nil || o.FreeSpaceMb == nil {
		return nil, false
	}
	return o.FreeSpaceMb, true
}

// HasFreeSpaceMb returns a boolean if a field has been set.
func (o *VCMonitorDatastore) HasFreeSpaceMb() bool {
	if o != nil && o.FreeSpaceMb != nil {
		return true
	}

	return false
}

// SetFreeSpaceMb gets a reference to the given int32 and assigns it to the FreeSpaceMb field.
func (o *VCMonitorDatastore) SetFreeSpaceMb(v int32) {
	o.FreeSpaceMb = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VCMonitorDatastore) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorDatastore) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VCMonitorDatastore) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VCMonitorDatastore) SetStatus(v string) {
	o.Status = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VCMonitorDatastore) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorDatastore) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VCMonitorDatastore) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VCMonitorDatastore) SetType(v string) {
	o.Type = &v
}

func (o VCMonitorDatastore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CapacityMb != nil {
		toSerialize["capacity_mb"] = o.CapacityMb
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.FreeSpaceMb != nil {
		toSerialize["free_space_mb"] = o.FreeSpaceMb
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableVCMonitorDatastore struct {
	value *VCMonitorDatastore
	isSet bool
}

func (v NullableVCMonitorDatastore) Get() *VCMonitorDatastore {
	return v.value
}

func (v *NullableVCMonitorDatastore) Set(val *VCMonitorDatastore) {
	v.value = val
	v.isSet = true
}

func (v NullableVCMonitorDatastore) IsSet() bool {
	return v.isSet
}

func (v *NullableVCMonitorDatastore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVCMonitorDatastore(val *VCMonitorDatastore) *NullableVCMonitorDatastore {
	return &NullableVCMonitorDatastore{value: val, isSet: true}
}

func (v NullableVCMonitorDatastore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVCMonitorDatastore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
