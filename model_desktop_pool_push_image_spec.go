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

// DesktopPoolPushImageSpec Specification for the push image operation.
type DesktopPoolPushImageSpec struct {
	// Indicates whether to add Virtual TPM device. Default: false
	AddVirtualTpm *bool `json:"add_virtual_tpm,omitempty"`
	// New image management stream for the desktop pool.<br>Either parent VM and snapshot or image management stream and tag are to be specified.
	ImStreamId *string `json:"im_stream_id,omitempty"`
	// New image management tag for the desktop pool. This must be a tag of the image management stream.
	ImTagId *string `json:"im_tag_id,omitempty"`
	// Determines when to perform the operation on machines which have an active session. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions.
	LogoffPolicy string `json:"logoff_policy"`
	// New base image virtual machine for the desktop pool. This must be in the same datacenter as the base image of the desktop pool.<br>Either parent VM and snapshot or image management stream and tag are to be specified.
	ParentVmId *string `json:"parent_vm_id,omitempty"`
	// New base image snapshot for the desktop pool. This must be a snapshot of the parent VM.
	SnapshotId *string `json:"snapshot_id,omitempty"`
	// When to start the operation. If unset or the time is in the past, the operation will begin immediately. Measured as epoch time.
	StartTime *int64 `json:"start_time,omitempty"`
	// Indicates that the operation should stop on first error. Default: true
	StopOnFirstError *bool `json:"stop_on_first_error,omitempty"`
}

// NewDesktopPoolPushImageSpec instantiates a new DesktopPoolPushImageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolPushImageSpec(logoffPolicy string) *DesktopPoolPushImageSpec {
	this := DesktopPoolPushImageSpec{}
	this.LogoffPolicy = logoffPolicy
	return &this
}

// NewDesktopPoolPushImageSpecWithDefaults instantiates a new DesktopPoolPushImageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolPushImageSpecWithDefaults() *DesktopPoolPushImageSpec {
	this := DesktopPoolPushImageSpec{}
	return &this
}

// GetAddVirtualTpm returns the AddVirtualTpm field value if set, zero value otherwise.
func (o *DesktopPoolPushImageSpec) GetAddVirtualTpm() bool {
	if o == nil || o.AddVirtualTpm == nil {
		var ret bool
		return ret
	}
	return *o.AddVirtualTpm
}

// GetAddVirtualTpmOk returns a tuple with the AddVirtualTpm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolPushImageSpec) GetAddVirtualTpmOk() (*bool, bool) {
	if o == nil || o.AddVirtualTpm == nil {
		return nil, false
	}
	return o.AddVirtualTpm, true
}

// HasAddVirtualTpm returns a boolean if a field has been set.
func (o *DesktopPoolPushImageSpec) HasAddVirtualTpm() bool {
	if o != nil && o.AddVirtualTpm != nil {
		return true
	}

	return false
}

// SetAddVirtualTpm gets a reference to the given bool and assigns it to the AddVirtualTpm field.
func (o *DesktopPoolPushImageSpec) SetAddVirtualTpm(v bool) {
	o.AddVirtualTpm = &v
}

// GetImStreamId returns the ImStreamId field value if set, zero value otherwise.
func (o *DesktopPoolPushImageSpec) GetImStreamId() string {
	if o == nil || o.ImStreamId == nil {
		var ret string
		return ret
	}
	return *o.ImStreamId
}

// GetImStreamIdOk returns a tuple with the ImStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolPushImageSpec) GetImStreamIdOk() (*string, bool) {
	if o == nil || o.ImStreamId == nil {
		return nil, false
	}
	return o.ImStreamId, true
}

// HasImStreamId returns a boolean if a field has been set.
func (o *DesktopPoolPushImageSpec) HasImStreamId() bool {
	if o != nil && o.ImStreamId != nil {
		return true
	}

	return false
}

// SetImStreamId gets a reference to the given string and assigns it to the ImStreamId field.
func (o *DesktopPoolPushImageSpec) SetImStreamId(v string) {
	o.ImStreamId = &v
}

// GetImTagId returns the ImTagId field value if set, zero value otherwise.
func (o *DesktopPoolPushImageSpec) GetImTagId() string {
	if o == nil || o.ImTagId == nil {
		var ret string
		return ret
	}
	return *o.ImTagId
}

// GetImTagIdOk returns a tuple with the ImTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolPushImageSpec) GetImTagIdOk() (*string, bool) {
	if o == nil || o.ImTagId == nil {
		return nil, false
	}
	return o.ImTagId, true
}

// HasImTagId returns a boolean if a field has been set.
func (o *DesktopPoolPushImageSpec) HasImTagId() bool {
	if o != nil && o.ImTagId != nil {
		return true
	}

	return false
}

