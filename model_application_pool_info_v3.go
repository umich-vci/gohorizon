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

// ApplicationPoolInfoV3 Information related to Application Pool.<br>List API returning this information can use search filter query to filter on specific fields supported by filters.<br>Supported Filters: 'And', 'Or', 'Equals', 'StartsWith', 'Contains' and 'EndsWith'.<br>See the field description to know the filter types it supports.
type ApplicationPoolInfoV3 struct {
	// Access groups can organize the entities (like application pools, desktop pools) in the organization. They can also be used for delegated administration. For application pool, this is the same as that of the farm or desktop pool that the application pool belongs to.<br>Supported Filters: 'Equals'.
	AccessGroupId    *string                      `json:"access_group_id,omitempty"`
	AntiAffinityData *ApplicationAntiAffinityData `json:"anti_affinity_data,omitempty"`
	// Name of the category folder in the user's OS containing a shortcut to the application. Unset if the application does not belong to a category.
	CategoryFolderName *string `json:"category_folder_name,omitempty"`
	// Indicates whether the application pool is cloud brokered.<br>Supported Filters: 'Equals'.
	CloudBrokered *bool `json:"cloud_brokered,omitempty"`
	// Connection server restrictions. Application pool can be accessed from only those connection server instances that have a matching tag in this list. Null or empty list means that the application pool can be accessed from any connection server.
	CsRestrictionTags *[]string `json:"cs_restriction_tags,omitempty"`
	// List of customized icon IDs associated with the application which the user has configured.
	CustomizedIconIds *[]string `json:"customized_icon_ids,omitempty"`
	// Notes about the application pool.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Description *string `json:"description,omitempty"`
	// ID of the desktop pool from which this application pool is created. Either this or farm id will be set.<br>Supported Filters: 'Equals'.
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`
	// The display name is the name that users will see when they connect to view client. If the display name is left blank, it defaults to name.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	DisplayName *string `json:"display_name,omitempty"`
	// Indicates whether client restrictions are to be applied to application pool. Currently it is valid for application pool created from farm.<br>Supported Filters: 'Equals'.
	EnableClientRestrictions *bool `json:"enable_client_restrictions,omitempty"`
	// Whether to pre-launch the application.<br>Supported Filters: 'Equals'.
	EnablePreLaunch *bool `json:"enable_pre_launch,omitempty"`
	// Indicates whether the application pool is enabled.<br>Supported Filters: 'Equals'.
	Enabled *bool `json:"enabled,omitempty"`
	// Path to application executable.<br>Supported Filters: 'Equals', 'StartsWith', 'EndsWith' and 'Contains'.
	ExecutablePath *string `json:"executable_path,omitempty"`
	// ID of the farm from which this application pool is created. Either this or desktop pool id will be set.<br>Supported Filters: 'Equals'.
	FarmId *string `json:"farm_id,omitempty"`
	// Global application entitlement for this application pool. Caller should have permission to FEDERATED_LDAP_VIEW privilege for this field to be populated or to use in filter.<br>Supported Filters: 'Equals'.
	GlobalApplicationEntitlementId *string `json:"global_application_entitlement_id,omitempty"`
	// List of icon IDs associated with the application which are fetched from the agent.
	IconIds *[]string `json:"icon_ids,omitempty"`
	// Unique ID representing application pool.<br>Supported Filters: 'Equals'.
	Id *string `json:"id,omitempty"`
	// Maximum number of multi-sessions a user can have in this application pool. This property is required if multi-session mode is set to \"ENABLED_DEFAULT_OFF\", \"ENABLED_DEFAULT_ON\", or \"ENABLED_ENFORCED\"
	MaxMultiSessions *int32 `json:"max_multi_sessions,omitempty"`
	// Multi-session mode for the application pool. An application launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. Multi-session mode should be disabled when pre-launch is enabled.<br>Supported Filters: 'Equals'. * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application.
	MultiSessionMode *string `json:"multi_session_mode,omitempty"`
	// The application name is the unique identifier used to identify this application pool. This property must contain only alphanumerics, underscores, and dashes. The maximum length is 64 characters.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Name *string `json:"name,omitempty"`
	// Parameters to pass to application when launching.
	Parameters *string `json:"parameters,omitempty"`
	// Application publisher.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Publisher *string `json:"publisher,omitempty"`
	// Locations of the category folder in the user's OS containing a shortcut to the application. The value must be set if category folder name is provided.
	ShortcutLocations *[]string `json:"shortcut_locations,omitempty"`
	// Starting folder for application.
	StartFolder            *string                            `json:"start_folder,omitempty"`
	SupportedFileTypesData *ApplicationSupportedFileTypesData `json:"supported_file_types_data,omitempty"`
	// Application version.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Version *string `json:"version,omitempty"`
}

// NewApplicationPoolInfoV3 instantiates a new ApplicationPoolInfoV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationPoolInfoV3() *ApplicationPoolInfoV3 {
	this := ApplicationPoolInfoV3{}
	return &this
}

// NewApplicationPoolInfoV3WithDefaults instantiates a new ApplicationPoolInfoV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationPoolInfoV3WithDefaults() *ApplicationPoolInfoV3 {
	this := ApplicationPoolInfoV3{}
	return &this
}

// GetAccessGroupId returns the AccessGroupId field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetAccessGroupId() string {
	if o == nil || o.AccessGroupId == nil {
		var ret string
		return ret
	}
	return *o.AccessGroupId
}

// GetAccessGroupIdOk returns a tuple with the AccessGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetAccessGroupIdOk() (*string, bool) {
	if o == nil || o.AccessGroupId == nil {
		return nil, false
	}
	return o.AccessGroupId, true
}

// HasAccessGroupId returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasAccessGroupId() bool {
	if o != nil && o.AccessGroupId != nil {
		return true
	}

	return false
}

// SetAccessGroupId gets a reference to the given string and assigns it to the AccessGroupId field.
func (o *ApplicationPoolInfoV3) SetAccessGroupId(v string) {
	o.AccessGroupId = &v
}

// GetAntiAffinityData returns the AntiAffinityData field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetAntiAffinityData() ApplicationAntiAffinityData {
	if o == nil || o.AntiAffinityData == nil {
		var ret ApplicationAntiAffinityData
		return ret
	}
	return *o.AntiAffinityData
}

// GetAntiAffinityDataOk returns a tuple with the AntiAffinityData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetAntiAffinityDataOk() (*ApplicationAntiAffinityData, bool) {
	if o == nil || o.AntiAffinityData == nil {
		return nil, false
	}
	return o.AntiAffinityData, true
}

// HasAntiAffinityData returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasAntiAffinityData() bool {
	if o != nil && o.AntiAffinityData != nil {
		return true
	}

	return false
}

// SetAntiAffinityData gets a reference to the given ApplicationAntiAffinityData and assigns it to the AntiAffinityData field.
func (o *ApplicationPoolInfoV3) SetAntiAffinityData(v ApplicationAntiAffinityData) {
	o.AntiAffinityData = &v
}

// GetCategoryFolderName returns the CategoryFolderName field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetCategoryFolderName() string {
	if o == nil || o.CategoryFolderName == nil {
		var ret string
		return ret
	}
	return *o.CategoryFolderName
}

// GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetCategoryFolderNameOk() (*string, bool) {
	if o == nil || o.CategoryFolderName == nil {
		return nil, false
	}
	return o.CategoryFolderName, true
}

// HasCategoryFolderName returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasCategoryFolderName() bool {
	if o != nil && o.CategoryFolderName != nil {
		return true
	}

	return false
}

// SetCategoryFolderName gets a reference to the given string and assigns it to the CategoryFolderName field.
func (o *ApplicationPoolInfoV3) SetCategoryFolderName(v string) {
	o.CategoryFolderName = &v
}

// GetCloudBrokered returns the CloudBrokered field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetCloudBrokered() bool {
	if o == nil || o.CloudBrokered == nil {
		var ret bool
		return ret
	}
	return *o.CloudBrokered
}

// GetCloudBrokeredOk returns a tuple with the CloudBrokered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetCloudBrokeredOk() (*bool, bool) {
	if o == nil || o.CloudBrokered == nil {
		return nil, false
	}
	return o.CloudBrokered, true
}

// HasCloudBrokered returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasCloudBrokered() bool {
	if o != nil && o.CloudBrokered != nil {
		return true
	}

	return false
}

// SetCloudBrokered gets a reference to the given bool and assigns it to the CloudBrokered field.
func (o *ApplicationPoolInfoV3) SetCloudBrokered(v bool) {
	o.CloudBrokered = &v
}

// GetCsRestrictionTags returns the CsRestrictionTags field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetCsRestrictionTags() []string {
	if o == nil || o.CsRestrictionTags == nil {
		var ret []string
		return ret
	}
	return *o.CsRestrictionTags
}

// GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetCsRestrictionTagsOk() (*[]string, bool) {
	if o == nil || o.CsRestrictionTags == nil {
		return nil, false
	}
	return o.CsRestrictionTags, true
}

// HasCsRestrictionTags returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasCsRestrictionTags() bool {
	if o != nil && o.CsRestrictionTags != nil {
		return true
	}

	return false
}

// SetCsRestrictionTags gets a reference to the given []string and assigns it to the CsRestrictionTags field.
func (o *ApplicationPoolInfoV3) SetCsRestrictionTags(v []string) {
	o.CsRestrictionTags = &v
}

// GetCustomizedIconIds returns the CustomizedIconIds field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetCustomizedIconIds() []string {
	if o == nil || o.CustomizedIconIds == nil {
		var ret []string
		return ret
	}
	return *o.CustomizedIconIds
}

// GetCustomizedIconIdsOk returns a tuple with the CustomizedIconIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetCustomizedIconIdsOk() (*[]string, bool) {
	if o == nil || o.CustomizedIconIds == nil {
		return nil, false
	}
	return o.CustomizedIconIds, true
}

// HasCustomizedIconIds returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasCustomizedIconIds() bool {
	if o != nil && o.CustomizedIconIds != nil {
		return true
	}

	return false
}

// SetCustomizedIconIds gets a reference to the given []string and assigns it to the CustomizedIconIds field.
func (o *ApplicationPoolInfoV3) SetCustomizedIconIds(v []string) {
	o.CustomizedIconIds = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ApplicationPoolInfoV3) SetDescription(v string) {
	o.Description = &v
}

// GetDesktopPoolId returns the DesktopPoolId field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetDesktopPoolId() string {
	if o == nil || o.DesktopPoolId == nil {
		var ret string
		return ret
	}
	return *o.DesktopPoolId
}

// GetDesktopPoolIdOk returns a tuple with the DesktopPoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetDesktopPoolIdOk() (*string, bool) {
	if o == nil || o.DesktopPoolId == nil {
		return nil, false
	}
	return o.DesktopPoolId, true
}

// HasDesktopPoolId returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasDesktopPoolId() bool {
	if o != nil && o.DesktopPoolId != nil {
		return true
	}

	return false
}

// SetDesktopPoolId gets a reference to the given string and assigns it to the DesktopPoolId field.
func (o *ApplicationPoolInfoV3) SetDesktopPoolId(v string) {
	o.DesktopPoolId = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ApplicationPoolInfoV3) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnableClientRestrictions returns the EnableClientRestrictions field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetEnableClientRestrictions() bool {
	if o == nil || o.EnableClientRestrictions == nil {
		var ret bool
		return ret
	}
	return *o.EnableClientRestrictions
}

// GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetEnableClientRestrictionsOk() (*bool, bool) {
	if o == nil || o.EnableClientRestrictions == nil {
		return nil, false
	}
	return o.EnableClientRestrictions, true
}

// HasEnableClientRestrictions returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasEnableClientRestrictions() bool {
	if o != nil && o.EnableClientRestrictions != nil {
		return true
	}

	return false
}

// SetEnableClientRestrictions gets a reference to the given bool and assigns it to the EnableClientRestrictions field.
func (o *ApplicationPoolInfoV3) SetEnableClientRestrictions(v bool) {
	o.EnableClientRestrictions = &v
}

// GetEnablePreLaunch returns the EnablePreLaunch field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetEnablePreLaunch() bool {
	if o == nil || o.EnablePreLaunch == nil {
		var ret bool
		return ret
	}
	return *o.EnablePreLaunch
}

// GetEnablePreLaunchOk returns a tuple with the EnablePreLaunch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetEnablePreLaunchOk() (*bool, bool) {
	if o == nil || o.EnablePreLaunch == nil {
		return nil, false
	}
	return o.EnablePreLaunch, true
}

// HasEnablePreLaunch returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasEnablePreLaunch() bool {
	if o != nil && o.EnablePreLaunch != nil {
		return true
	}

	return false
}

// SetEnablePreLaunch gets a reference to the given bool and assigns it to the EnablePreLaunch field.
func (o *ApplicationPoolInfoV3) SetEnablePreLaunch(v bool) {
	o.EnablePreLaunch = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ApplicationPoolInfoV3) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExecutablePath returns the ExecutablePath field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetExecutablePath() string {
	if o == nil || o.ExecutablePath == nil {
		var ret string
		return ret
	}
	return *o.ExecutablePath
}

// GetExecutablePathOk returns a tuple with the ExecutablePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetExecutablePathOk() (*string, bool) {
	if o == nil || o.ExecutablePath == nil {
		return nil, false
	}
	return o.ExecutablePath, true
}

// HasExecutablePath returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasExecutablePath() bool {
	if o != nil && o.ExecutablePath != nil {
		return true
	}

	return false
}

// SetExecutablePath gets a reference to the given string and assigns it to the ExecutablePath field.
func (o *ApplicationPoolInfoV3) SetExecutablePath(v string) {
	o.ExecutablePath = &v
}

// GetFarmId returns the FarmId field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetFarmId() string {
	if o == nil || o.FarmId == nil {
		var ret string
		return ret
	}
	return *o.FarmId
}

// GetFarmIdOk returns a tuple with the FarmId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetFarmIdOk() (*string, bool) {
	if o == nil || o.FarmId == nil {
		return nil, false
	}
	return o.FarmId, true
}

// HasFarmId returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasFarmId() bool {
	if o != nil && o.FarmId != nil {
		return true
	}

	return false
}

// SetFarmId gets a reference to the given string and assigns it to the FarmId field.
func (o *ApplicationPoolInfoV3) SetFarmId(v string) {
	o.FarmId = &v
}

// GetGlobalApplicationEntitlementId returns the GlobalApplicationEntitlementId field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetGlobalApplicationEntitlementId() string {
	if o == nil || o.GlobalApplicationEntitlementId == nil {
		var ret string
		return ret
	}
	return *o.GlobalApplicationEntitlementId
}

// GetGlobalApplicationEntitlementIdOk returns a tuple with the GlobalApplicationEntitlementId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetGlobalApplicationEntitlementIdOk() (*string, bool) {
	if o == nil || o.GlobalApplicationEntitlementId == nil {
		return nil, false
	}
	return o.GlobalApplicationEntitlementId, true
}

// HasGlobalApplicationEntitlementId returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasGlobalApplicationEntitlementId() bool {
	if o != nil && o.GlobalApplicationEntitlementId != nil {
		return true
	}

	return false
}

// SetGlobalApplicationEntitlementId gets a reference to the given string and assigns it to the GlobalApplicationEntitlementId field.
func (o *ApplicationPoolInfoV3) SetGlobalApplicationEntitlementId(v string) {
	o.GlobalApplicationEntitlementId = &v
}

// GetIconIds returns the IconIds field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetIconIds() []string {
	if o == nil || o.IconIds == nil {
		var ret []string
		return ret
	}
	return *o.IconIds
}

// GetIconIdsOk returns a tuple with the IconIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetIconIdsOk() (*[]string, bool) {
	if o == nil || o.IconIds == nil {
		return nil, false
	}
	return o.IconIds, true
}

// HasIconIds returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasIconIds() bool {
	if o != nil && o.IconIds != nil {
		return true
	}

	return false
}

// SetIconIds gets a reference to the given []string and assigns it to the IconIds field.
func (o *ApplicationPoolInfoV3) SetIconIds(v []string) {
	o.IconIds = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApplicationPoolInfoV3) SetId(v string) {
	o.Id = &v
}

// GetMaxMultiSessions returns the MaxMultiSessions field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetMaxMultiSessions() int32 {
	if o == nil || o.MaxMultiSessions == nil {
		var ret int32
		return ret
	}
	return *o.MaxMultiSessions
}

// GetMaxMultiSessionsOk returns a tuple with the MaxMultiSessions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetMaxMultiSessionsOk() (*int32, bool) {
	if o == nil || o.MaxMultiSessions == nil {
		return nil, false
	}
	return o.MaxMultiSessions, true
}

// HasMaxMultiSessions returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasMaxMultiSessions() bool {
	if o != nil && o.MaxMultiSessions != nil {
		return true
	}

	return false
}

// SetMaxMultiSessions gets a reference to the given int32 and assigns it to the MaxMultiSessions field.
func (o *ApplicationPoolInfoV3) SetMaxMultiSessions(v int32) {
	o.MaxMultiSessions = &v
}

// GetMultiSessionMode returns the MultiSessionMode field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetMultiSessionMode() string {
	if o == nil || o.MultiSessionMode == nil {
		var ret string
		return ret
	}
	return *o.MultiSessionMode
}

// GetMultiSessionModeOk returns a tuple with the MultiSessionMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetMultiSessionModeOk() (*string, bool) {
	if o == nil || o.MultiSessionMode == nil {
		return nil, false
	}
	return o.MultiSessionMode, true
}

// HasMultiSessionMode returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasMultiSessionMode() bool {
	if o != nil && o.MultiSessionMode != nil {
		return true
	}

	return false
}

// SetMultiSessionMode gets a reference to the given string and assigns it to the MultiSessionMode field.
func (o *ApplicationPoolInfoV3) SetMultiSessionMode(v string) {
	o.MultiSessionMode = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApplicationPoolInfoV3) SetName(v string) {
	o.Name = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetParameters() string {
	if o == nil || o.Parameters == nil {
		var ret string
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetParametersOk() (*string, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given string and assigns it to the Parameters field.
func (o *ApplicationPoolInfoV3) SetParameters(v string) {
	o.Parameters = &v
}

// GetPublisher returns the Publisher field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetPublisher() string {
	if o == nil || o.Publisher == nil {
		var ret string
		return ret
	}
	return *o.Publisher
}

// GetPublisherOk returns a tuple with the Publisher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetPublisherOk() (*string, bool) {
	if o == nil || o.Publisher == nil {
		return nil, false
	}
	return o.Publisher, true
}

// HasPublisher returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasPublisher() bool {
	if o != nil && o.Publisher != nil {
		return true
	}

	return false
}

// SetPublisher gets a reference to the given string and assigns it to the Publisher field.
func (o *ApplicationPoolInfoV3) SetPublisher(v string) {
	o.Publisher = &v
}

// GetShortcutLocations returns the ShortcutLocations field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetShortcutLocations() []string {
	if o == nil || o.ShortcutLocations == nil {
		var ret []string
		return ret
	}
	return *o.ShortcutLocations
}

// GetShortcutLocationsOk returns a tuple with the ShortcutLocations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetShortcutLocationsOk() (*[]string, bool) {
	if o == nil || o.ShortcutLocations == nil {
		return nil, false
	}
	return o.ShortcutLocations, true
}

// HasShortcutLocations returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasShortcutLocations() bool {
	if o != nil && o.ShortcutLocations != nil {
		return true
	}

	return false
}

// SetShortcutLocations gets a reference to the given []string and assigns it to the ShortcutLocations field.
func (o *ApplicationPoolInfoV3) SetShortcutLocations(v []string) {
	o.ShortcutLocations = &v
}

// GetStartFolder returns the StartFolder field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetStartFolder() string {
	if o == nil || o.StartFolder == nil {
		var ret string
		return ret
	}
	return *o.StartFolder
}

// GetStartFolderOk returns a tuple with the StartFolder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetStartFolderOk() (*string, bool) {
	if o == nil || o.StartFolder == nil {
		return nil, false
	}
	return o.StartFolder, true
}

// HasStartFolder returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasStartFolder() bool {
	if o != nil && o.StartFolder != nil {
		return true
	}

	return false
}

// SetStartFolder gets a reference to the given string and assigns it to the StartFolder field.
func (o *ApplicationPoolInfoV3) SetStartFolder(v string) {
	o.StartFolder = &v
}

// GetSupportedFileTypesData returns the SupportedFileTypesData field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetSupportedFileTypesData() ApplicationSupportedFileTypesData {
	if o == nil || o.SupportedFileTypesData == nil {
		var ret ApplicationSupportedFileTypesData
		return ret
	}
	return *o.SupportedFileTypesData
}

// GetSupportedFileTypesDataOk returns a tuple with the SupportedFileTypesData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetSupportedFileTypesDataOk() (*ApplicationSupportedFileTypesData, bool) {
	if o == nil || o.SupportedFileTypesData == nil {
		return nil, false
	}
	return o.SupportedFileTypesData, true
}

// HasSupportedFileTypesData returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasSupportedFileTypesData() bool {
	if o != nil && o.SupportedFileTypesData != nil {
		return true
	}

	return false
}

// SetSupportedFileTypesData gets a reference to the given ApplicationSupportedFileTypesData and assigns it to the SupportedFileTypesData field.
func (o *ApplicationPoolInfoV3) SetSupportedFileTypesData(v ApplicationSupportedFileTypesData) {
	o.SupportedFileTypesData = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ApplicationPoolInfoV3) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationPoolInfoV3) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ApplicationPoolInfoV3) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ApplicationPoolInfoV3) SetVersion(v string) {
	o.Version = &v
}

func (o ApplicationPoolInfoV3) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessGroupId != nil {
		toSerialize["access_group_id"] = o.AccessGroupId
	}
	if o.AntiAffinityData != nil {
		toSerialize["anti_affinity_data"] = o.AntiAffinityData
	}
	if o.CategoryFolderName != nil {
		toSerialize["category_folder_name"] = o.CategoryFolderName
	}
	if o.CloudBrokered != nil {
		toSerialize["cloud_brokered"] = o.CloudBrokered
	}
	if o.CsRestrictionTags != nil {
		toSerialize["cs_restriction_tags"] = o.CsRestrictionTags
	}
	if o.CustomizedIconIds != nil {
		toSerialize["customized_icon_ids"] = o.CustomizedIconIds
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DesktopPoolId != nil {
		toSerialize["desktop_pool_id"] = o.DesktopPoolId
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
	if o.ExecutablePath != nil {
		toSerialize["executable_path"] = o.ExecutablePath
	}
	if o.FarmId != nil {
		toSerialize["farm_id"] = o.FarmId
	}
	if o.GlobalApplicationEntitlementId != nil {
		toSerialize["global_application_entitlement_id"] = o.GlobalApplicationEntitlementId
	}
	if o.IconIds != nil {
		toSerialize["icon_ids"] = o.IconIds
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.MaxMultiSessions != nil {
		toSerialize["max_multi_sessions"] = o.MaxMultiSessions
	}
	if o.MultiSessionMode != nil {
		toSerialize["multi_session_mode"] = o.MultiSessionMode
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	if o.Publisher != nil {
		toSerialize["publisher"] = o.Publisher
	}
	if o.ShortcutLocations != nil {
		toSerialize["shortcut_locations"] = o.ShortcutLocations
	}
	if o.StartFolder != nil {
		toSerialize["start_folder"] = o.StartFolder
	}
	if o.SupportedFileTypesData != nil {
		toSerialize["supported_file_types_data"] = o.SupportedFileTypesData
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationPoolInfoV3 struct {
	value *ApplicationPoolInfoV3
	isSet bool
}

func (v NullableApplicationPoolInfoV3) Get() *ApplicationPoolInfoV3 {
	return v.value
}

func (v *NullableApplicationPoolInfoV3) Set(val *ApplicationPoolInfoV3) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationPoolInfoV3) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationPoolInfoV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationPoolInfoV3(val *ApplicationPoolInfoV3) *NullableApplicationPoolInfoV3 {
	return &NullableApplicationPoolInfoV3{value: val, isSet: true}
}

func (v NullableApplicationPoolInfoV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationPoolInfoV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
