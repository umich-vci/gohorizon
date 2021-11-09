# SessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessGroupId** | Pointer to **string** | Access group id associated with the session.  For a non-RDS desktop session, this is the desktop pool&#39;s access group id.  For an RDS desktop session, this is the RDS desktop pool&#39;s farm&#39;s access group id.  For an application session, this is the application&#39;s farm&#39;s access group id.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**AgentVersion** | Pointer to **string** | Version of agent This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**ApplicationNames** | Pointer to **[]string** | Names of the applications launched in the session.  It will be only set when session_type is APPLICATION.&lt;br&gt;Supported Filters : &#39;Contains&#39;. | [optional] 
**BrokerUserId** | Pointer to **string** | User SID for the broker user associated with the session.  It will be unset for non-broker sessions.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**BrokeredRemotely** | Pointer to **bool** | Indicates whether the session is brokered from a remote pod.  It is set only if the Horizon View agent where the session resides is version 6.0 or later.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**ClientData** | Pointer to [**ClientData**](ClientData.md) |  | [optional] 
**DesktopPoolId** | Pointer to **string** | Unique desktop pool id for the session.  This is unset if the session is not brokered through a desktop, such as for direct console access.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**DisconnectedTime** | Pointer to **int64** | Epoch time in milli seconds, when the session was last disconnected.  This will be unset if the session&#39;s machine has an error state, or if the session has never been disconnected.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**FarmId** | Pointer to **string** | Unique farm id for this RDS desktop or application session. This is unset if the session is not brokered through a farm, such as for application sessions or direct console access.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **string** | Unique id representing a session.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**IdleDuration** | Pointer to **int64** | Idle time duration in minutes, indicating how long the end user of the session has been idle for. This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**LastSessionDurationMs** | Pointer to **int64** | Duration of the last connection period of the session in milli seconds. If the session is currently connected, this is the duration that the session has been in connected state. If the session is currently disconnected, this is the duration of its previous connection period. This will be unset on error.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**MachineId** | Pointer to **string** | Unique machine id for the session.  This is unset for RDS Desktop or application sessions. If desktop pool id is unset, it is the id of registered un-managed physical machine.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**RdsServerId** | Pointer to **string** | Unique RDS server id for the RDS desktop or application session. This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**ResourcedRemotely** | Pointer to **bool** | Indicates whether the session is running on remote pod resource.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**SecurityGatewayData** | Pointer to [**SecurityGatewayData**](SecurityGatewayData.md) |  | [optional] 
**SessionProtocol** | Pointer to **string** | Protocol for the session.  It will be unset for disconnected sessions.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * PCOIP: Display protocol is PCoIP. * RDP: Display protocol is RDP. * BLAST: Display protocol is BLAST. * CONSOLE: Display protocol is console. * UNKNOWN: Display protocol is unknown. | [optional] 
**SessionState** | Pointer to **string** | State of session.&lt;br&gt;Supported Filters : &#39;Equals&#39;, &#39;NotEquals&#39;. * CONNECTED: Session is connected * DISCONNECTED: Session is disconnected * PENDING: Session is pending | [optional] 
**SessionType** | Pointer to **string** | Type of session.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * DESKTOP: Desktop or RDS desktop session. * APPLICATION: Application session. | [optional] 
**StartTime** | Pointer to **int64** | Epoch time in milli seconds, when the session was originally logged in.  The lifecycle of a session begins at login and ends at logout, with any number of connect and disconnect occurrences in between. The first connection time will be shortly after this time.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**Unauthenticated** | Pointer to **bool** | Indicates whether the session is of unauthenticated access user.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**UserId** | Pointer to **string** | Unique SID of the user logged into the session.  It may not match the broker user id for non-SSO scenarios.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 

## Methods

### NewSessionInfo

`func NewSessionInfo() *SessionInfo`

NewSessionInfo instantiates a new SessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionInfoWithDefaults

`func NewSessionInfoWithDefaults() *SessionInfo`

NewSessionInfoWithDefaults instantiates a new SessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessGroupId

`func (o *SessionInfo) GetAccessGroupId() string`

GetAccessGroupId returns the AccessGroupId field if non-nil, zero value otherwise.

### GetAccessGroupIdOk

`func (o *SessionInfo) GetAccessGroupIdOk() (*string, bool)`

GetAccessGroupIdOk returns a tuple with the AccessGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessGroupId

`func (o *SessionInfo) SetAccessGroupId(v string)`

SetAccessGroupId sets AccessGroupId field to given value.

### HasAccessGroupId

`func (o *SessionInfo) HasAccessGroupId() bool`

HasAccessGroupId returns a boolean if a field has been set.

### GetAgentVersion

`func (o *SessionInfo) GetAgentVersion() string`

GetAgentVersion returns the AgentVersion field if non-nil, zero value otherwise.

### GetAgentVersionOk

`func (o *SessionInfo) GetAgentVersionOk() (*string, bool)`

GetAgentVersionOk returns a tuple with the AgentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentVersion

`func (o *SessionInfo) SetAgentVersion(v string)`

SetAgentVersion sets AgentVersion field to given value.

### HasAgentVersion

