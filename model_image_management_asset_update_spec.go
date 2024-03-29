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

// ImageManagementAssetUpdateSpec Information related to image management asset.
type ImageManagementAssetUpdateSpec struct {
	// Additional details about image management asset.
	AdditionalDetails *map[string]string `json:"additional_details,omitempty"`
	// Image management asset clone type. * FULL_CLONE: Image management asset to be used in full clone automated desktop pool. * INSTANT_CLONE: Image management asset to be used in instant clone desktop pool/farm.
	CloneType string `json:"clone_type"`
	// Image management asset image type. * RDSH_APPS: Image management asset to be used for farm creation which is be used in application. * RDSH_DESKTOP: Image management asset is for farm creation to be created. * VDI_DESKTOP: Image management asset is available for desktops/farms to be created.
	ImageType string `json:"image_type"`
	// Image management asset status. * AVAILABLE: Image management asset is available for desktop pools/farms to be created. * DEPLOYING_VM: Image management asset is deploying VM on the virtual center. * DEPLOYMENT_DONE: Image management asset VM deployed on the virtual center. * DELETED: Image management asset has been deleted. * DISABLED: Image management asset has been disabled and no further pool/farm operation can be done using the same. * FAILED: Image management asset creation has failed. * REPLICATING: Copying the specialized images across all virtual centers. * RETRY_PENDING: When image management asset creation has failed, retry action is pending for asset to be created. * SPECIALIZING_VM: Image management asset is being published and specialized internally like installing agents etc.
	Status string `json:"status"`
}

// NewImageManagementAssetUpdateSpec instantiates a new ImageManagementAssetUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageManagementAssetUpdateSpec(cloneType string, imageType string, status string) *ImageManagementAssetUpdateSpec {
	this := ImageManagementAssetUpdateSpec{}
	this.CloneType = cloneType
	this.ImageType = imageType
	this.Status = status
	return &this
}

// NewImageManagementAssetUpdateSpecWithDefaults instantiates a new ImageManagementAssetUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageManagementAssetUpdateSpecWithDefaults() *ImageManagementAssetUpdateSpec {
	this := ImageManagementAssetUpdateSpec{}
	return &this
}

// GetAdditionalDetails returns the AdditionalDetails field value if set, zero value otherwise.
func (o *ImageManagementAssetUpdateSpec) GetAdditionalDetails() map[string]string {
	if o == nil || o.AdditionalDetails == nil {
		var ret map[string]string
		return ret
	}
	return *o.AdditionalDetails
}

// GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageManagementAssetUpdateSpec) GetAdditionalDetailsOk() (*map[string]string, bool) {
	if o == nil || o.AdditionalDetails == nil {
		return nil, false
	}
	return o.AdditionalDetails, true
}

// HasAdditionalDetails returns a boolean if a field has been set.
func (o *ImageManagementAssetUpdateSpec) HasAdditionalDetails() bool {
	if o != nil && o.AdditionalDetails != nil {
		return true
	}

	return false
}

// SetAdditionalDetails gets a reference to the given map[string]string and assigns it to the AdditionalDetails field.
func (o *ImageManagementAssetUpdateSpec) SetAdditionalDetails(v map[string]string) {
	o.AdditionalDetails = &v
}

// GetCloneType returns the CloneType field value
func (o *ImageManagementAssetUpdateSpec) GetCloneType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CloneType
}

// GetCloneTypeOk returns a tuple with the CloneType field value
// and a boolean to check if the value has been set.
func (o *ImageManagementAssetUpdateSpec) GetCloneTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloneType, true
}

// SetCloneType sets field value
func (o *ImageManagementAssetUpdateSpec) SetCloneType(v string) {
	o.CloneType = v
}

// GetImageType returns the ImageType field value
func (o *ImageManagementAssetUpdateSpec) GetImageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImageType
}

// GetImageTypeOk returns a tuple with the ImageType field value
// and a boolean to check if the value has been set.
func (o *ImageManagementAssetUpdateSpec) GetImageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImageType, true
}

// SetImageType sets field value
func (o *ImageManagementAssetUpdateSpec) SetImageType(v string) {
	o.ImageType = v
}

// GetStatus returns the Status field value
func (o *ImageManagementAssetUpdateSpec) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ImageManagementAssetUpdateSpec) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ImageManagementAssetUpdateSpec) SetStatus(v string) {
	o.Status = v
}

func (o ImageManagementAssetUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdditionalDetails != nil {
		toSerialize["additional_details"] = o.AdditionalDetails
	}
	if true {
		toSerialize["clone_type"] = o.CloneType
	}
	if true {
		toSerialize["image_type"] = o.ImageType
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableImageManagementAssetUpdateSpec struct {
	value *ImageManagementAssetUpdateSpec
	isSet bool
}

func (v NullableImageManagementAssetUpdateSpec) Get() *ImageManagementAssetUpdateSpec {
	return v.value
}

func (v *NullableImageManagementAssetUpdateSpec) Set(val *ImageManagementAssetUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableImageManagementAssetUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableImageManagementAssetUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageManagementAssetUpdateSpec(val *ImageManagementAssetUpdateSpec) *NullableImageManagementAssetUpdateSpec {
	return &NullableImageManagementAssetUpdateSpec{value: val, isSet: true}
}

func (v NullableImageManagementAssetUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageManagementAssetUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
