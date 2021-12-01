# ApplicationPoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | Access groups can organize the entities (like application pools, desktop pools) in the organization. They can also be used for delegated administration. For application pool, this is the same as that of the farm or desktop pool that the application pool belongs to.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**AntiAffinityData** | Pointer to [**ApplicationAntiAffinityData**](ApplicationAntiAffinityData.md) |  | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the application. Unset if the application does not belong to a category. | [optional] 
**CsRestrictionTags** | Pointer to **[]string** | Connection server restrictions. Application pool can be accessed from only those connection server instances that have a matching tag in this list. Null or empty list means that the application pool can be accessed from any connection server. | [optional] 
**CustomizedIconIds** | Pointer to **[]string** | List of customized icon IDs associated with the application which the user has configured. | [optional] 
**Description** | Pointer to **string** | Notes about the application pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DesktopPoolId** | Pointer to **string** | ID of the desktop pool from which this application pool is created. Either this or farm id will be set.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**DisplayName** | Pointer to **string** | The display name is the name that users will see when they connect to view client. If the display name is left blank, it defaults to name.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**EnableClientRestrictions** | Pointer to **bool** | Indicates whether client restrictions are to be applied to application pool. Currently it is valid for application pool created from farm.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**EnablePreLaunch** | Pointer to **bool** | Whether to pre-launch the application.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the application pool is enabled.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ExecutablePath** | Pointer to **string** | Path to application executable.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;EndsWith&#39; and &#39;Contains&#39;. | [optional] 
**FarmId** | Pointer to **string** | ID of the farm from which this application pool is created. Either this or desktop pool id will be set.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**IconIds** | Pointer to **[]string** | List of icon IDs associated with the application which are fetched from the agent. | [optional] 
**Id** | Pointer to **string** | Unique ID representing application pool.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**MaxMultiSessions** | Pointer to **int32** | Maximum number of multi-sessions a user can have in this application pool. This property is required if multi-session mode is set to \&quot;ENABLED_DEFAULT_OFF\&quot;, \&quot;ENABLED_DEFAULT_ON\&quot;, or \&quot;ENABLED_ENFORCED\&quot; | [optional] 
**MultiSessionMode** | Pointer to **string** | Multi-session mode for the application pool. An application launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. Multi-session mode should be disabled when pre-launch is enabled.&lt;br&gt;Supported Filters: &#39;Equals&#39;. * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application. | [optional] 
**Name** | Pointer to **string** | The application name is the unique identifier used to identify this application pool. This property must contain only alphanumerics, underscores, and dashes. The maximum length is 64 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**Parameters** | Pointer to **string** | Parameters to pass to application when launching. | [optional] 
**Publisher** | Pointer to **string** | Application publisher.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**ShortcutLocations** | Pointer to **[]string** | Locations of the category folder in the user&#39;s OS containing a shortcut to the application. The value must be set if category folder name is provided. | [optional] 
**StartFolder** | Pointer to **string** | Starting folder for application. | [optional] 
**SupportedFileTypesData** | Pointer to [**ApplicationSupportedFileTypesData**](ApplicationSupportedFileTypesData.md) |  | [optional] 
**Version** | Pointer to **string** | Application version.&lt;br&gt;Supported Filters: &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 

## Methods

### NewApplicationPoolInfo

`func NewApplicationPoolInfo() *ApplicationPoolInfo`

NewApplicationPoolInfo instantiates a new ApplicationPoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPoolInfoWithDefaults

`func NewApplicationPoolInfoWithDefaults() *ApplicationPoolInfo`

NewApplicationPoolInfoWithDefaults instantiates a new ApplicationPoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *ApplicationPoolInfo) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *ApplicationPoolInfo) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *ApplicationPoolInfo) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *ApplicationPoolInfo) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAntiAffinityData

`func (o *ApplicationPoolInfo) GetAntiAffinityData() ApplicationAntiAffinityData`

GetAntiAffinityData returns the AntiAffinityData field if non-nil, zero value otherwise.

### GetAntiAffinityDataOk

`func (o *ApplicationPoolInfo) GetAntiAffinityDataOk() (*ApplicationAntiAffinityData, bool)`

GetAntiAffinityDataOk returns a tuple with the AntiAffinityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinityData

`func (o *ApplicationPoolInfo) SetAntiAffinityData(v ApplicationAntiAffinityData)`

SetAntiAffinityData sets AntiAffinityData field to given value.

### HasAntiAffinityData

