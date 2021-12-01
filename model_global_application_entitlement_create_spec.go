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

// GlobalApplicationEntitlementCreateSpec Information required to create global application entitlement.
type GlobalApplicationEntitlementCreateSpec struct {
	// Indicates whether the users can choose the protocol used. Default value is true.
	AllowUsersToChooseProtocol *bool `json:"allow_users_to_choose_protocol,omitempty"`
	// Name of the category folder in the user's OS containing a shortcut to the entitlement. The property will not be set if the entitlement does not belong to a category.
	CategoryFolderName *string `json:"category_folder_name,omitempty"`
	// List of connection server restriction tags to which the access to the global application entitlement is restricted. Empty or null list means that entitlement can be accessed from any connection server.
	CsRestrictionTags *[]string `json:"cs_restriction_tags,omitempty"`
	// The default display protocol for the global application entitlement. This can only be set to \"PCOIP\" or \"BLAST\". If this application's Farm's or desktop pool's allow_users_to_choose_protocol is set to false, then default_display_protocol must match that default_display_protocol of farm or desktop pool. Default value is \"PCOIP\". * PCOIP: PCoIP protocol. * BLAST: BLAST protocol.
	DefaultDisplayProtocol *string `json:"default_display_protocol,omitempty"`
	// Description of global application entitlement.
	Description *string `json:"description,omitempty"`
	// The display name is the name that users will see when they connect using Horizon View Client. If display_name is left blank, it defaults to name.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether client restrictions to be applied to global application entitlement, currently it is valid for published application pools. Default value is false.
	EnableClientRestrictions *bool `json:"enable_client_restrictions,omitempty"`
	// If value is set to true, this global application entitlement can be pre-launched. This property can be set to true only if multi_session_mode is set to \"DISABLED\".  Default value is false.
	EnablePreLaunch *bool `json:"enable_pre_launch,omitempty"`
	// Indicates whether the global application entitlement is enabled. Default value is true.
	Enabled *bool `json:"enabled,omitempty"`
	// ID of the federated access group with which the global application entitlement is to be associated. They can also be used for delegated administration.
	FederatedAccessGroupId string `json:"federated_access_group_id"`
	// Multi-session mode for this entitlement. A global application entitlement launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. Default value is \"DISABLED\" * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application.
	MultiSessionMode *string `json:"multi_session_mode,omitempty"`
	// Indicates whether automatic session clean up is enabled. Default value is false.
	MultipleSessionAutoClean *bool `json:"multiple_session_auto_clean,omitempty"`
	// Unique name used to identify the global application entitlement.
	Name string `json:"name"`
	// Indicates whether it should fail if a home site isn't defined for this global application entitlement. This property cannot be set to true if use_home_site is set to false. Default value is false.
	RequireHomeSite *bool `json:"require_home_site,omitempty"`
	// Scope for this global application entitlement. Visibility and Placement policies are defined by this value. Default value of \"ALL_SITES\". * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set.
	Scope *string `json:"scope,omitempty"`
	// Locations of the category folder in the user's OS containing a shortcut to the desktop. This property is required if category_folder_name is set.
	ShortcutLocationsV2 *[]string `json:"shortcut_locations_v2,omitempty"`
	// Indicates whether a pod in the user's home site is used to start the search or the current site is used. Default value is false.
	UseHomeSite *bool `json:"use_home_site,omitempty"`
}

// NewGlobalApplicationEntitlementCreateSpec instantiates a new GlobalApplicationEntitlementCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalApplicationEntitlementCreateSpec(federatedAccessGroupId string, name string) *GlobalApplicationEntitlementCreateSpec {
	this := GlobalApplicationEntitlementCreateSpec{}
	this.FederatedAccessGroupId = federatedAccessGroupId
	this.Name = name
	return &this
}

// NewGlobalApplicationEntitlementCreateSpecWithDefaults instantiates a new GlobalApplicationEntitlementCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalApplicationEntitlementCreateSpecWithDefaults() *GlobalApplicationEntitlementCreateSpec {
	this := GlobalApplicationEntitlementCreateSpec{}
	return &this
}

// GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetAllowUsersToChooseProtocol() bool {
	if o == nil || o.AllowUsersToChooseProtocol == nil {
		var ret bool
		return ret
	}
	return *o.AllowUsersToChooseProtocol
}

// GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetAllowUsersToChooseProtocolOk() (*bool, bool) {
	if o == nil || o.AllowUsersToChooseProtocol == nil {
		return nil, false
	}
	return o.AllowUsersToChooseProtocol, true
}

// HasAllowUsersToChooseProtocol returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasAllowUsersToChooseProtocol() bool {
	if o != nil && o.AllowUsersToChooseProtocol != nil {
		return true
	}

	return false
}

// SetAllowUsersToChooseProtocol gets a reference to the given bool and assigns it to the AllowUsersToChooseProtocol field.
func (o *GlobalApplicationEntitlementCreateSpec) SetAllowUsersToChooseProtocol(v bool) {
	o.AllowUsersToChooseProtocol = &v
}

// GetCategoryFolderName returns the CategoryFolderName field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetCategoryFolderName() string {
	if o == nil || o.CategoryFolderName == nil {
		var ret string
		return ret
	}
	return *o.CategoryFolderName
}

// GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetCategoryFolderNameOk() (*string, bool) {
	if o == nil || o.CategoryFolderName == nil {
		return nil, false
	}
	return o.CategoryFolderName, true
}

// HasCategoryFolderName returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasCategoryFolderName() bool {
	if o != nil && o.CategoryFolderName != nil {
		return true
	}

	return false
}

// SetCategoryFolderName gets a reference to the given string and assigns it to the CategoryFolderName field.
func (o *GlobalApplicationEntitlementCreateSpec) SetCategoryFolderName(v string) {
	o.CategoryFolderName = &v
}

// GetCsRestrictionTags returns the CsRestrictionTags field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetCsRestrictionTags() []string {
	if o == nil || o.CsRestrictionTags == nil {
		var ret []string
		return ret
	}
	return *o.CsRestrictionTags
}

// GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetCsRestrictionTagsOk() (*[]string, bool) {
	if o == nil || o.CsRestrictionTags == nil {
		return nil, false
	}
	return o.CsRestrictionTags, true
}

// HasCsRestrictionTags returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasCsRestrictionTags() bool {
	if o != nil && o.CsRestrictionTags != nil {
		return true
	}

	return false
}

// SetCsRestrictionTags gets a reference to the given []string and assigns it to the CsRestrictionTags field.
func (o *GlobalApplicationEntitlementCreateSpec) SetCsRestrictionTags(v []string) {
	o.CsRestrictionTags = &v
}

// GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetDefaultDisplayProtocol() string {
	if o == nil || o.DefaultDisplayProtocol == nil {
		var ret string
		return ret
	}
	return *o.DefaultDisplayProtocol
}

// GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetDefaultDisplayProtocolOk() (*string, bool) {
	if o == nil || o.DefaultDisplayProtocol == nil {
		return nil, false
	}
	return o.DefaultDisplayProtocol, true
}

// HasDefaultDisplayProtocol returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasDefaultDisplayProtocol() bool {
	if o != nil && o.DefaultDisplayProtocol != nil {
		return true
	}

	return false
}

// SetDefaultDisplayProtocol gets a reference to the given string and assigns it to the DefaultDisplayProtocol field.
func (o *GlobalApplicationEntitlementCreateSpec) SetDefaultDisplayProtocol(v string) {
	o.DefaultDisplayProtocol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GlobalApplicationEntitlementCreateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GlobalApplicationEntitlementCreateSpec) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnableClientRestrictions returns the EnableClientRestrictions field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetEnableClientRestrictions() bool {
	if o == nil || o.EnableClientRestrictions == nil {
		var ret bool
		return ret
	}
	return *o.EnableClientRestrictions
}

// GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetEnableClientRestrictionsOk() (*bool, bool) {
	if o == nil || o.EnableClientRestrictions == nil {
		return nil, false
	}
	return o.EnableClientRestrictions, true
}

// HasEnableClientRestrictions returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasEnableClientRestrictions() bool {
	if o != nil && o.EnableClientRestrictions != nil {
		return true
	}

	return false
}

// SetEnableClientRestrictions gets a reference to the given bool and assigns it to the EnableClientRestrictions field.
func (o *GlobalApplicationEntitlementCreateSpec) SetEnableClientRestrictions(v bool) {
	o.EnableClientRestrictions = &v
}

