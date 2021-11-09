# DesktopPoolPushImageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddVirtualTpm** | Pointer to **bool** | Indicates whether to add Virtual TPM device. Default: false | [optional] 
**ImStreamId** | Pointer to **string** | New image management stream for the desktop pool.&lt;br&gt;Either parent VM and snapshot or image management stream and tag are to be specified. | [optional] 
**ImTagId** | Pointer to **string** | New image management tag for the desktop pool. This must be a tag of the image management stream. | [optional] 
**LogoffPolicy** | **string** | Determines when to perform the operation on machines which have an active session. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions. | 
**ParentVmId** | Pointer to **string** | New base image virtual machine for the desktop pool. This must be in the same datacenter as the base image of the desktop pool.&lt;br&gt;Either parent VM and snapshot or image management stream and tag are to be specified. | [optional] 
**SnapshotId** | Pointer to **string** | New base image snapshot for the desktop pool. This must be a snapshot of the parent VM. | [optional] 
**StartTime** | Pointer to **int64** | When to start the operation. If unset or the time is in the past, the operation will begin immediately. Measured as epoch time. | [optional] 
**StopOnFirstError** | Pointer to **bool** | Indicates that the operation should stop on first error. Default: true | [optional] 

## Methods

### NewDesktopPoolPushImageSpec

`func NewDesktopPoolPushImageSpec(logoffPolicy string, ) *DesktopPoolPushImageSpec`

NewDesktopPoolPushImageSpec instantiates a new DesktopPoolPushImageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolPushImageSpecWithDefaults

`func NewDesktopPoolPushImageSpecWithDefaults() *DesktopPoolPushImageSpec`

NewDesktopPoolPushImageSpecWithDefaults instantiates a new DesktopPoolPushImageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddVirtualTpm

`func (o *DesktopPoolPushImageSpec) GetAddVirtualTpm() bool`

GetAddVirtualTpm returns the AddVirtualTpm field if non-nil, zero value otherwise.

### GetAddVirtualTpmOk

`func (o *DesktopPoolPushImageSpec) GetAddVirtualTpmOk() (*bool, bool)`

GetAddVirtualTpmOk returns a tuple with the AddVirtualTpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVirtualTpm

`func (o *DesktopPoolPushImageSpec) SetAddVirtualTpm(v bool)`

SetAddVirtualTpm sets AddVirtualTpm field to given value.

### HasAddVirtualTpm

`func (o *DesktopPoolPushImageSpec) HasAddVirtualTpm() bool`

HasAddVirtualTpm returns a boolean if a field has been set.

### GetImStreamId

`func (o *DesktopPoolPushImageSpec) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *DesktopPoolPushImageSpec) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *DesktopPoolPushImageSpec) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.

### HasImStreamId

`func (o *DesktopPoolPushImageSpec) HasImStreamId() bool`

HasImStreamId returns a boolean if a field has been set.

### GetImTagId

`func (o *DesktopPoolPushImageSpec) GetImTagId() string`

GetImTagId returns the ImTagId field if non-nil, zero value otherwise.

### GetImTagIdOk

`func (o *DesktopPoolPushImageSpec) GetImTagIdOk() (*string, bool)`

GetImTagIdOk returns a tuple with the ImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImTagId

`func (o *DesktopPoolPushImageSpec) SetImTagId(v string)`

SetImTagId sets ImTagId field to given value.

### HasImTagId

`func (o *DesktopPoolPushImageSpec) HasImTagId() bool`

HasImTagId returns a boolean if a field has been set.

### GetLogoffPolicy

`func (o *DesktopPoolPushImageSpec) GetLogoffPolicy() string`

GetLogoffPolicy returns the LogoffPolicy field if non-nil, zero value otherwise.

### GetLogoffPolicyOk

`func (o *DesktopPoolPushImageSpec) GetLogoffPolicyOk() (*string, bool)`

GetLogoffPolicyOk returns a tuple with the LogoffPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPolicy

`func (o *DesktopPoolPushImageSpec) SetLogoffPolicy(v string)`

SetLogoffPolicy sets LogoffPolicy field to given value.


### GetParentVmId

`func (o *DesktopPoolPushImageSpec) GetParentVmId() string`

GetParentVmId returns the ParentVmId field if non-nil, zero value otherwise.

### GetParentVmIdOk

`func (o *DesktopPoolPushImageSpec) GetParentVmIdOk() (*string, bool)`

GetParentVmIdOk returns a tuple with the ParentVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmId

`func (o *DesktopPoolPushImageSpec) SetParentVmId(v string)`

SetParentVmId sets ParentVmId field to given value.

### HasParentVmId

`func (o *DesktopPoolPushImageSpec) HasParentVmId() bool`

HasParentVmId returns a boolean if a field has been set.

### GetSnapshotId

`func (o *DesktopPoolPushImageSpec) GetSnapshotId() string`

GetSnapshotId returns the SnapshotId field if non-nil, zero value otherwise.

### GetSnapshotIdOk

`func (o *DesktopPoolPushImageSpec) GetSnapshotIdOk() (*string, bool)`

GetSnapshotIdOk returns a tuple with the SnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotId

`func (o *DesktopPoolPushImageSpec) SetSnapshotId(v string)`

SetSnapshotId sets SnapshotId field to given value.

### HasSnapshotId

`func (o *DesktopPoolPushImageSpec) HasSnapshotId() bool`

HasSnapshotId returns a boolean if a field has been set.

### GetStartTime

`func (o *DesktopPoolPushImageSpec) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *DesktopPoolPushImageSpec) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *DesktopPoolPushImageSpec) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *DesktopPoolPushImageSpec) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStopOnFirstError

`func (o *DesktopPoolPushImageSpec) GetStopOnFirstError() bool`

GetStopOnFirstError returns the StopOnFirstError field if non-nil, zero value otherwise.

### GetStopOnFirstErrorOk

`func (o *DesktopPoolPushImageSpec) GetStopOnFirstErrorOk() (*bool, bool)`

GetStopOnFirstErrorOk returns a tuple with the StopOnFirstError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopOnFirstError

`func (o *DesktopPoolPushImageSpec) SetStopOnFirstError(v bool)`

SetStopOnFirstError sets StopOnFirstError field to given value.

### HasStopOnFirstError

`func (o *DesktopPoolPushImageSpec) HasStopOnFirstError() bool`

HasStopOnFirstError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


