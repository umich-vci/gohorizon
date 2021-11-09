# DesktopPoolInstantClonePushImageSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddVirtualTpm** | Pointer to **bool** | Whether to add Virtual TPM device. | [optional] 
**LogoffPolicy** | Pointer to **string** | Determines when to perform the operation on machines which have an active session. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions. | [optional] 
**StartTime** | Pointer to **int64** | When to start the operation. If unset or the time is in the past, the operation will begin immediately. Measured as epoch time. | [optional] 
**StopOnFirstError** | Pointer to **bool** | Indicates that the operation should stop on first error. | [optional] 

## Methods

### NewDesktopPoolInstantClonePushImageSettings

`func NewDesktopPoolInstantClonePushImageSettings() *DesktopPoolInstantClonePushImageSettings`

NewDesktopPoolInstantClonePushImageSettings instantiates a new DesktopPoolInstantClonePushImageSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolInstantClonePushImageSettingsWithDefaults

`func NewDesktopPoolInstantClonePushImageSettingsWithDefaults() *DesktopPoolInstantClonePushImageSettings`

NewDesktopPoolInstantClonePushImageSettingsWithDefaults instantiates a new DesktopPoolInstantClonePushImageSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddVirtualTpm

`func (o *DesktopPoolInstantClonePushImageSettings) GetAddVirtualTpm() bool`

GetAddVirtualTpm returns the AddVirtualTpm field if non-nil, zero value otherwise.

### GetAddVirtualTpmOk

`func (o *DesktopPoolInstantClonePushImageSettings) GetAddVirtualTpmOk() (*bool, bool)`

GetAddVirtualTpmOk returns a tuple with the AddVirtualTpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVirtualTpm

`func (o *DesktopPoolInstantClonePushImageSettings) SetAddVirtualTpm(v bool)`

SetAddVirtualTpm sets AddVirtualTpm field to given value.

### HasAddVirtualTpm

`func (o *DesktopPoolInstantClonePushImageSettings) HasAddVirtualTpm() bool`

HasAddVirtualTpm returns a boolean if a field has been set.

### GetLogoffPolicy

`func (o *DesktopPoolInstantClonePushImageSettings) GetLogoffPolicy() string`

GetLogoffPolicy returns the LogoffPolicy field if non-nil, zero value otherwise.

### GetLogoffPolicyOk

`func (o *DesktopPoolInstantClonePushImageSettings) GetLogoffPolicyOk() (*string, bool)`

GetLogoffPolicyOk returns a tuple with the LogoffPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPolicy

`func (o *DesktopPoolInstantClonePushImageSettings) SetLogoffPolicy(v string)`

SetLogoffPolicy sets LogoffPolicy field to given value.

### HasLogoffPolicy

`func (o *DesktopPoolInstantClonePushImageSettings) HasLogoffPolicy() bool`

HasLogoffPolicy returns a boolean if a field has been set.

### GetStartTime

`func (o *DesktopPoolInstantClonePushImageSettings) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DesktopPoolInstantClonePushImageSettings) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DesktopPoolInstantClonePushImageSettings) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DesktopPoolInstantClonePushImageSettings) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStopOnFirstError

`func (o *DesktopPoolInstantClonePushImageSettings) GetStopOnFirstError() bool`

GetStopOnFirstError returns the StopOnFirstError field if non-nil, zero value otherwise.

### GetStopOnFirstErrorOk

`func (o *DesktopPoolInstantClonePushImageSettings) GetStopOnFirstErrorOk() (*bool, bool)`

GetStopOnFirstErrorOk returns a tuple with the StopOnFirstError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnFirstError

`func (o *DesktopPoolInstantClonePushImageSettings) SetStopOnFirstError(v bool)`

SetStopOnFirstError sets StopOnFirstError field to given value.

### HasStopOnFirstError

`func (o *DesktopPoolInstantClonePushImageSettings) HasStopOnFirstError() bool`

HasStopOnFirstError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


