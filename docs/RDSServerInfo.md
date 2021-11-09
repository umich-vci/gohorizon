# RDSServerInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | The id of the Access Group that the RDS Server belongs to.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**AgentBuildNumber** | Pointer to **string** | The Horizon Agent build number.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**AgentVersion** | Pointer to **string** | The Horizon Agent version.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**BaseVmId** | Pointer to **string** | The base vm id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**BaseVmSnapshotId** | Pointer to **string** | The base vm snapshot id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**Description** | Pointer to **string** | Description of the RDS Server.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DnsName** | Pointer to **string** | DNS name of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;EndsWith&#39; and &#39;Contains&#39;. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if RDS server is enabled.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**FarmId** | Pointer to **string** | The id of the Farm that the RDS Server belongs to.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the RDS Server.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**ImageManagementStreamId** | Pointer to **string** | The id of the image management stream. This will be populated only for RDS server belonging to Instant Clone farms created using image catalog.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**ImageManagementTagId** | Pointer to **string** | The id of the image management tag. This will be populated only for RDS server belonging to Instant Clone farms created using image catalog.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**LoadIndex** | Pointer to **int32** | This value is similar to load_preference and represents the load on RDS Server in the range of 0 to 100. | [optional] 
**LoadPreference** | Pointer to **string** | Based on the current load of this RDS Server, gives a measure of how preferential this server is to be chosen for new application sessions. * BLOCK: RDS Server is blocked and new sessions will not be accepted. * HEAVY: RDS Server is experiencing heavy load and should likely not be chosen for new sessions. * NORMAL: RDS Server is experiencing normal load and is okay to be chosen for new sessions. * LIGHT: RDS Server is experiencing light load and is okay to be chosen for new sessions. * UNKNOWN: RDS Server did not report a load preference. This is potentially a configuration issue if other RDS Servers in the same Farm do report load preferences. | [optional] 
**LogoffPolicy** | Pointer to **string** | The user log off behavior at the time of maintenance. * FORCE_LOGOFF: Users will be forced to log off when the system is ready to execute the operation. Before being forcibly logged off, users may have a grace period in which to save their work which can be configured in Global Settings. * WAIT_FOR_LOGOFF: Wait for connected users to disconnect before the task starts. The operation starts immediately when there are no active sessions. | [optional] 
**MaxSessionsCount** | Pointer to **int32** | Maximum number of sessions for RDS server as reported by the Horizon Agent. This will be unset if the value is not configured. | [optional] 
**MaxSessionsCountConfigured** | Pointer to **int32** | Maximum number of sessions for RDS server as configured by administrator. This will be unset if the value is not configured. | [optional] 
**MaxSessionsType** | Pointer to **string** | RDS Server max sessions type as reported by the Horizon Agent. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions. | [optional] 
**MaxSessionsTypeConfigured** | Pointer to **string** | The configured RDS Server max sessions type. If the RDS Server is part of an automated farm, this value is inherited from the farm configuration. This property has a default value of UNCONFIGURED. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions. * UNCONFIGURED: The max number of sessions has not yet been defined for the RDSServer. | [optional] 
**MessageSecurityEnhancedModeSupported** | Pointer to **bool** | Indicates whether ENHANCED message security mode is currently supported by this machine. | [optional] 
**MessageSecurityMode** | Pointer to **string** | The current JMS message security mode used by this machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * DISABLED: Message security mode is disabled. * MIXED: Message security mode is enabled but not enforced. * ENABLED: Message security mode is enabled. Unsigned messages are rejected by Horizon components. * ENHANCED: Message Security mode is Enhanced. Message signing and validation is performed based on the current Security Level and desktop Message Security mode. | [optional] 
**Name** | Pointer to **string** | Name of the RDS Server.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**OperatingSystem** | Pointer to **string** | The machine operating system. * UNKNOWN: Unknown * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_SERVER_OTHER: Linux Server (other) | [optional] 
**Operation** | Pointer to **string** | The current maintenance operation on the RDS Server.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * PUSH_IMAGE: A push image operation. | [optional] 
**OperationState** | Pointer to **string** | The state of the current maintenance operation on the RDS Server.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * UNDEFINED: The operation state is unrecognized. * SCHEDULED: The operation is scheduled for future execution. * PROGRESSING: The operation is in progress. * COMPLETED: The operation has completed. * FAULT: The operation has encountered an error. * CANCELLING: The operation has been cancelled. * HOLDING: The operation has been paused. * CREATE: The operation is being initiated. | [optional] 
**PendingBaseVmId** | Pointer to **string** | The pending base vm id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**PendingBaseVmSnapshotId** | Pointer to **string** | The pending base vm snapshot id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**PendingImageManagementStreamId** | Pointer to **string** | The id of the pending image management stream. This will be populated only for RDS server belonging to Instant Clone farms created using image catalog.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**PendingImageManagementTagId** | Pointer to **string** | The id of the pending image management tag. This will be populated only for RDS server belonging to Instant Clone farms created using image catalog.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**RemoteExperienceAgentBuildNumber** | Pointer to **string** | The remote experience Horizon Agent build number.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**RemoteExperienceAgentVersion** | Pointer to **string** | The remote experience Horizon Agent version.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**SessionCount** | Pointer to **int32** | RDS server session count.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**State** | Pointer to **string** | The state of the RDS server.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * WAIT_FOR_AGENT: Connection Server is waiting to establish communication with View Agent on a RDS Server. * AGENT_UNREACHABLE: Connection Server cannot establish communication with View Agent on a RDS Server. * AGENT_CONFIG_ERROR: The RDS Server has configuration error. Ex: Remote Desktop Service role is not enabled. * AVAILABLE: The RDS Server is powered on and ready for an active connection. * DISABLED: The RDS Server is disabled. * DISABLE_IN_PROGRESS: Disabled RDS server still has some brokered sessions. It can still accept re-connections. * PROVISIONING: The RDS Server is being provisioned. * PROVISIONING_ERROR: An error occurred during provisioning. * CUSTOMIZING: The RDS Server is being customized. * DELETING: The RDS Server is marked for deletion. It will be deleted soon. * MAINTENANCE: The RDS Server is in maintenance mode. * ERROR: An unknown error occurred in the RDS Server. * PROVISIONED: The RDS Server has been provisioned. * CONNECTED: The RDS Server is in an active session. * DISCONNECTED: The RDS Server is in an active session but is disconnected. * AGENT_ERR_STARTUP_IN_PROGRESS: The Horizon Agent has started on the virtual machine, but other required services such as the display protocol are still starting. * AGENT_ERR_DISABLED: The Horizon Agent is disabled. * AGENT_ERR_INVALID_IP: The Horizon Agent has an invalid IP. * AGENT_ERR_NEED_REBOOT: The Horizon Agent needs reboot. * AGENT_ERR_PROTOCOL_FAILURE: Protocol such as RDP or PCoIP is not enabled. * AGENT_ERR_DOMAIN_FAILURE: The RDS Server has an invalid domain. * AGENT_DRAIN_MODE: The RDS Server is configured for drain mode. * AGENT_DRAIN_UNTIL_RESTART: The RDS Server is configured a mode to drain until restart. * ALREADY_USED: The RDS Server cannot accept new sessions. * IN_PROGRESS: There is a RDS Server operation in progress. * VALIDATING: The connection server is synchronizing state information with the agent. * UNKNOWN: Could not determine the state of the RDS Server. | [optional] 

