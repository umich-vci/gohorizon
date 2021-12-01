# DesktopPoolDatastoreSettingsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatastoreId** | **string** | Id of the datastore. | 
**SdrsCluster** | Pointer to **bool** | Applicable to full clone desktop pools with default value as false. Indicates whether datastore represents a Storage DRS cluster. | [optional] 

## Methods

### NewDesktopPoolDatastoreSettingsCreateSpec

`func NewDesktopPoolDatastoreSettingsCreateSpec(datastoreId string, ) *DesktopPoolDatastoreSettingsCreateSpec`

NewDesktopPoolDatastoreSettingsCreateSpec instantiates a new DesktopPoolDatastoreSettingsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDesktopPoolDatastoreSettingsCreateSpecWithDefaults

`func NewDesktopPoolDatastoreSettingsCreateSpecWithDefaults() *DesktopPoolDatastoreSettingsCreateSpec`

NewDesktopPoolDatastoreSettingsCreateSpecWithDefaults instantiates a new DesktopPoolDatastoreSettingsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastoreId

`func (o *DesktopPoolDatastoreSettingsCreateSpec) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *DesktopPoolDatastoreSettingsCreateSpec) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *DesktopPoolDatastoreSettingsCreateSpec) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.


### GetSdrsCluster

`func (o *DesktopPoolDatastoreSettingsCreateSpec) GetSdrsCluster() bool`

GetSdrsCluster returns the SdrsCluster field if non-nil, zero value otherwise.

### GetSdrsClusterOk

`func (o *DesktopPoolDatastoreSettingsCreateSpec) GetSdrsClusterOk() (*bool, bool)`

GetSdrsClusterOk returns a tuple with the SdrsCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdrsCluster

`func (o *DesktopPoolDatastoreSettingsCreateSpec) SetSdrsCluster(v bool)`

SetSdrsCluster sets SdrsCluster field to given value.

### HasSdrsCluster

`func (o *DesktopPoolDatastoreSettingsCreateSpec) HasSdrsCluster() bool`

HasSdrsCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


