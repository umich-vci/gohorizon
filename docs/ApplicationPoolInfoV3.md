# ApplicationPoolInfoV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | Access groups can organize the entities (like application pools, desktop pools) in the organization. They can also be used for delegated administration. For application pool, this is the same as that of the farm or desktop pool that the application pool belongs to.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**AntiAffinityData** | Pointer to [**ApplicationAntiAffinityData**](ApplicationAntiAffinityData.md) |  | [optional] 
**CategoryFolderName** | Pointer to **string** | Name of the category folder in the user&#39;s OS containing a shortcut to the application. Unset if the application does not belong to a category. | [optional] 
**CloudBrokered** | Pointer to **bool** | Indicates whether the application pool is cloud brokered.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
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
**GlobalApplicationEntitlementId** | Pointer to **string** | Global application entitlement for this application pool. Caller should have permission to FEDERATED_LDAP_VIEW privilege for this field to be populated or to use in filter.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
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

### NewApplicationPoolInfoV3

`func NewApplicationPoolInfoV3() *ApplicationPoolInfoV3`

NewApplicationPoolInfoV3 instantiates a new ApplicationPoolInfoV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationPoolInfoV3WithDefaults

`func NewApplicationPoolInfoV3WithDefaults() *ApplicationPoolInfoV3`

NewApplicationPoolInfoV3WithDefaults instantiates a new ApplicationPoolInfoV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *ApplicationPoolInfoV3) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *ApplicationPoolInfoV3) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *ApplicationPoolInfoV3) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *ApplicationPoolInfoV3) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAntiAffinityData

`func (o *ApplicationPoolInfoV3) GetAntiAffinityData() ApplicationAntiAffinityData`

GetAntiAffinityData returns the AntiAffinityData field if non-nil, zero value otherwise.

### GetAntiAffinityDataOk

`func (o *ApplicationPoolInfoV3) GetAntiAffinityDataOk() (*ApplicationAntiAffinityData, bool)`

GetAntiAffinityDataOk returns a tuple with the AntiAffinityData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntiAffinityData

`func (o *ApplicationPoolInfoV3) SetAntiAffinityData(v ApplicationAntiAffinityData)`

SetAntiAffinityData sets AntiAffinityData field to given value.

### HasAntiAffinityData

`func (o *ApplicationPoolInfoV3) HasAntiAffinityData() bool`

HasAntiAffinityData returns a boolean if a field has been set.

### GetCategoryFolderName

`func (o *ApplicationPoolInfoV3) GetCategoryFolderName() string`

GetCategoryFolderName returns the CategoryFolderName field if non-nil, zero value otherwise.

### GetCategoryFolderNameOk

`func (o *ApplicationPoolInfoV3) GetCategoryFolderNameOk() (*string, bool)`

GetCategoryFolderNameOk returns a tuple with the CategoryFolderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryFolderName

`func (o *ApplicationPoolInfoV3) SetCategoryFolderName(v string)`

SetCategoryFolderName sets CategoryFolderName field to given value.

### HasCategoryFolderName

`func (o *ApplicationPoolInfoV3) HasCategoryFolderName() bool`

HasCategoryFolderName returns a boolean if a field has been set.

### GetCloudBrokered

`func (o *ApplicationPoolInfoV3) GetCloudBrokered() bool`

GetCloudBrokered returns the CloudBrokered field if non-nil, zero value otherwise.

### GetCloudBrokeredOk

`func (o *ApplicationPoolInfoV3) GetCloudBrokeredOk() (*bool, bool)`

GetCloudBrokeredOk returns a tuple with the CloudBrokered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudBrokered

`func (o *ApplicationPoolInfoV3) SetCloudBrokered(v bool)`

SetCloudBrokered sets CloudBrokered field to given value.

### HasCloudBrokered

`func (o *ApplicationPoolInfoV3) HasCloudBrokered() bool`

HasCloudBrokered returns a boolean if a field has been set.

