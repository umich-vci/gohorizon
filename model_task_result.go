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

// TaskResult Information about a task result.
type TaskResult struct {
	// The result message.
	Message *string `json:"message,omitempty"`
	// The result message ID.
	MessageId *string `json:"message_id,omitempty"`
	// The result code of the task. * SUCCESS: Task is completed successfully. * WARN: Task is finished but has warning. * ERROR: Task is finished but the it has error.
	ResultCode *string `json:"result_code,omitempty"`
}

// NewTaskResult instantiates a new TaskResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResult() *TaskResult {
	this := TaskResult{}
	return &this
}

// NewTaskResultWithDefaults instantiates a new TaskResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskResultWithDefaults() *TaskResult {
	this := TaskResult{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TaskResult) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResult) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TaskResult) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TaskResult) SetMessage(v string) {
	o.Message = &v
}

// GetMessageId returns the MessageId field value if set, zero value otherwise.
func (o *TaskResult) GetMessageId() string {
	if o == nil || o.MessageId == nil {
		var ret string
		return ret
	}
	return *o.MessageId
}

// GetMessageIdOk returns a tuple with the MessageId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResult) GetMessageIdOk() (*string, bool) {
	if o == nil || o.MessageId == nil {
		return nil, false
	}
	return o.MessageId, true
}

// HasMessageId returns a boolean if a field has been set.
func (o *TaskResult) HasMessageId() bool {
	if o != nil && o.MessageId != nil {
		return true
	}

	return false
}

// SetMessageId gets a reference to the given string and assigns it to the MessageId field.
func (o *TaskResult) SetMessageId(v string) {
	o.MessageId = &v
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *TaskResult) GetResultCode() string {
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResult) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *TaskResult) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *TaskResult) SetResultCode(v string) {
	o.ResultCode = &v
}

func (o TaskResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.MessageId != nil {
		toSerialize["message_id"] = o.MessageId
	}
	if o.ResultCode != nil {
		toSerialize["result_code"] = o.ResultCode
	}
	return json.Marshal(toSerialize)
}

type NullableTaskResult struct {
	value *TaskResult
	isSet bool
}

func (v NullableTaskResult) Get() *TaskResult {
	return v.value
}

func (v *NullableTaskResult) Set(val *TaskResult) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskResult) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskResult(val *TaskResult) *NullableTaskResult {
	return &NullableTaskResult{value: val, isSet: true}
}

func (v NullableTaskResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
