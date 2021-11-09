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

// RDSHLoadBalancerSettingsInfo RDSH load balancer settings for the farm.
type RDSHLoadBalancerSettingsInfo struct {
	// Represents threshold of CPU usage, in percentage. If the value is 0, then this metric is not considered for load balancing.
	CpuThreshold *int32 `json:"cpu_threshold,omitempty"`
	// Represents the threshold of average number of both read and write requests that were queued for the selected disk during the sample interval. If the value is 0, then this metric is not considered for load balancing.
	DiskQueueLengthThreshold *int32 `json:"disk_queue_length_threshold,omitempty"`
	// Represents the threshold of average time, in milliseconds, of a read of data from the disk. If the value is 0, then this metric is not considered for load balancing.
	DiskReadLatencyThreshold *int32 `json:"disk_read_latency_threshold,omitempty"`
	// Represents the threshold of average time, in milliseconds, of a write of data to the disk. If the value is 0, then this metric is not considered for load balancing.
	DiskWriteLatencyThreshold *int32 `json:"disk_write_latency_threshold,omitempty"`
	// Indicates whether to include session count for load balancing.
	IncludeSessionCount *bool `json:"include_session_count,omitempty"`
	// Represents threshold of memory usage, in percentage. If the value is 0, then this metric is not considered for load balancing.
	MemoryThreshold *int32 `json:"memory_threshold,omitempty"`
}

// NewRDSHLoadBalancerSettingsInfo instantiates a new RDSHLoadBalancerSettingsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRDSHLoadBalancerSettingsInfo() *RDSHLoadBalancerSettingsInfo {
	this := RDSHLoadBalancerSettingsInfo{}
	return &this
}

// NewRDSHLoadBalancerSettingsInfoWithDefaults instantiates a new RDSHLoadBalancerSettingsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRDSHLoadBalancerSettingsInfoWithDefaults() *RDSHLoadBalancerSettingsInfo {
	this := RDSHLoadBalancerSettingsInfo{}
	return &this
}

// GetCpuThreshold returns the CpuThreshold field value if set, zero value otherwise.
func (o *RDSHLoadBalancerSettingsInfo) GetCpuThreshold() int32 {
	if o == nil || o.CpuThreshold == nil {
		var ret int32
		return ret
	}
	return *o.CpuThreshold
}

// GetCpuThresholdOk returns a tuple with the CpuThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSHLoadBalancerSettingsInfo) GetCpuThresholdOk() (*int32, bool) {
	if o == nil || o.CpuThreshold == nil {
		return nil, false
	}
	return o.CpuThreshold, true
}

// HasCpuThreshold returns a boolean if a field has been set.
func (o *RDSHLoadBalancerSettingsInfo) HasCpuThreshold() bool {
	if o != nil && o.CpuThreshold != nil {
		return true
	}

	return false
}

// SetCpuThreshold gets a reference to the given int32 and assigns it to the CpuThreshold field.
func (o *RDSHLoadBalancerSettingsInfo) SetCpuThreshold(v int32) {
	o.CpuThreshold = &v
}

// GetDiskQueueLengthThreshold returns the DiskQueueLengthThreshold field value if set, zero value otherwise.
func (o *RDSHLoadBalancerSettingsInfo) GetDiskQueueLengthThreshold() int32 {
	if o == nil || o.DiskQueueLengthThreshold == nil {
		var ret int32
		return ret
	}
	return *o.DiskQueueLengthThreshold
}

// GetDiskQueueLengthThresholdOk returns a tuple with the DiskQueueLengthThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSHLoadBalancerSettingsInfo) GetDiskQueueLengthThresholdOk() (*int32, bool) {
	if o == nil || o.DiskQueueLengthThreshold == nil {
		return nil, false
	}
	return o.DiskQueueLengthThreshold, true
}

// HasDiskQueueLengthThreshold returns a boolean if a field has been set.
func (o *RDSHLoadBalancerSettingsInfo) HasDiskQueueLengthThreshold() bool {
	if o != nil && o.DiskQueueLengthThreshold != nil {
		return true
	}

	return false
}

// SetDiskQueueLengthThreshold gets a reference to the given int32 and assigns it to the DiskQueueLengthThreshold field.
func (o *RDSHLoadBalancerSettingsInfo) SetDiskQueueLengthThreshold(v int32) {
	o.DiskQueueLengthThreshold = &v
}

// GetDiskReadLatencyThreshold returns the DiskReadLatencyThreshold field value if set, zero value otherwise.
func (o *RDSHLoadBalancerSettingsInfo) GetDiskReadLatencyThreshold() int32 {
	if o == nil || o.DiskReadLatencyThreshold == nil {
		var ret int32
		return ret
	}
	return *o.DiskReadLatencyThreshold
}

// GetDiskReadLatencyThresholdOk returns a tuple with the DiskReadLatencyThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSHLoadBalancerSettingsInfo) GetDiskReadLatencyThresholdOk() (*int32, bool) {
	if o == nil || o.DiskReadLatencyThreshold == nil {
		return nil, false
	}
	return o.DiskReadLatencyThreshold, true
}

// HasDiskReadLatencyThreshold returns a boolean if a field has been set.
func (o *RDSHLoadBalancerSettingsInfo) HasDiskReadLatencyThreshold() bool {
	if o != nil && o.DiskReadLatencyThreshold != nil {
		return true
	}

	return false
}

