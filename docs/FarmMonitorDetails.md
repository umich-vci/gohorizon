# FarmMonitorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | Pointer to **string** | Source of farm machines. This property is required if type is set to \&quot;AUTOMATED\&quot;. * LINKED_CLONE: Linked clone share the same base image and use less storage space than full RDS Servers. * INSTANT_CLONE: Instant clone engine uses vmfork technology to create the instant clones. These clones take much less time for provisioning. | [optional] 
**Type** | **string** | Farm type. * AUTOMATED: Automated Farm. * MANUAL: Manual farm. | 

## Methods

### NewFarmMonitorDetails

`func NewFarmMonitorDetails(type_ string, ) *FarmMonitorDetails`

NewFarmMonitorDetails instantiates a new FarmMonitorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmMonitorDetailsWithDefaults

`func NewFarmMonitorDetailsWithDefaults() *FarmMonitorDetails`

NewFarmMonitorDetailsWithDefaults instantiates a new FarmMonitorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *FarmMonitorDetails) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *FarmMonitorDetails) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *FarmMonitorDetails) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *FarmMonitorDetails) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetType

`func (o *FarmMonitorDetails) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FarmMonitorDetails) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FarmMonitorDetails) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


