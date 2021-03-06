/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// AdUserInfo Information related to AD User.
type AdUserInfo struct {
	// DNS name of the domain in which this user or group belongs to.
	Domain string `json:"domain,omitempty"`
	// Guids of the user's groups in RFC 4122 format.
	GroupGuids []string `json:"group_guids,omitempty"`
	// List of unique SIDs of the groups, this user or group belongs to.
	GroupSids []string `json:"group_sids,omitempty"`
	// GUID of the user in RFC 4122 format.
	UserGuid string `json:"user_guid"`
	// User Principal name(UPN) of this user.
	UserPrincipalName string `json:"user_principal_name,omitempty"`
	// Unique SID representing this AD User.
	UserSid string `json:"user_sid"`
	// Username of this user.
	Username string `json:"username,omitempty"`
}
