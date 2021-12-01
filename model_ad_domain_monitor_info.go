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

// ADDomainMonitorInfo Monitoring information related to AD Domain.
type ADDomainMonitorInfo struct {
	// Information about the AD Domain connections from each of the connection servers.
	ConnectionServers *[]ADDomainMonitorConnectionServer `json:"connection_servers,omitempty"`
	// The DNS name for the domain.
	DnsName *string `json:"dns_name,omitempty"`
	// The NetBIOS name for the domain.
	NetbiosName *string `json:"netbios_name,omitempty"`
	// If this is an NT4 domain or not.
	Nt4Domain *bool `json:"nt4_domain,omitempty"`
}

// NewADDomainMonitorInfo instantiates a new ADDomainMonitorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewADDomainMonitorInfo() *ADDomainMonitorInfo {
	this := ADDomainMonitorInfo{}
	return &this
}

// NewADDomainMonitorInfoWithDefaults instantiates a new ADDomainMonitorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewADDomainMonitorInfoWithDefaults() *ADDomainMonitorInfo {
	this := ADDomainMonitorInfo{}
	return &this
}

// GetConnectionServers returns the ConnectionServers field value if set, zero value otherwise.
func (o *ADDomainMonitorInfo) GetConnectionServers() []ADDomainMonitorConnectionServer {
	if o == nil || o.ConnectionServers == nil {
		var ret []ADDomainMonitorConnectionServer
		return ret
	}
	return *o.ConnectionServers
}

// GetConnectionServersOk returns a tuple with the ConnectionServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorInfo) GetConnectionServersOk() (*[]ADDomainMonitorConnectionServer, bool) {
	if o == nil || o.ConnectionServers == nil {
		return nil, false
	}
	return o.ConnectionServers, true
}

// HasConnectionServers returns a boolean if a field has been set.
func (o *ADDomainMonitorInfo) HasConnectionServers() bool {
	if o != nil && o.ConnectionServers != nil {
		return true
	}

	return false
}

// SetConnectionServers gets a reference to the given []ADDomainMonitorConnectionServer and assigns it to the ConnectionServers field.
func (o *ADDomainMonitorInfo) SetConnectionServers(v []ADDomainMonitorConnectionServer) {
	o.ConnectionServers = &v
}

// GetDnsName returns the DnsName field value if set, zero value otherwise.
func (o *ADDomainMonitorInfo) GetDnsName() string {
	if o == nil || o.DnsName == nil {
		var ret string
		return ret
	}
	return *o.DnsName
}

// GetDnsNameOk returns a tuple with the DnsName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorInfo) GetDnsNameOk() (*string, bool) {
	if o == nil || o.DnsName == nil {
		return nil, false
	}
	return o.DnsName, true
}

// HasDnsName returns a boolean if a field has been set.
func (o *ADDomainMonitorInfo) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// SetDnsName gets a reference to the given string and assigns it to the DnsName field.
func (o *ADDomainMonitorInfo) SetDnsName(v string) {
	o.DnsName = &v
}

// GetNetbiosName returns the NetbiosName field value if set, zero value otherwise.
func (o *ADDomainMonitorInfo) GetNetbiosName() string {
	if o == nil || o.NetbiosName == nil {
		var ret string
		return ret
	}
	return *o.NetbiosName
}

// GetNetbiosNameOk returns a tuple with the NetbiosName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorInfo) GetNetbiosNameOk() (*string, bool) {
	if o == nil || o.NetbiosName == nil {
		return nil, false
	}
	return o.NetbiosName, true
}

// HasNetbiosName returns a boolean if a field has been set.
func (o *ADDomainMonitorInfo) HasNetbiosName() bool {
	if o != nil && o.NetbiosName != nil {
		return true
	}

	return false
}

// SetNetbiosName gets a reference to the given string and assigns it to the NetbiosName field.
func (o *ADDomainMonitorInfo) SetNetbiosName(v string) {
	o.NetbiosName = &v
}

// GetNt4Domain returns the Nt4Domain field value if set, zero value otherwise.
func (o *ADDomainMonitorInfo) GetNt4Domain() bool {
	if o == nil || o.Nt4Domain == nil {
		var ret bool
		return ret
	}
	return *o.Nt4Domain
}

// GetNt4DomainOk returns a tuple with the Nt4Domain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorInfo) GetNt4DomainOk() (*bool, bool) {
	if o == nil || o.Nt4Domain == nil {
		return nil, false
	}
	return o.Nt4Domain, true
}

// HasNt4Domain returns a boolean if a field has been set.
func (o *ADDomainMonitorInfo) HasNt4Domain() bool {
	if o != nil && o.Nt4Domain != nil {
		return true
	}

	return false
}

// SetNt4Domain gets a reference to the given bool and assigns it to the Nt4Domain field.
func (o *ADDomainMonitorInfo) SetNt4Domain(v bool) {
	o.Nt4Domain = &v
}

func (o ADDomainMonitorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionServers != nil {
		toSerialize["connection_servers"] = o.ConnectionServers
	}
	if o.DnsName != nil {
		toSerialize["dns_name"] = o.DnsName
	}
	if o.NetbiosName != nil {
		toSerialize["netbios_name"] = o.NetbiosName
	}
	if o.Nt4Domain != nil {
		toSerialize["nt4_domain"] = o.Nt4Domain
	}
	return json.Marshal(toSerialize)
}

type NullableADDomainMonitorInfo struct {
	value *ADDomainMonitorInfo
	isSet bool
}

func (v NullableADDomainMonitorInfo) Get() *ADDomainMonitorInfo {
	return v.value
}

func (v *NullableADDomainMonitorInfo) Set(val *ADDomainMonitorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableADDomainMonitorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableADDomainMonitorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableADDomainMonitorInfo(val *ADDomainMonitorInfo) *NullableADDomainMonitorInfo {
	return &NullableADDomainMonitorInfo{value: val, isSet: true}
}

func (v NullableADDomainMonitorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableADDomainMonitorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
