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

// SAMLAuthMonitorDetails Details of the SAML authenticator.
type SAMLAuthMonitorDetails struct {
	// The administrator URL for the SAML authenticator.
	AdministratorUrl *string `json:"administrator_url,omitempty"`
	// The label of the SAML Authenticator.
	Label *string `json:"label,omitempty"`
	// The metadata URL of the SAML Authenticator.
	MetadataUrl *string `json:"metadata_url,omitempty"`
}

// NewSAMLAuthMonitorDetails instantiates a new SAMLAuthMonitorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSAMLAuthMonitorDetails() *SAMLAuthMonitorDetails {
	this := SAMLAuthMonitorDetails{}
	return &this
}

// NewSAMLAuthMonitorDetailsWithDefaults instantiates a new SAMLAuthMonitorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSAMLAuthMonitorDetailsWithDefaults() *SAMLAuthMonitorDetails {
	this := SAMLAuthMonitorDetails{}
	return &this
}

// GetAdministratorUrl returns the AdministratorUrl field value if set, zero value otherwise.
func (o *SAMLAuthMonitorDetails) GetAdministratorUrl() string {
	if o == nil || o.AdministratorUrl == nil {
		var ret string
		return ret
	}
	return *o.AdministratorUrl
}

// GetAdministratorUrlOk returns a tuple with the AdministratorUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAuthMonitorDetails) GetAdministratorUrlOk() (*string, bool) {
	if o == nil || o.AdministratorUrl == nil {
		return nil, false
	}
	return o.AdministratorUrl, true
}

// HasAdministratorUrl returns a boolean if a field has been set.
func (o *SAMLAuthMonitorDetails) HasAdministratorUrl() bool {
	if o != nil && o.AdministratorUrl != nil {
		return true
	}

	return false
}

// SetAdministratorUrl gets a reference to the given string and assigns it to the AdministratorUrl field.
func (o *SAMLAuthMonitorDetails) SetAdministratorUrl(v string) {
	o.AdministratorUrl = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *SAMLAuthMonitorDetails) GetLabel() string {
	if o == nil || o.Label == nil {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAuthMonitorDetails) GetLabelOk() (*string, bool) {
	if o == nil || o.Label == nil {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *SAMLAuthMonitorDetails) HasLabel() bool {
	if o != nil && o.Label != nil {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *SAMLAuthMonitorDetails) SetLabel(v string) {
	o.Label = &v
}

// GetMetadataUrl returns the MetadataUrl field value if set, zero value otherwise.
func (o *SAMLAuthMonitorDetails) GetMetadataUrl() string {
	if o == nil || o.MetadataUrl == nil {
		var ret string
		return ret
	}
	return *o.MetadataUrl
}

// GetMetadataUrlOk returns a tuple with the MetadataUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SAMLAuthMonitorDetails) GetMetadataUrlOk() (*string, bool) {
	if o == nil || o.MetadataUrl == nil {
		return nil, false
	}
	return o.MetadataUrl, true
}

// HasMetadataUrl returns a boolean if a field has been set.
func (o *SAMLAuthMonitorDetails) HasMetadataUrl() bool {
	if o != nil && o.MetadataUrl != nil {
		return true
	}

	return false
}

// SetMetadataUrl gets a reference to the given string and assigns it to the MetadataUrl field.
func (o *SAMLAuthMonitorDetails) SetMetadataUrl(v string) {
	o.MetadataUrl = &v
}

func (o SAMLAuthMonitorDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdministratorUrl != nil {
		toSerialize["administrator_url"] = o.AdministratorUrl
	}
	if o.Label != nil {
		toSerialize["label"] = o.Label
	}
	if o.MetadataUrl != nil {
		toSerialize["metadata_url"] = o.MetadataUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSAMLAuthMonitorDetails struct {
	value *SAMLAuthMonitorDetails
	isSet bool
}

func (v NullableSAMLAuthMonitorDetails) Get() *SAMLAuthMonitorDetails {
	return v.value
}

func (v *NullableSAMLAuthMonitorDetails) Set(val *SAMLAuthMonitorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableSAMLAuthMonitorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableSAMLAuthMonitorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSAMLAuthMonitorDetails(val *SAMLAuthMonitorDetails) *NullableSAMLAuthMonitorDetails {
	return &NullableSAMLAuthMonitorDetails{value: val, isSet: true}
}

func (v NullableSAMLAuthMonitorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSAMLAuthMonitorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
