# DesktopPoolDatastoreSettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreId** | **string** | Id of the datastore. | 
**SdrsCluster** | Pointer to **bool** | Indicates whether datastore represents a Storage DRS cluster. | [optional] 

## Methods

### NewDesktopPoolDatastoreSettingsUpdateSpec

`func NewDesktopPoolDatastoreSettingsUpdateSpec(datastoreId string, ) *DesktopPoolDatastoreSettingsUpdateSpec`

NewDesktopPoolDatastoreSettingsUpdateSpec instantiates a new DesktopPoolDatastoreSettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolDatastoreSettingsUpdateSpecWithDefaults

`func NewDesktopPoolDatastoreSettingsUpdateSpecWithDefaults() *DesktopPoolDatastoreSettingsUpdateSpec`

NewDesktopPoolDatastoreSettingsUpdateSpecWithDefaults instantiates a new DesktopPoolDatastoreSettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreId

`func (o *DesktopPoolDatastoreSettingsUpdateSpec) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *DesktopPoolDatastoreSettingsUpdateSpec) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *DesktopPoolDatastoreSettingsUpdateSpec) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.


### GetSdrsCluster

`func (o *DesktopPoolDatastoreSettingsUpdateSpec) GetSdrsCluster() bool`

GetSdrsCluster returns the SdrsCluster field if non-nil, zero value otherwise.

### GetSdrsClusterOk

`func (o *DesktopPoolDatastoreSettingsUpdateSpec) GetSdrsClusterOk() (*bool, bool)`

GetSdrsClusterOk returns a tuple with the SdrsCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdrsCluster

`func (o *DesktopPoolDatastoreSettingsUpdateSpec) SetSdrsCluster(v bool)`

SetSdrsCluster sets SdrsCluster field to given value.

### HasSdrsCluster

`func (o *DesktopPoolDatastoreSettingsUpdateSpec) HasSdrsCluster() bool`

HasSdrsCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


