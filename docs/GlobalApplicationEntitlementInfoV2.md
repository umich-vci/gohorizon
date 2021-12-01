# GlobalApplicationEntitlementInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | Pointer to **bool** | Indicates whether the users can choose the protocol used. If set to true, the application pools that are associated with this Global Application Entitlement must also allow users to choose display protocol with allowUsersToChooseProtocol. Supported Filters: &#39;Equals&#39;. | [optional] 
**ApplicationData** | Pointer to [**ApplicationData**](ApplicationData.md) |  | [optional] 
**ApplicationIconIds** | Pointer to **[]string** | Icons associated with the Global Application Entitlement | [optional] 
**BackupGaeId** | Pointer to **string** | Indicates the Global Application Entitlement that can be used as backup for this Global Application Entitlement. Supported Filters: &#39;Equals&#39;. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the entitlement. Unset if the entitlement does not belong to a category. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | Connection server restrictions. This is a list of tags that access to the entitlement is restricted to. No list means that the entitlement can be accessed from any connection server. | [optional] 
**DefaultDisplayProtocol** | Pointer to **string** | The default display protocol for the Global Application Entitlement. Must be a protocol in the supportedDisplayProtocols list. Clients connecting through this Global Application Entitlement that do not specify a protocol will use this value, not the value specified directly on the application pool to which they connect (if different). Supported Filters: &#39;Equals&#39;. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | [optional] 
**Description** | Pointer to **string** | Description of Global Application Entitlement. This property has a maximum length of 1024 characters. Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DisplayName** | Pointer to **string** | Name that users will see when they connect using Horizon Client. If the display name is left blank, while creating or updating the Global Application Entitlement display name value will be same as name. This property has a maximum length of 64 characters. Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**EnableClientRestrictions** | Pointer to **bool** | Indicates whether client restrictions to be applied to Global Application Entitlement. Currently it is valid for RDSH pools. Supported Filters: &#39;Equals&#39;. | [optional] 
**EnablePreLaunch** | Pointer to **bool** | Indicates whether Global Application Entitlement can be pre-launched Supported Filters: &#39;Equals&#39;. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if this Global Application Entitlement is enabled. Supported Filters: &#39;Equals&#39;. | [optional] 
**FederatedAccessGroupId** | Pointer to **string** | This represents id of the federated access group associated with the global application entitlement. | [optional] 
**GroupCount** | Pointer to **int32** | Count of groups that are associated with this Global Application Entitlement. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this Global Application Entitlement. Supported Filters: &#39;Equals&#39;. | [optional] 
**LocalApplicationPoolsCount** | Pointer to **int32** | Count of application pools local to this pod that are associated with this Global Application Entitlement. | [optional] 
**MemberPods** | Pointer to **[]string** | Pods that have application pools associated with this Global Application Entitlement. | [optional] 
**MultiSessionMode** | Pointer to **string** | Multi-session mode for the application pool. An application launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. Multi-session mode should be disabled when pre-launch is enabled. Supported Filters: &#39;Equals&#39;. * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application. | [optional] 
**MultipleSessionAutoClean** | Pointer to **bool** | Indicates if automatic session clean up is enabled. This cannot be enabled when this Global Application Entitlement is associated with a Application pool that has dedicated user assignment. Supported Filters: &#39;Equals&#39;. | [optional] 
**Name** | Pointer to **string** | Unique name used to identify the Global Application Entitlement. This property has a maximum length of 64 characters. Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**PrimaryGaeId** | Pointer to **string** | Indicates the Global Application Entitlement for which this Global Application Entitlement acts as backup. | [optional] 
**RemoteApplicationPoolsCount** | Pointer to **int32** | Count of application pools on remote pods that are associated with this Global Application Entitlement. | [optional] 
**RequireHomeSite** | Pointer to **bool** | Indicates whether we fail if a home site isn&#39;t defined for this Global Application Entitlement. Supported Filters: &#39;Equals&#39;. | [optional] 
**Scope** | Pointer to **string** | Scope for this global application entitlement. Visibility and Placement policies are defined by this value. Supported Filters: &#39;Equals&#39;. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set. | [optional] 
**ShortcutLocationsV2** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the application.  | [optional] 
**SupportedDisplayProtocols** | Pointer to **[]string** | The set of supported display protocols for the Global Application Entitlement. All the application pools associated with this Global Application Entitlement must support these protocols supportedDisplayProtocols . Clients connecting through this Global Application Entitlement that are allowed to select their protocol will see these display protocol options. | [optional] 
**UseHomeSite** | Pointer to **bool** | Indicates whether a pod in the user&#39;s home site is used to start the search or the current site is used. Supported Filters: &#39;Equals&#39;. | [optional] 
**UserCount** | Pointer to **int32** | Count of users that are associated with this Global Application Entitlement. | [optional] 
**UserOrGroupSiteOverrideCount** | Pointer to **int32** | Count of all User Home Site overrides associated with this Global Application Entitlement (for either users or groups). | [optional] 

