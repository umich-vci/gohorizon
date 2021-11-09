# SecurityGatewayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | IP address of the security gateway.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**DomainName** | Pointer to **string** | Computer machine name or DNS name of the security gateway.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. | [optional] 
**Location** | Pointer to **string** | Location of the security gateway.  This property need not be set.&lt;br&gt;Supported Filters : &#39;Equals&#39;. * EXTERNAL: Gateway location is external. * INTERNAL: Gateway location is internal. * UNKNOWN: Gateway location is unknown. | [optional] 

## Methods

### NewSecurityGatewayData

`func NewSecurityGatewayData() *SecurityGatewayData`

NewSecurityGatewayData instantiates a new SecurityGatewayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityGatewayDataWithDefaults

`func NewSecurityGatewayDataWithDefaults() *SecurityGatewayData`

NewSecurityGatewayDataWithDefaults instantiates a new SecurityGatewayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *SecurityGatewayData) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *SecurityGatewayData) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *SecurityGatewayData) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *SecurityGatewayData) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDomainName

`func (o *SecurityGatewayData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *SecurityGatewayData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *SecurityGatewayData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *SecurityGatewayData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetLocation

`func (o *SecurityGatewayData) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SecurityGatewayData) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SecurityGatewayData) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SecurityGatewayData) HasLocation() bool`

HasLocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


