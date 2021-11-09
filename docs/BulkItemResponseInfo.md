# BulkItemResponseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessages** | Pointer to **[]string** | Reasons for the failure of the operation. | [optional] 
**Id** | Pointer to **string** | Unique ID representing the entity on which the operation was performed. Will not be populated for create operation if the operation fails. | [optional] 
**Key** | Pointer to **string** | Key on which the operation was performed. | [optional] 
**StatusCode** | **int32** | Response HTTP status code of the operation. | 
**Timestamp** | **int64** | Timestamp in milliseconds when the operation failed. Measured as epoch time. | 

## Methods

### NewBulkItemResponseInfo

`func NewBulkItemResponseInfo(statusCode int32, timestamp int64, ) *BulkItemResponseInfo`

NewBulkItemResponseInfo instantiates a new BulkItemResponseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkItemResponseInfoWithDefaults

`func NewBulkItemResponseInfoWithDefaults() *BulkItemResponseInfo`

NewBulkItemResponseInfoWithDefaults instantiates a new BulkItemResponseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessages

`func (o *BulkItemResponseInfo) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *BulkItemResponseInfo) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *BulkItemResponseInfo) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *BulkItemResponseInfo) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetId

`func (o *BulkItemResponseInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkItemResponseInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkItemResponseInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BulkItemResponseInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *BulkItemResponseInfo) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *BulkItemResponseInfo) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *BulkItemResponseInfo) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *BulkItemResponseInfo) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetStatusCode

`func (o *BulkItemResponseInfo) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *BulkItemResponseInfo) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *BulkItemResponseInfo) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### GetTimestamp

`func (o *BulkItemResponseInfo) GetTimestamp() int64`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *BulkItemResponseInfo) GetTimestampOk() (*int64, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *BulkItemResponseInfo) SetTimestamp(v int64)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


