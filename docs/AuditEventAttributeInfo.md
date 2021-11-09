# AuditEventAttributeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventData** | Pointer to **map[string]map[string]interface{}** | Key value pairs representing Extended attributes related to the event.  | [optional] 
**Id** | Pointer to **int64** | Unique id representing an event.  | [optional] 

## Methods

### NewAuditEventAttributeInfo

`func NewAuditEventAttributeInfo() *AuditEventAttributeInfo`

NewAuditEventAttributeInfo instantiates a new AuditEventAttributeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventAttributeInfoWithDefaults

`func NewAuditEventAttributeInfoWithDefaults() *AuditEventAttributeInfo`

NewAuditEventAttributeInfoWithDefaults instantiates a new AuditEventAttributeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventData

`func (o *AuditEventAttributeInfo) GetEventData() map[string]map[string]interface{}`

GetEventData returns the EventData field if non-nil, zero value otherwise.

### GetEventDataOk

`func (o *AuditEventAttributeInfo) GetEventDataOk() (*map[string]map[string]interface{}, bool)`

GetEventDataOk returns a tuple with the EventData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventData

`func (o *AuditEventAttributeInfo) SetEventData(v map[string]map[string]interface{})`

SetEventData sets EventData field to given value.

### HasEventData

`func (o *AuditEventAttributeInfo) HasEventData() bool`

HasEventData returns a boolean if a field has been set.

### GetId

`func (o *AuditEventAttributeInfo) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditEventAttributeInfo) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditEventAttributeInfo) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuditEventAttributeInfo) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


