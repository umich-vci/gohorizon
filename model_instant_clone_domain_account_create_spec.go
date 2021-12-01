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

// InstantCloneDomainAccountCreateSpec Specification of the instant clone domain account.
type InstantCloneDomainAccountCreateSpec struct {
	// SID of the AD Domain that this account user belongs to.
	AdDomainId string `json:"ad_domain_id"`
	// Password of the account.
	Password []string `json:"password"`
	// User name of the account.
	Username string `json:"username"`
}

// NewInstantCloneDomainAccountCreateSpec instantiates a new InstantCloneDomainAccountCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInstantCloneDomainAccountCreateSpec(adDomainId string, password []string, username string) *InstantCloneDomainAccountCreateSpec {
	this := InstantCloneDomainAccountCreateSpec{}
	this.AdDomainId = adDomainId
	this.Password = password
	this.Username = username
	return &this
}

// NewInstantCloneDomainAccountCreateSpecWithDefaults instantiates a new InstantCloneDomainAccountCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInstantCloneDomainAccountCreateSpecWithDefaults() *InstantCloneDomainAccountCreateSpec {
	this := InstantCloneDomainAccountCreateSpec{}
	return &this
}

// GetAdDomainId returns the AdDomainId field value
func (o *InstantCloneDomainAccountCreateSpec) GetAdDomainId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdDomainId
}

// GetAdDomainIdOk returns a tuple with the AdDomainId field value
// and a boolean to check if the value has been set.
func (o *InstantCloneDomainAccountCreateSpec) GetAdDomainIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdDomainId, true
}

// SetAdDomainId sets field value
func (o *InstantCloneDomainAccountCreateSpec) SetAdDomainId(v string) {
	o.AdDomainId = v
}

// GetPassword returns the Password field value
func (o *InstantCloneDomainAccountCreateSpec) GetPassword() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *InstantCloneDomainAccountCreateSpec) GetPasswordOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *InstantCloneDomainAccountCreateSpec) SetPassword(v []string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *InstantCloneDomainAccountCreateSpec) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *InstantCloneDomainAccountCreateSpec) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *InstantCloneDomainAccountCreateSpec) SetUsername(v string) {
	o.Username = v
}

func (o InstantCloneDomainAccountCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ad_domain_id"] = o.AdDomainId
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableInstantCloneDomainAccountCreateSpec struct {
	value *InstantCloneDomainAccountCreateSpec
	isSet bool
}

func (v NullableInstantCloneDomainAccountCreateSpec) Get() *InstantCloneDomainAccountCreateSpec {
	return v.value
}

func (v *NullableInstantCloneDomainAccountCreateSpec) Set(val *InstantCloneDomainAccountCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableInstantCloneDomainAccountCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableInstantCloneDomainAccountCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInstantCloneDomainAccountCreateSpec(val *InstantCloneDomainAccountCreateSpec) *NullableInstantCloneDomainAccountCreateSpec {
	return &NullableInstantCloneDomainAccountCreateSpec{value: val, isSet: true}
}

func (v NullableInstantCloneDomainAccountCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInstantCloneDomainAccountCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
