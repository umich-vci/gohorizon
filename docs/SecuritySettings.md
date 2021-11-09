# SecuritySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterPublicKey** | **string** | The Base 64 encoded public key of the cluster in PEM format. | 
**ClusterPublicKeyId** | Pointer to **string** | Key Id to identify the cluster&#39;s active key pair. | [optional] 
**DataRecoveryPasswordConfigured** | **bool** | Indicates whether the backup recovery password has been configured. | 
**MessageSecurityMode** | **string** | Determines if signing and verification of the JMS messages passed between Horizon components takes place. * DISABLED: Message security mode is disabled. * MIXED: Message security mode is enabled but not enforced. * ENABLED: Message security mode is enabled. Unsigned messages are rejected by Horizon components. * ENHANCED: Message Security mode is Enhanced. Message signing and validation is performed based on the current Security Level and desktop Message Security mode. | 
**MessageSecurityStatus** | **string** | The status of the JMS message security. This tracks the application of changes to messageSecurityMode. * READY: The cluster is performing at the specified message security mode. * INITIALIZING_ENHANCED: The cluster is initializing a transition to the ENHANCED message security mode. * PENDING_ENHANCED: The cluster is propagating the change to ENHANCED message security mode to all nodes. * LEAVING_ENHANCED: The cluster is leaving the ENHANCED message security mode. | 
**ReAuthSecureTunnelAfterInterruption** | **bool** | Determines if user credentials must be re-authenticated after a network interruption when Horizon clients use secure tunnel connections to Horizon resources. When you select this setting, if a secure tunnel connection ends during a session, Horizon Client requires the user to re-authenticate before reconnecting. | 

## Methods

### NewSecuritySettings

`func NewSecuritySettings(clusterPublicKey string, dataRecoveryPasswordConfigured bool, messageSecurityMode string, messageSecurityStatus string, reAuthSecureTunnelAfterInterruption bool, ) *SecuritySettings`

NewSecuritySettings instantiates a new SecuritySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySettingsWithDefaults

`func NewSecuritySettingsWithDefaults() *SecuritySettings`

NewSecuritySettingsWithDefaults instantiates a new SecuritySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterPublicKey

`func (o *SecuritySettings) GetClusterPublicKey() string`

GetClusterPublicKey returns the ClusterPublicKey field if non-nil, zero value otherwise.

### GetClusterPublicKeyOk

`func (o *SecuritySettings) GetClusterPublicKeyOk() (*string, bool)`

GetClusterPublicKeyOk returns a tuple with the ClusterPublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPublicKey

`func (o *SecuritySettings) SetClusterPublicKey(v string)`

SetClusterPublicKey sets ClusterPublicKey field to given value.


### GetClusterPublicKeyId

`func (o *SecuritySettings) GetClusterPublicKeyId() string`

GetClusterPublicKeyId returns the ClusterPublicKeyId field if non-nil, zero value otherwise.

### GetClusterPublicKeyIdOk

`func (o *SecuritySettings) GetClusterPublicKeyIdOk() (*string, bool)`

GetClusterPublicKeyIdOk returns a tuple with the ClusterPublicKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPublicKeyId

`func (o *SecuritySettings) SetClusterPublicKeyId(v string)`

SetClusterPublicKeyId sets ClusterPublicKeyId field to given value.

### HasClusterPublicKeyId

`func (o *SecuritySettings) HasClusterPublicKeyId() bool`

HasClusterPublicKeyId returns a boolean if a field has been set.

### GetDataRecoveryPasswordConfigured

`func (o *SecuritySettings) GetDataRecoveryPasswordConfigured() bool`

GetDataRecoveryPasswordConfigured returns the DataRecoveryPasswordConfigured field if non-nil, zero value otherwise.

### GetDataRecoveryPasswordConfiguredOk

`func (o *SecuritySettings) GetDataRecoveryPasswordConfiguredOk() (*bool, bool)`

GetDataRecoveryPasswordConfiguredOk returns a tuple with the DataRecoveryPasswordConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRecoveryPasswordConfigured

`func (o *SecuritySettings) SetDataRecoveryPasswordConfigured(v bool)`

SetDataRecoveryPasswordConfigured sets DataRecoveryPasswordConfigured field to given value.


### GetMessageSecurityMode

`func (o *SecuritySettings) GetMessageSecurityMode() string`

GetMessageSecurityMode returns the MessageSecurityMode field if non-nil, zero value otherwise.

### GetMessageSecurityModeOk

`func (o *SecuritySettings) GetMessageSecurityModeOk() (*string, bool)`

GetMessageSecurityModeOk returns a tuple with the MessageSecurityMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityMode

`func (o *SecuritySettings) SetMessageSecurityMode(v string)`

SetMessageSecurityMode sets MessageSecurityMode field to given value.


### GetMessageSecurityStatus

`func (o *SecuritySettings) GetMessageSecurityStatus() string`

GetMessageSecurityStatus returns the MessageSecurityStatus field if non-nil, zero value otherwise.

### GetMessageSecurityStatusOk

`func (o *SecuritySettings) GetMessageSecurityStatusOk() (*string, bool)`

GetMessageSecurityStatusOk returns a tuple with the MessageSecurityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSecurityStatus

`func (o *SecuritySettings) SetMessageSecurityStatus(v string)`

SetMessageSecurityStatus sets MessageSecurityStatus field to given value.


### GetReAuthSecureTunnelAfterInterruption

`func (o *SecuritySettings) GetReAuthSecureTunnelAfterInterruption() bool`

GetReAuthSecureTunnelAfterInterruption returns the ReAuthSecureTunnelAfterInterruption field if non-nil, zero value otherwise.

### GetReAuthSecureTunnelAfterInterruptionOk

`func (o *SecuritySettings) GetReAuthSecureTunnelAfterInterruptionOk() (*bool, bool)`

GetReAuthSecureTunnelAfterInterruptionOk returns a tuple with the ReAuthSecureTunnelAfterInterruption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReAuthSecureTunnelAfterInterruption

`func (o *SecuritySettings) SetReAuthSecureTunnelAfterInterruption(v bool)`

SetReAuthSecureTunnelAfterInterruption sets ReAuthSecureTunnelAfterInterruption field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


