# MachineInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | Access group id of this Machine. | [optional] 
**AgentBuildNumber** | Pointer to **string** | The Horizon Agent build number.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**AgentVersion** | Pointer to **string** | The Horizon Agent version.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**Aliases** | Pointer to [**[]MachineAlias**](MachineAlias.md) | List of MachineAlias | [optional] 
**AttemptedTheftByConnectionServer** | Pointer to **[]string** | Names of the Horizon Connection Servers that attempted theft of pairing for this Agent. | [optional] 
**ConfiguredByConnectionServer** | Pointer to **[]string** | Names of the Horizon Connection Servers the Horizon Agent is paired with. | [optional] 
**DesktopPoolId** | Pointer to **string** | The id of the Desktop Pool that the machine belongs to.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**DnsName** | Pointer to **string** | DNS name of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;EndsWith&#39; and &#39;Contains&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**ManagedMachineData** | Pointer to [**ManagedMachineDataV2**](ManagedMachineDataV2.md) |  | [optional] 
**MessageSecurityEnhancedModeSupported** | Pointer to **bool** | Indicates whether ENHANCED message security mode is currently supported by this machine. | [optional] 
**MessageSecurityMode** | Pointer to **string** | The current JMS message security mode used by this machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * DISABLED: Message security mode is disabled. * MIXED: Message security mode is enabled but not enforced. * ENABLED: Message security mode is enabled. Unsigned messages are rejected by Horizon components. * ENHANCED: Message Security mode is Enhanced. Message signing and validation is performed based on the current Security Level and desktop Message Security mode. | [optional] 
**Name** | Pointer to **string** | Name of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**OperatingSystem** | Pointer to **string** | The guest operating system. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS) | [optional] 
**OperatingSystemArchitecture** | Pointer to **string** | The guest operating system architecture. * UNKNOWN: Operating System cannot be determined. * BIT_32: 32 bit Operating System Architecture. * BIT_64: 64 bit Operating System Architecture. | [optional] 
**PairingState** | Pointer to **string** | Horizon Agent pairing state. * NOT_AVAILABLE: Agent pairing state is not available. * IN_PAIRING: Agent pairing with Horizon Connection Server is in progress. * PAIRED_AND_SECURED: Agent is paired and secured with a Horizon Connection Server. | [optional] 
**RemoteExperienceAgentBuildNumber** | Pointer to **string** | The remote experience Horizon Agent build number.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**RemoteExperienceAgentVersion** | Pointer to **string** | The remote experience Horizon Agent version.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**SessionId** | Pointer to **string** | The ID of the session on the Machine (if one exists). | [optional] 
**State** | Pointer to **string** | The state of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * PROVISIONING: The machine is being provisioned. * PROVISIONING_ERROR: An error occurred during provisioning. * WAITING_FOR_AGENT: Horizon Connection Server is waiting to establish communication with Horizon Agent for one of these cases - a virtual machine in a manual desktop pool, unmanaged machine or RDS server. * CUSTOMIZING: The machine which is from an automated desktop pool is being customized after provisioning. * DELETING: The machine is marked for deletion. * MAINTENANCE: The machine is in maintenance mode. Users cannot log in or use the machine. * ERROR: An unknown error occurred in the machine. * PROVISIONED: The machine is powered off or suspended. * AGENT_UNREACHABLE: Horizon Connection Server cannot establish communication with Horizon Agent on the machine. * UNASSIGNED_USER_CONNECTED: A user other than the assigned user is logged in to the machine in a dedicated desktop pool. * CONNECTED: The machine is in an active session and has an active connection to a Horizon client. * UNASSIGNED_USER_DISCONNECTED: A user other than the assigned user is logged in and disconnected from the machine in a dedicated desktop pool. * DISCONNECTED: The machine is in an active session, but it is disconnected from the Horizon client. * AGENT_ERROR_STARTUP_IN_PROGRESS: Horizon Agent has started on the machine, but other required services such as the display protocol are still starting. * AGENT_ERROR_DISABLED: Horizon Agent is disabled. * AGENT_ERROR_INVALID_IP: Horizon Agent has an invalid IP address. * AGENT_ERROR_NEEDS_REBOOT: Horizon Agent needs reboot. * AGENT_ERROR_PROTOCOL_FAILURE: Protocol such as BLAST, RDP or PCoIP is not enabled. * AGENT_CONFIG_ERROR: The Remote Desktop Services role is not enabled on the windows server. * AGENT_DRAIN_MODE: RDS host is configured for drain mode. New connections are currently disabled. * AGENT_DRAIN_UNTIL_RESTART: RDS host is configured for drain-until-restart mode. * ALREADY_USED: The machine is configured to have only one session which is currently in progress and cannot accept new sessions. * AVAILABLE: The machine is powered on and ready for active connections. * IN_PROGRESS: There is a machine operation in progress. * DISABLED: The machine is disabled. * DISABLE_IN_PROGRESS: Disabled Horizon Connection Server still has some Horizon brokered sessions. It can still accept re-connections. * VALIDATING: The Horizon Connection Server is synchronizing state information with the agent. * UNKNOWN: Could not determine the state of the machine. | [optional] 
**Type** | Pointer to **string** | The type of machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * MANAGED_MACHINE: The machine is a managed virtual machine. * UNMANAGED_MACHINE: The machine is an unmanaged physical or virtual machine. | [optional] 
**UserIds** | Pointer to **[]string** | The unique SIDs of the users assigned to the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;NotEquals&#39; and &#39;Contains&#39;. | [optional] 

