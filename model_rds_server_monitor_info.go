/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// RdsServerMonitorInfo RDS Server Monitor Information.
type RdsServerMonitorInfo struct {
	Details RdsServerMonitorDetails `json:"details"`
	// Indicates if RDS server is enabled.
	Enabled bool `json:"enabled"`
	// Indicates the Farm ID to which the RDS Server belongs to.
	FarmId string `json:"farm_id"`
	// Unique ID of the RDS server.
	Id string `json:"id"`
	// This value is similar to load_preference and represents the load on RDS Server in the range of 0 to 100.
	LoadIndex int32 `json:"load_index,omitempty"`
	// Based on the current load of this RDS Server, gives a measure of how preferential this server is to be chosen for new application sessions. * BLOCK: RDS Server is blocked and new sessions will not be accepted. * HEAVY: RDS Server is experiencing heavy load and should likely not be chosen for new sessions. * NORMAL: RDS Server is experiencing normal load and is okay to be chosen for new sessions. * LIGHT: RDS Server is experiencing light load and is okay to be chosen for new sessions. * UNKNOWN: RDS Server did not report a load preference. This is potentially a configuration issue if other RDS Servers in the same Farm do report load preferences.
	LoadPreference string `json:"load_preference,omitempty"`
	// RDS Server name.
	Name string `json:"name"`
	// RDS server session count.
	SessionCount int32 `json:"session_count,omitempty"`
	// RDS server status. * OK: RDS Server is reachable. All applications (defined on its farm) are verified installed on the RDS Server. * WARNING: RDS Server is reachable. Some applications are detected as not installed on the RDS Server. * ERROR: RDS Server is unreachable, or, none of the applications are installed. * DISABLED: RDS Server is disabled.
	Status string `json:"status"`
}
