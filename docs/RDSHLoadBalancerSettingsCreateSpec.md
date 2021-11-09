# RDSHLoadBalancerSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuThreshold** | Pointer to **int32** | Represents threshold of CPU usage, in percentage. If the value is 0, then this metric is not considered for load balancing. Default value is 0. | [optional] 
**DiskQueueLengthThreshold** | Pointer to **int32** | Represents the threshold of average number of both read and write requests that were queued for the selected disk during the sample interval. If the value is 0, then this metric is not considered for load balancing. Default value is 0. | [optional] 
**DiskReadLatencyThreshold** | Pointer to **int32** | Represents the threshold of average time, in milliseconds, of a read of data from the disk. If the value is 0, then this metric is not considered for load balancing. Default value is 0. | [optional] 
**DiskWriteLatencyThreshold** | Pointer to **int32** | Represents the threshold of average time, in milliseconds, of a write of data to the disk. If the value is 0, then this metric is not considered for load balancing. Default value is 0. | [optional] 
**IncludeSessionCount** | Pointer to **bool** | Indicates whether to include session count for load balancing. Default value is true. | [optional] 
**MemoryThreshold** | Pointer to **int32** | Represents threshold of memory usage, in percentage. If the value is 0, then this metric is not considered for load balancing. Default value is 0. | [optional] 

## Methods

### NewRDSHLoadBalancerSettingsCreateSpec

`func NewRDSHLoadBalancerSettingsCreateSpec() *RDSHLoadBalancerSettingsCreateSpec`

NewRDSHLoadBalancerSettingsCreateSpec instantiates a new RDSHLoadBalancerSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSHLoadBalancerSettingsCreateSpecWithDefaults

`func NewRDSHLoadBalancerSettingsCreateSpecWithDefaults() *RDSHLoadBalancerSettingsCreateSpec`

NewRDSHLoadBalancerSettingsCreateSpecWithDefaults instantiates a new RDSHLoadBalancerSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetCpuThreshold() int32`

GetCpuThreshold returns the CpuThreshold field if non-nil, zero value otherwise.

### GetCpuThresholdOk

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetCpuThresholdOk() (*int32, bool)`

GetCpuThresholdOk returns a tuple with the CpuThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) SetCpuThreshold(v int32)`

SetCpuThreshold sets CpuThreshold field to given value.

### HasCpuThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) HasCpuThreshold() bool`

HasCpuThreshold returns a boolean if a field has been set.

### GetDiskQueueLengthThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetDiskQueueLengthThreshold() int32`

GetDiskQueueLengthThreshold returns the DiskQueueLengthThreshold field if non-nil, zero value otherwise.

### GetDiskQueueLengthThresholdOk

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetDiskQueueLengthThresholdOk() (*int32, bool)`

GetDiskQueueLengthThresholdOk returns a tuple with the DiskQueueLengthThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskQueueLengthThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) SetDiskQueueLengthThreshold(v int32)`

SetDiskQueueLengthThreshold sets DiskQueueLengthThreshold field to given value.

### HasDiskQueueLengthThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) HasDiskQueueLengthThreshold() bool`

HasDiskQueueLengthThreshold returns a boolean if a field has been set.

### GetDiskReadLatencyThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetDiskReadLatencyThreshold() int32`

GetDiskReadLatencyThreshold returns the DiskReadLatencyThreshold field if non-nil, zero value otherwise.

### GetDiskReadLatencyThresholdOk

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetDiskReadLatencyThresholdOk() (*int32, bool)`

GetDiskReadLatencyThresholdOk returns a tuple with the DiskReadLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskReadLatencyThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) SetDiskReadLatencyThreshold(v int32)`

SetDiskReadLatencyThreshold sets DiskReadLatencyThreshold field to given value.

### HasDiskReadLatencyThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) HasDiskReadLatencyThreshold() bool`

HasDiskReadLatencyThreshold returns a boolean if a field has been set.

### GetDiskWriteLatencyThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetDiskWriteLatencyThreshold() int32`

GetDiskWriteLatencyThreshold returns the DiskWriteLatencyThreshold field if non-nil, zero value otherwise.

### GetDiskWriteLatencyThresholdOk

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetDiskWriteLatencyThresholdOk() (*int32, bool)`

GetDiskWriteLatencyThresholdOk returns a tuple with the DiskWriteLatencyThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskWriteLatencyThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) SetDiskWriteLatencyThreshold(v int32)`

SetDiskWriteLatencyThreshold sets DiskWriteLatencyThreshold field to given value.

### HasDiskWriteLatencyThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) HasDiskWriteLatencyThreshold() bool`

HasDiskWriteLatencyThreshold returns a boolean if a field has been set.

### GetIncludeSessionCount

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetIncludeSessionCount() bool`

GetIncludeSessionCount returns the IncludeSessionCount field if non-nil, zero value otherwise.

### GetIncludeSessionCountOk

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetIncludeSessionCountOk() (*bool, bool)`

GetIncludeSessionCountOk returns a tuple with the IncludeSessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSessionCount

`func (o *RDSHLoadBalancerSettingsCreateSpec) SetIncludeSessionCount(v bool)`

SetIncludeSessionCount sets IncludeSessionCount field to given value.

### HasIncludeSessionCount

`func (o *RDSHLoadBalancerSettingsCreateSpec) HasIncludeSessionCount() bool`

HasIncludeSessionCount returns a boolean if a field has been set.

### GetMemoryThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetMemoryThreshold() int32`

GetMemoryThreshold returns the MemoryThreshold field if non-nil, zero value otherwise.

### GetMemoryThresholdOk

`func (o *RDSHLoadBalancerSettingsCreateSpec) GetMemoryThresholdOk() (*int32, bool)`

GetMemoryThresholdOk returns a tuple with the MemoryThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) SetMemoryThreshold(v int32)`

SetMemoryThreshold sets MemoryThreshold field to given value.

### HasMemoryThreshold

`func (o *RDSHLoadBalancerSettingsCreateSpec) HasMemoryThreshold() bool`

HasMemoryThreshold returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


