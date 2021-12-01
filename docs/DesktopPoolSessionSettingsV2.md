# DesktopPoolSessionSettingsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleSessionsPerUser** | Pointer to **bool** | Indicates whether multiple sessions are allowed per user in case of Floating User Assignment. Default value is false. | [optional] 
**AllowUsersToResetMachines** | Pointer to **bool** | Indicates whether the user can be allowed to reset or restart their machines. Default value is false. | [optional] 
**DeleteOrRefreshMachineAfterLogoff** | Pointer to **string** | Whether machines are to be deleted or refreshed after logoff in case of Floating User Assignment.This is applicable for automated desktops with virtual machines names based onpattern naming. This is not applicable for desktops that are using specified naming since dynamic creation and deletion of VMs is not supported.For Instant clone desktops this setting can only be set to DELETE. Default value is NEVER. * NEVER: Never delete or refresh the machine in the desktop pool. * DELETE: Delete the machine after user logoff. * REFRESH: Refresh the machine after user logoff. | [optional] 
**DisconnectedSessionTimeoutMinutes** | Pointer to **int32** | Disconnected sessions timeout (in minutes). Will be set when disconnected_session_timeout_policy is set to AFTER. | [optional] 
**DisconnectedSessionTimeoutPolicy** | Pointer to **string** | Log-off policy after disconnected session. Default value is NEVER. * IMMEDIATELY: Immediately Logoff after user disconnect. * AFTER: Logoff after the specified number of minutes after user disconnect. * NEVER: Do not logoff after user disconnect. | [optional] 
**EmptySessionTimeoutMinutes** | Pointer to **int32** | Application empty session timeout (in minutes). An empty session (that has no remote-ablewindow) is disconnected after the timeout. Default value is 1.Will be set when the empty_session_timeout_policy set to AFTER. | [optional] 
**EmptySessionTimeoutPolicy** | Pointer to **string** | Application empty session timeout policy. Default value is AFTER. * IMMEDIATE: Empty session will be disconnected immediately. * NEVER: Empty session will never disconnected. * AFTER: Empty session will be disconnected after specified number of minutes. | [optional] 
**LogoffAfterTimeout** | Pointer to **bool** | Indicates whether the empty application sessions are logged off (true) or disconnected (false) after timeout.Default value is false. | [optional] 
**PowerPolicy** | Pointer to **string** | Power policy for the machines in the desktop pool after logoff. This setting is only relevant for managed machines.Default value is TAKE_NO_POWER_ACTION.For Instant clone desktops this setting can only be set to ALWAYS_POWERED_ON. * TAKE_NO_POWER_ACTION: No action will be taken when user logs off. * ALWAYS_POWERED_ON: Ensure machines in the Desktop pool are always powered on. The connection server will monitor and power on machines as necessary. * SUSPEND: Suspend when a user logs off or when desktop pool is no longer keeping a machine as a spare. This does not affect spare and newly provisioned machines. * POWER_OFF: Power off when a user logs off or when desktop pool is no longer keeping a machine as a spare. This does not affect spare and newly provisioned machines. | [optional] 
**PreLaunchSessionTimeoutMinutes** | Pointer to **int32** | Application pre-launch session timeout (in minutes). A pre-launch session is disconnected after the timeout. Default value is 10.Will be required when the pre-launch session timeout policy is set to AFTER. | [optional] 
**PreLaunchSessionTimeoutPolicy** | Pointer to **string** | Application pre-launch session timeout policy. Default value is AFTER. * AFTER: Pre-launched session is disconnected after specified number of minutes. * NEVER: Pre-launched session is never disconnected. | [optional] 
**RefreshOsDiskAfterLogoff** | Pointer to **string** | Whether and when to refresh the OS disks for dedicated-assignment, linked-clone and instant-clone machines.Default value is NEVER. * NEVER: The OS disk is never refreshed. * ALWAYS: The OS disk is refreshed every time the user logs off. * EVERY: The OS disk is refreshed at regular intervals of a specified number of days. The number of days is counted from the last refresh, or from the initial provisioning if no refresh has occurred yet. For example, if the specified value is 3 days, and three days have passed since the last refresh, the machine is refreshed after the user logs off. * AT_SIZE: The OS disk is refreshed when its current size reaches a specified percentage of its maximum allowable size. The maximum size of a linked clone&#39;s OS disk is the size of the replica&#39;s OS disk. With this option, the size of the linked clone&#39;s OS disk in the datastore is compared to maximum allowable size. This disk-utilization percentage does not reflect disk usage that you might see inside the machine&#39;s guest operating system. | [optional] 
**RefreshPeriodDaysForReplicaOsDisk** | Pointer to **int32** | Regular interval at which to refresh the OS disk. Will be set when refresh_os_disk_after_logoff set to EVERY. | [optional] 
**RefreshThresholdPercentageForReplicaOsDisk** | Pointer to **int32** | With the &#39;AT_SIZE&#39; option for refreshOsDiskAfterLogoff, the size of the linked clone&#39;s OS diskin the datastore is compared to its maximum allowable size. This disk-utilization percentage does not reflect disk usage that you might see inside the machine&#39;s guest operating system.Will be set when refresh_os_disk_after_logoff set to AT_SIZE. | [optional] 

