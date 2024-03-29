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

// ResourcePoolInfo Information related to resource pool.
type ResourcePoolInfo struct {
	// Child nodes of the resource pool.
	Children *[]ResourcePoolInfo `json:"children,omitempty"`
	// Unique ID representing the resource pool.
	Id *string `json:"id,omitempty"`
	// Resource pool name.
	Name *string `json:"name,omitempty"`
	// Resource pool path.
	Path *string `json:"path,omitempty"`
	// Resource pool type. * HOST: Host used as a resource pool suitable for use in desktop pool/farm. * CLUSTER: Cluster used as a resource pool suitable for use in desktop pool/farm. * RESOURCE_POOL: Regular resource pool suitable for use in desktop pool/farm. * OTHER: Other resource type which cannot be used in desktop pool/farm.
	Type *string `json:"type,omitempty"`
}

// NewResourcePoolInfo instantiates a new ResourcePoolInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePoolInfo() *ResourcePoolInfo {
	this := ResourcePoolInfo{}
	return &this
}

// NewResourcePoolInfoWithDefaults instantiates a new ResourcePoolInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePoolInfoWithDefaults() *ResourcePoolInfo {
	this := ResourcePoolInfo{}
	return &this
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *ResourcePoolInfo) GetChildren() []ResourcePoolInfo {
	if o == nil || o.Children == nil {
		var ret []ResourcePoolInfo
		return ret
	}
	return *o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePoolInfo) GetChildrenOk() (*[]ResourcePoolInfo, bool) {
	if o == nil || o.Children == nil {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *ResourcePoolInfo) HasChildren() bool {
	if o != nil && o.Children != nil {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []ResourcePoolInfo and assigns it to the Children field.
func (o *ResourcePoolInfo) SetChildren(v []ResourcePoolInfo) {
	o.Children = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourcePoolInfo) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePoolInfo) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourcePoolInfo) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourcePoolInfo) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourcePoolInfo) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePoolInfo) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourcePoolInfo) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourcePoolInfo) SetName(v string) {
	o.Name = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ResourcePoolInfo) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePoolInfo) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ResourcePoolInfo) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ResourcePoolInfo) SetPath(v string) {
	o.Path = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ResourcePoolInfo) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePoolInfo) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ResourcePoolInfo) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ResourcePoolInfo) SetType(v string) {
	o.Type = &v
}

func (o ResourcePoolInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Children != nil {
		toSerialize["children"] = o.Children
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableResourcePoolInfo struct {
	value *ResourcePoolInfo
	isSet bool
}

func (v NullableResourcePoolInfo) Get() *ResourcePoolInfo {
	return v.value
}

func (v *NullableResourcePoolInfo) Set(val *ResourcePoolInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePoolInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePoolInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePoolInfo(val *ResourcePoolInfo) *NullableResourcePoolInfo {
	return &NullableResourcePoolInfo{value: val, isSet: true}
}

func (v NullableResourcePoolInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePoolInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
