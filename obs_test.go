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

func TestObservations(t *testing.T) {
	tests := []struct {
		name        string
		method      string
		endpoint    string
		regionCode  string
		speciesCode string
		input       string
		wantResult  interface{}
		wantErr     bool
		params      []RequestOption
	}{
		{
			name:       "RecentObservationsInRegion",
			method:     "RecentObservationsInRegion",
			endpoint:   "/data/obs/CA-AB/recent",
			regionCode: "CA-AB",
			input:      `[{"speciesCode":"rufgro","comName":"Ruffed Grouse","sciName":"Bonasa umbellus","locId":"L366849","locName":"Kananaskis--Gorge Creek Trail (North Access)","obsDt":"2023-10-05 21:42","howMany":1,"lat":50.7541732,"lng":-114.5556483,"obsValid":true,"obsReviewed":false,"locationPrivate":false,"subId":"S151517643"}]`,
			wantResult: []Observation{{SpeciesCode: "rufgro", ComName: "Ruffed Grouse", SciName: "Bonasa umbellus", LocId: "L366849", LocName: "Kananaskis--Gorge Creek Trail (North Access)", ObsDt: "2023-10-05 21:42", HowMany: 1, Lat: 50.7541732, Lng: -114.5556483, ObsValid: true, ObsReviewed: false, LocationPrivate: false, SubId: "S151517643"}},
		},
		{
			name:       "RecentNotableObservationsInRegion",
			method:     "RecentNotableObservationsInRegion",
			endpoint:   "/data/obs/CA-AB/recent/notable",
			regionCode: "CA-AB",
			input:      `[{"speciesCode":"rebwoo","comName":"Red-bellied Woodpecker","sciName":"Melanerpes carolinus","locId":"L3089074","locName":"Nacmine (community)","obsDt":"2023-10-05 17:46","howMany":1,"lat":51.4708064,"lng":-112.7871633,"obsValid":true,"obsReviewed":true,"locationPrivate":false,"subId":"S151515981"}]`,
			wantResult: []Observation{{SpeciesCode: "rebwoo", ComName: "Red-bellied Woodpecker", SciName: "Melanerpes carolinus", LocId: "L3089074", LocName: "Nacmine (community)", ObsDt: "2023-10-05 17:46", HowMany: 1, Lat: 51.4708064, Lng: -112.7871633, ObsValid: true, ObsReviewed: true, LocationPrivate: false, SubId: "S151515981"}},
		},
		{
			name:        "RecentObservationsOfSpeciesInRegion",
			method:      "RecentObservationsOfSpeciesInRegion",
			endpoint:    "/data/obs/US/recent/eutspa",
			regionCode:  "US",
			speciesCode: "eutspa",
			input:       `[{"speciesCode":"eutspa","comName":"Eurasian Tree Sparrow","sciName":"Passer montanus","locId":"L17449098","locName":"Home","obsDt":"2023-10-05 18:34","howMany":2,"lat":38.572064,"lng":-90.420457,"obsValid":true,"obsReviewed":false,"locationPrivate":true,"subId":"S151508867","exoticCategory":"N"}]`,
			wantResult:  []Observation{{SpeciesCode: "eutspa", ComName: "Eurasian Tree Sparrow", SciName: "Passer montanus", LocId: "L17449098", LocName: "Home", ObsDt: "2023-10-05 18:34", HowMany: 2, Lat: 38.572064, Lng: -90.420457, ObsValid: true, ObsReviewed: false, LocationPrivate: true, SubId: "S151508867", ExoticCategory: "N"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, tt.endpoint, r.URL.Path)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.input))
			}))
			defer server.Close()

			client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
			require.NoError(t, err)

			ctx := context.Background()
			var got interface{}

			switch tt.method {
			case "RecentObservationsInRegion":
				got, err = client.RecentObservationsInRegion(ctx, tt.regionCode, tt.params...)
			case "RecentNotableObservationsInRegion":
				got, err = client.RecentNotableObservationsInRegion(ctx, tt.regionCode, tt.params...)
			case "RecentObservationsOfSpeciesInRegion":
				got, err = client.RecentObservationsOfSpeciesInRegion(ctx, tt.regionCode, tt.speciesCode, tt.params...)
			}

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestRecentChecklistsFeed(t *testing.T) {
	input := `[{"locId":"L13557069","subId":"S151523092","userDisplayName":"Ed Bailey","numSpecies":1,"obsDt":"6 Oct 2023","obsTime":"03:35","isoObsDate":"2023-10-06 03:35","subID":"S151523092","loc":{"locId":"L13557069","name":"my yard","latitude":41.3242778,"longitude":-73.1988271,"countryCode":"US","countryName":"United States","subnational1Name":"Connecticut","subnational1Code":"US-CT","subnational2Code":"US-CT-001","subnational2Name":"Fairfield","isHotspot":false,"locID":"L13557069","locName":"my yard","lat":41.3242778,"lng":-73.1988271,"hierarchicalName":"my yard, Fairfield, Connecticut, US"}}]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/product/lists/US", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.RecentChecklistsFeed(ctx, "US", MaxResults(10))
	require.NoError(t, err)

	var want []RecentChecklistFeed
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestRecentNearbyObservations(t *testing.T) {
	input := `[{"speciesCode":"eutspa","comName":"Eurasian Tree Sparrow","sciName":"Passer montanus","locId":"L17449098","locName":"Home","obsDt":"2023-10-05 18:34","howMany":2,"lat":38.572064,"lng":-90.420457,"obsValid":true,"obsReviewed":false,"locationPrivate":true,"subId":"S151508867","exoticCategory":"N"}]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/data/obs/geo/recent", r.URL.Path)
		assert.Equal(t, "38.50", r.URL.Query().Get("lat"))
		assert.Equal(t, "-90.40", r.URL.Query().Get("lng"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.RecentNearbyObservations(ctx, Lat(38.5), Lng(-90.4))
	require.NoError(t, err)

	var want []Observation
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestNearestObservationsOfSpecies(t *testing.T) {
	input := `[{"speciesCode":"eutspa","comName":"Eurasian Tree Sparrow","sciName":"Passer montanus","locId":"L21181579","locName":"名古屋市--勅使池  (Nagoya--Chokushi-ike Park)","obsDt":"2023-09-24 08:00","howMany":2,"lat":35.083182,"lng":137.013097,"obsValid":true,"obsReviewed":false,"locationPrivate":false,"subId":"S150855601"}]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/data/nearest/geo/recent/eutspa", r.URL.Path)
		assert.Equal(t, "35.00", r.URL.Query().Get("lat"))
		assert.Equal(t, "137.00", r.URL.Query().Get("lng"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.NearestObservationsOfSpecies(ctx, "eutspa", Lat(35.0), Lng(137.0))
	require.NoError(t, err)

	var want []Observation
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestHistoricObservationsOnDate(t *testing.T) {
	input := `[{"speciesCode":"grhowl","comName":"Great Horned Owl","sciName":"Bubo virginianus","locId":"L4122129","locName":"Hays Home","obsDt":"2023-09-30 23:40","howMany":2,"lat":35.1337729,"lng":-120.5566585,"obsValid":true,"obsReviewed":false,"locationPrivate":true,"subId":"S151140860"}]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/data/obs/US-CA/historic/2023/9/30", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	date := time.Date(2023, 9, 30, 0, 0, 0, 0, time.UTC)
	got, err := client.HistoricObservationsOnDate(ctx, "US-CA", date)
	require.NoError(t, err)

	var want []Observation
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestRecentNearbyNotableObservations(t *testing.T) {
	input := `[{"speciesCode":"masboo","comName":"Masked Booby","sciName":"Sula dactylatra","locId":"L27325351","locName":"岐阜市--上佐波西９丁目の水田 (Gifu--Kamisabanishi 9-Chome ricefields)","obsDt":"2023-10-01 15:50","howMany":1,"lat":35.3662473,"lng":136.7154931,"obsValid":false,"obsReviewed":false,"locationPrivate":false,"subId":"S151144322"}]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/data/obs/geo/recent/notable", r.URL.Path)
		assert.Equal(t, "35.30", r.URL.Query().Get("lat"))
		assert.Equal(t, "136.70", r.URL.Query().Get("lng"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.RecentNearbyNotableObservations(ctx, Lat(35.3), Lng(136.7))
	require.NoError(t, err)

	var want []Observation
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}
