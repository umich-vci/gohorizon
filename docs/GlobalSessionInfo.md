# GlobalSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentVersion** | Pointer to **string** | Version of the agent for the session. | [optional] 
**ApplicationNames** | Pointer to **[]string** | Names of the applications launched in the session. &lt;br&gt;It will only be set when session_type is APPLICATION. | [optional] 
**BrokerUserId** | Pointer to **string** | SID for the broker user associated with the session.  It will be unset for non-broker sessions. | [optional] 
**BrokeredRemotely** | Pointer to **bool** | Indicates whether the session is brokered from a remote pod. &lt;br&gt;This will be set only if the Horizon View agent where the session resides is version 6.0 or later. | [optional] 
**BrokeringPodId** | Pointer to **string** | ID of the pod that brokered the session. | [optional] 
**ClientData** | Pointer to [**GlobalSessionClientData**](GlobalSessionClientData.md) |  | [optional] 
**DesktopPoolId** | Pointer to **string** | ID of the desktop pool for the desktop session.  This will be unset if the session is not brokered through a desktop pool, such as for direct console access.  This will be unset if the session is hosted by remote pod. | [optional] 
**DesktopPoolOrFarmName** | Pointer to **string** | Display name of the desktop pool or name of the farm for the session.  This will be unset if session is not brokered through a desktop pool or a farm, such as for direct console access. | [optional] 
**DesktopPoolOrFarmSource** | Pointer to **string** | Source of the desktop pool or farm for the session.  This will be unset if session is not brokered through a desktop pool or a farm, such as for direct console access. * INSTANT_CLONE: The Desktop Pool uses instant clone technology for provisioning the machines. Applicable for AUTOMATED type desktop pools. * LINKED_CLONE: The Desktop Pool uses linked clone technology for provisioning the machines. Applicable for AUTOMATED type desktop pools. * VIRTUAL_CENTER: The Desktop Pool uses Virtual Center as source for provisioning the machines. Applicable for AUTOMATED and MANUAL type desktop pools. * RDS: The Desktop Pool is backed by Farm. The Farm used in this Desktop Pool can be of any Source. * UNMANAGED: The Desktop Pool holds the non-vCenter source machines that includes physical computers, blade PCs and non-vCenter servers. Applicable for MANUAL type desktop pools. | [optional] 
**DesktopPoolOrFarmType** | Pointer to **string** | Type of the desktop pool or farm for the session.  This will be unset if session is not brokered through a desktop pool or a farm, such as for direct console access. * AUTOMATED: Automated Desktop Pool. * MANUAL: Manual Desktop Pool. * RDS: RDS Desktop Pool. | [optional] 
**DisconnectedTime** | Pointer to **int64** | Epoch time in milliseconds, when the session was last disconnected.  This will be unset if the machine on which the session resides has an error state, or if the session has never been disconnected. | [optional] 
**FarmId** | Pointer to **string** | ID of the farm for the RDS desktop or application session.  This will be unset if the session is not brokered through a farm, such as for direct console access.  This will be unset if the session is hosted by remote pod. | [optional] 
**FederatedAccessGroupIds** | Pointer to **[]string** | IDs of the federated access groups associated with the session. &lt;br&gt;This represents the federated access groups associated with the global desktop entitlement or global application entitlements used to launch the session. | [optional] 
**ForeverSession** | Pointer to **bool** | Indicates whether the application session will continue to run indefinitely on reaching global idle timeout or max session timeout. | [optional] 
**GlobalApplicationEntitlementIds** | Pointer to **[]string** | IDs of the global application entitlements used to launch applications in the session.  Either this or global_desktop_entitlement_id may be set, but not both. | [optional] 
**GlobalDesktopEntitlementId** | Pointer to **string** | ID of the global desktop entitlement used to launch the session.  Either this or global_application_entitlement_ids may be set, but not both. | [optional] 
**Id** | Pointer to **string** | Unique ID representing a session. | [optional] 
**IdleDuration** | Pointer to **int64** | Idle time duration in minutes, indicating how long the end user of the session has been idle for. | [optional] 
**LastSessionDurationMs** | Pointer to **int64** | Duration of the last connection period of the session in milliseconds.  If the session is currently connected, this is the duration that the session has been in connected state.  If the session is currently disconnected, this is the duration of its previous connection period. This will be unset on error. | [optional] 
**LocalAccessGroupId** | Pointer to **string** | ID of the local access group associated with the session.  For a non-RDS desktop session, this is the access group ID of the desktop pool.  For an RDS desktop session, this is access group ID of the farm of the RDS desktop pool.  For an application session, this is the access group ID of the farm of the application pool.  This will be unset if the session is hosted by remote pod. | [optional] 
**MachineId** | Pointer to **string** | ID of the machine for the session.  This will be unset for RDS desktop or application sessions.  If desktop_pool_id is unset, it is the id of registered unmanaged physical machine.  This will be unset if the session is hosted by remote pod. | [optional] 
**MachineOrRdsServerDnsName** | Pointer to **string** | DNS name of the machine or RDS server for the session. | [optional] 
**PodId** | Pointer to **string** | ID of the pod that provided the resource for the session. | [optional] 
**RdsServerId** | Pointer to **string** | ID of the RDS server for the RDS desktop or application session. &lt;br&gt;This will be unset if the session is hosted by remote pod. | [optional] 
**ResourcedRemotely** | Pointer to **bool** | Indicates whether the session is running on a remote pod resource. | [optional] 
**SecurityGatewayData** | Pointer to [**GlobalSessionSecurityGatewayData**](GlobalSessionSecurityGatewayData.md) |  | [optional] 
**SessionProtocol** | Pointer to **string** | Protocol for session. This will be unset for disconnected sessions. * PCOIP: Display protocol is PCoIP. * RDP: Display protocol is RDP. * BLAST: Display protocol is BLAST. * CONSOLE: Display protocol is console. * UNKNOWN: Display protocol is unknown. | [optional] 
**SessionState** | Pointer to **string** | State of session. * CONNECTED: Session is connected * DISCONNECTED: Session is disconnected * PENDING: Session is pending | [optional] 
**SessionType** | Pointer to **string** | Type of session. * DESKTOP: Desktop or RDS desktop session. * APPLICATION: Application session. | [optional] 
**SiteId** | Pointer to **string** | ID of the site where the pod that resourced the session belongs. | [optional] 
**StartTime** | Pointer to **int64** | Epoch time in milliseconds when this session was originally logged in.  The lifecycle of a session begins at login and ends at logout, with any number of connect and disconnect occurrences between.  The first connection time will be shortly after this time. | [optional] 
**Unauthenticated** | Pointer to **bool** | Indicates whether the session belongs to unauthenticated access user. | [optional] 
**UserId** | Pointer to **string** |  SID of the user logged into the session.  It may not match the broker user id for non-SSO scenarios. | [optional] 

