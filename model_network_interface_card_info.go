/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// NetworkInterfaceCardInfo Information related to network interface card.
type NetworkInterfaceCardInfo struct {
	// Unique ID representing the network interface card.
	Id string `json:"id,omitempty"`
	// Network interface card MAC address.
	MacAddress string `json:"mac_address,omitempty"`
	// Network interface card name.
	Name string `json:"name,omitempty"`
}