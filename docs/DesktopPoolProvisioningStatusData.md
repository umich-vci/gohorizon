# DesktopPoolProvisioningStatusData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstantCloneCurrentImageState** | Pointer to **string** | Applicable To: instant clone automated desktop pools.&lt;br&gt;This represents the state of the current image of this instant clone desktop pool. * READY: This is the state of the current image after successful completion of creation operation. At this stage the current image is ready to be used to create the instant clones. Please note that this state is also reached from UNPUBLISHING state on successful completion of editing of cluster or editing of datastore(s) operations. * FAILED: This is the state of the current image if instant clone delete operation has failed or timed out. * PENDING_UNPUBLISH: This is the state of the current image before instant clone delete or cluster edit or datastore(s) edit operation(s) begins. * UNPUBLISHING: This is the transient state of the current image when instant clone delete or cluster edit or datastore(s) edit operation(s) is going on. | [optional] 
**InstantCloneOperation** | Pointer to **string** | Applicable To: instant clone automated desktop pools.&lt;br&gt;The operation that this instant clone desktop pool is undergoing. * NONE: There is no current operation on the desktop pool. * INITIAL_PUBLISH: The desktop pool has just been created and is undergoing initial publishing. * SCHEDULE_PUSH_IMAGE: The push operation is scheduled on the desktop pool. * CANCEL_SCHEDULED_PUSH_IMAGE: The scheduled push operation on the desktop pool is being cancelled. * INFRASTRUCTURE_CHANGE: A cluster or datastore change operation was requested for the desktop pool. * FINAL_UNPUBLISH: A desktop pool has been deleted and is undergoing final unpublishing. | [optional] 
**InstantClonePendingImStreamId** | Pointer to **string** | Applicable To: instant clone automated desktop pools.&lt;br&gt;Pending image management stream for instant clone desktop pools.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**InstantClonePendingImTagId** | Pointer to **string** | Applicable To: instant clone automated desktop pools.&lt;br&gt;Pending image management tag for instant clone desktop pools&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**InstantClonePendingImageError** | Pointer to **string** | Applicable To: instant clone automated desktop pools.&lt;br&gt;This represents the error message if publishing of push image operation is failed. | [optional] 
**InstantClonePendingImageParentVmId** | Pointer to **string** | Applicable To: instant clone automated desktop pools.&lt;br&gt;Pending base image VM for instant clone desktop pools. This is used to return the information about the parent VM of the pending Image. | [optional] 
**InstantClonePendingImageProgress** | Pointer to **int32** | Applicable To: instant clone automated desktop pools.&lt;br&gt;This represents the pending image publish progress in percentage for an instant clone desktop pool. | [optional] 
**InstantClonePendingImageSnapshotId** | Pointer to **string** | Applicable To: instant clone automated desktop pools.&lt;br&gt;Pending base image snapshot for instant clone desktop pools. This is used to return the information about the snapshot of the pending image. | [optional] 
**InstantClonePendingImageState** | Pointer to **string** | Applicable To: instant clone automated desktop pools.&lt;br&gt;This represents the state of the pending image of this instant clone desktop pool. This will be null when there is no pending image for the desktop pool. * PENDING_PUBLISH: This is the initial transient state of the pending image before instant clone creation operation has started. * PUBLISHING: This is the transient state of the pending image when creation of instant clone operation is going on. * UNPUBLISHING: This is the transient state of the pending image when instant clone delete or cluster edit or datastore(s) edit operation(s) is going on. * READY: This is the state of the pending image after successful publish of the pending image and before that image has been upgraded to the current image. This is normally seen after successful publish for a push image which has been scheduled to trigger at a later time. * FAILED: This is the state of the pending image if creation of instant clone operation has failed or timed out. | [optional] 
**InstantClonePushImageSettings** | Pointer to [**DesktopPoolInstantClonePushImageSettings**](DesktopPoolInstantClonePushImageSettings.md) |  | [optional] 
**LastProvisioningError** | Pointer to **string** | String message detailing the last provisioning error on this desktop pool while stop_provisioning_on_error is enabled. This will be cleared when enable_provisioning is updated to true. | [optional] 
**LastProvisioningErrorTime** | Pointer to **int64** | Time the last provisioning error occurred on this desktop while stop_provisioning_on_error is enabled. This will be cleared when enable_provisioning is updated to true. Measured as epoch time. | [optional] 

## Methods

### NewDesktopPoolProvisioningStatusData

`func NewDesktopPoolProvisioningStatusData() *DesktopPoolProvisioningStatusData`

NewDesktopPoolProvisioningStatusData instantiates a new DesktopPoolProvisioningStatusData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolProvisioningStatusDataWithDefaults

`func NewDesktopPoolProvisioningStatusDataWithDefaults() *DesktopPoolProvisioningStatusData`

NewDesktopPoolProvisioningStatusDataWithDefaults instantiates a new DesktopPoolProvisioningStatusData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstantCloneCurrentImageState

`func (o *DesktopPoolProvisioningStatusData) GetInstantCloneCurrentImageState() string`

GetInstantCloneCurrentImageState returns the InstantCloneCurrentImageState field if non-nil, zero value otherwise.

### GetInstantCloneCurrentImageStateOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantCloneCurrentImageStateOk() (*string, bool)`

GetInstantCloneCurrentImageStateOk returns a tuple with the InstantCloneCurrentImageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneCurrentImageState

`func (o *DesktopPoolProvisioningStatusData) SetInstantCloneCurrentImageState(v string)`

SetInstantCloneCurrentImageState sets InstantCloneCurrentImageState field to given value.

### HasInstantCloneCurrentImageState

`func (o *DesktopPoolProvisioningStatusData) HasInstantCloneCurrentImageState() bool`

HasInstantCloneCurrentImageState returns a boolean if a field has been set.

### GetInstantCloneOperation

`func (o *DesktopPoolProvisioningStatusData) GetInstantCloneOperation() string`

GetInstantCloneOperation returns the InstantCloneOperation field if non-nil, zero value otherwise.

### GetInstantCloneOperationOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantCloneOperationOk() (*string, bool)`

GetInstantCloneOperationOk returns a tuple with the InstantCloneOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneOperation

`func (o *DesktopPoolProvisioningStatusData) SetInstantCloneOperation(v string)`

SetInstantCloneOperation sets InstantCloneOperation field to given value.

### HasInstantCloneOperation

`func (o *DesktopPoolProvisioningStatusData) HasInstantCloneOperation() bool`

HasInstantCloneOperation returns a boolean if a field has been set.

### GetInstantClonePendingImStreamId

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImStreamId() string`

GetInstantClonePendingImStreamId returns the InstantClonePendingImStreamId field if non-nil, zero value otherwise.

### GetInstantClonePendingImStreamIdOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImStreamIdOk() (*string, bool)`

GetInstantClonePendingImStreamIdOk returns a tuple with the InstantClonePendingImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImStreamId

`func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImStreamId(v string)`

SetInstantClonePendingImStreamId sets InstantClonePendingImStreamId field to given value.

### HasInstantClonePendingImStreamId

`func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImStreamId() bool`

HasInstantClonePendingImStreamId returns a boolean if a field has been set.

### GetInstantClonePendingImTagId

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImTagId() string`

GetInstantClonePendingImTagId returns the InstantClonePendingImTagId field if non-nil, zero value otherwise.

### GetInstantClonePendingImTagIdOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImTagIdOk() (*string, bool)`

GetInstantClonePendingImTagIdOk returns a tuple with the InstantClonePendingImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImTagId

`func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImTagId(v string)`

SetInstantClonePendingImTagId sets InstantClonePendingImTagId field to given value.

### HasInstantClonePendingImTagId

`func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImTagId() bool`

HasInstantClonePendingImTagId returns a boolean if a field has been set.

### GetInstantClonePendingImageError

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageError() string`

GetInstantClonePendingImageError returns the InstantClonePendingImageError field if non-nil, zero value otherwise.

### GetInstantClonePendingImageErrorOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageErrorOk() (*string, bool)`

GetInstantClonePendingImageErrorOk returns a tuple with the InstantClonePendingImageError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageError

`func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageError(v string)`

SetInstantClonePendingImageError sets InstantClonePendingImageError field to given value.

### HasInstantClonePendingImageError

`func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageError() bool`

HasInstantClonePendingImageError returns a boolean if a field has been set.

### GetInstantClonePendingImageParentVmId

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageParentVmId() string`

GetInstantClonePendingImageParentVmId returns the InstantClonePendingImageParentVmId field if non-nil, zero value otherwise.

### GetInstantClonePendingImageParentVmIdOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageParentVmIdOk() (*string, bool)`

GetInstantClonePendingImageParentVmIdOk returns a tuple with the InstantClonePendingImageParentVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageParentVmId

`func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageParentVmId(v string)`

SetInstantClonePendingImageParentVmId sets InstantClonePendingImageParentVmId field to given value.

### HasInstantClonePendingImageParentVmId

`func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageParentVmId() bool`

HasInstantClonePendingImageParentVmId returns a boolean if a field has been set.

### GetInstantClonePendingImageProgress

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageProgress() int32`

GetInstantClonePendingImageProgress returns the InstantClonePendingImageProgress field if non-nil, zero value otherwise.

### GetInstantClonePendingImageProgressOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageProgressOk() (*int32, bool)`

