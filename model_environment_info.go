/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// EnvironmentInfo Information related to Environment Settings.
type EnvironmentInfo struct {
	// The GUID of a group of Connection Servers sharing the same configuration.
	ClusterGuid string `json:"cluster_guid"`
	// The name of a group of Connection Servers sharing the same configuration.
	ClusterName string `json:"cluster_name"`
	// Indicates if FIPS mode is enabled.
	FipsModeEnabled bool `json:"fips_mode_enabled"`
	// Indicates the IP mode of the environment. * IPv4: The ip mode is IPv4. * IPv6: The ip mode is IPv6.
	IpMode string `json:"ip_mode"`
	// The name of the current pod in the Multi-DataCenter Horizon Pod, the value will be null when PodFederation is not initialized.
	LocalPodName string `json:"local_pod_name,omitempty"`
	// Represents the clusters time zone offset from UTC in seconds.
	TimezoneOffset int32 `json:"timezone_offset"`
}