## Methods

### NewRDSServerInfo

`func NewRDSServerInfo() *RDSServerInfo`

NewRDSServerInfo instantiates a new RDSServerInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSServerInfoWithDefaults

`func NewRDSServerInfoWithDefaults() *RDSServerInfo`

NewRDSServerInfoWithDefaults instantiates a new RDSServerInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *RDSServerInfo) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *RDSServerInfo) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *RDSServerInfo) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *RDSServerInfo) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAgentBuildNumber

`func (o *RDSServerInfo) GetAgentBuildNumber() string`

GetAgentBuildNumber returns the AgentBuildNumber field if non-nil, zero value otherwise.

### GetAgentBuildNumberOk

`func (o *RDSServerInfo) GetAgentBuildNumberOk() (*string, bool)`

GetAgentBuildNumberOk returns a tuple with the AgentBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBuildNumber

`func (o *RDSServerInfo) SetAgentBuildNumber(v string)`

SetAgentBuildNumber sets AgentBuildNumber field to given value.

### HasAgentBuildNumber

`func (o *RDSServerInfo) HasAgentBuildNumber() bool`

HasAgentBuildNumber returns a boolean if a field has been set.

### GetAgentVersion

`func (o *RDSServerInfo) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *RDSServerInfo) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *RDSServerInfo) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *RDSServerInfo) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetBaseVmId

`func (o *RDSServerInfo) GetBaseVmId() string`

GetBaseVmId returns the BaseVmId field if non-nil, zero value otherwise.

### GetBaseVmIdOk

`func (o *RDSServerInfo) GetBaseVmIdOk() (*string, bool)`

GetBaseVmIdOk returns a tuple with the BaseVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVmId

`func (o *RDSServerInfo) SetBaseVmId(v string)`

SetBaseVmId sets BaseVmId field to given value.

### HasBaseVmId

`func (o *RDSServerInfo) HasBaseVmId() bool`

HasBaseVmId returns a boolean if a field has been set.

### GetBaseVmSnapshotId

`func (o *RDSServerInfo) GetBaseVmSnapshotId() string`

GetBaseVmSnapshotId returns the BaseVmSnapshotId field if non-nil, zero value otherwise.

### GetBaseVmSnapshotIdOk

`func (o *RDSServerInfo) GetBaseVmSnapshotIdOk() (*string, bool)`

GetBaseVmSnapshotIdOk returns a tuple with the BaseVmSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVmSnapshotId

`func (o *RDSServerInfo) SetBaseVmSnapshotId(v string)`

SetBaseVmSnapshotId sets BaseVmSnapshotId field to given value.

### HasBaseVmSnapshotId

`func (o *RDSServerInfo) HasBaseVmSnapshotId() bool`

HasBaseVmSnapshotId returns a boolean if a field has been set.

### GetDescription

`func (o *RDSServerInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RDSServerInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RDSServerInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RDSServerInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsName

`func (o *RDSServerInfo) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *RDSServerInfo) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *RDSServerInfo) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *RDSServerInfo) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetEnabled

`func (o *RDSServerInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RDSServerInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RDSServerInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RDSServerInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFarmId

`func (o *RDSServerInfo) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *RDSServerInfo) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *RDSServerInfo) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *RDSServerInfo) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetId

`func (o *RDSServerInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RDSServerInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RDSServerInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RDSServerInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetImageManagementStreamId

`func (o *RDSServerInfo) GetImageManagementStreamId() string`

GetImageManagementStreamId returns the ImageManagementStreamId field if non-nil, zero value otherwise.

### GetImageManagementStreamIdOk

`func (o *RDSServerInfo) GetImageManagementStreamIdOk() (*string, bool)`

GetImageManagementStreamIdOk returns a tuple with the ImageManagementStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageManagementStreamId

`func (o *RDSServerInfo) SetImageManagementStreamId(v string)`

SetImageManagementStreamId sets ImageManagementStreamId field to given value.

### HasImageManagementStreamId

`func (o *RDSServerInfo) HasImageManagementStreamId() bool`

HasImageManagementStreamId returns a boolean if a field has been set.

### GetImageManagementTagId

`func (o *RDSServerInfo) GetImageManagementTagId() string`

GetImageManagementTagId returns the ImageManagementTagId field if non-nil, zero value otherwise.

### GetImageManagementTagIdOk

`func (o *RDSServerInfo) GetImageManagementTagIdOk() (*string, bool)`

GetImageManagementTagIdOk returns a tuple with the ImageManagementTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageManagementTagId

`func (o *RDSServerInfo) SetImageManagementTagId(v string)`

SetImageManagementTagId sets ImageManagementTagId field to given value.

### HasImageManagementTagId

`func (o *RDSServerInfo) HasImageManagementTagId() bool`

HasImageManagementTagId returns a boolean if a field has been set.

### GetLoadIndex

`func (o *RDSServerInfo) GetLoadIndex() int32`

GetLoadIndex returns the LoadIndex field if non-nil, zero value otherwise.

### GetLoadIndexOk

`func (o *RDSServerInfo) GetLoadIndexOk() (*int32, bool)`

GetLoadIndexOk returns a tuple with the LoadIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadIndex

`func (o *RDSServerInfo) SetLoadIndex(v int32)`

SetLoadIndex sets LoadIndex field to given value.

### HasLoadIndex

`func (o *RDSServerInfo) HasLoadIndex() bool`

HasLoadIndex returns a boolean if a field has been set.

### GetLoadPreference

`func (o *RDSServerInfo) GetLoadPreference() string`

GetLoadPreference returns the LoadPreference field if non-nil, zero value otherwise.

### GetLoadPreferenceOk

`func (o *RDSServerInfo) GetLoadPreferenceOk() (*string, bool)`

GetLoadPreferenceOk returns a tuple with the LoadPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadPreference

`func (o *RDSServerInfo) SetLoadPreference(v string)`

SetLoadPreference sets LoadPreference field to given value.

### HasLoadPreference

`func (o *RDSServerInfo) HasLoadPreference() bool`

HasLoadPreference returns a boolean if a field has been set.

### GetLogoffPolicy

`func (o *RDSServerInfo) GetLogoffPolicy() string`

GetLogoffPolicy returns the LogoffPolicy field if non-nil, zero value otherwise.

### GetLogoffPolicyOk

`func (o *RDSServerInfo) GetLogoffPolicyOk() (*string, bool)`

GetLogoffPolicyOk returns a tuple with the LogoffPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoffPolicy

`func (o *RDSServerInfo) SetLogoffPolicy(v string)`

SetLogoffPolicy sets LogoffPolicy field to given value.

### HasLogoffPolicy

`func (o *RDSServerInfo) HasLogoffPolicy() bool`

HasLogoffPolicy returns a boolean if a field has been set.

### GetMaxSessionsCount

`func (o *RDSServerInfo) GetMaxSessionsCount() int32`

GetMaxSessionsCount returns the MaxSessionsCount field if non-nil, zero value otherwise.

### GetMaxSessionsCountOk

`func (o *RDSServerInfo) GetMaxSessionsCountOk() (*int32, bool)`

GetMaxSessionsCountOk returns a tuple with the MaxSessionsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionsCount

`func (o *RDSServerInfo) SetMaxSessionsCount(v int32)`

SetMaxSessionsCount sets MaxSessionsCount field to given value.

### HasMaxSessionsCount

`func (o *RDSServerInfo) HasMaxSessionsCount() bool`

HasMaxSessionsCount returns a boolean if a field has been set.

### GetMaxSessionsCountConfigured

`func (o *RDSServerInfo) GetMaxSessionsCountConfigured() int32`

GetMaxSessionsCountConfigured returns the MaxSessionsCountConfigured field if non-nil, zero value otherwise.

### GetMaxSessionsCountConfiguredOk

`func (o *RDSServerInfo) GetMaxSessionsCountConfiguredOk() (*int32, bool)`

GetMaxSessionsCountConfiguredOk returns a tuple with the MaxSessionsCountConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionsCountConfigured

`func (o *RDSServerInfo) SetMaxSessionsCountConfigured(v int32)`

SetMaxSessionsCountConfigured sets MaxSessionsCountConfigured field to given value.

### HasMaxSessionsCountConfigured

`func (o *RDSServerInfo) HasMaxSessionsCountConfigured() bool`

HasMaxSessionsCountConfigured returns a boolean if a field has been set.

### GetMaxSessionsType

`func (o *RDSServerInfo) GetMaxSessionsType() string`

GetMaxSessionsType returns the MaxSessionsType field if non-nil, zero value otherwise.

### GetMaxSessionsTypeOk

`func (o *RDSServerInfo) GetMaxSessionsTypeOk() (*string, bool)`

GetMaxSessionsTypeOk returns a tuple with the MaxSessionsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionsType

`func (o *RDSServerInfo) SetMaxSessionsType(v string)`

SetMaxSessionsType sets MaxSessionsType field to given value.

### HasMaxSessionsType

`func (o *RDSServerInfo) HasMaxSessionsType() bool`

HasMaxSessionsType returns a boolean if a field has been set.

### GetMaxSessionsTypeConfigured

`func (o *RDSServerInfo) GetMaxSessionsTypeConfigured() string`

GetMaxSessionsTypeConfigured returns the MaxSessionsTypeConfigured field if non-nil, zero value otherwise.

### GetMaxSessionsTypeConfiguredOk

`func (o *RDSServerInfo) GetMaxSessionsTypeConfiguredOk() (*string, bool)`

GetMaxSessionsTypeConfiguredOk returns a tuple with the MaxSessionsTypeConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionsTypeConfigured

`func (o *RDSServerInfo) SetMaxSessionsTypeConfigured(v string)`

SetMaxSessionsTypeConfigured sets MaxSessionsTypeConfigured field to given value.

### HasMaxSessionsTypeConfigured

`func (o *RDSServerInfo) HasMaxSessionsTypeConfigured() bool`

HasMaxSessionsTypeConfigured returns a boolean if a field has been set.

### GetMessageSecurityEnhancedModeSupported

`func (o *RDSServerInfo) GetMessageSecurityEnhancedModeSupported() bool`

GetMessageSecurityEnhancedModeSupported returns the MessageSecurityEnhancedModeSupported field if non-nil, zero value otherwise.

### GetMessageSecurityEnhancedModeSupportedOk

`func (o *RDSServerInfo) GetMessageSecurityEnhancedModeSupportedOk() (*bool, bool)`

GetMessageSecurityEnhancedModeSupportedOk returns a tuple with the MessageSecurityEnhancedModeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityEnhancedModeSupported

`func (o *RDSServerInfo) SetMessageSecurityEnhancedModeSupported(v bool)`

SetMessageSecurityEnhancedModeSupported sets MessageSecurityEnhancedModeSupported field to given value.

### HasMessageSecurityEnhancedModeSupported

`func (o *RDSServerInfo) HasMessageSecurityEnhancedModeSupported() bool`

HasMessageSecurityEnhancedModeSupported returns a boolean if a field has been set.

### GetMessageSecurityMode

`func (o *RDSServerInfo) GetMessageSecurityMode() string`

GetMessageSecurityMode returns the MessageSecurityMode field if non-nil, zero value otherwise.

### GetMessageSecurityModeOk

`func (o *RDSServerInfo) GetMessageSecurityModeOk() (*string, bool)`

GetMessageSecurityModeOk returns a tuple with the MessageSecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityMode

`func (o *RDSServerInfo) SetMessageSecurityMode(v string)`

SetMessageSecurityMode sets MessageSecurityMode field to given value.

### HasMessageSecurityMode

`func (o *RDSServerInfo) HasMessageSecurityMode() bool`

HasMessageSecurityMode returns a boolean if a field has been set.

### GetName

`func (o *RDSServerInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RDSServerInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RDSServerInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RDSServerInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *RDSServerInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *RDSServerInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *RDSServerInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *RDSServerInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperation

`func (o *RDSServerInfo) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *RDSServerInfo) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *RDSServerInfo) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *RDSServerInfo) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOperationState

`func (o *RDSServerInfo) GetOperationState() string`

GetOperationState returns the OperationState field if non-nil, zero value otherwise.

### GetOperationStateOk

`func (o *RDSServerInfo) GetOperationStateOk() (*string, bool)`

GetOperationStateOk returns a tuple with the OperationState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationState

`func (o *RDSServerInfo) SetOperationState(v string)`

SetOperationState sets OperationState field to given value.

### HasOperationState

`func (o *RDSServerInfo) HasOperationState() bool`

HasOperationState returns a boolean if a field has been set.

### GetPendingBaseVmId

`func (o *RDSServerInfo) GetPendingBaseVmId() string`

GetPendingBaseVmId returns the PendingBaseVmId field if non-nil, zero value otherwise.

### GetPendingBaseVmIdOk

`func (o *RDSServerInfo) GetPendingBaseVmIdOk() (*string, bool)`

GetPendingBaseVmIdOk returns a tuple with the PendingBaseVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBaseVmId

`func (o *RDSServerInfo) SetPendingBaseVmId(v string)`

SetPendingBaseVmId sets PendingBaseVmId field to given value.

### HasPendingBaseVmId

`func (o *RDSServerInfo) HasPendingBaseVmId() bool`

HasPendingBaseVmId returns a boolean if a field has been set.

