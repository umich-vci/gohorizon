/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// ConnectionServerMonitorCsReplication Replication status with respect to Peer Connection Servers in the same cluster.
type ConnectionServerMonitorCsReplication struct {
	// Connection Server host name or IP address.
	ServerName string `json:"server_name"`
	// LDAP replication status. * OK: The connection to the Connection Server is working properly. * ERROR: There is a problem with LDAP replication to the Connection Server.
	Status string `json:"status"`
}