### GetCsRestrictionTags

`func (o *ApplicationPoolInfoV3) GetCsRestrictionTags() []string`

GetCsRestrictionTags returns the CsRestrictionTags field if non-nil, zero value otherwise.

### GetCsRestrictionTagsOk

`func (o *ApplicationPoolInfoV3) GetCsRestrictionTagsOk() (*[]string, bool)`

GetCsRestrictionTagsOk returns a tuple with the CsRestrictionTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsRestrictionTags

`func (o *ApplicationPoolInfoV3) SetCsRestrictionTags(v []string)`

SetCsRestrictionTags sets CsRestrictionTags field to given value.

### HasCsRestrictionTags

`func (o *ApplicationPoolInfoV3) HasCsRestrictionTags() bool`

HasCsRestrictionTags returns a boolean if a field has been set.

### GetCustomizedIconIds

`func (o *ApplicationPoolInfoV3) GetCustomizedIconIds() []string`

GetCustomizedIconIds returns the CustomizedIconIds field if non-nil, zero value otherwise.

### GetCustomizedIconIdsOk

`func (o *ApplicationPoolInfoV3) GetCustomizedIconIdsOk() (*[]string, bool)`

GetCustomizedIconIdsOk returns a tuple with the CustomizedIconIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizedIconIds

`func (o *ApplicationPoolInfoV3) SetCustomizedIconIds(v []string)`

SetCustomizedIconIds sets CustomizedIconIds field to given value.

### HasCustomizedIconIds

`func (o *ApplicationPoolInfoV3) HasCustomizedIconIds() bool`

HasCustomizedIconIds returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationPoolInfoV3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationPoolInfoV3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationPoolInfoV3) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationPoolInfoV3) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDesktopPoolId

`func (o *ApplicationPoolInfoV3) GetDesktopPoolId() string`

GetDesktopPoolId returns the DesktopPoolId field if non-nil, zero value otherwise.

### GetDesktopPoolIdOk

`func (o *ApplicationPoolInfoV3) GetDesktopPoolIdOk() (*string, bool)`

GetDesktopPoolIdOk returns a tuple with the DesktopPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolId

`func (o *ApplicationPoolInfoV3) SetDesktopPoolId(v string)`

SetDesktopPoolId sets DesktopPoolId field to given value.

### HasDesktopPoolId

`func (o *ApplicationPoolInfoV3) HasDesktopPoolId() bool`

HasDesktopPoolId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ApplicationPoolInfoV3) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ApplicationPoolInfoV3) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ApplicationPoolInfoV3) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ApplicationPoolInfoV3) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableClientRestrictions

`func (o *ApplicationPoolInfoV3) GetEnableClientRestrictions() bool`

GetEnableClientRestrictions returns the EnableClientRestrictions field if non-nil, zero value otherwise.

### GetEnableClientRestrictionsOk

`func (o *ApplicationPoolInfoV3) GetEnableClientRestrictionsOk() (*bool, bool)`

GetEnableClientRestrictionsOk returns a tuple with the EnableClientRestrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableClientRestrictions

`func (o *ApplicationPoolInfoV3) SetEnableClientRestrictions(v bool)`

SetEnableClientRestrictions sets EnableClientRestrictions field to given value.

### HasEnableClientRestrictions

`func (o *ApplicationPoolInfoV3) HasEnableClientRestrictions() bool`

HasEnableClientRestrictions returns a boolean if a field has been set.

### GetEnablePreLaunch

`func (o *ApplicationPoolInfoV3) GetEnablePreLaunch() bool`

GetEnablePreLaunch returns the EnablePreLaunch field if non-nil, zero value otherwise.

### GetEnablePreLaunchOk

`func (o *ApplicationPoolInfoV3) GetEnablePreLaunchOk() (*bool, bool)`

GetEnablePreLaunchOk returns a tuple with the EnablePreLaunch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnablePreLaunch

