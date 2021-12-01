# DesktopPoolUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | Access groups can organize the entities such as desktop pools in the organization. They can also be used for delegated administration. &lt;br&gt; This property is required for all the pools except for RDS desktop pool, which will be inherited from the corresponding Farm. | [optional] 
**AllowMultipleUserAssignments** | Pointer to **bool** | Applicable To: Dedicated manual and automated desktop pools. with manual user assignment.&lt;br&gt;Indicates whether assignment of multiple users to a single machine is allowed. This is required for Dedicated manual and automated desktop pools&lt;br&gt; | [optional] 
**AllowRdsPoolMultiSessionPerUser** | Pointer to **bool** | Applicable To: RDS desktop pools.&lt;br&gt;Indicates whether multiple sessions are allowed per user for this pool. This is required for RDS desktop pool. For other desktop pools, allow_multiple_sessions_per_user in session_settings will be applicable. | [optional] 
**AutomaticUserAssignment** | Pointer to **bool** | Applicable To: Dedicated desktop pools.&lt;br&gt;Automatic assignment of a user the first time they access the machine.&lt;br&gt; | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the desktop pool.This property will not be set if the desktop pool does not belong to a category.This property defines valid folder names with a max length of 64 characters and up to 4 subdirectory levels.The subdirectories can be specified using a backslash, e.g. (dir1\\dir2\\dir3\\dir4). Folder names can&#39;t start orend with a backslash nor can there be 2 or more backslashes together. Combinations such as(\\dir1, dir1\\dir2\\, dir1\\\\\\dir2, dir1\\\\\\\\\\dir2) are invalid. The windows reserved keywords(CON, PRN, NUL, AUX, COM1 - COM9, LPT1 - LPT9 etc.) are not allowed in subdirectory names. | [optional] 
**CloudAssigned** | **bool** | Indicates whether this desktop pool is assigned to a workspace in Horizon Cloud Services.&lt;br&gt;This can be set to true from cloud session only and only when cloud_managed is true.This can be changed to false only if there are no entitlements. | 
**CloudBrokered** | Pointer to **bool** | Applicable To: RDS Desktop Pools. &lt;br&gt;This is required for RDS Desktop Pools. | [optional] 
**CloudManaged** | **bool** | Indicates whether this desktop pool is managed by Horizon Cloud Services.This can be false only when cloud_assigned is false.&lt;br&gt;This cannot be set to true, if any of the conditions are satisfied:&lt;br&gt;1. user is provided.&lt;br&gt;2. enabled is false.&lt;br&gt;3. session_type is not DESKTOP.&lt;br&gt;4. global_entitlement is set.&lt;br&gt;5. user_assignment is DEDICATED.&lt;br&gt;6. automatic_user_assignment is false.&lt;br&gt;7. Local entitlements are configured.&lt;br&gt;8. Any of the machines in the pool have users assigned.&lt;br&gt;9. cs_restriction_tags is not set.&lt;br&gt;10. type is MANUAL.&lt;br&gt; | 
**CsRestrictionTags** | Pointer to **[]string** | List of Connection server restriction tags for which the access to the desktop pool is restricted to. If this property is not set then it indicates that desktop pool can be accessed from any connection server. | [optional] 
**CustomizationSettings** | Pointer to [**DesktopPoolCustomizationSettingsUpdateSpec**](DesktopPoolCustomizationSettingsUpdateSpec.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the desktop pool. | [optional] 
**DisplayAssignedMachineName** | **bool** | Applicable To: Dedicated desktop pools.&lt;br&gt;Indicates whether users should see the hostname of the machine assigned to them instead of display_name when they connect using Horizon Client. If no machine is assigned to the user then \&quot;display_name (No machine assigned)\&quot; will be displayed in the client. | 
**DisplayMachineAlias** | **bool** | Applicable To: Dedicated desktop pools.&lt;br&gt;Indicates whether users should see the machine alias of the machine assigned to them instead of displayName when they connect using Horizon Client. If no machine is assigned to the user then \&quot;display_name (No machine assigned)\&quot; will be displayed in the client. If both display_assigned_machine_name and this property is set to true, machine alias of the assigned machine is displayed if the user has machine alias set. Otherwise hostname will be displayed. | 
**DisplayName** | **string** | Display name of the desktop pool. | 
**DisplayProtocolSettings** | Pointer to [**DesktopPoolDisplayProtocolSettingsUpdateSpec**](DesktopPoolDisplayProtocolSettingsUpdateSpec.md) |  | [optional] 
**EnableClientRestrictions** | **bool** | Indicates whether client restrictions are to be applied to desktop pool.  | 
**EnableProvisioning** | Pointer to **bool** | Applicable To: Automated desktop pools.&lt;br&gt;Indicates whether provisioning is enabled. This is required for Automated desktop pools. | [optional] 
**Enabled** | **bool** | Indicates whether the desktop pool is enabled for brokering. | 
**Nics** | Pointer to [**[]DesktopPoolNetworkInterfaceCardSettingsUpdateSpec**](DesktopPoolNetworkInterfaceCardSettingsUpdateSpec.md) | Applicable To: Instant Clone desktop pools.&lt;br&gt;Network interface card settings for machines provisioned for this desktop pool. A NIC may appear at most once in these settings and must be present on this desktop pool&#39;s parent&#39;s snapshot or template. Not all NICs need be configured. If value is not configured than will use default settings. | [optional] 
**PatternNamingSettings** | Pointer to [**DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec**](DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec.md) |  | [optional] 
**ProvisioningSettings** | Pointer to [**DesktopPoolProvisioningSettingsUpdateSpec**](DesktopPoolProvisioningSettingsUpdateSpec.md) |  | [optional] 
**SessionSettings** | Pointer to [**DesktopPoolSessionSettingsUpdateSpec**](DesktopPoolSessionSettingsUpdateSpec.md) |  | [optional] 
**SessionType** | Pointer to **string** | Applicable To: Managed desktop pools.&lt;br&gt;Supported session types for this desktop pool. If this property is set to APPLICATION then this desktop pool can be used for application pool creation. This will beuseful when the machines in the pool support application remoting. This is required for managed desktop pools. * DESKTOP: Only desktop sessions are supported. * APPLICATION: Only application sessions are supported. * DESKTOP_AND_APPLICATION: Both desktop and application sessions are supported. | [optional] 
**ShortcutLocationsV2** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop pool. This is required if category_folder_name is set. | [optional] 
**SpecificNamingSettings** | Pointer to [**DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec**](DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec.md) |  | [optional] 
**StopProvisioningOnError** | Pointer to **bool** | Applicable To: Automated desktop pools.&lt;br&gt;Indicates whether provisioning on all machines stops on error. This is required for Automated desktop pools. | [optional] 
**StorageSettings** | Pointer to [**DesktopPoolStorageSettingsUpdateSpec**](DesktopPoolStorageSettingsUpdateSpec.md) |  | [optional] 
**TransparentPageSharingScope** | Pointer to **string** | Applicable To: Managed Manual and Automated desktop pools.&lt;br&gt;The transparent page sharing scope. This is required for Manual and Automated desktop pools. * VM: Inter-VM page sharing is not permitted. * DESKTOP_POOL: Inter-VM page sharing among VMs belonging to the same Desktop pool is permitted. * POD: Inter-VM page sharing among VMs belonging to the same Pod is permitted. * GLOBAL: Inter-VM page sharing among all VMs on the same host is permitted. | [optional] 
**ViewStorageAcceleratorSettings** | Pointer to [**DesktopPoolViewStorageAcceleratorSettingsUpdateSpec**](DesktopPoolViewStorageAcceleratorSettingsUpdateSpec.md) |  | [optional] 

## Methods

### NewDesktopPoolUpdateSpec

`func NewDesktopPoolUpdateSpec(cloudAssigned bool, cloudManaged bool, displayAssignedMachineName bool, displayMachineAlias bool, displayName string, enableClientRestrictions bool, enabled bool, ) *DesktopPoolUpdateSpec`

NewDesktopPoolUpdateSpec instantiates a new DesktopPoolUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolUpdateSpecWithDefaults

`func NewDesktopPoolUpdateSpecWithDefaults() *DesktopPoolUpdateSpec`

NewDesktopPoolUpdateSpecWithDefaults instantiates a new DesktopPoolUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *DesktopPoolUpdateSpec) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *DesktopPoolUpdateSpec) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *DesktopPoolUpdateSpec) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *DesktopPoolUpdateSpec) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAllowMultipleUserAssignments

`func (o *DesktopPoolUpdateSpec) GetAllowMultipleUserAssignments() bool`

GetAllowMultipleUserAssignments returns the AllowMultipleUserAssignments field if non-nil, zero value otherwise.

### GetAllowMultipleUserAssignmentsOk

`func (o *DesktopPoolUpdateSpec) GetAllowMultipleUserAssignmentsOk() (*bool, bool)`

GetAllowMultipleUserAssignmentsOk returns a tuple with the AllowMultipleUserAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleUserAssignments

`func (o *DesktopPoolUpdateSpec) SetAllowMultipleUserAssignments(v bool)`

SetAllowMultipleUserAssignments sets AllowMultipleUserAssignments field to given value.

### HasAllowMultipleUserAssignments

`func (o *DesktopPoolUpdateSpec) HasAllowMultipleUserAssignments() bool`

HasAllowMultipleUserAssignments returns a boolean if a field has been set.

### GetAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolUpdateSpec) GetAllowRdsPoolMultiSessionPerUser() bool`

GetAllowRdsPoolMultiSessionPerUser returns the AllowRdsPoolMultiSessionPerUser field if non-nil, zero value otherwise.

### GetAllowRdsPoolMultiSessionPerUserOk

`func (o *DesktopPoolUpdateSpec) GetAllowRdsPoolMultiSessionPerUserOk() (*bool, bool)`

