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

// DesktopPoolProvisioningStatusData Provisioning status data about this automated desktop pool.
type DesktopPoolProvisioningStatusData struct {
	// Applicable To: instant clone automated desktop pools.<br>This represents the state of the current image of this instant clone desktop pool. * READY: This is the state of the current image after successful completion of creation operation. At this stage the current image is ready to be used to create the instant clones. Please note that this state is also reached from UNPUBLISHING state on successful completion of editing of cluster or editing of datastore(s) operations. * FAILED: This is the state of the current image if instant clone delete operation has failed or timed out. * PENDING_UNPUBLISH: This is the state of the current image before instant clone delete or cluster edit or datastore(s) edit operation(s) begins. * UNPUBLISHING: This is the transient state of the current image when instant clone delete or cluster edit or datastore(s) edit operation(s) is going on.
	InstantCloneCurrentImageState *string `json:"instant_clone_current_image_state,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>The operation that this instant clone desktop pool is undergoing. * NONE: There is no current operation on the desktop pool. * INITIAL_PUBLISH: The desktop pool has just been created and is undergoing initial publishing. * SCHEDULE_PUSH_IMAGE: The push operation is scheduled on the desktop pool. * CANCEL_SCHEDULED_PUSH_IMAGE: The scheduled push operation on the desktop pool is being cancelled. * INFRASTRUCTURE_CHANGE: A cluster or datastore change operation was requested for the desktop pool. * FINAL_UNPUBLISH: A desktop pool has been deleted and is undergoing final unpublishing.
	InstantCloneOperation *string `json:"instant_clone_operation,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>Pending image management stream for instant clone desktop pools.<br>Supported Filters: 'Equals'.
	InstantClonePendingImStreamId *string `json:"instant_clone_pending_im_stream_id,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>Pending image management tag for instant clone desktop pools<br>Supported Filters: 'Equals'.
	InstantClonePendingImTagId *string `json:"instant_clone_pending_im_tag_id,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>This represents the error message if publishing of push image operation is failed.
	InstantClonePendingImageError *string `json:"instant_clone_pending_image_error,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>Pending base image VM for instant clone desktop pools. This is used to return the information about the parent VM of the pending Image.
	InstantClonePendingImageParentVmId *string `json:"instant_clone_pending_image_parent_vm_id,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>This represents the pending image publish progress in percentage for an instant clone desktop pool.
	InstantClonePendingImageProgress *int32 `json:"instant_clone_pending_image_progress,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>Pending base image snapshot for instant clone desktop pools. This is used to return the information about the snapshot of the pending image.
	InstantClonePendingImageSnapshotId *string `json:"instant_clone_pending_image_snapshot_id,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>This represents the state of the pending image of this instant clone desktop pool. This will be null when there is no pending image for the desktop pool. * PENDING_PUBLISH: This is the initial transient state of the pending image before instant clone creation operation has started. * PUBLISHING: This is the transient state of the pending image when creation of instant clone operation is going on. * UNPUBLISHING: This is the transient state of the pending image when instant clone delete or cluster edit or datastore(s) edit operation(s) is going on. * READY: This is the state of the pending image after successful publish of the pending image and before that image has been upgraded to the current image. This is normally seen after successful publish for a push image which has been scheduled to trigger at a later time. * FAILED: This is the state of the pending image if creation of instant clone operation has failed or timed out.
	InstantClonePendingImageState *string                                   `json:"instant_clone_pending_image_state,omitempty"`
	InstantClonePushImageSettings *DesktopPoolInstantClonePushImageSettings `json:"instant_clone_push_image_settings,omitempty"`
	// String message detailing the last provisioning error on this desktop pool while stop_provisioning_on_error is enabled. This will be cleared when enable_provisioning is updated to true.
	LastProvisioningError *string `json:"last_provisioning_error,omitempty"`
	// Time the last provisioning error occurred on this desktop while stop_provisioning_on_error is enabled. This will be cleared when enable_provisioning is updated to true. Measured as epoch time.
	LastProvisioningErrorTime *int64 `json:"last_provisioning_error_time,omitempty"`
}

// NewDesktopPoolProvisioningStatusData instantiates a new DesktopPoolProvisioningStatusData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolProvisioningStatusData() *DesktopPoolProvisioningStatusData {
	this := DesktopPoolProvisioningStatusData{}
	return &this
}

// NewDesktopPoolProvisioningStatusDataWithDefaults instantiates a new DesktopPoolProvisioningStatusData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolProvisioningStatusDataWithDefaults() *DesktopPoolProvisioningStatusData {
	this := DesktopPoolProvisioningStatusData{}
	return &this
}

// GetInstantCloneCurrentImageState returns the InstantCloneCurrentImageState field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantCloneCurrentImageState() string {
	if o == nil || o.InstantCloneCurrentImageState == nil {
		var ret string
		return ret
	}
	return *o.InstantCloneCurrentImageState
}

