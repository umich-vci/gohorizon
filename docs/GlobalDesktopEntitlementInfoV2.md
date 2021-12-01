# GlobalDesktopEntitlementInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMultipleSessionsPerUser** | Pointer to **bool** | Indicates whether users can have multiple sessions when accessed from different client devices, this is also called Class room mode and applicable only to floating user assignment. If value is set to true, the desktop pools that are associated with this Global Desktop Entitlement  must also allow users to have multiple sessions with allowMultipleSessionsPerUser Supported Filters: &#39;Equals&#39;. | [optional] 
**AllowUsersToChooseProtocol** | Pointer to **bool** | Indicates whether the users can choose the protocol used. If set to true, the desktop pools that are associated with this Global Desktop Entitlement must also allow users to choose display protocol with allowUsersToChooseProtocol. Supported Filters: &#39;Equals&#39;. | [optional] 
**AllowUsersToResetMachines** | Pointer to **bool** | Indicates whether users are allowed to reset/restart their machines. If set to true, the desktop pools that are associated with this Global Desktop Entitlement must also allow users to reset/restart machines with allowUsersToResetMachines. Supported Filters: &#39;Equals&#39;. | [optional] 
**BackupGdeId** | Pointer to **string** | Indicates the Global Desktop Entitlement that can be used as backup for this Global Desktop Entitlement. Supported Filters: &#39;Equals&#39;. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the entitlement. Unset if the entitlement does not belong to a category. | [optional] 
**CloudManaged** | Pointer to **bool** | Indicates whether this global desktop entitlement is managed from cloud. Supported Filters: &#39;Equals&#39;. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | Connection server restrictions. This is a list of tags that access to the entitlement is restricted to. No list means that the entitlement can be accessed from any connection server. | [optional] 
**Dedicated** | Pointer to **bool** | Indicates whether global desktop entitlement is dedicated. If so, only dedicated desktop pools can be associated with this Global Desktop Entitlement. Otherwise, only floating desktop pools, can be associated with it. Can only be set at time of creation. Supported Filters: &#39;Equals&#39;. | [optional] 
**DefaultDisplayProtocol** | Pointer to **string** | The default display protocol for the Global Desktop Entitlement. Must be a protocol in the supportedDisplayProtocols list. Clients connecting through this Global Desktop Entitlement that do not specify a protocol will use this value, not the value specified directly on the desktop pool to which they connect (if different). Supported Filters: &#39;Equals&#39;. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | [optional] 
**Description** | Pointer to **string** | Description of Global Desktop Entitlement. This property has a maximum length of 1024 characters. Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DisplayAssignedMachineName** | Pointer to **bool** | Indicates whether users should see the hostname of the machine assigned to them instead of displayName when they connect using Horizon Client. This is applicable for dedicated Global Desktop Entitlements only. If no machine is assigned to the user then \&quot;display_name (No machine assigned)\&quot; will be displayed in the client. Supported Filters: &#39;Equals&#39;. | [optional] 
**DisplayMachineAlias** | Pointer to **bool** | Decides the visibility of the machine alias to the user. | [optional] 
**DisplayName** | Pointer to **string** | Name that users will see when they connect using Horizon Client. If the display name is left blank, while creating or updating the Global Desktop Entitlement display name value will be same as name. This property has a maximum length of 64 characters. Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**EnableClientRestrictions** | Pointer to **bool** | Indicates whether client restrictions to be applied to Global Desktop Entitlement. Currently it is valid for RDSH pools. Supported Filters: &#39;Equals&#39;. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if this Global Desktop Entitlement is enabled. Supported Filters: &#39;Equals&#39;. | [optional] 
**FederatedAccessGroupId** | Pointer to **string** | This represents id of the federated access group associated with the global desktop entitlement. | [optional] 
**GroupCount** | Pointer to **int32** | Count of groups that are associated with this Global Desktop Entitlement. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this Global Desktop Entitlement. Supported Filters: &#39;Equals&#39;. | [optional] 
**LocalDesktopPoolsCount** | Pointer to **int32** | Count of desktop pools local to this pod that are associated with this Global Desktop Entitlement. | [optional] 
**MemberPods** | Pointer to **[]string** | Pods that have desktop pools associated with this Global Desktop Entitlement. | [optional] 
**MultipleSessionAutoClean** | Pointer to **bool** | Indicates if automatic session clean up is enabled. This cannot be enabled when this Global Desktop Entitlement is associated with a Desktop pool that has dedicated user assignment. Supported Filters: &#39;Equals&#39;. | [optional] 
**Name** | Pointer to **string** | Unique name used to identify the Global Desktop Entitlement. This property has a maximum length of 64 characters. Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**PrimaryGdeId** | Pointer to **string** | Indicates the Global Desktop Entitlement for which this Global Desktop Entitlement acts as backup. | [optional] 
**RemoteDesktopPoolsCount** | Pointer to **int32** | Count of desktop pools on remote pods that are associated with this Global Desktop Entitlement. | [optional] 
**RequireHomeSite** | Pointer to **bool** | Indicates whether we fail if a home site isn&#39;t defined for this Global Desktop Entitlement. Supported Filters: &#39;Equals&#39;. | [optional] 
**Scope** | Pointer to **string** | Scope for this global desktop entitlement. Visibility and Placement policies are defined by this value. Supported Filters: &#39;Equals&#39;. * WITHIN_POD: Within POD Policy: Local pod will be used for this policy. If this policy is for visibility, search for existing session will happen only in local pod. If this policy is for placement, session will always be placed on local pod. * WITHIN_SITE: Within Site Policy: Site will be used for this policy. If this policy is for visibility, search for existing session will happen only from site. If this policy is for placement, session will be placed on site. * ALL_SITES: All Sites Policy: Any pod can be used for this action. If this policy is for visibility, search for existing session will span all pods in LMV set. If this policy is for placement, session can be placed on any pod in LMV set. | [optional] 
**SessionCollaborationEnabled** | Pointer to **bool** | Session collaboration allows a user to share their remote session with other users. Blast must be configured as a supported protocol in supportedDisplayProtocols. Indicates if the desktop pools that are associated with this Global Desktop Entitlement must also have session collaboration enabled with enableCollaboration. Supported Filters: &#39;Equals&#39;. | [optional] 
**ShortcutLocationsV2** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop.  | [optional] 
**SupportedDisplayProtocols** | Pointer to **[]string** | The set of supported display protocols for the Global Desktop Entitlement. All the desktop pools associated with this Global Desktop Entitlement must support these protocols supportedDisplayProtocols . Clients connecting through this Global Desktop Entitlement that are allowed to select their protocol will see these display protocol options. | [optional] 
**UseHomeSite** | Pointer to **bool** | Indicates whether a pod in the user&#39;s home site is used to start the search or the current site is used. Supported Filters: &#39;Equals&#39;. | [optional] 
**UserCount** | Pointer to **int32** | Count of users that are associated with this Global Desktop Entitlement. | [optional] 
**UserOrGroupSiteOverrideCount** | Pointer to **int32** | Count of all User Home Site overrides associated with this Global Desktop Entitlement (for either users or groups). | [optional] 

