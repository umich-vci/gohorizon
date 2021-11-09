# FarmDisplayProtocolSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | Pointer to **bool** | Indicates whether the users can choose the protocol. | [optional] 
**DefaultDisplayProtocol** | Pointer to **string** | Indicates default server display protocol, when user is not allowed to choose protocol.Farms support PCOIP, RDP and BLAST. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | [optional] 
**GridVgpusEnabled** | Pointer to **bool** | Indicates whether RDSH instant clone farm must not support NVIDIA GRID vGPUs. If this is true, the host or cluster associated with the farm must support NVIDIA GRID and vGPU types required by the RDSH desktop pool virtualMachines, VmTemplate, or BaseImageSnapshot. If this is false, RDSH instant clone farm must not support NVIDIA GRID vGPUs. | [optional] 
**SessionCollaborationEnabled** | Pointer to **bool** | Indicates whether session collaboration feature is enabled. Session collaboration allows a user to share their remote session with other users. | [optional] 
**VgpuGridProfile** | Pointer to **string** | NVIDIA GRID vGPUs have multiple profiles and any one of the available profiles can be applied to newly created instant clone RDSH desktop pool. The profile specified in this field will be used in the newly created instant clone RDSH server. This is set when grid_vgpus_enabled is true. | [optional] 

## Methods

### NewFarmDisplayProtocolSettingsInfo

`func NewFarmDisplayProtocolSettingsInfo() *FarmDisplayProtocolSettingsInfo`

NewFarmDisplayProtocolSettingsInfo instantiates a new FarmDisplayProtocolSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmDisplayProtocolSettingsInfoWithDefaults

`func NewFarmDisplayProtocolSettingsInfoWithDefaults() *FarmDisplayProtocolSettingsInfo`

NewFarmDisplayProtocolSettingsInfoWithDefaults instantiates a new FarmDisplayProtocolSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsersToChooseProtocol

`func (o *FarmDisplayProtocolSettingsInfo) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *FarmDisplayProtocolSettingsInfo) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *FarmDisplayProtocolSettingsInfo) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.

### HasAllowUsersToChooseProtocol

`func (o *FarmDisplayProtocolSettingsInfo) HasAllowUsersToChooseProtocol() bool`

HasAllowUsersToChooseProtocol returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettingsInfo) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *FarmDisplayProtocolSettingsInfo) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettingsInfo) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.

### HasDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettingsInfo) HasDefaultDisplayProtocol() bool`

HasDefaultDisplayProtocol returns a boolean if a field has been set.

### GetGridVgpusEnabled

`func (o *FarmDisplayProtocolSettingsInfo) GetGridVgpusEnabled() bool`

GetGridVgpusEnabled returns the GridVgpusEnabled field if non-nil, zero value otherwise.

### GetGridVgpusEnabledOk

`func (o *FarmDisplayProtocolSettingsInfo) GetGridVgpusEnabledOk() (*bool, bool)`

GetGridVgpusEnabledOk returns a tuple with the GridVgpusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridVgpusEnabled

`func (o *FarmDisplayProtocolSettingsInfo) SetGridVgpusEnabled(v bool)`

SetGridVgpusEnabled sets GridVgpusEnabled field to given value.

### HasGridVgpusEnabled

`func (o *FarmDisplayProtocolSettingsInfo) HasGridVgpusEnabled() bool`

HasGridVgpusEnabled returns a boolean if a field has been set.

### GetSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettingsInfo) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *FarmDisplayProtocolSettingsInfo) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettingsInfo) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.

### HasSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettingsInfo) HasSessionCollaborationEnabled() bool`

HasSessionCollaborationEnabled returns a boolean if a field has been set.

### GetVgpuGridProfile

`func (o *FarmDisplayProtocolSettingsInfo) GetVgpuGridProfile() string`

GetVgpuGridProfile returns the VgpuGridProfile field if non-nil, zero value otherwise.

### GetVgpuGridProfileOk

`func (o *FarmDisplayProtocolSettingsInfo) GetVgpuGridProfileOk() (*string, bool)`

GetVgpuGridProfileOk returns a tuple with the VgpuGridProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuGridProfile

`func (o *FarmDisplayProtocolSettingsInfo) SetVgpuGridProfile(v string)`

SetVgpuGridProfile sets VgpuGridProfile field to given value.

### HasVgpuGridProfile

`func (o *FarmDisplayProtocolSettingsInfo) HasVgpuGridProfile() bool`

HasVgpuGridProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


