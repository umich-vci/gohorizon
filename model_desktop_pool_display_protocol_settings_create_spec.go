/*
Horizon Server API

Welcome to the Horizon Server API Reference documentation. This API reference provides comprehensive information about status of all Horizon Server components and resources. <br> Choose Latest spec from dropdown to view API reference on latest version available.

API version: 2111
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gohorizon

import (
	"encoding/json"
)

// DesktopPoolDisplayProtocolSettingsCreateSpec Applicable To: Automated and Manual desktop pool. <br>Display protocol settings for Automated and Manual desktop pool.
type DesktopPoolDisplayProtocolSettingsCreateSpec struct {
	// Indicates whether the users can choose the protocol. Default value is true.
	AllowUsersToChooseProtocol *bool `json:"allow_users_to_choose_protocol,omitempty"`
	// The default display protocol for the desktop pool. For a managed desktop pool, this will default to PCOIP and for unmanaged desktop pool, this will default to RDP. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol.
	DefaultDisplayProtocol *string `json:"default_display_protocol,omitempty"`
	// When 3D rendering is managed by the vSphere Client, this enables support for NVIDIA GRID vGPUs. This will be false if 3D rendering is not managed by the vSphere Client. If this is true, the host or cluster associated with the desktop pool must support NVIDIA GRID and vGPU types required by the desktop pool's VirtualMachines, VmTemplate or BaseImageSnapshot. If this is false, the desktop pool's VirtualMachines, VmTemplate or BaseImageSnapshot must not support NVIDIA GRID vGPUs. Since suspending VMs with passthrough devices such as vGPUs is not possible, power_policy cannot be set to SUSPEND if this is enabled. Default value is false.
	GridVgpusEnabled *bool `json:"grid_vgpus_enabled,omitempty"`
	// When render3D is disabled, the max_number_of_monitors and max_resolution_of_any_one_monitor settings determine the amount of vRAM assigned to machines in this desktop. The greater these values are, the more memory will be consume on the associated ESX hosts. Existing virtual machines must be powered off and subsequently powered on for the change to take effect. A restart will not cause the changes to take effect. If 3D is enabled  and managed by View, the maximum number of monitors must be 1 or 2. For Instant Clones, this value is inherited from snapshot of Master VM. This property has a default value of 2. <br> This property is required if renderer3D is set to AUTOMATIC, SOFTWARE, HARDWARE or DISABLED.
	MaxNumberOfMonitors *int32 `json:"max_number_of_monitors,omitempty"`
	// If 3D rendering is enabled and managed by View, this must be set to the default value. When 3D rendering is disabled, the max_number_of_monitors and max_resolution_of_any_one_monitor settings determine the amount of vRAM assigned to machines in this desktop. The greater these values are, the more memory will be consumed on the associated ESX hosts. This setting is only relevant on managed machines. Existing virtual machines must be powered off and subsequently powered on for the change to take effect. A restart will not cause the changes to take effect. For Instant Clones, this value is inherited from snapshot of Master VM. This property has a default value of WUXGA. <br> This property is required if renderer3D is set to AUTOMATIC, SOFTWARE, HARDWARE or DISABLED. * WSXGA_PLUS: 1680x1050 resolution. * WUXGA: 1920x1200 resolution. * WQXGA: 2560x1600 resolution. * UHD: 3840x2160 resolution. * UHD_5K: 5120x2880 resolution. * UHD_8K: 7680x4320 resolution.
	MaxResolutionOfAnyOneMonitor *string `json:"max_resolution_of_any_one_monitor,omitempty"`
	// 3D rendering is supported on Windows 7 or later guests running on VMs with virtual hardware version 8 or later. The default_display_protocol must set to PCOIP and allow_users_to_choose_protocol must be set to false to enable 3D rendering. For instant clone source desktop 3D rendering always mapped to MANAGE_BY_VSPHERE_CLIENT. Default value is DISABLED. * MANAGE_BY_VSPHERE_CLIENT: 3D rendering managed by vSphere Client. * AUTOMATIC: 3D rendering is automatic. * SOFTWARE: 3D rendering is software dependent. The software renderer is supported (at minimum) on virtual hardware version 8 in a vSphere 5.0 environment. * HARDWARE: 3D rendering is hardware dependent. The hardware-based renderer is supported (at minimum) on virtual hardware version 9 in a vSphere 5.1 environment. * DISABLED: 3D rendering is disabled.
	Renderer3d *string `json:"renderer3d,omitempty"`
	// Applicable To: Automated and Manual pools with default value of false. <br> Enable session collaboration feature. Session collaboration allows a user to share their remote session with other users. BLAST must be configured as a supported protocol in supported_display_protocols.
	SessionCollaborationEnabled *bool `json:"session_collaboration_enabled,omitempty"`
	// vRAM size for View managed 3D rendering. More VRAM can improve 3D performance. Size is in MB. On ESXi 5.0 hosts, the renderer allows a maximum VRAM size of 128MB. On ESXi 5.1 and later hosts, the maximum VRAM size is 512MB. For Instant Clones, this value is inherited from snapshot of Master VM. This property has a default value of 96. <br>This property is required if renderer3d is set to AUTOMATIC, SOFTWARE or HARDWARE.
	VramSizeMb *int32 `json:"vram_size_mb,omitempty"`
}

// NewDesktopPoolDisplayProtocolSettingsCreateSpec instantiates a new DesktopPoolDisplayProtocolSettingsCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesktopPoolDisplayProtocolSettingsCreateSpec() *DesktopPoolDisplayProtocolSettingsCreateSpec {
	this := DesktopPoolDisplayProtocolSettingsCreateSpec{}
	return &this
}

// NewDesktopPoolDisplayProtocolSettingsCreateSpecWithDefaults instantiates a new DesktopPoolDisplayProtocolSettingsCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesktopPoolDisplayProtocolSettingsCreateSpecWithDefaults() *DesktopPoolDisplayProtocolSettingsCreateSpec {
	this := DesktopPoolDisplayProtocolSettingsCreateSpec{}
	return &this
}

// GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetAllowUsersToChooseProtocol() bool {
	if o == nil || o.AllowUsersToChooseProtocol == nil {
		var ret bool
		return ret
	}
	return *o.AllowUsersToChooseProtocol
}

// GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetAllowUsersToChooseProtocolOk() (*bool, bool) {
	if o == nil || o.AllowUsersToChooseProtocol == nil {
		return nil, false
	}
	return o.AllowUsersToChooseProtocol, true
}

// HasAllowUsersToChooseProtocol returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasAllowUsersToChooseProtocol() bool {
	if o != nil && o.AllowUsersToChooseProtocol != nil {
		return true
	}

	return false
}

// SetAllowUsersToChooseProtocol gets a reference to the given bool and assigns it to the AllowUsersToChooseProtocol field.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetAllowUsersToChooseProtocol(v bool) {
	o.AllowUsersToChooseProtocol = &v
}

// GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetDefaultDisplayProtocol() string {
	if o == nil || o.DefaultDisplayProtocol == nil {
		var ret string
		return ret
	}
	return *o.DefaultDisplayProtocol
}

// GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetDefaultDisplayProtocolOk() (*string, bool) {
	if o == nil || o.DefaultDisplayProtocol == nil {
		return nil, false
	}
	return o.DefaultDisplayProtocol, true
}

// HasDefaultDisplayProtocol returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasDefaultDisplayProtocol() bool {
	if o != nil && o.DefaultDisplayProtocol != nil {
		return true
	}

	return false
}

// SetDefaultDisplayProtocol gets a reference to the given string and assigns it to the DefaultDisplayProtocol field.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetDefaultDisplayProtocol(v string) {
	o.DefaultDisplayProtocol = &v
}

// GetGridVgpusEnabled returns the GridVgpusEnabled field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetGridVgpusEnabled() bool {
	if o == nil || o.GridVgpusEnabled == nil {
		var ret bool
		return ret
	}
	return *o.GridVgpusEnabled
}

// GetGridVgpusEnabledOk returns a tuple with the GridVgpusEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetGridVgpusEnabledOk() (*bool, bool) {
	if o == nil || o.GridVgpusEnabled == nil {
		return nil, false
	}
	return o.GridVgpusEnabled, true
}

// HasGridVgpusEnabled returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasGridVgpusEnabled() bool {
	if o != nil && o.GridVgpusEnabled != nil {
		return true
	}

	return false
}

// SetGridVgpusEnabled gets a reference to the given bool and assigns it to the GridVgpusEnabled field.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetGridVgpusEnabled(v bool) {
	o.GridVgpusEnabled = &v
}

// GetMaxNumberOfMonitors returns the MaxNumberOfMonitors field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetMaxNumberOfMonitors() int32 {
	if o == nil || o.MaxNumberOfMonitors == nil {
		var ret int32
		return ret
	}
	return *o.MaxNumberOfMonitors
}

// GetMaxNumberOfMonitorsOk returns a tuple with the MaxNumberOfMonitors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetMaxNumberOfMonitorsOk() (*int32, bool) {
	if o == nil || o.MaxNumberOfMonitors == nil {
		return nil, false
	}
	return o.MaxNumberOfMonitors, true
}

// HasMaxNumberOfMonitors returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasMaxNumberOfMonitors() bool {
	if o != nil && o.MaxNumberOfMonitors != nil {
		return true
	}

	return false
}

// SetMaxNumberOfMonitors gets a reference to the given int32 and assigns it to the MaxNumberOfMonitors field.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetMaxNumberOfMonitors(v int32) {
	o.MaxNumberOfMonitors = &v
}

// GetMaxResolutionOfAnyOneMonitor returns the MaxResolutionOfAnyOneMonitor field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetMaxResolutionOfAnyOneMonitor() string {
	if o == nil || o.MaxResolutionOfAnyOneMonitor == nil {
		var ret string
		return ret
	}
	return *o.MaxResolutionOfAnyOneMonitor
}

// GetMaxResolutionOfAnyOneMonitorOk returns a tuple with the MaxResolutionOfAnyOneMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetMaxResolutionOfAnyOneMonitorOk() (*string, bool) {
	if o == nil || o.MaxResolutionOfAnyOneMonitor == nil {
		return nil, false
	}
	return o.MaxResolutionOfAnyOneMonitor, true
}

// HasMaxResolutionOfAnyOneMonitor returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasMaxResolutionOfAnyOneMonitor() bool {
	if o != nil && o.MaxResolutionOfAnyOneMonitor != nil {
		return true
	}

	return false
}

// SetMaxResolutionOfAnyOneMonitor gets a reference to the given string and assigns it to the MaxResolutionOfAnyOneMonitor field.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetMaxResolutionOfAnyOneMonitor(v string) {
	o.MaxResolutionOfAnyOneMonitor = &v
}

// GetRenderer3d returns the Renderer3d field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetRenderer3d() string {
	if o == nil || o.Renderer3d == nil {
		var ret string
		return ret
	}
	return *o.Renderer3d
}

// GetRenderer3dOk returns a tuple with the Renderer3d field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetRenderer3dOk() (*string, bool) {
	if o == nil || o.Renderer3d == nil {
		return nil, false
	}
	return o.Renderer3d, true
}

// HasRenderer3d returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasRenderer3d() bool {
	if o != nil && o.Renderer3d != nil {
		return true
	}

	return false
}

// SetRenderer3d gets a reference to the given string and assigns it to the Renderer3d field.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetRenderer3d(v string) {
	o.Renderer3d = &v
}

// GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetSessionCollaborationEnabled() bool {
	if o == nil || o.SessionCollaborationEnabled == nil {
		var ret bool
		return ret
	}
	return *o.SessionCollaborationEnabled
}

// GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetSessionCollaborationEnabledOk() (*bool, bool) {
	if o == nil || o.SessionCollaborationEnabled == nil {
		return nil, false
	}
	return o.SessionCollaborationEnabled, true
}

// HasSessionCollaborationEnabled returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasSessionCollaborationEnabled() bool {
	if o != nil && o.SessionCollaborationEnabled != nil {
		return true
	}

	return false
}

// SetSessionCollaborationEnabled gets a reference to the given bool and assigns it to the SessionCollaborationEnabled field.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetSessionCollaborationEnabled(v bool) {
	o.SessionCollaborationEnabled = &v
}

// GetVramSizeMb returns the VramSizeMb field value if set, zero value otherwise.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetVramSizeMb() int32 {
	if o == nil || o.VramSizeMb == nil {
		var ret int32
		return ret
	}
	return *o.VramSizeMb
}

// GetVramSizeMbOk returns a tuple with the VramSizeMb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetVramSizeMbOk() (*int32, bool) {
	if o == nil || o.VramSizeMb == nil {
		return nil, false
	}
	return o.VramSizeMb, true
}

// HasVramSizeMb returns a boolean if a field has been set.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasVramSizeMb() bool {
	if o != nil && o.VramSizeMb != nil {
		return true
	}

	return false
}

// SetVramSizeMb gets a reference to the given int32 and assigns it to the VramSizeMb field.
func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetVramSizeMb(v int32) {
	o.VramSizeMb = &v
}

func (o DesktopPoolDisplayProtocolSettingsCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllowUsersToChooseProtocol != nil {
		toSerialize["allow_users_to_choose_protocol"] = o.AllowUsersToChooseProtocol
	}
	if o.DefaultDisplayProtocol != nil {
		toSerialize["default_display_protocol"] = o.DefaultDisplayProtocol
	}
	if o.GridVgpusEnabled != nil {
		toSerialize["grid_vgpus_enabled"] = o.GridVgpusEnabled
	}
	if o.MaxNumberOfMonitors != nil {
		toSerialize["max_number_of_monitors"] = o.MaxNumberOfMonitors
	}
	if o.MaxResolutionOfAnyOneMonitor != nil {
		toSerialize["max_resolution_of_any_one_monitor"] = o.MaxResolutionOfAnyOneMonitor
	}
	if o.Renderer3d != nil {
		toSerialize["renderer3d"] = o.Renderer3d
	}
	if o.SessionCollaborationEnabled != nil {
		toSerialize["session_collaboration_enabled"] = o.SessionCollaborationEnabled
	}
	if o.VramSizeMb != nil {
		toSerialize["vram_size_mb"] = o.VramSizeMb
	}
	return json.Marshal(toSerialize)
}

type NullableDesktopPoolDisplayProtocolSettingsCreateSpec struct {
	value *DesktopPoolDisplayProtocolSettingsCreateSpec
	isSet bool
}

func (v NullableDesktopPoolDisplayProtocolSettingsCreateSpec) Get() *DesktopPoolDisplayProtocolSettingsCreateSpec {
	return v.value
}

func (v *NullableDesktopPoolDisplayProtocolSettingsCreateSpec) Set(val *DesktopPoolDisplayProtocolSettingsCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableDesktopPoolDisplayProtocolSettingsCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDesktopPoolDisplayProtocolSettingsCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesktopPoolDisplayProtocolSettingsCreateSpec(val *DesktopPoolDisplayProtocolSettingsCreateSpec) *NullableDesktopPoolDisplayProtocolSettingsCreateSpec {
	return &NullableDesktopPoolDisplayProtocolSettingsCreateSpec{value: val, isSet: true}
}

func (v NullableDesktopPoolDisplayProtocolSettingsCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesktopPoolDisplayProtocolSettingsCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}