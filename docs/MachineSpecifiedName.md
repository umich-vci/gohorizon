# MachineSpecifiedName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the machine. | 
**UserId** | Pointer to **string** | The unique SID of the user assigned to the machine. | [optional] 

## Methods

### NewMachineSpecifiedName

`func NewMachineSpecifiedName(name string, ) *MachineSpecifiedName`

NewMachineSpecifiedName instantiates a new MachineSpecifiedName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineSpecifiedNameWithDefaults

`func NewMachineSpecifiedNameWithDefaults() *MachineSpecifiedName`

NewMachineSpecifiedNameWithDefaults instantiates a new MachineSpecifiedName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MachineSpecifiedName) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MachineSpecifiedName) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MachineSpecifiedName) SetName(v string)`

SetName sets Name field to given value.


### GetUserId

`func (o *MachineSpecifiedName) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *MachineSpecifiedName) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *MachineSpecifiedName) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *MachineSpecifiedName) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