## Methods

### NewGlobalDesktopEntitlementInfoV2

`func NewGlobalDesktopEntitlementInfoV2() *GlobalDesktopEntitlementInfoV2`

NewGlobalDesktopEntitlementInfoV2 instantiates a new GlobalDesktopEntitlementInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalDesktopEntitlementInfoV2WithDefaults

`func NewGlobalDesktopEntitlementInfoV2WithDefaults() *GlobalDesktopEntitlementInfoV2`

NewGlobalDesktopEntitlementInfoV2WithDefaults instantiates a new GlobalDesktopEntitlementInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementInfoV2) GetAllowMultipleSessionsPerUser() bool`

GetAllowMultipleSessionsPerUser returns the AllowMultipleSessionsPerUser field if non-nil, zero value otherwise.

### GetAllowMultipleSessionsPerUserOk

`func (o *GlobalDesktopEntitlementInfoV2) GetAllowMultipleSessionsPerUserOk() (*bool, bool)`

GetAllowMultipleSessionsPerUserOk returns a tuple with the AllowMultipleSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementInfoV2) SetAllowMultipleSessionsPerUser(v bool)`

SetAllowMultipleSessionsPerUser sets AllowMultipleSessionsPerUser field to given value.

### HasAllowMultipleSessionsPerUser

`func (o *GlobalDesktopEntitlementInfoV2) HasAllowMultipleSessionsPerUser() bool`

HasAllowMultipleSessionsPerUser returns a boolean if a field has been set.

### GetAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementInfoV2) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *GlobalDesktopEntitlementInfoV2) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementInfoV2) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.

### HasAllowUsersToChooseProtocol

