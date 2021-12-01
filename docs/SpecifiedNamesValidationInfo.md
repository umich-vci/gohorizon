# SpecifiedNamesValidationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineName** | Pointer to **string** | The virtual machine name. | [optional] 
**SpecifiedNamesStatuses** | Pointer to **[]string** | The status information captured after validation. | [optional] 
**UserId** | Pointer to **string** | The SID of the user. | [optional] 
**UserName** | Pointer to **string** | The name of the user entitled to the VM. | [optional] 

## Methods

### NewSpecifiedNamesValidationInfo

`func NewSpecifiedNamesValidationInfo() *SpecifiedNamesValidationInfo`

NewSpecifiedNamesValidationInfo instantiates a new SpecifiedNamesValidationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecifiedNamesValidationInfoWithDefaults

`func NewSpecifiedNamesValidationInfoWithDefaults() *SpecifiedNamesValidationInfo`

NewSpecifiedNamesValidationInfoWithDefaults instantiates a new SpecifiedNamesValidationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineName

`func (o *SpecifiedNamesValidationInfo) GetMachineName() string`

GetMachineName returns the MachineName field if non-nil, zero value otherwise.

### GetMachineNameOk

`func (o *SpecifiedNamesValidationInfo) GetMachineNameOk() (*string, bool)`

GetMachineNameOk returns a tuple with the MachineName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineName

`func (o *SpecifiedNamesValidationInfo) SetMachineName(v string)`

SetMachineName sets MachineName field to given value.

### HasMachineName

`func (o *SpecifiedNamesValidationInfo) HasMachineName() bool`

HasMachineName returns a boolean if a field has been set.

### GetSpecifiedNamesStatuses

`func (o *SpecifiedNamesValidationInfo) GetSpecifiedNamesStatuses() []string`

GetSpecifiedNamesStatuses returns the SpecifiedNamesStatuses field if non-nil, zero value otherwise.

### GetSpecifiedNamesStatusesOk

`func (o *SpecifiedNamesValidationInfo) GetSpecifiedNamesStatusesOk() (*[]string, bool)`

GetSpecifiedNamesStatusesOk returns a tuple with the SpecifiedNamesStatuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecifiedNamesStatuses

`func (o *SpecifiedNamesValidationInfo) SetSpecifiedNamesStatuses(v []string)`

SetSpecifiedNamesStatuses sets SpecifiedNamesStatuses field to given value.

### HasSpecifiedNamesStatuses

`func (o *SpecifiedNamesValidationInfo) HasSpecifiedNamesStatuses() bool`

HasSpecifiedNamesStatuses returns a boolean if a field has been set.

### GetUserId

`func (o *SpecifiedNamesValidationInfo) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SpecifiedNamesValidationInfo) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SpecifiedNamesValidationInfo) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *SpecifiedNamesValidationInfo) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUserName

`func (o *SpecifiedNamesValidationInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *SpecifiedNamesValidationInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *SpecifiedNamesValidationInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *SpecifiedNamesValidationInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


