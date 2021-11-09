# BaseSnapshotInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedTimestamp** | Pointer to **int64** | Epoch time in milli seconds, when the VM snapshot was created. | [optional] 
**Description** | Pointer to **string** | Description of the VM snapshot. | [optional] 
**DiskSizeMb** | Pointer to **int64** | Sum of capacities of all the virtual disks in the VM snapshot, in MB. | [optional] 
**HardwareVersion** | Pointer to **int32** | VM snapshot hardware version | [optional] 
**Id** | Pointer to **string** | Unique ID representing the VM snapshot. | [optional] 
**IncompatibleReasons** | Pointer to **[]string** | Reasons that may preclude this VM snapshot from being used in linked/instant clone desktop pool or farm creation. | [optional] 
**MaxNumberOfMonitors** | Pointer to **int32** | Maximum number of monitors set in SVGA settings for the VM snapshot in vCenter. | [optional] 
**MaxResolutionOfAnyOneMonitor** | Pointer to **string** | Maximum resolution of any one monitor set in SVGA settings for the VM snapshot in vCenter. | [optional] 
**MemoryMb** | Pointer to **int32** | The physical memory size of VM snapshot, in MB | [optional] 
**MemoryReservationMb** | Pointer to **int64** | Amount of memory that is guaranteed available to the virtual machine, in MB. | [optional] 
**Name** | Pointer to **string** | VM snapshot name. | [optional] 
**Path** | Pointer to **string** | VM snapshot path. | [optional] 
**Renderer3d** | Pointer to **string** | Indicate how the virtual video device for the VM snapshot renders 3D graphics. Will be set only if VM snapshot supports 3D functions. * MANAGE_BY_VSPHERE_CLIENT: 3D rendering managed by vSphere Client. * AUTOMATIC: 3D rendering is automatic. * SOFTWARE: 3D rendering is software dependent. The software renderer is supported (at minimum) on virtual hardware version 8 in a vSphere 5.0 environment. * HARDWARE: 3D rendering is hardware dependent. The hardware-based renderer is supported (at minimum) on virtual hardware version 9 in a vSphere 5.1 environment. * DISABLED: 3D rendering is disabled. | [optional] 
**TotalVideoMemoryMb** | Pointer to **float64** | Total video memory in MB set in SVGA settings for the VM snapshot in vCenter. | [optional] 
**VcenterId** | Pointer to **string** | Virtual Center id for this VM snapshot. | [optional] 
**VgpuType** | Pointer to **string** | NVIDIA GRID vGPU type configured on this VM snapshot. | [optional] 

## Methods

### NewBaseSnapshotInfo

`func NewBaseSnapshotInfo() *BaseSnapshotInfo`

NewBaseSnapshotInfo instantiates a new BaseSnapshotInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseSnapshotInfoWithDefaults

`func NewBaseSnapshotInfoWithDefaults() *BaseSnapshotInfo`

NewBaseSnapshotInfoWithDefaults instantiates a new BaseSnapshotInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedTimestamp

`func (o *BaseSnapshotInfo) GetCreatedTimestamp() int64`

GetCreatedTimestamp returns the CreatedTimestamp field if non-nil, zero value otherwise.

### GetCreatedTimestampOk

`func (o *BaseSnapshotInfo) GetCreatedTimestampOk() (*int64, bool)`

GetCreatedTimestampOk returns a tuple with the CreatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTimestamp

`func (o *BaseSnapshotInfo) SetCreatedTimestamp(v int64)`

SetCreatedTimestamp sets CreatedTimestamp field to given value.

### HasCreatedTimestamp

`func (o *BaseSnapshotInfo) HasCreatedTimestamp() bool`

HasCreatedTimestamp returns a boolean if a field has been set.

### GetDescription

