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

// GlobalSessionSendMessageSpec Information required to send message to global sessions.
type GlobalSessionSendMessageSpec struct {
	// Sessions to which message is to be sent.
	GlobalSessionActionSpecs []GlobalSessionActionSpec `json:"global_session_action_specs"`
	// Message to be sent to sessions.
	Message string `json:"message"`
	// Type of message to be sent to sessions. * ERROR: Message is of error type. * WARNING: Message is of warning type. * INFO: Message is of information type.
	MessageType string `json:"message_type"`
}

// NewGlobalSessionSendMessageSpec instantiates a new GlobalSessionSendMessageSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalSessionSendMessageSpec(globalSessionActionSpecs []GlobalSessionActionSpec, message string, messageType string) *GlobalSessionSendMessageSpec {
	this := GlobalSessionSendMessageSpec{}
	this.GlobalSessionActionSpecs = globalSessionActionSpecs
	this.Message = message
	this.MessageType = messageType
	return &this
}

// NewGlobalSessionSendMessageSpecWithDefaults instantiates a new GlobalSessionSendMessageSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalSessionSendMessageSpecWithDefaults() *GlobalSessionSendMessageSpec {
	this := GlobalSessionSendMessageSpec{}
	return &this
}

// GetGlobalSessionActionSpecs returns the GlobalSessionActionSpecs field value
func (o *GlobalSessionSendMessageSpec) GetGlobalSessionActionSpecs() []GlobalSessionActionSpec {
	if o == nil {
		var ret []GlobalSessionActionSpec
		return ret
	}

	return o.GlobalSessionActionSpecs
}

// GetGlobalSessionActionSpecsOk returns a tuple with the GlobalSessionActionSpecs field value
// and a boolean to check if the value has been set.
func (o *GlobalSessionSendMessageSpec) GetGlobalSessionActionSpecsOk() (*[]GlobalSessionActionSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalSessionActionSpecs, true
}

// SetGlobalSessionActionSpecs sets field value
func (o *GlobalSessionSendMessageSpec) SetGlobalSessionActionSpecs(v []GlobalSessionActionSpec) {
	o.GlobalSessionActionSpecs = v
}

// GetMessage returns the Message field value
func (o *GlobalSessionSendMessageSpec) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *GlobalSessionSendMessageSpec) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *GlobalSessionSendMessageSpec) SetMessage(v string) {
	o.Message = v
}

// GetMessageType returns the MessageType field value
func (o *GlobalSessionSendMessageSpec) GetMessageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessageType
}

// GetMessageTypeOk returns a tuple with the MessageType field value
// and a boolean to check if the value has been set.
func (o *GlobalSessionSendMessageSpec) GetMessageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageType, true
}

// SetMessageType sets field value
func (o *GlobalSessionSendMessageSpec) SetMessageType(v string) {
	o.MessageType = v
}

func (o GlobalSessionSendMessageSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["global_session_action_specs"] = o.GlobalSessionActionSpecs
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["message_type"] = o.MessageType
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalSessionSendMessageSpec struct {
	value *GlobalSessionSendMessageSpec
	isSet bool
}

func (v NullableGlobalSessionSendMessageSpec) Get() *GlobalSessionSendMessageSpec {
	return v.value
}

func (v *NullableGlobalSessionSendMessageSpec) Set(val *GlobalSessionSendMessageSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalSessionSendMessageSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalSessionSendMessageSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalSessionSendMessageSpec(val *GlobalSessionSendMessageSpec) *NullableGlobalSessionSendMessageSpec {
	return &NullableGlobalSessionSendMessageSpec{value: val, isSet: true}
}

func (v NullableGlobalSessionSendMessageSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalSessionSendMessageSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
