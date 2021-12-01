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

// SAMLAuthenticatorMonitorConnectionServerV2 Information about the SAML authenticator connection from a connection server.
type SAMLAuthenticatorMonitorConnectionServerV2 struct {
	// Unique ID of the Connection Server.
	Id *string `json:"id,omitempty"`
	// The timestamp in milliseconds when the last update was obtained. Measured as epoch time.
	LastUpdatedTimestamp *int64 `json:"last_updated_timestamp,omitempty"`
	// Connection server host name or IP address.
	Name *string `json:"name,omitempty"`
	// Status of the SAML authenticator with respect to this Connection Server. * OK: The connection to SAML authenticator is working properly. * ERROR: Error occurred when connecting to SAML authenticator. * WARN: The connection to SAML authenticator has minor issues. * UNKNOWN: State of SAML authenticator is unknown.
	Status *string `json:"status,omitempty"`
	// Indicates if the thumbprint of the SAML authenticator was accepted.
	ThumbprintAccepted *bool `json:"thumbprint_accepted,omitempty"`
}

// NewSAMLAuthenticatorMonitorConnectionServerV2 instantiates a new SAMLAuthenticatorMonitorConnectionServerV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLAuthenticatorMonitorConnectionServerV2() *SAMLAuthenticatorMonitorConnectionServerV2 {
	this := SAMLAuthenticatorMonitorConnectionServerV2{}
	return &this
}

// NewSAMLAuthenticatorMonitorConnectionServerV2WithDefaults instantiates a new SAMLAuthenticatorMonitorConnectionServerV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLAuthenticatorMonitorConnectionServerV2WithDefaults() *SAMLAuthenticatorMonitorConnectionServerV2 {
	this := SAMLAuthenticatorMonitorConnectionServerV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field value if set, zero value otherwise.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetLastUpdatedTimestamp() int64 {
	if o == nil || o.LastUpdatedTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimestamp
}

// GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetLastUpdatedTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdatedTimestamp == nil {
		return nil, false
	}
	return o.LastUpdatedTimestamp, true
}

// HasLastUpdatedTimestamp returns a boolean if a field has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasLastUpdatedTimestamp() bool {
	if o != nil && o.LastUpdatedTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdatedTimestamp gets a reference to the given int64 and assigns it to the LastUpdatedTimestamp field.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetLastUpdatedTimestamp(v int64) {
	o.LastUpdatedTimestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetStatus(v string) {
	o.Status = &v
}

// GetThumbprintAccepted returns the ThumbprintAccepted field value if set, zero value otherwise.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetThumbprintAccepted() bool {
	if o == nil || o.ThumbprintAccepted == nil {
		var ret bool
		return ret
	}
	return *o.ThumbprintAccepted
}

// GetThumbprintAcceptedOk returns a tuple with the ThumbprintAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) GetThumbprintAcceptedOk() (*bool, bool) {
	if o == nil || o.ThumbprintAccepted == nil {
		return nil, false
	}
	return o.ThumbprintAccepted, true
}

// HasThumbprintAccepted returns a boolean if a field has been set.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) HasThumbprintAccepted() bool {
	if o != nil && o.ThumbprintAccepted != nil {
		return true
	}

	return false
}

// SetThumbprintAccepted gets a reference to the given bool and assigns it to the ThumbprintAccepted field.
func (o *SAMLAuthenticatorMonitorConnectionServerV2) SetThumbprintAccepted(v bool) {
	o.ThumbprintAccepted = &v
}

func (o SAMLAuthenticatorMonitorConnectionServerV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTimestamp != nil {
		toSerialize["last_updated_timestamp"] = o.LastUpdatedTimestamp
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ThumbprintAccepted != nil {
		toSerialize["thumbprint_accepted"] = o.ThumbprintAccepted
	}
	return json.Marshal(toSerialize)
}

type NullableSAMLAuthenticatorMonitorConnectionServerV2 struct {
	value *SAMLAuthenticatorMonitorConnectionServerV2
	isSet bool
}

func (v NullableSAMLAuthenticatorMonitorConnectionServerV2) Get() *SAMLAuthenticatorMonitorConnectionServerV2 {
	return v.value
}

func (v *NullableSAMLAuthenticatorMonitorConnectionServerV2) Set(val *SAMLAuthenticatorMonitorConnectionServerV2) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLAuthenticatorMonitorConnectionServerV2) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLAuthenticatorMonitorConnectionServerV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLAuthenticatorMonitorConnectionServerV2(val *SAMLAuthenticatorMonitorConnectionServerV2) *NullableSAMLAuthenticatorMonitorConnectionServerV2 {
	return &NullableSAMLAuthenticatorMonitorConnectionServerV2{value: val, isSet: true}
}

func (v NullableSAMLAuthenticatorMonitorConnectionServerV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLAuthenticatorMonitorConnectionServerV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
