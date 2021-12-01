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

// StorageAcceleratorData Information about the Storage Accelerator Data
type StorageAcceleratorData struct {
	// Default size of the cache in megabytes. This property has a default value of 1024. This property has a minimum value of 100. This property has a maximum value of 2048.
	DefaultCacheSizeMb *int32 `json:"default_cache_size_mb,omitempty"`
	// Is View Storage Accelerator enabled? This property has a default value of false.
	Enabled *bool `json:"enabled,omitempty"`
	// Cache size overrides for hosts which support View Storage Accelerator.
	HostOverrides *[]HostOverrideData `json:"host_overrides,omitempty"`
}

// NewStorageAcceleratorData instantiates a new StorageAcceleratorData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageAcceleratorData() *StorageAcceleratorData {
	this := StorageAcceleratorData{}
	return &this
}

// NewStorageAcceleratorDataWithDefaults instantiates a new StorageAcceleratorData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageAcceleratorDataWithDefaults() *StorageAcceleratorData {
	this := StorageAcceleratorData{}
	return &this
}

// GetDefaultCacheSizeMb returns the DefaultCacheSizeMb field value if set, zero value otherwise.
func (o *StorageAcceleratorData) GetDefaultCacheSizeMb() int32 {
	if o == nil || o.DefaultCacheSizeMb == nil {
		var ret int32
		return ret
	}
	return *o.DefaultCacheSizeMb
}

// GetDefaultCacheSizeMbOk returns a tuple with the DefaultCacheSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAcceleratorData) GetDefaultCacheSizeMbOk() (*int32, bool) {
	if o == nil || o.DefaultCacheSizeMb == nil {
		return nil, false
	}
	return o.DefaultCacheSizeMb, true
}

// HasDefaultCacheSizeMb returns a boolean if a field has been set.
func (o *StorageAcceleratorData) HasDefaultCacheSizeMb() bool {
	if o != nil && o.DefaultCacheSizeMb != nil {
		return true
	}

	return false
}

// SetDefaultCacheSizeMb gets a reference to the given int32 and assigns it to the DefaultCacheSizeMb field.
func (o *StorageAcceleratorData) SetDefaultCacheSizeMb(v int32) {
	o.DefaultCacheSizeMb = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StorageAcceleratorData) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAcceleratorData) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StorageAcceleratorData) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StorageAcceleratorData) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetHostOverrides returns the HostOverrides field value if set, zero value otherwise.
func (o *StorageAcceleratorData) GetHostOverrides() []HostOverrideData {
	if o == nil || o.HostOverrides == nil {
		var ret []HostOverrideData
		return ret
	}
	return *o.HostOverrides
}

// GetHostOverridesOk returns a tuple with the HostOverrides field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageAcceleratorData) GetHostOverridesOk() (*[]HostOverrideData, bool) {
	if o == nil || o.HostOverrides == nil {
		return nil, false
	}
	return o.HostOverrides, true
}

// HasHostOverrides returns a boolean if a field has been set.
func (o *StorageAcceleratorData) HasHostOverrides() bool {
	if o != nil && o.HostOverrides != nil {
		return true
	}

	return false
}

// SetHostOverrides gets a reference to the given []HostOverrideData and assigns it to the HostOverrides field.
func (o *StorageAcceleratorData) SetHostOverrides(v []HostOverrideData) {
	o.HostOverrides = &v
}

func (o StorageAcceleratorData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DefaultCacheSizeMb != nil {
		toSerialize["default_cache_size_mb"] = o.DefaultCacheSizeMb
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.HostOverrides != nil {
		toSerialize["host_overrides"] = o.HostOverrides
	}
	return json.Marshal(toSerialize)
}

type NullableStorageAcceleratorData struct {
	value *StorageAcceleratorData
	isSet bool
}

func (v NullableStorageAcceleratorData) Get() *StorageAcceleratorData {
	return v.value
}

func (v *NullableStorageAcceleratorData) Set(val *StorageAcceleratorData) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageAcceleratorData) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageAcceleratorData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageAcceleratorData(val *StorageAcceleratorData) *NullableStorageAcceleratorData {
	return &NullableStorageAcceleratorData{value: val, isSet: true}
}

func (v NullableStorageAcceleratorData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageAcceleratorData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
