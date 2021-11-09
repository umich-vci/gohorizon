# MachineInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBuildNumber** | Pointer to **string** | The Horizon Agent build number.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**AgentVersion** | Pointer to **string** | The Horizon Agent version.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**AttemptedTheftByConnectionServer** | Pointer to **[]string** | Names of the Horizon Connection Servers that attempted theft of pairing for this Agent. | [optional] 
**ConfiguredByConnectionServer** | Pointer to **[]string** | Names of the Horizon Connection Servers the Horizon Agent is paired with. | [optional] 
**DesktopPoolId** | Pointer to **string** | The id of the Desktop Pool that the machine belongs to.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**DnsName** | Pointer to **string** | DNS name of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;EndsWith&#39; and &#39;Contains&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**ManagedMachineData** | Pointer to [**ManagedMachineData**](ManagedMachineData.md) |  | [optional] 
**MessageSecurityEnhancedModeSupported** | Pointer to **bool** | Indicates whether ENHANCED message security mode is currently supported by this machine. | [optional] 
**MessageSecurityMode** | Pointer to **string** | The current JMS message security mode used by this machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * DISABLED: Message security mode is disabled. * MIXED: Message security mode is enabled but not enforced. * ENABLED: Message security mode is enabled. Unsigned messages are rejected by Horizon components. * ENHANCED: Message Security mode is Enhanced. Message signing and validation is performed based on the current Security Level and desktop Message Security mode. | [optional] 
**Name** | Pointer to **string** | Name of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**OperatingSystem** | Pointer to **string** | The guest operating system. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS) | [optional] 
**OperatingSystemArchitecture** | Pointer to **string** | The guest operating system architecture. * UNKNOWN: Operating System cannot be determined. * BIT_32: 32 bit Operating System Architecture. * BIT_64: 64 bit Operating System Architecture. | [optional] 
**PairingState** | Pointer to **string** | Horizon Agent pairing state. * NOT_AVAILABLE: Agent pairing state is not available. * IN_PAIRING: Agent pairing with Horizon Connection Server is in progress. * PAIRED_AND_SECURED: Agent is paired and secured with a Horizon Connection Server. | [optional] 
**RemoteExperienceAgentBuildNumber** | Pointer to **string** | The remote experience Horizon Agent build number.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**RemoteExperienceAgentVersion** | Pointer to **string** | The remote experience Horizon Agent version.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**State** | Pointer to **string** | The state of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * PROVISIONING: The machine is being provisioned. * PROVISIONING_ERROR: An error occurred during provisioning. * WAITING_FOR_AGENT: Horizon Connection Server is waiting to establish communication with Horizon Agent for one of these cases - a virtual machine in a manual desktop pool, unmanaged machine or RDS server. * CUSTOMIZING: The machine which is from an automated desktop pool is being customized after provisioning. * DELETING: The machine is marked for deletion. * MAINTENANCE: The machine is in maintenance mode. Users cannot log in or use the machine. * ERROR: An unknown error occurred in the machine. * PROVISIONED: The machine is powered off or suspended. * AGENT_UNREACHABLE: Horizon Connection Server cannot establish communication with Horizon Agent on the machine. * UNASSIGNED_USER_CONNECTED: A user other than the assigned user is logged in to the machine in a dedicated desktop pool. * CONNECTED: The machine is in an active session and has an active connection to a Horizon client. * UNASSIGNED_USER_DISCONNECTED: A user other than the assigned user is logged in and disconnected from the machine in a dedicated desktop pool. * DISCONNECTED: The machine is in an active session, but it is disconnected from the Horizon client. * AGENT_ERROR_STARTUP_IN_PROGRESS: Horizon Agent has started on the machine, but other required services such as the display protocol are still starting. * AGENT_ERROR_DISABLED: Horizon Agent is disabled. * AGENT_ERROR_INVALID_IP: Horizon Agent has an invalid IP address. * AGENT_ERROR_NEEDS_REBOOT: Horizon Agent needs reboot. * AGENT_ERROR_PROTOCOL_FAILURE: Protocol such as BLAST, RDP or PCoIP is not enabled. * AGENT_CONFIG_ERROR: The Remote Desktop Services role is not enabled on the windows server. * AGENT_DRAIN_MODE: RDS host is configured for drain mode. New connections are currently disabled. * AGENT_DRAIN_UNTIL_RESTART: RDS host is configured for drain-until-restart mode. * ALREADY_USED: The machine is configured to have only one session which is currently in progress and cannot accept new sessions. * AVAILABLE: The machine is powered on and ready for active connections. * IN_PROGRESS: There is a machine operation in progress. * DISABLED: The machine is disabled. * DISABLE_IN_PROGRESS: Disabled Horizon Connection Server still has some Horizon brokered sessions. It can still accept re-connections. * VALIDATING: The Horizon Connection Server is synchronizing state information with the agent. * UNKNOWN: Could not determine the state of the machine. | [optional] 
**Type** | Pointer to **string** | The type of machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * MANAGED_MACHINE: The machine is a managed virtual machine. * UNMANAGED_MACHINE: The machine is an unmanaged physical or virtual machine. | [optional] 
**UserIds** | Pointer to **[]string** | The unique SIDs of the users assigned to the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;NotEquals&#39; and &#39;Contains&#39;. | [optional] 

