# GlobalApplicationEntitlementUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupGaeId** | Pointer to **string** | Global Application Entitlement that can be used as backup for this Global Application Entitlement. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the entitlement. The property will not be set if the entitlement does not belong to a category. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | List of connection server restriction tags to which the access to the global application entitlement is restricted. Empty or null list means that entitlement can be accessed from any connection server. | [optional] 
**DefaultDisplayProtocol** | **string** | The default display protocol for the global application entitlement. This can only be set to \&quot;PCOIP\&quot; or \&quot;BLAST\&quot;. If this application&#39;s Farm&#39;s or desktop pool&#39;s allow_users_to_choose_protocol is set to false, then default_display_protocol must match that default_display_protocol of farm or desktop pool. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | 
**Description** | Pointer to **string** | Description of global application entitlement. | [optional] 
**DisplayName** | **string** | The display name is the name that users will see when they connect using Horizon View Client. If display_name is left blank, it defaults to name. | 
**EnableClientRestrictions** | **bool** | Indicates whether client restrictions to be applied to global application entitlement, currently it is valid for published application pools. | 
**EnablePreLaunch** | **bool** | If value is set to true, this global application entitlement can be pre-launched. This property can be set to true only if multi_session_mode is set to \&quot;DISABLED\&quot;. | 
**Enabled** | **bool** | Indicates whether the global application entitlement is enabled. | 
**FederatedAccessGroupId** | **string** | ID of the federated access group with which the global application entitlement is to be associated. They can also be used for delegated administration. | 
**MultiSessionMode** | **string** | Multi-session mode for this entitlement. A global application entitlement launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application. | 
**MultipleSessionAutoClean** | **bool** | Indicates whether automatic session clean up is enabled. | 
**Name** | **string** | Unique name used to identify the global application entitlement. | 
**RequireHomeSite** | **bool** | Indicates whether it should fail if a home site isn&#39;t defined for this global application entitlement. This property cannot be set to true if use_home_site is set to false. | 
**Scope** | **string** | Scope for this global application entitlement. Visibility and Placement policies are defined by this value. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set. | 
**ShortcutLocationsV2** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop. This property is required if category_folder_name is set. | [optional] 
**UseHomeSite** | **bool** | Indicates whether a pod in the user&#39;s home site is used to start the search or the current site is used. | 

## Methods

### NewGlobalApplicationEntitlementUpdateSpec

`func NewGlobalApplicationEntitlementUpdateSpec(defaultDisplayProtocol string, displayName string, enableClientRestrictions bool, enablePreLaunch bool, enabled bool, federatedAccessGroupId string, multiSessionMode string, multipleSessionAutoClean bool, name string, requireHomeSite bool, scope string, useHomeSite bool, ) *GlobalApplicationEntitlementUpdateSpec`

NewGlobalApplicationEntitlementUpdateSpec instantiates a new GlobalApplicationEntitlementUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalApplicationEntitlementUpdateSpecWithDefaults

`func NewGlobalApplicationEntitlementUpdateSpecWithDefaults() *GlobalApplicationEntitlementUpdateSpec`

NewGlobalApplicationEntitlementUpdateSpecWithDefaults instantiates a new GlobalApplicationEntitlementUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupGaeId

`func (o *GlobalApplicationEntitlementUpdateSpec) GetBackupGaeId() string`

GetBackupGaeId returns the BackupGaeId field if non-nil, zero value otherwise.

### GetBackupGaeIdOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetBackupGaeIdOk() (*string, bool)`

GetBackupGaeIdOk returns a tuple with the BackupGaeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGaeId

`func (o *GlobalApplicationEntitlementUpdateSpec) SetBackupGaeId(v string)`

SetBackupGaeId sets BackupGaeId field to given value.

### HasBackupGaeId

`func (o *GlobalApplicationEntitlementUpdateSpec) HasBackupGaeId() bool`

HasBackupGaeId returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *GlobalApplicationEntitlementUpdateSpec) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *GlobalApplicationEntitlementUpdateSpec) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *GlobalApplicationEntitlementUpdateSpec) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *GlobalApplicationEntitlementUpdateSpec) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *GlobalApplicationEntitlementUpdateSpec) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *GlobalApplicationEntitlementUpdateSpec) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *GlobalApplicationEntitlementUpdateSpec) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *GlobalApplicationEntitlementUpdateSpec) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.


### GetDescription

`func (o *GlobalApplicationEntitlementUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalApplicationEntitlementUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalApplicationEntitlementUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *GlobalApplicationEntitlementUpdateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GlobalApplicationEntitlementUpdateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnableClientRestrictions

`func (o *GlobalApplicationEntitlementUpdateSpec) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *GlobalApplicationEntitlementUpdateSpec) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.


