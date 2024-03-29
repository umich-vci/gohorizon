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

// VirtualCenterMonitorInfoV2 Monitoring information related to a Virtual Center
type VirtualCenterMonitorInfoV2 struct {
	// Information about the Virtual Center connections from each of the connection servers.
	ConnectionServers *[]VCMonitorConnectionServerV2 `json:"connection_servers,omitempty"`
	// Information about the datastores of the Virtual Center.
	Datastores *[]VCMonitorDatastore `json:"datastores,omitempty"`
	// Number of Desktop Pools And Farms managed by the virtual center
	DesktopPoolsAndFarmsCount *int32            `json:"desktop_pools_and_farms_count,omitempty"`
	Details                   *VCMonitorDetails `json:"details,omitempty"`
	// Information about the ESX hosts added in the Virtual Center.
	Hosts *[]VCMonitorHost `json:"hosts,omitempty"`
	// Unique ID of the Virtual Center.
	Id *string `json:"id,omitempty"`
	// Virtual Center server name or IP address.
	Name *string `json:"name,omitempty"`
}

// NewVirtualCenterMonitorInfoV2 instantiates a new VirtualCenterMonitorInfoV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVirtualCenterMonitorInfoV2() *VirtualCenterMonitorInfoV2 {
	this := VirtualCenterMonitorInfoV2{}
	return &this
}

// NewVirtualCenterMonitorInfoV2WithDefaults instantiates a new VirtualCenterMonitorInfoV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVirtualCenterMonitorInfoV2WithDefaults() *VirtualCenterMonitorInfoV2 {
	this := VirtualCenterMonitorInfoV2{}
	return &this
}

// GetConnectionServers returns the ConnectionServers field value if set, zero value otherwise.
func (o *VirtualCenterMonitorInfoV2) GetConnectionServers() []VCMonitorConnectionServerV2 {
	if o == nil || o.ConnectionServers == nil {
		var ret []VCMonitorConnectionServerV2
		return ret
	}
	return *o.ConnectionServers
}

// GetConnectionServersOk returns a tuple with the ConnectionServers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCenterMonitorInfoV2) GetConnectionServersOk() (*[]VCMonitorConnectionServerV2, bool) {
	if o == nil || o.ConnectionServers == nil {
		return nil, false
	}
	return o.ConnectionServers, true
}

// HasConnectionServers returns a boolean if a field has been set.
func (o *VirtualCenterMonitorInfoV2) HasConnectionServers() bool {
	if o != nil && o.ConnectionServers != nil {
		return true
	}

	return false
}

// SetConnectionServers gets a reference to the given []VCMonitorConnectionServerV2 and assigns it to the ConnectionServers field.
func (o *VirtualCenterMonitorInfoV2) SetConnectionServers(v []VCMonitorConnectionServerV2) {
	o.ConnectionServers = &v
}

// GetDatastores returns the Datastores field value if set, zero value otherwise.
func (o *VirtualCenterMonitorInfoV2) GetDatastores() []VCMonitorDatastore {
	if o == nil || o.Datastores == nil {
		var ret []VCMonitorDatastore
		return ret
	}
	return *o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCenterMonitorInfoV2) GetDatastoresOk() (*[]VCMonitorDatastore, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *VirtualCenterMonitorInfoV2) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []VCMonitorDatastore and assigns it to the Datastores field.
func (o *VirtualCenterMonitorInfoV2) SetDatastores(v []VCMonitorDatastore) {
	o.Datastores = &v
}

// GetDesktopPoolsAndFarmsCount returns the DesktopPoolsAndFarmsCount field value if set, zero value otherwise.
func (o *VirtualCenterMonitorInfoV2) GetDesktopPoolsAndFarmsCount() int32 {
	if o == nil || o.DesktopPoolsAndFarmsCount == nil {
		var ret int32
		return ret
	}
	return *o.DesktopPoolsAndFarmsCount
}

// GetDesktopPoolsAndFarmsCountOk returns a tuple with the DesktopPoolsAndFarmsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCenterMonitorInfoV2) GetDesktopPoolsAndFarmsCountOk() (*int32, bool) {
	if o == nil || o.DesktopPoolsAndFarmsCount == nil {
		return nil, false
	}
	return o.DesktopPoolsAndFarmsCount, true
}

// HasDesktopPoolsAndFarmsCount returns a boolean if a field has been set.
func (o *VirtualCenterMonitorInfoV2) HasDesktopPoolsAndFarmsCount() bool {
	if o != nil && o.DesktopPoolsAndFarmsCount != nil {
		return true
	}

	return false
}

// SetDesktopPoolsAndFarmsCount gets a reference to the given int32 and assigns it to the DesktopPoolsAndFarmsCount field.
func (o *VirtualCenterMonitorInfoV2) SetDesktopPoolsAndFarmsCount(v int32) {
	o.DesktopPoolsAndFarmsCount = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *VirtualCenterMonitorInfoV2) GetDetails() VCMonitorDetails {
	if o == nil || o.Details == nil {
		var ret VCMonitorDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCenterMonitorInfoV2) GetDetailsOk() (*VCMonitorDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *VirtualCenterMonitorInfoV2) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given VCMonitorDetails and assigns it to the Details field.
func (o *VirtualCenterMonitorInfoV2) SetDetails(v VCMonitorDetails) {
	o.Details = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *VirtualCenterMonitorInfoV2) GetHosts() []VCMonitorHost {
	if o == nil || o.Hosts == nil {
		var ret []VCMonitorHost
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCenterMonitorInfoV2) GetHostsOk() (*[]VCMonitorHost, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VirtualCenterMonitorInfoV2) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []VCMonitorHost and assigns it to the Hosts field.
func (o *VirtualCenterMonitorInfoV2) SetHosts(v []VCMonitorHost) {
	o.Hosts = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VirtualCenterMonitorInfoV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCenterMonitorInfoV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VirtualCenterMonitorInfoV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VirtualCenterMonitorInfoV2) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VirtualCenterMonitorInfoV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VirtualCenterMonitorInfoV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VirtualCenterMonitorInfoV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VirtualCenterMonitorInfoV2) SetName(v string) {
	o.Name = &v
}

func (o VirtualCenterMonitorInfoV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ConnectionServers != nil {
		toSerialize["connection_servers"] = o.ConnectionServers
	}
	if o.Datastores != nil {
		toSerialize["datastores"] = o.Datastores
	}
	if o.DesktopPoolsAndFarmsCount != nil {
		toSerialize["desktop_pools_and_farms_count"] = o.DesktopPoolsAndFarmsCount
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableVirtualCenterMonitorInfoV2 struct {
	value *VirtualCenterMonitorInfoV2
	isSet bool
}

func (v NullableVirtualCenterMonitorInfoV2) Get() *VirtualCenterMonitorInfoV2 {
	return v.value
}

func (v *NullableVirtualCenterMonitorInfoV2) Set(val *VirtualCenterMonitorInfoV2) {
	v.value = val
	v.isSet = true
}

func (v NullableVirtualCenterMonitorInfoV2) IsSet() bool {
	return v.isSet
}

func (v *NullableVirtualCenterMonitorInfoV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVirtualCenterMonitorInfoV2(val *VirtualCenterMonitorInfoV2) *NullableVirtualCenterMonitorInfoV2 {
	return &NullableVirtualCenterMonitorInfoV2{value: val, isSet: true}
}

func (v NullableVirtualCenterMonitorInfoV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVirtualCenterMonitorInfoV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
