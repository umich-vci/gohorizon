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

// ConnectionServerMonitorCSReplication Replication status with respect to Peer Connection Servers in the same cluster.
type ConnectionServerMonitorCSReplication struct {
	// Connection Server host name or IP address.
	ServerName *string `json:"server_name,omitempty"`
	// LDAP replication status. * OK: The connection to the Connection Server is working properly. * ERROR: There is a problem with LDAP replication to the Connection Server.
	Status *string `json:"status,omitempty"`
}

// NewConnectionServerMonitorCSReplication instantiates a new ConnectionServerMonitorCSReplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectionServerMonitorCSReplication() *ConnectionServerMonitorCSReplication {
	this := ConnectionServerMonitorCSReplication{}
	return &this
}

// NewConnectionServerMonitorCSReplicationWithDefaults instantiates a new ConnectionServerMonitorCSReplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectionServerMonitorCSReplicationWithDefaults() *ConnectionServerMonitorCSReplication {
	this := ConnectionServerMonitorCSReplication{}
	return &this
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *ConnectionServerMonitorCSReplication) GetServerName() string {
	if o == nil || o.ServerName == nil {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionServerMonitorCSReplication) GetServerNameOk() (*string, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *ConnectionServerMonitorCSReplication) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *ConnectionServerMonitorCSReplication) SetServerName(v string) {
	o.ServerName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ConnectionServerMonitorCSReplication) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectionServerMonitorCSReplication) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ConnectionServerMonitorCSReplication) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ConnectionServerMonitorCSReplication) SetStatus(v string) {
	o.Status = &v
}

func (o ConnectionServerMonitorCSReplication) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServerName != nil {
		toSerialize["server_name"] = o.ServerName
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableConnectionServerMonitorCSReplication struct {
	value *ConnectionServerMonitorCSReplication
	isSet bool
}

func (v NullableConnectionServerMonitorCSReplication) Get() *ConnectionServerMonitorCSReplication {
	return v.value
}

func (v *NullableConnectionServerMonitorCSReplication) Set(val *ConnectionServerMonitorCSReplication) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionServerMonitorCSReplication) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionServerMonitorCSReplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionServerMonitorCSReplication(val *ConnectionServerMonitorCSReplication) *NullableConnectionServerMonitorCSReplication {
	return &NullableConnectionServerMonitorCSReplication{value: val, isSet: true}
}

func (v NullableConnectionServerMonitorCSReplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionServerMonitorCSReplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