`func (o *ApplicationPoolInfo) HasAntiAffinityData() bool`

HasAntiAffinityData returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *ApplicationPoolInfo) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *ApplicationPoolInfo) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *ApplicationPoolInfo) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *ApplicationPoolInfo) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *ApplicationPoolInfo) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *ApplicationPoolInfo) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *ApplicationPoolInfo) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *ApplicationPoolInfo) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetCustomizedIconIds

`func (o *ApplicationPoolInfo) GetCustomizedIconIds() []string`

GetCustomizedIconIds returns the CustomizedIconIds field if non-nil, zero value otherwise.

### GetCustomizedIconIdsOk

`func (o *ApplicationPoolInfo) GetCustomizedIconIdsOk() (*[]string, bool)`

GetCustomizedIconIdsOk returns a tuple with the CustomizedIconIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedIconIds

`func (o *ApplicationPoolInfo) SetCustomizedIconIds(v []string)`

SetCustomizedIconIds sets CustomizedIconIds field to given value.

### HasCustomizedIconIds

`func (o *ApplicationPoolInfo) HasCustomizedIconIds() bool`

HasCustomizedIconIds returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationPoolInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationPoolInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationPoolInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationPoolInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesktopPoolId

`func (o *ApplicationPoolInfo) GetDesktopPoolId() string`

GetDesktopPoolId returns the DesktopPoolId field if non-nil, zero value otherwise.

### GetDesktopPoolIdOk

`func (o *ApplicationPoolInfo) GetDesktopPoolIdOk() (*string, bool)`

GetDesktopPoolIdOk returns a tuple with the DesktopPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolId

`func (o *ApplicationPoolInfo) SetDesktopPoolId(v string)`

SetDesktopPoolId sets DesktopPoolId field to given value.

### HasDesktopPoolId

`func (o *ApplicationPoolInfo) HasDesktopPoolId() bool`

HasDesktopPoolId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApplicationPoolInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApplicationPoolInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApplicationPoolInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApplicationPoolInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *ApplicationPoolInfo) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *ApplicationPoolInfo) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *ApplicationPoolInfo) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *ApplicationPoolInfo) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnablePreLaunch

`func (o *ApplicationPoolInfo) GetEnablePreLaunch() bool`

GetEnablePreLaunch returns the EnablePreLaunch field if non-nil, zero value otherwise.

### GetEnablePreLaunchOk

`func (o *ApplicationPoolInfo) GetEnablePreLaunchOk() (*bool, bool)`

GetEnablePreLaunchOk returns a tuple with the EnablePreLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePreLaunch

`func (o *ApplicationPoolInfo) SetEnablePreLaunch(v bool)`

SetEnablePreLaunch sets EnablePreLaunch field to given value.

### HasEnablePreLaunch

`func (o *ApplicationPoolInfo) HasEnablePreLaunch() bool`

HasEnablePreLaunch returns a boolean if a field has been set.

### GetEnabled

`func (o *ApplicationPoolInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationPoolInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationPoolInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationPoolInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExecutablePath

`func (o *ApplicationPoolInfo) GetExecutablePath() string`

GetExecutablePath returns the ExecutablePath field if non-nil, zero value otherwise.

### GetExecutablePathOk

`func (o *ApplicationPoolInfo) GetExecutablePathOk() (*string, bool)`

GetExecutablePathOk returns a tuple with the ExecutablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutablePath

`func (o *ApplicationPoolInfo) SetExecutablePath(v string)`

SetExecutablePath sets ExecutablePath field to given value.

### HasExecutablePath

`func (o *ApplicationPoolInfo) HasExecutablePath() bool`

HasExecutablePath returns a boolean if a field has been set.

### GetFarmId

`func (o *ApplicationPoolInfo) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *ApplicationPoolInfo) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *ApplicationPoolInfo) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *ApplicationPoolInfo) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetIconIds

`func (o *ApplicationPoolInfo) GetIconIds() []string`

GetIconIds returns the IconIds field if non-nil, zero value otherwise.

### GetIconIdsOk

`func (o *ApplicationPoolInfo) GetIconIdsOk() (*[]string, bool)`

GetIconIdsOk returns a tuple with the IconIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconIds

`func (o *ApplicationPoolInfo) SetIconIds(v []string)`

SetIconIds sets IconIds field to given value.

### HasIconIds

`func (o *ApplicationPoolInfo) HasIconIds() bool`

