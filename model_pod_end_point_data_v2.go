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

// PodEndPointDataV2 struct for PodEndPointDataV2
type PodEndPointDataV2 struct {
	// Indicates whether an endpoint is enabled. A disabled endpoint will be excluded from participating inter-pod communication.
	Enabled *bool `json:"enabled,omitempty"`
	// Unique ID for a pod endpoint.
	Id *string `json:"id,omitempty"`
	// The timestamp in milliseconds when the last update was obtained. Measured as epoch time.
	LastUpdatedTimestamp *int64 `json:"last_updated_timestamp,omitempty"`
	// Name for the pod endpoint.
	Name *string `json:"name,omitempty"`
	// Round trip time (in milliseconds) for ping request between the local pod endpoint and the remote pod.
	RoundtripTime *int64 `json:"roundtrip_time,omitempty"`
	// Status of the pod endpoint. * ONLINE: Pod endpoint is online and functional. * UNCHECKED: Pod endpoint was offline and it just came back online but the system has not verified that it is functional. * OFFLINE: Pod endpoint is offline or unreachable.
	Status *string `json:"status,omitempty"`
	// The URL for the pod endpoint. This address and special port will be used for inter-pod communication.
	Url *string `json:"url,omitempty"`
}

// NewPodEndPointDataV2 instantiates a new PodEndPointDataV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodEndPointDataV2() *PodEndPointDataV2 {
	this := PodEndPointDataV2{}
	return &this
}

// NewPodEndPointDataV2WithDefaults instantiates a new PodEndPointDataV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodEndPointDataV2WithDefaults() *PodEndPointDataV2 {
	this := PodEndPointDataV2{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *PodEndPointDataV2) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodEndPointDataV2) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *PodEndPointDataV2) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *PodEndPointDataV2) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PodEndPointDataV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodEndPointDataV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PodEndPointDataV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PodEndPointDataV2) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field value if set, zero value otherwise.
func (o *PodEndPointDataV2) GetLastUpdatedTimestamp() int64 {
	if o == nil || o.LastUpdatedTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimestamp
}

// GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodEndPointDataV2) GetLastUpdatedTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdatedTimestamp == nil {
		return nil, false
	}
	return o.LastUpdatedTimestamp, true
}

// HasLastUpdatedTimestamp returns a boolean if a field has been set.
func (o *PodEndPointDataV2) HasLastUpdatedTimestamp() bool {
	if o != nil && o.LastUpdatedTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdatedTimestamp gets a reference to the given int64 and assigns it to the LastUpdatedTimestamp field.
func (o *PodEndPointDataV2) SetLastUpdatedTimestamp(v int64) {
	o.LastUpdatedTimestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PodEndPointDataV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodEndPointDataV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PodEndPointDataV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PodEndPointDataV2) SetName(v string) {
	o.Name = &v
}

// GetRoundtripTime returns the RoundtripTime field value if set, zero value otherwise.
func (o *PodEndPointDataV2) GetRoundtripTime() int64 {
	if o == nil || o.RoundtripTime == nil {
		var ret int64
		return ret
	}
	return *o.RoundtripTime
}

// GetRoundtripTimeOk returns a tuple with the RoundtripTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodEndPointDataV2) GetRoundtripTimeOk() (*int64, bool) {
	if o == nil || o.RoundtripTime == nil {
		return nil, false
	}
	return o.RoundtripTime, true
}

// HasRoundtripTime returns a boolean if a field has been set.
func (o *PodEndPointDataV2) HasRoundtripTime() bool {
	if o != nil && o.RoundtripTime != nil {
		return true
	}

	return false
}

// SetRoundtripTime gets a reference to the given int64 and assigns it to the RoundtripTime field.
func (o *PodEndPointDataV2) SetRoundtripTime(v int64) {
	o.RoundtripTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PodEndPointDataV2) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodEndPointDataV2) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PodEndPointDataV2) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *PodEndPointDataV2) SetStatus(v string) {
	o.Status = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *PodEndPointDataV2) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodEndPointDataV2) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *PodEndPointDataV2) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *PodEndPointDataV2) SetUrl(v string) {
	o.Url = &v
}

func (o PodEndPointDataV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTimestamp != nil {
		toSerialize["last_updated_timestamp"] = o.LastUpdatedTimestamp
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.RoundtripTime != nil {
		toSerialize["roundtrip_time"] = o.RoundtripTime
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Url != nil {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullablePodEndPointDataV2 struct {
	value *PodEndPointDataV2
	isSet bool
}

func (v NullablePodEndPointDataV2) Get() *PodEndPointDataV2 {
	return v.value
}

func (v *NullablePodEndPointDataV2) Set(val *PodEndPointDataV2) {
	v.value = val
	v.isSet = true
}

func (v NullablePodEndPointDataV2) IsSet() bool {
	return v.isSet
}

func (v *NullablePodEndPointDataV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodEndPointDataV2(val *PodEndPointDataV2) *NullablePodEndPointDataV2 {
	return &NullablePodEndPointDataV2{value: val, isSet: true}
}

func (v NullablePodEndPointDataV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodEndPointDataV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
