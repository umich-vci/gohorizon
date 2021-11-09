# ViewComposerMonitorConnectionServerV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**CertificateMonitorInfo**](CertificateMonitorInfo.md) |  | [optional] 
**Id** | **string** | Unique ID of the Connection Server. | 
**LastUpdatedTimestamp** | Pointer to **int64** | The timestamp in milliseconds when the last update was obtained. Measured as epoch time. | [optional] 
**Name** | **string** | Connection server host name or IP address. | 
**Status** | **string** | Status of the View Composer with respect to this Connection Server. * OK: The connection to View Composer server is working properly. * MALFORMED_URL: The connection to View Composer server was not possible due to a malformed URL. * ERROR: Error occurred when connecting to View Composer server. * CERT_ERROR: Certificate validation error when connecting to the View Composer server. | 
**ThumbprintAccepted** | **bool** | Indicates if the thumbprint of the View Composer was accepted. | 

## Methods

### NewViewComposerMonitorConnectionServerV2

`func NewViewComposerMonitorConnectionServerV2(id string, name string, status string, thumbprintAccepted bool, ) *ViewComposerMonitorConnectionServerV2`

NewViewComposerMonitorConnectionServerV2 instantiates a new ViewComposerMonitorConnectionServerV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewComposerMonitorConnectionServerV2WithDefaults

`func NewViewComposerMonitorConnectionServerV2WithDefaults() *ViewComposerMonitorConnectionServerV2`

NewViewComposerMonitorConnectionServerV2WithDefaults instantiates a new ViewComposerMonitorConnectionServerV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *ViewComposerMonitorConnectionServerV2) GetCertificate() CertificateMonitorInfo`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ViewComposerMonitorConnectionServerV2) GetCertificateOk() (*CertificateMonitorInfo, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ViewComposerMonitorConnectionServerV2) SetCertificate(v CertificateMonitorInfo)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ViewComposerMonitorConnectionServerV2) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetId

`func (o *ViewComposerMonitorConnectionServerV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewComposerMonitorConnectionServerV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewComposerMonitorConnectionServerV2) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdatedTimestamp

`func (o *ViewComposerMonitorConnectionServerV2) GetLastUpdatedTimestamp() int64`

GetLastUpdatedTimestamp returns the LastUpdatedTimestamp field if non-nil, zero value otherwise.

### GetLastUpdatedTimestampOk

`func (o *ViewComposerMonitorConnectionServerV2) GetLastUpdatedTimestampOk() (*int64, bool)`

GetLastUpdatedTimestampOk returns a tuple with the LastUpdatedTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedTimestamp

`func (o *ViewComposerMonitorConnectionServerV2) SetLastUpdatedTimestamp(v int64)`

SetLastUpdatedTimestamp sets LastUpdatedTimestamp field to given value.

### HasLastUpdatedTimestamp

`func (o *ViewComposerMonitorConnectionServerV2) HasLastUpdatedTimestamp() bool`

HasLastUpdatedTimestamp returns a boolean if a field has been set.

### GetName

`func (o *ViewComposerMonitorConnectionServerV2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewComposerMonitorConnectionServerV2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewComposerMonitorConnectionServerV2) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ViewComposerMonitorConnectionServerV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewComposerMonitorConnectionServerV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewComposerMonitorConnectionServerV2) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetThumbprintAccepted

`func (o *ViewComposerMonitorConnectionServerV2) GetThumbprintAccepted() bool`

GetThumbprintAccepted returns the ThumbprintAccepted field if non-nil, zero value otherwise.

### GetThumbprintAcceptedOk

`func (o *ViewComposerMonitorConnectionServerV2) GetThumbprintAcceptedOk() (*bool, bool)`

GetThumbprintAcceptedOk returns a tuple with the ThumbprintAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprintAccepted

`func (o *ViewComposerMonitorConnectionServerV2) SetThumbprintAccepted(v bool)`

SetThumbprintAccepted sets ThumbprintAccepted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