## Methods

### NewMachineInfoV2

`func NewMachineInfoV2() *MachineInfoV2`

NewMachineInfoV2 instantiates a new MachineInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineInfoV2WithDefaults

`func NewMachineInfoV2WithDefaults() *MachineInfoV2`

NewMachineInfoV2WithDefaults instantiates a new MachineInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *MachineInfoV2) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *MachineInfoV2) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *MachineInfoV2) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *MachineInfoV2) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAgentBuildNumber

`func (o *MachineInfoV2) GetAgentBuildNumber() string`

GetAgentBuildNumber returns the AgentBuildNumber field if non-nil, zero value otherwise.

### GetAgentBuildNumberOk

`func (o *MachineInfoV2) GetAgentBuildNumberOk() (*string, bool)`

GetAgentBuildNumberOk returns a tuple with the AgentBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBuildNumber

`func (o *MachineInfoV2) SetAgentBuildNumber(v string)`

SetAgentBuildNumber sets AgentBuildNumber field to given value.

### HasAgentBuildNumber

`func (o *MachineInfoV2) HasAgentBuildNumber() bool`

HasAgentBuildNumber returns a boolean if a field has been set.

### GetAgentVersion

`func (o *MachineInfoV2) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *MachineInfoV2) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *MachineInfoV2) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *MachineInfoV2) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAliases

`func (o *MachineInfoV2) GetAliases() []MachineAlias`

GetAliases returns the Aliases field if non-nil, zero value otherwise.

### GetAliasesOk

`func (o *MachineInfoV2) GetAliasesOk() (*[]MachineAlias, bool)`

GetAliasesOk returns a tuple with the Aliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliases

`func (o *MachineInfoV2) SetAliases(v []MachineAlias)`

SetAliases sets Aliases field to given value.

### HasAliases

`func (o *MachineInfoV2) HasAliases() bool`

HasAliases returns a boolean if a field has been set.

### GetAttemptedTheftByConnectionServer

`func (o *MachineInfoV2) GetAttemptedTheftByConnectionServer() []string`

GetAttemptedTheftByConnectionServer returns the AttemptedTheftByConnectionServer field if non-nil, zero value otherwise.

### GetAttemptedTheftByConnectionServerOk

`func (o *MachineInfoV2) GetAttemptedTheftByConnectionServerOk() (*[]string, bool)`

GetAttemptedTheftByConnectionServerOk returns a tuple with the AttemptedTheftByConnectionServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedTheftByConnectionServer

`func (o *MachineInfoV2) SetAttemptedTheftByConnectionServer(v []string)`

SetAttemptedTheftByConnectionServer sets AttemptedTheftByConnectionServer field to given value.

### HasAttemptedTheftByConnectionServer

`func (o *MachineInfoV2) HasAttemptedTheftByConnectionServer() bool`

HasAttemptedTheftByConnectionServer returns a boolean if a field has been set.

### GetConfiguredByConnectionServer

`func (o *MachineInfoV2) GetConfiguredByConnectionServer() []string`

GetConfiguredByConnectionServer returns the ConfiguredByConnectionServer field if non-nil, zero value otherwise.

### GetConfiguredByConnectionServerOk

`func (o *MachineInfoV2) GetConfiguredByConnectionServerOk() (*[]string, bool)`

GetConfiguredByConnectionServerOk returns a tuple with the ConfiguredByConnectionServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredByConnectionServer

`func (o *MachineInfoV2) SetConfiguredByConnectionServer(v []string)`

SetConfiguredByConnectionServer sets ConfiguredByConnectionServer field to given value.

### HasConfiguredByConnectionServer

`func (o *MachineInfoV2) HasConfiguredByConnectionServer() bool`

HasConfiguredByConnectionServer returns a boolean if a field has been set.

### GetDesktopPoolId

`func (o *MachineInfoV2) GetDesktopPoolId() string`

GetDesktopPoolId returns the DesktopPoolId field if non-nil, zero value otherwise.

### GetDesktopPoolIdOk

`func (o *MachineInfoV2) GetDesktopPoolIdOk() (*string, bool)`

GetDesktopPoolIdOk returns a tuple with the DesktopPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolId

`func (o *MachineInfoV2) SetDesktopPoolId(v string)`

SetDesktopPoolId sets DesktopPoolId field to given value.

### HasDesktopPoolId

`func (o *MachineInfoV2) HasDesktopPoolId() bool`

HasDesktopPoolId returns a boolean if a field has been set.

### GetDnsName

`func (o *MachineInfoV2) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *MachineInfoV2) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *MachineInfoV2) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *MachineInfoV2) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetId

