/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon
// ApplicationPoolUpdateSpec Information required to update an application pool.
type ApplicationPoolUpdateSpec struct {
	AntiAffinityData ApplicationAntiAffinityData `json:"anti_affinity_data,omitempty"`
	// Name of the category folder in the user's OS containing a shortcut to the application. Unset if the application does not belong to a category.
	CategoryFolderName string `json:"category_folder_name,omitempty"`
	// Connection server restrictions. Application pool can be accessed from only those connection server instances that have a matching tag in this list. Null or empty list means that the application pool can be accessed from any connection server.
	CsRestrictionTags []string `json:"cs_restriction_tags,omitempty"`
	// Notes about the application pool.
	Description string `json:"description,omitempty"`
	// The display name is the name that users will see in Horizon client. If the display name is left blank, it defaults to name.
	DisplayName string `json:"display_name,omitempty"`
	// Indicates whether client restrictions are to be applied to application pool. Currently it is valid for application pool created from farm. Default value is false.
	EnableClientRestrictions bool `json:"enable_client_restrictions,omitempty"`
	// Whether to pre-launch the application. Default value is false.
	EnablePreLaunch bool `json:"enable_pre_launch,omitempty"`
	// Indicates whether the application pool is enabled. Default value is true
	Enabled bool `json:"enabled"`
	// Path to application executable.
	ExecutablePath string `json:"executable_path"`
	// Maximum number of multi-sessions a user can have in this application pool. This property is required if multi-session mode is set to \"ENABLED_DEFAULT_OFF\", \"ENABLED_DEFAULT_ON\", or \"ENABLED_ENFORCED\"Default value is 1.
	MaxMultiSessions int32 `json:"max_multi_sessions,omitempty"`
	// Multi-session mode for the application pool. An application launched in multi-session mode does not support reconnect behavior when user logs in from a different client instance. Multi-session mode should be disabled when pre-launch is enabled.Default value is \"DISABLED\" * DISABLED: Multi-session is not supported for this application. * ENABLED_DEFAULT_OFF: Multi-session is supported for this application but is disabled by default. The client would need to explicitly request multi-session launch, if wanted. If a legacy client is used, this will always result in a single-session application launch. * ENABLED_DEFAULT_ON: Multi-session mode is supported for this application and is enabled by default. The client can request explicitly for single-session launch, if wanted. If a legacy client is used, this will always result in a multi-session application launch. * ENABLED_ENFORCED: Multi-session is supported for this application and it is enforced. The client can not select to launch this application as a single-session application.
	MultiSessionMode string `json:"multi_session_mode,omitempty"`
	// Parameters to pass to application when launching.
	Parameters string `json:"parameters,omitempty"`
	// Application publisher
	Publisher string `json:"publisher,omitempty"`
	// Locations of the category folder in the user's OS containing a shortcut to the application. The value must be set if category folder name is provided.
	ShortcutLocations []string `json:"shortcut_locations,omitempty"`
	// Starting folder for application
	StartFolder string `json:"start_folder,omitempty"`
	SupportedFileTypesData ApplicationSupportedFileTypesData `json:"supported_file_types_data,omitempty"`
	// Application version.
	Version string `json:"version,omitempty"`
}
