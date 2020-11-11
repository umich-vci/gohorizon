/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// ImageManagementTagCreateSpec Information related to image management tag.
type ImageManagementTagCreateSpec struct {
	// Additional details about image management tag.
	AdditionalDetails map[string]string `json:"additional_details,omitempty"`
	// Image management stream ID to which this tag belongs.
	ImStreamId string `json:"im_stream_id"`
	// Image management version ID to which this tag belongs.
	ImVersionId string `json:"im_version_id"`
	// Image management tag name. It is unique among all the tags of a stream.
	Name string `json:"name"`
}