`func (o *MachineInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineInfoV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MachineInfoV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagedMachineData

`func (o *MachineInfoV2) GetManagedMachineData() ManagedMachineDataV2`

GetManagedMachineData returns the ManagedMachineData field if non-nil, zero value otherwise.

### GetManagedMachineDataOk

`func (o *MachineInfoV2) GetManagedMachineDataOk() (*ManagedMachineDataV2, bool)`

GetManagedMachineDataOk returns a tuple with the ManagedMachineData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedMachineData

`func (o *MachineInfoV2) SetManagedMachineData(v ManagedMachineDataV2)`

SetManagedMachineData sets ManagedMachineData field to given value.

### HasManagedMachineData

`func (o *MachineInfoV2) HasManagedMachineData() bool`

HasManagedMachineData returns a boolean if a field has been set.

### GetMessageSecurityEnhancedModeSupported

`func (o *MachineInfoV2) GetMessageSecurityEnhancedModeSupported() bool`

GetMessageSecurityEnhancedModeSupported returns the MessageSecurityEnhancedModeSupported field if non-nil, zero value otherwise.

### GetMessageSecurityEnhancedModeSupportedOk

`func (o *MachineInfoV2) GetMessageSecurityEnhancedModeSupportedOk() (*bool, bool)`

GetMessageSecurityEnhancedModeSupportedOk returns a tuple with the MessageSecurityEnhancedModeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityEnhancedModeSupported

`func (o *MachineInfoV2) SetMessageSecurityEnhancedModeSupported(v bool)`

SetMessageSecurityEnhancedModeSupported sets MessageSecurityEnhancedModeSupported field to given value.

### HasMessageSecurityEnhancedModeSupported

`func (o *MachineInfoV2) HasMessageSecurityEnhancedModeSupported() bool`

HasMessageSecurityEnhancedModeSupported returns a boolean if a field has been set.

### GetMessageSecurityMode

`func (o *MachineInfoV2) GetMessageSecurityMode() string`

GetMessageSecurityMode returns the MessageSecurityMode field if non-nil, zero value otherwise.

### GetMessageSecurityModeOk

`func (o *MachineInfoV2) GetMessageSecurityModeOk() (*string, bool)`

GetMessageSecurityModeOk returns a tuple with the MessageSecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityMode

`func (o *MachineInfoV2) SetMessageSecurityMode(v string)`

SetMessageSecurityMode sets MessageSecurityMode field to given value.

### HasMessageSecurityMode

`func (o *MachineInfoV2) HasMessageSecurityMode() bool`

HasMessageSecurityMode returns a boolean if a field has been set.

### GetName

`func (o *MachineInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineInfoV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineInfoV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *MachineInfoV2) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MachineInfoV2) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *MachineInfoV2) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *MachineInfoV2) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperatingSystemArchitecture

