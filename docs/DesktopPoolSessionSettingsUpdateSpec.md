# DesktopPoolSessionSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleSessionsPerUser** | **bool** | Indicates whether multiple sessions are allowed per user in case of Floating User Assignment. | 
**AllowUsersToResetMachines** | Pointer to **bool** | Indicates whether the user can be allowed to reset or restart their machines. | [optional] 
**DeleteOrRefreshMachineAfterLogoff** | Pointer to **string** | Indicates whether machines are to be deleted or refreshed after logoff in case of Floating User Assignment. This is applicable for automated desktop pools with virtual machines names based on pattern naming. This is not applicable for desktop pools that are using specified naming since dynamic creation and deletion of VMs is not supported. For Instant clone desktop pools this setting can only be set to DELETE. * NEVER: Never delete or refresh the machine in the desktop pool. * DELETE: Delete the machine after user logoff. * REFRESH: Refresh the machine after user logoff. | [optional] 
**DisconnectedSessionTimeoutMinutes** | Pointer to **int32** | Disconnected sessions timeout (in minutes). This is required if disconnected_session_timeout_policy is set to AFTER. | [optional] 
**DisconnectedSessionTimeoutPolicy** | **string** | Log-off policy after disconnected session. * IMMEDIATELY: Immediately Logoff after user disconnect. * AFTER: Logoff after the specified number of minutes after user disconnect. * NEVER: Do not logoff after user disconnect. | 
**EmptySessionTimeoutMinutes** | Pointer to **int32** | Desktop Pool empty session timeout (in minutes). An empty session (that has no remote-ablewindow) is disconnected after the timeout. This is required if empty_session_timeout_policy is set to AFTER. | [optional] 
**EmptySessionTimeoutPolicy** | Pointer to **string** | Applicable only when session_type is APPLICATION or DESKTOP_AND_APPLICATION. &lt;br&gt;Desktop Pool empty session timeout policy. * IMMEDIATE: Empty session will be disconnected immediately. * NEVER: Empty session will never disconnected. * AFTER: Empty session will be disconnected after specified number of minutes. | [optional] 
**LogoffAfterTimeout** | Pointer to **bool** | Applicable only when session_type is APPLICATION or DESKTOP_AND_APPLICATION. Indicates whether the empty desktop pool sessions are logged off (true) or disconnected (false) after timeout. | [optional] 
**PowerPolicy** | Pointer to **string** | Power policy for the machines in the desktop pool after logoff. This setting is only relevant for managed machines. For Instant clone desktop pools this setting can only be set to ALWAYS_POWERED_ON. * TAKE_NO_POWER_ACTION: No action will be taken when user logs off. * ALWAYS_POWERED_ON: Ensure machines in the Desktop pool are always powered on. The connection server will monitor and power on machines as necessary. * SUSPEND: Suspend when a user logs off or when desktop pool is no longer keeping a machine as a spare. This does not affect spare and newly provisioned machines. * POWER_OFF: Power off when a user logs off or when desktop pool is no longer keeping a machine as a spare. This does not affect spare and newly provisioned machines. | [optional] 
**PreLaunchSessionTimeoutMinutes** | Pointer to **int32** | Desktop Pool pre-launch session timeout (in minutes). A pre-launch session is disconnected after the timeout. This is required if pre-launch session timeout policy is set to AFTER. | [optional] 
**PreLaunchSessionTimeoutPolicy** | Pointer to **string** | Desktop Pool pre-launch session timeout policy. * AFTER: Pre-launched session is disconnected after specified number of minutes. * NEVER: Pre-launched session is never disconnected. | [optional] 
**RefreshOsDiskAfterLogoff** | Pointer to **string** | Indicates whether and when to refresh the OS disks for dedicated instant-clone desktop pools. * NEVER: The OS disk is never refreshed. * ALWAYS: The OS disk is refreshed every time the user logs off. * EVERY: The OS disk is refreshed at regular intervals of a specified number of days. The number of days is counted from the last refresh, or from the initial provisioning if no refresh has occurred yet. For example, if the specified value is 3 days, and three days have passed since the last refresh, the machine is refreshed after the user logs off. * AT_SIZE: The OS disk is refreshed when its current size reaches a specified percentage of its maximum allowable size. The maximum size of a linked clone&#39;s OS disk is the size of the replica&#39;s OS disk. With this option, the size of the linked clone&#39;s OS disk in the datastore is compared to maximum allowable size. This disk-utilization percentage does not reflect disk usage that you might see inside the machine&#39;s guest operating system. | [optional] 
**RefreshPeriodDaysForReplicaOsDisk** | Pointer to **int32** | Regular interval at which to refresh the OS disk. This is required if when refresh_os_disk_after_logoff set to EVERY. | [optional] 
**RefreshThresholdPercentageForReplicaOsDisk** | Pointer to **int32** | With the AT_SIZE option for refresh_os_disk_after_logoff, the size of the instant clone&#39;s OS diskin the datastore is compared to its maximum allowable size. This disk-utilization percentage does not reflect disk usage that you might see inside the machine&#39;s guest operating system.This is required if refresh_os_disk_after_logoff is set to AT_SIZE. | [optional] 
**SessionTimeoutPolicy** | Pointer to **string** | Specifies the session timeout policy for the applications published from the Desktop pool. This policy indicates whether the launched application session is a forever application session or not. * DEFAULT: Indicates application sessions will be disconnected either on reaching the global idle timeout or on reaching the max session timeout. * NEVER: Indicates application sessions will not be disconnected either on reaching the global idle timeout or on reaching the max session timeout. | [optional] 

