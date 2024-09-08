/*
Celenium API

Testing VestingAPIService

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

func Test_celenium_VestingAPIService(t *testing.T) {

	configuration := celeniumApi.NewConfiguration()
	apiClient := celeniumApi.NewAPIClient(configuration)

	t.Run("Test VestingAPIService GetVestingPeriods", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var id int32

		resp, httpRes, err := apiClient.VestingAPI.GetVestingPeriods(context.Background(), id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
