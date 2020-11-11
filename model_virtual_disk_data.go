/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// VirtualDiskData Information related to a virtual disk.
type VirtualDiskData struct {
	// The disk capacity in MB.
	CapacityMb int64 `json:"capacity_mb"`
	// The virtual disk's datastore.
	DatastorePath string `json:"datastore_path"`
	// The virtual disk path.
	Path string `json:"path"`
}