`func (o *MachineInfoV2) GetOperatingSystemArchitecture() string`

GetOperatingSystemArchitecture returns the OperatingSystemArchitecture field if non-nil, zero value otherwise.

### GetOperatingSystemArchitectureOk

`func (o *MachineInfoV2) GetOperatingSystemArchitectureOk() (*string, bool)`

GetOperatingSystemArchitectureOk returns a tuple with the OperatingSystemArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemArchitecture

`func (o *MachineInfoV2) SetOperatingSystemArchitecture(v string)`

SetOperatingSystemArchitecture sets OperatingSystemArchitecture field to given value.

### HasOperatingSystemArchitecture

`func (o *MachineInfoV2) HasOperatingSystemArchitecture() bool`

HasOperatingSystemArchitecture returns a boolean if a field has been set.

### GetPairingState

`func (o *MachineInfoV2) GetPairingState() string`

GetPairingState returns the PairingState field if non-nil, zero value otherwise.

### GetPairingStateOk

`func (o *MachineInfoV2) GetPairingStateOk() (*string, bool)`

GetPairingStateOk returns a tuple with the PairingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingState

`func (o *MachineInfoV2) SetPairingState(v string)`

SetPairingState sets PairingState field to given value.

### HasPairingState

`func (o *MachineInfoV2) HasPairingState() bool`

HasPairingState returns a boolean if a field has been set.

### GetRemoteExperienceAgentBuildNumber

`func (o *MachineInfoV2) GetRemoteExperienceAgentBuildNumber() string`

GetRemoteExperienceAgentBuildNumber returns the RemoteExperienceAgentBuildNumber field if non-nil, zero value otherwise.

### GetRemoteExperienceAgentBuildNumberOk

`func (o *MachineInfoV2) GetRemoteExperienceAgentBuildNumberOk() (*string, bool)`

GetRemoteExperienceAgentBuildNumberOk returns a tuple with the RemoteExperienceAgentBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteExperienceAgentBuildNumber

`func (o *MachineInfoV2) SetRemoteExperienceAgentBuildNumber(v string)`

SetRemoteExperienceAgentBuildNumber sets RemoteExperienceAgentBuildNumber field to given value.

### HasRemoteExperienceAgentBuildNumber

`func (o *MachineInfoV2) HasRemoteExperienceAgentBuildNumber() bool`

HasRemoteExperienceAgentBuildNumber returns a boolean if a field has been set.

### GetRemoteExperienceAgentVersion

`func (o *MachineInfoV2) GetRemoteExperienceAgentVersion() string`

GetRemoteExperienceAgentVersion returns the RemoteExperienceAgentVersion field if non-nil, zero value otherwise.

### GetRemoteExperienceAgentVersionOk

`func (o *MachineInfoV2) GetRemoteExperienceAgentVersionOk() (*string, bool)`

GetRemoteExperienceAgentVersionOk returns a tuple with the RemoteExperienceAgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteExperienceAgentVersion

`func (o *MachineInfoV2) SetRemoteExperienceAgentVersion(v string)`

SetRemoteExperienceAgentVersion sets RemoteExperienceAgentVersion field to given value.

### HasRemoteExperienceAgentVersion

`func (o *MachineInfoV2) HasRemoteExperienceAgentVersion() bool`

HasRemoteExperienceAgentVersion returns a boolean if a field has been set.

### GetSessionId

`func (o *MachineInfoV2) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *MachineInfoV2) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *MachineInfoV2) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *MachineInfoV2) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetState

`func (o *MachineInfoV2) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MachineInfoV2) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MachineInfoV2) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MachineInfoV2) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *MachineInfoV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MachineInfoV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MachineInfoV2) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MachineInfoV2) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserIds

`func (o *MachineInfoV2) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *MachineInfoV2) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *MachineInfoV2) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *MachineInfoV2) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


