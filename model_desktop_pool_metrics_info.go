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

// DesktopPoolMetricsInfo Metrics related to desktop pool.
type DesktopPoolMetricsInfo struct {
	// Unique ID of the desktop pool.
	Id *string `json:"id,omitempty"`
	// Number of connected sessions of the desktop pool.
	NumConnectedSessions *int32 `json:"num_connected_sessions,omitempty"`
	// Number of machines in the desktop pool.
	NumMachines *int32 `json:"num_machines,omitempty"`
	// Occupancy count for the desktop pool. * For dedicated assignment desktop, it is the number of assigned machine count. * For floating assignment desktop, it is the summation of the connected and disconnected sessions.
	OccupancyCount *int32 `json:"occupancy_count,omitempty"`
}

// NewDesktopPoolMetricsInfo instantiates a new DesktopPoolMetricsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolMetricsInfo() *DesktopPoolMetricsInfo {
	this := DesktopPoolMetricsInfo{}
	return &this
}

// NewDesktopPoolMetricsInfoWithDefaults instantiates a new DesktopPoolMetricsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolMetricsInfoWithDefaults() *DesktopPoolMetricsInfo {
	this := DesktopPoolMetricsInfo{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DesktopPoolMetricsInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolMetricsInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DesktopPoolMetricsInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DesktopPoolMetricsInfo) SetId(v string) {
	o.Id = &v
}

// GetNumConnectedSessions returns the NumConnectedSessions field value if set, zero value otherwise.
func (o *DesktopPoolMetricsInfo) GetNumConnectedSessions() int32 {
	if o == nil || o.NumConnectedSessions == nil {
		var ret int32
		return ret
	}
	return *o.NumConnectedSessions
}

// GetNumConnectedSessionsOk returns a tuple with the NumConnectedSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolMetricsInfo) GetNumConnectedSessionsOk() (*int32, bool) {
	if o == nil || o.NumConnectedSessions == nil {
		return nil, false
	}
	return o.NumConnectedSessions, true
}

// HasNumConnectedSessions returns a boolean if a field has been set.
func (o *DesktopPoolMetricsInfo) HasNumConnectedSessions() bool {
	if o != nil && o.NumConnectedSessions != nil {
		return true
	}

	return false
}

// SetNumConnectedSessions gets a reference to the given int32 and assigns it to the NumConnectedSessions field.
func (o *DesktopPoolMetricsInfo) SetNumConnectedSessions(v int32) {
	o.NumConnectedSessions = &v
}

// GetNumMachines returns the NumMachines field value if set, zero value otherwise.
func (o *DesktopPoolMetricsInfo) GetNumMachines() int32 {
	if o == nil || o.NumMachines == nil {
		var ret int32
		return ret
	}
	return *o.NumMachines
}

// GetNumMachinesOk returns a tuple with the NumMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolMetricsInfo) GetNumMachinesOk() (*int32, bool) {
	if o == nil || o.NumMachines == nil {
		return nil, false
	}
	return o.NumMachines, true
}

// HasNumMachines returns a boolean if a field has been set.
func (o *DesktopPoolMetricsInfo) HasNumMachines() bool {
	if o != nil && o.NumMachines != nil {
		return true
	}

	return false
}

// SetNumMachines gets a reference to the given int32 and assigns it to the NumMachines field.
func (o *DesktopPoolMetricsInfo) SetNumMachines(v int32) {
	o.NumMachines = &v
}

// GetOccupancyCount returns the OccupancyCount field value if set, zero value otherwise.
func (o *DesktopPoolMetricsInfo) GetOccupancyCount() int32 {
	if o == nil || o.OccupancyCount == nil {
		var ret int32
		return ret
	}
	return *o.OccupancyCount
}

// GetOccupancyCountOk returns a tuple with the OccupancyCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolMetricsInfo) GetOccupancyCountOk() (*int32, bool) {
	if o == nil || o.OccupancyCount == nil {
		return nil, false
	}
	return o.OccupancyCount, true
}

// HasOccupancyCount returns a boolean if a field has been set.
func (o *DesktopPoolMetricsInfo) HasOccupancyCount() bool {
	if o != nil && o.OccupancyCount != nil {
		return true
	}

	return false
}

// SetOccupancyCount gets a reference to the given int32 and assigns it to the OccupancyCount field.
func (o *DesktopPoolMetricsInfo) SetOccupancyCount(v int32) {
	o.OccupancyCount = &v
}

func (o DesktopPoolMetricsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.NumConnectedSessions != nil {
		toSerialize["num_connected_sessions"] = o.NumConnectedSessions
	}
	if o.NumMachines != nil {
		toSerialize["num_machines"] = o.NumMachines
	}
	if o.OccupancyCount != nil {
		toSerialize["occupancy_count"] = o.OccupancyCount
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolMetricsInfo struct {
	value *DesktopPoolMetricsInfo
	isSet bool
}

func (v NullableDesktopPoolMetricsInfo) Get() *DesktopPoolMetricsInfo {
	return v.value
}

func (v *NullableDesktopPoolMetricsInfo) Set(val *DesktopPoolMetricsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolMetricsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolMetricsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolMetricsInfo(val *DesktopPoolMetricsInfo) *NullableDesktopPoolMetricsInfo {
	return &NullableDesktopPoolMetricsInfo{value: val, isSet: true}
}

func (v NullableDesktopPoolMetricsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolMetricsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