## Methods

### NewGlobalSessionInfo

`func NewGlobalSessionInfo() *GlobalSessionInfo`

NewGlobalSessionInfo instantiates a new GlobalSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSessionInfoWithDefaults

`func NewGlobalSessionInfoWithDefaults() *GlobalSessionInfo`

NewGlobalSessionInfoWithDefaults instantiates a new GlobalSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentVersion

`func (o *GlobalSessionInfo) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *GlobalSessionInfo) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *GlobalSessionInfo) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *GlobalSessionInfo) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetApplicationNames

`func (o *GlobalSessionInfo) GetApplicationNames() []string`

GetApplicationNames returns the ApplicationNames field if non-nil, zero value otherwise.

### GetApplicationNamesOk

`func (o *GlobalSessionInfo) GetApplicationNamesOk() (*[]string, bool)`

GetApplicationNamesOk returns a tuple with the ApplicationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNames

`func (o *GlobalSessionInfo) SetApplicationNames(v []string)`

SetApplicationNames sets ApplicationNames field to given value.

### HasApplicationNames

`func (o *GlobalSessionInfo) HasApplicationNames() bool`

HasApplicationNames returns a boolean if a field has been set.

### GetBrokerUserId

`func (o *GlobalSessionInfo) GetBrokerUserId() string`

GetBrokerUserId returns the BrokerUserId field if non-nil, zero value otherwise.

### GetBrokerUserIdOk

`func (o *GlobalSessionInfo) GetBrokerUserIdOk() (*string, bool)`

GetBrokerUserIdOk returns a tuple with the BrokerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerUserId

`func (o *GlobalSessionInfo) SetBrokerUserId(v string)`

SetBrokerUserId sets BrokerUserId field to given value.

### HasBrokerUserId

`func (o *GlobalSessionInfo) HasBrokerUserId() bool`

HasBrokerUserId returns a boolean if a field has been set.

### GetBrokeredRemotely

`func (o *GlobalSessionInfo) GetBrokeredRemotely() bool`

GetBrokeredRemotely returns the BrokeredRemotely field if non-nil, zero value otherwise.

### GetBrokeredRemotelyOk

`func (o *GlobalSessionInfo) GetBrokeredRemotelyOk() (*bool, bool)`

GetBrokeredRemotelyOk returns a tuple with the BrokeredRemotely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokeredRemotely

`func (o *GlobalSessionInfo) SetBrokeredRemotely(v bool)`

SetBrokeredRemotely sets BrokeredRemotely field to given value.

### HasBrokeredRemotely

`func (o *GlobalSessionInfo) HasBrokeredRemotely() bool`

HasBrokeredRemotely returns a boolean if a field has been set.

### GetBrokeringPodId

`func (o *GlobalSessionInfo) GetBrokeringPodId() string`

GetBrokeringPodId returns the BrokeringPodId field if non-nil, zero value otherwise.

### GetBrokeringPodIdOk

`func (o *GlobalSessionInfo) GetBrokeringPodIdOk() (*string, bool)`

GetBrokeringPodIdOk returns a tuple with the BrokeringPodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokeringPodId

`func (o *GlobalSessionInfo) SetBrokeringPodId(v string)`

SetBrokeringPodId sets BrokeringPodId field to given value.

### HasBrokeringPodId

`func (o *GlobalSessionInfo) HasBrokeringPodId() bool`

HasBrokeringPodId returns a boolean if a field has been set.

### GetClientData

`func (o *GlobalSessionInfo) GetClientData() GlobalSessionClientData`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *GlobalSessionInfo) GetClientDataOk() (*GlobalSessionClientData, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *GlobalSessionInfo) SetClientData(v GlobalSessionClientData)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *GlobalSessionInfo) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetDesktopPoolId

`func (o *GlobalSessionInfo) GetDesktopPoolId() string`

GetDesktopPoolId returns the DesktopPoolId field if non-nil, zero value otherwise.

### GetDesktopPoolIdOk

`func (o *GlobalSessionInfo) GetDesktopPoolIdOk() (*string, bool)`

GetDesktopPoolIdOk returns a tuple with the DesktopPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolId

`func (o *GlobalSessionInfo) SetDesktopPoolId(v string)`

SetDesktopPoolId sets DesktopPoolId field to given value.

### HasDesktopPoolId

`func (o *GlobalSessionInfo) HasDesktopPoolId() bool`

HasDesktopPoolId returns a boolean if a field has been set.

### GetDesktopPoolOrFarmName

`func (o *GlobalSessionInfo) GetDesktopPoolOrFarmName() string`

GetDesktopPoolOrFarmName returns the DesktopPoolOrFarmName field if non-nil, zero value otherwise.

### GetDesktopPoolOrFarmNameOk

`func (o *GlobalSessionInfo) GetDesktopPoolOrFarmNameOk() (*string, bool)`

GetDesktopPoolOrFarmNameOk returns a tuple with the DesktopPoolOrFarmName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolOrFarmName

`func (o *GlobalSessionInfo) SetDesktopPoolOrFarmName(v string)`

SetDesktopPoolOrFarmName sets DesktopPoolOrFarmName field to given value.

### HasDesktopPoolOrFarmName

`func (o *GlobalSessionInfo) HasDesktopPoolOrFarmName() bool`

HasDesktopPoolOrFarmName returns a boolean if a field has been set.

### GetDesktopPoolOrFarmSource

`func (o *GlobalSessionInfo) GetDesktopPoolOrFarmSource() string`

GetDesktopPoolOrFarmSource returns the DesktopPoolOrFarmSource field if non-nil, zero value otherwise.

### GetDesktopPoolOrFarmSourceOk

`func (o *GlobalSessionInfo) GetDesktopPoolOrFarmSourceOk() (*string, bool)`

GetDesktopPoolOrFarmSourceOk returns a tuple with the DesktopPoolOrFarmSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolOrFarmSource

`func (o *GlobalSessionInfo) SetDesktopPoolOrFarmSource(v string)`

SetDesktopPoolOrFarmSource sets DesktopPoolOrFarmSource field to given value.

### HasDesktopPoolOrFarmSource

`func (o *GlobalSessionInfo) HasDesktopPoolOrFarmSource() bool`

HasDesktopPoolOrFarmSource returns a boolean if a field has been set.

### GetDesktopPoolOrFarmType

`func (o *GlobalSessionInfo) GetDesktopPoolOrFarmType() string`

GetDesktopPoolOrFarmType returns the DesktopPoolOrFarmType field if non-nil, zero value otherwise.

### GetDesktopPoolOrFarmTypeOk

`func (o *GlobalSessionInfo) GetDesktopPoolOrFarmTypeOk() (*string, bool)`

GetDesktopPoolOrFarmTypeOk returns a tuple with the DesktopPoolOrFarmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolOrFarmType

`func (o *GlobalSessionInfo) SetDesktopPoolOrFarmType(v string)`

SetDesktopPoolOrFarmType sets DesktopPoolOrFarmType field to given value.

### HasDesktopPoolOrFarmType

`func (o *GlobalSessionInfo) HasDesktopPoolOrFarmType() bool`

HasDesktopPoolOrFarmType returns a boolean if a field has been set.

### GetDisconnectedTime

`func (o *GlobalSessionInfo) GetDisconnectedTime() int64`

GetDisconnectedTime returns the DisconnectedTime field if non-nil, zero value otherwise.

### GetDisconnectedTimeOk

`func (o *GlobalSessionInfo) GetDisconnectedTimeOk() (*int64, bool)`

GetDisconnectedTimeOk returns a tuple with the DisconnectedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedTime

`func (o *GlobalSessionInfo) SetDisconnectedTime(v int64)`

SetDisconnectedTime sets DisconnectedTime field to given value.

### HasDisconnectedTime

`func (o *GlobalSessionInfo) HasDisconnectedTime() bool`

HasDisconnectedTime returns a boolean if a field has been set.

### GetFarmId

`func (o *GlobalSessionInfo) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *GlobalSessionInfo) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *GlobalSessionInfo) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *GlobalSessionInfo) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetFederatedAccessGroupIds

`func (o *GlobalSessionInfo) GetFederatedAccessGroupIds() []string`

GetFederatedAccessGroupIds returns the FederatedAccessGroupIds field if non-nil, zero value otherwise.

### GetFederatedAccessGroupIdsOk

`func (o *GlobalSessionInfo) GetFederatedAccessGroupIdsOk() (*[]string, bool)`

GetFederatedAccessGroupIdsOk returns a tuple with the FederatedAccessGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederatedAccessGroupIds

`func (o *GlobalSessionInfo) SetFederatedAccessGroupIds(v []string)`

SetFederatedAccessGroupIds sets FederatedAccessGroupIds field to given value.

### HasFederatedAccessGroupIds

`func (o *GlobalSessionInfo) HasFederatedAccessGroupIds() bool`

HasFederatedAccessGroupIds returns a boolean if a field has been set.

### GetForeverSession

`func (o *GlobalSessionInfo) GetForeverSession() bool`

GetForeverSession returns the ForeverSession field if non-nil, zero value otherwise.

### GetForeverSessionOk

`func (o *GlobalSessionInfo) GetForeverSessionOk() (*bool, bool)`

GetForeverSessionOk returns a tuple with the ForeverSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForeverSession

`func (o *GlobalSessionInfo) SetForeverSession(v bool)`

