# AuditEventSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationPoolName** | Pointer to **string** | Application Pool associated with this event. Will be unset if there is no application association for this event. Supported Filters : &#39;Equals&#39;. | [optional] 
**DesktopPoolName** | Pointer to **string** | Desktop Pool associated with this event. Will be unset if there is no desktop association for this event. Supported Filters : &#39;Equals&#39;. | [optional] 
**Id** | Pointer to **int64** | Unique id representing an event. Supported Filters : &#39;Equals&#39;. | [optional] 
**MachineDnsName** | Pointer to **string** | FQDN of the machine in the Pod that has logged this event. Supported Filters : &#39;Equals&#39;. | [optional] 
**MachineId** | Pointer to **string** | Machine associated with this event. Will be unset if there is no machine association for this event. Supported Filters : &#39;Equals&#39;. | [optional] 
**Message** | Pointer to **string** | Audit event message.  | [optional] 
**Module** | Pointer to **string** | Horizon component that has logged this event. Supported Filters : &#39;Equals&#39;. | [optional] 
**Severity** | Pointer to **string** | Severity type of the event. Supported Filters : &#39;Equals&#39;. * INFO: Audit event is of INFO severity. * WARNING: Audit event is of WARNING severity * ERROR: Audit event is of ERROR severity * AUDIT_SUCCESS: Audit event is of AUDIT_SUCCESS severity * AUDIT_FAIL: Audit event is of AUDIT_FAIL severity * UNKNOWN: Not able to identify severity | [optional] 
**Time** | Pointer to **int64** | Time at which the event occurred. Supported Filters : &#39;Between&#39;. | [optional] 
**Type** | Pointer to **string** | Event name that corresponds to an item in the message catalog. Supported Filters : &#39;Equals&#39;. | [optional] 
**UserId** | Pointer to **string** | Sid of the user associated with this event. Supported Filters : &#39;Equals&#39;. | [optional] 

## Methods

### NewAuditEventSummary

`func NewAuditEventSummary() *AuditEventSummary`

NewAuditEventSummary instantiates a new AuditEventSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditEventSummaryWithDefaults

`func NewAuditEventSummaryWithDefaults() *AuditEventSummary`

NewAuditEventSummaryWithDefaults instantiates a new AuditEventSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationPoolName

`func (o *AuditEventSummary) GetApplicationPoolName() string`

GetApplicationPoolName returns the ApplicationPoolName field if non-nil, zero value otherwise.

### GetApplicationPoolNameOk

`func (o *AuditEventSummary) GetApplicationPoolNameOk() (*string, bool)`

GetApplicationPoolNameOk returns a tuple with the ApplicationPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationPoolName

`func (o *AuditEventSummary) SetApplicationPoolName(v string)`

SetApplicationPoolName sets ApplicationPoolName field to given value.

### HasApplicationPoolName

`func (o *AuditEventSummary) HasApplicationPoolName() bool`

HasApplicationPoolName returns a boolean if a field has been set.

### GetDesktopPoolName

`func (o *AuditEventSummary) GetDesktopPoolName() string`

GetDesktopPoolName returns the DesktopPoolName field if non-nil, zero value otherwise.

### GetDesktopPoolNameOk

`func (o *AuditEventSummary) GetDesktopPoolNameOk() (*string, bool)`

GetDesktopPoolNameOk returns a tuple with the DesktopPoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesktopPoolName

`func (o *AuditEventSummary) SetDesktopPoolName(v string)`

SetDesktopPoolName sets DesktopPoolName field to given value.

### HasDesktopPoolName

`func (o *AuditEventSummary) HasDesktopPoolName() bool`

HasDesktopPoolName returns a boolean if a field has been set.

### GetId

`func (o *AuditEventSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuditEventSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuditEventSummary) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AuditEventSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMachineDnsName

`func (o *AuditEventSummary) GetMachineDnsName() string`

GetMachineDnsName returns the MachineDnsName field if non-nil, zero value otherwise.

### GetMachineDnsNameOk

`func (o *AuditEventSummary) GetMachineDnsNameOk() (*string, bool)`

GetMachineDnsNameOk returns a tuple with the MachineDnsName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineDnsName

`func (o *AuditEventSummary) SetMachineDnsName(v string)`

SetMachineDnsName sets MachineDnsName field to given value.

### HasMachineDnsName

`func (o *AuditEventSummary) HasMachineDnsName() bool`

HasMachineDnsName returns a boolean if a field has been set.

### GetMachineId

`func (o *AuditEventSummary) GetMachineId() string`

GetMachineId returns the MachineId field if non-nil, zero value otherwise.

### GetMachineIdOk

`func (o *AuditEventSummary) GetMachineIdOk() (*string, bool)`

GetMachineIdOk returns a tuple with the MachineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineId

`func (o *AuditEventSummary) SetMachineId(v string)`

SetMachineId sets MachineId field to given value.

### HasMachineId

`func (o *AuditEventSummary) HasMachineId() bool`

HasMachineId returns a boolean if a field has been set.

### GetMessage

`func (o *AuditEventSummary) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuditEventSummary) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuditEventSummary) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuditEventSummary) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetModule

`func (o *AuditEventSummary) GetModule() string`

GetModule returns the Module field if non-nil, zero value otherwise.

### GetModuleOk

`func (o *AuditEventSummary) GetModuleOk() (*string, bool)`

GetModuleOk returns a tuple with the Module field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModule

`func (o *AuditEventSummary) SetModule(v string)`

SetModule sets Module field to given value.

### HasModule

`func (o *AuditEventSummary) HasModule() bool`

HasModule returns a boolean if a field has been set.

### GetSeverity

`func (o *AuditEventSummary) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *AuditEventSummary) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *AuditEventSummary) SetSeverity(v string)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *AuditEventSummary) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetTime

`func (o *AuditEventSummary) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *AuditEventSummary) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *AuditEventSummary) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *AuditEventSummary) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *AuditEventSummary) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditEventSummary) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditEventSummary) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AuditEventSummary) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *AuditEventSummary) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AuditEventSummary) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AuditEventSummary) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AuditEventSummary) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


