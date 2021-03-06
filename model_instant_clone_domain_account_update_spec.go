/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// InstantCloneDomainAccountUpdateSpec Update specification for the instant clone domain account.
type InstantCloneDomainAccountUpdateSpec struct {
	// Password of the account.
	Password []string `json:"password"`
}
