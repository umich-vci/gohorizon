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

// FarmProvisioningSettingsUpdateSpec Virtual center provisioning settings for the farm.
type FarmProvisioningSettingsUpdateSpec struct {
	// Host or cluster where the machines are deployed in. For Instant clone farms it can be only be a cluster id. This can be modified only if there are no current operations ( operation is NONE).
	HostOrClusterId string `json:"host_or_cluster_id"`
	// Resource pool to deploy the RDS Servers.
	ResourcePoolId string `json:"resource_pool_id"`
}

// NewFarmProvisioningSettingsUpdateSpec instantiates a new FarmProvisioningSettingsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFarmProvisioningSettingsUpdateSpec(hostOrClusterId string, resourcePoolId string) *FarmProvisioningSettingsUpdateSpec {
	this := FarmProvisioningSettingsUpdateSpec{}
	this.HostOrClusterId = hostOrClusterId
	this.ResourcePoolId = resourcePoolId
	return &this
}

// NewFarmProvisioningSettingsUpdateSpecWithDefaults instantiates a new FarmProvisioningSettingsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFarmProvisioningSettingsUpdateSpecWithDefaults() *FarmProvisioningSettingsUpdateSpec {
	this := FarmProvisioningSettingsUpdateSpec{}
	return &this
}

// GetHostOrClusterId returns the HostOrClusterId field value
func (o *FarmProvisioningSettingsUpdateSpec) GetHostOrClusterId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HostOrClusterId
}

// GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field value
// and a boolean to check if the value has been set.
func (o *FarmProvisioningSettingsUpdateSpec) GetHostOrClusterIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HostOrClusterId, true
}

// SetHostOrClusterId sets field value
func (o *FarmProvisioningSettingsUpdateSpec) SetHostOrClusterId(v string) {
	o.HostOrClusterId = v
}

// GetResourcePoolId returns the ResourcePoolId field value
func (o *FarmProvisioningSettingsUpdateSpec) GetResourcePoolId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value
// and a boolean to check if the value has been set.
func (o *FarmProvisioningSettingsUpdateSpec) GetResourcePoolIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourcePoolId, true
}

// SetResourcePoolId sets field value
func (o *FarmProvisioningSettingsUpdateSpec) SetResourcePoolId(v string) {
	o.ResourcePoolId = v
}

func (o FarmProvisioningSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host_or_cluster_id"] = o.HostOrClusterId
	}
	if true {
		toSerialize["resource_pool_id"] = o.ResourcePoolId
	}
	return json.Marshal(toSerialize)
}

type NullableFarmProvisioningSettingsUpdateSpec struct {
	value *FarmProvisioningSettingsUpdateSpec
	isSet bool
}

func (v NullableFarmProvisioningSettingsUpdateSpec) Get() *FarmProvisioningSettingsUpdateSpec {
	return v.value
}

func (v *NullableFarmProvisioningSettingsUpdateSpec) Set(val *FarmProvisioningSettingsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableFarmProvisioningSettingsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableFarmProvisioningSettingsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFarmProvisioningSettingsUpdateSpec(val *FarmProvisioningSettingsUpdateSpec) *NullableFarmProvisioningSettingsUpdateSpec {
	return &NullableFarmProvisioningSettingsUpdateSpec{value: val, isSet: true}
}

func (v NullableFarmProvisioningSettingsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFarmProvisioningSettingsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
