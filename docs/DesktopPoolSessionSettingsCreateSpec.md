# DesktopPoolSessionSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleSessionsPerUser** | Pointer to **bool** | Indicates whether multiple sessions are allowed per user in case of Floating User Assignment. Default value is false. | [optional] 
**AllowUsersToResetMachines** | Pointer to **bool** | Indicates whether the user can be allowed to reset or restart their machines. Default value is false. | [optional] 
**DeleteOrRefreshMachineAfterLogoff** | Pointer to **string** | Whether machines are to be deleted or refreshed after logoff in case of Floating User Assignment. This is applicable for automated desktops with virtual machines names based on pattern naming. This is not applicable for desktops that are using specified naming since dynamic creation and deletion of VMs is not supported. For Instant clone desktops this setting can only be set to DELETE. Default value is NEVER. * NEVER: Never delete or refresh the machine in the desktop pool. * DELETE: Delete the machine after user logoff. * REFRESH: Refresh the machine after user logoff. | [optional] 
**DisconnectedSessionTimeoutMinutes** | Pointer to **int32** | Disconnected sessions timeout (in minutes). This is required if disconnected_session_timeout_policy is set to AFTER.&lt;br&gt; | [optional] 
**DisconnectedSessionTimeoutPolicy** | Pointer to **string** | Log-off policy after disconnected session. Default value is NEVER. * IMMEDIATELY: Immediately Logoff after user disconnect. * AFTER: Logoff after the specified number of minutes after user disconnect. * NEVER: Do not logoff after user disconnect. | [optional] 
**EmptySessionTimeoutMinutes** | Pointer to **int32** | Application empty session timeout (in minutes). An empty session (that has no remote-able window) is disconnected after the timeout. This is required if the empty_session_timeout_policy set to AFTER. Default value is 1. | [optional] 
**EmptySessionTimeoutPolicy** | Pointer to **string** | Applicable only when session_type is APPLICATION or DESKTOP_AND_APPLICATION. &lt;br&gt;Application empty session timeout policy. Default value is AFTER. * IMMEDIATE: Empty session will be disconnected immediately. * NEVER: Empty session will never disconnected. * AFTER: Empty session will be disconnected after specified number of minutes. | [optional] 
**LogoffAfterTimeout** | Pointer to **bool** | Applicable only when session_type is APPLICATION or DESKTOP_AND_APPLICATION. &lt;br&gt;Indicates whether the empty application sessions are logged off (true) or disconnected (false) after timeout. Default value is false. | [optional] 
**PowerPolicy** | Pointer to **string** | Applicable to managed machines with default value as TAKE_NO_POWER_ACTION. Power policy for the machines in the desktop pool after logoff. For Instant clone desktops this setting can only be set to ALWAYS_POWERED_ON. * TAKE_NO_POWER_ACTION: No action will be taken when user logs off. * ALWAYS_POWERED_ON: Ensure machines in the Desktop pool are always powered on. The connection server will monitor and power on machines as necessary. * SUSPEND: Suspend when a user logs off or when desktop pool is no longer keeping a machine as a spare. This does not affect spare and newly provisioned machines. * POWER_OFF: Power off when a user logs off or when desktop pool is no longer keeping a machine as a spare. This does not affect spare and newly provisioned machines. | [optional] 
**PreLaunchSessionTimeoutMinutes** | Pointer to **int32** | Application pre-launch session timeout (in minutes). A pre-launch session is disconnected after the timeout. This is required if pre_launch_session_timeout_policy is set to AFTER. Default value is 10. | [optional] 
**PreLaunchSessionTimeoutPolicy** | Pointer to **string** | Applicable only when session_type is APPLICATION or DESKTOP_AND_APPLICATION. &lt;br&gt;Application pre-launch session timeout policy. Default value is AFTER. * AFTER: Pre-launched session is disconnected after specified number of minutes. * NEVER: Pre-launched session is never disconnected. | [optional] 
**RefreshOsDiskAfterLogoff** | Pointer to **string** | Indicates whether and when to refresh the OS disks. Applicable To: dedicated instant-clone desktop pools with default value as NEVER. * NEVER: The OS disk is never refreshed. * ALWAYS: The OS disk is refreshed every time the user logs off. * EVERY: The OS disk is refreshed at regular intervals of a specified number of days. The number of days is counted from the last refresh, or from the initial provisioning if no refresh has occurred yet. For example, if the specified value is 3 days, and three days have passed since the last refresh, the machine is refreshed after the user logs off. * AT_SIZE: The OS disk is refreshed when its current size reaches a specified percentage of its maximum allowable size. The maximum size of a linked clone&#39;s OS disk is the size of the replica&#39;s OS disk. With this option, the size of the linked clone&#39;s OS disk in the datastore is compared to maximum allowable size. This disk-utilization percentage does not reflect disk usage that you might see inside the machine&#39;s guest operating system. | [optional] 
**RefreshPeriodDaysForReplicaOsDisk** | Pointer to **int32** | Regular interval at which to refresh the OS disk. This is required if refresh_os_disk_after_logoff is set to EVERY. | [optional] 
**RefreshThresholdPercentageForReplicaOsDisk** | Pointer to **int32** | With the AT_SIZE option for refresh_os_disk_after_logoff, the size of the instant clone&#39;s OS diskin the datastore is compared to its maximum allowable size. This disk-utilization percentage does not reflect disk usage that you might see inside the machine&#39;s guest operating system. This is required if refresh_os_disk_after_logoff set to AT_SIZE. | [optional] 
**SessionTimeoutPolicy** | Pointer to **string** | Session timeout policy. Applicable only when session_type is APPLICATION or DESKTOP_AND_APPLICATION with default value as DEFAULT. * DEFAULT: Indicates application sessions will be disconnected either on reaching the global idle timeout or on reaching the max session timeout. * NEVER: Indicates application sessions will not be disconnected either on reaching the global idle timeout or on reaching the max session timeout. | [optional] 

