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

// EventDatabaseMonitorDetails struct for EventDatabaseMonitorDetails
type EventDatabaseMonitorDetails struct {
	// The name of the database.
	DatabaseName string `json:"database_name"`
	// The port of the database server.
	Port int32 `json:"port"`
	// The prefix for event tables in the database.
	Prefix string `json:"prefix"`
	// The name or ip address of the database server.
	ServerName string `json:"server_name"`
	// The type of the database. * ORACLE: An Oracle database. * SQL_SERVER: A SQL server database. * POSTGRESQL: A PostgreSQL database.
	Type string `json:"type"`
	// The username used to connect to the database.
	UserName string `json:"user_name"`
}

// NewEventDatabaseMonitorDetails instantiates a new EventDatabaseMonitorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventDatabaseMonitorDetails(databaseName string, port int32, prefix string, serverName string, type_ string, userName string) *EventDatabaseMonitorDetails {
	this := EventDatabaseMonitorDetails{}
	this.DatabaseName = databaseName
	this.Port = port
	this.Prefix = prefix
	this.ServerName = serverName
	this.Type = type_
	this.UserName = userName
	return &this
}

// NewEventDatabaseMonitorDetailsWithDefaults instantiates a new EventDatabaseMonitorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventDatabaseMonitorDetailsWithDefaults() *EventDatabaseMonitorDetails {
	this := EventDatabaseMonitorDetails{}
	return &this
}

// GetDatabaseName returns the DatabaseName field value
func (o *EventDatabaseMonitorDetails) GetDatabaseName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatabaseName
}

// GetDatabaseNameOk returns a tuple with the DatabaseName field value
// and a boolean to check if the value has been set.
func (o *EventDatabaseMonitorDetails) GetDatabaseNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatabaseName, true
}

// SetDatabaseName sets field value
func (o *EventDatabaseMonitorDetails) SetDatabaseName(v string) {
	o.DatabaseName = v
}

// GetPort returns the Port field value
func (o *EventDatabaseMonitorDetails) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *EventDatabaseMonitorDetails) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *EventDatabaseMonitorDetails) SetPort(v int32) {
	o.Port = v
}

// GetPrefix returns the Prefix field value
func (o *EventDatabaseMonitorDetails) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *EventDatabaseMonitorDetails) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *EventDatabaseMonitorDetails) SetPrefix(v string) {
	o.Prefix = v
}

// GetServerName returns the ServerName field value
func (o *EventDatabaseMonitorDetails) GetServerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value
// and a boolean to check if the value has been set.
func (o *EventDatabaseMonitorDetails) GetServerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServerName, true
}

// SetServerName sets field value
func (o *EventDatabaseMonitorDetails) SetServerName(v string) {
	o.ServerName = v
}

// GetType returns the Type field value
func (o *EventDatabaseMonitorDetails) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *EventDatabaseMonitorDetails) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *EventDatabaseMonitorDetails) SetType(v string) {
	o.Type = v
}

// GetUserName returns the UserName field value
func (o *EventDatabaseMonitorDetails) GetUserName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value
// and a boolean to check if the value has been set.
func (o *EventDatabaseMonitorDetails) GetUserNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserName, true
}

// SetUserName sets field value
func (o *EventDatabaseMonitorDetails) SetUserName(v string) {
	o.UserName = v
}

func (o EventDatabaseMonitorDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["database_name"] = o.DatabaseName
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	if true {
		toSerialize["server_name"] = o.ServerName
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["user_name"] = o.UserName
	}
	return json.Marshal(toSerialize)
}

type NullableEventDatabaseMonitorDetails struct {
	value *EventDatabaseMonitorDetails
	isSet bool
}

func (v NullableEventDatabaseMonitorDetails) Get() *EventDatabaseMonitorDetails {
	return v.value
}

func (v *NullableEventDatabaseMonitorDetails) Set(val *EventDatabaseMonitorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableEventDatabaseMonitorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableEventDatabaseMonitorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventDatabaseMonitorDetails(val *EventDatabaseMonitorDetails) *NullableEventDatabaseMonitorDetails {
	return &NullableEventDatabaseMonitorDetails{value: val, isSet: true}
}

func (v NullableEventDatabaseMonitorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventDatabaseMonitorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
