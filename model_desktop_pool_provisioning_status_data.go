/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DesktopPoolProvisioningStatusData Provisioning status data about this automated desktop pool.
type DesktopPoolProvisioningStatusData struct {
	// Applicable To: instant clone automated desktop pools.<br>This represents the state of the current image of this instant clone desktop pool. * READY: This is the state of the current image after successful completion of pool creation operation. At this stage the current image is ready to be used to create the instant clones. Please note that this state is also reached from UNPUBLISHING state on successful completion of editing of cluster or editing of datastore(s) operations. * FAILED: This is the state of the current image if instant clone pool delete operation has failed or timed out. * PENDING_UNPUBLISH: This is the state of the current image before instant clone pool delete or cluster edit or datastore(s) edit operation(s) begins. * UNPUBLISHING: This is the transient state of the current image when instant clone pool delete or cluster edit or datastore(s) edit operation(s) is going on.
	InstantCloneCurrentImageState string `json:"instant_clone_current_image_state,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>The operation that this instant clone desktop pool is undergoing. * NONE: There is no current operation on the desktop pool. * INITIAL_PUBLISH: The desktop pool has just been created and is undergoing initial publishing. * SCHEDULE_PUSH_IMAGE: The push operation is scheduled on the desktop pool. * CANCEL_SCHEDULED_PUSH_IMAGE: The scheduled push operation on the desktop pool is being cancelled. * INFRASTRUCTURE_CHANGE: A cluster or datastore change operation was requested for the desktop pool. * FINAL_UNPUBLISH: A desktop pool has been deleted and is undergoing final unpublishing.
	InstantCloneOperation string `json:"instant_clone_operation,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>Pending image management stream for instant clone desktop pools.
	InstantClonePendingImStreamId string `json:"instant_clone_pending_im_stream_id,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>Pending image management tag for instant clone desktop pools
	InstantClonePendingImTagId string `json:"instant_clone_pending_im_tag_id,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>This represents the error message if publishing of push image operation is failed.
	InstantClonePendingImageError string `json:"instant_clone_pending_image_error,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>Pending base image VM for instant clone desktop pools. This is used to return the information about the parent VM of the pending Image.
	InstantClonePendingImageParentVmId string `json:"instant_clone_pending_image_parent_vm_id,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>This represents the pending image publish progress in percentage for an instant clone desktop pool.
	InstantClonePendingImageProgress int32 `json:"instant_clone_pending_image_progress,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>Pending base image snapshot for instant clone desktop pools. This is used to return the information about the snapshot of the pending image.
	InstantClonePendingImageSnapshotId string `json:"instant_clone_pending_image_snapshot_id,omitempty"`
	// Applicable To: instant clone automated desktop pools.<br>This represents the state of the pending image of this instant clone desktop pool. This will be null when there is no pending image for the desktop pool. * PENDING_PUBLISH: This is the initial transient state of the pending image before instant clone pool creation operation has started. * PUBLISHING: This is the transient state of the pending image when creation of instant clone pool operation is going on. * UNPUBLISHING: This is the transient state of the pending image when instant clone pool delete or cluster edit or datastore(s) edit operation(s) is going on. * READY: This is the state of the pending image after successful publish of the pending image and before that image has been upgraded to the current image. This is normally seen after successful publish for a push image which has been scheduled to trigger at a later time. * FAILED: This is the state of the pending image if creation of instant clone pool operation has failed or timed out.
	InstantClonePendingImageState string                                   `json:"instant_clone_pending_image_state,omitempty"`
	InstantClonePushImageSettings DesktopPoolInstantClonePushImageSettings `json:"instant_clone_push_image_settings,omitempty"`
	// String message detailing the last provisioning error on this desktop pool while stop_provisioning_on_error is enabled. This will be cleared when enable_provisioning is updated to true.
	LastProvisioningError string `json:"last_provisioning_error,omitempty"`
	// Time the last provisioning error occurred on this desktop while stop_provisioning_on_error is enabled. This will be cleared when enable_provisioning is updated to true. Measured as epoch time.
	LastProvisioningErrorTime int64 `json:"last_provisioning_error_time,omitempty"`
}
