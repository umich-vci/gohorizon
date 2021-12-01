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

// ApplicationAntiAffinityData Anti-affinity data required to create an application pool.
type ApplicationAntiAffinityData struct {
	// Maximum number of other applications that can be running on the RDS Server before the RDS Server is rejected for new application sessions.
	AntiAffinityCount int32 `json:"anti_affinity_count"`
	// Set of pattern strings to match against process names on a RDS Server when attempting to launch a session for this Application. For each application running on an RDSServer that matches one of the patterns, the tally against the count threshold is incremented.<br>Pattern strings may contain wildcard characters. '*' matches zero or more characters. '?' matches exactly one character.
	AntiAffinityPatterns []string `json:"anti_affinity_patterns"`
}

// NewApplicationAntiAffinityData instantiates a new ApplicationAntiAffinityData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationAntiAffinityData(antiAffinityCount int32, antiAffinityPatterns []string) *ApplicationAntiAffinityData {
	this := ApplicationAntiAffinityData{}
	this.AntiAffinityCount = antiAffinityCount
	this.AntiAffinityPatterns = antiAffinityPatterns
	return &this
}

// NewApplicationAntiAffinityDataWithDefaults instantiates a new ApplicationAntiAffinityData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationAntiAffinityDataWithDefaults() *ApplicationAntiAffinityData {
	this := ApplicationAntiAffinityData{}
	return &this
}

// GetAntiAffinityCount returns the AntiAffinityCount field value
func (o *ApplicationAntiAffinityData) GetAntiAffinityCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AntiAffinityCount
}

// GetAntiAffinityCountOk returns a tuple with the AntiAffinityCount field value
// and a boolean to check if the value has been set.
func (o *ApplicationAntiAffinityData) GetAntiAffinityCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AntiAffinityCount, true
}

// SetAntiAffinityCount sets field value
func (o *ApplicationAntiAffinityData) SetAntiAffinityCount(v int32) {
	o.AntiAffinityCount = v
}

// GetAntiAffinityPatterns returns the AntiAffinityPatterns field value
func (o *ApplicationAntiAffinityData) GetAntiAffinityPatterns() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AntiAffinityPatterns
}

// GetAntiAffinityPatternsOk returns a tuple with the AntiAffinityPatterns field value
// and a boolean to check if the value has been set.
func (o *ApplicationAntiAffinityData) GetAntiAffinityPatternsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AntiAffinityPatterns, true
}

// SetAntiAffinityPatterns sets field value
func (o *ApplicationAntiAffinityData) SetAntiAffinityPatterns(v []string) {
	o.AntiAffinityPatterns = v
}

func (o ApplicationAntiAffinityData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["anti_affinity_count"] = o.AntiAffinityCount
	}
	if true {
		toSerialize["anti_affinity_patterns"] = o.AntiAffinityPatterns
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationAntiAffinityData struct {
	value *ApplicationAntiAffinityData
	isSet bool
}

func (v NullableApplicationAntiAffinityData) Get() *ApplicationAntiAffinityData {
	return v.value
}

func (v *NullableApplicationAntiAffinityData) Set(val *ApplicationAntiAffinityData) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationAntiAffinityData) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationAntiAffinityData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationAntiAffinityData(val *ApplicationAntiAffinityData) *NullableApplicationAntiAffinityData {
	return &NullableApplicationAntiAffinityData{value: val, isSet: true}
}

func (v NullableApplicationAntiAffinityData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationAntiAffinityData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
