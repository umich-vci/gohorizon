/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// CertificateServerDetails struct for CertificateServerDetails
type CertificateServerDetails struct {
	// Certificate Server name.
	Name string `json:"name"`
	// Certificate Server status. * OK: The state of the certificate server is OK as reported by the enrollment servers. * WARN: At least one enrollment server reports a warning on this certificate server. * ERROR: At least one enrollment server reports an error on this certificate server.
	Status string `json:"status"`
}
