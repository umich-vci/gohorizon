/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// ApplicationSupportedFileTypesData Information about the file types supported by the application.
type ApplicationSupportedFileTypesData struct {
	// Whether or not the file types supported by this application should be allowed to automatically update to reflect changes reported by the agent. Typically this should be set to false if the application has manually configured supported file types. Default is true.
	EnableAutoUpdateFileTypes bool `json:"enable_auto_update_file_types"`
	// Whether or not the other file types supported by this application should be allowed to automatically update to reflect changes reported by the agent. Typically this should be set to false if the application has manually configured supported file types.
	EnableAutoUpdateOtherFileTypes bool `json:"enable_auto_update_other_file_types"`
	// Set of file types reported by the application as supported (if this application is discovered) or as specified by the administrator (if this application is manually configured). If unset, this application does not present any file type support.
	FileTypes []ApplicationFileTypeData `json:"file_types,omitempty"`
	// This represents the different file types reported by Application that can be passed from agent to client via broker or as specified by the administrator (if this application is manually configured). If unset, this application does not present any other file type support.
	OtherFileTypes []ApplicationOtherFileTypeData `json:"other_file_types,omitempty"`
}
