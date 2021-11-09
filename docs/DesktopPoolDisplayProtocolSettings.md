# DesktopPoolDisplayProtocolSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | **bool** | Indicates whether the users can choose the protocol. Default value is true. | 
**DefaultDisplayProtocol** | **string** | The default display protocol for the desktop pool. For a managed desktop pool, this will default to PCOIP.For an unmanaged desktop pool, this will default to RDP. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | 
**DisplayProtocols** | **[]string** | List of supported display protocols for this desktop pool.Default value is [PCOIP, RDP, BLAST]. | 
**GridVgpusEnabled** | Pointer to **bool** | When 3D rendering is managed by the vSphere Client, this enables support for NVIDIA GRID vGPUs.This will be false if 3D rendering is not managed by the vSphere Client. If this is true,the host or cluster associated with the desktop pool must support NVIDIA GRID and vGPU types required by the desktop pool&#39;s VirtualMachines,VmTemplate, or BaseImageSnapshot. If this is false, the desktop pool&#39;s VirtualMachines, VmTemplate, orBaseImageSnapshot must not support NVIDIA GRID vGPUs. Since suspending VMs with passthroughdevices such as vGPUs is not possible, power_policy cannot be set to SUSPEND if this is enabled.Default value is false. | [optional] 
**HtmlAccessEnabled** | **bool** | This property is no longer in use for Horizon Components. It is always set to true. HTML Access, enabled by VMware Blast technology, allows users to connect to machines from Web browsers. Horizon Client software does not have to be installed on the client devices. To enable HTML Access, you must install the HTML Machine Access feature pack. Also, Blast must be configured as a supported protocol in displayProtocols. | 
**MaxNumberOfMonitors** | Pointer to **int32** | When 3D is disabled, the &#39;Max number of monitors&#39; and &#39;Max resolution of any one monitor&#39; settings determine the amount ofvRAM assigned to machines in this desktop. The greater these values are, the more memory will be consumeon the associated ESX hosts. Existing virtual machines must be powered off and subsequently powered onfor the change to take effect. A restart will not cause the changes to take effect. If 3D is enabled and managedby View, the maximum number of monitors must be 1 or 2. For Instant Clones, this value is inherited from snapshot of Master VM. | [optional] 
**MaxResolutionOfAnyOneMonitor** | Pointer to **string** | If 3D rendering is enabled and managed by View, this must be set to the default value. When 3D rendering is disabled,the &#39;Max number of monitors&#39; and &#39;Max resolution of any one monitor&#39; settings determine the amount of vRAM assignedto machines in this desktop. The greater these values are, the more memory will be consumed on the associated ESX hosts.This setting is only relevant on managed machines. Existing virtual machines must be powered off and subsequently poweredon for the change to take effect. A restart will not cause the changes to take effect. For Instant Clones,this value is inherited from snapshot of Master VM. This property has a default value of WUXGA. * WSXGA_PLUS: 1680x1050 resolution. * WUXGA: 1920x1200 resolution. * WQXGA: 2560x1600 resolution. * UHD: 3840x2160 resolution. | [optional] 
**Renderer3d** | **string** | 3D rendering is supported on Windows 7 or later guests running on VMs with virtual hardware version8 or later. The default protocol must be PCoIP and users must not be allowed to choose their ownprotocol to enable 3D rendering. For instant clone source desktop 3D rendering always mapped to MANAGE_BY_VSPHERE_CLIENT.Default value is DISABLED. * MANAGE_BY_VSPHERE_CLIENT: 3D rendering managed by vSphere Client. * AUTOMATIC: 3D rendering is automatic. * SOFTWARE: 3D rendering is software dependent. The software renderer is supported (at minimum) on virtual hardware version 8 in a vSphere 5.0 environment. * HARDWARE: 3D rendering is hardware dependent. The hardware-based renderer is supported (at minimum) on virtual hardware version 9 in a vSphere 5.1 environment. * DISABLED: 3D rendering is disabled. | 
**SessionCollaborationEnabled** | **bool** | Enable session collaboration feature. Session collaborationallows a user to share their remote session with other users.BLAST must be configured as a supported protocol in supported_display_protocols.Default value is false. | 
**VgpuGridProfile** | Pointer to **string** | NVIDIA GRID vGPUs might have multiple profiles and any one of the available profiles can beapplied to newly created instant clone desktop. The profile specified in this field will beused in the newly created instant clone desktop. Will be set if enable_grid_vgpus set to true. | [optional] 
**VramSizeMb** | **int32** | vRAM size for View managed 3D rendering. More VRAM can improve 3D performance.Size is in MB. On ESXi 5.0 hosts, the renderer allows a maximum VRAM size of 128MB. On ESXi 5.1and later hosts, the maximum VRAM size is 512MB. For Instant Clones, this value is inherited from snapshot of Master VM.This property is applicable when 3D renderer is not disabled. This has a default value of 96. | 

