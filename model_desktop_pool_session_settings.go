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

// DesktopPoolSessionSettings Machine settings applicable when a user logs off or when a desktop pool is no longer keeping a machine as a spare.
type DesktopPoolSessionSettings struct {
	// Indicates whether multiple sessions are allowed per user in case of Floating User Assignment. Default value is false.
	AllowMultipleSessionsPerUser *bool `json:"allow_multiple_sessions_per_user,omitempty"`
	// Indicates whether the user can be allowed to reset or restart their machines. Default value is false.
	AllowUsersToResetMachines *bool `json:"allow_users_to_reset_machines,omitempty"`
	// Whether machines are to be deleted or refreshed after logoff in case of Floating User Assignment.This is applicable for automated desktops with virtual machines names based onpattern naming. This is not applicable for desktops that are using specified naming since dynamic creation and deletion of VMs is not supported.For Instant clone desktops this setting can only be set to DELETE. Default value is NEVER. * NEVER: Never delete or refresh the machine in the desktop pool. * DELETE: Delete the machine after user logoff. * REFRESH: Refresh the machine after user logoff.
	DeleteOrRefreshMachineAfterLogoff *string `json:"delete_or_refresh_machine_after_logoff,omitempty"`
	// Disconnected sessions timeout (in minutes). Will be set when disconnected_session_timeout_policy is set to AFTER.
	DisconnectedSessionTimeoutMinutes *int32 `json:"disconnected_session_timeout_minutes,omitempty"`
	// Log-off policy after disconnected session. Default value is NEVER. * IMMEDIATELY: Immediately Logoff after user disconnect. * AFTER: Logoff after the specified number of minutes after user disconnect. * NEVER: Do not logoff after user disconnect.
	DisconnectedSessionTimeoutPolicy *string `json:"disconnected_session_timeout_policy,omitempty"`
	// Application empty session timeout (in minutes). An empty session (that has no remote-ablewindow) is disconnected after the timeout. Default value is 1.Will be set when the empty_session_timeout_policy set to AFTER.
	EmptySessionTimeoutMinutes *int32 `json:"empty_session_timeout_minutes,omitempty"`
	// Application empty session timeout policy. Default value is AFTER. * IMMEDIATE: Empty session will be disconnected immediately. * NEVER: Empty session will never disconnected. * AFTER: Empty session will be disconnected after specified number of minutes.
	EmptySessionTimeoutPolicy *string `json:"empty_session_timeout_policy,omitempty"`
	// Indicates whether the empty application sessions are logged off (true) or disconnected (false) after timeout.Default value is false.
	LogoffAfterTimeout *bool `json:"logoff_after_timeout,omitempty"`
	// Power policy for the machines in the desktop pool after logoff. This setting is only relevant for managed machines.Default value is TAKE_NO_POWER_ACTION.For Instant clone desktops this setting can only be set to ALWAYS_POWERED_ON. * TAKE_NO_POWER_ACTION: No action will be taken when user logs off. * ALWAYS_POWERED_ON: Ensure machines in the Desktop pool are always powered on. The connection server will monitor and power on machines as necessary. * SUSPEND: Suspend when a user logs off or when desktop pool is no longer keeping a machine as a spare. This does not affect spare and newly provisioned machines. * POWER_OFF: Power off when a user logs off or when desktop pool is no longer keeping a machine as a spare. This does not affect spare and newly provisioned machines.
	PowerPolicy *string `json:"power_policy,omitempty"`
	// Whether and when to refresh the OS disks for dedicated-assignment, linked-clone and instant-clone machines.Default value is NEVER. * NEVER: The OS disk is never refreshed. * ALWAYS: The OS disk is refreshed every time the user logs off. * EVERY: The OS disk is refreshed at regular intervals of a specified number of days. The number of days is counted from the last refresh, or from the initial provisioning if no refresh has occurred yet. For example, if the specified value is 3 days, and three days have passed since the last refresh, the machine is refreshed after the user logs off. * AT_SIZE: The OS disk is refreshed when its current size reaches a specified percentage of its maximum allowable size. The maximum size of a linked clone's OS disk is the size of the replica's OS disk. With this option, the size of the linked clone's OS disk in the datastore is compared to maximum allowable size. This disk-utilization percentage does not reflect disk usage that you might see inside the machine's guest operating system.
	RefreshOsDiskAfterLogoff *string `json:"refresh_os_disk_after_logoff,omitempty"`
	// Regular interval at which to refresh the OS disk. Will be set when refresh_os_disk_after_logoff set to EVERY.
	RefreshPeriodDaysForReplicaOsDisk *int32 `json:"refresh_period_days_for_replica_os_disk,omitempty"`
	// With the 'AT_SIZE' option for refreshOsDiskAfterLogoff, the size of the linked clone's OS diskin the datastore is compared to its maximum allowable size. This disk-utilization percentage does not reflect disk usage that you might see inside the machine's guest operating system.Will be set when refresh_os_disk_after_logoff set to AT_SIZE.
	RefreshThresholdPercentageForReplicaOsDisk *int32 `json:"refresh_threshold_percentage_for_replica_os_disk,omitempty"`
}