GetInstantClonePendingImageProgressOk returns a tuple with the InstantClonePendingImageProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageProgress

`func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageProgress(v int32)`

SetInstantClonePendingImageProgress sets InstantClonePendingImageProgress field to given value.

### HasInstantClonePendingImageProgress

`func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageProgress() bool`

HasInstantClonePendingImageProgress returns a boolean if a field has been set.

### GetInstantClonePendingImageSnapshotId

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageSnapshotId() string`

GetInstantClonePendingImageSnapshotId returns the InstantClonePendingImageSnapshotId field if non-nil, zero value otherwise.

### GetInstantClonePendingImageSnapshotIdOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageSnapshotIdOk() (*string, bool)`

GetInstantClonePendingImageSnapshotIdOk returns a tuple with the InstantClonePendingImageSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageSnapshotId

`func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageSnapshotId(v string)`

SetInstantClonePendingImageSnapshotId sets InstantClonePendingImageSnapshotId field to given value.

### HasInstantClonePendingImageSnapshotId

`func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageSnapshotId() bool`

HasInstantClonePendingImageSnapshotId returns a boolean if a field has been set.

### GetInstantClonePendingImageState

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageState() string`

GetInstantClonePendingImageState returns the InstantClonePendingImageState field if non-nil, zero value otherwise.

### GetInstantClonePendingImageStateOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePendingImageStateOk() (*string, bool)`

GetInstantClonePendingImageStateOk returns a tuple with the InstantClonePendingImageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageState

`func (o *DesktopPoolProvisioningStatusData) SetInstantClonePendingImageState(v string)`

SetInstantClonePendingImageState sets InstantClonePendingImageState field to given value.

### HasInstantClonePendingImageState

`func (o *DesktopPoolProvisioningStatusData) HasInstantClonePendingImageState() bool`

HasInstantClonePendingImageState returns a boolean if a field has been set.

### GetInstantClonePushImageSettings

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePushImageSettings() DesktopPoolInstantClonePushImageSettings`

GetInstantClonePushImageSettings returns the InstantClonePushImageSettings field if non-nil, zero value otherwise.

### GetInstantClonePushImageSettingsOk

`func (o *DesktopPoolProvisioningStatusData) GetInstantClonePushImageSettingsOk() (*DesktopPoolInstantClonePushImageSettings, bool)`

GetInstantClonePushImageSettingsOk returns a tuple with the InstantClonePushImageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePushImageSettings

`func (o *DesktopPoolProvisioningStatusData) SetInstantClonePushImageSettings(v DesktopPoolInstantClonePushImageSettings)`

SetInstantClonePushImageSettings sets InstantClonePushImageSettings field to given value.

### HasInstantClonePushImageSettings

`func (o *DesktopPoolProvisioningStatusData) HasInstantClonePushImageSettings() bool`

HasInstantClonePushImageSettings returns a boolean if a field has been set.

### GetLastProvisioningError

`func (o *DesktopPoolProvisioningStatusData) GetLastProvisioningError() string`

GetLastProvisioningError returns the LastProvisioningError field if non-nil, zero value otherwise.

### GetLastProvisioningErrorOk

`func (o *DesktopPoolProvisioningStatusData) GetLastProvisioningErrorOk() (*string, bool)`

GetLastProvisioningErrorOk returns a tuple with the LastProvisioningError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProvisioningError

`func (o *DesktopPoolProvisioningStatusData) SetLastProvisioningError(v string)`

SetLastProvisioningError sets LastProvisioningError field to given value.

### HasLastProvisioningError

`func (o *DesktopPoolProvisioningStatusData) HasLastProvisioningError() bool`

HasLastProvisioningError returns a boolean if a field has been set.

### GetLastProvisioningErrorTime

`func (o *DesktopPoolProvisioningStatusData) GetLastProvisioningErrorTime() int64`

GetLastProvisioningErrorTime returns the LastProvisioningErrorTime field if non-nil, zero value otherwise.

### GetLastProvisioningErrorTimeOk

`func (o *DesktopPoolProvisioningStatusData) GetLastProvisioningErrorTimeOk() (*int64, bool)`

GetLastProvisioningErrorTimeOk returns a tuple with the LastProvisioningErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProvisioningErrorTime

`func (o *DesktopPoolProvisioningStatusData) SetLastProvisioningErrorTime(v int64)`

SetLastProvisioningErrorTime sets LastProvisioningErrorTime field to given value.

### HasLastProvisioningErrorTime

`func (o *DesktopPoolProvisioningStatusData) HasLastProvisioningErrorTime() bool`

HasLastProvisioningErrorTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


