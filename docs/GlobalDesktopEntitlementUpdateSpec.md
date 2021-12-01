# GlobalDesktopEntitlementUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleSessionsPerUser** | **bool** | Indicates whether users can have multiple sessions when accessed from different client devices, this is also called Class room mode and applicable only to floating user assignment. If value is set to true, the desktop pools that are associated with this Global Desktop Entitlement  must also allow users to have multiple sessions. | 
**BackupGdeId** | Pointer to **string** | Global Desktop Entitlement that can be used as backup for this Global Desktop Entitlement. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the entitlement. The property will not be set if the entitlement does not belong to a category. | [optional] 
**CloudManaged** | **bool** | Indicates whether this global desktop entitlement is managed from cloud. | 
**CsRestrictionTags** | Pointer to **[]string** | Connection server restrictions. This is a list of tags that access to the entitlement is restricted to. No list means that the entitlement can be accessed from any connection server. | [optional] 
**DefaultDisplayProtocol** | **string** | The default display protocol for the Global Desktop Entitlement. Clients connecting through this Global Desktop Entitlement that do not specify a protocol will use this value, not the value specified directly on the desktop pool to which they connect (if different). * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | 
**Description** | Pointer to **string** | Description of Global Desktop Entitlement.  | [optional] 
**DisplayAssignedMachineName** | **bool** | Indicates whether users should see the hostname of the machine assigned to them instead of display_name when they connect using Horizon Client. This is applicable for dedicated Global Desktop Entitlements only. If no machine is assigned to the user then \&quot;display_name (No machine assigned)\&quot; will be displayed in the client. | 
**DisplayMachineAlias** | **bool** | Indicates whether users should see the alias of the machine assigned to them instead of display_name when they connect using Horizon Client. This is applicable for dedicated Global Desktop Entitlements only.  | 
**DisplayName** | **string** | Name that users will see when they connect using Horizon Client. If the display name is left blank, while creating or updating the Global Desktop Entitlement display name value will be same as name.  | 
**EnableClientRestrictions** | **bool** | Indicates whether client restrictions to be applied to Global Desktop Entitlement. Currently it is valid for RDSH pools. | 
**Enabled** | **bool** | Indicates if this Global Desktop Entitlement is enabled. | 
**FederatedAccessGroupId** | **string** | ID of the federated access group with which the global desktop entitlement is to be associated. They can also be used for delegated administration. | 
**MultipleSessionAutoClean** | **bool** | Indicates if automatic session clean up is enabled. This cannot be enabled when this Global Desktop Entitlement is associated with a Desktop pool that has dedicated user assignment. | 
**Name** | **string** | Unique name used to identify the Global Desktop Entitlement.  | 
**RequireHomeSite** | **bool** | Indicates whether it should fail if a home site isn&#39;t defined for this Global Desktop Entitlement. | 
**Scope** | **string** | Scope for this global desktop entitlement. Visibility and Placement policies are defined by this value. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set. | 
**SessionCollaborationEnabled** | **bool** | Session collaboration allows a user to share their remote session with other users. Blast must be configured as a supported protocol. Indicates if the desktop pools that are associated with this Global Desktop Entitlement must also have session collaboration enabled with enableCollaboration. | 
**ShortcutLocationsV2** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop. The value must be set if category_folder_name is provided. | [optional] 
**UseHomeSite** | **bool** | Indicates whether a pod in the user&#39;s home site is used to start the search or the current site is used. | 

## Methods

### NewGlobalDesktopEntitlementUpdateSpec

`func NewGlobalDesktopEntitlementUpdateSpec(allowMultipleSessionsPerUser bool, cloudManaged bool, defaultDisplayProtocol string, displayAssignedMachineName bool, displayMachineAlias bool, displayName string, enableClientRestrictions bool, enabled bool, federatedAccessGroupId string, multipleSessionAutoClean bool, name string, requireHomeSite bool, scope string, sessionCollaborationEnabled bool, useHomeSite bool, ) *GlobalDesktopEntitlementUpdateSpec`

NewGlobalDesktopEntitlementUpdateSpec instantiates a new GlobalDesktopEntitlementUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalDesktopEntitlementUpdateSpecWithDefaults

`func NewGlobalDesktopEntitlementUpdateSpecWithDefaults() *GlobalDesktopEntitlementUpdateSpec`

NewGlobalDesktopEntitlementUpdateSpecWithDefaults instantiates a new GlobalDesktopEntitlementUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementUpdateSpec) GetAllowMultipleSessionsPerUser() bool`

GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field if non-nil, zero value otherwise.

### GetAllowMultipleSessionsPerUserOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetAllowMultipleSessionsPerUserOk() (*bool, bool)`

GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementUpdateSpec) SetAllowMultipleSessionsPerUser(v bool)`

SetAllowMultipleSessionsPerUser sets AllowMultipleSessionsPerUser field to given value.


### GetBackupGdeId

`func (o *GlobalDesktopEntitlementUpdateSpec) GetBackupGdeId() string`

GetBackupGdeId returns the BackupGdeId field if non-nil, zero value otherwise.

### GetBackupGdeIdOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetBackupGdeIdOk() (*string, bool)`

GetBackupGdeIdOk returns a tuple with the BackupGdeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGdeId

`func (o *GlobalDesktopEntitlementUpdateSpec) SetBackupGdeId(v string)`

SetBackupGdeId sets BackupGdeId field to given value.

### HasBackupGdeId

