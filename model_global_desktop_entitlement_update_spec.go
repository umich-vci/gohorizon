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

// GlobalDesktopEntitlementUpdateSpec Information required to update Global Desktop Entitlement.
type GlobalDesktopEntitlementUpdateSpec struct {
	// Indicates whether users can have multiple sessions when accessed from different client devices, this is also called Class room mode and applicable only to floating user assignment. If value is set to true, the desktop pools that are associated with this Global Desktop Entitlement  must also allow users to have multiple sessions.
	AllowMultipleSessionsPerUser bool `json:"allow_multiple_sessions_per_user"`
	// Global Desktop Entitlement that can be used as backup for this Global Desktop Entitlement.
	BackupGdeId *string `json:"backup_gde_id,omitempty"`
	// Name of the category folder in the user's OS containing a shortcut to the entitlement. The property will not be set if the entitlement does not belong to a category.
	CategoryFolderName *string `json:"category_folder_name,omitempty"`
	// Indicates whether this global desktop entitlement is managed from cloud.
	CloudManaged bool `json:"cloud_managed"`
	// Connection server restrictions. This is a list of tags that access to the entitlement is restricted to. No list means that the entitlement can be accessed from any connection server.
	CsRestrictionTags *[]string `json:"cs_restriction_tags,omitempty"`
	// The default display protocol for the Global Desktop Entitlement. Clients connecting through this Global Desktop Entitlement that do not specify a protocol will use this value, not the value specified directly on the desktop pool to which they connect (if different). * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol.
	DefaultDisplayProtocol string `json:"default_display_protocol"`
	// Description of Global Desktop Entitlement.
	Description *string `json:"description,omitempty"`
	// Indicates whether users should see the hostname of the machine assigned to them instead of display_name when they connect using Horizon Client. This is applicable for dedicated Global Desktop Entitlements only. If no machine is assigned to the user then \"display_name (No machine assigned)\" will be displayed in the client.
	DisplayAssignedMachineName bool `json:"display_assigned_machine_name"`
	// Indicates whether users should see the alias of the machine assigned to them instead of display_name when they connect using Horizon Client. This is applicable for dedicated Global Desktop Entitlements only.
	DisplayMachineAlias bool `json:"display_machine_alias"`
	// Name that users will see when they connect using Horizon Client. If the display name is left blank, while creating or updating the Global Desktop Entitlement display name value will be same as name.
	DisplayName string `json:"display_name"`
	// Indicates whether client restrictions to be applied to Global Desktop Entitlement. Currently it is valid for RDSH pools.
	EnableClientRestrictions bool `json:"enable_client_restrictions"`
	// Indicates if this Global Desktop Entitlement is enabled.
	Enabled bool `json:"enabled"`
	// ID of the federated access group with which the global desktop entitlement is to be associated. They can also be used for delegated administration.
	FederatedAccessGroupId string `json:"federated_access_group_id"`
	// Indicates if automatic session clean up is enabled. This cannot be enabled when this Global Desktop Entitlement is associated with a Desktop pool that has dedicated user assignment.
	MultipleSessionAutoClean bool `json:"multiple_session_auto_clean"`
	// Unique name used to identify the Global Desktop Entitlement.
	Name string `json:"name"`
	// Indicates whether it should fail if a home site isn't defined for this Global Desktop Entitlement.
	RequireHomeSite bool `json:"require_home_site"`
	// Scope for this global desktop entitlement. Visibility and Placement policies are defined by this value. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set.
	Scope string `json:"scope"`
	// Session collaboration allows a user to share their remote session with other users. Blast must be configured as a supported protocol. Indicates if the desktop pools that are associated with this Global Desktop Entitlement must also have session collaboration enabled with enableCollaboration.
	SessionCollaborationEnabled bool `json:"session_collaboration_enabled"`
	// Locations of the category folder in the user's OS containing a shortcut to the desktop. The value must be set if category_folder_name is provided.
	ShortcutLocationsV2 *[]string `json:"shortcut_locations_v2,omitempty"`
	// Indicates whether a pod in the user's home site is used to start the search or the current site is used.
	UseHomeSite bool `json:"use_home_site"`
}

