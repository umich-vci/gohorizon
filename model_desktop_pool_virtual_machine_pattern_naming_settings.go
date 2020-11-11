/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DesktopPoolVirtualMachinePatternNamingSettings Settings related to specification of VMs with a naming pattern.
type DesktopPoolVirtualMachinePatternNamingSettings struct {
	// Maximum number of machines in the desktop pool.
	MaxNumberOfMachines int32 `json:"max_number_of_machines,omitempty"`
	// The minimum number of machines to have provisioned if on demand provisioning is selected.
	MinNumberOfMachines int32 `json:"min_number_of_machines,omitempty"`
	// Virtual machines will be named according to the specified naming pattern.
	NamingPattern string `json:"naming_pattern,omitempty"`
	// Number of spare powered on machines.
	NumberOfSpareMachines int32 `json:"number_of_spare_machines,omitempty"`
	// Determines when the machines are provisioned. * ON_DEMAND: Provision machines on demand. * UP_FRONT: Provision all machines up-front.
	ProvisioningTime string `json:"provisioning_time,omitempty"`
}
