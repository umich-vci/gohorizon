# SecuritySettingsUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataRecoveryPasswordHash** | Pointer to **[]string** | The SHA-256 hash of the (UTF-8) data recovery password. | [optional] 
**DataRecoveryPasswordHint** | Pointer to **string** | The data recovery password hint. This property has a maximum length of 128 characters. | [optional] 
**MessageSecurityMode** | **string** | Determines if signing and verification of the JMS messages passed between Horizon components takes place. * DISABLED: Message security mode is disabled. * MIXED: Message security mode is enabled but not enforced. * ENABLED: Message security mode is enabled. Unsigned messages are rejected by Horizon components. * ENHANCED: Message Security mode is Enhanced. Message signing and validation is performed based on the current Security Level and desktop Message Security mode. | 
**ReAuthSecureTunnelAfterInterruption** | Pointer to **bool** | Determines if user credentials must be re-authenticated after a network interruption when Horizon clients use secure tunnel connections to Horizon resources. When you select this setting, if a secure tunnel connection ends during a session, Horizon Client requires the user to re-authenticate before reconnecting. | [optional] 

## Methods

### NewSecuritySettingsUpdateSpec

`func NewSecuritySettingsUpdateSpec(messageSecurityMode string, ) *SecuritySettingsUpdateSpec`

NewSecuritySettingsUpdateSpec instantiates a new SecuritySettingsUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySettingsUpdateSpecWithDefaults

`func NewSecuritySettingsUpdateSpecWithDefaults() *SecuritySettingsUpdateSpec`

NewSecuritySettingsUpdateSpecWithDefaults instantiates a new SecuritySettingsUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataRecoveryPasswordHash

`func (o *SecuritySettingsUpdateSpec) GetDataRecoveryPasswordHash() []string`

GetDataRecoveryPasswordHash returns the DataRecoveryPasswordHash field if non-nil, zero value otherwise.

### GetDataRecoveryPasswordHashOk

`func (o *SecuritySettingsUpdateSpec) GetDataRecoveryPasswordHashOk() (*[]string, bool)`

GetDataRecoveryPasswordHashOk returns a tuple with the DataRecoveryPasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRecoveryPasswordHash

`func (o *SecuritySettingsUpdateSpec) SetDataRecoveryPasswordHash(v []string)`

SetDataRecoveryPasswordHash sets DataRecoveryPasswordHash field to given value.

### HasDataRecoveryPasswordHash

`func (o *SecuritySettingsUpdateSpec) HasDataRecoveryPasswordHash() bool`

HasDataRecoveryPasswordHash returns a boolean if a field has been set.

### GetDataRecoveryPasswordHint

`func (o *SecuritySettingsUpdateSpec) GetDataRecoveryPasswordHint() string`

GetDataRecoveryPasswordHint returns the DataRecoveryPasswordHint field if non-nil, zero value otherwise.

### GetDataRecoveryPasswordHintOk

`func (o *SecuritySettingsUpdateSpec) GetDataRecoveryPasswordHintOk() (*string, bool)`

GetDataRecoveryPasswordHintOk returns a tuple with the DataRecoveryPasswordHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRecoveryPasswordHint

`func (o *SecuritySettingsUpdateSpec) SetDataRecoveryPasswordHint(v string)`

SetDataRecoveryPasswordHint sets DataRecoveryPasswordHint field to given value.

### HasDataRecoveryPasswordHint

`func (o *SecuritySettingsUpdateSpec) HasDataRecoveryPasswordHint() bool`

HasDataRecoveryPasswordHint returns a boolean if a field has been set.

### GetMessageSecurityMode

`func (o *SecuritySettingsUpdateSpec) GetMessageSecurityMode() string`

GetMessageSecurityMode returns the MessageSecurityMode field if non-nil, zero value otherwise.

### GetMessageSecurityModeOk

`func (o *SecuritySettingsUpdateSpec) GetMessageSecurityModeOk() (*string, bool)`

GetMessageSecurityModeOk returns a tuple with the MessageSecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityMode

`func (o *SecuritySettingsUpdateSpec) SetMessageSecurityMode(v string)`

SetMessageSecurityMode sets MessageSecurityMode field to given value.


### GetReAuthSecureTunnelAfterInterruption

`func (o *SecuritySettingsUpdateSpec) GetReAuthSecureTunnelAfterInterruption() bool`

GetReAuthSecureTunnelAfterInterruption returns the ReAuthSecureTunnelAfterInterruption field if non-nil, zero value otherwise.

### GetReAuthSecureTunnelAfterInterruptionOk

`func (o *SecuritySettingsUpdateSpec) GetReAuthSecureTunnelAfterInterruptionOk() (*bool, bool)`

GetReAuthSecureTunnelAfterInterruptionOk returns a tuple with the ReAuthSecureTunnelAfterInterruption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAuthSecureTunnelAfterInterruption

`func (o *SecuritySettingsUpdateSpec) SetReAuthSecureTunnelAfterInterruption(v bool)`

SetReAuthSecureTunnelAfterInterruption sets ReAuthSecureTunnelAfterInterruption field to given value.

### HasReAuthSecureTunnelAfterInterruption

`func (o *SecuritySettingsUpdateSpec) HasReAuthSecureTunnelAfterInterruption() bool`

HasReAuthSecureTunnelAfterInterruption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


