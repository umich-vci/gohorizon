# DesktopPoolProvisioningSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddVirtualTpm** | Pointer to **bool** | Indicates whether to add Virtual TPM device. &lt;br&gt; Default value is false. | [optional] 
**BaseSnapshotId** | Pointer to **string** | This property can be set only when source is set to INSTANT_CLONE, vm_template_id is unset and parent_vm_id is set. | [optional] 
**DatacenterId** | Pointer to **string** | Datacenter within which the desktop pool is configured. | [optional] 
**HostOrClusterId** | **string** | Host or cluster where the machines are deployed in. &lt;br&gt; For Instant clone desktops it can only be set to a cluster id. &lt;br&gt; | 
**ImStreamId** | Pointer to **string** | Applicable To: Automated desktop pools.&lt;br&gt;This is required when vm_template_id, parent_vm_id and base_snapshot_id are not set. | [optional] 
**ImTagId** | Pointer to **string** | Applicable To: Automated desktop pools.&lt;br&gt;This is required when im_stream_Id is set. | [optional] 
**ParentVmId** | Pointer to **string** | This property can be set only when source is set to INSTANT_CLONE. | [optional] 
**ResourcePoolId** | **string** | Resource pool to deploy the machines. | 
**VmFolderId** | **string** | VM folder where the machines are deployed to. | 
**VmTemplateId** | Pointer to **string** | Applicable To: Full clone desktop pools.&lt;br&gt;This is required if parent_vm_id and base_snapshot_id are not set. &lt;br&gt; | [optional] 

## Methods

### NewDesktopPoolProvisioningSettingsCreateSpec

`func NewDesktopPoolProvisioningSettingsCreateSpec(hostOrClusterId string, resourcePoolId string, vmFolderId string, ) *DesktopPoolProvisioningSettingsCreateSpec`

NewDesktopPoolProvisioningSettingsCreateSpec instantiates a new DesktopPoolProvisioningSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolProvisioningSettingsCreateSpecWithDefaults

`func NewDesktopPoolProvisioningSettingsCreateSpecWithDefaults() *DesktopPoolProvisioningSettingsCreateSpec`

NewDesktopPoolProvisioningSettingsCreateSpecWithDefaults instantiates a new DesktopPoolProvisioningSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddVirtualTpm

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetAddVirtualTpm() bool`

GetAddVirtualTpm returns the AddVirtualTpm field if non-nil, zero value otherwise.

### GetAddVirtualTpmOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetAddVirtualTpmOk() (*bool, bool)`

GetAddVirtualTpmOk returns a tuple with the AddVirtualTpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVirtualTpm

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetAddVirtualTpm(v bool)`

SetAddVirtualTpm sets AddVirtualTpm field to given value.

### HasAddVirtualTpm

`func (o *DesktopPoolProvisioningSettingsCreateSpec) HasAddVirtualTpm() bool`

HasAddVirtualTpm returns a boolean if a field has been set.

### GetBaseSnapshotId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetBaseSnapshotId() string`

GetBaseSnapshotId returns the BaseSnapshotId field if non-nil, zero value otherwise.

### GetBaseSnapshotIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetBaseSnapshotIdOk() (*string, bool)`

GetBaseSnapshotIdOk returns a tuple with the BaseSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetBaseSnapshotId(v string)`

SetBaseSnapshotId sets BaseSnapshotId field to given value.

### HasBaseSnapshotId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) HasBaseSnapshotId() bool`

HasBaseSnapshotId returns a boolean if a field has been set.

### GetDatacenterId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetHostOrClusterId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetHostOrClusterId() string`

GetHostOrClusterId returns the HostOrClusterId field if non-nil, zero value otherwise.

### GetHostOrClusterIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetHostOrClusterIdOk() (*string, bool)`

GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrClusterId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetHostOrClusterId(v string)`

SetHostOrClusterId sets HostOrClusterId field to given value.


### GetImStreamId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.

### HasImStreamId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) HasImStreamId() bool`

HasImStreamId returns a boolean if a field has been set.

### GetImTagId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetImTagId() string`

GetImTagId returns the ImTagId field if non-nil, zero value otherwise.

### GetImTagIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetImTagIdOk() (*string, bool)`

GetImTagIdOk returns a tuple with the ImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImTagId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetImTagId(v string)`

SetImTagId sets ImTagId field to given value.

### HasImTagId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) HasImTagId() bool`

HasImTagId returns a boolean if a field has been set.

### GetParentVmId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetParentVmId() string`

GetParentVmId returns the ParentVmId field if non-nil, zero value otherwise.

### GetParentVmIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetParentVmIdOk() (*string, bool)`

GetParentVmIdOk returns a tuple with the ParentVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetParentVmId(v string)`

SetParentVmId sets ParentVmId field to given value.

### HasParentVmId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) HasParentVmId() bool`

HasParentVmId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.


### GetVmFolderId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetVmFolderId() string`

GetVmFolderId returns the VmFolderId field if non-nil, zero value otherwise.

### GetVmFolderIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetVmFolderIdOk() (*string, bool)`

GetVmFolderIdOk returns a tuple with the VmFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolderId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetVmFolderId(v string)`

SetVmFolderId sets VmFolderId field to given value.


### GetVmTemplateId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *DesktopPoolProvisioningSettingsCreateSpec) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.

### HasVmTemplateId

`func (o *DesktopPoolProvisioningSettingsCreateSpec) HasVmTemplateId() bool`

HasVmTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


