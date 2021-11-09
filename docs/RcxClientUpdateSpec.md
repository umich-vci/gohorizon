# RCXClientUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | IP address of the RCX client. | [optional] 
**Thumbprints** | [**[]CertificateThumbprint**](CertificateThumbprint.md) | Thumbprints of the RCX client certificate. | 

## Methods

### NewRCXClientUpdateSpec

`func NewRCXClientUpdateSpec(thumbprints []CertificateThumbprint, ) *RCXClientUpdateSpec`

NewRCXClientUpdateSpec instantiates a new RCXClientUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCXClientUpdateSpecWithDefaults

`func NewRCXClientUpdateSpecWithDefaults() *RCXClientUpdateSpec`

NewRCXClientUpdateSpecWithDefaults instantiates a new RCXClientUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *RCXClientUpdateSpec) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RCXClientUpdateSpec) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RCXClientUpdateSpec) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *RCXClientUpdateSpec) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetThumbprints

`func (o *RCXClientUpdateSpec) GetThumbprints() []CertificateThumbprint`

GetThumbprints returns the Thumbprints field if non-nil, zero value otherwise.

### GetThumbprintsOk

`func (o *RCXClientUpdateSpec) GetThumbprintsOk() (*[]CertificateThumbprint, bool)`

GetThumbprintsOk returns a tuple with the Thumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprints

`func (o *RCXClientUpdateSpec) SetThumbprints(v []CertificateThumbprint)`

SetThumbprints sets Thumbprints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


