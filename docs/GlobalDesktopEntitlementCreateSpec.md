# GlobalDesktopEntitlementCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleSessionsPerUser** | Pointer to **bool** | Indicates whether users can have multiple sessions when accessed from different client devices, this is also called Class room mode and applicable only to floating user assignment. If value is set to true, the desktop pools that are associated with this Global Desktop Entitlement  must also allow users to have multiple sessions with allowMultipleSessionsPerUser. Default value is false. | [optional] 
**AllowUsersToChooseProtocol** | Pointer to **bool** | Indicates whether the users can choose the protocol used. If set to true, the desktop pools that are associated with this Global Desktop Entitlement must also allow users to choose display protocol with allowUsersToChooseProtocol. Default value is true. | [optional] 
**AllowUsersToResetMachines** | Pointer to **bool** | Indicates whether users are allowed to reset/restart their machines. If set to true, the desktop pools that are associated with this Global Desktop Entitlement must also allow users to reset/restart machines with allowUsersToResetMachines. Default value is false. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the entitlement. Unset if the entitlement does not belong to a category. | [optional] 
**CloudManaged** | Pointer to **bool** | Indicates whether this global desktop entitlement is managed from cloud. Default value is false. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | List of connection server restriction tags to which the access to the global desktop entitlement is restricted. Empty or null list means that entitlement can be accessed from any connection server | [optional] 
**Dedicated** | Pointer to **bool** | Indicates whether global desktop entitlement is dedicated. If so, only dedicated desktop pools can be associated with this Global Desktop Entitlement. Otherwise, only floating desktop pools, can be associated with it. Can only be set at time of creation. Default value is false. | [optional] 
**DefaultDisplayProtocol** | Pointer to **string** | The default display protocol for the Global Desktop Entitlement. Clients connecting through this Global Desktop Entitlement that do not specify a protocol will use this value, not the value specified directly on the desktop pool to which they connect (if different).This property has a default value of \&quot;PCOIP\&quot;. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | [optional] 
**Description** | Pointer to **string** | Description of Global Desktop Entitlement.  | [optional] 
**DisplayAssignedMachineName** | Pointer to **bool** | Indicates whether users should see the hostname of the machine assigned to them instead of displayName when they connect using Horizon Client. This is applicable for dedicated Global Desktop Entitlements only. If no machine is assigned to the user then \&quot;display_name (No machine assigned)\&quot; will be displayed in the client. Default value is false. | [optional] 
**DisplayMachineAlias** | Pointer to **bool** | Indicates whether users should see the alias of the machine assigned to them instead of display_name when they connect using Horizon Client. This is applicable for dedicated Global Desktop Entitlements only.  Default value is false. | [optional] 
**DisplayName** | Pointer to **string** | Name that users will see when they connect using Horizon Client. If display_name is left blank, it defaults to name.  | [optional] 
**EnableClientRestrictions** | Pointer to **bool** | Indicates whether client restrictions to be applied to Global Desktop Entitlement. Currently it is valid for RDSH pools. Default value is false. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if this Global Desktop Entitlement is enabled.Default value is true. | [optional] 
**MultipleSessionAutoClean** | Pointer to **bool** | Indicates if automatic session clean up is enabled. This cannot be enabled when this Global Desktop Entitlement is associated with a Desktop pool that has dedicated user assignment. Default value is false. | [optional] 
**Name** | **string** | Unique name used to identify the Global Desktop Entitlement.  | 
**RequireHomeSite** | Pointer to **bool** | Indicates whether we fail if a home site isn&#39;t defined for this Global Desktop Entitlement. Default value is false. | [optional] 
**Scope** | Pointer to **string** | Scope for this global desktop entitlement. Visibility and Placement policies are defined by this value. This property has a default value of \&quot;ALL_SITES\&quot;. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set. | [optional] 
**SessionCollaborationEnabled** | Pointer to **bool** | Session collaboration allows a user to share their remote session with other users. Blast must be configured as a supported protocol. Indicates if the desktop pools that are associated with this Global Desktop Entitlement must also have session collaboration enabled with enableCollaboration. Default value is false. | [optional] 
**ShortcutLocations** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop. The value must be set if category_folder_name is provided. | [optional] 
**UseHomeSite** | Pointer to **bool** | Indicates whether a pod in the user&#39;s home site is used to start the search or the current site is used. Default value is false. | [optional] 

