# FarmAutomatedSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomizationSettings** | [**FarmCustomizationSettingsCreateSpec**](FarmCustomizationSettingsCreateSpec.md) |  | 
**EnableProvisioning** | Pointer to **bool** | Indicates whether to enable provisioning immediately. Default value is true. | [optional] 
**MaxSessionType** | **string** | RDS Server type for max sessions. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions. | 
**MaxSessions** | Pointer to **int32** | Maximum number of sessions allowed for RDS Server. This is required if max_session_type is set to LIMITED. | [optional] 
**MinReadyVms** | Pointer to **int32** | Minimum number of ready (provisioned) RDS Servers during Instant clone maintenance operations. Use this setting to perform machine maintenance operations in a rolling fashion. Increasing this count may decrease the concurrency for Instant clone maintenance operations for the automated farm. Default value is 0. | [optional] 
**Nics** | Pointer to [**[]FarmNetworkInterfaceCardSettingsCreateSpec**](FarmNetworkInterfaceCardSettingsCreateSpec.md) | Network interface card settings for RDS Servers provisioned for this farm. A NIC may appear at most once in these settings and must be present on this RDS Server&#39;s parent&#39;s snapshot. Not all NICs need be configured. Any that are not will use default settings. | [optional] 
**PatternNamingSettings** | [**FarmRDSServersPatternNamingSettingsCreateSpec**](FarmRDSServersPatternNamingSettingsCreateSpec.md) |  | 
**ProvisioningSettings** | [**FarmProvisioningSettingsCreateSpec**](FarmProvisioningSettingsCreateSpec.md) |  | 
**StopProvisioningOnError** | Pointer to **bool** | Indicates whether provisioning on all VMs stops on error. Default value is true. | [optional] 
**StorageSettings** | [**FarmStorageSettingsCreateSpec**](FarmStorageSettingsCreateSpec.md) |  | 
**TransparentPageSharingScope** | Pointer to **string** | Transparent page sharing scope for the farm. Default value is VM. * VM: Inter-VM page sharing is not permitted. * FARM: Inter-VM page sharing among VMs belonging to the same automated farm is permitted. * POD: Inter-VM page sharing among VMs belonging to the same Pod is permitted. * GLOBAL: Inter-VM page sharing among all VMs on the same host is permitted. | [optional] 
**VcenterId** | **string** | ID of the virtual center server. | 

## Methods

### NewFarmAutomatedSettingsCreateSpec

`func NewFarmAutomatedSettingsCreateSpec(customizationSettings FarmCustomizationSettingsCreateSpec, maxSessionType string, patternNamingSettings FarmRDSServersPatternNamingSettingsCreateSpec, provisioningSettings FarmProvisioningSettingsCreateSpec, storageSettings FarmStorageSettingsCreateSpec, vcenterId string, ) *FarmAutomatedSettingsCreateSpec`

NewFarmAutomatedSettingsCreateSpec instantiates a new FarmAutomatedSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmAutomatedSettingsCreateSpecWithDefaults

`func NewFarmAutomatedSettingsCreateSpecWithDefaults() *FarmAutomatedSettingsCreateSpec`

NewFarmAutomatedSettingsCreateSpecWithDefaults instantiates a new FarmAutomatedSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomizationSettings

`func (o *FarmAutomatedSettingsCreateSpec) GetCustomizationSettings() FarmCustomizationSettingsCreateSpec`

GetCustomizationSettings returns the CustomizationSettings field if non-nil, zero value otherwise.

### GetCustomizationSettingsOk

`func (o *FarmAutomatedSettingsCreateSpec) GetCustomizationSettingsOk() (*FarmCustomizationSettingsCreateSpec, bool)`

GetCustomizationSettingsOk returns a tuple with the CustomizationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomizationSettings

`func (o *FarmAutomatedSettingsCreateSpec) SetCustomizationSettings(v FarmCustomizationSettingsCreateSpec)`

SetCustomizationSettings sets CustomizationSettings field to given value.


### GetEnableProvisioning

`func (o *FarmAutomatedSettingsCreateSpec) GetEnableProvisioning() bool`

