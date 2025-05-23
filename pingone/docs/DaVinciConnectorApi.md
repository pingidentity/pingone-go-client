# \DaVinciConnectorApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConnectorInstance**](DaVinciConnectorApi.md#CreateConnectorInstance) | **Post** /environments/{environmentID}/connectorInstances | _TO_BE_DEFINED_
[**CreateConnectorInstanceById**](DaVinciConnectorApi.md#CreateConnectorInstanceById) | **Post** /environments/{environmentID}/connectorInstances/{connectorInstanceID} | _TO_BE_DEFINED_
[**DeleteConnectorInstanceById**](DaVinciConnectorApi.md#DeleteConnectorInstanceById) | **Delete** /environments/{environmentID}/connectorInstances/{connectorInstanceID} | _TO_BE_DEFINED_
[**GetConnectorById**](DaVinciConnectorApi.md#GetConnectorById) | **Get** /environments/{environmentID}/connectors/{connectorID} | _TO_BE_DEFINED_
[**GetConnectorInstanceById**](DaVinciConnectorApi.md#GetConnectorInstanceById) | **Get** /environments/{environmentID}/connectorInstances/{connectorInstanceID} | _TO_BE_DEFINED_
[**GetConnectorInstances**](DaVinciConnectorApi.md#GetConnectorInstances) | **Get** /environments/{environmentID}/connectorInstances | _TO_BE_DEFINED_
[**GetConnectors**](DaVinciConnectorApi.md#GetConnectors) | **Get** /environments/{environmentID}/connectors | _TO_BE_DEFINED_
[**GetDetailsByConnectorId**](DaVinciConnectorApi.md#GetDetailsByConnectorId) | **Get** /environments/{environmentID}/connectors/{connectorID}/details | _TO_BE_DEFINED_
[**ReplaceConnectorInstanceById**](DaVinciConnectorApi.md#ReplaceConnectorInstanceById) | **Put** /environments/{environmentID}/connectorInstances/{connectorInstanceID} | _TO_BE_DEFINED_



## CreateConnectorInstance

> DaVinciConnectorInstance CreateConnectorInstance(ctx, environmentID).DaVinciConnectorInstanceCreateRequest(daVinciConnectorInstanceCreateRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	daVinciConnectorInstanceCreateRequest := *openapiclient.NewDaVinciConnectorInstanceCreateRequest(*openapiclient.NewDaVinciConnectorInstanceCreateRequestConnector(), "Name_example") // DaVinciConnectorInstanceCreateRequest | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorApi.CreateConnectorInstance(context.Background(), environmentID).DaVinciConnectorInstanceCreateRequest(daVinciConnectorInstanceCreateRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.CreateConnectorInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorInstance`: DaVinciConnectorInstance
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.CreateConnectorInstance`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:create:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daVinciConnectorInstanceCreateRequest** | [**DaVinciConnectorInstanceCreateRequest**](DaVinciConnectorInstanceCreateRequest.md) |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciConnectorInstance**](DaVinciConnectorInstance.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateConnectorInstanceById

> DaVinciConnectorInstance CreateConnectorInstanceById(ctx, environmentID, connectorInstanceID).RequestBody(requestBody).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	connectorInstanceID := "connectorInstanceID_example" // string | 
	requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorApi.CreateConnectorInstanceById(context.Background(), environmentID, connectorInstanceID).RequestBody(requestBody).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.CreateConnectorInstanceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateConnectorInstanceById`: DaVinciConnectorInstance
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.CreateConnectorInstanceById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:create:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**connectorInstanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConnectorInstanceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestBody** | **map[string]interface{}** |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciConnectorInstance**](DaVinciConnectorInstance.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.pingidentity.connectorInstance.clone+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteConnectorInstanceById

> DeleteConnectorInstanceById(ctx, environmentID, connectorInstanceID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	connectorInstanceID := "connectorInstanceID_example" // string | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciConnectorApi.DeleteConnectorInstanceById(context.Background(), environmentID, connectorInstanceID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.DeleteConnectorInstanceById``: %v\n", err)
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
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**connectorInstanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectorInstanceByIdRequest struct via the builder pattern


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


## GetConnectorById

> DaVinciConnectorMinimalResponse GetConnectorById(ctx, environmentID, connectorID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	connectorID := "connectorID_example" // string | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorApi.GetConnectorById(context.Background(), environmentID, connectorID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.GetConnectorById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorById`: DaVinciConnectorMinimalResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.GetConnectorById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connectors`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**connectorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciConnectorMinimalResponse**](DaVinciConnectorMinimalResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorInstanceById

> DaVinciConnectorInstance GetConnectorInstanceById(ctx, environmentID, connectorInstanceID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	connectorInstanceID := "connectorInstanceID_example" // string | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorApi.GetConnectorInstanceById(context.Background(), environmentID, connectorInstanceID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.GetConnectorInstanceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorInstanceById`: DaVinciConnectorInstance
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.GetConnectorInstanceById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**connectorInstanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorInstanceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciConnectorInstance**](DaVinciConnectorInstance.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectorInstances

_TO_BE_DEFINED_



### Paged Response (Recommended)

> PagedIterator[DaVinciConnectorInstanceCollection] GetConnectorInstances(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	pagedIterator := apiClient.DaVinciConnectorApi.GetConnectorInstances(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.GetConnectorInstances``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `GetConnectorInstances` page iteration: DaVinciConnectorInstanceCollection
		fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.GetConnectorInstances` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> DaVinciConnectorInstanceCollection GetConnectorInstances(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).ExecuteInitialPage()

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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorApi.GetConnectorInstances(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.GetConnectorInstances``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectorInstances`: DaVinciConnectorInstanceCollection
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.GetConnectorInstances`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorInstancesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

Page Iterator: PagedIterator[[**DaVinciConnectorInstanceCollection**](DaVinciConnectorInstanceCollection.md)]

PagedIterator[DaVinciConnectorInstanceCollection] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**DaVinciConnectorInstanceCollection**](DaVinciConnectorInstanceCollection.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**DaVinciConnectorInstanceCollection**](DaVinciConnectorInstanceCollection.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConnectors

_TO_BE_DEFINED_



### Paged Response (Recommended)

> PagedIterator[DaVinciConnectorCollectionMinimalResponse] GetConnectors(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	pagedIterator := apiClient.DaVinciConnectorApi.GetConnectors(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

	for pageCursor, err := range pagedIterator {
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.GetConnectors``: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", pageCursor.HTTPResponse)
		}

		// response from `GetConnectors` page iteration: DaVinciConnectorCollectionMinimalResponse
		fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.GetConnectors` page iteration: %v\n", pageCursor.Data)
	}
}
```

### Initial Page Response

> DaVinciConnectorCollectionMinimalResponse GetConnectors(ctx, environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).ExecuteInitialPage()

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
	environmentID := "environmentID_example" // uuid.UUID | An environment's unique identifier in UUID format.
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorApi.GetConnectors(context.Background(), environmentID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).ExecuteInitialPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.GetConnectors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConnectors`: DaVinciConnectorCollectionMinimalResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.GetConnectors`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connectors`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

Page Iterator: PagedIterator[[**DaVinciConnectorCollectionMinimalResponse**](DaVinciConnectorCollectionMinimalResponse.md)]

PagedIterator[DaVinciConnectorCollectionMinimalResponse] is a struct alias for iter.Seq2[[PagedCursor](PagedCursor.md)[[**DaVinciConnectorCollectionMinimalResponse**](DaVinciConnectorCollectionMinimalResponse.md)], error] using the standard `iter` package in go `1.23`.

Page Data: [**DaVinciConnectorCollectionMinimalResponse**](DaVinciConnectorCollectionMinimalResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDetailsByConnectorId

> DaVinciConnectorDetailsResponse GetDetailsByConnectorId(ctx, environmentID, connectorID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	connectorID := "connectorID_example" // string | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorApi.GetDetailsByConnectorId(context.Background(), environmentID, connectorID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.GetDetailsByConnectorId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDetailsByConnectorId`: DaVinciConnectorDetailsResponse
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.GetDetailsByConnectorId`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:connectors`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**connectorID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetailsByConnectorIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciConnectorDetailsResponse**](DaVinciConnectorDetailsResponse.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceConnectorInstanceById

> DaVinciConnectorInstance ReplaceConnectorInstanceById(ctx, environmentID, connectorInstanceID).DaVinciConnectorInstanceReplaceRequest(daVinciConnectorInstanceReplaceRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	connectorInstanceID := "connectorInstanceID_example" // string | 
	daVinciConnectorInstanceReplaceRequest := *openapiclient.NewDaVinciConnectorInstanceReplaceRequest("Name_example") // DaVinciConnectorInstanceReplaceRequest | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciConnectorApi.ReplaceConnectorInstanceById(context.Background(), environmentID, connectorInstanceID).DaVinciConnectorInstanceReplaceRequest(daVinciConnectorInstanceReplaceRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciConnectorApi.ReplaceConnectorInstanceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceConnectorInstanceById`: DaVinciConnectorInstance
	fmt.Fprintf(os.Stdout, "Response from `DaVinciConnectorApi.ReplaceConnectorInstanceById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:connections`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**connectorInstanceID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceConnectorInstanceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciConnectorInstanceReplaceRequest** | [**DaVinciConnectorInstanceReplaceRequest**](DaVinciConnectorInstanceReplaceRequest.md) |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciConnectorInstance**](DaVinciConnectorInstance.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

