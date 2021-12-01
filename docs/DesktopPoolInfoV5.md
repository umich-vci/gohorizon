# DesktopPoolInfoV5

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | Access groups can organize the entities (like application pools, desktop pools) in the organization. They can also be used for delegated administration.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**AllowMultipleUserAssignments** | Pointer to **bool** | Applicable To: Dedicated manual and automated desktop pools (except linked clone pools) with manual user assignment.&lt;br&gt;Whether assignment of multiple users to a single machine is allowed.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**AllowRdsPoolMultiSessionPerUser** | Pointer to **bool** | Applicable To: RDS desktop pools.&lt;br&gt;Whether multiple sessions are allowed per user for this pool.For other desktop pools, allow_multiple_sessions_per_user in sessionSettings will be applicable. | [optional] 
**AutomaticUserAssignment** | Pointer to **bool** | Applicable To: Dedicated manual and automated desktop pools.&lt;br&gt;Automatic assignment of a user the first time they access the machine.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the desktop pool.Will be unset if the desktop does not belong to a category.This property defines valid folder names with a max length of 64 characters and up to 4 subdirectory levels.The subdirectories can be specified using a backslash, e.g. (dir1\\dir2\\dir3\\dir4). Folder names can&#39;t start orend with a backslash nor can there be 2 or more backslashes together. Combinations such as(\\dir1, dir1\\dir2\\, dir1\\\\dir2, dir1\\\\\\dir2) are invalid. The windows reserved keywords(CON, PRN, NUL, AUX, COM1 - COM9, LPT1 - LPT9 etc.) are not allowed in subdirectory names. | [optional] 
**CloudAssigned** | Pointer to **bool** | Indicates whether this desktop is assigned to a workspace in Horizon Cloud Services.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**CloudBrokered** | Pointer to **bool** | Indicates whether the RDS desktop pool is cloud brokered. This property will be unset for non RDS desktop pools.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**CloudManaged** | Pointer to **bool** | Indicates whether this desktop is managed by Horizon Cloud Services.This can be false only when cloud_assigned is false.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | List of tags for which the access to the desktop pool is restricted to.No list indicates that desktop pool can be accessed from any connection server. | [optional] 
**CustomizationSettings** | Pointer to [**DesktopPoolCustomizationSettings**](DesktopPoolCustomizationSettings.md) |  | [optional] 
**DeleteInProgress** | Pointer to **bool** | Indicates whether the desktop pool is in the process of being deleted.Default value is false.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Description** | Pointer to **string** | Description of the Desktop Pool. The maximum length is 1024 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DisplayAssignedMachineName** | Pointer to **bool** | Applicable To: Dedicated manual and automated desktop pools.&lt;br&gt;Indicates whether users should see the hostname of the machine assigned to them instead of displayName when they connect using View Client. If no machine is assigned to the user then \&quot;displayName (No machine assigned)\&quot; will be displayed in the client.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**DisplayMachineAlias** | Pointer to **bool** | Decides the visibility of the machine alias to the user. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the Desktop Pool. The maximum length is 256 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DisplayProtocolSettings** | Pointer to [**DesktopPoolDisplayProtocolSettings**](DesktopPoolDisplayProtocolSettings.md) |  | [optional] 
**EnableClientRestrictions** | Pointer to **bool** | Client restrictions to be applied to the desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**EnableProvisioning** | Pointer to **bool** | Applicable To: Automated desktop pools.&lt;br&gt;Whether provisioning is enabled.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Desktop Pool is enabled for brokering.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**FarmId** | Pointer to **string** | Applicable To: RDS desktop pools.&lt;br&gt;Farm needed to create RDS desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**GlobalDesktopEntitlementId** | Pointer to **string** | Global desktop entitlement for this desktop pool. Caller should have permission to FEDERATED_LDAP_VIEW privilege for this field to be populated or to use in filter.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing Desktop Pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ImageSource** | Pointer to **string** | Applicable To: Automated desktop pools.&lt;br&gt;Source of image used in the desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * VIRTUAL_CENTER: Image was created in virtual center. * IMAGE_CATALOG: Image was created in image catalog. | [optional] 
**Name** | Pointer to **string** | Name of the Desktop Pool. The maximum length is 64 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**NamingMethod** | Pointer to **string** | Naming method for the desktop pool. * SPECIFIED: List of specified names. All provisioning is done up-front. * PATTERN: Naming pattern. | [optional] 
**Nics** | Pointer to [**[]DesktopPoolNetworkInterfaceCardSettings**](DesktopPoolNetworkInterfaceCardSettings.md) | Applicable To: Automated desktop pools.&lt;br&gt;Network interface card settings for machines provisioned for this desktop. A NIC may appear at most once in these settings and must be present on this desktop pool&#39;s parent&#39;s snapshot or template. Not all NICs need be configured. Any that are not will use default settings. | [optional] 
**PatternNamingSettings** | Pointer to [**DesktopPoolVirtualMachinePatternNamingSettings**](DesktopPoolVirtualMachinePatternNamingSettings.md) |  | [optional] 
**ProvisioningSettings** | Pointer to [**DesktopPoolProvisioningSettings**](DesktopPoolProvisioningSettings.md) |  | [optional] 
**ProvisioningStatusData** | Pointer to [**DesktopPoolProvisioningStatusData**](DesktopPoolProvisioningStatusData.md) |  | [optional] 
**SessionSettings** | Pointer to [**DesktopPoolSessionSettingsV3**](DesktopPoolSessionSettingsV3.md) |  | [optional] 
**SessionType** | Pointer to **string** | Supported session types for this desktop pool. If application sessions are selected to besupported then this desktop pool can be used for application pool creation. This will beuseful when the machines in the pool support application remoting. Default value of DESKTOP.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * DESKTOP: Only desktop sessions are supported. * APPLICATION: Only application sessions are supported. * DESKTOP_AND_APPLICATION: Both desktop and application sessions are supported. | [optional] 
**ShortcutLocationsV2** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop pool.  | [optional] 
**Source** | Pointer to **string** | Source of the Machines in this Desktop Pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * INSTANT_CLONE: The Desktop Pool uses instant clone technology for provisioning the machines. Applicable for AUTOMATED type desktop pools. * LINKED_CLONE: The Desktop Pool uses linked clone technology for provisioning the machines. Applicable for AUTOMATED type desktop pools. * VIRTUAL_CENTER: The Desktop Pool uses Virtual Center as source for provisioning the machines. Applicable for AUTOMATED and MANUAL type desktop pools. * RDS: The Desktop Pool is backed by Farm. The Farm used in this Desktop Pool can be of any Source. * UNMANAGED: The Desktop Pool holds the non-vCenter source machines that includes physical computers, blade PCs and non-vCenter servers. Applicable for MANUAL type desktop pools. | [optional] 
**SpecificNamingSettings** | Pointer to [**DesktopPoolVirtualMachineSpecifiedNamingSettings**](DesktopPoolVirtualMachineSpecifiedNamingSettings.md) |  | [optional] 
**StopProvisioningOnError** | Pointer to **bool** | Applicable To: Automated desktop pools.&lt;br&gt;Whether provisioning on all machines stops on error. | [optional] 
**StorageSettings** | Pointer to [**DesktopPoolStorageSettings**](DesktopPoolStorageSettings.md) |  | [optional] 
**TransparentPageSharingScope** | Pointer to **string** | Applicable To: Manual and automated desktop pools.&lt;br&gt;The transparent page sharing scope. * VM: Inter-VM page sharing is not permitted. * DESKTOP_POOL: Inter-VM page sharing among VMs belonging to the same Desktop pool is permitted. * POD: Inter-VM page sharing among VMs belonging to the same Pod is permitted. * GLOBAL: Inter-VM page sharing among all VMs on the same host is permitted. | [optional] 
**Type** | Pointer to **string** | Type of the desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * AUTOMATED: Automated Desktop Pool. * MANUAL: Manual Desktop Pool. * RDS: RDS Desktop Pool. | [optional] 
**UserAssignment** | Pointer to **string** | Applicable To: Manual and automated desktop pools.&lt;br&gt;User assignment scheme.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * DEDICATED: With dedicated assignment, a user returns to the same machine at each session. * FLOATING: With floating assignment, a user may return to one of the available machines for the next session. | [optional] 
**UserGroupCount** | Pointer to **int32** | Count of user or group entitlements present for the desktop pool. | [optional] 
**VcenterId** | Pointer to **string** | Applicable To: Manual and automated desktop pools.&lt;br&gt;Virtual Center that manages the machines of the desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ViewStorageAcceleratorSettings** | Pointer to [**DesktopPoolViewStorageAcceleratorSettings**](DesktopPoolViewStorageAcceleratorSettings.md) |  | [optional] 

