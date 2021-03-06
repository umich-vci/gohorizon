/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// ApplicationPoolInfo Application pool information.<br>List API returning this summary information can use search filter query to filter on specific fields supported by filters.<br> Supported Filters : 'And', 'Or', 'Equals', 'StartsWith', 'Contains' and 'EndsWith'.<br>See the field description to know the filter types it supports.
type ApplicationPoolInfo struct {
	// Access groups can organize the entities (like application pools, desktop pools) in the organization. They can also be used for delegated administration. For application pool, this is the same as that of the farm or desktop pool that the application pool belongs to.<br>Supported Filters: 'Equals'.
	AccessGroupId    string                      `json:"access_group_id"`
	AntiAffinityData ApplicationAntiAffinityData `json:"anti_affinity_data,omitempty"`
	// Name of the category folder in the user's OS containing a shortcut to the application. Unset if the application does not belong to a category.
	CategoryFolderName string `json:"category_folder_name,omitempty"`
	// Connection server restrictions. Application pool can be accessed from only those connection server instances that have a matching tag in this list. Null or empty list means that the application pool can be accessed from any connection server.
	CsRestrictionTags []string `json:"cs_restriction_tags,omitempty"`
	// List of customized icon IDs associated with the application which the user has configured.
	CustomizedIconIds []string `json:"customized_icon_ids,omitempty"`
	// Notes about the application pool.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Description string `json:"description,omitempty"`
	// ID of the desktop pool from which this application pool is created. Either this or farm id will be set.<br>Supported Filters: 'Equals'.
	DesktopPoolId string `json:"desktop_pool_id,omitempty"`
	// The display name is the name that users will see when they connect to view client. If the display name is left blank, it defaults to name.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	DisplayName string `json:"display_name,omitempty"`
	// Indicates whether client restrictions are to be applied to application pool. Currently it is valid for application pool created from farm.<br>Supported Filters: 'Equals'.
	EnableClientRestrictions bool `json:"enable_client_restrictions,omitempty"`
	// Whether to pre-launch the application.<br>Supported Filters: 'Equals'.
	EnablePreLaunch bool `json:"enable_pre_launch,omitempty"`
	// Indicates whether the application pool is enabled.<br>Supported Filters: 'Equals'.
	Enabled bool `json:"enabled"`
	// Path to application executable.<br>Supported Filters: 'Equals', 'StartsWith', 'EndsWith' and 'Contains'.
	ExecutablePath string `json:"executable_path"`
	// ID of the farm from which this application pool is created. Either this or desktop pool id will be set.<br>Supported Filters: 'Equals'.
	FarmId string `json:"farm_id,omitempty"`
	// List of icon IDs associated with the application which are fetched from the agent.
	IconIds []string `json:"icon_ids,omitempty"`
	// Unique ID representing application pool.
	Id string `json:"id"`
	// Maximum number of multi-sessions a user can have in this application pool. This property is required if multi-session mode is set to \"ENABLED_DEFAULT_OFF\", \"ENABLED_DEFAULT_ON\", or \"ENABLED_ENFORCED\"
	MaxMultiSessions int32 `json:"max_multi_sessions,omitempty"`
	// Multi-session mode for the application pool. An application launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. Multi-session mode should be disabled when pre-launch is enabled.<br>Supported Filters: 'Equals'. * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application.
	MultiSessionMode string `json:"multi_session_mode,omitempty"`
	// The application name is the unique identifier used to identify this application pool. This property must contain only alphanumerics, underscores, and dashes. The maximum length is 64 characters.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Name string `json:"name"`
	// Parameters to pass to application when launching.
	Parameters string `json:"parameters,omitempty"`
	// Application publisher.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Publisher string `json:"publisher,omitempty"`
	// Locations of the category folder in the user's OS containing a shortcut to the application. The value must be set if category folder name is provided.
	ShortcutLocations []string `json:"shortcut_locations,omitempty"`
	// Starting folder for application.
	StartFolder            string                            `json:"start_folder,omitempty"`
	SupportedFileTypesData ApplicationSupportedFileTypesData `json:"supported_file_types_data,omitempty"`
	// Application version.<br>Supported Filters: 'Equals', 'StartsWith' and 'Contains'.
	Version string `json:"version,omitempty"`
}
