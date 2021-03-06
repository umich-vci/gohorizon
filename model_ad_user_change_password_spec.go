/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// AdUserChangePasswordSpec Specification for changing AD user's password.
type AdUserChangePasswordSpec struct {
	// The domain of user. Note that domain is optional if UPN is supplied.
	Domain string `json:"domain,omitempty"`
	// The keyId of the cluster's SSO KeyPair used to encrypt the password key.
	KeyId string `json:"key_id"`
	// New password for the user in encrypted format.
	NewEncryptedPassword string `json:"new_encrypted_password"`
	// Old password for the user in encrypted format.
	OldEncryptedPassword string `json:"old_encrypted_password"`
	// Decryption key for the password. This key is itself encrypted with cluster's SSO keypair.
	ProtectedPasswordKey string `json:"protected_password_key"`
	// The username or UPN.
	Username string `json:"username"`
}
