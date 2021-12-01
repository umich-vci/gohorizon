# TrueSSOMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomainDnsName** | Pointer to **string** | The DNS name of the domain. | [optional] 
**AdDomainId** | Pointer to **string** | Unique SID of the AD Domain. | [optional] 
**AdDomainStatus** | Pointer to **string** | The state of the domain health, taken as the most severe reported by one of the enrollment servers. * OK: The state of the domain is OK as reported by the enrollment servers. * WARN: At least one of the enrollment servers has a warning. * ERROR: At least one of the enrollment servers is in an error state. | [optional] 
**CertificateServerDetails** | Pointer to [**[]CertificateServerDetails**](CertificateServerDetails.md) | Details of each certificate server. | [optional] 
**Enabled** | Pointer to **bool** | Indicates if the True SSO connector is enabled. | [optional] 
**Id** | Pointer to **string** | Unique ID of the True SSO Connector. | [optional] 
**Name** | Pointer to **string** | True SSO connector name. | [optional] 
**PrimaryEnrollmentServer** | Pointer to [**EnrollmentServerDetails**](EnrollmentServerDetails.md) |  | [optional] 
**SecondaryEnrollmentServer** | Pointer to [**EnrollmentServerDetails**](EnrollmentServerDetails.md) |  | [optional] 
**Status** | Pointer to **string** | Overall status of the True SSO connector. * OK: All the components of the True SSO connector are fine. * WARN: At least one component of the True SSO connector has a warning. * ERROR: At least one component of the True SSO connector has an error. | [optional] 
**TemplateName** | Pointer to **string** | Unique name for the True SSO template. | [optional] 
**TemplateStatus** | Pointer to **string** | The state of the template health, taken as the most severe reported by one of the enrollment servers. * OK: The state of the template is OK as reported by the enrollment servers. * WARN: At least one enrollment server reports a warning on this template. * ERROR: At least one enrollment server reports an error on this template. | [optional] 

## Methods

### NewTrueSSOMonitorInfo

`func NewTrueSSOMonitorInfo() *TrueSSOMonitorInfo`

NewTrueSSOMonitorInfo instantiates a new TrueSSOMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrueSSOMonitorInfoWithDefaults

`func NewTrueSSOMonitorInfoWithDefaults() *TrueSSOMonitorInfo`

NewTrueSSOMonitorInfoWithDefaults instantiates a new TrueSSOMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomainDnsName

`func (o *TrueSSOMonitorInfo) GetAdDomainDnsName() string`

GetAdDomainDnsName returns the AdDomainDnsName field if non-nil, zero value otherwise.

### GetAdDomainDnsNameOk

`func (o *TrueSSOMonitorInfo) GetAdDomainDnsNameOk() (*string, bool)`

GetAdDomainDnsNameOk returns a tuple with the AdDomainDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainDnsName

`func (o *TrueSSOMonitorInfo) SetAdDomainDnsName(v string)`

SetAdDomainDnsName sets AdDomainDnsName field to given value.

### HasAdDomainDnsName

`func (o *TrueSSOMonitorInfo) HasAdDomainDnsName() bool`

HasAdDomainDnsName returns a boolean if a field has been set.

### GetAdDomainId

`func (o *TrueSSOMonitorInfo) GetAdDomainId() string`

GetAdDomainId returns the AdDomainId field if non-nil, zero value otherwise.

### GetAdDomainIdOk

`func (o *TrueSSOMonitorInfo) GetAdDomainIdOk() (*string, bool)`

GetAdDomainIdOk returns a tuple with the AdDomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainId

`func (o *TrueSSOMonitorInfo) SetAdDomainId(v string)`

SetAdDomainId sets AdDomainId field to given value.

### HasAdDomainId

`func (o *TrueSSOMonitorInfo) HasAdDomainId() bool`

HasAdDomainId returns a boolean if a field has been set.

### GetAdDomainStatus

`func (o *TrueSSOMonitorInfo) GetAdDomainStatus() string`

GetAdDomainStatus returns the AdDomainStatus field if non-nil, zero value otherwise.

### GetAdDomainStatusOk

`func (o *TrueSSOMonitorInfo) GetAdDomainStatusOk() (*string, bool)`

GetAdDomainStatusOk returns a tuple with the AdDomainStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomainStatus

`func (o *TrueSSOMonitorInfo) SetAdDomainStatus(v string)`

SetAdDomainStatus sets AdDomainStatus field to given value.

### HasAdDomainStatus

`func (o *TrueSSOMonitorInfo) HasAdDomainStatus() bool`

HasAdDomainStatus returns a boolean if a field has been set.

### GetCertificateServerDetails

`func (o *TrueSSOMonitorInfo) GetCertificateServerDetails() []CertificateServerDetails`

