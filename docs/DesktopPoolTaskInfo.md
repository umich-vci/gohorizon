# DesktopPoolTaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelledVmtaskCount** | Pointer to **int32** | Number of VMs whose task is in cancelled state. | [optional] 
**Description** | Pointer to **string** | Description of the desktop pool task. | [optional] 
**ErrorVmtaskCount** | Pointer to **int32** | Number of VMs whose task is in fault state. | [optional] 
**HaltedVmtaskCount** | Pointer to **int32** | Number of VMs whose task is in holding state. | [optional] 
**Id** | Pointer to **string** | Unique ID representing Desktop Pool Task. | [optional] 
**Operation** | Pointer to **string** | The current desktop pool operation. * PUSH_IMAGE: A push image operation. | [optional] 
**RemainingVmtaskCount** | Pointer to **int32** | Number of VMs whose task is scheduled or running. | [optional] 
**ScheduleTime** | Pointer to **int64** | Time at which desktop pool task is scheduled to start. | [optional] 

## Methods

### NewDesktopPoolTaskInfo

`func NewDesktopPoolTaskInfo() *DesktopPoolTaskInfo`

NewDesktopPoolTaskInfo instantiates a new DesktopPoolTaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolTaskInfoWithDefaults

`func NewDesktopPoolTaskInfoWithDefaults() *DesktopPoolTaskInfo`

NewDesktopPoolTaskInfoWithDefaults instantiates a new DesktopPoolTaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelledVmtaskCount

`func (o *DesktopPoolTaskInfo) GetCancelledVmtaskCount() int32`

GetCancelledVmtaskCount returns the CancelledVmtaskCount field if non-nil, zero value otherwise.

### GetCancelledVmtaskCountOk

`func (o *DesktopPoolTaskInfo) GetCancelledVmtaskCountOk() (*int32, bool)`

GetCancelledVmtaskCountOk returns a tuple with the CancelledVmtaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledVmtaskCount

`func (o *DesktopPoolTaskInfo) SetCancelledVmtaskCount(v int32)`

SetCancelledVmtaskCount sets CancelledVmtaskCount field to given value.

### HasCancelledVmtaskCount

`func (o *DesktopPoolTaskInfo) HasCancelledVmtaskCount() bool`

HasCancelledVmtaskCount returns a boolean if a field has been set.

### GetDescription

`func (o *DesktopPoolTaskInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DesktopPoolTaskInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DesktopPoolTaskInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DesktopPoolTaskInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetErrorVmtaskCount

`func (o *DesktopPoolTaskInfo) GetErrorVmtaskCount() int32`

GetErrorVmtaskCount returns the ErrorVmtaskCount field if non-nil, zero value otherwise.

### GetErrorVmtaskCountOk

`func (o *DesktopPoolTaskInfo) GetErrorVmtaskCountOk() (*int32, bool)`

GetErrorVmtaskCountOk returns a tuple with the ErrorVmtaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorVmtaskCount

`func (o *DesktopPoolTaskInfo) SetErrorVmtaskCount(v int32)`

SetErrorVmtaskCount sets ErrorVmtaskCount field to given value.

### HasErrorVmtaskCount

`func (o *DesktopPoolTaskInfo) HasErrorVmtaskCount() bool`

HasErrorVmtaskCount returns a boolean if a field has been set.

### GetHaltedVmtaskCount

`func (o *DesktopPoolTaskInfo) GetHaltedVmtaskCount() int32`

GetHaltedVmtaskCount returns the HaltedVmtaskCount field if non-nil, zero value otherwise.

### GetHaltedVmtaskCountOk

`func (o *DesktopPoolTaskInfo) GetHaltedVmtaskCountOk() (*int32, bool)`

GetHaltedVmtaskCountOk returns a tuple with the HaltedVmtaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaltedVmtaskCount

`func (o *DesktopPoolTaskInfo) SetHaltedVmtaskCount(v int32)`

SetHaltedVmtaskCount sets HaltedVmtaskCount field to given value.

### HasHaltedVmtaskCount

`func (o *DesktopPoolTaskInfo) HasHaltedVmtaskCount() bool`

HasHaltedVmtaskCount returns a boolean if a field has been set.

### GetId

`func (o *DesktopPoolTaskInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DesktopPoolTaskInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DesktopPoolTaskInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DesktopPoolTaskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOperation

`func (o *DesktopPoolTaskInfo) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *DesktopPoolTaskInfo) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *DesktopPoolTaskInfo) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *DesktopPoolTaskInfo) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetRemainingVmtaskCount

`func (o *DesktopPoolTaskInfo) GetRemainingVmtaskCount() int32`

GetRemainingVmtaskCount returns the RemainingVmtaskCount field if non-nil, zero value otherwise.

### GetRemainingVmtaskCountOk

`func (o *DesktopPoolTaskInfo) GetRemainingVmtaskCountOk() (*int32, bool)`

GetRemainingVmtaskCountOk returns a tuple with the RemainingVmtaskCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingVmtaskCount

`func (o *DesktopPoolTaskInfo) SetRemainingVmtaskCount(v int32)`

SetRemainingVmtaskCount sets RemainingVmtaskCount field to given value.

### HasRemainingVmtaskCount

`func (o *DesktopPoolTaskInfo) HasRemainingVmtaskCount() bool`

HasRemainingVmtaskCount returns a boolean if a field has been set.

### GetScheduleTime

`func (o *DesktopPoolTaskInfo) GetScheduleTime() int64`

GetScheduleTime returns the ScheduleTime field if non-nil, zero value otherwise.

### GetScheduleTimeOk

`func (o *DesktopPoolTaskInfo) GetScheduleTimeOk() (*int64, bool)`

GetScheduleTimeOk returns a tuple with the ScheduleTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleTime

`func (o *DesktopPoolTaskInfo) SetScheduleTime(v int64)`

SetScheduleTime sets ScheduleTime field to given value.

### HasScheduleTime

`func (o *DesktopPoolTaskInfo) HasScheduleTime() bool`

HasScheduleTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


