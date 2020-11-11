/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// VirtualCenterMonitorInfo Monitoring information related to a Virtual Center
type VirtualCenterMonitorInfo struct {
	// Information about the Virtual Center connections from each of the connection servers.
	ConnectionServers []VcMonitorConnectionServer `json:"connection_servers"`
	// Information about the datastores of the Virtual Center.
	Datastores []VcMonitorDatastore `json:"datastores"`
	// Number of Desktop Pools And Farms managed by the virtual center.
	DesktopsCount int32            `json:"desktops_count"`
	Details       VcMonitorDetails `json:"details"`
	// Information about the ESX hosts added in the Virtual Center.
	Hosts []VcMonitorHost `json:"hosts"`
	// Unique ID of the Virtual Center.
	Id string `json:"id"`
	// Virtual Center server name or IP address.
	Name string `json:"name"`
}