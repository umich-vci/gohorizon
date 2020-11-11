/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// CertificateOverrideData Certificate details and type information, which can be used to override thumbprint details.
type CertificateOverrideData struct {
	// Virtual Center certificate
	Certificate string `json:"certificate,omitempty"`
	// Type of Certificate. * PEM: PEM encoded certificate type
	Type string `json:"type,omitempty"`
}