### GetPendingBaseVmSnapshotId

`func (o *RDSServerInfo) GetPendingBaseVmSnapshotId() string`

GetPendingBaseVmSnapshotId returns the PendingBaseVmSnapshotId field if non-nil, zero value otherwise.

### GetPendingBaseVmSnapshotIdOk

`func (o *RDSServerInfo) GetPendingBaseVmSnapshotIdOk() (*string, bool)`

GetPendingBaseVmSnapshotIdOk returns a tuple with the PendingBaseVmSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingBaseVmSnapshotId

`func (o *RDSServerInfo) SetPendingBaseVmSnapshotId(v string)`

SetPendingBaseVmSnapshotId sets PendingBaseVmSnapshotId field to given value.

### HasPendingBaseVmSnapshotId

`func (o *RDSServerInfo) HasPendingBaseVmSnapshotId() bool`

HasPendingBaseVmSnapshotId returns a boolean if a field has been set.

### GetPendingImageManagementStreamId

`func (o *RDSServerInfo) GetPendingImageManagementStreamId() string`

GetPendingImageManagementStreamId returns the PendingImageManagementStreamId field if non-nil, zero value otherwise.

### GetPendingImageManagementStreamIdOk

`func (o *RDSServerInfo) GetPendingImageManagementStreamIdOk() (*string, bool)`

