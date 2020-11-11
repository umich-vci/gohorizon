# \MonitorApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConnectionServerMonitorInfoV2**](MonitorApi.md#GetConnectionServerMonitorInfoV2) | **Get** /monitor/v1/connection-servers/{id} | Gets monitoring information related to Connection Server.
[**GetEventDatabaseMonitor**](MonitorApi.md#GetEventDatabaseMonitor) | **Get** /monitor/event-database | Returns monitoring information related to Event database of the environment.
[**GetFarmMonitorInfo**](MonitorApi.md#GetFarmMonitorInfo) | **Get** /monitor/v1/farms/{id} | Gets monitoring information related to farm.
[**GetGatewayMonitorInfo**](MonitorApi.md#GetGatewayMonitorInfo) | **Get** /monitor/v1/gateways/{id} | Gets monitoring information related to a Gateway.
[**GetPodMonitorInfoV2**](MonitorApi.md#GetPodMonitorInfoV2) | **Get** /monitor/v1/pods/{id} | Gets monitoring information related to the remote pod.
[**GetRDSServerMonitors**](MonitorApi.md#GetRDSServerMonitors) | **Get** /monitor/v1/rds-servers/{id} | Gets monitoring information related to RDS Server.
[**GetSAMLAuthenticatorMonitorInfo**](MonitorApi.md#GetSAMLAuthenticatorMonitorInfo) | **Get** /monitor/v1/saml-authenticators/{id} | Gets Monitoring Information related to a SAML Authenticator
[**GetTrueSSOMonitorInfo**](MonitorApi.md#GetTrueSSOMonitorInfo) | **Get** /monitor/v1/true-sso/{id} | Gets monitoring information related to a True SSO connector.
[**GetViewComposerByVCId**](MonitorApi.md#GetViewComposerByVCId) | **Get** /monitor/v1/view-composers/{vcId} | Gets monitoring information of view composer for a given virtual center
[**GetVirtualCenterMonitorInfo**](MonitorApi.md#GetVirtualCenterMonitorInfo) | **Get** /monitor/v1/virtual-centers/{id} | Gets monitoring information related to Virtual Center.
[**ListADDomainMonitorInfosV2**](MonitorApi.md#ListADDomainMonitorInfosV2) | **Get** /monitor/v2/ad-domains | Lists monitoring information related to AD Domains of the environment.
[**ListADDomainMonitors**](MonitorApi.md#ListADDomainMonitors) | **Get** /monitor/ad-domains | Lists monitoring information related to AD Domains of the environment.
[**ListConnectionServerMonitors**](MonitorApi.md#ListConnectionServerMonitors) | **Get** /monitor/connection-servers | Lists monitoring information related to Connection Servers of the environment.
[**ListConnectionServerMonitorsV2**](MonitorApi.md#ListConnectionServerMonitorsV2) | **Get** /monitor/v2/connection-servers | Lists monitoring information related to Connection Servers of the environment.
[**ListFarmMonitors**](MonitorApi.md#ListFarmMonitors) | **Get** /monitor/farms | Lists monitoring information related to Farms of the environment.
[**ListGatewayMonitorInfoV1**](MonitorApi.md#ListGatewayMonitorInfoV1) | **Get** /monitor/gateways | Lists monitoring information related to Gateways registered in the environment.
[**ListGatewayMonitorInfoV2**](MonitorApi.md#ListGatewayMonitorInfoV2) | **Get** /monitor/v2/gateways | Lists monitoring information related to Gateways registered in the environment.
[**ListPodMonitorInfosV1**](MonitorApi.md#ListPodMonitorInfosV1) | **Get** /monitor/v1/pods | Lists monitoring information related to the remote pods.
[**ListPodMonitorInfosV2**](MonitorApi.md#ListPodMonitorInfosV2) | **Get** /monitor/v2/pods | Lists monitoring information related to the remote pods.
[**ListRDSServerMonitors**](MonitorApi.md#ListRDSServerMonitors) | **Get** /monitor/rds-servers | Lists monitoring information related to RDS Servers of the environment.
[**ListSAMLAuthenticatorMonitorsV1**](MonitorApi.md#ListSAMLAuthenticatorMonitorsV1) | **Get** /monitor/saml-authenticators | Lists monitoring information related to SAML Authenticators of the environment.
[**ListSAMLAuthenticatorMonitorsV2**](MonitorApi.md#ListSAMLAuthenticatorMonitorsV2) | **Get** /monitor/v2/saml-authenticators | Lists monitoring information related to SAML Authenticators of the environment.
[**ListTrueSSOMonitorInfos**](MonitorApi.md#ListTrueSSOMonitorInfos) | **Get** /monitor/v1/true-sso | Lists monitoring information related to True SSO connectors.
[**ListViewComposerMonitorsV1**](MonitorApi.md#ListViewComposerMonitorsV1) | **Get** /monitor/view-composers | Lists monitoring information related to View Composers of the environment.
[**ListViewComposerMonitorsV2**](MonitorApi.md#ListViewComposerMonitorsV2) | **Get** /monitor/v2/view-composers | Lists monitoring information related to View Composers of the environment.
[**ListVirtualCenterMonitors**](MonitorApi.md#ListVirtualCenterMonitors) | **Get** /monitor/virtual-centers | Lists monitoring information related to Virtual Centers of the environment.
[**ListVirtualCenterMonitorsV2**](MonitorApi.md#ListVirtualCenterMonitorsV2) | **Get** /monitor/v2/virtual-centers | Lists monitoring information related to Virtual Centers of the environment.



## GetConnectionServerMonitorInfoV2

> ConnectionServerMonitorInfoV2 GetConnectionServerMonitorInfoV2(ctx, id)

Gets monitoring information related to Connection Server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**ConnectionServerMonitorInfoV2**](ConnectionServerMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventDatabaseMonitor

> EventDatabaseMonitorInfo GetEventDatabaseMonitor(ctx, )

Returns monitoring information related to Event database of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**EventDatabaseMonitorInfo**](EventDatabaseMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFarmMonitorInfo

> FarmMonitorInfo GetFarmMonitorInfo(ctx, id)

Gets monitoring information related to farm.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**FarmMonitorInfo**](FarmMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayMonitorInfo

> GatewayMonitorInfoV2 GetGatewayMonitorInfo(ctx, id)

Gets monitoring information related to a Gateway.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**GatewayMonitorInfoV2**](GatewayMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPodMonitorInfoV2

> PodMonitorInfoV2 GetPodMonitorInfoV2(ctx, id)

Gets monitoring information related to the remote pod.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**PodMonitorInfoV2**](PodMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRDSServerMonitors

> RdsServerMonitorInfo GetRDSServerMonitors(ctx, id)

Gets monitoring information related to RDS Server.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**RdsServerMonitorInfo**](RDSServerMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSAMLAuthenticatorMonitorInfo

> SamlAuthenticatorMonitorInfoV2 GetSAMLAuthenticatorMonitorInfo(ctx, id)

Gets Monitoring Information related to a SAML Authenticator

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**SamlAuthenticatorMonitorInfoV2**](SAMLAuthenticatorMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrueSSOMonitorInfo

> TrueSsoMonitorInfo GetTrueSSOMonitorInfo(ctx, id)

Gets monitoring information related to a True SSO connector.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**TrueSsoMonitorInfo**](TrueSSOMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewComposerByVCId

> ViewComposerMonitorInfoV2 GetViewComposerByVCId(ctx, vcId)

Gets monitoring information of view composer for a given virtual center

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vcId** | **string**| vcId | 

### Return type

[**ViewComposerMonitorInfoV2**](ViewComposerMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVirtualCenterMonitorInfo

> VirtualCenterMonitorInfoV2 GetVirtualCenterMonitorInfo(ctx, id)

Gets monitoring information related to Virtual Center.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**VirtualCenterMonitorInfoV2**](VirtualCenterMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADDomainMonitorInfosV2

> []AdDomainMonitorInfoV2 ListADDomainMonitorInfosV2(ctx, )

Lists monitoring information related to AD Domains of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AdDomainMonitorInfoV2**](ADDomainMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADDomainMonitors

> []AdDomainMonitorInfo ListADDomainMonitors(ctx, )

Lists monitoring information related to AD Domains of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]AdDomainMonitorInfo**](ADDomainMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionServerMonitors

> []ConnectionServerMonitorInfo ListConnectionServerMonitors(ctx, )

Lists monitoring information related to Connection Servers of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ConnectionServerMonitorInfo**](ConnectionServerMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionServerMonitorsV2

> []ConnectionServerMonitorInfoV2 ListConnectionServerMonitorsV2(ctx, )

Lists monitoring information related to Connection Servers of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ConnectionServerMonitorInfoV2**](ConnectionServerMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFarmMonitors

> []FarmMonitorInfo ListFarmMonitors(ctx, )

Lists monitoring information related to Farms of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]FarmMonitorInfo**](FarmMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGatewayMonitorInfoV1

> []GatewayMonitorInfo ListGatewayMonitorInfoV1(ctx, )

Lists monitoring information related to Gateways registered in the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]GatewayMonitorInfo**](GatewayMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGatewayMonitorInfoV2

> []GatewayMonitorInfoV2 ListGatewayMonitorInfoV2(ctx, )

Lists monitoring information related to Gateways registered in the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]GatewayMonitorInfoV2**](GatewayMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPodMonitorInfosV1

> []PodMonitorInfo ListPodMonitorInfosV1(ctx, )

Lists monitoring information related to the remote pods.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PodMonitorInfo**](PodMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPodMonitorInfosV2

> []PodMonitorInfoV2 ListPodMonitorInfosV2(ctx, )

Lists monitoring information related to the remote pods.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]PodMonitorInfoV2**](PodMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRDSServerMonitors

> []RdsServerMonitorInfo ListRDSServerMonitors(ctx, )

Lists monitoring information related to RDS Servers of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]RdsServerMonitorInfo**](RDSServerMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSAMLAuthenticatorMonitorsV1

> []SamlAuthenticatorMonitorInfo ListSAMLAuthenticatorMonitorsV1(ctx, )

Lists monitoring information related to SAML Authenticators of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SamlAuthenticatorMonitorInfo**](SAMLAuthenticatorMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSAMLAuthenticatorMonitorsV2

> []SamlAuthenticatorMonitorInfoV2 ListSAMLAuthenticatorMonitorsV2(ctx, )

Lists monitoring information related to SAML Authenticators of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]SamlAuthenticatorMonitorInfoV2**](SAMLAuthenticatorMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrueSSOMonitorInfos

> []TrueSsoMonitorInfo ListTrueSSOMonitorInfos(ctx, )

Lists monitoring information related to True SSO connectors.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]TrueSsoMonitorInfo**](TrueSSOMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListViewComposerMonitorsV1

> []ViewComposerMonitorInfo ListViewComposerMonitorsV1(ctx, )

Lists monitoring information related to View Composers of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ViewComposerMonitorInfo**](ViewComposerMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListViewComposerMonitorsV2

> []ViewComposerMonitorInfoV2 ListViewComposerMonitorsV2(ctx, )

Lists monitoring information related to View Composers of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]ViewComposerMonitorInfoV2**](ViewComposerMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualCenterMonitors

> []VirtualCenterMonitorInfo ListVirtualCenterMonitors(ctx, )

Lists monitoring information related to Virtual Centers of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]VirtualCenterMonitorInfo**](VirtualCenterMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListVirtualCenterMonitorsV2

> []VirtualCenterMonitorInfoV2 ListVirtualCenterMonitorsV2(ctx, )

Lists monitoring information related to Virtual Centers of the environment.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]VirtualCenterMonitorInfoV2**](VirtualCenterMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

