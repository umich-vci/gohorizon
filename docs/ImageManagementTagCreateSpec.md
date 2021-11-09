# ImageManagementTagCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalDetails** | Pointer to **map[string]string** | Additional details about image management tag. | [optional] 
**ImStreamId** | **string** | Image management stream ID to which this tag belongs. | 
**ImVersionId** | **string** | Image management version ID to which this tag belongs. | 
**Name** | **string** | Image management tag name. It is unique among all the tags of a stream. | 

## Methods

### NewImageManagementTagCreateSpec

`func NewImageManagementTagCreateSpec(imStreamId string, imVersionId string, name string, ) *ImageManagementTagCreateSpec`

NewImageManagementTagCreateSpec instantiates a new ImageManagementTagCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImageManagementTagCreateSpecWithDefaults

`func NewImageManagementTagCreateSpecWithDefaults() *ImageManagementTagCreateSpec`

NewImageManagementTagCreateSpecWithDefaults instantiates a new ImageManagementTagCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalDetails

`func (o *ImageManagementTagCreateSpec) GetAdditionalDetails() map[string]string`

GetAdditionalDetails returns the AdditionalDetails field if non-nil, zero value otherwise.

### GetAdditionalDetailsOk

`func (o *ImageManagementTagCreateSpec) GetAdditionalDetailsOk() (*map[string]string, bool)`

GetAdditionalDetailsOk returns a tuple with the AdditionalDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalDetails

`func (o *ImageManagementTagCreateSpec) SetAdditionalDetails(v map[string]string)`

SetAdditionalDetails sets AdditionalDetails field to given value.

### HasAdditionalDetails

`func (o *ImageManagementTagCreateSpec) HasAdditionalDetails() bool`

HasAdditionalDetails returns a boolean if a field has been set.

### GetImStreamId

`func (o *ImageManagementTagCreateSpec) GetImStreamId() string`

GetImStreamId returns the ImStreamId field if non-nil, zero value otherwise.

### GetImStreamIdOk

`func (o *ImageManagementTagCreateSpec) GetImStreamIdOk() (*string, bool)`

GetImStreamIdOk returns a tuple with the ImStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImStreamId

`func (o *ImageManagementTagCreateSpec) SetImStreamId(v string)`

SetImStreamId sets ImStreamId field to given value.


### GetImVersionId

`func (o *ImageManagementTagCreateSpec) GetImVersionId() string`

GetImVersionId returns the ImVersionId field if non-nil, zero value otherwise.

### GetImVersionIdOk

`func (o *ImageManagementTagCreateSpec) GetImVersionIdOk() (*string, bool)`

GetImVersionIdOk returns a tuple with the ImVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImVersionId

`func (o *ImageManagementTagCreateSpec) SetImVersionId(v string)`

SetImVersionId sets ImVersionId field to given value.


### GetName

`func (o *ImageManagementTagCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ImageManagementTagCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ImageManagementTagCreateSpec) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


