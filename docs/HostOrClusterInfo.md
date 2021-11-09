# HostOrClusterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Container** | Pointer to [**HostOrClusterContainer**](HostOrClusterContainer.md) |  | [optional] 
**Details** | Pointer to [**HostOrClusterDetails**](HostOrClusterDetails.md) |  | [optional] 
**Id** | **string** | Unique ID representing a host or cluster. | 

## Methods

### NewHostOrClusterInfo

`func NewHostOrClusterInfo(id string, ) *HostOrClusterInfo`

NewHostOrClusterInfo instantiates a new HostOrClusterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOrClusterInfoWithDefaults

`func NewHostOrClusterInfoWithDefaults() *HostOrClusterInfo`

NewHostOrClusterInfoWithDefaults instantiates a new HostOrClusterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainer

`func (o *HostOrClusterInfo) GetContainer() HostOrClusterContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *HostOrClusterInfo) GetContainerOk() (*HostOrClusterContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *HostOrClusterInfo) SetContainer(v HostOrClusterContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *HostOrClusterInfo) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### GetDetails

`func (o *HostOrClusterInfo) GetDetails() HostOrClusterDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *HostOrClusterInfo) GetDetailsOk() (*HostOrClusterDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *HostOrClusterInfo) SetDetails(v HostOrClusterDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *HostOrClusterInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *HostOrClusterInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostOrClusterInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostOrClusterInfo) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


