/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2106
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// ManagedMachineDataV2 Information related to Managed machines V2.
type ManagedMachineDataV2 struct {
	// The base VM id.<br>Supported Filters : 'Equals'.
	BaseVmId *string `json:"base_vm_id,omitempty"`
	// The base VM snapshot id.<br>Supported Filters : 'Equals'.
	BaseVmSnapshotId *string `json:"base_vm_snapshot_id,omitempty"`
	// Cloning error message for this machine. This will be set for machine belonging to automated desktop pools when the machine's state is in PROVISIONING_ERROR or ERROR state.
	CloneErrorMessage *string `json:"clone_error_message,omitempty"`
	// Cloning error time for this machine in milliseconds. Measured as epoch time. This will be set for machine belonging to automated desktop pools when the machine's state is in PROVISIONING_ERROR or ERROR state.
	CloneErrorTime *int64 `json:"clone_error_time,omitempty"`
	// Time at which the machine was created in milliseconds. Measured as epoch time.
	CreateTime *int64 `json:"create_time,omitempty"`
	// The ids of the datastores.
	DatastoreIds *[]string `json:"datastore_ids,omitempty"`
	// The name of the host on which this virtual machine is registered.
	HostName *string `json:"host_name,omitempty"`
	// The id of the image management stream. This will be populated only for instant clone machines provisioned from pools created using image catalog.
	ImageManagementStreamId *string `json:"image_management_stream_id,omitempty"`
	// The id of the image management tag. This will be populated only for instant clone machines provisioned from pools created using image catalog.
	ImageManagementTagId *string `json:"image_management_tag_id,omitempty"`
	// This condition determines if this virtual machine should be on hold before customization is started.<br>Supported Filters : 'Equals'.
	InHoldCustomization *bool `json:"in_hold_customization,omitempty"`
	// Indicates whether the Machine is in maintenance mode.
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`
	// The time of the last maintenance operation.
	LastMaintenanceTime *int64 `json:"last_maintenance_time,omitempty"`
	// The user log off behavior at the time of maintenance. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions.
	LogoffPolicy *string `json:"logoff_policy,omitempty"`
	// The virtual machine physical memory in MB.
	MemoryMb *int32 `json:"memory_mb,omitempty"`
	// This condition determines if the virtual machine is missing in vCenter Server.<br>Supported Filters : 'Equals'.
	MissingInVcenter *bool `json:"missing_in_vcenter,omitempty"`
	// The network label(s) associated with this Machine.<br>This information will only be populated if a network label is explicitly assigned to this machine.<br>Otherwise, the machine inherits these properties from the parent virtual machine.
	NetworkLabels *[]NetworkLabelData `json:"network_labels,omitempty"`
	// The current maintenance operation on the machine.<br>Supported Filters : 'Equals'. * PUSH_IMAGE: A push image operation.
	Operation *string `json:"operation,omitempty"`
	// The state of the current maintenance operation on the machine.<br>Supported Filters : 'Equals'. * UNDEFINED: The operation state is unrecognized. * SCHEDULED: The operation is scheduled for future execution. * PROGRESSING: The operation is in progress. * COMPLETED: The operation has completed. * FAULT: The operation has encountered an error. * CANCELLING: The operation has been cancelled. * HOLDING: The operation has been paused. * CREATE: The operation is being initiated.
	OperationState *string `json:"operation_state,omitempty"`
	// The virtual machine path.<br>Supported Filters : 'Equals', 'StartsWith', 'EndsWith' and 'Contains'.Field name to be used in filter : managedMachineData.path.
	Path *string `json:"path,omitempty"`
	// The pending base VM id.<br>Supported Filters : 'Equals'.
	PendingBaseVmId *string `json:"pending_base_vm_id,omitempty"`
	// The pending base VM snapshot id.<br>Supported Filters : 'Equals'.
	PendingBaseVmSnapshotId *string `json:"pending_base_vm_snapshot_id,omitempty"`
	// The id of the pending image management stream. This will be populated only for instant clone machines provisioned from pools created using image catalog.
	PendingImageManagementStreamId *string `json:"pending_image_management_stream_id,omitempty"`
	// The id of the pending image management tag. This will be populated only for machines belonging to Instant Clone farms created using image catalog.
	PendingImageManagementTagId *string `json:"pending_image_management_tag_id,omitempty"`
	// The Horizon Storage Accelerator state. Storage acceleration will be available for managed machines if configured.<br>Supported Filters : 'Equals'. * OFF: The Storage Accelerator is off. * CURRENT: The machine cached data is updated. * OUT_OF_DATE: The machine cached data is not updated and requires regeneration. * ERROR: The Storage Accelerator has encountered an error.
	StorageAcceleratorState *string `json:"storage_accelerator_state,omitempty"`
	// The ID of the Virtual Center managing this machine.<br>Supported Filters : 'Equals'.
	VirtualCenterId *string `json:"virtual_center_id,omitempty"`
	// The virtual disks comprising the virtual machine.
	VirtualDisks *[]VirtualDiskData `json:"virtual_disks,omitempty"`
	// The virtual machine power state.<br>Supported Filters : 'Equals'. * POWERED_OFF: The machine is powered off. * POWERED_ON: The machine is powered on. * SUSPENDED: The machine is suspended.
	VirtualMachinePowerState *string `json:"virtual_machine_power_state,omitempty"`
}

// NewManagedMachineDataV2 instantiates a new ManagedMachineDataV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManagedMachineDataV2() *ManagedMachineDataV2 {
	this := ManagedMachineDataV2{}
	return &this
}

// NewManagedMachineDataV2WithDefaults instantiates a new ManagedMachineDataV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManagedMachineDataV2WithDefaults() *ManagedMachineDataV2 {
	this := ManagedMachineDataV2{}
	return &this
}

// GetBaseVmId returns the BaseVmId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetBaseVmId() string {
	if o == nil || o.BaseVmId == nil {
		var ret string
		return ret
	}
	return *o.BaseVmId
}

// GetBaseVmIdOk returns a tuple with the BaseVmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetBaseVmIdOk() (*string, bool) {
	if o == nil || o.BaseVmId == nil {
		return nil, false
	}
	return o.BaseVmId, true
}

// HasBaseVmId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasBaseVmId() bool {
	if o != nil && o.BaseVmId != nil {
		return true
	}

	return false
}

// SetBaseVmId gets a reference to the given string and assigns it to the BaseVmId field.
func (o *ManagedMachineDataV2) SetBaseVmId(v string) {
	o.BaseVmId = &v
}

// GetBaseVmSnapshotId returns the BaseVmSnapshotId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetBaseVmSnapshotId() string {
	if o == nil || o.BaseVmSnapshotId == nil {
		var ret string
		return ret
	}
	return *o.BaseVmSnapshotId
}

// GetBaseVmSnapshotIdOk returns a tuple with the BaseVmSnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetBaseVmSnapshotIdOk() (*string, bool) {
	if o == nil || o.BaseVmSnapshotId == nil {
		return nil, false
	}
	return o.BaseVmSnapshotId, true
}

// HasBaseVmSnapshotId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasBaseVmSnapshotId() bool {
	if o != nil && o.BaseVmSnapshotId != nil {
		return true
	}

	return false
}

// SetBaseVmSnapshotId gets a reference to the given string and assigns it to the BaseVmSnapshotId field.
func (o *ManagedMachineDataV2) SetBaseVmSnapshotId(v string) {
	o.BaseVmSnapshotId = &v
}

// GetCloneErrorMessage returns the CloneErrorMessage field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetCloneErrorMessage() string {
	if o == nil || o.CloneErrorMessage == nil {
		var ret string
		return ret
	}
	return *o.CloneErrorMessage
}

// GetCloneErrorMessageOk returns a tuple with the CloneErrorMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetCloneErrorMessageOk() (*string, bool) {
	if o == nil || o.CloneErrorMessage == nil {
		return nil, false
	}
	return o.CloneErrorMessage, true
}

// HasCloneErrorMessage returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasCloneErrorMessage() bool {
	if o != nil && o.CloneErrorMessage != nil {
		return true
	}

	return false
}

// SetCloneErrorMessage gets a reference to the given string and assigns it to the CloneErrorMessage field.
func (o *ManagedMachineDataV2) SetCloneErrorMessage(v string) {
	o.CloneErrorMessage = &v
}

// GetCloneErrorTime returns the CloneErrorTime field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetCloneErrorTime() int64 {
	if o == nil || o.CloneErrorTime == nil {
		var ret int64
		return ret
	}
	return *o.CloneErrorTime
}

// GetCloneErrorTimeOk returns a tuple with the CloneErrorTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetCloneErrorTimeOk() (*int64, bool) {
	if o == nil || o.CloneErrorTime == nil {
		return nil, false
	}
	return o.CloneErrorTime, true
}

// HasCloneErrorTime returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasCloneErrorTime() bool {
	if o != nil && o.CloneErrorTime != nil {
		return true
	}

	return false
}

// SetCloneErrorTime gets a reference to the given int64 and assigns it to the CloneErrorTime field.
func (o *ManagedMachineDataV2) SetCloneErrorTime(v int64) {
	o.CloneErrorTime = &v
}

// GetCreateTime returns the CreateTime field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetCreateTime() int64 {
	if o == nil || o.CreateTime == nil {
		var ret int64
		return ret
	}
	return *o.CreateTime
}

// GetCreateTimeOk returns a tuple with the CreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetCreateTimeOk() (*int64, bool) {
	if o == nil || o.CreateTime == nil {
		return nil, false
	}
	return o.CreateTime, true
}

// HasCreateTime returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasCreateTime() bool {
	if o != nil && o.CreateTime != nil {
		return true
	}

	return false
}

// SetCreateTime gets a reference to the given int64 and assigns it to the CreateTime field.
func (o *ManagedMachineDataV2) SetCreateTime(v int64) {
	o.CreateTime = &v
}

// GetDatastoreIds returns the DatastoreIds field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetDatastoreIds() []string {
	if o == nil || o.DatastoreIds == nil {
		var ret []string
		return ret
	}
	return *o.DatastoreIds
}

// GetDatastoreIdsOk returns a tuple with the DatastoreIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetDatastoreIdsOk() (*[]string, bool) {
	if o == nil || o.DatastoreIds == nil {
		return nil, false
	}
	return o.DatastoreIds, true
}

// HasDatastoreIds returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasDatastoreIds() bool {
	if o != nil && o.DatastoreIds != nil {
		return true
	}

	return false
}

// SetDatastoreIds gets a reference to the given []string and assigns it to the DatastoreIds field.
func (o *ManagedMachineDataV2) SetDatastoreIds(v []string) {
	o.DatastoreIds = &v
}

// GetHostName returns the HostName field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetHostName() string {
	if o == nil || o.HostName == nil {
		var ret string
		return ret
	}
	return *o.HostName
}

// GetHostNameOk returns a tuple with the HostName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetHostNameOk() (*string, bool) {
	if o == nil || o.HostName == nil {
		return nil, false
	}
	return o.HostName, true
}

// HasHostName returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasHostName() bool {
	if o != nil && o.HostName != nil {
		return true
	}

	return false
}

// SetHostName gets a reference to the given string and assigns it to the HostName field.
func (o *ManagedMachineDataV2) SetHostName(v string) {
	o.HostName = &v
}

// GetImageManagementStreamId returns the ImageManagementStreamId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetImageManagementStreamId() string {
	if o == nil || o.ImageManagementStreamId == nil {
		var ret string
		return ret
	}
	return *o.ImageManagementStreamId
}

// GetImageManagementStreamIdOk returns a tuple with the ImageManagementStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetImageManagementStreamIdOk() (*string, bool) {
	if o == nil || o.ImageManagementStreamId == nil {
		return nil, false
	}
	return o.ImageManagementStreamId, true
}

// HasImageManagementStreamId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasImageManagementStreamId() bool {
	if o != nil && o.ImageManagementStreamId != nil {
		return true
	}

	return false
}

// SetImageManagementStreamId gets a reference to the given string and assigns it to the ImageManagementStreamId field.
func (o *ManagedMachineDataV2) SetImageManagementStreamId(v string) {
	o.ImageManagementStreamId = &v
}

// GetImageManagementTagId returns the ImageManagementTagId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetImageManagementTagId() string {
	if o == nil || o.ImageManagementTagId == nil {
		var ret string
		return ret
	}
	return *o.ImageManagementTagId
}

// GetImageManagementTagIdOk returns a tuple with the ImageManagementTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetImageManagementTagIdOk() (*string, bool) {
	if o == nil || o.ImageManagementTagId == nil {
		return nil, false
	}
	return o.ImageManagementTagId, true
}

// HasImageManagementTagId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasImageManagementTagId() bool {
	if o != nil && o.ImageManagementTagId != nil {
		return true
	}

	return false
}

// SetImageManagementTagId gets a reference to the given string and assigns it to the ImageManagementTagId field.
func (o *ManagedMachineDataV2) SetImageManagementTagId(v string) {
	o.ImageManagementTagId = &v
}

// GetInHoldCustomization returns the InHoldCustomization field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetInHoldCustomization() bool {
	if o == nil || o.InHoldCustomization == nil {
		var ret bool
		return ret
	}
	return *o.InHoldCustomization
}

// GetInHoldCustomizationOk returns a tuple with the InHoldCustomization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetInHoldCustomizationOk() (*bool, bool) {
	if o == nil || o.InHoldCustomization == nil {
		return nil, false
	}
	return o.InHoldCustomization, true
}

// HasInHoldCustomization returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasInHoldCustomization() bool {
	if o != nil && o.InHoldCustomization != nil {
		return true
	}

	return false
}

// SetInHoldCustomization gets a reference to the given bool and assigns it to the InHoldCustomization field.
func (o *ManagedMachineDataV2) SetInHoldCustomization(v bool) {
	o.InHoldCustomization = &v
}

// GetInMaintenanceMode returns the InMaintenanceMode field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetInMaintenanceMode() bool {
	if o == nil || o.InMaintenanceMode == nil {
		var ret bool
		return ret
	}
	return *o.InMaintenanceMode
}

// GetInMaintenanceModeOk returns a tuple with the InMaintenanceMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetInMaintenanceModeOk() (*bool, bool) {
	if o == nil || o.InMaintenanceMode == nil {
		return nil, false
	}
	return o.InMaintenanceMode, true
}

// HasInMaintenanceMode returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasInMaintenanceMode() bool {
	if o != nil && o.InMaintenanceMode != nil {
		return true
	}

	return false
}

// SetInMaintenanceMode gets a reference to the given bool and assigns it to the InMaintenanceMode field.
func (o *ManagedMachineDataV2) SetInMaintenanceMode(v bool) {
	o.InMaintenanceMode = &v
}

// GetLastMaintenanceTime returns the LastMaintenanceTime field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetLastMaintenanceTime() int64 {
	if o == nil || o.LastMaintenanceTime == nil {
		var ret int64
		return ret
	}
	return *o.LastMaintenanceTime
}

// GetLastMaintenanceTimeOk returns a tuple with the LastMaintenanceTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetLastMaintenanceTimeOk() (*int64, bool) {
	if o == nil || o.LastMaintenanceTime == nil {
		return nil, false
	}
	return o.LastMaintenanceTime, true
}

// HasLastMaintenanceTime returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasLastMaintenanceTime() bool {
	if o != nil && o.LastMaintenanceTime != nil {
		return true
	}

	return false
}

// SetLastMaintenanceTime gets a reference to the given int64 and assigns it to the LastMaintenanceTime field.
func (o *ManagedMachineDataV2) SetLastMaintenanceTime(v int64) {
	o.LastMaintenanceTime = &v
}

// GetLogoffPolicy returns the LogoffPolicy field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetLogoffPolicy() string {
	if o == nil || o.LogoffPolicy == nil {
		var ret string
		return ret
	}
	return *o.LogoffPolicy
}

// GetLogoffPolicyOk returns a tuple with the LogoffPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetLogoffPolicyOk() (*string, bool) {
	if o == nil || o.LogoffPolicy == nil {
		return nil, false
	}
	return o.LogoffPolicy, true
}

// HasLogoffPolicy returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasLogoffPolicy() bool {
	if o != nil && o.LogoffPolicy != nil {
		return true
	}

	return false
}

// SetLogoffPolicy gets a reference to the given string and assigns it to the LogoffPolicy field.
func (o *ManagedMachineDataV2) SetLogoffPolicy(v string) {
	o.LogoffPolicy = &v
}

// GetMemoryMb returns the MemoryMb field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetMemoryMb() int32 {
	if o == nil || o.MemoryMb == nil {
		var ret int32
		return ret
	}
	return *o.MemoryMb
}

// GetMemoryMbOk returns a tuple with the MemoryMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetMemoryMbOk() (*int32, bool) {
	if o == nil || o.MemoryMb == nil {
		return nil, false
	}
	return o.MemoryMb, true
}

// HasMemoryMb returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasMemoryMb() bool {
	if o != nil && o.MemoryMb != nil {
		return true
	}

	return false
}

// SetMemoryMb gets a reference to the given int32 and assigns it to the MemoryMb field.
func (o *ManagedMachineDataV2) SetMemoryMb(v int32) {
	o.MemoryMb = &v
}

// GetMissingInVcenter returns the MissingInVcenter field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetMissingInVcenter() bool {
	if o == nil || o.MissingInVcenter == nil {
		var ret bool
		return ret
	}
	return *o.MissingInVcenter
}

// GetMissingInVcenterOk returns a tuple with the MissingInVcenter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetMissingInVcenterOk() (*bool, bool) {
	if o == nil || o.MissingInVcenter == nil {
		return nil, false
	}
	return o.MissingInVcenter, true
}

// HasMissingInVcenter returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasMissingInVcenter() bool {
	if o != nil && o.MissingInVcenter != nil {
		return true
	}

	return false
}

// SetMissingInVcenter gets a reference to the given bool and assigns it to the MissingInVcenter field.
func (o *ManagedMachineDataV2) SetMissingInVcenter(v bool) {
	o.MissingInVcenter = &v
}

// GetNetworkLabels returns the NetworkLabels field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetNetworkLabels() []NetworkLabelData {
	if o == nil || o.NetworkLabels == nil {
		var ret []NetworkLabelData
		return ret
	}
	return *o.NetworkLabels
}

// GetNetworkLabelsOk returns a tuple with the NetworkLabels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetNetworkLabelsOk() (*[]NetworkLabelData, bool) {
	if o == nil || o.NetworkLabels == nil {
		return nil, false
	}
	return o.NetworkLabels, true
}

// HasNetworkLabels returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasNetworkLabels() bool {
	if o != nil && o.NetworkLabels != nil {
		return true
	}

	return false
}

// SetNetworkLabels gets a reference to the given []NetworkLabelData and assigns it to the NetworkLabels field.
func (o *ManagedMachineDataV2) SetNetworkLabels(v []NetworkLabelData) {
	o.NetworkLabels = &v
}

// GetOperation returns the Operation field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetOperation() string {
	if o == nil || o.Operation == nil {
		var ret string
		return ret
	}
	return *o.Operation
}

// GetOperationOk returns a tuple with the Operation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetOperationOk() (*string, bool) {
	if o == nil || o.Operation == nil {
		return nil, false
	}
	return o.Operation, true
}

// HasOperation returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasOperation() bool {
	if o != nil && o.Operation != nil {
		return true
	}

	return false
}

// SetOperation gets a reference to the given string and assigns it to the Operation field.
func (o *ManagedMachineDataV2) SetOperation(v string) {
	o.Operation = &v
}

// GetOperationState returns the OperationState field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetOperationState() string {
	if o == nil || o.OperationState == nil {
		var ret string
		return ret
	}
	return *o.OperationState
}

// GetOperationStateOk returns a tuple with the OperationState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetOperationStateOk() (*string, bool) {
	if o == nil || o.OperationState == nil {
		return nil, false
	}
	return o.OperationState, true
}

// HasOperationState returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasOperationState() bool {
	if o != nil && o.OperationState != nil {
		return true
	}

	return false
}

// SetOperationState gets a reference to the given string and assigns it to the OperationState field.
func (o *ManagedMachineDataV2) SetOperationState(v string) {
	o.OperationState = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ManagedMachineDataV2) SetPath(v string) {
	o.Path = &v
}

// GetPendingBaseVmId returns the PendingBaseVmId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetPendingBaseVmId() string {
	if o == nil || o.PendingBaseVmId == nil {
		var ret string
		return ret
	}
	return *o.PendingBaseVmId
}

// GetPendingBaseVmIdOk returns a tuple with the PendingBaseVmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetPendingBaseVmIdOk() (*string, bool) {
	if o == nil || o.PendingBaseVmId == nil {
		return nil, false
	}
	return o.PendingBaseVmId, true
}

// HasPendingBaseVmId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasPendingBaseVmId() bool {
	if o != nil && o.PendingBaseVmId != nil {
		return true
	}

	return false
}

// SetPendingBaseVmId gets a reference to the given string and assigns it to the PendingBaseVmId field.
func (o *ManagedMachineDataV2) SetPendingBaseVmId(v string) {
	o.PendingBaseVmId = &v
}

// GetPendingBaseVmSnapshotId returns the PendingBaseVmSnapshotId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetPendingBaseVmSnapshotId() string {
	if o == nil || o.PendingBaseVmSnapshotId == nil {
		var ret string
		return ret
	}
	return *o.PendingBaseVmSnapshotId
}

// GetPendingBaseVmSnapshotIdOk returns a tuple with the PendingBaseVmSnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetPendingBaseVmSnapshotIdOk() (*string, bool) {
	if o == nil || o.PendingBaseVmSnapshotId == nil {
		return nil, false
	}
	return o.PendingBaseVmSnapshotId, true
}

// HasPendingBaseVmSnapshotId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasPendingBaseVmSnapshotId() bool {
	if o != nil && o.PendingBaseVmSnapshotId != nil {
		return true
	}

	return false
}

// SetPendingBaseVmSnapshotId gets a reference to the given string and assigns it to the PendingBaseVmSnapshotId field.
func (o *ManagedMachineDataV2) SetPendingBaseVmSnapshotId(v string) {
	o.PendingBaseVmSnapshotId = &v
}

// GetPendingImageManagementStreamId returns the PendingImageManagementStreamId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetPendingImageManagementStreamId() string {
	if o == nil || o.PendingImageManagementStreamId == nil {
		var ret string
		return ret
	}
	return *o.PendingImageManagementStreamId
}

// GetPendingImageManagementStreamIdOk returns a tuple with the PendingImageManagementStreamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetPendingImageManagementStreamIdOk() (*string, bool) {
	if o == nil || o.PendingImageManagementStreamId == nil {
		return nil, false
	}
	return o.PendingImageManagementStreamId, true
}

// HasPendingImageManagementStreamId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasPendingImageManagementStreamId() bool {
	if o != nil && o.PendingImageManagementStreamId != nil {
		return true
	}

	return false
}

// SetPendingImageManagementStreamId gets a reference to the given string and assigns it to the PendingImageManagementStreamId field.
func (o *ManagedMachineDataV2) SetPendingImageManagementStreamId(v string) {
	o.PendingImageManagementStreamId = &v
}

// GetPendingImageManagementTagId returns the PendingImageManagementTagId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetPendingImageManagementTagId() string {
	if o == nil || o.PendingImageManagementTagId == nil {
		var ret string
		return ret
	}
	return *o.PendingImageManagementTagId
}

// GetPendingImageManagementTagIdOk returns a tuple with the PendingImageManagementTagId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetPendingImageManagementTagIdOk() (*string, bool) {
	if o == nil || o.PendingImageManagementTagId == nil {
		return nil, false
	}
	return o.PendingImageManagementTagId, true
}

// HasPendingImageManagementTagId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasPendingImageManagementTagId() bool {
	if o != nil && o.PendingImageManagementTagId != nil {
		return true
	}

	return false
}

// SetPendingImageManagementTagId gets a reference to the given string and assigns it to the PendingImageManagementTagId field.
func (o *ManagedMachineDataV2) SetPendingImageManagementTagId(v string) {
	o.PendingImageManagementTagId = &v
}

// GetStorageAcceleratorState returns the StorageAcceleratorState field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetStorageAcceleratorState() string {
	if o == nil || o.StorageAcceleratorState == nil {
		var ret string
		return ret
	}
	return *o.StorageAcceleratorState
}

// GetStorageAcceleratorStateOk returns a tuple with the StorageAcceleratorState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetStorageAcceleratorStateOk() (*string, bool) {
	if o == nil || o.StorageAcceleratorState == nil {
		return nil, false
	}
	return o.StorageAcceleratorState, true
}

// HasStorageAcceleratorState returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasStorageAcceleratorState() bool {
	if o != nil && o.StorageAcceleratorState != nil {
		return true
	}

	return false
}

// SetStorageAcceleratorState gets a reference to the given string and assigns it to the StorageAcceleratorState field.
func (o *ManagedMachineDataV2) SetStorageAcceleratorState(v string) {
	o.StorageAcceleratorState = &v
}

// GetVirtualCenterId returns the VirtualCenterId field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetVirtualCenterId() string {
	if o == nil || o.VirtualCenterId == nil {
		var ret string
		return ret
	}
	return *o.VirtualCenterId
}

// GetVirtualCenterIdOk returns a tuple with the VirtualCenterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetVirtualCenterIdOk() (*string, bool) {
	if o == nil || o.VirtualCenterId == nil {
		return nil, false
	}
	return o.VirtualCenterId, true
}

// HasVirtualCenterId returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasVirtualCenterId() bool {
	if o != nil && o.VirtualCenterId != nil {
		return true
	}

	return false
}

// SetVirtualCenterId gets a reference to the given string and assigns it to the VirtualCenterId field.
func (o *ManagedMachineDataV2) SetVirtualCenterId(v string) {
	o.VirtualCenterId = &v
}

// GetVirtualDisks returns the VirtualDisks field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetVirtualDisks() []VirtualDiskData {
	if o == nil || o.VirtualDisks == nil {
		var ret []VirtualDiskData
		return ret
	}
	return *o.VirtualDisks
}

// GetVirtualDisksOk returns a tuple with the VirtualDisks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetVirtualDisksOk() (*[]VirtualDiskData, bool) {
	if o == nil || o.VirtualDisks == nil {
		return nil, false
	}
	return o.VirtualDisks, true
}

// HasVirtualDisks returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasVirtualDisks() bool {
	if o != nil && o.VirtualDisks != nil {
		return true
	}

	return false
}

// SetVirtualDisks gets a reference to the given []VirtualDiskData and assigns it to the VirtualDisks field.
func (o *ManagedMachineDataV2) SetVirtualDisks(v []VirtualDiskData) {
	o.VirtualDisks = &v
}

// GetVirtualMachinePowerState returns the VirtualMachinePowerState field value if set, zero value otherwise.
func (o *ManagedMachineDataV2) GetVirtualMachinePowerState() string {
	if o == nil || o.VirtualMachinePowerState == nil {
		var ret string
		return ret
	}
	return *o.VirtualMachinePowerState
}

// GetVirtualMachinePowerStateOk returns a tuple with the VirtualMachinePowerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManagedMachineDataV2) GetVirtualMachinePowerStateOk() (*string, bool) {
	if o == nil || o.VirtualMachinePowerState == nil {
		return nil, false
	}
	return o.VirtualMachinePowerState, true
}

// HasVirtualMachinePowerState returns a boolean if a field has been set.
func (o *ManagedMachineDataV2) HasVirtualMachinePowerState() bool {
	if o != nil && o.VirtualMachinePowerState != nil {
		return true
	}

	return false
}

// SetVirtualMachinePowerState gets a reference to the given string and assigns it to the VirtualMachinePowerState field.
func (o *ManagedMachineDataV2) SetVirtualMachinePowerState(v string) {
	o.VirtualMachinePowerState = &v
}

func (o ManagedMachineDataV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BaseVmId != nil {
		toSerialize["base_vm_id"] = o.BaseVmId
	}
	if o.BaseVmSnapshotId != nil {
		toSerialize["base_vm_snapshot_id"] = o.BaseVmSnapshotId
	}
	if o.CloneErrorMessage != nil {
		toSerialize["clone_error_message"] = o.CloneErrorMessage
	}
	if o.CloneErrorTime != nil {
		toSerialize["clone_error_time"] = o.CloneErrorTime
	}
	if o.CreateTime != nil {
		toSerialize["create_time"] = o.CreateTime
	}
	if o.DatastoreIds != nil {
		toSerialize["datastore_ids"] = o.DatastoreIds
	}
	if o.HostName != nil {
		toSerialize["host_name"] = o.HostName
	}
	if o.ImageManagementStreamId != nil {
		toSerialize["image_management_stream_id"] = o.ImageManagementStreamId
	}
	if o.ImageManagementTagId != nil {
		toSerialize["image_management_tag_id"] = o.ImageManagementTagId
	}
	if o.InHoldCustomization != nil {
		toSerialize["in_hold_customization"] = o.InHoldCustomization
	}
	if o.InMaintenanceMode != nil {
		toSerialize["in_maintenance_mode"] = o.InMaintenanceMode
	}
	if o.LastMaintenanceTime != nil {
		toSerialize["last_maintenance_time"] = o.LastMaintenanceTime
	}
	if o.LogoffPolicy != nil {
		toSerialize["logoff_policy"] = o.LogoffPolicy
	}
	if o.MemoryMb != nil {
		toSerialize["memory_mb"] = o.MemoryMb
	}
	if o.MissingInVcenter != nil {
		toSerialize["missing_in_vcenter"] = o.MissingInVcenter
	}
	if o.NetworkLabels != nil {
		toSerialize["network_labels"] = o.NetworkLabels
	}
	if o.Operation != nil {
		toSerialize["operation"] = o.Operation
	}
	if o.OperationState != nil {
		toSerialize["operation_state"] = o.OperationState
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.PendingBaseVmId != nil {
		toSerialize["pending_base_vm_id"] = o.PendingBaseVmId
	}
	if o.PendingBaseVmSnapshotId != nil {
		toSerialize["pending_base_vm_snapshot_id"] = o.PendingBaseVmSnapshotId
	}
	if o.PendingImageManagementStreamId != nil {
		toSerialize["pending_image_management_stream_id"] = o.PendingImageManagementStreamId
	}
	if o.PendingImageManagementTagId != nil {
		toSerialize["pending_image_management_tag_id"] = o.PendingImageManagementTagId
	}
	if o.StorageAcceleratorState != nil {
		toSerialize["storage_accelerator_state"] = o.StorageAcceleratorState
	}
	if o.VirtualCenterId != nil {
		toSerialize["virtual_center_id"] = o.VirtualCenterId
	}
	if o.VirtualDisks != nil {
		toSerialize["virtual_disks"] = o.VirtualDisks
	}
	if o.VirtualMachinePowerState != nil {
		toSerialize["virtual_machine_power_state"] = o.VirtualMachinePowerState
	}
	return json.Marshal(toSerialize)
}

type NullableManagedMachineDataV2 struct {
	value *ManagedMachineDataV2
	isSet bool
}

func (v NullableManagedMachineDataV2) Get() *ManagedMachineDataV2 {
	return v.value
}

func (v *NullableManagedMachineDataV2) Set(val *ManagedMachineDataV2) {
	v.value = val
	v.isSet = true
}

func (v NullableManagedMachineDataV2) IsSet() bool {
	return v.isSet
}

func (v *NullableManagedMachineDataV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagedMachineDataV2(val *ManagedMachineDataV2) *NullableManagedMachineDataV2 {
	return &NullableManagedMachineDataV2{value: val, isSet: true}
}

func (v NullableManagedMachineDataV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagedMachineDataV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}