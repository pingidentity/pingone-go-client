# \DaVinciFlowVersionsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteVersionByIdUsingFlowId**](DaVinciFlowVersionsApi.md#DeleteVersionByIdUsingFlowId) | **Delete** /environments/{environmentID}/flows/{flowID}/versions/{versionID} | 
[**GetDetailsByFlowIdAndVersionId**](DaVinciFlowVersionsApi.md#GetDetailsByFlowIdAndVersionId) | **Get** /environments/{environmentID}/flows/{flowID}/versions/{versionID}/details | 
[**GetVersionByIdUsingFlowId**](DaVinciFlowVersionsApi.md#GetVersionByIdUsingFlowId) | **Get** /environments/{environmentID}/flows/{flowID}/versions/{versionID} | 
[**GetVersionsByFlowId**](DaVinciFlowVersionsApi.md#GetVersionsByFlowId) | **Get** /environments/{environmentID}/flows/{flowID}/versions | 
[**ReplaceAliasByFlowIdAndVersionId**](DaVinciFlowVersionsApi.md#ReplaceAliasByFlowIdAndVersionId) | **Put** /environments/{environmentID}/flows/{flowID}/versions/{versionID}/alias | 



## DeleteVersionByIdUsingFlowId

> DeleteVersionByIdUsingFlowId(ctx, environmentID, flowID, versionID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	flowID := "flowID_example" // string | 
	versionID := "versionID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciFlowVersionsApi.DeleteVersionByIdUsingFlowId(context.Background(), environmentID, flowID, versionID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowVersionsApi.DeleteVersionByIdUsingFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:delete:flowVersions`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 
**versionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVersionByIdUsingFlowIdRequest struct via the builder pattern


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


## GetDetailsByFlowIdAndVersionId

> DaVinciFlowVersionDetailResponse GetDetailsByFlowIdAndVersionId(ctx, environmentID, flowID, versionID).Expand(expand).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	flowID := "flowID_example" // string | 
	versionID := "versionID_example" // string | 
	expand := "expand_example" // string |  (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowVersionsApi.GetDetailsByFlowIdAndVersionId(context.Background(), environmentID, flowID, versionID).Expand(expand).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowVersionsApi.GetDetailsByFlowIdAndVersionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailsByFlowIdAndVersionId`: DaVinciFlowVersionDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowVersionsApi.GetDetailsByFlowIdAndVersionId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:flowVersions`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 
**versionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsByFlowIdAndVersionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowVersionDetailResponse**](DaVinciFlowVersionDetailResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionByIdUsingFlowId

> DaVinciFlowVersionResponse GetVersionByIdUsingFlowId(ctx, environmentID, flowID, versionID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	flowID := "flowID_example" // string | 
	versionID := "versionID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowVersionsApi.GetVersionByIdUsingFlowId(context.Background(), environmentID, flowID, versionID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowVersionsApi.GetVersionByIdUsingFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionByIdUsingFlowId`: DaVinciFlowVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowVersionsApi.GetVersionByIdUsingFlowId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:flowVersions`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 
**versionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionByIdUsingFlowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowVersionResponse**](DaVinciFlowVersionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionsByFlowId

> DaVinciFlowVersionCollectionResponse GetVersionsByFlowId(ctx, environmentID, flowID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	flowID := "flowID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowVersionsApi.GetVersionsByFlowId(context.Background(), environmentID, flowID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowVersionsApi.GetVersionsByFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionsByFlowId`: DaVinciFlowVersionCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowVersionsApi.GetVersionsByFlowId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:flowVersions`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionsByFlowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowVersionCollectionResponse**](DaVinciFlowVersionCollectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceAliasByFlowIdAndVersionId

> DaVinciFlowVersionAliasResponse ReplaceAliasByFlowIdAndVersionId(ctx, environmentID, flowID, versionID).DaVinciFlowVersionAliasRequest(daVinciFlowVersionAliasRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	flowID := "flowID_example" // string | 
	versionID := "versionID_example" // string | 
	daVinciFlowVersionAliasRequest := *openapiclient.NewDaVinciFlowVersionAliasRequest("Alias_example") // DaVinciFlowVersionAliasRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowVersionsApi.ReplaceAliasByFlowIdAndVersionId(context.Background(), environmentID, flowID, versionID).DaVinciFlowVersionAliasRequest(daVinciFlowVersionAliasRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowVersionsApi.ReplaceAliasByFlowIdAndVersionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceAliasByFlowIdAndVersionId`: DaVinciFlowVersionAliasResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowVersionsApi.ReplaceAliasByFlowIdAndVersionId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:flowVersions`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 
**versionID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceAliasByFlowIdAndVersionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **daVinciFlowVersionAliasRequest** | [**DaVinciFlowVersionAliasRequest**](DaVinciFlowVersionAliasRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowVersionAliasResponse**](DaVinciFlowVersionAliasResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