## Methods

### NewDesktopPoolSessionSettingsCreateSpec

`func NewDesktopPoolSessionSettingsCreateSpec() *DesktopPoolSessionSettingsCreateSpec`

NewDesktopPoolSessionSettingsCreateSpec instantiates a new DesktopPoolSessionSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolSessionSettingsCreateSpecWithDefaults

`func NewDesktopPoolSessionSettingsCreateSpecWithDefaults() *DesktopPoolSessionSettingsCreateSpec`

NewDesktopPoolSessionSettingsCreateSpecWithDefaults instantiates a new DesktopPoolSessionSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleSessionsPerUser

`func (o *DesktopPoolSessionSettingsCreateSpec) GetAllowMultipleSessionsPerUser() bool`

GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field if non-nil, zero value otherwise.

### GetAllowMultipleSessionsPerUserOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetAllowMultipleSessionsPerUserOk() (*bool, bool)`

GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSessionsPerUser

`func (o *DesktopPoolSessionSettingsCreateSpec) SetAllowMultipleSessionsPerUser(v bool)`

SetAllowMultipleSessionsPerUser sets AllowMultipleSessionsPerUser field to given value.

### HasAllowMultipleSessionsPerUser

`func (o *DesktopPoolSessionSettingsCreateSpec) HasAllowMultipleSessionsPerUser() bool`

HasAllowMultipleSessionsPerUser returns a boolean if a field has been set.

### GetAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsCreateSpec) GetAllowUsersToResetMachines() bool`

GetAllowUsersToResetMachines returns the AllowUsersToResetMachines field if non-nil, zero value otherwise.

### GetAllowUsersToResetMachinesOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetAllowUsersToResetMachinesOk() (*bool, bool)`

GetAllowUsersToResetMachinesOk returns a tuple with the AllowUsersToResetMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsCreateSpec) SetAllowUsersToResetMachines(v bool)`

SetAllowUsersToResetMachines sets AllowUsersToResetMachines field to given value.

### HasAllowUsersToResetMachines

`func (o *DesktopPoolSessionSettingsCreateSpec) HasAllowUsersToResetMachines() bool`

HasAllowUsersToResetMachines returns a boolean if a field has been set.

### GetDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsCreateSpec) GetDeleteOrRefreshMachineAfterLogoff() string`

GetDeleteOrRefreshMachineAfterLogoff returns the DeleteOrRefreshMachineAfterLogoff field if non-nil, zero value otherwise.

### GetDeleteOrRefreshMachineAfterLogoffOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetDeleteOrRefreshMachineAfterLogoffOk() (*string, bool)`

GetDeleteOrRefreshMachineAfterLogoffOk returns a tuple with the DeleteOrRefreshMachineAfterLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsCreateSpec) SetDeleteOrRefreshMachineAfterLogoff(v string)`

SetDeleteOrRefreshMachineAfterLogoff sets DeleteOrRefreshMachineAfterLogoff field to given value.

### HasDeleteOrRefreshMachineAfterLogoff

`func (o *DesktopPoolSessionSettingsCreateSpec) HasDeleteOrRefreshMachineAfterLogoff() bool`

HasDeleteOrRefreshMachineAfterLogoff returns a boolean if a field has been set.

### GetDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) GetDisconnectedSessionTimeoutMinutes() int32`

