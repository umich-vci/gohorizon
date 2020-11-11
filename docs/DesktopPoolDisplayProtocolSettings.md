# DesktopPoolDisplayProtocolSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | **bool** | Indicates whether the users can choose the protocol. Default value is true. | 
**DefaultDisplayProtocol** | **string** | The default display protocol for the desktop pool. For a managed desktop pool, this will default to PCOIP.For an unmanaged desktop pool, this will default to RDP. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | 
**DisplayProtocols** | **[]string** | List of supported display protocols for this desktop pool.Default value is [PCOIP, RDP, BLAST]. | 
**GridVgpusEnabled** | **bool** | When 3D rendering is managed by the vSphere Client, this enables support for NVIDIA GRID vGPUs.This will be false if 3D rendering is not managed by the vSphere Client. If this is true,the host or cluster associated with the desktop pool must support NVIDIA GRID and vGPU types required by the desktop pool&#39;s VirtualMachines,VmTemplate, or BaseImageSnapshot. If this is false, the desktop pool&#39;s VirtualMachines, VmTemplate, orBaseImageSnapshot must not support NVIDIA GRID vGPUs. Since suspending VMs with passthroughdevices such as vGPUs is not possible, power_policy cannot be set to SUSPEND if this is enabled.Default value is false. | [optional] 
**HtmlAccessEnabled** | **bool** | This property is no longer in use for Horizon Components. It is always set to true. HTML Access, enabled by VMware Blast technology, allows users to connect to machines from Web browsers. Horizon Client software does not have to be installed on the client devices. To enable HTML Access, you must install the HTML Machine Access feature pack. Also, Blast must be configured as a supported protocol in displayProtocols. | 
**MaxNumberOfMonitors** | **int32** | When 3D is disabled, the &#39;Max number of monitors&#39; and &#39;Max resolution of any one monitor&#39; settings determine the amount ofvRAM assigned to machines in this desktop. The greater these values are, the more memory will be consumeon the associated ESX hosts. Existing virtual machines must be powered off and subsequently powered onfor the change to take effect. A restart will not cause the changes to take effect. If 3D is enabled and managedby View, the maximum number of monitors must be 1 or 2. For Instant Clones, this value is inherited from snapshot of Master VM. | [optional] 
**MaxResolutionOfAnyOneMonitor** | **string** | If 3D rendering is enabled and managed by View, this must be set to the default value. When 3D rendering is disabled,the &#39;Max number of monitors&#39; and &#39;Max resolution of any one monitor&#39; settings determine the amount of vRAM assignedto machines in this desktop. The greater these values are, the more memory will be consumed on the associated ESX hosts.This setting is only relevant on managed machines. Existing virtual machines must be powered off and subsequently poweredon for the change to take effect. A restart will not cause the changes to take effect. For Instant Clones,this value is inherited from snapshot of Master VM. This property has a default value of WUXGA. * WSXGA_PLUS: 1680x1050 resolution. * WUXGA: 1920x1200 resolution. * WQXGA: 2560x1600 resolution. * UHD: 3840x2160 resolution. | [optional] 
**Renderer3d** | **string** | 3D rendering is supported on Windows 7 or later guests running on VMs with virtual hardware version8 or later. The default protocol must be PCoIP and users must not be allowed to choose their ownprotocol to enable 3D rendering. For instant clone source desktop 3D rendering always mapped to MANAGE_BY_VSPHERE_CLIENT.Default value is DISABLED. * MANAGE_BY_VSPHERE_CLIENT: 3D rendering managed by vSphere Client. * AUTOMATIC: 3D rendering is automatic. * SOFTWARE: 3D rendering is software dependent. The software renderer is supported (at minimum) on virtual hardware version 8 in a vSphere 5.0 environment. * HARDWARE: 3D rendering is hardware dependent. The hardware-based renderer is supported (at minimum) on virtual hardware version 9 in a vSphere 5.1 environment. * DISABLED: 3D rendering is disabled. | 
**SessionCollaborationEnabled** | **bool** | Enable session collaboration feature. Session collaborationallows a user to share their remote session with other users.BLAST must be configured as a supported protocol in supported_display_protocols.Default value is false. | 
**VgpuGridProfile** | **string** | NVIDIA GRID vGPUs might have multiple profiles and any one of the available profiles can beapplied to newly created instant clone desktop. The profile specified in this field will beused in the newly created instant clone desktop. Will be set if enable_grid_vgpus set to true. | [optional] 
**VramSizeMb** | **int32** | vRAM size for View managed 3D rendering. More VRAM can improve 3D performance.Size is in MB. On ESXi 5.0 hosts, the renderer allows a maximum VRAM size of 128MB. On ESXi 5.1and later hosts, the maximum VRAM size is 512MB. For Instant Clones, this value is inherited from snapshot of Master VM.This property is applicable when 3D renderer is not disabled. This has a default value of 96. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

