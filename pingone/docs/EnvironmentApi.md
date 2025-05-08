# \EnvironmentApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](EnvironmentApi.md#CreateEnvironment) | **Post** /environments | Create Environment
[**DeleteEnvironmentById**](EnvironmentApi.md#DeleteEnvironmentById) | **Delete** /environments/{environmentID} | Delete Environment
[**GetBillOfMaterialsByEnvironmentId**](EnvironmentApi.md#GetBillOfMaterialsByEnvironmentId) | **Get** /environments/{environmentID}/billOfMaterials | _TO_BE_DEFINED_
[**GetEnvironmentById**](EnvironmentApi.md#GetEnvironmentById) | **Get** /environments/{environmentID} | Read One Environment
[**ReplaceBillOfMaterialsByEnvironmentId**](EnvironmentApi.md#ReplaceBillOfMaterialsByEnvironmentId) | **Put** /environments/{environmentID}/billOfMaterials | _TO_BE_DEFINED_
[**ReplaceEnvironmentById**](EnvironmentApi.md#ReplaceEnvironmentById) | **Put** /environments/{environmentID} | _TO_BE_DEFINED_



## CreateEnvironment

> Environment CreateEnvironment(ctx).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).EnvironmentCreateRequest(environmentCreateRequest).Execute()

Create Environment



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
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)
	environmentCreateRequest := *openapiclient.NewEnvironmentCreateRequest(*openapiclient.NewResourceRelationshipPingOne("TODO"), "Name_example", openapiclient.EnvironmentType("PRODUCTION"), openapiclient.EnvironmentRegion{EnvironmentRegionCode: openapiclient.Environment_Region_Code("AP")}) // EnvironmentCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentApi.CreateEnvironment(context.Background()).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).EnvironmentCreateRequest(environmentCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.CreateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironment`: Environment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.CreateEnvironment`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:create:environment`

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnvironmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 
 **environmentCreateRequest** | [**EnvironmentCreateRequest**](EnvironmentCreateRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentById

> DeleteEnvironmentById(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Delete Environment



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
	r, err := apiClient.EnvironmentApi.DeleteEnvironmentById(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.DeleteEnvironmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:delete:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentByIdRequest struct via the builder pattern


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


## GetBillOfMaterialsByEnvironmentId

> EnvironmentBillOfMaterials GetBillOfMaterialsByEnvironmentId(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	resp, r, err := apiClient.EnvironmentApi.GetBillOfMaterialsByEnvironmentId(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.GetBillOfMaterialsByEnvironmentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillOfMaterialsByEnvironmentId`: EnvironmentBillOfMaterials
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.GetBillOfMaterialsByEnvironmentId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:read:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillOfMaterialsByEnvironmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**EnvironmentBillOfMaterials**](EnvironmentBillOfMaterials.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentById

> Environment GetEnvironmentById(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Read One Environment



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
	resp, r, err := apiClient.EnvironmentApi.GetEnvironmentById(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.GetEnvironmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentById`: Environment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.GetEnvironmentById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:read:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**Environment**](Environment.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBillOfMaterialsByEnvironmentId

> EnvironmentBillOfMaterials ReplaceBillOfMaterialsByEnvironmentId(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).EnvironmentBillOfMaterialsReplaceRequest(environmentBillOfMaterialsReplaceRequest).Execute()

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
	environmentBillOfMaterialsReplaceRequest := *openapiclient.NewEnvironmentBillOfMaterialsReplaceRequest([]openapiclient.EnvironmentBillOfMaterialsProduct{*openapiclient.NewEnvironmentBillOfMaterialsProduct(openapiclient.Environment_Bill_of_Materials_Product_type("PING_ACCESS"))}) // EnvironmentBillOfMaterialsReplaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentApi.ReplaceBillOfMaterialsByEnvironmentId(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).EnvironmentBillOfMaterialsReplaceRequest(environmentBillOfMaterialsReplaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ReplaceBillOfMaterialsByEnvironmentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceBillOfMaterialsByEnvironmentId`: EnvironmentBillOfMaterials
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ReplaceBillOfMaterialsByEnvironmentId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:update:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBillOfMaterialsByEnvironmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 
 **environmentBillOfMaterialsReplaceRequest** | [**EnvironmentBillOfMaterialsReplaceRequest**](EnvironmentBillOfMaterialsReplaceRequest.md) |  | 

### Return type

[**EnvironmentBillOfMaterials**](EnvironmentBillOfMaterials.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEnvironmentById

> Environment ReplaceEnvironmentById(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).EnvironmentReplaceRequest(environmentReplaceRequest).Execute()

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
	environmentReplaceRequest := *openapiclient.NewEnvironmentReplaceRequest(*openapiclient.NewResourceRelationshipPingOne("TODO"), "Name_example", openapiclient.EnvironmentType("PRODUCTION")) // EnvironmentReplaceRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentApi.ReplaceEnvironmentById(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).EnvironmentReplaceRequest(environmentReplaceRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentApi.ReplaceEnvironmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceEnvironmentById`: Environment
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentApi.ReplaceEnvironmentById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:update:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEnvironmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 
 **environmentReplaceRequest** | [**EnvironmentReplaceRequest**](EnvironmentReplaceRequest.md) |  | 

### Return type

[**Environment**](Environment.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