## Methods

### NewDesktopPoolInfoV5

`func NewDesktopPoolInfoV5() *DesktopPoolInfoV5`

NewDesktopPoolInfoV5 instantiates a new DesktopPoolInfoV5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolInfoV5WithDefaults

`func NewDesktopPoolInfoV5WithDefaults() *DesktopPoolInfoV5`

NewDesktopPoolInfoV5WithDefaults instantiates a new DesktopPoolInfoV5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *DesktopPoolInfoV5) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *DesktopPoolInfoV5) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *DesktopPoolInfoV5) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *DesktopPoolInfoV5) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAllowMultipleUserAssignments

`func (o *DesktopPoolInfoV5) GetAllowMultipleUserAssignments() bool`

GetAllowMultipleUserAssignments returns the AllowMultipleUserAssignments field if non-nil, zero value otherwise.

### GetAllowMultipleUserAssignmentsOk

`func (o *DesktopPoolInfoV5) GetAllowMultipleUserAssignmentsOk() (*bool, bool)`

GetAllowMultipleUserAssignmentsOk returns a tuple with the AllowMultipleUserAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleUserAssignments

`func (o *DesktopPoolInfoV5) SetAllowMultipleUserAssignments(v bool)`

SetAllowMultipleUserAssignments sets AllowMultipleUserAssignments field to given value.