GetDisconnectedSessionTimeoutMinutes returns the DisconnectedSessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetDisconnectedSessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetDisconnectedSessionTimeoutMinutesOk() (*int32, bool)`

GetDisconnectedSessionTimeoutMinutesOk returns a tuple with the DisconnectedSessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) SetDisconnectedSessionTimeoutMinutes(v int32)`

SetDisconnectedSessionTimeoutMinutes sets DisconnectedSessionTimeoutMinutes field to given value.

### HasDisconnectedSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) HasDisconnectedSessionTimeoutMinutes() bool`

HasDisconnectedSessionTimeoutMinutes returns a boolean if a field has been set.

### GetDisconnectedSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) GetDisconnectedSessionTimeoutPolicy() string`

GetDisconnectedSessionTimeoutPolicy returns the DisconnectedSessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetDisconnectedSessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetDisconnectedSessionTimeoutPolicyOk() (*string, bool)`

GetDisconnectedSessionTimeoutPolicyOk returns a tuple with the DisconnectedSessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) SetDisconnectedSessionTimeoutPolicy(v string)`

SetDisconnectedSessionTimeoutPolicy sets DisconnectedSessionTimeoutPolicy field to given value.

### HasDisconnectedSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) HasDisconnectedSessionTimeoutPolicy() bool`

HasDisconnectedSessionTimeoutPolicy returns a boolean if a field has been set.

### GetEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) GetEmptySessionTimeoutMinutes() int32`

GetEmptySessionTimeoutMinutes returns the EmptySessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetEmptySessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetEmptySessionTimeoutMinutesOk() (*int32, bool)`

GetEmptySessionTimeoutMinutesOk returns a tuple with the EmptySessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) SetEmptySessionTimeoutMinutes(v int32)`

SetEmptySessionTimeoutMinutes sets EmptySessionTimeoutMinutes field to given value.

### HasEmptySessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) HasEmptySessionTimeoutMinutes() bool`

HasEmptySessionTimeoutMinutes returns a boolean if a field has been set.

### GetEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) GetEmptySessionTimeoutPolicy() string`

GetEmptySessionTimeoutPolicy returns the EmptySessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetEmptySessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetEmptySessionTimeoutPolicyOk() (*string, bool)`

GetEmptySessionTimeoutPolicyOk returns a tuple with the EmptySessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) SetEmptySessionTimeoutPolicy(v string)`

SetEmptySessionTimeoutPolicy sets EmptySessionTimeoutPolicy field to given value.

### HasEmptySessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) HasEmptySessionTimeoutPolicy() bool`

HasEmptySessionTimeoutPolicy returns a boolean if a field has been set.

### GetLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsCreateSpec) GetLogoffAfterTimeout() bool`

GetLogoffAfterTimeout returns the LogoffAfterTimeout field if non-nil, zero value otherwise.

### GetLogoffAfterTimeoutOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetLogoffAfterTimeoutOk() (*bool, bool)`

GetLogoffAfterTimeoutOk returns a tuple with the LogoffAfterTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsCreateSpec) SetLogoffAfterTimeout(v bool)`

SetLogoffAfterTimeout sets LogoffAfterTimeout field to given value.

### HasLogoffAfterTimeout

`func (o *DesktopPoolSessionSettingsCreateSpec) HasLogoffAfterTimeout() bool`

HasLogoffAfterTimeout returns a boolean if a field has been set.

### GetPowerPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) GetPowerPolicy() string`

GetPowerPolicy returns the PowerPolicy field if non-nil, zero value otherwise.

### GetPowerPolicyOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetPowerPolicyOk() (*string, bool)`

GetPowerPolicyOk returns a tuple with the PowerPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) SetPowerPolicy(v string)`

SetPowerPolicy sets PowerPolicy field to given value.

### HasPowerPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) HasPowerPolicy() bool`

HasPowerPolicy returns a boolean if a field has been set.

### GetPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) GetPreLaunchSessionTimeoutMinutes() int32`

GetPreLaunchSessionTimeoutMinutes returns the PreLaunchSessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetPreLaunchSessionTimeoutMinutesOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetPreLaunchSessionTimeoutMinutesOk() (*int32, bool)`

GetPreLaunchSessionTimeoutMinutesOk returns a tuple with the PreLaunchSessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) SetPreLaunchSessionTimeoutMinutes(v int32)`

SetPreLaunchSessionTimeoutMinutes sets PreLaunchSessionTimeoutMinutes field to given value.

### HasPreLaunchSessionTimeoutMinutes

