# DesktopPoolInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | Access groups can organize the entities (like application pools, desktop pools) in the organization. They can also be used for delegated administration.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**AllowMultipleUserAssignments** | Pointer to **bool** | Applicable To: Dedicated manual and automated desktop pools (except linked clone pools) with manual user assignment.&lt;br&gt;Whether assignment of multiple users to a single machine is allowed.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**AllowRdsPoolMultiSessionPerUser** | Pointer to **bool** | Applicable To: RDS desktop pools.&lt;br&gt;Whether multiple sessions are allowed per user for this pool.For other desktop pools, allow_multiple_sessions_per_user in sessionSettings will be applicable. | [optional] 
**AutomaticUserAssignment** | Pointer to **bool** | Applicable To: Dedicated manual and automated desktop pools.&lt;br&gt;Automatic assignment of a user the first time they access the machine.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the desktop pool.Will be unset if the desktop does not belong to a category.This property defines valid folder names with a max length of 64 characters and up to 4 subdirectory levels.The subdirectories can be specified using a backslash, e.g. (dir1\\dir2\\dir3\\dir4). Folder names can&#39;t start orend with a backslash nor can there be 2 or more backslashes together. Combinations such as(\\dir1, dir1\\dir2\\, dir1\\\\dir2, dir1\\\\\\dir2) are invalid. The windows reserved keywords(CON, PRN, NUL, AUX, COM1 - COM9, LPT1 - LPT9 etc.) are not allowed in subdirectory names. | [optional] 
**CloudAssigned** | Pointer to **bool** | Indicates whether this desktop is assigned to a workspace in Horizon Cloud Services.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**CloudManaged** | Pointer to **bool** | Indicates whether this desktop is managed by Horizon Cloud Services.This can be false only when cloud_assigned is false.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | List of tags for which the access to the desktop pool is restricted to.No list indicates that desktop pool can be accessed from any connection server. | [optional] 
**CustomizationSettings** | Pointer to [**DesktopPoolCustomizationSettings**](DesktopPoolCustomizationSettings.md) |  | [optional] 
**DeleteInProgress** | Pointer to **bool** | Indicates whether the desktop pool is in the process of being deleted.Default value is false.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Description** | Pointer to **string** | Description of the Desktop Pool. The maximum length is 1024 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DisplayAssignedMachineName** | Pointer to **bool** | Applicable To: Dedicated manual and automated desktop pools.&lt;br&gt;Indicates whether users should see the hostname of the machine assigned to them instead of displayName when they connect using View Client. If no machine is assigned to the user then \&quot;displayName (No machine assigned)\&quot; will be displayed in the client.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**DisplayName** | Pointer to **string** | Display name of the Desktop Pool. The maximum length is 256 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DisplayProtocolSettings** | Pointer to [**DesktopPoolDisplayProtocolSettings**](DesktopPoolDisplayProtocolSettings.md) |  | [optional] 
**EnableClientRestrictions** | Pointer to **bool** | Client restrictions to be applied to the desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**EnableProvisioning** | Pointer to **bool** | Applicable To: Automated desktop pools.&lt;br&gt;Whether provisioning is enabled.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the Desktop Pool is enabled for brokering.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**FarmId** | Pointer to **string** | Applicable To: RDS desktop pools.&lt;br&gt;Farm needed to create RDS desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing Desktop Pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ImageSource** | Pointer to **string** | Applicable To: Automated desktop pools.&lt;br&gt;Source of image used in the desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * VIRTUAL_CENTER: Image was created in virtual center. * IMAGE_CATALOG: Image was created in image catalog. | [optional] 
**Name** | Pointer to **string** | Name of the Desktop Pool. The maximum length is 64 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**Nics** | Pointer to [**[]DesktopPoolNetworkInterfaceCardSettings**](DesktopPoolNetworkInterfaceCardSettings.md) | Applicable To: Automated desktop pools.&lt;br&gt;Network interface card settings for machines provisioned for this desktop. A NIC may appear at most once in these settings and must be present on this desktop pool&#39;s parent&#39;s snapshot or template. Not all NICs need be configured. Any that are not will use default settings. | [optional] 
**PatternNamingSettings** | Pointer to [**DesktopPoolVirtualMachinePatternNamingSettings**](DesktopPoolVirtualMachinePatternNamingSettings.md) |  | [optional] 
**ProvisioningSettings** | Pointer to [**DesktopPoolProvisioningSettings**](DesktopPoolProvisioningSettings.md) |  | [optional] 
**ProvisioningStatusData** | Pointer to [**DesktopPoolProvisioningStatusData**](DesktopPoolProvisioningStatusData.md) |  | [optional] 
**SessionSettings** | Pointer to [**DesktopPoolSessionSettingsV2**](DesktopPoolSessionSettingsV2.md) |  | [optional] 
**SessionType** | Pointer to **string** | Supported session types for this desktop pool. If application sessions are selected to besupported then this desktop pool can be used for application pool creation. This will beuseful when the machines in the pool support application remoting. Default value of DESKTOP.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * DESKTOP: Only desktop sessions are supported. * APPLICATION: Only application sessions are supported. * DESKTOP_AND_APPLICATION: Both desktop and application sessions are supported. | [optional] 
**ShortcutLocations** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop pool.The value will be present if categoryFolderName is set. | [optional] 
**Source** | Pointer to **string** | Source of the Machines in this Desktop Pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * INSTANT_CLONE: The Desktop Pool uses instant clone technology for provisioning the machines. Applicable for AUTOMATED type desktop pools. * LINKED_CLONE: The Desktop Pool uses linked clone technology for provisioning the machines. Applicable for AUTOMATED type desktop pools. * VIRTUAL_CENTER: The Desktop Pool uses Virtual Center as source for provisioning the machines. Applicable for AUTOMATED and MANUAL type desktop pools. * RDS: The Desktop Pool is backed by Farm. The Farm used in this Desktop Pool can be of any Source. * UNMANAGED: The Desktop Pool holds the non-vCenter source machines that includes physical computers, blade PCs and non-vCenter servers. Applicable for MANUAL type desktop pools. | [optional] 
**SpecificNamingSettings** | Pointer to [**DesktopPoolVirtualMachineSpecifiedNamingSettings**](DesktopPoolVirtualMachineSpecifiedNamingSettings.md) |  | [optional] 
**StopProvisioningOnError** | Pointer to **bool** | Applicable To: Automated desktop pools.&lt;br&gt;Whether provisioning on all machines stops on error. | [optional] 
**StorageSettings** | Pointer to [**DesktopPoolStorageSettings**](DesktopPoolStorageSettings.md) |  | [optional] 
**TransparentPageSharingScope** | Pointer to **string** | Applicable To: Manual and automated desktop pools.&lt;br&gt;The transparent page sharing scope. * VM: Inter-VM page sharing is not permitted. * DESKTOP_POOL: Inter-VM page sharing among VMs belonging to the same Desktop pool is permitted. * POD: Inter-VM page sharing among VMs belonging to the same Pod is permitted. * GLOBAL: Inter-VM page sharing among all VMs on the same host is permitted. | [optional] 
**Type** | Pointer to **string** | Type of the desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * AUTOMATED: Automated Desktop Pool. * MANUAL: Manual Desktop Pool. * RDS: RDS Desktop Pool. | [optional] 
**UserAssignment** | Pointer to **string** | Applicable To: Manual and automated desktop pools.&lt;br&gt;User assignment scheme.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * DEDICATED: With dedicated assignment, a user returns to the same machine at each session. * FLOATING: With floating assignment, a user may return to one of the available machines for the next session. | [optional] 
**VcenterId** | Pointer to **string** | Applicable To: Manual and automated desktop pools.&lt;br&gt;Virtual Center that manages the machines of the desktop pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ViewStorageAcceleratorSettings** | Pointer to [**DesktopPoolViewStorageAcceleratorSettings**](DesktopPoolViewStorageAcceleratorSettings.md) |  | [optional] 

