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

// CPAEjectSpec struct for CPAEjectSpec
type CPAEjectSpec struct {
	// The ID of the pod that has to be removed from Pod Federation.
	PodId string `json:"pod_id"`
}

// NewCPAEjectSpec instantiates a new CPAEjectSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCPAEjectSpec(podId string) *CPAEjectSpec {
	this := CPAEjectSpec{}
	this.PodId = podId
	return &this
}

// NewCPAEjectSpecWithDefaults instantiates a new CPAEjectSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCPAEjectSpecWithDefaults() *CPAEjectSpec {
	this := CPAEjectSpec{}
	return &this
}

// GetPodId returns the PodId field value
func (o *CPAEjectSpec) GetPodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodId
}

// GetPodIdOk returns a tuple with the PodId field value
// and a boolean to check if the value has been set.
func (o *CPAEjectSpec) GetPodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodId, true
}

// SetPodId sets field value
func (o *CPAEjectSpec) SetPodId(v string) {
	o.PodId = v
}

func (o CPAEjectSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pod_id"] = o.PodId
	}
	return json.Marshal(toSerialize)
}

type NullableCPAEjectSpec struct {
	value *CPAEjectSpec
	isSet bool
}

func (v NullableCPAEjectSpec) Get() *CPAEjectSpec {
	return v.value
}

func (v *NullableCPAEjectSpec) Set(val *CPAEjectSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableCPAEjectSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCPAEjectSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCPAEjectSpec(val *CPAEjectSpec) *NullableCPAEjectSpec {
	return &NullableCPAEjectSpec{value: val, isSet: true}
}

func (v NullableCPAEjectSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCPAEjectSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