GetAllowRdsPoolMultiSessionPerUserOk returns a tuple with the AllowRdsPoolMultiSessionPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolUpdateSpec) SetAllowRdsPoolMultiSessionPerUser(v bool)`

SetAllowRdsPoolMultiSessionPerUser sets AllowRdsPoolMultiSessionPerUser field to given value.

### HasAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolUpdateSpec) HasAllowRdsPoolMultiSessionPerUser() bool`

HasAllowRdsPoolMultiSessionPerUser returns a boolean if a field has been set.

### GetAutomaticUserAssignment

`func (o *DesktopPoolUpdateSpec) GetAutomaticUserAssignment() bool`

GetAutomaticUserAssignment returns the AutomaticUserAssignment field if non-nil, zero value otherwise.

### GetAutomaticUserAssignmentOk

`func (o *DesktopPoolUpdateSpec) GetAutomaticUserAssignmentOk() (*bool, bool)`

GetAutomaticUserAssignmentOk returns a tuple with the AutomaticUserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticUserAssignment

`func (o *DesktopPoolUpdateSpec) SetAutomaticUserAssignment(v bool)`

SetAutomaticUserAssignment sets AutomaticUserAssignment field to given value.

### HasAutomaticUserAssignment

`func (o *DesktopPoolUpdateSpec) HasAutomaticUserAssignment() bool`

HasAutomaticUserAssignment returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *DesktopPoolUpdateSpec) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *DesktopPoolUpdateSpec) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *DesktopPoolUpdateSpec) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *DesktopPoolUpdateSpec) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudAssigned

`func (o *DesktopPoolUpdateSpec) GetCloudAssigned() bool`

GetCloudAssigned returns the CloudAssigned field if non-nil, zero value otherwise.

### GetCloudAssignedOk

`func (o *DesktopPoolUpdateSpec) GetCloudAssignedOk() (*bool, bool)`

GetCloudAssignedOk returns a tuple with the CloudAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAssigned

`func (o *DesktopPoolUpdateSpec) SetCloudAssigned(v bool)`

SetCloudAssigned sets CloudAssigned field to given value.


### GetCloudBrokered

`func (o *DesktopPoolUpdateSpec) GetCloudBrokered() bool`

GetCloudBrokered returns the CloudBrokered field if non-nil, zero value otherwise.

### GetCloudBrokeredOk

`func (o *DesktopPoolUpdateSpec) GetCloudBrokeredOk() (*bool, bool)`

GetCloudBrokeredOk returns a tuple with the CloudBrokered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudBrokered

`func (o *DesktopPoolUpdateSpec) SetCloudBrokered(v bool)`

SetCloudBrokered sets CloudBrokered field to given value.

### HasCloudBrokered

`func (o *DesktopPoolUpdateSpec) HasCloudBrokered() bool`

HasCloudBrokered returns a boolean if a field has been set.

### GetCloudManaged

`func (o *DesktopPoolUpdateSpec) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *DesktopPoolUpdateSpec) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *DesktopPoolUpdateSpec) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.


### GetCsRestrictionTags

`func (o *DesktopPoolUpdateSpec) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *DesktopPoolUpdateSpec) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *DesktopPoolUpdateSpec) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *DesktopPoolUpdateSpec) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetCustomizationSettings

`func (o *DesktopPoolUpdateSpec) GetCustomizationSettings() DesktopPoolCustomizationSettingsUpdateSpec`

GetCustomizationSettings returns the CustomizationSettings field if non-nil, zero value otherwise.

### GetCustomizationSettingsOk

`func (o *DesktopPoolUpdateSpec) GetCustomizationSettingsOk() (*DesktopPoolCustomizationSettingsUpdateSpec, bool)`

GetCustomizationSettingsOk returns a tuple with the CustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationSettings

`func (o *DesktopPoolUpdateSpec) SetCustomizationSettings(v DesktopPoolCustomizationSettingsUpdateSpec)`

SetCustomizationSettings sets CustomizationSettings field to given value.

### HasCustomizationSettings

`func (o *DesktopPoolUpdateSpec) HasCustomizationSettings() bool`

HasCustomizationSettings returns a boolean if a field has been set.

### GetDescription

`func (o *DesktopPoolUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DesktopPoolUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DesktopPoolUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DesktopPoolUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayAssignedMachineName

`func (o *DesktopPoolUpdateSpec) GetDisplayAssignedMachineName() bool`

GetDisplayAssignedMachineName returns the DisplayAssignedMachineName field if non-nil, zero value otherwise.

### GetDisplayAssignedMachineNameOk

`func (o *DesktopPoolUpdateSpec) GetDisplayAssignedMachineNameOk() (*bool, bool)`

GetDisplayAssignedMachineNameOk returns a tuple with the DisplayAssignedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAssignedMachineName

`func (o *DesktopPoolUpdateSpec) SetDisplayAssignedMachineName(v bool)`

SetDisplayAssignedMachineName sets DisplayAssignedMachineName field to given value.


### GetDisplayMachineAlias

`func (o *DesktopPoolUpdateSpec) GetDisplayMachineAlias() bool`

GetDisplayMachineAlias returns the DisplayMachineAlias field if non-nil, zero value otherwise.

### GetDisplayMachineAliasOk

`func (o *DesktopPoolUpdateSpec) GetDisplayMachineAliasOk() (*bool, bool)`

GetDisplayMachineAliasOk returns a tuple with the DisplayMachineAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayMachineAlias

`func (o *DesktopPoolUpdateSpec) SetDisplayMachineAlias(v bool)`

SetDisplayMachineAlias sets DisplayMachineAlias field to given value.


### GetDisplayName

`func (o *DesktopPoolUpdateSpec) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DesktopPoolUpdateSpec) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DesktopPoolUpdateSpec) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDisplayProtocolSettings

`func (o *DesktopPoolUpdateSpec) GetDisplayProtocolSettings() DesktopPoolDisplayProtocolSettingsUpdateSpec`

GetDisplayProtocolSettings returns the DisplayProtocolSettings field if non-nil, zero value otherwise.

### GetDisplayProtocolSettingsOk

`func (o *DesktopPoolUpdateSpec) GetDisplayProtocolSettingsOk() (*DesktopPoolDisplayProtocolSettingsUpdateSpec, bool)`

GetDisplayProtocolSettingsOk returns a tuple with the DisplayProtocolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocolSettings

`func (o *DesktopPoolUpdateSpec) SetDisplayProtocolSettings(v DesktopPoolDisplayProtocolSettingsUpdateSpec)`

SetDisplayProtocolSettings sets DisplayProtocolSettings field to given value.

### HasDisplayProtocolSettings

`func (o *DesktopPoolUpdateSpec) HasDisplayProtocolSettings() bool`

HasDisplayProtocolSettings returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *DesktopPoolUpdateSpec) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *DesktopPoolUpdateSpec) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *DesktopPoolUpdateSpec) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.


### GetEnableProvisioning

`func (o *DesktopPoolUpdateSpec) GetEnableProvisioning() bool`

GetEnableProvisioning returns the EnableProvisioning field if non-nil, zero value otherwise.

### GetEnableProvisioningOk

`func (o *DesktopPoolUpdateSpec) GetEnableProvisioningOk() (*bool, bool)`

GetEnableProvisioningOk returns a tuple with the EnableProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProvisioning

`func (o *DesktopPoolUpdateSpec) SetEnableProvisioning(v bool)`

SetEnableProvisioning sets EnableProvisioning field to given value.

### HasEnableProvisioning

`func (o *DesktopPoolUpdateSpec) HasEnableProvisioning() bool`

