/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// VcLimits Information about the limits configured for Virtual Center
type VcLimits struct {
	// Maximum concurrent instant clone engine provisioning operations. This property has a default value of 20. This property has a minimum value of 1.
	InstantCloneEngineProvisioningLimit int32 `json:"instant_clone_engine_provisioning_limit"`
	// Maximum concurrent virtual center power operations. This property has a default value of 50. This property has a minimum value of 1.
	PowerOperationsLimit int32 `json:"power_operations_limit"`
	// Maximum concurrent virtual center provisioning operations. This property has a default value of 20. This property has a minimum value of 1.
	ProvisioningLimit int32 `json:"provisioning_limit"`
}