`func (o *ApplicationPoolInfoV3) SetEnablePreLaunch(v bool)`

SetEnablePreLaunch sets EnablePreLaunch field to given value.

### HasEnablePreLaunch

`func (o *ApplicationPoolInfoV3) HasEnablePreLaunch() bool`

HasEnablePreLaunch returns a boolean if a field has been set.

### GetEnabled

`func (o *ApplicationPoolInfoV3) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationPoolInfoV3) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationPoolInfoV3) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationPoolInfoV3) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExecutablePath

`func (o *ApplicationPoolInfoV3) GetExecutablePath() string`

GetExecutablePath returns the ExecutablePath field if non-nil, zero value otherwise.

### GetExecutablePathOk

`func (o *ApplicationPoolInfoV3) GetExecutablePathOk() (*string, bool)`

GetExecutablePathOk returns a tuple with the ExecutablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutablePath

`func (o *ApplicationPoolInfoV3) SetExecutablePath(v string)`

SetExecutablePath sets ExecutablePath field to given value.

### HasExecutablePath

`func (o *ApplicationPoolInfoV3) HasExecutablePath() bool`

HasExecutablePath returns a boolean if a field has been set.

### GetFarmId

`func (o *ApplicationPoolInfoV3) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *ApplicationPoolInfoV3) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *ApplicationPoolInfoV3) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *ApplicationPoolInfoV3) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetGlobalApplicationEntitlementId

`func (o *ApplicationPoolInfoV3) GetGlobalApplicationEntitlementId() string`

GetGlobalApplicationEntitlementId returns the GlobalApplicationEntitlementId field if non-nil, zero value otherwise.

### GetGlobalApplicationEntitlementIdOk

`func (o *ApplicationPoolInfoV3) GetGlobalApplicationEntitlementIdOk() (*string, bool)`

GetGlobalApplicationEntitlementIdOk returns a tuple with the GlobalApplicationEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApplicationEntitlementId

`func (o *ApplicationPoolInfoV3) SetGlobalApplicationEntitlementId(v string)`

SetGlobalApplicationEntitlementId sets GlobalApplicationEntitlementId field to given value.

### HasGlobalApplicationEntitlementId

`func (o *ApplicationPoolInfoV3) HasGlobalApplicationEntitlementId() bool`

HasGlobalApplicationEntitlementId returns a boolean if a field has been set.

### GetIconIds

`func (o *ApplicationPoolInfoV3) GetIconIds() []string`

GetIconIds returns the IconIds field if non-nil, zero value otherwise.

### GetIconIdsOk

`func (o *ApplicationPoolInfoV3) GetIconIdsOk() (*[]string, bool)`

GetIconIdsOk returns a tuple with the IconIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconIds

`func (o *ApplicationPoolInfoV3) SetIconIds(v []string)`

SetIconIds sets IconIds field to given value.

### HasIconIds

`func (o *ApplicationPoolInfoV3) HasIconIds() bool`

HasIconIds returns a boolean if a field has been set.

### GetId

