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

// ImageManagementVersionCreateSpec Information related to image management version.
type ImageManagementVersionCreateSpec struct {
	// Additional details about image management version.
	AdditionalDetails *map[string]string `json:"additional_details,omitempty"`
	// Image management version description.
	Description *string `json:"description,omitempty"`
	// Image management stream ID
	ImStreamId string `json:"im_stream_id"`
	// Image management version name.
	Name string `json:"name"`
	// Image management version status. * AVAILABLE: Image management version is available for desktop pools/farms to be created. * DEPLOYING_VM: Image management version is deploying VM on the selected pod. * DEPLOYMENT_DONE: Image management version status when VM deployment is done for the selected pod. * DELETED: Image management version has been deleted. * DISABLED: Image management version has been disabled and no further pool operation can be done using the same. * FAILED: Image management version creation has failed. * PARTIALLY_AVAILABLE: Some of the image management asset creation in some of the virtual centers have failed. * PUBLISHING: Image management version is being published and specialized internally like installing agents etc. * REPLICATING: Copying the specialized images across all virtual centers.
	Status string `json:"status"`
}

// NewImageManagementVersionCreateSpec instantiates a new ImageManagementVersionCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageManagementVersionCreateSpec(imStreamId string, name string, status string) *ImageManagementVersionCreateSpec {
	this := ImageManagementVersionCreateSpec{}
	this.ImStreamId = imStreamId
	this.Name = name
	this.Status = status
	return &this
}

// NewImageManagementVersionCreateSpecWithDefaults instantiates a new ImageManagementVersionCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageManagementVersionCreateSpecWithDefaults() *ImageManagementVersionCreateSpec {
	this := ImageManagementVersionCreateSpec{}
	return &this
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *ImageManagementVersionCreateSpec) GetAdditionalDetails() map[string]string {
	if o == nil || o.AdditionalDetails == nil {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementVersionCreateSpec) GetAdditionalDetailsOk() (*map[string]string, bool) {
	if o == nil || o.AdditionalDetails == nil {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *ImageManagementVersionCreateSpec) HasAdditionalDetails() bool {
	if o != nil && o.AdditionalDetails != nil {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given map[string]string and assigns it to the AdditionalDetails field.
func (o *ImageManagementVersionCreateSpec) SetAdditionalDetails(v map[string]string) {
	o.AdditionalDetails = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImageManagementVersionCreateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementVersionCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImageManagementVersionCreateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImageManagementVersionCreateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetImStreamId returns the ImStreamId field value
func (o *ImageManagementVersionCreateSpec) GetImStreamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImStreamId
}

// GetImStreamIdOk returns a tuple with the ImStreamId field value
// and a boolean to check if the value has been set.
func (o *ImageManagementVersionCreateSpec) GetImStreamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImStreamId, true
}

// SetImStreamId sets field value
func (o *ImageManagementVersionCreateSpec) SetImStreamId(v string) {
	o.ImStreamId = v
}

// GetName returns the Name field value
func (o *ImageManagementVersionCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImageManagementVersionCreateSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImageManagementVersionCreateSpec) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *ImageManagementVersionCreateSpec) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ImageManagementVersionCreateSpec) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ImageManagementVersionCreateSpec) SetStatus(v string) {
	o.Status = v
}

func (o ImageManagementVersionCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalDetails != nil {
		toSerialize["additional_details"] = o.AdditionalDetails
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["im_stream_id"] = o.ImStreamId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableImageManagementVersionCreateSpec struct {
	value *ImageManagementVersionCreateSpec
	isSet bool
}

func (v NullableImageManagementVersionCreateSpec) Get() *ImageManagementVersionCreateSpec {
	return v.value
}

func (v *NullableImageManagementVersionCreateSpec) Set(val *ImageManagementVersionCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableImageManagementVersionCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableImageManagementVersionCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageManagementVersionCreateSpec(val *ImageManagementVersionCreateSpec) *NullableImageManagementVersionCreateSpec {
	return &NullableImageManagementVersionCreateSpec{value: val, isSet: true}
}

func (v NullableImageManagementVersionCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageManagementVersionCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
