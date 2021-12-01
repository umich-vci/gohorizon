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

// GlobalApplicationEntitlementSummaryV2 Summary Information about Global Application Entitlement.
type GlobalApplicationEntitlementSummaryV2 struct {
	// Indicates whether the users can choose the protocol used. If set to true, the application pools that are associated with this Global Application Entitlement must also allow users to choose display protocol with allowUsersToChooseProtocol. Supported Filters: 'Equals'.
	AllowUsersToChooseProtocol *bool `json:"allow_users_to_choose_protocol,omitempty"`
	// Indicates the Global Application Entitlement that can be used as backup for this Global Application Entitlement. Supported Filters: 'Equals'.
	BackupGaeId *string `json:"backup_gae_id,omitempty"`
	// Name of the category folder in the user's OS containing a shortcut to the entitlement. Unset if the entitlement does not belong to a category.
	CategoryFolderName *string `json:"category_folder_name,omitempty"`
	// Connection server restrictions. This is a list of tags that access to the entitlement is restricted to. No list means that the entitlement can be accessed from any connection server.
	CsRestrictionTags *[]string `json:"cs_restriction_tags,omitempty"`
	// The default display protocol for the Global Application Entitlement. Must be a protocol in the supportedDisplayProtocols list. Clients connecting through this Global Application Entitlement that do not specify a protocol will use this value, not the value specified directly on the application pool to which they connect (if different). Supported Filters: 'Equals'. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol.
	DefaultDisplayProtocol *string `json:"default_display_protocol,omitempty"`
	// Description of Global Application Entitlement. This property has a maximum length of 1024 characters. Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Description *string `json:"description,omitempty"`
	// Name that users will see when they connect using Horizon Client. If the display name is left blank, while creating or updating the Global Application Entitlement display name value will be same as name. This property has a maximum length of 64 characters. Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether client restrictions to be applied to Global Application Entitlement. Currently it is valid for RDSH pools. Supported Filters: 'Equals'.
	EnableClientRestrictions *bool `json:"enable_client_restrictions,omitempty"`
	// Indicates whether Global Application Entitlement can be pre-launched Supported Filters: 'Equals'.
	EnablePreLaunch *bool `json:"enable_pre_launch,omitempty"`
	// Indicates if this Global Application Entitlement is enabled. Supported Filters: 'Equals'.
	Enabled *bool `json:"enabled,omitempty"`
	// This represents id of the federated access group associated with the global application entitlement.<br> Supported Filters: 'Equals'.
	FederatedAccessGroupId *string `json:"federated_access_group_id,omitempty"`
	// Unique ID representing this Global Application Entitlement. Supported Filters: 'Equals'.
	Id *string `json:"id,omitempty"`
	// Multi-session mode for the application pool. An application launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. Multi-session mode should be disabled when pre-launch is enabled. Supported Filters: 'Equals'. * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application.
	MultiSessionMode *string `json:"multi_session_mode,omitempty"`
	// Indicates if automatic session clean up is enabled. This cannot be enabled when this Global Application Entitlement is associated with a Application pool that has dedicated user assignment. Supported Filters: 'Equals'.
	MultipleSessionAutoClean *bool `json:"multiple_session_auto_clean,omitempty"`
	// Unique name used to identify the Global Application Entitlement. This property has a maximum length of 64 characters. Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Name *string `json:"name,omitempty"`
	// Indicates the Global Application Entitlement for which this Global Application Entitlement acts as backup.
	PrimaryGaeId *string `json:"primary_gae_id,omitempty"`
	// Indicates whether we fail if a home site isn't defined for this Global Application Entitlement. Supported Filters: 'Equals'.
	RequireHomeSite *bool `json:"require_home_site,omitempty"`
	// Scope for this global application entitlement. Visibility and Placement policies are defined by this value. Supported Filters: 'Equals'. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set.
	Scope *string `json:"scope,omitempty"`
	// Locations of the category folder in the user's OS containing a shortcut to the application.
	ShortcutLocationsV2 *[]string `json:"shortcut_locations_v2,omitempty"`
	// The set of supported display protocols for the Global Application Entitlement. All the application pools associated with this Global Application Entitlement must support these protocols supportedDisplayProtocols . Clients connecting through this Global Application Entitlement that are allowed to select their protocol will see these display protocol options.
	SupportedDisplayProtocols *[]string `json:"supported_display_protocols,omitempty"`
	// Indicates whether a pod in the user's home site is used to start the search or the current site is used. Supported Filters: 'Equals'.
	UseHomeSite *bool `json:"use_home_site,omitempty"`
}

