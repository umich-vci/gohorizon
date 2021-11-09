# ImageManagementVersionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **map[string]string** | Additional details about image management version. | [optional] 
**Description** | Pointer to **string** | Image management version description. | [optional] 
**Id** | **string** | Unique ID representing image management version. | 
**ImStreamId** | Pointer to **string** | Image management stream ID | [optional] 
**Name** | **string** | Image management version name. | 
**Status** | **string** | Image management version status. * AVAILABLE: Image management version is available for desktop pools/farms to be created. * DEPLOYING_VM: Image management version is deploying VM on the selected pod. * DEPLOYMENT_DONE: Image management version status when VM deployment is done for the selected pod. * DELETED: Image management version has been deleted. * DISABLED: Image management version has been disabled and no further pool operation can be done using the same. * FAILED: Image management version creation has failed. * PARTIALLY_AVAILABLE: Some of the image management asset creation in some of the virtual centers have failed. * PUBLISHING: Image management version is being published and specialized internally like installing agents etc. * REPLICATING: Copying the specialized images across all virtual centers. | 

## Methods

### NewImageManagementVersionInfo

`func NewImageManagementVersionInfo(id string, name string, status string, ) *ImageManagementVersionInfo`

NewImageManagementVersionInfo instantiates a new ImageManagementVersionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageManagementVersionInfoWithDefaults

`func NewImageManagementVersionInfoWithDefaults() *ImageManagementVersionInfo`

NewImageManagementVersionInfoWithDefaults instantiates a new ImageManagementVersionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *ImageManagementVersionInfo) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ImageManagementVersionInfo) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ImageManagementVersionInfo) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ImageManagementVersionInfo) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetDescription

`func (o *ImageManagementVersionInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ImageManagementVersionInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ImageManagementVersionInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ImageManagementVersionInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ImageManagementVersionInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ImageManagementVersionInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ImageManagementVersionInfo) SetId(v string)`

SetId sets Id field to given value.


### GetImStreamId

`func (o *ImageManagementVersionInfo) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *ImageManagementVersionInfo) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *ImageManagementVersionInfo) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.

### HasImStreamId

`func (o *ImageManagementVersionInfo) HasImStreamId() bool`

HasImStreamId returns a boolean if a field has been set.

### GetName

`func (o *ImageManagementVersionInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageManagementVersionInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageManagementVersionInfo) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ImageManagementVersionInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ImageManagementVersionInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ImageManagementVersionInfo) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


