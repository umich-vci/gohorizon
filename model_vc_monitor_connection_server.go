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

// VCMonitorConnectionServer Information about connection to Virtual Center from Connection Server.
type VCMonitorConnectionServer struct {
	Certificate *CertificateMonitorInfo `json:"certificate,omitempty"`
	// Unique ID of the Connection Server.
	Id *string `json:"id,omitempty"`
	// Connection server host name or IP address.
	Name *string `json:"name,omitempty"`
	// Status of the Virtual Center Connection with respect to this Connection Server. * OK: The connection to Virtual Center server is working properly. * DOWN: The connection to Virtual Center server is down. * RECONNECTING: The connection to Virtual Center server was lost and is being reconnected to. * UNKNOWN: Connection state to Virtual Center server is unknown. * INVALID_CREDENTIALS: The supplied credentials cannot be used to authenticate to the Virtual Center server. * CANNOT_LOGIN: The connection server cannot login to the Virtual Center server. * NOT_YET_CONNECTED: Connection server has not yet connected to Virtual Center server.
	Status *string `json:"status,omitempty"`
	// Indicates if the thumbprints of the Virtual Center was accepted.
	ThumbprintAccepted *bool `json:"thumbprint_accepted,omitempty"`
}

// NewVCMonitorConnectionServer instantiates a new VCMonitorConnectionServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVCMonitorConnectionServer() *VCMonitorConnectionServer {
	this := VCMonitorConnectionServer{}
	return &this
}

// NewVCMonitorConnectionServerWithDefaults instantiates a new VCMonitorConnectionServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVCMonitorConnectionServerWithDefaults() *VCMonitorConnectionServer {
	this := VCMonitorConnectionServer{}
	return &this
}

// GetCertificate returns the Certificate field value if set, zero value otherwise.
func (o *VCMonitorConnectionServer) GetCertificate() CertificateMonitorInfo {
	if o == nil || o.Certificate == nil {
		var ret CertificateMonitorInfo
		return ret
	}
	return *o.Certificate
}

// GetCertificateOk returns a tuple with the Certificate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorConnectionServer) GetCertificateOk() (*CertificateMonitorInfo, bool) {
	if o == nil || o.Certificate == nil {
		return nil, false
	}
	return o.Certificate, true
}

// HasCertificate returns a boolean if a field has been set.
func (o *VCMonitorConnectionServer) HasCertificate() bool {
	if o != nil && o.Certificate != nil {
		return true
	}

	return false
}

// SetCertificate gets a reference to the given CertificateMonitorInfo and assigns it to the Certificate field.
func (o *VCMonitorConnectionServer) SetCertificate(v CertificateMonitorInfo) {
	o.Certificate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VCMonitorConnectionServer) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorConnectionServer) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VCMonitorConnectionServer) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VCMonitorConnectionServer) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VCMonitorConnectionServer) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorConnectionServer) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VCMonitorConnectionServer) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VCMonitorConnectionServer) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VCMonitorConnectionServer) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorConnectionServer) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VCMonitorConnectionServer) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VCMonitorConnectionServer) SetStatus(v string) {
	o.Status = &v
}

// GetThumbprintAccepted returns the ThumbprintAccepted field value if set, zero value otherwise.
func (o *VCMonitorConnectionServer) GetThumbprintAccepted() bool {
	if o == nil || o.ThumbprintAccepted == nil {
		var ret bool
		return ret
	}
	return *o.ThumbprintAccepted
}

// GetThumbprintAcceptedOk returns a tuple with the ThumbprintAccepted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VCMonitorConnectionServer) GetThumbprintAcceptedOk() (*bool, bool) {
	if o == nil || o.ThumbprintAccepted == nil {
		return nil, false
	}
	return o.ThumbprintAccepted, true
}

// HasThumbprintAccepted returns a boolean if a field has been set.
func (o *VCMonitorConnectionServer) HasThumbprintAccepted() bool {
	if o != nil && o.ThumbprintAccepted != nil {
		return true
	}

	return false
}

// SetThumbprintAccepted gets a reference to the given bool and assigns it to the ThumbprintAccepted field.
func (o *VCMonitorConnectionServer) SetThumbprintAccepted(v bool) {
	o.ThumbprintAccepted = &v
}

func (o VCMonitorConnectionServer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Certificate != nil {
		toSerialize["certificate"] = o.Certificate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.ThumbprintAccepted != nil {
		toSerialize["thumbprint_accepted"] = o.ThumbprintAccepted
	}
	return json.Marshal(toSerialize)
}

type NullableVCMonitorConnectionServer struct {
	value *VCMonitorConnectionServer
	isSet bool
}

func (v NullableVCMonitorConnectionServer) Get() *VCMonitorConnectionServer {
	return v.value
}

func (v *NullableVCMonitorConnectionServer) Set(val *VCMonitorConnectionServer) {
	v.value = val
	v.isSet = true
}

func (v NullableVCMonitorConnectionServer) IsSet() bool {
	return v.isSet
}

func (v *NullableVCMonitorConnectionServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVCMonitorConnectionServer(val *VCMonitorConnectionServer) *NullableVCMonitorConnectionServer {
	return &NullableVCMonitorConnectionServer{value: val, isSet: true}
}

func (v NullableVCMonitorConnectionServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVCMonitorConnectionServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