### HasAllowMultipleUserAssignments

`func (o *DesktopPoolInfoV5) HasAllowMultipleUserAssignments() bool`

HasAllowMultipleUserAssignments returns a boolean if a field has been set.

### GetAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolInfoV5) GetAllowRdsPoolMultiSessionPerUser() bool`

GetAllowRdsPoolMultiSessionPerUser returns the AllowRdsPoolMultiSessionPerUser field if non-nil, zero value otherwise.

### GetAllowRdsPoolMultiSessionPerUserOk

`func (o *DesktopPoolInfoV5) GetAllowRdsPoolMultiSessionPerUserOk() (*bool, bool)`

GetAllowRdsPoolMultiSessionPerUserOk returns a tuple with the AllowRdsPoolMultiSessionPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolInfoV5) SetAllowRdsPoolMultiSessionPerUser(v bool)`

SetAllowRdsPoolMultiSessionPerUser sets AllowRdsPoolMultiSessionPerUser field to given value.

### HasAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolInfoV5) HasAllowRdsPoolMultiSessionPerUser() bool`

HasAllowRdsPoolMultiSessionPerUser returns a boolean if a field has been set.

### GetAutomaticUserAssignment

`func (o *DesktopPoolInfoV5) GetAutomaticUserAssignment() bool`

GetAutomaticUserAssignment returns the AutomaticUserAssignment field if non-nil, zero value otherwise.

### GetAutomaticUserAssignmentOk

`func (o *DesktopPoolInfoV5) GetAutomaticUserAssignmentOk() (*bool, bool)`

GetAutomaticUserAssignmentOk returns a tuple with the AutomaticUserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticUserAssignment

`func (o *DesktopPoolInfoV5) SetAutomaticUserAssignment(v bool)`

SetAutomaticUserAssignment sets AutomaticUserAssignment field to given value.

### HasAutomaticUserAssignment

`func (o *DesktopPoolInfoV5) HasAutomaticUserAssignment() bool`

HasAutomaticUserAssignment returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *DesktopPoolInfoV5) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *DesktopPoolInfoV5) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *DesktopPoolInfoV5) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *DesktopPoolInfoV5) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudAssigned

`func (o *DesktopPoolInfoV5) GetCloudAssigned() bool`

GetCloudAssigned returns the CloudAssigned field if non-nil, zero value otherwise.

### GetCloudAssignedOk

