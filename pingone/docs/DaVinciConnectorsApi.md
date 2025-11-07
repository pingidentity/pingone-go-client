# \DaVinciConnectorsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorInstance**](DaVinciConnectorsApi.md#CreateConnectorInstance) | **Post** /environments/{environmentID}/connectorInstances | 
[**CreateConnectorInstanceById**](DaVinciConnectorsApi.md#CreateConnectorInstanceById) | **Post** /environments/{environmentID}/connectorInstances/{connectorInstanceID} | 
[**DeleteConnectorInstanceById**](DaVinciConnectorsApi.md#DeleteConnectorInstanceById) | **Delete** /environments/{environmentID}/connectorInstances/{connectorInstanceID} | 
[**GetConnectorById**](DaVinciConnectorsApi.md#GetConnectorById) | **Get** /environments/{environmentID}/connectors/{connectorID} | 
[**GetConnectorInstanceById**](DaVinciConnectorsApi.md#GetConnectorInstanceById) | **Get** /environments/{environmentID}/connectorInstances/{connectorInstanceID} | 
[**GetConnectorInstances**](DaVinciConnectorsApi.md#GetConnectorInstances) | **Get** /environments/{environmentID}/connectorInstances | 
[**GetConnectors**](DaVinciConnectorsApi.md#GetConnectors) | **Get** /environments/{environmentID}/connectors | 
[**GetDetailsByConnectorId**](DaVinciConnectorsApi.md#GetDetailsByConnectorId) | **Get** /environments/{environmentID}/connectors/{connectorID}/details | 
[**ReplaceConnectorInstanceById**](DaVinciConnectorsApi.md#ReplaceConnectorInstanceById) | **Put** /environments/{environmentID}/connectorInstances/{connectorInstanceID} | 



## CreateConnectorInstance

> DaVinciConnectorInstanceResponse CreateConnectorInstance(ctx, environmentID).DaVinciConnectorInstanceCreateRequest(daVinciConnectorInstanceCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	daVinciConnectorInstanceCreateRequest := *openapiclient.NewDaVinciConnectorInstanceCreateRequest("Name_example", *openapiclient.NewResourceRelationshipDaVinci("Id_example")) // DaVinciConnectorInstanceCreateRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorsApi.CreateConnectorInstance(context.Background(), environmentID).DaVinciConnectorInstanceCreateRequest(daVinciConnectorInstanceCreateRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.CreateConnectorInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorInstance`: DaVinciConnectorInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorsApi.CreateConnectorInstance`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:create:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daVinciConnectorInstanceCreateRequest** | [**DaVinciConnectorInstanceCreateRequest**](DaVinciConnectorInstanceCreateRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciConnectorInstanceResponse**](DaVinciConnectorInstanceResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectorInstanceById

> DaVinciConnectorInstanceResponse CreateConnectorInstanceById(ctx, environmentID, connectorInstanceID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	connectorInstanceID := "connectorInstanceID_example" // string | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorsApi.CreateConnectorInstanceById(context.Background(), environmentID, connectorInstanceID).RequestBody(requestBody).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.CreateConnectorInstanceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorInstanceById`: DaVinciConnectorInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorsApi.CreateConnectorInstanceById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:create:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**connectorInstanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorInstanceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciConnectorInstanceResponse**](DaVinciConnectorInstanceResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.pingidentity.connectorInstance.clone+json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectorInstanceById

> DeleteConnectorInstanceById(ctx, environmentID, connectorInstanceID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	connectorInstanceID := "connectorInstanceID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciConnectorsApi.DeleteConnectorInstanceById(context.Background(), environmentID, connectorInstanceID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.DeleteConnectorInstanceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:delete:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**connectorInstanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorInstanceByIdRequest struct via the builder pattern


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


## GetConnectorById

> DaVinciConnectorMinimalResponse GetConnectorById(ctx, environmentID, connectorID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	connectorID := "connectorID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorsApi.GetConnectorById(context.Background(), environmentID, connectorID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.GetConnectorById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorById`: DaVinciConnectorMinimalResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorsApi.GetConnectorById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connectors`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**connectorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciConnectorMinimalResponse**](DaVinciConnectorMinimalResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorInstanceById

> DaVinciConnectorInstanceResponse GetConnectorInstanceById(ctx, environmentID, connectorInstanceID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	connectorInstanceID := "connectorInstanceID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorsApi.GetConnectorInstanceById(context.Background(), environmentID, connectorInstanceID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.GetConnectorInstanceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorInstanceById`: DaVinciConnectorInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorsApi.GetConnectorInstanceById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**connectorInstanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorInstanceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciConnectorInstanceResponse**](DaVinciConnectorInstanceResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorInstances

> DaVinciConnectorInstanceCollectionResponse GetConnectorInstances(ctx, environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	resp, r, err := apiClient.DaVinciConnectorsApi.GetConnectorInstances(context.Background(), environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.GetConnectorInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorInstances`: DaVinciConnectorInstanceCollectionResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorsApi.GetConnectorInstances`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciConnectorInstanceCollectionResponse**](DaVinciConnectorInstanceCollectionResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectors

> DaVinciConnectorCollectionMinimalResponse GetConnectors(ctx, environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	resp, r, err := apiClient.DaVinciConnectorsApi.GetConnectors(context.Background(), environmentID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.GetConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectors`: DaVinciConnectorCollectionMinimalResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorsApi.GetConnectors`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connectors`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciConnectorCollectionMinimalResponse**](DaVinciConnectorCollectionMinimalResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetailsByConnectorId

> DaVinciConnectorDetailsResponse GetDetailsByConnectorId(ctx, environmentID, connectorID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	connectorID := "connectorID_example" // string | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorsApi.GetDetailsByConnectorId(context.Background(), environmentID, connectorID).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.GetDetailsByConnectorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailsByConnectorId`: DaVinciConnectorDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorsApi.GetDetailsByConnectorId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connectors`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**connectorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsByConnectorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciConnectorDetailsResponse**](DaVinciConnectorDetailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceConnectorInstanceById

> DaVinciConnectorInstanceResponse ReplaceConnectorInstanceById(ctx, environmentID, connectorInstanceID).DaVinciConnectorInstanceReplaceRequest(daVinciConnectorInstanceReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()



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
	connectorInstanceID := "connectorInstanceID_example" // string | 
	daVinciConnectorInstanceReplaceRequest := *openapiclient.NewDaVinciConnectorInstanceReplaceRequest("Name_example", *openapiclient.NewResourceRelationshipDaVinci("Id_example")) // DaVinciConnectorInstanceReplaceRequest | 
	xPingExternalSessionID := "xPingExternalSessionID_example" // string |  (optional)
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorsApi.ReplaceConnectorInstanceById(context.Background(), environmentID, connectorInstanceID).DaVinciConnectorInstanceReplaceRequest(daVinciConnectorInstanceReplaceRequest).XPingExternalSessionID(xPingExternalSessionID).XPingExternalTransactionID(xPingExternalTransactionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorsApi.ReplaceConnectorInstanceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceConnectorInstanceById`: DaVinciConnectorInstanceResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorsApi.ReplaceConnectorInstanceById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** |  | 
**connectorInstanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceConnectorInstanceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciConnectorInstanceReplaceRequest** | [**DaVinciConnectorInstanceReplaceRequest**](DaVinciConnectorInstanceReplaceRequest.md) |  | 
 **xPingExternalSessionID** | **string** |  | 
 **xPingExternalTransactionID** | **string** |  | 

### Return type

[**DaVinciConnectorInstanceResponse**](DaVinciConnectorInstanceResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

