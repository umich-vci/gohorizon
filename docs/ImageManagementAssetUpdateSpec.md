# ImageManagementAssetUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **map[string]string** | Additional details about image management asset. | [optional] 
**CloneType** | **string** | Image management asset clone type. * FULL_CLONE: Image management asset to be used in full clone automated desktop pool. * INSTANT_CLONE: Image management asset to be used in instant clone desktop pool/farm. | 
**ImageType** | **string** | Image management asset image type. * RDSH_APPS: Image management asset to be used for farm creation which is be used in application. * RDSH_DESKTOP: Image management asset is for farm creation to be created. * VDI_DESKTOP: Image management asset is available for desktops/farms to be created. | 
**Status** | **string** | Image management asset status. * AVAILABLE: Image management asset is available for desktop pools/farms to be created. * DEPLOYING_VM: Image management asset is deploying VM on the virtual center. * DEPLOYMENT_DONE: Image management asset VM deployed on the virtual center. * DELETED: Image management asset has been deleted. * DISABLED: Image management asset has been disabled and no further pool/farm operation can be done using the same. * FAILED: Image management asset creation has failed. * REPLICATING: Copying the specialized images across all virtual centers. * RETRY_PENDING: When image management asset creation has failed, retry action is pending for asset to be created. * SPECIALIZING_VM: Image management asset is being published and specialized internally like installing agents etc. | 

## Methods

### NewImageManagementAssetUpdateSpec

`func NewImageManagementAssetUpdateSpec(cloneType string, imageType string, status string, ) *ImageManagementAssetUpdateSpec`

NewImageManagementAssetUpdateSpec instantiates a new ImageManagementAssetUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageManagementAssetUpdateSpecWithDefaults

`func NewImageManagementAssetUpdateSpecWithDefaults() *ImageManagementAssetUpdateSpec`

NewImageManagementAssetUpdateSpecWithDefaults instantiates a new ImageManagementAssetUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *ImageManagementAssetUpdateSpec) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ImageManagementAssetUpdateSpec) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ImageManagementAssetUpdateSpec) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ImageManagementAssetUpdateSpec) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetCloneType

`func (o *ImageManagementAssetUpdateSpec) GetCloneType() string`

GetCloneType returns the CloneType field if non-nil, zero value otherwise.

### GetCloneTypeOk

`func (o *ImageManagementAssetUpdateSpec) GetCloneTypeOk() (*string, bool)`

GetCloneTypeOk returns a tuple with the CloneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneType

`func (o *ImageManagementAssetUpdateSpec) SetCloneType(v string)`

SetCloneType sets CloneType field to given value.


### GetImageType

`func (o *ImageManagementAssetUpdateSpec) GetImageType() string`

GetImageType returns the ImageType field if non-nil, zero value otherwise.

### GetImageTypeOk

`func (o *ImageManagementAssetUpdateSpec) GetImageTypeOk() (*string, bool)`

GetImageTypeOk returns a tuple with the ImageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageType

`func (o *ImageManagementAssetUpdateSpec) SetImageType(v string)`

SetImageType sets ImageType field to given value.


### GetStatus

`func (o *ImageManagementAssetUpdateSpec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageManagementAssetUpdateSpec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageManagementAssetUpdateSpec) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


