# ADDomainMonitorConnectionServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID of the connection server. | 
**Name** | **string** | Connection server host name or IP address. | 
**Status** | **string** | Status of the connection to the domain. * UNCONTACTABLE: No domain controllers appear to be present on the network for this domain. * FULLY_ACCESSIBLE: The domain controller(s) are accepting bind operations. * CANNOT_BIND: The domain controller(s) are only accepting LDAP ping operations. * UNKNOWN: Cannot determine accessibility. | 
**TrustRelationship** | **string** | The trust relationship for the domain. * PRIMARY_DOMAIN: The domain is the domain that the broker is present in. * FROM_BROKER: The domain is trusted by the domain that the broker is in. * TO_BROKER: The domain trusts the broker&#39;s domain (this is for completeness and generally will not be used). * TWO_WAY: The domain has a two way trust relationship with the broker&#39;s domain. * TWO_WAY_FOREST: The domain is in the same forest as the broker&#39;s domain, implies two way trust. * MANUAL: The domain was manually configured (the trust has not been detected). * NOTRUST: The domain not having trust with broker domain. * UNKNOWN: The trust relationship could not be determined. | 

## Methods

### NewADDomainMonitorConnectionServer

`func NewADDomainMonitorConnectionServer(id string, name string, status string, trustRelationship string, ) *ADDomainMonitorConnectionServer`

NewADDomainMonitorConnectionServer instantiates a new ADDomainMonitorConnectionServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainMonitorConnectionServerWithDefaults

`func NewADDomainMonitorConnectionServerWithDefaults() *ADDomainMonitorConnectionServer`

NewADDomainMonitorConnectionServerWithDefaults instantiates a new ADDomainMonitorConnectionServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ADDomainMonitorConnectionServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ADDomainMonitorConnectionServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ADDomainMonitorConnectionServer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ADDomainMonitorConnectionServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ADDomainMonitorConnectionServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ADDomainMonitorConnectionServer) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ADDomainMonitorConnectionServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ADDomainMonitorConnectionServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ADDomainMonitorConnectionServer) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTrustRelationship

`func (o *ADDomainMonitorConnectionServer) GetTrustRelationship() string`

GetTrustRelationship returns the TrustRelationship field if non-nil, zero value otherwise.

### GetTrustRelationshipOk

`func (o *ADDomainMonitorConnectionServer) GetTrustRelationshipOk() (*string, bool)`

GetTrustRelationshipOk returns a tuple with the TrustRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustRelationship

`func (o *ADDomainMonitorConnectionServer) SetTrustRelationship(v string)`

SetTrustRelationship sets TrustRelationship field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


