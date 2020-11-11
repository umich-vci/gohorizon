/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// EntitlementResponseInfo Entitlement response info object corrspsonding to the given ad-user-or-group SID for the entitlement operation.
type EntitlementResponseInfo struct {
	// Unique SID representing the ad-user-or-group
	AdUserOrGroupId string `json:"ad_user_or_group_id,omitempty"`
	// Reasons for the failure of the operation.
	ErrorMessages []string `json:"error_messages,omitempty"`
	// Response HTTP status code of the operation.
	StatusCode int32 `json:"status_code,omitempty"`
	// Timestamp in milliseconds when the operation failed. Measured as epoch time.
	Timestamp int64 `json:"timestamp,omitempty"`
}
