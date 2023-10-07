package ebird

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type GeoTestSuite struct {
	suite.Suite
}

func TestGeoSuites(t *testing.T) {
	suite.Run(t, new(GeoTestSuite))
}

func (ts *GeoTestSuite) TestAdjacentRegions() {
	tests := []struct {
		Desc      string
		Input     string
		ExpRes    []AdjacentRegions
		ExpErr    error
		ExpParams RequestOption
	}{
		{
			Desc: "ensure that the request of AdjacentRegions is correct without query params",
			Input: `
[
  {
    "code": "US",
    "name": "United States"
  },
  {
    "code": "CA",
    "name": "Canada"
  }
]`,
			ExpRes: []AdjacentRegions{
				{
					Code: "US",
					Name: "United States",
				},
				{
					Code: "CA",
					Name: "Canada",
				},
			},
			ExpErr:    nil,
			ExpParams: nil,
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		regionCode := "US"

		ctx := context.Background()

		result, err := client.AdjacentRegions(ctx, regionCode)
		if err != nil {

		} else {
			ts.Equal(test.ExpRes, result, test.Desc)
		}
	}
}
