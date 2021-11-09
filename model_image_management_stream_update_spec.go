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

// ImageManagementStreamUpdateSpec Information related to image management stream.
type ImageManagementStreamUpdateSpec struct {
	// Additional details about image management stream.
	AdditionalDetails *map[string]string `json:"additional_details,omitempty"`
	// Image management stream description.
	Description *string `json:"description,omitempty"`
	// Image management stream name.
	Name string `json:"name"`
	// Operating system. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS)
	OperatingSystem string `json:"operating_system"`
	// Image management stream publisher
	Publisher *string `json:"publisher,omitempty"`
	// Image management stream source. * MARKET_PLACE: Image management stream is from market place. * UPLOADED: Image management stream is uploaded. * COPIED_FROM_STREAM: Image management stream is copied from another stream. * COPIED_FROM_VERSION: Image management stream is copied from a version.
	Source string `json:"source"`
	// Image management stream status. * AVAILABLE: Image management stream is available for desktop pools/farms to be created. * DELETED: Image management stream is deleted. * DISABLED: Image management stream is disabled and no further desktop pools/farms can be created using the same. * FAILED: Image management stream creation has failed. * IN_PROGRESS: Image management stream creation is in progress. * PARTIALLY_AVAILABLE: Image management version for this stream could not be created in one or more environments. * PENDING: Image management stream is in pending state.
	Status string `json:"status"`
}

// NewImageManagementStreamUpdateSpec instantiates a new ImageManagementStreamUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageManagementStreamUpdateSpec(name string, operatingSystem string, source string, status string) *ImageManagementStreamUpdateSpec {
	this := ImageManagementStreamUpdateSpec{}
	this.Name = name
	this.OperatingSystem = operatingSystem
	this.Source = source
	this.Status = status
	return &this
}

// NewImageManagementStreamUpdateSpecWithDefaults instantiates a new ImageManagementStreamUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageManagementStreamUpdateSpecWithDefaults() *ImageManagementStreamUpdateSpec {
	this := ImageManagementStreamUpdateSpec{}
	return &this
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *ImageManagementStreamUpdateSpec) GetAdditionalDetails() map[string]string {
	if o == nil || o.AdditionalDetails == nil {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementStreamUpdateSpec) GetAdditionalDetailsOk() (*map[string]string, bool) {
	if o == nil || o.AdditionalDetails == nil {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *ImageManagementStreamUpdateSpec) HasAdditionalDetails() bool {
	if o != nil && o.AdditionalDetails != nil {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given map[string]string and assigns it to the AdditionalDetails field.
func (o *ImageManagementStreamUpdateSpec) SetAdditionalDetails(v map[string]string) {
	o.AdditionalDetails = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImageManagementStreamUpdateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementStreamUpdateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImageManagementStreamUpdateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImageManagementStreamUpdateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetName returns the Name field value
func (o *ImageManagementStreamUpdateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImageManagementStreamUpdateSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImageManagementStreamUpdateSpec) SetName(v string) {
	o.Name = v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *ImageManagementStreamUpdateSpec) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *ImageManagementStreamUpdateSpec) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *ImageManagementStreamUpdateSpec) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *ImageManagementStreamUpdateSpec) GetPublisher() string {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementStreamUpdateSpec) GetPublisherOk() (*string, bool) {
	if o == nil || o.Publisher == nil {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *ImageManagementStreamUpdateSpec) HasPublisher() bool {
	if o != nil && o.Publisher != nil {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *ImageManagementStreamUpdateSpec) SetPublisher(v string) {
	o.Publisher = &v
}

// GetSource returns the Source field value
func (o *ImageManagementStreamUpdateSpec) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ImageManagementStreamUpdateSpec) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ImageManagementStreamUpdateSpec) SetSource(v string) {
	o.Source = v
}

// GetStatus returns the Status field value
func (o *ImageManagementStreamUpdateSpec) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ImageManagementStreamUpdateSpec) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ImageManagementStreamUpdateSpec) SetStatus(v string) {
	o.Status = v
}

func (o ImageManagementStreamUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalDetails != nil {
		toSerialize["additional_details"] = o.AdditionalDetails
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	if o.Publisher != nil {
		toSerialize["publisher"] = o.Publisher
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableImageManagementStreamUpdateSpec struct {
	value *ImageManagementStreamUpdateSpec
	isSet bool
}

func (v NullableImageManagementStreamUpdateSpec) Get() *ImageManagementStreamUpdateSpec {
	return v.value
}

func (v *NullableImageManagementStreamUpdateSpec) Set(val *ImageManagementStreamUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableImageManagementStreamUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableImageManagementStreamUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageManagementStreamUpdateSpec(val *ImageManagementStreamUpdateSpec) *NullableImageManagementStreamUpdateSpec {
	return &NullableImageManagementStreamUpdateSpec{value: val, isSet: true}
}

func (v NullableImageManagementStreamUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageManagementStreamUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
