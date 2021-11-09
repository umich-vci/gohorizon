# CertificateThumbprint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SslCertThumbprint** | **string** | A digest of the certificate. | 
**SslCertThumbprintAlgorithm** | **string** | Algorithm used to compute the thumbprint. * SHA_1: SHA-1 hashing algorithm. * SHA_256: SHA-256 hashing algorithm. | 

## Methods

### NewCertificateThumbprint

`func NewCertificateThumbprint(sslCertThumbprint string, sslCertThumbprintAlgorithm string, ) *CertificateThumbprint`

NewCertificateThumbprint instantiates a new CertificateThumbprint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateThumbprintWithDefaults

`func NewCertificateThumbprintWithDefaults() *CertificateThumbprint`

NewCertificateThumbprintWithDefaults instantiates a new CertificateThumbprint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSslCertThumbprint

`func (o *CertificateThumbprint) GetSslCertThumbprint() string`

GetSslCertThumbprint returns the SslCertThumbprint field if non-nil, zero value otherwise.

### GetSslCertThumbprintOk

`func (o *CertificateThumbprint) GetSslCertThumbprintOk() (*string, bool)`

GetSslCertThumbprintOk returns a tuple with the SslCertThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertThumbprint

`func (o *CertificateThumbprint) SetSslCertThumbprint(v string)`

SetSslCertThumbprint sets SslCertThumbprint field to given value.


### GetSslCertThumbprintAlgorithm

`func (o *CertificateThumbprint) GetSslCertThumbprintAlgorithm() string`

GetSslCertThumbprintAlgorithm returns the SslCertThumbprintAlgorithm field if non-nil, zero value otherwise.

### GetSslCertThumbprintAlgorithmOk

`func (o *CertificateThumbprint) GetSslCertThumbprintAlgorithmOk() (*string, bool)`

GetSslCertThumbprintAlgorithmOk returns a tuple with the SslCertThumbprintAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCertThumbprintAlgorithm

`func (o *CertificateThumbprint) SetSslCertThumbprintAlgorithm(v string)`

SetSslCertThumbprintAlgorithm sets SslCertThumbprintAlgorithm field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