## Methods

### NewDesktopPoolSessionSettingsV2

`func NewDesktopPoolSessionSettingsV2() *DesktopPoolSessionSettingsV2`

NewDesktopPoolSessionSettingsV2 instantiates a new DesktopPoolSessionSettingsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolSessionSettingsV2WithDefaults

`func NewDesktopPoolSessionSettingsV2WithDefaults() *DesktopPoolSessionSettingsV2`

NewDesktopPoolSessionSettingsV2WithDefaults instantiates a new DesktopPoolSessionSettingsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleSessionsPerUser

`func (o *DesktopPoolSessionSettingsV2) GetAllowMultipleSessionsPerUser() bool`

GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field if non-nil, zero value otherwise.

### GetAllowMultipleSessionsPerUserOk

`func (o *DesktopPoolSessionSettingsV2) GetAllowMultipleSessionsPerUserOk() (*bool, bool)`

GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSessionsPerUser

`func (o *DesktopPoolSessionSettingsV2) SetAllowMultipleSessionsPerUser(v bool)`

SetAllowMultipleSessionsPerUser sets AllowMultipleSessionsPerUser field to given value.

### HasAllowMultipleSessionsPerUser

`func (o *DesktopPoolSessionSettingsV2) HasAllowMultipleSessionsPerUser() bool`

HasAllowMultipleSessionsPerUser returns a boolean if a field has been set.

### GetAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsV2) GetAllowUsersToResetMachines() bool`

GetAllowUsersToResetMachines returns the AllowUsersToResetMachines field if non-nil, zero value otherwise.

### GetAllowUsersToResetMachinesOk

`func (o *DesktopPoolSessionSettingsV2) GetAllowUsersToResetMachinesOk() (*bool, bool)`

GetAllowUsersToResetMachinesOk returns a tuple with the AllowUsersToResetMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsV2) SetAllowUsersToResetMachines(v bool)`

SetAllowUsersToResetMachines sets AllowUsersToResetMachines field to given value.

### HasAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsV2) HasAllowUsersToResetMachines() bool`

HasAllowUsersToResetMachines returns a boolean if a field has been set.

### GetDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsV2) GetDeleteOrRefreshMachineAfterLogoff() string`

GetDeleteOrRefreshMachineAfterLogoff returns the DeleteOrRefreshMachineAfterLogoff field if non-nil, zero value otherwise.

### GetDeleteOrRefreshMachineAfterLogoffOk

`func (o *DesktopPoolSessionSettingsV2) GetDeleteOrRefreshMachineAfterLogoffOk() (*string, bool)`

GetDeleteOrRefreshMachineAfterLogoffOk returns a tuple with the DeleteOrRefreshMachineAfterLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsV2) SetDeleteOrRefreshMachineAfterLogoff(v string)`

SetDeleteOrRefreshMachineAfterLogoff sets DeleteOrRefreshMachineAfterLogoff field to given value.

### HasDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsV2) HasDeleteOrRefreshMachineAfterLogoff() bool`

