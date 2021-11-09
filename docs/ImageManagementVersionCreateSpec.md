# ImageManagementVersionCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **map[string]string** | Additional details about image management version. | [optional] 
**Description** | Pointer to **string** | Image management version description. | [optional] 
**ImStreamId** | **string** | Image management stream ID | 
**Name** | **string** | Image management version name. | 
**Status** | **string** | Image management version status. * AVAILABLE: Image management version is available for desktop pools/farms to be created. * DEPLOYING_VM: Image management version is deploying VM on the selected pod. * DEPLOYMENT_DONE: Image management version status when VM deployment is done for the selected pod. * DELETED: Image management version has been deleted. * DISABLED: Image management version has been disabled and no further pool operation can be done using the same. * FAILED: Image management version creation has failed. * PARTIALLY_AVAILABLE: Some of the image management asset creation in some of the virtual centers have failed. * PUBLISHING: Image management version is being published and specialized internally like installing agents etc. * REPLICATING: Copying the specialized images across all virtual centers. | 

## Methods

### NewImageManagementVersionCreateSpec

`func NewImageManagementVersionCreateSpec(imStreamId string, name string, status string, ) *ImageManagementVersionCreateSpec`

NewImageManagementVersionCreateSpec instantiates a new ImageManagementVersionCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageManagementVersionCreateSpecWithDefaults

`func NewImageManagementVersionCreateSpecWithDefaults() *ImageManagementVersionCreateSpec`

NewImageManagementVersionCreateSpecWithDefaults instantiates a new ImageManagementVersionCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *ImageManagementVersionCreateSpec) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ImageManagementVersionCreateSpec) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ImageManagementVersionCreateSpec) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ImageManagementVersionCreateSpec) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetDescription

`func (o *ImageManagementVersionCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageManagementVersionCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageManagementVersionCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageManagementVersionCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetImStreamId

`func (o *ImageManagementVersionCreateSpec) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *ImageManagementVersionCreateSpec) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *ImageManagementVersionCreateSpec) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.


### GetName

`func (o *ImageManagementVersionCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageManagementVersionCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageManagementVersionCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ImageManagementVersionCreateSpec) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageManagementVersionCreateSpec) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageManagementVersionCreateSpec) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