## Methods

### NewGlobalDesktopEntitlementCreateSpec

`func NewGlobalDesktopEntitlementCreateSpec(name string, ) *GlobalDesktopEntitlementCreateSpec`

NewGlobalDesktopEntitlementCreateSpec instantiates a new GlobalDesktopEntitlementCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalDesktopEntitlementCreateSpecWithDefaults

`func NewGlobalDesktopEntitlementCreateSpecWithDefaults() *GlobalDesktopEntitlementCreateSpec`

NewGlobalDesktopEntitlementCreateSpecWithDefaults instantiates a new GlobalDesktopEntitlementCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementCreateSpec) GetAllowMultipleSessionsPerUser() bool`

GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field if non-nil, zero value otherwise.

### GetAllowMultipleSessionsPerUserOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetAllowMultipleSessionsPerUserOk() (*bool, bool)`

GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementCreateSpec) SetAllowMultipleSessionsPerUser(v bool)`

SetAllowMultipleSessionsPerUser sets AllowMultipleSessionsPerUser field to given value.

### HasAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementCreateSpec) HasAllowMultipleSessionsPerUser() bool`

HasAllowMultipleSessionsPerUser returns a boolean if a field has been set.

### GetAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementCreateSpec) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementCreateSpec) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.

### HasAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementCreateSpec) HasAllowUsersToChooseProtocol() bool`

HasAllowUsersToChooseProtocol returns a boolean if a field has been set.

### GetAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementCreateSpec) GetAllowUsersToResetMachines() bool`

GetAllowUsersToResetMachines returns the AllowUsersToResetMachines field if non-nil, zero value otherwise.

### GetAllowUsersToResetMachinesOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetAllowUsersToResetMachinesOk() (*bool, bool)`

GetAllowUsersToResetMachinesOk returns a tuple with the AllowUsersToResetMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementCreateSpec) SetAllowUsersToResetMachines(v bool)`

SetAllowUsersToResetMachines sets AllowUsersToResetMachines field to given value.

### HasAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementCreateSpec) HasAllowUsersToResetMachines() bool`

HasAllowUsersToResetMachines returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *GlobalDesktopEntitlementCreateSpec) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *GlobalDesktopEntitlementCreateSpec) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *GlobalDesktopEntitlementCreateSpec) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudManaged

`func (o *GlobalDesktopEntitlementCreateSpec) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *GlobalDesktopEntitlementCreateSpec) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *GlobalDesktopEntitlementCreateSpec) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *GlobalDesktopEntitlementCreateSpec) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *GlobalDesktopEntitlementCreateSpec) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *GlobalDesktopEntitlementCreateSpec) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetDedicated

`func (o *GlobalDesktopEntitlementCreateSpec) GetDedicated() bool`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetDedicatedOk() (*bool, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *GlobalDesktopEntitlementCreateSpec) SetDedicated(v bool)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *GlobalDesktopEntitlementCreateSpec) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementCreateSpec) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementCreateSpec) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.

### HasDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementCreateSpec) HasDefaultDisplayProtocol() bool`

HasDefaultDisplayProtocol returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalDesktopEntitlementCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalDesktopEntitlementCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalDesktopEntitlementCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementCreateSpec) GetDisplayAssignedMachineName() bool`

GetDisplayAssignedMachineName returns the DisplayAssignedMachineName field if non-nil, zero value otherwise.

### GetDisplayAssignedMachineNameOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetDisplayAssignedMachineNameOk() (*bool, bool)`

GetDisplayAssignedMachineNameOk returns a tuple with the DisplayAssignedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementCreateSpec) SetDisplayAssignedMachineName(v bool)`

SetDisplayAssignedMachineName sets DisplayAssignedMachineName field to given value.

### HasDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementCreateSpec) HasDisplayAssignedMachineName() bool`

HasDisplayAssignedMachineName returns a boolean if a field has been set.

### GetDisplayMachineAlias

`func (o *GlobalDesktopEntitlementCreateSpec) GetDisplayMachineAlias() bool`

GetDisplayMachineAlias returns the DisplayMachineAlias field if non-nil, zero value otherwise.

### GetDisplayMachineAliasOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetDisplayMachineAliasOk() (*bool, bool)`

GetDisplayMachineAliasOk returns a tuple with the DisplayMachineAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMachineAlias

`func (o *GlobalDesktopEntitlementCreateSpec) SetDisplayMachineAlias(v bool)`

SetDisplayMachineAlias sets DisplayMachineAlias field to given value.

### HasDisplayMachineAlias

`func (o *GlobalDesktopEntitlementCreateSpec) HasDisplayMachineAlias() bool`

HasDisplayMachineAlias returns a boolean if a field has been set.

### GetDisplayName

`func (o *GlobalDesktopEntitlementCreateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GlobalDesktopEntitlementCreateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GlobalDesktopEntitlementCreateSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *GlobalDesktopEntitlementCreateSpec) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *GlobalDesktopEntitlementCreateSpec) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *GlobalDesktopEntitlementCreateSpec) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalDesktopEntitlementCreateSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalDesktopEntitlementCreateSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GlobalDesktopEntitlementCreateSpec) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementCreateSpec) GetMultipleSessionAutoClean() bool`

GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field if non-nil, zero value otherwise.

### GetMultipleSessionAutoCleanOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetMultipleSessionAutoCleanOk() (*bool, bool)`

GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementCreateSpec) SetMultipleSessionAutoClean(v bool)`

SetMultipleSessionAutoClean sets MultipleSessionAutoClean field to given value.

### HasMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementCreateSpec) HasMultipleSessionAutoClean() bool`

HasMultipleSessionAutoClean returns a boolean if a field has been set.

### GetName

`func (o *GlobalDesktopEntitlementCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalDesktopEntitlementCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetRequireHomeSite

`func (o *GlobalDesktopEntitlementCreateSpec) GetRequireHomeSite() bool`

GetRequireHomeSite returns the RequireHomeSite field if non-nil, zero value otherwise.

### GetRequireHomeSiteOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetRequireHomeSiteOk() (*bool, bool)`

GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHomeSite

`func (o *GlobalDesktopEntitlementCreateSpec) SetRequireHomeSite(v bool)`

SetRequireHomeSite sets RequireHomeSite field to given value.

### HasRequireHomeSite

`func (o *GlobalDesktopEntitlementCreateSpec) HasRequireHomeSite() bool`

HasRequireHomeSite returns a boolean if a field has been set.

### GetScope

`func (o *GlobalDesktopEntitlementCreateSpec) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalDesktopEntitlementCreateSpec) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GlobalDesktopEntitlementCreateSpec) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementCreateSpec) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementCreateSpec) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.

### HasSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementCreateSpec) HasSessionCollaborationEnabled() bool`

HasSessionCollaborationEnabled returns a boolean if a field has been set.

### GetShortcutLocations

`func (o *GlobalDesktopEntitlementCreateSpec) GetShortcutLocations() []string`

GetShortcutLocations returns the ShortcutLocations field if non-nil, zero value otherwise.

### GetShortcutLocationsOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetShortcutLocationsOk() (*[]string, bool)`

GetShortcutLocationsOk returns a tuple with the ShortcutLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocations

`func (o *GlobalDesktopEntitlementCreateSpec) SetShortcutLocations(v []string)`

SetShortcutLocations sets ShortcutLocations field to given value.

### HasShortcutLocations

`func (o *GlobalDesktopEntitlementCreateSpec) HasShortcutLocations() bool`

HasShortcutLocations returns a boolean if a field has been set.

### GetUseHomeSite

`func (o *GlobalDesktopEntitlementCreateSpec) GetUseHomeSite() bool`

GetUseHomeSite returns the UseHomeSite field if non-nil, zero value otherwise.

### GetUseHomeSiteOk

`func (o *GlobalDesktopEntitlementCreateSpec) GetUseHomeSiteOk() (*bool, bool)`

GetUseHomeSiteOk returns a tuple with the UseHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHomeSite

`func (o *GlobalDesktopEntitlementCreateSpec) SetUseHomeSite(v bool)`

SetUseHomeSite sets UseHomeSite field to given value.

### HasUseHomeSite

`func (o *GlobalDesktopEntitlementCreateSpec) HasUseHomeSite() bool`

HasUseHomeSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


