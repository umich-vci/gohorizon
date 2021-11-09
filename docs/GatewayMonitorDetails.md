# GatewayMonitorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **string** | Gateway host name or IP address. | 
**Internal** | **bool** | Indicates if the Gateway is internal. | 
**Type** | **string** | Type of the Gateway. * UAG: Unified Access Gateway type. * F5: F5 Gateway type. * UNKNOWN: Unknown type. | 
**Version** | **string** | Version of the Gateway. | 

## Methods

### NewGatewayMonitorDetails

`func NewGatewayMonitorDetails(address string, internal bool, type_ string, version string, ) *GatewayMonitorDetails`

NewGatewayMonitorDetails instantiates a new GatewayMonitorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayMonitorDetailsWithDefaults

`func NewGatewayMonitorDetailsWithDefaults() *GatewayMonitorDetails`

NewGatewayMonitorDetailsWithDefaults instantiates a new GatewayMonitorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *GatewayMonitorDetails) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GatewayMonitorDetails) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GatewayMonitorDetails) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetInternal

`func (o *GatewayMonitorDetails) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *GatewayMonitorDetails) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *GatewayMonitorDetails) SetInternal(v bool)`

SetInternal sets Internal field to given value.


### GetType

`func (o *GatewayMonitorDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayMonitorDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayMonitorDetails) SetType(v string)`

SetType sets Type field to given value.


### GetVersion

`func (o *GatewayMonitorDetails) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GatewayMonitorDetails) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GatewayMonitorDetails) SetVersion(v string)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


