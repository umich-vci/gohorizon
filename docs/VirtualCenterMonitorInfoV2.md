# VirtualCenterMonitorInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | [**[]VCMonitorConnectionServerV2**](VCMonitorConnectionServerV2.md) | Information about the Virtual Center connections from each of the connection servers. | 
**Datastores** | [**[]VCMonitorDatastore**](VCMonitorDatastore.md) | Information about the datastores of the Virtual Center. | 
**DesktopPoolsAndFarmsCount** | **int32** | Number of Desktop Pools And Farms managed by the virtual center | 
**Details** | [**VCMonitorDetails**](VCMonitorDetails.md) |  | 
**Hosts** | [**[]VCMonitorHost**](VCMonitorHost.md) | Information about the ESX hosts added in the Virtual Center. | 
**Id** | **string** | Unique ID of the Virtual Center. | 
**Name** | **string** | Virtual Center server name or IP address. | 

## Methods

### NewVirtualCenterMonitorInfoV2

`func NewVirtualCenterMonitorInfoV2(connectionServers []VCMonitorConnectionServerV2, datastores []VCMonitorDatastore, desktopPoolsAndFarmsCount int32, details VCMonitorDetails, hosts []VCMonitorHost, id string, name string, ) *VirtualCenterMonitorInfoV2`

NewVirtualCenterMonitorInfoV2 instantiates a new VirtualCenterMonitorInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCenterMonitorInfoV2WithDefaults

`func NewVirtualCenterMonitorInfoV2WithDefaults() *VirtualCenterMonitorInfoV2`

NewVirtualCenterMonitorInfoV2WithDefaults instantiates a new VirtualCenterMonitorInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *VirtualCenterMonitorInfoV2) GetConnectionServers() []VCMonitorConnectionServerV2`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *VirtualCenterMonitorInfoV2) GetConnectionServersOk() (*[]VCMonitorConnectionServerV2, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *VirtualCenterMonitorInfoV2) SetConnectionServers(v []VCMonitorConnectionServerV2)`

SetConnectionServers sets ConnectionServers field to given value.


### GetDatastores

`func (o *VirtualCenterMonitorInfoV2) GetDatastores() []VCMonitorDatastore`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VirtualCenterMonitorInfoV2) GetDatastoresOk() (*[]VCMonitorDatastore, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VirtualCenterMonitorInfoV2) SetDatastores(v []VCMonitorDatastore)`

SetDatastores sets Datastores field to given value.


### GetDesktopPoolsAndFarmsCount

`func (o *VirtualCenterMonitorInfoV2) GetDesktopPoolsAndFarmsCount() int32`

GetDesktopPoolsAndFarmsCount returns the DesktopPoolsAndFarmsCount field if non-nil, zero value otherwise.

### GetDesktopPoolsAndFarmsCountOk

`func (o *VirtualCenterMonitorInfoV2) GetDesktopPoolsAndFarmsCountOk() (*int32, bool)`

GetDesktopPoolsAndFarmsCountOk returns a tuple with the DesktopPoolsAndFarmsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolsAndFarmsCount

`func (o *VirtualCenterMonitorInfoV2) SetDesktopPoolsAndFarmsCount(v int32)`

SetDesktopPoolsAndFarmsCount sets DesktopPoolsAndFarmsCount field to given value.


### GetDetails

`func (o *VirtualCenterMonitorInfoV2) GetDetails() VCMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VirtualCenterMonitorInfoV2) GetDetailsOk() (*VCMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VirtualCenterMonitorInfoV2) SetDetails(v VCMonitorDetails)`

SetDetails sets Details field to given value.


### GetHosts

`func (o *VirtualCenterMonitorInfoV2) GetHosts() []VCMonitorHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualCenterMonitorInfoV2) GetHostsOk() (*[]VCMonitorHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualCenterMonitorInfoV2) SetHosts(v []VCMonitorHost)`

SetHosts sets Hosts field to given value.


### GetId

`func (o *VirtualCenterMonitorInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCenterMonitorInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCenterMonitorInfoV2) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VirtualCenterMonitorInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCenterMonitorInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCenterMonitorInfoV2) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


