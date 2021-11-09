# FederationAccessGroupInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deletable** | Pointer to **bool** | Indicates whether this federation access group can be deleted. | [optional] 
**Description** | Pointer to **string** | Federation Access group description. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this Federation Access Group. | [optional] 
**Name** | Pointer to **string** | Federation Access group name. | [optional] 

## Methods

### NewFederationAccessGroupInfo

`func NewFederationAccessGroupInfo() *FederationAccessGroupInfo`

NewFederationAccessGroupInfo instantiates a new FederationAccessGroupInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationAccessGroupInfoWithDefaults

`func NewFederationAccessGroupInfoWithDefaults() *FederationAccessGroupInfo`

NewFederationAccessGroupInfoWithDefaults instantiates a new FederationAccessGroupInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeletable

`func (o *FederationAccessGroupInfo) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *FederationAccessGroupInfo) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *FederationAccessGroupInfo) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *FederationAccessGroupInfo) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDescription

`func (o *FederationAccessGroupInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederationAccessGroupInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederationAccessGroupInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederationAccessGroupInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *FederationAccessGroupInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederationAccessGroupInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederationAccessGroupInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FederationAccessGroupInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FederationAccessGroupInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederationAccessGroupInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederationAccessGroupInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FederationAccessGroupInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