HasDeleteOrRefreshMachineAfterLogoff returns a boolean if a field has been set.

### GetDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) GetDisconnectedSessionTimeoutMinutes() int32`

GetDisconnectedSessionTimeoutMinutes returns the DisconnectedSessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetDisconnectedSessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsV2) GetDisconnectedSessionTimeoutMinutesOk() (*int32, bool)`

GetDisconnectedSessionTimeoutMinutesOk returns a tuple with the DisconnectedSessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) SetDisconnectedSessionTimeoutMinutes(v int32)`

SetDisconnectedSessionTimeoutMinutes sets DisconnectedSessionTimeoutMinutes field to given value.

### HasDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) HasDisconnectedSessionTimeoutMinutes() bool`

HasDisconnectedSessionTimeoutMinutes returns a boolean if a field has been set.

### GetDisconnectedSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) GetDisconnectedSessionTimeoutPolicy() string`

GetDisconnectedSessionTimeoutPolicy returns the DisconnectedSessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetDisconnectedSessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsV2) GetDisconnectedSessionTimeoutPolicyOk() (*string, bool)`

GetDisconnectedSessionTimeoutPolicyOk returns a tuple with the DisconnectedSessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) SetDisconnectedSessionTimeoutPolicy(v string)`

SetDisconnectedSessionTimeoutPolicy sets DisconnectedSessionTimeoutPolicy field to given value.

### HasDisconnectedSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) HasDisconnectedSessionTimeoutPolicy() bool`

HasDisconnectedSessionTimeoutPolicy returns a boolean if a field has been set.

### GetEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) GetEmptySessionTimeoutMinutes() int32`

GetEmptySessionTimeoutMinutes returns the EmptySessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetEmptySessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsV2) GetEmptySessionTimeoutMinutesOk() (*int32, bool)`

GetEmptySessionTimeoutMinutesOk returns a tuple with the EmptySessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) SetEmptySessionTimeoutMinutes(v int32)`

SetEmptySessionTimeoutMinutes sets EmptySessionTimeoutMinutes field to given value.

### HasEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) HasEmptySessionTimeoutMinutes() bool`

HasEmptySessionTimeoutMinutes returns a boolean if a field has been set.

### GetEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) GetEmptySessionTimeoutPolicy() string`

GetEmptySessionTimeoutPolicy returns the EmptySessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetEmptySessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsV2) GetEmptySessionTimeoutPolicyOk() (*string, bool)`

GetEmptySessionTimeoutPolicyOk returns a tuple with the EmptySessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) SetEmptySessionTimeoutPolicy(v string)`

SetEmptySessionTimeoutPolicy sets EmptySessionTimeoutPolicy field to given value.

### HasEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) HasEmptySessionTimeoutPolicy() bool`

HasEmptySessionTimeoutPolicy returns a boolean if a field has been set.

### GetLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsV2) GetLogoffAfterTimeout() bool`

GetLogoffAfterTimeout returns the LogoffAfterTimeout field if non-nil, zero value otherwise.

### GetLogoffAfterTimeoutOk

`func (o *DesktopPoolSessionSettingsV2) GetLogoffAfterTimeoutOk() (*bool, bool)`

GetLogoffAfterTimeoutOk returns a tuple with the LogoffAfterTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsV2) SetLogoffAfterTimeout(v bool)`

SetLogoffAfterTimeout sets LogoffAfterTimeout field to given value.

### HasLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsV2) HasLogoffAfterTimeout() bool`

HasLogoffAfterTimeout returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *DesktopPoolSessionSettingsV2) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *DesktopPoolSessionSettingsV2) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *DesktopPoolSessionSettingsV2) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *DesktopPoolSessionSettingsV2) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) GetPreLaunchSessionTimeoutMinutes() int32`

GetPreLaunchSessionTimeoutMinutes returns the PreLaunchSessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetPreLaunchSessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsV2) GetPreLaunchSessionTimeoutMinutesOk() (*int32, bool)`

