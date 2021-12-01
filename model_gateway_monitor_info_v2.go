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

// GatewayMonitorInfoV2 Monitoring information related to Gateways registered in the environment.
type GatewayMonitorInfoV2 struct {
	// Number of active connections for the gateway. Includes PCoIP and BLAST connection counts.
	ActiveConnectionCount *int32 `json:"active_connection_count,omitempty"`
	// Number of BLAST connections for the gateway.
	BlastConnectionCount *int32                 `json:"blast_connection_count,omitempty"`
	Details              *GatewayMonitorDetails `json:"details,omitempty"`
	// Unique ID of the Gateway.
	Id *string `json:"id,omitempty"`
	// The timestamp in milliseconds when the last update was obtained. Measured as epoch time.
	LastUpdatedTimestamp *int64 `json:"last_updated_timestamp,omitempty"`
	// Gateway name.
	Name *string `json:"name,omitempty"`
	// Number of PCoIP connections for the gateway.
	PcoipConnectionCount *int32 `json:"pcoip_connection_count,omitempty"`
	// Status of the Gateway. * NOT_CONTACTED: There has been no contact from the gateway. * PROBLEM: The gateway has reported a problem. * STALE: Gateway is stale. Gateway will be marked as stale when Connection Server does not receive any request from the Gateway in last two successive intervals. * OK: The Gateway is working as expected.
	Status *string `json:"status,omitempty"`
}

// NewGatewayMonitorInfoV2 instantiates a new GatewayMonitorInfoV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayMonitorInfoV2() *GatewayMonitorInfoV2 {
	this := GatewayMonitorInfoV2{}
	return &this
}

// NewGatewayMonitorInfoV2WithDefaults instantiates a new GatewayMonitorInfoV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayMonitorInfoV2WithDefaults() *GatewayMonitorInfoV2 {
	this := GatewayMonitorInfoV2{}
	return &this
}

// GetActiveConnectionCount returns the ActiveConnectionCount field value if set, zero value otherwise.
func (o *GatewayMonitorInfoV2) GetActiveConnectionCount() int32 {
	if o == nil || o.ActiveConnectionCount == nil {
		var ret int32
		return ret
	}
	return *o.ActiveConnectionCount
}

// GetActiveConnectionCountOk returns a tuple with the ActiveConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorInfoV2) GetActiveConnectionCountOk() (*int32, bool) {
	if o == nil || o.ActiveConnectionCount == nil {
		return nil, false
	}
	return o.ActiveConnectionCount, true
}

// HasActiveConnectionCount returns a boolean if a field has been set.
func (o *GatewayMonitorInfoV2) HasActiveConnectionCount() bool {
	if o != nil && o.ActiveConnectionCount != nil {
		return true
	}

	return false
}

// SetActiveConnectionCount gets a reference to the given int32 and assigns it to the ActiveConnectionCount field.
func (o *GatewayMonitorInfoV2) SetActiveConnectionCount(v int32) {
	o.ActiveConnectionCount = &v
}

// GetBlastConnectionCount returns the BlastConnectionCount field value if set, zero value otherwise.
func (o *GatewayMonitorInfoV2) GetBlastConnectionCount() int32 {
	if o == nil || o.BlastConnectionCount == nil {
		var ret int32
		return ret
	}
	return *o.BlastConnectionCount
}

// GetBlastConnectionCountOk returns a tuple with the BlastConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorInfoV2) GetBlastConnectionCountOk() (*int32, bool) {
	if o == nil || o.BlastConnectionCount == nil {
		return nil, false
	}
	return o.BlastConnectionCount, true
}

// HasBlastConnectionCount returns a boolean if a field has been set.
func (o *GatewayMonitorInfoV2) HasBlastConnectionCount() bool {
	if o != nil && o.BlastConnectionCount != nil {
		return true
	}

	return false
}

