# FarmScheduledMaintenanceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImmediateMaintenanceScheduled** | Pointer to **bool** | Indicates whether immediate maintenance is scheduled. | [optional] 
**LogoffPolicy** | Pointer to **string** | Determines when to perform the operation on RDS servers which have an active session. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions. | [optional] 
**NextScheduledTime** | Pointer to **int64** | Time when next scheduled maintenance would happen. | [optional] 
**RecurringMaintenanceSettings** | Pointer to [**FarmRecurringMaintenanceSettingsInfo**](FarmRecurringMaintenanceSettingsInfo.md) |  | [optional] 
**StopOnFirstError** | Pointer to **bool** | Indicates whether the operation should stop on first error. | [optional] 

## Methods

### NewFarmScheduledMaintenanceInfo

`func NewFarmScheduledMaintenanceInfo() *FarmScheduledMaintenanceInfo`

NewFarmScheduledMaintenanceInfo instantiates a new FarmScheduledMaintenanceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmScheduledMaintenanceInfoWithDefaults

`func NewFarmScheduledMaintenanceInfoWithDefaults() *FarmScheduledMaintenanceInfo`

NewFarmScheduledMaintenanceInfoWithDefaults instantiates a new FarmScheduledMaintenanceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImmediateMaintenanceScheduled

`func (o *FarmScheduledMaintenanceInfo) GetImmediateMaintenanceScheduled() bool`

GetImmediateMaintenanceScheduled returns the ImmediateMaintenanceScheduled field if non-nil, zero value otherwise.

### GetImmediateMaintenanceScheduledOk

`func (o *FarmScheduledMaintenanceInfo) GetImmediateMaintenanceScheduledOk() (*bool, bool)`

GetImmediateMaintenanceScheduledOk returns a tuple with the ImmediateMaintenanceScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmediateMaintenanceScheduled

`func (o *FarmScheduledMaintenanceInfo) SetImmediateMaintenanceScheduled(v bool)`

SetImmediateMaintenanceScheduled sets ImmediateMaintenanceScheduled field to given value.

### HasImmediateMaintenanceScheduled

`func (o *FarmScheduledMaintenanceInfo) HasImmediateMaintenanceScheduled() bool`

HasImmediateMaintenanceScheduled returns a boolean if a field has been set.

### GetLogoffPolicy

`func (o *FarmScheduledMaintenanceInfo) GetLogoffPolicy() string`

GetLogoffPolicy returns the LogoffPolicy field if non-nil, zero value otherwise.

### GetLogoffPolicyOk

`func (o *FarmScheduledMaintenanceInfo) GetLogoffPolicyOk() (*string, bool)`

GetLogoffPolicyOk returns a tuple with the LogoffPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPolicy

`func (o *FarmScheduledMaintenanceInfo) SetLogoffPolicy(v string)`

SetLogoffPolicy sets LogoffPolicy field to given value.

### HasLogoffPolicy

`func (o *FarmScheduledMaintenanceInfo) HasLogoffPolicy() bool`

HasLogoffPolicy returns a boolean if a field has been set.

### GetNextScheduledTime

`func (o *FarmScheduledMaintenanceInfo) GetNextScheduledTime() int64`

GetNextScheduledTime returns the NextScheduledTime field if non-nil, zero value otherwise.

### GetNextScheduledTimeOk

`func (o *FarmScheduledMaintenanceInfo) GetNextScheduledTimeOk() (*int64, bool)`

GetNextScheduledTimeOk returns a tuple with the NextScheduledTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextScheduledTime

`func (o *FarmScheduledMaintenanceInfo) SetNextScheduledTime(v int64)`

SetNextScheduledTime sets NextScheduledTime field to given value.

### HasNextScheduledTime

`func (o *FarmScheduledMaintenanceInfo) HasNextScheduledTime() bool`

HasNextScheduledTime returns a boolean if a field has been set.

### GetRecurringMaintenanceSettings

`func (o *FarmScheduledMaintenanceInfo) GetRecurringMaintenanceSettings() FarmRecurringMaintenanceSettingsInfo`

GetRecurringMaintenanceSettings returns the RecurringMaintenanceSettings field if non-nil, zero value otherwise.

### GetRecurringMaintenanceSettingsOk

`func (o *FarmScheduledMaintenanceInfo) GetRecurringMaintenanceSettingsOk() (*FarmRecurringMaintenanceSettingsInfo, bool)`

GetRecurringMaintenanceSettingsOk returns a tuple with the RecurringMaintenanceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecurringMaintenanceSettings

`func (o *FarmScheduledMaintenanceInfo) SetRecurringMaintenanceSettings(v FarmRecurringMaintenanceSettingsInfo)`

SetRecurringMaintenanceSettings sets RecurringMaintenanceSettings field to given value.

### HasRecurringMaintenanceSettings

`func (o *FarmScheduledMaintenanceInfo) HasRecurringMaintenanceSettings() bool`

HasRecurringMaintenanceSettings returns a boolean if a field has been set.

### GetStopOnFirstError

`func (o *FarmScheduledMaintenanceInfo) GetStopOnFirstError() bool`

GetStopOnFirstError returns the StopOnFirstError field if non-nil, zero value otherwise.

### GetStopOnFirstErrorOk

`func (o *FarmScheduledMaintenanceInfo) GetStopOnFirstErrorOk() (*bool, bool)`

GetStopOnFirstErrorOk returns a tuple with the StopOnFirstError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnFirstError

`func (o *FarmScheduledMaintenanceInfo) SetStopOnFirstError(v bool)`

SetStopOnFirstError sets StopOnFirstError field to given value.

### HasStopOnFirstError

`func (o *FarmScheduledMaintenanceInfo) HasStopOnFirstError() bool`

HasStopOnFirstError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