`func (o *DesktopPoolSessionSettingsCreateSpec) HasPreLaunchSessionTimeoutMinutes() bool`

HasPreLaunchSessionTimeoutMinutes returns a boolean if a field has been set.

### GetPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) GetPreLaunchSessionTimeoutPolicy() string`

GetPreLaunchSessionTimeoutPolicy returns the PreLaunchSessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetPreLaunchSessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetPreLaunchSessionTimeoutPolicyOk() (*string, bool)`

GetPreLaunchSessionTimeoutPolicyOk returns a tuple with the PreLaunchSessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) SetPreLaunchSessionTimeoutPolicy(v string)`

SetPreLaunchSessionTimeoutPolicy sets PreLaunchSessionTimeoutPolicy field to given value.

### HasPreLaunchSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) HasPreLaunchSessionTimeoutPolicy() bool`

HasPreLaunchSessionTimeoutPolicy returns a boolean if a field has been set.

### GetRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsCreateSpec) GetRefreshOsDiskAfterLogoff() string`

GetRefreshOsDiskAfterLogoff returns the RefreshOsDiskAfterLogoff field if non-nil, zero value otherwise.

### GetRefreshOsDiskAfterLogoffOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetRefreshOsDiskAfterLogoffOk() (*string, bool)`

GetRefreshOsDiskAfterLogoffOk returns a tuple with the RefreshOsDiskAfterLogoff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsCreateSpec) SetRefreshOsDiskAfterLogoff(v string)`

SetRefreshOsDiskAfterLogoff sets RefreshOsDiskAfterLogoff field to given value.

### HasRefreshOsDiskAfterLogoff

`func (o *DesktopPoolSessionSettingsCreateSpec) HasRefreshOsDiskAfterLogoff() bool`

HasRefreshOsDiskAfterLogoff returns a boolean if a field has been set.

### GetRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetRefreshPeriodDaysForReplicaOsDisk() int32`

GetRefreshPeriodDaysForReplicaOsDisk returns the RefreshPeriodDaysForReplicaOsDisk field if non-nil, zero value otherwise.

### GetRefreshPeriodDaysForReplicaOsDiskOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetRefreshPeriodDaysForReplicaOsDiskOk() (*int32, bool)`

GetRefreshPeriodDaysForReplicaOsDiskOk returns a tuple with the RefreshPeriodDaysForReplicaOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsCreateSpec) SetRefreshPeriodDaysForReplicaOsDisk(v int32)`

SetRefreshPeriodDaysForReplicaOsDisk sets RefreshPeriodDaysForReplicaOsDisk field to given value.

### HasRefreshPeriodDaysForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsCreateSpec) HasRefreshPeriodDaysForReplicaOsDisk() bool`

HasRefreshPeriodDaysForReplicaOsDisk returns a boolean if a field has been set.

### GetRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetRefreshThresholdPercentageForReplicaOsDisk() int32`

GetRefreshThresholdPercentageForReplicaOsDisk returns the RefreshThresholdPercentageForReplicaOsDisk field if non-nil, zero value otherwise.

### GetRefreshThresholdPercentageForReplicaOsDiskOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetRefreshThresholdPercentageForReplicaOsDiskOk() (*int32, bool)`

GetRefreshThresholdPercentageForReplicaOsDiskOk returns a tuple with the RefreshThresholdPercentageForReplicaOsDisk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsCreateSpec) SetRefreshThresholdPercentageForReplicaOsDisk(v int32)`

SetRefreshThresholdPercentageForReplicaOsDisk sets RefreshThresholdPercentageForReplicaOsDisk field to given value.

### HasRefreshThresholdPercentageForReplicaOsDisk

`func (o *DesktopPoolSessionSettingsCreateSpec) HasRefreshThresholdPercentageForReplicaOsDisk() bool`

HasRefreshThresholdPercentageForReplicaOsDisk returns a boolean if a field has been set.

### GetSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) GetSessionTimeoutPolicy() string`

GetSessionTimeoutPolicy returns the SessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetSessionTimeoutPolicyOk

`func (o *DesktopPoolSessionSettingsCreateSpec) GetSessionTimeoutPolicyOk() (*string, bool)`

GetSessionTimeoutPolicyOk returns a tuple with the SessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) SetSessionTimeoutPolicy(v string)`

SetSessionTimeoutPolicy sets SessionTimeoutPolicy field to given value.

### HasSessionTimeoutPolicy

`func (o *DesktopPoolSessionSettingsCreateSpec) HasSessionTimeoutPolicy() bool`

HasSessionTimeoutPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


