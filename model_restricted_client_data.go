/*
 * Horizon Server API
 *
 * Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources.
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

// RestrictedClientData Information related to Restricted Horizon Clients.
type RestrictedClientData struct {
	// The type of Horizon Client. * WINDOWS: The client is the Horizon Client for Windows. * MAC: The client is the Horizon Client for MacOS. * HTMLACCESS: The client is a Web client. * LINUX: The client is the Horizon Client for Linux. * IOS: The client is the Horizon Client for IOS devices. * ANDROID: The client is the Horizon Client for Android. * WINSTORE: The client is the Horizon Client for Windows 10 UWP. * CHROME: The client is the Horizon Client for Chrome Native OS.
	Type string `json:"type,omitempty"`
	// The version of Horizon Client.
	Version string `json:"version,omitempty"`
}