## Methods

### NewMachineInfo

`func NewMachineInfo() *MachineInfo`

NewMachineInfo instantiates a new MachineInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineInfoWithDefaults

`func NewMachineInfoWithDefaults() *MachineInfo`

NewMachineInfoWithDefaults instantiates a new MachineInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBuildNumber

`func (o *MachineInfo) GetAgentBuildNumber() string`

GetAgentBuildNumber returns the AgentBuildNumber field if non-nil, zero value otherwise.

### GetAgentBuildNumberOk

`func (o *MachineInfo) GetAgentBuildNumberOk() (*string, bool)`

GetAgentBuildNumberOk returns a tuple with the AgentBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBuildNumber

`func (o *MachineInfo) SetAgentBuildNumber(v string)`

SetAgentBuildNumber sets AgentBuildNumber field to given value.

### HasAgentBuildNumber

`func (o *MachineInfo) HasAgentBuildNumber() bool`

HasAgentBuildNumber returns a boolean if a field has been set.

### GetAgentVersion

`func (o *MachineInfo) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *MachineInfo) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *MachineInfo) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *MachineInfo) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetAttemptedTheftByConnectionServer

`func (o *MachineInfo) GetAttemptedTheftByConnectionServer() []string`

GetAttemptedTheftByConnectionServer returns the AttemptedTheftByConnectionServer field if non-nil, zero value otherwise.

### GetAttemptedTheftByConnectionServerOk

`func (o *MachineInfo) GetAttemptedTheftByConnectionServerOk() (*[]string, bool)`

GetAttemptedTheftByConnectionServerOk returns a tuple with the AttemptedTheftByConnectionServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedTheftByConnectionServer

`func (o *MachineInfo) SetAttemptedTheftByConnectionServer(v []string)`

SetAttemptedTheftByConnectionServer sets AttemptedTheftByConnectionServer field to given value.

### HasAttemptedTheftByConnectionServer

`func (o *MachineInfo) HasAttemptedTheftByConnectionServer() bool`

HasAttemptedTheftByConnectionServer returns a boolean if a field has been set.

### GetConfiguredByConnectionServer

`func (o *MachineInfo) GetConfiguredByConnectionServer() []string`

GetConfiguredByConnectionServer returns the ConfiguredByConnectionServer field if non-nil, zero value otherwise.

### GetConfiguredByConnectionServerOk

`func (o *MachineInfo) GetConfiguredByConnectionServerOk() (*[]string, bool)`

GetConfiguredByConnectionServerOk returns a tuple with the ConfiguredByConnectionServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredByConnectionServer

`func (o *MachineInfo) SetConfiguredByConnectionServer(v []string)`

SetConfiguredByConnectionServer sets ConfiguredByConnectionServer field to given value.

### HasConfiguredByConnectionServer

`func (o *MachineInfo) HasConfiguredByConnectionServer() bool`

HasConfiguredByConnectionServer returns a boolean if a field has been set.

### GetDesktopPoolId

`func (o *MachineInfo) GetDesktopPoolId() string`

GetDesktopPoolId returns the DesktopPoolId field if non-nil, zero value otherwise.

### GetDesktopPoolIdOk

`func (o *MachineInfo) GetDesktopPoolIdOk() (*string, bool)`

GetDesktopPoolIdOk returns a tuple with the DesktopPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolId

`func (o *MachineInfo) SetDesktopPoolId(v string)`

SetDesktopPoolId sets DesktopPoolId field to given value.

### HasDesktopPoolId

`func (o *MachineInfo) HasDesktopPoolId() bool`

HasDesktopPoolId returns a boolean if a field has been set.

### GetDnsName

`func (o *MachineInfo) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *MachineInfo) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *MachineInfo) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *MachineInfo) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetId

`func (o *MachineInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MachineInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MachineInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MachineInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetManagedMachineData

`func (o *MachineInfo) GetManagedMachineData() ManagedMachineData`

GetManagedMachineData returns the ManagedMachineData field if non-nil, zero value otherwise.

### GetManagedMachineDataOk

`func (o *MachineInfo) GetManagedMachineDataOk() (*ManagedMachineData, bool)`

GetManagedMachineDataOk returns a tuple with the ManagedMachineData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagedMachineData

`func (o *MachineInfo) SetManagedMachineData(v ManagedMachineData)`

SetManagedMachineData sets ManagedMachineData field to given value.

### HasManagedMachineData

`func (o *MachineInfo) HasManagedMachineData() bool`

HasManagedMachineData returns a boolean if a field has been set.

### GetMessageSecurityEnhancedModeSupported

`func (o *MachineInfo) GetMessageSecurityEnhancedModeSupported() bool`

GetMessageSecurityEnhancedModeSupported returns the MessageSecurityEnhancedModeSupported field if non-nil, zero value otherwise.

### GetMessageSecurityEnhancedModeSupportedOk

`func (o *MachineInfo) GetMessageSecurityEnhancedModeSupportedOk() (*bool, bool)`

GetMessageSecurityEnhancedModeSupportedOk returns a tuple with the MessageSecurityEnhancedModeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityEnhancedModeSupported

`func (o *MachineInfo) SetMessageSecurityEnhancedModeSupported(v bool)`

SetMessageSecurityEnhancedModeSupported sets MessageSecurityEnhancedModeSupported field to given value.

### HasMessageSecurityEnhancedModeSupported

`func (o *MachineInfo) HasMessageSecurityEnhancedModeSupported() bool`

HasMessageSecurityEnhancedModeSupported returns a boolean if a field has been set.

### GetMessageSecurityMode

`func (o *MachineInfo) GetMessageSecurityMode() string`

GetMessageSecurityMode returns the MessageSecurityMode field if non-nil, zero value otherwise.

### GetMessageSecurityModeOk

`func (o *MachineInfo) GetMessageSecurityModeOk() (*string, bool)`

GetMessageSecurityModeOk returns a tuple with the MessageSecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityMode

`func (o *MachineInfo) SetMessageSecurityMode(v string)`

SetMessageSecurityMode sets MessageSecurityMode field to given value.

### HasMessageSecurityMode

`func (o *MachineInfo) HasMessageSecurityMode() bool`

HasMessageSecurityMode returns a boolean if a field has been set.

### GetName

`func (o *MachineInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MachineInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *MachineInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *MachineInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *MachineInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *MachineInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetOperatingSystemArchitecture

`func (o *MachineInfo) GetOperatingSystemArchitecture() string`

GetOperatingSystemArchitecture returns the OperatingSystemArchitecture field if non-nil, zero value otherwise.

### GetOperatingSystemArchitectureOk

`func (o *MachineInfo) GetOperatingSystemArchitectureOk() (*string, bool)`

GetOperatingSystemArchitectureOk returns a tuple with the OperatingSystemArchitecture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystemArchitecture

`func (o *MachineInfo) SetOperatingSystemArchitecture(v string)`

SetOperatingSystemArchitecture sets OperatingSystemArchitecture field to given value.

### HasOperatingSystemArchitecture

`func (o *MachineInfo) HasOperatingSystemArchitecture() bool`

HasOperatingSystemArchitecture returns a boolean if a field has been set.

### GetPairingState

`func (o *MachineInfo) GetPairingState() string`

GetPairingState returns the PairingState field if non-nil, zero value otherwise.

### GetPairingStateOk

`func (o *MachineInfo) GetPairingStateOk() (*string, bool)`

GetPairingStateOk returns a tuple with the PairingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingState

`func (o *MachineInfo) SetPairingState(v string)`

SetPairingState sets PairingState field to given value.

### HasPairingState

`func (o *MachineInfo) HasPairingState() bool`

HasPairingState returns a boolean if a field has been set.

### GetRemoteExperienceAgentBuildNumber

`func (o *MachineInfo) GetRemoteExperienceAgentBuildNumber() string`

GetRemoteExperienceAgentBuildNumber returns the RemoteExperienceAgentBuildNumber field if non-nil, zero value otherwise.

### GetRemoteExperienceAgentBuildNumberOk

`func (o *MachineInfo) GetRemoteExperienceAgentBuildNumberOk() (*string, bool)`

GetRemoteExperienceAgentBuildNumberOk returns a tuple with the RemoteExperienceAgentBuildNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteExperienceAgentBuildNumber

`func (o *MachineInfo) SetRemoteExperienceAgentBuildNumber(v string)`

SetRemoteExperienceAgentBuildNumber sets RemoteExperienceAgentBuildNumber field to given value.

### HasRemoteExperienceAgentBuildNumber

`func (o *MachineInfo) HasRemoteExperienceAgentBuildNumber() bool`

HasRemoteExperienceAgentBuildNumber returns a boolean if a field has been set.

### GetRemoteExperienceAgentVersion

`func (o *MachineInfo) GetRemoteExperienceAgentVersion() string`

GetRemoteExperienceAgentVersion returns the RemoteExperienceAgentVersion field if non-nil, zero value otherwise.

### GetRemoteExperienceAgentVersionOk

`func (o *MachineInfo) GetRemoteExperienceAgentVersionOk() (*string, bool)`

GetRemoteExperienceAgentVersionOk returns a tuple with the RemoteExperienceAgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteExperienceAgentVersion

`func (o *MachineInfo) SetRemoteExperienceAgentVersion(v string)`

SetRemoteExperienceAgentVersion sets RemoteExperienceAgentVersion field to given value.

### HasRemoteExperienceAgentVersion

`func (o *MachineInfo) HasRemoteExperienceAgentVersion() bool`

HasRemoteExperienceAgentVersion returns a boolean if a field has been set.

### GetState

`func (o *MachineInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *MachineInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *MachineInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *MachineInfo) HasState() bool`

HasState returns a boolean if a field has been set.

### GetType

`func (o *MachineInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MachineInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MachineInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MachineInfo) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserIds

`func (o *MachineInfo) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *MachineInfo) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *MachineInfo) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *MachineInfo) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