// NewDesktopPoolSessionSettings instantiates a new DesktopPoolSessionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolSessionSettings() *DesktopPoolSessionSettings {
	this := DesktopPoolSessionSettings{}
	return &this
}

// NewDesktopPoolSessionSettingsWithDefaults instantiates a new DesktopPoolSessionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolSessionSettingsWithDefaults() *DesktopPoolSessionSettings {
	this := DesktopPoolSessionSettings{}
	return &this
}

// GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetAllowMultipleSessionsPerUser() bool {
	if o == nil || o.AllowMultipleSessionsPerUser == nil {
		var ret bool
		return ret
	}
	return *o.AllowMultipleSessionsPerUser
}

// GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetAllowMultipleSessionsPerUserOk() (*bool, bool) {
	if o == nil || o.AllowMultipleSessionsPerUser == nil {
		return nil, false
	}
	return o.AllowMultipleSessionsPerUser, true
}

// HasAllowMultipleSessionsPerUser returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasAllowMultipleSessionsPerUser() bool {
	if o != nil && o.AllowMultipleSessionsPerUser != nil {
		return true
	}

	return false
}

// SetAllowMultipleSessionsPerUser gets a reference to the given bool and assigns it to the AllowMultipleSessionsPerUser field.
func (o *DesktopPoolSessionSettings) SetAllowMultipleSessionsPerUser(v bool) {
	o.AllowMultipleSessionsPerUser = &v
}

// GetAllowUsersToResetMachines returns the AllowUsersToResetMachines field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetAllowUsersToResetMachines() bool {
	if o == nil || o.AllowUsersToResetMachines == nil {
		var ret bool
		return ret
	}
	return *o.AllowUsersToResetMachines
}

// GetAllowUsersToResetMachinesOk returns a tuple with the AllowUsersToResetMachines field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetAllowUsersToResetMachinesOk() (*bool, bool) {
	if o == nil || o.AllowUsersToResetMachines == nil {
		return nil, false
	}
	return o.AllowUsersToResetMachines, true
}

// HasAllowUsersToResetMachines returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasAllowUsersToResetMachines() bool {
	if o != nil && o.AllowUsersToResetMachines != nil {
		return true
	}

	return false
}

// SetAllowUsersToResetMachines gets a reference to the given bool and assigns it to the AllowUsersToResetMachines field.
func (o *DesktopPoolSessionSettings) SetAllowUsersToResetMachines(v bool) {
	o.AllowUsersToResetMachines = &v
}

// GetDeleteOrRefreshMachineAfterLogoff returns the DeleteOrRefreshMachineAfterLogoff field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetDeleteOrRefreshMachineAfterLogoff() string {
	if o == nil || o.DeleteOrRefreshMachineAfterLogoff == nil {
		var ret string
		return ret
	}
	return *o.DeleteOrRefreshMachineAfterLogoff
}

// GetDeleteOrRefreshMachineAfterLogoffOk returns a tuple with the DeleteOrRefreshMachineAfterLogoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetDeleteOrRefreshMachineAfterLogoffOk() (*string, bool) {
	if o == nil || o.DeleteOrRefreshMachineAfterLogoff == nil {
		return nil, false
	}
	return o.DeleteOrRefreshMachineAfterLogoff, true
}

// HasDeleteOrRefreshMachineAfterLogoff returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasDeleteOrRefreshMachineAfterLogoff() bool {
	if o != nil && o.DeleteOrRefreshMachineAfterLogoff != nil {
		return true
	}

	return false
}

// SetDeleteOrRefreshMachineAfterLogoff gets a reference to the given string and assigns it to the DeleteOrRefreshMachineAfterLogoff field.
func (o *DesktopPoolSessionSettings) SetDeleteOrRefreshMachineAfterLogoff(v string) {
	o.DeleteOrRefreshMachineAfterLogoff = &v
}