// NewGlobalApplicationEntitlementSummaryV2 instantiates a new GlobalApplicationEntitlementSummaryV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalApplicationEntitlementSummaryV2() *GlobalApplicationEntitlementSummaryV2 {
	this := GlobalApplicationEntitlementSummaryV2{}
	return &this
}

// NewGlobalApplicationEntitlementSummaryV2WithDefaults instantiates a new GlobalApplicationEntitlementSummaryV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalApplicationEntitlementSummaryV2WithDefaults() *GlobalApplicationEntitlementSummaryV2 {
	this := GlobalApplicationEntitlementSummaryV2{}
	return &this
}

// GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetAllowUsersToChooseProtocol() bool {
	if o == nil || o.AllowUsersToChooseProtocol == nil {
		var ret bool
		return ret
	}
	return *o.AllowUsersToChooseProtocol
}

// GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetAllowUsersToChooseProtocolOk() (*bool, bool) {
	if o == nil || o.AllowUsersToChooseProtocol == nil {
		return nil, false
	}
	return o.AllowUsersToChooseProtocol, true
}

// HasAllowUsersToChooseProtocol returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasAllowUsersToChooseProtocol() bool {
	if o != nil && o.AllowUsersToChooseProtocol != nil {
		return true
	}

	return false
}

// SetAllowUsersToChooseProtocol gets a reference to the given bool and assigns it to the AllowUsersToChooseProtocol field.
func (o *GlobalApplicationEntitlementSummaryV2) SetAllowUsersToChooseProtocol(v bool) {
	o.AllowUsersToChooseProtocol = &v
}

// GetBackupGaeId returns the BackupGaeId field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetBackupGaeId() string {
	if o == nil || o.BackupGaeId == nil {
		var ret string
		return ret
	}
	return *o.BackupGaeId
}

// GetBackupGaeIdOk returns a tuple with the BackupGaeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetBackupGaeIdOk() (*string, bool) {
	if o == nil || o.BackupGaeId == nil {
		return nil, false
	}
	return o.BackupGaeId, true
}

// HasBackupGaeId returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasBackupGaeId() bool {
	if o != nil && o.BackupGaeId != nil {
		return true
	}

	return false
}

// SetBackupGaeId gets a reference to the given string and assigns it to the BackupGaeId field.
func (o *GlobalApplicationEntitlementSummaryV2) SetBackupGaeId(v string) {
	o.BackupGaeId = &v
}

// GetCategoryFolderName returns the CategoryFolderName field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetCategoryFolderName() string {
	if o == nil || o.CategoryFolderName == nil {
		var ret string
		return ret
	}
	return *o.CategoryFolderName
}

// GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetCategoryFolderNameOk() (*string, bool) {
	if o == nil || o.CategoryFolderName == nil {
		return nil, false
	}
	return o.CategoryFolderName, true
}

// HasCategoryFolderName returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasCategoryFolderName() bool {
	if o != nil && o.CategoryFolderName != nil {
		return true
	}

	return false
}

// SetCategoryFolderName gets a reference to the given string and assigns it to the CategoryFolderName field.
func (o *GlobalApplicationEntitlementSummaryV2) SetCategoryFolderName(v string) {
	o.CategoryFolderName = &v
}

// GetCsRestrictionTags returns the CsRestrictionTags field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetCsRestrictionTags() []string {
	if o == nil || o.CsRestrictionTags == nil {
		var ret []string
		return ret
	}
	return *o.CsRestrictionTags
}

// GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetCsRestrictionTagsOk() (*[]string, bool) {
	if o == nil || o.CsRestrictionTags == nil {
		return nil, false
	}
	return o.CsRestrictionTags, true
}

// HasCsRestrictionTags returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasCsRestrictionTags() bool {
	if o != nil && o.CsRestrictionTags != nil {
		return true
	}

	return false
}

// SetCsRestrictionTags gets a reference to the given []string and assigns it to the CsRestrictionTags field.
func (o *GlobalApplicationEntitlementSummaryV2) SetCsRestrictionTags(v []string) {
	o.CsRestrictionTags = &v
}

// GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetDefaultDisplayProtocol() string {
	if o == nil || o.DefaultDisplayProtocol == nil {
		var ret string
		return ret
	}
	return *o.DefaultDisplayProtocol
}

// GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetDefaultDisplayProtocolOk() (*string, bool) {
	if o == nil || o.DefaultDisplayProtocol == nil {
		return nil, false
	}
	return o.DefaultDisplayProtocol, true
}

// HasDefaultDisplayProtocol returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasDefaultDisplayProtocol() bool {
	if o != nil && o.DefaultDisplayProtocol != nil {
		return true
	}

	return false
}

// SetDefaultDisplayProtocol gets a reference to the given string and assigns it to the DefaultDisplayProtocol field.
func (o *GlobalApplicationEntitlementSummaryV2) SetDefaultDisplayProtocol(v string) {
	o.DefaultDisplayProtocol = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GlobalApplicationEntitlementSummaryV2) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *GlobalApplicationEntitlementSummaryV2) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnableClientRestrictions returns the EnableClientRestrictions field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetEnableClientRestrictions() bool {
	if o == nil || o.EnableClientRestrictions == nil {
		var ret bool
		return ret
	}
	return *o.EnableClientRestrictions
}

// GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetEnableClientRestrictionsOk() (*bool, bool) {
	if o == nil || o.EnableClientRestrictions == nil {
		return nil, false
	}
	return o.EnableClientRestrictions, true
}

// HasEnableClientRestrictions returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasEnableClientRestrictions() bool {
	if o != nil && o.EnableClientRestrictions != nil {
		return true
	}

	return false
}

// SetEnableClientRestrictions gets a reference to the given bool and assigns it to the EnableClientRestrictions field.
func (o *GlobalApplicationEntitlementSummaryV2) SetEnableClientRestrictions(v bool) {
	o.EnableClientRestrictions = &v
}

// GetEnablePreLaunch returns the EnablePreLaunch field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetEnablePreLaunch() bool {
	if o == nil || o.EnablePreLaunch == nil {
		var ret bool
		return ret
	}
	return *o.EnablePreLaunch
}

// GetEnablePreLaunchOk returns a tuple with the EnablePreLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetEnablePreLaunchOk() (*bool, bool) {
	if o == nil || o.EnablePreLaunch == nil {
		return nil, false
	}
	return o.EnablePreLaunch, true
}

// HasEnablePreLaunch returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasEnablePreLaunch() bool {
	if o != nil && o.EnablePreLaunch != nil {
		return true
	}

	return false
}

// SetEnablePreLaunch gets a reference to the given bool and assigns it to the EnablePreLaunch field.
func (o *GlobalApplicationEntitlementSummaryV2) SetEnablePreLaunch(v bool) {
	o.EnablePreLaunch = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *GlobalApplicationEntitlementSummaryV2) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFederatedAccessGroupId returns the FederatedAccessGroupId field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetFederatedAccessGroupId() string {
	if o == nil || o.FederatedAccessGroupId == nil {
		var ret string
		return ret
	}
	return *o.FederatedAccessGroupId
}

// GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetFederatedAccessGroupIdOk() (*string, bool) {
	if o == nil || o.FederatedAccessGroupId == nil {
		return nil, false
	}
	return o.FederatedAccessGroupId, true
}

// HasFederatedAccessGroupId returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasFederatedAccessGroupId() bool {
	if o != nil && o.FederatedAccessGroupId != nil {
		return true
	}

	return false
}

// SetFederatedAccessGroupId gets a reference to the given string and assigns it to the FederatedAccessGroupId field.
func (o *GlobalApplicationEntitlementSummaryV2) SetFederatedAccessGroupId(v string) {
	o.FederatedAccessGroupId = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GlobalApplicationEntitlementSummaryV2) SetId(v string) {
	o.Id = &v
}

// GetMultiSessionMode returns the MultiSessionMode field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetMultiSessionMode() string {
	if o == nil || o.MultiSessionMode == nil {
		var ret string
		return ret
	}
	return *o.MultiSessionMode
}

// GetMultiSessionModeOk returns a tuple with the MultiSessionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetMultiSessionModeOk() (*string, bool) {
	if o == nil || o.MultiSessionMode == nil {
		return nil, false
	}
	return o.MultiSessionMode, true
}

// HasMultiSessionMode returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasMultiSessionMode() bool {
	if o != nil && o.MultiSessionMode != nil {
		return true
	}

	return false
}

