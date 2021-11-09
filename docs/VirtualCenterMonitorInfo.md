# VirtualCenterMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionServers** | [**[]VCMonitorConnectionServer**](VCMonitorConnectionServer.md) | Information about the Virtual Center connections from each of the connection servers. | 
**Datastores** | [**[]VCMonitorDatastore**](VCMonitorDatastore.md) | Information about the datastores of the Virtual Center. | 
**DesktopsCount** | **int32** | Number of Desktop Pools And Farms managed by the virtual center. | 
**Details** | [**VCMonitorDetails**](VCMonitorDetails.md) |  | 
**Hosts** | [**[]VCMonitorHost**](VCMonitorHost.md) | Information about the ESX hosts added in the Virtual Center. | 
**Id** | **string** | Unique ID of the Virtual Center. | 
**Name** | **string** | Virtual Center server name or IP address. | 

## Methods

### NewVirtualCenterMonitorInfo

`func NewVirtualCenterMonitorInfo(connectionServers []VCMonitorConnectionServer, datastores []VCMonitorDatastore, desktopsCount int32, details VCMonitorDetails, hosts []VCMonitorHost, id string, name string, ) *VirtualCenterMonitorInfo`

NewVirtualCenterMonitorInfo instantiates a new VirtualCenterMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCenterMonitorInfoWithDefaults

`func NewVirtualCenterMonitorInfoWithDefaults() *VirtualCenterMonitorInfo`

NewVirtualCenterMonitorInfoWithDefaults instantiates a new VirtualCenterMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionServers

`func (o *VirtualCenterMonitorInfo) GetConnectionServers() []VCMonitorConnectionServer`

GetConnectionServers returns the ConnectionServers field if non-nil, zero value otherwise.

### GetConnectionServersOk

`func (o *VirtualCenterMonitorInfo) GetConnectionServersOk() (*[]VCMonitorConnectionServer, bool)`

GetConnectionServersOk returns a tuple with the ConnectionServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionServers

`func (o *VirtualCenterMonitorInfo) SetConnectionServers(v []VCMonitorConnectionServer)`

SetConnectionServers sets ConnectionServers field to given value.


### GetDatastores

`func (o *VirtualCenterMonitorInfo) GetDatastores() []VCMonitorDatastore`

GetDatastores returns the Datastores field if non-nil, zero value otherwise.

### GetDatastoresOk

`func (o *VirtualCenterMonitorInfo) GetDatastoresOk() (*[]VCMonitorDatastore, bool)`

GetDatastoresOk returns a tuple with the Datastores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastores

`func (o *VirtualCenterMonitorInfo) SetDatastores(v []VCMonitorDatastore)`

SetDatastores sets Datastores field to given value.


### GetDesktopsCount

`func (o *VirtualCenterMonitorInfo) GetDesktopsCount() int32`

GetDesktopsCount returns the DesktopsCount field if non-nil, zero value otherwise.

### GetDesktopsCountOk

`func (o *VirtualCenterMonitorInfo) GetDesktopsCountOk() (*int32, bool)`

GetDesktopsCountOk returns a tuple with the DesktopsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopsCount

`func (o *VirtualCenterMonitorInfo) SetDesktopsCount(v int32)`

SetDesktopsCount sets DesktopsCount field to given value.


### GetDetails

`func (o *VirtualCenterMonitorInfo) GetDetails() VCMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VirtualCenterMonitorInfo) GetDetailsOk() (*VCMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VirtualCenterMonitorInfo) SetDetails(v VCMonitorDetails)`

SetDetails sets Details field to given value.


### GetHosts

`func (o *VirtualCenterMonitorInfo) GetHosts() []VCMonitorHost`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VirtualCenterMonitorInfo) GetHostsOk() (*[]VCMonitorHost, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VirtualCenterMonitorInfo) SetHosts(v []VCMonitorHost)`

SetHosts sets Hosts field to given value.


### GetId

`func (o *VirtualCenterMonitorInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCenterMonitorInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCenterMonitorInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VirtualCenterMonitorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCenterMonitorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCenterMonitorInfo) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


