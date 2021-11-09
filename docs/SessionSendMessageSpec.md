# SessionSendMessageSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | Message to be sent to session(s). | 
**MessageType** | **string** | Type of message to be sent to session(s). * ERROR: Message is of error type. * WARNING: Message is of warning type. * INFO: Message is of information type. | 
**SessionIds** | **[]string** | List of session ids to which message is to be sent | 

## Methods

### NewSessionSendMessageSpec

`func NewSessionSendMessageSpec(message string, messageType string, sessionIds []string, ) *SessionSendMessageSpec`

NewSessionSendMessageSpec instantiates a new SessionSendMessageSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionSendMessageSpecWithDefaults

`func NewSessionSendMessageSpecWithDefaults() *SessionSendMessageSpec`

NewSessionSendMessageSpecWithDefaults instantiates a new SessionSendMessageSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SessionSendMessageSpec) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SessionSendMessageSpec) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SessionSendMessageSpec) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetMessageType

`func (o *SessionSendMessageSpec) GetMessageType() string`

GetMessageType returns the MessageType field if non-nil, zero value otherwise.

### GetMessageTypeOk

`func (o *SessionSendMessageSpec) GetMessageTypeOk() (*string, bool)`

GetMessageTypeOk returns a tuple with the MessageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageType

`func (o *SessionSendMessageSpec) SetMessageType(v string)`

SetMessageType sets MessageType field to given value.


### GetSessionIds

`func (o *SessionSendMessageSpec) GetSessionIds() []string`

GetSessionIds returns the SessionIds field if non-nil, zero value otherwise.

### GetSessionIdsOk

`func (o *SessionSendMessageSpec) GetSessionIdsOk() (*[]string, bool)`

GetSessionIdsOk returns a tuple with the SessionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionIds

`func (o *SessionSendMessageSpec) SetSessionIds(v []string)`

SetSessionIds sets SessionIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


