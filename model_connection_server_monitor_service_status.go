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

// ConnectionServerMonitorServiceStatus Status of Connection Server related Windows service.
type ConnectionServerMonitorServiceStatus struct {
	// Service name of the Connection Server. * PCOIP_SECURE_GATEWAY: PCoIP Secure Gateway service. * BLAST_SECURE_GATEWAY: BLAST Secure Gateway service. * SECURITY_GATEWAY_COMPONENT: Security Gateway Component service.
	ServiceName *string `json:"service_name,omitempty"`
	// Status of the service. * UP: The Windows service is UP and running. * DOWN: The Windows service is not UP. * UNKNOWN: The Windows service state is Unknown.
	Status *string `json:"status,omitempty"`
}

// NewConnectionServerMonitorServiceStatus instantiates a new ConnectionServerMonitorServiceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionServerMonitorServiceStatus() *ConnectionServerMonitorServiceStatus {
	this := ConnectionServerMonitorServiceStatus{}
	return &this
}

// NewConnectionServerMonitorServiceStatusWithDefaults instantiates a new ConnectionServerMonitorServiceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionServerMonitorServiceStatusWithDefaults() *ConnectionServerMonitorServiceStatus {
	this := ConnectionServerMonitorServiceStatus{}
	return &this
}

// GetServiceName returns the ServiceName field value if set, zero value otherwise.
func (o *ConnectionServerMonitorServiceStatus) GetServiceName() string {
	if o == nil || o.ServiceName == nil {
		var ret string
		return ret
	}
	return *o.ServiceName
}

// GetServiceNameOk returns a tuple with the ServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionServerMonitorServiceStatus) GetServiceNameOk() (*string, bool) {
	if o == nil || o.ServiceName == nil {
		return nil, false
	}
	return o.ServiceName, true
}

// HasServiceName returns a boolean if a field has been set.
func (o *ConnectionServerMonitorServiceStatus) HasServiceName() bool {
	if o != nil && o.ServiceName != nil {
		return true
	}

	return false
}

// SetServiceName gets a reference to the given string and assigns it to the ServiceName field.
func (o *ConnectionServerMonitorServiceStatus) SetServiceName(v string) {
	o.ServiceName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConnectionServerMonitorServiceStatus) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionServerMonitorServiceStatus) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConnectionServerMonitorServiceStatus) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConnectionServerMonitorServiceStatus) SetStatus(v string) {
	o.Status = &v
}

func (o ConnectionServerMonitorServiceStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServiceName != nil {
		toSerialize["service_name"] = o.ServiceName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionServerMonitorServiceStatus struct {
	value *ConnectionServerMonitorServiceStatus
	isSet bool
}

func (v NullableConnectionServerMonitorServiceStatus) Get() *ConnectionServerMonitorServiceStatus {
	return v.value
}

func (v *NullableConnectionServerMonitorServiceStatus) Set(val *ConnectionServerMonitorServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionServerMonitorServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionServerMonitorServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionServerMonitorServiceStatus(val *ConnectionServerMonitorServiceStatus) *NullableConnectionServerMonitorServiceStatus {
	return &NullableConnectionServerMonitorServiceStatus{value: val, isSet: true}
}

func (v NullableConnectionServerMonitorServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionServerMonitorServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
