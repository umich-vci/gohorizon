# FarmDisplayProtocolSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowDisplayProtocolOverride** | **bool** | Indicates whether the display protocol settings could be overridden.If set to false, then default_display_protocol is used.Default value is true. | 
**DefaultDisplayProtocol** | **string** | The default server display protocol. Default value is PCOIP. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | 
**GridVgpusEnabled** | Pointer to **bool** | If this is true, the host or cluster associated with the farm must support NVIDIA GRID andvGPU types required by the RDSH desktop virtualMachines, VmTemplate, or BaseImageSnapshot.If this is false, RDSH instant clone farm must not support NVIDIA GRID vGPUs.Default value is false. | [optional] 
**HtmlAccessEnabled** | **bool** | This property is no longer in use for Horizon Components. It is always set to true. HTML Access, enabled by VMware Blast technology, allows users to connect to Horizon machines from Web browsers. Horizon Client software does not have to be installed on the client devices. To enable HTML Access, you must install the HTML Machine Access feature pack. | 
**SessionCollaborationEnabled** | **bool** | Enable session collaboration feature. Session collaborationallows a user to share their remote session with other users.BLAST must be configured as a supported protocol in supported_display_protocols.Default value is false. | 
**VgpuGridProfile** | Pointer to **string** | NVIDIA GRID vGPUs might have multiple profiles and any one of the available profiles can beapplied to newly created instant clone RDSH server. The profile specified in this field will beused in the newly created instant clone RDSH server. Will be set if enable_grid_vgpus set to true. | [optional] 

## Methods

### NewFarmDisplayProtocolSettings

`func NewFarmDisplayProtocolSettings(allowDisplayProtocolOverride bool, defaultDisplayProtocol string, htmlAccessEnabled bool, sessionCollaborationEnabled bool, ) *FarmDisplayProtocolSettings`

NewFarmDisplayProtocolSettings instantiates a new FarmDisplayProtocolSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmDisplayProtocolSettingsWithDefaults

`func NewFarmDisplayProtocolSettingsWithDefaults() *FarmDisplayProtocolSettings`

NewFarmDisplayProtocolSettingsWithDefaults instantiates a new FarmDisplayProtocolSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowDisplayProtocolOverride

`func (o *FarmDisplayProtocolSettings) GetAllowDisplayProtocolOverride() bool`

GetAllowDisplayProtocolOverride returns the AllowDisplayProtocolOverride field if non-nil, zero value otherwise.

### GetAllowDisplayProtocolOverrideOk

`func (o *FarmDisplayProtocolSettings) GetAllowDisplayProtocolOverrideOk() (*bool, bool)`

GetAllowDisplayProtocolOverrideOk returns a tuple with the AllowDisplayProtocolOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDisplayProtocolOverride

`func (o *FarmDisplayProtocolSettings) SetAllowDisplayProtocolOverride(v bool)`

SetAllowDisplayProtocolOverride sets AllowDisplayProtocolOverride field to given value.


### GetDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettings) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *FarmDisplayProtocolSettings) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettings) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.


### GetGridVgpusEnabled

`func (o *FarmDisplayProtocolSettings) GetGridVgpusEnabled() bool`

GetGridVgpusEnabled returns the GridVgpusEnabled field if non-nil, zero value otherwise.

### GetGridVgpusEnabledOk

`func (o *FarmDisplayProtocolSettings) GetGridVgpusEnabledOk() (*bool, bool)`

GetGridVgpusEnabledOk returns a tuple with the GridVgpusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridVgpusEnabled

`func (o *FarmDisplayProtocolSettings) SetGridVgpusEnabled(v bool)`

SetGridVgpusEnabled sets GridVgpusEnabled field to given value.

### HasGridVgpusEnabled

`func (o *FarmDisplayProtocolSettings) HasGridVgpusEnabled() bool`

HasGridVgpusEnabled returns a boolean if a field has been set.

### GetHtmlAccessEnabled

`func (o *FarmDisplayProtocolSettings) GetHtmlAccessEnabled() bool`

GetHtmlAccessEnabled returns the HtmlAccessEnabled field if non-nil, zero value otherwise.

### GetHtmlAccessEnabledOk

`func (o *FarmDisplayProtocolSettings) GetHtmlAccessEnabledOk() (*bool, bool)`

GetHtmlAccessEnabledOk returns a tuple with the HtmlAccessEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlAccessEnabled

`func (o *FarmDisplayProtocolSettings) SetHtmlAccessEnabled(v bool)`

SetHtmlAccessEnabled sets HtmlAccessEnabled field to given value.


### GetSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettings) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *FarmDisplayProtocolSettings) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettings) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.


### GetVgpuGridProfile

`func (o *FarmDisplayProtocolSettings) GetVgpuGridProfile() string`

GetVgpuGridProfile returns the VgpuGridProfile field if non-nil, zero value otherwise.

### GetVgpuGridProfileOk

`func (o *FarmDisplayProtocolSettings) GetVgpuGridProfileOk() (*string, bool)`

GetVgpuGridProfileOk returns a tuple with the VgpuGridProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuGridProfile

`func (o *FarmDisplayProtocolSettings) SetVgpuGridProfile(v string)`

SetVgpuGridProfile sets VgpuGridProfile field to given value.

### HasVgpuGridProfile

`func (o *FarmDisplayProtocolSettings) HasVgpuGridProfile() bool`

HasVgpuGridProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


