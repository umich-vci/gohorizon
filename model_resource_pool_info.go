/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// ResourcePoolInfo Information related to resource pool.
type ResourcePoolInfo struct {
	// Child nodes of the resource pool.
	Children []ResourcePoolInfo `json:"children,omitempty"`
	// Unique ID representing the resource pool.
	Id string `json:"id,omitempty"`
	// Resource pool name.
	Name string `json:"name,omitempty"`
	// Resource pool path.
	Path string `json:"path,omitempty"`
	// Resource pool type. * HOST: Host used as a resource pool suitable for use in desktop pool/farm. * CLUSTER: Cluster used as a resource pool suitable for use in desktop pool/farm. * RESOURCE_POOL: Regular resource pool suitable for use in desktop pool/farm. * OTHER: Other resource type which cannot be used in desktop pool/farm.
	Type string `json:"type,omitempty"`
}
