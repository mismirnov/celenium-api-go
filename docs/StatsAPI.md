# \StatsAPI

All URIs are relative to *https://api-mainnet.celenium.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Stats24hChanges**](StatsAPI.md#Stats24hChanges) | **Get** /stats/changes_24h | Get changes for 24 hours
[**StatsMessagesCount24h**](StatsAPI.md#StatsMessagesCount24h) | **Get** /stats/messages_count_24h | Get messages distribution for the last 24 hours
[**StatsNamespaceUsage**](StatsAPI.md#StatsNamespaceUsage) | **Get** /stats/namespace/usage | Get namespaces with sorting by size.
[**StatsNsSeries**](StatsAPI.md#StatsNsSeries) | **Get** /stats/namespace/series/{id}/{name}/{timeframe} | Get histogram for namespace with precomputed stats
[**StatsPriceCurrent**](StatsAPI.md#StatsPriceCurrent) | **Get** /stats/price/current | Get current TIA price
[**StatsPriceSeries**](StatsAPI.md#StatsPriceSeries) | **Get** /stats/price/series/{timeframe} | Get histogram with TIA price
[**StatsRollup24h**](StatsAPI.md#StatsRollup24h) | **Get** /stats/rollup_stats_24h | Get rollups stats for last 24 hours
[**StatsSeries**](StatsAPI.md#StatsSeries) | **Get** /stats/series/{name}/{timeframe} | Get histogram with precomputed stats
[**StatsSeriesCumulative**](StatsAPI.md#StatsSeriesCumulative) | **Get** /stats/series/{name}/{timeframe}/cumulative | Get cumulative histogram with precomputed stats
[**StatsSquareSize**](StatsAPI.md#StatsSquareSize) | **Get** /stats/square_size | Get histogram for square size distribution
[**StatsStakingSeries**](StatsAPI.md#StatsStakingSeries) | **Get** /stats/staking/series/{id}/{name}/{timeframe} | Get histogram for staking with precomputed stats
[**StatsSummary**](StatsAPI.md#StatsSummary) | **Get** /stats/summary/{table}/{function} | Get value by table and function



## Stats24hChanges

> []ResponsesChange24hBlockStats Stats24hChanges(ctx).Execute()

Get changes for 24 hours



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.Stats24hChanges(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.Stats24hChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Stats24hChanges`: []ResponsesChange24hBlockStats
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.Stats24hChanges`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStats24hChangesRequest struct via the builder pattern


### Return type

[**[]ResponsesChange24hBlockStats**](ResponsesChange24hBlockStats.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsMessagesCount24h

> []ResponsesCountItem StatsMessagesCount24h(ctx).Execute()

Get messages distribution for the last 24 hours



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsMessagesCount24h(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsMessagesCount24h``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsMessagesCount24h`: []ResponsesCountItem
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsMessagesCount24h`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatsMessagesCount24hRequest struct via the builder pattern


### Return type

[**[]ResponsesCountItem**](ResponsesCountItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsNamespaceUsage

> []ResponsesNamespaceUsage StatsNamespaceUsage(ctx).Top(top).Execute()

Get namespaces with sorting by size.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {
	top := int32(56) // int32 | Count of entities (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsNamespaceUsage(context.Background()).Top(top).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsNamespaceUsage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsNamespaceUsage`: []ResponsesNamespaceUsage
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsNamespaceUsage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatsNamespaceUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **top** | **int32** | Count of entities | 

### Return type

[**[]ResponsesNamespaceUsage**](ResponsesNamespaceUsage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsNsSeries

> []ResponsesSeriesItem StatsNsSeries(ctx, id, timeframe, name).From(from).To(to).Execute()

Get histogram for namespace with precomputed stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {
	id := "id_example" // string | Namespace id in hexadecimal
	timeframe := "timeframe_example" // string | Timeframe
	name := "name_example" // string | Series name
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsNsSeries(context.Background(), id, timeframe, name).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsNsSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsNsSeries`: []ResponsesSeriesItem
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsNsSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Namespace id in hexadecimal | 
**timeframe** | **string** | Timeframe | 
**name** | **string** | Series name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatsNsSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

[**[]ResponsesSeriesItem**](ResponsesSeriesItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsPriceCurrent

> ResponsesPrice StatsPriceCurrent(ctx).Execute()

Get current TIA price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsPriceCurrent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsPriceCurrent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsPriceCurrent`: ResponsesPrice
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsPriceCurrent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatsPriceCurrentRequest struct via the builder pattern


### Return type

[**ResponsesPrice**](ResponsesPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsPriceSeries

> []ResponsesPrice StatsPriceSeries(ctx, timeframe).From(from).To(to).Execute()

Get histogram with TIA price



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {
	timeframe := "timeframe_example" // string | Timeframe
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsPriceSeries(context.Background(), timeframe).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsPriceSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsPriceSeries`: []ResponsesPrice
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsPriceSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeframe** | **string** | Timeframe | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatsPriceSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

[**[]ResponsesPrice**](ResponsesPrice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsRollup24h

> []ResponsesRollupStats24h StatsRollup24h(ctx).Execute()

Get rollups stats for last 24 hours



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsRollup24h(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsRollup24h``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsRollup24h`: []ResponsesRollupStats24h
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsRollup24h`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStatsRollup24hRequest struct via the builder pattern


### Return type

[**[]ResponsesRollupStats24h**](ResponsesRollupStats24h.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsSeries

> []ResponsesSeriesItem StatsSeries(ctx, timeframe, name).From(from).To(to).Execute()

Get histogram with precomputed stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {
	timeframe := "timeframe_example" // string | Timeframe
	name := "name_example" // string | Series name
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsSeries(context.Background(), timeframe, name).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsSeries`: []ResponsesSeriesItem
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeframe** | **string** | Timeframe | 
**name** | **string** | Series name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatsSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

[**[]ResponsesSeriesItem**](ResponsesSeriesItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsSeriesCumulative

> []ResponsesSeriesItem StatsSeriesCumulative(ctx, timeframe, name).From(from).To(to).Execute()

Get cumulative histogram with precomputed stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {
	timeframe := "timeframe_example" // string | Timeframe
	name := "name_example" // string | Series name
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsSeriesCumulative(context.Background(), timeframe, name).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsSeriesCumulative``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsSeriesCumulative`: []ResponsesSeriesItem
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsSeriesCumulative`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**timeframe** | **string** | Timeframe | 
**name** | **string** | Series name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatsSeriesCumulativeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

[**[]ResponsesSeriesItem**](ResponsesSeriesItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsSquareSize

> []map[string][]ResponsesTimeValueItem StatsSquareSize(ctx).From(from).To(to).Execute()

Get histogram for square size distribution



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsSquareSize(context.Background()).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsSquareSize``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsSquareSize`: []map[string][]ResponsesTimeValueItem
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsSquareSize`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStatsSquareSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

[**[]map[string][]ResponsesTimeValueItem**](map.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsStakingSeries

> []ResponsesSeriesItem StatsStakingSeries(ctx, id, timeframe, name).From(from).To(to).Execute()

Get histogram for staking with precomputed stats



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {
	id := "id_example" // string | Validator id
	timeframe := "timeframe_example" // string | Timeframe
	name := "name_example" // string | Series name
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsStakingSeries(context.Background(), id, timeframe, name).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsStakingSeries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsStakingSeries`: []ResponsesSeriesItem
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsStakingSeries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Validator id | 
**timeframe** | **string** | Timeframe | 
**name** | **string** | Series name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatsStakingSeriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

[**[]ResponsesSeriesItem**](ResponsesSeriesItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StatsSummary

> string StatsSummary(ctx, table, function).Column(column).From(from).To(to).Execute()

Get value by table and function



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func main() {
	table := "table_example" // string | Table name
	function := "function_example" // string | Function name
	column := "column_example" // string | Column name which will be used for computation. Optional for count. (optional)
	from := int32(56) // int32 | Time from in unix timestamp (optional)
	to := int32(56) // int32 | Time to in unix timestamp (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatsAPI.StatsSummary(context.Background(), table, function).Column(column).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatsAPI.StatsSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StatsSummary`: string
	fmt.Fprintf(os.Stdout, "Response from `StatsAPI.StatsSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**table** | **string** | Table name | 
**function** | **string** | Function name | 

### Other Parameters

Other parameters are passed through a pointer to a apiStatsSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **column** | **string** | Column name which will be used for computation. Optional for count. | 
 **from** | **int32** | Time from in unix timestamp | 
 **to** | **int32** | Time to in unix timestamp | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

