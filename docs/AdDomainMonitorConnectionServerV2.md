# ADDomainMonitorConnectionServerV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the connection server. | [optional] 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp in milliseconds when the last update was obtained. Measured as epoch time. | [optional] 
**Name** | Pointer to **string** | Connection server host name or IP address. | [optional] 
**Status** | Pointer to **string** | Status of the connection to the domain. * UNCONTACTABLE: No domain controllers appear to be present on the network for this domain. * FULLY_ACCESSIBLE: The domain controller(s) are accepting bind operations. * CANNOT_BIND: The domain controller(s) are only accepting LDAP ping operations. * UNKNOWN: Cannot determine accessibility. | [optional] 
**TrustRelationship** | Pointer to **string** | The trust relationship for the domain. * PRIMARY_DOMAIN: The domain is the domain that the broker is present in. * FROM_BROKER: The domain is trusted by the domain that the broker is in. * TO_BROKER: The domain trusts the broker&#39;s domain (this is for completeness and generally will not be used). * TWO_WAY: The domain has a two way trust relationship with the broker&#39;s domain. * TWO_WAY_FOREST: The domain is in the same forest as the broker&#39;s domain, implies two way trust. * MANUAL: The domain was manually configured (the trust has not been detected). * NOTRUST: The domain not having trust with broker domain. * UNKNOWN: The trust relationship could not be determined. | [optional] 

## Methods

### NewADDomainMonitorConnectionServerV2

`func NewADDomainMonitorConnectionServerV2() *ADDomainMonitorConnectionServerV2`

NewADDomainMonitorConnectionServerV2 instantiates a new ADDomainMonitorConnectionServerV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewADDomainMonitorConnectionServerV2WithDefaults

`func NewADDomainMonitorConnectionServerV2WithDefaults() *ADDomainMonitorConnectionServerV2`

NewADDomainMonitorConnectionServerV2WithDefaults instantiates a new ADDomainMonitorConnectionServerV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ADDomainMonitorConnectionServerV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ADDomainMonitorConnectionServerV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ADDomainMonitorConnectionServerV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ADDomainMonitorConnectionServerV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *ADDomainMonitorConnectionServerV2) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *ADDomainMonitorConnectionServerV2) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *ADDomainMonitorConnectionServerV2) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *ADDomainMonitorConnectionServerV2) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ADDomainMonitorConnectionServerV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ADDomainMonitorConnectionServerV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ADDomainMonitorConnectionServerV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ADDomainMonitorConnectionServerV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ADDomainMonitorConnectionServerV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ADDomainMonitorConnectionServerV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ADDomainMonitorConnectionServerV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ADDomainMonitorConnectionServerV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTrustRelationship

`func (o *ADDomainMonitorConnectionServerV2) GetTrustRelationship() string`

GetTrustRelationship returns the TrustRelationship field if non-nil, zero value otherwise.

### GetTrustRelationshipOk

`func (o *ADDomainMonitorConnectionServerV2) GetTrustRelationshipOk() (*string, bool)`

GetTrustRelationshipOk returns a tuple with the TrustRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustRelationship

`func (o *ADDomainMonitorConnectionServerV2) SetTrustRelationship(v string)`

SetTrustRelationship sets TrustRelationship field to given value.

### HasTrustRelationship

`func (o *ADDomainMonitorConnectionServerV2) HasTrustRelationship() bool`

HasTrustRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