`func (o *SessionInfo) HasAgentVersion() bool`

HasAgentVersion returns a boolean if a field has been set.

### GetApplicationNames

`func (o *SessionInfo) GetApplicationNames() []string`

GetApplicationNames returns the ApplicationNames field if non-nil, zero value otherwise.

### GetApplicationNamesOk

`func (o *SessionInfo) GetApplicationNamesOk() (*[]string, bool)`

GetApplicationNamesOk returns a tuple with the ApplicationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationNames

`func (o *SessionInfo) SetApplicationNames(v []string)`

SetApplicationNames sets ApplicationNames field to given value.

### HasApplicationNames

`func (o *SessionInfo) HasApplicationNames() bool`

HasApplicationNames returns a boolean if a field has been set.

### GetBrokerUserId

`func (o *SessionInfo) GetBrokerUserId() string`

GetBrokerUserId returns the BrokerUserId field if non-nil, zero value otherwise.

### GetBrokerUserIdOk

`func (o *SessionInfo) GetBrokerUserIdOk() (*string, bool)`

GetBrokerUserIdOk returns a tuple with the BrokerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokerUserId

`func (o *SessionInfo) SetBrokerUserId(v string)`

SetBrokerUserId sets BrokerUserId field to given value.

### HasBrokerUserId

`func (o *SessionInfo) HasBrokerUserId() bool`

HasBrokerUserId returns a boolean if a field has been set.

### GetBrokeredRemotely

`func (o *SessionInfo) GetBrokeredRemotely() bool`

GetBrokeredRemotely returns the BrokeredRemotely field if non-nil, zero value otherwise.

### GetBrokeredRemotelyOk

`func (o *SessionInfo) GetBrokeredRemotelyOk() (*bool, bool)`

GetBrokeredRemotelyOk returns a tuple with the BrokeredRemotely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokeredRemotely

`func (o *SessionInfo) SetBrokeredRemotely(v bool)`

SetBrokeredRemotely sets BrokeredRemotely field to given value.

### HasBrokeredRemotely

`func (o *SessionInfo) HasBrokeredRemotely() bool`

HasBrokeredRemotely returns a boolean if a field has been set.

### GetClientData

`func (o *SessionInfo) GetClientData() ClientData`

GetClientData returns the ClientData field if non-nil, zero value otherwise.

### GetClientDataOk

`func (o *SessionInfo) GetClientDataOk() (*ClientData, bool)`

GetClientDataOk returns a tuple with the ClientData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientData

`func (o *SessionInfo) SetClientData(v ClientData)`

SetClientData sets ClientData field to given value.

### HasClientData

`func (o *SessionInfo) HasClientData() bool`

HasClientData returns a boolean if a field has been set.

### GetDesktopPoolId

`func (o *SessionInfo) GetDesktopPoolId() string`

GetDesktopPoolId returns the DesktopPoolId field if non-nil, zero value otherwise.

### GetDesktopPoolIdOk

`func (o *SessionInfo) GetDesktopPoolIdOk() (*string, bool)`

GetDesktopPoolIdOk returns a tuple with the DesktopPoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolId

`func (o *SessionInfo) SetDesktopPoolId(v string)`

SetDesktopPoolId sets DesktopPoolId field to given value.

### HasDesktopPoolId

`func (o *SessionInfo) HasDesktopPoolId() bool`

HasDesktopPoolId returns a boolean if a field has been set.

### GetDisconnectedTime

`func (o *SessionInfo) GetDisconnectedTime() int64`

GetDisconnectedTime returns the DisconnectedTime field if non-nil, zero value otherwise.

### GetDisconnectedTimeOk

`func (o *SessionInfo) GetDisconnectedTimeOk() (*int64, bool)`

GetDisconnectedTimeOk returns a tuple with the DisconnectedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnectedTime

`func (o *SessionInfo) SetDisconnectedTime(v int64)`

SetDisconnectedTime sets DisconnectedTime field to given value.

### HasDisconnectedTime

`func (o *SessionInfo) HasDisconnectedTime() bool`

HasDisconnectedTime returns a boolean if a field has been set.

### GetFarmId