GetEnableProvisioning returns the EnableProvisioning field if non-nil, zero value otherwise.

### GetEnableProvisioningOk

`func (o *FarmAutomatedSettingsCreateSpec) GetEnableProvisioningOk() (*bool, bool)`

GetEnableProvisioningOk returns a tuple with the EnableProvisioning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableProvisioning

`func (o *FarmAutomatedSettingsCreateSpec) SetEnableProvisioning(v bool)`

SetEnableProvisioning sets EnableProvisioning field to given value.

### HasEnableProvisioning

`func (o *FarmAutomatedSettingsCreateSpec) HasEnableProvisioning() bool`

HasEnableProvisioning returns a boolean if a field has been set.

### GetMaxSessionType

`func (o *FarmAutomatedSettingsCreateSpec) GetMaxSessionType() string`

GetMaxSessionType returns the MaxSessionType field if non-nil, zero value otherwise.

### GetMaxSessionTypeOk

`func (o *FarmAutomatedSettingsCreateSpec) GetMaxSessionTypeOk() (*string, bool)`

GetMaxSessionTypeOk returns a tuple with the MaxSessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionType

`func (o *FarmAutomatedSettingsCreateSpec) SetMaxSessionType(v string)`

SetMaxSessionType sets MaxSessionType field to given value.


### GetMaxSessions

`func (o *FarmAutomatedSettingsCreateSpec) GetMaxSessions() int32`

GetMaxSessions returns the MaxSessions field if non-nil, zero value otherwise.

### GetMaxSessionsOk

`func (o *FarmAutomatedSettingsCreateSpec) GetMaxSessionsOk() (*int32, bool)`

GetMaxSessionsOk returns a tuple with the MaxSessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessions

`func (o *FarmAutomatedSettingsCreateSpec) SetMaxSessions(v int32)`

SetMaxSessions sets MaxSessions field to given value.

### HasMaxSessions

`func (o *FarmAutomatedSettingsCreateSpec) HasMaxSessions() bool`

HasMaxSessions returns a boolean if a field has been set.

### GetMinReadyVms

`func (o *FarmAutomatedSettingsCreateSpec) GetMinReadyVms() int32`

GetMinReadyVms returns the MinReadyVms field if non-nil, zero value otherwise.

### GetMinReadyVmsOk

`func (o *FarmAutomatedSettingsCreateSpec) GetMinReadyVmsOk() (*int32, bool)`

GetMinReadyVmsOk returns a tuple with the MinReadyVms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReadyVms

`func (o *FarmAutomatedSettingsCreateSpec) SetMinReadyVms(v int32)`

SetMinReadyVms sets MinReadyVms field to given value.

### HasMinReadyVms

`func (o *FarmAutomatedSettingsCreateSpec) HasMinReadyVms() bool`

HasMinReadyVms returns a boolean if a field has been set.

### GetNics

`func (o *FarmAutomatedSettingsCreateSpec) GetNics() []FarmNetworkInterfaceCardSettingsCreateSpec`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *FarmAutomatedSettingsCreateSpec) GetNicsOk() (*[]FarmNetworkInterfaceCardSettingsCreateSpec, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *FarmAutomatedSettingsCreateSpec) SetNics(v []FarmNetworkInterfaceCardSettingsCreateSpec)`

SetNics sets Nics field to given value.

### HasNics

`func (o *FarmAutomatedSettingsCreateSpec) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetPatternNamingSettings

`func (o *FarmAutomatedSettingsCreateSpec) GetPatternNamingSettings() FarmRDSServersPatternNamingSettingsCreateSpec`

GetPatternNamingSettings returns the PatternNamingSettings field if non-nil, zero value otherwise.

### GetPatternNamingSettingsOk

`func (o *FarmAutomatedSettingsCreateSpec) GetPatternNamingSettingsOk() (*FarmRDSServersPatternNamingSettingsCreateSpec, bool)`

GetPatternNamingSettingsOk returns a tuple with the PatternNamingSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternNamingSettings

`func (o *FarmAutomatedSettingsCreateSpec) SetPatternNamingSettings(v FarmRDSServersPatternNamingSettingsCreateSpec)`

SetPatternNamingSettings sets PatternNamingSettings field to given value.


### GetProvisioningSettings

`func (o *FarmAutomatedSettingsCreateSpec) GetProvisioningSettings() FarmProvisioningSettingsCreateSpec`

GetProvisioningSettings returns the ProvisioningSettings field if non-nil, zero value otherwise.

### GetProvisioningSettingsOk

`func (o *FarmAutomatedSettingsCreateSpec) GetProvisioningSettingsOk() (*FarmProvisioningSettingsCreateSpec, bool)`

GetProvisioningSettingsOk returns a tuple with the ProvisioningSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningSettings

`func (o *FarmAutomatedSettingsCreateSpec) SetProvisioningSettings(v FarmProvisioningSettingsCreateSpec)`

SetProvisioningSettings sets ProvisioningSettings field to given value.


### GetStopProvisioningOnError

`func (o *FarmAutomatedSettingsCreateSpec) GetStopProvisioningOnError() bool`

GetStopProvisioningOnError returns the StopProvisioningOnError field if non-nil, zero value otherwise.

### GetStopProvisioningOnErrorOk

`func (o *FarmAutomatedSettingsCreateSpec) GetStopProvisioningOnErrorOk() (*bool, bool)`

GetStopProvisioningOnErrorOk returns a tuple with the StopProvisioningOnError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopProvisioningOnError

`func (o *FarmAutomatedSettingsCreateSpec) SetStopProvisioningOnError(v bool)`

SetStopProvisioningOnError sets StopProvisioningOnError field to given value.

### HasStopProvisioningOnError

`func (o *FarmAutomatedSettingsCreateSpec) HasStopProvisioningOnError() bool`

HasStopProvisioningOnError returns a boolean if a field has been set.

### GetStorageSettings

`func (o *FarmAutomatedSettingsCreateSpec) GetStorageSettings() FarmStorageSettingsCreateSpec`

GetStorageSettings returns the StorageSettings field if non-nil, zero value otherwise.

### GetStorageSettingsOk

`func (o *FarmAutomatedSettingsCreateSpec) GetStorageSettingsOk() (*FarmStorageSettingsCreateSpec, bool)`

GetStorageSettingsOk returns a tuple with the StorageSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSettings

`func (o *FarmAutomatedSettingsCreateSpec) SetStorageSettings(v FarmStorageSettingsCreateSpec)`

SetStorageSettings sets StorageSettings field to given value.


### GetTransparentPageSharingScope

`func (o *FarmAutomatedSettingsCreateSpec) GetTransparentPageSharingScope() string`

GetTransparentPageSharingScope returns the TransparentPageSharingScope field if non-nil, zero value otherwise.

### GetTransparentPageSharingScopeOk

`func (o *FarmAutomatedSettingsCreateSpec) GetTransparentPageSharingScopeOk() (*string, bool)`

GetTransparentPageSharingScopeOk returns a tuple with the TransparentPageSharingScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransparentPageSharingScope

`func (o *FarmAutomatedSettingsCreateSpec) SetTransparentPageSharingScope(v string)`

SetTransparentPageSharingScope sets TransparentPageSharingScope field to given value.

### HasTransparentPageSharingScope

`func (o *FarmAutomatedSettingsCreateSpec) HasTransparentPageSharingScope() bool`

HasTransparentPageSharingScope returns a boolean if a field has been set.

### GetVcenterId

`func (o *FarmAutomatedSettingsCreateSpec) GetVcenterId() string`

GetVcenterId returns the VcenterId field if non-nil, zero value otherwise.

### GetVcenterIdOk

`func (o *FarmAutomatedSettingsCreateSpec) GetVcenterIdOk() (*string, bool)`

GetVcenterIdOk returns a tuple with the VcenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVcenterId

`func (o *FarmAutomatedSettingsCreateSpec) SetVcenterId(v string)`

SetVcenterId sets VcenterId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


