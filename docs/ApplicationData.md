# ApplicationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExecutablePath** | Pointer to **string** | Path to application executable. | [optional] 
**Publisher** | Pointer to **string** | Application publisher. | [optional] 
**Version** | Pointer to **string** | Application version. | [optional] 

## Methods

### NewApplicationData

`func NewApplicationData() *ApplicationData`

NewApplicationData instantiates a new ApplicationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationDataWithDefaults

`func NewApplicationDataWithDefaults() *ApplicationData`

NewApplicationDataWithDefaults instantiates a new ApplicationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExecutablePath

`func (o *ApplicationData) GetExecutablePath() string`

GetExecutablePath returns the ExecutablePath field if non-nil, zero value otherwise.

### GetExecutablePathOk

`func (o *ApplicationData) GetExecutablePathOk() (*string, bool)`

GetExecutablePathOk returns a tuple with the ExecutablePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutablePath

`func (o *ApplicationData) SetExecutablePath(v string)`

SetExecutablePath sets ExecutablePath field to given value.

### HasExecutablePath

`func (o *ApplicationData) HasExecutablePath() bool`

HasExecutablePath returns a boolean if a field has been set.

### GetPublisher

`func (o *ApplicationData) GetPublisher() string`

GetPublisher returns the Publisher field if non-nil, zero value otherwise.

### GetPublisherOk

`func (o *ApplicationData) GetPublisherOk() (*string, bool)`

GetPublisherOk returns a tuple with the Publisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisher

`func (o *ApplicationData) SetPublisher(v string)`

SetPublisher sets Publisher field to given value.

### HasPublisher

`func (o *ApplicationData) HasPublisher() bool`

HasPublisher returns a boolean if a field has been set.

### GetVersion

`func (o *ApplicationData) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApplicationData) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApplicationData) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApplicationData) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


