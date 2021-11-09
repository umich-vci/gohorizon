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

// EnvironmentInfoV2 Information related to Environment Settings.
type EnvironmentInfoV2 struct {
	// The GUID of a group of Connection Servers sharing the same configuration.
	ClusterGuid string `json:"cluster_guid"`
	// The name of a group of Connection Servers sharing the same configuration.
	ClusterName string `json:"cluster_name"`
	// Indicates different environments that Horizon can be deployed into. * GENERAL: Horizon is deployed on On-premises. * AZURE: Horizon is deployed on Azure. * AWS: Horizon is deployed on AWS. * DELL_EMC: Horizon is deployed on Dell EMC. * GOOGLE: Horizon is deployed on Google Cloud. * ORACLE: Horizon is deployed on Oracle Cloud.
	DeploymentType string `json:"deployment_type"`
	// Indicates if FIPS mode is enabled.
	FipsModeEnabled bool `json:"fips_mode_enabled"`
	// Indicates the IP mode of the environment. * IPv4: The ip mode is IPv4. * IPv6: The ip mode is IPv6.
	IpMode string `json:"ip_mode"`
	// Local connection Server build number.
	LocalConnectionServerBuild *string `json:"local_connection_server_build,omitempty"`
	// Local connection Server version number.
	LocalConnectionServerVersion *string `json:"local_connection_server_version,omitempty"`
	// The name of the current pod in the Multi-DataCenter Horizon Pod, the value will be null when PodFederation is not initialized.
	LocalPodName *string `json:"local_pod_name,omitempty"`
	// Represents the clusters time zone offset from UTC in seconds.
	TimezoneOffset int32 `json:"timezone_offset"`
}

// NewEnvironmentInfoV2 instantiates a new EnvironmentInfoV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentInfoV2(clusterGuid string, clusterName string, deploymentType string, fipsModeEnabled bool, ipMode string, timezoneOffset int32) *EnvironmentInfoV2 {
	this := EnvironmentInfoV2{}
	this.ClusterGuid = clusterGuid
	this.ClusterName = clusterName
	this.DeploymentType = deploymentType
	this.FipsModeEnabled = fipsModeEnabled
	this.IpMode = ipMode
	this.TimezoneOffset = timezoneOffset
	return &this
}

// NewEnvironmentInfoV2WithDefaults instantiates a new EnvironmentInfoV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentInfoV2WithDefaults() *EnvironmentInfoV2 {
	this := EnvironmentInfoV2{}
	return &this
}

// GetClusterGuid returns the ClusterGuid field value
func (o *EnvironmentInfoV2) GetClusterGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterGuid
}

// GetClusterGuidOk returns a tuple with the ClusterGuid field value
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetClusterGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterGuid, true
}

// SetClusterGuid sets field value
func (o *EnvironmentInfoV2) SetClusterGuid(v string) {
	o.ClusterGuid = v
}

// GetClusterName returns the ClusterName field value
func (o *EnvironmentInfoV2) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetClusterNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *EnvironmentInfoV2) SetClusterName(v string) {
	o.ClusterName = v
}

// GetDeploymentType returns the DeploymentType field value
func (o *EnvironmentInfoV2) GetDeploymentType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DeploymentType
}

// GetDeploymentTypeOk returns a tuple with the DeploymentType field value
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetDeploymentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeploymentType, true
}

// SetDeploymentType sets field value
func (o *EnvironmentInfoV2) SetDeploymentType(v string) {
	o.DeploymentType = v
}

// GetFipsModeEnabled returns the FipsModeEnabled field value
func (o *EnvironmentInfoV2) GetFipsModeEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.FipsModeEnabled
}

// GetFipsModeEnabledOk returns a tuple with the FipsModeEnabled field value
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetFipsModeEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FipsModeEnabled, true
}

// SetFipsModeEnabled sets field value
func (o *EnvironmentInfoV2) SetFipsModeEnabled(v bool) {
	o.FipsModeEnabled = v
}

// GetIpMode returns the IpMode field value
func (o *EnvironmentInfoV2) GetIpMode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpMode
}

// GetIpModeOk returns a tuple with the IpMode field value
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetIpModeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpMode, true
}

// SetIpMode sets field value
func (o *EnvironmentInfoV2) SetIpMode(v string) {
	o.IpMode = v
}

// GetLocalConnectionServerBuild returns the LocalConnectionServerBuild field value if set, zero value otherwise.
func (o *EnvironmentInfoV2) GetLocalConnectionServerBuild() string {
	if o == nil || o.LocalConnectionServerBuild == nil {
		var ret string
		return ret
	}
	return *o.LocalConnectionServerBuild
}

