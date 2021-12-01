# DesktopPoolSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowMutilpleSessionsPerUser** | Pointer to **bool** | Whether multiple sessions are allowed per user for this pool. This is valid for RDS desktop pools.For other Desktops, allow_multiple_sessions_per_user in settings will be applicable.Default value is false. | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the desktop pool.Will be unset if the desktop does not belong to a category.This property defines valid folder names with a max length of 64 characters and up to 4 subdirectory levels.The subdirectories can be specified using a backslash, e.g. (dir1\\dir2\\dir3\\dir4). Folder names can&#39;t start orend with a backslash nor can there be 2 or more backslashes together. Combinations such as(\\dir1, dir1\\dir2\\, dir1\\\\dir2, dir1\\\\\\dir2) are invalid. The windows reserved keywords(CON, PRN, NUL, AUX, COM1 - COM9, LPT1 - LPT9 etc.) are not allowed in subdirectory names. | [optional] 
**CloudAssigned** | Pointer to **bool** | Indicates whether this desktop is assigned to a workspace in Horizon Cloud Services. Default value is false. | [optional] 
**CloudManaged** | Pointer to **bool** | Indicates whether this desktop is managed by Horizon Cloud Services.This can be false only when cloud_assigned is false. Default value is false. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | List of tags for which the access to the desktop pool is restricted to.No list indicates that desktop pool can be accessed from any connection server. | [optional] 
**DeleteInProgress** | Pointer to **bool** | Indicates whether the desktop pool is in the process of being deleted.Default value is false. | [optional] 
**DisplayProtocolSettings** | Pointer to [**DesktopPoolDisplayProtocolSettings**](DesktopPoolDisplayProtocolSettings.md) |  | [optional] 
**EnableClientRestrictions** | Pointer to **bool** | Client restrictions to be applied to the desktop pool.Currently it is valid for RDS desktop pools only. Default value is false. | [optional] 
**SessionSettings** | Pointer to [**DesktopPoolSessionSettings**](DesktopPoolSessionSettings.md) |  | [optional] 
**SessionType** | Pointer to **string** | Supported session types for this desktop pool. If application sessions are selected to besupported then this desktop pool can be used for application pool creation. This will beuseful when the machines in the pool support application remoting. Default value of DESKTOP. * DESKTOP: Only desktop sessions are supported. * APPLICATION: Only application sessions are supported. * DESKTOP_AND_APPLICATION: Both desktop and application sessions are supported. | [optional] 
**ShortcutLocations** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the desktop pool.The value will be present if categoryFolderName is set. | [optional] 

## Methods

### NewDesktopPoolSettings

`func NewDesktopPoolSettings() *DesktopPoolSettings`

NewDesktopPoolSettings instantiates a new DesktopPoolSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolSettingsWithDefaults

`func NewDesktopPoolSettingsWithDefaults() *DesktopPoolSettings`

NewDesktopPoolSettingsWithDefaults instantiates a new DesktopPoolSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowMutilpleSessionsPerUser

`func (o *DesktopPoolSettings) GetAllowMutilpleSessionsPerUser() bool`

GetAllowMutilpleSessionsPerUser returns the AllowMutilpleSessionsPerUser field if non-nil, zero value otherwise.

### GetAllowMutilpleSessionsPerUserOk

`func (o *DesktopPoolSettings) GetAllowMutilpleSessionsPerUserOk() (*bool, bool)`

GetAllowMutilpleSessionsPerUserOk returns a tuple with the AllowMutilpleSessionsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMutilpleSessionsPerUser

`func (o *DesktopPoolSettings) SetAllowMutilpleSessionsPerUser(v bool)`

SetAllowMutilpleSessionsPerUser sets AllowMutilpleSessionsPerUser field to given value.

### HasAllowMutilpleSessionsPerUser

`func (o *DesktopPoolSettings) HasAllowMutilpleSessionsPerUser() bool`

HasAllowMutilpleSessionsPerUser returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *DesktopPoolSettings) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *DesktopPoolSettings) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *DesktopPoolSettings) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *DesktopPoolSettings) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudAssigned

`func (o *DesktopPoolSettings) GetCloudAssigned() bool`

GetCloudAssigned returns the CloudAssigned field if non-nil, zero value otherwise.

### GetCloudAssignedOk

`func (o *DesktopPoolSettings) GetCloudAssignedOk() (*bool, bool)`

GetCloudAssignedOk returns a tuple with the CloudAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudAssigned

`func (o *DesktopPoolSettings) SetCloudAssigned(v bool)`

SetCloudAssigned sets CloudAssigned field to given value.

### HasCloudAssigned

`func (o *DesktopPoolSettings) HasCloudAssigned() bool`

HasCloudAssigned returns a boolean if a field has been set.

### GetCloudManaged

`func (o *DesktopPoolSettings) GetCloudManaged() bool`

GetCloudManaged returns the CloudManaged field if non-nil, zero value otherwise.

### GetCloudManagedOk

`func (o *DesktopPoolSettings) GetCloudManagedOk() (*bool, bool)`

GetCloudManagedOk returns a tuple with the CloudManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudManaged

`func (o *DesktopPoolSettings) SetCloudManaged(v bool)`

SetCloudManaged sets CloudManaged field to given value.

### HasCloudManaged

`func (o *DesktopPoolSettings) HasCloudManaged() bool`

HasCloudManaged returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *DesktopPoolSettings) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *DesktopPoolSettings) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *DesktopPoolSettings) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *DesktopPoolSettings) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetDeleteInProgress

`func (o *DesktopPoolSettings) GetDeleteInProgress() bool`

GetDeleteInProgress returns the DeleteInProgress field if non-nil, zero value otherwise.

### GetDeleteInProgressOk

`func (o *DesktopPoolSettings) GetDeleteInProgressOk() (*bool, bool)`

GetDeleteInProgressOk returns a tuple with the DeleteInProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteInProgress

`func (o *DesktopPoolSettings) SetDeleteInProgress(v bool)`

SetDeleteInProgress sets DeleteInProgress field to given value.

### HasDeleteInProgress

`func (o *DesktopPoolSettings) HasDeleteInProgress() bool`

HasDeleteInProgress returns a boolean if a field has been set.

### GetDisplayProtocolSettings

`func (o *DesktopPoolSettings) GetDisplayProtocolSettings() DesktopPoolDisplayProtocolSettings`

GetDisplayProtocolSettings returns the DisplayProtocolSettings field if non-nil, zero value otherwise.

### GetDisplayProtocolSettingsOk

`func (o *DesktopPoolSettings) GetDisplayProtocolSettingsOk() (*DesktopPoolDisplayProtocolSettings, bool)`

GetDisplayProtocolSettingsOk returns a tuple with the DisplayProtocolSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocolSettings

`func (o *DesktopPoolSettings) SetDisplayProtocolSettings(v DesktopPoolDisplayProtocolSettings)`

SetDisplayProtocolSettings sets DisplayProtocolSettings field to given value.

### HasDisplayProtocolSettings

`func (o *DesktopPoolSettings) HasDisplayProtocolSettings() bool`

HasDisplayProtocolSettings returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *DesktopPoolSettings) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *DesktopPoolSettings) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *DesktopPoolSettings) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *DesktopPoolSettings) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetSessionSettings

`func (o *DesktopPoolSettings) GetSessionSettings() DesktopPoolSessionSettings`

GetSessionSettings returns the SessionSettings field if non-nil, zero value otherwise.

### GetSessionSettingsOk

`func (o *DesktopPoolSettings) GetSessionSettingsOk() (*DesktopPoolSessionSettings, bool)`

GetSessionSettingsOk returns a tuple with the SessionSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionSettings

`func (o *DesktopPoolSettings) SetSessionSettings(v DesktopPoolSessionSettings)`

SetSessionSettings sets SessionSettings field to given value.

### HasSessionSettings

`func (o *DesktopPoolSettings) HasSessionSettings() bool`

HasSessionSettings returns a boolean if a field has been set.

### GetSessionType

`func (o *DesktopPoolSettings) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *DesktopPoolSettings) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *DesktopPoolSettings) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *DesktopPoolSettings) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetShortcutLocations

`func (o *DesktopPoolSettings) GetShortcutLocations() []string`

GetShortcutLocations returns the ShortcutLocations field if non-nil, zero value otherwise.

### GetShortcutLocationsOk

`func (o *DesktopPoolSettings) GetShortcutLocationsOk() (*[]string, bool)`

GetShortcutLocationsOk returns a tuple with the ShortcutLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocations

`func (o *DesktopPoolSettings) SetShortcutLocations(v []string)`

SetShortcutLocations sets ShortcutLocations field to given value.

### HasShortcutLocations

`func (o *DesktopPoolSettings) HasShortcutLocations() bool`

HasShortcutLocations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