HasEnableProvisioning returns a boolean if a field has been set.

### GetEnabled

`func (o *DesktopPoolUpdateSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DesktopPoolUpdateSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DesktopPoolUpdateSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetNics

`func (o *DesktopPoolUpdateSpec) GetNics() []DesktopPoolNetworkInterfaceCardSettingsUpdateSpec`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *DesktopPoolUpdateSpec) GetNicsOk() (*[]DesktopPoolNetworkInterfaceCardSettingsUpdateSpec, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *DesktopPoolUpdateSpec) SetNics(v []DesktopPoolNetworkInterfaceCardSettingsUpdateSpec)`

SetNics sets Nics field to given value.

### HasNics

`func (o *DesktopPoolUpdateSpec) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetPatternNamingSettings

`func (o *DesktopPoolUpdateSpec) GetPatternNamingSettings() DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec`

GetPatternNamingSettings returns the PatternNamingSettings field if non-nil, zero value otherwise.

### GetPatternNamingSettingsOk

`func (o *DesktopPoolUpdateSpec) GetPatternNamingSettingsOk() (*DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec, bool)`

GetPatternNamingSettingsOk returns a tuple with the PatternNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternNamingSettings

`func (o *DesktopPoolUpdateSpec) SetPatternNamingSettings(v DesktopPoolVirtualMachinePatternNamingSettingsUpdateSpec)`

SetPatternNamingSettings sets PatternNamingSettings field to given value.

### HasPatternNamingSettings

`func (o *DesktopPoolUpdateSpec) HasPatternNamingSettings() bool`

HasPatternNamingSettings returns a boolean if a field has been set.

### GetProvisioningSettings

`func (o *DesktopPoolUpdateSpec) GetProvisioningSettings() DesktopPoolProvisioningSettingsUpdateSpec`

GetProvisioningSettings returns the ProvisioningSettings field if non-nil, zero value otherwise.

### GetProvisioningSettingsOk

`func (o *DesktopPoolUpdateSpec) GetProvisioningSettingsOk() (*DesktopPoolProvisioningSettingsUpdateSpec, bool)`

GetProvisioningSettingsOk returns a tuple with the ProvisioningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSettings

`func (o *DesktopPoolUpdateSpec) SetProvisioningSettings(v DesktopPoolProvisioningSettingsUpdateSpec)`

SetProvisioningSettings sets ProvisioningSettings field to given value.

### HasProvisioningSettings

`func (o *DesktopPoolUpdateSpec) HasProvisioningSettings() bool`

HasProvisioningSettings returns a boolean if a field has been set.

### GetSessionSettings

`func (o *DesktopPoolUpdateSpec) GetSessionSettings() DesktopPoolSessionSettingsUpdateSpec`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *DesktopPoolUpdateSpec) GetSessionSettingsOk() (*DesktopPoolSessionSettingsUpdateSpec, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *DesktopPoolUpdateSpec) SetSessionSettings(v DesktopPoolSessionSettingsUpdateSpec)`

SetSessionSettings sets SessionSettings field to given value.

### HasSessionSettings

`func (o *DesktopPoolUpdateSpec) HasSessionSettings() bool`

HasSessionSettings returns a boolean if a field has been set.

### GetSessionType

