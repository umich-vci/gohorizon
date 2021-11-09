/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2106
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// DatacenterInfo Information related to datacenter.
type DatacenterInfo struct {
	// Unique ID representing a datacenter.
	Id string `json:"id"`
	// Name of the datacenter.
	Name string `json:"name"`
	// Datacenter path.
	Path string `json:"path"`
}

// NewDatacenterInfo instantiates a new DatacenterInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatacenterInfo(id string, name string, path string) *DatacenterInfo {
	this := DatacenterInfo{}
	this.Id = id
	this.Name = name
	this.Path = path
	return &this
}

// NewDatacenterInfoWithDefaults instantiates a new DatacenterInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatacenterInfoWithDefaults() *DatacenterInfo {
	this := DatacenterInfo{}
	return &this
}

// GetId returns the Id field value
func (o *DatacenterInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DatacenterInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DatacenterInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *DatacenterInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatacenterInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatacenterInfo) SetName(v string) {
	o.Name = v
}

// GetPath returns the Path field value
func (o *DatacenterInfo) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DatacenterInfo) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DatacenterInfo) SetPath(v string) {
	o.Path = v
}

func (o DatacenterInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableDatacenterInfo struct {
	value *DatacenterInfo
	isSet bool
}

func (v NullableDatacenterInfo) Get() *DatacenterInfo {
	return v.value
}

func (v *NullableDatacenterInfo) Set(val *DatacenterInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDatacenterInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDatacenterInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatacenterInfo(val *DatacenterInfo) *NullableDatacenterInfo {
	return &NullableDatacenterInfo{value: val, isSet: true}
}

func (v NullableDatacenterInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatacenterInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