// GetDisconnectedSessionTimeoutMinutes returns the DisconnectedSessionTimeoutMinutes field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetDisconnectedSessionTimeoutMinutes() int32 {
	if o == nil || o.DisconnectedSessionTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.DisconnectedSessionTimeoutMinutes
}

// GetDisconnectedSessionTimeoutMinutesOk returns a tuple with the DisconnectedSessionTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetDisconnectedSessionTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.DisconnectedSessionTimeoutMinutes == nil {
		return nil, false
	}
	return o.DisconnectedSessionTimeoutMinutes, true
}

// HasDisconnectedSessionTimeoutMinutes returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasDisconnectedSessionTimeoutMinutes() bool {
	if o != nil && o.DisconnectedSessionTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetDisconnectedSessionTimeoutMinutes gets a reference to the given int32 and assigns it to the DisconnectedSessionTimeoutMinutes field.
func (o *DesktopPoolSessionSettings) SetDisconnectedSessionTimeoutMinutes(v int32) {
	o.DisconnectedSessionTimeoutMinutes = &v
}

// GetDisconnectedSessionTimeoutPolicy returns the DisconnectedSessionTimeoutPolicy field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetDisconnectedSessionTimeoutPolicy() string {
	if o == nil || o.DisconnectedSessionTimeoutPolicy == nil {
		var ret string
		return ret
	}
	return *o.DisconnectedSessionTimeoutPolicy
}

// GetDisconnectedSessionTimeoutPolicyOk returns a tuple with the DisconnectedSessionTimeoutPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetDisconnectedSessionTimeoutPolicyOk() (*string, bool) {
	if o == nil || o.DisconnectedSessionTimeoutPolicy == nil {
		return nil, false
	}
	return o.DisconnectedSessionTimeoutPolicy, true
}

// HasDisconnectedSessionTimeoutPolicy returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasDisconnectedSessionTimeoutPolicy() bool {
	if o != nil && o.DisconnectedSessionTimeoutPolicy != nil {
		return true
	}

	return false
}

// SetDisconnectedSessionTimeoutPolicy gets a reference to the given string and assigns it to the DisconnectedSessionTimeoutPolicy field.
func (o *DesktopPoolSessionSettings) SetDisconnectedSessionTimeoutPolicy(v string) {
	o.DisconnectedSessionTimeoutPolicy = &v
}

// GetEmptySessionTimeoutMinutes returns the EmptySessionTimeoutMinutes field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetEmptySessionTimeoutMinutes() int32 {
	if o == nil || o.EmptySessionTimeoutMinutes == nil {
		var ret int32
		return ret
	}
	return *o.EmptySessionTimeoutMinutes
}

// GetEmptySessionTimeoutMinutesOk returns a tuple with the EmptySessionTimeoutMinutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetEmptySessionTimeoutMinutesOk() (*int32, bool) {
	if o == nil || o.EmptySessionTimeoutMinutes == nil {
		return nil, false
	}
	return o.EmptySessionTimeoutMinutes, true
}

// HasEmptySessionTimeoutMinutes returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasEmptySessionTimeoutMinutes() bool {
	if o != nil && o.EmptySessionTimeoutMinutes != nil {
		return true
	}

	return false
}

// SetEmptySessionTimeoutMinutes gets a reference to the given int32 and assigns it to the EmptySessionTimeoutMinutes field.
func (o *DesktopPoolSessionSettings) SetEmptySessionTimeoutMinutes(v int32) {
	o.EmptySessionTimeoutMinutes = &v
}

// GetEmptySessionTimeoutPolicy returns the EmptySessionTimeoutPolicy field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetEmptySessionTimeoutPolicy() string {
	if o == nil || o.EmptySessionTimeoutPolicy == nil {
		var ret string
		return ret
	}
	return *o.EmptySessionTimeoutPolicy
}

// GetEmptySessionTimeoutPolicyOk returns a tuple with the EmptySessionTimeoutPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetEmptySessionTimeoutPolicyOk() (*string, bool) {
	if o == nil || o.EmptySessionTimeoutPolicy == nil {
		return nil, false
	}
	return o.EmptySessionTimeoutPolicy, true
}

// HasEmptySessionTimeoutPolicy returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasEmptySessionTimeoutPolicy() bool {
	if o != nil && o.EmptySessionTimeoutPolicy != nil {
		return true
	}

	return false
}

