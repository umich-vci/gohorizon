# EntitlementResponseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdUserOrGroupId** | Pointer to **string** | Unique SID representing the ad-user-or-group | [optional] 
**ErrorMessages** | Pointer to **[]string** | Reasons for the failure of the operation. | [optional] 
**StatusCode** | Pointer to **int32** | Response HTTP status code of the operation. | [optional] 
**Timestamp** | Pointer to **int64** | Timestamp in milliseconds when the operation failed. Measured as epoch time. | [optional] 

## Methods

### NewEntitlementResponseInfo

`func NewEntitlementResponseInfo() *EntitlementResponseInfo`

NewEntitlementResponseInfo instantiates a new EntitlementResponseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementResponseInfoWithDefaults

`func NewEntitlementResponseInfoWithDefaults() *EntitlementResponseInfo`

NewEntitlementResponseInfoWithDefaults instantiates a new EntitlementResponseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdUserOrGroupId

`func (o *EntitlementResponseInfo) GetAdUserOrGroupId() string`

GetAdUserOrGroupId returns the AdUserOrGroupId field if non-nil, zero value otherwise.

### GetAdUserOrGroupIdOk

`func (o *EntitlementResponseInfo) GetAdUserOrGroupIdOk() (*string, bool)`

GetAdUserOrGroupIdOk returns a tuple with the AdUserOrGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUserOrGroupId

`func (o *EntitlementResponseInfo) SetAdUserOrGroupId(v string)`

SetAdUserOrGroupId sets AdUserOrGroupId field to given value.

### HasAdUserOrGroupId

`func (o *EntitlementResponseInfo) HasAdUserOrGroupId() bool`

HasAdUserOrGroupId returns a boolean if a field has been set.

### GetErrorMessages

`func (o *EntitlementResponseInfo) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *EntitlementResponseInfo) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *EntitlementResponseInfo) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *EntitlementResponseInfo) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetStatusCode

`func (o *EntitlementResponseInfo) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *EntitlementResponseInfo) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *EntitlementResponseInfo) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *EntitlementResponseInfo) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *EntitlementResponseInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EntitlementResponseInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EntitlementResponseInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EntitlementResponseInfo) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


