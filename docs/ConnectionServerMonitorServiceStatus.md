# ConnectionServerMonitorServiceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceName** | **string** | Service name of the Connection Server. * PCOIP_SECURE_GATEWAY: PCoIP Secure Gateway service. * BLAST_SECURE_GATEWAY: BLAST Secure Gateway service. * SECURITY_GATEWAY_COMPONENT: Security Gateway Component service. | 
**Status** | **string** | Status of the service. * UP: The Windows service is UP and running. * DOWN: The Windows service is not UP. * UNKNOWN: The Windows service state is Unknown. | 

## Methods

### NewConnectionServerMonitorServiceStatus

`func NewConnectionServerMonitorServiceStatus(serviceName string, status string, ) *ConnectionServerMonitorServiceStatus`

NewConnectionServerMonitorServiceStatus instantiates a new ConnectionServerMonitorServiceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionServerMonitorServiceStatusWithDefaults

`func NewConnectionServerMonitorServiceStatusWithDefaults() *ConnectionServerMonitorServiceStatus`

NewConnectionServerMonitorServiceStatusWithDefaults instantiates a new ConnectionServerMonitorServiceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceName

`func (o *ConnectionServerMonitorServiceStatus) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *ConnectionServerMonitorServiceStatus) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *ConnectionServerMonitorServiceStatus) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.


### GetStatus

`func (o *ConnectionServerMonitorServiceStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ConnectionServerMonitorServiceStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ConnectionServerMonitorServiceStatus) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


