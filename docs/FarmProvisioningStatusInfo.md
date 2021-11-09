# FarmProvisioningStatusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstantCloneCurrentImageState** | Pointer to **string** | This represents the state of the current image of this instant clone farms. * READY: This is the state of the current image after successful completion of creation operation. At this stage the current image is ready to be used to create the instant clones. Please note that this state is also reached from UNPUBLISHING state on successful completion of editing of cluster or editing of datastore(s) operations. * FAILED: This is the state of the current image if instant clone delete operation has failed or timed out. * PENDING_UNPUBLISH: This is the state of the current image before instant clone delete or cluster edit or datastore(s) edit operation(s) begins. * UNPUBLISHING: This is the transient state of the current image when instant clone delete or cluster edit or datastore(s) edit operation(s) is going on. | [optional] 
**InstantCloneOperation** | Pointer to **string** | The operation that this instant clone farm is undergoing. * NONE: There is no current operation on the farm. * INITIAL_PUBLISH: The farm has just been created and is undergoing initial publishing. * RECURRING_SCHEDULED_MAINTENANCE: Recurring maintenance operation is scheduled on the farm. * CANCEL_SCHEDULED_MAINTENANCE: The recurring maintenance operation on the farm is being cancelled. * INFRASTRUCTURE_CHANGE: A cluster or datastore change operation was requested for the farm. * FINAL_UNPUBLISH: A farm has been deleted and is undergoing final unpublishing. | [optional] 
**InstantCloneOperationTime** | Pointer to **int64** | Time of the operation that instant clone farm is undergoing. | [optional] 
**InstantClonePendingImStreamId** | Pointer to **string** | Pending image management stream for instant clone farms.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**InstantClonePendingImTagId** | Pointer to **string** | Pending image management tag for instant clone farms.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**InstantClonePendingImageError** | Pointer to **string** | This represents the error message if publishing of push image operation is failed. | [optional] 
**InstantClonePendingImageParentVmId** | Pointer to **string** | Pending base image VM for instant clone farms. This is used to return the information about the parent VM of the pending Image. | [optional] 
**InstantClonePendingImageProgress** | Pointer to **int32** | This represents the pending image publish progress in percentage for an instant clone farm. | [optional] 
**InstantClonePendingImageSnapshotId** | Pointer to **string** | Pending base image snapshot for instant clone farms. This is used to return the information about the snapshot of the pending image. | [optional] 
**InstantClonePendingImageState** | Pointer to **string** | This represents the state of the pending image of this instant clone farms. This will be null when there is no pending image for the farm. * PENDING_PUBLISH: This is the initial transient state of the pending image before instant clone creation operation has started. * PUBLISHING: This is the transient state of the pending image when creation of instant clone operation is going on. * UNPUBLISHING: This is the transient state of the pending image when instant clone delete or cluster edit or datastore(s) edit operation(s) is going on. * READY: This is the state of the pending image after successful publish of the pending image and before that image has been upgraded to the current image. This is normally seen after successful publish for a push image which has been scheduled to trigger at a later time. * FAILED: This is the state of the pending image if creation of instant clone operation has failed or timed out. | [optional] 
**InstantCloneScheduledMaintenanceData** | Pointer to [**FarmScheduledMaintenanceInfo**](FarmScheduledMaintenanceInfo.md) |  | [optional] 
**LastProvisioningError** | Pointer to **string** | String message detailing the last provisioning error on this farm while stop_provisioning_on_error is enabled. | [optional] 
**LastProvisioningErrorTime** | Pointer to **int64** | Time the last provisioning error occurred on this farm while stop_provisioning_on_error is enabled. Measured as epoch time. | [optional] 

## Methods

### NewFarmProvisioningStatusInfo

`func NewFarmProvisioningStatusInfo() *FarmProvisioningStatusInfo`

NewFarmProvisioningStatusInfo instantiates a new FarmProvisioningStatusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmProvisioningStatusInfoWithDefaults

`func NewFarmProvisioningStatusInfoWithDefaults() *FarmProvisioningStatusInfo`

NewFarmProvisioningStatusInfoWithDefaults instantiates a new FarmProvisioningStatusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstantCloneCurrentImageState

`func (o *FarmProvisioningStatusInfo) GetInstantCloneCurrentImageState() string`

GetInstantCloneCurrentImageState returns the InstantCloneCurrentImageState field if non-nil, zero value otherwise.

### GetInstantCloneCurrentImageStateOk

`func (o *FarmProvisioningStatusInfo) GetInstantCloneCurrentImageStateOk() (*string, bool)`

GetInstantCloneCurrentImageStateOk returns a tuple with the InstantCloneCurrentImageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneCurrentImageState

`func (o *FarmProvisioningStatusInfo) SetInstantCloneCurrentImageState(v string)`

SetInstantCloneCurrentImageState sets InstantCloneCurrentImageState field to given value.

### HasInstantCloneCurrentImageState

`func (o *FarmProvisioningStatusInfo) HasInstantCloneCurrentImageState() bool`

HasInstantCloneCurrentImageState returns a boolean if a field has been set.

### GetInstantCloneOperation

