# ConnectionServerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID of the connection server to which this status pertains. | [optional] 
**Message** | Pointer to **string** | The CPA error message for the connection server if any, is populated, or a success message. | [optional] 
**MessageCode** | Pointer to **string** | Message code of the message. * LMV_OP_OK: Pod Federation operation is successful. * LMV_OP_ERROR: Pod Federation operation failed. * LMV_OP_ERROR_UNKNOWN: Pod Federation operation failed with unknown error. * LMV_OP_CANNOT_APPLY_DATA: Pod Federation operation could not apply data. * LMV_OP_CANNOT_APPLY_SCHEMA: Pod Federation operation could not apply schema. * LMV_OP_FILESYSTEM_ERROR: Pod Federation operation encountered file system error. * LMV_OP_NO_TIME_SYNC: Pod Federation operation encountered time synchronization error. * LMV_OP_NO_PERMISSION: No permission to perform Pod Federation operation. * LMV_OP_REPLICATION_ERROR: Pod Federation operation encountered replication error. * LMV_OP_UNREACHABLE_SERVER: The server is unreachable to perform Pod Federation operation. | [optional] 
**Name** | Pointer to **string** | Name of the connection server. | [optional] 
**PendingPercentage** | Pointer to **int32** |  Value between 0 and 100 representing CPA operation completion percentage when connection server status is pending. | [optional] 
**Status** | Pointer to **string** | CPA status of the connection server. * ENABLED: CPA is enabled. * DISABLED: CPA is disabled. * PENDING: CPA is undergoing an operation related to initialization, uninitialization, joining, or unjoining. * ENABLE_ERROR: The connection server has failed to reach the ENABLED status in a timely manner. This may also indicate the current connection server was recently installed. * DISABLE_ERROR: The connection server has failed to reach the DISABLED status in a timely manner. | [optional] 

## Methods

### NewConnectionServerStatus

`func NewConnectionServerStatus() *ConnectionServerStatus`

NewConnectionServerStatus instantiates a new ConnectionServerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServerStatusWithDefaults

`func NewConnectionServerStatusWithDefaults() *ConnectionServerStatus`

NewConnectionServerStatusWithDefaults instantiates a new ConnectionServerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConnectionServerStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConnectionServerStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConnectionServerStatus) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConnectionServerStatus) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessage

`func (o *ConnectionServerStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ConnectionServerStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ConnectionServerStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ConnectionServerStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetMessageCode

`func (o *ConnectionServerStatus) GetMessageCode() string`

GetMessageCode returns the MessageCode field if non-nil, zero value otherwise.

### GetMessageCodeOk

`func (o *ConnectionServerStatus) GetMessageCodeOk() (*string, bool)`

GetMessageCodeOk returns a tuple with the MessageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCode

`func (o *ConnectionServerStatus) SetMessageCode(v string)`

SetMessageCode sets MessageCode field to given value.

### HasMessageCode

`func (o *ConnectionServerStatus) HasMessageCode() bool`

HasMessageCode returns a boolean if a field has been set.

### GetName

`func (o *ConnectionServerStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConnectionServerStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConnectionServerStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConnectionServerStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPendingPercentage

`func (o *ConnectionServerStatus) GetPendingPercentage() int32`

GetPendingPercentage returns the PendingPercentage field if non-nil, zero value otherwise.

### GetPendingPercentageOk

`func (o *ConnectionServerStatus) GetPendingPercentageOk() (*int32, bool)`

GetPendingPercentageOk returns a tuple with the PendingPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingPercentage

`func (o *ConnectionServerStatus) SetPendingPercentage(v int32)`

SetPendingPercentage sets PendingPercentage field to given value.

### HasPendingPercentage

`func (o *ConnectionServerStatus) HasPendingPercentage() bool`

HasPendingPercentage returns a boolean if a field has been set.

### GetStatus

`func (o *ConnectionServerStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionServerStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionServerStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ConnectionServerStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


