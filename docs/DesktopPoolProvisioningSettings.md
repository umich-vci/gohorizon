# DesktopPoolProvisioningSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddVirtualTpm** | Pointer to **bool** | Whether to add Virtual TPM device. | [optional] 
**BaseSnapshotId** | Pointer to **string** | Applicable To: Linked/instant clone automated desktop pools.&lt;br&gt;Base image snapshot for linked clone desktop pool and current image snapshot for instant clone desktop pool. | [optional] 
**DatacenterId** | Pointer to **string** | Datacenter within which the desktop pool is configured. | [optional] 
**HostOrClusterId** | Pointer to **string** | Host or cluster where the machines are deployed in. | [optional] 
**ImStreamId** | Pointer to **string** | Applicable To: Full/instant clone automated desktop pools.&lt;br&gt;Image management stream used in desktop pool when Image Management feature is enabled.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ImTagId** | Pointer to **string** | Applicable To: Full/instant clone automated desktop pools.&lt;br&gt;Image management tag associated with the selected image management stream which is used in desktop pool when Image Management feature is enabled.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**MinReadyVmsOnVcomposerMaintenance** | Pointer to **int32** | Applicable To: Linked clone automated desktop pools.&lt;br&gt;Minimum number of ready (provisioned) machines during View Composer maintenance operations. | [optional] 
**ParentVmId** | Pointer to **string** | Applicable To: Linked/instant clone automated desktop pools.&lt;br&gt;Base image VM for linked clone desktop pool and current image for instant clone desktop pool. | [optional] 
**ResourcePoolId** | Pointer to **string** | Resource pool to deploy the machines. | [optional] 
**VmFolderId** | Pointer to **string** | VM folder where the machines are deployed to. | [optional] 
**VmTemplateId** | Pointer to **string** | Applicable To: Full clone automated desktop pools.&lt;br&gt;Template from which full clone machines are deployed. | [optional] 

## Methods

### NewDesktopPoolProvisioningSettings

`func NewDesktopPoolProvisioningSettings() *DesktopPoolProvisioningSettings`

NewDesktopPoolProvisioningSettings instantiates a new DesktopPoolProvisioningSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolProvisioningSettingsWithDefaults

`func NewDesktopPoolProvisioningSettingsWithDefaults() *DesktopPoolProvisioningSettings`

NewDesktopPoolProvisioningSettingsWithDefaults instantiates a new DesktopPoolProvisioningSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddVirtualTpm

`func (o *DesktopPoolProvisioningSettings) GetAddVirtualTpm() bool`

GetAddVirtualTpm returns the AddVirtualTpm field if non-nil, zero value otherwise.

### GetAddVirtualTpmOk

`func (o *DesktopPoolProvisioningSettings) GetAddVirtualTpmOk() (*bool, bool)`

GetAddVirtualTpmOk returns a tuple with the AddVirtualTpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVirtualTpm

`func (o *DesktopPoolProvisioningSettings) SetAddVirtualTpm(v bool)`

SetAddVirtualTpm sets AddVirtualTpm field to given value.

### HasAddVirtualTpm

`func (o *DesktopPoolProvisioningSettings) HasAddVirtualTpm() bool`

HasAddVirtualTpm returns a boolean if a field has been set.

### GetBaseSnapshotId

`func (o *DesktopPoolProvisioningSettings) GetBaseSnapshotId() string`

GetBaseSnapshotId returns the BaseSnapshotId field if non-nil, zero value otherwise.

### GetBaseSnapshotIdOk

`func (o *DesktopPoolProvisioningSettings) GetBaseSnapshotIdOk() (*string, bool)`

GetBaseSnapshotIdOk returns a tuple with the BaseSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotId

`func (o *DesktopPoolProvisioningSettings) SetBaseSnapshotId(v string)`

SetBaseSnapshotId sets BaseSnapshotId field to given value.

### HasBaseSnapshotId

`func (o *DesktopPoolProvisioningSettings) HasBaseSnapshotId() bool`

HasBaseSnapshotId returns a boolean if a field has been set.

### GetDatacenterId