`func (o *FarmProvisioningStatusInfo) GetInstantCloneOperation() string`

GetInstantCloneOperation returns the InstantCloneOperation field if non-nil, zero value otherwise.

### GetInstantCloneOperationOk

`func (o *FarmProvisioningStatusInfo) GetInstantCloneOperationOk() (*string, bool)`

GetInstantCloneOperationOk returns a tuple with the InstantCloneOperation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneOperation

`func (o *FarmProvisioningStatusInfo) SetInstantCloneOperation(v string)`

SetInstantCloneOperation sets InstantCloneOperation field to given value.

### HasInstantCloneOperation

`func (o *FarmProvisioningStatusInfo) HasInstantCloneOperation() bool`

HasInstantCloneOperation returns a boolean if a field has been set.

### GetInstantCloneOperationTime

`func (o *FarmProvisioningStatusInfo) GetInstantCloneOperationTime() int64`

GetInstantCloneOperationTime returns the InstantCloneOperationTime field if non-nil, zero value otherwise.

### GetInstantCloneOperationTimeOk

`func (o *FarmProvisioningStatusInfo) GetInstantCloneOperationTimeOk() (*int64, bool)`

GetInstantCloneOperationTimeOk returns a tuple with the InstantCloneOperationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneOperationTime

`func (o *FarmProvisioningStatusInfo) SetInstantCloneOperationTime(v int64)`

SetInstantCloneOperationTime sets InstantCloneOperationTime field to given value.

### HasInstantCloneOperationTime

`func (o *FarmProvisioningStatusInfo) HasInstantCloneOperationTime() bool`

HasInstantCloneOperationTime returns a boolean if a field has been set.

### GetInstantClonePendingImStreamId

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImStreamId() string`

GetInstantClonePendingImStreamId returns the InstantClonePendingImStreamId field if non-nil, zero value otherwise.

### GetInstantClonePendingImStreamIdOk

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImStreamIdOk() (*string, bool)`

GetInstantClonePendingImStreamIdOk returns a tuple with the InstantClonePendingImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImStreamId

`func (o *FarmProvisioningStatusInfo) SetInstantClonePendingImStreamId(v string)`

SetInstantClonePendingImStreamId sets InstantClonePendingImStreamId field to given value.

### HasInstantClonePendingImStreamId

`func (o *FarmProvisioningStatusInfo) HasInstantClonePendingImStreamId() bool`

HasInstantClonePendingImStreamId returns a boolean if a field has been set.

### GetInstantClonePendingImTagId

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImTagId() string`

GetInstantClonePendingImTagId returns the InstantClonePendingImTagId field if non-nil, zero value otherwise.

### GetInstantClonePendingImTagIdOk

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImTagIdOk() (*string, bool)`

GetInstantClonePendingImTagIdOk returns a tuple with the InstantClonePendingImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImTagId

`func (o *FarmProvisioningStatusInfo) SetInstantClonePendingImTagId(v string)`

SetInstantClonePendingImTagId sets InstantClonePendingImTagId field to given value.

### HasInstantClonePendingImTagId

`func (o *FarmProvisioningStatusInfo) HasInstantClonePendingImTagId() bool`

HasInstantClonePendingImTagId returns a boolean if a field has been set.

### GetInstantClonePendingImageError

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageError() string`

GetInstantClonePendingImageError returns the InstantClonePendingImageError field if non-nil, zero value otherwise.

### GetInstantClonePendingImageErrorOk

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageErrorOk() (*string, bool)`

GetInstantClonePendingImageErrorOk returns a tuple with the InstantClonePendingImageError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageError

`func (o *FarmProvisioningStatusInfo) SetInstantClonePendingImageError(v string)`

SetInstantClonePendingImageError sets InstantClonePendingImageError field to given value.

### HasInstantClonePendingImageError

`func (o *FarmProvisioningStatusInfo) HasInstantClonePendingImageError() bool`

HasInstantClonePendingImageError returns a boolean if a field has been set.

### GetInstantClonePendingImageParentVmId

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageParentVmId() string`

GetInstantClonePendingImageParentVmId returns the InstantClonePendingImageParentVmId field if non-nil, zero value otherwise.

### GetInstantClonePendingImageParentVmIdOk

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageParentVmIdOk() (*string, bool)`

GetInstantClonePendingImageParentVmIdOk returns a tuple with the InstantClonePendingImageParentVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageParentVmId

`func (o *FarmProvisioningStatusInfo) SetInstantClonePendingImageParentVmId(v string)`

SetInstantClonePendingImageParentVmId sets InstantClonePendingImageParentVmId field to given value.

### HasInstantClonePendingImageParentVmId

`func (o *FarmProvisioningStatusInfo) HasInstantClonePendingImageParentVmId() bool`

HasInstantClonePendingImageParentVmId returns a boolean if a field has been set.

### GetInstantClonePendingImageProgress

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageProgress() int32`

GetInstantClonePendingImageProgress returns the InstantClonePendingImageProgress field if non-nil, zero value otherwise.

### GetInstantClonePendingImageProgressOk

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageProgressOk() (*int32, bool)`

