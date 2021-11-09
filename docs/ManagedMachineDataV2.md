# ManagedMachineDataV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseVmId** | Pointer to **string** | The base VM id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**BaseVmSnapshotId** | Pointer to **string** | The base VM snapshot id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**CloneErrorMessage** | Pointer to **string** | Cloning error message for this machine. This will be set for machine belonging to automated desktop pools when the machine&#39;s state is in PROVISIONING_ERROR or ERROR state. | [optional] 
**CloneErrorTime** | Pointer to **int64** | Cloning error time for this machine in milliseconds. Measured as epoch time. This will be set for machine belonging to automated desktop pools when the machine&#39;s state is in PROVISIONING_ERROR or ERROR state. | [optional] 
**CreateTime** | Pointer to **int64** | Time at which the machine was created in milliseconds. Measured as epoch time. | [optional] 
**DatastoreIds** | Pointer to **[]string** | The ids of the datastores. | [optional] 
**HostName** | Pointer to **string** | The name of the host on which this virtual machine is registered. | [optional] 
**ImageManagementStreamId** | Pointer to **string** | The id of the image management stream. This will be populated only for instant clone machines provisioned from pools created using image catalog. | [optional] 
**ImageManagementTagId** | Pointer to **string** | The id of the image management tag. This will be populated only for instant clone machines provisioned from pools created using image catalog. | [optional] 
**InHoldCustomization** | Pointer to **bool** | This condition determines if this virtual machine should be on hold before customization is started.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**InMaintenanceMode** | Pointer to **bool** | Indicates whether the Machine is in maintenance mode. | [optional] 
**LastMaintenanceTime** | Pointer to **int64** | The time of the last maintenance operation. | [optional] 
**LogoffPolicy** | Pointer to **string** | The user log off behavior at the time of maintenance. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions. | [optional] 
**MemoryMb** | Pointer to **int32** | The virtual machine physical memory in MB. | [optional] 
**MissingInVcenter** | Pointer to **bool** | This condition determines if the virtual machine is missing in vCenter Server.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**NetworkLabels** | Pointer to [**[]NetworkLabelData**](NetworkLabelData.md) | The network label(s) associated with this Machine.&lt;br&gt;This information will only be populated if a network label is explicitly assigned to this machine.&lt;br&gt;Otherwise, the machine inherits these properties from the parent virtual machine. | [optional] 
**Operation** | Pointer to **string** | The current maintenance operation on the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * PUSH_IMAGE: A push image operation. | [optional] 
**OperationState** | Pointer to **string** | The state of the current maintenance operation on the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * UNDEFINED: The operation state is unrecognized. * SCHEDULED: The operation is scheduled for future execution. * PROGRESSING: The operation is in progress. * COMPLETED: The operation has completed. * FAULT: The operation has encountered an error. * CANCELLING: The operation has been cancelled. * HOLDING: The operation has been paused. * CREATE: The operation is being initiated. | [optional] 
**Path** | Pointer to **string** | The virtual machine path.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;EndsWith&#39; and &#39;Contains&#39;.Field name to be used in filter : managedMachineData.path. | [optional] 
**PendingBaseVmId** | Pointer to **string** | The pending base VM id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**PendingBaseVmSnapshotId** | Pointer to **string** | The pending base VM snapshot id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**PendingImageManagementStreamId** | Pointer to **string** | The id of the pending image management stream. This will be populated only for instant clone machines provisioned from pools created using image catalog. | [optional] 
**PendingImageManagementTagId** | Pointer to **string** | The id of the pending image management tag. This will be populated only for machines belonging to Instant Clone farms created using image catalog. | [optional] 
**StorageAcceleratorState** | Pointer to **string** | The Horizon Storage Accelerator state. Storage acceleration will be available for managed machines if configured.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * OFF: The Storage Accelerator is off. * CURRENT: The machine cached data is updated. * OUT_OF_DATE: The machine cached data is not updated and requires regeneration. * ERROR: The Storage Accelerator has encountered an error. | [optional] 
**VirtualCenterId** | Pointer to **string** | The ID of the Virtual Center managing this machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**VirtualDisks** | Pointer to [**[]VirtualDiskData**](VirtualDiskData.md) | The virtual disks comprising the virtual machine. | [optional] 
**VirtualMachinePowerState** | Pointer to **string** | The virtual machine power state.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * POWERED_OFF: The machine is powered off. * POWERED_ON: The machine is powered on. * SUSPENDED: The machine is suspended. | [optional] 

