# GlobalApplicationEntitlementCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | Pointer to **bool** | Indicates whether the users can choose the protocol used. Default value is true. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the entitlement. The property will not be set if the entitlement does not belong to a category. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | List of connection server restriction tags to which the access to the global application entitlement is restricted. Empty or null list means that entitlement can be accessed from any connection server. | [optional] 
**DefaultDisplayProtocol** | Pointer to **string** | The default display protocol for the global application entitlement. This can only be set to \&quot;PCOIP\&quot; or \&quot;BLAST\&quot;. If this application&#39;s Farm&#39;s or desktop pool&#39;s allow_users_to_choose_protocol is set to false, then default_display_protocol must match that default_display_protocol of farm or desktop pool. Default value is \&quot;PCOIP\&quot;. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | [optional] 
**Description** | Pointer to **string** | Description of global application entitlement. | [optional] 
**DisplayName** | Pointer to **string** | The display name is the name that users will see when they connect using Horizon View Client. If display_name is left blank, it defaults to name. | [optional] 
**EnableClientRestrictions** | Pointer to **bool** | Indicates whether client restrictions to be applied to global application entitlement, currently it is valid for published application pools. Default value is false. | [optional] 
**EnablePreLaunch** | Pointer to **bool** | If value is set to true, this global application entitlement can be pre-launched. This property can be set to true only if multi_session_mode is set to \&quot;DISABLED\&quot;.  Default value is false. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the global application entitlement is enabled. Default value is true. | [optional] 
**FederatedAccessGroupId** | **string** | ID of the federated access group with which the global application entitlement is to be associated. They can also be used for delegated administration. | 
**MultiSessionMode** | Pointer to **string** | Multi-session mode for this entitlement. A global application entitlement launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. Default value is \&quot;DISABLED\&quot; * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application. | [optional] 
**MultipleSessionAutoClean** | Pointer to **bool** | Indicates whether automatic session clean up is enabled. Default value is false. | [optional] 
**Name** | **string** | Unique name used to identify the global application entitlement. | 
**RequireHomeSite** | Pointer to **bool** | Indicates whether it should fail if a home site isn&#39;t defined for this global application entitlement. This property cannot be set to true if use_home_site is set to false. Default value is false. | [optional] 
**Scope** | Pointer to **string** | Scope for this global application entitlement. Visibility and Placement policies are defined by this value. Default value of \&quot;ALL_SITES\&quot;. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set. | [optional] 
**ShortcutLocationsV2** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop. This property is required if category_folder_name is set. | [optional] 
**UseHomeSite** | Pointer to **bool** | Indicates whether a pod in the user&#39;s home site is used to start the search or the current site is used. Default value is false. | [optional] 

## Methods

### NewGlobalApplicationEntitlementCreateSpec

`func NewGlobalApplicationEntitlementCreateSpec(federatedAccessGroupId string, name string, ) *GlobalApplicationEntitlementCreateSpec`

NewGlobalApplicationEntitlementCreateSpec instantiates a new GlobalApplicationEntitlementCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalApplicationEntitlementCreateSpecWithDefaults

`func NewGlobalApplicationEntitlementCreateSpecWithDefaults() *GlobalApplicationEntitlementCreateSpec`

NewGlobalApplicationEntitlementCreateSpecWithDefaults instantiates a new GlobalApplicationEntitlementCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsersToChooseProtocol

`func (o *GlobalApplicationEntitlementCreateSpec) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *GlobalApplicationEntitlementCreateSpec) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.

### HasAllowUsersToChooseProtocol

`func (o *GlobalApplicationEntitlementCreateSpec) HasAllowUsersToChooseProtocol() bool`

HasAllowUsersToChooseProtocol returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *GlobalApplicationEntitlementCreateSpec) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *GlobalApplicationEntitlementCreateSpec) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *GlobalApplicationEntitlementCreateSpec) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *GlobalApplicationEntitlementCreateSpec) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *GlobalApplicationEntitlementCreateSpec) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *GlobalApplicationEntitlementCreateSpec) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *GlobalApplicationEntitlementCreateSpec) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *GlobalApplicationEntitlementCreateSpec) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.

### HasDefaultDisplayProtocol

`func (o *GlobalApplicationEntitlementCreateSpec) HasDefaultDisplayProtocol() bool`

HasDefaultDisplayProtocol returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalApplicationEntitlementCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalApplicationEntitlementCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalApplicationEntitlementCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *GlobalApplicationEntitlementCreateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GlobalApplicationEntitlementCreateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GlobalApplicationEntitlementCreateSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *GlobalApplicationEntitlementCreateSpec) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *GlobalApplicationEntitlementCreateSpec) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *GlobalApplicationEntitlementCreateSpec) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnablePreLaunch

`func (o *GlobalApplicationEntitlementCreateSpec) GetEnablePreLaunch() bool`

GetEnablePreLaunch returns the EnablePreLaunch field if non-nil, zero value otherwise.

### GetEnablePreLaunchOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetEnablePreLaunchOk() (*bool, bool)`

GetEnablePreLaunchOk returns a tuple with the EnablePreLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePreLaunch

`func (o *GlobalApplicationEntitlementCreateSpec) SetEnablePreLaunch(v bool)`

SetEnablePreLaunch sets EnablePreLaunch field to given value.

### HasEnablePreLaunch

`func (o *GlobalApplicationEntitlementCreateSpec) HasEnablePreLaunch() bool`

HasEnablePreLaunch returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalApplicationEntitlementCreateSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalApplicationEntitlementCreateSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GlobalApplicationEntitlementCreateSpec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFederatedAccessGroupId

`func (o *GlobalApplicationEntitlementCreateSpec) GetFederatedAccessGroupId() string`

GetFederatedAccessGroupId returns the FederatedAccessGroupId field if non-nil, zero value otherwise.

### GetFederatedAccessGroupIdOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetFederatedAccessGroupIdOk() (*string, bool)`

GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAccessGroupId

`func (o *GlobalApplicationEntitlementCreateSpec) SetFederatedAccessGroupId(v string)`

SetFederatedAccessGroupId sets FederatedAccessGroupId field to given value.


### GetMultiSessionMode

`func (o *GlobalApplicationEntitlementCreateSpec) GetMultiSessionMode() string`

GetMultiSessionMode returns the MultiSessionMode field if non-nil, zero value otherwise.

### GetMultiSessionModeOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetMultiSessionModeOk() (*string, bool)`

GetMultiSessionModeOk returns a tuple with the MultiSessionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSessionMode

`func (o *GlobalApplicationEntitlementCreateSpec) SetMultiSessionMode(v string)`

SetMultiSessionMode sets MultiSessionMode field to given value.

### HasMultiSessionMode

`func (o *GlobalApplicationEntitlementCreateSpec) HasMultiSessionMode() bool`

HasMultiSessionMode returns a boolean if a field has been set.

### GetMultipleSessionAutoClean

`func (o *GlobalApplicationEntitlementCreateSpec) GetMultipleSessionAutoClean() bool`

GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field if non-nil, zero value otherwise.

### GetMultipleSessionAutoCleanOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetMultipleSessionAutoCleanOk() (*bool, bool)`

GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSessionAutoClean

`func (o *GlobalApplicationEntitlementCreateSpec) SetMultipleSessionAutoClean(v bool)`

SetMultipleSessionAutoClean sets MultipleSessionAutoClean field to given value.

### HasMultipleSessionAutoClean

`func (o *GlobalApplicationEntitlementCreateSpec) HasMultipleSessionAutoClean() bool`

HasMultipleSessionAutoClean returns a boolean if a field has been set.

### GetName

`func (o *GlobalApplicationEntitlementCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalApplicationEntitlementCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetRequireHomeSite

`func (o *GlobalApplicationEntitlementCreateSpec) GetRequireHomeSite() bool`

GetRequireHomeSite returns the RequireHomeSite field if non-nil, zero value otherwise.

### GetRequireHomeSiteOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetRequireHomeSiteOk() (*bool, bool)`

GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHomeSite

`func (o *GlobalApplicationEntitlementCreateSpec) SetRequireHomeSite(v bool)`

SetRequireHomeSite sets RequireHomeSite field to given value.

### HasRequireHomeSite

`func (o *GlobalApplicationEntitlementCreateSpec) HasRequireHomeSite() bool`

HasRequireHomeSite returns a boolean if a field has been set.

### GetScope

`func (o *GlobalApplicationEntitlementCreateSpec) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalApplicationEntitlementCreateSpec) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GlobalApplicationEntitlementCreateSpec) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetShortcutLocationsV2

`func (o *GlobalApplicationEntitlementCreateSpec) GetShortcutLocationsV2() []string`

GetShortcutLocationsV2 returns the ShortcutLocationsV2 field if non-nil, zero value otherwise.

### GetShortcutLocationsV2Ok

`func (o *GlobalApplicationEntitlementCreateSpec) GetShortcutLocationsV2Ok() (*[]string, bool)`

GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocationsV2

`func (o *GlobalApplicationEntitlementCreateSpec) SetShortcutLocationsV2(v []string)`

SetShortcutLocationsV2 sets ShortcutLocationsV2 field to given value.

### HasShortcutLocationsV2

`func (o *GlobalApplicationEntitlementCreateSpec) HasShortcutLocationsV2() bool`

HasShortcutLocationsV2 returns a boolean if a field has been set.

### GetUseHomeSite

`func (o *GlobalApplicationEntitlementCreateSpec) GetUseHomeSite() bool`

GetUseHomeSite returns the UseHomeSite field if non-nil, zero value otherwise.

### GetUseHomeSiteOk

`func (o *GlobalApplicationEntitlementCreateSpec) GetUseHomeSiteOk() (*bool, bool)`

GetUseHomeSiteOk returns a tuple with the UseHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHomeSite

`func (o *GlobalApplicationEntitlementCreateSpec) SetUseHomeSite(v bool)`

SetUseHomeSite sets UseHomeSite field to given value.

### HasUseHomeSite

`func (o *GlobalApplicationEntitlementCreateSpec) HasUseHomeSite() bool`

HasUseHomeSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