## Methods

### NewDesktopPoolInfoV2

`func NewDesktopPoolInfoV2() *DesktopPoolInfoV2`

NewDesktopPoolInfoV2 instantiates a new DesktopPoolInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolInfoV2WithDefaults

`func NewDesktopPoolInfoV2WithDefaults() *DesktopPoolInfoV2`

NewDesktopPoolInfoV2WithDefaults instantiates a new DesktopPoolInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *DesktopPoolInfoV2) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *DesktopPoolInfoV2) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *DesktopPoolInfoV2) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *DesktopPoolInfoV2) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAllowMultipleUserAssignments

`func (o *DesktopPoolInfoV2) GetAllowMultipleUserAssignments() bool`

GetAllowMultipleUserAssignments returns the AllowMultipleUserAssignments field if non-nil, zero value otherwise.

### GetAllowMultipleUserAssignmentsOk

`func (o *DesktopPoolInfoV2) GetAllowMultipleUserAssignmentsOk() (*bool, bool)`

GetAllowMultipleUserAssignmentsOk returns a tuple with the AllowMultipleUserAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultipleUserAssignments

`func (o *DesktopPoolInfoV2) SetAllowMultipleUserAssignments(v bool)`

SetAllowMultipleUserAssignments sets AllowMultipleUserAssignments field to given value.

