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

// PodAssignmentInfo Information related to pod assignments in a pod federation.<br>Only one of global Desktop Entitlement ID and global Application Entitlement ID will be set.<br>Supported Filters : 'And', 'Or' and 'Equals'.<br>See the field description to know the filter types it supports.
type PodAssignmentInfo struct {
	// ID of the Global Application Entitlement associated with this pod assignment.<br>Supported Filters: 'Equals'.
	GlobalApplicationEntitlementId *string `json:"global_application_entitlement_id,omitempty"`
	// ID of the Global Desktop Entitlement associated with this pod assignment.<br>Supported Filters: 'Equals'.
	GlobalDesktopEntitlementId *string `json:"global_desktop_entitlement_id,omitempty"`
	// Unique ID representing this pod assignment.<br>Supported Filters: 'Equals'.
	Id *string `json:"id,omitempty"`
	// ID representing the pod associated with this pod assignment.<br>Supported Filters: 'Equals'.
	PodId *string `json:"pod_id,omitempty"`
	// SID of the user associated with this pod assignment.<br>Supported Filters: 'Equals'.
	UserId *string `json:"user_id,omitempty"`
}

// NewPodAssignmentInfo instantiates a new PodAssignmentInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodAssignmentInfo() *PodAssignmentInfo {
	this := PodAssignmentInfo{}
	return &this
}

// NewPodAssignmentInfoWithDefaults instantiates a new PodAssignmentInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodAssignmentInfoWithDefaults() *PodAssignmentInfo {
	this := PodAssignmentInfo{}
	return &this
}

// GetGlobalApplicationEntitlementId returns the GlobalApplicationEntitlementId field value if set, zero value otherwise.
func (o *PodAssignmentInfo) GetGlobalApplicationEntitlementId() string {
	if o == nil || o.GlobalApplicationEntitlementId == nil {
		var ret string
		return ret
	}
	return *o.GlobalApplicationEntitlementId
}

// GetGlobalApplicationEntitlementIdOk returns a tuple with the GlobalApplicationEntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAssignmentInfo) GetGlobalApplicationEntitlementIdOk() (*string, bool) {
	if o == nil || o.GlobalApplicationEntitlementId == nil {
		return nil, false
	}
	return o.GlobalApplicationEntitlementId, true
}

// HasGlobalApplicationEntitlementId returns a boolean if a field has been set.
func (o *PodAssignmentInfo) HasGlobalApplicationEntitlementId() bool {
	if o != nil && o.GlobalApplicationEntitlementId != nil {
		return true
	}

	return false
}

// SetGlobalApplicationEntitlementId gets a reference to the given string and assigns it to the GlobalApplicationEntitlementId field.
func (o *PodAssignmentInfo) SetGlobalApplicationEntitlementId(v string) {
	o.GlobalApplicationEntitlementId = &v
}

// GetGlobalDesktopEntitlementId returns the GlobalDesktopEntitlementId field value if set, zero value otherwise.
func (o *PodAssignmentInfo) GetGlobalDesktopEntitlementId() string {
	if o == nil || o.GlobalDesktopEntitlementId == nil {
		var ret string
		return ret
	}
	return *o.GlobalDesktopEntitlementId
}

// GetGlobalDesktopEntitlementIdOk returns a tuple with the GlobalDesktopEntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAssignmentInfo) GetGlobalDesktopEntitlementIdOk() (*string, bool) {
	if o == nil || o.GlobalDesktopEntitlementId == nil {
		return nil, false
	}
	return o.GlobalDesktopEntitlementId, true
}

// HasGlobalDesktopEntitlementId returns a boolean if a field has been set.
func (o *PodAssignmentInfo) HasGlobalDesktopEntitlementId() bool {
	if o != nil && o.GlobalDesktopEntitlementId != nil {
		return true
	}

	return false
}

// SetGlobalDesktopEntitlementId gets a reference to the given string and assigns it to the GlobalDesktopEntitlementId field.
func (o *PodAssignmentInfo) SetGlobalDesktopEntitlementId(v string) {
	o.GlobalDesktopEntitlementId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PodAssignmentInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAssignmentInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PodAssignmentInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PodAssignmentInfo) SetId(v string) {
	o.Id = &v
}

// GetPodId returns the PodId field value if set, zero value otherwise.
func (o *PodAssignmentInfo) GetPodId() string {
	if o == nil || o.PodId == nil {
		var ret string
		return ret
	}
	return *o.PodId
}

// GetPodIdOk returns a tuple with the PodId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAssignmentInfo) GetPodIdOk() (*string, bool) {
	if o == nil || o.PodId == nil {
		return nil, false
	}
	return o.PodId, true
}

// HasPodId returns a boolean if a field has been set.
func (o *PodAssignmentInfo) HasPodId() bool {
	if o != nil && o.PodId != nil {
		return true
	}

	return false
}

// SetPodId gets a reference to the given string and assigns it to the PodId field.
func (o *PodAssignmentInfo) SetPodId(v string) {
	o.PodId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *PodAssignmentInfo) GetUserId() string {
	if o == nil || o.UserId == nil {
		var ret string
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodAssignmentInfo) GetUserIdOk() (*string, bool) {
	if o == nil || o.UserId == nil {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *PodAssignmentInfo) HasUserId() bool {
	if o != nil && o.UserId != nil {
		return true
	}

	return false
}

// SetUserId gets a reference to the given string and assigns it to the UserId field.
func (o *PodAssignmentInfo) SetUserId(v string) {
	o.UserId = &v
}

func (o PodAssignmentInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.GlobalApplicationEntitlementId != nil {
		toSerialize["global_application_entitlement_id"] = o.GlobalApplicationEntitlementId
	}
	if o.GlobalDesktopEntitlementId != nil {
		toSerialize["global_desktop_entitlement_id"] = o.GlobalDesktopEntitlementId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.PodId != nil {
		toSerialize["pod_id"] = o.PodId
	}
	if o.UserId != nil {
		toSerialize["user_id"] = o.UserId
	}
	return json.Marshal(toSerialize)
}

type NullablePodAssignmentInfo struct {
	value *PodAssignmentInfo
	isSet bool
}

func (v NullablePodAssignmentInfo) Get() *PodAssignmentInfo {
	return v.value
}

func (v *NullablePodAssignmentInfo) Set(val *PodAssignmentInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePodAssignmentInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePodAssignmentInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodAssignmentInfo(val *PodAssignmentInfo) *NullablePodAssignmentInfo {
	return &NullablePodAssignmentInfo{value: val, isSet: true}
}

func (v NullablePodAssignmentInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodAssignmentInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}