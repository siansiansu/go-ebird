package ebird

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRegionInfo(t *testing.T) {
	input := `{
		"bounds": {"minX": -125.0, "maxX": -66.934570, "minY": 24.396308, "maxY": 49.384358},
		"result": "Success",
		"code": "US",
		"type": "country",
		"longitude": -95.712891,
		"latitude": 37.09024
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ref/region/info/US", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.RegionInfo(ctx, "US")
	require.NoError(t, err)

	var want RegionInfo
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, &want, got)
}

func TestSubRegionList(t *testing.T) {
	input := `[
		{"code": "US-TX", "name": "Texas"},
		{"code": "US-CA", "name": "California"}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ref/region/list/subnational1/US", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.SubRegionList(ctx, "subnational1", "US")
	require.NoError(t, err)

	var want []SubRegion
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestRegionInfoWithEmptyRegionCode(t *testing.T) {
	client, err := NewClient("test-api-key")
	require.NoError(t, err)

	ctx := context.Background()
	_, err = client.RegionInfo(ctx, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "regionCode cannot be empty")
}

func TestSubRegionListWithEmptyParameters(t *testing.T) {
	client, err := NewClient("test-api-key")
	require.NoError(t, err)

	ctx := context.Background()

	_, err = client.SubRegionList(ctx, "", "US")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "regionType cannot be empty")

	_, err = client.SubRegionList(ctx, "subnational1", "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "parentRegionCode cannot be empty")
}