// SetMultiSessionMode gets a reference to the given string and assigns it to the MultiSessionMode field.
func (o *GlobalApplicationEntitlementSummaryV2) SetMultiSessionMode(v string) {
	o.MultiSessionMode = &v
}

// GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetMultipleSessionAutoClean() bool {
	if o == nil || o.MultipleSessionAutoClean == nil {
		var ret bool
		return ret
	}
	return *o.MultipleSessionAutoClean
}

// GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetMultipleSessionAutoCleanOk() (*bool, bool) {
	if o == nil || o.MultipleSessionAutoClean == nil {
		return nil, false
	}
	return o.MultipleSessionAutoClean, true
}

// HasMultipleSessionAutoClean returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasMultipleSessionAutoClean() bool {
	if o != nil && o.MultipleSessionAutoClean != nil {
		return true
	}

	return false
}

// SetMultipleSessionAutoClean gets a reference to the given bool and assigns it to the MultipleSessionAutoClean field.
func (o *GlobalApplicationEntitlementSummaryV2) SetMultipleSessionAutoClean(v bool) {
	o.MultipleSessionAutoClean = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GlobalApplicationEntitlementSummaryV2) SetName(v string) {
	o.Name = &v
}

// GetPrimaryGaeId returns the PrimaryGaeId field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetPrimaryGaeId() string {
	if o == nil || o.PrimaryGaeId == nil {
		var ret string
		return ret
	}
	return *o.PrimaryGaeId
}

// GetPrimaryGaeIdOk returns a tuple with the PrimaryGaeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetPrimaryGaeIdOk() (*string, bool) {
	if o == nil || o.PrimaryGaeId == nil {
		return nil, false
	}
	return o.PrimaryGaeId, true
}

// HasPrimaryGaeId returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasPrimaryGaeId() bool {
	if o != nil && o.PrimaryGaeId != nil {
		return true
	}

	return false
}

// SetPrimaryGaeId gets a reference to the given string and assigns it to the PrimaryGaeId field.
func (o *GlobalApplicationEntitlementSummaryV2) SetPrimaryGaeId(v string) {
	o.PrimaryGaeId = &v
}

// GetRequireHomeSite returns the RequireHomeSite field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetRequireHomeSite() bool {
	if o == nil || o.RequireHomeSite == nil {
		var ret bool
		return ret
	}
	return *o.RequireHomeSite
}

// GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetRequireHomeSiteOk() (*bool, bool) {
	if o == nil || o.RequireHomeSite == nil {
		return nil, false
	}
	return o.RequireHomeSite, true
}

// HasRequireHomeSite returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasRequireHomeSite() bool {
	if o != nil && o.RequireHomeSite != nil {
		return true
	}

	return false
}

// SetRequireHomeSite gets a reference to the given bool and assigns it to the RequireHomeSite field.
func (o *GlobalApplicationEntitlementSummaryV2) SetRequireHomeSite(v bool) {
	o.RequireHomeSite = &v
}

// GetScope returns the Scope field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetScope() string {
	if o == nil || o.Scope == nil {
		var ret string
		return ret
	}
	return *o.Scope
}

// GetScopeOk returns a tuple with the Scope field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetScopeOk() (*string, bool) {
	if o == nil || o.Scope == nil {
		return nil, false
	}
	return o.Scope, true
}

// HasScope returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasScope() bool {
	if o != nil && o.Scope != nil {
		return true
	}

	return false
}

// SetScope gets a reference to the given string and assigns it to the Scope field.
func (o *GlobalApplicationEntitlementSummaryV2) SetScope(v string) {
	o.Scope = &v
}

// GetShortcutLocationsV2 returns the ShortcutLocationsV2 field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetShortcutLocationsV2() []string {
	if o == nil || o.ShortcutLocationsV2 == nil {
		var ret []string
		return ret
	}
	return *o.ShortcutLocationsV2
}

// GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetShortcutLocationsV2Ok() (*[]string, bool) {
	if o == nil || o.ShortcutLocationsV2 == nil {
		return nil, false
	}
	return o.ShortcutLocationsV2, true
}

// HasShortcutLocationsV2 returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasShortcutLocationsV2() bool {
	if o != nil && o.ShortcutLocationsV2 != nil {
		return true
	}

	return false
}

