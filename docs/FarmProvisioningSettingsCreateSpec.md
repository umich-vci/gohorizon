# FarmProvisioningSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseSnapshotId** | Pointer to **string** | Current image snapshot for instant clone farm. | [optional] 
**DatacenterId** | **string** | Datacenter within which the farm is configured | 
**HostOrClusterId** | **string** | Host or cluster where the machines are deployed in. For Instant clone farms it can be only be a cluster id. | 
**ImStreamId** | Pointer to **string** | Image management stream used in the farm. This is required if parent_vm_id and base_snapshot_id are not set. | [optional] 
**ImTagId** | Pointer to **string** | Image management tag used in the farm. This is required if im_stream_id is set. | [optional] 
**ParentVmId** | Pointer to **string** | Current VM image for instant clone farm. | [optional] 
**ResourcePoolId** | **string** | Resource pool to deploy the RDS Servers. | 
**VmFolderId** | **string** | VM folder to deploy the RDS Servers to. | 

## Methods

### NewFarmProvisioningSettingsCreateSpec

`func NewFarmProvisioningSettingsCreateSpec(datacenterId string, hostOrClusterId string, resourcePoolId string, vmFolderId string, ) *FarmProvisioningSettingsCreateSpec`

NewFarmProvisioningSettingsCreateSpec instantiates a new FarmProvisioningSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmProvisioningSettingsCreateSpecWithDefaults

`func NewFarmProvisioningSettingsCreateSpecWithDefaults() *FarmProvisioningSettingsCreateSpec`

NewFarmProvisioningSettingsCreateSpecWithDefaults instantiates a new FarmProvisioningSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseSnapshotId

`func (o *FarmProvisioningSettingsCreateSpec) GetBaseSnapshotId() string`

GetBaseSnapshotId returns the BaseSnapshotId field if non-nil, zero value otherwise.

### GetBaseSnapshotIdOk

`func (o *FarmProvisioningSettingsCreateSpec) GetBaseSnapshotIdOk() (*string, bool)`

GetBaseSnapshotIdOk returns a tuple with the BaseSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotId

`func (o *FarmProvisioningSettingsCreateSpec) SetBaseSnapshotId(v string)`

SetBaseSnapshotId sets BaseSnapshotId field to given value.

### HasBaseSnapshotId

`func (o *FarmProvisioningSettingsCreateSpec) HasBaseSnapshotId() bool`

HasBaseSnapshotId returns a boolean if a field has been set.

### GetDatacenterId

`func (o *FarmProvisioningSettingsCreateSpec) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *FarmProvisioningSettingsCreateSpec) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *FarmProvisioningSettingsCreateSpec) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetHostOrClusterId

`func (o *FarmProvisioningSettingsCreateSpec) GetHostOrClusterId() string`

GetHostOrClusterId returns the HostOrClusterId field if non-nil, zero value otherwise.

### GetHostOrClusterIdOk

`func (o *FarmProvisioningSettingsCreateSpec) GetHostOrClusterIdOk() (*string, bool)`

GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrClusterId

`func (o *FarmProvisioningSettingsCreateSpec) SetHostOrClusterId(v string)`

SetHostOrClusterId sets HostOrClusterId field to given value.


### GetImStreamId

`func (o *FarmProvisioningSettingsCreateSpec) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *FarmProvisioningSettingsCreateSpec) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *FarmProvisioningSettingsCreateSpec) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.

### HasImStreamId

`func (o *FarmProvisioningSettingsCreateSpec) HasImStreamId() bool`

HasImStreamId returns a boolean if a field has been set.

### GetImTagId

`func (o *FarmProvisioningSettingsCreateSpec) GetImTagId() string`

GetImTagId returns the ImTagId field if non-nil, zero value otherwise.

### GetImTagIdOk

`func (o *FarmProvisioningSettingsCreateSpec) GetImTagIdOk() (*string, bool)`

GetImTagIdOk returns a tuple with the ImTagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImTagId

`func (o *FarmProvisioningSettingsCreateSpec) SetImTagId(v string)`

SetImTagId sets ImTagId field to given value.

### HasImTagId

`func (o *FarmProvisioningSettingsCreateSpec) HasImTagId() bool`

HasImTagId returns a boolean if a field has been set.

### GetParentVmId

`func (o *FarmProvisioningSettingsCreateSpec) GetParentVmId() string`

GetParentVmId returns the ParentVmId field if non-nil, zero value otherwise.

### GetParentVmIdOk

`func (o *FarmProvisioningSettingsCreateSpec) GetParentVmIdOk() (*string, bool)`

GetParentVmIdOk returns a tuple with the ParentVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVmId

`func (o *FarmProvisioningSettingsCreateSpec) SetParentVmId(v string)`

SetParentVmId sets ParentVmId field to given value.

### HasParentVmId

`func (o *FarmProvisioningSettingsCreateSpec) HasParentVmId() bool`

HasParentVmId returns a boolean if a field has been set.

### GetResourcePoolId

`func (o *FarmProvisioningSettingsCreateSpec) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *FarmProvisioningSettingsCreateSpec) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *FarmProvisioningSettingsCreateSpec) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.


### GetVmFolderId

`func (o *FarmProvisioningSettingsCreateSpec) GetVmFolderId() string`

GetVmFolderId returns the VmFolderId field if non-nil, zero value otherwise.

### GetVmFolderIdOk

`func (o *FarmProvisioningSettingsCreateSpec) GetVmFolderIdOk() (*string, bool)`

GetVmFolderIdOk returns a tuple with the VmFolderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmFolderId

`func (o *FarmProvisioningSettingsCreateSpec) SetVmFolderId(v string)`

SetVmFolderId sets VmFolderId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


