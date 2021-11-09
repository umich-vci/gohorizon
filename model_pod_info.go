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

// PodInfo Information related to pods in a pod federation.
type PodInfo struct {
	// List of Global Application Entitlements with mappings to application pools within this pod.
	ActiveGlobalApplicationEntitlements *[]string `json:"active_global_application_entitlements,omitempty"`
	// List of Global Desktop Entitlements with mappings to desktop pools within this pod.
	ActiveGlobalDesktopEntitlements *[]string `json:"active_global_desktop_entitlements,omitempty"`
	// Indicates whether this pod is managed from cloud.
	CloudManaged *bool `json:"cloud_managed,omitempty"`
	// Description of this pod.
	Description *string `json:"description,omitempty"`
	// List of endpoints within this pod.
	Endpoints *[]string `json:"endpoints,omitempty"`
	// Unique ID representing this pod.
	Id *string `json:"id,omitempty"`
	// Indicates whether this is the local pod the request was made from.<br>Only one pod in the pod federation will return true.
	LocalPod *bool `json:"local_pod,omitempty"`
	// Name of this pod.
	Name *string `json:"name,omitempty"`
	// ID of the site this pod belongs to.
	SiteId *string `json:"site_id,omitempty"`
}

// NewPodInfo instantiates a new PodInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPodInfo() *PodInfo {
	this := PodInfo{}
	return &this
}

// NewPodInfoWithDefaults instantiates a new PodInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodInfoWithDefaults() *PodInfo {
	this := PodInfo{}
	return &this
}

// GetActiveGlobalApplicationEntitlements returns the ActiveGlobalApplicationEntitlements field value if set, zero value otherwise.
func (o *PodInfo) GetActiveGlobalApplicationEntitlements() []string {
	if o == nil || o.ActiveGlobalApplicationEntitlements == nil {
		var ret []string
		return ret
	}
	return *o.ActiveGlobalApplicationEntitlements
}

// GetActiveGlobalApplicationEntitlementsOk returns a tuple with the ActiveGlobalApplicationEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetActiveGlobalApplicationEntitlementsOk() (*[]string, bool) {
	if o == nil || o.ActiveGlobalApplicationEntitlements == nil {
		return nil, false
	}
	return o.ActiveGlobalApplicationEntitlements, true
}

// HasActiveGlobalApplicationEntitlements returns a boolean if a field has been set.
func (o *PodInfo) HasActiveGlobalApplicationEntitlements() bool {
	if o != nil && o.ActiveGlobalApplicationEntitlements != nil {
		return true
	}

	return false
}

// SetActiveGlobalApplicationEntitlements gets a reference to the given []string and assigns it to the ActiveGlobalApplicationEntitlements field.
func (o *PodInfo) SetActiveGlobalApplicationEntitlements(v []string) {
	o.ActiveGlobalApplicationEntitlements = &v
}

// GetActiveGlobalDesktopEntitlements returns the ActiveGlobalDesktopEntitlements field value if set, zero value otherwise.
func (o *PodInfo) GetActiveGlobalDesktopEntitlements() []string {
	if o == nil || o.ActiveGlobalDesktopEntitlements == nil {
		var ret []string
		return ret
	}
	return *o.ActiveGlobalDesktopEntitlements
}

// GetActiveGlobalDesktopEntitlementsOk returns a tuple with the ActiveGlobalDesktopEntitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetActiveGlobalDesktopEntitlementsOk() (*[]string, bool) {
	if o == nil || o.ActiveGlobalDesktopEntitlements == nil {
		return nil, false
	}
	return o.ActiveGlobalDesktopEntitlements, true
}

// HasActiveGlobalDesktopEntitlements returns a boolean if a field has been set.
func (o *PodInfo) HasActiveGlobalDesktopEntitlements() bool {
	if o != nil && o.ActiveGlobalDesktopEntitlements != nil {
		return true
	}

	return false
}

// SetActiveGlobalDesktopEntitlements gets a reference to the given []string and assigns it to the ActiveGlobalDesktopEntitlements field.
func (o *PodInfo) SetActiveGlobalDesktopEntitlements(v []string) {
	o.ActiveGlobalDesktopEntitlements = &v
}

