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

// GlobalSessionActionSpec Information required to perform an action on global sessions.
type GlobalSessionActionSpec struct {
	// IDs of the sessions on which action is to be performed.
	Ids []string `json:"ids"`
	// ID of the hosting pod for the sessions.
	PodId string `json:"pod_id"`
}

// NewGlobalSessionActionSpec instantiates a new GlobalSessionActionSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalSessionActionSpec(ids []string, podId string) *GlobalSessionActionSpec {
	this := GlobalSessionActionSpec{}
	this.Ids = ids
	this.PodId = podId
	return &this
}

// NewGlobalSessionActionSpecWithDefaults instantiates a new GlobalSessionActionSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalSessionActionSpecWithDefaults() *GlobalSessionActionSpec {
	this := GlobalSessionActionSpec{}
	return &this
}

// GetIds returns the Ids field value
func (o *GlobalSessionActionSpec) GetIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Ids
}

// GetIdsOk returns a tuple with the Ids field value
// and a boolean to check if the value has been set.
func (o *GlobalSessionActionSpec) GetIdsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ids, true
}

// SetIds sets field value
func (o *GlobalSessionActionSpec) SetIds(v []string) {
	o.Ids = v
}

// GetPodId returns the PodId field value
func (o *GlobalSessionActionSpec) GetPodId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PodId
}

// GetPodIdOk returns a tuple with the PodId field value
// and a boolean to check if the value has been set.
func (o *GlobalSessionActionSpec) GetPodIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PodId, true
}

// SetPodId sets field value
func (o *GlobalSessionActionSpec) SetPodId(v string) {
	o.PodId = v
}

func (o GlobalSessionActionSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ids"] = o.Ids
	}
	if true {
		toSerialize["pod_id"] = o.PodId
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalSessionActionSpec struct {
	value *GlobalSessionActionSpec
	isSet bool
}

func (v NullableGlobalSessionActionSpec) Get() *GlobalSessionActionSpec {
	return v.value
}

func (v *NullableGlobalSessionActionSpec) Set(val *GlobalSessionActionSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalSessionActionSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalSessionActionSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalSessionActionSpec(val *GlobalSessionActionSpec) *NullableGlobalSessionActionSpec {
	return &NullableGlobalSessionActionSpec{value: val, isSet: true}
}

func (v NullableGlobalSessionActionSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalSessionActionSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}