`func (o *DesktopPoolProvisioningSettings) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *DesktopPoolProvisioningSettings) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *DesktopPoolProvisioningSettings) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *DesktopPoolProvisioningSettings) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetHostOrClusterId

`func (o *DesktopPoolProvisioningSettings) GetHostOrClusterId() string`

GetHostOrClusterId returns the HostOrClusterId field if non-nil, zero value otherwise.

### GetHostOrClusterIdOk

`func (o *DesktopPoolProvisioningSettings) GetHostOrClusterIdOk() (*string, bool)`

GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrClusterId

`func (o *DesktopPoolProvisioningSettings) SetHostOrClusterId(v string)`

SetHostOrClusterId sets HostOrClusterId field to given value.

### HasHostOrClusterId

`func (o *DesktopPoolProvisioningSettings) HasHostOrClusterId() bool`

HasHostOrClusterId returns a boolean if a field has been set.

### GetImStreamId

`func (o *DesktopPoolProvisioningSettings) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *DesktopPoolProvisioningSettings) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *DesktopPoolProvisioningSettings) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.

### HasImStreamId

`func (o *DesktopPoolProvisioningSettings) HasImStreamId() bool`

HasImStreamId returns a boolean if a field has been set.

### GetImTagId

`func (o *DesktopPoolProvisioningSettings) GetImTagId() string`

GetImTagId returns the ImTagId field if non-nil, zero value otherwise.

### GetImTagIdOk

`func (o *DesktopPoolProvisioningSettings) GetImTagIdOk() (*string, bool)`

GetImTagIdOk returns a tuple with the ImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImTagId

`func (o *DesktopPoolProvisioningSettings) SetImTagId(v string)`

SetImTagId sets ImTagId field to given value.

### HasImTagId

`func (o *DesktopPoolProvisioningSettings) HasImTagId() bool`

HasImTagId returns a boolean if a field has been set.

### GetMinReadyVmsOnVcomposerMaintenance

`func (o *DesktopPoolProvisioningSettings) GetMinReadyVmsOnVcomposerMaintenance() int32`

GetMinReadyVmsOnVcomposerMaintenance returns the MinReadyVmsOnVcomposerMaintenance field if non-nil, zero value otherwise.

### GetMinReadyVmsOnVcomposerMaintenanceOk

`func (o *DesktopPoolProvisioningSettings) GetMinReadyVmsOnVcomposerMaintenanceOk() (*int32, bool)`

GetMinReadyVmsOnVcomposerMaintenanceOk returns a tuple with the MinReadyVmsOnVcomposerMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadyVmsOnVcomposerMaintenance

`func (o *DesktopPoolProvisioningSettings) SetMinReadyVmsOnVcomposerMaintenance(v int32)`

SetMinReadyVmsOnVcomposerMaintenance sets MinReadyVmsOnVcomposerMaintenance field to given value.

### HasMinReadyVmsOnVcomposerMaintenance

`func (o *DesktopPoolProvisioningSettings) HasMinReadyVmsOnVcomposerMaintenance() bool`

HasMinReadyVmsOnVcomposerMaintenance returns a boolean if a field has been set.

### GetParentVmId

`func (o *DesktopPoolProvisioningSettings) GetParentVmId() string`

GetParentVmId returns the ParentVmId field if non-nil, zero value otherwise.

### GetParentVmIdOk

`func (o *DesktopPoolProvisioningSettings) GetParentVmIdOk() (*string, bool)`

GetParentVmIdOk returns a tuple with the ParentVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmId

`func (o *DesktopPoolProvisioningSettings) SetParentVmId(v string)`

SetParentVmId sets ParentVmId field to given value.

### HasParentVmId

`func (o *DesktopPoolProvisioningSettings) HasParentVmId() bool`

HasParentVmId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *DesktopPoolProvisioningSettings) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *DesktopPoolProvisioningSettings) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *DesktopPoolProvisioningSettings) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *DesktopPoolProvisioningSettings) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetVmFolderId

`func (o *DesktopPoolProvisioningSettings) GetVmFolderId() string`

GetVmFolderId returns the VmFolderId field if non-nil, zero value otherwise.

### GetVmFolderIdOk

`func (o *DesktopPoolProvisioningSettings) GetVmFolderIdOk() (*string, bool)`

GetVmFolderIdOk returns a tuple with the VmFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolderId

`func (o *DesktopPoolProvisioningSettings) SetVmFolderId(v string)`

SetVmFolderId sets VmFolderId field to given value.

### HasVmFolderId

`func (o *DesktopPoolProvisioningSettings) HasVmFolderId() bool`

HasVmFolderId returns a boolean if a field has been set.

### GetVmTemplateId

`func (o *DesktopPoolProvisioningSettings) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *DesktopPoolProvisioningSettings) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *DesktopPoolProvisioningSettings) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.

### HasVmTemplateId

`func (o *DesktopPoolProvisioningSettings) HasVmTemplateId() bool`

HasVmTemplateId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