// GetCloudManaged returns the CloudManaged field value if set, zero value otherwise.
func (o *PodInfo) GetCloudManaged() bool {
	if o == nil || o.CloudManaged == nil {
		var ret bool
		return ret
	}
	return *o.CloudManaged
}

// GetCloudManagedOk returns a tuple with the CloudManaged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetCloudManagedOk() (*bool, bool) {
	if o == nil || o.CloudManaged == nil {
		return nil, false
	}
	return o.CloudManaged, true
}

// HasCloudManaged returns a boolean if a field has been set.
func (o *PodInfo) HasCloudManaged() bool {
	if o != nil && o.CloudManaged != nil {
		return true
	}

	return false
}

// SetCloudManaged gets a reference to the given bool and assigns it to the CloudManaged field.
func (o *PodInfo) SetCloudManaged(v bool) {
	o.CloudManaged = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PodInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PodInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PodInfo) SetDescription(v string) {
	o.Description = &v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *PodInfo) GetEndpoints() []string {
	if o == nil || o.Endpoints == nil {
		var ret []string
		return ret
	}
	return *o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetEndpointsOk() (*[]string, bool) {
	if o == nil || o.Endpoints == nil {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *PodInfo) HasEndpoints() bool {
	if o != nil && o.Endpoints != nil {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []string and assigns it to the Endpoints field.
func (o *PodInfo) SetEndpoints(v []string) {
	o.Endpoints = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PodInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PodInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PodInfo) SetId(v string) {
	o.Id = &v
}

// GetLocalPod returns the LocalPod field value if set, zero value otherwise.
func (o *PodInfo) GetLocalPod() bool {
	if o == nil || o.LocalPod == nil {
		var ret bool
		return ret
	}
	return *o.LocalPod
}

// GetLocalPodOk returns a tuple with the LocalPod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetLocalPodOk() (*bool, bool) {
	if o == nil || o.LocalPod == nil {
		return nil, false
	}
	return o.LocalPod, true
}

// HasLocalPod returns a boolean if a field has been set.
func (o *PodInfo) HasLocalPod() bool {
	if o != nil && o.LocalPod != nil {
		return true
	}

	return false
}

// SetLocalPod gets a reference to the given bool and assigns it to the LocalPod field.
func (o *PodInfo) SetLocalPod(v bool) {
	o.LocalPod = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PodInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PodInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PodInfo) SetName(v string) {
	o.Name = &v
}

// GetSiteId returns the SiteId field value if set, zero value otherwise.
func (o *PodInfo) GetSiteId() string {
	if o == nil || o.SiteId == nil {
		var ret string
		return ret
	}
	return *o.SiteId
}

// GetSiteIdOk returns a tuple with the SiteId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PodInfo) GetSiteIdOk() (*string, bool) {
	if o == nil || o.SiteId == nil {
		return nil, false
	}
	return o.SiteId, true
}

// HasSiteId returns a boolean if a field has been set.
func (o *PodInfo) HasSiteId() bool {
	if o != nil && o.SiteId != nil {
		return true
	}

	return false
}

// SetSiteId gets a reference to the given string and assigns it to the SiteId field.
func (o *PodInfo) SetSiteId(v string) {
	o.SiteId = &v
}

func (o PodInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveGlobalApplicationEntitlements != nil {
		toSerialize["active_global_application_entitlements"] = o.ActiveGlobalApplicationEntitlements
	}
	if o.ActiveGlobalDesktopEntitlements != nil {
		toSerialize["active_global_desktop_entitlements"] = o.ActiveGlobalDesktopEntitlements
	}
	if o.CloudManaged != nil {
		toSerialize["cloud_managed"] = o.CloudManaged
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Endpoints != nil {
		toSerialize["endpoints"] = o.Endpoints
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LocalPod != nil {
		toSerialize["local_pod"] = o.LocalPod
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SiteId != nil {
		toSerialize["site_id"] = o.SiteId
	}
	return json.Marshal(toSerialize)
}

type NullablePodInfo struct {
	value *PodInfo
	isSet bool
}

func (v NullablePodInfo) Get() *PodInfo {
	return v.value
}

func (v *NullablePodInfo) Set(val *PodInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePodInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePodInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodInfo(val *PodInfo) *NullablePodInfo {
	return &NullablePodInfo{value: val, isSet: true}
}

func (v NullablePodInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}