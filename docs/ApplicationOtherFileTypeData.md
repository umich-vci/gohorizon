# ApplicationOtherFileTypeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Friendly name for the file type. If unset, no friendly name will be displayed. | [optional] 
**Name** | **string** | The name for other file type data. | 
**Type** | **string** | Other file type currently supported. * URL: URL scheme types supported by application. * OTHER: Other scheme types supported by application. | 

## Methods

### NewApplicationOtherFileTypeData

`func NewApplicationOtherFileTypeData(name string, type_ string, ) *ApplicationOtherFileTypeData`

NewApplicationOtherFileTypeData instantiates a new ApplicationOtherFileTypeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationOtherFileTypeDataWithDefaults

`func NewApplicationOtherFileTypeDataWithDefaults() *ApplicationOtherFileTypeData`

NewApplicationOtherFileTypeDataWithDefaults instantiates a new ApplicationOtherFileTypeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApplicationOtherFileTypeData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationOtherFileTypeData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationOtherFileTypeData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationOtherFileTypeData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *ApplicationOtherFileTypeData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationOtherFileTypeData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationOtherFileTypeData) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ApplicationOtherFileTypeData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationOtherFileTypeData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationOtherFileTypeData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


