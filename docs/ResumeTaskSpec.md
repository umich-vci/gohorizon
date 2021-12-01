# ResumeTaskSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RetryFailedVms** | Pointer to **bool** | Indicates whether to restart the task for virtual machines whose task status is in error state. Default value is false. | [optional] 
**StopOnError** | Pointer to **bool** | Indicates whether to stop the task at first error. Default value is true. | [optional] 

## Methods

### NewResumeTaskSpec

`func NewResumeTaskSpec() *ResumeTaskSpec`

NewResumeTaskSpec instantiates a new ResumeTaskSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResumeTaskSpecWithDefaults

`func NewResumeTaskSpecWithDefaults() *ResumeTaskSpec`

NewResumeTaskSpecWithDefaults instantiates a new ResumeTaskSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetryFailedVms

`func (o *ResumeTaskSpec) GetRetryFailedVms() bool`

GetRetryFailedVms returns the RetryFailedVms field if non-nil, zero value otherwise.

### GetRetryFailedVmsOk

`func (o *ResumeTaskSpec) GetRetryFailedVmsOk() (*bool, bool)`

GetRetryFailedVmsOk returns a tuple with the RetryFailedVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryFailedVms

`func (o *ResumeTaskSpec) SetRetryFailedVms(v bool)`

SetRetryFailedVms sets RetryFailedVms field to given value.

### HasRetryFailedVms

`func (o *ResumeTaskSpec) HasRetryFailedVms() bool`

HasRetryFailedVms returns a boolean if a field has been set.

### GetStopOnError

`func (o *ResumeTaskSpec) GetStopOnError() bool`

GetStopOnError returns the StopOnError field if non-nil, zero value otherwise.

### GetStopOnErrorOk

`func (o *ResumeTaskSpec) GetStopOnErrorOk() (*bool, bool)`

GetStopOnErrorOk returns a tuple with the StopOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnError

`func (o *ResumeTaskSpec) SetStopOnError(v bool)`

SetStopOnError sets StopOnError field to given value.

### HasStopOnError

`func (o *ResumeTaskSpec) HasStopOnError() bool`

HasStopOnError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


