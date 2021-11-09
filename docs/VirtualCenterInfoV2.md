# VirtualCenterInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateOverride** | Pointer to [**CertificateOverrideData**](CertificateOverrideData.md) |  | [optional] 
**DeploymentType** | **string** | Indicates different environments that Horizon can be deployed into. * GENERAL: Horizon is deployed on On-premises. * AZURE: Horizon is deployed on Azure. * AWS: Horizon is deployed on AWS. * DELL_EMC: Horizon is deployed on Dell EMC. * GOOGLE: Horizon is deployed on Google Cloud. * ORACLE: Horizon is deployed on Oracle Cloud. | 
**Description** | Pointer to **string** | Human readable description of the Virtual Center instance. | [optional] 
**DisplayName** | Pointer to **string** | Human readable name of the Virtual Center instance. | [optional] 
**Enabled** | **bool** | Indicates if the virtual center is enabled. | 
**HasVirtualTpmPools** | Pointer to **bool** | Indicates if there is any instant clone Desktop pool associated with this Virtual Center which has addVirtualTPM set | [optional] 
**Id** | **string** | Unique ID of the Virtual Center. | 
**InstanceUuid** | Pointer to **string** | Virtual center&#39;s instanceUuid. | [optional] 
**Limits** | [**VCLimits**](VCLimits.md) |  | 
**MaintenanceMode** | Pointer to **bool** | Indicates if maintenance or upgrade task is scheduled on Virtual center or hosts | [optional] 
**Port** | **int32** | Port of the virtual center to connect to. | 
**SeSparseReclamationEnabled** | **bool** | Indicates if Storage Efficiency Sparse (seSparse) reclamation is enabled. | 
**ServerName** | **string** | Virtual Center&#39;s server name or IP address. | 
**StorageAcceleratorData** | [**StorageAcceleratorData**](StorageAcceleratorData.md) |  | 
**UseSsl** | **bool** | Indicates if SSL should be used when connecting to the server. | 
**UserName** | **string** | User name to use for the connection. | 
**Version** | **string** | Version of the Virtual Center. | 

## Methods

### NewVirtualCenterInfoV2

`func NewVirtualCenterInfoV2(deploymentType string, enabled bool, id string, limits VCLimits, port int32, seSparseReclamationEnabled bool, serverName string, storageAcceleratorData StorageAcceleratorData, useSsl bool, userName string, version string, ) *VirtualCenterInfoV2`

NewVirtualCenterInfoV2 instantiates a new VirtualCenterInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCenterInfoV2WithDefaults

`func NewVirtualCenterInfoV2WithDefaults() *VirtualCenterInfoV2`

NewVirtualCenterInfoV2WithDefaults instantiates a new VirtualCenterInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateOverride

`func (o *VirtualCenterInfoV2) GetCertificateOverride() CertificateOverrideData`

GetCertificateOverride returns the CertificateOverride field if non-nil, zero value otherwise.

### GetCertificateOverrideOk

`func (o *VirtualCenterInfoV2) GetCertificateOverrideOk() (*CertificateOverrideData, bool)`

GetCertificateOverrideOk returns a tuple with the CertificateOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOverride

`func (o *VirtualCenterInfoV2) SetCertificateOverride(v CertificateOverrideData)`

SetCertificateOverride sets CertificateOverride field to given value.

### HasCertificateOverride

`func (o *VirtualCenterInfoV2) HasCertificateOverride() bool`

HasCertificateOverride returns a boolean if a field has been set.

### GetDeploymentType

`func (o *VirtualCenterInfoV2) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *VirtualCenterInfoV2) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *VirtualCenterInfoV2) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.


### GetDescription

`func (o *VirtualCenterInfoV2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCenterInfoV2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCenterInfoV2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCenterInfoV2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *VirtualCenterInfoV2) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VirtualCenterInfoV2) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VirtualCenterInfoV2) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VirtualCenterInfoV2) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *VirtualCenterInfoV2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VirtualCenterInfoV2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VirtualCenterInfoV2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetHasVirtualTpmPools

`func (o *VirtualCenterInfoV2) GetHasVirtualTpmPools() bool`

GetHasVirtualTpmPools returns the HasVirtualTpmPools field if non-nil, zero value otherwise.

### GetHasVirtualTpmPoolsOk

`func (o *VirtualCenterInfoV2) GetHasVirtualTpmPoolsOk() (*bool, bool)`

GetHasVirtualTpmPoolsOk returns a tuple with the HasVirtualTpmPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVirtualTpmPools

`func (o *VirtualCenterInfoV2) SetHasVirtualTpmPools(v bool)`

SetHasVirtualTpmPools sets HasVirtualTpmPools field to given value.

### HasHasVirtualTpmPools

`func (o *VirtualCenterInfoV2) HasHasVirtualTpmPools() bool`

HasHasVirtualTpmPools returns a boolean if a field has been set.

### GetId

`func (o *VirtualCenterInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCenterInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCenterInfoV2) SetId(v string)`

SetId sets Id field to given value.


### GetInstanceUuid

`func (o *VirtualCenterInfoV2) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *VirtualCenterInfoV2) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *VirtualCenterInfoV2) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.

### HasInstanceUuid

`func (o *VirtualCenterInfoV2) HasInstanceUuid() bool`

HasInstanceUuid returns a boolean if a field has been set.

### GetLimits

`func (o *VirtualCenterInfoV2) GetLimits() VCLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *VirtualCenterInfoV2) GetLimitsOk() (*VCLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *VirtualCenterInfoV2) SetLimits(v VCLimits)`

SetLimits sets Limits field to given value.


### GetMaintenanceMode

`func (o *VirtualCenterInfoV2) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *VirtualCenterInfoV2) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *VirtualCenterInfoV2) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.

### HasMaintenanceMode

`func (o *VirtualCenterInfoV2) HasMaintenanceMode() bool`

HasMaintenanceMode returns a boolean if a field has been set.

### GetPort

`func (o *VirtualCenterInfoV2) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VirtualCenterInfoV2) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VirtualCenterInfoV2) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSeSparseReclamationEnabled

`func (o *VirtualCenterInfoV2) GetSeSparseReclamationEnabled() bool`

GetSeSparseReclamationEnabled returns the SeSparseReclamationEnabled field if non-nil, zero value otherwise.

### GetSeSparseReclamationEnabledOk

`func (o *VirtualCenterInfoV2) GetSeSparseReclamationEnabledOk() (*bool, bool)`

GetSeSparseReclamationEnabledOk returns a tuple with the SeSparseReclamationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeSparseReclamationEnabled

`func (o *VirtualCenterInfoV2) SetSeSparseReclamationEnabled(v bool)`

SetSeSparseReclamationEnabled sets SeSparseReclamationEnabled field to given value.


### GetServerName

`func (o *VirtualCenterInfoV2) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *VirtualCenterInfoV2) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *VirtualCenterInfoV2) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetStorageAcceleratorData

`func (o *VirtualCenterInfoV2) GetStorageAcceleratorData() StorageAcceleratorData`

GetStorageAcceleratorData returns the StorageAcceleratorData field if non-nil, zero value otherwise.

### GetStorageAcceleratorDataOk

`func (o *VirtualCenterInfoV2) GetStorageAcceleratorDataOk() (*StorageAcceleratorData, bool)`

GetStorageAcceleratorDataOk returns a tuple with the StorageAcceleratorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAcceleratorData

`func (o *VirtualCenterInfoV2) SetStorageAcceleratorData(v StorageAcceleratorData)`

SetStorageAcceleratorData sets StorageAcceleratorData field to given value.


### GetUseSsl

`func (o *VirtualCenterInfoV2) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *VirtualCenterInfoV2) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *VirtualCenterInfoV2) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.


### GetUserName

`func (o *VirtualCenterInfoV2) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VirtualCenterInfoV2) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VirtualCenterInfoV2) SetUserName(v string)`

SetUserName sets UserName field to given value.


### GetVersion

`func (o *VirtualCenterInfoV2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualCenterInfoV2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualCenterInfoV2) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


