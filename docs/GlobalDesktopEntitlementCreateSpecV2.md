# GlobalDesktopEntitlementCreateSpecV2

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
**FederatedAccessGroupId** | **string** | ID of the federated access group with which the global desktop entitlement is to be associated. They can also be used for delegated administration. | 
**MultipleSessionAutoClean** | Pointer to **bool** | Indicates if automatic session clean up is enabled. This cannot be enabled when this Global Desktop Entitlement is associated with a Desktop pool that has dedicated user assignment. Default value is false. | [optional] 
**Name** | **string** | Unique name used to identify the Global Desktop Entitlement.  | 
**RequireHomeSite** | Pointer to **bool** | Indicates whether we fail if a home site isn&#39;t defined for this Global Desktop Entitlement. Default value is false. | [optional] 
**Scope** | Pointer to **string** | Scope for this global desktop entitlement. Visibility and Placement policies are defined by this value. This property has a default value of \&quot;ALL_SITES\&quot;. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set. | [optional] 
**SessionCollaborationEnabled** | Pointer to **bool** | Session collaboration allows a user to share their remote session with other users. Blast must be configured as a supported protocol. Indicates if the desktop pools that are associated with this Global Desktop Entitlement must also have session collaboration enabled with enableCollaboration. Default value is false. | [optional] 
**ShortcutLocationsV2** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop. This property is required if category_folder_name is set. | [optional] 
**UseHomeSite** | Pointer to **bool** | Indicates whether a pod in the user&#39;s home site is used to start the search or the current site is used. Default value is false. | [optional] 

## Methods

### NewGlobalDesktopEntitlementCreateSpecV2

`func NewGlobalDesktopEntitlementCreateSpecV2(federatedAccessGroupId string, name string, ) *GlobalDesktopEntitlementCreateSpecV2`

NewGlobalDesktopEntitlementCreateSpecV2 instantiates a new GlobalDesktopEntitlementCreateSpecV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalDesktopEntitlementCreateSpecV2WithDefaults

`func NewGlobalDesktopEntitlementCreateSpecV2WithDefaults() *GlobalDesktopEntitlementCreateSpecV2`

NewGlobalDesktopEntitlementCreateSpecV2WithDefaults instantiates a new GlobalDesktopEntitlementCreateSpecV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetAllowMultipleSessionsPerUser() bool`

GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field if non-nil, zero value otherwise.

### GetAllowMultipleSessionsPerUserOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetAllowMultipleSessionsPerUserOk() (*bool, bool)`

GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetAllowMultipleSessionsPerUser(v bool)`

SetAllowMultipleSessionsPerUser sets AllowMultipleSessionsPerUser field to given value.

### HasAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasAllowMultipleSessionsPerUser() bool`

HasAllowMultipleSessionsPerUser returns a boolean if a field has been set.

### GetAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.

### HasAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasAllowUsersToChooseProtocol() bool`

HasAllowUsersToChooseProtocol returns a boolean if a field has been set.

### GetAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetAllowUsersToResetMachines() bool`

GetAllowUsersToResetMachines returns the AllowUsersToResetMachines field if non-nil, zero value otherwise.

### GetAllowUsersToResetMachinesOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetAllowUsersToResetMachinesOk() (*bool, bool)`

GetAllowUsersToResetMachinesOk returns a tuple with the AllowUsersToResetMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetAllowUsersToResetMachines(v bool)`

SetAllowUsersToResetMachines sets AllowUsersToResetMachines field to given value.

### HasAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasAllowUsersToResetMachines() bool`

HasAllowUsersToResetMachines returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudManaged

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetDedicated

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDedicated() bool`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDedicatedOk() (*bool, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetDedicated(v bool)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.

### HasDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasDefaultDisplayProtocol() bool`

HasDefaultDisplayProtocol returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDisplayAssignedMachineName() bool`

GetDisplayAssignedMachineName returns the DisplayAssignedMachineName field if non-nil, zero value otherwise.

### GetDisplayAssignedMachineNameOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDisplayAssignedMachineNameOk() (*bool, bool)`

