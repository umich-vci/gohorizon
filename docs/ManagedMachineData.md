# ManagedMachineData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CloneErrorMessage** | Pointer to **string** | Cloning error message for this machine. This will be set for machine belonging to automated desktop pools when the machine&#39;s state is in PROVISIONING_ERROR or ERROR state. | [optional] 
**CloneErrorTime** | Pointer to **int64** | Cloning error time for this machine in milliseconds. Measured as epoch time. This will be set for machine belonging to automated desktop pools when the machine&#39;s state is in PROVISIONING_ERROR or ERROR state. | [optional] 
**CreateTime** | Pointer to **int64** | Time at which the machine was created in milliseconds. Measured as epoch time. | [optional] 
**HostName** | Pointer to **string** | The name of the host on which this virtual machine is registered. | [optional] 
**InHoldCustomization** | Pointer to **bool** | This condition determines if this virtual machine should be on hold before customization is started.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**InMaintenanceMode** | Pointer to **bool** | Indicates whether the Machine is in maintenance mode. | [optional] 
**MemoryMb** | Pointer to **int32** | The virtual machine physical memory in MB. | [optional] 
**MissingInVcenter** | Pointer to **bool** | This condition determines if the virtual machine is missing in vCenter Server.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**NetworkLabels** | Pointer to [**[]NetworkLabelData**](NetworkLabelData.md) | The network label(s) associated with this Machine.&lt;br&gt;This information will only be populated if a network label is explicitly assigned to this machine.&lt;br&gt;Otherwise, the machine inherits these properties from the parent virtual machine. | [optional] 
**Path** | Pointer to **string** | The virtual machine path.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;EndsWith&#39; and &#39;Contains&#39;.Field name to be used in filter : managedMachineData.path. | [optional] 
**StorageAcceleratorState** | Pointer to **string** | The Horizon Storage Accelerator state. Storage acceleration will be available for managed machines if configured.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * OFF: The Storage Accelerator is off. * CURRENT: The machine cached data is updated. * OUT_OF_DATE: The machine cached data is not updated and requires regeneration. * ERROR: The Storage Accelerator has encountered an error. | [optional] 
**VirtualCenterId** | Pointer to **string** | The ID of the Virtual Center managing this machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**VirtualDisks** | Pointer to [**[]VirtualDiskData**](VirtualDiskData.md) | The virtual disks comprising the virtual machine. | [optional] 
**VirtualMachinePowerState** | Pointer to **string** | The virtual machine power state.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * POWERED_OFF: The machine is powered off. * POWERED_ON: The machine is powered on. * SUSPENDED: The machine is suspended. | [optional] 

## Methods

### NewManagedMachineData

`func NewManagedMachineData() *ManagedMachineData`

NewManagedMachineData instantiates a new ManagedMachineData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedMachineDataWithDefaults

`func NewManagedMachineDataWithDefaults() *ManagedMachineData`

NewManagedMachineDataWithDefaults instantiates a new ManagedMachineData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCloneErrorMessage

`func (o *ManagedMachineData) GetCloneErrorMessage() string`

GetCloneErrorMessage returns the CloneErrorMessage field if non-nil, zero value otherwise.

### GetCloneErrorMessageOk

`func (o *ManagedMachineData) GetCloneErrorMessageOk() (*string, bool)`

GetCloneErrorMessageOk returns a tuple with the CloneErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneErrorMessage

`func (o *ManagedMachineData) SetCloneErrorMessage(v string)`

SetCloneErrorMessage sets CloneErrorMessage field to given value.

### HasCloneErrorMessage

`func (o *ManagedMachineData) HasCloneErrorMessage() bool`

HasCloneErrorMessage returns a boolean if a field has been set.

### GetCloneErrorTime

`func (o *ManagedMachineData) GetCloneErrorTime() int64`

GetCloneErrorTime returns the CloneErrorTime field if non-nil, zero value otherwise.

### GetCloneErrorTimeOk

`func (o *ManagedMachineData) GetCloneErrorTimeOk() (*int64, bool)`

GetCloneErrorTimeOk returns a tuple with the CloneErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneErrorTime

`func (o *ManagedMachineData) SetCloneErrorTime(v int64)`

SetCloneErrorTime sets CloneErrorTime field to given value.

### HasCloneErrorTime

`func (o *ManagedMachineData) HasCloneErrorTime() bool`

HasCloneErrorTime returns a boolean if a field has been set.

### GetCreateTime

`func (o *ManagedMachineData) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ManagedMachineData) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ManagedMachineData) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ManagedMachineData) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetHostName

`func (o *ManagedMachineData) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ManagedMachineData) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ManagedMachineData) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ManagedMachineData) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetInHoldCustomization

`func (o *ManagedMachineData) GetInHoldCustomization() bool`

GetInHoldCustomization returns the InHoldCustomization field if non-nil, zero value otherwise.

### GetInHoldCustomizationOk

`func (o *ManagedMachineData) GetInHoldCustomizationOk() (*bool, bool)`

GetInHoldCustomizationOk returns a tuple with the InHoldCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInHoldCustomization

`func (o *ManagedMachineData) SetInHoldCustomization(v bool)`

SetInHoldCustomization sets InHoldCustomization field to given value.

### HasInHoldCustomization

`func (o *ManagedMachineData) HasInHoldCustomization() bool`

HasInHoldCustomization returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *ManagedMachineData) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *ManagedMachineData) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *ManagedMachineData) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *ManagedMachineData) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### GetMemoryMb

`func (o *ManagedMachineData) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *ManagedMachineData) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *ManagedMachineData) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *ManagedMachineData) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetMissingInVcenter

`func (o *ManagedMachineData) GetMissingInVcenter() bool`

GetMissingInVcenter returns the MissingInVcenter field if non-nil, zero value otherwise.

### GetMissingInVcenterOk

`func (o *ManagedMachineData) GetMissingInVcenterOk() (*bool, bool)`

GetMissingInVcenterOk returns a tuple with the MissingInVcenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingInVcenter

`func (o *ManagedMachineData) SetMissingInVcenter(v bool)`

SetMissingInVcenter sets MissingInVcenter field to given value.

### HasMissingInVcenter

`func (o *ManagedMachineData) HasMissingInVcenter() bool`

HasMissingInVcenter returns a boolean if a field has been set.

### GetNetworkLabels

`func (o *ManagedMachineData) GetNetworkLabels() []NetworkLabelData`

GetNetworkLabels returns the NetworkLabels field if non-nil, zero value otherwise.

### GetNetworkLabelsOk

`func (o *ManagedMachineData) GetNetworkLabelsOk() (*[]NetworkLabelData, bool)`

GetNetworkLabelsOk returns a tuple with the NetworkLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLabels

`func (o *ManagedMachineData) SetNetworkLabels(v []NetworkLabelData)`

SetNetworkLabels sets NetworkLabels field to given value.

### HasNetworkLabels

`func (o *ManagedMachineData) HasNetworkLabels() bool`

HasNetworkLabels returns a boolean if a field has been set.

### GetPath

`func (o *ManagedMachineData) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManagedMachineData) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManagedMachineData) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManagedMachineData) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetStorageAcceleratorState

`func (o *ManagedMachineData) GetStorageAcceleratorState() string`

GetStorageAcceleratorState returns the StorageAcceleratorState field if non-nil, zero value otherwise.

### GetStorageAcceleratorStateOk

`func (o *ManagedMachineData) GetStorageAcceleratorStateOk() (*string, bool)`

GetStorageAcceleratorStateOk returns a tuple with the StorageAcceleratorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAcceleratorState

`func (o *ManagedMachineData) SetStorageAcceleratorState(v string)`

SetStorageAcceleratorState sets StorageAcceleratorState field to given value.

### HasStorageAcceleratorState

`func (o *ManagedMachineData) HasStorageAcceleratorState() bool`

HasStorageAcceleratorState returns a boolean if a field has been set.

### GetVirtualCenterId

`func (o *ManagedMachineData) GetVirtualCenterId() string`

GetVirtualCenterId returns the VirtualCenterId field if non-nil, zero value otherwise.

### GetVirtualCenterIdOk

`func (o *ManagedMachineData) GetVirtualCenterIdOk() (*string, bool)`

GetVirtualCenterIdOk returns a tuple with the VirtualCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCenterId

`func (o *ManagedMachineData) SetVirtualCenterId(v string)`

SetVirtualCenterId sets VirtualCenterId field to given value.

### HasVirtualCenterId

`func (o *ManagedMachineData) HasVirtualCenterId() bool`

HasVirtualCenterId returns a boolean if a field has been set.

### GetVirtualDisks

`func (o *ManagedMachineData) GetVirtualDisks() []VirtualDiskData`

GetVirtualDisks returns the VirtualDisks field if non-nil, zero value otherwise.

### GetVirtualDisksOk

`func (o *ManagedMachineData) GetVirtualDisksOk() (*[]VirtualDiskData, bool)`

GetVirtualDisksOk returns a tuple with the VirtualDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisks

`func (o *ManagedMachineData) SetVirtualDisks(v []VirtualDiskData)`

SetVirtualDisks sets VirtualDisks field to given value.

### HasVirtualDisks

`func (o *ManagedMachineData) HasVirtualDisks() bool`

HasVirtualDisks returns a boolean if a field has been set.

### GetVirtualMachinePowerState

`func (o *ManagedMachineData) GetVirtualMachinePowerState() string`

GetVirtualMachinePowerState returns the VirtualMachinePowerState field if non-nil, zero value otherwise.

### GetVirtualMachinePowerStateOk

`func (o *ManagedMachineData) GetVirtualMachinePowerStateOk() (*string, bool)`

GetVirtualMachinePowerStateOk returns a tuple with the VirtualMachinePowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachinePowerState

`func (o *ManagedMachineData) SetVirtualMachinePowerState(v string)`

SetVirtualMachinePowerState sets VirtualMachinePowerState field to given value.

### HasVirtualMachinePowerState

`func (o *ManagedMachineData) HasVirtualMachinePowerState() bool`

HasVirtualMachinePowerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