SetForeverSession sets ForeverSession field to given value.

### HasForeverSession

`func (o *GlobalSessionInfo) HasForeverSession() bool`

HasForeverSession returns a boolean if a field has been set.

### GetGlobalApplicationEntitlementIds

`func (o *GlobalSessionInfo) GetGlobalApplicationEntitlementIds() []string`

GetGlobalApplicationEntitlementIds returns the GlobalApplicationEntitlementIds field if non-nil, zero value otherwise.

### GetGlobalApplicationEntitlementIdsOk

`func (o *GlobalSessionInfo) GetGlobalApplicationEntitlementIdsOk() (*[]string, bool)`

GetGlobalApplicationEntitlementIdsOk returns a tuple with the GlobalApplicationEntitlementIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalApplicationEntitlementIds

`func (o *GlobalSessionInfo) SetGlobalApplicationEntitlementIds(v []string)`

SetGlobalApplicationEntitlementIds sets GlobalApplicationEntitlementIds field to given value.

### HasGlobalApplicationEntitlementIds

`func (o *GlobalSessionInfo) HasGlobalApplicationEntitlementIds() bool`

HasGlobalApplicationEntitlementIds returns a boolean if a field has been set.

### GetGlobalDesktopEntitlementId

`func (o *GlobalSessionInfo) GetGlobalDesktopEntitlementId() string`

GetGlobalDesktopEntitlementId returns the GlobalDesktopEntitlementId field if non-nil, zero value otherwise.

### GetGlobalDesktopEntitlementIdOk

