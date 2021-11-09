# TaskResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** | The result message. | [optional] 
**MessageId** | Pointer to **string** | The result message ID. | [optional] 
**ResultCode** | Pointer to **string** | The result code of the task. * SUCCESS: Task is completed successfully. * WARN: Task is finished but has warning. * ERROR: Task is finished but the it has error. | [optional] 

## Methods

### NewTaskResult

`func NewTaskResult() *TaskResult`

NewTaskResult instantiates a new TaskResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskResultWithDefaults

`func NewTaskResultWithDefaults() *TaskResult`

NewTaskResultWithDefaults instantiates a new TaskResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *TaskResult) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *TaskResult) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *TaskResult) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *TaskResult) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageId

`func (o *TaskResult) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *TaskResult) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *TaskResult) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *TaskResult) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetResultCode

`func (o *TaskResult) GetResultCode() string`

GetResultCode returns the ResultCode field if non-nil, zero value otherwise.

### GetResultCodeOk

`func (o *TaskResult) GetResultCodeOk() (*string, bool)`

GetResultCodeOk returns a tuple with the ResultCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultCode

`func (o *TaskResult) SetResultCode(v string)`

SetResultCode sets ResultCode field to given value.

### HasResultCode

`func (o *TaskResult) HasResultCode() bool`

HasResultCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


