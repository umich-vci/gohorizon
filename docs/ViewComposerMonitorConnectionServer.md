# ViewComposerMonitorConnectionServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**CertificateMonitorInfo**](CertificateMonitorInfo.md) |  | [optional] 
**Id** | **string** | Unique ID of the Connection Server. | 
**Name** | **string** | Connection server host name or IP address. | 
**Status** | **string** | Status of the View Composer with respect to this Connection Server. * OK: The connection to View Composer server is working properly. * MALFORMED_URL: The connection to View Composer server was not possible due to a malformed URL. * ERROR: Error occurred when connecting to View Composer server. * CERT_ERROR: Certificate validation error when connecting to the View Composer server. | 
**ThumbprintAccepted** | **bool** | Indicates if the thumbprint of the View Composer was accepted. | 

## Methods

### NewViewComposerMonitorConnectionServer

`func NewViewComposerMonitorConnectionServer(id string, name string, status string, thumbprintAccepted bool, ) *ViewComposerMonitorConnectionServer`

NewViewComposerMonitorConnectionServer instantiates a new ViewComposerMonitorConnectionServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewViewComposerMonitorConnectionServerWithDefaults

`func NewViewComposerMonitorConnectionServerWithDefaults() *ViewComposerMonitorConnectionServer`

NewViewComposerMonitorConnectionServerWithDefaults instantiates a new ViewComposerMonitorConnectionServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *ViewComposerMonitorConnectionServer) GetCertificate() CertificateMonitorInfo`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ViewComposerMonitorConnectionServer) GetCertificateOk() (*CertificateMonitorInfo, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ViewComposerMonitorConnectionServer) SetCertificate(v CertificateMonitorInfo)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *ViewComposerMonitorConnectionServer) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetId

`func (o *ViewComposerMonitorConnectionServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ViewComposerMonitorConnectionServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ViewComposerMonitorConnectionServer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ViewComposerMonitorConnectionServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ViewComposerMonitorConnectionServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ViewComposerMonitorConnectionServer) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ViewComposerMonitorConnectionServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ViewComposerMonitorConnectionServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ViewComposerMonitorConnectionServer) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetThumbprintAccepted

`func (o *ViewComposerMonitorConnectionServer) GetThumbprintAccepted() bool`

GetThumbprintAccepted returns the ThumbprintAccepted field if non-nil, zero value otherwise.

### GetThumbprintAcceptedOk

`func (o *ViewComposerMonitorConnectionServer) GetThumbprintAcceptedOk() (*bool, bool)`

GetThumbprintAcceptedOk returns a tuple with the ThumbprintAccepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprintAccepted

`func (o *ViewComposerMonitorConnectionServer) SetThumbprintAccepted(v bool)`

SetThumbprintAccepted sets ThumbprintAccepted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


