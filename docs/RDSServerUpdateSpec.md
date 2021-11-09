# RDSServerUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Indicates if RDS server is enabled. | 
**MaxSessionsCountConfigured** | Pointer to **int32** | Maximum number of sessions for RDS server as configured by administrator. | [optional] 
**MaxSessionsTypeConfigured** | **string** | The configured RDS Server max sessions type. * UNLIMITED: The RDS Server has an unlimited number of sessions. * LIMITED: The RDS Server has a limited number of sessions. * UNCONFIGURED: The max number of sessions has not yet been defined for the RDSServer. | 

## Methods

### NewRDSServerUpdateSpec

`func NewRDSServerUpdateSpec(enabled bool, maxSessionsTypeConfigured string, ) *RDSServerUpdateSpec`

NewRDSServerUpdateSpec instantiates a new RDSServerUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSServerUpdateSpecWithDefaults

`func NewRDSServerUpdateSpecWithDefaults() *RDSServerUpdateSpec`

NewRDSServerUpdateSpecWithDefaults instantiates a new RDSServerUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RDSServerUpdateSpec) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RDSServerUpdateSpec) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RDSServerUpdateSpec) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxSessionsCountConfigured

`func (o *RDSServerUpdateSpec) GetMaxSessionsCountConfigured() int32`

GetMaxSessionsCountConfigured returns the MaxSessionsCountConfigured field if non-nil, zero value otherwise.

### GetMaxSessionsCountConfiguredOk

`func (o *RDSServerUpdateSpec) GetMaxSessionsCountConfiguredOk() (*int32, bool)`

GetMaxSessionsCountConfiguredOk returns a tuple with the MaxSessionsCountConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionsCountConfigured

`func (o *RDSServerUpdateSpec) SetMaxSessionsCountConfigured(v int32)`

SetMaxSessionsCountConfigured sets MaxSessionsCountConfigured field to given value.

### HasMaxSessionsCountConfigured

`func (o *RDSServerUpdateSpec) HasMaxSessionsCountConfigured() bool`

HasMaxSessionsCountConfigured returns a boolean if a field has been set.

### GetMaxSessionsTypeConfigured

`func (o *RDSServerUpdateSpec) GetMaxSessionsTypeConfigured() string`

GetMaxSessionsTypeConfigured returns the MaxSessionsTypeConfigured field if non-nil, zero value otherwise.

### GetMaxSessionsTypeConfiguredOk

`func (o *RDSServerUpdateSpec) GetMaxSessionsTypeConfiguredOk() (*string, bool)`

GetMaxSessionsTypeConfiguredOk returns a tuple with the MaxSessionsTypeConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSessionsTypeConfigured

`func (o *RDSServerUpdateSpec) SetMaxSessionsTypeConfigured(v string)`

SetMaxSessionsTypeConfigured sets MaxSessionsTypeConfigured field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