## Methods

### NewGlobalApplicationEntitlementInfoV2

`func NewGlobalApplicationEntitlementInfoV2() *GlobalApplicationEntitlementInfoV2`

NewGlobalApplicationEntitlementInfoV2 instantiates a new GlobalApplicationEntitlementInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalApplicationEntitlementInfoV2WithDefaults

`func NewGlobalApplicationEntitlementInfoV2WithDefaults() *GlobalApplicationEntitlementInfoV2`

NewGlobalApplicationEntitlementInfoV2WithDefaults instantiates a new GlobalApplicationEntitlementInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsersToChooseProtocol

`func (o *GlobalApplicationEntitlementInfoV2) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *GlobalApplicationEntitlementInfoV2) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *GlobalApplicationEntitlementInfoV2) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.

### HasAllowUsersToChooseProtocol

`func (o *GlobalApplicationEntitlementInfoV2) HasAllowUsersToChooseProtocol() bool`

HasAllowUsersToChooseProtocol returns a boolean if a field has been set.

### GetApplicationData

`func (o *GlobalApplicationEntitlementInfoV2) GetApplicationData() ApplicationData`

GetApplicationData returns the ApplicationData field if non-nil, zero value otherwise.

### GetApplicationDataOk

`func (o *GlobalApplicationEntitlementInfoV2) GetApplicationDataOk() (*ApplicationData, bool)`

GetApplicationDataOk returns a tuple with the ApplicationData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationData

`func (o *GlobalApplicationEntitlementInfoV2) SetApplicationData(v ApplicationData)`

SetApplicationData sets ApplicationData field to given value.

### HasApplicationData

`func (o *GlobalApplicationEntitlementInfoV2) HasApplicationData() bool`

HasApplicationData returns a boolean if a field has been set.

### GetApplicationIconIds

`func (o *GlobalApplicationEntitlementInfoV2) GetApplicationIconIds() []string`

GetApplicationIconIds returns the ApplicationIconIds field if non-nil, zero value otherwise.

### GetApplicationIconIdsOk

`func (o *GlobalApplicationEntitlementInfoV2) GetApplicationIconIdsOk() (*[]string, bool)`

GetApplicationIconIdsOk returns a tuple with the ApplicationIconIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationIconIds

`func (o *GlobalApplicationEntitlementInfoV2) SetApplicationIconIds(v []string)`

SetApplicationIconIds sets ApplicationIconIds field to given value.

### HasApplicationIconIds

`func (o *GlobalApplicationEntitlementInfoV2) HasApplicationIconIds() bool`

HasApplicationIconIds returns a boolean if a field has been set.

### GetBackupGaeId

`func (o *GlobalApplicationEntitlementInfoV2) GetBackupGaeId() string`

GetBackupGaeId returns the BackupGaeId field if non-nil, zero value otherwise.

### GetBackupGaeIdOk

`func (o *GlobalApplicationEntitlementInfoV2) GetBackupGaeIdOk() (*string, bool)`

GetBackupGaeIdOk returns a tuple with the BackupGaeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGaeId

`func (o *GlobalApplicationEntitlementInfoV2) SetBackupGaeId(v string)`

SetBackupGaeId sets BackupGaeId field to given value.

### HasBackupGaeId

`func (o *GlobalApplicationEntitlementInfoV2) HasBackupGaeId() bool`

HasBackupGaeId returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *GlobalApplicationEntitlementInfoV2) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *GlobalApplicationEntitlementInfoV2) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *GlobalApplicationEntitlementInfoV2) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *GlobalApplicationEntitlementInfoV2) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *GlobalApplicationEntitlementInfoV2) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *GlobalApplicationEntitlementInfoV2) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *GlobalApplicationEntitlementInfoV2) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *GlobalApplicationEntitlementInfoV2) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *GlobalApplicationEntitlementInfoV2) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *GlobalApplicationEntitlementInfoV2) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *GlobalApplicationEntitlementInfoV2) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.

### HasDefaultDisplayProtocol

`func (o *GlobalApplicationEntitlementInfoV2) HasDefaultDisplayProtocol() bool`

HasDefaultDisplayProtocol returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalApplicationEntitlementInfoV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalApplicationEntitlementInfoV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalApplicationEntitlementInfoV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalApplicationEntitlementInfoV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *GlobalApplicationEntitlementInfoV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GlobalApplicationEntitlementInfoV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GlobalApplicationEntitlementInfoV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GlobalApplicationEntitlementInfoV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *GlobalApplicationEntitlementInfoV2) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *GlobalApplicationEntitlementInfoV2) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *GlobalApplicationEntitlementInfoV2) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *GlobalApplicationEntitlementInfoV2) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnablePreLaunch

`func (o *GlobalApplicationEntitlementInfoV2) GetEnablePreLaunch() bool`

GetEnablePreLaunch returns the EnablePreLaunch field if non-nil, zero value otherwise.

### GetEnablePreLaunchOk

`func (o *GlobalApplicationEntitlementInfoV2) GetEnablePreLaunchOk() (*bool, bool)`

GetEnablePreLaunchOk returns a tuple with the EnablePreLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePreLaunch

`func (o *GlobalApplicationEntitlementInfoV2) SetEnablePreLaunch(v bool)`

SetEnablePreLaunch sets EnablePreLaunch field to given value.

### HasEnablePreLaunch

`func (o *GlobalApplicationEntitlementInfoV2) HasEnablePreLaunch() bool`

HasEnablePreLaunch returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalApplicationEntitlementInfoV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalApplicationEntitlementInfoV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalApplicationEntitlementInfoV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GlobalApplicationEntitlementInfoV2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFederatedAccessGroupId

`func (o *GlobalApplicationEntitlementInfoV2) GetFederatedAccessGroupId() string`

GetFederatedAccessGroupId returns the FederatedAccessGroupId field if non-nil, zero value otherwise.

### GetFederatedAccessGroupIdOk

`func (o *GlobalApplicationEntitlementInfoV2) GetFederatedAccessGroupIdOk() (*string, bool)`

GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAccessGroupId

`func (o *GlobalApplicationEntitlementInfoV2) SetFederatedAccessGroupId(v string)`

SetFederatedAccessGroupId sets FederatedAccessGroupId field to given value.

### HasFederatedAccessGroupId

`func (o *GlobalApplicationEntitlementInfoV2) HasFederatedAccessGroupId() bool`

HasFederatedAccessGroupId returns a boolean if a field has been set.

### GetGroupCount

`func (o *GlobalApplicationEntitlementInfoV2) GetGroupCount() int32`

GetGroupCount returns the GroupCount field if non-nil, zero value otherwise.

### GetGroupCountOk

`func (o *GlobalApplicationEntitlementInfoV2) GetGroupCountOk() (*int32, bool)`

GetGroupCountOk returns a tuple with the GroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCount

`func (o *GlobalApplicationEntitlementInfoV2) SetGroupCount(v int32)`

SetGroupCount sets GroupCount field to given value.

### HasGroupCount

`func (o *GlobalApplicationEntitlementInfoV2) HasGroupCount() bool`

HasGroupCount returns a boolean if a field has been set.

### GetId

`func (o *GlobalApplicationEntitlementInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalApplicationEntitlementInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalApplicationEntitlementInfoV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalApplicationEntitlementInfoV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalApplicationPoolsCount

`func (o *GlobalApplicationEntitlementInfoV2) GetLocalApplicationPoolsCount() int32`

GetLocalApplicationPoolsCount returns the LocalApplicationPoolsCount field if non-nil, zero value otherwise.

### GetLocalApplicationPoolsCountOk

`func (o *GlobalApplicationEntitlementInfoV2) GetLocalApplicationPoolsCountOk() (*int32, bool)`

GetLocalApplicationPoolsCountOk returns a tuple with the LocalApplicationPoolsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalApplicationPoolsCount

`func (o *GlobalApplicationEntitlementInfoV2) SetLocalApplicationPoolsCount(v int32)`

SetLocalApplicationPoolsCount sets LocalApplicationPoolsCount field to given value.

### HasLocalApplicationPoolsCount

`func (o *GlobalApplicationEntitlementInfoV2) HasLocalApplicationPoolsCount() bool`

HasLocalApplicationPoolsCount returns a boolean if a field has been set.

### GetMemberPods

`func (o *GlobalApplicationEntitlementInfoV2) GetMemberPods() []string`

GetMemberPods returns the MemberPods field if non-nil, zero value otherwise.

### GetMemberPodsOk

`func (o *GlobalApplicationEntitlementInfoV2) GetMemberPodsOk() (*[]string, bool)`

GetMemberPodsOk returns a tuple with the MemberPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPods

`func (o *GlobalApplicationEntitlementInfoV2) SetMemberPods(v []string)`

SetMemberPods sets MemberPods field to given value.

### HasMemberPods

`func (o *GlobalApplicationEntitlementInfoV2) HasMemberPods() bool`

HasMemberPods returns a boolean if a field has been set.

### GetMultiSessionMode

`func (o *GlobalApplicationEntitlementInfoV2) GetMultiSessionMode() string`

GetMultiSessionMode returns the MultiSessionMode field if non-nil, zero value otherwise.

### GetMultiSessionModeOk

`func (o *GlobalApplicationEntitlementInfoV2) GetMultiSessionModeOk() (*string, bool)`

GetMultiSessionModeOk returns a tuple with the MultiSessionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSessionMode

`func (o *GlobalApplicationEntitlementInfoV2) SetMultiSessionMode(v string)`

SetMultiSessionMode sets MultiSessionMode field to given value.

### HasMultiSessionMode

`func (o *GlobalApplicationEntitlementInfoV2) HasMultiSessionMode() bool`

HasMultiSessionMode returns a boolean if a field has been set.

### GetMultipleSessionAutoClean

`func (o *GlobalApplicationEntitlementInfoV2) GetMultipleSessionAutoClean() bool`

GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field if non-nil, zero value otherwise.

### GetMultipleSessionAutoCleanOk

`func (o *GlobalApplicationEntitlementInfoV2) GetMultipleSessionAutoCleanOk() (*bool, bool)`

GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSessionAutoClean

`func (o *GlobalApplicationEntitlementInfoV2) SetMultipleSessionAutoClean(v bool)`

SetMultipleSessionAutoClean sets MultipleSessionAutoClean field to given value.

### HasMultipleSessionAutoClean

`func (o *GlobalApplicationEntitlementInfoV2) HasMultipleSessionAutoClean() bool`

HasMultipleSessionAutoClean returns a boolean if a field has been set.

### GetName

`func (o *GlobalApplicationEntitlementInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalApplicationEntitlementInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalApplicationEntitlementInfoV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalApplicationEntitlementInfoV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryGaeId

`func (o *GlobalApplicationEntitlementInfoV2) GetPrimaryGaeId() string`

GetPrimaryGaeId returns the PrimaryGaeId field if non-nil, zero value otherwise.

### GetPrimaryGaeIdOk

`func (o *GlobalApplicationEntitlementInfoV2) GetPrimaryGaeIdOk() (*string, bool)`

GetPrimaryGaeIdOk returns a tuple with the PrimaryGaeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGaeId

`func (o *GlobalApplicationEntitlementInfoV2) SetPrimaryGaeId(v string)`

SetPrimaryGaeId sets PrimaryGaeId field to given value.

### HasPrimaryGaeId

`func (o *GlobalApplicationEntitlementInfoV2) HasPrimaryGaeId() bool`

HasPrimaryGaeId returns a boolean if a field has been set.

### GetRemoteApplicationPoolsCount

`func (o *GlobalApplicationEntitlementInfoV2) GetRemoteApplicationPoolsCount() int32`

GetRemoteApplicationPoolsCount returns the RemoteApplicationPoolsCount field if non-nil, zero value otherwise.

### GetRemoteApplicationPoolsCountOk

`func (o *GlobalApplicationEntitlementInfoV2) GetRemoteApplicationPoolsCountOk() (*int32, bool)`

GetRemoteApplicationPoolsCountOk returns a tuple with the RemoteApplicationPoolsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteApplicationPoolsCount

`func (o *GlobalApplicationEntitlementInfoV2) SetRemoteApplicationPoolsCount(v int32)`

SetRemoteApplicationPoolsCount sets RemoteApplicationPoolsCount field to given value.

### HasRemoteApplicationPoolsCount

`func (o *GlobalApplicationEntitlementInfoV2) HasRemoteApplicationPoolsCount() bool`

HasRemoteApplicationPoolsCount returns a boolean if a field has been set.

### GetRequireHomeSite

`func (o *GlobalApplicationEntitlementInfoV2) GetRequireHomeSite() bool`

GetRequireHomeSite returns the RequireHomeSite field if non-nil, zero value otherwise.

### GetRequireHomeSiteOk

`func (o *GlobalApplicationEntitlementInfoV2) GetRequireHomeSiteOk() (*bool, bool)`

GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHomeSite

`func (o *GlobalApplicationEntitlementInfoV2) SetRequireHomeSite(v bool)`

SetRequireHomeSite sets RequireHomeSite field to given value.

### HasRequireHomeSite

`func (o *GlobalApplicationEntitlementInfoV2) HasRequireHomeSite() bool`

HasRequireHomeSite returns a boolean if a field has been set.

### GetScope

`func (o *GlobalApplicationEntitlementInfoV2) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalApplicationEntitlementInfoV2) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalApplicationEntitlementInfoV2) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GlobalApplicationEntitlementInfoV2) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetShortcutLocationsV2

