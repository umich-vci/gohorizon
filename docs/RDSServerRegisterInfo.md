# RDSServerRegisterInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique ID representing the RDS Server. | [optional] 
**PairingToken** | Pointer to **string** | The pairing token for the RDS Server. | [optional] 

## Methods

### NewRDSServerRegisterInfo

`func NewRDSServerRegisterInfo() *RDSServerRegisterInfo`

NewRDSServerRegisterInfo instantiates a new RDSServerRegisterInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSServerRegisterInfoWithDefaults

`func NewRDSServerRegisterInfoWithDefaults() *RDSServerRegisterInfo`

NewRDSServerRegisterInfoWithDefaults instantiates a new RDSServerRegisterInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RDSServerRegisterInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RDSServerRegisterInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RDSServerRegisterInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RDSServerRegisterInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPairingToken

`func (o *RDSServerRegisterInfo) GetPairingToken() string`

GetPairingToken returns the PairingToken field if non-nil, zero value otherwise.

### GetPairingTokenOk

`func (o *RDSServerRegisterInfo) GetPairingTokenOk() (*string, bool)`

GetPairingTokenOk returns a tuple with the PairingToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingToken

`func (o *RDSServerRegisterInfo) SetPairingToken(v string)`

SetPairingToken sets PairingToken field to given value.

### HasPairingToken

`func (o *RDSServerRegisterInfo) HasPairingToken() bool`

HasPairingToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


