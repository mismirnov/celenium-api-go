/*
Celenium API

Celenium API is a powerful tool to access all blockchain data that is processed and indexed by our proprietary indexer. With Celenium API you can retrieve all historical data, off-chain data, blobs and statistics through our REST API. Celenium API indexer are open source, which allows you to not depend on third-party services. You can clone, build and run them independently, giving you full control over all components. If you have any questions or feature requests, please feel free to contact us. We appreciate your feedback!

API version: 1.0
Contact: celenium@pklabs.me
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package celenium

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// SearchAPIService SearchAPI service
type SearchAPIService service

type ApiSearchRequest struct {
	ctx context.Context
	ApiService *SearchAPIService
	query *string
}

// Search string
func (r ApiSearchRequest) Query(query string) ApiSearchRequest {
	r.query = &query
	return r
}

func (r ApiSearchRequest) Execute() ([]ResponsesSearchItem, *http.Response, error) {
	return r.ApiService.SearchExecute(r)
}

/*
Search Search by hash

Endpoint finds entity by hash (block, address, validator, namespace and tx). It returns array of `responses.SearchItem` entities.

### Block

Block will be found by its hash. Hash example: `652452A670018D629CC116E510BA88C1CABE061336661B1F3D206D248BD558AF`.
Hash should be hexadecimal and has a length of 64.

#### Example response 

```json
{
    "type": "block",
    "result": {
        "id": 1,
        "hash": "652452A670018D629CC116E510BA88C1CABE061336661B1F3D206D248BD558AF",
        // ... rest fields from response.Block type
    }
}
```

### Tx

Tx will be found by its hash. Hash example: `652452A670018D629CC116E510BA88C1CABE061336661B1F3D206D248BD558AF`.
Tx should be hexadecimal and has a length of 64.

#### Example response 

```json
{
    "type": "tx",
    "result": {
        "id": 1,
        "hash": "652452A670018D629CC116E510BA88C1CABE061336661B1F3D206D248BD558AF",
        // ... rest fields from response.Tx type
    }
}
```

### Address

The Address will be found by its hash.
Hash example: `celestia1jc92qdnty48pafummfr8ava2tjtuhfdw774w60`.
Address has prefix `celestia` and has length 47.
Also, it should be decoded by `bech32`.

#### Example response 

```json
{
    "type": "address",
    "result": {
        "id": 1,
        "hash": "celestia1jc92qdnty48pafummfr8ava2tjtuhfdw774w60",
        "height": 100,
        "balance": "6525472354"
    }
}
```

### Namespace

Namespace can be found by base64 hash and identity pair version + namespace id. 
Hash example: `U3dhZ2dlciByb2Nrcw==`. 
Identity pair example: `014723ce10b187716adfc55ff7e6d9179c226e6b5440b02577cca49d02`

#### Example response 

```json
{
    "type": "namespace",
    "result": {
        "id": 1,
        "hash": "U3dhZ2dlciByb2Nrcw==",
        "version": 1,
        "namespace_id": "4723ce10b187716adfc55ff7e6d9179c226e6b5440b02577cca49d02"
        // ... rest fields from response.Namespace type
    }
}
```

### Validator

Validator can be found by moniker prefix. 
For example: names `Node 1` and `Node 2` can be found with query string `Node`

#### Example response 

```json
{
    "type": "validator",
    "result": {
        "id": 1,
        "moniker": "Node 1",
        // ... rest fields from response.Validator type
    }
}
```

### Rollup

Rollup can be found by name prefix. 
For example: rollup with names `Rollup 1` and `Rollup 2` can be found with query string `Rol`

#### Example response 

```json
{
    "type": "rollup",
    "result": {
        "id": 1,
        "moniker": "Rollup 1",
        // ... rest fields from response.Rollup type
    }
}
```

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiSearchRequest
*/
func (a *SearchAPIService) Search(ctx context.Context) ApiSearchRequest {
	return ApiSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ResponsesSearchItem
func (a *SearchAPIService) SearchExecute(r ApiSearchRequest) ([]ResponsesSearchItem, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ResponsesSearchItem
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SearchAPIService.Search")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.query == nil {
		return localVarReturnValue, nil, reportError("query is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "query", r.query, "", "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v HandlerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v HandlerError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
