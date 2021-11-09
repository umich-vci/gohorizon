# ApplicationFileTypeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Friendly name for the file type. If unset, no friendly name will be displayed. | [optional] 
**Type** | **string** | File type supported by this application. This value is case insensitive. If multiple file types are specified using the same (case insensitive) name and description, all but one will be ignored. | 

## Methods

### NewApplicationFileTypeData

`func NewApplicationFileTypeData(type_ string, ) *ApplicationFileTypeData`

NewApplicationFileTypeData instantiates a new ApplicationFileTypeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationFileTypeDataWithDefaults

`func NewApplicationFileTypeDataWithDefaults() *ApplicationFileTypeData`

NewApplicationFileTypeDataWithDefaults instantiates a new ApplicationFileTypeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ApplicationFileTypeData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationFileTypeData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationFileTypeData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationFileTypeData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ApplicationFileTypeData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationFileTypeData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationFileTypeData) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


