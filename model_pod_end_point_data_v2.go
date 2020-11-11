/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// PodEndPointDataV2 struct for PodEndPointDataV2
type PodEndPointDataV2 struct {
	// Indicates whether an endpoint is enabled. A disabled endpoint will be excluded from participating inter-pod communication.
	Enabled bool `json:"enabled"`
	// Unique ID for a pod endpoint.
	Id string `json:"id"`
	// The timestamp in milliseconds when the last update was obtained. Measured as epoch time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Name for the pod endpoint.
	Name string `json:"name"`
	// Round trip time (in milliseconds) for ping request between the local pod endpoint and the remote pod.
	RoundtripTime int64 `json:"roundtrip_time,omitempty"`
	// Status of the pod endpoint. * ONLINE: Pod endpoint is online and functional. * UNCHECKED: Pod endpoint was offline and it just came back online but the system has not verified that it is functional. * OFFLINE: Pod endpoint is offline or unreachable.
	Status string `json:"status"`
	// The URL for the pod endpoint. This address and special port will be used for inter-pod communication.
	Url string `json:"url"`
}