// GetInstantCloneCurrentImageStateOk returns a tuple with the InstantCloneCurrentImageState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantCloneCurrentImageStateOk() (*string, bool) {
	if o == nil || o.InstantCloneCurrentImageState == nil {
		return nil, false
	}
	return o.InstantCloneCurrentImageState, true
}

// HasInstantCloneCurrentImageState returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantCloneCurrentImageState() bool {
	if o != nil && o.InstantCloneCurrentImageState != nil {
		return true
	}

	return false
}

// SetInstantCloneCurrentImageState gets a reference to the given string and assigns it to the InstantCloneCurrentImageState field.
func (o *DesktopPoolProvisioningStatusData) SetInstantCloneCurrentImageState(v string) {
	o.InstantCloneCurrentImageState = &v
}

// GetInstantCloneOperation returns the InstantCloneOperation field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantCloneOperation() string {
	if o == nil || o.InstantCloneOperation == nil {
		var ret string
		return ret
	}
	return *o.InstantCloneOperation
}

// GetInstantCloneOperationOk returns a tuple with the InstantCloneOperation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantCloneOperationOk() (*string, bool) {
	if o == nil || o.InstantCloneOperation == nil {
		return nil, false
	}
	return o.InstantCloneOperation, true
}

// HasInstantCloneOperation returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantCloneOperation() bool {
	if o != nil && o.InstantCloneOperation != nil {
		return true
	}

	return false
}

// SetInstantCloneOperation gets a reference to the given string and assigns it to the InstantCloneOperation field.
func (o *DesktopPoolProvisioningStatusData) SetInstantCloneOperation(v string) {
	o.InstantCloneOperation = &v
}

// GetInstantClonePendingImStreamId returns the InstantClonePendingImStreamId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImStreamId() string {
	if o == nil || o.InstantClonePendingImStreamId == nil {
		var ret string
		return ret
	}
	return *o.InstantClonePendingImStreamId
}

// GetInstantClonePendingImStreamIdOk returns a tuple with the InstantClonePendingImStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImStreamIdOk() (*string, bool) {
	if o == nil || o.InstantClonePendingImStreamId == nil {
		return nil, false
	}
	return o.InstantClonePendingImStreamId, true
}

// HasInstantClonePendingImStreamId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImStreamId() bool {
	if o != nil && o.InstantClonePendingImStreamId != nil {
		return true
	}

	return false
}

// SetInstantClonePendingImStreamId gets a reference to the given string and assigns it to the InstantClonePendingImStreamId field.
func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImStreamId(v string) {
	o.InstantClonePendingImStreamId = &v
}

// GetInstantClonePendingImTagId returns the InstantClonePendingImTagId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImTagId() string {
	if o == nil || o.InstantClonePendingImTagId == nil {
		var ret string
		return ret
	}
	return *o.InstantClonePendingImTagId
}

// GetInstantClonePendingImTagIdOk returns a tuple with the InstantClonePendingImTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImTagIdOk() (*string, bool) {
	if o == nil || o.InstantClonePendingImTagId == nil {
		return nil, false
	}
	return o.InstantClonePendingImTagId, true
}

// HasInstantClonePendingImTagId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImTagId() bool {
	if o != nil && o.InstantClonePendingImTagId != nil {
		return true
	}

	return false
}

// SetInstantClonePendingImTagId gets a reference to the given string and assigns it to the InstantClonePendingImTagId field.
func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImTagId(v string) {
	o.InstantClonePendingImTagId = &v
}

// GetInstantClonePendingImageError returns the InstantClonePendingImageError field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageError() string {
	if o == nil || o.InstantClonePendingImageError == nil {
		var ret string
		return ret
	}
	return *o.InstantClonePendingImageError
}

// GetInstantClonePendingImageErrorOk returns a tuple with the InstantClonePendingImageError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageErrorOk() (*string, bool) {
	if o == nil || o.InstantClonePendingImageError == nil {
		return nil, false
	}
	return o.InstantClonePendingImageError, true
}

// HasInstantClonePendingImageError returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageError() bool {
	if o != nil && o.InstantClonePendingImageError != nil {
		return true
	}

	return false
}

// SetInstantClonePendingImageError gets a reference to the given string and assigns it to the InstantClonePendingImageError field.
func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageError(v string) {
	o.InstantClonePendingImageError = &v
}

// GetInstantClonePendingImageParentVmId returns the InstantClonePendingImageParentVmId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageParentVmId() string {
	if o == nil || o.InstantClonePendingImageParentVmId == nil {
		var ret string
		return ret
	}
	return *o.InstantClonePendingImageParentVmId
}

