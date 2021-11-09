/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2106
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// DesktopPoolDisplayProtocolSettings Settings for the networking protocol to display the remote machine.
type DesktopPoolDisplayProtocolSettings struct {
	// Indicates whether the users can choose the protocol. Default value is true.
	AllowUsersToChooseProtocol bool `json:"allow_users_to_choose_protocol"`
	// The default display protocol for the desktop pool. For a managed desktop pool, this will default to PCOIP.For an unmanaged desktop pool, this will default to RDP. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol.
	DefaultDisplayProtocol string `json:"default_display_protocol"`
	// List of supported display protocols for this desktop pool.Default value is [PCOIP, RDP, BLAST].
	DisplayProtocols []string `json:"display_protocols"`
	// When 3D rendering is managed by the vSphere Client, this enables support for NVIDIA GRID vGPUs.This will be false if 3D rendering is not managed by the vSphere Client. If this is true,the host or cluster associated with the desktop pool must support NVIDIA GRID and vGPU types required by the desktop pool's VirtualMachines,VmTemplate, or BaseImageSnapshot. If this is false, the desktop pool's VirtualMachines, VmTemplate, orBaseImageSnapshot must not support NVIDIA GRID vGPUs. Since suspending VMs with passthroughdevices such as vGPUs is not possible, power_policy cannot be set to SUSPEND if this is enabled.Default value is false.
	GridVgpusEnabled *bool `json:"grid_vgpus_enabled,omitempty"`
	// This property is no longer in use for Horizon Components. It is always set to true. HTML Access, enabled by VMware Blast technology, allows users to connect to machines from Web browsers. Horizon Client software does not have to be installed on the client devices. To enable HTML Access, you must install the HTML Machine Access feature pack. Also, Blast must be configured as a supported protocol in displayProtocols.
	HtmlAccessEnabled bool `json:"html_access_enabled"`
	// When 3D is disabled, the 'Max number of monitors' and 'Max resolution of any one monitor' settings determine the amount ofvRAM assigned to machines in this desktop. The greater these values are, the more memory will be consumeon the associated ESX hosts. Existing virtual machines must be powered off and subsequently powered onfor the change to take effect. A restart will not cause the changes to take effect. If 3D is enabled and managedby View, the maximum number of monitors must be 1 or 2. For Instant Clones, this value is inherited from snapshot of Master VM.
	MaxNumberOfMonitors *int32 `json:"max_number_of_monitors,omitempty"`
	// If 3D rendering is enabled and managed by View, this must be set to the default value. When 3D rendering is disabled,the 'Max number of monitors' and 'Max resolution of any one monitor' settings determine the amount of vRAM assignedto machines in this desktop. The greater these values are, the more memory will be consumed on the associated ESX hosts.This setting is only relevant on managed machines. Existing virtual machines must be powered off and subsequently poweredon for the change to take effect. A restart will not cause the changes to take effect. For Instant Clones,this value is inherited from snapshot of Master VM. This property has a default value of WUXGA. * WSXGA_PLUS: 1680x1050 resolution. * WUXGA: 1920x1200 resolution. * WQXGA: 2560x1600 resolution. * UHD: 3840x2160 resolution.
	MaxResolutionOfAnyOneMonitor *string `json:"max_resolution_of_any_one_monitor,omitempty"`
	// 3D rendering is supported on Windows 7 or later guests running on VMs with virtual hardware version8 or later. The default protocol must be PCoIP and users must not be allowed to choose their ownprotocol to enable 3D rendering. For instant clone source desktop 3D rendering always mapped to MANAGE_BY_VSPHERE_CLIENT.Default value is DISABLED. * MANAGE_BY_VSPHERE_CLIENT: 3D rendering managed by vSphere Client. * AUTOMATIC: 3D rendering is automatic. * SOFTWARE: 3D rendering is software dependent. The software renderer is supported (at minimum) on virtual hardware version 8 in a vSphere 5.0 environment. * HARDWARE: 3D rendering is hardware dependent. The hardware-based renderer is supported (at minimum) on virtual hardware version 9 in a vSphere 5.1 environment. * DISABLED: 3D rendering is disabled.
	Renderer3d string `json:"renderer3d"`
	// Enable session collaboration feature. Session collaborationallows a user to share their remote session with other users.BLAST must be configured as a supported protocol in supported_display_protocols.Default value is false.
	SessionCollaborationEnabled bool `json:"session_collaboration_enabled"`
	// NVIDIA GRID vGPUs might have multiple profiles and any one of the available profiles can beapplied to newly created instant clone desktop. The profile specified in this field will beused in the newly created instant clone desktop. Will be set if enable_grid_vgpus set to true.
	VgpuGridProfile *string `json:"vgpu_grid_profile,omitempty"`
	// vRAM size for View managed 3D rendering. More VRAM can improve 3D performance.Size is in MB. On ESXi 5.0 hosts, the renderer allows a maximum VRAM size of 128MB. On ESXi 5.1and later hosts, the maximum VRAM size is 512MB. For Instant Clones, this value is inherited from snapshot of Master VM.This property is applicable when 3D renderer is not disabled. This has a default value of 96.
	VramSizeMb int32 `json:"vram_size_mb"`
}