`func (o *SessionInfo) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *SessionInfo) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *SessionInfo) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.

### HasFarmId

`func (o *SessionInfo) HasFarmId() bool`

HasFarmId returns a boolean if a field has been set.

### GetId

`func (o *SessionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SessionInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdleDuration

`func (o *SessionInfo) GetIdleDuration() int64`

GetIdleDuration returns the IdleDuration field if non-nil, zero value otherwise.

### GetIdleDurationOk

`func (o *SessionInfo) GetIdleDurationOk() (*int64, bool)`

GetIdleDurationOk returns a tuple with the IdleDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleDuration

`func (o *SessionInfo) SetIdleDuration(v int64)`

SetIdleDuration sets IdleDuration field to given value.

### HasIdleDuration

`func (o *SessionInfo) HasIdleDuration() bool`

HasIdleDuration returns a boolean if a field has been set.

### GetLastSessionDurationMs

`func (o *SessionInfo) GetLastSessionDurationMs() int64`

GetLastSessionDurationMs returns the LastSessionDurationMs field if non-nil, zero value otherwise.

### GetLastSessionDurationMsOk

`func (o *SessionInfo) GetLastSessionDurationMsOk() (*int64, bool)`

GetLastSessionDurationMsOk returns a tuple with the LastSessionDurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSessionDurationMs

`func (o *SessionInfo) SetLastSessionDurationMs(v int64)`

SetLastSessionDurationMs sets LastSessionDurationMs field to given value.

### HasLastSessionDurationMs

`func (o *SessionInfo) HasLastSessionDurationMs() bool`

HasLastSessionDurationMs returns a boolean if a field has been set.

### GetMachineId

`func (o *SessionInfo) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *SessionInfo) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *SessionInfo) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *SessionInfo) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetRdsServerId

`func (o *SessionInfo) GetRdsServerId() string`

GetRdsServerId returns the RdsServerId field if non-nil, zero value otherwise.

### GetRdsServerIdOk

`func (o *SessionInfo) GetRdsServerIdOk() (*string, bool)`

GetRdsServerIdOk returns a tuple with the RdsServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsServerId

`func (o *SessionInfo) SetRdsServerId(v string)`

SetRdsServerId sets RdsServerId field to given value.

### HasRdsServerId

`func (o *SessionInfo) HasRdsServerId() bool`

HasRdsServerId returns a boolean if a field has been set.

### GetResourcedRemotely

`func (o *SessionInfo) GetResourcedRemotely() bool`

GetResourcedRemotely returns the ResourcedRemotely field if non-nil, zero value otherwise.

### GetResourcedRemotelyOk

`func (o *SessionInfo) GetResourcedRemotelyOk() (*bool, bool)`

GetResourcedRemotelyOk returns a tuple with the ResourcedRemotely field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcedRemotely

`func (o *SessionInfo) SetResourcedRemotely(v bool)`

SetResourcedRemotely sets ResourcedRemotely field to given value.

### HasResourcedRemotely

`func (o *SessionInfo) HasResourcedRemotely() bool`

HasResourcedRemotely returns a boolean if a field has been set.

### GetSecurityGatewayData

`func (o *SessionInfo) GetSecurityGatewayData() SecurityGatewayData`

GetSecurityGatewayData returns the SecurityGatewayData field if non-nil, zero value otherwise.

### GetSecurityGatewayDataOk

`func (o *SessionInfo) GetSecurityGatewayDataOk() (*SecurityGatewayData, bool)`

GetSecurityGatewayDataOk returns a tuple with the SecurityGatewayData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGatewayData

`func (o *SessionInfo) SetSecurityGatewayData(v SecurityGatewayData)`

SetSecurityGatewayData sets SecurityGatewayData field to given value.

### HasSecurityGatewayData

`func (o *SessionInfo) HasSecurityGatewayData() bool`

HasSecurityGatewayData returns a boolean if a field has been set.

### GetSessionProtocol

`func (o *SessionInfo) GetSessionProtocol() string`

GetSessionProtocol returns the SessionProtocol field if non-nil, zero value otherwise.

### GetSessionProtocolOk

`func (o *SessionInfo) GetSessionProtocolOk() (*string, bool)`

GetSessionProtocolOk returns a tuple with the SessionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocol

`func (o *SessionInfo) SetSessionProtocol(v string)`

SetSessionProtocol sets SessionProtocol field to given value.

### HasSessionProtocol

`func (o *SessionInfo) HasSessionProtocol() bool`

HasSessionProtocol returns a boolean if a field has been set.

### GetSessionState

`func (o *SessionInfo) GetSessionState() string`

GetSessionState returns the SessionState field if non-nil, zero value otherwise.

### GetSessionStateOk

`func (o *SessionInfo) GetSessionStateOk() (*string, bool)`

GetSessionStateOk returns a tuple with the SessionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionState

`func (o *SessionInfo) SetSessionState(v string)`

SetSessionState sets SessionState field to given value.

### HasSessionState

`func (o *SessionInfo) HasSessionState() bool`

HasSessionState returns a boolean if a field has been set.

### GetSessionType

`func (o *SessionInfo) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *SessionInfo) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *SessionInfo) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.

### HasSessionType

`func (o *SessionInfo) HasSessionType() bool`

HasSessionType returns a boolean if a field has been set.

### GetStartTime

`func (o *SessionInfo) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *SessionInfo) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *SessionInfo) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *SessionInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUnauthenticated

`func (o *SessionInfo) GetUnauthenticated() bool`

GetUnauthenticated returns the Unauthenticated field if non-nil, zero value otherwise.

### GetUnauthenticatedOk

`func (o *SessionInfo) GetUnauthenticatedOk() (*bool, bool)`

GetUnauthenticatedOk returns a tuple with the Unauthenticated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnauthenticated

`func (o *SessionInfo) SetUnauthenticated(v bool)`

SetUnauthenticated sets Unauthenticated field to given value.

### HasUnauthenticated

`func (o *SessionInfo) HasUnauthenticated() bool`

HasUnauthenticated returns a boolean if a field has been set.

### GetUserId

`func (o *SessionInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SessionInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SessionInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SessionInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


