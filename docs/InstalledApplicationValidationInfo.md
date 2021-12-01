# InstalledApplicationValidationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationExePath** | Pointer to **string** | Application executable path. | [optional] 
**MachineOrRdsServerIds** | Pointer to **[]string** | Machine/RDS server ids on which this application has not been installed on. This will be populated only when application is not installed on one or more RDS servers/machines. | [optional] 
**Status** | Pointer to **string** | Status of application installation on all RDS servers/machines. * SUCCESS: The given application is installed on all RDS Servers/machines of a given farm/desktop pool. * NOT_INSTALLED: The given application is not installed on one or more RDS Servers/machines of a given farm/desktop pool. | [optional] 

## Methods

### NewInstalledApplicationValidationInfo

`func NewInstalledApplicationValidationInfo() *InstalledApplicationValidationInfo`

NewInstalledApplicationValidationInfo instantiates a new InstalledApplicationValidationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstalledApplicationValidationInfoWithDefaults

`func NewInstalledApplicationValidationInfoWithDefaults() *InstalledApplicationValidationInfo`

NewInstalledApplicationValidationInfoWithDefaults instantiates a new InstalledApplicationValidationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationExePath

`func (o *InstalledApplicationValidationInfo) GetApplicationExePath() string`

GetApplicationExePath returns the ApplicationExePath field if non-nil, zero value otherwise.

### GetApplicationExePathOk

`func (o *InstalledApplicationValidationInfo) GetApplicationExePathOk() (*string, bool)`

GetApplicationExePathOk returns a tuple with the ApplicationExePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationExePath

`func (o *InstalledApplicationValidationInfo) SetApplicationExePath(v string)`

SetApplicationExePath sets ApplicationExePath field to given value.

### HasApplicationExePath

`func (o *InstalledApplicationValidationInfo) HasApplicationExePath() bool`

HasApplicationExePath returns a boolean if a field has been set.

### GetMachineOrRdsServerIds

`func (o *InstalledApplicationValidationInfo) GetMachineOrRdsServerIds() []string`

GetMachineOrRdsServerIds returns the MachineOrRdsServerIds field if non-nil, zero value otherwise.

### GetMachineOrRdsServerIdsOk

`func (o *InstalledApplicationValidationInfo) GetMachineOrRdsServerIdsOk() (*[]string, bool)`

GetMachineOrRdsServerIdsOk returns a tuple with the MachineOrRdsServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineOrRdsServerIds

`func (o *InstalledApplicationValidationInfo) SetMachineOrRdsServerIds(v []string)`

SetMachineOrRdsServerIds sets MachineOrRdsServerIds field to given value.

### HasMachineOrRdsServerIds

`func (o *InstalledApplicationValidationInfo) HasMachineOrRdsServerIds() bool`

HasMachineOrRdsServerIds returns a boolean if a field has been set.

### GetStatus

`func (o *InstalledApplicationValidationInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InstalledApplicationValidationInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InstalledApplicationValidationInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InstalledApplicationValidationInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


