package ebird

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTop100(t *testing.T) {
	input := `[
		{"profileHandle": "MTE3NDA2NA", "userDisplayName": "Rajan Rao", "numSpecies": 132, "numCompleteChecklists": 0, "rowNum": 1, "userId": "USER1174064"},
		{"profileHandle": "MTIzMjAy", "userDisplayName": "Aaron Haiman", "numSpecies": 132, "numCompleteChecklists": 0, "rowNum": 1, "userId": "USER123202"}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/product/top100/CA-AB/2023/10/6", r.URL.Path)
		assert.Equal(t, "spp", r.URL.Query().Get("rankedBy"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	date := time.Date(2023, 10, 6, 0, 0, 0, 0, time.UTC)
	got, err := client.Top100(ctx, "CA-AB", date, RankedBy("spp"))
	require.NoError(t, err)

	var want []Top100
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestChecklistFeedOnDate(t *testing.T) {
	input := `[
		{
			"locId": "L4122129",
			"subId": "S151140860",
			"userDisplayName": "Mark Hays",
			"numSpecies": 1,
			"obsDt": "30 Sep 2023",
			"obsTime": "23:40",
			"isoObsDate": "2023-09-30 23:40",
			"subID": "S151140860",
			"loc": {
				"locId": "L4122129",
				"name": "Hays Home",
				"latitude": 35.1337729,
				"longitude": -120.5566585,
				"countryCode": "US",
				"countryName": "United States",
				"subnational1Name": "California",
				"subnational1Code": "US-CA",
				"subnational2Code": "US-CA-079",
				"subnational2Name": "San Luis Obispo",
				"isHotspot": false,
				"locName": "Hays Home",
				"lat": 35.1337729,
				"lng": -120.5566585,
				"hierarchicalName": "Hays Home, San Luis Obispo, California, US",
				"locID": "L4122129"
			}
		}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/product/lists/US-CA/2023/9/30", r.URL.Path)
		assert.Equal(t, "10", r.URL.Query().Get("maxResults"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	date := time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC)
	got, err := client.ChecklistFeedOnDate(ctx, "US-CA", date, MaxResults(10))
	require.NoError(t, err)

	var want []ChecklistFeedOnDate
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestRegionalStatisticsOnDate(t *testing.T) {
	input := `{"numChecklists": 100, "numContributors": 50, "numSpecies": 200}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/product/stats/US-CA/2023/9/30", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	date := time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC)
	got, err := client.RegionalStatisticsOnDate(ctx, "US-CA", date)
	require.NoError(t, err)

	want := &RegionalStatisticsOnDate{
		NumChecklists:   100,
		NumContributors: 50,
		NumSpecies:      200,
	}

	assert.Equal(t, want, got)
}

func TestSpeciesListForRegion(t *testing.T) {
	input := `["specode1", "specode2", "specode3"]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/product/spplist/US-CA", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.SpeciesListForRegion(ctx, "US-CA")
	require.NoError(t, err)

	want := []string{"specode1", "specode2", "specode3"}
	assert.Equal(t, want, got)
}

func TestViewChecklist(t *testing.T) {
	input := `{
		"subId": "S12345678",
		"protocolId": "P22",
		"locId": "L123456",
		"durationHrs": 2.5,
		"allObsReported": true,
		"creationDt": "2023-10-07T10:00:00Z",
		"obsDt": "2023-10-07",
		"numObservers": 2,
		"numSpecies": 15,
		"obs": [
			{
				"speciesCode": "amecro",
				"howManyStr": "5",
				"present": true
			},
			{
				"speciesCode": "norcad",
				"howManyStr": "2",
				"present": true
			}
		]
	}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/product/checklist/view/S12345678", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.ViewChecklist(ctx, "S12345678")
	require.NoError(t, err)

	var want ViewChecklist
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, &want, got)
}

func TestTop100WithError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid region code"}`))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	date := time.Date(2023, 10, 6, 0, 0, 0, 0, time.UTC)
	_, err = client.Top100(ctx, "INVALID", date)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get Top 100")
}

func TestChecklistFeedOnDateWithEmptyRegionCode(t *testing.T) {
	client, err := NewClient("test-api-key")
	require.NoError(t, err)

	ctx := context.Background()
	date := time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC)
	_, err = client.ChecklistFeedOnDate(ctx, "", date)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "regionCode cannot be empty")
}

func TestRegionalStatisticsOnDateWithInvalidDate(t *testing.T) {
	client, err := NewClient("test-api-key")
	require.NoError(t, err)

	ctx := context.Background()
	invalidDate := time.Time{}
	_, err = client.RegionalStatisticsOnDate(ctx, "US-CA", invalidDate)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to get regional statistics")
}