// SetEmptySessionTimeoutPolicy gets a reference to the given string and assigns it to the EmptySessionTimeoutPolicy field.
func (o *DesktopPoolSessionSettings) SetEmptySessionTimeoutPolicy(v string) {
	o.EmptySessionTimeoutPolicy = &v
}

// GetLogoffAfterTimeout returns the LogoffAfterTimeout field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetLogoffAfterTimeout() bool {
	if o == nil || o.LogoffAfterTimeout == nil {
		var ret bool
		return ret
	}
	return *o.LogoffAfterTimeout
}

// GetLogoffAfterTimeoutOk returns a tuple with the LogoffAfterTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetLogoffAfterTimeoutOk() (*bool, bool) {
	if o == nil || o.LogoffAfterTimeout == nil {
		return nil, false
	}
	return o.LogoffAfterTimeout, true
}

// HasLogoffAfterTimeout returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasLogoffAfterTimeout() bool {
	if o != nil && o.LogoffAfterTimeout != nil {
		return true
	}

	return false
}

// SetLogoffAfterTimeout gets a reference to the given bool and assigns it to the LogoffAfterTimeout field.
func (o *DesktopPoolSessionSettings) SetLogoffAfterTimeout(v bool) {
	o.LogoffAfterTimeout = &v
}

// GetPowerPolicy returns the PowerPolicy field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetPowerPolicy() string {
	if o == nil || o.PowerPolicy == nil {
		var ret string
		return ret
	}
	return *o.PowerPolicy
}

// GetPowerPolicyOk returns a tuple with the PowerPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetPowerPolicyOk() (*string, bool) {
	if o == nil || o.PowerPolicy == nil {
		return nil, false
	}
	return o.PowerPolicy, true
}

// HasPowerPolicy returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasPowerPolicy() bool {
	if o != nil && o.PowerPolicy != nil {
		return true
	}

	return false
}

// SetPowerPolicy gets a reference to the given string and assigns it to the PowerPolicy field.
func (o *DesktopPoolSessionSettings) SetPowerPolicy(v string) {
	o.PowerPolicy = &v
}

// GetRefreshOsDiskAfterLogoff returns the RefreshOsDiskAfterLogoff field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetRefreshOsDiskAfterLogoff() string {
	if o == nil || o.RefreshOsDiskAfterLogoff == nil {
		var ret string
		return ret
	}
	return *o.RefreshOsDiskAfterLogoff
}

// GetRefreshOsDiskAfterLogoffOk returns a tuple with the RefreshOsDiskAfterLogoff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetRefreshOsDiskAfterLogoffOk() (*string, bool) {
	if o == nil || o.RefreshOsDiskAfterLogoff == nil {
		return nil, false
	}
	return o.RefreshOsDiskAfterLogoff, true
}

// HasRefreshOsDiskAfterLogoff returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasRefreshOsDiskAfterLogoff() bool {
	if o != nil && o.RefreshOsDiskAfterLogoff != nil {
		return true
	}

	return false
}

// SetRefreshOsDiskAfterLogoff gets a reference to the given string and assigns it to the RefreshOsDiskAfterLogoff field.
func (o *DesktopPoolSessionSettings) SetRefreshOsDiskAfterLogoff(v string) {
	o.RefreshOsDiskAfterLogoff = &v
}

// GetRefreshPeriodDaysForReplicaOsDisk returns the RefreshPeriodDaysForReplicaOsDisk field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetRefreshPeriodDaysForReplicaOsDisk() int32 {
	if o == nil || o.RefreshPeriodDaysForReplicaOsDisk == nil {
		var ret int32
		return ret
	}
	return *o.RefreshPeriodDaysForReplicaOsDisk
}

// GetRefreshPeriodDaysForReplicaOsDiskOk returns a tuple with the RefreshPeriodDaysForReplicaOsDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetRefreshPeriodDaysForReplicaOsDiskOk() (*int32, bool) {
	if o == nil || o.RefreshPeriodDaysForReplicaOsDisk == nil {
		return nil, false
	}
	return o.RefreshPeriodDaysForReplicaOsDisk, true
}

// HasRefreshPeriodDaysForReplicaOsDisk returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasRefreshPeriodDaysForReplicaOsDisk() bool {
	if o != nil && o.RefreshPeriodDaysForReplicaOsDisk != nil {
		return true
	}

	return false
}

// SetRefreshPeriodDaysForReplicaOsDisk gets a reference to the given int32 and assigns it to the RefreshPeriodDaysForReplicaOsDisk field.
func (o *DesktopPoolSessionSettings) SetRefreshPeriodDaysForReplicaOsDisk(v int32) {
	o.RefreshPeriodDaysForReplicaOsDisk = &v
}

