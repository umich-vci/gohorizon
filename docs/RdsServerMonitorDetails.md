# RDSServerMonitorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBuild** | **string** | Agent build number. | 
**AgentVersion** | **string** | Agent version. | 
**MaxSessionsCountConfigured** | Pointer to **int32** | Maximum number of sessions for RDS server as configured by administrator. Will be unset if the value is not configured. | [optional] 
**OperatingSystem** | **string** | Operating System version. * UNKNOWN: Unknown * WINDOWS_SERVER_2003: Windows Server 2003 * WINDOWS_SERVER_2008: Windows Server 2008 * WINDOWS_SERVER_2008_R2: Windows Server 2008 R2 * WINDOWS_SERVER_2012: Windows Server 2012 * WINDOWS_SERVER_2012_R2: Windows Server 2012 R2 * WINDOWS_SERVER_2016_OR_ABOVE: Windows Server 2016 or above * LINUX_SERVER_OTHER: Linux Server (other) | 
**State** | **string** | State of RDS Server. * WAIT_FOR_AGENT: Connection Server is waiting to establish communication with View Agent on a RDS Server. * AGENT_UNREACHABLE: Connection Server cannot establish communication with View Agent on a RDS Server. * AGENT_CONFIG_ERROR: The RDS Server has configuration error. Ex: Remote Desktop Service role is not enabled. * AVAILABLE: The RDS Server is powered on and ready for an active connection. * DISABLED: The RDS Server is disabled. * DISABLE_IN_PROGRESS: Disabled RDS server still has some brokered sessions. It can still accept re-connections. * PROVISIONING: The RDS Server is being provisioned. * PROVISIONING_ERROR: An error occurred during provisioning. * CUSTOMIZING: The RDS Server is being customized. * DELETING: The RDS Server is marked for deletion. It will be deleted soon. * MAINTENANCE: The RDS Server is in maintenance mode. * ERROR: An unknown error occurred in the RDS Server. * PROVISIONED: The RDS Server has been provisioned. * CONNECTED: The RDS Server is in an active session. * DISCONNECTED: The RDS Server is in an active session but is disconnected. * AGENT_ERR_STARTUP_IN_PROGRESS: The Horizon Agent has started on the virtual machine, but other required services such as the display protocol are still starting. * AGENT_ERR_DISABLED: The Horizon Agent is disabled. * AGENT_ERR_INVALID_IP: The Horizon Agent has an invalid IP. * AGENT_ERR_NEED_REBOOT: The Horizon Agent needs reboot. * AGENT_ERR_PROTOCOL_FAILURE: Protocol such as RDP or PCoIP is not enabled. * AGENT_ERR_DOMAIN_FAILURE: The RDS Server has an invalid domain. * AGENT_DRAIN_MODE: The RDS Server is configured for drain mode. * AGENT_DRAIN_UNTIL_RESTART: The RDS Server is configured a mode to drain until restart. * ALREADY_USED: The RDS Server cannot accept new sessions. * IN_PROGRESS: There is a RDS Server operation in progress. * VALIDATING: The connection server is synchronizing state information with the agent. * UNKNOWN: Could not determine the state of the RDS Server. | 

## Methods

### NewRDSServerMonitorDetails

`func NewRDSServerMonitorDetails(agentBuild string, agentVersion string, operatingSystem string, state string, ) *RDSServerMonitorDetails`

NewRDSServerMonitorDetails instantiates a new RDSServerMonitorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSServerMonitorDetailsWithDefaults

`func NewRDSServerMonitorDetailsWithDefaults() *RDSServerMonitorDetails`

NewRDSServerMonitorDetailsWithDefaults instantiates a new RDSServerMonitorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBuild

`func (o *RDSServerMonitorDetails) GetAgentBuild() string`

GetAgentBuild returns the AgentBuild field if non-nil, zero value otherwise.

### GetAgentBuildOk

`func (o *RDSServerMonitorDetails) GetAgentBuildOk() (*string, bool)`

GetAgentBuildOk returns a tuple with the AgentBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBuild

`func (o *RDSServerMonitorDetails) SetAgentBuild(v string)`

SetAgentBuild sets AgentBuild field to given value.


### GetAgentVersion

`func (o *RDSServerMonitorDetails) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *RDSServerMonitorDetails) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *RDSServerMonitorDetails) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.


### GetMaxSessionsCountConfigured

`func (o *RDSServerMonitorDetails) GetMaxSessionsCountConfigured() int32`

GetMaxSessionsCountConfigured returns the MaxSessionsCountConfigured field if non-nil, zero value otherwise.

### GetMaxSessionsCountConfiguredOk

`func (o *RDSServerMonitorDetails) GetMaxSessionsCountConfiguredOk() (*int32, bool)`

GetMaxSessionsCountConfiguredOk returns a tuple with the MaxSessionsCountConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionsCountConfigured

`func (o *RDSServerMonitorDetails) SetMaxSessionsCountConfigured(v int32)`

SetMaxSessionsCountConfigured sets MaxSessionsCountConfigured field to given value.

### HasMaxSessionsCountConfigured

`func (o *RDSServerMonitorDetails) HasMaxSessionsCountConfigured() bool`

HasMaxSessionsCountConfigured returns a boolean if a field has been set.

### GetOperatingSystem

`func (o *RDSServerMonitorDetails) GetOperatingSystem() string`

GetOperatingSystem returns the OperatingSystem field if non-nil, zero value otherwise.

### GetOperatingSystemOk

`func (o *RDSServerMonitorDetails) GetOperatingSystemOk() (*string, bool)`

GetOperatingSystemOk returns a tuple with the OperatingSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatingSystem

`func (o *RDSServerMonitorDetails) SetOperatingSystem(v string)`

SetOperatingSystem sets OperatingSystem field to given value.


### GetState

`func (o *RDSServerMonitorDetails) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *RDSServerMonitorDetails) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *RDSServerMonitorDetails) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