## Methods

### NewManagedMachineDataV2

`func NewManagedMachineDataV2() *ManagedMachineDataV2`

NewManagedMachineDataV2 instantiates a new ManagedMachineDataV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedMachineDataV2WithDefaults

`func NewManagedMachineDataV2WithDefaults() *ManagedMachineDataV2`

NewManagedMachineDataV2WithDefaults instantiates a new ManagedMachineDataV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseVmId

`func (o *ManagedMachineDataV2) GetBaseVmId() string`

GetBaseVmId returns the BaseVmId field if non-nil, zero value otherwise.

### GetBaseVmIdOk

`func (o *ManagedMachineDataV2) GetBaseVmIdOk() (*string, bool)`

GetBaseVmIdOk returns a tuple with the BaseVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVmId

`func (o *ManagedMachineDataV2) SetBaseVmId(v string)`

SetBaseVmId sets BaseVmId field to given value.

### HasBaseVmId

`func (o *ManagedMachineDataV2) HasBaseVmId() bool`

HasBaseVmId returns a boolean if a field has been set.

### GetBaseVmSnapshotId

`func (o *ManagedMachineDataV2) GetBaseVmSnapshotId() string`

GetBaseVmSnapshotId returns the BaseVmSnapshotId field if non-nil, zero value otherwise.

### GetBaseVmSnapshotIdOk

`func (o *ManagedMachineDataV2) GetBaseVmSnapshotIdOk() (*string, bool)`

GetBaseVmSnapshotIdOk returns a tuple with the BaseVmSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVmSnapshotId

`func (o *ManagedMachineDataV2) SetBaseVmSnapshotId(v string)`

SetBaseVmSnapshotId sets BaseVmSnapshotId field to given value.

### HasBaseVmSnapshotId

`func (o *ManagedMachineDataV2) HasBaseVmSnapshotId() bool`

HasBaseVmSnapshotId returns a boolean if a field has been set.

### GetCloneErrorMessage

`func (o *ManagedMachineDataV2) GetCloneErrorMessage() string`

GetCloneErrorMessage returns the CloneErrorMessage field if non-nil, zero value otherwise.

### GetCloneErrorMessageOk

`func (o *ManagedMachineDataV2) GetCloneErrorMessageOk() (*string, bool)`

GetCloneErrorMessageOk returns a tuple with the CloneErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneErrorMessage

`func (o *ManagedMachineDataV2) SetCloneErrorMessage(v string)`

SetCloneErrorMessage sets CloneErrorMessage field to given value.

### HasCloneErrorMessage

`func (o *ManagedMachineDataV2) HasCloneErrorMessage() bool`

HasCloneErrorMessage returns a boolean if a field has been set.

### GetCloneErrorTime

`func (o *ManagedMachineDataV2) GetCloneErrorTime() int64`

GetCloneErrorTime returns the CloneErrorTime field if non-nil, zero value otherwise.

### GetCloneErrorTimeOk

`func (o *ManagedMachineDataV2) GetCloneErrorTimeOk() (*int64, bool)`

GetCloneErrorTimeOk returns a tuple with the CloneErrorTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneErrorTime

`func (o *ManagedMachineDataV2) SetCloneErrorTime(v int64)`

SetCloneErrorTime sets CloneErrorTime field to given value.

### HasCloneErrorTime

`func (o *ManagedMachineDataV2) HasCloneErrorTime() bool`

HasCloneErrorTime returns a boolean if a field has been set.

