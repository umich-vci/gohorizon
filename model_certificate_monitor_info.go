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

// CertificateMonitorInfo Monitoring data about server's certificate.
type CertificateMonitorInfo struct {
	// Indicates if the certificate is valid.
	Valid *bool `json:"valid,omitempty"`
	// Start time of the certificate validity in milliseconds. Measured as epoch time.
	ValidFrom *int64 `json:"valid_from,omitempty"`
	// Expiration time of the certificate validity in milliseconds. Measured as epoch time.
	ValidTo *int64 `json:"valid_to,omitempty"`
}

// NewCertificateMonitorInfo instantiates a new CertificateMonitorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCertificateMonitorInfo() *CertificateMonitorInfo {
	this := CertificateMonitorInfo{}
	return &this
}

// NewCertificateMonitorInfoWithDefaults instantiates a new CertificateMonitorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCertificateMonitorInfoWithDefaults() *CertificateMonitorInfo {
	this := CertificateMonitorInfo{}
	return &this
}

// GetValid returns the Valid field value if set, zero value otherwise.
func (o *CertificateMonitorInfo) GetValid() bool {
	if o == nil || o.Valid == nil {
		var ret bool
		return ret
	}
	return *o.Valid
}

// GetValidOk returns a tuple with the Valid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMonitorInfo) GetValidOk() (*bool, bool) {
	if o == nil || o.Valid == nil {
		return nil, false
	}
	return o.Valid, true
}

// HasValid returns a boolean if a field has been set.
func (o *CertificateMonitorInfo) HasValid() bool {
	if o != nil && o.Valid != nil {
		return true
	}

	return false
}

// SetValid gets a reference to the given bool and assigns it to the Valid field.
func (o *CertificateMonitorInfo) SetValid(v bool) {
	o.Valid = &v
}

// GetValidFrom returns the ValidFrom field value if set, zero value otherwise.
func (o *CertificateMonitorInfo) GetValidFrom() int64 {
	if o == nil || o.ValidFrom == nil {
		var ret int64
		return ret
	}
	return *o.ValidFrom
}

// GetValidFromOk returns a tuple with the ValidFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMonitorInfo) GetValidFromOk() (*int64, bool) {
	if o == nil || o.ValidFrom == nil {
		return nil, false
	}
	return o.ValidFrom, true
}

// HasValidFrom returns a boolean if a field has been set.
func (o *CertificateMonitorInfo) HasValidFrom() bool {
	if o != nil && o.ValidFrom != nil {
		return true
	}

	return false
}

// SetValidFrom gets a reference to the given int64 and assigns it to the ValidFrom field.
func (o *CertificateMonitorInfo) SetValidFrom(v int64) {
	o.ValidFrom = &v
}

// GetValidTo returns the ValidTo field value if set, zero value otherwise.
func (o *CertificateMonitorInfo) GetValidTo() int64 {
	if o == nil || o.ValidTo == nil {
		var ret int64
		return ret
	}
	return *o.ValidTo
}

// GetValidToOk returns a tuple with the ValidTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CertificateMonitorInfo) GetValidToOk() (*int64, bool) {
	if o == nil || o.ValidTo == nil {
		return nil, false
	}
	return o.ValidTo, true
}

// HasValidTo returns a boolean if a field has been set.
func (o *CertificateMonitorInfo) HasValidTo() bool {
	if o != nil && o.ValidTo != nil {
		return true
	}

	return false
}

// SetValidTo gets a reference to the given int64 and assigns it to the ValidTo field.
func (o *CertificateMonitorInfo) SetValidTo(v int64) {
	o.ValidTo = &v
}

func (o CertificateMonitorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Valid != nil {
		toSerialize["valid"] = o.Valid
	}
	if o.ValidFrom != nil {
		toSerialize["valid_from"] = o.ValidFrom
	}
	if o.ValidTo != nil {
		toSerialize["valid_to"] = o.ValidTo
	}
	return json.Marshal(toSerialize)
}

type NullableCertificateMonitorInfo struct {
	value *CertificateMonitorInfo
	isSet bool
}

func (v NullableCertificateMonitorInfo) Get() *CertificateMonitorInfo {
	return v.value
}

func (v *NullableCertificateMonitorInfo) Set(val *CertificateMonitorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCertificateMonitorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCertificateMonitorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCertificateMonitorInfo(val *CertificateMonitorInfo) *NullableCertificateMonitorInfo {
	return &NullableCertificateMonitorInfo{value: val, isSet: true}
}

func (v NullableCertificateMonitorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCertificateMonitorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
