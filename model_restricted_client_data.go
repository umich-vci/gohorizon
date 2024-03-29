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

// RestrictedClientData Information related to Restricted Horizon Clients.
type RestrictedClientData struct {
	// The type of Horizon Client. * WINDOWS: The client is the Horizon Client for Windows. * MAC: The client is the Horizon Client for MacOS. * HTMLACCESS: The client is a Web client. * LINUX: The client is the Horizon Client for Linux. * IOS: The client is the Horizon Client for IOS devices. * ANDROID: The client is the Horizon Client for Android. * WINSTORE: The client is the Horizon Client for Windows 10 UWP. * CHROME: The client is the Horizon Client for Chrome Native OS. * OTHER: Client type is other.
	Type *string `json:"type,omitempty"`
	// The version of Horizon Client.
	Version *string `json:"version,omitempty"`
}

// NewRestrictedClientData instantiates a new RestrictedClientData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRestrictedClientData() *RestrictedClientData {
	this := RestrictedClientData{}
	return &this
}

// NewRestrictedClientDataWithDefaults instantiates a new RestrictedClientData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRestrictedClientDataWithDefaults() *RestrictedClientData {
	this := RestrictedClientData{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RestrictedClientData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictedClientData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RestrictedClientData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RestrictedClientData) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *RestrictedClientData) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RestrictedClientData) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *RestrictedClientData) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *RestrictedClientData) SetVersion(v string) {
	o.Version = &v
}

func (o RestrictedClientData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableRestrictedClientData struct {
	value *RestrictedClientData
	isSet bool
}

func (v NullableRestrictedClientData) Get() *RestrictedClientData {
	return v.value
}

func (v *NullableRestrictedClientData) Set(val *RestrictedClientData) {
	v.value = val
	v.isSet = true
}

func (v NullableRestrictedClientData) IsSet() bool {
	return v.isSet
}

func (v *NullableRestrictedClientData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRestrictedClientData(val *RestrictedClientData) *NullableRestrictedClientData {
	return &NullableRestrictedClientData{value: val, isSet: true}
}

func (v NullableRestrictedClientData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRestrictedClientData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
