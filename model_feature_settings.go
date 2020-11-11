/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// FeatureSettings Information related to Feature Settings.
type FeatureSettings struct {
	// Indicates whether this cluster/pod is managed by Horizon Cloud Services. This will be false only when there are no cloud managed machines.
	CloudManaged bool `json:"cloud_managed,omitempty"`
	// Determines whether Helpdesk feature is enabled or not. By default Helpdesk is enabled.
	EnableHelpdesk bool `json:"enable_helpdesk,omitempty"`
	// Determines whether Image Management feature is enabled or not. By default Image Management is disabled.
	EnableImageManagement bool `json:"enable_image_management,omitempty"`
}
