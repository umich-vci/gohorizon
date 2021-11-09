# FarmRDSServersPatternNamingSettingsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxNumberOfRdsServers** | Pointer to **int32** | Maximum number of RDS Servers in the farm. | [optional] 
**NamingPattern** | Pointer to **string** | RDS Servers will be named according to the specified naming pattern. By default, Horizon appends a unique number to the specified pattern to provide a unique name for each RDS Server. To place this unique number elsewhere in the pattern, use &#39;{n}&#39;. (For example: rds-{n}-sales.) The unique number can also be made a fixed length. (For example: rds-{n:fixed&#x3D;3}-sales will name RDS Servers from rds-001-sales to rds-999-sales). RDS Server names are constrained to a maximum size of 15 characters including the unique number. Therefore, care must be taken when choosing a pattern. If the maximum farm size is 9 RDS servers, the pattern must be at most 14 characters. For 99 RDS servers, 13 characters, for 999 RDS servers, 12 characters. For 9999 RDS servers, 11 characters. If using a fixed size token, use a maximum of 14 characters for \&quot;n&#x3D;1\&quot;, 13 characters for \&quot;n&#x3D;2\&quot;, 12 characters for \&quot;n&#x3D;3\&quot;, and 11 characters for \&quot;n&#x3D;4\&quot;. If {n} is specified with no size, a size of 2 is automatically used and if no {} is specified, {n&#x3D;2} is automatically appended to the end of the pattern. If the numeric placement token not present, then maximum length is 14 characters.&lt;br&gt;Supported Filters: &#39;Equals&#39;. | [optional] 

## Methods

### NewFarmRDSServersPatternNamingSettingsInfo

`func NewFarmRDSServersPatternNamingSettingsInfo() *FarmRDSServersPatternNamingSettingsInfo`

NewFarmRDSServersPatternNamingSettingsInfo instantiates a new FarmRDSServersPatternNamingSettingsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFarmRDSServersPatternNamingSettingsInfoWithDefaults

`func NewFarmRDSServersPatternNamingSettingsInfoWithDefaults() *FarmRDSServersPatternNamingSettingsInfo`

NewFarmRDSServersPatternNamingSettingsInfoWithDefaults instantiates a new FarmRDSServersPatternNamingSettingsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxNumberOfRdsServers

`func (o *FarmRDSServersPatternNamingSettingsInfo) GetMaxNumberOfRdsServers() int32`

GetMaxNumberOfRdsServers returns the MaxNumberOfRdsServers field if non-nil, zero value otherwise.

### GetMaxNumberOfRdsServersOk

`func (o *FarmRDSServersPatternNamingSettingsInfo) GetMaxNumberOfRdsServersOk() (*int32, bool)`

GetMaxNumberOfRdsServersOk returns a tuple with the MaxNumberOfRdsServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNumberOfRdsServers

`func (o *FarmRDSServersPatternNamingSettingsInfo) SetMaxNumberOfRdsServers(v int32)`

SetMaxNumberOfRdsServers sets MaxNumberOfRdsServers field to given value.

### HasMaxNumberOfRdsServers

`func (o *FarmRDSServersPatternNamingSettingsInfo) HasMaxNumberOfRdsServers() bool`

HasMaxNumberOfRdsServers returns a boolean if a field has been set.

### GetNamingPattern

`func (o *FarmRDSServersPatternNamingSettingsInfo) GetNamingPattern() string`

GetNamingPattern returns the NamingPattern field if non-nil, zero value otherwise.

### GetNamingPatternOk

`func (o *FarmRDSServersPatternNamingSettingsInfo) GetNamingPatternOk() (*string, bool)`

GetNamingPatternOk returns a tuple with the NamingPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamingPattern

`func (o *FarmRDSServersPatternNamingSettingsInfo) SetNamingPattern(v string)`

SetNamingPattern sets NamingPattern field to given value.

### HasNamingPattern

`func (o *FarmRDSServersPatternNamingSettingsInfo) HasNamingPattern() bool`

HasNamingPattern returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


