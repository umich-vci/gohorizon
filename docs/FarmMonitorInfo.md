# FarmMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationCount** | Pointer to **int32** | Number of Applications published from the farm. | [optional] 
**Details** | Pointer to [**FarmMonitorDetails**](FarmMonitorDetails.md) |  | [optional] 
**Id** | Pointer to **string** | Unique ID of the Farm. | [optional] 
**Name** | Pointer to **string** | Farm name. | [optional] 
**RdsServerCount** | Pointer to **int32** | Number of RDS Servers in the farm. | [optional] 
**Status** | Pointer to **string** | Status of Farm. * OK: Farm is enabled and no servers are in WARNING or ERROR state. One or more server(s) may be DISABLED (including the case where all of them are DISABLED). * WARNING: Farm is enabled. One or more of the following is true: 1) One or more server(s) is either in WARNING or ERROR (not exceeding the predefined threshold) state. 2) The RDS Servers in this Farm present a mix of both known and unknown load preferences. * ERROR: Farm is enabled. One or more server(s) (exceeding the predefined threshold) is in ERROR state, or, for Automated Farms, there could be a provisioning error. * DISABLED: Farm is disabled. | [optional] 

## Methods

### NewFarmMonitorInfo

`func NewFarmMonitorInfo() *FarmMonitorInfo`

NewFarmMonitorInfo instantiates a new FarmMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmMonitorInfoWithDefaults

`func NewFarmMonitorInfoWithDefaults() *FarmMonitorInfo`

NewFarmMonitorInfoWithDefaults instantiates a new FarmMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationCount

`func (o *FarmMonitorInfo) GetApplicationCount() int32`

GetApplicationCount returns the ApplicationCount field if non-nil, zero value otherwise.

### GetApplicationCountOk

`func (o *FarmMonitorInfo) GetApplicationCountOk() (*int32, bool)`

GetApplicationCountOk returns a tuple with the ApplicationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationCount

`func (o *FarmMonitorInfo) SetApplicationCount(v int32)`

SetApplicationCount sets ApplicationCount field to given value.

### HasApplicationCount

`func (o *FarmMonitorInfo) HasApplicationCount() bool`

HasApplicationCount returns a boolean if a field has been set.

### GetDetails

`func (o *FarmMonitorInfo) GetDetails() FarmMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *FarmMonitorInfo) GetDetailsOk() (*FarmMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *FarmMonitorInfo) SetDetails(v FarmMonitorDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *FarmMonitorInfo) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *FarmMonitorInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FarmMonitorInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FarmMonitorInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FarmMonitorInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *FarmMonitorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FarmMonitorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FarmMonitorInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FarmMonitorInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRdsServerCount

`func (o *FarmMonitorInfo) GetRdsServerCount() int32`

GetRdsServerCount returns the RdsServerCount field if non-nil, zero value otherwise.

### GetRdsServerCountOk

`func (o *FarmMonitorInfo) GetRdsServerCountOk() (*int32, bool)`

GetRdsServerCountOk returns a tuple with the RdsServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdsServerCount

`func (o *FarmMonitorInfo) SetRdsServerCount(v int32)`

SetRdsServerCount sets RdsServerCount field to given value.

### HasRdsServerCount

`func (o *FarmMonitorInfo) HasRdsServerCount() bool`

HasRdsServerCount returns a boolean if a field has been set.

### GetStatus

`func (o *FarmMonitorInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FarmMonitorInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FarmMonitorInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FarmMonitorInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