## Methods

### NewDesktopPoolDisplayProtocolSettings

`func NewDesktopPoolDisplayProtocolSettings(allowUsersToChooseProtocol bool, defaultDisplayProtocol string, displayProtocols []string, htmlAccessEnabled bool, renderer3d string, sessionCollaborationEnabled bool, vramSizeMb int32, ) *DesktopPoolDisplayProtocolSettings`

NewDesktopPoolDisplayProtocolSettings instantiates a new DesktopPoolDisplayProtocolSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolDisplayProtocolSettingsWithDefaults

`func NewDesktopPoolDisplayProtocolSettingsWithDefaults() *DesktopPoolDisplayProtocolSettings`

NewDesktopPoolDisplayProtocolSettingsWithDefaults instantiates a new DesktopPoolDisplayProtocolSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsersToChooseProtocol

`func (o *DesktopPoolDisplayProtocolSettings) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *DesktopPoolDisplayProtocolSettings) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *DesktopPoolDisplayProtocolSettings) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.


### GetDefaultDisplayProtocol

`func (o *DesktopPoolDisplayProtocolSettings) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *DesktopPoolDisplayProtocolSettings) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *DesktopPoolDisplayProtocolSettings) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.


### GetDisplayProtocols

`func (o *DesktopPoolDisplayProtocolSettings) GetDisplayProtocols() []string`

GetDisplayProtocols returns the DisplayProtocols field if non-nil, zero value otherwise.

### GetDisplayProtocolsOk

`func (o *DesktopPoolDisplayProtocolSettings) GetDisplayProtocolsOk() (*[]string, bool)`

GetDisplayProtocolsOk returns a tuple with the DisplayProtocols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayProtocols

`func (o *DesktopPoolDisplayProtocolSettings) SetDisplayProtocols(v []string)`

SetDisplayProtocols sets DisplayProtocols field to given value.


### GetGridVgpusEnabled

`func (o *DesktopPoolDisplayProtocolSettings) GetGridVgpusEnabled() bool`

GetGridVgpusEnabled returns the GridVgpusEnabled field if non-nil, zero value otherwise.

### GetGridVgpusEnabledOk

`func (o *DesktopPoolDisplayProtocolSettings) GetGridVgpusEnabledOk() (*bool, bool)`

GetGridVgpusEnabledOk returns a tuple with the GridVgpusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridVgpusEnabled

`func (o *DesktopPoolDisplayProtocolSettings) SetGridVgpusEnabled(v bool)`

SetGridVgpusEnabled sets GridVgpusEnabled field to given value.

### HasGridVgpusEnabled

`func (o *DesktopPoolDisplayProtocolSettings) HasGridVgpusEnabled() bool`

HasGridVgpusEnabled returns a boolean if a field has been set.

### GetHtmlAccessEnabled

`func (o *DesktopPoolDisplayProtocolSettings) GetHtmlAccessEnabled() bool`

GetHtmlAccessEnabled returns the HtmlAccessEnabled field if non-nil, zero value otherwise.

### GetHtmlAccessEnabledOk

`func (o *DesktopPoolDisplayProtocolSettings) GetHtmlAccessEnabledOk() (*bool, bool)`

GetHtmlAccessEnabledOk returns a tuple with the HtmlAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlAccessEnabled

`func (o *DesktopPoolDisplayProtocolSettings) SetHtmlAccessEnabled(v bool)`

SetHtmlAccessEnabled sets HtmlAccessEnabled field to given value.


### GetMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettings) GetMaxNumberOfMonitors() int32`

GetMaxNumberOfMonitors returns the MaxNumberOfMonitors field if non-nil, zero value otherwise.

### GetMaxNumberOfMonitorsOk

`func (o *DesktopPoolDisplayProtocolSettings) GetMaxNumberOfMonitorsOk() (*int32, bool)`

GetMaxNumberOfMonitorsOk returns a tuple with the MaxNumberOfMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettings) SetMaxNumberOfMonitors(v int32)`

SetMaxNumberOfMonitors sets MaxNumberOfMonitors field to given value.

### HasMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettings) HasMaxNumberOfMonitors() bool`

HasMaxNumberOfMonitors returns a boolean if a field has been set.

### GetMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettings) GetMaxResolutionOfAnyOneMonitor() string`

GetMaxResolutionOfAnyOneMonitor returns the MaxResolutionOfAnyOneMonitor field if non-nil, zero value otherwise.

### GetMaxResolutionOfAnyOneMonitorOk

`func (o *DesktopPoolDisplayProtocolSettings) GetMaxResolutionOfAnyOneMonitorOk() (*string, bool)`

GetMaxResolutionOfAnyOneMonitorOk returns a tuple with the MaxResolutionOfAnyOneMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettings) SetMaxResolutionOfAnyOneMonitor(v string)`

SetMaxResolutionOfAnyOneMonitor sets MaxResolutionOfAnyOneMonitor field to given value.

### HasMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettings) HasMaxResolutionOfAnyOneMonitor() bool`

HasMaxResolutionOfAnyOneMonitor returns a boolean if a field has been set.

### GetRenderer3d

`func (o *DesktopPoolDisplayProtocolSettings) GetRenderer3d() string`

GetRenderer3d returns the Renderer3d field if non-nil, zero value otherwise.

### GetRenderer3dOk

`func (o *DesktopPoolDisplayProtocolSettings) GetRenderer3dOk() (*string, bool)`

GetRenderer3dOk returns a tuple with the Renderer3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer3d

`func (o *DesktopPoolDisplayProtocolSettings) SetRenderer3d(v string)`

SetRenderer3d sets Renderer3d field to given value.


### GetSessionCollaborationEnabled

`func (o *DesktopPoolDisplayProtocolSettings) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *DesktopPoolDisplayProtocolSettings) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *DesktopPoolDisplayProtocolSettings) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.


### GetVgpuGridProfile

`func (o *DesktopPoolDisplayProtocolSettings) GetVgpuGridProfile() string`

GetVgpuGridProfile returns the VgpuGridProfile field if non-nil, zero value otherwise.

### GetVgpuGridProfileOk

`func (o *DesktopPoolDisplayProtocolSettings) GetVgpuGridProfileOk() (*string, bool)`

GetVgpuGridProfileOk returns a tuple with the VgpuGridProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuGridProfile

`func (o *DesktopPoolDisplayProtocolSettings) SetVgpuGridProfile(v string)`

SetVgpuGridProfile sets VgpuGridProfile field to given value.

### HasVgpuGridProfile

`func (o *DesktopPoolDisplayProtocolSettings) HasVgpuGridProfile() bool`

HasVgpuGridProfile returns a boolean if a field has been set.

### GetVramSizeMb

`func (o *DesktopPoolDisplayProtocolSettings) GetVramSizeMb() int32`

GetVramSizeMb returns the VramSizeMb field if non-nil, zero value otherwise.

### GetVramSizeMbOk

`func (o *DesktopPoolDisplayProtocolSettings) GetVramSizeMbOk() (*int32, bool)`

GetVramSizeMbOk returns a tuple with the VramSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVramSizeMb

`func (o *DesktopPoolDisplayProtocolSettings) SetVramSizeMb(v int32)`

SetVramSizeMb sets VramSizeMb field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