## Methods

### NewDesktopPoolSessionSettingsUpdateSpec

`func NewDesktopPoolSessionSettingsUpdateSpec(allowMultipleSessionsPerUser bool, disconnectedSessionTimeoutPolicy string, ) *DesktopPoolSessionSettingsUpdateSpec`

NewDesktopPoolSessionSettingsUpdateSpec instantiates a new DesktopPoolSessionSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolSessionSettingsUpdateSpecWithDefaults

`func NewDesktopPoolSessionSettingsUpdateSpecWithDefaults() *DesktopPoolSessionSettingsUpdateSpec`

NewDesktopPoolSessionSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolSessionSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleSessionsPerUser

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetAllowMultipleSessionsPerUser() bool`

GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field if non-nil, zero value otherwise.

### GetAllowMultipleSessionsPerUserOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetAllowMultipleSessionsPerUserOk() (*bool, bool)`

GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSessionsPerUser

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetAllowMultipleSessionsPerUser(v bool)`

SetAllowMultipleSessionsPerUser sets AllowMultipleSessionsPerUser field to given value.


### GetAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetAllowUsersToResetMachines() bool`

GetAllowUsersToResetMachines returns the AllowUsersToResetMachines field if non-nil, zero value otherwise.

### GetAllowUsersToResetMachinesOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetAllowUsersToResetMachinesOk() (*bool, bool)`

GetAllowUsersToResetMachinesOk returns a tuple with the AllowUsersToResetMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetAllowUsersToResetMachines(v bool)`

SetAllowUsersToResetMachines sets AllowUsersToResetMachines field to given value.

### HasAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasAllowUsersToResetMachines() bool`

HasAllowUsersToResetMachines returns a boolean if a field has been set.

### GetDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetDeleteOrRefreshMachineAfterLogoff() string`

GetDeleteOrRefreshMachineAfterLogoff returns the DeleteOrRefreshMachineAfterLogoff field if non-nil, zero value otherwise.

### GetDeleteOrRefreshMachineAfterLogoffOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetDeleteOrRefreshMachineAfterLogoffOk() (*string, bool)`

