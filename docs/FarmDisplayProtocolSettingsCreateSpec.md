# FarmDisplayProtocolSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowUsersToChooseProtocol** | Pointer to **bool** | Indicates whether the users can choose the protocol. Default value is true. | [optional] 
**DefaultDisplayProtocol** | Pointer to **string** | Indicates default server display protocol, when user is not allowed to choose protocol.Farms support PCOIP, RDP and BLAST. Default value is PCOIP. * RDP: Microsoft Remote Desktop Protocol. * PCOIP: PCoIP protocol. * BLAST: BLAST protocol. | [optional] 
**GridVgpusEnabled** | Pointer to **bool** | Indicates whether RDSH instant clone farm must not support NVIDIA GRID vGPUs. If this is true, the host or cluster associated with the farm must support NVIDIA GRID and vGPU types required by the RDSH desktop pool virtualMachines, VmTemplate, or BaseImageSnapshot. If this is false, RDSH instant clone farm must not support NVIDIA GRID vGPUs. This property is only applicable to instant clone farm. Default value is false. | [optional] 
**SessionCollaborationEnabled** | Pointer to **bool** | Indicates whether session collaboration feature is enabled. Session collaboration allows a user to share their remote session with other users. Default value is false. | [optional] 

## Methods

### NewFarmDisplayProtocolSettingsCreateSpec

`func NewFarmDisplayProtocolSettingsCreateSpec() *FarmDisplayProtocolSettingsCreateSpec`

NewFarmDisplayProtocolSettingsCreateSpec instantiates a new FarmDisplayProtocolSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmDisplayProtocolSettingsCreateSpecWithDefaults

`func NewFarmDisplayProtocolSettingsCreateSpecWithDefaults() *FarmDisplayProtocolSettingsCreateSpec`

NewFarmDisplayProtocolSettingsCreateSpecWithDefaults instantiates a new FarmDisplayProtocolSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowUsersToChooseProtocol

`func (o *FarmDisplayProtocolSettingsCreateSpec) GetAllowUsersToChooseProtocol() bool`

GetAllowUsersToChooseProtocol returns the AllowUsersToChooseProtocol field if non-nil, zero value otherwise.

### GetAllowUsersToChooseProtocolOk

`func (o *FarmDisplayProtocolSettingsCreateSpec) GetAllowUsersToChooseProtocolOk() (*bool, bool)`

GetAllowUsersToChooseProtocolOk returns a tuple with the AllowUsersToChooseProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowUsersToChooseProtocol

`func (o *FarmDisplayProtocolSettingsCreateSpec) SetAllowUsersToChooseProtocol(v bool)`

SetAllowUsersToChooseProtocol sets AllowUsersToChooseProtocol field to given value.

### HasAllowUsersToChooseProtocol

`func (o *FarmDisplayProtocolSettingsCreateSpec) HasAllowUsersToChooseProtocol() bool`

HasAllowUsersToChooseProtocol returns a boolean if a field has been set.

### GetDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettingsCreateSpec) GetDefaultDisplayProtocol() string`

GetDefaultDisplayProtocol returns the DefaultDisplayProtocol field if non-nil, zero value otherwise.

### GetDefaultDisplayProtocolOk

`func (o *FarmDisplayProtocolSettingsCreateSpec) GetDefaultDisplayProtocolOk() (*string, bool)`

GetDefaultDisplayProtocolOk returns a tuple with the DefaultDisplayProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettingsCreateSpec) SetDefaultDisplayProtocol(v string)`

SetDefaultDisplayProtocol sets DefaultDisplayProtocol field to given value.

### HasDefaultDisplayProtocol

`func (o *FarmDisplayProtocolSettingsCreateSpec) HasDefaultDisplayProtocol() bool`

HasDefaultDisplayProtocol returns a boolean if a field has been set.

### GetGridVgpusEnabled

`func (o *FarmDisplayProtocolSettingsCreateSpec) GetGridVgpusEnabled() bool`

GetGridVgpusEnabled returns the GridVgpusEnabled field if non-nil, zero value otherwise.

### GetGridVgpusEnabledOk

`func (o *FarmDisplayProtocolSettingsCreateSpec) GetGridVgpusEnabledOk() (*bool, bool)`

GetGridVgpusEnabledOk returns a tuple with the GridVgpusEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGridVgpusEnabled

`func (o *FarmDisplayProtocolSettingsCreateSpec) SetGridVgpusEnabled(v bool)`

SetGridVgpusEnabled sets GridVgpusEnabled field to given value.

### HasGridVgpusEnabled

`func (o *FarmDisplayProtocolSettingsCreateSpec) HasGridVgpusEnabled() bool`

HasGridVgpusEnabled returns a boolean if a field has been set.

### GetSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettingsCreateSpec) GetSessionCollaborationEnabled() bool`

GetSessionCollaborationEnabled returns the SessionCollaborationEnabled field if non-nil, zero value otherwise.

### GetSessionCollaborationEnabledOk

`func (o *FarmDisplayProtocolSettingsCreateSpec) GetSessionCollaborationEnabledOk() (*bool, bool)`

GetSessionCollaborationEnabledOk returns a tuple with the SessionCollaborationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettingsCreateSpec) SetSessionCollaborationEnabled(v bool)`

SetSessionCollaborationEnabled sets SessionCollaborationEnabled field to given value.

### HasSessionCollaborationEnabled

`func (o *FarmDisplayProtocolSettingsCreateSpec) HasSessionCollaborationEnabled() bool`

HasSessionCollaborationEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


