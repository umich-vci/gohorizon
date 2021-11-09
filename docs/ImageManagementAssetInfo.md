# ImageManagementAssetInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **map[string]string** | Additional details about image management asset. | [optional] 
**BaseSnapshotId** | **string** | Virtual machine snapshot. Must be set if vm_template_id is unset. | 
**BaseVmId** | **string** | Virtual machine ID. Must be set if vm_template_id is unset. | 
**CloneType** | **string** | Image management asset clone type. * FULL_CLONE: Image management asset to be used in full clone automated desktop pool. * INSTANT_CLONE: Image management asset to be used in instant clone desktop pool/farm. | 
**DatacenterId** | **string** | Datacenter where this asset is created. | 
**Id** | **string** | Unique ID representing image management asset. | 
**ImStreamId** | **string** | Image management stream to which this asset belongs to. | 
**ImVersionId** | **string** | Image management version to which this asset belongs to. | 
**ImageType** | **string** | Image management asset image type. * RDSH_APPS: Image management asset to be used for farm creation which is be used in application. * RDSH_DESKTOP: Image management asset is for farm creation to be created. * VDI_DESKTOP: Image management asset is available for desktops/farms to be created. | 
**Status** | **string** | Image management asset status. * AVAILABLE: Image management asset is available for desktop pools/farms to be created. * DEPLOYING_VM: Image management asset is deploying VM on the virtual center. * DEPLOYMENT_DONE: Image management asset VM deployed on the virtual center. * DELETED: Image management asset has been deleted. * DISABLED: Image management asset has been disabled and no further pool/farm operation can be done using the same. * FAILED: Image management asset creation has failed. * REPLICATING: Copying the specialized images across all virtual centers. * RETRY_PENDING: When image management asset creation has failed, retry action is pending for asset to be created. * SPECIALIZING_VM: Image management asset is being published and specialized internally like installing agents etc. | 
**VcenterId** | **string** | Virtual Center where this asset is created. | 
**VmTemplateId** | **string** | Virtual machine template ID. | 

## Methods

### NewImageManagementAssetInfo

`func NewImageManagementAssetInfo(baseSnapshotId string, baseVmId string, cloneType string, datacenterId string, id string, imStreamId string, imVersionId string, imageType string, status string, vcenterId string, vmTemplateId string, ) *ImageManagementAssetInfo`

NewImageManagementAssetInfo instantiates a new ImageManagementAssetInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageManagementAssetInfoWithDefaults

`func NewImageManagementAssetInfoWithDefaults() *ImageManagementAssetInfo`

NewImageManagementAssetInfoWithDefaults instantiates a new ImageManagementAssetInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *ImageManagementAssetInfo) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ImageManagementAssetInfo) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ImageManagementAssetInfo) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ImageManagementAssetInfo) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetBaseSnapshotId

`func (o *ImageManagementAssetInfo) GetBaseSnapshotId() string`

GetBaseSnapshotId returns the BaseSnapshotId field if non-nil, zero value otherwise.

### GetBaseSnapshotIdOk

`func (o *ImageManagementAssetInfo) GetBaseSnapshotIdOk() (*string, bool)`

GetBaseSnapshotIdOk returns a tuple with the BaseSnapshotId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseSnapshotId

`func (o *ImageManagementAssetInfo) SetBaseSnapshotId(v string)`

SetBaseSnapshotId sets BaseSnapshotId field to given value.


### GetBaseVmId

`func (o *ImageManagementAssetInfo) GetBaseVmId() string`

GetBaseVmId returns the BaseVmId field if non-nil, zero value otherwise.

### GetBaseVmIdOk

`func (o *ImageManagementAssetInfo) GetBaseVmIdOk() (*string, bool)`

GetBaseVmIdOk returns a tuple with the BaseVmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseVmId

`func (o *ImageManagementAssetInfo) SetBaseVmId(v string)`

SetBaseVmId sets BaseVmId field to given value.


### GetCloneType

`func (o *ImageManagementAssetInfo) GetCloneType() string`

GetCloneType returns the CloneType field if non-nil, zero value otherwise.

### GetCloneTypeOk

`func (o *ImageManagementAssetInfo) GetCloneTypeOk() (*string, bool)`

GetCloneTypeOk returns a tuple with the CloneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneType

`func (o *ImageManagementAssetInfo) SetCloneType(v string)`

SetCloneType sets CloneType field to given value.


### GetDatacenterId

`func (o *ImageManagementAssetInfo) GetDatacenterId() string`

GetDatacenterId returns the DatacenterId field if non-nil, zero value otherwise.

### GetDatacenterIdOk

`func (o *ImageManagementAssetInfo) GetDatacenterIdOk() (*string, bool)`

GetDatacenterIdOk returns a tuple with the DatacenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenterId

`func (o *ImageManagementAssetInfo) SetDatacenterId(v string)`

SetDatacenterId sets DatacenterId field to given value.


### GetId

`func (o *ImageManagementAssetInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageManagementAssetInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageManagementAssetInfo) SetId(v string)`

SetId sets Id field to given value.


### GetImStreamId

`func (o *ImageManagementAssetInfo) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *ImageManagementAssetInfo) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *ImageManagementAssetInfo) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.


### GetImVersionId

`func (o *ImageManagementAssetInfo) GetImVersionId() string`

GetImVersionId returns the ImVersionId field if non-nil, zero value otherwise.

### GetImVersionIdOk

`func (o *ImageManagementAssetInfo) GetImVersionIdOk() (*string, bool)`

GetImVersionIdOk returns a tuple with the ImVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImVersionId

`func (o *ImageManagementAssetInfo) SetImVersionId(v string)`

SetImVersionId sets ImVersionId field to given value.


### GetImageType

`func (o *ImageManagementAssetInfo) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ImageManagementAssetInfo) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ImageManagementAssetInfo) SetImageType(v string)`

SetImageType sets ImageType field to given value.


### GetStatus

`func (o *ImageManagementAssetInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageManagementAssetInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageManagementAssetInfo) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVcenterId

`func (o *ImageManagementAssetInfo) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *ImageManagementAssetInfo) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *ImageManagementAssetInfo) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.


### GetVmTemplateId

`func (o *ImageManagementAssetInfo) GetVmTemplateId() string`

GetVmTemplateId returns the VmTemplateId field if non-nil, zero value otherwise.

### GetVmTemplateIdOk

`func (o *ImageManagementAssetInfo) GetVmTemplateIdOk() (*string, bool)`

GetVmTemplateIdOk returns a tuple with the VmTemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmTemplateId

`func (o *ImageManagementAssetInfo) SetVmTemplateId(v string)`

SetVmTemplateId sets VmTemplateId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