`func (o *GlobalDesktopEntitlementUpdateSpec) HasBackupGdeId() bool`

HasBackupGdeId returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *GlobalDesktopEntitlementUpdateSpec) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *GlobalDesktopEntitlementUpdateSpec) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *GlobalDesktopEntitlementUpdateSpec) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudManaged

`func (o *GlobalDesktopEntitlementUpdateSpec) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *GlobalDesktopEntitlementUpdateSpec) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.


### GetCsRestrictionTags

`func (o *GlobalDesktopEntitlementUpdateSpec) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *GlobalDesktopEntitlementUpdateSpec) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *GlobalDesktopEntitlementUpdateSpec) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementUpdateSpec) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.


### GetDescription

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalDesktopEntitlementUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalDesktopEntitlementUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayAssignedMachineName() bool`

GetDisplayAssignedMachineName returns the DisplayAssignedMachineName field if non-nil, zero value otherwise.

### GetDisplayAssignedMachineNameOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayAssignedMachineNameOk() (*bool, bool)`

GetDisplayAssignedMachineNameOk returns a tuple with the DisplayAssignedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementUpdateSpec) SetDisplayAssignedMachineName(v bool)`

SetDisplayAssignedMachineName sets DisplayAssignedMachineName field to given value.


### GetDisplayMachineAlias

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayMachineAlias() bool`

GetDisplayMachineAlias returns the DisplayMachineAlias field if non-nil, zero value otherwise.

### GetDisplayMachineAliasOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayMachineAliasOk() (*bool, bool)`

GetDisplayMachineAliasOk returns a tuple with the DisplayMachineAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMachineAlias

`func (o *GlobalDesktopEntitlementUpdateSpec) SetDisplayMachineAlias(v bool)`

SetDisplayMachineAlias sets DisplayMachineAlias field to given value.


### GetDisplayName

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GlobalDesktopEntitlementUpdateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetEnableClientRestrictions

`func (o *GlobalDesktopEntitlementUpdateSpec) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *GlobalDesktopEntitlementUpdateSpec) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.


### GetEnabled

`func (o *GlobalDesktopEntitlementUpdateSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalDesktopEntitlementUpdateSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFederatedAccessGroupId

`func (o *GlobalDesktopEntitlementUpdateSpec) GetFederatedAccessGroupId() string`

GetFederatedAccessGroupId returns the FederatedAccessGroupId field if non-nil, zero value otherwise.

### GetFederatedAccessGroupIdOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetFederatedAccessGroupIdOk() (*string, bool)`

GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAccessGroupId

`func (o *GlobalDesktopEntitlementUpdateSpec) SetFederatedAccessGroupId(v string)`

SetFederatedAccessGroupId sets FederatedAccessGroupId field to given value.


### GetMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementUpdateSpec) GetMultipleSessionAutoClean() bool`

GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field if non-nil, zero value otherwise.

### GetMultipleSessionAutoCleanOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetMultipleSessionAutoCleanOk() (*bool, bool)`

GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementUpdateSpec) SetMultipleSessionAutoClean(v bool)`

SetMultipleSessionAutoClean sets MultipleSessionAutoClean field to given value.


### GetName

`func (o *GlobalDesktopEntitlementUpdateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalDesktopEntitlementUpdateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetRequireHomeSite

`func (o *GlobalDesktopEntitlementUpdateSpec) GetRequireHomeSite() bool`

GetRequireHomeSite returns the RequireHomeSite field if non-nil, zero value otherwise.

### GetRequireHomeSiteOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetRequireHomeSiteOk() (*bool, bool)`

GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHomeSite

`func (o *GlobalDesktopEntitlementUpdateSpec) SetRequireHomeSite(v bool)`

SetRequireHomeSite sets RequireHomeSite field to given value.


### GetScope

`func (o *GlobalDesktopEntitlementUpdateSpec) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalDesktopEntitlementUpdateSpec) SetScope(v string)`

SetScope sets Scope field to given value.


### GetSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementUpdateSpec) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementUpdateSpec) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.


### GetShortcutLocationsV2

`func (o *GlobalDesktopEntitlementUpdateSpec) GetShortcutLocationsV2() []string`

GetShortcutLocationsV2 returns the ShortcutLocationsV2 field if non-nil, zero value otherwise.

### GetShortcutLocationsV2Ok

`func (o *GlobalDesktopEntitlementUpdateSpec) GetShortcutLocationsV2Ok() (*[]string, bool)`

GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocationsV2

`func (o *GlobalDesktopEntitlementUpdateSpec) SetShortcutLocationsV2(v []string)`

SetShortcutLocationsV2 sets ShortcutLocationsV2 field to given value.

### HasShortcutLocationsV2

`func (o *GlobalDesktopEntitlementUpdateSpec) HasShortcutLocationsV2() bool`

HasShortcutLocationsV2 returns a boolean if a field has been set.

### GetUseHomeSite

`func (o *GlobalDesktopEntitlementUpdateSpec) GetUseHomeSite() bool`

GetUseHomeSite returns the UseHomeSite field if non-nil, zero value otherwise.

### GetUseHomeSiteOk

`func (o *GlobalDesktopEntitlementUpdateSpec) GetUseHomeSiteOk() (*bool, bool)`

GetUseHomeSiteOk returns a tuple with the UseHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHomeSite

`func (o *GlobalDesktopEntitlementUpdateSpec) SetUseHomeSite(v bool)`

SetUseHomeSite sets UseHomeSite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


