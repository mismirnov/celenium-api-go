# \GasAPI

All URIs are relative to *https://api.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GasEstimateForPfb**](GasAPI.md#GasEstimateForPfb) | **Get** /gas/estimate_for_pfb | Get estimated gas for pay for blob
[**GasPrice**](GasAPI.md#GasPrice) | **Get** /gas/price | Get estimated gas price



## GasEstimateForPfb

> int32 GasEstimateForPfb(ctx).Sizes(sizes).Execute()

Get estimated gas for pay for blob



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	celenium "github.com/mismirnov/celenium-api-go"
)

func main() {
	sizes := "sizes_example" // string | Comma-separated array of blob sizes

	configuration := celenium.NewConfiguration()
	apiClient := celenium.NewAPIClient(configuration)
	resp, r, err := apiClient.GasAPI.GasEstimateForPfb(context.Background()).Sizes(sizes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GasAPI.GasEstimateForPfb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GasEstimateForPfb`: int32
	fmt.Fprintf(os.Stdout, "Response from `GasAPI.GasEstimateForPfb`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGasEstimateForPfbRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sizes** | **string** | Comma-separated array of blob sizes | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GasPrice

> ResponsesGasPrice GasPrice(ctx).Execute()

Get estimated gas price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	celenium "github.com/mismirnov/celenium-api-go"
)

func main() {

	configuration := celenium.NewConfiguration()
	apiClient := celenium.NewAPIClient(configuration)
	resp, r, err := apiClient.GasAPI.GasPrice(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GasAPI.GasPrice``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GasPrice`: ResponsesGasPrice
	fmt.Fprintf(os.Stdout, "Response from `GasAPI.GasPrice`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGasPriceRequest struct via the builder pattern


### Return type

[**ResponsesGasPrice**](ResponsesGasPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

