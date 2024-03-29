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

// ViewComposerMonitorDetails Details of the View Composer.
type ViewComposerMonitorDetails struct {
	// The version of the View Composer API used to communicate with the View Composer server.
	ApiVerion *string `json:"api_verion,omitempty"`
	// The build of the View Composer server.
	Build *string `json:"build,omitempty"`
	// The minimum ESX version required for compatibility with this View Composer server.
	MinEsxVersion *string `json:"min_esx_version,omitempty"`
	// The minimum Virtual Center version required for compatibility with this View Composer server.
	MinVcVersion *string `json:"min_vc_version,omitempty"`
	// The Virtual Center servers referencing to this View Composer.
	ReferencedVcs *[]string `json:"referenced_vcs,omitempty"`
	// The version of the View Composer server.
	Version *string `json:"version,omitempty"`
}

// NewViewComposerMonitorDetails instantiates a new ViewComposerMonitorDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewViewComposerMonitorDetails() *ViewComposerMonitorDetails {
	this := ViewComposerMonitorDetails{}
	return &this
}

// NewViewComposerMonitorDetailsWithDefaults instantiates a new ViewComposerMonitorDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewViewComposerMonitorDetailsWithDefaults() *ViewComposerMonitorDetails {
	this := ViewComposerMonitorDetails{}
	return &this
}

// GetApiVerion returns the ApiVerion field value if set, zero value otherwise.
func (o *ViewComposerMonitorDetails) GetApiVerion() string {
	if o == nil || o.ApiVerion == nil {
		var ret string
		return ret
	}
	return *o.ApiVerion
}

// GetApiVerionOk returns a tuple with the ApiVerion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComposerMonitorDetails) GetApiVerionOk() (*string, bool) {
	if o == nil || o.ApiVerion == nil {
		return nil, false
	}
	return o.ApiVerion, true
}

// HasApiVerion returns a boolean if a field has been set.
func (o *ViewComposerMonitorDetails) HasApiVerion() bool {
	if o != nil && o.ApiVerion != nil {
		return true
	}

	return false
}

// SetApiVerion gets a reference to the given string and assigns it to the ApiVerion field.
func (o *ViewComposerMonitorDetails) SetApiVerion(v string) {
	o.ApiVerion = &v
}

// GetBuild returns the Build field value if set, zero value otherwise.
func (o *ViewComposerMonitorDetails) GetBuild() string {
	if o == nil || o.Build == nil {
		var ret string
		return ret
	}
	return *o.Build
}

// GetBuildOk returns a tuple with the Build field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComposerMonitorDetails) GetBuildOk() (*string, bool) {
	if o == nil || o.Build == nil {
		return nil, false
	}
	return o.Build, true
}

// HasBuild returns a boolean if a field has been set.
func (o *ViewComposerMonitorDetails) HasBuild() bool {
	if o != nil && o.Build != nil {
		return true
	}

	return false
}

// SetBuild gets a reference to the given string and assigns it to the Build field.
func (o *ViewComposerMonitorDetails) SetBuild(v string) {
	o.Build = &v
}

// GetMinEsxVersion returns the MinEsxVersion field value if set, zero value otherwise.
func (o *ViewComposerMonitorDetails) GetMinEsxVersion() string {
	if o == nil || o.MinEsxVersion == nil {
		var ret string
		return ret
	}
	return *o.MinEsxVersion
}

// GetMinEsxVersionOk returns a tuple with the MinEsxVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComposerMonitorDetails) GetMinEsxVersionOk() (*string, bool) {
	if o == nil || o.MinEsxVersion == nil {
		return nil, false
	}
	return o.MinEsxVersion, true
}

// HasMinEsxVersion returns a boolean if a field has been set.
func (o *ViewComposerMonitorDetails) HasMinEsxVersion() bool {
	if o != nil && o.MinEsxVersion != nil {
		return true
	}

	return false
}

// SetMinEsxVersion gets a reference to the given string and assigns it to the MinEsxVersion field.
func (o *ViewComposerMonitorDetails) SetMinEsxVersion(v string) {
	o.MinEsxVersion = &v
}

// GetMinVcVersion returns the MinVcVersion field value if set, zero value otherwise.
func (o *ViewComposerMonitorDetails) GetMinVcVersion() string {
	if o == nil || o.MinVcVersion == nil {
		var ret string
		return ret
	}
	return *o.MinVcVersion
}

// GetMinVcVersionOk returns a tuple with the MinVcVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComposerMonitorDetails) GetMinVcVersionOk() (*string, bool) {
	if o == nil || o.MinVcVersion == nil {
		return nil, false
	}
	return o.MinVcVersion, true
}

// HasMinVcVersion returns a boolean if a field has been set.
func (o *ViewComposerMonitorDetails) HasMinVcVersion() bool {
	if o != nil && o.MinVcVersion != nil {
		return true
	}

	return false
}

// SetMinVcVersion gets a reference to the given string and assigns it to the MinVcVersion field.
func (o *ViewComposerMonitorDetails) SetMinVcVersion(v string) {
	o.MinVcVersion = &v
}

// GetReferencedVcs returns the ReferencedVcs field value if set, zero value otherwise.
func (o *ViewComposerMonitorDetails) GetReferencedVcs() []string {
	if o == nil || o.ReferencedVcs == nil {
		var ret []string
		return ret
	}
	return *o.ReferencedVcs
}

// GetReferencedVcsOk returns a tuple with the ReferencedVcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComposerMonitorDetails) GetReferencedVcsOk() (*[]string, bool) {
	if o == nil || o.ReferencedVcs == nil {
		return nil, false
	}
	return o.ReferencedVcs, true
}

// HasReferencedVcs returns a boolean if a field has been set.
func (o *ViewComposerMonitorDetails) HasReferencedVcs() bool {
	if o != nil && o.ReferencedVcs != nil {
		return true
	}

	return false
}

// SetReferencedVcs gets a reference to the given []string and assigns it to the ReferencedVcs field.
func (o *ViewComposerMonitorDetails) SetReferencedVcs(v []string) {
	o.ReferencedVcs = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ViewComposerMonitorDetails) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ViewComposerMonitorDetails) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ViewComposerMonitorDetails) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ViewComposerMonitorDetails) SetVersion(v string) {
	o.Version = &v
}

func (o ViewComposerMonitorDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApiVerion != nil {
		toSerialize["api_verion"] = o.ApiVerion
	}
	if o.Build != nil {
		toSerialize["build"] = o.Build
	}
	if o.MinEsxVersion != nil {
		toSerialize["min_esx_version"] = o.MinEsxVersion
	}
	if o.MinVcVersion != nil {
		toSerialize["min_vc_version"] = o.MinVcVersion
	}
	if o.ReferencedVcs != nil {
		toSerialize["referenced_vcs"] = o.ReferencedVcs
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableViewComposerMonitorDetails struct {
	value *ViewComposerMonitorDetails
	isSet bool
}

func (v NullableViewComposerMonitorDetails) Get() *ViewComposerMonitorDetails {
	return v.value
}

func (v *NullableViewComposerMonitorDetails) Set(val *ViewComposerMonitorDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableViewComposerMonitorDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableViewComposerMonitorDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableViewComposerMonitorDetails(val *ViewComposerMonitorDetails) *NullableViewComposerMonitorDetails {
	return &NullableViewComposerMonitorDetails{value: val, isSet: true}
}

func (v NullableViewComposerMonitorDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableViewComposerMonitorDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
