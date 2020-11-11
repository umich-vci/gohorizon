/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// VcMonitorConnectionServerV2 Information about connection to Virtual Center from Connection Server.
type VcMonitorConnectionServerV2 struct {
	Certificate CertificateMonitorInfo `json:"certificate,omitempty"`
	// Unique ID of the Connection Server.
	Id string `json:"id"`
	// The timestamp in milliseconds when the last update was obtained. Measured as epoch time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Connection server host name or IP address.
	Name string `json:"name"`
	// Status of the Virtual Center Connection with respect to this Connection Server. * OK: The connection to Virtual Center server is working properly. * DOWN: The connection to Virtual Center server is down. * RECONNECTING: The connection to Virtual Center server was lost and is being reconnected to. * UNKNOWN: Connection state to Virtual Center server is unknown. * INVALID_CREDENTIALS: The supplied credentials cannot be used to authenticate to the Virtual Center server. * CANNOT_LOGIN: The connection server cannot login to the Virtual Center server. * NOT_YET_CONNECTED: Connection server has not yet connected to Virtual Center server.
	Status string `json:"status"`
	// Indicates if the thumbprints of the Virtual Center was accepted.
	ThumbprintAccepted bool `json:"thumbprint_accepted"`
}