# VMTemplateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatacenterId** | Pointer to **string** | Datacenter id for this VM template. | [optional] 
**DiskSizeInBytes** | Pointer to **int64** | Sum of capacities of all the virtual disks in the template, in bytes. | [optional] 
**Id** | Pointer to **string** | Unique ID representing a VM template. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this VM template from being used in full clone desktop pool creation. | [optional] 
**MemoryMb** | Pointer to **int32** | Memory size of the VM template, in MB | [optional] 
**Name** | Pointer to **string** | VM template name. | [optional] 
**OperatingSystem** | Pointer to **string** | Operating system. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS) | [optional] 
**OperatingSystemDisplayName** | Pointer to **string** | Operating system display name from Virtual Center. | [optional] 
**Path** | Pointer to **string** | VM template path. | [optional] 
**VcenterId** | Pointer to **string** | ID of the vCenter to which this VM template belongs to. | [optional] 
**VgpuType** | Pointer to **string** | NVIDIA GRID vGPU type configured on this VM template. | [optional] 

## Methods

### NewVMTemplateInfo

`func NewVMTemplateInfo() *VMTemplateInfo`

NewVMTemplateInfo instantiates a new VMTemplateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVMTemplateInfoWithDefaults

`func NewVMTemplateInfoWithDefaults() *VMTemplateInfo`

NewVMTemplateInfoWithDefaults instantiates a new VMTemplateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatacenterId

`func (o *VMTemplateInfo) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *VMTemplateInfo) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *VMTemplateInfo) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *VMTemplateInfo) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetDiskSizeInBytes

`func (o *VMTemplateInfo) GetDiskSizeInBytes() int64`

GetDiskSizeInBytes returns the DiskSizeInBytes field if non-nil, zero value otherwise.

### GetDiskSizeInBytesOk

`func (o *VMTemplateInfo) GetDiskSizeInBytesOk() (*int64, bool)`

GetDiskSizeInBytesOk returns a tuple with the DiskSizeInBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeInBytes

`func (o *VMTemplateInfo) SetDiskSizeInBytes(v int64)`

SetDiskSizeInBytes sets DiskSizeInBytes field to given value.

### HasDiskSizeInBytes

`func (o *VMTemplateInfo) HasDiskSizeInBytes() bool`

HasDiskSizeInBytes returns a boolean if a field has been set.

### GetId

`func (o *VMTemplateInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VMTemplateInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VMTemplateInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VMTemplateInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *VMTemplateInfo) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *VMTemplateInfo) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *VMTemplateInfo) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *VMTemplateInfo) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetMemoryMb

`func (o *VMTemplateInfo) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *VMTemplateInfo) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *VMTemplateInfo) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *VMTemplateInfo) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetName

`func (o *VMTemplateInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VMTemplateInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VMTemplateInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VMTemplateInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *VMTemplateInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *VMTemplateInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *VMTemplateInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *VMTemplateInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperatingSystemDisplayName

`func (o *VMTemplateInfo) GetOperatingSystemDisplayName() string`

GetOperatingSystemDisplayName returns the OperatingSystemDisplayName field if non-nil, zero value otherwise.

### GetOperatingSystemDisplayNameOk

`func (o *VMTemplateInfo) GetOperatingSystemDisplayNameOk() (*string, bool)`

GetOperatingSystemDisplayNameOk returns a tuple with the OperatingSystemDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemDisplayName

`func (o *VMTemplateInfo) SetOperatingSystemDisplayName(v string)`

SetOperatingSystemDisplayName sets OperatingSystemDisplayName field to given value.

### HasOperatingSystemDisplayName

`func (o *VMTemplateInfo) HasOperatingSystemDisplayName() bool`

HasOperatingSystemDisplayName returns a boolean if a field has been set.

### GetPath

`func (o *VMTemplateInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *VMTemplateInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *VMTemplateInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *VMTemplateInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVcenterId

`func (o *VMTemplateInfo) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *VMTemplateInfo) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *VMTemplateInfo) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *VMTemplateInfo) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.

### GetVgpuType

`func (o *VMTemplateInfo) GetVgpuType() string`

GetVgpuType returns the VgpuType field if non-nil, zero value otherwise.

### GetVgpuTypeOk

`func (o *VMTemplateInfo) GetVgpuTypeOk() (*string, bool)`

GetVgpuTypeOk returns a tuple with the VgpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuType

`func (o *VMTemplateInfo) SetVgpuType(v string)`

SetVgpuType sets VgpuType field to given value.

### HasVgpuType

`func (o *VMTemplateInfo) HasVgpuType() bool`

HasVgpuType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


