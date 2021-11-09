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
[**GetViewComposerByVCId**](MonitorApi.md#GetViewComposerByVCId) | **Get** /monitor/v1/view-composers/{vcId} | View Composer is no longer supported from Horizon version 2012 onwards.
[**GetVirtualCenterMonitorInfo**](MonitorApi.md#GetVirtualCenterMonitorInfo) | **Get** /monitor/v1/virtual-centers/{id} | Gets monitoring information related to Virtual Center.
[**ListADDomainMonitorInfosV2**](MonitorApi.md#ListADDomainMonitorInfosV2) | **Get** /monitor/v2/ad-domains | Lists monitoring information related to AD Domains of the environment.
[**ListADDomainMonitorInfosV3**](MonitorApi.md#ListADDomainMonitorInfosV3) | **Get** /monitor/v3/ad-domains | Lists monitoring information related to AD Domains of the environment.
[**ListADDomainMonitors**](MonitorApi.md#ListADDomainMonitors) | **Get** /monitor/ad-domains | Lists monitoring information related to AD Domains of the environment.
[**ListConnectionServerMonitors**](MonitorApi.md#ListConnectionServerMonitors) | **Get** /monitor/connection-servers | Lists monitoring information related to Connection Servers of the environment.
[**ListConnectionServerMonitorsV2**](MonitorApi.md#ListConnectionServerMonitorsV2) | **Get** /monitor/v2/connection-servers | Lists monitoring information related to Connection Servers of the environment.
[**ListDesktopPoolMetrics**](MonitorApi.md#ListDesktopPoolMetrics) | **Get** /monitor/v1/desktop-pools/metrics | Lists metrics of desktop pools (except RDS desktop pools).
[**ListFarmMonitors**](MonitorApi.md#ListFarmMonitors) | **Get** /monitor/farms | Lists monitoring information related to Farms of the environment.
[**ListGatewayMonitorInfoV1**](MonitorApi.md#ListGatewayMonitorInfoV1) | **Get** /monitor/gateways | Lists monitoring information related to Gateways registered in the environment.
[**ListGatewayMonitorInfoV2**](MonitorApi.md#ListGatewayMonitorInfoV2) | **Get** /monitor/v2/gateways | Lists monitoring information related to Gateways registered in the environment.
[**ListPodMonitorInfosV1**](MonitorApi.md#ListPodMonitorInfosV1) | **Get** /monitor/v1/pods | Lists monitoring information related to the remote pods.
[**ListPodMonitorInfosV2**](MonitorApi.md#ListPodMonitorInfosV2) | **Get** /monitor/v2/pods | Lists monitoring information related to the remote pods.
[**ListRDSServerMonitors**](MonitorApi.md#ListRDSServerMonitors) | **Get** /monitor/rds-servers | Lists monitoring information related to RDS Servers of the environment.
[**ListSAMLAuthenticatorMonitorsV1**](MonitorApi.md#ListSAMLAuthenticatorMonitorsV1) | **Get** /monitor/saml-authenticators | Lists monitoring information related to SAML Authenticators of the environment.
[**ListSAMLAuthenticatorMonitorsV2**](MonitorApi.md#ListSAMLAuthenticatorMonitorsV2) | **Get** /monitor/v2/saml-authenticators | Lists monitoring information related to SAML Authenticators of the environment.
[**ListTrueSSOMonitorInfos**](MonitorApi.md#ListTrueSSOMonitorInfos) | **Get** /monitor/v1/true-sso | Lists monitoring information related to True SSO connectors.
[**ListViewComposerMonitorsV1**](MonitorApi.md#ListViewComposerMonitorsV1) | **Get** /monitor/view-composers | View Composer is no longer supported from Horizon version 2012 onwards.
[**ListViewComposerMonitorsV2**](MonitorApi.md#ListViewComposerMonitorsV2) | **Get** /monitor/v2/view-composers | View Composer is no longer supported from Horizon version 2012 onwards.
[**ListVirtualCenterMonitors**](MonitorApi.md#ListVirtualCenterMonitors) | **Get** /monitor/virtual-centers | Lists monitoring information related to Virtual Centers of the environment.
[**ListVirtualCenterMonitorsV2**](MonitorApi.md#ListVirtualCenterMonitorsV2) | **Get** /monitor/v2/virtual-centers | Lists monitoring information related to Virtual Centers of the environment.



## GetConnectionServerMonitorInfoV2

> ConnectionServerMonitorInfoV2 GetConnectionServerMonitorInfoV2(ctx, id).Execute()

Gets monitoring information related to Connection Server.

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
    resp, r, err := api_client.MonitorApi.GetConnectionServerMonitorInfoV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetConnectionServerMonitorInfoV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectionServerMonitorInfoV2`: ConnectionServerMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetConnectionServerMonitorInfoV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectionServerMonitorInfoV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> EventDatabaseMonitorInfo GetEventDatabaseMonitor(ctx).Execute()

Returns monitoring information related to Event database of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.GetEventDatabaseMonitor(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetEventDatabaseMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventDatabaseMonitor`: EventDatabaseMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetEventDatabaseMonitor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventDatabaseMonitorRequest struct via the builder pattern


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

> FarmMonitorInfo GetFarmMonitorInfo(ctx, id).Execute()

Gets monitoring information related to farm.

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
    resp, r, err := api_client.MonitorApi.GetFarmMonitorInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetFarmMonitorInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFarmMonitorInfo`: FarmMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetFarmMonitorInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFarmMonitorInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> GatewayMonitorInfoV2 GetGatewayMonitorInfo(ctx, id).Execute()

Gets monitoring information related to a Gateway.

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
    resp, r, err := api_client.MonitorApi.GetGatewayMonitorInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetGatewayMonitorInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGatewayMonitorInfo`: GatewayMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetGatewayMonitorInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayMonitorInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> PodMonitorInfoV2 GetPodMonitorInfoV2(ctx, id).Execute()

Gets monitoring information related to the remote pod.



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
    resp, r, err := api_client.MonitorApi.GetPodMonitorInfoV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetPodMonitorInfoV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPodMonitorInfoV2`: PodMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetPodMonitorInfoV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPodMonitorInfoV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> RDSServerMonitorInfo GetRDSServerMonitors(ctx, id).Execute()

Gets monitoring information related to RDS Server.

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
    resp, r, err := api_client.MonitorApi.GetRDSServerMonitors(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetRDSServerMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRDSServerMonitors`: RDSServerMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetRDSServerMonitors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRDSServerMonitorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RDSServerMonitorInfo**](RDSServerMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSAMLAuthenticatorMonitorInfo

> SAMLAuthenticatorMonitorInfoV2 GetSAMLAuthenticatorMonitorInfo(ctx, id).Execute()

Gets Monitoring Information related to a SAML Authenticator

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
    resp, r, err := api_client.MonitorApi.GetSAMLAuthenticatorMonitorInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetSAMLAuthenticatorMonitorInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSAMLAuthenticatorMonitorInfo`: SAMLAuthenticatorMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetSAMLAuthenticatorMonitorInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSAMLAuthenticatorMonitorInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SAMLAuthenticatorMonitorInfoV2**](SAMLAuthenticatorMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrueSSOMonitorInfo

> TrueSSOMonitorInfo GetTrueSSOMonitorInfo(ctx, id).Execute()

Gets monitoring information related to a True SSO connector.

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
    resp, r, err := api_client.MonitorApi.GetTrueSSOMonitorInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetTrueSSOMonitorInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTrueSSOMonitorInfo`: TrueSSOMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetTrueSSOMonitorInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTrueSSOMonitorInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TrueSSOMonitorInfo**](TrueSSOMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetViewComposerByVCId

> ViewComposerMonitorInfoV2 GetViewComposerByVCId(ctx, vcId).Execute()

View Composer is no longer supported from Horizon version 2012 onwards.

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
    vcId := "vcId_example" // string | vcId

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.GetViewComposerByVCId(context.Background(), vcId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetViewComposerByVCId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetViewComposerByVCId`: ViewComposerMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetViewComposerByVCId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vcId** | **string** | vcId | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetViewComposerByVCIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> VirtualCenterMonitorInfoV2 GetVirtualCenterMonitorInfo(ctx, id).Execute()

Gets monitoring information related to Virtual Center.

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
    resp, r, err := api_client.MonitorApi.GetVirtualCenterMonitorInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.GetVirtualCenterMonitorInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVirtualCenterMonitorInfo`: VirtualCenterMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.GetVirtualCenterMonitorInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVirtualCenterMonitorInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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

> []ADDomainMonitorInfoV2 ListADDomainMonitorInfosV2(ctx).Execute()

Lists monitoring information related to AD Domains of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListADDomainMonitorInfosV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListADDomainMonitorInfosV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListADDomainMonitorInfosV2`: []ADDomainMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListADDomainMonitorInfosV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListADDomainMonitorInfosV2Request struct via the builder pattern


### Return type

[**[]ADDomainMonitorInfoV2**](ADDomainMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADDomainMonitorInfosV3

> []ADDomainMonitorInfoV3 ListADDomainMonitorInfosV3(ctx).Execute()

Lists monitoring information related to AD Domains of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListADDomainMonitorInfosV3(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListADDomainMonitorInfosV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListADDomainMonitorInfosV3`: []ADDomainMonitorInfoV3
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListADDomainMonitorInfosV3`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListADDomainMonitorInfosV3Request struct via the builder pattern


### Return type

[**[]ADDomainMonitorInfoV3**](ADDomainMonitorInfoV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListADDomainMonitors

> []ADDomainMonitorInfo ListADDomainMonitors(ctx).Execute()

Lists monitoring information related to AD Domains of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListADDomainMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListADDomainMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListADDomainMonitors`: []ADDomainMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListADDomainMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListADDomainMonitorsRequest struct via the builder pattern


### Return type

[**[]ADDomainMonitorInfo**](ADDomainMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectionServerMonitors

> []ConnectionServerMonitorInfo ListConnectionServerMonitors(ctx).Execute()

Lists monitoring information related to Connection Servers of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListConnectionServerMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListConnectionServerMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionServerMonitors`: []ConnectionServerMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListConnectionServerMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionServerMonitorsRequest struct via the builder pattern


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

> []ConnectionServerMonitorInfoV2 ListConnectionServerMonitorsV2(ctx).Execute()

Lists monitoring information related to Connection Servers of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListConnectionServerMonitorsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListConnectionServerMonitorsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectionServerMonitorsV2`: []ConnectionServerMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListConnectionServerMonitorsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectionServerMonitorsV2Request struct via the builder pattern


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


## ListDesktopPoolMetrics

> []DesktopPoolMetricsInfo ListDesktopPoolMetrics(ctx).Ids(ids).Execute()

Lists metrics of desktop pools (except RDS desktop pools).



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
    ids := []string{"Inner_example"} // []string | Desktop pool IDs

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListDesktopPoolMetrics(context.Background()).Ids(ids).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListDesktopPoolMetrics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDesktopPoolMetrics`: []DesktopPoolMetricsInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListDesktopPoolMetrics`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDesktopPoolMetricsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **[]string** | Desktop pool IDs | 

### Return type

[**[]DesktopPoolMetricsInfo**](DesktopPoolMetricsInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFarmMonitors

> []FarmMonitorInfo ListFarmMonitors(ctx).Execute()

Lists monitoring information related to Farms of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListFarmMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListFarmMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFarmMonitors`: []FarmMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListFarmMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFarmMonitorsRequest struct via the builder pattern


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

> []GatewayMonitorInfo ListGatewayMonitorInfoV1(ctx).Execute()

Lists monitoring information related to Gateways registered in the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListGatewayMonitorInfoV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListGatewayMonitorInfoV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGatewayMonitorInfoV1`: []GatewayMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListGatewayMonitorInfoV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGatewayMonitorInfoV1Request struct via the builder pattern


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

> []GatewayMonitorInfoV2 ListGatewayMonitorInfoV2(ctx).Execute()

Lists monitoring information related to Gateways registered in the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListGatewayMonitorInfoV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListGatewayMonitorInfoV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGatewayMonitorInfoV2`: []GatewayMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListGatewayMonitorInfoV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGatewayMonitorInfoV2Request struct via the builder pattern


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

> []PodMonitorInfo ListPodMonitorInfosV1(ctx).Execute()

Lists monitoring information related to the remote pods.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListPodMonitorInfosV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListPodMonitorInfosV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPodMonitorInfosV1`: []PodMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListPodMonitorInfosV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPodMonitorInfosV1Request struct via the builder pattern


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

> []PodMonitorInfoV2 ListPodMonitorInfosV2(ctx).Execute()

Lists monitoring information related to the remote pods.



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListPodMonitorInfosV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListPodMonitorInfosV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPodMonitorInfosV2`: []PodMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListPodMonitorInfosV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPodMonitorInfosV2Request struct via the builder pattern


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

> []RDSServerMonitorInfo ListRDSServerMonitors(ctx).Execute()

Lists monitoring information related to RDS Servers of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListRDSServerMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListRDSServerMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRDSServerMonitors`: []RDSServerMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListRDSServerMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRDSServerMonitorsRequest struct via the builder pattern


### Return type

[**[]RDSServerMonitorInfo**](RDSServerMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSAMLAuthenticatorMonitorsV1

> []SAMLAuthenticatorMonitorInfo ListSAMLAuthenticatorMonitorsV1(ctx).Execute()

Lists monitoring information related to SAML Authenticators of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListSAMLAuthenticatorMonitorsV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListSAMLAuthenticatorMonitorsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSAMLAuthenticatorMonitorsV1`: []SAMLAuthenticatorMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListSAMLAuthenticatorMonitorsV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSAMLAuthenticatorMonitorsV1Request struct via the builder pattern


### Return type

[**[]SAMLAuthenticatorMonitorInfo**](SAMLAuthenticatorMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSAMLAuthenticatorMonitorsV2

> []SAMLAuthenticatorMonitorInfoV2 ListSAMLAuthenticatorMonitorsV2(ctx).Execute()

Lists monitoring information related to SAML Authenticators of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListSAMLAuthenticatorMonitorsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListSAMLAuthenticatorMonitorsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSAMLAuthenticatorMonitorsV2`: []SAMLAuthenticatorMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListSAMLAuthenticatorMonitorsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSAMLAuthenticatorMonitorsV2Request struct via the builder pattern


### Return type

[**[]SAMLAuthenticatorMonitorInfoV2**](SAMLAuthenticatorMonitorInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTrueSSOMonitorInfos

> []TrueSSOMonitorInfo ListTrueSSOMonitorInfos(ctx).Execute()

Lists monitoring information related to True SSO connectors.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListTrueSSOMonitorInfos(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListTrueSSOMonitorInfos``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTrueSSOMonitorInfos`: []TrueSSOMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListTrueSSOMonitorInfos`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTrueSSOMonitorInfosRequest struct via the builder pattern


### Return type

[**[]TrueSSOMonitorInfo**](TrueSSOMonitorInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListViewComposerMonitorsV1

> []ViewComposerMonitorInfo ListViewComposerMonitorsV1(ctx).Execute()

View Composer is no longer supported from Horizon version 2012 onwards.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListViewComposerMonitorsV1(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListViewComposerMonitorsV1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListViewComposerMonitorsV1`: []ViewComposerMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListViewComposerMonitorsV1`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListViewComposerMonitorsV1Request struct via the builder pattern


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

> []ViewComposerMonitorInfoV2 ListViewComposerMonitorsV2(ctx).Execute()

View Composer is no longer supported from Horizon version 2012 onwards.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListViewComposerMonitorsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListViewComposerMonitorsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListViewComposerMonitorsV2`: []ViewComposerMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListViewComposerMonitorsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListViewComposerMonitorsV2Request struct via the builder pattern


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

> []VirtualCenterMonitorInfo ListVirtualCenterMonitors(ctx).Execute()

Lists monitoring information related to Virtual Centers of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListVirtualCenterMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListVirtualCenterMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualCenterMonitors`: []VirtualCenterMonitorInfo
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListVirtualCenterMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualCenterMonitorsRequest struct via the builder pattern


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

> []VirtualCenterMonitorInfoV2 ListVirtualCenterMonitorsV2(ctx).Execute()

Lists monitoring information related to Virtual Centers of the environment.

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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.MonitorApi.ListVirtualCenterMonitorsV2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MonitorApi.ListVirtualCenterMonitorsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVirtualCenterMonitorsV2`: []VirtualCenterMonitorInfoV2
    fmt.Fprintf(os.Stdout, "Response from `MonitorApi.ListVirtualCenterMonitorsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVirtualCenterMonitorsV2Request struct via the builder pattern


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

