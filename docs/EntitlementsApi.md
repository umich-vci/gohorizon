# \EntitlementsApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkCreateApplicationPoolEntitlements**](EntitlementsApi.md#BulkCreateApplicationPoolEntitlements) | **Post** /entitlements/v1/application-pools | Create the bulk entitlements for a set of application pools
[**BulkCreateDesktopPoolEntitlements**](EntitlementsApi.md#BulkCreateDesktopPoolEntitlements) | **Post** /entitlements/v1/desktop-pools | Create the bulk entitlements for a set of desktop pools
[**BulkCreateGAEEntitlements**](EntitlementsApi.md#BulkCreateGAEEntitlements) | **Post** /entitlements/v1/global-application-entitlements | Create the bulk entitlements for a set of Global Application Entitlements
[**BulkCreateGDEEntitlements**](EntitlementsApi.md#BulkCreateGDEEntitlements) | **Post** /entitlements/v1/global-desktop-entitlements | Create the bulk entitlements for a set of Global Desktop Entitlements
[**BulkDeleteApplicationPoolEntitlements**](EntitlementsApi.md#BulkDeleteApplicationPoolEntitlements) | **Delete** /entitlements/v1/application-pools | Delete the bulk entitlements for a set of application pools
[**BulkDeleteDesktopPoolEntitlements**](EntitlementsApi.md#BulkDeleteDesktopPoolEntitlements) | **Delete** /entitlements/v1/desktop-pools | Delete the bulk entitlements for a set of desktop pools
[**BulkDeleteGAEEntitlements**](EntitlementsApi.md#BulkDeleteGAEEntitlements) | **Delete** /entitlements/v1/global-application-entitlements | Delete the bulk entitlements for a set of Global Application Entitlements
[**BulkDeleteGDEEntitlements**](EntitlementsApi.md#BulkDeleteGDEEntitlements) | **Delete** /entitlements/v1/global-desktop-entitlements | Delete the bulk entitlements for a set of Global Desktop Entitlements
[**GetApplicationPoolEntitlements**](EntitlementsApi.md#GetApplicationPoolEntitlements) | **Get** /entitlements/v1/application-pools/{id} | Returns the IDs of users or groups entitled to a given application pool.
[**GetDesktopPoolEntitlements**](EntitlementsApi.md#GetDesktopPoolEntitlements) | **Get** /entitlements/v1/desktop-pools/{id} | Returns the IDs of users or groups entitled to a given desktop pool.
[**GetGAEEntitlement**](EntitlementsApi.md#GetGAEEntitlement) | **Get** /entitlements/v1/global-application-entitlements/{id} | Gets the user or group entitlements for a Global Application Entitlement.
[**GetGDEEntitlement**](EntitlementsApi.md#GetGDEEntitlement) | **Get** /entitlements/v1/global-desktop-entitlements/{id} | Gets the user or group entitlements for a Global Desktop Entitlement.
[**ListApplicationPoolEntitlements**](EntitlementsApi.md#ListApplicationPoolEntitlements) | **Get** /entitlements/v1/application-pools | Lists the entitlements for Application Pools in the environment.
[**ListDesktopPoolEntitlements**](EntitlementsApi.md#ListDesktopPoolEntitlements) | **Get** /entitlements/v1/desktop-pools | Lists the entitlements for Desktop Pools in the environment.
[**ListGAEEntitlements**](EntitlementsApi.md#ListGAEEntitlements) | **Get** /entitlements/v1/global-application-entitlements | Lists the user or group entitlements for Global Application Entitlements in the environment.
[**ListGDEEntitlements**](EntitlementsApi.md#ListGDEEntitlements) | **Get** /entitlements/v1/global-desktop-entitlements | Lists the user or group entitlements for Global Desktop Entitlements in the environment.



## BulkCreateApplicationPoolEntitlements

> []BulkEntitlementResponseInfo BulkCreateApplicationPoolEntitlements(ctx).Body(body).Execute()

Create the bulk entitlements for a set of application pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.EntitlementSpec{*openapiclient.NewEntitlementSpec()} // []EntitlementSpec | Specifications for bulk application entitlements to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.BulkCreateApplicationPoolEntitlements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.BulkCreateApplicationPoolEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateApplicationPoolEntitlements`: []BulkEntitlementResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.BulkCreateApplicationPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateApplicationPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]EntitlementSpec**](EntitlementSpec.md) | Specifications for bulk application entitlements to be created. | 

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

> []BulkEntitlementResponseInfo BulkCreateDesktopPoolEntitlements(ctx).Body(body).Execute()

Create the bulk entitlements for a set of desktop pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.EntitlementSpec{*openapiclient.NewEntitlementSpec()} // []EntitlementSpec | Specifications for bulk desktop entitlements to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.BulkCreateDesktopPoolEntitlements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.BulkCreateDesktopPoolEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateDesktopPoolEntitlements`: []BulkEntitlementResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.BulkCreateDesktopPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateDesktopPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]EntitlementSpec**](EntitlementSpec.md) | Specifications for bulk desktop entitlements to be created. | 

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


## BulkCreateGAEEntitlements

> []BulkEntitlementResponseInfo BulkCreateGAEEntitlements(ctx).Body(body).Execute()

Create the bulk entitlements for a set of Global Application Entitlements



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.EntitlementSpec{*openapiclient.NewEntitlementSpec()} // []EntitlementSpec | Specifications for bulk GAE Entitlements to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.BulkCreateGAEEntitlements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.BulkCreateGAEEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateGAEEntitlements`: []BulkEntitlementResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.BulkCreateGAEEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateGAEEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]EntitlementSpec**](EntitlementSpec.md) | Specifications for bulk GAE Entitlements to be created. | 

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


## BulkCreateGDEEntitlements

> []BulkEntitlementResponseInfo BulkCreateGDEEntitlements(ctx).Body(body).Execute()

Create the bulk entitlements for a set of Global Desktop Entitlements



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.EntitlementSpec{*openapiclient.NewEntitlementSpec()} // []EntitlementSpec | Specifications for bulk GDE Entitlements to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.BulkCreateGDEEntitlements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.BulkCreateGDEEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkCreateGDEEntitlements`: []BulkEntitlementResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.BulkCreateGDEEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkCreateGDEEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]EntitlementSpec**](EntitlementSpec.md) | Specifications for bulk GDE Entitlements to be created. | 

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

> []BulkEntitlementResponseInfo BulkDeleteApplicationPoolEntitlements(ctx).Body(body).Execute()

Delete the bulk entitlements for a set of application pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.EntitlementSpec{*openapiclient.NewEntitlementSpec()} // []EntitlementSpec | Specifications for bulk application entitlements to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.BulkDeleteApplicationPoolEntitlements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.BulkDeleteApplicationPoolEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteApplicationPoolEntitlements`: []BulkEntitlementResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.BulkDeleteApplicationPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteApplicationPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]EntitlementSpec**](EntitlementSpec.md) | Specifications for bulk application entitlements to be deleted. | 

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

> []BulkEntitlementResponseInfo BulkDeleteDesktopPoolEntitlements(ctx).Body(body).Execute()

Delete the bulk entitlements for a set of desktop pools



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.EntitlementSpec{*openapiclient.NewEntitlementSpec()} // []EntitlementSpec | Specifications for bulk desktop entitlements to be deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.BulkDeleteDesktopPoolEntitlements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.BulkDeleteDesktopPoolEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteDesktopPoolEntitlements`: []BulkEntitlementResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.BulkDeleteDesktopPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteDesktopPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]EntitlementSpec**](EntitlementSpec.md) | Specifications for bulk desktop entitlements to be deleted. | 

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


## BulkDeleteGAEEntitlements

> []BulkEntitlementResponseInfo BulkDeleteGAEEntitlements(ctx).Body(body).Execute()

Delete the bulk entitlements for a set of Global Application Entitlements



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.EntitlementSpec{*openapiclient.NewEntitlementSpec()} // []EntitlementSpec | Specifications for bulk GAE Entitlements to be Deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.BulkDeleteGAEEntitlements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.BulkDeleteGAEEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteGAEEntitlements`: []BulkEntitlementResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.BulkDeleteGAEEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteGAEEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]EntitlementSpec**](EntitlementSpec.md) | Specifications for bulk GAE Entitlements to be Deleted. | 

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


## BulkDeleteGDEEntitlements

> []BulkEntitlementResponseInfo BulkDeleteGDEEntitlements(ctx).Body(body).Execute()

Delete the bulk entitlements for a set of Global Desktop Entitlements



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := []openapiclient.EntitlementSpec{*openapiclient.NewEntitlementSpec()} // []EntitlementSpec | Specifications for bulk GDE Entitlements to be Deleted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.BulkDeleteGDEEntitlements(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.BulkDeleteGDEEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BulkDeleteGDEEntitlements`: []BulkEntitlementResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.BulkDeleteGDEEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteGDEEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**[]EntitlementSpec**](EntitlementSpec.md) | Specifications for bulk GDE Entitlements to be Deleted. | 

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

> EntitlementInfo GetApplicationPoolEntitlements(ctx, id).Execute()

Returns the IDs of users or groups entitled to a given application pool.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.GetApplicationPoolEntitlements(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetApplicationPoolEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationPoolEntitlements`: EntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetApplicationPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> EntitlementInfo GetDesktopPoolEntitlements(ctx, id).Execute()

Returns the IDs of users or groups entitled to a given desktop pool.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.GetDesktopPoolEntitlements(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetDesktopPoolEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDesktopPoolEntitlements`: EntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetDesktopPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetGAEEntitlement

> EntitlementInfo GetGAEEntitlement(ctx, id).Execute()

Gets the user or group entitlements for a Global Application Entitlement.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.GetGAEEntitlement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetGAEEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGAEEntitlement`: EntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetGAEEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGAEEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetGDEEntitlement

> EntitlementInfo GetGDEEntitlement(ctx, id).Execute()

Gets the user or group entitlements for a Global Desktop Entitlement.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | id

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.GetGDEEntitlement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.GetGDEEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGDEEntitlement`: EntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.GetGDEEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGDEEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []EntitlementInfo ListApplicationPoolEntitlements(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the entitlements for Application Pools in the environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "{"type":"And", "filters":[{"type":"Equals", "name":"<>", "value":"<>"}] }" // string | filter expression built using fields with <b>'supported filters'</b> as described in output <b>model</b> schema of this API. (optional)
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    size := int32(10) // int32 | size, if passed should be > 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.ListApplicationPoolEntitlements(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListApplicationPoolEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationPoolEntitlements`: []EntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListApplicationPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0 | 

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

> []EntitlementInfo ListDesktopPoolEntitlements(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the entitlements for Desktop Pools in the environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "{"type":"And", "filters":[{"type":"Equals", "name":"<>", "value":"<>"}] }" // string | filter expression built using fields with <b>'supported filters'</b> as described in output <b>model</b> schema of this API. (optional)
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    size := int32(10) // int32 | size, if passed should be > 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.ListDesktopPoolEntitlements(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListDesktopPoolEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDesktopPoolEntitlements`: []EntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListDesktopPoolEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDesktopPoolEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0 | 

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


## ListGAEEntitlements

> []EntitlementInfo ListGAEEntitlements(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the user or group entitlements for Global Application Entitlements in the environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "{"type":"And", "filters":[{"type":"Equals", "name":"<>", "value":"<>"}] }" // string | filter expression built using fields with <b>'supported filters'</b> as described in output <b>model</b> schema of this API. (optional)
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    size := int32(10) // int32 | size, if passed should be > 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.ListGAEEntitlements(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListGAEEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGAEEntitlements`: []EntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListGAEEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGAEEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0 | 

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


## ListGDEEntitlements

> []EntitlementInfo ListGDEEntitlements(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the user or group entitlements for Global Desktop Entitlements in the environment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    filter := "{"type":"And", "filters":[{"type":"Equals", "name":"<>", "value":"<>"}] }" // string | filter expression built using fields with <b>'supported filters'</b> as described in output <b>model</b> schema of this API. (optional)
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    size := int32(10) // int32 | size, if passed should be > 0 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EntitlementsApi.ListGDEEntitlements(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsApi.ListGDEEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGDEEntitlements`: []EntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `EntitlementsApi.ListGDEEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGDEEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0 | 

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