// SetBlastConnectionCount gets a reference to the given int32 and assigns it to the BlastConnectionCount field.
func (o *GatewayMonitorInfoV2) SetBlastConnectionCount(v int32) {
	o.BlastConnectionCount = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GatewayMonitorInfoV2) GetDetails() GatewayMonitorDetails {
	if o == nil || o.Details == nil {
		var ret GatewayMonitorDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorInfoV2) GetDetailsOk() (*GatewayMonitorDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GatewayMonitorInfoV2) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given GatewayMonitorDetails and assigns it to the Details field.
func (o *GatewayMonitorInfoV2) SetDetails(v GatewayMonitorDetails) {
	o.Details = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayMonitorInfoV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorInfoV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayMonitorInfoV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GatewayMonitorInfoV2) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field value if set, zero value otherwise.
func (o *GatewayMonitorInfoV2) GetLastUpdatedTimestamp() int64 {
	if o == nil || o.LastUpdatedTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimestamp
}

// GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorInfoV2) GetLastUpdatedTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdatedTimestamp == nil {
		return nil, false
	}
	return o.LastUpdatedTimestamp, true
}

// HasLastUpdatedTimestamp returns a boolean if a field has been set.
func (o *GatewayMonitorInfoV2) HasLastUpdatedTimestamp() bool {
	if o != nil && o.LastUpdatedTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdatedTimestamp gets a reference to the given int64 and assigns it to the LastUpdatedTimestamp field.
func (o *GatewayMonitorInfoV2) SetLastUpdatedTimestamp(v int64) {
	o.LastUpdatedTimestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GatewayMonitorInfoV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorInfoV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GatewayMonitorInfoV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GatewayMonitorInfoV2) SetName(v string) {
	o.Name = &v
}

// GetPcoipConnectionCount returns the PcoipConnectionCount field value if set, zero value otherwise.
func (o *GatewayMonitorInfoV2) GetPcoipConnectionCount() int32 {
	if o == nil || o.PcoipConnectionCount == nil {
		var ret int32
		return ret
	}
	return *o.PcoipConnectionCount
}

// GetPcoipConnectionCountOk returns a tuple with the PcoipConnectionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorInfoV2) GetPcoipConnectionCountOk() (*int32, bool) {
	if o == nil || o.PcoipConnectionCount == nil {
		return nil, false
	}
	return o.PcoipConnectionCount, true
}

// HasPcoipConnectionCount returns a boolean if a field has been set.
func (o *GatewayMonitorInfoV2) HasPcoipConnectionCount() bool {
	if o != nil && o.PcoipConnectionCount != nil {
		return true
	}

	return false
}

// SetPcoipConnectionCount gets a reference to the given int32 and assigns it to the PcoipConnectionCount field.
func (o *GatewayMonitorInfoV2) SetPcoipConnectionCount(v int32) {
	o.PcoipConnectionCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GatewayMonitorInfoV2) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayMonitorInfoV2) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GatewayMonitorInfoV2) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GatewayMonitorInfoV2) SetStatus(v string) {
	o.Status = &v
}

func (o GatewayMonitorInfoV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActiveConnectionCount != nil {
		toSerialize["active_connection_count"] = o.ActiveConnectionCount
	}
	if o.BlastConnectionCount != nil {
		toSerialize["blast_connection_count"] = o.BlastConnectionCount
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTimestamp != nil {
		toSerialize["last_updated_timestamp"] = o.LastUpdatedTimestamp
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PcoipConnectionCount != nil {
		toSerialize["pcoip_connection_count"] = o.PcoipConnectionCount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayMonitorInfoV2 struct {
	value *GatewayMonitorInfoV2
	isSet bool
}

func (v NullableGatewayMonitorInfoV2) Get() *GatewayMonitorInfoV2 {
	return v.value
}

func (v *NullableGatewayMonitorInfoV2) Set(val *GatewayMonitorInfoV2) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayMonitorInfoV2) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayMonitorInfoV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayMonitorInfoV2(val *GatewayMonitorInfoV2) *NullableGatewayMonitorInfoV2 {
	return &NullableGatewayMonitorInfoV2{value: val, isSet: true}
}

func (v NullableGatewayMonitorInfoV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayMonitorInfoV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
