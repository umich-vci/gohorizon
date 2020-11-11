/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// HostOrClusterInfo Information related to host or cluster.
type HostOrClusterInfo struct {
	Container HostOrClusterContainer `json:"container,omitempty"`
	Details   HostOrClusterDetails   `json:"details,omitempty"`
	// Unique ID representing a host or cluster.
	Id string `json:"id"`
}