// NewGlobalDesktopEntitlementUpdateSpec instantiates a new GlobalDesktopEntitlementUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalDesktopEntitlementUpdateSpec(allowMultipleSessionsPerUser bool, cloudManaged bool, defaultDisplayProtocol string, displayAssignedMachineName bool, displayMachineAlias bool, displayName string, enableClientRestrictions bool, enabled bool, federatedAccessGroupId string, multipleSessionAutoClean bool, name string, requireHomeSite bool, scope string, sessionCollaborationEnabled bool, useHomeSite bool) *GlobalDesktopEntitlementUpdateSpec {
	this := GlobalDesktopEntitlementUpdateSpec{}
	this.AllowMultipleSessionsPerUser = allowMultipleSessionsPerUser
	this.CloudManaged = cloudManaged
	this.DefaultDisplayProtocol = defaultDisplayProtocol
	this.DisplayAssignedMachineName = displayAssignedMachineName
	this.DisplayMachineAlias = displayMachineAlias
	this.DisplayName = displayName
	this.EnableClientRestrictions = enableClientRestrictions
	this.Enabled = enabled
	this.FederatedAccessGroupId = federatedAccessGroupId
	this.MultipleSessionAutoClean = multipleSessionAutoClean
	this.Name = name
	this.RequireHomeSite = requireHomeSite
	this.Scope = scope
	this.SessionCollaborationEnabled = sessionCollaborationEnabled
	this.UseHomeSite = useHomeSite
	return &this
}

// NewGlobalDesktopEntitlementUpdateSpecWithDefaults instantiates a new GlobalDesktopEntitlementUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalDesktopEntitlementUpdateSpecWithDefaults() *GlobalDesktopEntitlementUpdateSpec {
	this := GlobalDesktopEntitlementUpdateSpec{}
	return &this
}

// GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetAllowMultipleSessionsPerUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowMultipleSessionsPerUser
}

// GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetAllowMultipleSessionsPerUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowMultipleSessionsPerUser, true
}

// SetAllowMultipleSessionsPerUser sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetAllowMultipleSessionsPerUser(v bool) {
	o.AllowMultipleSessionsPerUser = v
}

// GetBackupGdeId returns the BackupGdeId field value if set, zero value otherwise.
func (o *GlobalDesktopEntitlementUpdateSpec) GetBackupGdeId() string {
	if o == nil || o.BackupGdeId == nil {
		var ret string
		return ret
	}
	return *o.BackupGdeId
}

// GetBackupGdeIdOk returns a tuple with the BackupGdeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetBackupGdeIdOk() (*string, bool) {
	if o == nil || o.BackupGdeId == nil {
		return nil, false
	}
	return o.BackupGdeId, true
}

// HasBackupGdeId returns a boolean if a field has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) HasBackupGdeId() bool {
	if o != nil && o.BackupGdeId != nil {
		return true
	}

	return false
}

// SetBackupGdeId gets a reference to the given string and assigns it to the BackupGdeId field.
func (o *GlobalDesktopEntitlementUpdateSpec) SetBackupGdeId(v string) {
	o.BackupGdeId = &v
}

// GetCategoryFolderName returns the CategoryFolderName field value if set, zero value otherwise.
func (o *GlobalDesktopEntitlementUpdateSpec) GetCategoryFolderName() string {
	if o == nil || o.CategoryFolderName == nil {
		var ret string
		return ret
	}
	return *o.CategoryFolderName
}

// GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetCategoryFolderNameOk() (*string, bool) {
	if o == nil || o.CategoryFolderName == nil {
		return nil, false
	}
	return o.CategoryFolderName, true
}

// HasCategoryFolderName returns a boolean if a field has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) HasCategoryFolderName() bool {
	if o != nil && o.CategoryFolderName != nil {
		return true
	}

	return false
}

// SetCategoryFolderName gets a reference to the given string and assigns it to the CategoryFolderName field.
func (o *GlobalDesktopEntitlementUpdateSpec) SetCategoryFolderName(v string) {
	o.CategoryFolderName = &v
}

// GetCloudManaged returns the CloudManaged field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetCloudManaged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CloudManaged
}

