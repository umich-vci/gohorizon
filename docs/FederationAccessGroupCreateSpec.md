# FederationAccessGroupCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Federation Access Group Description. | [optional] 
**Name** | **string** | Federation Access Group Name. This property must contain only alphanumerics, underscores, and dashes. | 

## Methods

### NewFederationAccessGroupCreateSpec

`func NewFederationAccessGroupCreateSpec(name string, ) *FederationAccessGroupCreateSpec`

NewFederationAccessGroupCreateSpec instantiates a new FederationAccessGroupCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationAccessGroupCreateSpecWithDefaults

`func NewFederationAccessGroupCreateSpecWithDefaults() *FederationAccessGroupCreateSpec`

NewFederationAccessGroupCreateSpecWithDefaults instantiates a new FederationAccessGroupCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FederationAccessGroupCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FederationAccessGroupCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FederationAccessGroupCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FederationAccessGroupCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *FederationAccessGroupCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FederationAccessGroupCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FederationAccessGroupCreateSpec) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