// SetImTagId gets a reference to the given string and assigns it to the ImTagId field.
func (o *DesktopPoolPushImageSpec) SetImTagId(v string) {
	o.ImTagId = &v
}

// GetLogoffPolicy returns the LogoffPolicy field value
func (o *DesktopPoolPushImageSpec) GetLogoffPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LogoffPolicy
}

// GetLogoffPolicyOk returns a tuple with the LogoffPolicy field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolPushImageSpec) GetLogoffPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LogoffPolicy, true
}

// SetLogoffPolicy sets field value
func (o *DesktopPoolPushImageSpec) SetLogoffPolicy(v string) {
	o.LogoffPolicy = v
}

// GetParentVmId returns the ParentVmId field value if set, zero value otherwise.
func (o *DesktopPoolPushImageSpec) GetParentVmId() string {
	if o == nil || o.ParentVmId == nil {
		var ret string
		return ret
	}
	return *o.ParentVmId
}

// GetParentVmIdOk returns a tuple with the ParentVmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolPushImageSpec) GetParentVmIdOk() (*string, bool) {
	if o == nil || o.ParentVmId == nil {
		return nil, false
	}
	return o.ParentVmId, true
}

// HasParentVmId returns a boolean if a field has been set.
func (o *DesktopPoolPushImageSpec) HasParentVmId() bool {
	if o != nil && o.ParentVmId != nil {
		return true
	}

	return false
}

// SetParentVmId gets a reference to the given string and assigns it to the ParentVmId field.
func (o *DesktopPoolPushImageSpec) SetParentVmId(v string) {
	o.ParentVmId = &v
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *DesktopPoolPushImageSpec) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolPushImageSpec) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *DesktopPoolPushImageSpec) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *DesktopPoolPushImageSpec) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *DesktopPoolPushImageSpec) GetStartTime() int64 {
	if o == nil || o.StartTime == nil {
		var ret int64
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolPushImageSpec) GetStartTimeOk() (*int64, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *DesktopPoolPushImageSpec) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given int64 and assigns it to the StartTime field.
func (o *DesktopPoolPushImageSpec) SetStartTime(v int64) {
	o.StartTime = &v
}

// GetStopOnFirstError returns the StopOnFirstError field value if set, zero value otherwise.
func (o *DesktopPoolPushImageSpec) GetStopOnFirstError() bool {
	if o == nil || o.StopOnFirstError == nil {
		var ret bool
		return ret
	}
	return *o.StopOnFirstError
}

// GetStopOnFirstErrorOk returns a tuple with the StopOnFirstError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolPushImageSpec) GetStopOnFirstErrorOk() (*bool, bool) {
	if o == nil || o.StopOnFirstError == nil {
		return nil, false
	}
	return o.StopOnFirstError, true
}

// HasStopOnFirstError returns a boolean if a field has been set.
func (o *DesktopPoolPushImageSpec) HasStopOnFirstError() bool {
	if o != nil && o.StopOnFirstError != nil {
		return true
	}

	return false
}

// SetStopOnFirstError gets a reference to the given bool and assigns it to the StopOnFirstError field.
func (o *DesktopPoolPushImageSpec) SetStopOnFirstError(v bool) {
	o.StopOnFirstError = &v
}

func (o DesktopPoolPushImageSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddVirtualTpm != nil {
		toSerialize["add_virtual_tpm"] = o.AddVirtualTpm
	}
	if o.ImStreamId != nil {
		toSerialize["im_stream_id"] = o.ImStreamId
	}
	if o.ImTagId != nil {
		toSerialize["im_tag_id"] = o.ImTagId
	}
	if true {
		toSerialize["logoff_policy"] = o.LogoffPolicy
	}
	if o.ParentVmId != nil {
		toSerialize["parent_vm_id"] = o.ParentVmId
	}
	if o.SnapshotId != nil {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.StopOnFirstError != nil {
		toSerialize["stop_on_first_error"] = o.StopOnFirstError
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolPushImageSpec struct {
	value *DesktopPoolPushImageSpec
	isSet bool
}

func (v NullableDesktopPoolPushImageSpec) Get() *DesktopPoolPushImageSpec {
	return v.value
}

func (v *NullableDesktopPoolPushImageSpec) Set(val *DesktopPoolPushImageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolPushImageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolPushImageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolPushImageSpec(val *DesktopPoolPushImageSpec) *NullableDesktopPoolPushImageSpec {
	return &NullableDesktopPoolPushImageSpec{value: val, isSet: true}
}

func (v NullableDesktopPoolPushImageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolPushImageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
