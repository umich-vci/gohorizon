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

// RDSServerMonitorInfo RDS Server Monitor Information.
type RDSServerMonitorInfo struct {
	Details *RDSServerMonitorDetails `json:"details,omitempty"`
	// Indicates if RDS server is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// Indicates the Farm ID to which the RDS Server belongs to.
	FarmId *string `json:"farm_id,omitempty"`
	// Unique ID of the RDS server.
	Id *string `json:"id,omitempty"`
	// This value is similar to load_preference and represents the load on RDS Server in the range of 0 to 100.
	LoadIndex *int32 `json:"load_index,omitempty"`
	// Based on the current load of this RDS Server, gives a measure of how preferential this server is to be chosen for new application sessions. * BLOCK: RDS Server is blocked and new sessions will not be accepted. * HEAVY: RDS Server is experiencing heavy load and should likely not be chosen for new sessions. * NORMAL: RDS Server is experiencing normal load and is okay to be chosen for new sessions. * LIGHT: RDS Server is experiencing light load and is okay to be chosen for new sessions. * UNKNOWN: RDS Server did not report a load preference. This is potentially a configuration issue if other RDS Servers in the same Farm do report load preferences.
	LoadPreference *string `json:"load_preference,omitempty"`
	// RDS Server name.
	Name *string `json:"name,omitempty"`
	// RDS server session count.
	SessionCount *int32 `json:"session_count,omitempty"`
	// RDS server status. * OK: RDS Server is reachable. All applications (defined on its farm) are verified installed on the RDS Server. * WARNING: RDS Server is reachable. Some applications are detected as not installed on the RDS Server. * ERROR: RDS Server is unreachable, or, none of the applications are installed. * DISABLED: RDS Server is disabled.
	Status *string `json:"status,omitempty"`
}

// NewRDSServerMonitorInfo instantiates a new RDSServerMonitorInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRDSServerMonitorInfo() *RDSServerMonitorInfo {
	this := RDSServerMonitorInfo{}
	return &this
}

// NewRDSServerMonitorInfoWithDefaults instantiates a new RDSServerMonitorInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRDSServerMonitorInfoWithDefaults() *RDSServerMonitorInfo {
	this := RDSServerMonitorInfo{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetDetails() RDSServerMonitorDetails {
	if o == nil || o.Details == nil {
		var ret RDSServerMonitorDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetDetailsOk() (*RDSServerMonitorDetails, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given RDSServerMonitorDetails and assigns it to the Details field.
func (o *RDSServerMonitorInfo) SetDetails(v RDSServerMonitorDetails) {
	o.Details = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *RDSServerMonitorInfo) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFarmId returns the FarmId field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetFarmId() string {
	if o == nil || o.FarmId == nil {
		var ret string
		return ret
	}
	return *o.FarmId
}

// GetFarmIdOk returns a tuple with the FarmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetFarmIdOk() (*string, bool) {
	if o == nil || o.FarmId == nil {
		return nil, false
	}
	return o.FarmId, true
}

// HasFarmId returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasFarmId() bool {
	if o != nil && o.FarmId != nil {
		return true
	}

	return false
}

// SetFarmId gets a reference to the given string and assigns it to the FarmId field.
func (o *RDSServerMonitorInfo) SetFarmId(v string) {
	o.FarmId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RDSServerMonitorInfo) SetId(v string) {
	o.Id = &v
}

// GetLoadIndex returns the LoadIndex field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetLoadIndex() int32 {
	if o == nil || o.LoadIndex == nil {
		var ret int32
		return ret
	}
	return *o.LoadIndex
}

// GetLoadIndexOk returns a tuple with the LoadIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetLoadIndexOk() (*int32, bool) {
	if o == nil || o.LoadIndex == nil {
		return nil, false
	}
	return o.LoadIndex, true
}

// HasLoadIndex returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasLoadIndex() bool {
	if o != nil && o.LoadIndex != nil {
		return true
	}

	return false
}

// SetLoadIndex gets a reference to the given int32 and assigns it to the LoadIndex field.
func (o *RDSServerMonitorInfo) SetLoadIndex(v int32) {
	o.LoadIndex = &v
}

// GetLoadPreference returns the LoadPreference field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetLoadPreference() string {
	if o == nil || o.LoadPreference == nil {
		var ret string
		return ret
	}
	return *o.LoadPreference
}

// GetLoadPreferenceOk returns a tuple with the LoadPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetLoadPreferenceOk() (*string, bool) {
	if o == nil || o.LoadPreference == nil {
		return nil, false
	}
	return o.LoadPreference, true
}

// HasLoadPreference returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasLoadPreference() bool {
	if o != nil && o.LoadPreference != nil {
		return true
	}

	return false
}

// SetLoadPreference gets a reference to the given string and assigns it to the LoadPreference field.
func (o *RDSServerMonitorInfo) SetLoadPreference(v string) {
	o.LoadPreference = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RDSServerMonitorInfo) SetName(v string) {
	o.Name = &v
}

// GetSessionCount returns the SessionCount field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetSessionCount() int32 {
	if o == nil || o.SessionCount == nil {
		var ret int32
		return ret
	}
	return *o.SessionCount
}

// GetSessionCountOk returns a tuple with the SessionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetSessionCountOk() (*int32, bool) {
	if o == nil || o.SessionCount == nil {
		return nil, false
	}
	return o.SessionCount, true
}

// HasSessionCount returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasSessionCount() bool {
	if o != nil && o.SessionCount != nil {
		return true
	}

	return false
}

// SetSessionCount gets a reference to the given int32 and assigns it to the SessionCount field.
func (o *RDSServerMonitorInfo) SetSessionCount(v int32) {
	o.SessionCount = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *RDSServerMonitorInfo) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RDSServerMonitorInfo) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *RDSServerMonitorInfo) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *RDSServerMonitorInfo) SetStatus(v string) {
	o.Status = &v
}

func (o RDSServerMonitorInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FarmId != nil {
		toSerialize["farm_id"] = o.FarmId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LoadIndex != nil {
		toSerialize["load_index"] = o.LoadIndex
	}
	if o.LoadPreference != nil {
		toSerialize["load_preference"] = o.LoadPreference
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.SessionCount != nil {
		toSerialize["session_count"] = o.SessionCount
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableRDSServerMonitorInfo struct {
	value *RDSServerMonitorInfo
	isSet bool
}

func (v NullableRDSServerMonitorInfo) Get() *RDSServerMonitorInfo {
	return v.value
}

func (v *NullableRDSServerMonitorInfo) Set(val *RDSServerMonitorInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableRDSServerMonitorInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableRDSServerMonitorInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRDSServerMonitorInfo(val *RDSServerMonitorInfo) *NullableRDSServerMonitorInfo {
	return &NullableRDSServerMonitorInfo{value: val, isSet: true}
}

func (v NullableRDSServerMonitorInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRDSServerMonitorInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