// GetInstantClonePendingImageParentVmIdOk returns a tuple with the InstantClonePendingImageParentVmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageParentVmIdOk() (*string, bool) {
	if o == nil || o.InstantClonePendingImageParentVmId == nil {
		return nil, false
	}
	return o.InstantClonePendingImageParentVmId, true
}

// HasInstantClonePendingImageParentVmId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageParentVmId() bool {
	if o != nil && o.InstantClonePendingImageParentVmId != nil {
		return true
	}

	return false
}

// SetInstantClonePendingImageParentVmId gets a reference to the given string and assigns it to the InstantClonePendingImageParentVmId field.
func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageParentVmId(v string) {
	o.InstantClonePendingImageParentVmId = &v
}

// GetInstantClonePendingImageProgress returns the InstantClonePendingImageProgress field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageProgress() int32 {
	if o == nil || o.InstantClonePendingImageProgress == nil {
		var ret int32
		return ret
	}
	return *o.InstantClonePendingImageProgress
}

// GetInstantClonePendingImageProgressOk returns a tuple with the InstantClonePendingImageProgress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageProgressOk() (*int32, bool) {
	if o == nil || o.InstantClonePendingImageProgress == nil {
		return nil, false
	}
	return o.InstantClonePendingImageProgress, true
}

// HasInstantClonePendingImageProgress returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageProgress() bool {
	if o != nil && o.InstantClonePendingImageProgress != nil {
		return true
	}

	return false
}

// SetInstantClonePendingImageProgress gets a reference to the given int32 and assigns it to the InstantClonePendingImageProgress field.
func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageProgress(v int32) {
	o.InstantClonePendingImageProgress = &v
}

// GetInstantClonePendingImageSnapshotId returns the InstantClonePendingImageSnapshotId field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageSnapshotId() string {
	if o == nil || o.InstantClonePendingImageSnapshotId == nil {
		var ret string
		return ret
	}
	return *o.InstantClonePendingImageSnapshotId
}

// GetInstantClonePendingImageSnapshotIdOk returns a tuple with the InstantClonePendingImageSnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageSnapshotIdOk() (*string, bool) {
	if o == nil || o.InstantClonePendingImageSnapshotId == nil {
		return nil, false
	}
	return o.InstantClonePendingImageSnapshotId, true
}

// HasInstantClonePendingImageSnapshotId returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageSnapshotId() bool {
	if o != nil && o.InstantClonePendingImageSnapshotId != nil {
		return true
	}

	return false
}

// SetInstantClonePendingImageSnapshotId gets a reference to the given string and assigns it to the InstantClonePendingImageSnapshotId field.
func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageSnapshotId(v string) {
	o.InstantClonePendingImageSnapshotId = &v
}

// GetInstantClonePendingImageState returns the InstantClonePendingImageState field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageState() string {
	if o == nil || o.InstantClonePendingImageState == nil {
		var ret string
		return ret
	}
	return *o.InstantClonePendingImageState
}

// GetInstantClonePendingImageStateOk returns a tuple with the InstantClonePendingImageState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageStateOk() (*string, bool) {
	if o == nil || o.InstantClonePendingImageState == nil {
		return nil, false
	}
	return o.InstantClonePendingImageState, true
}

// HasInstantClonePendingImageState returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageState() bool {
	if o != nil && o.InstantClonePendingImageState != nil {
		return true
	}

	return false
}

// SetInstantClonePendingImageState gets a reference to the given string and assigns it to the InstantClonePendingImageState field.
func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageState(v string) {
	o.InstantClonePendingImageState = &v
}

// GetInstantClonePushImageSettings returns the InstantClonePushImageSettings field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePushImageSettings() DesktopPoolInstantClonePushImageSettings {
	if o == nil || o.InstantClonePushImageSettings == nil {
		var ret DesktopPoolInstantClonePushImageSettings
		return ret
	}
	return *o.InstantClonePushImageSettings
}

// GetInstantClonePushImageSettingsOk returns a tuple with the InstantClonePushImageSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetInstantClonePushImageSettingsOk() (*DesktopPoolInstantClonePushImageSettings, bool) {
	if o == nil || o.InstantClonePushImageSettings == nil {
		return nil, false
	}
	return o.InstantClonePushImageSettings, true
}

// HasInstantClonePushImageSettings returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasInstantClonePushImageSettings() bool {
	if o != nil && o.InstantClonePushImageSettings != nil {
		return true
	}

	return false
}

// SetInstantClonePushImageSettings gets a reference to the given DesktopPoolInstantClonePushImageSettings and assigns it to the InstantClonePushImageSettings field.
func (o *DesktopPoolProvisioningStatusData) SetInstantClonePushImageSettings(v DesktopPoolInstantClonePushImageSettings) {
	o.InstantClonePushImageSettings = &v
}

