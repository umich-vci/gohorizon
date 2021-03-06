/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DatastorePathInfo Information related to datastore paths.
type DatastorePathInfo struct {
	// Unique ID representing the datastore path.
	Id string `json:"id,omitempty"`
	// Datastore path.
	Path string `json:"path,omitempty"`
}
