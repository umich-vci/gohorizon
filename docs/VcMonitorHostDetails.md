# VCMonitorHostDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | **string** | API Version of the ESX Host. | 
**ClusterName** | **string** | Name of the cluster of the ESX Host in the Virtual Center. | 
**Name** | **string** | ESX Host name or IP address. | 
**Version** | **string** | ESX Host version. | 
**VgpuTypes** | Pointer to **[]string** | Types of vGPUs supported by this host. | [optional] 

## Methods

### NewVCMonitorHostDetails

`func NewVCMonitorHostDetails(apiVersion string, clusterName string, name string, version string, ) *VCMonitorHostDetails`

NewVCMonitorHostDetails instantiates a new VCMonitorHostDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVCMonitorHostDetailsWithDefaults

`func NewVCMonitorHostDetailsWithDefaults() *VCMonitorHostDetails`

NewVCMonitorHostDetailsWithDefaults instantiates a new VCMonitorHostDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *VCMonitorHostDetails) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *VCMonitorHostDetails) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *VCMonitorHostDetails) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.


### GetClusterName

`func (o *VCMonitorHostDetails) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VCMonitorHostDetails) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VCMonitorHostDetails) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetName

`func (o *VCMonitorHostDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VCMonitorHostDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VCMonitorHostDetails) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *VCMonitorHostDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VCMonitorHostDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VCMonitorHostDetails) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetVgpuTypes

`func (o *VCMonitorHostDetails) GetVgpuTypes() []string`

GetVgpuTypes returns the VgpuTypes field if non-nil, zero value otherwise.

### GetVgpuTypesOk

`func (o *VCMonitorHostDetails) GetVgpuTypesOk() (*[]string, bool)`

GetVgpuTypesOk returns a tuple with the VgpuTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVgpuTypes

`func (o *VCMonitorHostDetails) SetVgpuTypes(v []string)`

SetVgpuTypes sets VgpuTypes field to given value.

### HasVgpuTypes

`func (o *VCMonitorHostDetails) HasVgpuTypes() bool`

HasVgpuTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