GetDeleteOrRefreshMachineAfterLogoffOk returns a tuple with the DeleteOrRefreshMachineAfterLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetDeleteOrRefreshMachineAfterLogoff(v string)`

SetDeleteOrRefreshMachineAfterLogoff sets DeleteOrRefreshMachineAfterLogoff field to given value.

### HasDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasDeleteOrRefreshMachineAfterLogoff() bool`

HasDeleteOrRefreshMachineAfterLogoff returns a boolean if a field has been set.

### GetDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetDisconnectedSessionTimeoutMinutes() int32`

GetDisconnectedSessionTimeoutMinutes returns the DisconnectedSessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetDisconnectedSessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetDisconnectedSessionTimeoutMinutesOk() (*int32, bool)`

GetDisconnectedSessionTimeoutMinutesOk returns a tuple with the DisconnectedSessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetDisconnectedSessionTimeoutMinutes(v int32)`

SetDisconnectedSessionTimeoutMinutes sets DisconnectedSessionTimeoutMinutes field to given value.

### HasDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasDisconnectedSessionTimeoutMinutes() bool`

HasDisconnectedSessionTimeoutMinutes returns a boolean if a field has been set.

### GetDisconnectedSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetDisconnectedSessionTimeoutPolicy() string`

GetDisconnectedSessionTimeoutPolicy returns the DisconnectedSessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetDisconnectedSessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetDisconnectedSessionTimeoutPolicyOk() (*string, bool)`

GetDisconnectedSessionTimeoutPolicyOk returns a tuple with the DisconnectedSessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetDisconnectedSessionTimeoutPolicy(v string)`

SetDisconnectedSessionTimeoutPolicy sets DisconnectedSessionTimeoutPolicy field to given value.


### GetEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetEmptySessionTimeoutMinutes() int32`

GetEmptySessionTimeoutMinutes returns the EmptySessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetEmptySessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetEmptySessionTimeoutMinutesOk() (*int32, bool)`

GetEmptySessionTimeoutMinutesOk returns a tuple with the EmptySessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetEmptySessionTimeoutMinutes(v int32)`

SetEmptySessionTimeoutMinutes sets EmptySessionTimeoutMinutes field to given value.

### HasEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasEmptySessionTimeoutMinutes() bool`

HasEmptySessionTimeoutMinutes returns a boolean if a field has been set.

### GetEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetEmptySessionTimeoutPolicy() string`

GetEmptySessionTimeoutPolicy returns the EmptySessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetEmptySessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetEmptySessionTimeoutPolicyOk() (*string, bool)`

GetEmptySessionTimeoutPolicyOk returns a tuple with the EmptySessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetEmptySessionTimeoutPolicy(v string)`

SetEmptySessionTimeoutPolicy sets EmptySessionTimeoutPolicy field to given value.

### HasEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasEmptySessionTimeoutPolicy() bool`

HasEmptySessionTimeoutPolicy returns a boolean if a field has been set.

### GetLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetLogoffAfterTimeout() bool`

GetLogoffAfterTimeout returns the LogoffAfterTimeout field if non-nil, zero value otherwise.

### GetLogoffAfterTimeoutOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetLogoffAfterTimeoutOk() (*bool, bool)`

GetLogoffAfterTimeoutOk returns a tuple with the LogoffAfterTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetLogoffAfterTimeout(v bool)`

SetLogoffAfterTimeout sets LogoffAfterTimeout field to given value.

### HasLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasLogoffAfterTimeout() bool`

HasLogoffAfterTimeout returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetPreLaunchSessionTimeoutMinutes() int32`

GetPreLaunchSessionTimeoutMinutes returns the PreLaunchSessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetPreLaunchSessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetPreLaunchSessionTimeoutMinutesOk() (*int32, bool)`

GetPreLaunchSessionTimeoutMinutesOk returns a tuple with the PreLaunchSessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetPreLaunchSessionTimeoutMinutes(v int32)`

SetPreLaunchSessionTimeoutMinutes sets PreLaunchSessionTimeoutMinutes field to given value.

### HasPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasPreLaunchSessionTimeoutMinutes() bool`

