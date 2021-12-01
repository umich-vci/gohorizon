# VirtualCenterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateOverride** | Pointer to [**CertificateOverrideData**](CertificateOverrideData.md) |  | [optional] 
**Description** | Pointer to **string** | Human readable description of the Virtual Center instance. | [optional] 
**DisplayName** | Pointer to **string** | Human readable name of the Virtual Center instance. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the virtual center is enabled. | [optional] 
**Id** | Pointer to **string** | Unique ID of the Virtual Center. | [optional] 
**InstanceUuid** | Pointer to **string** | Virtual center&#39;s instanceUuid. | [optional] 
**Limits** | Pointer to [**VCLimits**](VCLimits.md) |  | [optional] 
**Port** | Pointer to **int32** | Port of the virtual center to connect to. | [optional] 
**SeSparseReclamationEnabled** | Pointer to **bool** | Indicates if Storage Efficiency Sparse (seSparse) reclamation is enabled. | [optional] 
**ServerName** | Pointer to **string** | Virtual Center&#39;s server name or IP address. | [optional] 
**StorageAcceleratorData** | Pointer to [**StorageAcceleratorData**](StorageAcceleratorData.md) |  | [optional] 
**UseSsl** | Pointer to **bool** | Indicates if SSL should be used when connecting to the server. | [optional] 
**UserName** | Pointer to **string** | User name to use for the connection. | [optional] 
**Version** | Pointer to **string** | Version of the Virtual Center. | [optional] 
**VmcDeployment** | Pointer to **bool** | Indicates if virtual center is deployed in VMC. | [optional] 

## Methods

### NewVirtualCenterInfo

`func NewVirtualCenterInfo() *VirtualCenterInfo`

NewVirtualCenterInfo instantiates a new VirtualCenterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCenterInfoWithDefaults

`func NewVirtualCenterInfoWithDefaults() *VirtualCenterInfo`

NewVirtualCenterInfoWithDefaults instantiates a new VirtualCenterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateOverride

`func (o *VirtualCenterInfo) GetCertificateOverride() CertificateOverrideData`

GetCertificateOverride returns the CertificateOverride field if non-nil, zero value otherwise.

### GetCertificateOverrideOk

`func (o *VirtualCenterInfo) GetCertificateOverrideOk() (*CertificateOverrideData, bool)`

GetCertificateOverrideOk returns a tuple with the CertificateOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateOverride

`func (o *VirtualCenterInfo) SetCertificateOverride(v CertificateOverrideData)`

SetCertificateOverride sets CertificateOverride field to given value.

### HasCertificateOverride

`func (o *VirtualCenterInfo) HasCertificateOverride() bool`

HasCertificateOverride returns a boolean if a field has been set.

### GetDescription

