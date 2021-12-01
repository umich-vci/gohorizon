# DesktopPoolDisplayProtocolSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | **bool** | Indicates whether the users can choose the protocol. | 
**DefaultDisplayProtocol** | **string** | The default display protocol for the desktop pool. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | 
**MaxNumberOfMonitors** | Pointer to **int32** | When renderer3d is disabled, the max_number_of_monitors and max_resolution_of_any_one_monitor&#39; settings determine the amount ofvRAM assigned to machines in this desktop pool. The greater these values are, the more memory will be consumeon the associated ESX hosts. Existing virtual machines must be powered off and subsequently powered onfor the change to take effect. A restart will not cause the changes to take effect. If 3D is enabled and managedby Horizon, the maximum number of monitors must be 1 or 2. For Instant Clones, this value is inherited from snapshot of Master VM. This property is required if renderer3D is set to AUTOMATIC, SOFTWARE, HARDWARE, or DISABLED. | [optional] 
**MaxResolutionOfAnyOneMonitor** | Pointer to **string** | If 3D rendering is enabled and managed by Horizon, this must be set to the default value. When 3D rendering is disabled,the max_number_of_monitors and max_resolution_of_any_one_monitor&#39; settings determine the amount of vRAM assignedto machines in this desktop pool. The greater these values are, the more memory will be consumed on the associated ESX hosts.This setting is only relevant on managed machines. Existing virtual machines must be powered off and subsequently poweredon for the change to take effect. A restart will not cause the changes to take effect. For Instant Clones,this value is inherited from snapshot of Master VM. This property has a default value of WUXGA.&lt;br&gt; This property is required if renderer3D is set to AUTOMATIC, SOFTWARE, HARDWARE or DISABLED. * WSXGA_PLUS: 1680x1050 resolution. * WUXGA: 1920x1200 resolution. * WQXGA: 2560x1600 resolution. * UHD: 3840x2160 resolution. * UHD_5K: 5120x2880 resolution. * UHD_8K: 7680x4320 resolution. | [optional] 
**Renderer3d** | Pointer to **string** | 3D rendering is supported on Windows 7 or later guests running on VMs with virtual hardware version8 or later. The default_display_protocol must set to PCOIP or BLAST and allow_users_to_choose_protocol must be set to false to enable 3D rendering protocol to enable 3D rendering. For instant clone source desktop pool 3D rendering always mapped to MANAGE_BY_VSPHERE_CLIENT. * MANAGE_BY_VSPHERE_CLIENT: 3D rendering managed by vSphere Client. * AUTOMATIC: 3D rendering is automatic. * SOFTWARE: 3D rendering is software dependent. The software renderer is supported (at minimum) on virtual hardware version 8 in a vSphere 5.0 environment. * HARDWARE: 3D rendering is hardware dependent. The hardware-based renderer is supported (at minimum) on virtual hardware version 9 in a vSphere 5.1 environment. * DISABLED: 3D rendering is disabled. | [optional] 
**SessionCollaborationEnabled** | Pointer to **bool** | Applicable To: Automated and Manual pools.&lt;br&gt;Enable session collaboration feature. Session collaborationallows a user to share their remote session with other users.BLAST must be configured as a supported protocol in supported_display_protocols. | [optional] 
**VramSizeMb** | Pointer to **int32** | vRAM size for Horizon managed 3D rendering. More VRAM can improve 3D performance. Size is in MB. On ESXi 5.0 hosts, the renderer allows a maximum VRAM size of 128MB. On ESXi 5.1 and later hosts, the maximum VRAM size is 512MB. For Instant Clones, this value is inherited from snapshot of Master VM.This property is applicable when 3D renderer is not disabled. This property is required if renderer3d is set to AUTOMATIC, SOFTWARE, or HARDWARE. | [optional] 

## Methods

### NewDesktopPoolDisplayProtocolSettingsUpdateSpec

`func NewDesktopPoolDisplayProtocolSettingsUpdateSpec(allowUsersToChooseProtocol bool, defaultDisplayProtocol string, ) *DesktopPoolDisplayProtocolSettingsUpdateSpec`

NewDesktopPoolDisplayProtocolSettingsUpdateSpec instantiates a new DesktopPoolDisplayProtocolSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolDisplayProtocolSettingsUpdateSpecWithDefaults

`func NewDesktopPoolDisplayProtocolSettingsUpdateSpecWithDefaults() *DesktopPoolDisplayProtocolSettingsUpdateSpec`

NewDesktopPoolDisplayProtocolSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolDisplayProtocolSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsersToChooseProtocol

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.


### GetDefaultDisplayProtocol

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.


### GetMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetMaxNumberOfMonitors() int32`

GetMaxNumberOfMonitors returns the MaxNumberOfMonitors field if non-nil, zero value otherwise.

### GetMaxNumberOfMonitorsOk

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetMaxNumberOfMonitorsOk() (*int32, bool)`

GetMaxNumberOfMonitorsOk returns a tuple with the MaxNumberOfMonitors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) SetMaxNumberOfMonitors(v int32)`

SetMaxNumberOfMonitors sets MaxNumberOfMonitors field to given value.

### HasMaxNumberOfMonitors

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) HasMaxNumberOfMonitors() bool`

HasMaxNumberOfMonitors returns a boolean if a field has been set.

### GetMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetMaxResolutionOfAnyOneMonitor() string`

GetMaxResolutionOfAnyOneMonitor returns the MaxResolutionOfAnyOneMonitor field if non-nil, zero value otherwise.

### GetMaxResolutionOfAnyOneMonitorOk

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetMaxResolutionOfAnyOneMonitorOk() (*string, bool)`

GetMaxResolutionOfAnyOneMonitorOk returns a tuple with the MaxResolutionOfAnyOneMonitor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) SetMaxResolutionOfAnyOneMonitor(v string)`

SetMaxResolutionOfAnyOneMonitor sets MaxResolutionOfAnyOneMonitor field to given value.

### HasMaxResolutionOfAnyOneMonitor

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) HasMaxResolutionOfAnyOneMonitor() bool`

HasMaxResolutionOfAnyOneMonitor returns a boolean if a field has been set.

### GetRenderer3d

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetRenderer3d() string`

GetRenderer3d returns the Renderer3d field if non-nil, zero value otherwise.

### GetRenderer3dOk

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetRenderer3dOk() (*string, bool)`

GetRenderer3dOk returns a tuple with the Renderer3d field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenderer3d

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) SetRenderer3d(v string)`

SetRenderer3d sets Renderer3d field to given value.

### HasRenderer3d

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) HasRenderer3d() bool`

HasRenderer3d returns a boolean if a field has been set.

### GetSessionCollaborationEnabled

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.

### HasSessionCollaborationEnabled

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) HasSessionCollaborationEnabled() bool`

HasSessionCollaborationEnabled returns a boolean if a field has been set.

### GetVramSizeMb

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetVramSizeMb() int32`

GetVramSizeMb returns the VramSizeMb field if non-nil, zero value otherwise.

### GetVramSizeMbOk

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) GetVramSizeMbOk() (*int32, bool)`

GetVramSizeMbOk returns a tuple with the VramSizeMb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVramSizeMb

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) SetVramSizeMb(v int32)`

SetVramSizeMb sets VramSizeMb field to given value.

### HasVramSizeMb

`func (o *DesktopPoolDisplayProtocolSettingsUpdateSpec) HasVramSizeMb() bool`

HasVramSizeMb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