HasPreLaunchSessionTimeoutMinutes returns a boolean if a field has been set.

### GetPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetPreLaunchSessionTimeoutPolicy() string`

GetPreLaunchSessionTimeoutPolicy returns the PreLaunchSessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetPreLaunchSessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetPreLaunchSessionTimeoutPolicyOk() (*string, bool)`

GetPreLaunchSessionTimeoutPolicyOk returns a tuple with the PreLaunchSessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetPreLaunchSessionTimeoutPolicy(v string)`

SetPreLaunchSessionTimeoutPolicy sets PreLaunchSessionTimeoutPolicy field to given value.

### HasPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasPreLaunchSessionTimeoutPolicy() bool`

HasPreLaunchSessionTimeoutPolicy returns a boolean if a field has been set.

### GetRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetRefreshOsDiskAfterLogoff() string`

GetRefreshOsDiskAfterLogoff returns the RefreshOsDiskAfterLogoff field if non-nil, zero value otherwise.

### GetRefreshOsDiskAfterLogoffOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetRefreshOsDiskAfterLogoffOk() (*string, bool)`

GetRefreshOsDiskAfterLogoffOk returns a tuple with the RefreshOsDiskAfterLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetRefreshOsDiskAfterLogoff(v string)`

SetRefreshOsDiskAfterLogoff sets RefreshOsDiskAfterLogoff field to given value.

### HasRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasRefreshOsDiskAfterLogoff() bool`

HasRefreshOsDiskAfterLogoff returns a boolean if a field has been set.

### GetRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetRefreshPeriodDaysForReplicaOsDisk() int32`

GetRefreshPeriodDaysForReplicaOsDisk returns the RefreshPeriodDaysForReplicaOsDisk field if non-nil, zero value otherwise.

### GetRefreshPeriodDaysForReplicaOsDiskOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetRefreshPeriodDaysForReplicaOsDiskOk() (*int32, bool)`

GetRefreshPeriodDaysForReplicaOsDiskOk returns a tuple with the RefreshPeriodDaysForReplicaOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetRefreshPeriodDaysForReplicaOsDisk(v int32)`

SetRefreshPeriodDaysForReplicaOsDisk sets RefreshPeriodDaysForReplicaOsDisk field to given value.

### HasRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasRefreshPeriodDaysForReplicaOsDisk() bool`

HasRefreshPeriodDaysForReplicaOsDisk returns a boolean if a field has been set.

### GetRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetRefreshThresholdPercentageForReplicaOsDisk() int32`

GetRefreshThresholdPercentageForReplicaOsDisk returns the RefreshThresholdPercentageForReplicaOsDisk field if non-nil, zero value otherwise.

### GetRefreshThresholdPercentageForReplicaOsDiskOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetRefreshThresholdPercentageForReplicaOsDiskOk() (*int32, bool)`

GetRefreshThresholdPercentageForReplicaOsDiskOk returns a tuple with the RefreshThresholdPercentageForReplicaOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetRefreshThresholdPercentageForReplicaOsDisk(v int32)`

SetRefreshThresholdPercentageForReplicaOsDisk sets RefreshThresholdPercentageForReplicaOsDisk field to given value.

### HasRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasRefreshThresholdPercentageForReplicaOsDisk() bool`

HasRefreshThresholdPercentageForReplicaOsDisk returns a boolean if a field has been set.

### GetSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetSessionTimeoutPolicy() string`

GetSessionTimeoutPolicy returns the SessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetSessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsUpdateSpec) GetSessionTimeoutPolicyOk() (*string, bool)`

GetSessionTimeoutPolicyOk returns a tuple with the SessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) SetSessionTimeoutPolicy(v string)`

SetSessionTimeoutPolicy sets SessionTimeoutPolicy field to given value.

### HasSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsUpdateSpec) HasSessionTimeoutPolicy() bool`

HasSessionTimeoutPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


