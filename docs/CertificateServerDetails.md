# CertificateServerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Certificate Server name. | 
**Status** | **string** | Certificate Server status. * OK: The state of the certificate server is OK as reported by the enrollment servers. * WARN: At least one enrollment server reports a warning on this certificate server. * ERROR: At least one enrollment server reports an error on this certificate server. | 

## Methods

### NewCertificateServerDetails

`func NewCertificateServerDetails(name string, status string, ) *CertificateServerDetails`

NewCertificateServerDetails instantiates a new CertificateServerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateServerDetailsWithDefaults

`func NewCertificateServerDetailsWithDefaults() *CertificateServerDetails`

NewCertificateServerDetailsWithDefaults instantiates a new CertificateServerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateServerDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateServerDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateServerDetails) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *CertificateServerDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CertificateServerDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CertificateServerDetails) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


