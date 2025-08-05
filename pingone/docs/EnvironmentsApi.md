# \EnvironmentsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnvironment**](EnvironmentsApi.md#CreateEnvironment) | **Post** /environments | 
[**DeleteEnvironmentById**](EnvironmentsApi.md#DeleteEnvironmentById) | **Delete** /environments/{environmentID} | 
[**GetBillOfMaterialsByEnvironmentId**](EnvironmentsApi.md#GetBillOfMaterialsByEnvironmentId) | **Get** /environments/{environmentID}/billOfMaterials | 
[**GetEnvironmentById**](EnvironmentsApi.md#GetEnvironmentById) | **Get** /environments/{environmentID} | 
[**GetEnvironments**](EnvironmentsApi.md#GetEnvironments) | **Get** /environments | 
[**ReplaceBillOfMaterialsByEnvironmentId**](EnvironmentsApi.md#ReplaceBillOfMaterialsByEnvironmentId) | **Put** /environments/{environmentID}/billOfMaterials | 
[**ReplaceEnvironmentById**](EnvironmentsApi.md#ReplaceEnvironmentById) | **Put** /environments/{environmentID} | 



## CreateEnvironment

> EnvironmentResponse CreateEnvironment(ctx).EnvironmentCreateRequest(environmentCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	environmentCreateRequest := *openapiclient.NewEnvironmentCreateRequest("Name_example", openapiclient.EnvironmentRegionCode("AP"), openapiclient.EnvironmentTypeValue("PRODUCTION"), *openapiclient.NewEnvironmentLicense("TODO")) // EnvironmentCreateRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsApi.CreateEnvironment(context.Background()).EnvironmentCreateRequest(environmentCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.CreateEnvironment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnvironment`: EnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.CreateEnvironment`: %v\n", resp)
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
 **environmentCreateRequest** | [**EnvironmentCreateRequest**](EnvironmentCreateRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEnvironmentById

> DeleteEnvironmentById(ctx, environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	r, err := apiClient.EnvironmentsApi.DeleteEnvironmentById(context.Background(), environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.DeleteEnvironmentById``: %v\n", err)
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
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnvironmentByIdRequest struct via the builder pattern


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


## GetBillOfMaterialsByEnvironmentId

> EnvironmentBillOfMaterialsResponse GetBillOfMaterialsByEnvironmentId(ctx, environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	resp, r, err := apiClient.EnvironmentsApi.GetBillOfMaterialsByEnvironmentId(context.Background(), environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetBillOfMaterialsByEnvironmentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBillOfMaterialsByEnvironmentId`: EnvironmentBillOfMaterialsResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetBillOfMaterialsByEnvironmentId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:read:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBillOfMaterialsByEnvironmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**EnvironmentBillOfMaterialsResponse**](EnvironmentBillOfMaterialsResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironmentById

> EnvironmentResponse GetEnvironmentById(ctx, environmentID).Expand(expand).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	expand := "expand_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsApi.GetEnvironmentById(context.Background(), environmentID).Expand(expand).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironmentById`: EnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironmentById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:read:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnvironments



### Paged Response (Recommended)

> PagedIterator[EnvironmentsCollectionResponse] GetEnvironments(ctx).Expand(expand).Filter(filter).Order(order).Limit(limit).Cursor(cursor).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	expand := "expand_example" // string | 
	filter := "filter_example" // string |  (optional)
	order := "order_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 1000)
	cursor := "cursor_example" // string |  (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	pagedIterator := apiClient.EnvironmentsApi.GetEnvironments(context.Background()).Expand(expand).Filter(filter).Order(order).Limit(limit).Cursor(cursor).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironments``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `GetEnvironments` page iteration: EnvironmentsCollectionResponse
		fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironments` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> EnvironmentsCollectionResponse GetEnvironments(ctx).Expand(expand).Filter(filter).Order(order).Limit(limit).Cursor(cursor).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).ExecuteInitialPage()

#### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/pingidentity/pingone-go-client/pingone"
)

func main() {
	expand := "expand_example" // string | 
	filter := "filter_example" // string |  (optional)
	order := "order_example" // string |  (optional)
	limit := int32(56) // int32 |  (optional) (default to 1000)
	cursor := "cursor_example" // string |  (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsApi.GetEnvironments(context.Background()).Expand(expand).Filter(filter).Order(order).Limit(limit).Cursor(cursor).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.GetEnvironments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnvironments`: EnvironmentsCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.GetEnvironments`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:read:environment`

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEnvironmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** |  | 
 **filter** | **string** |  | 
 **order** | **string** |  | 
 **limit** | **int32** |  | [default to 1000]
 **cursor** | **string** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

Page Iterator: PagedIterator[[**EnvironmentsCollectionResponse**](EnvironmentsCollectionResponse.md)]

PagedIterator[EnvironmentsCollectionResponse] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**EnvironmentsCollectionResponse**](EnvironmentsCollectionResponse.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**EnvironmentsCollectionResponse**](EnvironmentsCollectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceBillOfMaterialsByEnvironmentId

> EnvironmentBillOfMaterialsResponse ReplaceBillOfMaterialsByEnvironmentId(ctx, environmentID).EnvironmentBillOfMaterialsReplaceRequest(environmentBillOfMaterialsReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	environmentBillOfMaterialsReplaceRequest := *openapiclient.NewEnvironmentBillOfMaterialsReplaceRequest([]openapiclient.EnvironmentBillOfMaterialsProduct{*openapiclient.NewEnvironmentBillOfMaterialsProduct(openapiclient.Environment_Bill_of_Materials_Product_Type("IDENTITY_CLOUD"))}) // EnvironmentBillOfMaterialsReplaceRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsApi.ReplaceBillOfMaterialsByEnvironmentId(context.Background(), environmentID).EnvironmentBillOfMaterialsReplaceRequest(environmentBillOfMaterialsReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ReplaceBillOfMaterialsByEnvironmentId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceBillOfMaterialsByEnvironmentId`: EnvironmentBillOfMaterialsResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ReplaceBillOfMaterialsByEnvironmentId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:update:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceBillOfMaterialsByEnvironmentIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentBillOfMaterialsReplaceRequest** | [**EnvironmentBillOfMaterialsReplaceRequest**](EnvironmentBillOfMaterialsReplaceRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**EnvironmentBillOfMaterialsResponse**](EnvironmentBillOfMaterialsResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEnvironmentById

> EnvironmentResponse ReplaceEnvironmentById(ctx, environmentID).EnvironmentReplaceRequest(environmentReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	environmentReplaceRequest := *openapiclient.NewEnvironmentReplaceRequest("Name_example", openapiclient.EnvironmentRegionCode("AP"), openapiclient.EnvironmentTypeValue("PRODUCTION")) // EnvironmentReplaceRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnvironmentsApi.ReplaceEnvironmentById(context.Background(), environmentID).EnvironmentReplaceRequest(environmentReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnvironmentsApi.ReplaceEnvironmentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceEnvironmentById`: EnvironmentResponse
	fmt.Fprintf(os.Stdout, "Response from `EnvironmentsApi.ReplaceEnvironmentById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `orgmgt:update:environment`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEnvironmentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **environmentReplaceRequest** | [**EnvironmentReplaceRequest**](EnvironmentReplaceRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**EnvironmentResponse**](EnvironmentResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/hal+json, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

