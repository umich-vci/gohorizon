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

// ApplicationIconAssociateSpec struct for ApplicationIconAssociateSpec
type ApplicationIconAssociateSpec struct {
	// Unique ID representing an application icon.
	IconId string `json:"icon_id"`
}

// NewApplicationIconAssociateSpec instantiates a new ApplicationIconAssociateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationIconAssociateSpec(iconId string) *ApplicationIconAssociateSpec {
	this := ApplicationIconAssociateSpec{}
	this.IconId = iconId
	return &this
}

// NewApplicationIconAssociateSpecWithDefaults instantiates a new ApplicationIconAssociateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationIconAssociateSpecWithDefaults() *ApplicationIconAssociateSpec {
	this := ApplicationIconAssociateSpec{}
	return &this
}

// GetIconId returns the IconId field value
func (o *ApplicationIconAssociateSpec) GetIconId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IconId
}

// GetIconIdOk returns a tuple with the IconId field value
// and a boolean to check if the value has been set.
func (o *ApplicationIconAssociateSpec) GetIconIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IconId, true
}

// SetIconId sets field value
func (o *ApplicationIconAssociateSpec) SetIconId(v string) {
	o.IconId = v
}

func (o ApplicationIconAssociateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["icon_id"] = o.IconId
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationIconAssociateSpec struct {
	value *ApplicationIconAssociateSpec
	isSet bool
}

func (v NullableApplicationIconAssociateSpec) Get() *ApplicationIconAssociateSpec {
	return v.value
}

func (v *NullableApplicationIconAssociateSpec) Set(val *ApplicationIconAssociateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationIconAssociateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationIconAssociateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationIconAssociateSpec(val *ApplicationIconAssociateSpec) *NullableApplicationIconAssociateSpec {
	return &NullableApplicationIconAssociateSpec{value: val, isSet: true}
}

func (v NullableApplicationIconAssociateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationIconAssociateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}