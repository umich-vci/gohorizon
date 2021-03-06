/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// SecuritySettingsUpdateSpec Information related to Security Settings.
type SecuritySettingsUpdateSpec struct {
	// The SHA-256 hash of the (UTF-8) data recovery password.
	DataRecoveryPasswordHash []string `json:"data_recovery_password_hash,omitempty"`
	// The data recovery password hint. This property has a maximum length of 128 characters.
	DataRecoveryPasswordHint string `json:"data_recovery_password_hint,omitempty"`
	// Determines if signing and verification of the JMS messages passed between Horizon components takes place. * DISABLED: Message security mode is disabled. * MIXED: Message security mode is enabled but not enforced. * ENABLED: Message security mode is enabled. Unsigned messages are rejected by Horizon components. * ENHANCED: Message Security mode is Enhanced. Message signing and validation is performed based on the current Security Level and desktop Message Security mode.
	MessageSecurityMode string `json:"message_security_mode"`
	// Determines if user credentials must be re-authenticated after a network interruption when Horizon clients use secure tunnel connections to Horizon resources. When you select this setting, if a secure tunnel connection ends during a session, Horizon Client requires the user to re-authenticate before reconnecting.
	ReAuthSecureTunnelAfterInterruption bool `json:"re_auth_secure_tunnel_after_interruption,omitempty"`
}