### HasAllowMultipleUserAssignments

`func (o *DesktopPoolInfoV2) HasAllowMultipleUserAssignments() bool`

HasAllowMultipleUserAssignments returns a boolean if a field has been set.

### GetAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolInfoV2) GetAllowRdsPoolMultiSessionPerUser() bool`

GetAllowRdsPoolMultiSessionPerUser returns the AllowRdsPoolMultiSessionPerUser field if non-nil, zero value otherwise.

### GetAllowRdsPoolMultiSessionPerUserOk

`func (o *DesktopPoolInfoV2) GetAllowRdsPoolMultiSessionPerUserOk() (*bool, bool)`

GetAllowRdsPoolMultiSessionPerUserOk returns a tuple with the AllowRdsPoolMultiSessionPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolInfoV2) SetAllowRdsPoolMultiSessionPerUser(v bool)`

SetAllowRdsPoolMultiSessionPerUser sets AllowRdsPoolMultiSessionPerUser field to given value.

### HasAllowRdsPoolMultiSessionPerUser

`func (o *DesktopPoolInfoV2) HasAllowRdsPoolMultiSessionPerUser() bool`

HasAllowRdsPoolMultiSessionPerUser returns a boolean if a field has been set.

### GetAutomaticUserAssignment

`func (o *DesktopPoolInfoV2) GetAutomaticUserAssignment() bool`

GetAutomaticUserAssignment returns the AutomaticUserAssignment field if non-nil, zero value otherwise.

### GetAutomaticUserAssignmentOk

`func (o *DesktopPoolInfoV2) GetAutomaticUserAssignmentOk() (*bool, bool)`

GetAutomaticUserAssignmentOk returns a tuple with the AutomaticUserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomaticUserAssignment

`func (o *DesktopPoolInfoV2) SetAutomaticUserAssignment(v bool)`

SetAutomaticUserAssignment sets AutomaticUserAssignment field to given value.

### HasAutomaticUserAssignment

`func (o *DesktopPoolInfoV2) HasAutomaticUserAssignment() bool`

HasAutomaticUserAssignment returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *DesktopPoolInfoV2) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *DesktopPoolInfoV2) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *DesktopPoolInfoV2) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *DesktopPoolInfoV2) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudAssigned

`func (o *DesktopPoolInfoV2) GetCloudAssigned() bool`

GetCloudAssigned returns the CloudAssigned field if non-nil, zero value otherwise.

### GetCloudAssignedOk

`func (o *DesktopPoolInfoV2) GetCloudAssignedOk() (*bool, bool)`

GetCloudAssignedOk returns a tuple with the CloudAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAssigned

`func (o *DesktopPoolInfoV2) SetCloudAssigned(v bool)`

SetCloudAssigned sets CloudAssigned field to given value.

### HasCloudAssigned

`func (o *DesktopPoolInfoV2) HasCloudAssigned() bool`

HasCloudAssigned returns a boolean if a field has been set.

### GetCloudManaged

`func (o *DesktopPoolInfoV2) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *DesktopPoolInfoV2) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *DesktopPoolInfoV2) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *DesktopPoolInfoV2) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *DesktopPoolInfoV2) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *DesktopPoolInfoV2) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *DesktopPoolInfoV2) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *DesktopPoolInfoV2) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetCustomizationSettings

