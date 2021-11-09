# GatewayMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveConnectionCount** | Pointer to **int32** | Number of active connections for the gateway. Includes PCoIP and BLAST connection counts. | [optional] 
**BlastConnectionCount** | Pointer to **int32** | Number of BLAST connections for the gateway. | [optional] 
**Details** | [**GatewayMonitorDetails**](GatewayMonitorDetails.md) |  | 
**Id** | **string** | Unique ID of the Gateway. | 
**Name** | **string** | Gateway name. | 
**PcoipConnectionCount** | Pointer to **int32** | Number of PCoIP connections for the gateway. | [optional] 
**Status** | **string** | Status of the Gateway. * NOT_CONTACTED: There has been no contact from the gateway. * PROBLEM: The gateway has reported a problem. * STALE: Gateway is stale. Gateway will be marked as stale when Connection Server does not receive any request from the Gateway in last two successive intervals. * OK: The Gateway is working as expected. | 

## Methods

### NewGatewayMonitorInfo

`func NewGatewayMonitorInfo(details GatewayMonitorDetails, id string, name string, status string, ) *GatewayMonitorInfo`

NewGatewayMonitorInfo instantiates a new GatewayMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayMonitorInfoWithDefaults

`func NewGatewayMonitorInfoWithDefaults() *GatewayMonitorInfo`

NewGatewayMonitorInfoWithDefaults instantiates a new GatewayMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveConnectionCount

`func (o *GatewayMonitorInfo) GetActiveConnectionCount() int32`

GetActiveConnectionCount returns the ActiveConnectionCount field if non-nil, zero value otherwise.

### GetActiveConnectionCountOk

`func (o *GatewayMonitorInfo) GetActiveConnectionCountOk() (*int32, bool)`

GetActiveConnectionCountOk returns a tuple with the ActiveConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveConnectionCount

`func (o *GatewayMonitorInfo) SetActiveConnectionCount(v int32)`

SetActiveConnectionCount sets ActiveConnectionCount field to given value.

### HasActiveConnectionCount

`func (o *GatewayMonitorInfo) HasActiveConnectionCount() bool`

HasActiveConnectionCount returns a boolean if a field has been set.

### GetBlastConnectionCount

`func (o *GatewayMonitorInfo) GetBlastConnectionCount() int32`

GetBlastConnectionCount returns the BlastConnectionCount field if non-nil, zero value otherwise.

### GetBlastConnectionCountOk

`func (o *GatewayMonitorInfo) GetBlastConnectionCountOk() (*int32, bool)`

GetBlastConnectionCountOk returns a tuple with the BlastConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastConnectionCount

`func (o *GatewayMonitorInfo) SetBlastConnectionCount(v int32)`

SetBlastConnectionCount sets BlastConnectionCount field to given value.

### HasBlastConnectionCount

`func (o *GatewayMonitorInfo) HasBlastConnectionCount() bool`

HasBlastConnectionCount returns a boolean if a field has been set.

### GetDetails

`func (o *GatewayMonitorInfo) GetDetails() GatewayMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GatewayMonitorInfo) GetDetailsOk() (*GatewayMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GatewayMonitorInfo) SetDetails(v GatewayMonitorDetails)`

SetDetails sets Details field to given value.


### GetId

`func (o *GatewayMonitorInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayMonitorInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayMonitorInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GatewayMonitorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayMonitorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayMonitorInfo) SetName(v string)`

SetName sets Name field to given value.


### GetPcoipConnectionCount

`func (o *GatewayMonitorInfo) GetPcoipConnectionCount() int32`

GetPcoipConnectionCount returns the PcoipConnectionCount field if non-nil, zero value otherwise.

### GetPcoipConnectionCountOk

`func (o *GatewayMonitorInfo) GetPcoipConnectionCountOk() (*int32, bool)`

GetPcoipConnectionCountOk returns a tuple with the PcoipConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcoipConnectionCount

`func (o *GatewayMonitorInfo) SetPcoipConnectionCount(v int32)`

SetPcoipConnectionCount sets PcoipConnectionCount field to given value.

### HasPcoipConnectionCount

`func (o *GatewayMonitorInfo) HasPcoipConnectionCount() bool`

HasPcoipConnectionCount returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayMonitorInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayMonitorInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayMonitorInfo) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


