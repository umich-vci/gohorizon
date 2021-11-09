# CertificateMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** | Indicates if the certificate is valid. | 
**ValidFrom** | **int64** | Start time of the certificate validity in milliseconds. Measured as epoch time. | 
**ValidTo** | **int64** | Expiration time of the certificate validity in milliseconds. Measured as epoch time. | 

## Methods

### NewCertificateMonitorInfo

`func NewCertificateMonitorInfo(valid bool, validFrom int64, validTo int64, ) *CertificateMonitorInfo`

NewCertificateMonitorInfo instantiates a new CertificateMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateMonitorInfoWithDefaults

`func NewCertificateMonitorInfoWithDefaults() *CertificateMonitorInfo`

NewCertificateMonitorInfoWithDefaults instantiates a new CertificateMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *CertificateMonitorInfo) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CertificateMonitorInfo) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CertificateMonitorInfo) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetValidFrom

`func (o *CertificateMonitorInfo) GetValidFrom() int64`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *CertificateMonitorInfo) GetValidFromOk() (*int64, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *CertificateMonitorInfo) SetValidFrom(v int64)`

SetValidFrom sets ValidFrom field to given value.


### GetValidTo

`func (o *CertificateMonitorInfo) GetValidTo() int64`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *CertificateMonitorInfo) GetValidToOk() (*int64, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *CertificateMonitorInfo) SetValidTo(v int64)`

SetValidTo sets ValidTo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


