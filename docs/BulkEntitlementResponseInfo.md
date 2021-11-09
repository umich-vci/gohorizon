# BulkEntitlementResponseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]EntitlementResponseInfo**](EntitlementResponseInfo.md) | List of entitlement response info objects correpsonding to each of the given ad-user-or-group SID. | [optional] 
**ErrorMessages** | Pointer to **[]string** | Reasons for the failure of the operation. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the resource for the entitlement operation. | [optional] 
**StatusCode** | Pointer to **int32** | Response HTTP status code of the operation. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp in milliseconds when the operation failed. Measured as epoch time. | [optional] 

## Methods

### NewBulkEntitlementResponseInfo

`func NewBulkEntitlementResponseInfo() *BulkEntitlementResponseInfo`

NewBulkEntitlementResponseInfo instantiates a new BulkEntitlementResponseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEntitlementResponseInfoWithDefaults

`func NewBulkEntitlementResponseInfoWithDefaults() *BulkEntitlementResponseInfo`

NewBulkEntitlementResponseInfoWithDefaults instantiates a new BulkEntitlementResponseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *BulkEntitlementResponseInfo) GetDetails() []EntitlementResponseInfo`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BulkEntitlementResponseInfo) GetDetailsOk() (*[]EntitlementResponseInfo, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BulkEntitlementResponseInfo) SetDetails(v []EntitlementResponseInfo)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BulkEntitlementResponseInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetErrorMessages

`func (o *BulkEntitlementResponseInfo) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *BulkEntitlementResponseInfo) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *BulkEntitlementResponseInfo) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *BulkEntitlementResponseInfo) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetId

`func (o *BulkEntitlementResponseInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkEntitlementResponseInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkEntitlementResponseInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkEntitlementResponseInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatusCode

`func (o *BulkEntitlementResponseInfo) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BulkEntitlementResponseInfo) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BulkEntitlementResponseInfo) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BulkEntitlementResponseInfo) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *BulkEntitlementResponseInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BulkEntitlementResponseInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BulkEntitlementResponseInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BulkEntitlementResponseInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