`func (o *DesktopPoolInfoV5) GetCloudAssignedOk() (*bool, bool)`

GetCloudAssignedOk returns a tuple with the CloudAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAssigned

`func (o *DesktopPoolInfoV5) SetCloudAssigned(v bool)`

SetCloudAssigned sets CloudAssigned field to given value.

### HasCloudAssigned

`func (o *DesktopPoolInfoV5) HasCloudAssigned() bool`

HasCloudAssigned returns a boolean if a field has been set.

### GetCloudBrokered

`func (o *DesktopPoolInfoV5) GetCloudBrokered() bool`

GetCloudBrokered returns the CloudBrokered field if non-nil, zero value otherwise.

### GetCloudBrokeredOk

`func (o *DesktopPoolInfoV5) GetCloudBrokeredOk() (*bool, bool)`

GetCloudBrokeredOk returns a tuple with the CloudBrokered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudBrokered

`func (o *DesktopPoolInfoV5) SetCloudBrokered(v bool)`

SetCloudBrokered sets CloudBrokered field to given value.

### HasCloudBrokered

`func (o *DesktopPoolInfoV5) HasCloudBrokered() bool`

HasCloudBrokered returns a boolean if a field has been set.

### GetCloudManaged

`func (o *DesktopPoolInfoV5) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *DesktopPoolInfoV5) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *DesktopPoolInfoV5) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *DesktopPoolInfoV5) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *DesktopPoolInfoV5) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *DesktopPoolInfoV5) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *DesktopPoolInfoV5) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *DesktopPoolInfoV5) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetCustomizationSettings

`func (o *DesktopPoolInfoV5) GetCustomizationSettings() DesktopPoolCustomizationSettings`

GetCustomizationSettings returns the CustomizationSettings field if non-nil, zero value otherwise.

### GetCustomizationSettingsOk

`func (o *DesktopPoolInfoV5) GetCustomizationSettingsOk() (*DesktopPoolCustomizationSettings, bool)`

GetCustomizationSettingsOk returns a tuple with the CustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationSettings

`func (o *DesktopPoolInfoV5) SetCustomizationSettings(v DesktopPoolCustomizationSettings)`

SetCustomizationSettings sets CustomizationSettings field to given value.

### HasCustomizationSettings

`func (o *DesktopPoolInfoV5) HasCustomizationSettings() bool`

HasCustomizationSettings returns a boolean if a field has been set.

### GetDeleteInProgress

`func (o *DesktopPoolInfoV5) GetDeleteInProgress() bool`

GetDeleteInProgress returns the DeleteInProgress field if non-nil, zero value otherwise.

### GetDeleteInProgressOk

`func (o *DesktopPoolInfoV5) GetDeleteInProgressOk() (*bool, bool)`

GetDeleteInProgressOk returns a tuple with the DeleteInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInProgress

`func (o *DesktopPoolInfoV5) SetDeleteInProgress(v bool)`

SetDeleteInProgress sets DeleteInProgress field to given value.

### HasDeleteInProgress

`func (o *DesktopPoolInfoV5) HasDeleteInProgress() bool`

HasDeleteInProgress returns a boolean if a field has been set.

### GetDescription

`func (o *DesktopPoolInfoV5) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DesktopPoolInfoV5) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DesktopPoolInfoV5) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DesktopPoolInfoV5) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayAssignedMachineName

`func (o *DesktopPoolInfoV5) GetDisplayAssignedMachineName() bool`

GetDisplayAssignedMachineName returns the DisplayAssignedMachineName field if non-nil, zero value otherwise.

### GetDisplayAssignedMachineNameOk

`func (o *DesktopPoolInfoV5) GetDisplayAssignedMachineNameOk() (*bool, bool)`

GetDisplayAssignedMachineNameOk returns a tuple with the DisplayAssignedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAssignedMachineName

`func (o *DesktopPoolInfoV5) SetDisplayAssignedMachineName(v bool)`

SetDisplayAssignedMachineName sets DisplayAssignedMachineName field to given value.

### HasDisplayAssignedMachineName

`func (o *DesktopPoolInfoV5) HasDisplayAssignedMachineName() bool`

HasDisplayAssignedMachineName returns a boolean if a field has been set.