// GetEnablePreLaunch returns the EnablePreLaunch field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetEnablePreLaunch() bool {
	if o == nil || o.EnablePreLaunch == nil {
		var ret bool
		return ret
	}
	return *o.EnablePreLaunch
}

// GetEnablePreLaunchOk returns a tuple with the EnablePreLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetEnablePreLaunchOk() (*bool, bool) {
	if o == nil || o.EnablePreLaunch == nil {
		return nil, false
	}
	return o.EnablePreLaunch, true
}

// HasEnablePreLaunch returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasEnablePreLaunch() bool {
	if o != nil && o.EnablePreLaunch != nil {
		return true
	}

	return false
}

// SetEnablePreLaunch gets a reference to the given bool and assigns it to the EnablePreLaunch field.
func (o *GlobalApplicationEntitlementCreateSpec) SetEnablePreLaunch(v bool) {
	o.EnablePreLaunch = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GlobalApplicationEntitlementCreateSpec) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFederatedAccessGroupId returns the FederatedAccessGroupId field value
func (o *GlobalApplicationEntitlementCreateSpec) GetFederatedAccessGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FederatedAccessGroupId
}

// GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetFederatedAccessGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FederatedAccessGroupId, true
}

// SetFederatedAccessGroupId sets field value
func (o *GlobalApplicationEntitlementCreateSpec) SetFederatedAccessGroupId(v string) {
	o.FederatedAccessGroupId = v
}

// GetMultiSessionMode returns the MultiSessionMode field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetMultiSessionMode() string {
	if o == nil || o.MultiSessionMode == nil {
		var ret string
		return ret
	}
	return *o.MultiSessionMode
}

// GetMultiSessionModeOk returns a tuple with the MultiSessionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetMultiSessionModeOk() (*string, bool) {
	if o == nil || o.MultiSessionMode == nil {
		return nil, false
	}
	return o.MultiSessionMode, true
}

// HasMultiSessionMode returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasMultiSessionMode() bool {
	if o != nil && o.MultiSessionMode != nil {
		return true
	}

	return false
}

// SetMultiSessionMode gets a reference to the given string and assigns it to the MultiSessionMode field.
func (o *GlobalApplicationEntitlementCreateSpec) SetMultiSessionMode(v string) {
	o.MultiSessionMode = &v
}

// GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetMultipleSessionAutoClean() bool {
	if o == nil || o.MultipleSessionAutoClean == nil {
		var ret bool
		return ret
	}
	return *o.MultipleSessionAutoClean
}

// GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetMultipleSessionAutoCleanOk() (*bool, bool) {
	if o == nil || o.MultipleSessionAutoClean == nil {
		return nil, false
	}
	return o.MultipleSessionAutoClean, true
}

// HasMultipleSessionAutoClean returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasMultipleSessionAutoClean() bool {
	if o != nil && o.MultipleSessionAutoClean != nil {
		return true
	}

	return false
}

// SetMultipleSessionAutoClean gets a reference to the given bool and assigns it to the MultipleSessionAutoClean field.
func (o *GlobalApplicationEntitlementCreateSpec) SetMultipleSessionAutoClean(v bool) {
	o.MultipleSessionAutoClean = &v
}

// GetName returns the Name field value
func (o *GlobalApplicationEntitlementCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GlobalApplicationEntitlementCreateSpec) SetName(v string) {
	o.Name = v
}

// GetRequireHomeSite returns the RequireHomeSite field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetRequireHomeSite() bool {
	if o == nil || o.RequireHomeSite == nil {
		var ret bool
		return ret
	}
	return *o.RequireHomeSite
}

// GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetRequireHomeSiteOk() (*bool, bool) {
	if o == nil || o.RequireHomeSite == nil {
		return nil, false
	}
	return o.RequireHomeSite, true
}

// HasRequireHomeSite returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasRequireHomeSite() bool {
	if o != nil && o.RequireHomeSite != nil {
		return true
	}

	return false
}

// SetRequireHomeSite gets a reference to the given bool and assigns it to the RequireHomeSite field.
func (o *GlobalApplicationEntitlementCreateSpec) SetRequireHomeSite(v bool) {
	o.RequireHomeSite = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *GlobalApplicationEntitlementCreateSpec) SetScope(v string) {
	o.Scope = &v
}

