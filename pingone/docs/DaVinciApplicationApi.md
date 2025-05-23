# \DaVinciApplicationApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDavinciApplication**](DaVinciApplicationApi.md#CreateDavinciApplication) | **Post** /environments/{environmentID}/davinciApplications | _TO_BE_DEFINED_
[**CreateFlowPolicyByDavinciApplicationId**](DaVinciApplicationApi.md#CreateFlowPolicyByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{daVinciApplicationID}/flowPolicies | _TO_BE_DEFINED_
[**CreateKeyByDavinciApplicationId**](DaVinciApplicationApi.md#CreateKeyByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{daVinciApplicationID}/key | _TO_BE_DEFINED_
[**CreateSecretByDavinciApplicationId**](DaVinciApplicationApi.md#CreateSecretByDavinciApplicationId) | **Post** /environments/{environmentID}/davinciApplications/{daVinciApplicationID}/secret | _TO_BE_DEFINED_
[**DeleteDavinciApplicationById**](DaVinciApplicationApi.md#DeleteDavinciApplicationById) | **Delete** /environments/{environmentID}/davinciApplications/{daVinciApplicationID} | _TO_BE_DEFINED_
[**DeleteFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationApi.md#DeleteFlowPolicyByIdUsingDavinciApplicationId) | **Delete** /environments/{environmentID}/davinciApplications/{daVinciApplicationID}/flowPolicies/{flowPolicyID} | _TO_BE_DEFINED_
[**GetDavinciApplicationById**](DaVinciApplicationApi.md#GetDavinciApplicationById) | **Get** /environments/{environmentID}/davinciApplications/{daVinciApplicationID} | _TO_BE_DEFINED_
[**GetDavinciApplications**](DaVinciApplicationApi.md#GetDavinciApplications) | **Get** /environments/{environmentID}/davinciApplications | _TO_BE_DEFINED_
[**GetFlowPoliciesByDavinciApplicationId**](DaVinciApplicationApi.md#GetFlowPoliciesByDavinciApplicationId) | **Get** /environments/{environmentID}/davinciApplications/{daVinciApplicationID}/flowPolicies | _TO_BE_DEFINED_
[**GetFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationApi.md#GetFlowPolicyByIdUsingDavinciApplicationId) | **Get** /environments/{environmentID}/davinciApplications/{daVinciApplicationID}/flowPolicies/{flowPolicyID} | _TO_BE_DEFINED_
[**ReplaceDavinciApplicationById**](DaVinciApplicationApi.md#ReplaceDavinciApplicationById) | **Put** /environments/{environmentID}/davinciApplications/{daVinciApplicationID} | _TO_BE_DEFINED_
[**ReplaceFlowPolicyByIdUsingDavinciApplicationId**](DaVinciApplicationApi.md#ReplaceFlowPolicyByIdUsingDavinciApplicationId) | **Put** /environments/{environmentID}/davinciApplications/{daVinciApplicationID}/flowPolicies/{flowPolicyID} | _TO_BE_DEFINED_



## CreateDavinciApplication

> DaVinciApplication CreateDavinciApplication(ctx, environmentID).DaVinciApplicationCreateRequest(daVinciApplicationCreateRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationCreateRequest := *openapiclient.NewDaVinciApplicationCreateRequest("Name_example") // DaVinciApplicationCreateRequest | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.CreateDavinciApplication(context.Background(), environmentID).DaVinciApplicationCreateRequest(daVinciApplicationCreateRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDavinciApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daVinciApplicationCreateRequest** | [**DaVinciApplicationCreateRequest**](DaVinciApplicationCreateRequest.md) |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFlowPolicyByDavinciApplicationId

> DaVinciFlowPolicyResponse CreateFlowPolicyByDavinciApplicationId(ctx, environmentID, daVinciApplicationID).DaVinciFlowPolicyCreateRequest(daVinciFlowPolicyCreateRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	daVinciFlowPolicyCreateRequest := *openapiclient.NewDaVinciFlowPolicyCreateRequest([]openapiclient.DaVinciFlowPolicyCreateRequestFlowDistribution{*openapiclient.NewDaVinciFlowPolicyCreateRequestFlowDistribution("Id_example", float32(123))}) // DaVinciFlowPolicyCreateRequest | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.CreateFlowPolicyByDavinciApplicationId(context.Background(), environmentID, daVinciApplicationID).DaVinciFlowPolicyCreateRequest(daVinciFlowPolicyCreateRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowPolicyByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciFlowPolicyCreateRequest** | [**DaVinciFlowPolicyCreateRequest**](DaVinciFlowPolicyCreateRequest.md) |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciFlowPolicyResponse**](DaVinciFlowPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKeyByDavinciApplicationId

> DaVinciApplication CreateKeyByDavinciApplicationId(ctx, environmentID, daVinciApplicationID).RequestBody(requestBody).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.CreateKeyByDavinciApplicationId(context.Background(), environmentID, daVinciApplicationID).RequestBody(requestBody).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateKeyByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.pingidentity.davinciApplication.rotateKey+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSecretByDavinciApplicationId

> DaVinciApplication CreateSecretByDavinciApplicationId(ctx, environmentID, daVinciApplicationID).RequestBody(requestBody).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.CreateSecretByDavinciApplicationId(context.Background(), environmentID, daVinciApplicationID).RequestBody(requestBody).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecretByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.pingidentity.davinciApplication.rotateSecret+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDavinciApplicationById

> DeleteDavinciApplicationById(ctx, environmentID, daVinciApplicationID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciApplicationApi.DeleteDavinciApplicationById(context.Background(), environmentID, daVinciApplicationID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDavinciApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlowPolicyByIdUsingDavinciApplicationId

> DeleteFlowPolicyByIdUsingDavinciApplicationId(ctx, environmentID, daVinciApplicationID, flowPolicyID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	flowPolicyID := "flowPolicyID_example" // uuid.UUID | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciApplicationApi.DeleteFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, daVinciApplicationID, flowPolicyID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 
**flowPolicyID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlowPolicyByIdUsingDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

 (empty response body)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDavinciApplicationById

> DaVinciApplication GetDavinciApplicationById(ctx, environmentID, daVinciApplicationID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetDavinciApplicationById(context.Background(), environmentID, daVinciApplicationID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDavinciApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDavinciApplications

> DaVinciApplicationCollection GetDavinciApplications(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetDavinciApplications(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDavinciApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciApplicationCollection**](DaVinciApplicationCollection.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowPoliciesByDavinciApplicationId

> DaVinciFlowPolicyCollection GetFlowPoliciesByDavinciApplicationId(ctx, environmentID, daVinciApplicationID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetFlowPoliciesByDavinciApplicationId(context.Background(), environmentID, daVinciApplicationID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowPoliciesByDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciFlowPolicyCollection**](DaVinciFlowPolicyCollection.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlowPolicyByIdUsingDavinciApplicationId

> DaVinciFlowPolicyResponse GetFlowPolicyByIdUsingDavinciApplicationId(ctx, environmentID, daVinciApplicationID, flowPolicyID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	flowPolicyID := "flowPolicyID_example" // uuid.UUID | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.GetFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, daVinciApplicationID, flowPolicyID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 
**flowPolicyID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowPolicyByIdUsingDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciFlowPolicyResponse**](DaVinciFlowPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceDavinciApplicationById

> DaVinciApplication ReplaceDavinciApplicationById(ctx, environmentID, daVinciApplicationID).DaVinciApplicationReplaceRequest(daVinciApplicationReplaceRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	daVinciApplicationReplaceRequest := *openapiclient.NewDaVinciApplicationReplaceRequest("Name_example") // DaVinciApplicationReplaceRequest | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.ReplaceDavinciApplicationById(context.Background(), environmentID, daVinciApplicationID).DaVinciApplicationReplaceRequest(daVinciApplicationReplaceRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceDavinciApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciApplicationReplaceRequest** | [**DaVinciApplicationReplaceRequest**](DaVinciApplicationReplaceRequest.md) |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciApplication**](DaVinciApplication.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceFlowPolicyByIdUsingDavinciApplicationId

> DaVinciFlowPolicyResponse ReplaceFlowPolicyByIdUsingDavinciApplicationId(ctx, environmentID, daVinciApplicationID, flowPolicyID).DaVinciFlowPolicyReplaceRequest(daVinciFlowPolicyReplaceRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

_TO_BE_DEFINED_



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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	daVinciApplicationID := "daVinciApplicationID_example" // uuid.UUID | 
	flowPolicyID := "flowPolicyID_example" // uuid.UUID | 
	daVinciFlowPolicyReplaceRequest := *openapiclient.NewDaVinciFlowPolicyReplaceRequest("Name_example") // DaVinciFlowPolicyReplaceRequest | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciApplicationApi.ReplaceFlowPolicyByIdUsingDavinciApplicationId(context.Background(), environmentID, daVinciApplicationID, flowPolicyID).DaVinciFlowPolicyReplaceRequest(daVinciFlowPolicyReplaceRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**daVinciApplicationID** | **uuid.UUID** |  | 
**flowPolicyID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFlowPolicyByIdUsingDavinciApplicationIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **daVinciFlowPolicyReplaceRequest** | [**DaVinciFlowPolicyReplaceRequest**](DaVinciFlowPolicyReplaceRequest.md) |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciFlowPolicyResponse**](DaVinciFlowPolicyResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

