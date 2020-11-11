/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// CustomizationSpecInfo Information related to customization specification created in vCenter.
type CustomizationSpecInfo struct {
	// Customization specification description.
	Description string `json:"description,omitempty"`
	// Guest Operating system. * UNKNOWN: Unknown * WINDOWS: Windows * LINUX: Linux
	GuestOs string `json:"guest_os,omitempty"`
	// Unique ID representing the customization specification.
	Id string `json:"id,omitempty"`
	// Reasons that may preclude this customization specification from being used in desktop pool creation.
	IncompatibleReasons []string `json:"incompatible_reasons,omitempty"`
	// Name of the customization specification.
	Name string `json:"name,omitempty"`
}