// SetDiskReadLatencyThreshold gets a reference to the given int32 and assigns it to the DiskReadLatencyThreshold field.
func (o *RDSHLoadBalancerSettingsInfo) SetDiskReadLatencyThreshold(v int32) {
	o.DiskReadLatencyThreshold = &v
}

// GetDiskWriteLatencyThreshold returns the DiskWriteLatencyThreshold field value if set, zero value otherwise.
func (o *RDSHLoadBalancerSettingsInfo) GetDiskWriteLatencyThreshold() int32 {
	if o == nil || o.DiskWriteLatencyThreshold == nil {
		var ret int32
		return ret
	}
	return *o.DiskWriteLatencyThreshold
}

// GetDiskWriteLatencyThresholdOk returns a tuple with the DiskWriteLatencyThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSHLoadBalancerSettingsInfo) GetDiskWriteLatencyThresholdOk() (*int32, bool) {
	if o == nil || o.DiskWriteLatencyThreshold == nil {
		return nil, false
	}
	return o.DiskWriteLatencyThreshold, true
}

// HasDiskWriteLatencyThreshold returns a boolean if a field has been set.
func (o *RDSHLoadBalancerSettingsInfo) HasDiskWriteLatencyThreshold() bool {
	if o != nil && o.DiskWriteLatencyThreshold != nil {
		return true
	}

	return false
}

// SetDiskWriteLatencyThreshold gets a reference to the given int32 and assigns it to the DiskWriteLatencyThreshold field.
func (o *RDSHLoadBalancerSettingsInfo) SetDiskWriteLatencyThreshold(v int32) {
	o.DiskWriteLatencyThreshold = &v
}

// GetIncludeSessionCount returns the IncludeSessionCount field value if set, zero value otherwise.
func (o *RDSHLoadBalancerSettingsInfo) GetIncludeSessionCount() bool {
	if o == nil || o.IncludeSessionCount == nil {
		var ret bool
		return ret
	}
	return *o.IncludeSessionCount
}

// GetIncludeSessionCountOk returns a tuple with the IncludeSessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSHLoadBalancerSettingsInfo) GetIncludeSessionCountOk() (*bool, bool) {
	if o == nil || o.IncludeSessionCount == nil {
		return nil, false
	}
	return o.IncludeSessionCount, true
}

// HasIncludeSessionCount returns a boolean if a field has been set.
func (o *RDSHLoadBalancerSettingsInfo) HasIncludeSessionCount() bool {
	if o != nil && o.IncludeSessionCount != nil {
		return true
	}

	return false
}

// SetIncludeSessionCount gets a reference to the given bool and assigns it to the IncludeSessionCount field.
func (o *RDSHLoadBalancerSettingsInfo) SetIncludeSessionCount(v bool) {
	o.IncludeSessionCount = &v
}

// GetMemoryThreshold returns the MemoryThreshold field value if set, zero value otherwise.
func (o *RDSHLoadBalancerSettingsInfo) GetMemoryThreshold() int32 {
	if o == nil || o.MemoryThreshold == nil {
		var ret int32
		return ret
	}
	return *o.MemoryThreshold
}

// GetMemoryThresholdOk returns a tuple with the MemoryThreshold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSHLoadBalancerSettingsInfo) GetMemoryThresholdOk() (*int32, bool) {
	if o == nil || o.MemoryThreshold == nil {
		return nil, false
	}
	return o.MemoryThreshold, true
}

// HasMemoryThreshold returns a boolean if a field has been set.
func (o *RDSHLoadBalancerSettingsInfo) HasMemoryThreshold() bool {
	if o != nil && o.MemoryThreshold != nil {
		return true
	}

	return false
}

// SetMemoryThreshold gets a reference to the given int32 and assigns it to the MemoryThreshold field.
func (o *RDSHLoadBalancerSettingsInfo) SetMemoryThreshold(v int32) {
	o.MemoryThreshold = &v
}

func (o RDSHLoadBalancerSettingsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CpuThreshold != nil {
		toSerialize["cpu_threshold"] = o.CpuThreshold
	}
	if o.DiskQueueLengthThreshold != nil {
		toSerialize["disk_queue_length_threshold"] = o.DiskQueueLengthThreshold
	}
	if o.DiskReadLatencyThreshold != nil {
		toSerialize["disk_read_latency_threshold"] = o.DiskReadLatencyThreshold
	}
	if o.DiskWriteLatencyThreshold != nil {
		toSerialize["disk_write_latency_threshold"] = o.DiskWriteLatencyThreshold
	}
	if o.IncludeSessionCount != nil {
		toSerialize["include_session_count"] = o.IncludeSessionCount
	}
	if o.MemoryThreshold != nil {
		toSerialize["memory_threshold"] = o.MemoryThreshold
	}
	return json.Marshal(toSerialize)
}

type NullableRDSHLoadBalancerSettingsInfo struct {
	value *RDSHLoadBalancerSettingsInfo
	isSet bool
}

func (v NullableRDSHLoadBalancerSettingsInfo) Get() *RDSHLoadBalancerSettingsInfo {
	return v.value
}

func (v *NullableRDSHLoadBalancerSettingsInfo) Set(val *RDSHLoadBalancerSettingsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRDSHLoadBalancerSettingsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRDSHLoadBalancerSettingsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRDSHLoadBalancerSettingsInfo(val *RDSHLoadBalancerSettingsInfo) *NullableRDSHLoadBalancerSettingsInfo {
	return &NullableRDSHLoadBalancerSettingsInfo{value: val, isSet: true}
}

func (v NullableRDSHLoadBalancerSettingsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRDSHLoadBalancerSettingsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