### GetDisplayMachineAlias

`func (o *DesktopPoolInfoV5) GetDisplayMachineAlias() bool`

GetDisplayMachineAlias returns the DisplayMachineAlias field if non-nil, zero value otherwise.

### GetDisplayMachineAliasOk

`func (o *DesktopPoolInfoV5) GetDisplayMachineAliasOk() (*bool, bool)`

GetDisplayMachineAliasOk returns a tuple with the DisplayMachineAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMachineAlias

`func (o *DesktopPoolInfoV5) SetDisplayMachineAlias(v bool)`

SetDisplayMachineAlias sets DisplayMachineAlias field to given value.

### HasDisplayMachineAlias

`func (o *DesktopPoolInfoV5) HasDisplayMachineAlias() bool`

HasDisplayMachineAlias returns a boolean if a field has been set.

### GetDisplayName

`func (o *DesktopPoolInfoV5) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DesktopPoolInfoV5) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DesktopPoolInfoV5) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DesktopPoolInfoV5) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayProtocolSettings

`func (o *DesktopPoolInfoV5) GetDisplayProtocolSettings() DesktopPoolDisplayProtocolSettings`

GetDisplayProtocolSettings returns the DisplayProtocolSettings field if non-nil, zero value otherwise.

### GetDisplayProtocolSettingsOk

`func (o *DesktopPoolInfoV5) GetDisplayProtocolSettingsOk() (*DesktopPoolDisplayProtocolSettings, bool)`

GetDisplayProtocolSettingsOk returns a tuple with the DisplayProtocolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocolSettings

`func (o *DesktopPoolInfoV5) SetDisplayProtocolSettings(v DesktopPoolDisplayProtocolSettings)`

SetDisplayProtocolSettings sets DisplayProtocolSettings field to given value.

### HasDisplayProtocolSettings

`func (o *DesktopPoolInfoV5) HasDisplayProtocolSettings() bool`

HasDisplayProtocolSettings returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *DesktopPoolInfoV5) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *DesktopPoolInfoV5) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *DesktopPoolInfoV5) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *DesktopPoolInfoV5) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnableProvisioning

`func (o *DesktopPoolInfoV5) GetEnableProvisioning() bool`

GetEnableProvisioning returns the EnableProvisioning field if non-nil, zero value otherwise.

### GetEnableProvisioningOk

`func (o *DesktopPoolInfoV5) GetEnableProvisioningOk() (*bool, bool)`

GetEnableProvisioningOk returns a tuple with the EnableProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProvisioning

`func (o *DesktopPoolInfoV5) SetEnableProvisioning(v bool)`

SetEnableProvisioning sets EnableProvisioning field to given value.

### HasEnableProvisioning

`func (o *DesktopPoolInfoV5) HasEnableProvisioning() bool`

HasEnableProvisioning returns a boolean if a field has been set.

### GetEnabled

`func (o *DesktopPoolInfoV5) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DesktopPoolInfoV5) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DesktopPoolInfoV5) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DesktopPoolInfoV5) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFarmId

`func (o *DesktopPoolInfoV5) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *DesktopPoolInfoV5) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *DesktopPoolInfoV5) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *DesktopPoolInfoV5) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetGlobalDesktopEntitlementId

`func (o *DesktopPoolInfoV5) GetGlobalDesktopEntitlementId() string`

GetGlobalDesktopEntitlementId returns the GlobalDesktopEntitlementId field if non-nil, zero value otherwise.

### GetGlobalDesktopEntitlementIdOk

`func (o *DesktopPoolInfoV5) GetGlobalDesktopEntitlementIdOk() (*string, bool)`

GetGlobalDesktopEntitlementIdOk returns a tuple with the GlobalDesktopEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDesktopEntitlementId

`func (o *DesktopPoolInfoV5) SetGlobalDesktopEntitlementId(v string)`

SetGlobalDesktopEntitlementId sets GlobalDesktopEntitlementId field to given value.

### HasGlobalDesktopEntitlementId

`func (o *DesktopPoolInfoV5) HasGlobalDesktopEntitlementId() bool`

HasGlobalDesktopEntitlementId returns a boolean if a field has been set.

### GetId

`func (o *DesktopPoolInfoV5) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DesktopPoolInfoV5) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DesktopPoolInfoV5) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DesktopPoolInfoV5) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageSource

`func (o *DesktopPoolInfoV5) GetImageSource() string`

GetImageSource returns the ImageSource field if non-nil, zero value otherwise.

### GetImageSourceOk

`func (o *DesktopPoolInfoV5) GetImageSourceOk() (*string, bool)`

GetImageSourceOk returns a tuple with the ImageSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSource

`func (o *DesktopPoolInfoV5) SetImageSource(v string)`

SetImageSource sets ImageSource field to given value.

### HasImageSource

`func (o *DesktopPoolInfoV5) HasImageSource() bool`

HasImageSource returns a boolean if a field has been set.

### GetName

`func (o *DesktopPoolInfoV5) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DesktopPoolInfoV5) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DesktopPoolInfoV5) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DesktopPoolInfoV5) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamingMethod

`func (o *DesktopPoolInfoV5) GetNamingMethod() string`

GetNamingMethod returns the NamingMethod field if non-nil, zero value otherwise.

### GetNamingMethodOk

`func (o *DesktopPoolInfoV5) GetNamingMethodOk() (*string, bool)`

GetNamingMethodOk returns a tuple with the NamingMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingMethod

`func (o *DesktopPoolInfoV5) SetNamingMethod(v string)`

SetNamingMethod sets NamingMethod field to given value.

### HasNamingMethod

`func (o *DesktopPoolInfoV5) HasNamingMethod() bool`

HasNamingMethod returns a boolean if a field has been set.

### GetNics

`func (o *DesktopPoolInfoV5) GetNics() []DesktopPoolNetworkInterfaceCardSettings`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *DesktopPoolInfoV5) GetNicsOk() (*[]DesktopPoolNetworkInterfaceCardSettings, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *DesktopPoolInfoV5) SetNics(v []DesktopPoolNetworkInterfaceCardSettings)`

SetNics sets Nics field to given value.

### HasNics

`func (o *DesktopPoolInfoV5) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetPatternNamingSettings

`func (o *DesktopPoolInfoV5) GetPatternNamingSettings() DesktopPoolVirtualMachinePatternNamingSettings`

GetPatternNamingSettings returns the PatternNamingSettings field if non-nil, zero value otherwise.

### GetPatternNamingSettingsOk

`func (o *DesktopPoolInfoV5) GetPatternNamingSettingsOk() (*DesktopPoolVirtualMachinePatternNamingSettings, bool)`

GetPatternNamingSettingsOk returns a tuple with the PatternNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternNamingSettings

`func (o *DesktopPoolInfoV5) SetPatternNamingSettings(v DesktopPoolVirtualMachinePatternNamingSettings)`

SetPatternNamingSettings sets PatternNamingSettings field to given value.

### HasPatternNamingSettings

`func (o *DesktopPoolInfoV5) HasPatternNamingSettings() bool`

HasPatternNamingSettings returns a boolean if a field has been set.

### GetProvisioningSettings

`func (o *DesktopPoolInfoV5) GetProvisioningSettings() DesktopPoolProvisioningSettings`

GetProvisioningSettings returns the ProvisioningSettings field if non-nil, zero value otherwise.

### GetProvisioningSettingsOk

`func (o *DesktopPoolInfoV5) GetProvisioningSettingsOk() (*DesktopPoolProvisioningSettings, bool)`

GetProvisioningSettingsOk returns a tuple with the ProvisioningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSettings

`func (o *DesktopPoolInfoV5) SetProvisioningSettings(v DesktopPoolProvisioningSettings)`

SetProvisioningSettings sets ProvisioningSettings field to given value.

### HasProvisioningSettings

`func (o *DesktopPoolInfoV5) HasProvisioningSettings() bool`

HasProvisioningSettings returns a boolean if a field has been set.

### GetProvisioningStatusData

`func (o *DesktopPoolInfoV5) GetProvisioningStatusData() DesktopPoolProvisioningStatusData`

GetProvisioningStatusData returns the ProvisioningStatusData field if non-nil, zero value otherwise.

### GetProvisioningStatusDataOk

`func (o *DesktopPoolInfoV5) GetProvisioningStatusDataOk() (*DesktopPoolProvisioningStatusData, bool)`

GetProvisioningStatusDataOk returns a tuple with the ProvisioningStatusData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatusData

`func (o *DesktopPoolInfoV5) SetProvisioningStatusData(v DesktopPoolProvisioningStatusData)`

SetProvisioningStatusData sets ProvisioningStatusData field to given value.

### HasProvisioningStatusData

`func (o *DesktopPoolInfoV5) HasProvisioningStatusData() bool`

HasProvisioningStatusData returns a boolean if a field has been set.

### GetSessionSettings

`func (o *DesktopPoolInfoV5) GetSessionSettings() DesktopPoolSessionSettingsV3`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *DesktopPoolInfoV5) GetSessionSettingsOk() (*DesktopPoolSessionSettingsV3, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *DesktopPoolInfoV5) SetSessionSettings(v DesktopPoolSessionSettingsV3)`

