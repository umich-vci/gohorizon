# FarmMaintenanceSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImStreamId** | Pointer to **string** | New image management stream for the farm. Either parent_vm_id and snapshot_id or im_stream_id and im_tag_id are to be specified. | [optional] 
**ImTagId** | Pointer to **string** | New image management tag for the farm. This tag must be within the im_stream_id. Either parent_vm_id and snapshot_id or im_stream_id and im_tag_id are to be specified. | [optional] 
**LogoffPolicy** | **string** | Determines when to perform the operation on RDS servers which have an active session. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions. | 
**MaintenanceMode** | **string** | The mode of schedule maintenance for Instant Clone Farm. * IMMEDIATE: All server VMs will be refreshed once, immediately or at user scheduled time. * RECURRING: All server VMs will be periodically refreshed based on FarmInstantCloneRecurringMaintenancePeriod and StartTime | 
**NextScheduledTime** | Pointer to **int64** | Time when next scheduled maintenance would happen. If maintenance_mode is set to IMMEDIATE and next_scheduled_time is not set, maintenance will begin immediately. If maintenance_mode is set to RECURRING and next_scheduled_time is not set, it will be calculated based on current recurring maintenance configuration. If the value is in the past, maintenance will begin immediately. Measured as epoch time. | [optional] 
**ParentVmId** | Pointer to **string** | New base image VM for the instant clone farm. This must be in the same datacenter as the base image of the farm. Either parent_vm_id and snapshot_id or im_stream_id and im_tag_id are to be specified. | [optional] 
**RecurringMaintenanceSettings** | Pointer to [**FarmRecurringMaintenanceSettingsSpec**](FarmRecurringMaintenanceSettingsSpec.md) |  | [optional] 
**SnapshotId** | Pointer to **string** | New base image snapshot for the instant clone farm. This must be a snapshot of the parent_vm_id. Either parent_vm_id and snapshot_id or im_stream_id and im_tag_id are to be specified. | [optional] 
**StopOnFirstError** | Pointer to **bool** | Indicates whether the operation should stop on first error. Default value is true. | [optional] 

## Methods

### NewFarmMaintenanceSpec

`func NewFarmMaintenanceSpec(logoffPolicy string, maintenanceMode string, ) *FarmMaintenanceSpec`

NewFarmMaintenanceSpec instantiates a new FarmMaintenanceSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmMaintenanceSpecWithDefaults

`func NewFarmMaintenanceSpecWithDefaults() *FarmMaintenanceSpec`

NewFarmMaintenanceSpecWithDefaults instantiates a new FarmMaintenanceSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImStreamId

`func (o *FarmMaintenanceSpec) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *FarmMaintenanceSpec) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *FarmMaintenanceSpec) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.

### HasImStreamId

`func (o *FarmMaintenanceSpec) HasImStreamId() bool`

HasImStreamId returns a boolean if a field has been set.

### GetImTagId

`func (o *FarmMaintenanceSpec) GetImTagId() string`

GetImTagId returns the ImTagId field if non-nil, zero value otherwise.

### GetImTagIdOk

`func (o *FarmMaintenanceSpec) GetImTagIdOk() (*string, bool)`

GetImTagIdOk returns a tuple with the ImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImTagId

`func (o *FarmMaintenanceSpec) SetImTagId(v string)`

SetImTagId sets ImTagId field to given value.

### HasImTagId

`func (o *FarmMaintenanceSpec) HasImTagId() bool`

HasImTagId returns a boolean if a field has been set.

### GetLogoffPolicy

`func (o *FarmMaintenanceSpec) GetLogoffPolicy() string`

GetLogoffPolicy returns the LogoffPolicy field if non-nil, zero value otherwise.

### GetLogoffPolicyOk

`func (o *FarmMaintenanceSpec) GetLogoffPolicyOk() (*string, bool)`

GetLogoffPolicyOk returns a tuple with the LogoffPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPolicy

`func (o *FarmMaintenanceSpec) SetLogoffPolicy(v string)`

SetLogoffPolicy sets LogoffPolicy field to given value.


### GetMaintenanceMode

`func (o *FarmMaintenanceSpec) GetMaintenanceMode() string`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *FarmMaintenanceSpec) GetMaintenanceModeOk() (*string, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *FarmMaintenanceSpec) SetMaintenanceMode(v string)`

