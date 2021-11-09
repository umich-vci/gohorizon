# VCMonitorHost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCoreCount** | Pointer to **int32** | Number of physical CPU cores on the host. | [optional] 
**CpuMhz** | Pointer to **int32** | CPU speed per core in Mhz. This might be an averaged value if the speed is not uniform across all cores. | [optional] 
**Details** | [**VCMonitorHostDetails**](VCMonitorHostDetails.md) |  | 
**MemorySizeMb** | Pointer to **int32** | The physical memory size in mega bytes. | [optional] 
**Status** | **string** | Status of the host connection. * CONNECTED: The host is successfully connected to Virtual Center server. * DISCONNECTED: The host is disconnected from Virtual Center server. * NOT_RESPONDING: The host is not responding. | 

## Methods

### NewVCMonitorHost

`func NewVCMonitorHost(details VCMonitorHostDetails, status string, ) *VCMonitorHost`

NewVCMonitorHost instantiates a new VCMonitorHost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCMonitorHostWithDefaults

`func NewVCMonitorHostWithDefaults() *VCMonitorHost`

NewVCMonitorHostWithDefaults instantiates a new VCMonitorHost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCoreCount

`func (o *VCMonitorHost) GetCpuCoreCount() int32`

GetCpuCoreCount returns the CpuCoreCount field if non-nil, zero value otherwise.

### GetCpuCoreCountOk

`func (o *VCMonitorHost) GetCpuCoreCountOk() (*int32, bool)`

GetCpuCoreCountOk returns a tuple with the CpuCoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCoreCount

`func (o *VCMonitorHost) SetCpuCoreCount(v int32)`

SetCpuCoreCount sets CpuCoreCount field to given value.

### HasCpuCoreCount

`func (o *VCMonitorHost) HasCpuCoreCount() bool`

HasCpuCoreCount returns a boolean if a field has been set.

### GetCpuMhz

`func (o *VCMonitorHost) GetCpuMhz() int32`

GetCpuMhz returns the CpuMhz field if non-nil, zero value otherwise.

### GetCpuMhzOk

`func (o *VCMonitorHost) GetCpuMhzOk() (*int32, bool)`

GetCpuMhzOk returns a tuple with the CpuMhz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMhz

`func (o *VCMonitorHost) SetCpuMhz(v int32)`

SetCpuMhz sets CpuMhz field to given value.

### HasCpuMhz

`func (o *VCMonitorHost) HasCpuMhz() bool`

HasCpuMhz returns a boolean if a field has been set.

### GetDetails

`func (o *VCMonitorHost) GetDetails() VCMonitorHostDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VCMonitorHost) GetDetailsOk() (*VCMonitorHostDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VCMonitorHost) SetDetails(v VCMonitorHostDetails)`

SetDetails sets Details field to given value.


### GetMemorySizeMb

`func (o *VCMonitorHost) GetMemorySizeMb() int32`

GetMemorySizeMb returns the MemorySizeMb field if non-nil, zero value otherwise.

### GetMemorySizeMbOk

`func (o *VCMonitorHost) GetMemorySizeMbOk() (*int32, bool)`

GetMemorySizeMbOk returns a tuple with the MemorySizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMb

`func (o *VCMonitorHost) SetMemorySizeMb(v int32)`

SetMemorySizeMb sets MemorySizeMb field to given value.

### HasMemorySizeMb

`func (o *VCMonitorHost) HasMemorySizeMb() bool`

HasMemorySizeMb returns a boolean if a field has been set.

### GetStatus

`func (o *VCMonitorHost) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VCMonitorHost) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VCMonitorHost) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


