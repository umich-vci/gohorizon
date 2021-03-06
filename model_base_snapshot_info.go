/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// BaseSnapshotInfo Information related to VM snapshot.
type BaseSnapshotInfo struct {
	// Epoch time in milli seconds, when the VM snapshot was created.
	CreatedTimestamp int64 `json:"created_timestamp,omitempty"`
	// Description of the VM snapshot.
	Description string `json:"description,omitempty"`
	// Sum of capacities of all the virtual disks in the VM snapshot, in MB.
	DiskSizeMb int64 `json:"disk_size_mb,omitempty"`
	// VM snapshot hardware version
	HardwareVersion int32 `json:"hardware_version,omitempty"`
	// Unique ID representing the VM snapshot.
	Id string `json:"id,omitempty"`
	// Reasons that may preclude this VM snapshot from being used in linked/instant clone desktop pool or farm creation.
	IncompatibleReasons []string `json:"incompatible_reasons,omitempty"`
	// Maximum number of monitors set in SVGA settings for the VM snapshot in vCenter.
	MaxNumberOfMonitors int32 `json:"max_number_of_monitors,omitempty"`
	// Maximum resolution of any one monitor set in SVGA settings for the VM snapshot in vCenter.
	MaxResolutionOfAnyOneMonitor string `json:"max_resolution_of_any_one_monitor,omitempty"`
	// The physical memory size of VM snapshot, in MB
	MemoryMb int32 `json:"memory_mb,omitempty"`
	// Amount of memory that is guaranteed available to the virtual machine, in MB.
	MemoryReservationMb int64 `json:"memory_reservation_mb,omitempty"`
	// VM snapshot name.
	Name string `json:"name,omitempty"`
	// VM snapshot path.
	Path string `json:"path,omitempty"`
	// Indicate how the virtual video device for the VM snapshot renders 3D graphics. Will be set only if VM snapshot supports 3D functions. * MANAGE_BY_VSPHERE_CLIENT: 3D rendering managed by vSphere Client. * AUTOMATIC: 3D rendering is automatic. * SOFTWARE: 3D rendering is software dependent. The software renderer is supported (at minimum) on virtual hardware version 8 in a vSphere 5.0 environment. * HARDWARE: 3D rendering is hardware dependent. The hardware-based renderer is supported (at minimum) on virtual hardware version 9 in a vSphere 5.1 environment. * DISABLED: 3D rendering is disabled.
	Renderer3d string `json:"renderer3d,omitempty"`
	// Total video memory in MB set in SVGA settings for the VM snapshot in vCenter.
	TotalVideoMemoryMb float64 `json:"total_video_memory_mb,omitempty"`
	// Virtual Center id for this VM snapshot.
	VcenterId string `json:"vcenter_id,omitempty"`
	// NVIDIA GRID vGPU type configured on this VM snapshot.
	VgpuType string `json:"vgpu_type,omitempty"`
}