// GetLastProvisioningError returns the LastProvisioningError field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetLastProvisioningError() string {
	if o == nil || o.LastProvisioningError == nil {
		var ret string
		return ret
	}
	return *o.LastProvisioningError
}

// GetLastProvisioningErrorOk returns a tuple with the LastProvisioningError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetLastProvisioningErrorOk() (*string, bool) {
	if o == nil || o.LastProvisioningError == nil {
		return nil, false
	}
	return o.LastProvisioningError, true
}

// HasLastProvisioningError returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasLastProvisioningError() bool {
	if o != nil && o.LastProvisioningError != nil {
		return true
	}

	return false
}

// SetLastProvisioningError gets a reference to the given string and assigns it to the LastProvisioningError field.
func (o *DesktopPoolProvisioningStatusData) SetLastProvisioningError(v string) {
	o.LastProvisioningError = &v
}

// GetLastProvisioningErrorTime returns the LastProvisioningErrorTime field value if set, zero value otherwise.
func (o *DesktopPoolProvisioningStatusData) GetLastProvisioningErrorTime() int64 {
	if o == nil || o.LastProvisioningErrorTime == nil {
		var ret int64
		return ret
	}
	return *o.LastProvisioningErrorTime
}

// GetLastProvisioningErrorTimeOk returns a tuple with the LastProvisioningErrorTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolProvisioningStatusData) GetLastProvisioningErrorTimeOk() (*int64, bool) {
	if o == nil || o.LastProvisioningErrorTime == nil {
		return nil, false
	}
	return o.LastProvisioningErrorTime, true
}

// HasLastProvisioningErrorTime returns a boolean if a field has been set.
func (o *DesktopPoolProvisioningStatusData) HasLastProvisioningErrorTime() bool {
	if o != nil && o.LastProvisioningErrorTime != nil {
		return true
	}

	return false
}

// SetLastProvisioningErrorTime gets a reference to the given int64 and assigns it to the LastProvisioningErrorTime field.
func (o *DesktopPoolProvisioningStatusData) SetLastProvisioningErrorTime(v int64) {
	o.LastProvisioningErrorTime = &v
}

func (o DesktopPoolProvisioningStatusData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.InstantCloneCurrentImageState != nil {
		toSerialize["instant_clone_current_image_state"] = o.InstantCloneCurrentImageState
	}
	if o.InstantCloneOperation != nil {
		toSerialize["instant_clone_operation"] = o.InstantCloneOperation
	}
	if o.InstantClonePendingImStreamId != nil {
		toSerialize["instant_clone_pending_im_stream_id"] = o.InstantClonePendingImStreamId
	}
	if o.InstantClonePendingImTagId != nil {
		toSerialize["instant_clone_pending_im_tag_id"] = o.InstantClonePendingImTagId
	}
	if o.InstantClonePendingImageError != nil {
		toSerialize["instant_clone_pending_image_error"] = o.InstantClonePendingImageError
	}
	if o.InstantClonePendingImageParentVmId != nil {
		toSerialize["instant_clone_pending_image_parent_vm_id"] = o.InstantClonePendingImageParentVmId
	}
	if o.InstantClonePendingImageProgress != nil {
		toSerialize["instant_clone_pending_image_progress"] = o.InstantClonePendingImageProgress
	}
	if o.InstantClonePendingImageSnapshotId != nil {
		toSerialize["instant_clone_pending_image_snapshot_id"] = o.InstantClonePendingImageSnapshotId
	}
	if o.InstantClonePendingImageState != nil {
		toSerialize["instant_clone_pending_image_state"] = o.InstantClonePendingImageState
	}
	if o.InstantClonePushImageSettings != nil {
		toSerialize["instant_clone_push_image_settings"] = o.InstantClonePushImageSettings
	}
	if o.LastProvisioningError != nil {
		toSerialize["last_provisioning_error"] = o.LastProvisioningError
	}
	if o.LastProvisioningErrorTime != nil {
		toSerialize["last_provisioning_error_time"] = o.LastProvisioningErrorTime
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolProvisioningStatusData struct {
	value *DesktopPoolProvisioningStatusData
	isSet bool
}

func (v NullableDesktopPoolProvisioningStatusData) Get() *DesktopPoolProvisioningStatusData {
	return v.value
}

func (v *NullableDesktopPoolProvisioningStatusData) Set(val *DesktopPoolProvisioningStatusData) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolProvisioningStatusData) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolProvisioningStatusData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolProvisioningStatusData(val *DesktopPoolProvisioningStatusData) *NullableDesktopPoolProvisioningStatusData {
	return &NullableDesktopPoolProvisioningStatusData{value: val, isSet: true}
}

func (v NullableDesktopPoolProvisioningStatusData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolProvisioningStatusData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