// NewDesktopPoolDisplayProtocolSettings instantiates a new DesktopPoolDisplayProtocolSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolDisplayProtocolSettings(allowUsersToChooseProtocol bool, defaultDisplayProtocol string, displayProtocols []string, htmlAccessEnabled bool, renderer3d string, sessionCollaborationEnabled bool, vramSizeMb int32) *DesktopPoolDisplayProtocolSettings {
	this := DesktopPoolDisplayProtocolSettings{}
	this.AllowUsersToChooseProtocol = allowUsersToChooseProtocol
	this.DefaultDisplayProtocol = defaultDisplayProtocol
	this.DisplayProtocols = displayProtocols
	this.HtmlAccessEnabled = htmlAccessEnabled
	this.Renderer3d = renderer3d
	this.SessionCollaborationEnabled = sessionCollaborationEnabled
	this.VramSizeMb = vramSizeMb
	return &this
}

// NewDesktopPoolDisplayProtocolSettingsWithDefaults instantiates a new DesktopPoolDisplayProtocolSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolDisplayProtocolSettingsWithDefaults() *DesktopPoolDisplayProtocolSettings {
	this := DesktopPoolDisplayProtocolSettings{}
	return &this
}

// GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field value
func (o *DesktopPoolDisplayProtocolSettings) GetAllowUsersToChooseProtocol() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowUsersToChooseProtocol
}

// GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetAllowUsersToChooseProtocolOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowUsersToChooseProtocol, true
}

// SetAllowUsersToChooseProtocol sets field value
func (o *DesktopPoolDisplayProtocolSettings) SetAllowUsersToChooseProtocol(v bool) {
	o.AllowUsersToChooseProtocol = v
}

// GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field value
func (o *DesktopPoolDisplayProtocolSettings) GetDefaultDisplayProtocol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultDisplayProtocol
}

// GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetDefaultDisplayProtocolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultDisplayProtocol, true
}

// SetDefaultDisplayProtocol sets field value
func (o *DesktopPoolDisplayProtocolSettings) SetDefaultDisplayProtocol(v string) {
	o.DefaultDisplayProtocol = v
}

// GetDisplayProtocols returns the DisplayProtocols field value
func (o *DesktopPoolDisplayProtocolSettings) GetDisplayProtocols() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.DisplayProtocols
}

// GetDisplayProtocolsOk returns a tuple with the DisplayProtocols field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetDisplayProtocolsOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayProtocols, true
}

// SetDisplayProtocols sets field value
func (o *DesktopPoolDisplayProtocolSettings) SetDisplayProtocols(v []string) {
	o.DisplayProtocols = v
}

// GetGridVgpusEnabled returns the GridVgpusEnabled field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettings) GetGridVgpusEnabled() bool {
	if o == nil || o.GridVgpusEnabled == nil {
		var ret bool
		return ret
	}
	return *o.GridVgpusEnabled
}

