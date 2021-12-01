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

// FarmDatastoreSettingsUpdateSpec Datastore settings for the automated farm.
type FarmDatastoreSettingsUpdateSpec struct {
	// Id of the datastore.
	DatastoreId string `json:"datastore_id"`
}

// NewFarmDatastoreSettingsUpdateSpec instantiates a new FarmDatastoreSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmDatastoreSettingsUpdateSpec(datastoreId string) *FarmDatastoreSettingsUpdateSpec {
	this := FarmDatastoreSettingsUpdateSpec{}
	this.DatastoreId = datastoreId
	return &this
}

// NewFarmDatastoreSettingsUpdateSpecWithDefaults instantiates a new FarmDatastoreSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmDatastoreSettingsUpdateSpecWithDefaults() *FarmDatastoreSettingsUpdateSpec {
	this := FarmDatastoreSettingsUpdateSpec{}
	return &this
}

// GetDatastoreId returns the DatastoreId field value
func (o *FarmDatastoreSettingsUpdateSpec) GetDatastoreId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatastoreId
}

// GetDatastoreIdOk returns a tuple with the DatastoreId field value
// and a boolean to check if the value has been set.
func (o *FarmDatastoreSettingsUpdateSpec) GetDatastoreIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatastoreId, true
}

// SetDatastoreId sets field value
func (o *FarmDatastoreSettingsUpdateSpec) SetDatastoreId(v string) {
	o.DatastoreId = v
}

func (o FarmDatastoreSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["datastore_id"] = o.DatastoreId
	}
	return json.Marshal(toSerialize)
}

type NullableFarmDatastoreSettingsUpdateSpec struct {
	value *FarmDatastoreSettingsUpdateSpec
	isSet bool
}

func (v NullableFarmDatastoreSettingsUpdateSpec) Get() *FarmDatastoreSettingsUpdateSpec {
	return v.value
}

func (v *NullableFarmDatastoreSettingsUpdateSpec) Set(val *FarmDatastoreSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmDatastoreSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmDatastoreSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmDatastoreSettingsUpdateSpec(val *FarmDatastoreSettingsUpdateSpec) *NullableFarmDatastoreSettingsUpdateSpec {
	return &NullableFarmDatastoreSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableFarmDatastoreSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmDatastoreSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
