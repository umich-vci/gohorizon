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

// ImageManagementTagInfo Information related to image management tag.
type ImageManagementTagInfo struct {
	// Additional details about image management tag.
	AdditionalDetails *map[string]string `json:"additional_details,omitempty"`
	// Image management tag description.
	Description *string `json:"description,omitempty"`
	// Unique ID representing image management tag.
	Id *string `json:"id,omitempty"`
	// Image management stream ID to which this tag belongs.
	ImStreamId *string `json:"im_stream_id,omitempty"`
	// Image management version ID to which this tag belongs.
	ImVersionId *string `json:"im_version_id,omitempty"`
	// Image management tag name.
	Name *string `json:"name,omitempty"`
}

// NewImageManagementTagInfo instantiates a new ImageManagementTagInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageManagementTagInfo() *ImageManagementTagInfo {
	this := ImageManagementTagInfo{}
	return &this
}

// NewImageManagementTagInfoWithDefaults instantiates a new ImageManagementTagInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageManagementTagInfoWithDefaults() *ImageManagementTagInfo {
	this := ImageManagementTagInfo{}
	return &this
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *ImageManagementTagInfo) GetAdditionalDetails() map[string]string {
	if o == nil || o.AdditionalDetails == nil {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementTagInfo) GetAdditionalDetailsOk() (*map[string]string, bool) {
	if o == nil || o.AdditionalDetails == nil {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *ImageManagementTagInfo) HasAdditionalDetails() bool {
	if o != nil && o.AdditionalDetails != nil {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given map[string]string and assigns it to the AdditionalDetails field.
func (o *ImageManagementTagInfo) SetAdditionalDetails(v map[string]string) {
	o.AdditionalDetails = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ImageManagementTagInfo) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementTagInfo) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ImageManagementTagInfo) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ImageManagementTagInfo) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ImageManagementTagInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementTagInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ImageManagementTagInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ImageManagementTagInfo) SetId(v string) {
	o.Id = &v
}

// GetImStreamId returns the ImStreamId field value if set, zero value otherwise.
func (o *ImageManagementTagInfo) GetImStreamId() string {
	if o == nil || o.ImStreamId == nil {
		var ret string
		return ret
	}
	return *o.ImStreamId
}

// GetImStreamIdOk returns a tuple with the ImStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementTagInfo) GetImStreamIdOk() (*string, bool) {
	if o == nil || o.ImStreamId == nil {
		return nil, false
	}
	return o.ImStreamId, true
}

// HasImStreamId returns a boolean if a field has been set.
func (o *ImageManagementTagInfo) HasImStreamId() bool {
	if o != nil && o.ImStreamId != nil {
		return true
	}

	return false
}

// SetImStreamId gets a reference to the given string and assigns it to the ImStreamId field.
func (o *ImageManagementTagInfo) SetImStreamId(v string) {
	o.ImStreamId = &v
}

// GetImVersionId returns the ImVersionId field value if set, zero value otherwise.
func (o *ImageManagementTagInfo) GetImVersionId() string {
	if o == nil || o.ImVersionId == nil {
		var ret string
		return ret
	}
	return *o.ImVersionId
}

// GetImVersionIdOk returns a tuple with the ImVersionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementTagInfo) GetImVersionIdOk() (*string, bool) {
	if o == nil || o.ImVersionId == nil {
		return nil, false
	}
	return o.ImVersionId, true
}

// HasImVersionId returns a boolean if a field has been set.
func (o *ImageManagementTagInfo) HasImVersionId() bool {
	if o != nil && o.ImVersionId != nil {
		return true
	}

	return false
}

// SetImVersionId gets a reference to the given string and assigns it to the ImVersionId field.
func (o *ImageManagementTagInfo) SetImVersionId(v string) {
	o.ImVersionId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ImageManagementTagInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementTagInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ImageManagementTagInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ImageManagementTagInfo) SetName(v string) {
	o.Name = &v
}

func (o ImageManagementTagInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalDetails != nil {
		toSerialize["additional_details"] = o.AdditionalDetails
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ImStreamId != nil {
		toSerialize["im_stream_id"] = o.ImStreamId
	}
	if o.ImVersionId != nil {
		toSerialize["im_version_id"] = o.ImVersionId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableImageManagementTagInfo struct {
	value *ImageManagementTagInfo
	isSet bool
}

func (v NullableImageManagementTagInfo) Get() *ImageManagementTagInfo {
	return v.value
}

func (v *NullableImageManagementTagInfo) Set(val *ImageManagementTagInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableImageManagementTagInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableImageManagementTagInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageManagementTagInfo(val *ImageManagementTagInfo) *NullableImageManagementTagInfo {
	return &NullableImageManagementTagInfo{value: val, isSet: true}
}

func (v NullableImageManagementTagInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageManagementTagInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
