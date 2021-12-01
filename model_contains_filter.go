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

// ContainsFilter struct for ContainsFilter
type ContainsFilter struct {
	BaseFilter
	Name  *string                 `json:"name,omitempty"`
	Type  *string                 `json:"type,omitempty"`
	Value *map[string]interface{} `json:"value,omitempty"`
}

// NewContainsFilter instantiates a new ContainsFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainsFilter() *ContainsFilter {
	this := ContainsFilter{}
	return &this
}

// NewContainsFilterWithDefaults instantiates a new ContainsFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainsFilterWithDefaults() *ContainsFilter {
	this := ContainsFilter{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ContainsFilter) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainsFilter) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ContainsFilter) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ContainsFilter) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ContainsFilter) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainsFilter) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ContainsFilter) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ContainsFilter) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ContainsFilter) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainsFilter) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ContainsFilter) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *ContainsFilter) SetValue(v map[string]interface{}) {
	o.Value = &v
}

func (o ContainsFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseFilter, errBaseFilter := json.Marshal(o.BaseFilter)
	if errBaseFilter != nil {
		return []byte{}, errBaseFilter
	}
	errBaseFilter = json.Unmarshal([]byte(serializedBaseFilter), &toSerialize)
	if errBaseFilter != nil {
		return []byte{}, errBaseFilter
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableContainsFilter struct {
	value *ContainsFilter
	isSet bool
}

func (v NullableContainsFilter) Get() *ContainsFilter {
	return v.value
}

func (v *NullableContainsFilter) Set(val *ContainsFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableContainsFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableContainsFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContainsFilter(val *ContainsFilter) *NullableContainsFilter {
	return &NullableContainsFilter{value: val, isSet: true}
}

func (v NullableContainsFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContainsFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
