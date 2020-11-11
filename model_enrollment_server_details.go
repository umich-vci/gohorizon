/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// EnrollmentServerDetails struct for EnrollmentServerDetails
type EnrollmentServerDetails struct {
	// Enrollment server dns name.
	DnsName string `json:"dns_name"`
	// Unique ID of the Enrollment Server.
	Id string `json:"id"`
	// Enrollment server status. * OK: The state of enrollment server is OK. * ERROR: The enrollment server has an error.
	Status string `json:"status"`
}