`func (o *GlobalApplicationEntitlementInfoV2) GetShortcutLocationsV2() []string`

GetShortcutLocationsV2 returns the ShortcutLocationsV2 field if non-nil, zero value otherwise.

### GetShortcutLocationsV2Ok

`func (o *GlobalApplicationEntitlementInfoV2) GetShortcutLocationsV2Ok() (*[]string, bool)`

GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocationsV2

`func (o *GlobalApplicationEntitlementInfoV2) SetShortcutLocationsV2(v []string)`

SetShortcutLocationsV2 sets ShortcutLocationsV2 field to given value.

### HasShortcutLocationsV2

`func (o *GlobalApplicationEntitlementInfoV2) HasShortcutLocationsV2() bool`

HasShortcutLocationsV2 returns a boolean if a field has been set.

### GetSupportedDisplayProtocols

`func (o *GlobalApplicationEntitlementInfoV2) GetSupportedDisplayProtocols() []string`

GetSupportedDisplayProtocols returns the SupportedDisplayProtocols field if non-nil, zero value otherwise.

### GetSupportedDisplayProtocolsOk

`func (o *GlobalApplicationEntitlementInfoV2) GetSupportedDisplayProtocolsOk() (*[]string, bool)`

GetSupportedDisplayProtocolsOk returns a tuple with the SupportedDisplayProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDisplayProtocols

`func (o *GlobalApplicationEntitlementInfoV2) SetSupportedDisplayProtocols(v []string)`

SetSupportedDisplayProtocols sets SupportedDisplayProtocols field to given value.

### HasSupportedDisplayProtocols

`func (o *GlobalApplicationEntitlementInfoV2) HasSupportedDisplayProtocols() bool`

HasSupportedDisplayProtocols returns a boolean if a field has been set.

### GetUseHomeSite

`func (o *GlobalApplicationEntitlementInfoV2) GetUseHomeSite() bool`

GetUseHomeSite returns the UseHomeSite field if non-nil, zero value otherwise.

### GetUseHomeSiteOk

`func (o *GlobalApplicationEntitlementInfoV2) GetUseHomeSiteOk() (*bool, bool)`

GetUseHomeSiteOk returns a tuple with the UseHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHomeSite

`func (o *GlobalApplicationEntitlementInfoV2) SetUseHomeSite(v bool)`

SetUseHomeSite sets UseHomeSite field to given value.

### HasUseHomeSite

`func (o *GlobalApplicationEntitlementInfoV2) HasUseHomeSite() bool`

HasUseHomeSite returns a boolean if a field has been set.

### GetUserCount

`func (o *GlobalApplicationEntitlementInfoV2) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *GlobalApplicationEntitlementInfoV2) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *GlobalApplicationEntitlementInfoV2) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *GlobalApplicationEntitlementInfoV2) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetUserOrGroupSiteOverrideCount

`func (o *GlobalApplicationEntitlementInfoV2) GetUserOrGroupSiteOverrideCount() int32`

GetUserOrGroupSiteOverrideCount returns the UserOrGroupSiteOverrideCount field if non-nil, zero value otherwise.

### GetUserOrGroupSiteOverrideCountOk

`func (o *GlobalApplicationEntitlementInfoV2) GetUserOrGroupSiteOverrideCountOk() (*int32, bool)`

GetUserOrGroupSiteOverrideCountOk returns a tuple with the UserOrGroupSiteOverrideCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOrGroupSiteOverrideCount

`func (o *GlobalApplicationEntitlementInfoV2) SetUserOrGroupSiteOverrideCount(v int32)`

SetUserOrGroupSiteOverrideCount sets UserOrGroupSiteOverrideCount field to given value.

### HasUserOrGroupSiteOverrideCount

`func (o *GlobalApplicationEntitlementInfoV2) HasUserOrGroupSiteOverrideCount() bool`

HasUserOrGroupSiteOverrideCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