`func (o *DesktopPoolUpdateSpec) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *DesktopPoolUpdateSpec) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *DesktopPoolUpdateSpec) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *DesktopPoolUpdateSpec) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetShortcutLocationsV2

`func (o *DesktopPoolUpdateSpec) GetShortcutLocationsV2() []string`

GetShortcutLocationsV2 returns the ShortcutLocationsV2 field if non-nil, zero value otherwise.

### GetShortcutLocationsV2Ok

`func (o *DesktopPoolUpdateSpec) GetShortcutLocationsV2Ok() (*[]string, bool)`

GetShortcutLocationsV2Ok returns a tuple with the ShortcutLocationsV2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocationsV2

`func (o *DesktopPoolUpdateSpec) SetShortcutLocationsV2(v []string)`

SetShortcutLocationsV2 sets ShortcutLocationsV2 field to given value.

### HasShortcutLocationsV2

`func (o *DesktopPoolUpdateSpec) HasShortcutLocationsV2() bool`

HasShortcutLocationsV2 returns a boolean if a field has been set.

### GetSpecificNamingSettings

`func (o *DesktopPoolUpdateSpec) GetSpecificNamingSettings() DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec`

GetSpecificNamingSettings returns the SpecificNamingSettings field if non-nil, zero value otherwise.

### GetSpecificNamingSettingsOk

`func (o *DesktopPoolUpdateSpec) GetSpecificNamingSettingsOk() (*DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec, bool)`

GetSpecificNamingSettingsOk returns a tuple with the SpecificNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificNamingSettings

`func (o *DesktopPoolUpdateSpec) SetSpecificNamingSettings(v DesktopPoolVirtualMachineSpecifiedNamingSettingsUpdateSpec)`

SetSpecificNamingSettings sets SpecificNamingSettings field to given value.

### HasSpecificNamingSettings

`func (o *DesktopPoolUpdateSpec) HasSpecificNamingSettings() bool`

HasSpecificNamingSettings returns a boolean if a field has been set.

### GetStopProvisioningOnError

`func (o *DesktopPoolUpdateSpec) GetStopProvisioningOnError() bool`

GetStopProvisioningOnError returns the StopProvisioningOnError field if non-nil, zero value otherwise.

### GetStopProvisioningOnErrorOk

`func (o *DesktopPoolUpdateSpec) GetStopProvisioningOnErrorOk() (*bool, bool)`

GetStopProvisioningOnErrorOk returns a tuple with the StopProvisioningOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProvisioningOnError

`func (o *DesktopPoolUpdateSpec) SetStopProvisioningOnError(v bool)`

SetStopProvisioningOnError sets StopProvisioningOnError field to given value.

### HasStopProvisioningOnError

`func (o *DesktopPoolUpdateSpec) HasStopProvisioningOnError() bool`

HasStopProvisioningOnError returns a boolean if a field has been set.

### GetStorageSettings

`func (o *DesktopPoolUpdateSpec) GetStorageSettings() DesktopPoolStorageSettingsUpdateSpec`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *DesktopPoolUpdateSpec) GetStorageSettingsOk() (*DesktopPoolStorageSettingsUpdateSpec, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *DesktopPoolUpdateSpec) SetStorageSettings(v DesktopPoolStorageSettingsUpdateSpec)`

SetStorageSettings sets StorageSettings field to given value.

### HasStorageSettings

`func (o *DesktopPoolUpdateSpec) HasStorageSettings() bool`

HasStorageSettings returns a boolean if a field has been set.

### GetTransparentPageSharingScope

`func (o *DesktopPoolUpdateSpec) GetTransparentPageSharingScope() string`

GetTransparentPageSharingScope returns the TransparentPageSharingScope field if non-nil, zero value otherwise.

### GetTransparentPageSharingScopeOk

`func (o *DesktopPoolUpdateSpec) GetTransparentPageSharingScopeOk() (*string, bool)`

GetTransparentPageSharingScopeOk returns a tuple with the TransparentPageSharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentPageSharingScope

`func (o *DesktopPoolUpdateSpec) SetTransparentPageSharingScope(v string)`

SetTransparentPageSharingScope sets TransparentPageSharingScope field to given value.

### HasTransparentPageSharingScope

`func (o *DesktopPoolUpdateSpec) HasTransparentPageSharingScope() bool`

HasTransparentPageSharingScope returns a boolean if a field has been set.

### GetViewStorageAcceleratorSettings

`func (o *DesktopPoolUpdateSpec) GetViewStorageAcceleratorSettings() DesktopPoolViewStorageAcceleratorSettingsUpdateSpec`

GetViewStorageAcceleratorSettings returns the ViewStorageAcceleratorSettings field if non-nil, zero value otherwise.

### GetViewStorageAcceleratorSettingsOk

`func (o *DesktopPoolUpdateSpec) GetViewStorageAcceleratorSettingsOk() (*DesktopPoolViewStorageAcceleratorSettingsUpdateSpec, bool)`

GetViewStorageAcceleratorSettingsOk returns a tuple with the ViewStorageAcceleratorSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStorageAcceleratorSettings

`func (o *DesktopPoolUpdateSpec) SetViewStorageAcceleratorSettings(v DesktopPoolViewStorageAcceleratorSettingsUpdateSpec)`

SetViewStorageAcceleratorSettings sets ViewStorageAcceleratorSettings field to given value.

### HasViewStorageAcceleratorSettings

`func (o *DesktopPoolUpdateSpec) HasViewStorageAcceleratorSettings() bool`

HasViewStorageAcceleratorSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


