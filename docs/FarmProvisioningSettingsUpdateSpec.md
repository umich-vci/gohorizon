# FarmProvisioningSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostOrClusterId** | **string** | Host or cluster where the machines are deployed in. For Instant clone farms it can be only be a cluster id. This can be modified only if there are no current operations ( operation is NONE). | 
**ResourcePoolId** | **string** | Resource pool to deploy the RDS Servers. | 

## Methods

### NewFarmProvisioningSettingsUpdateSpec

`func NewFarmProvisioningSettingsUpdateSpec(hostOrClusterId string, resourcePoolId string, ) *FarmProvisioningSettingsUpdateSpec`

NewFarmProvisioningSettingsUpdateSpec instantiates a new FarmProvisioningSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmProvisioningSettingsUpdateSpecWithDefaults

`func NewFarmProvisioningSettingsUpdateSpecWithDefaults() *FarmProvisioningSettingsUpdateSpec`

NewFarmProvisioningSettingsUpdateSpecWithDefaults instantiates a new FarmProvisioningSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostOrClusterId

`func (o *FarmProvisioningSettingsUpdateSpec) GetHostOrClusterId() string`

GetHostOrClusterId returns the HostOrClusterId field if non-nil, zero value otherwise.

### GetHostOrClusterIdOk

`func (o *FarmProvisioningSettingsUpdateSpec) GetHostOrClusterIdOk() (*string, bool)`

GetHostOrClusterIdOk returns a tuple with the HostOrClusterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostOrClusterId

`func (o *FarmProvisioningSettingsUpdateSpec) SetHostOrClusterId(v string)`

SetHostOrClusterId sets HostOrClusterId field to given value.


### GetResourcePoolId

`func (o *FarmProvisioningSettingsUpdateSpec) GetResourcePoolId() string`

GetResourcePoolId returns the ResourcePoolId field if non-nil, zero value otherwise.

### GetResourcePoolIdOk

`func (o *FarmProvisioningSettingsUpdateSpec) GetResourcePoolIdOk() (*string, bool)`

GetResourcePoolIdOk returns a tuple with the ResourcePoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePoolId

`func (o *FarmProvisioningSettingsUpdateSpec) SetResourcePoolId(v string)`

SetResourcePoolId sets ResourcePoolId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


