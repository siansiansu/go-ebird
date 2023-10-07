package ebird

import (
	"bytes"
	"context"
	"net/http"
	"testing"
)

func TestRegionInfo(t *testing.T) {
	client, server := testClient(http.StatusOK, bytes.NewBufferString(`{"bounds": {"minX": -125.0, "maxX": -66.934570, "minY": 24.396308, "maxY": 49.384358}, "result": "Success", "code": "US", "type": "country", "longitude": -95.712891, "latitude": 37.09024}`))
	defer server.Close()

	client.httpClient = server.Client()

	regionCode := "US"

	ctx := context.Background()

	result, err := client.RegionInfo(ctx, regionCode)
	if err != nil {
		t.Fatalf("Error making API request: %v", err)
	}

	if result.Code != "US" || result.Result != "Success" || result.Type != "country" || result.Longitude != -95.712891 || result.Latitude != 37.09024 {
		t.Errorf("Unexpected values for RegionInfo. Got: %v", result)
	}
}

func TestSubRegionList(t *testing.T) {
	client, server := testClient(http.StatusOK, bytes.NewBufferString(`[{"code": "US-TX", "name": "Texas"}, {"code": "US-CA", "name": "California"}]`))
	defer server.Close()

	client.httpClient = server.Client()

	regionType := "subnational1"
	parentRegionCode := "US"

	ctx := context.Background()

	result, err := client.SubRegionList(ctx, regionType, parentRegionCode)
	if err != nil {
		t.Fatalf("Error making API request: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Unexpected number of subregions. Expected: 2, Got: %d", len(result))
	}

	if result[0].Code != "US-TX" || result[0].Name != "Texas" {
		t.Errorf("Unexpected values for the first subregion. Got: %v", result[0])
	}

	if result[1].Code != "US-CA" || result[1].Name != "California" {
		t.Errorf("Unexpected values for the second subregion. Got: %v", result[1])
	}
}
