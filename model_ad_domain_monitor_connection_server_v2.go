/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// AdDomainMonitorConnectionServerV2 Information about the AD Domain connection from connection server.
type AdDomainMonitorConnectionServerV2 struct {
	// Unique ID of the connection server.
	Id string `json:"id"`
	// The timestamp in milliseconds when the last update was obtained. Measured as epoch time.
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp,omitempty"`
	// Connection server host name or IP address.
	Name string `json:"name"`
	// Status of the connection to the domain. * UNCONTACTABLE: No domain controllers appear to be present on the network for this domain. * FULLY_ACCESSIBLE: The domain controller(s) are accepting bind operations. * CANNOT_BIND: The domain controller(s) are only accepting LDAP ping operations. * UNKNOWN: Cannot determine accessibility.
	Status string `json:"status"`
	// The trust relationship for the domain. * PRIMARY_DOMAIN: The domain is the domain that the broker is present in. * FROM_BROKER: The domain is trusted by the domain that the broker is in. * TO_BROKER: The domain trusts the broker's domain (this is for completeness and generally will not be used). * TWO_WAY: The domain has a two way trust relationship with the broker's domain. * TWO_WAY_FOREST: The domain is in the same forest as the broker's domain, implies two way trust. * MANUAL: The domain was manually configured (the trust has not been detected). * UNKNOWN: The trust relationship could not be determined.
	TrustRelationship string `json:"trust_relationship"`
}
