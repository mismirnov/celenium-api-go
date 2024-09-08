/*
Celenium API

Testing SearchAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package celenium

import (
	"context"
	celeniumApi "github.com/mismirnov/celenium-api-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_celenium_SearchAPIService(t *testing.T) {

	configuration := celeniumApi.NewConfiguration()
	apiClient := celeniumApi.NewAPIClient(configuration)

	t.Run("Test SearchAPIService Search", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.SearchAPI.Search(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