// GetCloudManagedOk returns a tuple with the CloudManaged field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetCloudManagedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CloudManaged, true
}

// SetCloudManaged sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetCloudManaged(v bool) {
	o.CloudManaged = v
}

// GetCsRestrictionTags returns the CsRestrictionTags field value if set, zero value otherwise.
func (o *GlobalDesktopEntitlementUpdateSpec) GetCsRestrictionTags() []string {
	if o == nil || o.CsRestrictionTags == nil {
		var ret []string
		return ret
	}
	return *o.CsRestrictionTags
}

// GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetCsRestrictionTagsOk() (*[]string, bool) {
	if o == nil || o.CsRestrictionTags == nil {
		return nil, false
	}
	return o.CsRestrictionTags, true
}

// HasCsRestrictionTags returns a boolean if a field has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) HasCsRestrictionTags() bool {
	if o != nil && o.CsRestrictionTags != nil {
		return true
	}

	return false
}

// SetCsRestrictionTags gets a reference to the given []string and assigns it to the CsRestrictionTags field.
func (o *GlobalDesktopEntitlementUpdateSpec) SetCsRestrictionTags(v []string) {
	o.CsRestrictionTags = &v
}

// GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetDefaultDisplayProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultDisplayProtocol
}

// GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetDefaultDisplayProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDisplayProtocol, true
}

// SetDefaultDisplayProtocol sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetDefaultDisplayProtocol(v string) {
	o.DefaultDisplayProtocol = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GlobalDesktopEntitlementUpdateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GlobalDesktopEntitlementUpdateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayAssignedMachineName returns the DisplayAssignedMachineName field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayAssignedMachineName() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisplayAssignedMachineName
}

// GetDisplayAssignedMachineNameOk returns a tuple with the DisplayAssignedMachineName field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayAssignedMachineNameOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayAssignedMachineName, true
}

// SetDisplayAssignedMachineName sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetDisplayAssignedMachineName(v bool) {
	o.DisplayAssignedMachineName = v
}

// GetDisplayMachineAlias returns the DisplayMachineAlias field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayMachineAlias() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DisplayMachineAlias
}

// GetDisplayMachineAliasOk returns a tuple with the DisplayMachineAlias field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayMachineAliasOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayMachineAlias, true
}

// SetDisplayMachineAlias sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetDisplayMachineAlias(v bool) {
	o.DisplayMachineAlias = v
}

// GetDisplayName returns the DisplayName field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnableClientRestrictions returns the EnableClientRestrictions field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetEnableClientRestrictions() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnableClientRestrictions
}

// GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetEnableClientRestrictionsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnableClientRestrictions, true
}

// SetEnableClientRestrictions sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetEnableClientRestrictions(v bool) {
	o.EnableClientRestrictions = v
}

// GetEnabled returns the Enabled field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetEnabled(v bool) {
	o.Enabled = v
}

// GetFederatedAccessGroupId returns the FederatedAccessGroupId field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetFederatedAccessGroupId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FederatedAccessGroupId
}

// GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetFederatedAccessGroupIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FederatedAccessGroupId, true
}

// SetFederatedAccessGroupId sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetFederatedAccessGroupId(v string) {
	o.FederatedAccessGroupId = v
}

// GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetMultipleSessionAutoClean() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.MultipleSessionAutoClean
}

// GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetMultipleSessionAutoCleanOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultipleSessionAutoClean, true
}

// SetMultipleSessionAutoClean sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetMultipleSessionAutoClean(v bool) {
	o.MultipleSessionAutoClean = v
}

// GetName returns the Name field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetName(v string) {
	o.Name = v
}

// GetRequireHomeSite returns the RequireHomeSite field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetRequireHomeSite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RequireHomeSite
}

// GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetRequireHomeSiteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequireHomeSite, true
}

// SetRequireHomeSite sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetRequireHomeSite(v bool) {
	o.RequireHomeSite = v
}

// GetScope returns the Scope field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetScope(v string) {
	o.Scope = v
}

// GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetSessionCollaborationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SessionCollaborationEnabled
}

// GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetSessionCollaborationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionCollaborationEnabled, true
}