GetDisplayAssignedMachineNameOk returns a tuple with the DisplayAssignedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetDisplayAssignedMachineName(v bool)`

SetDisplayAssignedMachineName sets DisplayAssignedMachineName field to given value.

### HasDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasDisplayAssignedMachineName() bool`

HasDisplayAssignedMachineName returns a boolean if a field has been set.

### GetDisplayMachineAlias

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDisplayMachineAlias() bool`

GetDisplayMachineAlias returns the DisplayMachineAlias field if non-nil, zero value otherwise.

### GetDisplayMachineAliasOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDisplayMachineAliasOk() (*bool, bool)`

GetDisplayMachineAliasOk returns a tuple with the DisplayMachineAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMachineAlias

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetDisplayMachineAlias(v bool)`

SetDisplayMachineAlias sets DisplayMachineAlias field to given value.

### HasDisplayMachineAlias

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasDisplayMachineAlias() bool`

HasDisplayMachineAlias returns a boolean if a field has been set.

### GetDisplayName

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFederatedAccessGroupId

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetFederatedAccessGroupId() string`

GetFederatedAccessGroupId returns the FederatedAccessGroupId field if non-nil, zero value otherwise.

### GetFederatedAccessGroupIdOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetFederatedAccessGroupIdOk() (*string, bool)`

GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAccessGroupId

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetFederatedAccessGroupId(v string)`

SetFederatedAccessGroupId sets FederatedAccessGroupId field to given value.


### GetMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetMultipleSessionAutoClean() bool`

GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field if non-nil, zero value otherwise.

### GetMultipleSessionAutoCleanOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetMultipleSessionAutoCleanOk() (*bool, bool)`

GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetMultipleSessionAutoClean(v bool)`

SetMultipleSessionAutoClean sets MultipleSessionAutoClean field to given value.

### HasMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasMultipleSessionAutoClean() bool`

HasMultipleSessionAutoClean returns a boolean if a field has been set.

### GetName

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetName(v string)`

SetName sets Name field to given value.


### GetRequireHomeSite

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetRequireHomeSite() bool`

GetRequireHomeSite returns the RequireHomeSite field if non-nil, zero value otherwise.

### GetRequireHomeSiteOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetRequireHomeSiteOk() (*bool, bool)`

GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHomeSite

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetRequireHomeSite(v bool)`

SetRequireHomeSite sets RequireHomeSite field to given value.

### HasRequireHomeSite

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasRequireHomeSite() bool`

HasRequireHomeSite returns a boolean if a field has been set.

### GetScope

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.

### HasSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasSessionCollaborationEnabled() bool`

HasSessionCollaborationEnabled returns a boolean if a field has been set.

### GetShortcutLocationsV2

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetShortcutLocationsV2() []string`

GetShortcutLocationsV2 returns the ShortcutLocationsV2 field if non-nil, zero value otherwise.

### GetShortcutLocationsV2Ok

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetShortcutLocationsV2Ok() (*[]string, bool)`

GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocationsV2

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetShortcutLocationsV2(v []string)`

SetShortcutLocationsV2 sets ShortcutLocationsV2 field to given value.

### HasShortcutLocationsV2

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasShortcutLocationsV2() bool`

HasShortcutLocationsV2 returns a boolean if a field has been set.

### GetUseHomeSite

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetUseHomeSite() bool`

GetUseHomeSite returns the UseHomeSite field if non-nil, zero value otherwise.

### GetUseHomeSiteOk

`func (o *GlobalDesktopEntitlementCreateSpecV2) GetUseHomeSiteOk() (*bool, bool)`

GetUseHomeSiteOk returns a tuple with the UseHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHomeSite

`func (o *GlobalDesktopEntitlementCreateSpecV2) SetUseHomeSite(v bool)`

SetUseHomeSite sets UseHomeSite field to given value.

### HasUseHomeSite

`func (o *GlobalDesktopEntitlementCreateSpecV2) HasUseHomeSite() bool`

HasUseHomeSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


