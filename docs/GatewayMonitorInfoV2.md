# GatewayMonitorInfoV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveConnectionCount** | Pointer to **int32** | Number of active connections for the gateway. Includes PCoIP and BLAST connection counts. | [optional] 
**BlastConnectionCount** | Pointer to **int32** | Number of BLAST connections for the gateway. | [optional] 
**Details** | Pointer to [**GatewayMonitorDetails**](GatewayMonitorDetails.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the Gateway. | [optional] 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp in milliseconds when the last update was obtained. Measured as epoch time. | [optional] 
**Name** | Pointer to **string** | Gateway name. | [optional] 
**PcoipConnectionCount** | Pointer to **int32** | Number of PCoIP connections for the gateway. | [optional] 
**Status** | Pointer to **string** | Status of the Gateway. * NOT_CONTACTED: There has been no contact from the gateway. * PROBLEM: The gateway has reported a problem. * STALE: Gateway is stale. Gateway will be marked as stale when Connection Server does not receive any request from the Gateway in last two successive intervals. * OK: The Gateway is working as expected. | [optional] 

## Methods

### NewGatewayMonitorInfoV2

`func NewGatewayMonitorInfoV2() *GatewayMonitorInfoV2`

NewGatewayMonitorInfoV2 instantiates a new GatewayMonitorInfoV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayMonitorInfoV2WithDefaults

`func NewGatewayMonitorInfoV2WithDefaults() *GatewayMonitorInfoV2`

NewGatewayMonitorInfoV2WithDefaults instantiates a new GatewayMonitorInfoV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveConnectionCount

`func (o *GatewayMonitorInfoV2) GetActiveConnectionCount() int32`

GetActiveConnectionCount returns the ActiveConnectionCount field if non-nil, zero value otherwise.

### GetActiveConnectionCountOk

`func (o *GatewayMonitorInfoV2) GetActiveConnectionCountOk() (*int32, bool)`

GetActiveConnectionCountOk returns a tuple with the ActiveConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveConnectionCount

`func (o *GatewayMonitorInfoV2) SetActiveConnectionCount(v int32)`

SetActiveConnectionCount sets ActiveConnectionCount field to given value.

### HasActiveConnectionCount

`func (o *GatewayMonitorInfoV2) HasActiveConnectionCount() bool`

HasActiveConnectionCount returns a boolean if a field has been set.

### GetBlastConnectionCount

`func (o *GatewayMonitorInfoV2) GetBlastConnectionCount() int32`

GetBlastConnectionCount returns the BlastConnectionCount field if non-nil, zero value otherwise.

### GetBlastConnectionCountOk

`func (o *GatewayMonitorInfoV2) GetBlastConnectionCountOk() (*int32, bool)`

GetBlastConnectionCountOk returns a tuple with the BlastConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlastConnectionCount

`func (o *GatewayMonitorInfoV2) SetBlastConnectionCount(v int32)`

SetBlastConnectionCount sets BlastConnectionCount field to given value.

### HasBlastConnectionCount

`func (o *GatewayMonitorInfoV2) HasBlastConnectionCount() bool`

HasBlastConnectionCount returns a boolean if a field has been set.

### GetDetails

`func (o *GatewayMonitorInfoV2) GetDetails() GatewayMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *GatewayMonitorInfoV2) GetDetailsOk() (*GatewayMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *GatewayMonitorInfoV2) SetDetails(v GatewayMonitorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *GatewayMonitorInfoV2) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *GatewayMonitorInfoV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayMonitorInfoV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayMonitorInfoV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayMonitorInfoV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastUpdatedTimestamp

`func (o *GatewayMonitorInfoV2) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *GatewayMonitorInfoV2) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *GatewayMonitorInfoV2) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *GatewayMonitorInfoV2) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *GatewayMonitorInfoV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayMonitorInfoV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayMonitorInfoV2) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayMonitorInfoV2) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcoipConnectionCount

`func (o *GatewayMonitorInfoV2) GetPcoipConnectionCount() int32`

GetPcoipConnectionCount returns the PcoipConnectionCount field if non-nil, zero value otherwise.

### GetPcoipConnectionCountOk

`func (o *GatewayMonitorInfoV2) GetPcoipConnectionCountOk() (*int32, bool)`

GetPcoipConnectionCountOk returns a tuple with the PcoipConnectionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcoipConnectionCount

`func (o *GatewayMonitorInfoV2) SetPcoipConnectionCount(v int32)`

SetPcoipConnectionCount sets PcoipConnectionCount field to given value.

### HasPcoipConnectionCount

`func (o *GatewayMonitorInfoV2) HasPcoipConnectionCount() bool`

HasPcoipConnectionCount returns a boolean if a field has been set.

### GetStatus

`func (o *GatewayMonitorInfoV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayMonitorInfoV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayMonitorInfoV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GatewayMonitorInfoV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