`func (o *GlobalDesktopEntitlementInfoV2) HasAllowUsersToChooseProtocol() bool`

HasAllowUsersToChooseProtocol returns a boolean if a field has been set.

### GetAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementInfoV2) GetAllowUsersToResetMachines() bool`

GetAllowUsersToResetMachines returns the AllowUsersToResetMachines field if non-nil, zero value otherwise.

### GetAllowUsersToResetMachinesOk

`func (o *GlobalDesktopEntitlementInfoV2) GetAllowUsersToResetMachinesOk() (*bool, bool)`

GetAllowUsersToResetMachinesOk returns a tuple with the AllowUsersToResetMachines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementInfoV2) SetAllowUsersToResetMachines(v bool)`

SetAllowUsersToResetMachines sets AllowUsersToResetMachines field to given value.

### HasAllowUsersToResetMachines

`func (o *GlobalDesktopEntitlementInfoV2) HasAllowUsersToResetMachines() bool`

HasAllowUsersToResetMachines returns a boolean if a field has been set.

### GetBackupGdeId

`func (o *GlobalDesktopEntitlementInfoV2) GetBackupGdeId() string`

GetBackupGdeId returns the BackupGdeId field if non-nil, zero value otherwise.

### GetBackupGdeIdOk

`func (o *GlobalDesktopEntitlementInfoV2) GetBackupGdeIdOk() (*string, bool)`

GetBackupGdeIdOk returns a tuple with the BackupGdeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupGdeId

`func (o *GlobalDesktopEntitlementInfoV2) SetBackupGdeId(v string)`

SetBackupGdeId sets BackupGdeId field to given value.

### HasBackupGdeId

`func (o *GlobalDesktopEntitlementInfoV2) HasBackupGdeId() bool`

HasBackupGdeId returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *GlobalDesktopEntitlementInfoV2) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *GlobalDesktopEntitlementInfoV2) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *GlobalDesktopEntitlementInfoV2) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *GlobalDesktopEntitlementInfoV2) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudManaged

`func (o *GlobalDesktopEntitlementInfoV2) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *GlobalDesktopEntitlementInfoV2) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *GlobalDesktopEntitlementInfoV2) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *GlobalDesktopEntitlementInfoV2) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *GlobalDesktopEntitlementInfoV2) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *GlobalDesktopEntitlementInfoV2) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *GlobalDesktopEntitlementInfoV2) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *GlobalDesktopEntitlementInfoV2) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetDedicated

`func (o *GlobalDesktopEntitlementInfoV2) GetDedicated() bool`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *GlobalDesktopEntitlementInfoV2) GetDedicatedOk() (*bool, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *GlobalDesktopEntitlementInfoV2) SetDedicated(v bool)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *GlobalDesktopEntitlementInfoV2) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementInfoV2) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *GlobalDesktopEntitlementInfoV2) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementInfoV2) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.

### HasDefaultDisplayProtocol

`func (o *GlobalDesktopEntitlementInfoV2) HasDefaultDisplayProtocol() bool`

HasDefaultDisplayProtocol returns a boolean if a field has been set.

### GetDescription

`func (o *GlobalDesktopEntitlementInfoV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GlobalDesktopEntitlementInfoV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GlobalDesktopEntitlementInfoV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GlobalDesktopEntitlementInfoV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementInfoV2) GetDisplayAssignedMachineName() bool`

GetDisplayAssignedMachineName returns the DisplayAssignedMachineName field if non-nil, zero value otherwise.

### GetDisplayAssignedMachineNameOk

`func (o *GlobalDesktopEntitlementInfoV2) GetDisplayAssignedMachineNameOk() (*bool, bool)`