`func (o *DesktopPoolInfoV2) GetCustomizationSettings() DesktopPoolCustomizationSettings`

GetCustomizationSettings returns the CustomizationSettings field if non-nil, zero value otherwise.

### GetCustomizationSettingsOk

`func (o *DesktopPoolInfoV2) GetCustomizationSettingsOk() (*DesktopPoolCustomizationSettings, bool)`

GetCustomizationSettingsOk returns a tuple with the CustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationSettings

`func (o *DesktopPoolInfoV2) SetCustomizationSettings(v DesktopPoolCustomizationSettings)`

SetCustomizationSettings sets CustomizationSettings field to given value.

### HasCustomizationSettings

`func (o *DesktopPoolInfoV2) HasCustomizationSettings() bool`

HasCustomizationSettings returns a boolean if a field has been set.

### GetDeleteInProgress

`func (o *DesktopPoolInfoV2) GetDeleteInProgress() bool`

GetDeleteInProgress returns the DeleteInProgress field if non-nil, zero value otherwise.

### GetDeleteInProgressOk

`func (o *DesktopPoolInfoV2) GetDeleteInProgressOk() (*bool, bool)`

GetDeleteInProgressOk returns a tuple with the DeleteInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInProgress

`func (o *DesktopPoolInfoV2) SetDeleteInProgress(v bool)`

SetDeleteInProgress sets DeleteInProgress field to given value.

### HasDeleteInProgress

`func (o *DesktopPoolInfoV2) HasDeleteInProgress() bool`

HasDeleteInProgress returns a boolean if a field has been set.

### GetDescription

`func (o *DesktopPoolInfoV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DesktopPoolInfoV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DesktopPoolInfoV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DesktopPoolInfoV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayAssignedMachineName

`func (o *DesktopPoolInfoV2) GetDisplayAssignedMachineName() bool`

GetDisplayAssignedMachineName returns the DisplayAssignedMachineName field if non-nil, zero value otherwise.

### GetDisplayAssignedMachineNameOk

`func (o *DesktopPoolInfoV2) GetDisplayAssignedMachineNameOk() (*bool, bool)`

GetDisplayAssignedMachineNameOk returns a tuple with the DisplayAssignedMachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayAssignedMachineName

`func (o *DesktopPoolInfoV2) SetDisplayAssignedMachineName(v bool)`

SetDisplayAssignedMachineName sets DisplayAssignedMachineName field to given value.

### HasDisplayAssignedMachineName

`func (o *DesktopPoolInfoV2) HasDisplayAssignedMachineName() bool`

HasDisplayAssignedMachineName returns a boolean if a field has been set.

### GetDisplayName

`func (o *DesktopPoolInfoV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DesktopPoolInfoV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DesktopPoolInfoV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DesktopPoolInfoV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDisplayProtocolSettings

`func (o *DesktopPoolInfoV2) GetDisplayProtocolSettings() DesktopPoolDisplayProtocolSettings`

GetDisplayProtocolSettings returns the DisplayProtocolSettings field if non-nil, zero value otherwise.

### GetDisplayProtocolSettingsOk

`func (o *DesktopPoolInfoV2) GetDisplayProtocolSettingsOk() (*DesktopPoolDisplayProtocolSettings, bool)`

GetDisplayProtocolSettingsOk returns a tuple with the DisplayProtocolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocolSettings

`func (o *DesktopPoolInfoV2) SetDisplayProtocolSettings(v DesktopPoolDisplayProtocolSettings)`

SetDisplayProtocolSettings sets DisplayProtocolSettings field to given value.

### HasDisplayProtocolSettings

`func (o *DesktopPoolInfoV2) HasDisplayProtocolSettings() bool`

HasDisplayProtocolSettings returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *DesktopPoolInfoV2) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *DesktopPoolInfoV2) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *DesktopPoolInfoV2) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *DesktopPoolInfoV2) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnableProvisioning

`func (o *DesktopPoolInfoV2) GetEnableProvisioning() bool`

GetEnableProvisioning returns the EnableProvisioning field if non-nil, zero value otherwise.

### GetEnableProvisioningOk

`func (o *DesktopPoolInfoV2) GetEnableProvisioningOk() (*bool, bool)`

GetEnableProvisioningOk returns a tuple with the EnableProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProvisioning

`func (o *DesktopPoolInfoV2) SetEnableProvisioning(v bool)`

SetEnableProvisioning sets EnableProvisioning field to given value.

### HasEnableProvisioning

`func (o *DesktopPoolInfoV2) HasEnableProvisioning() bool`

HasEnableProvisioning returns a boolean if a field has been set.

### GetEnabled

`func (o *DesktopPoolInfoV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DesktopPoolInfoV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DesktopPoolInfoV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DesktopPoolInfoV2) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFarmId

`func (o *DesktopPoolInfoV2) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *DesktopPoolInfoV2) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *DesktopPoolInfoV2) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *DesktopPoolInfoV2) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetId

`func (o *DesktopPoolInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DesktopPoolInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DesktopPoolInfoV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DesktopPoolInfoV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageSource

`func (o *DesktopPoolInfoV2) GetImageSource() string`

GetImageSource returns the ImageSource field if non-nil, zero value otherwise.

### GetImageSourceOk

`func (o *DesktopPoolInfoV2) GetImageSourceOk() (*string, bool)`

GetImageSourceOk returns a tuple with the ImageSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageSource

`func (o *DesktopPoolInfoV2) SetImageSource(v string)`

SetImageSource sets ImageSource field to given value.

### HasImageSource

`func (o *DesktopPoolInfoV2) HasImageSource() bool`

HasImageSource returns a boolean if a field has been set.

### GetName

`func (o *DesktopPoolInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DesktopPoolInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DesktopPoolInfoV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DesktopPoolInfoV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNics

`func (o *DesktopPoolInfoV2) GetNics() []DesktopPoolNetworkInterfaceCardSettings`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *DesktopPoolInfoV2) GetNicsOk() (*[]DesktopPoolNetworkInterfaceCardSettings, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *DesktopPoolInfoV2) SetNics(v []DesktopPoolNetworkInterfaceCardSettings)`

SetNics sets Nics field to given value.

### HasNics

`func (o *DesktopPoolInfoV2) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetPatternNamingSettings

`func (o *DesktopPoolInfoV2) GetPatternNamingSettings() DesktopPoolVirtualMachinePatternNamingSettings`

GetPatternNamingSettings returns the PatternNamingSettings field if non-nil, zero value otherwise.

### GetPatternNamingSettingsOk

`func (o *DesktopPoolInfoV2) GetPatternNamingSettingsOk() (*DesktopPoolVirtualMachinePatternNamingSettings, bool)`

GetPatternNamingSettingsOk returns a tuple with the PatternNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternNamingSettings

`func (o *DesktopPoolInfoV2) SetPatternNamingSettings(v DesktopPoolVirtualMachinePatternNamingSettings)`

SetPatternNamingSettings sets PatternNamingSettings field to given value.

### HasPatternNamingSettings

`func (o *DesktopPoolInfoV2) HasPatternNamingSettings() bool`

HasPatternNamingSettings returns a boolean if a field has been set.

### GetProvisioningSettings

`func (o *DesktopPoolInfoV2) GetProvisioningSettings() DesktopPoolProvisioningSettings`

GetProvisioningSettings returns the ProvisioningSettings field if non-nil, zero value otherwise.

### GetProvisioningSettingsOk

`func (o *DesktopPoolInfoV2) GetProvisioningSettingsOk() (*DesktopPoolProvisioningSettings, bool)`

GetProvisioningSettingsOk returns a tuple with the ProvisioningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSettings

`func (o *DesktopPoolInfoV2) SetProvisioningSettings(v DesktopPoolProvisioningSettings)`

SetProvisioningSettings sets ProvisioningSettings field to given value.

### HasProvisioningSettings

