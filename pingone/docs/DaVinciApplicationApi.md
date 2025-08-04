# \DaVinciApplicationApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDavinciApplication**](DaVinciApplicationApi.md#CreateDavinciApplication) | **Post** /environments/{environmentID}/davinciApplications | 
[**CreateFlowPolicyByDavinciApplicationId**](DaVinciApplicationApi.md#CreateFlowPolicyByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies | 
[**CreateKeyByDavinciApplicationId**](DaVinciApplicationApi.md#CreateKeyByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/key | 
[**CreateSecretByDavinciApplicationId**](DaVinciApplicationApi.md#CreateSecretByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/secret | 
[**DeleteDavinciApplicationById**](DaVinciApplicationApi.md#DeleteDavinciApplicationById) | **Delete** /environments/{environmentID}/davinciApplications/{davinciApplicationID} | 
[**DeleteFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationApi.md#DeleteFlowPolicyByIdUsingDavinciApplicationId) | **Delete** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies/{flowPolicyID} | 
[**GetDavinciApplicationById**](DaVinciApplicationApi.md#GetDavinciApplicationById) | **Get** /environments/{environmentID}/davinciApplications/{davinciApplicationID} | 
[**GetDavinciApplications**](DaVinciApplicationApi.md#GetDavinciApplications) | **Get** /environments/{environmentID}/davinciApplications | 
[**GetEventsByDavinciApplicationIdAndFlowPolicyId**](DaVinciApplicationApi.md#GetEventsByDavinciApplicationIdAndFlowPolicyId) | **Get** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies/{flowPolicyID}/events | 
[**GetFlowPoliciesByDavinciApplicationId**](DaVinciApplicationApi.md#GetFlowPoliciesByDavinciApplicationId) | **Get** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies | 
[**GetFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationApi.md#GetFlowPolicyByIdUsingDavinciApplicationId) | **Get** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies/{flowPolicyID} | 
[**ReplaceDavinciApplicationById**](DaVinciApplicationApi.md#ReplaceDavinciApplicationById) | **Put** /environments/{environmentID}/davinciApplications/{davinciApplicationID} | 
[**ReplaceFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationApi.md#ReplaceFlowPolicyByIdUsingDavinciApplicationId) | **Put** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies/{flowPolicyID} | 



## CreateDavinciApplication

> DaVinciApplication CreateDavinciApplication(ctx, environmentID).DaVinciApplicationCreateRequest(daVinciApplicationCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	daVinciApplicationCreateRequest := *openapiclient.NewDaVinciApplicationCreateRequest("Name_example") // DaVinciApplicationCreateRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.CreateDavinciApplication(context.Background(), environmentID).DaVinciApplicationCreateRequest(daVinciApplicationCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.CreateDavinciApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDavinciApplication`: DaVinciApplication
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.CreateDavinciApplication`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:create:applications`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDavinciApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daVinciApplicationCreateRequest** | [**DaVinciApplicationCreateRequest**](DaVinciApplicationCreateRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlowPolicyByDavinciApplicationId

> DaVinciFlowPolicyResponse CreateFlowPolicyByDavinciApplicationId(ctx, environmentID, davinciApplicationID).DaVinciFlowPolicyCreateRequest(daVinciFlowPolicyCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	daVinciFlowPolicyCreateRequest := *openapiclient.NewDaVinciFlowPolicyCreateRequest("Name_example") // DaVinciFlowPolicyCreateRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.CreateFlowPolicyByDavinciApplicationId(context.Background(), environmentID, davinciApplicationID).DaVinciFlowPolicyCreateRequest(daVinciFlowPolicyCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.CreateFlowPolicyByDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlowPolicyByDavinciApplicationId`: DaVinciFlowPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.CreateFlowPolicyByDavinciApplicationId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:create:flowPolicies`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowPolicyByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciFlowPolicyCreateRequest** | [**DaVinciFlowPolicyCreateRequest**](DaVinciFlowPolicyCreateRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowPolicyResponse**](DaVinciFlowPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKeyByDavinciApplicationId

> DaVinciApplication CreateKeyByDavinciApplicationId(ctx, environmentID, davinciApplicationID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.CreateKeyByDavinciApplicationId(context.Background(), environmentID, davinciApplicationID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.CreateKeyByDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateKeyByDavinciApplicationId`: DaVinciApplication
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.CreateKeyByDavinciApplicationId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:applications`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.pingidentity.davinciApplication.rotateKey+json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecretByDavinciApplicationId

> DaVinciApplication CreateSecretByDavinciApplicationId(ctx, environmentID, davinciApplicationID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.CreateSecretByDavinciApplicationId(context.Background(), environmentID, davinciApplicationID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.CreateSecretByDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecretByDavinciApplicationId`: DaVinciApplication
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.CreateSecretByDavinciApplicationId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:applications`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.pingidentity.davinciApplication.rotateSecret+json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDavinciApplicationById

> DeleteDavinciApplicationById(ctx, environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciApplicationApi.DeleteDavinciApplicationById(context.Background(), environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.DeleteDavinciApplicationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:delete:applications`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDavinciApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlowPolicyByIdUsingDavinciApplicationId

> DeleteFlowPolicyByIdUsingDavinciApplicationId(ctx, environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	flowPolicyID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciApplicationApi.DeleteFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.DeleteFlowPolicyByIdUsingDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:delete:flowPolicies`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 
**flowPolicyID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlowPolicyByIdUsingDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDavinciApplicationById

> DaVinciApplication GetDavinciApplicationById(ctx, environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetDavinciApplicationById(context.Background(), environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.GetDavinciApplicationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDavinciApplicationById`: DaVinciApplication
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.GetDavinciApplicationById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:applications`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDavinciApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDavinciApplications

> DaVinciApplicationCollection GetDavinciApplications(ctx, environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetDavinciApplications(context.Background(), environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.GetDavinciApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDavinciApplications`: DaVinciApplicationCollection
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.GetDavinciApplications`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:applications`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDavinciApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplicationCollection**](DaVinciApplicationCollection.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEventsByDavinciApplicationIdAndFlowPolicyId

> DaVinciFlowPolicyEventsCollection GetEventsByDavinciApplicationIdAndFlowPolicyId(ctx, environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	flowPolicyID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetEventsByDavinciApplicationIdAndFlowPolicyId(context.Background(), environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.GetEventsByDavinciApplicationIdAndFlowPolicyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsByDavinciApplicationIdAndFlowPolicyId`: DaVinciFlowPolicyEventsCollection
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.GetEventsByDavinciApplicationIdAndFlowPolicyId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:flowPolicies`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 
**flowPolicyID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsByDavinciApplicationIdAndFlowPolicyIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowPolicyEventsCollection**](DaVinciFlowPolicyEventsCollection.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowPoliciesByDavinciApplicationId

> DaVinciFlowPolicyCollection GetFlowPoliciesByDavinciApplicationId(ctx, environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetFlowPoliciesByDavinciApplicationId(context.Background(), environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.GetFlowPoliciesByDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowPoliciesByDavinciApplicationId`: DaVinciFlowPolicyCollection
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.GetFlowPoliciesByDavinciApplicationId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:flowPolicies`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowPoliciesByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowPolicyCollection**](DaVinciFlowPolicyCollection.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowPolicyByIdUsingDavinciApplicationId

> DaVinciFlowPolicyResponse GetFlowPolicyByIdUsingDavinciApplicationId(ctx, environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	flowPolicyID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.GetFlowPolicyByIdUsingDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowPolicyByIdUsingDavinciApplicationId`: DaVinciFlowPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.GetFlowPolicyByIdUsingDavinciApplicationId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:flowPolicies`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 
**flowPolicyID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowPolicyByIdUsingDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowPolicyResponse**](DaVinciFlowPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDavinciApplicationById

> DaVinciApplication ReplaceDavinciApplicationById(ctx, environmentID, davinciApplicationID).DaVinciApplicationReplaceRequest(daVinciApplicationReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	daVinciApplicationReplaceRequest := *openapiclient.NewDaVinciApplicationReplaceRequest("Name_example") // DaVinciApplicationReplaceRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.ReplaceDavinciApplicationById(context.Background(), environmentID, davinciApplicationID).DaVinciApplicationReplaceRequest(daVinciApplicationReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.ReplaceDavinciApplicationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceDavinciApplicationById`: DaVinciApplication
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.ReplaceDavinciApplicationById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:applications`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDavinciApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciApplicationReplaceRequest** | [**DaVinciApplicationReplaceRequest**](DaVinciApplicationReplaceRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceFlowPolicyByIdUsingDavinciApplicationId

> DaVinciFlowPolicyResponse ReplaceFlowPolicyByIdUsingDavinciApplicationId(ctx, environmentID, davinciApplicationID, flowPolicyID).DaVinciFlowPolicyReplaceRequest(daVinciFlowPolicyReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	environmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	davinciApplicationID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	flowPolicyID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	daVinciFlowPolicyReplaceRequest := *openapiclient.NewDaVinciFlowPolicyReplaceRequest([]openapiclient.DaVinciFlowPolicyReplaceRequestFlowDistribution{*openapiclient.NewDaVinciFlowPolicyReplaceRequestFlowDistribution("Id_example", float32(123))}) // DaVinciFlowPolicyReplaceRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.ReplaceFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, davinciApplicationID, flowPolicyID).DaVinciFlowPolicyReplaceRequest(daVinciFlowPolicyReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationApi.ReplaceFlowPolicyByIdUsingDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceFlowPolicyByIdUsingDavinciApplicationId`: DaVinciFlowPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationApi.ReplaceFlowPolicyByIdUsingDavinciApplicationId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:flowPolicies`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**davinciApplicationID** | **uuid.UUID** |  | 
**flowPolicyID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFlowPolicyByIdUsingDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **daVinciFlowPolicyReplaceRequest** | [**DaVinciFlowPolicyReplaceRequest**](DaVinciFlowPolicyReplaceRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowPolicyResponse**](DaVinciFlowPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