// GetGridVgpusEnabledOk returns a tuple with the GridVgpusEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetGridVgpusEnabledOk() (*bool, bool) {
	if o == nil || o.GridVgpusEnabled == nil {
		return nil, false
	}
	return o.GridVgpusEnabled, true
}

// HasGridVgpusEnabled returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettings) HasGridVgpusEnabled() bool {
	if o != nil && o.GridVgpusEnabled != nil {
		return true
	}

	return false
}

// SetGridVgpusEnabled gets a reference to the given bool and assigns it to the GridVgpusEnabled field.
func (o *DesktopPoolDisplayProtocolSettings) SetGridVgpusEnabled(v bool) {
	o.GridVgpusEnabled = &v
}

// GetHtmlAccessEnabled returns the HtmlAccessEnabled field value
func (o *DesktopPoolDisplayProtocolSettings) GetHtmlAccessEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HtmlAccessEnabled
}

// GetHtmlAccessEnabledOk returns a tuple with the HtmlAccessEnabled field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetHtmlAccessEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HtmlAccessEnabled, true
}

// SetHtmlAccessEnabled sets field value
func (o *DesktopPoolDisplayProtocolSettings) SetHtmlAccessEnabled(v bool) {
	o.HtmlAccessEnabled = v
}

// GetMaxNumberOfMonitors returns the MaxNumberOfMonitors field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettings) GetMaxNumberOfMonitors() int32 {
	if o == nil || o.MaxNumberOfMonitors == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfMonitors
}

// GetMaxNumberOfMonitorsOk returns a tuple with the MaxNumberOfMonitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetMaxNumberOfMonitorsOk() (*int32, bool) {
	if o == nil || o.MaxNumberOfMonitors == nil {
		return nil, false
	}
	return o.MaxNumberOfMonitors, true
}

// HasMaxNumberOfMonitors returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettings) HasMaxNumberOfMonitors() bool {
	if o != nil && o.MaxNumberOfMonitors != nil {
		return true
	}

	return false
}

// SetMaxNumberOfMonitors gets a reference to the given int32 and assigns it to the MaxNumberOfMonitors field.
func (o *DesktopPoolDisplayProtocolSettings) SetMaxNumberOfMonitors(v int32) {
	o.MaxNumberOfMonitors = &v
}

// GetMaxResolutionOfAnyOneMonitor returns the MaxResolutionOfAnyOneMonitor field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettings) GetMaxResolutionOfAnyOneMonitor() string {
	if o == nil || o.MaxResolutionOfAnyOneMonitor == nil {
		var ret string
		return ret
	}
	return *o.MaxResolutionOfAnyOneMonitor
}

// GetMaxResolutionOfAnyOneMonitorOk returns a tuple with the MaxResolutionOfAnyOneMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetMaxResolutionOfAnyOneMonitorOk() (*string, bool) {
	if o == nil || o.MaxResolutionOfAnyOneMonitor == nil {
		return nil, false
	}
	return o.MaxResolutionOfAnyOneMonitor, true
}

// HasMaxResolutionOfAnyOneMonitor returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettings) HasMaxResolutionOfAnyOneMonitor() bool {
	if o != nil && o.MaxResolutionOfAnyOneMonitor != nil {
		return true
	}

	return false
}

// SetMaxResolutionOfAnyOneMonitor gets a reference to the given string and assigns it to the MaxResolutionOfAnyOneMonitor field.
func (o *DesktopPoolDisplayProtocolSettings) SetMaxResolutionOfAnyOneMonitor(v string) {
	o.MaxResolutionOfAnyOneMonitor = &v
}

// GetRenderer3d returns the Renderer3d field value
func (o *DesktopPoolDisplayProtocolSettings) GetRenderer3d() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Renderer3d
}

// GetRenderer3dOk returns a tuple with the Renderer3d field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetRenderer3dOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Renderer3d, true
}