`func (o *DesktopPoolInfoV2) HasProvisioningSettings() bool`

HasProvisioningSettings returns a boolean if a field has been set.

### GetProvisioningStatusData

`func (o *DesktopPoolInfoV2) GetProvisioningStatusData() DesktopPoolProvisioningStatusData`

GetProvisioningStatusData returns the ProvisioningStatusData field if non-nil, zero value otherwise.

### GetProvisioningStatusDataOk

`func (o *DesktopPoolInfoV2) GetProvisioningStatusDataOk() (*DesktopPoolProvisioningStatusData, bool)`

GetProvisioningStatusDataOk returns a tuple with the ProvisioningStatusData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningStatusData

`func (o *DesktopPoolInfoV2) SetProvisioningStatusData(v DesktopPoolProvisioningStatusData)`

SetProvisioningStatusData sets ProvisioningStatusData field to given value.

### HasProvisioningStatusData

`func (o *DesktopPoolInfoV2) HasProvisioningStatusData() bool`

HasProvisioningStatusData returns a boolean if a field has been set.

### GetSessionSettings

`func (o *DesktopPoolInfoV2) GetSessionSettings() DesktopPoolSessionSettingsV2`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *DesktopPoolInfoV2) GetSessionSettingsOk() (*DesktopPoolSessionSettingsV2, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *DesktopPoolInfoV2) SetSessionSettings(v DesktopPoolSessionSettingsV2)`

SetSessionSettings sets SessionSettings field to given value.

### HasSessionSettings

`func (o *DesktopPoolInfoV2) HasSessionSettings() bool`

HasSessionSettings returns a boolean if a field has been set.

### GetSessionType

`func (o *DesktopPoolInfoV2) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *DesktopPoolInfoV2) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *DesktopPoolInfoV2) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *DesktopPoolInfoV2) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetShortcutLocations

`func (o *DesktopPoolInfoV2) GetShortcutLocations() []string`

GetShortcutLocations returns the ShortcutLocations field if non-nil, zero value otherwise.

### GetShortcutLocationsOk

`func (o *DesktopPoolInfoV2) GetShortcutLocationsOk() (*[]string, bool)`

GetShortcutLocationsOk returns a tuple with the ShortcutLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocations

`func (o *DesktopPoolInfoV2) SetShortcutLocations(v []string)`

SetShortcutLocations sets ShortcutLocations field to given value.

### HasShortcutLocations

`func (o *DesktopPoolInfoV2) HasShortcutLocations() bool`

HasShortcutLocations returns a boolean if a field has been set.

### GetSource

`func (o *DesktopPoolInfoV2) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *DesktopPoolInfoV2) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *DesktopPoolInfoV2) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *DesktopPoolInfoV2) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSpecificNamingSettings

`func (o *DesktopPoolInfoV2) GetSpecificNamingSettings() DesktopPoolVirtualMachineSpecifiedNamingSettings`

GetSpecificNamingSettings returns the SpecificNamingSettings field if non-nil, zero value otherwise.

### GetSpecificNamingSettingsOk

`func (o *DesktopPoolInfoV2) GetSpecificNamingSettingsOk() (*DesktopPoolVirtualMachineSpecifiedNamingSettings, bool)`

GetSpecificNamingSettingsOk returns a tuple with the SpecificNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecificNamingSettings

`func (o *DesktopPoolInfoV2) SetSpecificNamingSettings(v DesktopPoolVirtualMachineSpecifiedNamingSettings)`

SetSpecificNamingSettings sets SpecificNamingSettings field to given value.

### HasSpecificNamingSettings

`func (o *DesktopPoolInfoV2) HasSpecificNamingSettings() bool`

HasSpecificNamingSettings returns a boolean if a field has been set.

### GetStopProvisioningOnError

`func (o *DesktopPoolInfoV2) GetStopProvisioningOnError() bool`

GetStopProvisioningOnError returns the StopProvisioningOnError field if non-nil, zero value otherwise.

### GetStopProvisioningOnErrorOk

`func (o *DesktopPoolInfoV2) GetStopProvisioningOnErrorOk() (*bool, bool)`

GetStopProvisioningOnErrorOk returns a tuple with the StopProvisioningOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProvisioningOnError

`func (o *DesktopPoolInfoV2) SetStopProvisioningOnError(v bool)`

SetStopProvisioningOnError sets StopProvisioningOnError field to given value.

### HasStopProvisioningOnError

`func (o *DesktopPoolInfoV2) HasStopProvisioningOnError() bool`

HasStopProvisioningOnError returns a boolean if a field has been set.

### GetStorageSettings

`func (o *DesktopPoolInfoV2) GetStorageSettings() DesktopPoolStorageSettings`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *DesktopPoolInfoV2) GetStorageSettingsOk() (*DesktopPoolStorageSettings, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *DesktopPoolInfoV2) SetStorageSettings(v DesktopPoolStorageSettings)`

