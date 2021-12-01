# \InventoryApi

All URIs are relative to *http://localhost/rest*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCustomIcon**](InventoryApi.md#AddCustomIcon) | **Post** /inventory/v1/application-pools/{id}/action/add-custom-icon | Associates a custom icon to the application pool.
[**AddLocalApplicationPoolsToGAE**](InventoryApi.md#AddLocalApplicationPoolsToGAE) | **Post** /inventory/v1/global-application-entitlements/{id}/local-application-pools | Adds Local Application Pools to Global Application Entitlement.
[**AddLocalDesktopPoolsToGDE**](InventoryApi.md#AddLocalDesktopPoolsToGDE) | **Post** /inventory/v1/global-desktop-entitlements/{id}/local-desktop-pools | Adds Local Desktop Pools to Global Desktop Entitlement.
[**AddMachines**](InventoryApi.md#AddMachines) | **Post** /inventory/v1/desktop-pools/{id}/action/add-machines | Adds machines to the given manual desktop pool.
[**AddMachinesByName**](InventoryApi.md#AddMachinesByName) | **Post** /inventory/v1/desktop-pools/{id}/action/add-machines-by-name | Adds the named machines to the given desktop pool.
[**AddRdsServers**](InventoryApi.md#AddRdsServers) | **Post** /inventory/v1/farms/{id}/action/add-rds-servers | Add RDS servers to the specified manual farm.
[**AssignMachineAliases**](InventoryApi.md#AssignMachineAliases) | **Post** /inventory/v1/machines/{id}/action/assign-aliases | Assigns the specified aliases to the assigned users of the machine.
[**AssignUsers**](InventoryApi.md#AssignUsers) | **Post** /inventory/v1/machines/{id}/action/assign-users | Assigns the specified users to the machine.
[**CancelDesktopPoolTask**](InventoryApi.md#CancelDesktopPoolTask) | **Post** /inventory/v1/desktop-pools/{id}/tasks/{taskId}/action/cancel | Cancels the instant clone desktop pool push image task.
[**CancelScheduledMaintenance**](InventoryApi.md#CancelScheduledMaintenance) | **Post** /inventory/v1/farms/{id}/action/cancel-scheduled-maintenance | Requests cancellation of the current scheduled maintenance on the specified Instant Clone farm.
[**CancelScheduledPushImage**](InventoryApi.md#CancelScheduledPushImage) | **Post** /inventory/v1/desktop-pools/{id}/action/cancel-scheduled-push-image | Request the cancellation of the current scheduled push image operation on the specified instant clone desktop pool.
[**CheckApplicationPoolNameAvailability**](InventoryApi.md#CheckApplicationPoolNameAvailability) | **Post** /inventory/v1/application-pools/action/check-name-availability | Checks if the given name is available for application pool creation.
[**CheckDesktopPoolNameAvailability**](InventoryApi.md#CheckDesktopPoolNameAvailability) | **Post** /inventory/v1/desktop-pools/action/check-name-availability | Checks if the given name is available for desktop pool creation.
[**CheckFarmNameAvailability**](InventoryApi.md#CheckFarmNameAvailability) | **Post** /inventory/v1/farms/action/check-name-availability | Checks if the given name is available for farm creation.
[**CheckMachinePrefixAvailability**](InventoryApi.md#CheckMachinePrefixAvailability) | **Post** /inventory/v1/machines/action/check-name-availability | Checks if the given prefix is available for machine creation.
[**CheckRDSServerPrefixAvailability**](InventoryApi.md#CheckRDSServerPrefixAvailability) | **Post** /inventory/v1/rds-servers/action/check-name-availability | Checks if the given prefix is available for RDS Server creation.
[**CreateApplicationIcon**](InventoryApi.md#CreateApplicationIcon) | **Post** /inventory/v1/application-icons | Creates an application icon.
[**CreateApplicationPool**](InventoryApi.md#CreateApplicationPool) | **Post** /inventory/v1/application-pools | Creates an application pool.
[**CreateApplicationPoolV2**](InventoryApi.md#CreateApplicationPoolV2) | **Post** /inventory/v2/application-pools | Creates an application pool.
[**CreateDesktopPool**](InventoryApi.md#CreateDesktopPool) | **Post** /inventory/v1/desktop-pools | Creates a desktop pool.
[**CreateFarm**](InventoryApi.md#CreateFarm) | **Post** /inventory/v1/farms | Creates a farm.
[**CreateFarmV2**](InventoryApi.md#CreateFarmV2) | **Post** /inventory/v2/farms | Creates a farm.
[**CreateGlobalApplicationEntitlement**](InventoryApi.md#CreateGlobalApplicationEntitlement) | **Post** /inventory/v1/global-application-entitlements | Creates a Global Application Entitlement.
[**CreateGlobalDesktopEntitlement**](InventoryApi.md#CreateGlobalDesktopEntitlement) | **Post** /inventory/v1/global-desktop-entitlements | Creates a Global Desktop Entitlement.
[**CreateGlobalDesktopEntitlementV2**](InventoryApi.md#CreateGlobalDesktopEntitlementV2) | **Post** /inventory/v2/global-desktop-entitlements | Creates a Global Desktop Entitlement.
[**DeleteApplicationPool**](InventoryApi.md#DeleteApplicationPool) | **Delete** /inventory/v1/application-pools/{id} | Deletes application pool.
[**DeleteDesktopPool**](InventoryApi.md#DeleteDesktopPool) | **Delete** /inventory/v1/desktop-pools/{id} | Deletes a desktop pool.
[**DeleteFarm**](InventoryApi.md#DeleteFarm) | **Delete** /inventory/v1/farms/{id} | Deletes a farm.
[**DeleteGlobalApplicationEntitlement**](InventoryApi.md#DeleteGlobalApplicationEntitlement) | **Delete** /inventory/v1/global-application-entitlements/{id} | Deletes a Global Application Entitlement.
[**DeleteGlobalDesktopEntitlement**](InventoryApi.md#DeleteGlobalDesktopEntitlement) | **Delete** /inventory/v1/global-desktop-entitlements/{id} | Deletes a Global Desktop Entitlement.
[**DeleteMachine**](InventoryApi.md#DeleteMachine) | **Delete** /inventory/v1/machines/{id} | Deletes the machine.
[**DeleteMachines**](InventoryApi.md#DeleteMachines) | **Delete** /inventory/v1/machines | Deletes the specified machines.
[**DeletePhysicalMachine**](InventoryApi.md#DeletePhysicalMachine) | **Delete** /inventory/v1/physical-machines/{id} | Deletes the Physical Machine.
[**DeleteRDSServer**](InventoryApi.md#DeleteRDSServer) | **Delete** /inventory/v1/rds-servers/{id} | Deletes the RDS Server.
[**DisconnectGlobalSessions**](InventoryApi.md#DisconnectGlobalSessions) | **Post** /inventory/v1/global-sessions/action/disconnect | Disconnects global sessions in the environment.
[**DisconnectSessions**](InventoryApi.md#DisconnectSessions) | **Post** /inventory/v1/sessions/action/disconnect | Disconnects locally resourced user sessions.
[**EnterMaintenance**](InventoryApi.md#EnterMaintenance) | **Post** /inventory/v1/machines/action/enter-maintenance | Puts the machines into maintenance mode.
[**ExitMaintenance**](InventoryApi.md#ExitMaintenance) | **Post** /inventory/v1/machines/action/exit-maintenance | Puts the machines out of maintenance mode.
[**GetApplicationIcon**](InventoryApi.md#GetApplicationIcon) | **Get** /inventory/v1/application-icons/{id} | Gets application icon.
[**GetApplicationPool**](InventoryApi.md#GetApplicationPool) | **Get** /inventory/v1/application-pools/{id} | Gets application pool.
[**GetApplicationPoolV2**](InventoryApi.md#GetApplicationPoolV2) | **Get** /inventory/v2/application-pools/{id} | Gets application pool.
[**GetApplicationPoolV3**](InventoryApi.md#GetApplicationPoolV3) | **Get** /inventory/v3/application-pools/{id} | Gets application pool.
[**GetDesktopPool**](InventoryApi.md#GetDesktopPool) | **Get** /inventory/v1/desktop-pools/{id} | Gets the Desktop Pool information.
[**GetDesktopPoolTask**](InventoryApi.md#GetDesktopPoolTask) | **Get** /inventory/v1/desktop-pools/{id}/tasks/{taskId} | Gets the task information on the desktop pool.
[**GetDesktopPoolV2**](InventoryApi.md#GetDesktopPoolV2) | **Get** /inventory/v2/desktop-pools/{id} | Gets the desktop pool information.
[**GetDesktopPoolV3**](InventoryApi.md#GetDesktopPoolV3) | **Get** /inventory/v3/desktop-pools/{id} | Gets the desktop pool information.
[**GetDesktopPoolV4**](InventoryApi.md#GetDesktopPoolV4) | **Get** /inventory/v4/desktop-pools/{id} | Gets the desktop pool information.
[**GetDesktopPoolV5**](InventoryApi.md#GetDesktopPoolV5) | **Get** /inventory/v5/desktop-pools/{id} | Gets the desktop pool information.
[**GetFarm**](InventoryApi.md#GetFarm) | **Get** /inventory/v1/farms/{id} | Gets the Farm information.
[**GetFarmV2**](InventoryApi.md#GetFarmV2) | **Get** /inventory/v2/farms/{id} | Gets the Farm information.
[**GetFarmV3**](InventoryApi.md#GetFarmV3) | **Get** /inventory/v3/farms/{id} | Gets the Farm information.
[**GetGlobalApplicationEntitlement**](InventoryApi.md#GetGlobalApplicationEntitlement) | **Get** /inventory/v1/global-application-entitlements/{id} | Gets the Global Application Entitlement in the environment.
[**GetGlobalApplicationEntitlementV2**](InventoryApi.md#GetGlobalApplicationEntitlementV2) | **Get** /inventory/v2/global-application-entitlements/{id} | Gets the Global Application Entitlement in the environment.
[**GetGlobalDesktopEntitlement**](InventoryApi.md#GetGlobalDesktopEntitlement) | **Get** /inventory/v1/global-desktop-entitlements/{id} | Gets the Global Desktop Entitlement in the environment.
[**GetGlobalDesktopEntitlementV2**](InventoryApi.md#GetGlobalDesktopEntitlementV2) | **Get** /inventory/v2/global-desktop-entitlements/{id} | Gets the Global Desktop Entitlement in the environment.
[**GetMachine**](InventoryApi.md#GetMachine) | **Get** /inventory/v1/machines/{id} | Gets the Machine information.
[**GetMachineV2**](InventoryApi.md#GetMachineV2) | **Get** /inventory/v2/machines/{id} | Gets the Machine information.
[**GetPhysicalMachine**](InventoryApi.md#GetPhysicalMachine) | **Get** /inventory/v1/physical-machines/{id} | Gets the Physical Machine information.
[**GetRDSServer**](InventoryApi.md#GetRDSServer) | **Get** /inventory/v1/rds-servers/{id} | Gets the RDS Server information.
[**GetSessionInfo**](InventoryApi.md#GetSessionInfo) | **Get** /inventory/v1/sessions/{id} | Gets the Session information for locally resourced session.
[**ListApplicationIcons**](InventoryApi.md#ListApplicationIcons) | **Get** /inventory/v1/application-icons | Lists the application icons for the given application pool.
[**ListApplicationPools**](InventoryApi.md#ListApplicationPools) | **Get** /inventory/v1/application-pools | Lists the application pools in the environment.
[**ListApplicationPoolsV2**](InventoryApi.md#ListApplicationPoolsV2) | **Get** /inventory/v2/application-pools | Lists the application pools in the environment.
[**ListApplicationPoolsV3**](InventoryApi.md#ListApplicationPoolsV3) | **Get** /inventory/v3/application-pools | Lists the application pools in the environment.
[**ListCompatibleBackupGAEs**](InventoryApi.md#ListCompatibleBackupGAEs) | **Get** /inventory/v1/global-application-entitlements/{id}/compatible-backup-global-application-entitlements | Lists the Global Application Entitlements that can be associated as backup Global Application Entitlement.
[**ListCompatibleBackupGDEs**](InventoryApi.md#ListCompatibleBackupGDEs) | **Get** /inventory/v1/global-desktop-entitlements/{id}/compatible-backup-global-desktop-entitlements | Lists the Global Desktop Entitlements that can be associated as backup Global Desktop Entitlement.
[**ListCompatibleLocalApplicationPools**](InventoryApi.md#ListCompatibleLocalApplicationPools) | **Get** /inventory/v1/global-application-entitlements/{id}/compatible-local-application-pools | Lists Local Application Pools which are compatible with Global Application Entitlement.
[**ListCompatibleLocalDesktopPools**](InventoryApi.md#ListCompatibleLocalDesktopPools) | **Get** /inventory/v1/global-desktop-entitlements/{id}/compatible-local-desktop-pools | Lists Local Desktop Pools which are compatible with Global Desktop Entitlement.
[**ListDesktopPoolTasks**](InventoryApi.md#ListDesktopPoolTasks) | **Get** /inventory/v1/desktop-pools/{id}/tasks | Lists the tasks on the desktop pool.
[**ListDesktopPools**](InventoryApi.md#ListDesktopPools) | **Get** /inventory/v1/desktop-pools | Lists the Desktop Pools in the environment.
[**ListDesktopPoolsV2**](InventoryApi.md#ListDesktopPoolsV2) | **Get** /inventory/v2/desktop-pools | Lists the desktop pools in the environment.
[**ListDesktopPoolsV3**](InventoryApi.md#ListDesktopPoolsV3) | **Get** /inventory/v3/desktop-pools | Lists the desktop pools in the environment.
[**ListDesktopPoolsV4**](InventoryApi.md#ListDesktopPoolsV4) | **Get** /inventory/v4/desktop-pools | Lists the desktop pools in the environment.
[**ListDesktopPoolsV5**](InventoryApi.md#ListDesktopPoolsV5) | **Get** /inventory/v5/desktop-pools | Lists the desktop pools in the environment.
[**ListFarms**](InventoryApi.md#ListFarms) | **Get** /inventory/v1/farms | Lists the Farms in the environment.
[**ListFarmsV2**](InventoryApi.md#ListFarmsV2) | **Get** /inventory/v2/farms | Lists the Farms in the environment.
[**ListFarmsV3**](InventoryApi.md#ListFarmsV3) | **Get** /inventory/v3/farms | Lists the Farms in the environment.
[**ListGlobalApplicationEntitlements**](InventoryApi.md#ListGlobalApplicationEntitlements) | **Get** /inventory/v1/global-application-entitlements | Lists the Global Application Entitlements in the environment.
[**ListGlobalApplicationEntitlementsV2**](InventoryApi.md#ListGlobalApplicationEntitlementsV2) | **Get** /inventory/v2/global-application-entitlements | Lists the Global Application Entitlements in the environment.
[**ListGlobalDesktopEntitlements**](InventoryApi.md#ListGlobalDesktopEntitlements) | **Get** /inventory/v1/global-desktop-entitlements | Lists the Global Desktop Entitlements in the environment.
[**ListGlobalDesktopEntitlementsV2**](InventoryApi.md#ListGlobalDesktopEntitlementsV2) | **Get** /inventory/v2/global-desktop-entitlements | Lists the Global Desktop Entitlements in the environment.
[**ListInstalledApplicationsOnDesktopPool**](InventoryApi.md#ListInstalledApplicationsOnDesktopPool) | **Get** /inventory/v1/desktop-pools/{id}/installed-applications | Lists the installed applications on the given desktop pool.
[**ListInstalledApplicationsOnFarm**](InventoryApi.md#ListInstalledApplicationsOnFarm) | **Get** /inventory/v1/farms/{id}/installed-applications | Lists the installed applications on the given farm.
[**ListLocalApplicationPools**](InventoryApi.md#ListLocalApplicationPools) | **Get** /inventory/v1/global-application-entitlements/{id}/local-application-pools | Lists Local Application Pools which are assigned to Global Application Entitlement.
[**ListLocalDesktopPools**](InventoryApi.md#ListLocalDesktopPools) | **Get** /inventory/v1/global-desktop-entitlements/{id}/local-desktop-pools | Lists Local Desktop Pools which are assigned to Global Desktop Entitlement.
[**ListMachines**](InventoryApi.md#ListMachines) | **Get** /inventory/v1/machines | Lists the Machines in the environment.
[**ListMachinesV2**](InventoryApi.md#ListMachinesV2) | **Get** /inventory/v2/machines | Lists the Machines in the environment.
[**ListPhysicalMachines**](InventoryApi.md#ListPhysicalMachines) | **Get** /inventory/v1/physical-machines | Lists the Physical Machines in the environment.
[**ListRDSServers**](InventoryApi.md#ListRDSServers) | **Get** /inventory/v1/rds-servers | Lists the RDS Servers in the environment.
[**ListSessionInfo**](InventoryApi.md#ListSessionInfo) | **Get** /inventory/v1/sessions | Lists the locally resourced Sessions in the environment
[**LogOffGlobalSessions**](InventoryApi.md#LogOffGlobalSessions) | **Post** /inventory/v1/global-sessions/action/logoff | Logs off global sessions in the environment.
[**LogOffSessions**](InventoryApi.md#LogOffSessions) | **Post** /inventory/v1/sessions/action/logoff | Logs off locally resourced user sessions, if they are not locked.
[**PauseDesktopPoolTask**](InventoryApi.md#PauseDesktopPoolTask) | **Post** /inventory/v1/desktop-pools/{id}/tasks/{taskId}/action/pause | Pause the instant clone desktop pool push image task.
[**QueryGlobalSessions**](InventoryApi.md#QueryGlobalSessions) | **Get** /inventory/v1/global-sessions | Lists global sessions in the environment for the given user, pod or brokering pod.
[**RebuildMachines**](InventoryApi.md#RebuildMachines) | **Post** /inventory/v1/machines/action/rebuild | Rebuilds the specified machines.
[**RecoverMachines**](InventoryApi.md#RecoverMachines) | **Post** /inventory/v1/machines/action/recover | Recovers the specified machines.
[**RecoverRDSServers**](InventoryApi.md#RecoverRDSServers) | **Post** /inventory/v1/rds-servers/action/recover | Recovers the specified RDS Servers.
[**RegisterPhysicalMachine**](InventoryApi.md#RegisterPhysicalMachine) | **Post** /inventory/v1/physical-machines/action/register | Registers the Physical Machine.
[**RegisterRDSServer**](InventoryApi.md#RegisterRDSServer) | **Post** /inventory/v1/rds-servers/action/register | Registers the RDS Server.
[**RemoveCustomIcon**](InventoryApi.md#RemoveCustomIcon) | **Post** /inventory/v1/application-pools/{id}/action/remove-custom-icon | Removes the associated custom icon from the application pool.
[**RemoveLocalApplicationPoolsFromGAE**](InventoryApi.md#RemoveLocalApplicationPoolsFromGAE) | **Delete** /inventory/v1/global-application-entitlements/{id}/local-application-pools | Removes Local Application Pools from Global Application Entitlement.
[**RemoveLocalDesktopPoolsFromGDE**](InventoryApi.md#RemoveLocalDesktopPoolsFromGDE) | **Delete** /inventory/v1/global-desktop-entitlements/{id}/local-desktop-pools | Removes Local Desktop Pools from Global Desktop Entitlement.
[**RemoveMachines**](InventoryApi.md#RemoveMachines) | **Post** /inventory/v1/desktop-pools/{id}/action/remove-machines | Removes machines from the given manual desktop pool.
[**RemoveRdsServers**](InventoryApi.md#RemoveRdsServers) | **Post** /inventory/v1/farms/{id}/action/remove-rds-servers | Remove RDS servers from the specified farm.
[**ResetGlobalSessions**](InventoryApi.md#ResetGlobalSessions) | **Post** /inventory/v1/global-sessions/action/reset | Resets machines of global sessions in the environment.
[**ResetMachines**](InventoryApi.md#ResetMachines) | **Post** /inventory/v1/machines/action/reset | Resets the specified machines.
[**ResetSessions**](InventoryApi.md#ResetSessions) | **Post** /inventory/v1/sessions/action/reset | Resets machine of locally resourced user sessions. The machine must be managed by Virtual Center and the session cannot be an application or an RDS desktop session.
[**RestartGlobalSessions**](InventoryApi.md#RestartGlobalSessions) | **Post** /inventory/v1/global-sessions/action/restart | Restarts machines of global sessions in the environment.
[**RestartMachines**](InventoryApi.md#RestartMachines) | **Post** /inventory/v1/machines/action/restart | Restarts the specified machines.
[**RestartSessions**](InventoryApi.md#RestartSessions) | **Post** /inventory/v1/sessions/action/restart | Restarts machine of locally resourced user sessions. The machine must be managed by Virtual Center and the session cannot be an application or an RDS desktop session.
[**ResumeDesktopPoolTask**](InventoryApi.md#ResumeDesktopPoolTask) | **Post** /inventory/v1/desktop-pools/{id}/tasks/{taskId}/action/resume | Resume the instant clone desktop pool push image task.
[**ScheduleMaintenance**](InventoryApi.md#ScheduleMaintenance) | **Post** /inventory/v1/farms/{id}/action/schedule-maintenance | Creates maintenance schedule for the specified farm.
[**SchedulePushImage**](InventoryApi.md#SchedulePushImage) | **Post** /inventory/v1/desktop-pools/{id}/action/schedule-push-image | Schedule/reschedule a request to update the image in an instant clone desktop pool.
[**SendMessageToGlobalSessions**](InventoryApi.md#SendMessageToGlobalSessions) | **Post** /inventory/v1/global-sessions/action/send-message | Sends message to global sessions in the environment.
[**SendMessageToSessions**](InventoryApi.md#SendMessageToSessions) | **Post** /inventory/v1/sessions/action/send-message | Sends the message to locally resourced user sessions.
[**UnassignMachineAliases**](InventoryApi.md#UnassignMachineAliases) | **Post** /inventory/v1/machines/{id}/action/unassign-aliases | Un-assigns the aliases for the specified users from the machine.
[**UnassignUsers**](InventoryApi.md#UnassignUsers) | **Post** /inventory/v1/machines/{id}/action/unassign-users | Un-assigns the specified users from the machine.
[**UpdateApplicationPool**](InventoryApi.md#UpdateApplicationPool) | **Put** /inventory/v1/application-pools/{id} | Updates application pool.
[**UpdateApplicationPoolV2**](InventoryApi.md#UpdateApplicationPoolV2) | **Put** /inventory/v2/application-pools/{id} | Updates application pool.
[**UpdateDesktopPool**](InventoryApi.md#UpdateDesktopPool) | **Put** /inventory/v1/desktop-pools/{id} | Updates desktop pool.
[**UpdateFarm**](InventoryApi.md#UpdateFarm) | **Put** /inventory/v1/farms/{id} | Updates farm.
[**UpdateFarmV2**](InventoryApi.md#UpdateFarmV2) | **Put** /inventory/v2/farms/{id} | Updates farm.
[**UpdateGlobalApplicationEntitlement**](InventoryApi.md#UpdateGlobalApplicationEntitlement) | **Put** /inventory/v1/global-application-entitlements/{id} | Updates a Global Application Entitlement.
[**UpdateGlobalDesktopEntitlement**](InventoryApi.md#UpdateGlobalDesktopEntitlement) | **Put** /inventory/v1/global-desktop-entitlements/{id} | Updates a Global Desktop Entitlement.
[**UpdateRDSServer**](InventoryApi.md#UpdateRDSServer) | **Put** /inventory/v1/rds-servers/{id} | Updates the RDS Server.
[**ValidateInstalledApplicationsOnDesktopPool**](InventoryApi.md#ValidateInstalledApplicationsOnDesktopPool) | **Post** /inventory/v1/desktop-pools/{id}/action/validate-installed-applications | Validates that each application in the given list is installed on the machines belonging to the specified desktop pool.
[**ValidateInstalledApplicationsOnFarm**](InventoryApi.md#ValidateInstalledApplicationsOnFarm) | **Post** /inventory/v1/farms/{id}/action/validate-installed-applications | Validates that each application in the given list is installed on the RDS Servers belonging to the specified Farm.
[**ValidateSpecifiedNames**](InventoryApi.md#ValidateSpecifiedNames) | **Post** /inventory/v1/desktop-pools/action/validate-specified-names | Validates manually specified virtual machines. Ensures machine and user names are valid and aren&#39;t duplicated in the given desktop pool.



## AddCustomIcon

> AddCustomIcon(ctx, id).Body(body).Execute()

Associates a custom icon to the application pool.



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
    id := "id_example" // string | Application pool ID
    body := *openapiclient.NewApplicationIconAssociateSpec("6f85b3a5-e7d0-4ad6-a1e3-37168dd1ed51") // ApplicationIconAssociateSpec | Icon id to be associated with the application pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.AddCustomIcon(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.AddCustomIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Application pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCustomIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationIconAssociateSpec**](ApplicationIconAssociateSpec.md) | Icon id to be associated with the application pool. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLocalApplicationPoolsToGAE

> []BulkItemResponseInfo AddLocalApplicationPoolsToGAE(ctx, id).Body(body).Execute()

Adds Local Application Pools to Global Application Entitlement.



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
    body := []string{"Property_example"} // []string | List of local application pool ids to be added.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.AddLocalApplicationPoolsToGAE(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.AddLocalApplicationPoolsToGAE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocalApplicationPoolsToGAE`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.AddLocalApplicationPoolsToGAE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLocalApplicationPoolsToGAERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of local application pool ids to be added. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddLocalDesktopPoolsToGDE

> []BulkItemResponseInfo AddLocalDesktopPoolsToGDE(ctx, id).Body(body).Execute()

Adds Local Desktop Pools to Global Desktop Entitlement.



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
    body := []string{"Property_example"} // []string | List of local desktop pool ids to be added.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.AddLocalDesktopPoolsToGDE(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.AddLocalDesktopPoolsToGDE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddLocalDesktopPoolsToGDE`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.AddLocalDesktopPoolsToGDE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddLocalDesktopPoolsToGDERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of local desktop pool ids to be added. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMachines

> []BulkItemResponseInfo AddMachines(ctx, id).Body(body).Execute()

Adds machines to the given manual desktop pool.



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
    body := []string{"Property_example"} // []string | List of Machine Ids representing the machines to be added to the desktop pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.AddMachines(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.AddMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMachines`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.AddMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of Machine Ids representing the machines to be added to the desktop pool. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddMachinesByName

> []BulkItemResponseInfo AddMachinesByName(ctx, id).Body(body).Execute()

Adds the named machines to the given desktop pool.



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
    body := []openapiclient.MachineSpecifiedName{*openapiclient.NewMachineSpecifiedName("machine1")} // []MachineSpecifiedName | List of MachineSpecifiedName representing the machines to be added to the desktop pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.AddMachinesByName(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.AddMachinesByName``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddMachinesByName`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.AddMachinesByName`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMachinesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]MachineSpecifiedName**](MachineSpecifiedName.md) | List of MachineSpecifiedName representing the machines to be added to the desktop pool. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddRdsServers

> []BulkItemResponseInfo AddRdsServers(ctx, id).Body(body).Execute()

Add RDS servers to the specified manual farm.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Farm ID
    body := []string{"Property_example"} // []string | List of RDS server IDs representing the RDS servers to be added to the farm.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.AddRdsServers(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.AddRdsServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddRdsServers`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.AddRdsServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Farm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRdsServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of RDS server IDs representing the RDS servers to be added to the farm. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignMachineAliases

> []BulkItemResponseInfo AssignMachineAliases(ctx, id).Body(body).Execute()

Assigns the specified aliases to the assigned users of the machine.



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
    body := []openapiclient.MachineAliasSpec{*openapiclient.NewMachineAliasSpec()} // []MachineAliasSpec | List of MachineAlias. If a user is assigned to the machine we can set the corresponding aliases.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.AssignMachineAliases(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.AssignMachineAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignMachineAliases`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.AssignMachineAliases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignMachineAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**[]MachineAliasSpec**](MachineAliasSpec.md) | List of MachineAlias. If a user is assigned to the machine we can set the corresponding aliases. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignUsers

> []BulkItemResponseInfo AssignUsers(ctx, id).Body(body).Execute()

Assigns the specified users to the machine.



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
    body := []string{"Property_example"} // []string | List of User SIDs representing the users to be assigned to the machine.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.AssignUsers(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.AssignUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssignUsers`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.AssignUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of User SIDs representing the users to be assigned to the machine. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelDesktopPoolTask

> CancelDesktopPoolTask(ctx, id, taskId).Execute()

Cancels the instant clone desktop pool push image task.



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
    id := "id_example" // string | Desktop pool ID
    taskId := "taskId_example" // string | Desktop pool task ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CancelDesktopPoolTask(context.Background(), id, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CancelDesktopPoolTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Desktop pool ID | 
**taskId** | **string** | Desktop pool task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelDesktopPoolTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelScheduledMaintenance

> CancelScheduledMaintenance(ctx, id).Body(body).Execute()

Requests cancellation of the current scheduled maintenance on the specified Instant Clone farm.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Farm ID
    body := *openapiclient.NewFarmCancelMaintenanceSpec("IMMEDIATE") // FarmCancelMaintenanceSpec | The specification to cancel the scheduled maintenance operation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CancelScheduledMaintenance(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CancelScheduledMaintenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Farm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelScheduledMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FarmCancelMaintenanceSpec**](FarmCancelMaintenanceSpec.md) | The specification to cancel the scheduled maintenance operation. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelScheduledPushImage

> CancelScheduledPushImage(ctx, id).Execute()

Request the cancellation of the current scheduled push image operation on the specified instant clone desktop pool.



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
    resp, r, err := api_client.InventoryApi.CancelScheduledPushImage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CancelScheduledPushImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelScheduledPushImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckApplicationPoolNameAvailability

> NameAvailabilityInfo CheckApplicationPoolNameAvailability(ctx).Body(body).Execute()

Checks if the given name is available for application pool creation.



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
    body := *openapiclient.NewNameAvailabilitySpec("Name_example") // NameAvailabilitySpec | Name Availability Spec.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CheckApplicationPoolNameAvailability(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CheckApplicationPoolNameAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckApplicationPoolNameAvailability`: NameAvailabilityInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CheckApplicationPoolNameAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckApplicationPoolNameAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NameAvailabilitySpec**](NameAvailabilitySpec.md) | Name Availability Spec. | 

### Return type

[**NameAvailabilityInfo**](NameAvailabilityInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckDesktopPoolNameAvailability

> NameAvailabilityInfo CheckDesktopPoolNameAvailability(ctx).Body(body).Execute()

Checks if the given name is available for desktop pool creation.



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
    body := *openapiclient.NewNameAvailabilitySpec("Name_example") // NameAvailabilitySpec | Name Availability Spec.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CheckDesktopPoolNameAvailability(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CheckDesktopPoolNameAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckDesktopPoolNameAvailability`: NameAvailabilityInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CheckDesktopPoolNameAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckDesktopPoolNameAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NameAvailabilitySpec**](NameAvailabilitySpec.md) | Name Availability Spec. | 

### Return type

[**NameAvailabilityInfo**](NameAvailabilityInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckFarmNameAvailability

> NameAvailabilityInfo CheckFarmNameAvailability(ctx).Body(body).Execute()

Checks if the given name is available for farm creation.



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
    body := *openapiclient.NewNameAvailabilitySpec("Name_example") // NameAvailabilitySpec | Name Availability Spec.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CheckFarmNameAvailability(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CheckFarmNameAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckFarmNameAvailability`: NameAvailabilityInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CheckFarmNameAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckFarmNameAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NameAvailabilitySpec**](NameAvailabilitySpec.md) | Name Availability Spec. | 

### Return type

[**NameAvailabilityInfo**](NameAvailabilityInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckMachinePrefixAvailability

> NameAvailabilityInfo CheckMachinePrefixAvailability(ctx).Body(body).Execute()

Checks if the given prefix is available for machine creation.



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
    body := *openapiclient.NewNameAvailabilitySpec("Name_example") // NameAvailabilitySpec | Name Availability Spec.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CheckMachinePrefixAvailability(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CheckMachinePrefixAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckMachinePrefixAvailability`: NameAvailabilityInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CheckMachinePrefixAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckMachinePrefixAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NameAvailabilitySpec**](NameAvailabilitySpec.md) | Name Availability Spec. | 

### Return type

[**NameAvailabilityInfo**](NameAvailabilityInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckRDSServerPrefixAvailability

> NameAvailabilityInfo CheckRDSServerPrefixAvailability(ctx).Body(body).Execute()

Checks if the given prefix is available for RDS Server creation.



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
    body := *openapiclient.NewNameAvailabilitySpec("Name_example") // NameAvailabilitySpec | Name Availability Spec.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CheckRDSServerPrefixAvailability(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CheckRDSServerPrefixAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckRDSServerPrefixAvailability`: NameAvailabilityInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.CheckRDSServerPrefixAvailability`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCheckRDSServerPrefixAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**NameAvailabilitySpec**](NameAvailabilitySpec.md) | Name Availability Spec. | 

### Return type

[**NameAvailabilityInfo**](NameAvailabilityInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationIcon

> CreateApplicationIcon(ctx).Body(body).Execute()

Creates an application icon.



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
    body := *openapiclient.NewApplicationIconCreateSpec("Data_example", int64(16), int64(16)) // ApplicationIconCreateSpec | Application icon object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateApplicationIcon(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateApplicationIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationIconCreateSpec**](ApplicationIconCreateSpec.md) | Application icon object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationPool

> CreateApplicationPool(ctx).Body(body).Execute()

Creates an application pool.



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
    body := *openapiclient.NewApplicationPoolCreateSpec("C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Firefox.lnk", "Firefox") // ApplicationPoolCreateSpec | Application pool object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateApplicationPool(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateApplicationPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationPoolCreateSpec**](ApplicationPoolCreateSpec.md) | Application pool object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateApplicationPoolV2

> CreateApplicationPoolV2(ctx).Body(body).Execute()

Creates an application pool.



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
    body := *openapiclient.NewApplicationPoolCreateSpecV2("C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Firefox.lnk", "Firefox") // ApplicationPoolCreateSpecV2 | Application pool object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateApplicationPoolV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateApplicationPoolV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationPoolV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ApplicationPoolCreateSpecV2**](ApplicationPoolCreateSpecV2.md) | Application pool object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDesktopPool

> CreateDesktopPool(ctx).Body(body).Execute()

Creates a desktop pool.



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
    body := *openapiclient.NewDesktopPoolCreateSpec("ManualPool", "AUTOMATED") // DesktopPoolCreateSpec | Desktop pool object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateDesktopPool(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateDesktopPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDesktopPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DesktopPoolCreateSpec**](DesktopPoolCreateSpec.md) | Desktop pool object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFarm

> CreateFarm(ctx).Body(body).Execute()

Creates a farm.



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
    body := *openapiclient.NewFarmCreateSpec("6fd4638a-381f-4518-aed6-042aa3d9f14c", "ManualFarm", "MANUAL") // FarmCreateSpec | Farm object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateFarm(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FarmCreateSpec**](FarmCreateSpec.md) | Farm object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFarmV2

> CreateFarmV2(ctx).Body(body).Execute()

Creates a farm.



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
    body := *openapiclient.NewFarmCreateSpecV2("6fd4638a-381f-4518-aed6-042aa3d9f14c", "ManualFarm", "MANUAL") // FarmCreateSpecV2 | Farm object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateFarmV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateFarmV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFarmV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FarmCreateSpecV2**](FarmCreateSpecV2.md) | Farm object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlobalApplicationEntitlement

> CreateGlobalApplicationEntitlement(ctx).Body(body).Execute()

Creates a Global Application Entitlement.



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
    body := *openapiclient.NewGlobalApplicationEntitlementCreateSpec("310927ee-1195-44c3-b877-5b194b4b4ea8", "global-application-entitlement") // GlobalApplicationEntitlementCreateSpec | Global Application Entitlement object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateGlobalApplicationEntitlement(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateGlobalApplicationEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalApplicationEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GlobalApplicationEntitlementCreateSpec**](GlobalApplicationEntitlementCreateSpec.md) | Global Application Entitlement object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlobalDesktopEntitlement

> CreateGlobalDesktopEntitlement(ctx).Body(body).Execute()

Creates a Global Desktop Entitlement.



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
    body := *openapiclient.NewGlobalDesktopEntitlementCreateSpec("global-desktop-entitlement") // GlobalDesktopEntitlementCreateSpec | Global Desktop Entitlement object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateGlobalDesktopEntitlement(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateGlobalDesktopEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalDesktopEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GlobalDesktopEntitlementCreateSpec**](GlobalDesktopEntitlementCreateSpec.md) | Global Desktop Entitlement object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGlobalDesktopEntitlementV2

> CreateGlobalDesktopEntitlementV2(ctx).Body(body).Execute()

Creates a Global Desktop Entitlement.



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
    body := *openapiclient.NewGlobalDesktopEntitlementCreateSpecV2("310927ee-1195-44c3-b877-5b194b4b4ea8", "global-desktop-entitlement") // GlobalDesktopEntitlementCreateSpecV2 | Global Desktop Entitlement object to be created.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.CreateGlobalDesktopEntitlementV2(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.CreateGlobalDesktopEntitlementV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGlobalDesktopEntitlementV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GlobalDesktopEntitlementCreateSpecV2**](GlobalDesktopEntitlementCreateSpecV2.md) | Global Desktop Entitlement object to be created. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApplicationPool

> DeleteApplicationPool(ctx, id).Execute()

Deletes application pool.



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
    resp, r, err := api_client.InventoryApi.DeleteApplicationPool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeleteApplicationPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApplicationPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDesktopPool

> DeleteDesktopPool(ctx, id).Body(body).Execute()

Deletes a desktop pool.



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
    id := "0103796c-102b-4ed3-953f-3dfe3d23e0fe" // string | ID of the desktop pool to be deleted.
    body := *openapiclient.NewDesktopPoolDeleteSpec() // DesktopPoolDeleteSpec | Desktop pool object to be deleted. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DeleteDesktopPool(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeleteDesktopPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of the desktop pool to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDesktopPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DesktopPoolDeleteSpec**](DesktopPoolDeleteSpec.md) | Desktop pool object to be deleted. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFarm

> DeleteFarm(ctx, id).Execute()

Deletes a farm.



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
    resp, r, err := api_client.InventoryApi.DeleteFarm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeleteFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalApplicationEntitlement

> DeleteGlobalApplicationEntitlement(ctx, id).Execute()

Deletes a Global Application Entitlement.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Global Application Entitlement ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DeleteGlobalApplicationEntitlement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeleteGlobalApplicationEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Global Application Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalApplicationEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGlobalDesktopEntitlement

> DeleteGlobalDesktopEntitlement(ctx, id).Execute()

Deletes a Global Desktop Entitlement.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Global Desktop Entitlement ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DeleteGlobalDesktopEntitlement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeleteGlobalDesktopEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Global Desktop Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalDesktopEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMachine

> DeleteMachine(ctx, id).Body(body).Execute()

Deletes the machine.



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
    body := *openapiclient.NewMachineDeleteData() // MachineDeleteData | The specification applicable to deleting the machine.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DeleteMachine(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeleteMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MachineDeleteData**](MachineDeleteData.md) | The specification applicable to deleting the machine. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMachines

> []BulkItemResponseInfo DeleteMachines(ctx).Body(body).Execute()

Deletes the specified machines.



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
    body := *openapiclient.NewMachineDeleteSpec([]string{"MachineIds_example"}) // MachineDeleteSpec | The machines and specification for deletion.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DeleteMachines(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeleteMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMachines`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.DeleteMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MachineDeleteSpec**](MachineDeleteSpec.md) | The machines and specification for deletion. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePhysicalMachine

> DeletePhysicalMachine(ctx, id).Execute()

Deletes the Physical Machine.



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
    id := "id_example" // string | Physical machine ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DeletePhysicalMachine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeletePhysicalMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Physical machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePhysicalMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRDSServer

> DeleteRDSServer(ctx, id).Execute()

Deletes the RDS Server.



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
    id := "id_example" // string | RDS Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DeleteRDSServer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DeleteRDSServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | RDS Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRDSServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectGlobalSessions

> []BulkGlobalSessionActionResponseInfo DisconnectGlobalSessions(ctx).GlobalSessionActionSpecs(globalSessionActionSpecs).Execute()

Disconnects global sessions in the environment.



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
    globalSessionActionSpecs := []openapiclient.GlobalSessionActionSpec{*openapiclient.NewGlobalSessionActionSpec([]string{"Ids_example"}, "54c122e0-d6b5-476c-af91-fd00392e5983")} // []GlobalSessionActionSpec | Sessions which are to be disconnected.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DisconnectGlobalSessions(context.Background()).GlobalSessionActionSpecs(globalSessionActionSpecs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DisconnectGlobalSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisconnectGlobalSessions`: []BulkGlobalSessionActionResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.DisconnectGlobalSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectGlobalSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalSessionActionSpecs** | [**[]GlobalSessionActionSpec**](GlobalSessionActionSpec.md) | Sessions which are to be disconnected. | 

### Return type

[**[]BulkGlobalSessionActionResponseInfo**](BulkGlobalSessionActionResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DisconnectSessions

> []BulkItemResponseInfo DisconnectSessions(ctx).Body(body).Execute()

Disconnects locally resourced user sessions.



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
    body := []string{"Property_example"} // []string | List of session ids to be disconnected.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.DisconnectSessions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.DisconnectSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DisconnectSessions`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.DisconnectSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDisconnectSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of session ids to be disconnected. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnterMaintenance

> []BulkItemResponseInfo EnterMaintenance(ctx).Body(body).Execute()

Puts the machines into maintenance mode.



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
    body := []string{"Property_example"} // []string | List of Machine Ids representing the machines to be put into maintenance mode.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.EnterMaintenance(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.EnterMaintenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EnterMaintenance`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.EnterMaintenance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEnterMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of Machine Ids representing the machines to be put into maintenance mode. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExitMaintenance

> []BulkItemResponseInfo ExitMaintenance(ctx).Body(body).Execute()

Puts the machines out of maintenance mode.



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
    body := []string{"Property_example"} // []string | List of Machine Ids representing the machines to be put out of maintenance mode.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ExitMaintenance(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ExitMaintenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ExitMaintenance`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ExitMaintenance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExitMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of Machine Ids representing the machines to be put out of maintenance mode. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationIcon

> ApplicationIconInfo GetApplicationIcon(ctx, id).Execute()

Gets application icon.



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
    resp, r, err := api_client.InventoryApi.GetApplicationIcon(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetApplicationIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationIcon`: ApplicationIconInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetApplicationIcon`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationIconInfo**](ApplicationIconInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationPool

> ApplicationPoolInfo GetApplicationPool(ctx, id).Execute()

Gets application pool.



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
    resp, r, err := api_client.InventoryApi.GetApplicationPool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetApplicationPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationPool`: ApplicationPoolInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetApplicationPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationPoolInfo**](ApplicationPoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationPoolV2

> ApplicationPoolInfoV2 GetApplicationPoolV2(ctx, id).Execute()

Gets application pool.



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
    resp, r, err := api_client.InventoryApi.GetApplicationPoolV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetApplicationPoolV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationPoolV2`: ApplicationPoolInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetApplicationPoolV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPoolV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationPoolInfoV2**](ApplicationPoolInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationPoolV3

> ApplicationPoolInfoV3 GetApplicationPoolV3(ctx, id).Execute()

Gets application pool.



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
    resp, r, err := api_client.InventoryApi.GetApplicationPoolV3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetApplicationPoolV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApplicationPoolV3`: ApplicationPoolInfoV3
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetApplicationPoolV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPoolV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationPoolInfoV3**](ApplicationPoolInfoV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPool

> DesktopPoolInfo GetDesktopPool(ctx, id).Execute()

Gets the Desktop Pool information.



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
    resp, r, err := api_client.InventoryApi.GetDesktopPool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetDesktopPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDesktopPool`: DesktopPoolInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetDesktopPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DesktopPoolInfo**](DesktopPoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPoolTask

> DesktopPoolTaskInfo GetDesktopPoolTask(ctx, id, taskId).Execute()

Gets the task information on the desktop pool.



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
    id := "id_example" // string | Desktop pool ID
    taskId := "taskId_example" // string | Desktop pool task ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.GetDesktopPoolTask(context.Background(), id, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetDesktopPoolTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDesktopPoolTask`: DesktopPoolTaskInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetDesktopPoolTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Desktop pool ID | 
**taskId** | **string** | Desktop pool task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopPoolTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DesktopPoolTaskInfo**](DesktopPoolTaskInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPoolV2

> DesktopPoolInfoV2 GetDesktopPoolV2(ctx, id).Execute()

Gets the desktop pool information.



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
    resp, r, err := api_client.InventoryApi.GetDesktopPoolV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetDesktopPoolV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDesktopPoolV2`: DesktopPoolInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetDesktopPoolV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopPoolV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DesktopPoolInfoV2**](DesktopPoolInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPoolV3

> DesktopPoolInfoV3 GetDesktopPoolV3(ctx, id).Execute()

Gets the desktop pool information.



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
    resp, r, err := api_client.InventoryApi.GetDesktopPoolV3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetDesktopPoolV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDesktopPoolV3`: DesktopPoolInfoV3
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetDesktopPoolV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopPoolV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DesktopPoolInfoV3**](DesktopPoolInfoV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPoolV4

> DesktopPoolInfoV4 GetDesktopPoolV4(ctx, id).Execute()

Gets the desktop pool information.



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
    resp, r, err := api_client.InventoryApi.GetDesktopPoolV4(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetDesktopPoolV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDesktopPoolV4`: DesktopPoolInfoV4
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetDesktopPoolV4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopPoolV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DesktopPoolInfoV4**](DesktopPoolInfoV4.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDesktopPoolV5

> DesktopPoolInfoV5 GetDesktopPoolV5(ctx, id).Execute()

Gets the desktop pool information.



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
    resp, r, err := api_client.InventoryApi.GetDesktopPoolV5(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetDesktopPoolV5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDesktopPoolV5`: DesktopPoolInfoV5
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetDesktopPoolV5`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDesktopPoolV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DesktopPoolInfoV5**](DesktopPoolInfoV5.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFarm

> FarmInfo GetFarm(ctx, id).Execute()

Gets the Farm information.



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
    resp, r, err := api_client.InventoryApi.GetFarm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFarm`: FarmInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetFarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FarmInfo**](FarmInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFarmV2

> FarmInfoV2 GetFarmV2(ctx, id).Execute()

Gets the Farm information.



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
    resp, r, err := api_client.InventoryApi.GetFarmV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetFarmV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFarmV2`: FarmInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetFarmV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFarmV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FarmInfoV2**](FarmInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFarmV3

> FarmInfoV3 GetFarmV3(ctx, id).Execute()

Gets the Farm information.



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
    resp, r, err := api_client.InventoryApi.GetFarmV3(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetFarmV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFarmV3`: FarmInfoV3
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetFarmV3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFarmV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FarmInfoV3**](FarmInfoV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalApplicationEntitlement

> GlobalApplicationEntitlementInfo GetGlobalApplicationEntitlement(ctx, id).Execute()

Gets the Global Application Entitlement in the environment.



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
    resp, r, err := api_client.InventoryApi.GetGlobalApplicationEntitlement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetGlobalApplicationEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalApplicationEntitlement`: GlobalApplicationEntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetGlobalApplicationEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalApplicationEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalApplicationEntitlementInfo**](GlobalApplicationEntitlementInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalApplicationEntitlementV2

> GlobalApplicationEntitlementInfoV2 GetGlobalApplicationEntitlementV2(ctx, id).Execute()

Gets the Global Application Entitlement in the environment.



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
    resp, r, err := api_client.InventoryApi.GetGlobalApplicationEntitlementV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetGlobalApplicationEntitlementV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalApplicationEntitlementV2`: GlobalApplicationEntitlementInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetGlobalApplicationEntitlementV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalApplicationEntitlementV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalApplicationEntitlementInfoV2**](GlobalApplicationEntitlementInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalDesktopEntitlement

> GlobalDesktopEntitlementInfo GetGlobalDesktopEntitlement(ctx, id).Execute()

Gets the Global Desktop Entitlement in the environment.



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
    resp, r, err := api_client.InventoryApi.GetGlobalDesktopEntitlement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetGlobalDesktopEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalDesktopEntitlement`: GlobalDesktopEntitlementInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetGlobalDesktopEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalDesktopEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalDesktopEntitlementInfo**](GlobalDesktopEntitlementInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalDesktopEntitlementV2

> GlobalDesktopEntitlementInfoV2 GetGlobalDesktopEntitlementV2(ctx, id).Execute()

Gets the Global Desktop Entitlement in the environment.



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
    resp, r, err := api_client.InventoryApi.GetGlobalDesktopEntitlementV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetGlobalDesktopEntitlementV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalDesktopEntitlementV2`: GlobalDesktopEntitlementInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetGlobalDesktopEntitlementV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalDesktopEntitlementV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GlobalDesktopEntitlementInfoV2**](GlobalDesktopEntitlementInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachine

> MachineInfo GetMachine(ctx, id).Execute()

Gets the Machine information.



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
    resp, r, err := api_client.InventoryApi.GetMachine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachine`: MachineInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MachineInfo**](MachineInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMachineV2

> MachineInfoV2 GetMachineV2(ctx, id).Execute()

Gets the Machine information.



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
    resp, r, err := api_client.InventoryApi.GetMachineV2(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetMachineV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMachineV2`: MachineInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetMachineV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMachineV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MachineInfoV2**](MachineInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPhysicalMachine

> PhysicalMachineInfo GetPhysicalMachine(ctx, id).Execute()

Gets the Physical Machine information.



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
    id := "id_example" // string | Physical machine ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.GetPhysicalMachine(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetPhysicalMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPhysicalMachine`: PhysicalMachineInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetPhysicalMachine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Physical machine ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPhysicalMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PhysicalMachineInfo**](PhysicalMachineInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRDSServer

> RDSServerInfo GetRDSServer(ctx, id).Execute()

Gets the RDS Server information.



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
    id := "id_example" // string | RDS Server ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.GetRDSServer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetRDSServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRDSServer`: RDSServerInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetRDSServer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | RDS Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRDSServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RDSServerInfo**](RDSServerInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSessionInfo

> SessionInfo GetSessionInfo(ctx, id).Execute()

Gets the Session information for locally resourced session.



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
    resp, r, err := api_client.InventoryApi.GetSessionInfo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.GetSessionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSessionInfo`: SessionInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.GetSessionInfo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSessionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SessionInfo**](SessionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationIcons

> []ApplicationIconInfo ListApplicationIcons(ctx).ApplicationPoolId(applicationPoolId).Execute()

Lists the application icons for the given application pool.



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
    applicationPoolId := "applicationPoolId_example" // string | Application Pool ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListApplicationIcons(context.Background()).ApplicationPoolId(applicationPoolId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListApplicationIcons``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationIcons`: []ApplicationIconInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListApplicationIcons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationIconsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationPoolId** | **string** | Application Pool ID | 

### Return type

[**[]ApplicationIconInfo**](ApplicationIconInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationPools

> []ApplicationPoolInfo ListApplicationPools(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the application pools in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListApplicationPools(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListApplicationPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationPools`: []ApplicationPoolInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListApplicationPools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]ApplicationPoolInfo**](ApplicationPoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationPoolsV2

> []ApplicationPoolInfoV2 ListApplicationPoolsV2(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the application pools in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListApplicationPoolsV2(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListApplicationPoolsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationPoolsV2`: []ApplicationPoolInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListApplicationPoolsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationPoolsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]ApplicationPoolInfoV2**](ApplicationPoolInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListApplicationPoolsV3

> []ApplicationPoolInfoV3 ListApplicationPoolsV3(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the application pools in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListApplicationPoolsV3(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListApplicationPoolsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApplicationPoolsV3`: []ApplicationPoolInfoV3
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListApplicationPoolsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListApplicationPoolsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]ApplicationPoolInfoV3**](ApplicationPoolInfoV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompatibleBackupGAEs

> []string ListCompatibleBackupGAEs(ctx, id).Execute()

Lists the Global Application Entitlements that can be associated as backup Global Application Entitlement.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Global Application Entitlement ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListCompatibleBackupGAEs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListCompatibleBackupGAEs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompatibleBackupGAEs`: []string
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListCompatibleBackupGAEs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Global Application Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompatibleBackupGAEsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompatibleBackupGDEs

> []string ListCompatibleBackupGDEs(ctx, id).Execute()

Lists the Global Desktop Entitlements that can be associated as backup Global Desktop Entitlement.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Global Desktop Entitlement ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListCompatibleBackupGDEs(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListCompatibleBackupGDEs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompatibleBackupGDEs`: []string
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListCompatibleBackupGDEs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Global Desktop Entitlement ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompatibleBackupGDEsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompatibleLocalApplicationPools

> []string ListCompatibleLocalApplicationPools(ctx, id).Execute()

Lists Local Application Pools which are compatible with Global Application Entitlement.



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
    resp, r, err := api_client.InventoryApi.ListCompatibleLocalApplicationPools(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListCompatibleLocalApplicationPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompatibleLocalApplicationPools`: []string
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListCompatibleLocalApplicationPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompatibleLocalApplicationPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCompatibleLocalDesktopPools

> []string ListCompatibleLocalDesktopPools(ctx, id).Execute()

Lists Local Desktop Pools which are compatible with Global Desktop Entitlement.



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
    resp, r, err := api_client.InventoryApi.ListCompatibleLocalDesktopPools(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListCompatibleLocalDesktopPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompatibleLocalDesktopPools`: []string
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListCompatibleLocalDesktopPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompatibleLocalDesktopPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPoolTasks

> []DesktopPoolTaskInfo ListDesktopPoolTasks(ctx, id).Execute()

Lists the tasks on the desktop pool.



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
    resp, r, err := api_client.InventoryApi.ListDesktopPoolTasks(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListDesktopPoolTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDesktopPoolTasks`: []DesktopPoolTaskInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListDesktopPoolTasks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDesktopPoolTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DesktopPoolTaskInfo**](DesktopPoolTaskInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPools

> []DesktopPoolInfo ListDesktopPools(ctx).Execute()

Lists the Desktop Pools in the environment.



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
    resp, r, err := api_client.InventoryApi.ListDesktopPools(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListDesktopPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDesktopPools`: []DesktopPoolInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListDesktopPools`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDesktopPoolsRequest struct via the builder pattern


### Return type

[**[]DesktopPoolInfo**](DesktopPoolInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPoolsV2

> []DesktopPoolInfoV2 ListDesktopPoolsV2(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the desktop pools in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListDesktopPoolsV2(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListDesktopPoolsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDesktopPoolsV2`: []DesktopPoolInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListDesktopPoolsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDesktopPoolsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]DesktopPoolInfoV2**](DesktopPoolInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPoolsV3

> []DesktopPoolInfoV3 ListDesktopPoolsV3(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the desktop pools in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListDesktopPoolsV3(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListDesktopPoolsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDesktopPoolsV3`: []DesktopPoolInfoV3
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListDesktopPoolsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDesktopPoolsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]DesktopPoolInfoV3**](DesktopPoolInfoV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPoolsV4

> []DesktopPoolInfoV4 ListDesktopPoolsV4(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the desktop pools in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListDesktopPoolsV4(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListDesktopPoolsV4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDesktopPoolsV4`: []DesktopPoolInfoV4
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListDesktopPoolsV4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDesktopPoolsV4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]DesktopPoolInfoV4**](DesktopPoolInfoV4.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDesktopPoolsV5

> []DesktopPoolInfoV5 ListDesktopPoolsV5(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the desktop pools in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListDesktopPoolsV5(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListDesktopPoolsV5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDesktopPoolsV5`: []DesktopPoolInfoV5
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListDesktopPoolsV5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDesktopPoolsV5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]DesktopPoolInfoV5**](DesktopPoolInfoV5.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFarms

> []FarmInfo ListFarms(ctx).Execute()

Lists the Farms in the environment.



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
    resp, r, err := api_client.InventoryApi.ListFarms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListFarms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFarms`: []FarmInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListFarms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFarmsRequest struct via the builder pattern


### Return type

[**[]FarmInfo**](FarmInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFarmsV2

> []FarmInfoV2 ListFarmsV2(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Farms in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListFarmsV2(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListFarmsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFarmsV2`: []FarmInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListFarmsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFarmsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]FarmInfoV2**](FarmInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFarmsV3

> []FarmInfoV3 ListFarmsV3(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Farms in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListFarmsV3(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListFarmsV3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFarmsV3`: []FarmInfoV3
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListFarmsV3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFarmsV3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]FarmInfoV3**](FarmInfoV3.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalApplicationEntitlements

> []GlobalApplicationEntitlementSummary ListGlobalApplicationEntitlements(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Global Application Entitlements in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListGlobalApplicationEntitlements(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListGlobalApplicationEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlobalApplicationEntitlements`: []GlobalApplicationEntitlementSummary
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListGlobalApplicationEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalApplicationEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]GlobalApplicationEntitlementSummary**](GlobalApplicationEntitlementSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalApplicationEntitlementsV2

> []GlobalApplicationEntitlementSummaryV2 ListGlobalApplicationEntitlementsV2(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Global Application Entitlements in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListGlobalApplicationEntitlementsV2(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListGlobalApplicationEntitlementsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlobalApplicationEntitlementsV2`: []GlobalApplicationEntitlementSummaryV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListGlobalApplicationEntitlementsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalApplicationEntitlementsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]GlobalApplicationEntitlementSummaryV2**](GlobalApplicationEntitlementSummaryV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalDesktopEntitlements

> []GlobalDesktopEntitlementSummary ListGlobalDesktopEntitlements(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Global Desktop Entitlements in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListGlobalDesktopEntitlements(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListGlobalDesktopEntitlements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlobalDesktopEntitlements`: []GlobalDesktopEntitlementSummary
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListGlobalDesktopEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalDesktopEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]GlobalDesktopEntitlementSummary**](GlobalDesktopEntitlementSummary.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGlobalDesktopEntitlementsV2

> []GlobalDesktopEntitlementSummaryV2 ListGlobalDesktopEntitlementsV2(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Global Desktop Entitlements in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListGlobalDesktopEntitlementsV2(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListGlobalDesktopEntitlementsV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGlobalDesktopEntitlementsV2`: []GlobalDesktopEntitlementSummaryV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListGlobalDesktopEntitlementsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGlobalDesktopEntitlementsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]GlobalDesktopEntitlementSummaryV2**](GlobalDesktopEntitlementSummaryV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledApplicationsOnDesktopPool

> []InstalledApplicationInfo ListInstalledApplicationsOnDesktopPool(ctx, id).Execute()

Lists the installed applications on the given desktop pool.



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
    resp, r, err := api_client.InventoryApi.ListInstalledApplicationsOnDesktopPool(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListInstalledApplicationsOnDesktopPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstalledApplicationsOnDesktopPool`: []InstalledApplicationInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListInstalledApplicationsOnDesktopPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstalledApplicationsOnDesktopPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InstalledApplicationInfo**](InstalledApplicationInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInstalledApplicationsOnFarm

> []InstalledApplicationInfo ListInstalledApplicationsOnFarm(ctx, id).Execute()

Lists the installed applications on the given farm.



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
    resp, r, err := api_client.InventoryApi.ListInstalledApplicationsOnFarm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListInstalledApplicationsOnFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInstalledApplicationsOnFarm`: []InstalledApplicationInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListInstalledApplicationsOnFarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInstalledApplicationsOnFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]InstalledApplicationInfo**](InstalledApplicationInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocalApplicationPools

> []string ListLocalApplicationPools(ctx, id).Execute()

Lists Local Application Pools which are assigned to Global Application Entitlement.



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
    resp, r, err := api_client.InventoryApi.ListLocalApplicationPools(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListLocalApplicationPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocalApplicationPools`: []string
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListLocalApplicationPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocalApplicationPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLocalDesktopPools

> []string ListLocalDesktopPools(ctx, id).Execute()

Lists Local Desktop Pools which are assigned to Global Desktop Entitlement.



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
    resp, r, err := api_client.InventoryApi.ListLocalDesktopPools(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListLocalDesktopPools``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLocalDesktopPools`: []string
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListLocalDesktopPools`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLocalDesktopPoolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMachines

> []MachineInfo ListMachines(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Machines in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListMachines(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMachines`: []MachineInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]MachineInfo**](MachineInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMachinesV2

> []MachineInfoV2 ListMachinesV2(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Machines in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListMachinesV2(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListMachinesV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMachinesV2`: []MachineInfoV2
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListMachinesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMachinesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]MachineInfoV2**](MachineInfoV2.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPhysicalMachines

> []PhysicalMachineInfo ListPhysicalMachines(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the Physical Machines in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListPhysicalMachines(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListPhysicalMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPhysicalMachines`: []PhysicalMachineInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListPhysicalMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPhysicalMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]PhysicalMachineInfo**](PhysicalMachineInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRDSServers

> []RDSServerInfo ListRDSServers(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the RDS Servers in the environment.



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListRDSServers(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListRDSServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRDSServers`: []RDSServerInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListRDSServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRDSServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]RDSServerInfo**](RDSServerInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSessionInfo

> []SessionInfo ListSessionInfo(ctx).Filter(filter).Page(page).Size(size).Execute()

Lists the locally resourced Sessions in the environment



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
    size := int32(10) // int32 | size, if passed should be > 0. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ListSessionInfo(context.Background()).Filter(filter).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ListSessionInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSessionInfo`: []SessionInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ListSessionInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSessionInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | filter expression built using fields with &lt;b&gt;&#39;supported filters&#39;&lt;/b&gt; as described in output &lt;b&gt;model&lt;/b&gt; schema of this API. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 

### Return type

[**[]SessionInfo**](SessionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogOffGlobalSessions

> []BulkGlobalSessionActionResponseInfo LogOffGlobalSessions(ctx).GlobalSessionActionSpecs(globalSessionActionSpecs).Forced(forced).Execute()

Logs off global sessions in the environment.



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
    globalSessionActionSpecs := []openapiclient.GlobalSessionActionSpec{*openapiclient.NewGlobalSessionActionSpec([]string{"Ids_example"}, "54c122e0-d6b5-476c-af91-fd00392e5983")} // []GlobalSessionActionSpec | Sessions which are to be logged off.
    forced := true // bool | Indicates whether sessions should be logged off forcibly.  If set to true, sessions would be logged off forcibly even if they are locked.  If set to false, sessions would be logged off if they are not locked.   (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.LogOffGlobalSessions(context.Background()).GlobalSessionActionSpecs(globalSessionActionSpecs).Forced(forced).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.LogOffGlobalSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogOffGlobalSessions`: []BulkGlobalSessionActionResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.LogOffGlobalSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogOffGlobalSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalSessionActionSpecs** | [**[]GlobalSessionActionSpec**](GlobalSessionActionSpec.md) | Sessions which are to be logged off. | 
 **forced** | **bool** | Indicates whether sessions should be logged off forcibly.  If set to true, sessions would be logged off forcibly even if they are locked.  If set to false, sessions would be logged off if they are not locked.   | [default to false]

### Return type

[**[]BulkGlobalSessionActionResponseInfo**](BulkGlobalSessionActionResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogOffSessions

> []BulkItemResponseInfo LogOffSessions(ctx).Body(body).Forced(forced).Execute()

Logs off locally resourced user sessions, if they are not locked.



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
    body := []string{"Property_example"} // []string | List of session ids to be logged off.
    forced := false // bool | Indicates to Log off session forcibly.  If passed as \"true\", then sessions are logoff forcibly, even if they are locked.  If passed as \"false\" or not passed at all, then sessions will be normally logged off, if they are not locked. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.LogOffSessions(context.Background()).Body(body).Forced(forced).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.LogOffSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LogOffSessions`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.LogOffSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogOffSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of session ids to be logged off. | 
 **forced** | **bool** | Indicates to Log off session forcibly.  If passed as \&quot;true\&quot;, then sessions are logoff forcibly, even if they are locked.  If passed as \&quot;false\&quot; or not passed at all, then sessions will be normally logged off, if they are not locked. | [default to false]

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PauseDesktopPoolTask

> PauseDesktopPoolTask(ctx, id, taskId).Execute()

Pause the instant clone desktop pool push image task.



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
    id := "id_example" // string | Desktop pool ID
    taskId := "taskId_example" // string | Desktop pool task ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.PauseDesktopPoolTask(context.Background(), id, taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.PauseDesktopPoolTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Desktop pool ID | 
**taskId** | **string** | Desktop pool task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiPauseDesktopPoolTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## QueryGlobalSessions

> []GlobalSessionInfo QueryGlobalSessions(ctx).BrokeringPodId(brokeringPodId).Page(page).PodId(podId).Size(size).UserId(userId).Execute()

Lists global sessions in the environment for the given user, pod or brokering pod.



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
    brokeringPodId := "806ca35f-bcab-4a42-9006-04ba27cad900" // string | ID of the pod that brokered the session. (optional)
    page := int32(1) // int32 | page, if passed should be > 0. (optional)
    podId := "96d708fe-badb-40f8-9b80-9029b0ab2ef2" // string | ID of the pod that hosted the session. (optional)
    size := int32(10) // int32 | size, if passed should be > 0. (optional)
    userId := "S-1-5-21-1111111111-2222222222-3333333333-500" // string | SID of the user. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.QueryGlobalSessions(context.Background()).BrokeringPodId(brokeringPodId).Page(page).PodId(podId).Size(size).UserId(userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.QueryGlobalSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `QueryGlobalSessions`: []GlobalSessionInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.QueryGlobalSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiQueryGlobalSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brokeringPodId** | **string** | ID of the pod that brokered the session. | 
 **page** | **int32** | page, if passed should be &gt; 0. | 
 **podId** | **string** | ID of the pod that hosted the session. | 
 **size** | **int32** | size, if passed should be &gt; 0. | 
 **userId** | **string** | SID of the user. | 

### Return type

[**[]GlobalSessionInfo**](GlobalSessionInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RebuildMachines

> []BulkItemResponseInfo RebuildMachines(ctx).Body(body).Execute()

Rebuilds the specified machines.



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
    body := []string{"Property_example"} // []string | List of Machine Ids representing the machines to be rebuilt.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RebuildMachines(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RebuildMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RebuildMachines`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RebuildMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRebuildMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of Machine Ids representing the machines to be rebuilt. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoverMachines

> []BulkItemResponseInfo RecoverMachines(ctx).Body(body).Execute()

Recovers the specified machines.



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
    body := []string{"Property_example"} // []string | List of Machine Ids representing the machines to be recovered.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RecoverMachines(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RecoverMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoverMachines`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RecoverMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecoverMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of Machine Ids representing the machines to be recovered. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecoverRDSServers

> []BulkItemResponseInfo RecoverRDSServers(ctx).Body(body).Execute()

Recovers the specified RDS Servers.



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
    body := []string{"Property_example"} // []string | List of RDS Server Ids representing the RDS Servers to be recovered.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RecoverRDSServers(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RecoverRDSServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RecoverRDSServers`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RecoverRDSServers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecoverRDSServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of RDS Server Ids representing the RDS Servers to be recovered. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterPhysicalMachine

> PhysicalMachineRegisterInfo RegisterPhysicalMachine(ctx).Body(body).Execute()

Registers the Physical Machine.



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
    body := *openapiclient.NewPhysicalMachineRegisterSpec("machine1.example.com", "WINDOWS_10") // PhysicalMachineRegisterSpec | The specification for registering the physical machine.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RegisterPhysicalMachine(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RegisterPhysicalMachine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterPhysicalMachine`: PhysicalMachineRegisterInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RegisterPhysicalMachine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterPhysicalMachineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PhysicalMachineRegisterSpec**](PhysicalMachineRegisterSpec.md) | The specification for registering the physical machine. | 

### Return type

[**PhysicalMachineRegisterInfo**](PhysicalMachineRegisterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterRDSServer

> RDSServerRegisterInfo RegisterRDSServer(ctx).Body(body).Execute()

Registers the RDS Server.



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
    body := *openapiclient.NewRDSServerRegisterSpec("rdsserver1.example.com", "WINDOWS_SERVER_2012") // RDSServerRegisterSpec | The specification for registering the RDS Server.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RegisterRDSServer(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RegisterRDSServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterRDSServer`: RDSServerRegisterInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RegisterRDSServer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterRDSServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RDSServerRegisterSpec**](RDSServerRegisterSpec.md) | The specification for registering the RDS Server. | 

### Return type

[**RDSServerRegisterInfo**](RDSServerRegisterInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomIcon

> RemoveCustomIcon(ctx, id).Execute()

Removes the associated custom icon from the application pool.



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
    id := "id_example" // string | Application pool ID

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RemoveCustomIcon(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RemoveCustomIcon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Application pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCustomIconRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLocalApplicationPoolsFromGAE

> []BulkItemResponseInfo RemoveLocalApplicationPoolsFromGAE(ctx, id).Body(body).Execute()

Removes Local Application Pools from Global Application Entitlement.



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
    body := []string{"Property_example"} // []string | List of local application pool ids to be removed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RemoveLocalApplicationPoolsFromGAE(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RemoveLocalApplicationPoolsFromGAE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveLocalApplicationPoolsFromGAE`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RemoveLocalApplicationPoolsFromGAE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLocalApplicationPoolsFromGAERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of local application pool ids to be removed. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLocalDesktopPoolsFromGDE

> []BulkItemResponseInfo RemoveLocalDesktopPoolsFromGDE(ctx, id).Body(body).Execute()

Removes Local Desktop Pools from Global Desktop Entitlement.



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
    body := []string{"Property_example"} // []string | List of local desktop pool ids to be removed.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RemoveLocalDesktopPoolsFromGDE(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RemoveLocalDesktopPoolsFromGDE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveLocalDesktopPoolsFromGDE`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RemoveLocalDesktopPoolsFromGDE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLocalDesktopPoolsFromGDERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of local desktop pool ids to be removed. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMachines

> []BulkItemResponseInfo RemoveMachines(ctx, id).Body(body).Execute()

Removes machines from the given manual desktop pool.



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
    body := []string{"Property_example"} // []string | List of Machine Ids representing the machines to be removed from the desktop pool.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RemoveMachines(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RemoveMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveMachines`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RemoveMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of Machine Ids representing the machines to be removed from the desktop pool. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveRdsServers

> []BulkItemResponseInfo RemoveRdsServers(ctx, id).Body(body).Execute()

Remove RDS servers from the specified farm.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Farm ID
    body := []string{"Property_example"} // []string | List of RDS server IDs representing the RDS servers to be removed from the farm.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RemoveRdsServers(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RemoveRdsServers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveRdsServers`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RemoveRdsServers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Farm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRdsServersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of RDS server IDs representing the RDS servers to be removed from the farm. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetGlobalSessions

> []BulkGlobalSessionActionResponseInfo ResetGlobalSessions(ctx).GlobalSessionActionSpecs(globalSessionActionSpecs).Execute()

Resets machines of global sessions in the environment.



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
    globalSessionActionSpecs := []openapiclient.GlobalSessionActionSpec{*openapiclient.NewGlobalSessionActionSpec([]string{"Ids_example"}, "54c122e0-d6b5-476c-af91-fd00392e5983")} // []GlobalSessionActionSpec | Sessions which are to be reset.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ResetGlobalSessions(context.Background()).GlobalSessionActionSpecs(globalSessionActionSpecs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ResetGlobalSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetGlobalSessions`: []BulkGlobalSessionActionResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ResetGlobalSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetGlobalSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalSessionActionSpecs** | [**[]GlobalSessionActionSpec**](GlobalSessionActionSpec.md) | Sessions which are to be reset. | 

### Return type

[**[]BulkGlobalSessionActionResponseInfo**](BulkGlobalSessionActionResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetMachines

> []BulkItemResponseInfo ResetMachines(ctx).Body(body).Execute()

Resets the specified machines.



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
    body := []string{"Property_example"} // []string | List of Machine Ids representing the machines to be reset.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ResetMachines(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ResetMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetMachines`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ResetMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of Machine Ids representing the machines to be reset. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetSessions

> []BulkItemResponseInfo ResetSessions(ctx).Body(body).Execute()

Resets machine of locally resourced user sessions. The machine must be managed by Virtual Center and the session cannot be an application or an RDS desktop session.



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
    body := []string{"Property_example"} // []string | List of session ids to be reset.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ResetSessions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ResetSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ResetSessions`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ResetSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of session ids to be reset. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartGlobalSessions

> []BulkGlobalSessionActionResponseInfo RestartGlobalSessions(ctx).GlobalSessionActionSpecs(globalSessionActionSpecs).Execute()

Restarts machines of global sessions in the environment.



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
    globalSessionActionSpecs := []openapiclient.GlobalSessionActionSpec{*openapiclient.NewGlobalSessionActionSpec([]string{"Ids_example"}, "54c122e0-d6b5-476c-af91-fd00392e5983")} // []GlobalSessionActionSpec | Sessions which are to be restarted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RestartGlobalSessions(context.Background()).GlobalSessionActionSpecs(globalSessionActionSpecs).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RestartGlobalSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartGlobalSessions`: []BulkGlobalSessionActionResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RestartGlobalSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestartGlobalSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **globalSessionActionSpecs** | [**[]GlobalSessionActionSpec**](GlobalSessionActionSpec.md) | Sessions which are to be restarted. | 

### Return type

[**[]BulkGlobalSessionActionResponseInfo**](BulkGlobalSessionActionResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartMachines

> []BulkItemResponseInfo RestartMachines(ctx).Body(body).Execute()

Restarts the specified machines.



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
    body := []string{"Property_example"} // []string | List of Machine Ids representing the machines to be restarted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RestartMachines(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RestartMachines``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartMachines`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RestartMachines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestartMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of Machine Ids representing the machines to be restarted. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestartSessions

> []BulkItemResponseInfo RestartSessions(ctx).Body(body).Execute()

Restarts machine of locally resourced user sessions. The machine must be managed by Virtual Center and the session cannot be an application or an RDS desktop session.



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
    body := []string{"Property_example"} // []string | List of session ids to be restarted.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.RestartSessions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.RestartSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RestartSessions`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.RestartSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRestartSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **[]string** | List of session ids to be restarted. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResumeDesktopPoolTask

> ResumeDesktopPoolTask(ctx, id, taskId).Body(body).Execute()

Resume the instant clone desktop pool push image task.



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
    id := "id_example" // string | Desktop pool ID
    taskId := "taskId_example" // string | Desktop pool task ID
    body := *openapiclient.NewResumeTaskSpec() // ResumeTaskSpec | Specification for resuming the task. If not provided, specification will be set to default values. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ResumeDesktopPoolTask(context.Background(), id, taskId).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ResumeDesktopPoolTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Desktop pool ID | 
**taskId** | **string** | Desktop pool task ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiResumeDesktopPoolTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**ResumeTaskSpec**](ResumeTaskSpec.md) | Specification for resuming the task. If not provided, specification will be set to default values. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleMaintenance

> ScheduleMaintenance(ctx, id).Body(body).Execute()

Creates maintenance schedule for the specified farm.



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
    id := "id_example" // string | Farm ID
    body := *openapiclient.NewFarmMaintenanceSpec("FORCE_LOGOFF", "IMMEDIATE") // FarmMaintenanceSpec | The specification to schedule maintenance.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ScheduleMaintenance(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ScheduleMaintenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Farm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiScheduleMaintenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FarmMaintenanceSpec**](FarmMaintenanceSpec.md) | The specification to schedule maintenance. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulePushImage

> SchedulePushImage(ctx, id).Body(body).Execute()

Schedule/reschedule a request to update the image in an instant clone desktop pool.



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
    body := *openapiclient.NewDesktopPoolPushImageSpec("WAIT_FOR_LOGOFF") // DesktopPoolPushImageSpec | Specification for the push image operation.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.SchedulePushImage(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.SchedulePushImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiSchedulePushImageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DesktopPoolPushImageSpec**](DesktopPoolPushImageSpec.md) | Specification for the push image operation. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessageToGlobalSessions

> []BulkGlobalSessionActionResponseInfo SendMessageToGlobalSessions(ctx).Body(body).Execute()

Sends message to global sessions in the environment.



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
    body := *openapiclient.NewGlobalSessionSendMessageSpec([]openapiclient.GlobalSessionActionSpec{*openapiclient.NewGlobalSessionActionSpec([]string{"Ids_example"}, "54c122e0-d6b5-476c-af91-fd00392e5983")}, "Sample Info Message", "INFO") // GlobalSessionSendMessageSpec | Global sessions to which message is to be sent.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.SendMessageToGlobalSessions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.SendMessageToGlobalSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendMessageToGlobalSessions`: []BulkGlobalSessionActionResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.SendMessageToGlobalSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageToGlobalSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GlobalSessionSendMessageSpec**](GlobalSessionSendMessageSpec.md) | Global sessions to which message is to be sent. | 

### Return type

[**[]BulkGlobalSessionActionResponseInfo**](BulkGlobalSessionActionResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SendMessageToSessions

> []BulkItemResponseInfo SendMessageToSessions(ctx).Body(body).Execute()

Sends the message to locally resourced user sessions.



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
    body := *openapiclient.NewSessionSendMessageSpec("Sample Info Message", "INFO", []string{"SessionIds_example"}) // SessionSendMessageSpec | Message information object to be sent to sessions.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.SendMessageToSessions(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.SendMessageToSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendMessageToSessions`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.SendMessageToSessions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageToSessionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SessionSendMessageSpec**](SessionSendMessageSpec.md) | Message information object to be sent to sessions. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignMachineAliases

> []BulkItemResponseInfo UnassignMachineAliases(ctx, id).Body(body).Execute()

Un-assigns the aliases for the specified users from the machine.



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
    body := []string{"Property_example"} // []string | List of User SIDs whose aliases will be un-assigned from the machine.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UnassignMachineAliases(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UnassignMachineAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignMachineAliases`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.UnassignMachineAliases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignMachineAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of User SIDs whose aliases will be un-assigned from the machine. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnassignUsers

> []BulkItemResponseInfo UnassignUsers(ctx, id).Body(body).Execute()

Un-assigns the specified users from the machine.



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
    body := []string{"Property_example"} // []string | List of User SIDs representing the users to be un-assigned from the machine.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UnassignUsers(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UnassignUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnassignUsers`: []BulkItemResponseInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.UnassignUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnassignUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of User SIDs representing the users to be un-assigned from the machine. | 

### Return type

[**[]BulkItemResponseInfo**](BulkItemResponseInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationPool

> UpdateApplicationPool(ctx, id).Body(body).Execute()

Updates application pool.



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
    body := *openapiclient.NewApplicationPoolUpdateSpec(false, true, "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Firefox.lnk") // ApplicationPoolUpdateSpec | Application pool object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UpdateApplicationPool(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateApplicationPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationPoolUpdateSpec**](ApplicationPoolUpdateSpec.md) | Application pool object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApplicationPoolV2

> UpdateApplicationPoolV2(ctx, id).Body(body).Execute()

Updates application pool.



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
    body := *openapiclient.NewApplicationPoolUpdateSpecV2(false, true, "C:\ProgramData\Microsoft\Windows\Start Menu\Programs\Firefox.lnk") // ApplicationPoolUpdateSpecV2 | Application pool object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UpdateApplicationPoolV2(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateApplicationPoolV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApplicationPoolV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ApplicationPoolUpdateSpecV2**](ApplicationPoolUpdateSpecV2.md) | Application pool object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDesktopPool

> UpdateDesktopPool(ctx, id).Body(body).Execute()

Updates desktop pool.



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
    body := *openapiclient.NewDesktopPoolUpdateSpec(false, false, false, true, "pool", false, true) // DesktopPoolUpdateSpec | Desktop Pool object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UpdateDesktopPool(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateDesktopPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDesktopPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DesktopPoolUpdateSpec**](DesktopPoolUpdateSpec.md) | Desktop Pool object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFarm

> UpdateFarm(ctx, id).Body(body).Execute()

Updates farm.



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
    body := *openapiclient.NewFarmUpdateSpec("6fd4638a-381f-4518-aed6-042aa3d9f14c", "ManualFarm", *openapiclient.NewFarmDisplayProtocolSettingsUpdateSpec(true, "PCOIP", false), true, int32(0), *openapiclient.NewFarmSessionSettingsUpdateSpec("NEVER", "AFTER"), false) // FarmUpdateSpec | Farm object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UpdateFarm(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FarmUpdateSpec**](FarmUpdateSpec.md) | Farm object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFarmV2

> UpdateFarmV2(ctx, id).Body(body).Execute()

Updates farm.



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
    body := *openapiclient.NewFarmUpdateSpecV2("6fd4638a-381f-4518-aed6-042aa3d9f14c", "ManualFarm", *openapiclient.NewFarmDisplayProtocolSettingsUpdateSpec(true, "PCOIP", false), true, int32(0), *openapiclient.NewFarmSessionSettingsUpdateSpec("NEVER", "AFTER"), false) // FarmUpdateSpecV2 | Farm object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UpdateFarmV2(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateFarmV2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFarmV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**FarmUpdateSpecV2**](FarmUpdateSpecV2.md) | Farm object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalApplicationEntitlement

> UpdateGlobalApplicationEntitlement(ctx, id).Body(body).Execute()

Updates a Global Application Entitlement.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Global Application Entitlement object to be updated
    body := *openapiclient.NewGlobalApplicationEntitlementUpdateSpec("PCOIP", "global-application-entitlement", false, false, true, "310927ee-1195-44c3-b877-5b194b4b4ea8", "ENABLED_ENFORCED", true, "global-application-entitlement", false, "ALL_SITES", false) // GlobalApplicationEntitlementUpdateSpec | Global Application Entitlement object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UpdateGlobalApplicationEntitlement(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateGlobalApplicationEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Global Application Entitlement object to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalApplicationEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GlobalApplicationEntitlementUpdateSpec**](GlobalApplicationEntitlementUpdateSpec.md) | Global Application Entitlement object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalDesktopEntitlement

> UpdateGlobalDesktopEntitlement(ctx, id).Body(body).Execute()

Updates a Global Desktop Entitlement.



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
    id := "5134796a-322g-5fe5-343f-4daa5d25ebfe" // string | Global Desktop Entitlement object to be updated
    body := *openapiclient.NewGlobalDesktopEntitlementUpdateSpec(true, false, "PCOIP", false, false, "global-desktop-entitlement", false, true, "310927ee-1195-44c3-b877-5b194b4b4ea8", false, "global-desktop-entitlement", false, "ALL_SITES", false, false) // GlobalDesktopEntitlementUpdateSpec | Global Desktop Entitlement object to be updated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UpdateGlobalDesktopEntitlement(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateGlobalDesktopEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Global Desktop Entitlement object to be updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalDesktopEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**GlobalDesktopEntitlementUpdateSpec**](GlobalDesktopEntitlementUpdateSpec.md) | Global Desktop Entitlement object to be updated. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRDSServer

> UpdateRDSServer(ctx, id).Body(body).Execute()

Updates the RDS Server.



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
    id := "id_example" // string | RDS Server ID
    body := *openapiclient.NewRDSServerUpdateSpec(true, "LIMITED") // RDSServerUpdateSpec | The specification for updating the RDS Server.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.UpdateRDSServer(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.UpdateRDSServer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | RDS Server ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRDSServerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**RDSServerUpdateSpec**](RDSServerUpdateSpec.md) | The specification for updating the RDS Server. | 

### Return type

 (empty response body)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateInstalledApplicationsOnDesktopPool

> []InstalledApplicationValidationInfo ValidateInstalledApplicationsOnDesktopPool(ctx, id).Body(body).Execute()

Validates that each application in the given list is installed on the machines belonging to the specified desktop pool.



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
    id := "id_example" // string | Desktop Pool ID
    body := []string{"Property_example"} // []string | List of application exe paths for the applications to be validated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ValidateInstalledApplicationsOnDesktopPool(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ValidateInstalledApplicationsOnDesktopPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateInstalledApplicationsOnDesktopPool`: []InstalledApplicationValidationInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ValidateInstalledApplicationsOnDesktopPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Desktop Pool ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateInstalledApplicationsOnDesktopPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of application exe paths for the applications to be validated. | 

### Return type

[**[]InstalledApplicationValidationInfo**](InstalledApplicationValidationInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateInstalledApplicationsOnFarm

> []InstalledApplicationValidationInfo ValidateInstalledApplicationsOnFarm(ctx, id).Body(body).Execute()

Validates that each application in the given list is installed on the RDS Servers belonging to the specified Farm.



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
    id := "id_example" // string | Farm ID
    body := []string{"Property_example"} // []string | List of application exe paths for the applications to be validated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ValidateInstalledApplicationsOnFarm(context.Background(), id).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ValidateInstalledApplicationsOnFarm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateInstalledApplicationsOnFarm`: []InstalledApplicationValidationInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ValidateInstalledApplicationsOnFarm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Farm ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateInstalledApplicationsOnFarmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **[]string** | List of application exe paths for the applications to be validated. | 

### Return type

[**[]InstalledApplicationValidationInfo**](InstalledApplicationValidationInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateSpecifiedNames

> []SpecifiedNamesValidationInfo ValidateSpecifiedNames(ctx).Body(body).Execute()

Validates manually specified virtual machines. Ensures machine and user names are valid and aren't duplicated in the given desktop pool.



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
    body := *openapiclient.NewSpecifiedNamesValidationSpec([]openapiclient.NamesSpec{*openapiclient.NewNamesSpec("Machine-002-jdoe")}) // SpecifiedNamesValidationSpec | The specification for specified names validation that needs to be validated.

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.InventoryApi.ValidateSpecifiedNames(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InventoryApi.ValidateSpecifiedNames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateSpecifiedNames`: []SpecifiedNamesValidationInfo
    fmt.Fprintf(os.Stdout, "Response from `InventoryApi.ValidateSpecifiedNames`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateSpecifiedNamesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SpecifiedNamesValidationSpec**](SpecifiedNamesValidationSpec.md) | The specification for specified names validation that needs to be validated. | 

### Return type

[**[]SpecifiedNamesValidationInfo**](SpecifiedNamesValidationInfo.md)

### Authorization

[Bearer](../README.md#Bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