### GetCreateTime

`func (o *ManagedMachineDataV2) GetCreateTime() int64`

GetCreateTime returns the CreateTime field if non-nil, zero value otherwise.

### GetCreateTimeOk

`func (o *ManagedMachineDataV2) GetCreateTimeOk() (*int64, bool)`

GetCreateTimeOk returns a tuple with the CreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateTime

`func (o *ManagedMachineDataV2) SetCreateTime(v int64)`

SetCreateTime sets CreateTime field to given value.

### HasCreateTime

`func (o *ManagedMachineDataV2) HasCreateTime() bool`

HasCreateTime returns a boolean if a field has been set.

### GetDatastoreIds

`func (o *ManagedMachineDataV2) GetDatastoreIds() []string`

GetDatastoreIds returns the DatastoreIds field if non-nil, zero value otherwise.

### GetDatastoreIdsOk

`func (o *ManagedMachineDataV2) GetDatastoreIdsOk() (*[]string, bool)`

GetDatastoreIdsOk returns a tuple with the DatastoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreIds

`func (o *ManagedMachineDataV2) SetDatastoreIds(v []string)`

SetDatastoreIds sets DatastoreIds field to given value.

### HasDatastoreIds

`func (o *ManagedMachineDataV2) HasDatastoreIds() bool`

HasDatastoreIds returns a boolean if a field has been set.

### GetHostName