GetDisplayAssignedMachineNameOk returns a tuple with the DisplayAssignedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementInfoV2) SetDisplayAssignedMachineName(v bool)`

SetDisplayAssignedMachineName sets DisplayAssignedMachineName field to given value.

### HasDisplayAssignedMachineName

`func (o *GlobalDesktopEntitlementInfoV2) HasDisplayAssignedMachineName() bool`

HasDisplayAssignedMachineName returns a boolean if a field has been set.

### GetDisplayMachineAlias

`func (o *GlobalDesktopEntitlementInfoV2) GetDisplayMachineAlias() bool`

GetDisplayMachineAlias returns the DisplayMachineAlias field if non-nil, zero value otherwise.

### GetDisplayMachineAliasOk

`func (o *GlobalDesktopEntitlementInfoV2) GetDisplayMachineAliasOk() (*bool, bool)`

GetDisplayMachineAliasOk returns a tuple with the DisplayMachineAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMachineAlias

`func (o *GlobalDesktopEntitlementInfoV2) SetDisplayMachineAlias(v bool)`

SetDisplayMachineAlias sets DisplayMachineAlias field to given value.

### HasDisplayMachineAlias

`func (o *GlobalDesktopEntitlementInfoV2) HasDisplayMachineAlias() bool`

HasDisplayMachineAlias returns a boolean if a field has been set.

### GetDisplayName

`func (o *GlobalDesktopEntitlementInfoV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GlobalDesktopEntitlementInfoV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GlobalDesktopEntitlementInfoV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GlobalDesktopEntitlementInfoV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *GlobalDesktopEntitlementInfoV2) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *GlobalDesktopEntitlementInfoV2) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *GlobalDesktopEntitlementInfoV2) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *GlobalDesktopEntitlementInfoV2) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnabled

`func (o *GlobalDesktopEntitlementInfoV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GlobalDesktopEntitlementInfoV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GlobalDesktopEntitlementInfoV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GlobalDesktopEntitlementInfoV2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFederatedAccessGroupId

`func (o *GlobalDesktopEntitlementInfoV2) GetFederatedAccessGroupId() string`

GetFederatedAccessGroupId returns the FederatedAccessGroupId field if non-nil, zero value otherwise.

### GetFederatedAccessGroupIdOk

`func (o *GlobalDesktopEntitlementInfoV2) GetFederatedAccessGroupIdOk() (*string, bool)`

GetFederatedAccessGroupIdOk returns a tuple with the FederatedAccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAccessGroupId

`func (o *GlobalDesktopEntitlementInfoV2) SetFederatedAccessGroupId(v string)`

SetFederatedAccessGroupId sets FederatedAccessGroupId field to given value.

### HasFederatedAccessGroupId

`func (o *GlobalDesktopEntitlementInfoV2) HasFederatedAccessGroupId() bool`

HasFederatedAccessGroupId returns a boolean if a field has been set.

### GetGroupCount

`func (o *GlobalDesktopEntitlementInfoV2) GetGroupCount() int32`

GetGroupCount returns the GroupCount field if non-nil, zero value otherwise.

### GetGroupCountOk

`func (o *GlobalDesktopEntitlementInfoV2) GetGroupCountOk() (*int32, bool)`

GetGroupCountOk returns a tuple with the GroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupCount

`func (o *GlobalDesktopEntitlementInfoV2) SetGroupCount(v int32)`

SetGroupCount sets GroupCount field to given value.

### HasGroupCount

`func (o *GlobalDesktopEntitlementInfoV2) HasGroupCount() bool`

HasGroupCount returns a boolean if a field has been set.

### GetId

`func (o *GlobalDesktopEntitlementInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalDesktopEntitlementInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalDesktopEntitlementInfoV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalDesktopEntitlementInfoV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocalDesktopPoolsCount

`func (o *GlobalDesktopEntitlementInfoV2) GetLocalDesktopPoolsCount() int32`

GetLocalDesktopPoolsCount returns the LocalDesktopPoolsCount field if non-nil, zero value otherwise.

### GetLocalDesktopPoolsCountOk

`func (o *GlobalDesktopEntitlementInfoV2) GetLocalDesktopPoolsCountOk() (*int32, bool)`

GetLocalDesktopPoolsCountOk returns a tuple with the LocalDesktopPoolsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalDesktopPoolsCount

`func (o *GlobalDesktopEntitlementInfoV2) SetLocalDesktopPoolsCount(v int32)`

SetLocalDesktopPoolsCount sets LocalDesktopPoolsCount field to given value.

### HasLocalDesktopPoolsCount

`func (o *GlobalDesktopEntitlementInfoV2) HasLocalDesktopPoolsCount() bool`

HasLocalDesktopPoolsCount returns a boolean if a field has been set.

### GetMemberPods

`func (o *GlobalDesktopEntitlementInfoV2) GetMemberPods() []string`

GetMemberPods returns the MemberPods field if non-nil, zero value otherwise.

### GetMemberPodsOk

`func (o *GlobalDesktopEntitlementInfoV2) GetMemberPodsOk() (*[]string, bool)`

GetMemberPodsOk returns a tuple with the MemberPods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPods

`func (o *GlobalDesktopEntitlementInfoV2) SetMemberPods(v []string)`

SetMemberPods sets MemberPods field to given value.

### HasMemberPods

`func (o *GlobalDesktopEntitlementInfoV2) HasMemberPods() bool`

HasMemberPods returns a boolean if a field has been set.

### GetMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementInfoV2) GetMultipleSessionAutoClean() bool`

GetMultipleSessionAutoClean returns the MultipleSessionAutoClean field if non-nil, zero value otherwise.

### GetMultipleSessionAutoCleanOk

`func (o *GlobalDesktopEntitlementInfoV2) GetMultipleSessionAutoCleanOk() (*bool, bool)`

GetMultipleSessionAutoCleanOk returns a tuple with the MultipleSessionAutoClean field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementInfoV2) SetMultipleSessionAutoClean(v bool)`

SetMultipleSessionAutoClean sets MultipleSessionAutoClean field to given value.

### HasMultipleSessionAutoClean

`func (o *GlobalDesktopEntitlementInfoV2) HasMultipleSessionAutoClean() bool`

HasMultipleSessionAutoClean returns a boolean if a field has been set.

### GetName

`func (o *GlobalDesktopEntitlementInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalDesktopEntitlementInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalDesktopEntitlementInfoV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalDesktopEntitlementInfoV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryGdeId

`func (o *GlobalDesktopEntitlementInfoV2) GetPrimaryGdeId() string`

GetPrimaryGdeId returns the PrimaryGdeId field if non-nil, zero value otherwise.

### GetPrimaryGdeIdOk

`func (o *GlobalDesktopEntitlementInfoV2) GetPrimaryGdeIdOk() (*string, bool)`

GetPrimaryGdeIdOk returns a tuple with the PrimaryGdeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryGdeId

`func (o *GlobalDesktopEntitlementInfoV2) SetPrimaryGdeId(v string)`

SetPrimaryGdeId sets PrimaryGdeId field to given value.

### HasPrimaryGdeId

`func (o *GlobalDesktopEntitlementInfoV2) HasPrimaryGdeId() bool`

HasPrimaryGdeId returns a boolean if a field has been set.

### GetRemoteDesktopPoolsCount

`func (o *GlobalDesktopEntitlementInfoV2) GetRemoteDesktopPoolsCount() int32`

GetRemoteDesktopPoolsCount returns the RemoteDesktopPoolsCount field if non-nil, zero value otherwise.

### GetRemoteDesktopPoolsCountOk

`func (o *GlobalDesktopEntitlementInfoV2) GetRemoteDesktopPoolsCountOk() (*int32, bool)`

GetRemoteDesktopPoolsCountOk returns a tuple with the RemoteDesktopPoolsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteDesktopPoolsCount

`func (o *GlobalDesktopEntitlementInfoV2) SetRemoteDesktopPoolsCount(v int32)`

SetRemoteDesktopPoolsCount sets RemoteDesktopPoolsCount field to given value.

### HasRemoteDesktopPoolsCount

`func (o *GlobalDesktopEntitlementInfoV2) HasRemoteDesktopPoolsCount() bool`

HasRemoteDesktopPoolsCount returns a boolean if a field has been set.

### GetRequireHomeSite

`func (o *GlobalDesktopEntitlementInfoV2) GetRequireHomeSite() bool`

GetRequireHomeSite returns the RequireHomeSite field if non-nil, zero value otherwise.

### GetRequireHomeSiteOk

`func (o *GlobalDesktopEntitlementInfoV2) GetRequireHomeSiteOk() (*bool, bool)`

GetRequireHomeSiteOk returns a tuple with the RequireHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireHomeSite

`func (o *GlobalDesktopEntitlementInfoV2) SetRequireHomeSite(v bool)`

SetRequireHomeSite sets RequireHomeSite field to given value.

### HasRequireHomeSite

`func (o *GlobalDesktopEntitlementInfoV2) HasRequireHomeSite() bool`

HasRequireHomeSite returns a boolean if a field has been set.

### GetScope

`func (o *GlobalDesktopEntitlementInfoV2) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GlobalDesktopEntitlementInfoV2) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GlobalDesktopEntitlementInfoV2) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *GlobalDesktopEntitlementInfoV2) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementInfoV2) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *GlobalDesktopEntitlementInfoV2) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementInfoV2) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.

### HasSessionCollaborationEnabled

`func (o *GlobalDesktopEntitlementInfoV2) HasSessionCollaborationEnabled() bool`

HasSessionCollaborationEnabled returns a boolean if a field has been set.

### GetShortcutLocationsV2

`func (o *GlobalDesktopEntitlementInfoV2) GetShortcutLocationsV2() []string`

GetShortcutLocationsV2 returns the ShortcutLocationsV2 field if non-nil, zero value otherwise.

### GetShortcutLocationsV2Ok

`func (o *GlobalDesktopEntitlementInfoV2) GetShortcutLocationsV2Ok() (*[]string, bool)`

GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocationsV2

`func (o *GlobalDesktopEntitlementInfoV2) SetShortcutLocationsV2(v []string)`

SetShortcutLocationsV2 sets ShortcutLocationsV2 field to given value.

### HasShortcutLocationsV2

`func (o *GlobalDesktopEntitlementInfoV2) HasShortcutLocationsV2() bool`

HasShortcutLocationsV2 returns a boolean if a field has been set.

### GetSupportedDisplayProtocols

`func (o *GlobalDesktopEntitlementInfoV2) GetSupportedDisplayProtocols() []string`

GetSupportedDisplayProtocols returns the SupportedDisplayProtocols field if non-nil, zero value otherwise.

### GetSupportedDisplayProtocolsOk

`func (o *GlobalDesktopEntitlementInfoV2) GetSupportedDisplayProtocolsOk() (*[]string, bool)`

GetSupportedDisplayProtocolsOk returns a tuple with the SupportedDisplayProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedDisplayProtocols

`func (o *GlobalDesktopEntitlementInfoV2) SetSupportedDisplayProtocols(v []string)`

SetSupportedDisplayProtocols sets SupportedDisplayProtocols field to given value.

### HasSupportedDisplayProtocols

`func (o *GlobalDesktopEntitlementInfoV2) HasSupportedDisplayProtocols() bool`

HasSupportedDisplayProtocols returns a boolean if a field has been set.

### GetUseHomeSite

`func (o *GlobalDesktopEntitlementInfoV2) GetUseHomeSite() bool`

GetUseHomeSite returns the UseHomeSite field if non-nil, zero value otherwise.

### GetUseHomeSiteOk

`func (o *GlobalDesktopEntitlementInfoV2) GetUseHomeSiteOk() (*bool, bool)`

GetUseHomeSiteOk returns a tuple with the UseHomeSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHomeSite

`func (o *GlobalDesktopEntitlementInfoV2) SetUseHomeSite(v bool)`

SetUseHomeSite sets UseHomeSite field to given value.

### HasUseHomeSite

`func (o *GlobalDesktopEntitlementInfoV2) HasUseHomeSite() bool`

HasUseHomeSite returns a boolean if a field has been set.

### GetUserCount

`func (o *GlobalDesktopEntitlementInfoV2) GetUserCount() int32`

GetUserCount returns the UserCount field if non-nil, zero value otherwise.

### GetUserCountOk

`func (o *GlobalDesktopEntitlementInfoV2) GetUserCountOk() (*int32, bool)`

GetUserCountOk returns a tuple with the UserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCount

`func (o *GlobalDesktopEntitlementInfoV2) SetUserCount(v int32)`

SetUserCount sets UserCount field to given value.

### HasUserCount

`func (o *GlobalDesktopEntitlementInfoV2) HasUserCount() bool`

HasUserCount returns a boolean if a field has been set.

### GetUserOrGroupSiteOverrideCount

`func (o *GlobalDesktopEntitlementInfoV2) GetUserOrGroupSiteOverrideCount() int32`

GetUserOrGroupSiteOverrideCount returns the UserOrGroupSiteOverrideCount field if non-nil, zero value otherwise.

### GetUserOrGroupSiteOverrideCountOk

`func (o *GlobalDesktopEntitlementInfoV2) GetUserOrGroupSiteOverrideCountOk() (*int32, bool)`

GetUserOrGroupSiteOverrideCountOk returns a tuple with the UserOrGroupSiteOverrideCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserOrGroupSiteOverrideCount

`func (o *GlobalDesktopEntitlementInfoV2) SetUserOrGroupSiteOverrideCount(v int32)`

SetUserOrGroupSiteOverrideCount sets UserOrGroupSiteOverrideCount field to given value.

### HasUserOrGroupSiteOverrideCount

`func (o *GlobalDesktopEntitlementInfoV2) HasUserOrGroupSiteOverrideCount() bool`

HasUserOrGroupSiteOverrideCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