GetPendingImageManagementStreamIdOk returns a tuple with the PendingImageManagementStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingImageManagementStreamId

`func (o *RDSServerInfo) SetPendingImageManagementStreamId(v string)`

SetPendingImageManagementStreamId sets PendingImageManagementStreamId field to given value.

### HasPendingImageManagementStreamId

`func (o *RDSServerInfo) HasPendingImageManagementStreamId() bool`

HasPendingImageManagementStreamId returns a boolean if a field has been set.

### GetPendingImageManagementTagId

`func (o *RDSServerInfo) GetPendingImageManagementTagId() string`

GetPendingImageManagementTagId returns the PendingImageManagementTagId field if non-nil, zero value otherwise.

### GetPendingImageManagementTagIdOk

`func (o *RDSServerInfo) GetPendingImageManagementTagIdOk() (*string, bool)`

GetPendingImageManagementTagIdOk returns a tuple with the PendingImageManagementTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingImageManagementTagId

`func (o *RDSServerInfo) SetPendingImageManagementTagId(v string)`

SetPendingImageManagementTagId sets PendingImageManagementTagId field to given value.

### HasPendingImageManagementTagId

`func (o *RDSServerInfo) HasPendingImageManagementTagId() bool`

HasPendingImageManagementTagId returns a boolean if a field has been set.