SetMaintenanceMode sets MaintenanceMode field to given value.


### GetNextScheduledTime

`func (o *FarmMaintenanceSpec) GetNextScheduledTime() int64`

GetNextScheduledTime returns the NextScheduledTime field if non-nil, zero value otherwise.

### GetNextScheduledTimeOk

`func (o *FarmMaintenanceSpec) GetNextScheduledTimeOk() (*int64, bool)`

GetNextScheduledTimeOk returns a tuple with the NextScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledTime

`func (o *FarmMaintenanceSpec) SetNextScheduledTime(v int64)`

SetNextScheduledTime sets NextScheduledTime field to given value.

### HasNextScheduledTime

`func (o *FarmMaintenanceSpec) HasNextScheduledTime() bool`

HasNextScheduledTime returns a boolean if a field has been set.

### GetParentVmId

`func (o *FarmMaintenanceSpec) GetParentVmId() string`

GetParentVmId returns the ParentVmId field if non-nil, zero value otherwise.

### GetParentVmIdOk

`func (o *FarmMaintenanceSpec) GetParentVmIdOk() (*string, bool)`

GetParentVmIdOk returns a tuple with the ParentVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmId

`func (o *FarmMaintenanceSpec) SetParentVmId(v string)`

SetParentVmId sets ParentVmId field to given value.

### HasParentVmId

`func (o *FarmMaintenanceSpec) HasParentVmId() bool`

HasParentVmId returns a boolean if a field has been set.

### GetRecurringMaintenanceSettings

`func (o *FarmMaintenanceSpec) GetRecurringMaintenanceSettings() FarmRecurringMaintenanceSettingsSpec`

GetRecurringMaintenanceSettings returns the RecurringMaintenanceSettings field if non-nil, zero value otherwise.

### GetRecurringMaintenanceSettingsOk

`func (o *FarmMaintenanceSpec) GetRecurringMaintenanceSettingsOk() (*FarmRecurringMaintenanceSettingsSpec, bool)`

GetRecurringMaintenanceSettingsOk returns a tuple with the RecurringMaintenanceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringMaintenanceSettings

`func (o *FarmMaintenanceSpec) SetRecurringMaintenanceSettings(v FarmRecurringMaintenanceSettingsSpec)`

SetRecurringMaintenanceSettings sets RecurringMaintenanceSettings field to given value.

### HasRecurringMaintenanceSettings

`func (o *FarmMaintenanceSpec) HasRecurringMaintenanceSettings() bool`

HasRecurringMaintenanceSettings returns a boolean if a field has been set.

### GetSnapshotId

`func (o *FarmMaintenanceSpec) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *FarmMaintenanceSpec) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *FarmMaintenanceSpec) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *FarmMaintenanceSpec) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetStopOnFirstError

`func (o *FarmMaintenanceSpec) GetStopOnFirstError() bool`

GetStopOnFirstError returns the StopOnFirstError field if non-nil, zero value otherwise.

### GetStopOnFirstErrorOk

`func (o *FarmMaintenanceSpec) GetStopOnFirstErrorOk() (*bool, bool)`

GetStopOnFirstErrorOk returns a tuple with the StopOnFirstError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnFirstError

`func (o *FarmMaintenanceSpec) SetStopOnFirstError(v bool)`

SetStopOnFirstError sets StopOnFirstError field to given value.

### HasStopOnFirstError

`func (o *FarmMaintenanceSpec) HasStopOnFirstError() bool`

HasStopOnFirstError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


