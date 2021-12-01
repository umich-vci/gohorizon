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

// ResumeTaskSpec Information required for resuming a task.
type ResumeTaskSpec struct {
	// Indicates whether to restart the task for virtual machines whose task status is in error state. Default value is false.
	RetryFailedVms *bool `json:"retry_failed_vms,omitempty"`
	// Indicates whether to stop the task at first error. Default value is true.
	StopOnError *bool `json:"stop_on_error,omitempty"`
}

// NewResumeTaskSpec instantiates a new ResumeTaskSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResumeTaskSpec() *ResumeTaskSpec {
	this := ResumeTaskSpec{}
	return &this
}

// NewResumeTaskSpecWithDefaults instantiates a new ResumeTaskSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResumeTaskSpecWithDefaults() *ResumeTaskSpec {
	this := ResumeTaskSpec{}
	return &this
}

// GetRetryFailedVms returns the RetryFailedVms field value if set, zero value otherwise.
func (o *ResumeTaskSpec) GetRetryFailedVms() bool {
	if o == nil || o.RetryFailedVms == nil {
		var ret bool
		return ret
	}
	return *o.RetryFailedVms
}

// GetRetryFailedVmsOk returns a tuple with the RetryFailedVms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeTaskSpec) GetRetryFailedVmsOk() (*bool, bool) {
	if o == nil || o.RetryFailedVms == nil {
		return nil, false
	}
	return o.RetryFailedVms, true
}

// HasRetryFailedVms returns a boolean if a field has been set.
func (o *ResumeTaskSpec) HasRetryFailedVms() bool {
	if o != nil && o.RetryFailedVms != nil {
		return true
	}

	return false
}

// SetRetryFailedVms gets a reference to the given bool and assigns it to the RetryFailedVms field.
func (o *ResumeTaskSpec) SetRetryFailedVms(v bool) {
	o.RetryFailedVms = &v
}

// GetStopOnError returns the StopOnError field value if set, zero value otherwise.
func (o *ResumeTaskSpec) GetStopOnError() bool {
	if o == nil || o.StopOnError == nil {
		var ret bool
		return ret
	}
	return *o.StopOnError
}

// GetStopOnErrorOk returns a tuple with the StopOnError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResumeTaskSpec) GetStopOnErrorOk() (*bool, bool) {
	if o == nil || o.StopOnError == nil {
		return nil, false
	}
	return o.StopOnError, true
}

// HasStopOnError returns a boolean if a field has been set.
func (o *ResumeTaskSpec) HasStopOnError() bool {
	if o != nil && o.StopOnError != nil {
		return true
	}

	return false
}

// SetStopOnError gets a reference to the given bool and assigns it to the StopOnError field.
func (o *ResumeTaskSpec) SetStopOnError(v bool) {
	o.StopOnError = &v
}

func (o ResumeTaskSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RetryFailedVms != nil {
		toSerialize["retry_failed_vms"] = o.RetryFailedVms
	}
	if o.StopOnError != nil {
		toSerialize["stop_on_error"] = o.StopOnError
	}
	return json.Marshal(toSerialize)
}

type NullableResumeTaskSpec struct {
	value *ResumeTaskSpec
	isSet bool
}

func (v NullableResumeTaskSpec) Get() *ResumeTaskSpec {
	return v.value
}

func (v *NullableResumeTaskSpec) Set(val *ResumeTaskSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableResumeTaskSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableResumeTaskSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResumeTaskSpec(val *ResumeTaskSpec) *NullableResumeTaskSpec {
	return &NullableResumeTaskSpec{value: val, isSet: true}
}

func (v NullableResumeTaskSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResumeTaskSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