GetCertificateServerDetails returns the CertificateServerDetails field if non-nil, zero value otherwise.

### GetCertificateServerDetailsOk

`func (o *TrueSSOMonitorInfo) GetCertificateServerDetailsOk() (*[]CertificateServerDetails, bool)`

GetCertificateServerDetailsOk returns a tuple with the CertificateServerDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateServerDetails

`func (o *TrueSSOMonitorInfo) SetCertificateServerDetails(v []CertificateServerDetails)`

SetCertificateServerDetails sets CertificateServerDetails field to given value.

### HasCertificateServerDetails

`func (o *TrueSSOMonitorInfo) HasCertificateServerDetails() bool`

HasCertificateServerDetails returns a boolean if a field has been set.

### GetEnabled

`func (o *TrueSSOMonitorInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *TrueSSOMonitorInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *TrueSSOMonitorInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *TrueSSOMonitorInfo) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *TrueSSOMonitorInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrueSSOMonitorInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrueSSOMonitorInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrueSSOMonitorInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TrueSSOMonitorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TrueSSOMonitorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TrueSSOMonitorInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TrueSSOMonitorInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimaryEnrollmentServer

`func (o *TrueSSOMonitorInfo) GetPrimaryEnrollmentServer() EnrollmentServerDetails`

GetPrimaryEnrollmentServer returns the PrimaryEnrollmentServer field if non-nil, zero value otherwise.

### GetPrimaryEnrollmentServerOk

`func (o *TrueSSOMonitorInfo) GetPrimaryEnrollmentServerOk() (*EnrollmentServerDetails, bool)`

GetPrimaryEnrollmentServerOk returns a tuple with the PrimaryEnrollmentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEnrollmentServer

`func (o *TrueSSOMonitorInfo) SetPrimaryEnrollmentServer(v EnrollmentServerDetails)`

SetPrimaryEnrollmentServer sets PrimaryEnrollmentServer field to given value.

### HasPrimaryEnrollmentServer

`func (o *TrueSSOMonitorInfo) HasPrimaryEnrollmentServer() bool`

HasPrimaryEnrollmentServer returns a boolean if a field has been set.

### GetSecondaryEnrollmentServer

`func (o *TrueSSOMonitorInfo) GetSecondaryEnrollmentServer() EnrollmentServerDetails`

GetSecondaryEnrollmentServer returns the SecondaryEnrollmentServer field if non-nil, zero value otherwise.

### GetSecondaryEnrollmentServerOk

`func (o *TrueSSOMonitorInfo) GetSecondaryEnrollmentServerOk() (*EnrollmentServerDetails, bool)`

GetSecondaryEnrollmentServerOk returns a tuple with the SecondaryEnrollmentServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEnrollmentServer

`func (o *TrueSSOMonitorInfo) SetSecondaryEnrollmentServer(v EnrollmentServerDetails)`

SetSecondaryEnrollmentServer sets SecondaryEnrollmentServer field to given value.

### HasSecondaryEnrollmentServer

`func (o *TrueSSOMonitorInfo) HasSecondaryEnrollmentServer() bool`

HasSecondaryEnrollmentServer returns a boolean if a field has been set.

### GetStatus

`func (o *TrueSSOMonitorInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TrueSSOMonitorInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TrueSSOMonitorInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TrueSSOMonitorInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTemplateName

`func (o *TrueSSOMonitorInfo) GetTemplateName() string`

GetTemplateName returns the TemplateName field if non-nil, zero value otherwise.

### GetTemplateNameOk

`func (o *TrueSSOMonitorInfo) GetTemplateNameOk() (*string, bool)`

GetTemplateNameOk returns a tuple with the TemplateName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateName

`func (o *TrueSSOMonitorInfo) SetTemplateName(v string)`

SetTemplateName sets TemplateName field to given value.

### HasTemplateName

`func (o *TrueSSOMonitorInfo) HasTemplateName() bool`

HasTemplateName returns a boolean if a field has been set.

### GetTemplateStatus

`func (o *TrueSSOMonitorInfo) GetTemplateStatus() string`

GetTemplateStatus returns the TemplateStatus field if non-nil, zero value otherwise.

### GetTemplateStatusOk

`func (o *TrueSSOMonitorInfo) GetTemplateStatusOk() (*string, bool)`

GetTemplateStatusOk returns a tuple with the TemplateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateStatus

`func (o *TrueSSOMonitorInfo) SetTemplateStatus(v string)`

SetTemplateStatus sets TemplateStatus field to given value.

### HasTemplateStatus

`func (o *TrueSSOMonitorInfo) HasTemplateStatus() bool`

HasTemplateStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