`func (o *GlobalSessionInfo) GetGlobalDesktopEntitlementIdOk() (*string, bool)`

GetGlobalDesktopEntitlementIdOk returns a tuple with the GlobalDesktopEntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalDesktopEntitlementId

`func (o *GlobalSessionInfo) SetGlobalDesktopEntitlementId(v string)`

SetGlobalDesktopEntitlementId sets GlobalDesktopEntitlementId field to given value.

### HasGlobalDesktopEntitlementId

`func (o *GlobalSessionInfo) HasGlobalDesktopEntitlementId() bool`

HasGlobalDesktopEntitlementId returns a boolean if a field has been set.

### GetId

`func (o *GlobalSessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalSessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalSessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalSessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdleDuration

`func (o *GlobalSessionInfo) GetIdleDuration() int64`

GetIdleDuration returns the IdleDuration field if non-nil, zero value otherwise.

### GetIdleDurationOk

`func (o *GlobalSessionInfo) GetIdleDurationOk() (*int64, bool)`

GetIdleDurationOk returns a tuple with the IdleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleDuration

`func (o *GlobalSessionInfo) SetIdleDuration(v int64)`

SetIdleDuration sets IdleDuration field to given value.

### HasIdleDuration

`func (o *GlobalSessionInfo) HasIdleDuration() bool`

HasIdleDuration returns a boolean if a field has been set.

### GetLastSessionDurationMs

`func (o *GlobalSessionInfo) GetLastSessionDurationMs() int64`

GetLastSessionDurationMs returns the LastSessionDurationMs field if non-nil, zero value otherwise.

### GetLastSessionDurationMsOk

`func (o *GlobalSessionInfo) GetLastSessionDurationMsOk() (*int64, bool)`

GetLastSessionDurationMsOk returns a tuple with the LastSessionDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSessionDurationMs

`func (o *GlobalSessionInfo) SetLastSessionDurationMs(v int64)`

SetLastSessionDurationMs sets LastSessionDurationMs field to given value.

### HasLastSessionDurationMs

`func (o *GlobalSessionInfo) HasLastSessionDurationMs() bool`

HasLastSessionDurationMs returns a boolean if a field has been set.

### GetLocalAccessGroupId

`func (o *GlobalSessionInfo) GetLocalAccessGroupId() string`

GetLocalAccessGroupId returns the LocalAccessGroupId field if non-nil, zero value otherwise.

### GetLocalAccessGroupIdOk

`func (o *GlobalSessionInfo) GetLocalAccessGroupIdOk() (*string, bool)`

GetLocalAccessGroupIdOk returns a tuple with the LocalAccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalAccessGroupId

`func (o *GlobalSessionInfo) SetLocalAccessGroupId(v string)`

SetLocalAccessGroupId sets LocalAccessGroupId field to given value.

### HasLocalAccessGroupId

`func (o *GlobalSessionInfo) HasLocalAccessGroupId() bool`

HasLocalAccessGroupId returns a boolean if a field has been set.

### GetMachineId

`func (o *GlobalSessionInfo) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *GlobalSessionInfo) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *GlobalSessionInfo) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *GlobalSessionInfo) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetMachineOrRdsServerDnsName

`func (o *GlobalSessionInfo) GetMachineOrRdsServerDnsName() string`

GetMachineOrRdsServerDnsName returns the MachineOrRdsServerDnsName field if non-nil, zero value otherwise.

### GetMachineOrRdsServerDnsNameOk

`func (o *GlobalSessionInfo) GetMachineOrRdsServerDnsNameOk() (*string, bool)`

GetMachineOrRdsServerDnsNameOk returns a tuple with the MachineOrRdsServerDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineOrRdsServerDnsName

`func (o *GlobalSessionInfo) SetMachineOrRdsServerDnsName(v string)`

SetMachineOrRdsServerDnsName sets MachineOrRdsServerDnsName field to given value.

### HasMachineOrRdsServerDnsName

`func (o *GlobalSessionInfo) HasMachineOrRdsServerDnsName() bool`

HasMachineOrRdsServerDnsName returns a boolean if a field has been set.

### GetPodId

