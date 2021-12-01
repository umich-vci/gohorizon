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

// BulkGlobalSessionActionResponseInfo Response corresponding to each pod in a bulk global session action operation.
type BulkGlobalSessionActionResponseInfo struct {
	// List of BulkItemResponseInfo corresponding to each session id in the action operation.
	Details *[]BulkItemResponseInfo `json:"details,omitempty"`
	// Reasons for failure of the operation.
	ErrorMessages *[]string `json:"error_messages,omitempty"`
	// ID of the hosting pod for the sessions.
	PodId *string `json:"pod_id,omitempty"`
	// HTTP Status Code of the operation.
	StatusCode *int32 `json:"status_code,omitempty"`
	// Timestamp in milliseconds when the operation failed.  Measured as epoch time.
	Timestamp *int64 `json:"timestamp,omitempty"`
}

// NewBulkGlobalSessionActionResponseInfo instantiates a new BulkGlobalSessionActionResponseInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkGlobalSessionActionResponseInfo() *BulkGlobalSessionActionResponseInfo {
	this := BulkGlobalSessionActionResponseInfo{}
	return &this
}

// NewBulkGlobalSessionActionResponseInfoWithDefaults instantiates a new BulkGlobalSessionActionResponseInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkGlobalSessionActionResponseInfoWithDefaults() *BulkGlobalSessionActionResponseInfo {
	this := BulkGlobalSessionActionResponseInfo{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *BulkGlobalSessionActionResponseInfo) GetDetails() []BulkItemResponseInfo {
	if o == nil || o.Details == nil {
		var ret []BulkItemResponseInfo
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGlobalSessionActionResponseInfo) GetDetailsOk() (*[]BulkItemResponseInfo, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *BulkGlobalSessionActionResponseInfo) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []BulkItemResponseInfo and assigns it to the Details field.
func (o *BulkGlobalSessionActionResponseInfo) SetDetails(v []BulkItemResponseInfo) {
	o.Details = &v
}

// GetErrorMessages returns the ErrorMessages field value if set, zero value otherwise.
func (o *BulkGlobalSessionActionResponseInfo) GetErrorMessages() []string {
	if o == nil || o.ErrorMessages == nil {
		var ret []string
		return ret
	}
	return *o.ErrorMessages
}

// GetErrorMessagesOk returns a tuple with the ErrorMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGlobalSessionActionResponseInfo) GetErrorMessagesOk() (*[]string, bool) {
	if o == nil || o.ErrorMessages == nil {
		return nil, false
	}
	return o.ErrorMessages, true
}

// HasErrorMessages returns a boolean if a field has been set.
func (o *BulkGlobalSessionActionResponseInfo) HasErrorMessages() bool {
	if o != nil && o.ErrorMessages != nil {
		return true
	}

	return false
}

// SetErrorMessages gets a reference to the given []string and assigns it to the ErrorMessages field.
func (o *BulkGlobalSessionActionResponseInfo) SetErrorMessages(v []string) {
	o.ErrorMessages = &v
}

// GetPodId returns the PodId field value if set, zero value otherwise.
func (o *BulkGlobalSessionActionResponseInfo) GetPodId() string {
	if o == nil || o.PodId == nil {
		var ret string
		return ret
	}
	return *o.PodId
}

// GetPodIdOk returns a tuple with the PodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGlobalSessionActionResponseInfo) GetPodIdOk() (*string, bool) {
	if o == nil || o.PodId == nil {
		return nil, false
	}
	return o.PodId, true
}

// HasPodId returns a boolean if a field has been set.
func (o *BulkGlobalSessionActionResponseInfo) HasPodId() bool {
	if o != nil && o.PodId != nil {
		return true
	}

	return false
}

// SetPodId gets a reference to the given string and assigns it to the PodId field.
func (o *BulkGlobalSessionActionResponseInfo) SetPodId(v string) {
	o.PodId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *BulkGlobalSessionActionResponseInfo) GetStatusCode() int32 {
	if o == nil || o.StatusCode == nil {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGlobalSessionActionResponseInfo) GetStatusCodeOk() (*int32, bool) {
	if o == nil || o.StatusCode == nil {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *BulkGlobalSessionActionResponseInfo) HasStatusCode() bool {
	if o != nil && o.StatusCode != nil {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *BulkGlobalSessionActionResponseInfo) SetStatusCode(v int32) {
	o.StatusCode = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BulkGlobalSessionActionResponseInfo) GetTimestamp() int64 {
	if o == nil || o.Timestamp == nil {
		var ret int64
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkGlobalSessionActionResponseInfo) GetTimestampOk() (*int64, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BulkGlobalSessionActionResponseInfo) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int64 and assigns it to the Timestamp field.
func (o *BulkGlobalSessionActionResponseInfo) SetTimestamp(v int64) {
	o.Timestamp = &v
}

func (o BulkGlobalSessionActionResponseInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.ErrorMessages != nil {
		toSerialize["error_messages"] = o.ErrorMessages
	}
	if o.PodId != nil {
		toSerialize["pod_id"] = o.PodId
	}
	if o.StatusCode != nil {
		toSerialize["status_code"] = o.StatusCode
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableBulkGlobalSessionActionResponseInfo struct {
	value *BulkGlobalSessionActionResponseInfo
	isSet bool
}

func (v NullableBulkGlobalSessionActionResponseInfo) Get() *BulkGlobalSessionActionResponseInfo {
	return v.value
}

func (v *NullableBulkGlobalSessionActionResponseInfo) Set(val *BulkGlobalSessionActionResponseInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkGlobalSessionActionResponseInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkGlobalSessionActionResponseInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkGlobalSessionActionResponseInfo(val *BulkGlobalSessionActionResponseInfo) *NullableBulkGlobalSessionActionResponseInfo {
	return &NullableBulkGlobalSessionActionResponseInfo{value: val, isSet: true}
}

func (v NullableBulkGlobalSessionActionResponseInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkGlobalSessionActionResponseInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
