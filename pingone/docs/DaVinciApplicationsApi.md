# \DaVinciApplicationsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDavinciApplication**](DaVinciApplicationsApi.md#CreateDavinciApplication) | **Post** /environments/{environmentID}/davinciApplications | 
[**CreateFlowPolicyByDavinciApplicationId**](DaVinciApplicationsApi.md#CreateFlowPolicyByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies | 
[**DeleteDavinciApplicationById**](DaVinciApplicationsApi.md#DeleteDavinciApplicationById) | **Delete** /environments/{environmentID}/davinciApplications/{davinciApplicationID} | 
[**DeleteFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationsApi.md#DeleteFlowPolicyByIdUsingDavinciApplicationId) | **Delete** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies/{flowPolicyID} | 
[**GetDavinciApplicationById**](DaVinciApplicationsApi.md#GetDavinciApplicationById) | **Get** /environments/{environmentID}/davinciApplications/{davinciApplicationID} | 
[**GetDavinciApplications**](DaVinciApplicationsApi.md#GetDavinciApplications) | **Get** /environments/{environmentID}/davinciApplications | 
[**GetEventsByDavinciApplicationIdAndFlowPolicyId**](DaVinciApplicationsApi.md#GetEventsByDavinciApplicationIdAndFlowPolicyId) | **Get** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies/{flowPolicyID}/events | 
[**GetFlowPoliciesByDavinciApplicationId**](DaVinciApplicationsApi.md#GetFlowPoliciesByDavinciApplicationId) | **Get** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies | 
[**GetFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationsApi.md#GetFlowPolicyByIdUsingDavinciApplicationId) | **Get** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies/{flowPolicyID} | 
[**ReplaceDavinciApplicationById**](DaVinciApplicationsApi.md#ReplaceDavinciApplicationById) | **Put** /environments/{environmentID}/davinciApplications/{davinciApplicationID} | 
[**ReplaceFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationsApi.md#ReplaceFlowPolicyByIdUsingDavinciApplicationId) | **Put** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/flowPolicies/{flowPolicyID} | 
[**RotateKeyByDavinciApplicationId**](DaVinciApplicationsApi.md#RotateKeyByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/key | 
[**RotateSecretByDavinciApplicationId**](DaVinciApplicationsApi.md#RotateSecretByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{davinciApplicationID}/secret | 



## CreateDavinciApplication

> DaVinciApplicationResponse CreateDavinciApplication(ctx, environmentID).DaVinciApplicationCreateRequest(daVinciApplicationCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	resp, r, err := apiClient.DaVinciApplicationsApi.CreateDavinciApplication(context.Background(), environmentID).DaVinciApplicationCreateRequest(daVinciApplicationCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.CreateDavinciApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDavinciApplication`: DaVinciApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.CreateDavinciApplication`: %v\n", resp)
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

[**DaVinciApplicationResponse**](DaVinciApplicationResponse.md)

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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	daVinciFlowPolicyCreateRequest := *openapiclient.NewDaVinciFlowPolicyCreateRequest([]openapiclient.DaVinciFlowPolicyCreateRequestFlowDistribution{*openapiclient.NewDaVinciFlowPolicyCreateRequestFlowDistribution("Id_example", float32(123))}) // DaVinciFlowPolicyCreateRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.CreateFlowPolicyByDavinciApplicationId(context.Background(), environmentID, davinciApplicationID).DaVinciFlowPolicyCreateRequest(daVinciFlowPolicyCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.CreateFlowPolicyByDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlowPolicyByDavinciApplicationId`: DaVinciFlowPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.CreateFlowPolicyByDavinciApplicationId`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 

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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciApplicationsApi.DeleteDavinciApplicationById(context.Background(), environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.DeleteDavinciApplicationById``: %v\n", err)
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
**davinciApplicationID** | **string** |  | 

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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	flowPolicyID := "flowPolicyID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciApplicationsApi.DeleteFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.DeleteFlowPolicyByIdUsingDavinciApplicationId``: %v\n", err)
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
**davinciApplicationID** | **string** |  | 
**flowPolicyID** | **string** |  | 

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

> DaVinciApplicationResponse GetDavinciApplicationById(ctx, environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.GetDavinciApplicationById(context.Background(), environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.GetDavinciApplicationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDavinciApplicationById`: DaVinciApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.GetDavinciApplicationById`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDavinciApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplicationResponse**](DaVinciApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDavinciApplications

> DaVinciApplicationCollectionResponse GetDavinciApplications(ctx, environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	resp, r, err := apiClient.DaVinciApplicationsApi.GetDavinciApplications(context.Background(), environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.GetDavinciApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDavinciApplications`: DaVinciApplicationCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.GetDavinciApplications`: %v\n", resp)
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

[**DaVinciApplicationCollectionResponse**](DaVinciApplicationCollectionResponse.md)

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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	flowPolicyID := "flowPolicyID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.GetEventsByDavinciApplicationIdAndFlowPolicyId(context.Background(), environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.GetEventsByDavinciApplicationIdAndFlowPolicyId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEventsByDavinciApplicationIdAndFlowPolicyId`: DaVinciFlowPolicyEventsCollection
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.GetEventsByDavinciApplicationIdAndFlowPolicyId`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 
**flowPolicyID** | **string** |  | 

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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.GetFlowPoliciesByDavinciApplicationId(context.Background(), environmentID, davinciApplicationID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.GetFlowPoliciesByDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowPoliciesByDavinciApplicationId`: DaVinciFlowPolicyCollection
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.GetFlowPoliciesByDavinciApplicationId`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 

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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	flowPolicyID := "flowPolicyID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.GetFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, davinciApplicationID, flowPolicyID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.GetFlowPolicyByIdUsingDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowPolicyByIdUsingDavinciApplicationId`: DaVinciFlowPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.GetFlowPolicyByIdUsingDavinciApplicationId`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 
**flowPolicyID** | **string** |  | 

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

> DaVinciApplicationResponse ReplaceDavinciApplicationById(ctx, environmentID, davinciApplicationID).DaVinciApplicationReplaceRequest(daVinciApplicationReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	daVinciApplicationReplaceRequest := *openapiclient.NewDaVinciApplicationReplaceRequest("Name_example") // DaVinciApplicationReplaceRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.ReplaceDavinciApplicationById(context.Background(), environmentID, davinciApplicationID).DaVinciApplicationReplaceRequest(daVinciApplicationReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.ReplaceDavinciApplicationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceDavinciApplicationById`: DaVinciApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.ReplaceDavinciApplicationById`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDavinciApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciApplicationReplaceRequest** | [**DaVinciApplicationReplaceRequest**](DaVinciApplicationReplaceRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplicationResponse**](DaVinciApplicationResponse.md)

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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	flowPolicyID := "flowPolicyID_example" // string | 
	daVinciFlowPolicyReplaceRequest := *openapiclient.NewDaVinciFlowPolicyReplaceRequest("Name_example") // DaVinciFlowPolicyReplaceRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.ReplaceFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, davinciApplicationID, flowPolicyID).DaVinciFlowPolicyReplaceRequest(daVinciFlowPolicyReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.ReplaceFlowPolicyByIdUsingDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceFlowPolicyByIdUsingDavinciApplicationId`: DaVinciFlowPolicyResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.ReplaceFlowPolicyByIdUsingDavinciApplicationId`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 
**flowPolicyID** | **string** |  | 

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


## RotateKeyByDavinciApplicationId

> DaVinciApplicationResponse RotateKeyByDavinciApplicationId(ctx, environmentID, davinciApplicationID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.RotateKeyByDavinciApplicationId(context.Background(), environmentID, davinciApplicationID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.RotateKeyByDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateKeyByDavinciApplicationId`: DaVinciApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.RotateKeyByDavinciApplicationId`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateKeyByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplicationResponse**](DaVinciApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.pingidentity.davinciApplication.rotateKey+json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RotateSecretByDavinciApplicationId

> DaVinciApplicationResponse RotateSecretByDavinciApplicationId(ctx, environmentID, davinciApplicationID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	davinciApplicationID := "davinciApplicationID_example" // string | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationsApi.RotateSecretByDavinciApplicationId(context.Background(), environmentID, davinciApplicationID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciApplicationsApi.RotateSecretByDavinciApplicationId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RotateSecretByDavinciApplicationId`: DaVinciApplicationResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciApplicationsApi.RotateSecretByDavinciApplicationId`: %v\n", resp)
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
**davinciApplicationID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRotateSecretByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciApplicationResponse**](DaVinciApplicationResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.pingidentity.davinciApplication.rotateSecret+json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

