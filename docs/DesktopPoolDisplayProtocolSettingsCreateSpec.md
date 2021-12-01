# DesktopPoolDisplayProtocolSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | Pointer to **bool** | Indicates whether the users can choose the protocol. Default value is true. | [optional] 
**DefaultDisplayProtocol** | Pointer to **string** | The default display protocol for the desktop pool. For a managed desktop pool, this will default to PCOIP and for unmanaged desktop pool, this will default to RDP. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | [optional] 
**GridVgpusEnabled** | Pointer to **bool** | When 3D rendering is managed by the vSphere Client, this enables support for NVIDIA GRID vGPUs. This will be false if 3D rendering is not managed by the vSphere Client. If this is true, the host or cluster associated with the desktop pool must support NVIDIA GRID and vGPU types required by the desktop pool&#39;s VirtualMachines, VmTemplate or BaseImageSnapshot. If this is false, the desktop pool&#39;s VirtualMachines, VmTemplate or BaseImageSnapshot must not support NVIDIA GRID vGPUs. Since suspending VMs with passthrough devices such as vGPUs is not possible, power_policy cannot be set to SUSPEND if this is enabled. Default value is false. | [optional] 
**MaxNumberOfMonitors** | Pointer to **int32** | When render3D is disabled, the max_number_of_monitors and max_resolution_of_any_one_monitor settings determine the amount of vRAM assigned to machines in this desktop. The greater these values are, the more memory will be consume on the associated ESX hosts. Existing virtual machines must be powered off and subsequently powered on for the change to take effect. A restart will not cause the changes to take effect. If 3D is enabled  and managed by View, the maximum number of monitors must be 1 or 2. For Instant Clones, this value is inherited from snapshot of Master VM. This property has a default value of 2. &lt;br&gt; This property is required if renderer3D is set to AUTOMATIC, SOFTWARE, HARDWARE or DISABLED. | [optional] 
**MaxResolutionOfAnyOneMonitor** | Pointer to **string** | If 3D rendering is enabled and managed by View, this must be set to the default value. When 3D rendering is disabled, the max_number_of_monitors and max_resolution_of_any_one_monitor settings determine the amount of vRAM assigned to machines in this desktop. The greater these values are, the more memory will be consumed on the associated ESX hosts. This setting is only relevant on managed machines. Existing virtual machines must be powered off and subsequently powered on for the change to take effect. A restart will not cause the changes to take effect. For Instant Clones, this value is inherited from snapshot of Master VM. This property has a default value of WUXGA. &lt;br&gt; This property is required if renderer3D is set to AUTOMATIC, SOFTWARE, HARDWARE or DISABLED. * WSXGA_PLUS: 1680x1050 resolution. * WUXGA: 1920x1200 resolution. * WQXGA: 2560x1600 resolution. * UHD: 3840x2160 resolution. * UHD_5K: 5120x2880 resolution. * UHD_8K: 7680x4320 resolution. | [optional] 
**Renderer3d** | Pointer to **string** | 3D rendering is supported on Windows 7 or later guests running on VMs with virtual hardware version 8 or later. The default_display_protocol must set to PCOIP and allow_users_to_choose_protocol must be set to false to enable 3D rendering. For instant clone source desktop 3D rendering always mapped to MANAGE_BY_VSPHERE_CLIENT. Default value is DISABLED. * MANAGE_BY_VSPHERE_CLIENT: 3D rendering managed by vSphere Client. * AUTOMATIC: 3D rendering is automatic. * SOFTWARE: 3D rendering is software dependent. The software renderer is supported (at minimum) on virtual hardware version 8 in a vSphere 5.0 environment. * HARDWARE: 3D rendering is hardware dependent. The hardware-based renderer is supported (at minimum) on virtual hardware version 9 in a vSphere 5.1 environment. * DISABLED: 3D rendering is disabled. | [optional] 
**SessionCollaborationEnabled** | Pointer to **bool** | Applicable To: Automated and Manual pools with default value of false. &lt;br&gt; Enable session collaboration feature. Session collaboration allows a user to share their remote session with other users. BLAST must be configured as a supported protocol in supported_display_protocols. | [optional] 
**VramSizeMb** | Pointer to **int32** | vRAM size for View managed 3D rendering. More VRAM can improve 3D performance. Size is in MB. On ESXi 5.0 hosts, the renderer allows a maximum VRAM size of 128MB. On ESXi 5.1 and later hosts, the maximum VRAM size is 512MB. For Instant Clones, this value is inherited from snapshot of Master VM. This property has a default value of 96. &lt;br&gt;This property is required if renderer3d is set to AUTOMATIC, SOFTWARE or HARDWARE. | [optional] 

