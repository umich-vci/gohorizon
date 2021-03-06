/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// EntitlementSpec Specification for entitlement operations on a given resource id.
type EntitlementSpec struct {
	// List of ad-user-or-group SIDs for the entitlement operations on the given resource.
	AdUserOrGroupIds []string `json:"ad_user_or_group_ids,omitempty"`
	// Unique ID representing the resource.
	Id string `json:"id,omitempty"`
}
