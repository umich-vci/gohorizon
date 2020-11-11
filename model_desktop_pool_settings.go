/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// DesktopPoolSettings Information related to Desktop Pool Settings.
type DesktopPoolSettings struct {
	// Whether multiple sessions are allowed per user for this pool. This is valid for RDS desktop pools.For other Desktops, allow_multiple_sessions_per_user in settings will be applicable.Default value is false.
	AllowMutilpleSessionsPerUser bool `json:"allow_mutilple_sessions_per_user,omitempty"`
	// Name of the category folder in the user's OS containing a shortcut to the desktop pool.Will be unset if the desktop does not belong to a category.This property defines valid folder names with a max length of 64 characters and up to 4 subdirectory levels.The subdirectories can be specified using a backslash, e.g. (dir1\\dir2\\dir3\\dir4). Folder names can't start orend with a backslash nor can there be 2 or more backslashes together. Combinations such as(\\dir1, dir1\\dir2\\, dir1\\\\dir2, dir1\\\\\\dir2) are invalid. The windows reserved keywords(CON, PRN, NUL, AUX, COM1 - COM9, LPT1 - LPT9 etc.) are not allowed in subdirectory names.
	CategoryFolderName string `json:"category_folder_name,omitempty"`
	// Indicates whether this desktop is assigned to a workspace in Horizon Cloud Services. Default value is false.
	CloudAssigned bool `json:"cloud_assigned,omitempty"`
	// Indicates whether this desktop is managed by Horizon Cloud Services.This can be false only when cloud_assigned is false. Default value is false.
	CloudManaged bool `json:"cloud_managed,omitempty"`
	// List of tags for which the access to the desktop pool is restricted to.No list indicates that desktop pool can be accessed from any connection server.
	CsRestrictionTags []string `json:"cs_restriction_tags,omitempty"`
	// Indicates whether the desktop pool is in the process of being deleted.Default value is false.
	DeleteInProgress        bool                               `json:"delete_in_progress"`
	DisplayProtocolSettings DesktopPoolDisplayProtocolSettings `json:"display_protocol_settings,omitempty"`
	// Client restrictions to be applied to the desktop pool.Currently it is valid for RDS desktop pools only. Default value is false.
	EnableClientRestrictions bool                       `json:"enable_client_restrictions,omitempty"`
	SessionSettings          DesktopPoolSessionSettings `json:"session_settings,omitempty"`
	// Supported session types for this desktop pool. If application sessions are selected to besupported then this desktop pool can be used for application pool creation. This will beuseful when the machines in the pool support application remoting. Default value of DESKTOP. * DESKTOP: Only desktop sessions are supported. * APPLICATION: Only application sessions are supported. * DESKTOP_AND_APPLICATION: Both desktop and application sessions are supported.
	SessionType string `json:"session_type,omitempty"`
	// Locations of the category folder in the user's OS containing a shortcut to the desktop pool.The value will be present if categoryFolderName is set.
	ShortcutLocations []string `json:"shortcut_locations,omitempty"`
}