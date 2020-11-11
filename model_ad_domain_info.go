/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// AdDomainInfo Information related to AD Domains of the environment.
type AdDomainInfo struct {
	// DNS name of the AD Domain.
	DnsName string `json:"dns_name"`
	// Unique ID representing AD Domain.
	Id string `json:"id"`
	// NetBIOS name of the AD Domain.
	NetbiosName string `json:"netbios_name"`
}