`func (o *BaseSnapshotInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseSnapshotInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseSnapshotInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseSnapshotInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDiskSizeMb

`func (o *BaseSnapshotInfo) GetDiskSizeMb() int64`

GetDiskSizeMb returns the DiskSizeMb field if non-nil, zero value otherwise.

### GetDiskSizeMbOk

`func (o *BaseSnapshotInfo) GetDiskSizeMbOk() (*int64, bool)`

GetDiskSizeMbOk returns a tuple with the DiskSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskSizeMb

`func (o *BaseSnapshotInfo) SetDiskSizeMb(v int64)`

SetDiskSizeMb sets DiskSizeMb field to given value.

### HasDiskSizeMb

`func (o *BaseSnapshotInfo) HasDiskSizeMb() bool`

HasDiskSizeMb returns a boolean if a field has been set.

### GetHardwareVersion

`func (o *BaseSnapshotInfo) GetHardwareVersion() int32`

GetHardwareVersion returns the HardwareVersion field if non-nil, zero value otherwise.

### GetHardwareVersionOk

`func (o *BaseSnapshotInfo) GetHardwareVersionOk() (*int32, bool)`

GetHardwareVersionOk returns a tuple with the HardwareVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareVersion

`func (o *BaseSnapshotInfo) SetHardwareVersion(v int32)`

SetHardwareVersion sets HardwareVersion field to given value.

### HasHardwareVersion

`func (o *BaseSnapshotInfo) HasHardwareVersion() bool`

HasHardwareVersion returns a boolean if a field has been set.

### GetId

`func (o *BaseSnapshotInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseSnapshotInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseSnapshotInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BaseSnapshotInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIncompatibleReasons

`func (o *BaseSnapshotInfo) GetIncompatibleReasons() []string`

GetIncompatibleReasons returns the IncompatibleReasons field if non-nil, zero value otherwise.

### GetIncompatibleReasonsOk

`func (o *BaseSnapshotInfo) GetIncompatibleReasonsOk() (*[]string, bool)`

GetIncompatibleReasonsOk returns a tuple with the IncompatibleReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncompatibleReasons

`func (o *BaseSnapshotInfo) SetIncompatibleReasons(v []string)`

SetIncompatibleReasons sets IncompatibleReasons field to given value.

### HasIncompatibleReasons

`func (o *BaseSnapshotInfo) HasIncompatibleReasons() bool`

HasIncompatibleReasons returns a boolean if a field has been set.

### GetMaxNumberOfMonitors

`func (o *BaseSnapshotInfo) GetMaxNumberOfMonitors() int32`

GetMaxNumberOfMonitors returns the MaxNumberOfMonitors field if non-nil, zero value otherwise.

### GetMaxNumberOfMonitorsOk

`func (o *BaseSnapshotInfo) GetMaxNumberOfMonitorsOk() (*int32, bool)`

GetMaxNumberOfMonitorsOk returns a tuple with the MaxNumberOfMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfMonitors

`func (o *BaseSnapshotInfo) SetMaxNumberOfMonitors(v int32)`

SetMaxNumberOfMonitors sets MaxNumberOfMonitors field to given value.

### HasMaxNumberOfMonitors

`func (o *BaseSnapshotInfo) HasMaxNumberOfMonitors() bool`

HasMaxNumberOfMonitors returns a boolean if a field has been set.

### GetMaxResolutionOfAnyOneMonitor

`func (o *BaseSnapshotInfo) GetMaxResolutionOfAnyOneMonitor() string`

GetMaxResolutionOfAnyOneMonitor returns the MaxResolutionOfAnyOneMonitor field if non-nil, zero value otherwise.

### GetMaxResolutionOfAnyOneMonitorOk

`func (o *BaseSnapshotInfo) GetMaxResolutionOfAnyOneMonitorOk() (*string, bool)`

GetMaxResolutionOfAnyOneMonitorOk returns a tuple with the MaxResolutionOfAnyOneMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResolutionOfAnyOneMonitor

`func (o *BaseSnapshotInfo) SetMaxResolutionOfAnyOneMonitor(v string)`

SetMaxResolutionOfAnyOneMonitor sets MaxResolutionOfAnyOneMonitor field to given value.

### HasMaxResolutionOfAnyOneMonitor

`func (o *BaseSnapshotInfo) HasMaxResolutionOfAnyOneMonitor() bool`

HasMaxResolutionOfAnyOneMonitor returns a boolean if a field has been set.

### GetMemoryMb

`func (o *BaseSnapshotInfo) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *BaseSnapshotInfo) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *BaseSnapshotInfo) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *BaseSnapshotInfo) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetMemoryReservationMb

`func (o *BaseSnapshotInfo) GetMemoryReservationMb() int64`

GetMemoryReservationMb returns the MemoryReservationMb field if non-nil, zero value otherwise.

### GetMemoryReservationMbOk

`func (o *BaseSnapshotInfo) GetMemoryReservationMbOk() (*int64, bool)`

GetMemoryReservationMbOk returns a tuple with the MemoryReservationMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryReservationMb

`func (o *BaseSnapshotInfo) SetMemoryReservationMb(v int64)`

SetMemoryReservationMb sets MemoryReservationMb field to given value.

### HasMemoryReservationMb

`func (o *BaseSnapshotInfo) HasMemoryReservationMb() bool`

HasMemoryReservationMb returns a boolean if a field has been set.

### GetName

`func (o *BaseSnapshotInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BaseSnapshotInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BaseSnapshotInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BaseSnapshotInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *BaseSnapshotInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *BaseSnapshotInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *BaseSnapshotInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *BaseSnapshotInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetRenderer3d

`func (o *BaseSnapshotInfo) GetRenderer3d() string`

GetRenderer3d returns the Renderer3d field if non-nil, zero value otherwise.

### GetRenderer3dOk

`func (o *BaseSnapshotInfo) GetRenderer3dOk() (*string, bool)`

GetRenderer3dOk returns a tuple with the Renderer3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer3d

`func (o *BaseSnapshotInfo) SetRenderer3d(v string)`

SetRenderer3d sets Renderer3d field to given value.

### HasRenderer3d

`func (o *BaseSnapshotInfo) HasRenderer3d() bool`

HasRenderer3d returns a boolean if a field has been set.

### GetTotalVideoMemoryMb

`func (o *BaseSnapshotInfo) GetTotalVideoMemoryMb() float64`

GetTotalVideoMemoryMb returns the TotalVideoMemoryMb field if non-nil, zero value otherwise.

### GetTotalVideoMemoryMbOk

`func (o *BaseSnapshotInfo) GetTotalVideoMemoryMbOk() (*float64, bool)`

GetTotalVideoMemoryMbOk returns a tuple with the TotalVideoMemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVideoMemoryMb

`func (o *BaseSnapshotInfo) SetTotalVideoMemoryMb(v float64)`

SetTotalVideoMemoryMb sets TotalVideoMemoryMb field to given value.

### HasTotalVideoMemoryMb

`func (o *BaseSnapshotInfo) HasTotalVideoMemoryMb() bool`

HasTotalVideoMemoryMb returns a boolean if a field has been set.

### GetVcenterId

`func (o *BaseSnapshotInfo) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *BaseSnapshotInfo) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *BaseSnapshotInfo) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *BaseSnapshotInfo) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.

### GetVgpuType

`func (o *BaseSnapshotInfo) GetVgpuType() string`

GetVgpuType returns the VgpuType field if non-nil, zero value otherwise.

### GetVgpuTypeOk

`func (o *BaseSnapshotInfo) GetVgpuTypeOk() (*string, bool)`

GetVgpuTypeOk returns a tuple with the VgpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuType

`func (o *BaseSnapshotInfo) SetVgpuType(v string)`

SetVgpuType sets VgpuType field to given value.

### HasVgpuType

`func (o *BaseSnapshotInfo) HasVgpuType() bool`

HasVgpuType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


