/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2106
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// PhysicalMachineRegisterSpec Information required to register a physical machine.
type PhysicalMachineRegisterSpec struct {
	// An optional string to describe how and why this physical machine was registered.
	Description *string `json:"description,omitempty"`
	// The DNS name for the physical machine.
	DnsName string `json:"dns_name"`
	// The Operating System of the physical machine. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS)
	OperatingSystem string `json:"operating_system"`
}

// NewPhysicalMachineRegisterSpec instantiates a new PhysicalMachineRegisterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhysicalMachineRegisterSpec(dnsName string, operatingSystem string) *PhysicalMachineRegisterSpec {
	this := PhysicalMachineRegisterSpec{}
	this.DnsName = dnsName
	this.OperatingSystem = operatingSystem
	return &this
}

// NewPhysicalMachineRegisterSpecWithDefaults instantiates a new PhysicalMachineRegisterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhysicalMachineRegisterSpecWithDefaults() *PhysicalMachineRegisterSpec {
	this := PhysicalMachineRegisterSpec{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PhysicalMachineRegisterSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhysicalMachineRegisterSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PhysicalMachineRegisterSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PhysicalMachineRegisterSpec) SetDescription(v string) {
	o.Description = &v
}

// GetDnsName returns the DnsName field value
func (o *PhysicalMachineRegisterSpec) GetDnsName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value
// and a boolean to check if the value has been set.
func (o *PhysicalMachineRegisterSpec) GetDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DnsName, true
}

// SetDnsName sets field value
func (o *PhysicalMachineRegisterSpec) SetDnsName(v string) {
	o.DnsName = v
}

// GetOperatingSystem returns the OperatingSystem field value
func (o *PhysicalMachineRegisterSpec) GetOperatingSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OperatingSystem
}

// GetOperatingSystemOk returns a tuple with the OperatingSystem field value
// and a boolean to check if the value has been set.
func (o *PhysicalMachineRegisterSpec) GetOperatingSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingSystem, true
}

// SetOperatingSystem sets field value
func (o *PhysicalMachineRegisterSpec) SetOperatingSystem(v string) {
	o.OperatingSystem = v
}

func (o PhysicalMachineRegisterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["dns_name"] = o.DnsName
	}
	if true {
		toSerialize["operating_system"] = o.OperatingSystem
	}
	return json.Marshal(toSerialize)
}

type NullablePhysicalMachineRegisterSpec struct {
	value *PhysicalMachineRegisterSpec
	isSet bool
}

func (v NullablePhysicalMachineRegisterSpec) Get() *PhysicalMachineRegisterSpec {
	return v.value
}

func (v *NullablePhysicalMachineRegisterSpec) Set(val *PhysicalMachineRegisterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullablePhysicalMachineRegisterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullablePhysicalMachineRegisterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhysicalMachineRegisterSpec(val *PhysicalMachineRegisterSpec) *NullablePhysicalMachineRegisterSpec {
	return &NullablePhysicalMachineRegisterSpec{value: val, isSet: true}
}

func (v NullablePhysicalMachineRegisterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhysicalMachineRegisterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