### GetRemoteExperienceAgentBuildNumber

`func (o *RDSServerInfo) GetRemoteExperienceAgentBuildNumber() string`

GetRemoteExperienceAgentBuildNumber returns the RemoteExperienceAgentBuildNumber field if non-nil, zero value otherwise.

### GetRemoteExperienceAgentBuildNumberOk

`func (o *RDSServerInfo) GetRemoteExperienceAgentBuildNumberOk() (*string, bool)`

GetRemoteExperienceAgentBuildNumberOk returns a tuple with the RemoteExperienceAgentBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteExperienceAgentBuildNumber

`func (o *RDSServerInfo) SetRemoteExperienceAgentBuildNumber(v string)`

SetRemoteExperienceAgentBuildNumber sets RemoteExperienceAgentBuildNumber field to given value.

### HasRemoteExperienceAgentBuildNumber

`func (o *RDSServerInfo) HasRemoteExperienceAgentBuildNumber() bool`

HasRemoteExperienceAgentBuildNumber returns a boolean if a field has been set.

### GetRemoteExperienceAgentVersion

`func (o *RDSServerInfo) GetRemoteExperienceAgentVersion() string`

GetRemoteExperienceAgentVersion returns the RemoteExperienceAgentVersion field if non-nil, zero value otherwise.

### GetRemoteExperienceAgentVersionOk

`func (o *RDSServerInfo) GetRemoteExperienceAgentVersionOk() (*string, bool)`

GetRemoteExperienceAgentVersionOk returns a tuple with the RemoteExperienceAgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteExperienceAgentVersion

`func (o *RDSServerInfo) SetRemoteExperienceAgentVersion(v string)`

SetRemoteExperienceAgentVersion sets RemoteExperienceAgentVersion field to given value.

### HasRemoteExperienceAgentVersion

`func (o *RDSServerInfo) HasRemoteExperienceAgentVersion() bool`

HasRemoteExperienceAgentVersion returns a boolean if a field has been set.

### GetSessionCount

`func (o *RDSServerInfo) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *RDSServerInfo) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *RDSServerInfo) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *RDSServerInfo) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetState

`func (o *RDSServerInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RDSServerInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RDSServerInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *RDSServerInfo) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