GetInstantClonePendingImageProgressOk returns a tuple with the InstantClonePendingImageProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageProgress

`func (o *FarmProvisioningStatusInfo) SetInstantClonePendingImageProgress(v int32)`

SetInstantClonePendingImageProgress sets InstantClonePendingImageProgress field to given value.

### HasInstantClonePendingImageProgress

`func (o *FarmProvisioningStatusInfo) HasInstantClonePendingImageProgress() bool`

HasInstantClonePendingImageProgress returns a boolean if a field has been set.

### GetInstantClonePendingImageSnapshotId

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageSnapshotId() string`

GetInstantClonePendingImageSnapshotId returns the InstantClonePendingImageSnapshotId field if non-nil, zero value otherwise.

### GetInstantClonePendingImageSnapshotIdOk

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageSnapshotIdOk() (*string, bool)`

GetInstantClonePendingImageSnapshotIdOk returns a tuple with the InstantClonePendingImageSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageSnapshotId

`func (o *FarmProvisioningStatusInfo) SetInstantClonePendingImageSnapshotId(v string)`

SetInstantClonePendingImageSnapshotId sets InstantClonePendingImageSnapshotId field to given value.

### HasInstantClonePendingImageSnapshotId

`func (o *FarmProvisioningStatusInfo) HasInstantClonePendingImageSnapshotId() bool`

HasInstantClonePendingImageSnapshotId returns a boolean if a field has been set.

### GetInstantClonePendingImageState

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageState() string`

GetInstantClonePendingImageState returns the InstantClonePendingImageState field if non-nil, zero value otherwise.

### GetInstantClonePendingImageStateOk

`func (o *FarmProvisioningStatusInfo) GetInstantClonePendingImageStateOk() (*string, bool)`

GetInstantClonePendingImageStateOk returns a tuple with the InstantClonePendingImageState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantClonePendingImageState

`func (o *FarmProvisioningStatusInfo) SetInstantClonePendingImageState(v string)`

SetInstantClonePendingImageState sets InstantClonePendingImageState field to given value.

### HasInstantClonePendingImageState

`func (o *FarmProvisioningStatusInfo) HasInstantClonePendingImageState() bool`

HasInstantClonePendingImageState returns a boolean if a field has been set.

### GetInstantCloneScheduledMaintenanceData

`func (o *FarmProvisioningStatusInfo) GetInstantCloneScheduledMaintenanceData() FarmScheduledMaintenanceInfo`

GetInstantCloneScheduledMaintenanceData returns the InstantCloneScheduledMaintenanceData field if non-nil, zero value otherwise.

### GetInstantCloneScheduledMaintenanceDataOk

`func (o *FarmProvisioningStatusInfo) GetInstantCloneScheduledMaintenanceDataOk() (*FarmScheduledMaintenanceInfo, bool)`

GetInstantCloneScheduledMaintenanceDataOk returns a tuple with the InstantCloneScheduledMaintenanceData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneScheduledMaintenanceData

`func (o *FarmProvisioningStatusInfo) SetInstantCloneScheduledMaintenanceData(v FarmScheduledMaintenanceInfo)`

SetInstantCloneScheduledMaintenanceData sets InstantCloneScheduledMaintenanceData field to given value.

### HasInstantCloneScheduledMaintenanceData

`func (o *FarmProvisioningStatusInfo) HasInstantCloneScheduledMaintenanceData() bool`

HasInstantCloneScheduledMaintenanceData returns a boolean if a field has been set.

### GetLastProvisioningError

`func (o *FarmProvisioningStatusInfo) GetLastProvisioningError() string`

GetLastProvisioningError returns the LastProvisioningError field if non-nil, zero value otherwise.

### GetLastProvisioningErrorOk

`func (o *FarmProvisioningStatusInfo) GetLastProvisioningErrorOk() (*string, bool)`

GetLastProvisioningErrorOk returns a tuple with the LastProvisioningError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProvisioningError

`func (o *FarmProvisioningStatusInfo) SetLastProvisioningError(v string)`

SetLastProvisioningError sets LastProvisioningError field to given value.

### HasLastProvisioningError

`func (o *FarmProvisioningStatusInfo) HasLastProvisioningError() bool`

HasLastProvisioningError returns a boolean if a field has been set.

### GetLastProvisioningErrorTime

`func (o *FarmProvisioningStatusInfo) GetLastProvisioningErrorTime() int64`

GetLastProvisioningErrorTime returns the LastProvisioningErrorTime field if non-nil, zero value otherwise.

### GetLastProvisioningErrorTimeOk

`func (o *FarmProvisioningStatusInfo) GetLastProvisioningErrorTimeOk() (*int64, bool)`

GetLastProvisioningErrorTimeOk returns a tuple with the LastProvisioningErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastProvisioningErrorTime

`func (o *FarmProvisioningStatusInfo) SetLastProvisioningErrorTime(v int64)`

SetLastProvisioningErrorTime sets LastProvisioningErrorTime field to given value.

### HasLastProvisioningErrorTime

`func (o *FarmProvisioningStatusInfo) HasLastProvisioningErrorTime() bool`

HasLastProvisioningErrorTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


