# MachineDeleteSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MachineDeleteData** | Pointer to [**MachineDeleteData**](MachineDeleteData.md) |  | [optional] 
**MachineIds** | **[]string** | Machine Ids representing the machines to be deleted. | 

## Methods

### NewMachineDeleteSpec

`func NewMachineDeleteSpec(machineIds []string, ) *MachineDeleteSpec`

NewMachineDeleteSpec instantiates a new MachineDeleteSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMachineDeleteSpecWithDefaults

`func NewMachineDeleteSpecWithDefaults() *MachineDeleteSpec`

NewMachineDeleteSpecWithDefaults instantiates a new MachineDeleteSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMachineDeleteData

`func (o *MachineDeleteSpec) GetMachineDeleteData() MachineDeleteData`

GetMachineDeleteData returns the MachineDeleteData field if non-nil, zero value otherwise.

### GetMachineDeleteDataOk

`func (o *MachineDeleteSpec) GetMachineDeleteDataOk() (*MachineDeleteData, bool)`

GetMachineDeleteDataOk returns a tuple with the MachineDeleteData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDeleteData

`func (o *MachineDeleteSpec) SetMachineDeleteData(v MachineDeleteData)`

SetMachineDeleteData sets MachineDeleteData field to given value.

### HasMachineDeleteData

`func (o *MachineDeleteSpec) HasMachineDeleteData() bool`

HasMachineDeleteData returns a boolean if a field has been set.

### GetMachineIds

`func (o *MachineDeleteSpec) GetMachineIds() []string`

GetMachineIds returns the MachineIds field if non-nil, zero value otherwise.

### GetMachineIdsOk

`func (o *MachineDeleteSpec) GetMachineIdsOk() (*[]string, bool)`

GetMachineIdsOk returns a tuple with the MachineIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineIds

`func (o *MachineDeleteSpec) SetMachineIds(v []string)`

SetMachineIds sets MachineIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


