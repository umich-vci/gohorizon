# \AuthApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**LoginUser**](AuthApi.md#LoginUser) | **Post** /login | Logs in a user. Returns access token and refresh token.
[**LogoutUser**](AuthApi.md#LogoutUser) | **Post** /logout | Logs out a user.
[**RefreshAccessToken**](AuthApi.md#RefreshAccessToken) | **Post** /refresh | Refreshes access token from refresh token.



## LoginUser

> AuthTokens LoginUser(ctx, body)

Logs in a user. Returns access token and refresh token.

Note: UPN(e.g. testadmin@example.com) based login is not supported.<br/> Only Administrators on Root access group are allowed to login.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**AuthLogin**](AuthLogin.md)| Login credentials needed for Authentication | 

### Return type

[**AuthTokens**](AuthTokens.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutUser

> LogoutUser(ctx, body)

Logs out a user.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RefreshToken**](RefreshToken.md)| Refresh token needed for Logout | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshAccessToken

> AccessToken RefreshAccessToken(ctx, body)

Refreshes access token from refresh token.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**body** | [**RefreshToken**](RefreshToken.md)| Refresh token needed to generate new Access Token | 

### Return type

[**AccessToken**](AccessToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

