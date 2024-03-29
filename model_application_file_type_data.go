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

// ApplicationFileTypeData Information about the file type supported by the application.
type ApplicationFileTypeData struct {
	// Friendly name for the file type. If unset, no friendly name will be displayed.
	Description *string `json:"description,omitempty"`
	// File type supported by this application. This value is case insensitive. If multiple file types are specified using the same (case insensitive) name and description, all but one will be ignored.
	Type string `json:"type"`
}

// NewApplicationFileTypeData instantiates a new ApplicationFileTypeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationFileTypeData(type_ string) *ApplicationFileTypeData {
	this := ApplicationFileTypeData{}
	this.Type = type_
	return &this
}

// NewApplicationFileTypeDataWithDefaults instantiates a new ApplicationFileTypeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationFileTypeDataWithDefaults() *ApplicationFileTypeData {
	this := ApplicationFileTypeData{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationFileTypeData) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationFileTypeData) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationFileTypeData) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationFileTypeData) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *ApplicationFileTypeData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationFileTypeData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationFileTypeData) SetType(v string) {
	o.Type = v
}

func (o ApplicationFileTypeData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationFileTypeData struct {
	value *ApplicationFileTypeData
	isSet bool
}

func (v NullableApplicationFileTypeData) Get() *ApplicationFileTypeData {
	return v.value
}

func (v *NullableApplicationFileTypeData) Set(val *ApplicationFileTypeData) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationFileTypeData) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationFileTypeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationFileTypeData(val *ApplicationFileTypeData) *NullableApplicationFileTypeData {
	return &NullableApplicationFileTypeData{value: val, isSet: true}
}

func (v NullableApplicationFileTypeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationFileTypeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
