# ImageManagementTagUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **map[string]string** | Additional details about image management tag. | [optional] 
**ImVersionId** | **string** | Image management version ID to which this tag belongs. | 
**Name** | **string** | Image management tag name. | 

## Methods

### NewImageManagementTagUpdateSpec

`func NewImageManagementTagUpdateSpec(imVersionId string, name string, ) *ImageManagementTagUpdateSpec`

NewImageManagementTagUpdateSpec instantiates a new ImageManagementTagUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageManagementTagUpdateSpecWithDefaults

`func NewImageManagementTagUpdateSpecWithDefaults() *ImageManagementTagUpdateSpec`

NewImageManagementTagUpdateSpecWithDefaults instantiates a new ImageManagementTagUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *ImageManagementTagUpdateSpec) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ImageManagementTagUpdateSpec) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ImageManagementTagUpdateSpec) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ImageManagementTagUpdateSpec) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetImVersionId

`func (o *ImageManagementTagUpdateSpec) GetImVersionId() string`

GetImVersionId returns the ImVersionId field if non-nil, zero value otherwise.

### GetImVersionIdOk

`func (o *ImageManagementTagUpdateSpec) GetImVersionIdOk() (*string, bool)`

GetImVersionIdOk returns a tuple with the ImVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImVersionId

`func (o *ImageManagementTagUpdateSpec) SetImVersionId(v string)`

SetImVersionId sets ImVersionId field to given value.


### GetName

`func (o *ImageManagementTagUpdateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageManagementTagUpdateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageManagementTagUpdateSpec) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