`func (o *ManagedMachineDataV2) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ManagedMachineDataV2) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ManagedMachineDataV2) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ManagedMachineDataV2) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetImageManagementStreamId

`func (o *ManagedMachineDataV2) GetImageManagementStreamId() string`

GetImageManagementStreamId returns the ImageManagementStreamId field if non-nil, zero value otherwise.

### GetImageManagementStreamIdOk

`func (o *ManagedMachineDataV2) GetImageManagementStreamIdOk() (*string, bool)`

GetImageManagementStreamIdOk returns a tuple with the ImageManagementStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageManagementStreamId

`func (o *ManagedMachineDataV2) SetImageManagementStreamId(v string)`

SetImageManagementStreamId sets ImageManagementStreamId field to given value.

### HasImageManagementStreamId

`func (o *ManagedMachineDataV2) HasImageManagementStreamId() bool`

HasImageManagementStreamId returns a boolean if a field has been set.

### GetImageManagementTagId

`func (o *ManagedMachineDataV2) GetImageManagementTagId() string`

GetImageManagementTagId returns the ImageManagementTagId field if non-nil, zero value otherwise.

### GetImageManagementTagIdOk

`func (o *ManagedMachineDataV2) GetImageManagementTagIdOk() (*string, bool)`

GetImageManagementTagIdOk returns a tuple with the ImageManagementTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageManagementTagId

`func (o *ManagedMachineDataV2) SetImageManagementTagId(v string)`

SetImageManagementTagId sets ImageManagementTagId field to given value.

### HasImageManagementTagId

`func (o *ManagedMachineDataV2) HasImageManagementTagId() bool`

HasImageManagementTagId returns a boolean if a field has been set.

### GetInHoldCustomization

`func (o *ManagedMachineDataV2) GetInHoldCustomization() bool`

GetInHoldCustomization returns the InHoldCustomization field if non-nil, zero value otherwise.

### GetInHoldCustomizationOk

`func (o *ManagedMachineDataV2) GetInHoldCustomizationOk() (*bool, bool)`

GetInHoldCustomizationOk returns a tuple with the InHoldCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInHoldCustomization

`func (o *ManagedMachineDataV2) SetInHoldCustomization(v bool)`

SetInHoldCustomization sets InHoldCustomization field to given value.

### HasInHoldCustomization

`func (o *ManagedMachineDataV2) HasInHoldCustomization() bool`

HasInHoldCustomization returns a boolean if a field has been set.

### GetInMaintenanceMode

`func (o *ManagedMachineDataV2) GetInMaintenanceMode() bool`

GetInMaintenanceMode returns the InMaintenanceMode field if non-nil, zero value otherwise.

### GetInMaintenanceModeOk

`func (o *ManagedMachineDataV2) GetInMaintenanceModeOk() (*bool, bool)`

GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInMaintenanceMode

`func (o *ManagedMachineDataV2) SetInMaintenanceMode(v bool)`

SetInMaintenanceMode sets InMaintenanceMode field to given value.

### HasInMaintenanceMode

`func (o *ManagedMachineDataV2) HasInMaintenanceMode() bool`

HasInMaintenanceMode returns a boolean if a field has been set.

### GetLastMaintenanceTime

`func (o *ManagedMachineDataV2) GetLastMaintenanceTime() int64`

GetLastMaintenanceTime returns the LastMaintenanceTime field if non-nil, zero value otherwise.

### GetLastMaintenanceTimeOk

`func (o *ManagedMachineDataV2) GetLastMaintenanceTimeOk() (*int64, bool)`

GetLastMaintenanceTimeOk returns a tuple with the LastMaintenanceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMaintenanceTime

`func (o *ManagedMachineDataV2) SetLastMaintenanceTime(v int64)`

SetLastMaintenanceTime sets LastMaintenanceTime field to given value.

### HasLastMaintenanceTime

`func (o *ManagedMachineDataV2) HasLastMaintenanceTime() bool`

HasLastMaintenanceTime returns a boolean if a field has been set.

### GetLogoffPolicy

`func (o *ManagedMachineDataV2) GetLogoffPolicy() string`

GetLogoffPolicy returns the LogoffPolicy field if non-nil, zero value otherwise.

### GetLogoffPolicyOk

`func (o *ManagedMachineDataV2) GetLogoffPolicyOk() (*string, bool)`

GetLogoffPolicyOk returns a tuple with the LogoffPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPolicy

`func (o *ManagedMachineDataV2) SetLogoffPolicy(v string)`

SetLogoffPolicy sets LogoffPolicy field to given value.

### HasLogoffPolicy

`func (o *ManagedMachineDataV2) HasLogoffPolicy() bool`

HasLogoffPolicy returns a boolean if a field has been set.

### GetMemoryMb

`func (o *ManagedMachineDataV2) GetMemoryMb() int32`

GetMemoryMb returns the MemoryMb field if non-nil, zero value otherwise.

### GetMemoryMbOk

`func (o *ManagedMachineDataV2) GetMemoryMbOk() (*int32, bool)`

GetMemoryMbOk returns a tuple with the MemoryMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMb

`func (o *ManagedMachineDataV2) SetMemoryMb(v int32)`

SetMemoryMb sets MemoryMb field to given value.

### HasMemoryMb

`func (o *ManagedMachineDataV2) HasMemoryMb() bool`

HasMemoryMb returns a boolean if a field has been set.

### GetMissingInVcenter

`func (o *ManagedMachineDataV2) GetMissingInVcenter() bool`

GetMissingInVcenter returns the MissingInVcenter field if non-nil, zero value otherwise.

### GetMissingInVcenterOk

`func (o *ManagedMachineDataV2) GetMissingInVcenterOk() (*bool, bool)`

GetMissingInVcenterOk returns a tuple with the MissingInVcenter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissingInVcenter

`func (o *ManagedMachineDataV2) SetMissingInVcenter(v bool)`

SetMissingInVcenter sets MissingInVcenter field to given value.

### HasMissingInVcenter

`func (o *ManagedMachineDataV2) HasMissingInVcenter() bool`

HasMissingInVcenter returns a boolean if a field has been set.

### GetNetworkLabels

`func (o *ManagedMachineDataV2) GetNetworkLabels() []NetworkLabelData`

GetNetworkLabels returns the NetworkLabels field if non-nil, zero value otherwise.

### GetNetworkLabelsOk

`func (o *ManagedMachineDataV2) GetNetworkLabelsOk() (*[]NetworkLabelData, bool)`

GetNetworkLabelsOk returns a tuple with the NetworkLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLabels

`func (o *ManagedMachineDataV2) SetNetworkLabels(v []NetworkLabelData)`

SetNetworkLabels sets NetworkLabels field to given value.

### HasNetworkLabels

`func (o *ManagedMachineDataV2) HasNetworkLabels() bool`

HasNetworkLabels returns a boolean if a field has been set.

### GetOperation

`func (o *ManagedMachineDataV2) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ManagedMachineDataV2) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ManagedMachineDataV2) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ManagedMachineDataV2) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOperationState

`func (o *ManagedMachineDataV2) GetOperationState() string`

GetOperationState returns the OperationState field if non-nil, zero value otherwise.

### GetOperationStateOk

`func (o *ManagedMachineDataV2) GetOperationStateOk() (*string, bool)`

GetOperationStateOk returns a tuple with the OperationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationState

`func (o *ManagedMachineDataV2) SetOperationState(v string)`

SetOperationState sets OperationState field to given value.

### HasOperationState

`func (o *ManagedMachineDataV2) HasOperationState() bool`

HasOperationState returns a boolean if a field has been set.

### GetPath

`func (o *ManagedMachineDataV2) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ManagedMachineDataV2) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ManagedMachineDataV2) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ManagedMachineDataV2) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetPendingBaseVmId

`func (o *ManagedMachineDataV2) GetPendingBaseVmId() string`

GetPendingBaseVmId returns the PendingBaseVmId field if non-nil, zero value otherwise.

### GetPendingBaseVmIdOk

`func (o *ManagedMachineDataV2) GetPendingBaseVmIdOk() (*string, bool)`

GetPendingBaseVmIdOk returns a tuple with the PendingBaseVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBaseVmId

`func (o *ManagedMachineDataV2) SetPendingBaseVmId(v string)`

SetPendingBaseVmId sets PendingBaseVmId field to given value.

### HasPendingBaseVmId

`func (o *ManagedMachineDataV2) HasPendingBaseVmId() bool`

HasPendingBaseVmId returns a boolean if a field has been set.

### GetPendingBaseVmSnapshotId

`func (o *ManagedMachineDataV2) GetPendingBaseVmSnapshotId() string`

GetPendingBaseVmSnapshotId returns the PendingBaseVmSnapshotId field if non-nil, zero value otherwise.

### GetPendingBaseVmSnapshotIdOk

`func (o *ManagedMachineDataV2) GetPendingBaseVmSnapshotIdOk() (*string, bool)`

GetPendingBaseVmSnapshotIdOk returns a tuple with the PendingBaseVmSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBaseVmSnapshotId

`func (o *ManagedMachineDataV2) SetPendingBaseVmSnapshotId(v string)`

SetPendingBaseVmSnapshotId sets PendingBaseVmSnapshotId field to given value.

### HasPendingBaseVmSnapshotId

`func (o *ManagedMachineDataV2) HasPendingBaseVmSnapshotId() bool`

HasPendingBaseVmSnapshotId returns a boolean if a field has been set.

### GetPendingImageManagementStreamId

`func (o *ManagedMachineDataV2) GetPendingImageManagementStreamId() string`

GetPendingImageManagementStreamId returns the PendingImageManagementStreamId field if non-nil, zero value otherwise.

### GetPendingImageManagementStreamIdOk

`func (o *ManagedMachineDataV2) GetPendingImageManagementStreamIdOk() (*string, bool)`

GetPendingImageManagementStreamIdOk returns a tuple with the PendingImageManagementStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingImageManagementStreamId

`func (o *ManagedMachineDataV2) SetPendingImageManagementStreamId(v string)`

SetPendingImageManagementStreamId sets PendingImageManagementStreamId field to given value.

### HasPendingImageManagementStreamId

`func (o *ManagedMachineDataV2) HasPendingImageManagementStreamId() bool`

HasPendingImageManagementStreamId returns a boolean if a field has been set.

### GetPendingImageManagementTagId

`func (o *ManagedMachineDataV2) GetPendingImageManagementTagId() string`

GetPendingImageManagementTagId returns the PendingImageManagementTagId field if non-nil, zero value otherwise.

### GetPendingImageManagementTagIdOk

`func (o *ManagedMachineDataV2) GetPendingImageManagementTagIdOk() (*string, bool)`

GetPendingImageManagementTagIdOk returns a tuple with the PendingImageManagementTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingImageManagementTagId

`func (o *ManagedMachineDataV2) SetPendingImageManagementTagId(v string)`

SetPendingImageManagementTagId sets PendingImageManagementTagId field to given value.

### HasPendingImageManagementTagId

`func (o *ManagedMachineDataV2) HasPendingImageManagementTagId() bool`

HasPendingImageManagementTagId returns a boolean if a field has been set.

### GetStorageAcceleratorState

`func (o *ManagedMachineDataV2) GetStorageAcceleratorState() string`

GetStorageAcceleratorState returns the StorageAcceleratorState field if non-nil, zero value otherwise.

### GetStorageAcceleratorStateOk

`func (o *ManagedMachineDataV2) GetStorageAcceleratorStateOk() (*string, bool)`

GetStorageAcceleratorStateOk returns a tuple with the StorageAcceleratorState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAcceleratorState

`func (o *ManagedMachineDataV2) SetStorageAcceleratorState(v string)`

SetStorageAcceleratorState sets StorageAcceleratorState field to given value.

### HasStorageAcceleratorState

`func (o *ManagedMachineDataV2) HasStorageAcceleratorState() bool`

HasStorageAcceleratorState returns a boolean if a field has been set.

### GetVirtualCenterId

`func (o *ManagedMachineDataV2) GetVirtualCenterId() string`

GetVirtualCenterId returns the VirtualCenterId field if non-nil, zero value otherwise.

### GetVirtualCenterIdOk

`func (o *ManagedMachineDataV2) GetVirtualCenterIdOk() (*string, bool)`

GetVirtualCenterIdOk returns a tuple with the VirtualCenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualCenterId

`func (o *ManagedMachineDataV2) SetVirtualCenterId(v string)`

SetVirtualCenterId sets VirtualCenterId field to given value.

### HasVirtualCenterId

`func (o *ManagedMachineDataV2) HasVirtualCenterId() bool`

HasVirtualCenterId returns a boolean if a field has been set.

### GetVirtualDisks

`func (o *ManagedMachineDataV2) GetVirtualDisks() []VirtualDiskData`

GetVirtualDisks returns the VirtualDisks field if non-nil, zero value otherwise.

### GetVirtualDisksOk

`func (o *ManagedMachineDataV2) GetVirtualDisksOk() (*[]VirtualDiskData, bool)`

GetVirtualDisksOk returns a tuple with the VirtualDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualDisks

`func (o *ManagedMachineDataV2) SetVirtualDisks(v []VirtualDiskData)`

SetVirtualDisks sets VirtualDisks field to given value.

### HasVirtualDisks

`func (o *ManagedMachineDataV2) HasVirtualDisks() bool`

HasVirtualDisks returns a boolean if a field has been set.

### GetVirtualMachinePowerState

`func (o *ManagedMachineDataV2) GetVirtualMachinePowerState() string`

GetVirtualMachinePowerState returns the VirtualMachinePowerState field if non-nil, zero value otherwise.

### GetVirtualMachinePowerStateOk

`func (o *ManagedMachineDataV2) GetVirtualMachinePowerStateOk() (*string, bool)`

GetVirtualMachinePowerStateOk returns a tuple with the VirtualMachinePowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualMachinePowerState

`func (o *ManagedMachineDataV2) SetVirtualMachinePowerState(v string)`

SetVirtualMachinePowerState sets VirtualMachinePowerState field to given value.

### HasVirtualMachinePowerState

`func (o *ManagedMachineDataV2) HasVirtualMachinePowerState() bool`

HasVirtualMachinePowerState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


