# LocalAccessGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deletable** | Pointer to **bool** | Indicates whether this access group can be deleted. | [optional] 
**Description** | Pointer to **string** | Access group description. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this access group. | [optional] 
**Name** | Pointer to **string** | Access group name. | [optional] 

## Methods

### NewLocalAccessGroupInfo

`func NewLocalAccessGroupInfo() *LocalAccessGroupInfo`

NewLocalAccessGroupInfo instantiates a new LocalAccessGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalAccessGroupInfoWithDefaults

`func NewLocalAccessGroupInfoWithDefaults() *LocalAccessGroupInfo`

NewLocalAccessGroupInfoWithDefaults instantiates a new LocalAccessGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletable

`func (o *LocalAccessGroupInfo) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *LocalAccessGroupInfo) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *LocalAccessGroupInfo) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *LocalAccessGroupInfo) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDescription

`func (o *LocalAccessGroupInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LocalAccessGroupInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LocalAccessGroupInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LocalAccessGroupInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *LocalAccessGroupInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalAccessGroupInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalAccessGroupInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LocalAccessGroupInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LocalAccessGroupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalAccessGroupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalAccessGroupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LocalAccessGroupInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