// GetShortcutLocationsV2 returns the ShortcutLocationsV2 field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetShortcutLocationsV2() []string {
	if o == nil || o.ShortcutLocationsV2 == nil {
		var ret []string
		return ret
	}
	return *o.ShortcutLocationsV2
}

// GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetShortcutLocationsV2Ok() (*[]string, bool) {
	if o == nil || o.ShortcutLocationsV2 == nil {
		return nil, false
	}
	return o.ShortcutLocationsV2, true
}

// HasShortcutLocationsV2 returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasShortcutLocationsV2() bool {
	if o != nil && o.ShortcutLocationsV2 != nil {
		return true
	}

	return false
}

// SetShortcutLocationsV2 gets a reference to the given []string and assigns it to the ShortcutLocationsV2 field.
func (o *GlobalApplicationEntitlementCreateSpec) SetShortcutLocationsV2(v []string) {
	o.ShortcutLocationsV2 = &v
}

// GetUseHomeSite returns the UseHomeSite field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementCreateSpec) GetUseHomeSite() bool {
	if o == nil || o.UseHomeSite == nil {
		var ret bool
		return ret
	}
	return *o.UseHomeSite
}

// GetUseHomeSiteOk returns a tuple with the UseHomeSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementCreateSpec) GetUseHomeSiteOk() (*bool, bool) {
	if o == nil || o.UseHomeSite == nil {
		return nil, false
	}
	return o.UseHomeSite, true
}

// HasUseHomeSite returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementCreateSpec) HasUseHomeSite() bool {
	if o != nil && o.UseHomeSite != nil {
		return true
	}

	return false
}

// SetUseHomeSite gets a reference to the given bool and assigns it to the UseHomeSite field.
func (o *GlobalApplicationEntitlementCreateSpec) SetUseHomeSite(v bool) {
	o.UseHomeSite = &v
}

func (o GlobalApplicationEntitlementCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowUsersToChooseProtocol != nil {
		toSerialize["allow_users_to_choose_protocol"] = o.AllowUsersToChooseProtocol
	}
	if o.CategoryFolderName != nil {
		toSerialize["category_folder_name"] = o.CategoryFolderName
	}
	if o.CsRestrictionTags != nil {
		toSerialize["cs_restriction_tags"] = o.CsRestrictionTags
	}
	if o.DefaultDisplayProtocol != nil {
		toSerialize["default_display_protocol"] = o.DefaultDisplayProtocol
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["display_name"] = o.DisplayName
	}
	if o.EnableClientRestrictions != nil {
		toSerialize["enable_client_restrictions"] = o.EnableClientRestrictions
	}
	if o.EnablePreLaunch != nil {
		toSerialize["enable_pre_launch"] = o.EnablePreLaunch
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["federated_access_group_id"] = o.FederatedAccessGroupId
	}
	if o.MultiSessionMode != nil {
		toSerialize["multi_session_mode"] = o.MultiSessionMode
	}
	if o.MultipleSessionAutoClean != nil {
		toSerialize["multiple_session_auto_clean"] = o.MultipleSessionAutoClean
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.RequireHomeSite != nil {
		toSerialize["require_home_site"] = o.RequireHomeSite
	}
	if o.Scope != nil {
		toSerialize["scope"] = o.Scope
	}
	if o.ShortcutLocationsV2 != nil {
		toSerialize["shortcut_locations_v2"] = o.ShortcutLocationsV2
	}
	if o.UseHomeSite != nil {
		toSerialize["use_home_site"] = o.UseHomeSite
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalApplicationEntitlementCreateSpec struct {
	value *GlobalApplicationEntitlementCreateSpec
	isSet bool
}

func (v NullableGlobalApplicationEntitlementCreateSpec) Get() *GlobalApplicationEntitlementCreateSpec {
	return v.value
}

func (v *NullableGlobalApplicationEntitlementCreateSpec) Set(val *GlobalApplicationEntitlementCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalApplicationEntitlementCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalApplicationEntitlementCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalApplicationEntitlementCreateSpec(val *GlobalApplicationEntitlementCreateSpec) *NullableGlobalApplicationEntitlementCreateSpec {
	return &NullableGlobalApplicationEntitlementCreateSpec{value: val, isSet: true}
}

func (v NullableGlobalApplicationEntitlementCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalApplicationEntitlementCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}