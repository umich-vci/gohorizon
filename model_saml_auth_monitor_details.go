/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// SamlAuthMonitorDetails Details of the SAML authenticator.
type SamlAuthMonitorDetails struct {
	// The administrator URL for the SAML authenticator.
	AdministratorUrl string `json:"administrator_url,omitempty"`
	// The label of the SAML Authenticator.
	Label string `json:"label"`
	// The metadata URL of the SAML Authenticator.
	MetadataUrl string `json:"metadata_url"`
}
