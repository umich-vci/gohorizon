/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DatacenterInfo Information related to datacenter.
type DatacenterInfo struct {
	// Unique ID representing a datacenter.
	Id string `json:"id"`
	// Name of the datacenter.
	Name string `json:"name"`
	// Datacenter path.
	Path string `json:"path"`
}