// SetSessionCollaborationEnabled sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetSessionCollaborationEnabled(v bool) {
	o.SessionCollaborationEnabled = v
}

// GetShortcutLocationsV2 returns the ShortcutLocationsV2 field value if set, zero value otherwise.
func (o *GlobalDesktopEntitlementUpdateSpec) GetShortcutLocationsV2() []string {
	if o == nil || o.ShortcutLocationsV2 == nil {
		var ret []string
		return ret
	}
	return *o.ShortcutLocationsV2
}

// GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetShortcutLocationsV2Ok() (*[]string, bool) {
	if o == nil || o.ShortcutLocationsV2 == nil {
		return nil, false
	}
	return o.ShortcutLocationsV2, true
}

// HasShortcutLocationsV2 returns a boolean if a field has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) HasShortcutLocationsV2() bool {
	if o != nil && o.ShortcutLocationsV2 != nil {
		return true
	}

	return false
}

// SetShortcutLocationsV2 gets a reference to the given []string and assigns it to the ShortcutLocationsV2 field.
func (o *GlobalDesktopEntitlementUpdateSpec) SetShortcutLocationsV2(v []string) {
	o.ShortcutLocationsV2 = &v
}

// GetUseHomeSite returns the UseHomeSite field value
func (o *GlobalDesktopEntitlementUpdateSpec) GetUseHomeSite() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UseHomeSite
}

// GetUseHomeSiteOk returns a tuple with the UseHomeSite field value
// and a boolean to check if the value has been set.
func (o *GlobalDesktopEntitlementUpdateSpec) GetUseHomeSiteOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseHomeSite, true
}

// SetUseHomeSite sets field value
func (o *GlobalDesktopEntitlementUpdateSpec) SetUseHomeSite(v bool) {
	o.UseHomeSite = v
}

func (o GlobalDesktopEntitlementUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allow_multiple_sessions_per_user"] = o.AllowMultipleSessionsPerUser
	}
	if o.BackupGdeId != nil {
		toSerialize["backup_gde_id"] = o.BackupGdeId
	}
	if o.CategoryFolderName != nil {
		toSerialize["category_folder_name"] = o.CategoryFolderName
	}
	if true {
		toSerialize["cloud_managed"] = o.CloudManaged
	}
	if o.CsRestrictionTags != nil {
		toSerialize["cs_restriction_tags"] = o.CsRestrictionTags
	}
	if true {
		toSerialize["default_display_protocol"] = o.DefaultDisplayProtocol
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["display_assigned_machine_name"] = o.DisplayAssignedMachineName
	}
	if true {
		toSerialize["display_machine_alias"] = o.DisplayMachineAlias
	}
	if true {
		toSerialize["display_name"] = o.DisplayName
	}
	if true {
		toSerialize["enable_client_restrictions"] = o.EnableClientRestrictions
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if true {
		toSerialize["federated_access_group_id"] = o.FederatedAccessGroupId
	}
	if true {
		toSerialize["multiple_session_auto_clean"] = o.MultipleSessionAutoClean
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["require_home_site"] = o.RequireHomeSite
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["session_collaboration_enabled"] = o.SessionCollaborationEnabled
	}
	if o.ShortcutLocationsV2 != nil {
		toSerialize["shortcut_locations_v2"] = o.ShortcutLocationsV2
	}
	if true {
		toSerialize["use_home_site"] = o.UseHomeSite
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalDesktopEntitlementUpdateSpec struct {
	value *GlobalDesktopEntitlementUpdateSpec
	isSet bool
}

func (v NullableGlobalDesktopEntitlementUpdateSpec) Get() *GlobalDesktopEntitlementUpdateSpec {
	return v.value
}

func (v *NullableGlobalDesktopEntitlementUpdateSpec) Set(val *GlobalDesktopEntitlementUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalDesktopEntitlementUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalDesktopEntitlementUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalDesktopEntitlementUpdateSpec(val *GlobalDesktopEntitlementUpdateSpec) *NullableGlobalDesktopEntitlementUpdateSpec {
	return &NullableGlobalDesktopEntitlementUpdateSpec{value: val, isSet: true}
}

func (v NullableGlobalDesktopEntitlementUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalDesktopEntitlementUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}