# DatacenterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID representing a datacenter. | [optional] 
**Name** | Pointer to **string** | Name of the datacenter. | [optional] 
**Path** | Pointer to **string** | Datacenter path. | [optional] 

## Methods

### NewDatacenterInfo

`func NewDatacenterInfo() *DatacenterInfo`

NewDatacenterInfo instantiates a new DatacenterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatacenterInfoWithDefaults

`func NewDatacenterInfoWithDefaults() *DatacenterInfo`

NewDatacenterInfoWithDefaults instantiates a new DatacenterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DatacenterInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DatacenterInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DatacenterInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DatacenterInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DatacenterInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatacenterInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatacenterInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DatacenterInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPath

`func (o *DatacenterInfo) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DatacenterInfo) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DatacenterInfo) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DatacenterInfo) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


