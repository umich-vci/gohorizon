/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// ImageManagementVersionInfo Information related to image management version.
type ImageManagementVersionInfo struct {
	// Additional details about image management version.
	AdditionalDetails map[string]string `json:"additional_details,omitempty"`
	// Image management version description.
	Description string `json:"description,omitempty"`
	// Unique ID representing image management version.
	Id string `json:"id"`
	// Image management stream ID
	ImStreamId string `json:"im_stream_id,omitempty"`
	// Image management version name.
	Name string `json:"name"`
	// Image management version status. * AVAILABLE: Image management version is available for desktop pools/farms to be created. * DEPLOYING_VM: Image management version is deploying VM on the selected pod. * DEPLOYMENT_DONE: Image management version status when VM deployment is done for the selected pod. * DELETED: Image management version has been deleted. * DISABLED: Image management version has been disabled and no further pool operation can be done using the same. * FAILED: Image management version creation has failed. * PARTIALLY_AVAILABLE: Some of the image management asset creation in some of the virtual centers have failed. * PUBLISHING: Image management version is being published and specialized internally like installing agents etc. * REPLICATING: Copying the specialized images across all virtual centers.
	Status string `json:"status"`
}
