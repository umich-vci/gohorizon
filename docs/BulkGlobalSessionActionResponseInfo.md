# BulkGlobalSessionActionResponseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**[]BulkItemResponseInfo**](BulkItemResponseInfo.md) | List of BulkItemResponseInfo corresponding to each session id in the action operation. | [optional] 
**ErrorMessages** | Pointer to **[]string** | Reasons for failure of the operation. | [optional] 
**PodId** | Pointer to **string** | ID of the hosting pod for the sessions. | [optional] 
**StatusCode** | Pointer to **int32** | HTTP Status Code of the operation. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp in milliseconds when the operation failed.  Measured as epoch time. | [optional] 

## Methods

### NewBulkGlobalSessionActionResponseInfo

`func NewBulkGlobalSessionActionResponseInfo() *BulkGlobalSessionActionResponseInfo`

NewBulkGlobalSessionActionResponseInfo instantiates a new BulkGlobalSessionActionResponseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkGlobalSessionActionResponseInfoWithDefaults

`func NewBulkGlobalSessionActionResponseInfoWithDefaults() *BulkGlobalSessionActionResponseInfo`

NewBulkGlobalSessionActionResponseInfoWithDefaults instantiates a new BulkGlobalSessionActionResponseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *BulkGlobalSessionActionResponseInfo) GetDetails() []BulkItemResponseInfo`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *BulkGlobalSessionActionResponseInfo) GetDetailsOk() (*[]BulkItemResponseInfo, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *BulkGlobalSessionActionResponseInfo) SetDetails(v []BulkItemResponseInfo)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *BulkGlobalSessionActionResponseInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetErrorMessages

`func (o *BulkGlobalSessionActionResponseInfo) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *BulkGlobalSessionActionResponseInfo) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *BulkGlobalSessionActionResponseInfo) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *BulkGlobalSessionActionResponseInfo) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetPodId

`func (o *BulkGlobalSessionActionResponseInfo) GetPodId() string`

GetPodId returns the PodId field if non-nil, zero value otherwise.

### GetPodIdOk

`func (o *BulkGlobalSessionActionResponseInfo) GetPodIdOk() (*string, bool)`

GetPodIdOk returns a tuple with the PodId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodId

`func (o *BulkGlobalSessionActionResponseInfo) SetPodId(v string)`

SetPodId sets PodId field to given value.

### HasPodId

`func (o *BulkGlobalSessionActionResponseInfo) HasPodId() bool`

HasPodId returns a boolean if a field has been set.

### GetStatusCode

`func (o *BulkGlobalSessionActionResponseInfo) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BulkGlobalSessionActionResponseInfo) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BulkGlobalSessionActionResponseInfo) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *BulkGlobalSessionActionResponseInfo) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *BulkGlobalSessionActionResponseInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BulkGlobalSessionActionResponseInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BulkGlobalSessionActionResponseInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *BulkGlobalSessionActionResponseInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