// SetShortcutLocationsV2 gets a reference to the given []string and assigns it to the ShortcutLocationsV2 field.
func (o *GlobalApplicationEntitlementSummaryV2) SetShortcutLocationsV2(v []string) {
	o.ShortcutLocationsV2 = &v
}

// GetSupportedDisplayProtocols returns the SupportedDisplayProtocols field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetSupportedDisplayProtocols() []string {
	if o == nil || o.SupportedDisplayProtocols == nil {
		var ret []string
		return ret
	}
	return *o.SupportedDisplayProtocols
}

// GetSupportedDisplayProtocolsOk returns a tuple with the SupportedDisplayProtocols field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetSupportedDisplayProtocolsOk() (*[]string, bool) {
	if o == nil || o.SupportedDisplayProtocols == nil {
		return nil, false
	}
	return o.SupportedDisplayProtocols, true
}

// HasSupportedDisplayProtocols returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasSupportedDisplayProtocols() bool {
	if o != nil && o.SupportedDisplayProtocols != nil {
		return true
	}

	return false
}

// SetSupportedDisplayProtocols gets a reference to the given []string and assigns it to the SupportedDisplayProtocols field.
func (o *GlobalApplicationEntitlementSummaryV2) SetSupportedDisplayProtocols(v []string) {
	o.SupportedDisplayProtocols = &v
}

// GetUseHomeSite returns the UseHomeSite field value if set, zero value otherwise.
func (o *GlobalApplicationEntitlementSummaryV2) GetUseHomeSite() bool {
	if o == nil || o.UseHomeSite == nil {
		var ret bool
		return ret
	}
	return *o.UseHomeSite
}

// GetUseHomeSiteOk returns a tuple with the UseHomeSite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalApplicationEntitlementSummaryV2) GetUseHomeSiteOk() (*bool, bool) {
	if o == nil || o.UseHomeSite == nil {
		return nil, false
	}
	return o.UseHomeSite, true
}

// HasUseHomeSite returns a boolean if a field has been set.
func (o *GlobalApplicationEntitlementSummaryV2) HasUseHomeSite() bool {
	if o != nil && o.UseHomeSite != nil {
		return true
	}

	return false
}

// SetUseHomeSite gets a reference to the given bool and assigns it to the UseHomeSite field.
func (o *GlobalApplicationEntitlementSummaryV2) SetUseHomeSite(v bool) {
	o.UseHomeSite = &v
}

func (o GlobalApplicationEntitlementSummaryV2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowUsersToChooseProtocol != nil {
		toSerialize["allow_users_to_choose_protocol"] = o.AllowUsersToChooseProtocol
	}
	if o.BackupGaeId != nil {
		toSerialize["backup_gae_id"] = o.BackupGaeId
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
	if o.FederatedAccessGroupId != nil {
		toSerialize["federated_access_group_id"] = o.FederatedAccessGroupId
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MultiSessionMode != nil {
		toSerialize["multi_session_mode"] = o.MultiSessionMode
	}
	if o.MultipleSessionAutoClean != nil {
		toSerialize["multiple_session_auto_clean"] = o.MultipleSessionAutoClean
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.PrimaryGaeId != nil {
		toSerialize["primary_gae_id"] = o.PrimaryGaeId
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
	if o.SupportedDisplayProtocols != nil {
		toSerialize["supported_display_protocols"] = o.SupportedDisplayProtocols
	}
	if o.UseHomeSite != nil {
		toSerialize["use_home_site"] = o.UseHomeSite
	}
	return json.Marshal(toSerialize)
}

type NullableGlobalApplicationEntitlementSummaryV2 struct {
	value *GlobalApplicationEntitlementSummaryV2
	isSet bool
}

func (v NullableGlobalApplicationEntitlementSummaryV2) Get() *GlobalApplicationEntitlementSummaryV2 {
	return v.value
}

func (v *NullableGlobalApplicationEntitlementSummaryV2) Set(val *GlobalApplicationEntitlementSummaryV2) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalApplicationEntitlementSummaryV2) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalApplicationEntitlementSummaryV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalApplicationEntitlementSummaryV2(val *GlobalApplicationEntitlementSummaryV2) *NullableGlobalApplicationEntitlementSummaryV2 {
	return &NullableGlobalApplicationEntitlementSummaryV2{value: val, isSet: true}
}

func (v NullableGlobalApplicationEntitlementSummaryV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalApplicationEntitlementSummaryV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
