# ResourcePoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**[]ResourcePoolInfo**](ResourcePoolInfo.md) | Child nodes of the resource pool. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the resource pool. | [optional] 
**Name** | Pointer to **string** | Resource pool name. | [optional] 
**Path** | Pointer to **string** | Resource pool path. | [optional] 
**Type** | Pointer to **string** | Resource pool type. * HOST: Host used as a resource pool suitable for use in desktop pool/farm. * CLUSTER: Cluster used as a resource pool suitable for use in desktop pool/farm. * RESOURCE_POOL: Regular resource pool suitable for use in desktop pool/farm. * OTHER: Other resource type which cannot be used in desktop pool/farm. | [optional] 

## Methods

### NewResourcePoolInfo

`func NewResourcePoolInfo() *ResourcePoolInfo`

NewResourcePoolInfo instantiates a new ResourcePoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourcePoolInfoWithDefaults

`func NewResourcePoolInfoWithDefaults() *ResourcePoolInfo`

NewResourcePoolInfoWithDefaults instantiates a new ResourcePoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *ResourcePoolInfo) GetChildren() []ResourcePoolInfo`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *ResourcePoolInfo) GetChildrenOk() (*[]ResourcePoolInfo, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *ResourcePoolInfo) SetChildren(v []ResourcePoolInfo)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *ResourcePoolInfo) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetId

`func (o *ResourcePoolInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourcePoolInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourcePoolInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourcePoolInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourcePoolInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourcePoolInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourcePoolInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourcePoolInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *ResourcePoolInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ResourcePoolInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ResourcePoolInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ResourcePoolInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetType

`func (o *ResourcePoolInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourcePoolInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourcePoolInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ResourcePoolInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