// SetRenderer3d sets field value
func (o *DesktopPoolDisplayProtocolSettings) SetRenderer3d(v string) {
	o.Renderer3d = v
}

// GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field value
func (o *DesktopPoolDisplayProtocolSettings) GetSessionCollaborationEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SessionCollaborationEnabled
}

// GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetSessionCollaborationEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionCollaborationEnabled, true
}

// SetSessionCollaborationEnabled sets field value
func (o *DesktopPoolDisplayProtocolSettings) SetSessionCollaborationEnabled(v bool) {
	o.SessionCollaborationEnabled = v
}

// GetVgpuGridProfile returns the VgpuGridProfile field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettings) GetVgpuGridProfile() string {
	if o == nil || o.VgpuGridProfile == nil {
		var ret string
		return ret
	}
	return *o.VgpuGridProfile
}

// GetVgpuGridProfileOk returns a tuple with the VgpuGridProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetVgpuGridProfileOk() (*string, bool) {
	if o == nil || o.VgpuGridProfile == nil {
		return nil, false
	}
	return o.VgpuGridProfile, true
}

// HasVgpuGridProfile returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettings) HasVgpuGridProfile() bool {
	if o != nil && o.VgpuGridProfile != nil {
		return true
	}

	return false
}

// SetVgpuGridProfile gets a reference to the given string and assigns it to the VgpuGridProfile field.
func (o *DesktopPoolDisplayProtocolSettings) SetVgpuGridProfile(v string) {
	o.VgpuGridProfile = &v
}

// GetVramSizeMb returns the VramSizeMb field value
func (o *DesktopPoolDisplayProtocolSettings) GetVramSizeMb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.VramSizeMb
}

// GetVramSizeMbOk returns a tuple with the VramSizeMb field value
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettings) GetVramSizeMbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VramSizeMb, true
}

// SetVramSizeMb sets field value
func (o *DesktopPoolDisplayProtocolSettings) SetVramSizeMb(v int32) {
	o.VramSizeMb = v
}

func (o DesktopPoolDisplayProtocolSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["allow_users_to_choose_protocol"] = o.AllowUsersToChooseProtocol
	}
	if true {
		toSerialize["default_display_protocol"] = o.DefaultDisplayProtocol
	}
	if true {
		toSerialize["display_protocols"] = o.DisplayProtocols
	}
	if o.GridVgpusEnabled != nil {
		toSerialize["grid_vgpus_enabled"] = o.GridVgpusEnabled
	}
	if true {
		toSerialize["html_access_enabled"] = o.HtmlAccessEnabled
	}
	if o.MaxNumberOfMonitors != nil {
		toSerialize["max_number_of_monitors"] = o.MaxNumberOfMonitors
	}
	if o.MaxResolutionOfAnyOneMonitor != nil {
		toSerialize["max_resolution_of_any_one_monitor"] = o.MaxResolutionOfAnyOneMonitor
	}
	if true {
		toSerialize["renderer3d"] = o.Renderer3d
	}
	if true {
		toSerialize["session_collaboration_enabled"] = o.SessionCollaborationEnabled
	}
	if o.VgpuGridProfile != nil {
		toSerialize["vgpu_grid_profile"] = o.VgpuGridProfile
	}
	if true {
		toSerialize["vram_size_mb"] = o.VramSizeMb
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolDisplayProtocolSettings struct {
	value *DesktopPoolDisplayProtocolSettings
	isSet bool
}

func (v NullableDesktopPoolDisplayProtocolSettings) Get() *DesktopPoolDisplayProtocolSettings {
	return v.value
}

func (v *NullableDesktopPoolDisplayProtocolSettings) Set(val *DesktopPoolDisplayProtocolSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolDisplayProtocolSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolDisplayProtocolSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolDisplayProtocolSettings(val *DesktopPoolDisplayProtocolSettings) *NullableDesktopPoolDisplayProtocolSettings {
	return &NullableDesktopPoolDisplayProtocolSettings{value: val, isSet: true}
}

func (v NullableDesktopPoolDisplayProtocolSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolDisplayProtocolSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
