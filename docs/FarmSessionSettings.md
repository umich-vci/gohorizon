# FarmSessionSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisconnectedSessionTimeoutMinutes** | Pointer to **int32** | Disconnected sessions timeout (in minutes).Will be set when disconnected_session_timeout_policy is set to AFTER. | [optional] 
**DisconnectedSessionTimeoutPolicy** | Pointer to **string** | Log-off policy after disconnected session. Default value is NEVER. * IMMEDIATELY: Immediately Logoff after user disconnect. * AFTER: Logoff after the specified number of minutes after user disconnect. * NEVER: Do not logoff after user disconnect. | [optional] 
**EmptySessionTimeoutMinutes** | Pointer to **int32** | Application empty session timeout in minutes. An empty session that has no remote-ablewindow is disconnected after the timeout. Default value is 1.Will be set when the empty_session_timeout_policy set to AFTER. | [optional] 
**EmptySessionTimeoutPolicy** | Pointer to **string** | Application empty session timeout policy. Default value is AFTER. * IMMEDIATE: Empty session will be disconnected immediately. * NEVER: Empty session will never disconnected. * AFTER: Empty session will be disconnected after specified number of minutes. | [optional] 
**LogoffAfterTimeout** | Pointer to **bool** | After timeout, empty application sessions are logged off when set to true. Otherwise sessions are disconnected.Default value is false. | [optional] 
**PreLaunchSessionTimeoutMinutes** | Pointer to **int32** | Application pre-launch session timeout in minutes. A pre-launch session is disconnected after the timeout.Will be set only when pre_launch_timeout_policy is set to AFTER. | [optional] 
**PreLaunchSessionTimeoutPolicy** | Pointer to **string** | Pre-launch session timeout policy for the application sessions on this Farm. * AFTER: Pre-launched session is disconnected after specified number of minutes. * NEVER: Pre-launched session is never disconnected. | [optional] 

## Methods

### NewFarmSessionSettings

`func NewFarmSessionSettings() *FarmSessionSettings`

NewFarmSessionSettings instantiates a new FarmSessionSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmSessionSettingsWithDefaults

`func NewFarmSessionSettingsWithDefaults() *FarmSessionSettings`

NewFarmSessionSettingsWithDefaults instantiates a new FarmSessionSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisconnectedSessionTimeoutMinutes

`func (o *FarmSessionSettings) GetDisconnectedSessionTimeoutMinutes() int32`

GetDisconnectedSessionTimeoutMinutes returns the DisconnectedSessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetDisconnectedSessionTimeoutMinutesOk

`func (o *FarmSessionSettings) GetDisconnectedSessionTimeoutMinutesOk() (*int32, bool)`

GetDisconnectedSessionTimeoutMinutesOk returns a tuple with the DisconnectedSessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessionTimeoutMinutes

`func (o *FarmSessionSettings) SetDisconnectedSessionTimeoutMinutes(v int32)`

SetDisconnectedSessionTimeoutMinutes sets DisconnectedSessionTimeoutMinutes field to given value.

### HasDisconnectedSessionTimeoutMinutes

`func (o *FarmSessionSettings) HasDisconnectedSessionTimeoutMinutes() bool`

HasDisconnectedSessionTimeoutMinutes returns a boolean if a field has been set.

### GetDisconnectedSessionTimeoutPolicy

`func (o *FarmSessionSettings) GetDisconnectedSessionTimeoutPolicy() string`

GetDisconnectedSessionTimeoutPolicy returns the DisconnectedSessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetDisconnectedSessionTimeoutPolicyOk

`func (o *FarmSessionSettings) GetDisconnectedSessionTimeoutPolicyOk() (*string, bool)`

GetDisconnectedSessionTimeoutPolicyOk returns a tuple with the DisconnectedSessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedSessionTimeoutPolicy

`func (o *FarmSessionSettings) SetDisconnectedSessionTimeoutPolicy(v string)`

SetDisconnectedSessionTimeoutPolicy sets DisconnectedSessionTimeoutPolicy field to given value.

### HasDisconnectedSessionTimeoutPolicy

`func (o *FarmSessionSettings) HasDisconnectedSessionTimeoutPolicy() bool`

HasDisconnectedSessionTimeoutPolicy returns a boolean if a field has been set.

### GetEmptySessionTimeoutMinutes

`func (o *FarmSessionSettings) GetEmptySessionTimeoutMinutes() int32`

GetEmptySessionTimeoutMinutes returns the EmptySessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetEmptySessionTimeoutMinutesOk

`func (o *FarmSessionSettings) GetEmptySessionTimeoutMinutesOk() (*int32, bool)`

GetEmptySessionTimeoutMinutesOk returns a tuple with the EmptySessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySessionTimeoutMinutes

`func (o *FarmSessionSettings) SetEmptySessionTimeoutMinutes(v int32)`

SetEmptySessionTimeoutMinutes sets EmptySessionTimeoutMinutes field to given value.

### HasEmptySessionTimeoutMinutes

`func (o *FarmSessionSettings) HasEmptySessionTimeoutMinutes() bool`

HasEmptySessionTimeoutMinutes returns a boolean if a field has been set.

### GetEmptySessionTimeoutPolicy

`func (o *FarmSessionSettings) GetEmptySessionTimeoutPolicy() string`

GetEmptySessionTimeoutPolicy returns the EmptySessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetEmptySessionTimeoutPolicyOk

`func (o *FarmSessionSettings) GetEmptySessionTimeoutPolicyOk() (*string, bool)`

GetEmptySessionTimeoutPolicyOk returns a tuple with the EmptySessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptySessionTimeoutPolicy

`func (o *FarmSessionSettings) SetEmptySessionTimeoutPolicy(v string)`

SetEmptySessionTimeoutPolicy sets EmptySessionTimeoutPolicy field to given value.

### HasEmptySessionTimeoutPolicy

`func (o *FarmSessionSettings) HasEmptySessionTimeoutPolicy() bool`

HasEmptySessionTimeoutPolicy returns a boolean if a field has been set.

### GetLogoffAfterTimeout

`func (o *FarmSessionSettings) GetLogoffAfterTimeout() bool`

GetLogoffAfterTimeout returns the LogoffAfterTimeout field if non-nil, zero value otherwise.

### GetLogoffAfterTimeoutOk

`func (o *FarmSessionSettings) GetLogoffAfterTimeoutOk() (*bool, bool)`

GetLogoffAfterTimeoutOk returns a tuple with the LogoffAfterTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffAfterTimeout

`func (o *FarmSessionSettings) SetLogoffAfterTimeout(v bool)`

SetLogoffAfterTimeout sets LogoffAfterTimeout field to given value.

### HasLogoffAfterTimeout

`func (o *FarmSessionSettings) HasLogoffAfterTimeout() bool`

HasLogoffAfterTimeout returns a boolean if a field has been set.

### GetPreLaunchSessionTimeoutMinutes

`func (o *FarmSessionSettings) GetPreLaunchSessionTimeoutMinutes() int32`

GetPreLaunchSessionTimeoutMinutes returns the PreLaunchSessionTimeoutMinutes field if non-nil, zero value otherwise.

### GetPreLaunchSessionTimeoutMinutesOk

`func (o *FarmSessionSettings) GetPreLaunchSessionTimeoutMinutesOk() (*int32, bool)`

GetPreLaunchSessionTimeoutMinutesOk returns a tuple with the PreLaunchSessionTimeoutMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLaunchSessionTimeoutMinutes

`func (o *FarmSessionSettings) SetPreLaunchSessionTimeoutMinutes(v int32)`

SetPreLaunchSessionTimeoutMinutes sets PreLaunchSessionTimeoutMinutes field to given value.

### HasPreLaunchSessionTimeoutMinutes

`func (o *FarmSessionSettings) HasPreLaunchSessionTimeoutMinutes() bool`

HasPreLaunchSessionTimeoutMinutes returns a boolean if a field has been set.

### GetPreLaunchSessionTimeoutPolicy

`func (o *FarmSessionSettings) GetPreLaunchSessionTimeoutPolicy() string`

GetPreLaunchSessionTimeoutPolicy returns the PreLaunchSessionTimeoutPolicy field if non-nil, zero value otherwise.

### GetPreLaunchSessionTimeoutPolicyOk

`func (o *FarmSessionSettings) GetPreLaunchSessionTimeoutPolicyOk() (*string, bool)`

GetPreLaunchSessionTimeoutPolicyOk returns a tuple with the PreLaunchSessionTimeoutPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreLaunchSessionTimeoutPolicy

`func (o *FarmSessionSettings) SetPreLaunchSessionTimeoutPolicy(v string)`

SetPreLaunchSessionTimeoutPolicy sets PreLaunchSessionTimeoutPolicy field to given value.

### HasPreLaunchSessionTimeoutPolicy

`func (o *FarmSessionSettings) HasPreLaunchSessionTimeoutPolicy() bool`

HasPreLaunchSessionTimeoutPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


