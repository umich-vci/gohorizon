# RDSHLoadBalancerSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuThreshold** | **int32** | Represents threshold of CPU usage, in percentage. If the value is 0, then this metric is not considered for load balancing. | 
**DiskQueueLengthThreshold** | **int32** | Represents the threshold of average number of both read and write requests that were queued for the selected disk during the sample interval. If the value is 0, then this metric is not considered for load balancing.  | 
**DiskReadLatencyThreshold** | **int32** | Represents the threshold of average time, in milliseconds, of a read of data from the disk. If the value is 0, then this metric is not considered for load balancing. | 
**DiskWriteLatencyThreshold** | **int32** | Represents the threshold of average time, in milliseconds, of a write of data to the disk. If the value is 0, then this metric is not considered for load balancing.  | 
**IncludeSessionCount** | **bool** | Indicates whether to include session count for load balancing. | 
**MemoryThreshold** | **int32** | Represents threshold of memory usage, in percentage. If the value is 0, then this metric is not considered for load balancing.  | 

## Methods

### NewRDSHLoadBalancerSettingsUpdateSpec

`func NewRDSHLoadBalancerSettingsUpdateSpec(cpuThreshold int32, diskQueueLengthThreshold int32, diskReadLatencyThreshold int32, diskWriteLatencyThreshold int32, includeSessionCount bool, memoryThreshold int32, ) *RDSHLoadBalancerSettingsUpdateSpec`

NewRDSHLoadBalancerSettingsUpdateSpec instantiates a new RDSHLoadBalancerSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSHLoadBalancerSettingsUpdateSpecWithDefaults

`func NewRDSHLoadBalancerSettingsUpdateSpecWithDefaults() *RDSHLoadBalancerSettingsUpdateSpec`

NewRDSHLoadBalancerSettingsUpdateSpecWithDefaults instantiates a new RDSHLoadBalancerSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetCpuThreshold() int32`

GetCpuThreshold returns the CpuThreshold field if non-nil, zero value otherwise.

### GetCpuThresholdOk

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetCpuThresholdOk() (*int32, bool)`

GetCpuThresholdOk returns a tuple with the CpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) SetCpuThreshold(v int32)`

SetCpuThreshold sets CpuThreshold field to given value.


### GetDiskQueueLengthThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetDiskQueueLengthThreshold() int32`

GetDiskQueueLengthThreshold returns the DiskQueueLengthThreshold field if non-nil, zero value otherwise.

### GetDiskQueueLengthThresholdOk

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetDiskQueueLengthThresholdOk() (*int32, bool)`

GetDiskQueueLengthThresholdOk returns a tuple with the DiskQueueLengthThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskQueueLengthThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) SetDiskQueueLengthThreshold(v int32)`

SetDiskQueueLengthThreshold sets DiskQueueLengthThreshold field to given value.


### GetDiskReadLatencyThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetDiskReadLatencyThreshold() int32`

GetDiskReadLatencyThreshold returns the DiskReadLatencyThreshold field if non-nil, zero value otherwise.

### GetDiskReadLatencyThresholdOk

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetDiskReadLatencyThresholdOk() (*int32, bool)`

GetDiskReadLatencyThresholdOk returns a tuple with the DiskReadLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadLatencyThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) SetDiskReadLatencyThreshold(v int32)`

SetDiskReadLatencyThreshold sets DiskReadLatencyThreshold field to given value.


### GetDiskWriteLatencyThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetDiskWriteLatencyThreshold() int32`

GetDiskWriteLatencyThreshold returns the DiskWriteLatencyThreshold field if non-nil, zero value otherwise.

### GetDiskWriteLatencyThresholdOk

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetDiskWriteLatencyThresholdOk() (*int32, bool)`

GetDiskWriteLatencyThresholdOk returns a tuple with the DiskWriteLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteLatencyThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) SetDiskWriteLatencyThreshold(v int32)`

SetDiskWriteLatencyThreshold sets DiskWriteLatencyThreshold field to given value.


### GetIncludeSessionCount

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetIncludeSessionCount() bool`

GetIncludeSessionCount returns the IncludeSessionCount field if non-nil, zero value otherwise.

### GetIncludeSessionCountOk

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetIncludeSessionCountOk() (*bool, bool)`

GetIncludeSessionCountOk returns a tuple with the IncludeSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSessionCount

`func (o *RDSHLoadBalancerSettingsUpdateSpec) SetIncludeSessionCount(v bool)`

SetIncludeSessionCount sets IncludeSessionCount field to given value.


### GetMemoryThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetMemoryThreshold() int32`

GetMemoryThreshold returns the MemoryThreshold field if non-nil, zero value otherwise.

### GetMemoryThresholdOk

`func (o *RDSHLoadBalancerSettingsUpdateSpec) GetMemoryThresholdOk() (*int32, bool)`

GetMemoryThresholdOk returns a tuple with the MemoryThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryThreshold

`func (o *RDSHLoadBalancerSettingsUpdateSpec) SetMemoryThreshold(v int32)`

SetMemoryThreshold sets MemoryThreshold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


