# MachineAliasSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdUserId** | Pointer to **string** | Sid of the user | [optional] 
**AliasName** | Pointer to **string** | Alias name of the user. | [optional] 

## Methods

### NewMachineAliasSpec

`func NewMachineAliasSpec() *MachineAliasSpec`

NewMachineAliasSpec instantiates a new MachineAliasSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineAliasSpecWithDefaults

`func NewMachineAliasSpecWithDefaults() *MachineAliasSpec`

NewMachineAliasSpecWithDefaults instantiates a new MachineAliasSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdUserId

`func (o *MachineAliasSpec) GetAdUserId() string`

GetAdUserId returns the AdUserId field if non-nil, zero value otherwise.

### GetAdUserIdOk

`func (o *MachineAliasSpec) GetAdUserIdOk() (*string, bool)`

GetAdUserIdOk returns a tuple with the AdUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserId

`func (o *MachineAliasSpec) SetAdUserId(v string)`

SetAdUserId sets AdUserId field to given value.

### HasAdUserId

`func (o *MachineAliasSpec) HasAdUserId() bool`

HasAdUserId returns a boolean if a field has been set.

### GetAliasName

`func (o *MachineAliasSpec) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *MachineAliasSpec) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *MachineAliasSpec) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *MachineAliasSpec) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


