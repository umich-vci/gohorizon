# GlobalSessionSendMessageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalSessionActionSpecs** | [**[]GlobalSessionActionSpec**](GlobalSessionActionSpec.md) | Sessions to which message is to be sent. | 
**Message** | **string** | Message to be sent to sessions. | 
**MessageType** | **string** | Type of message to be sent to sessions. * ERROR: Message is of error type. * WARNING: Message is of warning type. * INFO: Message is of information type. | 

## Methods

### NewGlobalSessionSendMessageSpec

`func NewGlobalSessionSendMessageSpec(globalSessionActionSpecs []GlobalSessionActionSpec, message string, messageType string, ) *GlobalSessionSendMessageSpec`

NewGlobalSessionSendMessageSpec instantiates a new GlobalSessionSendMessageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalSessionSendMessageSpecWithDefaults

`func NewGlobalSessionSendMessageSpecWithDefaults() *GlobalSessionSendMessageSpec`

NewGlobalSessionSendMessageSpecWithDefaults instantiates a new GlobalSessionSendMessageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalSessionActionSpecs

`func (o *GlobalSessionSendMessageSpec) GetGlobalSessionActionSpecs() []GlobalSessionActionSpec`

GetGlobalSessionActionSpecs returns the GlobalSessionActionSpecs field if non-nil, zero value otherwise.

### GetGlobalSessionActionSpecsOk

`func (o *GlobalSessionSendMessageSpec) GetGlobalSessionActionSpecsOk() (*[]GlobalSessionActionSpec, bool)`

GetGlobalSessionActionSpecsOk returns a tuple with the GlobalSessionActionSpecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalSessionActionSpecs

`func (o *GlobalSessionSendMessageSpec) SetGlobalSessionActionSpecs(v []GlobalSessionActionSpec)`

SetGlobalSessionActionSpecs sets GlobalSessionActionSpecs field to given value.


### GetMessage

`func (o *GlobalSessionSendMessageSpec) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GlobalSessionSendMessageSpec) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GlobalSessionSendMessageSpec) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageType

`func (o *GlobalSessionSendMessageSpec) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *GlobalSessionSendMessageSpec) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *GlobalSessionSendMessageSpec) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


