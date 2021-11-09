# TaskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the task. * POD_FEDERATION_CATEGORY: Category for PodFederation related tasks. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the task. | [optional] 
**PercentageComplete** | Pointer to **int32** | The amount of task completed, in percentage. | [optional] 
**Result** | Pointer to [**TaskResult**](TaskResult.md) |  | [optional] 
**Status** | Pointer to **string** | The state of the task. * RUNNING: The task is currently running. * WAITING: The task is currently waiting to execute. * COMPLETED: The task execution has completed. * FAILED: The task execution has failed. * PAUSED: The task execution has been paused. * CANCELLED: The task execution has been cancelled. | [optional] 
**Type** | Pointer to **string** | The type of the task. * POD_FEDERATION_INITIALIZING: Task performing PodFederation initialize operation. * POD_FEDERATION_UNINITIALIZING: Task performing PodFederation uninitialize operation. * POD_FEDERATION_JOINING: Task performing PodFederation join operation. * POD_FEDERATION_UNJOINING: Task performing PodFederation unjoin operation. | [optional] 

## Methods

### NewTaskInfo

`func NewTaskInfo() *TaskInfo`

NewTaskInfo instantiates a new TaskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskInfoWithDefaults

`func NewTaskInfoWithDefaults() *TaskInfo`

NewTaskInfoWithDefaults instantiates a new TaskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TaskInfo) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TaskInfo) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TaskInfo) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TaskInfo) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetId

`func (o *TaskInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TaskInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TaskInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TaskInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPercentageComplete

`func (o *TaskInfo) GetPercentageComplete() int32`

GetPercentageComplete returns the PercentageComplete field if non-nil, zero value otherwise.

### GetPercentageCompleteOk

`func (o *TaskInfo) GetPercentageCompleteOk() (*int32, bool)`

GetPercentageCompleteOk returns a tuple with the PercentageComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageComplete

`func (o *TaskInfo) SetPercentageComplete(v int32)`

SetPercentageComplete sets PercentageComplete field to given value.

### HasPercentageComplete

`func (o *TaskInfo) HasPercentageComplete() bool`

HasPercentageComplete returns a boolean if a field has been set.

### GetResult

`func (o *TaskInfo) GetResult() TaskResult`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *TaskInfo) GetResultOk() (*TaskResult, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *TaskInfo) SetResult(v TaskResult)`

SetResult sets Result field to given value.

### HasResult

`func (o *TaskInfo) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *TaskInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *TaskInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TaskInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TaskInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TaskInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


