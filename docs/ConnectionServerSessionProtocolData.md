# ConnectionServerSessionProtocolData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SessionCount** | **int32** | Number of active sessions launched using the protocol. | 
**SessionProtocol** | **string** | Protocol used in launching the session. * PCOIP: Display protocol is PCoIP. * RDP: Display protocol is RDP. * BLAST: Display protocol is BLAST. * CONSOLE: Display protocol is console. * UNKNOWN: Display protocol is unknown. | 

## Methods

### NewConnectionServerSessionProtocolData

`func NewConnectionServerSessionProtocolData(sessionCount int32, sessionProtocol string, ) *ConnectionServerSessionProtocolData`

NewConnectionServerSessionProtocolData instantiates a new ConnectionServerSessionProtocolData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServerSessionProtocolDataWithDefaults

`func NewConnectionServerSessionProtocolDataWithDefaults() *ConnectionServerSessionProtocolData`

NewConnectionServerSessionProtocolDataWithDefaults instantiates a new ConnectionServerSessionProtocolData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessionCount

`func (o *ConnectionServerSessionProtocolData) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *ConnectionServerSessionProtocolData) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *ConnectionServerSessionProtocolData) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.


### GetSessionProtocol

`func (o *ConnectionServerSessionProtocolData) GetSessionProtocol() string`

GetSessionProtocol returns the SessionProtocol field if non-nil, zero value otherwise.

### GetSessionProtocolOk

`func (o *ConnectionServerSessionProtocolData) GetSessionProtocolOk() (*string, bool)`

GetSessionProtocolOk returns a tuple with the SessionProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionProtocol

`func (o *ConnectionServerSessionProtocolData) SetSessionProtocol(v string)`

SetSessionProtocol sets SessionProtocol field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


