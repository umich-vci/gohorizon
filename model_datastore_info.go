/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DatastoreInfo Information related to datastore.
type DatastoreInfo struct {
	// Maximum capacity of this datastore, in MB.
	CapacityMb int64 `json:"capacity_mb,omitempty"`
	// Datacenter id for this datastore.
	DatacenterId string `json:"datacenter_id,omitempty"`
	// Disk type of the datastore. * SSD: Solid State Drive disk type. * NON_SSD: NON-Solid State Drive disk type. * UNKNOWN: Unknown disk type. * NON_VMFS: NON-VMFS disk type.
	DiskType string `json:"disk_type,omitempty"`
	// File system type of the datastore. * VMFS: Virtual Machine File System. * NFS: Network File System. * VSAN: vSAN File System. * VVOL: Virtual Volumes.
	FileSystemType string `json:"file_system_type,omitempty"`
	// Available capacity of this datastore, in MB.
	FreeSpaceMb int64 `json:"free_space_mb,omitempty"`
	// Host or Cluster id for this datastore.
	HostOrClusterId string `json:"host_or_cluster_id,omitempty"`
	// Unique ID representing the datastore.
	Id string `json:"id,omitempty"`
	// Reasons that may preclude this Datastore from being used in desktop pool/farm.
	IncompatibleReasons []string `json:"incompatible_reasons,omitempty"`
	// Indicates if this datastore is local to a single host.
	LocalDatastore bool `json:"local_datastore,omitempty"`
	// Datastore name.
	Name string `json:"name,omitempty"`
	// Indicates the number of virtual machines the datastore has for desktop pool/farm when applicable
	NumberOfVms int32 `json:"number_of_vms,omitempty"`
	// Datastore path.
	Path string `json:"path,omitempty"`
	// Virtual Center id for this datastore.
	VcenterId string `json:"vcenter_id,omitempty"`
	// The VMFS major version number.
	VmfsMajorVersion string `json:"vmfs_major_version,omitempty"`
}
