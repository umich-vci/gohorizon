/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DesktopPoolViewStorageAcceleratorSettings Settings related to the View Storage Accelerator feature.
type DesktopPoolViewStorageAcceleratorSettings struct {
	// A list of blackout times. Storage accelerator regeneration and machine disk space reclamation do not occur during blackout times. The same blackout policy applies to both operations. If unset, no blackout times are used.
	BlackoutTimes []ViewStorageAcceleratorBlackoutTimeSettings `json:"blackout_times,omitempty"`
	// How often to regenerate the View Storage Accelerator cache. Measured in Days.
	RegenerateViewStorageAcceleratorDays int32 `json:"regenerate_view_storage_accelerator_days,omitempty"`
	// Whether to use View Storage Accelerator.
	UseViewStorageAccelerator bool `json:"use_view_storage_accelerator,omitempty"`
	// Disk types to enable for the View Storage Accelerator feature. This is only applicable to linked clone desktop pools. * OS_DISKS: OS disks. * OS_AND_PERSISTENT_DISKS: OS and persistent disks.
	ViewStorageAcceleratorDiskTypes string `json:"view_storage_accelerator_disk_types,omitempty"`
}