SetSessionSettings sets SessionSettings field to given value.

### HasSessionSettings

`func (o *DesktopPoolInfoV5) HasSessionSettings() bool`

HasSessionSettings returns a boolean if a field has been set.

### GetSessionType

`func (o *DesktopPoolInfoV5) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *DesktopPoolInfoV5) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *DesktopPoolInfoV5) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *DesktopPoolInfoV5) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetShortcutLocationsV2

`func (o *DesktopPoolInfoV5) GetShortcutLocationsV2() []string`

GetShortcutLocationsV2 returns the ShortcutLocationsV2 field if non-nil, zero value otherwise.

### GetShortcutLocationsV2Ok

`func (o *DesktopPoolInfoV5) GetShortcutLocationsV2Ok() (*[]string, bool)`

GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocationsV2

`func (o *DesktopPoolInfoV5) SetShortcutLocationsV2(v []string)`

SetShortcutLocationsV2 sets ShortcutLocationsV2 field to given value.

### HasShortcutLocationsV2

`func (o *DesktopPoolInfoV5) HasShortcutLocationsV2() bool`

HasShortcutLocationsV2 returns a boolean if a field has been set.

### GetSource

`func (o *DesktopPoolInfoV5) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DesktopPoolInfoV5) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DesktopPoolInfoV5) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DesktopPoolInfoV5) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpecificNamingSettings

`func (o *DesktopPoolInfoV5) GetSpecificNamingSettings() DesktopPoolVirtualMachineSpecifiedNamingSettings`

GetSpecificNamingSettings returns the SpecificNamingSettings field if non-nil, zero value otherwise.

### GetSpecificNamingSettingsOk

`func (o *DesktopPoolInfoV5) GetSpecificNamingSettingsOk() (*DesktopPoolVirtualMachineSpecifiedNamingSettings, bool)`

GetSpecificNamingSettingsOk returns a tuple with the SpecificNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificNamingSettings

`func (o *DesktopPoolInfoV5) SetSpecificNamingSettings(v DesktopPoolVirtualMachineSpecifiedNamingSettings)`

SetSpecificNamingSettings sets SpecificNamingSettings field to given value.

### HasSpecificNamingSettings

`func (o *DesktopPoolInfoV5) HasSpecificNamingSettings() bool`

HasSpecificNamingSettings returns a boolean if a field has been set.

### GetStopProvisioningOnError

`func (o *DesktopPoolInfoV5) GetStopProvisioningOnError() bool`

GetStopProvisioningOnError returns the StopProvisioningOnError field if non-nil, zero value otherwise.

### GetStopProvisioningOnErrorOk

`func (o *DesktopPoolInfoV5) GetStopProvisioningOnErrorOk() (*bool, bool)`

GetStopProvisioningOnErrorOk returns a tuple with the StopProvisioningOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProvisioningOnError

`func (o *DesktopPoolInfoV5) SetStopProvisioningOnError(v bool)`

SetStopProvisioningOnError sets StopProvisioningOnError field to given value.

### HasStopProvisioningOnError

`func (o *DesktopPoolInfoV5) HasStopProvisioningOnError() bool`

HasStopProvisioningOnError returns a boolean if a field has been set.

### GetStorageSettings

