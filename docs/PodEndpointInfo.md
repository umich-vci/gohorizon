# PodEndpointInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether an endpoint is enabled. A disabled endpoint is excluded from participating in inter-pod communication. | [optional] 
**Id** | Pointer to **string** | Unique ID representing this pod endpoint. | [optional] 
**Name** | Pointer to **string** | Name of the pod endpoint. | [optional] 
**ServerAddress** | Pointer to **string** | The URL for the pod endpoint. This address and port is used for inter-pod communication. | [optional] 

## Methods

### NewPodEndpointInfo

`func NewPodEndpointInfo() *PodEndpointInfo`

NewPodEndpointInfo instantiates a new PodEndpointInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPodEndpointInfoWithDefaults

`func NewPodEndpointInfoWithDefaults() *PodEndpointInfo`

NewPodEndpointInfoWithDefaults instantiates a new PodEndpointInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PodEndpointInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PodEndpointInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PodEndpointInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PodEndpointInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *PodEndpointInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PodEndpointInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PodEndpointInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PodEndpointInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PodEndpointInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PodEndpointInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PodEndpointInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PodEndpointInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServerAddress

`func (o *PodEndpointInfo) GetServerAddress() string`

GetServerAddress returns the ServerAddress field if non-nil, zero value otherwise.

### GetServerAddressOk

`func (o *PodEndpointInfo) GetServerAddressOk() (*string, bool)`

GetServerAddressOk returns a tuple with the ServerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddress

`func (o *PodEndpointInfo) SetServerAddress(v string)`

SetServerAddress sets ServerAddress field to given value.

### HasServerAddress

`func (o *PodEndpointInfo) HasServerAddress() bool`

HasServerAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


