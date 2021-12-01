# NamesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineName** | **string** | The virtual machine name. | 
**UserName** | Pointer to **string** | The name of the user entitled to the VM, used only in case of persistent desktop pools. | [optional] 

## Methods

### NewNamesSpec

`func NewNamesSpec(machineName string, ) *NamesSpec`

NewNamesSpec instantiates a new NamesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNamesSpecWithDefaults

`func NewNamesSpecWithDefaults() *NamesSpec`

NewNamesSpecWithDefaults instantiates a new NamesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineName

`func (o *NamesSpec) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *NamesSpec) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *NamesSpec) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.


### GetUserName

`func (o *NamesSpec) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *NamesSpec) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *NamesSpec) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *NamesSpec) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


