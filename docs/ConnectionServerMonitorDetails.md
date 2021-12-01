# ConnectionServerMonitorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Build** | Pointer to **string** | Connection Server build number. | [optional] 
**Version** | Pointer to **string** | Connection Server version number. | [optional] 

## Methods

### NewConnectionServerMonitorDetails

`func NewConnectionServerMonitorDetails() *ConnectionServerMonitorDetails`

NewConnectionServerMonitorDetails instantiates a new ConnectionServerMonitorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServerMonitorDetailsWithDefaults

`func NewConnectionServerMonitorDetailsWithDefaults() *ConnectionServerMonitorDetails`

NewConnectionServerMonitorDetailsWithDefaults instantiates a new ConnectionServerMonitorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuild

`func (o *ConnectionServerMonitorDetails) GetBuild() string`

GetBuild returns the Build field if non-nil, zero value otherwise.

### GetBuildOk

`func (o *ConnectionServerMonitorDetails) GetBuildOk() (*string, bool)`

GetBuildOk returns a tuple with the Build field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuild

`func (o *ConnectionServerMonitorDetails) SetBuild(v string)`

SetBuild sets Build field to given value.

### HasBuild

`func (o *ConnectionServerMonitorDetails) HasBuild() bool`

HasBuild returns a boolean if a field has been set.

### GetVersion

`func (o *ConnectionServerMonitorDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ConnectionServerMonitorDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ConnectionServerMonitorDetails) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ConnectionServerMonitorDetails) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


