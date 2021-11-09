# ApplicationSupportedFileTypesData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableAutoUpdateFileTypes** | **bool** | Whether or not the file types supported by this application should be allowed to automatically update to reflect changes reported by the agent. Typically this should be set to false if the application has manually configured supported file types. Default is true. | 
**EnableAutoUpdateOtherFileTypes** | **bool** | Whether or not the other file types supported by this application should be allowed to automatically update to reflect changes reported by the agent. Typically this should be set to false if the application has manually configured supported file types. | 
**FileTypes** | Pointer to [**[]ApplicationFileTypeData**](ApplicationFileTypeData.md) | Set of file types reported by the application as supported (if this application is discovered) or as specified by the administrator (if this application is manually configured). If unset, this application does not present any file type support. | [optional] 
**OtherFileTypes** | Pointer to [**[]ApplicationOtherFileTypeData**](ApplicationOtherFileTypeData.md) | This represents the different file types reported by Application that can be passed from agent to client via broker or as specified by the administrator (if this application is manually configured). If unset, this application does not present any other file type support. | [optional] 

## Methods

### NewApplicationSupportedFileTypesData

`func NewApplicationSupportedFileTypesData(enableAutoUpdateFileTypes bool, enableAutoUpdateOtherFileTypes bool, ) *ApplicationSupportedFileTypesData`

NewApplicationSupportedFileTypesData instantiates a new ApplicationSupportedFileTypesData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSupportedFileTypesDataWithDefaults

`func NewApplicationSupportedFileTypesDataWithDefaults() *ApplicationSupportedFileTypesData`

NewApplicationSupportedFileTypesDataWithDefaults instantiates a new ApplicationSupportedFileTypesData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableAutoUpdateFileTypes

`func (o *ApplicationSupportedFileTypesData) GetEnableAutoUpdateFileTypes() bool`

GetEnableAutoUpdateFileTypes returns the EnableAutoUpdateFileTypes field if non-nil, zero value otherwise.

### GetEnableAutoUpdateFileTypesOk

`func (o *ApplicationSupportedFileTypesData) GetEnableAutoUpdateFileTypesOk() (*bool, bool)`

GetEnableAutoUpdateFileTypesOk returns a tuple with the EnableAutoUpdateFileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoUpdateFileTypes

`func (o *ApplicationSupportedFileTypesData) SetEnableAutoUpdateFileTypes(v bool)`

SetEnableAutoUpdateFileTypes sets EnableAutoUpdateFileTypes field to given value.


### GetEnableAutoUpdateOtherFileTypes

`func (o *ApplicationSupportedFileTypesData) GetEnableAutoUpdateOtherFileTypes() bool`

GetEnableAutoUpdateOtherFileTypes returns the EnableAutoUpdateOtherFileTypes field if non-nil, zero value otherwise.

### GetEnableAutoUpdateOtherFileTypesOk

`func (o *ApplicationSupportedFileTypesData) GetEnableAutoUpdateOtherFileTypesOk() (*bool, bool)`

GetEnableAutoUpdateOtherFileTypesOk returns a tuple with the EnableAutoUpdateOtherFileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableAutoUpdateOtherFileTypes

`func (o *ApplicationSupportedFileTypesData) SetEnableAutoUpdateOtherFileTypes(v bool)`

SetEnableAutoUpdateOtherFileTypes sets EnableAutoUpdateOtherFileTypes field to given value.


### GetFileTypes

`func (o *ApplicationSupportedFileTypesData) GetFileTypes() []ApplicationFileTypeData`

GetFileTypes returns the FileTypes field if non-nil, zero value otherwise.

### GetFileTypesOk

`func (o *ApplicationSupportedFileTypesData) GetFileTypesOk() (*[]ApplicationFileTypeData, bool)`

GetFileTypesOk returns a tuple with the FileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTypes

`func (o *ApplicationSupportedFileTypesData) SetFileTypes(v []ApplicationFileTypeData)`

SetFileTypes sets FileTypes field to given value.

### HasFileTypes

`func (o *ApplicationSupportedFileTypesData) HasFileTypes() bool`

HasFileTypes returns a boolean if a field has been set.

### GetOtherFileTypes

`func (o *ApplicationSupportedFileTypesData) GetOtherFileTypes() []ApplicationOtherFileTypeData`

GetOtherFileTypes returns the OtherFileTypes field if non-nil, zero value otherwise.

### GetOtherFileTypesOk

`func (o *ApplicationSupportedFileTypesData) GetOtherFileTypesOk() (*[]ApplicationOtherFileTypeData, bool)`

GetOtherFileTypesOk returns a tuple with the OtherFileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFileTypes

`func (o *ApplicationSupportedFileTypesData) SetOtherFileTypes(v []ApplicationOtherFileTypeData)`

SetOtherFileTypes sets OtherFileTypes field to given value.

### HasOtherFileTypes

`func (o *ApplicationSupportedFileTypesData) HasOtherFileTypes() bool`

HasOtherFileTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