`func (o *VirtualCenterInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VirtualCenterInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VirtualCenterInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VirtualCenterInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *VirtualCenterInfo) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VirtualCenterInfo) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VirtualCenterInfo) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VirtualCenterInfo) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *VirtualCenterInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *VirtualCenterInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *VirtualCenterInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *VirtualCenterInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *VirtualCenterInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCenterInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCenterInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualCenterInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceUuid

`func (o *VirtualCenterInfo) GetInstanceUuid() string`

GetInstanceUuid returns the InstanceUuid field if non-nil, zero value otherwise.

### GetInstanceUuidOk

`func (o *VirtualCenterInfo) GetInstanceUuidOk() (*string, bool)`

GetInstanceUuidOk returns a tuple with the InstanceUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceUuid

`func (o *VirtualCenterInfo) SetInstanceUuid(v string)`

SetInstanceUuid sets InstanceUuid field to given value.

### HasInstanceUuid

`func (o *VirtualCenterInfo) HasInstanceUuid() bool`

HasInstanceUuid returns a boolean if a field has been set.

### GetLimits

`func (o *VirtualCenterInfo) GetLimits() VCLimits`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *VirtualCenterInfo) GetLimitsOk() (*VCLimits, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *VirtualCenterInfo) SetLimits(v VCLimits)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *VirtualCenterInfo) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetPort

`func (o *VirtualCenterInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VirtualCenterInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VirtualCenterInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VirtualCenterInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSeSparseReclamationEnabled

`func (o *VirtualCenterInfo) GetSeSparseReclamationEnabled() bool`

GetSeSparseReclamationEnabled returns the SeSparseReclamationEnabled field if non-nil, zero value otherwise.

### GetSeSparseReclamationEnabledOk

`func (o *VirtualCenterInfo) GetSeSparseReclamationEnabledOk() (*bool, bool)`

GetSeSparseReclamationEnabledOk returns a tuple with the SeSparseReclamationEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeSparseReclamationEnabled

`func (o *VirtualCenterInfo) SetSeSparseReclamationEnabled(v bool)`

SetSeSparseReclamationEnabled sets SeSparseReclamationEnabled field to given value.

### HasSeSparseReclamationEnabled

`func (o *VirtualCenterInfo) HasSeSparseReclamationEnabled() bool`

HasSeSparseReclamationEnabled returns a boolean if a field has been set.

### GetServerName

`func (o *VirtualCenterInfo) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *VirtualCenterInfo) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *VirtualCenterInfo) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *VirtualCenterInfo) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetStorageAcceleratorData

`func (o *VirtualCenterInfo) GetStorageAcceleratorData() StorageAcceleratorData`

GetStorageAcceleratorData returns the StorageAcceleratorData field if non-nil, zero value otherwise.

### GetStorageAcceleratorDataOk

`func (o *VirtualCenterInfo) GetStorageAcceleratorDataOk() (*StorageAcceleratorData, bool)`

GetStorageAcceleratorDataOk returns a tuple with the StorageAcceleratorData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAcceleratorData

`func (o *VirtualCenterInfo) SetStorageAcceleratorData(v StorageAcceleratorData)`

SetStorageAcceleratorData sets StorageAcceleratorData field to given value.

### HasStorageAcceleratorData

`func (o *VirtualCenterInfo) HasStorageAcceleratorData() bool`

HasStorageAcceleratorData returns a boolean if a field has been set.

### GetUseSsl

`func (o *VirtualCenterInfo) GetUseSsl() bool`

GetUseSsl returns the UseSsl field if non-nil, zero value otherwise.

### GetUseSslOk

`func (o *VirtualCenterInfo) GetUseSslOk() (*bool, bool)`

GetUseSslOk returns a tuple with the UseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSsl

`func (o *VirtualCenterInfo) SetUseSsl(v bool)`

SetUseSsl sets UseSsl field to given value.

### HasUseSsl

`func (o *VirtualCenterInfo) HasUseSsl() bool`

HasUseSsl returns a boolean if a field has been set.

### GetUserName

`func (o *VirtualCenterInfo) GetUserName() string`

GetUserName returns the UserName field if non-nil, zero value otherwise.

### GetUserNameOk

`func (o *VirtualCenterInfo) GetUserNameOk() (*string, bool)`

GetUserNameOk returns a tuple with the UserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserName

`func (o *VirtualCenterInfo) SetUserName(v string)`

SetUserName sets UserName field to given value.

### HasUserName

`func (o *VirtualCenterInfo) HasUserName() bool`

HasUserName returns a boolean if a field has been set.

### GetVersion

`func (o *VirtualCenterInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VirtualCenterInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VirtualCenterInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VirtualCenterInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetVmcDeployment

`func (o *VirtualCenterInfo) GetVmcDeployment() bool`

GetVmcDeployment returns the VmcDeployment field if non-nil, zero value otherwise.

### GetVmcDeploymentOk

`func (o *VirtualCenterInfo) GetVmcDeploymentOk() (*bool, bool)`

GetVmcDeploymentOk returns a tuple with the VmcDeployment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmcDeployment

`func (o *VirtualCenterInfo) SetVmcDeployment(v bool)`

SetVmcDeployment sets VmcDeployment field to given value.

### HasVmcDeployment

`func (o *VirtualCenterInfo) HasVmcDeployment() bool`

HasVmcDeployment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


