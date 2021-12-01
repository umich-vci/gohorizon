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

// HostOverrideData Details of the Hypervisor hosts involved in Content Based Read Caching (CBRC)
type HostOverrideData struct {
	// Size of the cache in megabytes. This property has a minimum value of 100. This property has a maximum value of 2048.
	CacheSizeMb *int32 `json:"cache_size_mb,omitempty"`
	// The path of the host that supports View Storage Accelerator.
	Path *string `json:"path,omitempty"`
}

// NewHostOverrideData instantiates a new HostOverrideData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostOverrideData() *HostOverrideData {
	this := HostOverrideData{}
	return &this
}

// NewHostOverrideDataWithDefaults instantiates a new HostOverrideData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostOverrideDataWithDefaults() *HostOverrideData {
	this := HostOverrideData{}
	return &this
}

// GetCacheSizeMb returns the CacheSizeMb field value if set, zero value otherwise.
func (o *HostOverrideData) GetCacheSizeMb() int32 {
	if o == nil || o.CacheSizeMb == nil {
		var ret int32
		return ret
	}
	return *o.CacheSizeMb
}

// GetCacheSizeMbOk returns a tuple with the CacheSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOverrideData) GetCacheSizeMbOk() (*int32, bool) {
	if o == nil || o.CacheSizeMb == nil {
		return nil, false
	}
	return o.CacheSizeMb, true
}

// HasCacheSizeMb returns a boolean if a field has been set.
func (o *HostOverrideData) HasCacheSizeMb() bool {
	if o != nil && o.CacheSizeMb != nil {
		return true
	}

	return false
}

// SetCacheSizeMb gets a reference to the given int32 and assigns it to the CacheSizeMb field.
func (o *HostOverrideData) SetCacheSizeMb(v int32) {
	o.CacheSizeMb = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HostOverrideData) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostOverrideData) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HostOverrideData) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *HostOverrideData) SetPath(v string) {
	o.Path = &v
}

func (o HostOverrideData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CacheSizeMb != nil {
		toSerialize["cache_size_mb"] = o.CacheSizeMb
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableHostOverrideData struct {
	value *HostOverrideData
	isSet bool
}

func (v NullableHostOverrideData) Get() *HostOverrideData {
	return v.value
}

func (v *NullableHostOverrideData) Set(val *HostOverrideData) {
	v.value = val
	v.isSet = true
}

func (v NullableHostOverrideData) IsSet() bool {
	return v.isSet
}

func (v *NullableHostOverrideData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostOverrideData(val *HostOverrideData) *NullableHostOverrideData {
	return &NullableHostOverrideData{value: val, isSet: true}
}

func (v NullableHostOverrideData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostOverrideData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
