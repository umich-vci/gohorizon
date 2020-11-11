/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// SessionInfo Information related to Session.
type SessionInfo struct {
	// Access group id associated with the session.  For a non-RDS desktop session, this is the desktop pool's access group id.  For an RDS desktop session, this is the RDS desktop pool's farm's access group id.  For an application session, this is the application's farm's access group id.
	AccessGroupId string `json:"access_group_id,omitempty"`
	// Version of agent This property need not be set.
	AgentVersion string `json:"agent_version,omitempty"`
	// Names of the applications launched in the session.  It will be only set when session_type is APPLICATION.
	ApplicationNames []string `json:"application_names,omitempty"`
	// User SID for the broker user associated with the session.  It will be unset for non-broker sessions.
	BrokerUserId string `json:"broker_user_id,omitempty"`
	// Indicates whether the session is brokered from a remote pod.  It is set only if the Horizon View agent where the session resides is version 6.0 or later.
	BrokeredRemotely bool       `json:"brokered_remotely,omitempty"`
	ClientData       ClientData `json:"client_data,omitempty"`
	// Unique desktop pool id for the session.  This is unset if the session is not brokered through a desktop, such as for direct console access.
	DesktopPoolId string `json:"desktop_pool_id,omitempty"`
	// Epoch time in milli seconds, when the session was last disconnected.  This will be unset if the session's machine has an error state, or if the session has never been disconnected.
	DisconnectedTime int64 `json:"disconnected_time,omitempty"`
	// Unique farm id for this RDS desktop or application session. This is unset if the session is not brokered through a farm, such as for application sessions or direct console access.
	FarmId string `json:"farm_id,omitempty"`
	// Unique id representing a session.
	Id string `json:"id,omitempty"`
	// Idle time duration in minutes, indicating how long the end user of the session has been idle for. This property need not be set.
	IdleDuration int64 `json:"idle_duration,omitempty"`
	// Duration of the last connection period of the session in milli seconds. If the session is currently connected, this is the duration that the session has been in connected state. If the session is currently disconnected, this is the duration of its previous connection period. This will be unset on error.
	LastSessionDurationMs int64 `json:"last_session_duration_ms,omitempty"`
	// Unique machine id for the session.  This is unset for RDS Desktop or application sessions. If desktop pool id is unset, it is the id of registered un-managed physical machine.
	MachineId string `json:"machine_id,omitempty"`
	// Unique RDS server id for the RDS desktop or application session. This property need not be set.
	RdsServerId string `json:"rds_server_id,omitempty"`
	// Indicates whether the session is running on remote pod resource.
	ResourcedRemotely   bool                `json:"resourced_remotely,omitempty"`
	SecurityGatewayData SecurityGatewayData `json:"security_gateway_data,omitempty"`
	// Protocol for the session.  It will be unset for disconnected sessions. * PCOIP: Display protocol is PCoIP. * RDP: Display protocol is RDP. * BLAST: Display protocol is BLAST. * CONSOLE: Display protocol is console. * UNKNOWN: Display protocol is unknown.
	SessionProtocol string `json:"session_protocol,omitempty"`
	// State of session. * CONNECTED: Session is connected * DISCONNECTED: Session is disconnected * PENDING: Session is pending
	SessionState string `json:"session_state,omitempty"`
	// Type of session. * DESKTOP: Desktop or RDS desktop session. * APPLICATION: Application session.
	SessionType string `json:"session_type,omitempty"`
	// Epoch time in milli seconds, when the session was originally logged in.  The lifecycle of a session begins at login and ends at logout, with any number of connect and disconnect occurrences in between. The first connection time will be shortly after this time.  This property need not be set.
	StartTime int64 `json:"start_time,omitempty"`
	// Indicates whether the session is of unauthenticated access user.  This property need not be set.
	Unauthenticated bool `json:"unauthenticated,omitempty"`
	// Unique SID of the user logged into the session.  It may not match the broker user id for non-SSO scenarios.
	UserId string `json:"user_id,omitempty"`
}