### GetEnablePreLaunch

`func (o *GlobalApplicationEntitlementUpdateSpec) GetEnablePreLaunch() bool`

GetEnablePreLaunch returns the EnablePreLaunch field if non-nil, zero value otherwise.

### GetEnablePreLaunchOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetEnablePreLaunchOk() (*bool, bool)`

GetEnablePreLaunchOk returns a tuple with the EnablePreLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePreLaunch

`func (o *GlobalApplicationEntitlementUpdateSpec) SetEnablePreLaunch(v bool)`

SetEnablePreLaunch sets EnablePreLaunch field to given value.


### GetEnabled

`func (o *GlobalApplicationEntitlementUpdateSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalApplicationEntitlementUpdateSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFederatedAccessGroupId

`func (o *GlobalApplicationEntitlementUpdateSpec) GetFederatedAccessGroupId() string`

GetFederatedAccessGroupId returns the FederatedAccessGroupId field if non-nil, zero value otherwise.

### GetFederatedAccessGroupIdOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetFederatedAccessGroupIdOk() (*string, bool)`

GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAccessGroupId

`func (o *GlobalApplicationEntitlementUpdateSpec) SetFederatedAccessGroupId(v string)`

SetFederatedAccessGroupId sets FederatedAccessGroupId field to given value.


### GetMultiSessionMode

`func (o *GlobalApplicationEntitlementUpdateSpec) GetMultiSessionMode() string`

GetMultiSessionMode returns the MultiSessionMode field if non-nil, zero value otherwise.

### GetMultiSessionModeOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetMultiSessionModeOk() (*string, bool)`

GetMultiSessionModeOk returns a tuple with the MultiSessionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSessionMode

`func (o *GlobalApplicationEntitlementUpdateSpec) SetMultiSessionMode(v string)`

SetMultiSessionMode sets MultiSessionMode field to given value.


### GetMultipleSessionAutoClean

`func (o *GlobalApplicationEntitlementUpdateSpec) GetMultipleSessionAutoClean() bool`

GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field if non-nil, zero value otherwise.

### GetMultipleSessionAutoCleanOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetMultipleSessionAutoCleanOk() (*bool, bool)`

GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSessionAutoClean

`func (o *GlobalApplicationEntitlementUpdateSpec) SetMultipleSessionAutoClean(v bool)`

SetMultipleSessionAutoClean sets MultipleSessionAutoClean field to given value.


### GetName

`func (o *GlobalApplicationEntitlementUpdateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalApplicationEntitlementUpdateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetRequireHomeSite

`func (o *GlobalApplicationEntitlementUpdateSpec) GetRequireHomeSite() bool`

GetRequireHomeSite returns the RequireHomeSite field if non-nil, zero value otherwise.

### GetRequireHomeSiteOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetRequireHomeSiteOk() (*bool, bool)`

GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHomeSite

`func (o *GlobalApplicationEntitlementUpdateSpec) SetRequireHomeSite(v bool)`

SetRequireHomeSite sets RequireHomeSite field to given value.


### GetScope

`func (o *GlobalApplicationEntitlementUpdateSpec) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalApplicationEntitlementUpdateSpec) SetScope(v string)`

SetScope sets Scope field to given value.


### GetShortcutLocationsV2

`func (o *GlobalApplicationEntitlementUpdateSpec) GetShortcutLocationsV2() []string`

GetShortcutLocationsV2 returns the ShortcutLocationsV2 field if non-nil, zero value otherwise.

### GetShortcutLocationsV2Ok

`func (o *GlobalApplicationEntitlementUpdateSpec) GetShortcutLocationsV2Ok() (*[]string, bool)`

GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocationsV2

`func (o *GlobalApplicationEntitlementUpdateSpec) SetShortcutLocationsV2(v []string)`

SetShortcutLocationsV2 sets ShortcutLocationsV2 field to given value.

### HasShortcutLocationsV2

`func (o *GlobalApplicationEntitlementUpdateSpec) HasShortcutLocationsV2() bool`

HasShortcutLocationsV2 returns a boolean if a field has been set.

### GetUseHomeSite

`func (o *GlobalApplicationEntitlementUpdateSpec) GetUseHomeSite() bool`

GetUseHomeSite returns the UseHomeSite field if non-nil, zero value otherwise.

### GetUseHomeSiteOk

`func (o *GlobalApplicationEntitlementUpdateSpec) GetUseHomeSiteOk() (*bool, bool)`

GetUseHomeSiteOk returns a tuple with the UseHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHomeSite

`func (o *GlobalApplicationEntitlementUpdateSpec) SetUseHomeSite(v bool)`

SetUseHomeSite sets UseHomeSite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


