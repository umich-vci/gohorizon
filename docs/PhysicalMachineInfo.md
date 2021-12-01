# PhysicalMachineInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**DnsName** | Pointer to **string** | DNS name of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39;, &#39;EndsWith&#39; and &#39;Contains&#39;. | [optional] 
**Id** | Pointer to **string** | Unique ID representing machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**MessageSecurityEnhancedModeSupported** | Pointer to **bool** | Indicates whether ENHANCED message security mode is currently supported by this machine. | [optional] 
**MessageSecurityMode** | Pointer to **string** | The current JMS message security mode used by this machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * DISABLED: Message security mode is disabled. * MIXED: Message security mode is enabled but not enforced. * ENABLED: Message security mode is enabled. Unsigned messages are rejected by Horizon components. * ENHANCED: Message Security mode is Enhanced. Message signing and validation is performed based on the current Security Level and desktop Message Security mode. | [optional] 
**Name** | Pointer to **string** | Name of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;StartsWith&#39; and &#39;Contains&#39;. | [optional] 
**OperatingSystem** | Pointer to **string** | The machine operating system. * UNKNOWN: Unknown * WINDOWS_XP: Windows XP * WINDOWS_VISTA: Windows Vista * WINDOWS_7: Windows 7 * WINDOWS_8: Windows 8 * WINDOWS_10: Windows 10 * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_OTHER: Linux (other) * LINUX_SERVER_OTHER: Linux server (other) * LINUX_UBUNTU: Linux (Ubuntu) * LINUX_RHEL: Linux (Red Hat Enterprise) * LINUX_SUSE: Linux (Suse) * LINUX_CENTOS: Linux (CentOS) | [optional] 
**State** | Pointer to **string** | The state of the machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * AGENT_UNREACHABLE: Horizon Connection Server cannot establish communication with Horizon Agent on the machine. * UNASSIGNED_USER_CONNECTED: A user other than the assigned user is logged in to the machine in a dedicated desktop pool. * CONNECTED: The machine is in an active session and has an active connection to a Horizon client. * UNASSIGNED_USER_DISCONNECTED: A user other than the assigned user is logged in and disconnected from the machine in a dedicated desktop pool. * DISCONNECTED: The machine is in an active session, but it is disconnected from the Horizon client. * AGENT_ERROR_STARTUP_IN_PROGRESS: Horizon Agent has started on the machine, but other required services such as the display protocol are still starting. * AGENT_ERROR_DISABLED: Horizon Agent is disabled. * AGENT_ERROR_INVALID_IP: Horizon Agent has an invalid IP address. * AGENT_ERROR_NEEDS_REBOOT: Horizon Agent needs reboot. * AGENT_ERROR_PROTOCOL_FAILURE: Protocol such as BLAST, RDP or PCoIP is not enabled. * AGENT_CONFIG_ERROR: The Remote Desktop Services role is not enabled on the windows server. * AVAILABLE: The machine is powered on and ready for active connections. * VALIDATING: The Horizon Connection Server is synchronizing state information with the agent. * DISABLED: The machine is disabled. * UNKNOWN: Could not determine the state of the machine. | [optional] 

## Methods

### NewPhysicalMachineInfo

`func NewPhysicalMachineInfo() *PhysicalMachineInfo`

NewPhysicalMachineInfo instantiates a new PhysicalMachineInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhysicalMachineInfoWithDefaults

`func NewPhysicalMachineInfoWithDefaults() *PhysicalMachineInfo`

NewPhysicalMachineInfoWithDefaults instantiates a new PhysicalMachineInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PhysicalMachineInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PhysicalMachineInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PhysicalMachineInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PhysicalMachineInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDnsName

`func (o *PhysicalMachineInfo) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *PhysicalMachineInfo) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *PhysicalMachineInfo) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.

### HasDnsName

`func (o *PhysicalMachineInfo) HasDnsName() bool`

HasDnsName returns a boolean if a field has been set.

### GetId

`func (o *PhysicalMachineInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhysicalMachineInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhysicalMachineInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhysicalMachineInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessageSecurityEnhancedModeSupported

`func (o *PhysicalMachineInfo) GetMessageSecurityEnhancedModeSupported() bool`

GetMessageSecurityEnhancedModeSupported returns the MessageSecurityEnhancedModeSupported field if non-nil, zero value otherwise.

### GetMessageSecurityEnhancedModeSupportedOk

`func (o *PhysicalMachineInfo) GetMessageSecurityEnhancedModeSupportedOk() (*bool, bool)`

GetMessageSecurityEnhancedModeSupportedOk returns a tuple with the MessageSecurityEnhancedModeSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityEnhancedModeSupported

`func (o *PhysicalMachineInfo) SetMessageSecurityEnhancedModeSupported(v bool)`

SetMessageSecurityEnhancedModeSupported sets MessageSecurityEnhancedModeSupported field to given value.

### HasMessageSecurityEnhancedModeSupported

`func (o *PhysicalMachineInfo) HasMessageSecurityEnhancedModeSupported() bool`

HasMessageSecurityEnhancedModeSupported returns a boolean if a field has been set.

### GetMessageSecurityMode

`func (o *PhysicalMachineInfo) GetMessageSecurityMode() string`

GetMessageSecurityMode returns the MessageSecurityMode field if non-nil, zero value otherwise.

### GetMessageSecurityModeOk

`func (o *PhysicalMachineInfo) GetMessageSecurityModeOk() (*string, bool)`

GetMessageSecurityModeOk returns a tuple with the MessageSecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityMode

`func (o *PhysicalMachineInfo) SetMessageSecurityMode(v string)`

SetMessageSecurityMode sets MessageSecurityMode field to given value.

### HasMessageSecurityMode

`func (o *PhysicalMachineInfo) HasMessageSecurityMode() bool`

HasMessageSecurityMode returns a boolean if a field has been set.

### GetName

`func (o *PhysicalMachineInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhysicalMachineInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhysicalMachineInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PhysicalMachineInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *PhysicalMachineInfo) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *PhysicalMachineInfo) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *PhysicalMachineInfo) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.

### HasOperatingSystem

`func (o *PhysicalMachineInfo) HasOperatingSystem() bool`

HasOperatingSystem returns a boolean if a field has been set.

### GetState

`func (o *PhysicalMachineInfo) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *PhysicalMachineInfo) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *PhysicalMachineInfo) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *PhysicalMachineInfo) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


