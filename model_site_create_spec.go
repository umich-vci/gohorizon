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

// SiteCreateSpec Information required to create a site.
type SiteCreateSpec struct {
	// Detailed description of the site.
	Description *string `json:"description,omitempty"`
	// The name of the site.
	Name string `json:"name"`
}

// NewSiteCreateSpec instantiates a new SiteCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiteCreateSpec(name string) *SiteCreateSpec {
	this := SiteCreateSpec{}
	this.Name = name
	return &this
}

// NewSiteCreateSpecWithDefaults instantiates a new SiteCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiteCreateSpecWithDefaults() *SiteCreateSpec {
	this := SiteCreateSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SiteCreateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SiteCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SiteCreateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SiteCreateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *SiteCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SiteCreateSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SiteCreateSpec) SetName(v string) {
	o.Name = v
}

func (o SiteCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableSiteCreateSpec struct {
	value *SiteCreateSpec
	isSet bool
}

func (v NullableSiteCreateSpec) Get() *SiteCreateSpec {
	return v.value
}

func (v *NullableSiteCreateSpec) Set(val *SiteCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableSiteCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSiteCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiteCreateSpec(val *SiteCreateSpec) *NullableSiteCreateSpec {
	return &NullableSiteCreateSpec{value: val, isSet: true}
}

func (v NullableSiteCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiteCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}