## Methods

### NewDesktopPoolDisplayProtocolSettingsCreateSpec

`func NewDesktopPoolDisplayProtocolSettingsCreateSpec() *DesktopPoolDisplayProtocolSettingsCreateSpec`

NewDesktopPoolDisplayProtocolSettingsCreateSpec instantiates a new DesktopPoolDisplayProtocolSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolDisplayProtocolSettingsCreateSpecWithDefaults

`func NewDesktopPoolDisplayProtocolSettingsCreateSpecWithDefaults() *DesktopPoolDisplayProtocolSettingsCreateSpec`

NewDesktopPoolDisplayProtocolSettingsCreateSpecWithDefaults instantiates a new DesktopPoolDisplayProtocolSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsersToChooseProtocol

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.

### HasAllowUsersToChooseProtocol

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasAllowUsersToChooseProtocol() bool`

HasAllowUsersToChooseProtocol returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.

### HasDefaultDisplayProtocol

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasDefaultDisplayProtocol() bool`

HasDefaultDisplayProtocol returns a boolean if a field has been set.

### GetGridVgpusEnabled

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetGridVgpusEnabled() bool`

GetGridVgpusEnabled returns the GridVgpusEnabled field if non-nil, zero value otherwise.

### GetGridVgpusEnabledOk

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetGridVgpusEnabledOk() (*bool, bool)`

GetGridVgpusEnabledOk returns a tuple with the GridVgpusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridVgpusEnabled

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetGridVgpusEnabled(v bool)`

SetGridVgpusEnabled sets GridVgpusEnabled field to given value.

### HasGridVgpusEnabled

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasGridVgpusEnabled() bool`

HasGridVgpusEnabled returns a boolean if a field has been set.

### GetMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetMaxNumberOfMonitors() int32`

GetMaxNumberOfMonitors returns the MaxNumberOfMonitors field if non-nil, zero value otherwise.

### GetMaxNumberOfMonitorsOk

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetMaxNumberOfMonitorsOk() (*int32, bool)`

GetMaxNumberOfMonitorsOk returns a tuple with the MaxNumberOfMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetMaxNumberOfMonitors(v int32)`

SetMaxNumberOfMonitors sets MaxNumberOfMonitors field to given value.

### HasMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasMaxNumberOfMonitors() bool`

HasMaxNumberOfMonitors returns a boolean if a field has been set.

### GetMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetMaxResolutionOfAnyOneMonitor() string`

GetMaxResolutionOfAnyOneMonitor returns the MaxResolutionOfAnyOneMonitor field if non-nil, zero value otherwise.

### GetMaxResolutionOfAnyOneMonitorOk

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetMaxResolutionOfAnyOneMonitorOk() (*string, bool)`

GetMaxResolutionOfAnyOneMonitorOk returns a tuple with the MaxResolutionOfAnyOneMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetMaxResolutionOfAnyOneMonitor(v string)`

SetMaxResolutionOfAnyOneMonitor sets MaxResolutionOfAnyOneMonitor field to given value.

### HasMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasMaxResolutionOfAnyOneMonitor() bool`

HasMaxResolutionOfAnyOneMonitor returns a boolean if a field has been set.

### GetRenderer3d

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetRenderer3d() string`

GetRenderer3d returns the Renderer3d field if non-nil, zero value otherwise.

### GetRenderer3dOk

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetRenderer3dOk() (*string, bool)`

GetRenderer3dOk returns a tuple with the Renderer3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer3d

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetRenderer3d(v string)`

SetRenderer3d sets Renderer3d field to given value.

### HasRenderer3d

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasRenderer3d() bool`

HasRenderer3d returns a boolean if a field has been set.

### GetSessionCollaborationEnabled

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.

### HasSessionCollaborationEnabled

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasSessionCollaborationEnabled() bool`

HasSessionCollaborationEnabled returns a boolean if a field has been set.

### GetVramSizeMb

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetVramSizeMb() int32`

GetVramSizeMb returns the VramSizeMb field if non-nil, zero value otherwise.

### GetVramSizeMbOk

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) GetVramSizeMbOk() (*int32, bool)`

GetVramSizeMbOk returns a tuple with the VramSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVramSizeMb

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) SetVramSizeMb(v int32)`

SetVramSizeMb sets VramSizeMb field to given value.

### HasVramSizeMb

`func (o *DesktopPoolDisplayProtocolSettingsCreateSpec) HasVramSizeMb() bool`

HasVramSizeMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


