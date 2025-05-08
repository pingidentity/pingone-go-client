# \DaVinciVariableApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateVariable**](DaVinciVariableApi.md#CreateVariable) | **Post** /environments/{environmentID}/variables | Create DaVinci Variable
[**DeleteVariableById**](DaVinciVariableApi.md#DeleteVariableById) | **Delete** /environments/{environmentID}/variables/{variableID} | _TO_BE_DEFINED_
[**GetVariableById**](DaVinciVariableApi.md#GetVariableById) | **Get** /environments/{environmentID}/variables/{variableID} | _TO_BE_DEFINED_
[**ReplaceVariableById**](DaVinciVariableApi.md#ReplaceVariableById) | **Put** /environments/{environmentID}/variables/{variableID} | _TO_BE_DEFINED_



## CreateVariable

> DaVinciVariable CreateVariable(ctx, environmentID).DaVinciVariableCreateRequest(daVinciVariableCreateRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Create DaVinci Variable



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
	daVinciVariableCreateRequest := *openapiclient.NewDaVinciVariableCreateRequest(openapiclient.DaVinci_Variable_Create_Request_Context("company"), openapiclient.DaVinci_Variable_Create_Request_Data_Type("boolean"), false, "Name_example") // DaVinciVariableCreateRequest | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciVariableApi.CreateVariable(context.Background(), environmentID).DaVinciVariableCreateRequest(daVinciVariableCreateRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciVariableApi.CreateVariable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVariable`: DaVinciVariable
	fmt.Fprintf(os.Stdout, "Response from `DaVinciVariableApi.CreateVariable`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:create:constructs`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **daVinciVariableCreateRequest** | [**DaVinciVariableCreateRequest**](DaVinciVariableCreateRequest.md) |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciVariable**](DaVinciVariable.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVariableById

> DeleteVariableById(ctx, environmentID, variableID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	variableID := "variableID_example" // uuid.UUID | _TO_BE_DEFINED_
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DaVinciVariableApi.DeleteVariableById(context.Background(), environmentID, variableID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciVariableApi.DeleteVariableById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:delete:constructs`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**variableID** | **uuid.UUID** | _TO_BE_DEFINED_ | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVariableByIdRequest struct via the builder pattern


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


## GetVariableById

> DaVinciVariable GetVariableById(ctx, environmentID, variableID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	variableID := "variableID_example" // uuid.UUID | _TO_BE_DEFINED_
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciVariableApi.GetVariableById(context.Background(), environmentID, variableID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciVariableApi.GetVariableById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVariableById`: DaVinciVariable
	fmt.Fprintf(os.Stdout, "Response from `DaVinciVariableApi.GetVariableById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:read:constructs`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**variableID** | **uuid.UUID** | _TO_BE_DEFINED_ | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVariableByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciVariable**](DaVinciVariable.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceVariableById

> DaVinciVariable ReplaceVariableById(ctx, environmentID, variableID).DaVinciVariableReplaceRequest(daVinciVariableReplaceRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

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
	variableID := "variableID_example" // uuid.UUID | _TO_BE_DEFINED_
	daVinciVariableReplaceRequest := *openapiclient.NewDaVinciVariableReplaceRequest(openapiclient.DaVinci_Variable_Replace_Request_Context("company"), openapiclient.DaVinci_Variable_Replace_Request_Data_Type("boolean"), false, "Name_example") // DaVinciVariableReplaceRequest | 
	xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. (optional)
	xPingExternalSessionID := "xPingExternalSessionID_example" // string | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DaVinciVariableApi.ReplaceVariableById(context.Background(), environmentID, variableID).DaVinciVariableReplaceRequest(daVinciVariableReplaceRequest).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DaVinciVariableApi.ReplaceVariableById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceVariableById`: DaVinciVariable
	fmt.Fprintf(os.Stdout, "Response from `DaVinciVariableApi.ReplaceVariableById`: %v\n", resp)
}
```

### Required Permission(s)

The following admin role permissions are required to call this endpoint:

- `davinci:update:constructs`

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **uuid.UUID** | An environment&#39;s unique identifier in UUID format. | 
**variableID** | **uuid.UUID** | _TO_BE_DEFINED_ | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceVariableByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **daVinciVariableReplaceRequest** | [**DaVinciVariableReplaceRequest**](DaVinciVariableReplaceRequest.md) |  | 
 **xPingExternalTransactionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single PingOne API request. It identifies one transaction that encompasses multiple API requests to PingOne. | 
 **xPingExternalSessionID** | **string** | In order to help track transactions, the PingOne platform supports this custom HTTP header that represents a scope larger than a single transaction. It identifies multiple transactions in the context of a session. For example, an end user completed an authentication request and several transactions one hour ago and now needs to re-authenticate. The session should be the same. | 

### Return type

[**DaVinciVariable**](DaVinciVariable.md)

### Authorization

[oauth2](../README.md#oauth2), [oauth2](../README.md#oauth2), [bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

