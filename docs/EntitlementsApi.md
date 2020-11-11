# \EntitlementsApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateApplicationPoolEntitlements**](EntitlementsApi.md#BulkCreateApplicationPoolEntitlements) | **Post** /entitlements/v1/application-pools | Create the bulk entitlements for a set of application pools
[**BulkCreateDesktopPoolEntitlements**](EntitlementsApi.md#BulkCreateDesktopPoolEntitlements) | **Post** /entitlements/v1/desktop-pools | Create the bulk entitlements for a set of desktop pools
[**BulkDeleteApplicationPoolEntitlements**](EntitlementsApi.md#BulkDeleteApplicationPoolEntitlements) | **Delete** /entitlements/v1/application-pools | Delete the bulk entitlements for a set of application pools
[**BulkDeleteDesktopPoolEntitlements**](EntitlementsApi.md#BulkDeleteDesktopPoolEntitlements) | **Delete** /entitlements/v1/desktop-pools | Delete the bulk entitlements for a set of desktop pools
[**GetApplicationPoolEntitlements**](EntitlementsApi.md#GetApplicationPoolEntitlements) | **Get** /entitlements/v1/application-pools/{id} | Returns the IDs of users or groups entitled to a given application pool.
[**GetDesktopPoolEntitlements**](EntitlementsApi.md#GetDesktopPoolEntitlements) | **Get** /entitlements/v1/desktop-pools/{id} | Returns the IDs of users or groups entitled to a given desktop pool.
[**ListApplicationPoolEntitlements**](EntitlementsApi.md#ListApplicationPoolEntitlements) | **Get** /entitlements/v1/application-pools | Lists the entitlements for Application Pools in the environment.
[**ListDesktopPoolEntitlements**](EntitlementsApi.md#ListDesktopPoolEntitlements) | **Get** /entitlements/v1/desktop-pools | Lists the entitlements for Desktop Pools in the environment.



## BulkCreateApplicationPoolEntitlements

> []BulkEntitlementResponseInfo BulkCreateApplicationPoolEntitlements(ctx, body)

Create the bulk entitlements for a set of application pools

The input spec must not contain duplicate Ids.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]EntitlementSpec**](EntitlementSpec.md)| Specifications for bulk application entitlements to be created. | 

### Return type

[**[]BulkEntitlementResponseInfo**](BulkEntitlementResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkCreateDesktopPoolEntitlements

> []BulkEntitlementResponseInfo BulkCreateDesktopPoolEntitlements(ctx, body)

Create the bulk entitlements for a set of desktop pools

The input spec must not contain duplicate Ids.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]EntitlementSpec**](EntitlementSpec.md)| Specifications for bulk desktop entitlements to be created. | 

### Return type

[**[]BulkEntitlementResponseInfo**](BulkEntitlementResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteApplicationPoolEntitlements

> []BulkEntitlementResponseInfo BulkDeleteApplicationPoolEntitlements(ctx, body)

Delete the bulk entitlements for a set of application pools

The input spec must not contain duplicate Ids.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]EntitlementSpec**](EntitlementSpec.md)| Specifications for bulk application entitlements to be deleted. | 

### Return type

[**[]BulkEntitlementResponseInfo**](BulkEntitlementResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteDesktopPoolEntitlements

> []BulkEntitlementResponseInfo BulkDeleteDesktopPoolEntitlements(ctx, body)

Delete the bulk entitlements for a set of desktop pools

The input spec must not contain duplicate Ids.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**[]EntitlementSpec**](EntitlementSpec.md)| Specifications for bulk desktop entitlements to be deleted. | 

### Return type

[**[]BulkEntitlementResponseInfo**](BulkEntitlementResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationPoolEntitlements

> EntitlementInfo GetApplicationPoolEntitlements(ctx, id)

Returns the IDs of users or groups entitled to a given application pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**EntitlementInfo**](EntitlementInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPoolEntitlements

> EntitlementInfo GetDesktopPoolEntitlements(ctx, id)

Returns the IDs of users or groups entitled to a given desktop pool.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| id | 

### Return type

[**EntitlementInfo**](EntitlementInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationPoolEntitlements

> []EntitlementInfo ListApplicationPoolEntitlements(ctx, )

Lists the entitlements for Application Pools in the environment.

This API supports <b>Pagination</b> and <b>Filters.</b><br/>For Pagination, optional query params of 'page' and 'size' need to be sent.<br/>For Filters, refer to 'EntitlementInfo' model description to find supported filters on specific field.<br/>For full information on using Filters, refer to 'Horizon Server REST Pagination and Filter Guide' of 'VMware Horizon Server API' in documentation.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EntitlementInfo**](EntitlementInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPoolEntitlements

> []EntitlementInfo ListDesktopPoolEntitlements(ctx, )

Lists the entitlements for Desktop Pools in the environment.

This API supports <b>Pagination</b> and <b>Filters.</b><br/>For Pagination, optional query params of 'page' and 'size' need to be sent.<br/>For Filters, refer to 'EntitlementInfo' model description to find supported filters on specific field.<br/>For full information on using Filters, refer to 'Horizon Server REST Pagination and Filter Guide' of 'VMware Horizon Server API' in documentation.

### Required Parameters

This endpoint does not need any parameter.

### Return type

[**[]EntitlementInfo**](EntitlementInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

