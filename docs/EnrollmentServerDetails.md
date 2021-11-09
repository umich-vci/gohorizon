# EnrollmentServerDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DnsName** | **string** | Enrollment server dns name. | 
**Id** | **string** | Unique ID of the Enrollment Server. | 
**Status** | **string** | Enrollment server status. * OK: The state of enrollment server is OK. * ERROR: The enrollment server has an error. | 

## Methods

### NewEnrollmentServerDetails

`func NewEnrollmentServerDetails(dnsName string, id string, status string, ) *EnrollmentServerDetails`

NewEnrollmentServerDetails instantiates a new EnrollmentServerDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentServerDetailsWithDefaults

`func NewEnrollmentServerDetailsWithDefaults() *EnrollmentServerDetails`

NewEnrollmentServerDetailsWithDefaults instantiates a new EnrollmentServerDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnsName

`func (o *EnrollmentServerDetails) GetDnsName() string`

GetDnsName returns the DnsName field if non-nil, zero value otherwise.

### GetDnsNameOk

`func (o *EnrollmentServerDetails) GetDnsNameOk() (*string, bool)`

GetDnsNameOk returns a tuple with the DnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsName

`func (o *EnrollmentServerDetails) SetDnsName(v string)`

SetDnsName sets DnsName field to given value.


### GetId

`func (o *EnrollmentServerDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EnrollmentServerDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EnrollmentServerDetails) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *EnrollmentServerDetails) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EnrollmentServerDetails) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EnrollmentServerDetails) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


