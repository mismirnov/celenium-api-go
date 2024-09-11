/*
Celenium API

Testing BlockAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package celenium

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/mismirnov/celenium-api-go"
)

func Test_celenium_BlockAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test BlockAPIService BlockBlobsCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var height int32

		resp, httpRes, err := apiClient.BlockAPI.BlockBlobsCount(context.Background(), height).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAPIService GetBlock", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var height int32

		resp, httpRes, err := apiClient.BlockAPI.GetBlock(context.Background(), height).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAPIService GetBlockBlobs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var height int32

		resp, httpRes, err := apiClient.BlockAPI.GetBlockBlobs(context.Background(), height).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAPIService GetBlockCount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockAPI.GetBlockCount(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAPIService GetBlockEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var height int32

		resp, httpRes, err := apiClient.BlockAPI.GetBlockEvents(context.Background(), height).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAPIService GetBlockMessages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var height int32

		resp, httpRes, err := apiClient.BlockAPI.GetBlockMessages(context.Background(), height).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAPIService GetBlockStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var height int32

		resp, httpRes, err := apiClient.BlockAPI.GetBlockStats(context.Background(), height).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test BlockAPIService ListBlock", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.BlockAPI.ListBlock(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
