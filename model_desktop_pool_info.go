/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DesktopPoolInfo Information related to Desktop Pool.
type DesktopPoolInfo struct {
	// Description of the Desktop Pool. The maximum length is 1024 characters.
	Description string `json:"description,omitempty"`
	// Display name of the Desktop Pool. The maximum length is 256 characters.
	DisplayName string `json:"display_name,omitempty"`
	// Indicates whether the Desktop Pool is enabled for brokering.
	Enabled bool `json:"enabled"`
	// Unique ID representing Desktop Pool.
	Id string `json:"id"`
	// Name of the Desktop Pool. The maximum length is 64 characters.
	Name     string              `json:"name"`
	Settings DesktopPoolSettings `json:"settings"`
	// Source of the Machines in this Desktop Pool. * INSTANT_CLONE: The Desktop Pool uses instant clone technology for provisioning the machines.Applicable for AUTOMATED type desktop pools. * LINKED_CLONE: The Desktop Pool uses linked clone technology for provisioning the machines.Applicable for AUTOMATED type desktop pools. * VIRTUAL_CENTER: The Desktop Pool uses Virtual Center as source for provisioning the machines.Applicable for AUTOMATED and MANUAL type desktop pools. * RDS: The Desktop Pool is backed by Farm. The Farm used in this Desktop Pool can be of any Source. * UNMANAGED: The Desktop Pool holds the non-vCenter source machines that includes physical computers,blade PCs and non-vCenter servers. Applicable for MANUAL type desktop pools.
	Source string `json:"source"`
	// Type of the Desktop Pool. * AUTOMATED: Automated Desktop Pool. * MANUAL: Manual Desktop Pool. * RDS: RDS Desktop Pool.
	Type string `json:"type"`
}
