# CertificateOverrideData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** | Virtual Center certificate | [optional] 
**Type** | Pointer to **string** | Type of Certificate. * PEM: PEM encoded certificate type | [optional] 

## Methods

### NewCertificateOverrideData

`func NewCertificateOverrideData() *CertificateOverrideData`

NewCertificateOverrideData instantiates a new CertificateOverrideData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateOverrideDataWithDefaults

`func NewCertificateOverrideDataWithDefaults() *CertificateOverrideData`

NewCertificateOverrideDataWithDefaults instantiates a new CertificateOverrideData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CertificateOverrideData) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CertificateOverrideData) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CertificateOverrideData) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CertificateOverrideData) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetType

`func (o *CertificateOverrideData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CertificateOverrideData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CertificateOverrideData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CertificateOverrideData) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


