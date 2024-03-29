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

// GatewayMonitorDetails Details of the Gateway.
type GatewayMonitorDetails struct {
	// Gateway host name or IP address.
	Address *string `json:"address,omitempty"`
	// Indicates if the Gateway is internal.
	Internal *bool `json:"internal,omitempty"`
	// Type of the Gateway. * UAG: Unified Access Gateway type. * F5: F5 Gateway type. * UNKNOWN: Unknown type.
	Type *string `json:"type,omitempty"`
	// Version of the Gateway.
	Version *string `json:"version,omitempty"`
}

// NewGatewayMonitorDetails instantiates a new GatewayMonitorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMonitorDetails() *GatewayMonitorDetails {
	this := GatewayMonitorDetails{}
	return &this
}

// NewGatewayMonitorDetailsWithDefaults instantiates a new GatewayMonitorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMonitorDetailsWithDefaults() *GatewayMonitorDetails {
	this := GatewayMonitorDetails{}
	return &this
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *GatewayMonitorDetails) GetAddress() string {
	if o == nil || o.Address == nil {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorDetails) GetAddressOk() (*string, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *GatewayMonitorDetails) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *GatewayMonitorDetails) SetAddress(v string) {
	o.Address = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *GatewayMonitorDetails) GetInternal() bool {
	if o == nil || o.Internal == nil {
		var ret bool
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorDetails) GetInternalOk() (*bool, bool) {
	if o == nil || o.Internal == nil {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *GatewayMonitorDetails) HasInternal() bool {
	if o != nil && o.Internal != nil {
		return true
	}

	return false
}

// SetInternal gets a reference to the given bool and assigns it to the Internal field.
func (o *GatewayMonitorDetails) SetInternal(v bool) {
	o.Internal = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GatewayMonitorDetails) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorDetails) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *GatewayMonitorDetails) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GatewayMonitorDetails) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *GatewayMonitorDetails) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorDetails) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *GatewayMonitorDetails) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *GatewayMonitorDetails) SetVersion(v string) {
	o.Version = &v
}

func (o GatewayMonitorDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Internal != nil {
		toSerialize["internal"] = o.Internal
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayMonitorDetails struct {
	value *GatewayMonitorDetails
	isSet bool
}

func (v NullableGatewayMonitorDetails) Get() *GatewayMonitorDetails {
	return v.value
}

func (v *NullableGatewayMonitorDetails) Set(val *GatewayMonitorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMonitorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMonitorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMonitorDetails(val *GatewayMonitorDetails) *NullableGatewayMonitorDetails {
	return &NullableGatewayMonitorDetails{value: val, isSet: true}
}

func (v NullableGatewayMonitorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMonitorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
