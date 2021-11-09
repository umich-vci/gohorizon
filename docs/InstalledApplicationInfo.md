# InstalledApplicationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutablePath** | Pointer to **string** | Path to application executable. | [optional] 
**FileTypes** | Pointer to [**[]ApplicationFileTypeData**](ApplicationFileTypeData.md) | Set of file types reported by the application as supported. If unset, this application does not present any file type support. | [optional] 
**Name** | Pointer to **string** | Application name information, as sent by RDSServer/machine during application discovery. | [optional] 
**OtherFileTypes** | Pointer to [**[]ApplicationOtherFileTypeData**](ApplicationOtherFileTypeData.md) | This represents the different file types reported by Application that can be passed from horizon agent to horizon client via connection server. If unset, this application does not present any other file type support. | [optional] 
**Publisher** | Pointer to **string** | Application publisher | [optional] 
**Version** | Pointer to **string** | Application version. | [optional] 

## Methods

### NewInstalledApplicationInfo

`func NewInstalledApplicationInfo() *InstalledApplicationInfo`

NewInstalledApplicationInfo instantiates a new InstalledApplicationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstalledApplicationInfoWithDefaults

`func NewInstalledApplicationInfoWithDefaults() *InstalledApplicationInfo`

NewInstalledApplicationInfoWithDefaults instantiates a new InstalledApplicationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutablePath

`func (o *InstalledApplicationInfo) GetExecutablePath() string`

GetExecutablePath returns the ExecutablePath field if non-nil, zero value otherwise.

### GetExecutablePathOk

`func (o *InstalledApplicationInfo) GetExecutablePathOk() (*string, bool)`

GetExecutablePathOk returns a tuple with the ExecutablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutablePath

`func (o *InstalledApplicationInfo) SetExecutablePath(v string)`

SetExecutablePath sets ExecutablePath field to given value.

### HasExecutablePath

`func (o *InstalledApplicationInfo) HasExecutablePath() bool`

HasExecutablePath returns a boolean if a field has been set.

### GetFileTypes

`func (o *InstalledApplicationInfo) GetFileTypes() []ApplicationFileTypeData`

GetFileTypes returns the FileTypes field if non-nil, zero value otherwise.

### GetFileTypesOk

`func (o *InstalledApplicationInfo) GetFileTypesOk() (*[]ApplicationFileTypeData, bool)`

GetFileTypesOk returns a tuple with the FileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTypes

`func (o *InstalledApplicationInfo) SetFileTypes(v []ApplicationFileTypeData)`

SetFileTypes sets FileTypes field to given value.

### HasFileTypes

`func (o *InstalledApplicationInfo) HasFileTypes() bool`

HasFileTypes returns a boolean if a field has been set.

### GetName

`func (o *InstalledApplicationInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstalledApplicationInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstalledApplicationInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstalledApplicationInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOtherFileTypes

`func (o *InstalledApplicationInfo) GetOtherFileTypes() []ApplicationOtherFileTypeData`

GetOtherFileTypes returns the OtherFileTypes field if non-nil, zero value otherwise.

### GetOtherFileTypesOk

`func (o *InstalledApplicationInfo) GetOtherFileTypesOk() (*[]ApplicationOtherFileTypeData, bool)`

GetOtherFileTypesOk returns a tuple with the OtherFileTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFileTypes

`func (o *InstalledApplicationInfo) SetOtherFileTypes(v []ApplicationOtherFileTypeData)`

SetOtherFileTypes sets OtherFileTypes field to given value.

### HasOtherFileTypes

`func (o *InstalledApplicationInfo) HasOtherFileTypes() bool`

HasOtherFileTypes returns a boolean if a field has been set.

### GetPublisher

`func (o *InstalledApplicationInfo) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *InstalledApplicationInfo) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *InstalledApplicationInfo) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *InstalledApplicationInfo) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetVersion

`func (o *InstalledApplicationInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstalledApplicationInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstalledApplicationInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstalledApplicationInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


