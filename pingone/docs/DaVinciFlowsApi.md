# \DaVinciFlowsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFlow**](DaVinciFlowsApi.md#CreateFlow) | **Post** /environments/{environmentID}/flows | 
[**DeleteFlowById**](DaVinciFlowsApi.md#DeleteFlowById) | **Delete** /environments/{environmentID}/flows/{flowID} | 
[**DeleteVersionByIdUsingFlowId**](DaVinciFlowsApi.md#DeleteVersionByIdUsingFlowId) | **Delete** /environments/{environmentID}/flows/{flowID}/versions/{versionID} | 
[**GetDetailsByFlowIdAndVersionId**](DaVinciFlowsApi.md#GetDetailsByFlowIdAndVersionId) | **Get** /environments/{environmentID}/flows/{flowID}/versions/{versionID}/details | 
[**GetFlowById**](DaVinciFlowsApi.md#GetFlowById) | **Get** /environments/{environmentID}/flows/{flowID} | 
[**GetFlows**](DaVinciFlowsApi.md#GetFlows) | **Get** /environments/{environmentID}/flows | 
[**GetVersionByIdUsingFlowId**](DaVinciFlowsApi.md#GetVersionByIdUsingFlowId) | **Get** /environments/{environmentID}/flows/{flowID}/versions/{versionID} | 
[**GetVersionsByFlowId**](DaVinciFlowsApi.md#GetVersionsByFlowId) | **Get** /environments/{environmentID}/flows/{flowID}/versions | 
[**ReplaceAliasByFlowIdAndVersionId**](DaVinciFlowsApi.md#ReplaceAliasByFlowIdAndVersionId) | **Put** /environments/{environmentID}/flows/{flowID}/versions/{versionID}/alias | 
[**ReplaceFlowById**](DaVinciFlowsApi.md#ReplaceFlowById) | **Put** /environments/{environmentID}/flows/{flowID} | 
[**UpdateEnabledByFlowId**](DaVinciFlowsApi.md#UpdateEnabledByFlowId) | **Put** /environments/{environmentID}/flows/{flowID}/enabled | 



## CreateFlow

> DaVinciFlowResponse CreateFlow(ctx, environmentID).DaVinciFlowCreateRequest(daVinciFlowCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	daVinciFlowCreateRequest := *openapiclient.NewDaVinciFlowCreateRequest("Name_example") // DaVinciFlowCreateRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowsApi.CreateFlow(context.Background(), environmentID).DaVinciFlowCreateRequest(daVinciFlowCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.CreateFlow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFlow`: DaVinciFlowResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.CreateFlow`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:create:dvFlows`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFlowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daVinciFlowCreateRequest** | [**DaVinciFlowCreateRequest**](DaVinciFlowCreateRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowResponse**](DaVinciFlowResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFlowById

> DeleteFlowById(ctx, environmentID, flowID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	r, err := apiClient.DaVinciFlowsApi.DeleteFlowById(context.Background(), environmentID, flowID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.DeleteFlowById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:delete:dvFlows`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFlowByIdRequest struct via the builder pattern


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
	r, err := apiClient.DaVinciFlowsApi.DeleteVersionByIdUsingFlowId(context.Background(), environmentID, flowID, versionID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.DeleteVersionByIdUsingFlowId``: %v\n", err)
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
	resp, r, err := apiClient.DaVinciFlowsApi.GetDetailsByFlowIdAndVersionId(context.Background(), environmentID, flowID, versionID).Expand(expand).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.GetDetailsByFlowIdAndVersionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailsByFlowIdAndVersionId`: DaVinciFlowVersionDetailResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.GetDetailsByFlowIdAndVersionId`: %v\n", resp)
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


## GetFlowById

> DaVinciFlowResponse GetFlowById(ctx, environmentID, flowID).Attributes(attributes).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	attributes := "attributes_example" // string |  (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowsApi.GetFlowById(context.Background(), environmentID, flowID).Attributes(attributes).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.GetFlowById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlowById`: DaVinciFlowResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.GetFlowById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:dvFlows`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **attributes** | **string** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowResponse**](DaVinciFlowResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFlows

> DaVinciFlowCollection GetFlows(ctx, environmentID).Attributes(attributes).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	attributes := "attributes_example" // string |  (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowsApi.GetFlows(context.Background(), environmentID).Attributes(attributes).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.GetFlows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFlows`: DaVinciFlowCollection
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.GetFlows`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:dvFlows`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFlowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **attributes** | **string** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowCollection**](DaVinciFlowCollection.md)

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
	resp, r, err := apiClient.DaVinciFlowsApi.GetVersionByIdUsingFlowId(context.Background(), environmentID, flowID, versionID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.GetVersionByIdUsingFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionByIdUsingFlowId`: DaVinciFlowVersionResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.GetVersionByIdUsingFlowId`: %v\n", resp)
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
	resp, r, err := apiClient.DaVinciFlowsApi.GetVersionsByFlowId(context.Background(), environmentID, flowID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.GetVersionsByFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionsByFlowId`: DaVinciFlowVersionCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.GetVersionsByFlowId`: %v\n", resp)
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
	resp, r, err := apiClient.DaVinciFlowsApi.ReplaceAliasByFlowIdAndVersionId(context.Background(), environmentID, flowID, versionID).DaVinciFlowVersionAliasRequest(daVinciFlowVersionAliasRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.ReplaceAliasByFlowIdAndVersionId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceAliasByFlowIdAndVersionId`: DaVinciFlowVersionAliasResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.ReplaceAliasByFlowIdAndVersionId`: %v\n", resp)
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


## ReplaceFlowById

> DaVinciFlowResponse ReplaceFlowById(ctx, environmentID, flowID).DaVinciFlowReplaceRequest(daVinciFlowReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	daVinciFlowReplaceRequest := *openapiclient.NewDaVinciFlowReplaceRequest("Name_example") // DaVinciFlowReplaceRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowsApi.ReplaceFlowById(context.Background(), environmentID, flowID).DaVinciFlowReplaceRequest(daVinciFlowReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.ReplaceFlowById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceFlowById`: DaVinciFlowResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.ReplaceFlowById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:dvFlows`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceFlowByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciFlowReplaceRequest** | [**DaVinciFlowReplaceRequest**](DaVinciFlowReplaceRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowResponse**](DaVinciFlowResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnabledByFlowId

> DaVinciFlowEnabledResponse UpdateEnabledByFlowId(ctx, environmentID, flowID).DaVinciFlowEnableRequest(daVinciFlowEnableRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	daVinciFlowEnableRequest := *openapiclient.NewDaVinciFlowEnableRequest(false) // DaVinciFlowEnableRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciFlowsApi.UpdateEnabledByFlowId(context.Background(), environmentID, flowID).DaVinciFlowEnableRequest(daVinciFlowEnableRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciFlowsApi.UpdateEnabledByFlowId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnabledByFlowId`: DaVinciFlowEnabledResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciFlowsApi.UpdateEnabledByFlowId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:dvFlows`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**flowID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnabledByFlowIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciFlowEnableRequest** | [**DaVinciFlowEnableRequest**](DaVinciFlowEnableRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciFlowEnabledResponse**](DaVinciFlowEnabledResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

