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

// ADDomainMonitorConnectionServerV2 Information about the AD Domain connection from connection server.
type ADDomainMonitorConnectionServerV2 struct {
	// Unique ID of the connection server.
	Id *string `json:"id,omitempty"`
	// The timestamp in milliseconds when the last update was obtained. Measured as epoch time.
	LastUpdatedTimestamp *int64 `json:"last_updated_timestamp,omitempty"`
	// Connection server host name or IP address.
	Name *string `json:"name,omitempty"`
	// Status of the connection to the domain. * UNCONTACTABLE: No domain controllers appear to be present on the network for this domain. * FULLY_ACCESSIBLE: The domain controller(s) are accepting bind operations. * CANNOT_BIND: The domain controller(s) are only accepting LDAP ping operations. * UNKNOWN: Cannot determine accessibility.
	Status *string `json:"status,omitempty"`
	// The trust relationship for the domain. * PRIMARY_DOMAIN: The domain is the domain that the broker is present in. * FROM_BROKER: The domain is trusted by the domain that the broker is in. * TO_BROKER: The domain trusts the broker's domain (this is for completeness and generally will not be used). * TWO_WAY: The domain has a two way trust relationship with the broker's domain. * TWO_WAY_FOREST: The domain is in the same forest as the broker's domain, implies two way trust. * MANUAL: The domain was manually configured (the trust has not been detected). * NOTRUST: The domain not having trust with broker domain. * UNKNOWN: The trust relationship could not be determined.
	TrustRelationship *string `json:"trust_relationship,omitempty"`
}

// NewADDomainMonitorConnectionServerV2 instantiates a new ADDomainMonitorConnectionServerV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewADDomainMonitorConnectionServerV2() *ADDomainMonitorConnectionServerV2 {
	this := ADDomainMonitorConnectionServerV2{}
	return &this
}

// NewADDomainMonitorConnectionServerV2WithDefaults instantiates a new ADDomainMonitorConnectionServerV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewADDomainMonitorConnectionServerV2WithDefaults() *ADDomainMonitorConnectionServerV2 {
	this := ADDomainMonitorConnectionServerV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ADDomainMonitorConnectionServerV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorConnectionServerV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ADDomainMonitorConnectionServerV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ADDomainMonitorConnectionServerV2) SetId(v string) {
	o.Id = &v
}

// GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field value if set, zero value otherwise.
func (o *ADDomainMonitorConnectionServerV2) GetLastUpdatedTimestamp() int64 {
	if o == nil || o.LastUpdatedTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.LastUpdatedTimestamp
}

// GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorConnectionServerV2) GetLastUpdatedTimestampOk() (*int64, bool) {
	if o == nil || o.LastUpdatedTimestamp == nil {
		return nil, false
	}
	return o.LastUpdatedTimestamp, true
}

// HasLastUpdatedTimestamp returns a boolean if a field has been set.
func (o *ADDomainMonitorConnectionServerV2) HasLastUpdatedTimestamp() bool {
	if o != nil && o.LastUpdatedTimestamp != nil {
		return true
	}

	return false
}

// SetLastUpdatedTimestamp gets a reference to the given int64 and assigns it to the LastUpdatedTimestamp field.
func (o *ADDomainMonitorConnectionServerV2) SetLastUpdatedTimestamp(v int64) {
	o.LastUpdatedTimestamp = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ADDomainMonitorConnectionServerV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorConnectionServerV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ADDomainMonitorConnectionServerV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ADDomainMonitorConnectionServerV2) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ADDomainMonitorConnectionServerV2) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorConnectionServerV2) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ADDomainMonitorConnectionServerV2) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ADDomainMonitorConnectionServerV2) SetStatus(v string) {
	o.Status = &v
}

// GetTrustRelationship returns the TrustRelationship field value if set, zero value otherwise.
func (o *ADDomainMonitorConnectionServerV2) GetTrustRelationship() string {
	if o == nil || o.TrustRelationship == nil {
		var ret string
		return ret
	}
	return *o.TrustRelationship
}

// GetTrustRelationshipOk returns a tuple with the TrustRelationship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ADDomainMonitorConnectionServerV2) GetTrustRelationshipOk() (*string, bool) {
	if o == nil || o.TrustRelationship == nil {
		return nil, false
	}
	return o.TrustRelationship, true
}

// HasTrustRelationship returns a boolean if a field has been set.
func (o *ADDomainMonitorConnectionServerV2) HasTrustRelationship() bool {
	if o != nil && o.TrustRelationship != nil {
		return true
	}

	return false
}

// SetTrustRelationship gets a reference to the given string and assigns it to the TrustRelationship field.
func (o *ADDomainMonitorConnectionServerV2) SetTrustRelationship(v string) {
	o.TrustRelationship = &v
}

func (o ADDomainMonitorConnectionServerV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastUpdatedTimestamp != nil {
		toSerialize["last_updated_timestamp"] = o.LastUpdatedTimestamp
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.TrustRelationship != nil {
		toSerialize["trust_relationship"] = o.TrustRelationship
	}
	return json.Marshal(toSerialize)
}

type NullableADDomainMonitorConnectionServerV2 struct {
	value *ADDomainMonitorConnectionServerV2
	isSet bool
}

func (v NullableADDomainMonitorConnectionServerV2) Get() *ADDomainMonitorConnectionServerV2 {
	return v.value
}

func (v *NullableADDomainMonitorConnectionServerV2) Set(val *ADDomainMonitorConnectionServerV2) {
	v.value = val
	v.isSet = true
}

func (v NullableADDomainMonitorConnectionServerV2) IsSet() bool {
	return v.isSet
}

func (v *NullableADDomainMonitorConnectionServerV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableADDomainMonitorConnectionServerV2(val *ADDomainMonitorConnectionServerV2) *NullableADDomainMonitorConnectionServerV2 {
	return &NullableADDomainMonitorConnectionServerV2{value: val, isSet: true}
}

func (v NullableADDomainMonitorConnectionServerV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableADDomainMonitorConnectionServerV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