`func (o *DesktopPoolInfoV5) GetStorageSettings() DesktopPoolStorageSettings`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *DesktopPoolInfoV5) GetStorageSettingsOk() (*DesktopPoolStorageSettings, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *DesktopPoolInfoV5) SetStorageSettings(v DesktopPoolStorageSettings)`

SetStorageSettings sets StorageSettings field to given value.

### HasStorageSettings

`func (o *DesktopPoolInfoV5) HasStorageSettings() bool`

HasStorageSettings returns a boolean if a field has been set.

### GetTransparentPageSharingScope

`func (o *DesktopPoolInfoV5) GetTransparentPageSharingScope() string`

GetTransparentPageSharingScope returns the TransparentPageSharingScope field if non-nil, zero value otherwise.

### GetTransparentPageSharingScopeOk

`func (o *DesktopPoolInfoV5) GetTransparentPageSharingScopeOk() (*string, bool)`

GetTransparentPageSharingScopeOk returns a tuple with the TransparentPageSharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentPageSharingScope

`func (o *DesktopPoolInfoV5) SetTransparentPageSharingScope(v string)`

SetTransparentPageSharingScope sets TransparentPageSharingScope field to given value.

### HasTransparentPageSharingScope

`func (o *DesktopPoolInfoV5) HasTransparentPageSharingScope() bool`

HasTransparentPageSharingScope returns a boolean if a field has been set.

### GetType

`func (o *DesktopPoolInfoV5) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DesktopPoolInfoV5) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DesktopPoolInfoV5) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DesktopPoolInfoV5) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserAssignment

`func (o *DesktopPoolInfoV5) GetUserAssignment() string`

GetUserAssignment returns the UserAssignment field if non-nil, zero value otherwise.

### GetUserAssignmentOk

`func (o *DesktopPoolInfoV5) GetUserAssignmentOk() (*string, bool)`

GetUserAssignmentOk returns a tuple with the UserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignment

`func (o *DesktopPoolInfoV5) SetUserAssignment(v string)`

SetUserAssignment sets UserAssignment field to given value.

### HasUserAssignment

`func (o *DesktopPoolInfoV5) HasUserAssignment() bool`

HasUserAssignment returns a boolean if a field has been set.

### GetUserGroupCount

`func (o *DesktopPoolInfoV5) GetUserGroupCount() int32`

GetUserGroupCount returns the UserGroupCount field if non-nil, zero value otherwise.

### GetUserGroupCountOk

`func (o *DesktopPoolInfoV5) GetUserGroupCountOk() (*int32, bool)`

GetUserGroupCountOk returns a tuple with the UserGroupCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroupCount

`func (o *DesktopPoolInfoV5) SetUserGroupCount(v int32)`

SetUserGroupCount sets UserGroupCount field to given value.

### HasUserGroupCount

`func (o *DesktopPoolInfoV5) HasUserGroupCount() bool`

HasUserGroupCount returns a boolean if a field has been set.

### GetVcenterId

`func (o *DesktopPoolInfoV5) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *DesktopPoolInfoV5) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *DesktopPoolInfoV5) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *DesktopPoolInfoV5) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.

### GetViewStorageAcceleratorSettings

`func (o *DesktopPoolInfoV5) GetViewStorageAcceleratorSettings() DesktopPoolViewStorageAcceleratorSettings`

GetViewStorageAcceleratorSettings returns the ViewStorageAcceleratorSettings field if non-nil, zero value otherwise.

### GetViewStorageAcceleratorSettingsOk

`func (o *DesktopPoolInfoV5) GetViewStorageAcceleratorSettingsOk() (*DesktopPoolViewStorageAcceleratorSettings, bool)`

GetViewStorageAcceleratorSettingsOk returns a tuple with the ViewStorageAcceleratorSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStorageAcceleratorSettings

`func (o *DesktopPoolInfoV5) SetViewStorageAcceleratorSettings(v DesktopPoolViewStorageAcceleratorSettings)`

SetViewStorageAcceleratorSettings sets ViewStorageAcceleratorSettings field to given value.

### HasViewStorageAcceleratorSettings

`func (o *DesktopPoolInfoV5) HasViewStorageAcceleratorSettings() bool`

HasViewStorageAcceleratorSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


