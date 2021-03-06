/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// InstalledApplicationInfo Information about the application installed on RDS server/machine of a farm/desktop pool.
type InstalledApplicationInfo struct {
	// Path to application executable.
	ExecutablePath string `json:"executable_path,omitempty"`
	// Set of file types reported by the application as supported. If unset, this application does not present any file type support.
	FileTypes []ApplicationFileTypeData `json:"file_types,omitempty"`
	// Application name information, as sent by RDSServer/machine during application discovery.
	Name string `json:"name,omitempty"`
	// This represents the different file types reported by Application that can be passed from horizon agent to horizon client via connection server. If unset, this application does not present any other file type support.
	OtherFileTypes []ApplicationOtherFileTypeData `json:"other_file_types,omitempty"`
	// Application publisher
	Publisher string `json:"publisher,omitempty"`
	// Application version.
	Version string `json:"version,omitempty"`
}
