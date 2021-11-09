# FarmProvisioningSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseSnapshotId** | Pointer to **string** | Current image snapshot for instant clone farm. | [optional] 
**DatacenterId** | Pointer to **string** | Datacenter within which the farm is configured | [optional] 
**HostOrClusterId** | Pointer to **string** | Host or cluster where the machines are deployed in. | [optional] 
**ImStreamId** | Pointer to **string** | Image management stream used in the farm.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ImTagId** | Pointer to **string** | Image management tag used in the farm.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 
**ParentVmId** | Pointer to **string** | Current VM image for instant clone farm. | [optional] 
**ResourcePoolId** | Pointer to **string** | Resource pool to deploy the RDS Servers. | [optional] 
**VmFolderId** | Pointer to **string** | VM folder to deploy the RDS Servers to. | [optional] 

## Methods

### NewFarmProvisioningSettingsInfo

`func NewFarmProvisioningSettingsInfo() *FarmProvisioningSettingsInfo`

NewFarmProvisioningSettingsInfo instantiates a new FarmProvisioningSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmProvisioningSettingsInfoWithDefaults

`func NewFarmProvisioningSettingsInfoWithDefaults() *FarmProvisioningSettingsInfo`

NewFarmProvisioningSettingsInfoWithDefaults instantiates a new FarmProvisioningSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseSnapshotId

`func (o *FarmProvisioningSettingsInfo) GetBaseSnapshotId() string`

GetBaseSnapshotId returns the BaseSnapshotId field if non-nil, zero value otherwise.

### GetBaseSnapshotIdOk

`func (o *FarmProvisioningSettingsInfo) GetBaseSnapshotIdOk() (*string, bool)`

GetBaseSnapshotIdOk returns a tuple with the BaseSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotId

`func (o *FarmProvisioningSettingsInfo) SetBaseSnapshotId(v string)`

SetBaseSnapshotId sets BaseSnapshotId field to given value.

### HasBaseSnapshotId

`func (o *FarmProvisioningSettingsInfo) HasBaseSnapshotId() bool`

HasBaseSnapshotId returns a boolean if a field has been set.

### GetDatacenterId

`func (o *FarmProvisioningSettingsInfo) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *FarmProvisioningSettingsInfo) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *FarmProvisioningSettingsInfo) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.

### HasDatacenterId

`func (o *FarmProvisioningSettingsInfo) HasDatacenterId() bool`

HasDatacenterId returns a boolean if a field has been set.

### GetHostOrClusterId

`func (o *FarmProvisioningSettingsInfo) GetHostOrClusterId() string`

GetHostOrClusterId returns the HostOrClusterId field if non-nil, zero value otherwise.

### GetHostOrClusterIdOk

`func (o *FarmProvisioningSettingsInfo) GetHostOrClusterIdOk() (*string, bool)`

GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrClusterId

`func (o *FarmProvisioningSettingsInfo) SetHostOrClusterId(v string)`

SetHostOrClusterId sets HostOrClusterId field to given value.

### HasHostOrClusterId

`func (o *FarmProvisioningSettingsInfo) HasHostOrClusterId() bool`

HasHostOrClusterId returns a boolean if a field has been set.

### GetImStreamId

`func (o *FarmProvisioningSettingsInfo) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *FarmProvisioningSettingsInfo) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *FarmProvisioningSettingsInfo) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.

### HasImStreamId

`func (o *FarmProvisioningSettingsInfo) HasImStreamId() bool`

HasImStreamId returns a boolean if a field has been set.

### GetImTagId

`func (o *FarmProvisioningSettingsInfo) GetImTagId() string`

GetImTagId returns the ImTagId field if non-nil, zero value otherwise.

### GetImTagIdOk

`func (o *FarmProvisioningSettingsInfo) GetImTagIdOk() (*string, bool)`

GetImTagIdOk returns a tuple with the ImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImTagId

`func (o *FarmProvisioningSettingsInfo) SetImTagId(v string)`

SetImTagId sets ImTagId field to given value.

### HasImTagId

`func (o *FarmProvisioningSettingsInfo) HasImTagId() bool`

HasImTagId returns a boolean if a field has been set.

### GetParentVmId

`func (o *FarmProvisioningSettingsInfo) GetParentVmId() string`

GetParentVmId returns the ParentVmId field if non-nil, zero value otherwise.

### GetParentVmIdOk

`func (o *FarmProvisioningSettingsInfo) GetParentVmIdOk() (*string, bool)`

GetParentVmIdOk returns a tuple with the ParentVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmId

`func (o *FarmProvisioningSettingsInfo) SetParentVmId(v string)`

SetParentVmId sets ParentVmId field to given value.

### HasParentVmId

`func (o *FarmProvisioningSettingsInfo) HasParentVmId() bool`

HasParentVmId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *FarmProvisioningSettingsInfo) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *FarmProvisioningSettingsInfo) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *FarmProvisioningSettingsInfo) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.

### HasResourcePoolId

`func (o *FarmProvisioningSettingsInfo) HasResourcePoolId() bool`

HasResourcePoolId returns a boolean if a field has been set.

### GetVmFolderId

`func (o *FarmProvisioningSettingsInfo) GetVmFolderId() string`

GetVmFolderId returns the VmFolderId field if non-nil, zero value otherwise.

### GetVmFolderIdOk

`func (o *FarmProvisioningSettingsInfo) GetVmFolderIdOk() (*string, bool)`

GetVmFolderIdOk returns a tuple with the VmFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolderId

`func (o *FarmProvisioningSettingsInfo) SetVmFolderId(v string)`

SetVmFolderId sets VmFolderId field to given value.

### HasVmFolderId

`func (o *FarmProvisioningSettingsInfo) HasVmFolderId() bool`

HasVmFolderId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


