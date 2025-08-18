# \ConfigurationManagementApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSnapshot**](ConfigurationManagementApi.md#CreateSnapshot) | **Post** /environments/{environmentID}/snapshots | 
[**GetSnapshotById**](ConfigurationManagementApi.md#GetSnapshotById) | **Get** /environments/{environmentID}/snapshots/{snapshotID} | 
[**GetVersionByIdUsingSnapshotId**](ConfigurationManagementApi.md#GetVersionByIdUsingSnapshotId) | **Get** /environments/{environmentID}/snapshots/{snapshotID}/versions/{versionID} | 
[**GetVersionsBySnapshotId**](ConfigurationManagementApi.md#GetVersionsBySnapshotId) | **Get** /environments/{environmentID}/snapshots/{snapshotID}/versions | 



## CreateSnapshot

> SnapshotView CreateSnapshot(ctx, environmentID).SnapshotRequest(snapshotRequest).Expand(expand).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	snapshotRequest := *openapiclient.NewSnapshotRequest("BaseResourceURL_example") // SnapshotRequest | 
	expand := "expand_example" // string |  (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationManagementApi.CreateSnapshot(context.Background(), environmentID).SnapshotRequest(snapshotRequest).Expand(expand).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationManagementApi.CreateSnapshot``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSnapshot`: SnapshotView
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationManagementApi.CreateSnapshot`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `promotion:create:snapshot`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **snapshotRequest** | [**SnapshotRequest**](SnapshotRequest.md) |  | 
 **expand** | **string** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**SnapshotView**](SnapshotView.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSnapshotById

> SnapshotView GetSnapshotById(ctx, environmentID, snapshotID).Expand(expand).Filter(filter).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	snapshotID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	expand := "expand_example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationManagementApi.GetSnapshotById(context.Background(), environmentID, snapshotID).Expand(expand).Filter(filter).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationManagementApi.GetSnapshotById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSnapshotById`: SnapshotView
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationManagementApi.GetSnapshotById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `promotion:read:snapshot`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**snapshotID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSnapshotByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 
 **filter** | **string** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**SnapshotView**](SnapshotView.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionByIdUsingSnapshotId

> SnapshotView GetVersionByIdUsingSnapshotId(ctx, environmentID, snapshotID, versionID).Expand(expand).Filter(filter).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	snapshotID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	versionID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	expand := "expand_example" // string |  (optional)
	filter := "filter_example" // string |  (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationManagementApi.GetVersionByIdUsingSnapshotId(context.Background(), environmentID, snapshotID, versionID).Expand(expand).Filter(filter).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationManagementApi.GetVersionByIdUsingSnapshotId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionByIdUsingSnapshotId`: SnapshotView
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationManagementApi.GetVersionByIdUsingSnapshotId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `promotion:read:snapshot`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**snapshotID** | **uuid.UUID** |  | 
**versionID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionByIdUsingSnapshotIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** |  | 
 **filter** | **string** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**SnapshotView**](SnapshotView.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionsBySnapshotId

> SnapshotVersionCollectionResponse GetVersionsBySnapshotId(ctx, environmentID, snapshotID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	snapshotID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // uuid.UUID | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ConfigurationManagementApi.GetVersionsBySnapshotId(context.Background(), environmentID, snapshotID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ConfigurationManagementApi.GetVersionsBySnapshotId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionsBySnapshotId`: SnapshotVersionCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `ConfigurationManagementApi.GetVersionsBySnapshotId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `promotion:read:snapshot`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**snapshotID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionsBySnapshotIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**SnapshotVersionCollectionResponse**](SnapshotVersionCollectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/hal+json, application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

