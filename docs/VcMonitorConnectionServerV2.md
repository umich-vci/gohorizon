# VcMonitorConnectionServerV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**CertificateMonitorInfo**](CertificateMonitorInfo.md) |  | [optional] 
**Id** | **string** | Unique ID of the Connection Server. | 
**LastUpdatedTimestamp** | **int64** | The timestamp in milliseconds when the last update was obtained. Measured as epoch time. | [optional] 
**Name** | **string** | Connection server host name or IP address. | 
**Status** | **string** | Status of the Virtual Center Connection with respect to this Connection Server. * OK: The connection to Virtual Center server is working properly. * DOWN: The connection to Virtual Center server is down. * RECONNECTING: The connection to Virtual Center server was lost and is being reconnected to. * UNKNOWN: Connection state to Virtual Center server is unknown. * INVALID_CREDENTIALS: The supplied credentials cannot be used to authenticate to the Virtual Center server. * CANNOT_LOGIN: The connection server cannot login to the Virtual Center server. * NOT_YET_CONNECTED: Connection server has not yet connected to Virtual Center server. | 
**ThumbprintAccepted** | **bool** | Indicates if the thumbprints of the Virtual Center was accepted. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