// GetLocalConnectionServerBuildOk returns a tuple with the LocalConnectionServerBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetLocalConnectionServerBuildOk() (*string, bool) {
	if o == nil || o.LocalConnectionServerBuild == nil {
		return nil, false
	}
	return o.LocalConnectionServerBuild, true
}

// HasLocalConnectionServerBuild returns a boolean if a field has been set.
func (o *EnvironmentInfoV2) HasLocalConnectionServerBuild() bool {
	if o != nil && o.LocalConnectionServerBuild != nil {
		return true
	}

	return false
}

// SetLocalConnectionServerBuild gets a reference to the given string and assigns it to the LocalConnectionServerBuild field.
func (o *EnvironmentInfoV2) SetLocalConnectionServerBuild(v string) {
	o.LocalConnectionServerBuild = &v
}

// GetLocalConnectionServerVersion returns the LocalConnectionServerVersion field value if set, zero value otherwise.
func (o *EnvironmentInfoV2) GetLocalConnectionServerVersion() string {
	if o == nil || o.LocalConnectionServerVersion == nil {
		var ret string
		return ret
	}
	return *o.LocalConnectionServerVersion
}

// GetLocalConnectionServerVersionOk returns a tuple with the LocalConnectionServerVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetLocalConnectionServerVersionOk() (*string, bool) {
	if o == nil || o.LocalConnectionServerVersion == nil {
		return nil, false
	}
	return o.LocalConnectionServerVersion, true
}

// HasLocalConnectionServerVersion returns a boolean if a field has been set.
func (o *EnvironmentInfoV2) HasLocalConnectionServerVersion() bool {
	if o != nil && o.LocalConnectionServerVersion != nil {
		return true
	}

	return false
}

// SetLocalConnectionServerVersion gets a reference to the given string and assigns it to the LocalConnectionServerVersion field.
func (o *EnvironmentInfoV2) SetLocalConnectionServerVersion(v string) {
	o.LocalConnectionServerVersion = &v
}

// GetLocalPodName returns the LocalPodName field value if set, zero value otherwise.
func (o *EnvironmentInfoV2) GetLocalPodName() string {
	if o == nil || o.LocalPodName == nil {
		var ret string
		return ret
	}
	return *o.LocalPodName
}

// GetLocalPodNameOk returns a tuple with the LocalPodName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetLocalPodNameOk() (*string, bool) {
	if o == nil || o.LocalPodName == nil {
		return nil, false
	}
	return o.LocalPodName, true
}

// HasLocalPodName returns a boolean if a field has been set.
func (o *EnvironmentInfoV2) HasLocalPodName() bool {
	if o != nil && o.LocalPodName != nil {
		return true
	}

	return false
}

// SetLocalPodName gets a reference to the given string and assigns it to the LocalPodName field.
func (o *EnvironmentInfoV2) SetLocalPodName(v string) {
	o.LocalPodName = &v
}

// GetTimezoneOffset returns the TimezoneOffset field value
func (o *EnvironmentInfoV2) GetTimezoneOffset() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TimezoneOffset
}

// GetTimezoneOffsetOk returns a tuple with the TimezoneOffset field value
// and a boolean to check if the value has been set.
func (o *EnvironmentInfoV2) GetTimezoneOffsetOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimezoneOffset, true
}

// SetTimezoneOffset sets field value
func (o *EnvironmentInfoV2) SetTimezoneOffset(v int32) {
	o.TimezoneOffset = v
}

func (o EnvironmentInfoV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster_guid"] = o.ClusterGuid
	}
	if true {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if true {
		toSerialize["deployment_type"] = o.DeploymentType
	}
	if true {
		toSerialize["fips_mode_enabled"] = o.FipsModeEnabled
	}
	if true {
		toSerialize["ip_mode"] = o.IpMode
	}
	if o.LocalConnectionServerBuild != nil {
		toSerialize["local_connection_server_build"] = o.LocalConnectionServerBuild
	}
	if o.LocalConnectionServerVersion != nil {
		toSerialize["local_connection_server_version"] = o.LocalConnectionServerVersion
	}
	if o.LocalPodName != nil {
		toSerialize["local_pod_name"] = o.LocalPodName
	}
	if true {
		toSerialize["timezone_offset"] = o.TimezoneOffset
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentInfoV2 struct {
	value *EnvironmentInfoV2
	isSet bool
}

func (v NullableEnvironmentInfoV2) Get() *EnvironmentInfoV2 {
	return v.value
}

func (v *NullableEnvironmentInfoV2) Set(val *EnvironmentInfoV2) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentInfoV2) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentInfoV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentInfoV2(val *EnvironmentInfoV2) *NullableEnvironmentInfoV2 {
	return &NullableEnvironmentInfoV2{value: val, isSet: true}
}

func (v NullableEnvironmentInfoV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentInfoV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