`func (o *GlobalSessionInfo) GetPodId() string`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *GlobalSessionInfo) GetPodIdOk() (*string, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *GlobalSessionInfo) SetPodId(v string)`

SetPodId sets PodId field to given value.

### HasPodId

`func (o *GlobalSessionInfo) HasPodId() bool`

HasPodId returns a boolean if a field has been set.

### GetRdsServerId

`func (o *GlobalSessionInfo) GetRdsServerId() string`

GetRdsServerId returns the RdsServerId field if non-nil, zero value otherwise.

### GetRdsServerIdOk

`func (o *GlobalSessionInfo) GetRdsServerIdOk() (*string, bool)`

GetRdsServerIdOk returns a tuple with the RdsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsServerId

`func (o *GlobalSessionInfo) SetRdsServerId(v string)`

SetRdsServerId sets RdsServerId field to given value.

### HasRdsServerId

`func (o *GlobalSessionInfo) HasRdsServerId() bool`

HasRdsServerId returns a boolean if a field has been set.

### GetResourcedRemotely

`func (o *GlobalSessionInfo) GetResourcedRemotely() bool`

GetResourcedRemotely returns the ResourcedRemotely field if non-nil, zero value otherwise.

### GetResourcedRemotelyOk

`func (o *GlobalSessionInfo) GetResourcedRemotelyOk() (*bool, bool)`

GetResourcedRemotelyOk returns a tuple with the ResourcedRemotely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcedRemotely

`func (o *GlobalSessionInfo) SetResourcedRemotely(v bool)`

SetResourcedRemotely sets ResourcedRemotely field to given value.

### HasResourcedRemotely

`func (o *GlobalSessionInfo) HasResourcedRemotely() bool`

HasResourcedRemotely returns a boolean if a field has been set.

### GetSecurityGatewayData

`func (o *GlobalSessionInfo) GetSecurityGatewayData() GlobalSessionSecurityGatewayData`

GetSecurityGatewayData returns the SecurityGatewayData field if non-nil, zero value otherwise.

### GetSecurityGatewayDataOk

`func (o *GlobalSessionInfo) GetSecurityGatewayDataOk() (*GlobalSessionSecurityGatewayData, bool)`

GetSecurityGatewayDataOk returns a tuple with the SecurityGatewayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGatewayData

`func (o *GlobalSessionInfo) SetSecurityGatewayData(v GlobalSessionSecurityGatewayData)`

SetSecurityGatewayData sets SecurityGatewayData field to given value.

### HasSecurityGatewayData

`func (o *GlobalSessionInfo) HasSecurityGatewayData() bool`

HasSecurityGatewayData returns a boolean if a field has been set.

### GetSessionProtocol

`func (o *GlobalSessionInfo) GetSessionProtocol() string`

GetSessionProtocol returns the SessionProtocol field if non-nil, zero value otherwise.

### GetSessionProtocolOk

`func (o *GlobalSessionInfo) GetSessionProtocolOk() (*string, bool)`

GetSessionProtocolOk returns a tuple with the SessionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocol

`func (o *GlobalSessionInfo) SetSessionProtocol(v string)`

SetSessionProtocol sets SessionProtocol field to given value.

### HasSessionProtocol

`func (o *GlobalSessionInfo) HasSessionProtocol() bool`

HasSessionProtocol returns a boolean if a field has been set.

### GetSessionState

`func (o *GlobalSessionInfo) GetSessionState() string`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *GlobalSessionInfo) GetSessionStateOk() (*string, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *GlobalSessionInfo) SetSessionState(v string)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *GlobalSessionInfo) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetSessionType

`func (o *GlobalSessionInfo) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *GlobalSessionInfo) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *GlobalSessionInfo) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *GlobalSessionInfo) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetSiteId

`func (o *GlobalSessionInfo) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *GlobalSessionInfo) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *GlobalSessionInfo) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *GlobalSessionInfo) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStartTime

`func (o *GlobalSessionInfo) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *GlobalSessionInfo) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *GlobalSessionInfo) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *GlobalSessionInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUnauthenticated

`func (o *GlobalSessionInfo) GetUnauthenticated() bool`

GetUnauthenticated returns the Unauthenticated field if non-nil, zero value otherwise.

### GetUnauthenticatedOk

`func (o *GlobalSessionInfo) GetUnauthenticatedOk() (*bool, bool)`

GetUnauthenticatedOk returns a tuple with the Unauthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticated

`func (o *GlobalSessionInfo) SetUnauthenticated(v bool)`

SetUnauthenticated sets Unauthenticated field to given value.

### HasUnauthenticated

`func (o *GlobalSessionInfo) HasUnauthenticated() bool`

HasUnauthenticated returns a boolean if a field has been set.

### GetUserId

`func (o *GlobalSessionInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GlobalSessionInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GlobalSessionInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GlobalSessionInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