HasIconIds returns a boolean if a field has been set.

### GetId

`func (o *ApplicationPoolInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationPoolInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationPoolInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationPoolInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxMultiSessions

`func (o *ApplicationPoolInfo) GetMaxMultiSessions() int32`

GetMaxMultiSessions returns the MaxMultiSessions field if non-nil, zero value otherwise.

### GetMaxMultiSessionsOk

`func (o *ApplicationPoolInfo) GetMaxMultiSessionsOk() (*int32, bool)`

GetMaxMultiSessionsOk returns a tuple with the MaxMultiSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMultiSessions

`func (o *ApplicationPoolInfo) SetMaxMultiSessions(v int32)`

SetMaxMultiSessions sets MaxMultiSessions field to given value.

### HasMaxMultiSessions

`func (o *ApplicationPoolInfo) HasMaxMultiSessions() bool`

HasMaxMultiSessions returns a boolean if a field has been set.

### GetMultiSessionMode

`func (o *ApplicationPoolInfo) GetMultiSessionMode() string`

GetMultiSessionMode returns the MultiSessionMode field if non-nil, zero value otherwise.

### GetMultiSessionModeOk

`func (o *ApplicationPoolInfo) GetMultiSessionModeOk() (*string, bool)`

GetMultiSessionModeOk returns a tuple with the MultiSessionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSessionMode

`func (o *ApplicationPoolInfo) SetMultiSessionMode(v string)`

SetMultiSessionMode sets MultiSessionMode field to given value.

### HasMultiSessionMode

`func (o *ApplicationPoolInfo) HasMultiSessionMode() bool`

HasMultiSessionMode returns a boolean if a field has been set.

### GetName

`func (o *ApplicationPoolInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationPoolInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationPoolInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationPoolInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ApplicationPoolInfo) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ApplicationPoolInfo) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ApplicationPoolInfo) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ApplicationPoolInfo) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPublisher

`func (o *ApplicationPoolInfo) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ApplicationPoolInfo) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ApplicationPoolInfo) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ApplicationPoolInfo) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetShortcutLocations

`func (o *ApplicationPoolInfo) GetShortcutLocations() []string`

GetShortcutLocations returns the ShortcutLocations field if non-nil, zero value otherwise.

### GetShortcutLocationsOk

`func (o *ApplicationPoolInfo) GetShortcutLocationsOk() (*[]string, bool)`

GetShortcutLocationsOk returns a tuple with the ShortcutLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocations

`func (o *ApplicationPoolInfo) SetShortcutLocations(v []string)`

SetShortcutLocations sets ShortcutLocations field to given value.

### HasShortcutLocations

`func (o *ApplicationPoolInfo) HasShortcutLocations() bool`

HasShortcutLocations returns a boolean if a field has been set.

### GetStartFolder

`func (o *ApplicationPoolInfo) GetStartFolder() string`

GetStartFolder returns the StartFolder field if non-nil, zero value otherwise.

### GetStartFolderOk

`func (o *ApplicationPoolInfo) GetStartFolderOk() (*string, bool)`

GetStartFolderOk returns a tuple with the StartFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFolder

`func (o *ApplicationPoolInfo) SetStartFolder(v string)`

SetStartFolder sets StartFolder field to given value.

### HasStartFolder

`func (o *ApplicationPoolInfo) HasStartFolder() bool`

HasStartFolder returns a boolean if a field has been set.

### GetSupportedFileTypesData

`func (o *ApplicationPoolInfo) GetSupportedFileTypesData() ApplicationSupportedFileTypesData`

GetSupportedFileTypesData returns the SupportedFileTypesData field if non-nil, zero value otherwise.

### GetSupportedFileTypesDataOk

`func (o *ApplicationPoolInfo) GetSupportedFileTypesDataOk() (*ApplicationSupportedFileTypesData, bool)`

GetSupportedFileTypesDataOk returns a tuple with the SupportedFileTypesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFileTypesData

`func (o *ApplicationPoolInfo) SetSupportedFileTypesData(v ApplicationSupportedFileTypesData)`

SetSupportedFileTypesData sets SupportedFileTypesData field to given value.

### HasSupportedFileTypesData

`func (o *ApplicationPoolInfo) HasSupportedFileTypesData() bool`

HasSupportedFileTypesData returns a boolean if a field has been set.

### GetVersion

`func (o *ApplicationPoolInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationPoolInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationPoolInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplicationPoolInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


