/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DesktopPoolNetworkInterfaceCardSettings Network interface card settings for machines provisioned for this desktop pool.
type DesktopPoolNetworkInterfaceCardSettings struct {
	// The network interface card id for these settings.
	NetworkInterfaceCardId string `json:"network_interface_card_id,omitempty"`
	// The network interface card name.
	NetworkInterfaceCardName string `json:"network_interface_card_name,omitempty"`
	// Automatic network label assignment feature settings for this NIC. By default, newly provisioned machines of an automated desktop pool retain their parent image's network labels on each of their network interface cards. In certain circumstances, notably dealing with VLAN subset sizing and DHCP IP address availability, it may be desirable for the desktop pool to instead use different network labels for these newly provisioned machines. This feature allows an administrator to provide a per NIC list of network labels and their maximum availability to be automatically distributed to newly provisioned machines. <br> If this is unset, the feature is disabled.<br> Starting at the alphabetically first network label specification in the list that has not yet been assigned its maximum count for this NIC on this desktop pool, the desktop pool will have its next provisioned machine's NIC assigned that label. If all network labels in this list have reached their maximum count, this desktop pool will have further provisioned machines assigned the last label in the list over capacity, and an error will be logged. Not all labels need be configured. <br>
	NetworkLabelAssignmentSpecs []NetworkLabelAssignmentSettings `json:"network_label_assignment_specs,omitempty"`
}