`func (o *ApplicationPoolInfoV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationPoolInfoV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationPoolInfoV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationPoolInfoV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMaxMultiSessions

`func (o *ApplicationPoolInfoV3) GetMaxMultiSessions() int32`

GetMaxMultiSessions returns the MaxMultiSessions field if non-nil, zero value otherwise.

### GetMaxMultiSessionsOk

`func (o *ApplicationPoolInfoV3) GetMaxMultiSessionsOk() (*int32, bool)`

GetMaxMultiSessionsOk returns a tuple with the MaxMultiSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMultiSessions

`func (o *ApplicationPoolInfoV3) SetMaxMultiSessions(v int32)`

SetMaxMultiSessions sets MaxMultiSessions field to given value.

### HasMaxMultiSessions

`func (o *ApplicationPoolInfoV3) HasMaxMultiSessions() bool`

HasMaxMultiSessions returns a boolean if a field has been set.

### GetMultiSessionMode

`func (o *ApplicationPoolInfoV3) GetMultiSessionMode() string`

GetMultiSessionMode returns the MultiSessionMode field if non-nil, zero value otherwise.

### GetMultiSessionModeOk

`func (o *ApplicationPoolInfoV3) GetMultiSessionModeOk() (*string, bool)`

GetMultiSessionModeOk returns a tuple with the MultiSessionMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSessionMode

`func (o *ApplicationPoolInfoV3) SetMultiSessionMode(v string)`

SetMultiSessionMode sets MultiSessionMode field to given value.

### HasMultiSessionMode

`func (o *ApplicationPoolInfoV3) HasMultiSessionMode() bool`

HasMultiSessionMode returns a boolean if a field has been set.

### GetName

`func (o *ApplicationPoolInfoV3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationPoolInfoV3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationPoolInfoV3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApplicationPoolInfoV3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParameters

`func (o *ApplicationPoolInfoV3) GetParameters() string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ApplicationPoolInfoV3) GetParametersOk() (*string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ApplicationPoolInfoV3) SetParameters(v string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ApplicationPoolInfoV3) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPublisher

`func (o *ApplicationPoolInfoV3) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ApplicationPoolInfoV3) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ApplicationPoolInfoV3) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ApplicationPoolInfoV3) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetShortcutLocations

`func (o *ApplicationPoolInfoV3) GetShortcutLocations() []string`

GetShortcutLocations returns the ShortcutLocations field if non-nil, zero value otherwise.

### GetShortcutLocationsOk

`func (o *ApplicationPoolInfoV3) GetShortcutLocationsOk() (*[]string, bool)`

GetShortcutLocationsOk returns a tuple with the ShortcutLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortcutLocations

`func (o *ApplicationPoolInfoV3) SetShortcutLocations(v []string)`

SetShortcutLocations sets ShortcutLocations field to given value.

### HasShortcutLocations

`func (o *ApplicationPoolInfoV3) HasShortcutLocations() bool`

HasShortcutLocations returns a boolean if a field has been set.

### GetStartFolder

`func (o *ApplicationPoolInfoV3) GetStartFolder() string`

GetStartFolder returns the StartFolder field if non-nil, zero value otherwise.

### GetStartFolderOk

`func (o *ApplicationPoolInfoV3) GetStartFolderOk() (*string, bool)`

GetStartFolderOk returns a tuple with the StartFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartFolder

`func (o *ApplicationPoolInfoV3) SetStartFolder(v string)`

SetStartFolder sets StartFolder field to given value.

### HasStartFolder

`func (o *ApplicationPoolInfoV3) HasStartFolder() bool`

HasStartFolder returns a boolean if a field has been set.

### GetSupportedFileTypesData

`func (o *ApplicationPoolInfoV3) GetSupportedFileTypesData() ApplicationSupportedFileTypesData`

GetSupportedFileTypesData returns the SupportedFileTypesData field if non-nil, zero value otherwise.

### GetSupportedFileTypesDataOk

`func (o *ApplicationPoolInfoV3) GetSupportedFileTypesDataOk() (*ApplicationSupportedFileTypesData, bool)`

GetSupportedFileTypesDataOk returns a tuple with the SupportedFileTypesData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFileTypesData

`func (o *ApplicationPoolInfoV3) SetSupportedFileTypesData(v ApplicationSupportedFileTypesData)`

SetSupportedFileTypesData sets SupportedFileTypesData field to given value.

### HasSupportedFileTypesData

`func (o *ApplicationPoolInfoV3) HasSupportedFileTypesData() bool`

HasSupportedFileTypesData returns a boolean if a field has been set.

### GetVersion

`func (o *ApplicationPoolInfoV3) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationPoolInfoV3) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationPoolInfoV3) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplicationPoolInfoV3) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


