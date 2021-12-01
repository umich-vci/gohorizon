# VCLimits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstantCloneEngineProvisioningLimit** | Pointer to **int32** | Maximum concurrent instant clone engine provisioning operations. This property has a default value of 20. This property has a minimum value of 1. | [optional] 
**PowerOperationsLimit** | Pointer to **int32** | Maximum concurrent virtual center power operations. This property has a default value of 50. This property has a minimum value of 1. | [optional] 
**ProvisioningLimit** | Pointer to **int32** | Maximum concurrent virtual center provisioning operations. This property has a default value of 20. This property has a minimum value of 1. | [optional] 

## Methods

### NewVCLimits

`func NewVCLimits() *VCLimits`

NewVCLimits instantiates a new VCLimits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCLimitsWithDefaults

`func NewVCLimitsWithDefaults() *VCLimits`

NewVCLimitsWithDefaults instantiates a new VCLimits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstantCloneEngineProvisioningLimit

`func (o *VCLimits) GetInstantCloneEngineProvisioningLimit() int32`

GetInstantCloneEngineProvisioningLimit returns the InstantCloneEngineProvisioningLimit field if non-nil, zero value otherwise.

### GetInstantCloneEngineProvisioningLimitOk

`func (o *VCLimits) GetInstantCloneEngineProvisioningLimitOk() (*int32, bool)`

GetInstantCloneEngineProvisioningLimitOk returns a tuple with the InstantCloneEngineProvisioningLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantCloneEngineProvisioningLimit

`func (o *VCLimits) SetInstantCloneEngineProvisioningLimit(v int32)`

SetInstantCloneEngineProvisioningLimit sets InstantCloneEngineProvisioningLimit field to given value.

### HasInstantCloneEngineProvisioningLimit

`func (o *VCLimits) HasInstantCloneEngineProvisioningLimit() bool`

HasInstantCloneEngineProvisioningLimit returns a boolean if a field has been set.

### GetPowerOperationsLimit

`func (o *VCLimits) GetPowerOperationsLimit() int32`

GetPowerOperationsLimit returns the PowerOperationsLimit field if non-nil, zero value otherwise.

### GetPowerOperationsLimitOk

`func (o *VCLimits) GetPowerOperationsLimitOk() (*int32, bool)`

GetPowerOperationsLimitOk returns a tuple with the PowerOperationsLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOperationsLimit

`func (o *VCLimits) SetPowerOperationsLimit(v int32)`

SetPowerOperationsLimit sets PowerOperationsLimit field to given value.

### HasPowerOperationsLimit

`func (o *VCLimits) HasPowerOperationsLimit() bool`

HasPowerOperationsLimit returns a boolean if a field has been set.

### GetProvisioningLimit

`func (o *VCLimits) GetProvisioningLimit() int32`

GetProvisioningLimit returns the ProvisioningLimit field if non-nil, zero value otherwise.

### GetProvisioningLimitOk

`func (o *VCLimits) GetProvisioningLimitOk() (*int32, bool)`

GetProvisioningLimitOk returns a tuple with the ProvisioningLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningLimit

`func (o *VCLimits) SetProvisioningLimit(v int32)`

SetProvisioningLimit sets ProvisioningLimit field to given value.

### HasProvisioningLimit

`func (o *VCLimits) HasProvisioningLimit() bool`

HasProvisioningLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