// GetRefreshThresholdPercentageForReplicaOsDisk returns the RefreshThresholdPercentageForReplicaOsDisk field value if set, zero value otherwise.
func (o *DesktopPoolSessionSettings) GetRefreshThresholdPercentageForReplicaOsDisk() int32 {
	if o == nil || o.RefreshThresholdPercentageForReplicaOsDisk == nil {
		var ret int32
		return ret
	}
	return *o.RefreshThresholdPercentageForReplicaOsDisk
}

// GetRefreshThresholdPercentageForReplicaOsDiskOk returns a tuple with the RefreshThresholdPercentageForReplicaOsDisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolSessionSettings) GetRefreshThresholdPercentageForReplicaOsDiskOk() (*int32, bool) {
	if o == nil || o.RefreshThresholdPercentageForReplicaOsDisk == nil {
		return nil, false
	}
	return o.RefreshThresholdPercentageForReplicaOsDisk, true
}

// HasRefreshThresholdPercentageForReplicaOsDisk returns a boolean if a field has been set.
func (o *DesktopPoolSessionSettings) HasRefreshThresholdPercentageForReplicaOsDisk() bool {
	if o != nil && o.RefreshThresholdPercentageForReplicaOsDisk != nil {
		return true
	}

	return false
}

// SetRefreshThresholdPercentageForReplicaOsDisk gets a reference to the given int32 and assigns it to the RefreshThresholdPercentageForReplicaOsDisk field.
func (o *DesktopPoolSessionSettings) SetRefreshThresholdPercentageForReplicaOsDisk(v int32) {
	o.RefreshThresholdPercentageForReplicaOsDisk = &v
}

func (o DesktopPoolSessionSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowMultipleSessionsPerUser != nil {
		toSerialize["allow_multiple_sessions_per_user"] = o.AllowMultipleSessionsPerUser
	}
	if o.AllowUsersToResetMachines != nil {
		toSerialize["allow_users_to_reset_machines"] = o.AllowUsersToResetMachines
	}
	if o.DeleteOrRefreshMachineAfterLogoff != nil {
		toSerialize["delete_or_refresh_machine_after_logoff"] = o.DeleteOrRefreshMachineAfterLogoff
	}
	if o.DisconnectedSessionTimeoutMinutes != nil {
		toSerialize["disconnected_session_timeout_minutes"] = o.DisconnectedSessionTimeoutMinutes
	}
	if o.DisconnectedSessionTimeoutPolicy != nil {
		toSerialize["disconnected_session_timeout_policy"] = o.DisconnectedSessionTimeoutPolicy
	}
	if o.EmptySessionTimeoutMinutes != nil {
		toSerialize["empty_session_timeout_minutes"] = o.EmptySessionTimeoutMinutes
	}
	if o.EmptySessionTimeoutPolicy != nil {
		toSerialize["empty_session_timeout_policy"] = o.EmptySessionTimeoutPolicy
	}
	if o.LogoffAfterTimeout != nil {
		toSerialize["logoff_after_timeout"] = o.LogoffAfterTimeout
	}
	if o.PowerPolicy != nil {
		toSerialize["power_policy"] = o.PowerPolicy
	}
	if o.RefreshOsDiskAfterLogoff != nil {
		toSerialize["refresh_os_disk_after_logoff"] = o.RefreshOsDiskAfterLogoff
	}
	if o.RefreshPeriodDaysForReplicaOsDisk != nil {
		toSerialize["refresh_period_days_for_replica_os_disk"] = o.RefreshPeriodDaysForReplicaOsDisk
	}
	if o.RefreshThresholdPercentageForReplicaOsDisk != nil {
		toSerialize["refresh_threshold_percentage_for_replica_os_disk"] = o.RefreshThresholdPercentageForReplicaOsDisk
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolSessionSettings struct {
	value *DesktopPoolSessionSettings
	isSet bool
}

func (v NullableDesktopPoolSessionSettings) Get() *DesktopPoolSessionSettings {
	return v.value
}

func (v *NullableDesktopPoolSessionSettings) Set(val *DesktopPoolSessionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolSessionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolSessionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolSessionSettings(val *DesktopPoolSessionSettings) *NullableDesktopPoolSessionSettings {
	return &NullableDesktopPoolSessionSettings{value: val, isSet: true}
}

func (v NullableDesktopPoolSessionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolSessionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
