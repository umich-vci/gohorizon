/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// RcxClientRegisterSpec Specification of the RCX client.
type RcxClientRegisterSpec struct {
	// IP address of the RCX client.
	IpAddress string `json:"ip_address,omitempty"`
	// The RCX client certificate name.
	Name string `json:"name"`
	// Thumbprints of the RCX client certificate.
	Thumbprints []CertificateThumbprint `json:"thumbprints"`
}
