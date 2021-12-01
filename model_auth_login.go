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

// AuthLogin Login Request
type AuthLogin struct {
	// Domain
	Domain string `json:"domain"`
	// User password
	Password string `json:"password"`
	// User Name
	Username string `json:"username"`
}

// NewAuthLogin instantiates a new AuthLogin object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthLogin(domain string, password string, username string) *AuthLogin {
	this := AuthLogin{}
	this.Domain = domain
	this.Password = password
	this.Username = username
	return &this
}

// NewAuthLoginWithDefaults instantiates a new AuthLogin object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthLoginWithDefaults() *AuthLogin {
	this := AuthLogin{}
	return &this
}

// GetDomain returns the Domain field value
func (o *AuthLogin) GetDomain() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Domain
}

// GetDomainOk returns a tuple with the Domain field value
// and a boolean to check if the value has been set.
func (o *AuthLogin) GetDomainOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Domain, true
}

// SetDomain sets field value
func (o *AuthLogin) SetDomain(v string) {
	o.Domain = v
}

// GetPassword returns the Password field value
func (o *AuthLogin) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AuthLogin) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AuthLogin) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *AuthLogin) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AuthLogin) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AuthLogin) SetUsername(v string) {
	o.Username = v
}

func (o AuthLogin) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["domain"] = o.Domain
	}
	if true {
		toSerialize["password"] = o.Password
	}
	if true {
		toSerialize["username"] = o.Username
	}
	return json.Marshal(toSerialize)
}

type NullableAuthLogin struct {
	value *AuthLogin
	isSet bool
}

func (v NullableAuthLogin) Get() *AuthLogin {
	return v.value
}

func (v *NullableAuthLogin) Set(val *AuthLogin) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthLogin) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthLogin) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthLogin(val *AuthLogin) *NullableAuthLogin {
	return &NullableAuthLogin{value: val, isSet: true}
}

func (v NullableAuthLogin) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthLogin) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
