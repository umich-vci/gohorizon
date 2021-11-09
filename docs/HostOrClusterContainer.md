# HostOrClusterContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]HostOrClusterInfo**](HostOrClusterInfo.md) | Contents of the container. These may be hosts or clusters or further nested containers. | [optional] 
**Name** | **string** | Host or cluster container node display name. | 
**Path** | **string** | Host or cluster container node path. | 
**Type** | **string** | Type of container. * FOLDER: A folder container. * OTHER: Other container type. | 

## Methods

### NewHostOrClusterContainer

`func NewHostOrClusterContainer(name string, path string, type_ string, ) *HostOrClusterContainer`

NewHostOrClusterContainer instantiates a new HostOrClusterContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostOrClusterContainerWithDefaults

`func NewHostOrClusterContainerWithDefaults() *HostOrClusterContainer`

NewHostOrClusterContainerWithDefaults instantiates a new HostOrClusterContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *HostOrClusterContainer) GetChildren() []HostOrClusterInfo`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *HostOrClusterContainer) GetChildrenOk() (*[]HostOrClusterInfo, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *HostOrClusterContainer) SetChildren(v []HostOrClusterInfo)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *HostOrClusterContainer) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetName

`func (o *HostOrClusterContainer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HostOrClusterContainer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HostOrClusterContainer) SetName(v string)`

SetName sets Name field to given value.


### GetPath

`func (o *HostOrClusterContainer) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *HostOrClusterContainer) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *HostOrClusterContainer) SetPath(v string)`

SetPath sets Path field to given value.


### GetType

`func (o *HostOrClusterContainer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *HostOrClusterContainer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *HostOrClusterContainer) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