GetPreLaunchSessionTimeoutMinutesOk returns a tuple with the PreLaunchSessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) SetPreLaunchSessionTimeoutMinutes(v int32)`

SetPreLaunchSessionTimeoutMinutes sets PreLaunchSessionTimeoutMinutes field to given value.

### HasPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsV2) HasPreLaunchSessionTimeoutMinutes() bool`

HasPreLaunchSessionTimeoutMinutes returns a boolean if a field has been set.

### GetPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) GetPreLaunchSessionTimeoutPolicy() string`

GetPreLaunchSessionTimeoutPolicy returns the PreLaunchSessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetPreLaunchSessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsV2) GetPreLaunchSessionTimeoutPolicyOk() (*string, bool)`

GetPreLaunchSessionTimeoutPolicyOk returns a tuple with the PreLaunchSessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) SetPreLaunchSessionTimeoutPolicy(v string)`

SetPreLaunchSessionTimeoutPolicy sets PreLaunchSessionTimeoutPolicy field to given value.

### HasPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsV2) HasPreLaunchSessionTimeoutPolicy() bool`

HasPreLaunchSessionTimeoutPolicy returns a boolean if a field has been set.

### GetRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsV2) GetRefreshOsDiskAfterLogoff() string`

GetRefreshOsDiskAfterLogoff returns the RefreshOsDiskAfterLogoff field if non-nil, zero value otherwise.

### GetRefreshOsDiskAfterLogoffOk

`func (o *DesktopPoolSessionSettingsV2) GetRefreshOsDiskAfterLogoffOk() (*string, bool)`

GetRefreshOsDiskAfterLogoffOk returns a tuple with the RefreshOsDiskAfterLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsV2) SetRefreshOsDiskAfterLogoff(v string)`

SetRefreshOsDiskAfterLogoff sets RefreshOsDiskAfterLogoff field to given value.

### HasRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsV2) HasRefreshOsDiskAfterLogoff() bool`

HasRefreshOsDiskAfterLogoff returns a boolean if a field has been set.

### GetRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsV2) GetRefreshPeriodDaysForReplicaOsDisk() int32`

GetRefreshPeriodDaysForReplicaOsDisk returns the RefreshPeriodDaysForReplicaOsDisk field if non-nil, zero value otherwise.

### GetRefreshPeriodDaysForReplicaOsDiskOk

`func (o *DesktopPoolSessionSettingsV2) GetRefreshPeriodDaysForReplicaOsDiskOk() (*int32, bool)`

GetRefreshPeriodDaysForReplicaOsDiskOk returns a tuple with the RefreshPeriodDaysForReplicaOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsV2) SetRefreshPeriodDaysForReplicaOsDisk(v int32)`

SetRefreshPeriodDaysForReplicaOsDisk sets RefreshPeriodDaysForReplicaOsDisk field to given value.

### HasRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsV2) HasRefreshPeriodDaysForReplicaOsDisk() bool`

HasRefreshPeriodDaysForReplicaOsDisk returns a boolean if a field has been set.

### GetRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsV2) GetRefreshThresholdPercentageForReplicaOsDisk() int32`

GetRefreshThresholdPercentageForReplicaOsDisk returns the RefreshThresholdPercentageForReplicaOsDisk field if non-nil, zero value otherwise.

### GetRefreshThresholdPercentageForReplicaOsDiskOk

`func (o *DesktopPoolSessionSettingsV2) GetRefreshThresholdPercentageForReplicaOsDiskOk() (*int32, bool)`

GetRefreshThresholdPercentageForReplicaOsDiskOk returns a tuple with the RefreshThresholdPercentageForReplicaOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsV2) SetRefreshThresholdPercentageForReplicaOsDisk(v int32)`

SetRefreshThresholdPercentageForReplicaOsDisk sets RefreshThresholdPercentageForReplicaOsDisk field to given value.

### HasRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsV2) HasRefreshThresholdPercentageForReplicaOsDisk() bool`

HasRefreshThresholdPercentageForReplicaOsDisk returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


