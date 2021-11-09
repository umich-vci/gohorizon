# RCXClientRegisterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpAddress** | Pointer to **string** | IP address of the RCX client. | [optional] 
**Name** | **string** | The RCX client certificate name. | 
**Thumbprints** | [**[]CertificateThumbprint**](CertificateThumbprint.md) | Thumbprints of the RCX client certificate. | 

## Methods

### NewRCXClientRegisterSpec

`func NewRCXClientRegisterSpec(name string, thumbprints []CertificateThumbprint, ) *RCXClientRegisterSpec`

NewRCXClientRegisterSpec instantiates a new RCXClientRegisterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCXClientRegisterSpecWithDefaults

`func NewRCXClientRegisterSpecWithDefaults() *RCXClientRegisterSpec`

NewRCXClientRegisterSpecWithDefaults instantiates a new RCXClientRegisterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpAddress

`func (o *RCXClientRegisterSpec) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *RCXClientRegisterSpec) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *RCXClientRegisterSpec) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *RCXClientRegisterSpec) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetName

`func (o *RCXClientRegisterSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RCXClientRegisterSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RCXClientRegisterSpec) SetName(v string)`

SetName sets Name field to given value.


### GetThumbprints

`func (o *RCXClientRegisterSpec) GetThumbprints() []CertificateThumbprint`

GetThumbprints returns the Thumbprints field if non-nil, zero value otherwise.

### GetThumbprintsOk

`func (o *RCXClientRegisterSpec) GetThumbprintsOk() (*[]CertificateThumbprint, bool)`

GetThumbprintsOk returns a tuple with the Thumbprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprints

`func (o *RCXClientRegisterSpec) SetThumbprints(v []CertificateThumbprint)`

SetThumbprints sets Thumbprints field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