SetStorageSettings sets StorageSettings field to given value.

### HasStorageSettings

`func (o *DesktopPoolInfoV2) HasStorageSettings() bool`

HasStorageSettings returns a boolean if a field has been set.

### GetTransparentPageSharingScope

`func (o *DesktopPoolInfoV2) GetTransparentPageSharingScope() string`

GetTransparentPageSharingScope returns the TransparentPageSharingScope field if non-nil, zero value otherwise.

### GetTransparentPageSharingScopeOk

`func (o *DesktopPoolInfoV2) GetTransparentPageSharingScopeOk() (*string, bool)`

GetTransparentPageSharingScopeOk returns a tuple with the TransparentPageSharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentPageSharingScope

`func (o *DesktopPoolInfoV2) SetTransparentPageSharingScope(v string)`

SetTransparentPageSharingScope sets TransparentPageSharingScope field to given value.

### HasTransparentPageSharingScope

`func (o *DesktopPoolInfoV2) HasTransparentPageSharingScope() bool`

HasTransparentPageSharingScope returns a boolean if a field has been set.

### GetType

`func (o *DesktopPoolInfoV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DesktopPoolInfoV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DesktopPoolInfoV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DesktopPoolInfoV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserAssignment

`func (o *DesktopPoolInfoV2) GetUserAssignment() string`

GetUserAssignment returns the UserAssignment field if non-nil, zero value otherwise.

### GetUserAssignmentOk

`func (o *DesktopPoolInfoV2) GetUserAssignmentOk() (*string, bool)`

GetUserAssignmentOk returns a tuple with the UserAssignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAssignment

`func (o *DesktopPoolInfoV2) SetUserAssignment(v string)`

SetUserAssignment sets UserAssignment field to given value.

### HasUserAssignment

`func (o *DesktopPoolInfoV2) HasUserAssignment() bool`

HasUserAssignment returns a boolean if a field has been set.

### GetVcenterId

`func (o *DesktopPoolInfoV2) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *DesktopPoolInfoV2) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *DesktopPoolInfoV2) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.

### HasVcenterId

`func (o *DesktopPoolInfoV2) HasVcenterId() bool`

HasVcenterId returns a boolean if a field has been set.

### GetViewStorageAcceleratorSettings

`func (o *DesktopPoolInfoV2) GetViewStorageAcceleratorSettings() DesktopPoolViewStorageAcceleratorSettings`

GetViewStorageAcceleratorSettings returns the ViewStorageAcceleratorSettings field if non-nil, zero value otherwise.

### GetViewStorageAcceleratorSettingsOk

`func (o *DesktopPoolInfoV2) GetViewStorageAcceleratorSettingsOk() (*DesktopPoolViewStorageAcceleratorSettings, bool)`

GetViewStorageAcceleratorSettingsOk returns a tuple with the ViewStorageAcceleratorSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewStorageAcceleratorSettings

`func (o *DesktopPoolInfoV2) SetViewStorageAcceleratorSettings(v DesktopPoolViewStorageAcceleratorSettings)`

SetViewStorageAcceleratorSettings sets ViewStorageAcceleratorSettings field to given value.

### HasViewStorageAcceleratorSettings

`func (o *DesktopPoolInfoV2) HasViewStorageAcceleratorSettings() bool`

HasViewStorageAcceleratorSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


