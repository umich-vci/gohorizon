/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2111
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// DesktopPoolNetworkInterfaceCardSettingsUpdateSpec Applicable To: Instant Clone desktop pool. <br>Network interface card settings for machines provisioned for Instant Clone desktop pool.
type DesktopPoolNetworkInterfaceCardSettingsUpdateSpec struct {
	// The network interface card id for these settings.
	NetworkInterfaceCardId string `json:"network_interface_card_id"`
	// Automatic network label assignment feature settings for this NIC. By default, newly provisioned machines of an automated desktop pool retain their parent image's network labels on each of their network interface cards. In certain circumstances, notably dealing with VLAN subset sizing and DHCP IP address availability, it may be desirable for the desktop pool to instead use different network labels for these newly provisioned machines. This feature allows an administrator to provide a per NIC list of network labels and their maximum availability to be automatically distributed to newly provisioned machines. <br> If this is unset, the feature is disabled.<br> Starting at the alphabetically first network label specification in the list that has not yet been assigned its maximum count for this NIC on this desktop pool, the desktop pool will have its next provisioned machine's NIC assigned that label. If all network labels in this list have reached their maximum count, this desktop pool will have further provisioned machines assigned the last label in the list over capacity, and an error will be logged. Not all labels need be configured. <br>
	NetworkLabelAssignmentSpecs *[]NetworkLabelAssignmentSettingsUpdateSpec `json:"network_label_assignment_specs,omitempty"`
}

// NewDesktopPoolNetworkInterfaceCardSettingsUpdateSpec instantiates a new DesktopPoolNetworkInterfaceCardSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolNetworkInterfaceCardSettingsUpdateSpec(networkInterfaceCardId string) *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec {
	this := DesktopPoolNetworkInterfaceCardSettingsUpdateSpec{}
	this.NetworkInterfaceCardId = networkInterfaceCardId
	return &this
}

// NewDesktopPoolNetworkInterfaceCardSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolNetworkInterfaceCardSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolNetworkInterfaceCardSettingsUpdateSpecWithDefaults() *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec {
	this := DesktopPoolNetworkInterfaceCardSettingsUpdateSpec{}
	return &this
}

// GetNetworkInterfaceCardId returns the NetworkInterfaceCardId field value
func (o *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) GetNetworkInterfaceCardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NetworkInterfaceCardId
}

// GetNetworkInterfaceCardIdOk returns a tuple with the NetworkInterfaceCardId field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) GetNetworkInterfaceCardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkInterfaceCardId, true
}

// SetNetworkInterfaceCardId sets field value
func (o *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) SetNetworkInterfaceCardId(v string) {
	o.NetworkInterfaceCardId = v
}

// GetNetworkLabelAssignmentSpecs returns the NetworkLabelAssignmentSpecs field value if set, zero value otherwise.
func (o *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) GetNetworkLabelAssignmentSpecs() []NetworkLabelAssignmentSettingsUpdateSpec {
	if o == nil || o.NetworkLabelAssignmentSpecs == nil {
		var ret []NetworkLabelAssignmentSettingsUpdateSpec
		return ret
	}
	return *o.NetworkLabelAssignmentSpecs
}

// GetNetworkLabelAssignmentSpecsOk returns a tuple with the NetworkLabelAssignmentSpecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) GetNetworkLabelAssignmentSpecsOk() (*[]NetworkLabelAssignmentSettingsUpdateSpec, bool) {
	if o == nil || o.NetworkLabelAssignmentSpecs == nil {
		return nil, false
	}
	return o.NetworkLabelAssignmentSpecs, true
}

// HasNetworkLabelAssignmentSpecs returns a boolean if a field has been set.
func (o *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) HasNetworkLabelAssignmentSpecs() bool {
	if o != nil && o.NetworkLabelAssignmentSpecs != nil {
		return true
	}

	return false
}

// SetNetworkLabelAssignmentSpecs gets a reference to the given []NetworkLabelAssignmentSettingsUpdateSpec and assigns it to the NetworkLabelAssignmentSpecs field.
func (o *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) SetNetworkLabelAssignmentSpecs(v []NetworkLabelAssignmentSettingsUpdateSpec) {
	o.NetworkLabelAssignmentSpecs = &v
}

func (o DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["network_interface_card_id"] = o.NetworkInterfaceCardId
	}
	if o.NetworkLabelAssignmentSpecs != nil {
		toSerialize["network_label_assignment_specs"] = o.NetworkLabelAssignmentSpecs
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec struct {
	value *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec
	isSet bool
}

func (v NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec) Get() *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec {
	return v.value
}

func (v *NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec) Set(val *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec(val *DesktopPoolNetworkInterfaceCardSettingsUpdateSpec) *NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec {
	return &NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolNetworkInterfaceCardSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
