# RDSServerMonitorInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | [**RDSServerMonitorDetails**](RDSServerMonitorDetails.md) |  | 
**Enabled** | **bool** | Indicates if RDS server is enabled. | 
**FarmId** | **string** | Indicates the Farm ID to which the RDS Server belongs to. | 
**Id** | **string** | Unique ID of the RDS server. | 
**LoadIndex** | Pointer to **int32** | This value is similar to load_preference and represents the load on RDS Server in the range of 0 to 100. | [optional] 
**LoadPreference** | Pointer to **string** | Based on the current load of this RDS Server, gives a measure of how preferential this server is to be chosen for new application sessions. * BLOCK: RDS Server is blocked and new sessions will not be accepted. * HEAVY: RDS Server is experiencing heavy load and should likely not be chosen for new sessions. * NORMAL: RDS Server is experiencing normal load and is okay to be chosen for new sessions. * LIGHT: RDS Server is experiencing light load and is okay to be chosen for new sessions. * UNKNOWN: RDS Server did not report a load preference. This is potentially a configuration issue if other RDS Servers in the same Farm do report load preferences. | [optional] 
**Name** | **string** | RDS Server name. | 
**SessionCount** | Pointer to **int32** | RDS server session count. | [optional] 
**Status** | **string** | RDS server status. * OK: RDS Server is reachable. All applications (defined on its farm) are verified installed on the RDS Server. * WARNING: RDS Server is reachable. Some applications are detected as not installed on the RDS Server. * ERROR: RDS Server is unreachable, or, none of the applications are installed. * DISABLED: RDS Server is disabled. | 

## Methods

### NewRDSServerMonitorInfo

`func NewRDSServerMonitorInfo(details RDSServerMonitorDetails, enabled bool, farmId string, id string, name string, status string, ) *RDSServerMonitorInfo`

NewRDSServerMonitorInfo instantiates a new RDSServerMonitorInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRDSServerMonitorInfoWithDefaults

`func NewRDSServerMonitorInfoWithDefaults() *RDSServerMonitorInfo`

NewRDSServerMonitorInfoWithDefaults instantiates a new RDSServerMonitorInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *RDSServerMonitorInfo) GetDetails() RDSServerMonitorDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *RDSServerMonitorInfo) GetDetailsOk() (*RDSServerMonitorDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *RDSServerMonitorInfo) SetDetails(v RDSServerMonitorDetails)`

SetDetails sets Details field to given value.


### GetEnabled

`func (o *RDSServerMonitorInfo) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RDSServerMonitorInfo) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RDSServerMonitorInfo) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFarmId

`func (o *RDSServerMonitorInfo) GetFarmId() string`

GetFarmId returns the FarmId field if non-nil, zero value otherwise.

### GetFarmIdOk

`func (o *RDSServerMonitorInfo) GetFarmIdOk() (*string, bool)`

GetFarmIdOk returns a tuple with the FarmId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFarmId

`func (o *RDSServerMonitorInfo) SetFarmId(v string)`

SetFarmId sets FarmId field to given value.


### GetId

`func (o *RDSServerMonitorInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RDSServerMonitorInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RDSServerMonitorInfo) SetId(v string)`

SetId sets Id field to given value.


### GetLoadIndex

`func (o *RDSServerMonitorInfo) GetLoadIndex() int32`

GetLoadIndex returns the LoadIndex field if non-nil, zero value otherwise.

### GetLoadIndexOk

`func (o *RDSServerMonitorInfo) GetLoadIndexOk() (*int32, bool)`

GetLoadIndexOk returns a tuple with the LoadIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadIndex

`func (o *RDSServerMonitorInfo) SetLoadIndex(v int32)`

SetLoadIndex sets LoadIndex field to given value.

### HasLoadIndex

`func (o *RDSServerMonitorInfo) HasLoadIndex() bool`

HasLoadIndex returns a boolean if a field has been set.

### GetLoadPreference

`func (o *RDSServerMonitorInfo) GetLoadPreference() string`

GetLoadPreference returns the LoadPreference field if non-nil, zero value otherwise.

### GetLoadPreferenceOk

`func (o *RDSServerMonitorInfo) GetLoadPreferenceOk() (*string, bool)`

GetLoadPreferenceOk returns a tuple with the LoadPreference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadPreference

`func (o *RDSServerMonitorInfo) SetLoadPreference(v string)`

SetLoadPreference sets LoadPreference field to given value.

### HasLoadPreference

`func (o *RDSServerMonitorInfo) HasLoadPreference() bool`

HasLoadPreference returns a boolean if a field has been set.

### GetName

`func (o *RDSServerMonitorInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RDSServerMonitorInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RDSServerMonitorInfo) SetName(v string)`

SetName sets Name field to given value.


### GetSessionCount

`func (o *RDSServerMonitorInfo) GetSessionCount() int32`

GetSessionCount returns the SessionCount field if non-nil, zero value otherwise.

### GetSessionCountOk

`func (o *RDSServerMonitorInfo) GetSessionCountOk() (*int32, bool)`

GetSessionCountOk returns a tuple with the SessionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionCount

`func (o *RDSServerMonitorInfo) SetSessionCount(v int32)`

SetSessionCount sets SessionCount field to given value.

### HasSessionCount

`func (o *RDSServerMonitorInfo) HasSessionCount() bool`

HasSessionCount returns a boolean if a field has been set.

### GetStatus

`func (o *RDSServerMonitorInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RDSServerMonitorInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RDSServerMonitorInfo) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


