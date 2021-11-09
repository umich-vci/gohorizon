# NetworkInterfaceCardInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID representing the network interface card. | [optional] 
**MacAddress** | Pointer to **string** | Network interface card MAC address. | [optional] 
**Name** | Pointer to **string** | Network interface card name. | [optional] 

## Methods

### NewNetworkInterfaceCardInfo

`func NewNetworkInterfaceCardInfo() *NetworkInterfaceCardInfo`

NewNetworkInterfaceCardInfo instantiates a new NetworkInterfaceCardInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkInterfaceCardInfoWithDefaults

`func NewNetworkInterfaceCardInfoWithDefaults() *NetworkInterfaceCardInfo`

NewNetworkInterfaceCardInfoWithDefaults instantiates a new NetworkInterfaceCardInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkInterfaceCardInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkInterfaceCardInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkInterfaceCardInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkInterfaceCardInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMacAddress

`func (o *NetworkInterfaceCardInfo) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *NetworkInterfaceCardInfo) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *NetworkInterfaceCardInfo) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *NetworkInterfaceCardInfo) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetName

`func (o *NetworkInterfaceCardInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkInterfaceCardInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkInterfaceCardInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkInterfaceCardInfo) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


