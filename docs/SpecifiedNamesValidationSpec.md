# SpecifiedNamesValidationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dedicated** | Pointer to **bool** | Indicates whether desktop pool is dedicated or floating. Default value is false. | [optional] 
**Id** | Pointer to **string** | ID of the desktop pool to which the manually defined virtual machines will belong. This is required only if virtual machines are being added to an existing pool. | [optional] 
**NamesSpec** | [**[]NamesSpec**](NamesSpec.md) | List of manually defined virtual machines and users. | 

## Methods

### NewSpecifiedNamesValidationSpec

`func NewSpecifiedNamesValidationSpec(namesSpec []NamesSpec, ) *SpecifiedNamesValidationSpec`

NewSpecifiedNamesValidationSpec instantiates a new SpecifiedNamesValidationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecifiedNamesValidationSpecWithDefaults

`func NewSpecifiedNamesValidationSpecWithDefaults() *SpecifiedNamesValidationSpec`

NewSpecifiedNamesValidationSpecWithDefaults instantiates a new SpecifiedNamesValidationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDedicated

`func (o *SpecifiedNamesValidationSpec) GetDedicated() bool`

GetDedicated returns the Dedicated field if non-nil, zero value otherwise.

### GetDedicatedOk

`func (o *SpecifiedNamesValidationSpec) GetDedicatedOk() (*bool, bool)`

GetDedicatedOk returns a tuple with the Dedicated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDedicated

`func (o *SpecifiedNamesValidationSpec) SetDedicated(v bool)`

SetDedicated sets Dedicated field to given value.

### HasDedicated

`func (o *SpecifiedNamesValidationSpec) HasDedicated() bool`

HasDedicated returns a boolean if a field has been set.

### GetId

`func (o *SpecifiedNamesValidationSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SpecifiedNamesValidationSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SpecifiedNamesValidationSpec) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SpecifiedNamesValidationSpec) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNamesSpec

`func (o *SpecifiedNamesValidationSpec) GetNamesSpec() []NamesSpec`

GetNamesSpec returns the NamesSpec field if non-nil, zero value otherwise.

### GetNamesSpecOk

`func (o *SpecifiedNamesValidationSpec) GetNamesSpecOk() (*[]NamesSpec, bool)`

GetNamesSpecOk returns a tuple with the NamesSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamesSpec

`func (o *SpecifiedNamesValidationSpec) SetNamesSpec(v []NamesSpec)`

SetNamesSpec sets NamesSpec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


